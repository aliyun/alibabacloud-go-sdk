// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateCostCenterResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *AllocateCostCenterResourceResponseBody
	GetCostCenterId() *int64
	SetIsSuccess(v bool) *AllocateCostCenterResourceResponseBody
	GetIsSuccess() *bool
	SetMetadata(v interface{}) *AllocateCostCenterResourceResponseBody
	GetMetadata() interface{}
	SetOwnerAccountId(v int64) *AllocateCostCenterResourceResponseBody
	GetOwnerAccountId() *int64
	SetRequestId(v string) *AllocateCostCenterResourceResponseBody
	GetRequestId() *string
}

type AllocateCostCenterResourceResponseBody struct {
	// example:
	//
	// 640403
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 1529600453335198
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateCostCenterResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostCenterResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateCostCenterResourceResponseBody) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *AllocateCostCenterResourceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *AllocateCostCenterResourceResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *AllocateCostCenterResourceResponseBody) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *AllocateCostCenterResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateCostCenterResourceResponseBody) SetCostCenterId(v int64) *AllocateCostCenterResourceResponseBody {
	s.CostCenterId = &v
	return s
}

func (s *AllocateCostCenterResourceResponseBody) SetIsSuccess(v bool) *AllocateCostCenterResourceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *AllocateCostCenterResourceResponseBody) SetMetadata(v interface{}) *AllocateCostCenterResourceResponseBody {
	s.Metadata = v
	return s
}

func (s *AllocateCostCenterResourceResponseBody) SetOwnerAccountId(v int64) *AllocateCostCenterResourceResponseBody {
	s.OwnerAccountId = &v
	return s
}

func (s *AllocateCostCenterResourceResponseBody) SetRequestId(v string) *AllocateCostCenterResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateCostCenterResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
