// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableThemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListTableThemeResponseBodyData) *ListTableThemeResponseBody
	GetData() *ListTableThemeResponseBodyData
	SetErrorCode(v string) *ListTableThemeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTableThemeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListTableThemeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListTableThemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTableThemeResponseBody
	GetSuccess() *bool
}

type ListTableThemeResponseBody struct {
	// The returned result.
	Data *ListTableThemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTableThemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTableThemeResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableThemeResponseBody) GetData() *ListTableThemeResponseBodyData {
	return s.Data
}

func (s *ListTableThemeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTableThemeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTableThemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTableThemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTableThemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTableThemeResponseBody) SetData(v *ListTableThemeResponseBodyData) *ListTableThemeResponseBody {
	s.Data = v
	return s
}

func (s *ListTableThemeResponseBody) SetErrorCode(v string) *ListTableThemeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTableThemeResponseBody) SetErrorMessage(v string) *ListTableThemeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTableThemeResponseBody) SetHttpStatusCode(v int32) *ListTableThemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTableThemeResponseBody) SetRequestId(v string) *ListTableThemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableThemeResponseBody) SetSuccess(v bool) *ListTableThemeResponseBody {
	s.Success = &v
	return s
}

func (s *ListTableThemeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTableThemeResponseBodyData struct {
	// The list of table levels.
	ThemeList []*ListTableThemeResponseBodyDataThemeList `json:"ThemeList,omitempty" xml:"ThemeList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTableThemeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTableThemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTableThemeResponseBodyData) GetThemeList() []*ListTableThemeResponseBodyDataThemeList {
	return s.ThemeList
}

func (s *ListTableThemeResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTableThemeResponseBodyData) SetThemeList(v []*ListTableThemeResponseBodyDataThemeList) *ListTableThemeResponseBodyData {
	s.ThemeList = v
	return s
}

func (s *ListTableThemeResponseBodyData) SetTotalCount(v int64) *ListTableThemeResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTableThemeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListTableThemeResponseBodyDataThemeList struct {
	// The time when the table level was created.
	//
	// example:
	//
	// 123432343243
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The creator of the table level.
	//
	// example:
	//
	// 123455
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The level of the table folder. Valid values: 1 and 2. The value 1 indicates the first level. The value 2 indicates the second level.
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of the table level.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ancestor node ID.
	//
	// example:
	//
	// 122
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The table theme ID.
	//
	// example:
	//
	// 123
	ThemeId *int64 `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
}

func (s ListTableThemeResponseBodyDataThemeList) String() string {
	return dara.Prettify(s)
}

func (s ListTableThemeResponseBodyDataThemeList) GoString() string {
	return s.String()
}

func (s *ListTableThemeResponseBodyDataThemeList) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *ListTableThemeResponseBodyDataThemeList) GetCreator() *string {
	return s.Creator
}

func (s *ListTableThemeResponseBodyDataThemeList) GetLevel() *int32 {
	return s.Level
}

func (s *ListTableThemeResponseBodyDataThemeList) GetName() *string {
	return s.Name
}

func (s *ListTableThemeResponseBodyDataThemeList) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListTableThemeResponseBodyDataThemeList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTableThemeResponseBodyDataThemeList) GetThemeId() *int64 {
	return s.ThemeId
}

func (s *ListTableThemeResponseBodyDataThemeList) SetCreateTimeStamp(v int64) *ListTableThemeResponseBodyDataThemeList {
	s.CreateTimeStamp = &v
	return s
}

func (s *ListTableThemeResponseBodyDataThemeList) SetCreator(v string) *ListTableThemeResponseBodyDataThemeList {
	s.Creator = &v
	return s
}

func (s *ListTableThemeResponseBodyDataThemeList) SetLevel(v int32) *ListTableThemeResponseBodyDataThemeList {
	s.Level = &v
	return s
}

func (s *ListTableThemeResponseBodyDataThemeList) SetName(v string) *ListTableThemeResponseBodyDataThemeList {
	s.Name = &v
	return s
}

func (s *ListTableThemeResponseBodyDataThemeList) SetParentId(v int64) *ListTableThemeResponseBodyDataThemeList {
	s.ParentId = &v
	return s
}

func (s *ListTableThemeResponseBodyDataThemeList) SetProjectId(v int64) *ListTableThemeResponseBodyDataThemeList {
	s.ProjectId = &v
	return s
}

func (s *ListTableThemeResponseBodyDataThemeList) SetThemeId(v int64) *ListTableThemeResponseBodyDataThemeList {
	s.ThemeId = &v
	return s
}

func (s *ListTableThemeResponseBodyDataThemeList) Validate() error {
	return dara.Validate(s)
}
