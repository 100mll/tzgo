// Copyright (c) 2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc
//

package micheline

import (
	"testing"
)

type typedefTest struct {
	Name string
	Spec string
	Want string
}

var typedefInfo = []typedefTest{
	// scalars
	//   int
	{
		Name: "int",
		Spec: `{"annots": ["%payoutDelay"],"prim": "int"}`,
		Want: `{"name":"payoutDelay","type":"int"}`,
	},
	//   nat
	{
		Name: "nat",
		Spec: `{"annots": ["%payoutFrequency"],"prim": "nat"}`,
		Want: `{"name":"payoutFrequency","type":"nat"}`,
	},
	//   string
	{
		Name: "string",
		Spec: `{"annots": ["%name"],"prim": "string"}`,
		Want: `{"name":"name","type":"string"}`,
	},
	//   bytes
	{
		Name: "bytes",
		Spec: `{"annots": ["%bakerName"],"prim": "bytes"}`,
		Want: `{"name":"bakerName","type":"bytes"}`,
	},
	//   mutez
	{
		Name: "mutez",
		Spec: `{"annots": ["%signup_fee"],"prim": "mutez"}`,
		Want: `{"name":"signup_fee","type":"mutez"}`,
	},
	//   bool
	{
		Name: "bool",
		Spec: `{"annots": ["%bakerChargesTransactionFee"],"prim": "bool"}`,
		Want: `{"name":"bakerChargesTransactionFee","type":"bool"}`,
	},
	//   key_hash
	{
		Name: "key_hash",
		Spec: `{"annots": ["%baker"],"prim": "key_hash"}`,
		Want: `{"name":"baker","type":"key_hash"}`,
	},
	//   timestamp
	{
		Name: "timestamp",
		Spec: `{"annots": ["%last_update"],"prim": "timestamp"}`,
		Want: `{"name":"last_update","type":"timestamp"}`,
	},
	//   address
	{
		Name: "address",
		Spec: `{"annots": ["%reporterAccount"],"prim": "address"}`,
		Want: `{"name":"reporterAccount","type":"address"}`,
	},
	//   key
	{
		Name: "key",
		Spec: `{"annots": ["%pour_authorizer"],"prim": "key"}`,
		Want: `{"name":"pour_authorizer","type":"key"}`,
	},
	//   unit
	//   signature
	{
		Name: "signature",
		Spec: `{"args":[{"args":[{"prim":"nat"},{"args":[{"prim":"key"},{"prim":"signature"}],"prim": "pair"}],"prim": "pair"}],"prim": "pair"}`,
		Want: `{"name":"","type":"struct","args":[{"name":"0","type":"nat"},{"name":"1","type":"key"},{"name":"2","type":"signature"}]}`,
	},
	//   chain_id
	//   bls12_381_g1
	//   bls12_381_g2
	//   bls12_381_fr
	{
		Name: "bls",
		Spec: `{"annots":["%g2"],"prim":"bls12_381_g2"}`,
		Want: `{"name":"g2","type":"bls12_381_g2"}`,
	},
	//   sapling_state
	{
		Name: "sapling_state",
		Spec: `{"prim":"sapling_state","args":[{"int":"8"}]}`,
		Want: `{"name":"","type":"sapling_state(8)"}`,
	},
	//   sapling_transaction
	{
		Name: "sapling_transaction",
		Spec: `{"prim":"sapling_transaction","args":[{"int":"8"}]}`,
		Want: `{"name":"","type":"sapling_transaction(8)"}`,
	},
	//   never
	{
		Name: "never",
		Spec: `{"prim":"never"}`,
		Want: `{"name":"","type":"never"}`,
	},
	// set
	{
		Name: "set",
		Spec: `{"annots": ["%admins"],"prim": "set", "args": [{"prim": "key_hash"}]}`,
		Want: `{"name":"admins","type":"set","args":[{"name":"@item","type":"key_hash"}]}`,
	},
	// map
	{
		Name: "map",
		Spec: `{"annots":["%approvals"],"prim":"map","args":[{"prim":"address"},{"prim":"nat"}]}`,
		Want: `{"name":"approvals","type":"map","args":[{"name":"@key","type":"address"},{"name":"@value","type":"nat"}]}`,
	},
	// bigmap with scalar key
	// bigmap with pair key
	{
		Name: "bigmap",
		Spec: `{"annots": ["%ledger"],"args": [{"args": [{"prim": "address"},{"prim": "nat"}],"prim": "pair"},{"prim": "nat"}],"prim": "big_map"}`,
		Want: `{"name": "ledger", "type": "big_map", "args":[{"name":"@key","type":"struct","args":[{"name":"0","type":"address"},{"name":"1","type":"nat"}]},{"name":"@value","type":"nat"}]}`,
	},
	// contract
	{
		Name: "contract",
		Spec: `{"annots": ["%pour_dest"],"args": [{"prim": "unit"}],"prim": "contract"}`,
		Want: `{"name":"pour_dest","type":"contract","args":[{"name":"0","type":"unit"}]}`,
	},
	// lambda, list, operation
	{
		Name: "lambda",
		Spec: `{"args": [{"args": [{"args": [{"prim": "string"},{"prim": "bytes"}],"prim": "pair"},{"args": [{"prim": "bytes"},{"prim": "bytes"}],"prim": "big_map"}],"prim": "pair"},{"args": [{"args": [{"prim": "operation"}],"prim": "list"},{"args": [{"prim": "bytes"},{"prim": "bytes"}],"prim": "big_map"}],"prim": "pair"}],"prim": "lambda"}`,
		Want: `{"name":"","type":"lambda","args":[{"name":"@param","type":"struct","args":[{"name":"0","type":"string"},{"name":"1","type":"bytes"},{"name":"2","type":"big_map","args":[{"name":"@key","type":"bytes"},{"name":"@value","type":"bytes"}]}]},{"name":"@return","type":"struct","args":[{"name":"0","type":"list","args":[{"name":"@item","type":"operation"}]},{"name":"1","type":"big_map","args":[{"name":"@key","type":"bytes"},{"name":"@value","type":"bytes"}]}]}]}`,
	},
	// ticket
	{
		Name: "ticket",
		Spec: `{"prim": "ticket", "args":[{"prim":"timestamp"}]}`,
		Want: `{"name":"","type":"ticket","args":[{"name":"@value","type":"timestamp"}]}`,
	},
	// option
	{
		Name: "option",
		Spec: `{"annots":["%reporterAccount"],"prim":"option","args":[{"prim":"address"}]}`,
		Want: `{"name":"reporterAccount","type":"address","optional":true}`,
	},
	// named union type
	{
		Name: "named-union",
		Spec: `{"args":[{"annots":["%do"],"args":[{"prim":"unit"},{"args":[{"prim":"operation"}],"prim":"list"}],"prim":"lambda"},{"annots":["%default"],"prim":"unit"}],"prim":"or"}`,
		Want: `{"name":"","type":"union","args":[{"name":"do","type":"lambda","args":[{"name":"@param","type":"unit"},{"name":"@return","type":"list","args":[{"name":"@item","type":"operation"}]}]},{"name":"default","type":"unit"}]}`,
	},
	// anonymous union type
	{
		Name: "anon-union",
		Spec: `{"args":[{"args":[{"prim":"unit"},{"prim":"operation"}],"prim":"lambda"},{"args":[{"prim":"key_hash"}],"prim":"set"}],"prim":"or"}`,
		Want: `{"name":"","type":"union","args":[{"name":"@or_0","type":"lambda","args":[{"name":"@param","type":"unit"},{"name":"@return","type":"operation"}]},{"name":"@or_1","type":"set","args":[{"name":"@item","type":"key_hash"}]}]}`,
	},
	// nested map
	{
		Name: "nested_map",
		Spec: `{"annots": ["%deck"],"args": [{"prim": "int"},{"args": [{"prim": "int"},{"prim": "int"}],"prim": "map"}],"prim": "map"}`,
		Want: `{"name":"deck","type":"map","args":[{"name":"@key","type":"int"},{"name":"@value","type":"map","args":[{"name":"@key","type":"int"},{"name":"@value","type":"int"}]}]}`,
	},
	// nested list (FA2)
	{
		Name: "nested_list",
		Spec: `{"annots": ["%transfer"],"args": [{"args": [{"annots": ["%from_"],"prim": "address"},{"annots": ["%txs"],"args": [{"args": [{"annots": ["%to_"],"prim": "address"},{"args": [{"annots": ["%token_id"],"prim": "nat"},{"annots": ["%amount"],"prim": "nat"}],"prim": "pair"}],"prim": "pair"}],"prim": "list"}],"prim": "pair"}],"prim": "list"}`,
		Want: `{"name":"transfer","type":"list","args":[{"name":"@item","type":"struct","args":[{"name":"from_","type":"address"},{"name":"txs","type":"list","args":[{"name":"@item","type":"struct","args":[{"name":"to_","type":"address"},{"name":"token_id","type":"nat"},{"name":"amount","type":"nat"}]}]}]}]}`,
	},
	// right-hand pair tree
	{
		Name: "right_hand_pair_tree",
		Spec: `{"args":[{"annots":["%tokenPool"],"prim":"nat"},{"args":[{"annots":["%xtzPool"],"prim":"mutez"},{"args":[{"annots":["%lqtTotal"],"prim":"nat"},{"args":[{"annots":["%tokenAddress"],"prim":"address"},{"annots":["%lqtAddress"],"prim":"address"}],"prim":"pair"}],"prim":"pair"}],"prim":"pair"}],"prim":"pair"}`,
		Want: `{"name":"","type":"struct","args":[{"name":"tokenPool","type":"nat"},{"name":"xtzPool","type":"mutez"},{"name":"lqtTotal","type":"nat"},{"name":"tokenAddress","type":"address"},{"name":"lqtAddress","type":"address"}]}`,
	},
}

