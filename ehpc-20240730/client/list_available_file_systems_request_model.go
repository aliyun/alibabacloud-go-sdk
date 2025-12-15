// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableFileSystemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAvailableFileSystemsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAvailableFileSystemsRequest
	GetPageSize() *int32
}

type ListAvailableFileSystemsRequest struct {
	// The page number of the page to return. Page starts from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAvailableFileSystemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAvailableFileSystemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAvailableFileSystemsRequest) SetPageNumber(v int32) *ListAvailableFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableFileSystemsRequest) SetPageSize(v int32) *ListAvailableFileSystemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAvailableFileSystemsRequest) Validate() error {
	return dara.Validate(s)
}
