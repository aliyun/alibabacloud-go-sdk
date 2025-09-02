// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListFileTypeRequest
	GetKeyword() *string
	SetLocale(v string) *ListFileTypeRequest
	GetLocale() *string
	SetPageNumber(v int32) *ListFileTypeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFileTypeRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListFileTypeRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *ListFileTypeRequest
	GetProjectIdentifier() *string
}

type ListFileTypeRequest struct {
	// The name of the node type. You can log on to the DataWorks console, go to the DataStudio page, and then view the name of a specific node type on the left side of the page. Take note of the following items when you configure this parameter:
	//
	// 	- You can view the name of a specific node type, but the language specified by this parameter to present the name must be the same as the language specified by the Locale parameter.
	//
	// 	- Fuzzy match is supported.
	//
	// 	- If this parameter is not configured, the names of all node types are returned.
	//
	// example:
	//
	// ODPS SQL
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language that you use for the query. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to view the workspace ID. You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace page to view the workspace name. You must configure either this parameter or the ProjectId parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s ListFileTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileTypeRequest) GoString() string {
	return s.String()
}

func (s *ListFileTypeRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListFileTypeRequest) GetLocale() *string {
	return s.Locale
}

func (s *ListFileTypeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFileTypeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileTypeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListFileTypeRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *ListFileTypeRequest) SetKeyword(v string) *ListFileTypeRequest {
	s.Keyword = &v
	return s
}

func (s *ListFileTypeRequest) SetLocale(v string) *ListFileTypeRequest {
	s.Locale = &v
	return s
}

func (s *ListFileTypeRequest) SetPageNumber(v int32) *ListFileTypeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFileTypeRequest) SetPageSize(v int32) *ListFileTypeRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileTypeRequest) SetProjectId(v int64) *ListFileTypeRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFileTypeRequest) SetProjectIdentifier(v string) *ListFileTypeRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *ListFileTypeRequest) Validate() error {
	return dara.Validate(s)
}
