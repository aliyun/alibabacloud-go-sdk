// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourceServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceServicesRequest
	GetPageSize() *int32
}

type ListResourceServicesRequest struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListResourceServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServicesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceServicesRequest) SetPageNumber(v int32) *ListResourceServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceServicesRequest) SetPageSize(v int32) *ListResourceServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceServicesRequest) Validate() error {
	return dara.Validate(s)
}
