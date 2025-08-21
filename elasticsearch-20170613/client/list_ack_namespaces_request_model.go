// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAckNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListAckNamespacesRequest
	GetPage() *int32
	SetSize(v int32) *ListAckNamespacesRequest
	GetSize() *int32
}

type ListAckNamespacesRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAckNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAckNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListAckNamespacesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListAckNamespacesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAckNamespacesRequest) SetPage(v int32) *ListAckNamespacesRequest {
	s.Page = &v
	return s
}

func (s *ListAckNamespacesRequest) SetSize(v int32) *ListAckNamespacesRequest {
	s.Size = &v
	return s
}

func (s *ListAckNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
