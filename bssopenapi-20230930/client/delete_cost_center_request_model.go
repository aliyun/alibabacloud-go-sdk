// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostCenterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *DeleteCostCenterRequest
	GetCostCenterId() *int64
	SetNbid(v string) *DeleteCostCenterRequest
	GetNbid() *string
	SetOwnerAccountId(v int64) *DeleteCostCenterRequest
	GetOwnerAccountId() *int64
}

type DeleteCostCenterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s DeleteCostCenterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostCenterRequest) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *DeleteCostCenterRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DeleteCostCenterRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *DeleteCostCenterRequest) SetCostCenterId(v int64) *DeleteCostCenterRequest {
	s.CostCenterId = &v
	return s
}

func (s *DeleteCostCenterRequest) SetNbid(v string) *DeleteCostCenterRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteCostCenterRequest) SetOwnerAccountId(v int64) *DeleteCostCenterRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *DeleteCostCenterRequest) Validate() error {
	return dara.Validate(s)
}
