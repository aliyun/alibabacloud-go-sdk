// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYikeWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListYikeWorkspacesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListYikeWorkspacesRequest
	GetPageSize() *int32
}

type ListYikeWorkspacesRequest struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListYikeWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListYikeWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListYikeWorkspacesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListYikeWorkspacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListYikeWorkspacesRequest) SetPageNo(v int32) *ListYikeWorkspacesRequest {
	s.PageNo = &v
	return s
}

func (s *ListYikeWorkspacesRequest) SetPageSize(v int32) *ListYikeWorkspacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListYikeWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
