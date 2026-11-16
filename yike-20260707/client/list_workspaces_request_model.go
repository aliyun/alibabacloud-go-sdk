// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListWorkspacesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListWorkspacesRequest
	GetPageSize() *int32
}

type ListWorkspacesRequest struct {
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListWorkspacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkspacesRequest) SetPageNo(v int32) *ListWorkspacesRequest {
	s.PageNo = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageSize(v int32) *ListWorkspacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
