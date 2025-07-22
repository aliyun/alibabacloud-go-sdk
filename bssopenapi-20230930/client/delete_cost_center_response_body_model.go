// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostCenterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *DeleteCostCenterResponseBody
	GetCostCenterId() *int64
	SetIsSuccess(v bool) *DeleteCostCenterResponseBody
	GetIsSuccess() *bool
	SetMetadata(v interface{}) *DeleteCostCenterResponseBody
	GetMetadata() interface{}
	SetOwnerAccountId(v int64) *DeleteCostCenterResponseBody
	GetOwnerAccountId() *int64
	SetRequestId(v string) *DeleteCostCenterResponseBody
	GetRequestId() *string
}

type DeleteCostCenterResponseBody struct {
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// C1BD134E-D914-6AE0-1901-AEB2A99FA205
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCostCenterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterResponseBody) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *DeleteCostCenterResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteCostCenterResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *DeleteCostCenterResponseBody) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *DeleteCostCenterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCostCenterResponseBody) SetCostCenterId(v int64) *DeleteCostCenterResponseBody {
	s.CostCenterId = &v
	return s
}

func (s *DeleteCostCenterResponseBody) SetIsSuccess(v bool) *DeleteCostCenterResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteCostCenterResponseBody) SetMetadata(v interface{}) *DeleteCostCenterResponseBody {
	s.Metadata = v
	return s
}

func (s *DeleteCostCenterResponseBody) SetOwnerAccountId(v int64) *DeleteCostCenterResponseBody {
	s.OwnerAccountId = &v
	return s
}

func (s *DeleteCostCenterResponseBody) SetRequestId(v string) *DeleteCostCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCostCenterResponseBody) Validate() error {
	return dara.Validate(s)
}
