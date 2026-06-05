// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSpecificationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSpecificationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSpecificationsRequest
	GetPageSize() *int32
	SetSpecification(v string) *ListSpecificationsRequest
	GetSpecification() *string
}

type ListSpecificationsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ew.gn8t6xlarge-rb.x1p
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s ListSpecificationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSpecificationsRequest) GoString() string {
	return s.String()
}

func (s *ListSpecificationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSpecificationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSpecificationsRequest) GetSpecification() *string {
	return s.Specification
}

func (s *ListSpecificationsRequest) SetPageNumber(v int32) *ListSpecificationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSpecificationsRequest) SetPageSize(v int32) *ListSpecificationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSpecificationsRequest) SetSpecification(v string) *ListSpecificationsRequest {
	s.Specification = &v
	return s
}

func (s *ListSpecificationsRequest) Validate() error {
	return dara.Validate(s)
}
