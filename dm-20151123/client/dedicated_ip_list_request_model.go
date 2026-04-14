// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DedicatedIpListRequest
	GetKeyword() *string
	SetPageIndex(v int32) *DedicatedIpListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *DedicatedIpListRequest
	GetPageSize() *int32
}

type DedicatedIpListRequest struct {
	// IP search keyword
	//
	// example:
	//
	// xxx
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Pagination index, starting from 1
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DedicatedIpListRequest) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpListRequest) GoString() string {
	return s.String()
}

func (s *DedicatedIpListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DedicatedIpListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DedicatedIpListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DedicatedIpListRequest) SetKeyword(v string) *DedicatedIpListRequest {
	s.Keyword = &v
	return s
}

func (s *DedicatedIpListRequest) SetPageIndex(v int32) *DedicatedIpListRequest {
	s.PageIndex = &v
	return s
}

func (s *DedicatedIpListRequest) SetPageSize(v int32) *DedicatedIpListRequest {
	s.PageSize = &v
	return s
}

func (s *DedicatedIpListRequest) Validate() error {
	return dara.Validate(s)
}
