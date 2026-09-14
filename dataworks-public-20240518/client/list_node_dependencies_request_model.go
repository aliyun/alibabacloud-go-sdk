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
	// The unique identifier of the DataStudio node.
	//
	// 	Notice: This field was of the Long type in SDK versions earlier than 8.0.0 and was changed to the String type in SDK 8.0.0 and later. **This change does not affect normal SDK usage, and the parameter is still returned in the type defined in the SDK**. Only when you upgrade across SDK version 8.0.0, the type change may cause project compilation failures, and you must manually correct the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
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
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace settings page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace for this API call.
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
