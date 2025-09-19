// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageEventOperationPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageEventOperationPageResponseBody
	GetCode() *string
	SetData(v *DescribeImageEventOperationPageResponseBodyData) *DescribeImageEventOperationPageResponseBody
	GetData() *DescribeImageEventOperationPageResponseBodyData
	SetMessage(v string) *DescribeImageEventOperationPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeImageEventOperationPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeImageEventOperationPageResponseBody
	GetSuccess() *bool
}

type DescribeImageEventOperationPageResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeImageEventOperationPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3D7C47D-3F11-57BB-90E8-E5C20C61****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeImageEventOperationPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageEventOperationPageResponseBody) GetData() *DescribeImageEventOperationPageResponseBodyData {
	return s.Data
}

func (s *DescribeImageEventOperationPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeImageEventOperationPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageEventOperationPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImageEventOperationPageResponseBody) SetCode(v string) *DescribeImageEventOperationPageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBody) SetData(v *DescribeImageEventOperationPageResponseBodyData) *DescribeImageEventOperationPageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageEventOperationPageResponseBody) SetMessage(v string) *DescribeImageEventOperationPageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBody) SetRequestId(v string) *DescribeImageEventOperationPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBody) SetSuccess(v bool) *DescribeImageEventOperationPageResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageEventOperationPageResponseBodyData struct {
	// The alert handling rules.
	List []*DescribeImageEventOperationPageResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageEventOperationPageResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s DescribeImageEventOperationPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationPageResponseBodyData) GetList() []*DescribeImageEventOperationPageResponseBodyDataList {
	return s.List
}

func (s *DescribeImageEventOperationPageResponseBodyData) GetPageInfo() *DescribeImageEventOperationPageResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeImageEventOperationPageResponseBodyData) SetList(v []*DescribeImageEventOperationPageResponseBodyDataList) *DescribeImageEventOperationPageResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyData) SetPageInfo(v *DescribeImageEventOperationPageResponseBodyDataPageInfo) *DescribeImageEventOperationPageResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeImageEventOperationPageResponseBodyDataList struct {
	// The rule conditions. The value is in the JSON format. Valid values of keys:
	//
	// 	- **condition**: the matching condition.
	//
	// 	- **type**: the matching type.
	//
	// 	- **value**: the matching value.
	//
	// example:
	//
	// [{\\"condition\\": \\"MD5\\", \\"type\\": \\"equals\\", \\"value\\": \\"0083a31cc0083a31ccf7c10367a6e783e\\"}]
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The keyword of the alert item.
	//
	// example:
	//
	// PEM
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// The name of the alert item.
	//
	// example:
	//
	// PEM
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The alert type.
	//
	// 	- Only **sensitiveFile*	- may be returned.
	//
	// example:
	//
	// sensitiveFile
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The primary key of the alert handling rule.
	//
	// example:
	//
	// 2646624
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The remarks.
	//
	// example:
	//
	// xxx
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The operation code.
	//
	// 	- Only **whitelist*	- may be returned, which means that the alert item is added to the whitelist.
	//
	// example:
	//
	// whitelist
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The application scope of the rule. The value is in the JSON format. Valid values of keys:
	//
	// 	- **type**
	//
	// 	- **value**
	//
	// example:
	//
	// {\\"type\\": \\"repo\\", \\"value\\": \\"test-aaa/shenzhen-repo-01\\"}
	Scenarios *string `json:"Scenarios,omitempty" xml:"Scenarios,omitempty"`
	// The source of the whitelist. Valid values:
	//
	// 	- **image**: image.
	//
	// 	- **agentless**: agentless detection.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeImageEventOperationPageResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationPageResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) GetConditions() *string {
	return s.Conditions
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) GetEventKey() *string {
	return s.EventKey
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) GetEventType() *string {
	return s.EventType
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) GetNote() *string {
	return s.Note
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) GetOperationCode() *string {
	return s.OperationCode
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) GetScenarios() *string {
	return s.Scenarios
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) GetSource() *string {
	return s.Source
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) SetConditions(v string) *DescribeImageEventOperationPageResponseBodyDataList {
	s.Conditions = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) SetEventKey(v string) *DescribeImageEventOperationPageResponseBodyDataList {
	s.EventKey = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) SetEventName(v string) *DescribeImageEventOperationPageResponseBodyDataList {
	s.EventName = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) SetEventType(v string) *DescribeImageEventOperationPageResponseBodyDataList {
	s.EventType = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) SetId(v int64) *DescribeImageEventOperationPageResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) SetNote(v string) *DescribeImageEventOperationPageResponseBodyDataList {
	s.Note = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) SetOperationCode(v string) *DescribeImageEventOperationPageResponseBodyDataList {
	s.OperationCode = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) SetScenarios(v string) *DescribeImageEventOperationPageResponseBodyDataList {
	s.Scenarios = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) SetSource(v string) *DescribeImageEventOperationPageResponseBodyDataList {
	s.Source = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeImageEventOperationPageResponseBodyDataPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 109
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageEventOperationPageResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationPageResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationPageResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageEventOperationPageResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageEventOperationPageResponseBodyDataPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageEventOperationPageResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeImageEventOperationPageResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeImageEventOperationPageResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataPageInfo) SetTotalCount(v int32) *DescribeImageEventOperationPageResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageEventOperationPageResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
