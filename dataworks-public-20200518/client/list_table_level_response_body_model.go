// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListTableLevelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTableLevelResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListTableLevelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListTableLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTableLevelResponseBody
	GetSuccess() *bool
	SetTableLevelInfo(v *ListTableLevelResponseBodyTableLevelInfo) *ListTableLevelResponseBody
	GetTableLevelInfo() *ListTableLevelResponseBodyTableLevelInfo
}

type ListTableLevelResponseBody struct {
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
	// abc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about the table levels.
	TableLevelInfo *ListTableLevelResponseBodyTableLevelInfo `json:"TableLevelInfo,omitempty" xml:"TableLevelInfo,omitempty" type:"Struct"`
}

func (s ListTableLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTableLevelResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableLevelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTableLevelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTableLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTableLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTableLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTableLevelResponseBody) GetTableLevelInfo() *ListTableLevelResponseBodyTableLevelInfo {
	return s.TableLevelInfo
}

func (s *ListTableLevelResponseBody) SetErrorCode(v string) *ListTableLevelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTableLevelResponseBody) SetErrorMessage(v string) *ListTableLevelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTableLevelResponseBody) SetHttpStatusCode(v int32) *ListTableLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTableLevelResponseBody) SetRequestId(v string) *ListTableLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableLevelResponseBody) SetSuccess(v bool) *ListTableLevelResponseBody {
	s.Success = &v
	return s
}

func (s *ListTableLevelResponseBody) SetTableLevelInfo(v *ListTableLevelResponseBodyTableLevelInfo) *ListTableLevelResponseBody {
	s.TableLevelInfo = v
	return s
}

func (s *ListTableLevelResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTableLevelResponseBodyTableLevelInfo struct {
	// The list of table levels.
	LevelList []*ListTableLevelResponseBodyTableLevelInfoLevelList `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
	// The total number of table levels returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTableLevelResponseBodyTableLevelInfo) String() string {
	return dara.Prettify(s)
}

func (s ListTableLevelResponseBodyTableLevelInfo) GoString() string {
	return s.String()
}

func (s *ListTableLevelResponseBodyTableLevelInfo) GetLevelList() []*ListTableLevelResponseBodyTableLevelInfoLevelList {
	return s.LevelList
}

func (s *ListTableLevelResponseBodyTableLevelInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTableLevelResponseBodyTableLevelInfo) SetLevelList(v []*ListTableLevelResponseBodyTableLevelInfoLevelList) *ListTableLevelResponseBodyTableLevelInfo {
	s.LevelList = v
	return s
}

func (s *ListTableLevelResponseBodyTableLevelInfo) SetTotalCount(v int64) *ListTableLevelResponseBodyTableLevelInfo {
	s.TotalCount = &v
	return s
}

func (s *ListTableLevelResponseBodyTableLevelInfo) Validate() error {
	return dara.Validate(s)
}

type ListTableLevelResponseBodyTableLevelInfoLevelList struct {
	// The description of the table level.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The table level ID.
	//
	// example:
	//
	// 123
	LevelId *int64 `json:"LevelId,omitempty" xml:"LevelId,omitempty"`
	// The table level type. Valid values: 1 and 2. The value 1 indicates the logical level. The value 2 indicates the physical level.
	//
	// example:
	//
	// 1
	LevelType *int32 `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	// The name of the table level.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListTableLevelResponseBodyTableLevelInfoLevelList) String() string {
	return dara.Prettify(s)
}

func (s ListTableLevelResponseBodyTableLevelInfoLevelList) GoString() string {
	return s.String()
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) GetDescription() *string {
	return s.Description
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) GetLevelId() *int64 {
	return s.LevelId
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) GetLevelType() *int32 {
	return s.LevelType
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) GetName() *string {
	return s.Name
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) SetDescription(v string) *ListTableLevelResponseBodyTableLevelInfoLevelList {
	s.Description = &v
	return s
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) SetLevelId(v int64) *ListTableLevelResponseBodyTableLevelInfoLevelList {
	s.LevelId = &v
	return s
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) SetLevelType(v int32) *ListTableLevelResponseBodyTableLevelInfoLevelList {
	s.LevelType = &v
	return s
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) SetName(v string) *ListTableLevelResponseBodyTableLevelInfoLevelList {
	s.Name = &v
	return s
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) SetProjectId(v int64) *ListTableLevelResponseBodyTableLevelInfoLevelList {
	s.ProjectId = &v
	return s
}

func (s *ListTableLevelResponseBodyTableLevelInfoLevelList) Validate() error {
	return dara.Validate(s)
}
