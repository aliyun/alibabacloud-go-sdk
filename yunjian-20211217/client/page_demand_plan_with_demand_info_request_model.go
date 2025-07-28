// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageDemandPlanWithDemandInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalStatus(v string) *PageDemandPlanWithDemandInfoRequest
	GetApprovalStatus() *string
	SetCreateTimeEnd(v string) *PageDemandPlanWithDemandInfoRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *PageDemandPlanWithDemandInfoRequest
	GetCreateTimeStart() *string
	SetCreatorStaffId(v string) *PageDemandPlanWithDemandInfoRequest
	GetCreatorStaffId() *string
	SetDemandDeliveryStatus(v string) *PageDemandPlanWithDemandInfoRequest
	GetDemandDeliveryStatus() *string
	SetDemandId(v []*int64) *PageDemandPlanWithDemandInfoRequest
	GetDemandId() []*int64
	SetDemandPlanId(v []*int64) *PageDemandPlanWithDemandInfoRequest
	GetDemandPlanId() []*int64
	SetDemandPlanStatus(v string) *PageDemandPlanWithDemandInfoRequest
	GetDemandPlanStatus() *string
	SetOperator(v string) *PageDemandPlanWithDemandInfoRequest
	GetOperator() *string
	SetPageNum(v int32) *PageDemandPlanWithDemandInfoRequest
	GetPageNum() *int32
	SetPageSize(v int32) *PageDemandPlanWithDemandInfoRequest
	GetPageSize() *int32
	SetRoCode(v string) *PageDemandPlanWithDemandInfoRequest
	GetRoCode() *string
}

type PageDemandPlanWithDemandInfoRequest struct {
	ApprovalStatus       *string  `json:"approvalStatus,omitempty" xml:"approvalStatus,omitempty"`
	CreateTimeEnd        *string  `json:"createTimeEnd,omitempty" xml:"createTimeEnd,omitempty"`
	CreateTimeStart      *string  `json:"createTimeStart,omitempty" xml:"createTimeStart,omitempty"`
	CreatorStaffId       *string  `json:"creatorStaffId,omitempty" xml:"creatorStaffId,omitempty"`
	DemandDeliveryStatus *string  `json:"demandDeliveryStatus,omitempty" xml:"demandDeliveryStatus,omitempty"`
	DemandId             []*int64 `json:"demandId,omitempty" xml:"demandId,omitempty" type:"Repeated"`
	DemandPlanId         []*int64 `json:"demandPlanId,omitempty" xml:"demandPlanId,omitempty" type:"Repeated"`
	DemandPlanStatus     *string  `json:"demandPlanStatus,omitempty" xml:"demandPlanStatus,omitempty"`
	// This parameter is required.
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// This parameter is required.
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// This parameter is required.
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RoCode   *string `json:"roCode,omitempty" xml:"roCode,omitempty"`
}

func (s PageDemandPlanWithDemandInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s PageDemandPlanWithDemandInfoRequest) GoString() string {
	return s.String()
}

func (s *PageDemandPlanWithDemandInfoRequest) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *PageDemandPlanWithDemandInfoRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *PageDemandPlanWithDemandInfoRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *PageDemandPlanWithDemandInfoRequest) GetCreatorStaffId() *string {
	return s.CreatorStaffId
}

func (s *PageDemandPlanWithDemandInfoRequest) GetDemandDeliveryStatus() *string {
	return s.DemandDeliveryStatus
}

func (s *PageDemandPlanWithDemandInfoRequest) GetDemandId() []*int64 {
	return s.DemandId
}

func (s *PageDemandPlanWithDemandInfoRequest) GetDemandPlanId() []*int64 {
	return s.DemandPlanId
}

func (s *PageDemandPlanWithDemandInfoRequest) GetDemandPlanStatus() *string {
	return s.DemandPlanStatus
}

func (s *PageDemandPlanWithDemandInfoRequest) GetOperator() *string {
	return s.Operator
}

func (s *PageDemandPlanWithDemandInfoRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *PageDemandPlanWithDemandInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageDemandPlanWithDemandInfoRequest) GetRoCode() *string {
	return s.RoCode
}

func (s *PageDemandPlanWithDemandInfoRequest) SetApprovalStatus(v string) *PageDemandPlanWithDemandInfoRequest {
	s.ApprovalStatus = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetCreateTimeEnd(v string) *PageDemandPlanWithDemandInfoRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetCreateTimeStart(v string) *PageDemandPlanWithDemandInfoRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetCreatorStaffId(v string) *PageDemandPlanWithDemandInfoRequest {
	s.CreatorStaffId = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetDemandDeliveryStatus(v string) *PageDemandPlanWithDemandInfoRequest {
	s.DemandDeliveryStatus = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetDemandId(v []*int64) *PageDemandPlanWithDemandInfoRequest {
	s.DemandId = v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetDemandPlanId(v []*int64) *PageDemandPlanWithDemandInfoRequest {
	s.DemandPlanId = v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetDemandPlanStatus(v string) *PageDemandPlanWithDemandInfoRequest {
	s.DemandPlanStatus = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetOperator(v string) *PageDemandPlanWithDemandInfoRequest {
	s.Operator = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetPageNum(v int32) *PageDemandPlanWithDemandInfoRequest {
	s.PageNum = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetPageSize(v int32) *PageDemandPlanWithDemandInfoRequest {
	s.PageSize = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) SetRoCode(v string) *PageDemandPlanWithDemandInfoRequest {
	s.RoCode = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoRequest) Validate() error {
	return dara.Validate(s)
}
