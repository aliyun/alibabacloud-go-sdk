// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusinessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListBusinessRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListBusinessRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBusinessRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListBusinessRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *ListBusinessRequest
	GetProjectIdentifier() *string
}

type ListBusinessRequest struct {
	// The keyword that is used to perform a fuzzy match.
	//
	// example:
	//
	// my
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace Management page to obtain the workspace ID. You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace Management page to obtain the workspace name. You must configure either this parameter or ProjectId parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s ListBusinessRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessRequest) GoString() string {
	return s.String()
}

func (s *ListBusinessRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListBusinessRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBusinessRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBusinessRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBusinessRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *ListBusinessRequest) SetKeyword(v string) *ListBusinessRequest {
	s.Keyword = &v
	return s
}

func (s *ListBusinessRequest) SetPageNumber(v int32) *ListBusinessRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBusinessRequest) SetPageSize(v int32) *ListBusinessRequest {
	s.PageSize = &v
	return s
}

func (s *ListBusinessRequest) SetProjectId(v int64) *ListBusinessRequest {
	s.ProjectId = &v
	return s
}

func (s *ListBusinessRequest) SetProjectIdentifier(v string) *ListBusinessRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *ListBusinessRequest) Validate() error {
	return dara.Validate(s)
}
