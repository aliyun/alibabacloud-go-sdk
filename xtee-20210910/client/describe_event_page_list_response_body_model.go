// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventPageListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeEventPageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeEventPageListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeEventPageListResponseBodyResultObject) *DescribeEventPageListResponseBody
	GetResultObject() []*DescribeEventPageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeEventPageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeEventPageListResponseBody
	GetTotalPage() *int32
}

type DescribeEventPageListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeEventPageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 9
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeEventPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeEventPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventPageListResponseBody) GetResultObject() []*DescribeEventPageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventPageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeEventPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeEventPageListResponseBody) SetRequestId(v string) *DescribeEventPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventPageListResponseBody) SetCurrentPage(v int32) *DescribeEventPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventPageListResponseBody) SetPageSize(v int32) *DescribeEventPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventPageListResponseBody) SetResultObject(v []*DescribeEventPageListResponseBodyResultObject) *DescribeEventPageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventPageListResponseBody) SetTotalItem(v int32) *DescribeEventPageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeEventPageListResponseBody) SetTotalPage(v int32) *DescribeEventPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeEventPageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventPageListResponseBodyResultObject struct {
	// Object
	Children []*DescribeEventPageListResponseBodyResultObjectChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Event status.
	//
	// example:
	//
	// ONLINE
	EventStatus *string `json:"eventStatus,omitempty" xml:"eventStatus,omitempty"`
	// Event type.
	//
	// example:
	//
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Database ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Total number of rules.
	//
	// example:
	//
	// 10
	RuleCount *int32 `json:"ruleCount,omitempty" xml:"ruleCount,omitempty"`
	// Template code
	//
	// example:
	//
	// register
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// Template name.
	//
	// example:
	//
	// 注册模版
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// Template type
	//
	// example:
	//
	// UNIVERSAL
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
	// Number of customer authorizations
	//
	// example:
	//
	// 1
	UserCount *int32 `json:"userCount,omitempty" xml:"userCount,omitempty"`
}

func (s DescribeEventPageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventPageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventPageListResponseBodyResultObject) GetChildren() []*DescribeEventPageListResponseBodyResultObjectChildren {
	return s.Children
}

func (s *DescribeEventPageListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeEventPageListResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEventPageListResponseBodyResultObject) GetEventStatus() *string {
	return s.EventStatus
}

func (s *DescribeEventPageListResponseBodyResultObject) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEventPageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeEventPageListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeEventPageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventPageListResponseBodyResultObject) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeEventPageListResponseBodyResultObject) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeEventPageListResponseBodyResultObject) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeEventPageListResponseBodyResultObject) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeEventPageListResponseBodyResultObject) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DescribeEventPageListResponseBodyResultObject) SetChildren(v []*DescribeEventPageListResponseBodyResultObjectChildren) *DescribeEventPageListResponseBodyResultObject {
	s.Children = v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetEventCode(v string) *DescribeEventPageListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetEventName(v string) *DescribeEventPageListResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetEventStatus(v string) *DescribeEventPageListResponseBodyResultObject {
	s.EventStatus = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetEventType(v string) *DescribeEventPageListResponseBodyResultObject {
	s.EventType = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeEventPageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetGmtModified(v int64) *DescribeEventPageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetId(v int64) *DescribeEventPageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetRuleCount(v int32) *DescribeEventPageListResponseBodyResultObject {
	s.RuleCount = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetTemplateCode(v string) *DescribeEventPageListResponseBodyResultObject {
	s.TemplateCode = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetTemplateName(v string) *DescribeEventPageListResponseBodyResultObject {
	s.TemplateName = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetTemplateType(v string) *DescribeEventPageListResponseBodyResultObject {
	s.TemplateType = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) SetUserCount(v int32) *DescribeEventPageListResponseBodyResultObject {
	s.UserCount = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeEventPageListResponseBodyResultObjectChildren struct {
	// Event code.
	//
	// example:
	//
	// de_aamexg3015
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 测试
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Event status.
	//
	// example:
	//
	// ONLINE
	EventStatus *string `json:"eventStatus,omitempty" xml:"eventStatus,omitempty"`
	// Event type.
	//
	// example:
	//
	// BYPASS
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time
	//
	// example:
	//
	// 1621578648000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 334
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Total number of rules.
	//
	// example:
	//
	// 10
	RuleCount *int64 `json:"ruleCount,omitempty" xml:"ruleCount,omitempty"`
}

func (s DescribeEventPageListResponseBodyResultObjectChildren) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventPageListResponseBodyResultObjectChildren) GoString() string {
	return s.String()
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) GetEventStatus() *string {
	return s.EventStatus
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) GetRuleCount() *int64 {
	return s.RuleCount
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) SetEventCode(v string) *DescribeEventPageListResponseBodyResultObjectChildren {
	s.EventCode = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) SetEventName(v string) *DescribeEventPageListResponseBodyResultObjectChildren {
	s.EventName = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) SetEventStatus(v string) *DescribeEventPageListResponseBodyResultObjectChildren {
	s.EventStatus = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) SetEventType(v string) *DescribeEventPageListResponseBodyResultObjectChildren {
	s.EventType = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) SetGmtCreate(v int64) *DescribeEventPageListResponseBodyResultObjectChildren {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) SetGmtModified(v int64) *DescribeEventPageListResponseBodyResultObjectChildren {
	s.GmtModified = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) SetId(v int64) *DescribeEventPageListResponseBodyResultObjectChildren {
	s.Id = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) SetRuleCount(v int64) *DescribeEventPageListResponseBodyResultObjectChildren {
	s.RuleCount = &v
	return s
}

func (s *DescribeEventPageListResponseBodyResultObjectChildren) Validate() error {
	return dara.Validate(s)
}
