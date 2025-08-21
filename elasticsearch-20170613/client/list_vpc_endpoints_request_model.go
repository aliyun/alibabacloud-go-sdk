// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListVpcEndpointsRequest
	GetPage() *int32
	SetSize(v int32) *ListVpcEndpointsRequest
	GetSize() *int32
}

type ListVpcEndpointsRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListVpcEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListVpcEndpointsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListVpcEndpointsRequest) SetPage(v int32) *ListVpcEndpointsRequest {
	s.Page = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetSize(v int32) *ListVpcEndpointsRequest {
	s.Size = &v
	return s
}

func (s *ListVpcEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
