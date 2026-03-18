// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmRecoveryPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeGtmRecoveryPlansResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmRecoveryPlansResponseBody
	GetPageSize() *int32
	SetRecoveryPlans(v *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) *DescribeGtmRecoveryPlansResponseBody
	GetRecoveryPlans() *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans
	SetRequestId(v string) *DescribeGtmRecoveryPlansResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeGtmRecoveryPlansResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeGtmRecoveryPlansResponseBody
	GetTotalPages() *int32
}

type DescribeGtmRecoveryPlansResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize      *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RecoveryPlans *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans `json:"RecoveryPlans,omitempty" xml:"RecoveryPlans,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2BA072CF-CA21-4A34-B6C2-227BE2C58079
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeGtmRecoveryPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmRecoveryPlansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmRecoveryPlansResponseBody) GetRecoveryPlans() *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	return s.RecoveryPlans
}

func (s *DescribeGtmRecoveryPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmRecoveryPlansResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeGtmRecoveryPlansResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetPageNumber(v int32) *DescribeGtmRecoveryPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetPageSize(v int32) *DescribeGtmRecoveryPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetRecoveryPlans(v *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) *DescribeGtmRecoveryPlansResponseBody {
	s.RecoveryPlans = v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetRequestId(v string) *DescribeGtmRecoveryPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetTotalItems(v int32) *DescribeGtmRecoveryPlansResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) SetTotalPages(v int32) *DescribeGtmRecoveryPlansResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBody) Validate() error {
	if s.RecoveryPlans != nil {
		if err := s.RecoveryPlans.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGtmRecoveryPlansResponseBodyRecoveryPlans struct {
	RecoveryPlan []*DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan `json:"RecoveryPlan,omitempty" xml:"RecoveryPlan,omitempty" type:"Repeated"`
}

func (s DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) GetRecoveryPlan() []*DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	return s.RecoveryPlan
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) SetRecoveryPlan(v []*DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans {
	s.RecoveryPlan = v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlans) Validate() error {
	if s.RecoveryPlan != nil {
		for _, item := range s.RecoveryPlan {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan struct {
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp       *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	FaultAddrPoolNum      *int32  `json:"FaultAddrPoolNum,omitempty" xml:"FaultAddrPoolNum,omitempty"`
	LastExecuteTime       *string `json:"LastExecuteTime,omitempty" xml:"LastExecuteTime,omitempty"`
	LastExecuteTimestamp  *int64  `json:"LastExecuteTimestamp,omitempty" xml:"LastExecuteTimestamp,omitempty"`
	LastRollbackTime      *string `json:"LastRollbackTime,omitempty" xml:"LastRollbackTime,omitempty"`
	LastRollbackTimestamp *int64  `json:"LastRollbackTimestamp,omitempty" xml:"LastRollbackTimestamp,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RecoveryPlanId        *int64  `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	Remark                *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime            *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp       *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetFaultAddrPoolNum() *int32 {
	return s.FaultAddrPoolNum
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetLastExecuteTime() *string {
	return s.LastExecuteTime
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetLastExecuteTimestamp() *int64 {
	return s.LastExecuteTimestamp
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetLastRollbackTime() *string {
	return s.LastRollbackTime
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetLastRollbackTimestamp() *int64 {
	return s.LastRollbackTimestamp
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetName() *string {
	return s.Name
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetRecoveryPlanId() *int64 {
	return s.RecoveryPlanId
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetRemark() *string {
	return s.Remark
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetStatus() *string {
	return s.Status
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetCreateTime(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetCreateTimestamp(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetFaultAddrPoolNum(v int32) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.FaultAddrPoolNum = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetLastExecuteTime(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.LastExecuteTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetLastExecuteTimestamp(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.LastExecuteTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetLastRollbackTime(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.LastRollbackTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetLastRollbackTimestamp(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.LastRollbackTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetName(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.Name = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetRecoveryPlanId(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.RecoveryPlanId = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetRemark(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.Remark = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetStatus(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.Status = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetUpdateTime(v string) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) SetUpdateTimestamp(v int64) *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlansResponseBodyRecoveryPlansRecoveryPlan) Validate() error {
	return dara.Validate(s)
}
