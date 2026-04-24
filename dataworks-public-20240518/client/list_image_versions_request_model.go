// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListImageVersionsRequest
	GetId() *string
	SetPageNumber(v int32) *ListImageVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImageVersionsRequest
	GetPageSize() *int32
}

type ListImageVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListImageVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListImageVersionsRequest) GetId() *string {
	return s.Id
}

func (s *ListImageVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageVersionsRequest) SetId(v string) *ListImageVersionsRequest {
	s.Id = &v
	return s
}

func (s *ListImageVersionsRequest) SetPageNumber(v int32) *ListImageVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImageVersionsRequest) SetPageSize(v int32) *ListImageVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListImageVersionsRequest) Validate() error {
	return dara.Validate(s)
}
