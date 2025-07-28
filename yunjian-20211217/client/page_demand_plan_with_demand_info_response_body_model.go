// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageDemandPlanWithDemandInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PageDemandPlanWithDemandInfoResponseBodyData) *PageDemandPlanWithDemandInfoResponseBody
	GetData() *PageDemandPlanWithDemandInfoResponseBodyData
	SetRequestId(v string) *PageDemandPlanWithDemandInfoResponseBody
	GetRequestId() *string
	SetTraceId(v string) *PageDemandPlanWithDemandInfoResponseBody
	GetTraceId() *string
}

type PageDemandPlanWithDemandInfoResponseBody struct {
	Data      *PageDemandPlanWithDemandInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TraceId   *string                                       `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s PageDemandPlanWithDemandInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PageDemandPlanWithDemandInfoResponseBody) GoString() string {
	return s.String()
}

func (s *PageDemandPlanWithDemandInfoResponseBody) GetData() *PageDemandPlanWithDemandInfoResponseBodyData {
	return s.Data
}

func (s *PageDemandPlanWithDemandInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PageDemandPlanWithDemandInfoResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *PageDemandPlanWithDemandInfoResponseBody) SetData(v *PageDemandPlanWithDemandInfoResponseBodyData) *PageDemandPlanWithDemandInfoResponseBody {
	s.Data = v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBody) SetRequestId(v string) *PageDemandPlanWithDemandInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBody) SetTraceId(v string) *PageDemandPlanWithDemandInfoResponseBody {
	s.TraceId = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type PageDemandPlanWithDemandInfoResponseBodyData struct {
	Data     []*PageDemandPlanWithDemandInfoResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNum  *int32                                              `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                              `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int64                                              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PageDemandPlanWithDemandInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PageDemandPlanWithDemandInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *PageDemandPlanWithDemandInfoResponseBodyData) GetData() []*PageDemandPlanWithDemandInfoResponseBodyDataData {
	return s.Data
}

func (s *PageDemandPlanWithDemandInfoResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *PageDemandPlanWithDemandInfoResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageDemandPlanWithDemandInfoResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *PageDemandPlanWithDemandInfoResponseBodyData) SetData(v []*PageDemandPlanWithDemandInfoResponseBodyDataData) *PageDemandPlanWithDemandInfoResponseBodyData {
	s.Data = v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyData) SetPageNum(v int32) *PageDemandPlanWithDemandInfoResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyData) SetPageSize(v int32) *PageDemandPlanWithDemandInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyData) SetTotal(v int64) *PageDemandPlanWithDemandInfoResponseBodyData {
	s.Total = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type PageDemandPlanWithDemandInfoResponseBodyDataData struct {
	AccountId             *string  `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountName           *string  `json:"accountName,omitempty" xml:"accountName,omitempty"`
	ApprovalNodeStatus    *string  `json:"approvalNodeStatus,omitempty" xml:"approvalNodeStatus,omitempty"`
	ApprovalStatus        *string  `json:"approvalStatus,omitempty" xml:"approvalStatus,omitempty"`
	Creator               *string  `json:"creator,omitempty" xml:"creator,omitempty"`
	DeliveryPlanId        *string  `json:"deliveryPlanId,omitempty" xml:"deliveryPlanId,omitempty"`
	DeliveryStatus        *string  `json:"deliveryStatus,omitempty" xml:"deliveryStatus,omitempty"`
	DemandDesc            *string  `json:"demandDesc,omitempty" xml:"demandDesc,omitempty"`
	DemandId              *int64   `json:"demandId,omitempty" xml:"demandId,omitempty"`
	DemandName            *string  `json:"demandName,omitempty" xml:"demandName,omitempty"`
	DemandPlanId          *float64 `json:"demandPlanId,omitempty" xml:"demandPlanId,omitempty"`
	GmtCreate             *string  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified           *string  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Modifier              *string  `json:"modifier,omitempty" xml:"modifier,omitempty"`
	RequirementObjectCode *string  `json:"requirementObjectCode,omitempty" xml:"requirementObjectCode,omitempty"`
	Uid                   *int64   `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s PageDemandPlanWithDemandInfoResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s PageDemandPlanWithDemandInfoResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetAccountId() *string {
	return s.AccountId
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetAccountName() *string {
	return s.AccountName
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetApprovalNodeStatus() *string {
	return s.ApprovalNodeStatus
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetCreator() *string {
	return s.Creator
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetDeliveryPlanId() *string {
	return s.DeliveryPlanId
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetDeliveryStatus() *string {
	return s.DeliveryStatus
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetDemandDesc() *string {
	return s.DemandDesc
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetDemandId() *int64 {
	return s.DemandId
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetDemandName() *string {
	return s.DemandName
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetDemandPlanId() *float64 {
	return s.DemandPlanId
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetModifier() *string {
	return s.Modifier
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetRequirementObjectCode() *string {
	return s.RequirementObjectCode
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) GetUid() *int64 {
	return s.Uid
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetAccountId(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.AccountId = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetAccountName(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.AccountName = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetApprovalNodeStatus(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.ApprovalNodeStatus = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetApprovalStatus(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.ApprovalStatus = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetCreator(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.Creator = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetDeliveryPlanId(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.DeliveryPlanId = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetDeliveryStatus(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.DeliveryStatus = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetDemandDesc(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.DemandDesc = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetDemandId(v int64) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.DemandId = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetDemandName(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.DemandName = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetDemandPlanId(v float64) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.DemandPlanId = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetGmtCreate(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.GmtCreate = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetGmtModified(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.GmtModified = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetModifier(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.Modifier = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetRequirementObjectCode(v string) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.RequirementObjectCode = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) SetUid(v int64) *PageDemandPlanWithDemandInfoResponseBodyDataData {
	s.Uid = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}
