// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeDependenciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListNodeDependenciesRequest
	GetId() *string
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
	// 	Notice:
	//
	// The data type of this parameter is Long in SDKs earlier than V8.0.0, and is String in SDKs of V8.0.0 and later versions. **The change does not affect the normal use of the SDKs. The parameter is still returned as the type defined in the SDKs.*	- When you upgrade an SDK to a version later than V8.0.0, a compilation error may occur due to the type change. In this case, you must manually change the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of the page to return. The value of this parameter must be a positive integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. You can go to the Workspace Management page in the [DataWorks console](https://workbench.data.aliyun.com/console) to obtain the workspace ID.
	//
	// This parameter is used to specify the DataWorks workspace for the API call.
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

func (s *ListNodeDependenciesRequest) GetId() *string {
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

func (s *ListNodeDependenciesRequest) SetId(v string) *ListNodeDependenciesRequest {
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
