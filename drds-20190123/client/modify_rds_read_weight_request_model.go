// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRdsReadWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ModifyRdsReadWeightRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *ModifyRdsReadWeightRequest
	GetDrdsInstanceId() *string
	SetInstanceNames(v string) *ModifyRdsReadWeightRequest
	GetInstanceNames() *string
	SetWeights(v string) *ModifyRdsReadWeightRequest
	GetWeights() *string
}

type ModifyRdsReadWeightRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The names of the ApsaraDB RDS for MySQL instances. Separate the names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-****************,rm-****************
	InstanceNames *string `json:"InstanceNames,omitempty" xml:"InstanceNames,omitempty"`
	// The weights of the ApsaraDB RDS for MySQL instances. Separate the weights with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 14,86
	Weights *string `json:"Weights,omitempty" xml:"Weights,omitempty"`
}

func (s ModifyRdsReadWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRdsReadWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightRequest) GetDbName() *string {
	return s.DbName
}

func (s *ModifyRdsReadWeightRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ModifyRdsReadWeightRequest) GetInstanceNames() *string {
	return s.InstanceNames
}

func (s *ModifyRdsReadWeightRequest) GetWeights() *string {
	return s.Weights
}

func (s *ModifyRdsReadWeightRequest) SetDbName(v string) *ModifyRdsReadWeightRequest {
	s.DbName = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetDrdsInstanceId(v string) *ModifyRdsReadWeightRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetInstanceNames(v string) *ModifyRdsReadWeightRequest {
	s.InstanceNames = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) SetWeights(v string) *ModifyRdsReadWeightRequest {
	s.Weights = &v
	return s
}

func (s *ModifyRdsReadWeightRequest) Validate() error {
	return dara.Validate(s)
}
