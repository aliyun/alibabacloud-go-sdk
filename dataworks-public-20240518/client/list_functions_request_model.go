// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListFunctionsRequest
	GetName() *string
	SetOwner(v string) *ListFunctionsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListFunctionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFunctionsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListFunctionsRequest
	GetProjectId() *int64
	SetType(v string) *ListFunctionsRequest
	GetType() *string
}

type ListFunctionsRequest struct {
	// Filter criteria: UDF name. Supports fuzzy search.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the owner of the UDF. This parameter specifies a filter condition.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Default value: 1. Minimum value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The user-defined function (UDF) type. This parameter specifies a filter condition.
	//
	// Valid values:
	//
	// 	- Math: mathematical operation function
	//
	// 	- Aggregate: aggregate function
	//
	// 	- String: string processing function
	//
	// 	- Date: date function
	//
	// 	- Analytic: window function
	//
	// 	- Other: other functions
	//
	// example:
	//
	// MATH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) GetName() *string {
	return s.Name
}

func (s *ListFunctionsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListFunctionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFunctionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFunctionsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListFunctionsRequest) GetType() *string {
	return s.Type
}

func (s *ListFunctionsRequest) SetName(v string) *ListFunctionsRequest {
	s.Name = &v
	return s
}

func (s *ListFunctionsRequest) SetOwner(v string) *ListFunctionsRequest {
	s.Owner = &v
	return s
}

func (s *ListFunctionsRequest) SetPageNumber(v int32) *ListFunctionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionsRequest) SetPageSize(v int32) *ListFunctionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionsRequest) SetProjectId(v int64) *ListFunctionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFunctionsRequest) SetType(v string) *ListFunctionsRequest {
	s.Type = &v
	return s
}

func (s *ListFunctionsRequest) Validate() error {
	return dara.Validate(s)
}
