// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListComponentsRequest
	GetName() *string
	SetPageNumber(v int32) *ListComponentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListComponentsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListComponentsRequest
	GetProjectId() *int64
}

type ListComponentsRequest struct {
	// example:
	//
	// xm_create_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentsRequest) GetName() *string {
	return s.Name
}

func (s *ListComponentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListComponentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComponentsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListComponentsRequest) SetName(v string) *ListComponentsRequest {
	s.Name = &v
	return s
}

func (s *ListComponentsRequest) SetPageNumber(v int32) *ListComponentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComponentsRequest) SetPageSize(v int32) *ListComponentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentsRequest) SetProjectId(v int64) *ListComponentsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListComponentsRequest) Validate() error {
	return dara.Validate(s)
}
