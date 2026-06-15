// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *DedicatedIpPoolListRequest
	GetAll() *bool
	SetKeyword(v string) *DedicatedIpPoolListRequest
	GetKeyword() *string
	SetPageIndex(v int32) *DedicatedIpPoolListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *DedicatedIpPoolListRequest
	GetPageSize() *int32
	SetPoolId(v string) *DedicatedIpPoolListRequest
	GetPoolId() *string
}

type DedicatedIpPoolListRequest struct {
	// Specifies whether to return all entries.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The keyword to search for IP pools by name.
	//
	// example:
	//
	// xxx
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number to return, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 44fb3ec8-1f87-42e4-866d-e23dad9e7c9a
	PoolId *string `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
}

func (s DedicatedIpPoolListRequest) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolListRequest) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolListRequest) GetAll() *bool {
	return s.All
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

func (s *DedicatedIpPoolListRequest) GetPoolId() *string {
	return s.PoolId
}

func (s *DedicatedIpPoolListRequest) SetAll(v bool) *DedicatedIpPoolListRequest {
	s.All = &v
	return s
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

func (s *DedicatedIpPoolListRequest) SetPoolId(v string) *DedicatedIpPoolListRequest {
	s.PoolId = &v
	return s
}

func (s *DedicatedIpPoolListRequest) Validate() error {
	return dara.Validate(s)
}
