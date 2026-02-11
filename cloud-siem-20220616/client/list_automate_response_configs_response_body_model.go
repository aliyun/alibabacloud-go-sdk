// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutomateResponseConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAutomateResponseConfigsResponseBody
	GetCode() *int32
	SetData(v *ListAutomateResponseConfigsResponseBodyData) *ListAutomateResponseConfigsResponseBody
	GetData() *ListAutomateResponseConfigsResponseBodyData
	SetMessage(v string) *ListAutomateResponseConfigsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAutomateResponseConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAutomateResponseConfigsResponseBody
	GetSuccess() *bool
}

type ListAutomateResponseConfigsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *ListAutomateResponseConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAutomateResponseConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAutomateResponseConfigsResponseBody) GetData() *ListAutomateResponseConfigsResponseBodyData {
	return s.Data
}

func (s *ListAutomateResponseConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAutomateResponseConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutomateResponseConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAutomateResponseConfigsResponseBody) SetCode(v int32) *ListAutomateResponseConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetData(v *ListAutomateResponseConfigsResponseBodyData) *ListAutomateResponseConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetMessage(v string) *ListAutomateResponseConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetRequestId(v string) *ListAutomateResponseConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetSuccess(v bool) *ListAutomateResponseConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAutomateResponseConfigsResponseBodyData struct {
	// The pagination information.
	PageInfo *ListAutomateResponseConfigsResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListAutomateResponseConfigsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListAutomateResponseConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBodyData) GetPageInfo() *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *ListAutomateResponseConfigsResponseBodyData) GetResponseData() []*ListAutomateResponseConfigsResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *ListAutomateResponseConfigsResponseBodyData) SetPageInfo(v *ListAutomateResponseConfigsResponseBodyDataPageInfo) *ListAutomateResponseConfigsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyData) SetResponseData(v []*ListAutomateResponseConfigsResponseBodyDataResponseData) *ListAutomateResponseConfigsResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyData) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseData != nil {
		for _, item := range s.ResponseData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAutomateResponseConfigsResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutomateResponseConfigsResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) SetPageSize(v int32) *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) SetTotalCount(v int64) *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListAutomateResponseConfigsResponseBodyDataResponseData struct {
	// The configuration of the action that is performed after the automated response rule is hit. The value is in the JSON format.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "actionType": "doPlaybook",
	//
	//             "playbookName": "WafBlockIP",
	//
	//             "playbookUuid": "bdad6220-6584-41b2-9704-fc6584568758"
	//
	//       }
	//
	// ]
	ActionConfig *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
	// The type of the handling action. Multiple types are separated by commas (,). Valid values:
	//
	// 	- **doPlaybook**: runs the playbook.
	//
	// 	- **changeEventStatus**: changes the event status.
	//
	// 	- **changeThreatLevel**: changes the risk level of the event.
	//
	// example:
	//
	// doPlaybook,changeEventStatus
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the rule in SIEM.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The type of the automated response rule. Valid values:
	//
	// 	- **event**
	//
	// 	- **alert**
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The type of the view. Valid values:
	//
	// 0: the current Alibaba Cloud account
	//
	// 1: the global account
	//
	// example:
	//
	// 1
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The trigger condition of the automated response rule. The value is in the JSON format.
	//
	// example:
	//
	// [{"left":{"value":"alert_name"},"operator":"containsString","right":{"value":"webshell_online"}}]
	ExecutionCondition *string `json:"ExecutionCondition,omitempty" xml:"ExecutionCondition,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the automated response rule.
	//
	// example:
	//
	// 123
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ResponseRuleType *string `json:"ResponseRuleType,omitempty" xml:"ResponseRuleType,omitempty"`
	// The name of the automated response rule.
	//
	// example:
	//
	// cfw kill quara book
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **100**: enabled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user who created the rule.
	//
	// example:
	//
	// 17108579417****
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAutomateResponseConfigsResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetActionConfig() *string {
	return s.ActionConfig
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetActionType() *string {
	return s.ActionType
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetAutoResponseType() *string {
	return s.AutoResponseType
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetDataType() *int32 {
	return s.DataType
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetExecutionCondition() *string {
	return s.ExecutionCondition
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetResponseRuleType() *string {
	return s.ResponseRuleType
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetStatus() *int32 {
	return s.Status
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetActionConfig(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ActionConfig = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetActionType(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ActionType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetAliuid(v int64) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetAutoResponseType(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.AutoResponseType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetDataType(v int32) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.DataType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetExecutionCondition(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ExecutionCondition = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetGmtCreate(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetGmtModified(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetId(v int64) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetResponseRuleType(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ResponseRuleType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetRuleName(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.RuleName = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetStatus(v int32) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetSubUserId(v int64) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}
