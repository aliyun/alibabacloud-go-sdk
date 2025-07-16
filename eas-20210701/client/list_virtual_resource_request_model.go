// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListVirtualResourceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVirtualResourceRequest
	GetPageSize() *int32
	SetVirtualResourceId(v string) *ListVirtualResourceRequest
	GetVirtualResourceId() *string
	SetVirtualResourceName(v string) *ListVirtualResourceRequest
	GetVirtualResourceName() *string
}

type ListVirtualResourceRequest struct {
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
	// The ID of the virtual resource group.
	//
	// example:
	//
	// eas-vr-npovr28onap1xxxxxx
	VirtualResourceId *string `json:"VirtualResourceId,omitempty" xml:"VirtualResourceId,omitempty"`
	// The name of the virtual resource group.
	//
	// example:
	//
	// MyVirtualResource
	VirtualResourceName *string `json:"VirtualResourceName,omitempty" xml:"VirtualResourceName,omitempty"`
}

func (s ListVirtualResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualResourceRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualResourceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVirtualResourceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVirtualResourceRequest) GetVirtualResourceId() *string {
	return s.VirtualResourceId
}

func (s *ListVirtualResourceRequest) GetVirtualResourceName() *string {
	return s.VirtualResourceName
}

func (s *ListVirtualResourceRequest) SetPageNumber(v int32) *ListVirtualResourceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVirtualResourceRequest) SetPageSize(v int32) *ListVirtualResourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListVirtualResourceRequest) SetVirtualResourceId(v string) *ListVirtualResourceRequest {
	s.VirtualResourceId = &v
	return s
}

func (s *ListVirtualResourceRequest) SetVirtualResourceName(v string) *ListVirtualResourceRequest {
	s.VirtualResourceName = &v
	return s
}

func (s *ListVirtualResourceRequest) Validate() error {
	return dara.Validate(s)
}
