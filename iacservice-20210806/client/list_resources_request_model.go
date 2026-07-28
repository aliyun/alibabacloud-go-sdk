// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourcesRequest
	GetPageSize() *int32
	SetSourceType(v string) *ListResourcesRequest
	GetSourceType() *string
	SetSourceValue(v string) *ListResourcesRequest
	GetSourceValue() *string
	SetSpecType(v string) *ListResourcesRequest
	GetSpecType() *string
}

type ListResourcesRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The type of the resource source. Valid values:
	//
	// - ExportTaskId: resource export ID
	//
	// - TaskId: Module execution task ID
	//
	// - StatePath: the OSS path where the resource state is stored.
	//
	// This parameter is required.
	//
	// example:
	//
	// TaskId
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The specific value of the resource source.
	//
	// - If sourceType is set to ExportTaskId, the format is ExportTaskId:Version.
	//
	// - If sourceType is set to TaskId, the format is TaskId.
	//
	// - If sourceType is set to StatePath, the format is the download URL of the State file.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-235436dsfdgd
	SourceValue *string `json:"sourceValue,omitempty" xml:"sourceValue,omitempty"`
	// The specification that resource properties follow in the response. Valid values: CloudSpec, Terraform.
	//
	// Default value: CloudSpec.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudSpec
	SpecType *string `json:"specType,omitempty" xml:"specType,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListResourcesRequest) GetSourceValue() *string {
	return s.SourceValue
}

func (s *ListResourcesRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetSourceType(v string) *ListResourcesRequest {
	s.SourceType = &v
	return s
}

func (s *ListResourcesRequest) SetSourceValue(v string) *ListResourcesRequest {
	s.SourceValue = &v
	return s
}

func (s *ListResourcesRequest) SetSpecType(v string) *ListResourcesRequest {
	s.SpecType = &v
	return s
}

func (s *ListResourcesRequest) Validate() error {
	return dara.Validate(s)
}