func TestTypeRendering(t *testing.T) {
	for _, test := range typedefInfo {
		t.Run(test.Name, func(T *testing.T) {
			prim := Prim{}
			err := prim.UnmarshalJSON([]byte(test.Spec))
			if err != nil {
				T.Errorf("unmarshal error: %v", err)
			}
			have, err := Type{prim}.MarshalJSON()
			if err != nil {
				T.Errorf("render error: %v", err)
			}
			if !jsonDiff(T, have, []byte(test.Want)) {
				T.Error("render mismatch, see log for details")
				t.FailNow()
			}
		})
	}
}

type interfaceTest struct {
	Name   string
	Type   string
	Value  string
	Expect bool
}

const (
	fa1TransferType = `{"prim":"pair","args":[{"prim":"address","annots":[":from"]},{"prim":"pair","args":[{"prim":"address","annots":[":to"]},{"prim":"nat","annots":[":value"]}]}]}`
	fa2TransferType = `{"prim":"list","annots":["%transfer"],"args":[{"prim":"pair","args":[{"prim":"address","annots":["%from_"]},{"prim":"list","annots":["%txs"],"args":[{"prim":"pair","args":[{"prim":"address","annots":["%to_"]},{"prim":"pair","args":[{"prim":"nat","annots":["%token_id"]},{"prim":"nat","annots":["%amount"]}]}]}]}]}]}`
)

