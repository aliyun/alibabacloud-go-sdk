// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutput(v string) *ListFunctionResourcesRequest
	GetOutput() *string
	SetPageNumber(v int32) *ListFunctionResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFunctionResourcesRequest
	GetPageSize() *int32
	SetResourceType(v string) *ListFunctionResourcesRequest
	GetResourceType() *string
}

type ListFunctionResourcesRequest struct {
	// The output level.
	//
	// Valid values:
	//
	// 	- simple
	//
	// 	- normal
	//
	// 	- detail
	//
	// example:
	//
	// detail
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The type of the resource.
	//
	// Valid values:
	//
	// 	- feature_generator
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- raw_file
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// feature_generator
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListFunctionResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesRequest) GetOutput() *string {
	return s.Output
}

func (s *ListFunctionResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFunctionResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFunctionResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListFunctionResourcesRequest) SetOutput(v string) *ListFunctionResourcesRequest {
	s.Output = &v
	return s
}

func (s *ListFunctionResourcesRequest) SetPageNumber(v int32) *ListFunctionResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionResourcesRequest) SetPageSize(v int32) *ListFunctionResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionResourcesRequest) SetResourceType(v string) *ListFunctionResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListFunctionResourcesRequest) Validate() error {
	return dara.Validate(s)
}
