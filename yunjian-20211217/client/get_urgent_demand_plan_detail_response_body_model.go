// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandPlanDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetUrgentDemandPlanDetailResponseBody
	GetCode() *int64
	SetData(v *GetUrgentDemandPlanDetailResponseBodyData) *GetUrgentDemandPlanDetailResponseBody
	GetData() *GetUrgentDemandPlanDetailResponseBodyData
	SetMessage(v string) *GetUrgentDemandPlanDetailResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetUrgentDemandPlanDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GetUrgentDemandPlanDetailResponseBody
	GetTraceId() *string
}

type GetUrgentDemandPlanDetailResponseBody struct {
	// code
	//
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// body
	Data *GetUrgentDemandPlanDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 1e2b798516402440016572132e1459
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetUrgentDemandPlanDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetUrgentDemandPlanDetailResponseBody) GetData() *GetUrgentDemandPlanDetailResponseBodyData {
	return s.Data
}

func (s *GetUrgentDemandPlanDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUrgentDemandPlanDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUrgentDemandPlanDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetCode(v int64) *GetUrgentDemandPlanDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetData(v *GetUrgentDemandPlanDetailResponseBodyData) *GetUrgentDemandPlanDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetMessage(v string) *GetUrgentDemandPlanDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetSuccess(v bool) *GetUrgentDemandPlanDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetTraceId(v string) *GetUrgentDemandPlanDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUrgentDemandPlanDetailResponseBodyData struct {
	AccountDept *string `json:"accountDept,omitempty" xml:"accountDept,omitempty"`
	// example:
	//
	// 1065737167271819
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// larus_prd
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// https://xxxxx
	ApprovalUrl *string `json:"approvalUrl,omitempty" xml:"approvalUrl,omitempty"`
	// example:
	//
	// {}
	BpmSubstate           map[string]interface{} `json:"bpmSubstate,omitempty" xml:"bpmSubstate,omitempty"`
	CommodityTypeCodeList []*string              `json:"commodityTypeCodeList,omitempty" xml:"commodityTypeCodeList,omitempty" type:"Repeated"`
	// example:
	//
	// 262940
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// xxx
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// xxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// URGENT
	DetailType *string `json:"detailType,omitempty" xml:"detailType,omitempty"`
	// example:
	//
	// 2021-12-17 16:53:21
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2021-12-17 16:53:21
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 262940
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// xx
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// 11223
	PlanId   *int64  `json:"planId,omitempty" xml:"planId,omitempty"`
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// example:
	//
	// 220
	Status            *int64  `json:"status,omitempty" xml:"status,omitempty"`
	YunzhiProductName *string `json:"yunzhiProductName,omitempty" xml:"yunzhiProductName,omitempty"`
}

func (s GetUrgentDemandPlanDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetAccountDept() *string {
	return s.AccountDept
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetApprovalUrl() *string {
	return s.ApprovalUrl
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetBpmSubstate() map[string]interface{} {
	return s.BpmSubstate
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetCommodityTypeCodeList() []*string {
	return s.CommodityTypeCodeList
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetDetailType() *string {
	return s.DetailType
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetModifier() *string {
	return s.Modifier
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetModifierName() *string {
	return s.ModifierName
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetPlanId() *int64 {
	return s.PlanId
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetPlanName() *string {
	return s.PlanName
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetStatus() *int64 {
	return s.Status
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) GetYunzhiProductName() *string {
	return s.YunzhiProductName
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetAccountDept(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.AccountDept = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetAccountId(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetAccountName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetApprovalUrl(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.ApprovalUrl = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetBpmSubstate(v map[string]interface{}) *GetUrgentDemandPlanDetailResponseBodyData {
	s.BpmSubstate = v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetCommodityTypeCodeList(v []*string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.CommodityTypeCodeList = v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetCreator(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetCreatorName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetDescription(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetDetailType(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.DetailType = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetGmtCreate(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetGmtModified(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetModifier(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetModifierName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.ModifierName = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetPlanId(v int64) *GetUrgentDemandPlanDetailResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetPlanName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.PlanName = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetStatus(v int64) *GetUrgentDemandPlanDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetYunzhiProductName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.YunzhiProductName = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
