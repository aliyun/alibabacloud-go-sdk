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
	// The unique identifier of the Data Studio node.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the Data Studio node.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number, starting from 1. Default value: 1.
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
