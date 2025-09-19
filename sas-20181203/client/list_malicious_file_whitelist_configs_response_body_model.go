// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaliciousFileWhitelistConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMaliciousFileWhitelistConfigsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMaliciousFileWhitelistConfigsResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListMaliciousFileWhitelistConfigsResponseBodyList) *ListMaliciousFileWhitelistConfigsResponseBody
	GetList() []*ListMaliciousFileWhitelistConfigsResponseBodyList
	SetMessage(v string) *ListMaliciousFileWhitelistConfigsResponseBody
	GetMessage() *string
	SetPageInfo(v *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) *ListMaliciousFileWhitelistConfigsResponseBody
	GetPageInfo() *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo
	SetRequestId(v string) *ListMaliciousFileWhitelistConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMaliciousFileWhitelistConfigsResponseBody
	GetSuccess() *bool
}

type ListMaliciousFileWhitelistConfigsResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The alert whitelist rules of sensitive files that are detected by using the agentless detection feature.
	List []*ListMaliciousFileWhitelistConfigsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5CF78A7-30AA-59DB-847F-13EE3AE7****
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

func (s ListMaliciousFileWhitelistConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMaliciousFileWhitelistConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) GetList() []*ListMaliciousFileWhitelistConfigsResponseBodyList {
	return s.List
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) GetPageInfo() *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) SetCode(v string) *ListMaliciousFileWhitelistConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) SetHttpStatusCode(v int32) *ListMaliciousFileWhitelistConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) SetList(v []*ListMaliciousFileWhitelistConfigsResponseBodyList) *ListMaliciousFileWhitelistConfigsResponseBody {
	s.List = v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) SetMessage(v string) *ListMaliciousFileWhitelistConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) SetPageInfo(v *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) *ListMaliciousFileWhitelistConfigsResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) SetRequestId(v string) *ListMaliciousFileWhitelistConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) SetSuccess(v bool) *ListMaliciousFileWhitelistConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMaliciousFileWhitelistConfigsResponseBodyList struct {
	// The number of the assets on which the rule takes effect.
	//
	// >  The value of this parameter is returned only if the value of TargetType is SELECTION_KEY.
	//
	// example:
	//
	// ALL
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the alert.
	//
	// 	- The value is fixed as ALL, which indicates all alert types.
	//
	// example:
	//
	// ALL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The field that is used in the whitelist rule.
	//
	// example:
	//
	// fileMd5
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The value of the field that is used in the whitelist rule.
	//
	// example:
	//
	// b2cf9747ee49d8d9b105cf16e078cc16
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The time when the rule was created.
	//
	// example:
	//
	// 1691719662000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the rule was modified.
	//
	// example:
	//
	// 1691719662000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The logical operator of the whitelist rule.
	//
	// 	- The value is fixed as strEqual, which indicates the equality operator (=).
	//
	// example:
	//
	// strEqual
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The feature to which this operation belongs.
	//
	// 	- The value is fixed as agentless, which indicates the agentless detection feature.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the assets on which the rule takes effect. Valid values:
	//
	// 	- ALL: all assets
	//
	// 	- SELECTION_KEY: selected assets
	//
	// example:
	//
	// ALL
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The assets on which the rule takes effect. Valid values:
	//
	// 	- ALL: all assets
	//
	// 	- Others: selected assets
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s ListMaliciousFileWhitelistConfigsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListMaliciousFileWhitelistConfigsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetCount() *string {
	return s.Count
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetEventName() *string {
	return s.EventName
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetField() *string {
	return s.Field
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetFieldValue() *string {
	return s.FieldValue
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetId() *int64 {
	return s.Id
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetOperator() *string {
	return s.Operator
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetSource() *string {
	return s.Source
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetTargetType() *string {
	return s.TargetType
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) GetTargetValue() *string {
	return s.TargetValue
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetCount(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.Count = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetEventName(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.EventName = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetField(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.Field = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetFieldValue(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.FieldValue = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetGmtCreate(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetGmtModified(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetId(v int64) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetOperator(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.Operator = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetSource(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.Source = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetTargetType(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.TargetType = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) SetTargetValue(v string) *ListMaliciousFileWhitelistConfigsResponseBodyList {
	s.TargetValue = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListMaliciousFileWhitelistConfigsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
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
	// 29
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) SetCount(v int32) *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) SetCurrentPage(v int32) *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) SetPageSize(v int32) *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) SetTotalCount(v int32) *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListMaliciousFileWhitelistConfigsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
