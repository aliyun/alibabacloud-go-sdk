// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandPlanListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetUrgentDemandPlanListResponseBody
	GetCode() *int64
	SetData(v *GetUrgentDemandPlanListResponseBodyData) *GetUrgentDemandPlanListResponseBody
	GetData() *GetUrgentDemandPlanListResponseBodyData
	SetMessage(v string) *GetUrgentDemandPlanListResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetUrgentDemandPlanListResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GetUrgentDemandPlanListResponseBody
	GetTraceId() *string
}

type GetUrgentDemandPlanListResponseBody struct {
	// example:
	//
	// 0
	Code *int64                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetUrgentDemandPlanListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2127968716405850615204514e9332
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetUrgentDemandPlanListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetUrgentDemandPlanListResponseBody) GetData() *GetUrgentDemandPlanListResponseBodyData {
	return s.Data
}

func (s *GetUrgentDemandPlanListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUrgentDemandPlanListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUrgentDemandPlanListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetUrgentDemandPlanListResponseBody) SetCode(v int64) *GetUrgentDemandPlanListResponseBody {
	s.Code = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBody) SetData(v *GetUrgentDemandPlanListResponseBodyData) *GetUrgentDemandPlanListResponseBody {
	s.Data = v
	return s
}

func (s *GetUrgentDemandPlanListResponseBody) SetMessage(v string) *GetUrgentDemandPlanListResponseBody {
	s.Message = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBody) SetSuccess(v bool) *GetUrgentDemandPlanListResponseBody {
	s.Success = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBody) SetTraceId(v string) *GetUrgentDemandPlanListResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUrgentDemandPlanListResponseBodyData struct {
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 2
	Pages   *int64                                            `json:"pages,omitempty" xml:"pages,omitempty"`
	Records []*GetUrgentDemandPlanListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 15
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetUrgentDemandPlanListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListResponseBodyData) GetCurrent() *int64 {
	return s.Current
}

func (s *GetUrgentDemandPlanListResponseBodyData) GetPages() *int64 {
	return s.Pages
}

func (s *GetUrgentDemandPlanListResponseBodyData) GetRecords() []*GetUrgentDemandPlanListResponseBodyDataRecords {
	return s.Records
}

func (s *GetUrgentDemandPlanListResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *GetUrgentDemandPlanListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetCurrent(v int64) *GetUrgentDemandPlanListResponseBodyData {
	s.Current = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetPages(v int64) *GetUrgentDemandPlanListResponseBodyData {
	s.Pages = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetRecords(v []*GetUrgentDemandPlanListResponseBodyDataRecords) *GetUrgentDemandPlanListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetSize(v int64) *GetUrgentDemandPlanListResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetTotal(v int64) *GetUrgentDemandPlanListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetUrgentDemandPlanListResponseBodyDataRecords struct {
	// example:
	//
	// 1705524002740212
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// xxxx
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// ALIYUN
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// example:
	//
	// https://xxx
	ApprovalUrl *string `json:"approvalUrl,omitempty" xml:"approvalUrl,omitempty"`
	// example:
	//
	// 1111
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// xxxx
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// xxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2021-12-20 10:29:50
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2021-12-20 10:29:50
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// xxxx
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// xxxx
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	PlanId       *int64  `json:"planId,omitempty" xml:"planId,omitempty"`
	PlanName     *string `json:"planName,omitempty" xml:"planName,omitempty"`
	Status       *int64  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetUrgentDemandPlanListResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetAccountName() *string {
	return s.AccountName
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetAccountType() *string {
	return s.AccountType
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetApprovalUrl() *string {
	return s.ApprovalUrl
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetCreator() *string {
	return s.Creator
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetDescription() *string {
	return s.Description
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetModifier() *string {
	return s.Modifier
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetModifierName() *string {
	return s.ModifierName
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetPlanId() *int64 {
	return s.PlanId
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetPlanName() *string {
	return s.PlanName
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) GetStatus() *int64 {
	return s.Status
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetAccountId(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.AccountId = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetAccountName(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.AccountName = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetAccountType(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.AccountType = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetApprovalUrl(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.ApprovalUrl = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetCreator(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetCreatorName(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.CreatorName = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetDescription(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetGmtCreate(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetGmtModified(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetModifier(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.Modifier = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetModifierName(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.ModifierName = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetPlanId(v int64) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.PlanId = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetPlanName(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.PlanName = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetStatus(v int64) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
