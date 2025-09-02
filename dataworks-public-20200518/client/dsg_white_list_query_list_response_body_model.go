// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListQueryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DsgWhiteListQueryListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgWhiteListQueryListResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgWhiteListQueryListResponseBody
	GetHttpStatusCode() *int32
	SetPageData(v *DsgWhiteListQueryListResponseBodyPageData) *DsgWhiteListQueryListResponseBody
	GetPageData() *DsgWhiteListQueryListResponseBodyPageData
	SetRequestId(v string) *DsgWhiteListQueryListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgWhiteListQueryListResponseBody
	GetSuccess() *bool
}

type DsgWhiteListQueryListResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The pagination information.
	PageData *DsgWhiteListQueryListResponseBodyPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Struct"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 102400001
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

func (s DsgWhiteListQueryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListQueryListResponseBody) GoString() string {
	return s.String()
}

func (s *DsgWhiteListQueryListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgWhiteListQueryListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgWhiteListQueryListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgWhiteListQueryListResponseBody) GetPageData() *DsgWhiteListQueryListResponseBodyPageData {
	return s.PageData
}

func (s *DsgWhiteListQueryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgWhiteListQueryListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgWhiteListQueryListResponseBody) SetErrorCode(v string) *DsgWhiteListQueryListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBody) SetErrorMessage(v string) *DsgWhiteListQueryListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBody) SetHttpStatusCode(v int32) *DsgWhiteListQueryListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBody) SetPageData(v *DsgWhiteListQueryListResponseBodyPageData) *DsgWhiteListQueryListResponseBody {
	s.PageData = v
	return s
}

func (s *DsgWhiteListQueryListResponseBody) SetRequestId(v string) *DsgWhiteListQueryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBody) SetSuccess(v bool) *DsgWhiteListQueryListResponseBody {
	s.Success = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DsgWhiteListQueryListResponseBodyPageData struct {
	// A collection of whitelists.
	Data []*DsgWhiteListQueryListResponseBodyPageDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of data masking whitelists.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DsgWhiteListQueryListResponseBodyPageData) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListQueryListResponseBodyPageData) GoString() string {
	return s.String()
}

func (s *DsgWhiteListQueryListResponseBodyPageData) GetData() []*DsgWhiteListQueryListResponseBodyPageDataData {
	return s.Data
}

func (s *DsgWhiteListQueryListResponseBodyPageData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgWhiteListQueryListResponseBodyPageData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgWhiteListQueryListResponseBodyPageData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DsgWhiteListQueryListResponseBodyPageData) SetData(v []*DsgWhiteListQueryListResponseBodyPageDataData) *DsgWhiteListQueryListResponseBodyPageData {
	s.Data = v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageData) SetPageNumber(v int32) *DsgWhiteListQueryListResponseBodyPageData {
	s.PageNumber = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageData) SetPageSize(v int32) *DsgWhiteListQueryListResponseBodyPageData {
	s.PageSize = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageData) SetTotalCount(v int32) *DsgWhiteListQueryListResponseBodyPageData {
	s.TotalCount = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageData) Validate() error {
	return dara.Validate(s)
}

type DsgWhiteListQueryListResponseBodyPageDataData struct {
	// The expiration time of the data masking whitelist cannot be earlier than the time when the data masking whitelist takes effect. Unit: days.
	//
	// example:
	//
	// 2024-05-10 15:46:20
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the whitelist was created.
	//
	// example:
	//
	// 2024-05-09 15:46:20
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the whitelist was modified.
	//
	// example:
	//
	// 2024-05-09 15:46:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the data masking whitelist.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the data masking rule.
	//
	// example:
	//
	// 123
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the level-2 data masking scenario.
	//
	// example:
	//
	// 123
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The time when the data masking whitelist takes effect cannot be earlier than the current time. Unit: days.
	//
	// example:
	//
	// 2024-05-09 15:46:20
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The sensitive field type.
	//
	// example:
	//
	// phone
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// A collection of user group names.
	UserGroups []*string `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s DsgWhiteListQueryListResponseBodyPageDataData) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListQueryListResponseBodyPageDataData) GoString() string {
	return s.String()
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) GetEndTime() *string {
	return s.EndTime
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) GetId() *int64 {
	return s.Id
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) GetStartTime() *string {
	return s.StartTime
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) GetType() *string {
	return s.Type
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) GetUserGroups() []*string {
	return s.UserGroups
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) SetEndTime(v string) *DsgWhiteListQueryListResponseBodyPageDataData {
	s.EndTime = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) SetGmtCreate(v string) *DsgWhiteListQueryListResponseBodyPageDataData {
	s.GmtCreate = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) SetGmtModified(v string) *DsgWhiteListQueryListResponseBodyPageDataData {
	s.GmtModified = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) SetId(v int64) *DsgWhiteListQueryListResponseBodyPageDataData {
	s.Id = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) SetRuleId(v int64) *DsgWhiteListQueryListResponseBodyPageDataData {
	s.RuleId = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) SetSceneId(v int64) *DsgWhiteListQueryListResponseBodyPageDataData {
	s.SceneId = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) SetStartTime(v string) *DsgWhiteListQueryListResponseBodyPageDataData {
	s.StartTime = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) SetType(v string) *DsgWhiteListQueryListResponseBodyPageDataData {
	s.Type = &v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) SetUserGroups(v []*string) *DsgWhiteListQueryListResponseBodyPageDataData {
	s.UserGroups = v
	return s
}

func (s *DsgWhiteListQueryListResponseBodyPageDataData) Validate() error {
	return dara.Validate(s)
}
