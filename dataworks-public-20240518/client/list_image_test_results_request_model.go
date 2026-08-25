// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageTestResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListImageTestResultsRequest
	GetId() *string
	SetPageNumber(v int32) *ListImageTestResultsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImageTestResultsRequest
	GetPageSize() *int32
}

type ListImageTestResultsRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// img_123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListImageTestResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageTestResultsRequest) GoString() string {
	return s.String()
}

func (s *ListImageTestResultsRequest) GetId() *string {
	return s.Id
}

func (s *ListImageTestResultsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageTestResultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageTestResultsRequest) SetId(v string) *ListImageTestResultsRequest {
	s.Id = &v
	return s
}

func (s *ListImageTestResultsRequest) SetPageNumber(v int32) *ListImageTestResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImageTestResultsRequest) SetPageSize(v int32) *ListImageTestResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListImageTestResultsRequest) Validate() error {
	return dara.Validate(s)
}
