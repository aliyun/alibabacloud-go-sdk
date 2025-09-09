// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarDbReadWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *ModifyPolarDbReadWeightRequest
	GetDbInstanceId() *string
	SetDbName(v string) *ModifyPolarDbReadWeightRequest
	GetDbName() *string
	SetDbNodeIds(v string) *ModifyPolarDbReadWeightRequest
	GetDbNodeIds() *string
	SetDrdsInstanceId(v string) *ModifyPolarDbReadWeightRequest
	GetDrdsInstanceId() *string
	SetWeights(v string) *ModifyPolarDbReadWeightRequest
	GetWeights() *string
}

type ModifyPolarDbReadWeightRequest struct {
	// Polar cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The node list in the destination apsaradb for PolarDB cluster. The nodes in each cluster are separated with commas (,) and colons (:).
	//
	// This parameter is required.
	//
	// example:
	//
	// pi-****************,pi-****************:pi-****************
	DbNodeIds *string `json:"DbNodeIds,omitempty" xml:"DbNodeIds,omitempty"`
	// The ID of a DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The weight of the PolarDB cluster. Separate multiple weights with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 14,86
	Weights *string `json:"Weights,omitempty" xml:"Weights,omitempty"`
}

func (s ModifyPolarDbReadWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarDbReadWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolarDbReadWeightRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ModifyPolarDbReadWeightRequest) GetDbName() *string {
	return s.DbName
}

func (s *ModifyPolarDbReadWeightRequest) GetDbNodeIds() *string {
	return s.DbNodeIds
}

func (s *ModifyPolarDbReadWeightRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *ModifyPolarDbReadWeightRequest) GetWeights() *string {
	return s.Weights
}

func (s *ModifyPolarDbReadWeightRequest) SetDbInstanceId(v string) *ModifyPolarDbReadWeightRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetDbName(v string) *ModifyPolarDbReadWeightRequest {
	s.DbName = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetDbNodeIds(v string) *ModifyPolarDbReadWeightRequest {
	s.DbNodeIds = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetDrdsInstanceId(v string) *ModifyPolarDbReadWeightRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) SetWeights(v string) *ModifyPolarDbReadWeightRequest {
	s.Weights = &v
	return s
}

func (s *ModifyPolarDbReadWeightRequest) Validate() error {
	return dara.Validate(s)
}
