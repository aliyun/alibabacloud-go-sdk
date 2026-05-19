// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostByCostCenterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsumeAmountList(v []*QueryCostByCostCenterResponseBodyConsumeAmountList) *QueryCostByCostCenterResponseBody
	GetConsumeAmountList() []*QueryCostByCostCenterResponseBodyConsumeAmountList
	SetMetadata(v interface{}) *QueryCostByCostCenterResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *QueryCostByCostCenterResponseBody
	GetRequestId() *string
	SetTotalAmount(v string) *QueryCostByCostCenterResponseBody
	GetTotalAmount() *string
}

type QueryCostByCostCenterResponseBody struct {
	ConsumeAmountList []*QueryCostByCostCenterResponseBodyConsumeAmountList `json:"ConsumeAmountList,omitempty" xml:"ConsumeAmountList,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 130296.64
	TotalAmount *string `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
}

func (s QueryCostByCostCenterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCostByCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostByCostCenterResponseBody) GetConsumeAmountList() []*QueryCostByCostCenterResponseBodyConsumeAmountList {
	return s.ConsumeAmountList
}

func (s *QueryCostByCostCenterResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *QueryCostByCostCenterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCostByCostCenterResponseBody) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *QueryCostByCostCenterResponseBody) SetConsumeAmountList(v []*QueryCostByCostCenterResponseBodyConsumeAmountList) *QueryCostByCostCenterResponseBody {
	s.ConsumeAmountList = v
	return s
}

func (s *QueryCostByCostCenterResponseBody) SetMetadata(v interface{}) *QueryCostByCostCenterResponseBody {
	s.Metadata = v
	return s
}

func (s *QueryCostByCostCenterResponseBody) SetRequestId(v string) *QueryCostByCostCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostByCostCenterResponseBody) SetTotalAmount(v string) *QueryCostByCostCenterResponseBody {
	s.TotalAmount = &v
	return s
}

func (s *QueryCostByCostCenterResponseBody) Validate() error {
	if s.ConsumeAmountList != nil {
		for _, item := range s.ConsumeAmountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCostByCostCenterResponseBodyConsumeAmountList struct {
	// example:
	//
	// 20
	AllocatedAmount *string `json:"AllocatedAmount,omitempty" xml:"AllocatedAmount,omitempty"`
	// example:
	//
	// 15945703968#
	CostCenterCode *string `json:"CostCenterCode,omitempty" xml:"CostCenterCode,omitempty"`
	// example:
	//
	// 530658
	CostCenterId   *int64  `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// example:
	//
	// 80
	DirectAmount *string `json:"DirectAmount,omitempty" xml:"DirectAmount,omitempty"`
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 1857464601594004
	OwnerAccountId   *int64  `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	OwnerAccountName *string `json:"OwnerAccountName,omitempty" xml:"OwnerAccountName,omitempty"`
	// example:
	//
	// 637537
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
	// example:
	//
	// 6375371
	PreCostCenterId *int64 `json:"PreCostCenterId,omitempty" xml:"PreCostCenterId,omitempty"`
	// example:
	//
	// 100
	TotalAllocatedAmount *string `json:"TotalAllocatedAmount,omitempty" xml:"TotalAllocatedAmount,omitempty"`
	// example:
	//
	// 0.01
	TotalAllocatedAmountPercent *string `json:"TotalAllocatedAmountPercent,omitempty" xml:"TotalAllocatedAmountPercent,omitempty"`
}

func (s QueryCostByCostCenterResponseBodyConsumeAmountList) String() string {
	return dara.Prettify(s)
}

func (s QueryCostByCostCenterResponseBodyConsumeAmountList) GoString() string {
	return s.String()
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetAllocatedAmount() *string {
	return s.AllocatedAmount
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetCostCenterCode() *string {
	return s.CostCenterCode
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetDirectAmount() *string {
	return s.DirectAmount
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetLevel() *int32 {
	return s.Level
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetOwnerAccountName() *string {
	return s.OwnerAccountName
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetParentCostCenterId() *int64 {
	return s.ParentCostCenterId
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetPreCostCenterId() *int64 {
	return s.PreCostCenterId
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetTotalAllocatedAmount() *string {
	return s.TotalAllocatedAmount
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) GetTotalAllocatedAmountPercent() *string {
	return s.TotalAllocatedAmountPercent
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetAllocatedAmount(v string) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.AllocatedAmount = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetCostCenterCode(v string) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.CostCenterCode = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetCostCenterId(v int64) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetCostCenterName(v string) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.CostCenterName = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetDirectAmount(v string) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.DirectAmount = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetLevel(v int32) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.Level = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetOwnerAccountId(v int64) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetOwnerAccountName(v string) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.OwnerAccountName = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetParentCostCenterId(v int64) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.ParentCostCenterId = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetPreCostCenterId(v int64) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.PreCostCenterId = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetTotalAllocatedAmount(v string) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.TotalAllocatedAmount = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) SetTotalAllocatedAmountPercent(v string) *QueryCostByCostCenterResponseBodyConsumeAmountList {
	s.TotalAllocatedAmountPercent = &v
	return s
}

func (s *QueryCostByCostCenterResponseBodyConsumeAmountList) Validate() error {
	return dara.Validate(s)
}
