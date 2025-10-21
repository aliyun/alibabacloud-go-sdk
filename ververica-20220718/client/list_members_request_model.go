// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *ListMembersRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMembersRequest
	GetPageSize() *int32
}

type ListMembersRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMembersRequest) GoString() string {
	return s.String()
}

func (s *ListMembersRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMembersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMembersRequest) SetPageIndex(v int32) *ListMembersRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMembersRequest) SetPageSize(v int32) *ListMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListMembersRequest) Validate() error {
	return dara.Validate(s)
}
