// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeDependenciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListNodeDependenciesRequest
	GetId() *int64
	SetPageNumber(v int32) *ListNodeDependenciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodeDependenciesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListNodeDependenciesRequest
	GetProjectId() *int64
}

type ListNodeDependenciesRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The page number, starting from 1. Default value: 1.
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
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListNodeDependenciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesRequest) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesRequest) GetId() *int64 {
	return s.Id
}

func (s *ListNodeDependenciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodeDependenciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodeDependenciesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodeDependenciesRequest) SetId(v int64) *ListNodeDependenciesRequest {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetPageNumber(v int32) *ListNodeDependenciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetPageSize(v int32) *ListNodeDependenciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetProjectId(v int64) *ListNodeDependenciesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListNodeDependenciesRequest) Validate() error {
	return dara.Validate(s)
}
