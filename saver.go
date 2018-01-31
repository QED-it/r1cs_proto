package main

import (
	"encoding/hex"
	"fmt"

	r1cs "github.com/QED-it/r1cs_proto/gen"
	jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	system := r1cs.ConstraintSystem{
		Variables: []string{"ONE", "res", "A_0", "A_1", "A_2", "B_0", "B_1", "B_2", "compute_inner_product S_0"},
		Constraints: []*r1cs.Constraint{
			&r1cs.Constraint{
				A: map[int32]int32{2: 1},
				B: map[int32]int32{5: 1},
				C: map[int32]int32{0: 0, 8: 1},
			},
			&r1cs.Constraint{
				A: map[int32]int32{3: 1},
				B: map[int32]int32{6: 1},
				C: map[int32]int32{8: -1, 9: 1},
			},
			&r1cs.Constraint{
				A: map[int32]int32{4: 1},
				B: map[int32]int32{7: 1},
				C: map[int32]int32{9: -1, 1: 1},
			},
		},
	}
	serialized, err := proto.Marshal(&system)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(serialized))
	jsonMarshaler := jsonpb.Marshaler{}
	serializedJson, err := jsonMarshaler.MarshalToString(&system)
	if err != nil {
		panic(err)
	}
	fmt.Println(serializedJson)
}
