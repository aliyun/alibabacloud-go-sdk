// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DedicatedIpPoolListRequest
	GetKeyword() *string
	SetPageIndex(v int32) *DedicatedIpPoolListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *DedicatedIpPoolListRequest
	GetPageSize() *int32
}

type DedicatedIpPoolListRequest struct {
	// Search keyword for the name
	//
	// example:
	//
	// xxx
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Page index, starting from 1
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// Number of items per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DedicatedIpPoolListRequest) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolListRequest) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DedicatedIpPoolListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DedicatedIpPoolListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DedicatedIpPoolListRequest) SetKeyword(v string) *DedicatedIpPoolListRequest {
	s.Keyword = &v
	return s
}

func (s *DedicatedIpPoolListRequest) SetPageIndex(v int32) *DedicatedIpPoolListRequest {
	s.PageIndex = &v
	return s
}

func (s *DedicatedIpPoolListRequest) SetPageSize(v int32) *DedicatedIpPoolListRequest {
	s.PageSize = &v
	return s
}

func (s *DedicatedIpPoolListRequest) Validate() error {
	return dara.Validate(s)
}
