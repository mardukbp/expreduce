package cas

import "bytes"

func (this *Expression) ToStringRule() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(this.Parts[1].String())
	buffer.WriteString(") -> (")
	buffer.WriteString(this.Parts[2].String())
	buffer.WriteString(")")
	return buffer.String()
}

func ReplacePD(this Ex, es *EvalState) Ex {
	es.Infof("In ReplacePD(%v, es.patternDefined=%v)", this, es.patternDefined)
	toReturn := this.DeepCopy()
	for nameStr, def := range es.patternDefined {
		toReturn = Replace(toReturn,
			&Expression{[]Ex{
				&Symbol{"Rule"},
				&Symbol{nameStr},
				def,
			}}, es)
	}
	es.Infof("Finished ReplacePD with toReturn=%v", toReturn)
	return toReturn
}

func Replace(this Ex, r *Expression, es *EvalState) Ex {
	_, isFlt := this.(*Flt)
	_, isInteger := this.(*Integer)
	_, isString := this.(*String)
	asExpression, isExpression := this.(*Expression)
	_, isSymbol := this.(*Symbol)

	if isFlt || isInteger || isString || isSymbol {
		if IsMatchQ(this, r.Parts[1], es) {
			return r.Parts[2]
		}
		return this
	} else if isExpression {
		return asExpression.Replace(r, es)
	}
	return &Symbol{"ReplaceFailed"}
}

func (this *Expression) EvalReplace(es *EvalState) Ex {
	if len(this.Parts) != 3 {
		return this
	}
	this.Parts[1] = this.Parts[1].Eval(es)
	this.Parts[2] = this.Parts[2].Eval(es)
	//_, ok := this.Parts[2].(*Rule)
	rulesRule, ok := HeadAssertion(this.Parts[2], "Rule")
	if ok {
		//oldVars := es.GetDefinedSnapshot()
		newEx := Replace(this.Parts[1], rulesRule, es)
		//newEx = ReplacePD(newEx, es)
		es.ClearPD()
		newEx = newEx.Eval(es)
		//es.defined = oldVars
		return newEx
		//return this
	}
	return this
}

func (this *Expression) ToStringReplace() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(this.Parts[1].String())
	buffer.WriteString(") /. (")
	buffer.WriteString(this.Parts[2].String())
	buffer.WriteString(")")
	return buffer.String()
}

func (this *Expression) EvalReplaceRepeated(es *EvalState) Ex {
	if len(this.Parts) != 3 {
		return this
	}
	es.Infof("Starting ReplaceRepeated.")
	this.Parts[1] = this.Parts[1].Eval(es)
	this.Parts[2] = this.Parts[2].Eval(es)
	//_, ok := this.Parts[2].(*Rule)
	rulesRule, ok := HeadAssertion(this.Parts[2], "Rule")
	if ok {
		isSame := false
		oldEx := this.Parts[1]
		es.Infof("In ReplaceRepeated. Initial expr: %v", oldEx)
		for !isSame {
			//oldVars := es.GetDefinedSnapshot()
			newEx := Replace(oldEx.DeepCopy(), rulesRule, es)
			es.ClearPD()
			newEx = newEx.Eval(es)
			//es.defined = oldVars
			es.Infof("In ReplaceRepeated. New expr: %v", newEx)

			//oldVars := es.GetDefinedSnapshot()
			if IsSameQ(oldEx, newEx, es) {
				isSame = true
			}
			es.ClearPD()
			//es.defined = oldVars
			oldEx = newEx
		}
		return oldEx
		//return this
	}
	return this
}

func (this *Expression) ToStringReplaceRepeated() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(this.Parts[1].String())
	buffer.WriteString(") //. (")
	buffer.WriteString(this.Parts[2].String())
	buffer.WriteString(")")
	return buffer.String()
}
