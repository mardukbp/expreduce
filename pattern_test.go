package cas

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {

	fmt.Println("Testing pattern")

	es := NewEvalState()

	// Matching patterns
	CasAssertSame(t, es, "True", "MatchQ[1, _Integer]")
	CasAssertSame(t, es, "False", "MatchQ[s, _Integer]")
	CasAssertSame(t, es, "True", "MatchQ[s, _Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[1, _Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[_Symbol, _Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[_Integer, _Integer]")
	CasAssertSame(t, es, "True", "MatchQ[_Symbol, _Blank]")
	CasAssertSame(t, es, "True", "MatchQ[_Symbol, test_Blank]")
	CasAssertSame(t, es, "True", "MatchQ[_Symbol, test_Blank]")
	CasAssertSame(t, es, "False", "MatchQ[_Symbol, test_Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[name_Symbol, test_Blank]")
	CasAssertSame(t, es, "True", "MatchQ[name_Symbol, test_Pattern]")
	CasAssertSame(t, es, "True", "MatchQ[_Symbol, test_Blank]")
	CasAssertSame(t, es, "False", "MatchQ[_Symbol, test_Pattern]")
	CasAssertSame(t, es, "False", "MatchQ[1.5, _Integer]")
	CasAssertSame(t, es, "True", "MatchQ[1.5, _Real]")
	CasAssertSame(t, es, "True", "_Symbol == _Symbol")
	CasAssertSame(t, es, "_Symbol == _Integer", "_Symbol == _Integer")
	CasAssertSame(t, es, "False", "MatchQ[_Symbol, s]")
	CasAssertSame(t, es, "False", "MatchQ[_Integer, 1]")
	CasAssertSame(t, es, "_Integer == 1", "_Integer == 1")
	CasAssertSame(t, es, "1 == _Integer", "1 == _Integer")

	CasAssertSame(t, es, "False", "_Integer === 1")
	CasAssertSame(t, es, "False", "1 === _Integer")
	CasAssertSame(t, es, "True", "_Integer === _Integer")
	CasAssertSame(t, es, "False", "_Symbol === a")
	CasAssertSame(t, es, "False", "a === _Symbol")
	CasAssertSame(t, es, "True", "_Symbol === _Symbol")

	CasAssertSame(t, es, "a == b", "a == b")
	CasAssertSame(t, es, "2", "a == b /. _Equal -> 2")
	CasAssertSame(t, es, "If[1 == k, itstrue, itsfalse]", "If[1 == k, itstrue, itsfalse]")
	CasAssertSame(t, es, "99", "If[1 == k, itstrue, itsfalse] /. _If -> 99")
	CasAssertSame(t, es, "False", "MatchQ[kfdsfdsf[], _Function]")
	CasAssertSame(t, es, "True", "MatchQ[kfdsfdsf[], _kfdsfdsf]")
	CasAssertSame(t, es, "99", "kfdsfdsf[] /. _kfdsfdsf -> 99")
	CasAssertSame(t, es, "a + b", "a + b")
	CasAssertSame(t, es, "2", "a + b /. _Plus -> 2")
	CasAssertSame(t, es, "2", "a*b /. _Times -> 2")
	CasAssertSame(t, es, "2", "a^b /. _Power -> 2")
	CasAssertSame(t, es, "2", "a -> b /. _Rule -> 2")
	CasAssertSame(t, es, "2", "a*b*c*d /. _Times -> 2")

	es.ClearAll()
	CasAssertSame(t, es, "True", "MatchQ[x*3., c1match_Real*matcha_]")
	CasAssertSame(t, es, "True", "MatchQ[3.*x, c1match_Real*matcha_]")
	CasAssertSame(t, es, "True", "MatchQ[x+3., c1match_Real+matcha_]")
	CasAssertSame(t, es, "True", "MatchQ[3.+x, c1match_Real+matcha_]")
	CasAssertSame(t, es, "True", "MatchQ[y + x, matcha_]")
	CasAssertSame(t, es, "True", "MatchQ[y*x, matcha_]")

	// Test BlankSequence
	// Be wary of the false matches - the default is usually false.
	CasAssertSame(t, es, "True", "MatchQ[a, __]")
	CasAssertSame(t, es, "True", "MatchQ[a + b, __]")
	CasAssertSame(t, es, "True", "MatchQ[a*b, __]")
	CasAssertSame(t, es, "True", "MatchQ[a*b, ___]")
	CasAssertSame(t, es, "False", "MatchQ[a, a*__]")
	CasAssertSame(t, es, "False", "MatchQ[a, a*___]")
	CasAssertSame(t, es, "False", "MatchQ[a, a + ___]")
	CasAssertSame(t, es, "True", "MatchQ[a + b, a + b + ___]")
	CasAssertSame(t, es, "True", "MatchQ[a + b + c, a + b + __]")
	CasAssertSame(t, es, "True", "MatchQ[a + b + c + d, a + b + __]")
	CasAssertSame(t, es, "False", "MatchQ[a*b, ___Integer]")
	CasAssertSame(t, es, "False", "MatchQ[a*b, ___Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[a*b, __Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[a*b, x__Symbol]")
	CasAssertSame(t, es, "True", "MatchQ[a, __Symbol]")
	CasAssertSame(t, es, "True", "MatchQ[a, ___Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[a + b, ___Symbol]")
	CasAssertSame(t, es, "True", "MatchQ[a*b, x__Times]")
	CasAssertSame(t, es, "False", "MatchQ[a*b, x__Plus]")
	CasAssertSame(t, es, "True", "MatchQ[a + b, x__Plus]")
	CasAssertSame(t, es, "True", "MatchQ[a + b + c, a + x__Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[a + b + c, a + x__Plus]")
	CasAssertSame(t, es, "True", "MatchQ[a + b + c, a + x___Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[a + b + c, a + x___Plus]")
	CasAssertSame(t, es, "True", "MatchQ[a + b, a + x__Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[a + b, a + x__Plus]")
	CasAssertSame(t, es, "False", "MatchQ[a + b, a + b + x__Symbol]")
	CasAssertSame(t, es, "False", "MatchQ[a + b, a + b + x__Plus]")
	CasAssertSame(t, es, "True", "MatchQ[a + b, a + b + x___Symbol]")
	CasAssertSame(t, es, "True", "MatchQ[a + b, a + b + x___Plus]")
	CasAssertSame(t, es, "True", "MatchQ[4*a*b*c*d*e*f, __]")
	CasAssertSame(t, es, "True", "MatchQ[4*a*b*c*d*e*f, 4*__]")
	CasAssertSame(t, es, "False", "MatchQ[4*a*b*c*4*d*e*f, 4*__]")
	CasAssertSame(t, es, "False", "MatchQ[4*a*b*c*4*d*e*f, 4*__]")
	CasAssertSame(t, es, "True", "MatchQ[a*b*c*4*d*e*f, 4*__]")
	CasAssertSame(t, es, "True", "MatchQ[a*b*c*4*d*e*f, 4*__]")
	CasAssertSame(t, es, "False", "MatchQ[a*b*c*4*d*e*f, 5*__]")
	CasAssertSame(t, es, "False", "MatchQ[a*b*c*4*d*e*f + 5, 4*__]")
	CasAssertSame(t, es, "False", "MatchQ[a*b*c*4*d*e*f + 5*k, 4*__]")
	CasAssertSame(t, es, "False", "MatchQ[a*b*c*4*d*e*f + 5*k, 4*__]")
	CasAssertSame(t, es, "True", "MatchQ[a*b*c*4*d*e*f + 5*k, 4*__ + 5*k]")
	CasAssertSame(t, es, "False", "MatchQ[a*b*c*4*d*e*f + 2 + 5*k, 4*__ + 5*k]")
	CasAssertSame(t, es, "True", "MatchQ[(a*b*c)^e, __^e]")
	CasAssertSame(t, es, "False", "MatchQ[(a*b*c)^e, __^f + __^e]")
	CasAssertSame(t, es, "True", "MatchQ[(a*b*c)^e + (a*b*c)^f, __^f + __^e]")
	CasAssertSame(t, es, "True", "MatchQ[(a*b*c)^e + (a + b + c)^f, __^f + __^e]")
	CasAssertSame(t, es, "False", "MatchQ[(a*b*c)^e + (a + b + c)^f, amatch__^f + amatch__^e]")
	CasAssertSame(t, es, "True", "MatchQ[(a*b*c)^e + (a*b*c)^f, amatch__^f + amatch__^e]")

	// Warm up for combining like terms
	CasAssertSame(t, es, "True", "MatchQ[bar[1, foo[a, b]], bar[amatch_Integer, foo[cmatch__]]]")
	CasAssertSame(t, es, "True", "MatchQ[bar[1, foo[a, b, c]], bar[amatch_Integer, foo[cmatch__]]]")
	CasAssertSame(t, es, "False", "MatchQ[bar[1, foo[]], bar[amatch_Integer, foo[cmatch__]]]")
	CasAssertSame(t, es, "2", "bar[1, foo[a, b]] /. bar[amatch_Integer, foo[cmatch__]] -> 2")
	CasAssertSame(t, es, "4", "bar[1, foo[a, b]] + bar[5, foo[a, b]] /. bar[amatch_Integer, foo[cmatch__]] -> 2")
	CasAssertSame(t, es, "6 * buzz[a, b]", "bar[1, foo[a, b]] + bar[5, foo[a, b]] /. bar[amatch_Integer, foo[cmatch__]] -> amatch*buzz[cmatch]")
	CasAssertSame(t, es, "bar[3, foo[a, b]]", "bar[1, foo[a, b]] + bar[2, foo[a, b]] /. bar[amatch_Integer, foo[cmatch__]] + bar[bmatch_Integer, foo[cmatch__]] -> bar[amatch + bmatch, foo[cmatch]]")

	// Test Except
	CasAssertSame(t, es, "{5, 2, x, y, 4}", "Cases[{5, 2, 3.5, x, y, 4}, Except[_Real]]")
	CasAssertSame(t, es, "{5, 2, x, y, 4}", "Cases[{5, 2, a^b, x, y, 4}, Except[_Symbol^_Symbol]]")
	CasAssertSame(t, es, "{a, b, 0, foo[1], foo[2], x, y}", "{a, b, 0, 1, 2, x, y} /. Except[0, a_Integer] -> foo[a]")

	// Test PatternTest
	CasAssertSame(t, es, "True", "MatchQ[1, _?NumberQ]")
	CasAssertSame(t, es, "False", "MatchQ[a, _?NumberQ]")
	CasAssertSame(t, es, "True", "MatchQ[1, 1?NumberQ]")
	CasAssertSame(t, es, "False", "MatchQ[1, 1.5?NumberQ]")
	CasAssertSame(t, es, "True", "MatchQ[1.5, 1.5?NumberQ]")
	CasAssertSame(t, es, "{5,2,4.5}", "Cases[{5, 2, a^b, x, y, 4.5}, _?NumberQ]")

	// Test Condition
	CasAssertSame(t, es, "True", "MatchQ[5, _ /; True]")
	CasAssertSame(t, es, "False", "MatchQ[5, _ /; False]")
	CasAssertSame(t, es, "True", "MatchQ[5, y_ /; True]")
	CasAssertSame(t, es, "False", "MatchQ[5, y_Real /; True]")
	CasAssertSame(t, es, "True", "MatchQ[5, y_Integer /; True]")
	CasAssertSame(t, es, "False", "MatchQ[5, y_ /; y == 0]")
	CasAssertSame(t, es, "True", "MatchQ[5, y_ /; y == 5]")
	//CasAssertSame(t, es, "{1,2,3,5}", "{3, 5, 2, 1} //. {x___, y_, z_, k___} /; (Order[y, z] == -1) -> {x, z, y, k}")

	// Test special case of Orderless sequence matches
	CasAssertSame(t, es, "Null", "fooPlus[Plus[addends__]] := Hold[addends]")
	CasAssertSame(t, es, "Null", "fooList[List[addends__]] := Hold[addends]")
	CasAssertSame(t, es, "Null", "fooBlank[_[addends__]] := Hold[addends]")
	CasAssertSame(t, es, "Hold[Plus[a,b,c]]", "fooPlus[Plus[a, b, c]]")
	CasAssertSame(t, es, "Hold[a, b, c]", "fooList[List[a, b, c]]")
	CasAssertSame(t, es, "Hold[a, b, c]", "fooBlank[Plus[a, b, c]]")
}