var interfaceInfo = []interfaceTest{
	// FA1
	{
		Name:   "fa1",
		Type:   fa1TransferType,
		Value:  `{"prim":"Pair","args":[{"bytes":"019c0931f0ebac06db1063abe651e773db6c353ce900"},{"prim":"Pair","args":[{"bytes": "0000c2bb77ac9a2c86ca05fdaea1888408471b9c1468"},{"int":"740596954"}]}]}`,
		Expect: true,
	},
	{
		Name:   "fa1_wrong_value",
		Type:   fa1TransferType,
		Value:  `[{"prim":"Pair","args":[{"string":"tz1bq5QazD43hBGxYxX8mv4sa1r1kz5sRGWz"},{"string":"tz1QgLtW1kJehxYhtypPyWJcutUTY6xZfDQf"},{"int":"21367500000"}]},{"prim":"Pair","args":[{"string":"tz1bq5QazD43hBGxYxX8mv4sa1r1kz5sRGWz"},{"string":"tz1QgLtW1kJehxYhtypPyWJcutUTY6xZfDQf"},{"int":"21367500000"}]}]`,
		Expect: false,
	},
	// FA2 single (nocomb)
	{
		Name:   "fa2_single_recv",
		Type:   fa2TransferType,
		Value:  `[{"prim":"Pair","args":[{"bytes":"01b39686f116bb35115559f7e781200850e02854c400"},[{"prim":"Pair","args":[{"bytes":"0000f0ddca1cdfa0c48c92d162f3f72b8144ee2045ba"},{"prim":"Pair","args":[{"int":"0"},{"int":"1027681"}]}]}]]}]`,
		Expect: true,
	},
	// FA2 multi (nocomb)
	{
		Name:   "fa2_multi_recv",
		Type:   fa2TransferType,
		Value:  `[{"prim":"Pair","args":[{"bytes":"01b39686f116bb35115559f7e781200850e02854c400"},[{"prim":"Pair","args":[{"bytes":"0000f0ddca1cdfa0c48c92d162f3f72b8144ee2045ba"},{"prim":"Pair","args":[{"int":"0"},{"int":"1027681"}]}]},{"prim":"Pair","args":[{"bytes":"0000f0ddca1cdfa0c48c92d162f3f72b8144ee2045ba"},{"prim":"Pair","args":[{"int":"0"},{"int":"1027681"}]}]}]]}]`,
		Expect: true,
	},
	// FA2 single (comb)
	{
		Name:   "fa2_multi_recv_comb",
		Type:   fa2TransferType,
		Value:  `[{"prim":"Pair","args":[{"bytes":"01b39686f116bb35115559f7e781200850e02854c400"},[{"prim":"Pair","args":[{"bytes":"0000f0ddca1cdfa0c48c92d162f3f72b8144ee2045ba"},{"int":"0"},{"int":"1027681"}]}]]}]`,
		Expect: true,
	},
	// TODO
	// union type
	// optional flag
	// set
	// map
	// bigmap
	// lambda
	// ticket
	// sapling
}

func TestInterfaceCheck(t *testing.T) {
	for _, test := range interfaceInfo {
		t.Run(test.Name, func(T *testing.T) {
			var typ Type
			if err := typ.Prim.UnmarshalJSON([]byte(test.Type)); err != nil {
				T.Fatalf("unmarshal type: %v", err)
			}
			var val Prim
			if err := val.UnmarshalJSON([]byte(test.Value)); err != nil {
				T.Fatalf("unmarshal value: %v", err)
			}
			if have, want := val.Implements(typ), test.Expect; have != want {
				T.Errorf("mismatch want=%t have=%t", want, have)
			}
		})
	}
}
