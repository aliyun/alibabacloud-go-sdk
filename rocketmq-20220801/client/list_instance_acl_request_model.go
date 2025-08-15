// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListInstanceAclRequest
	GetFilter() *string
	SetPageNumber(v int32) *ListInstanceAclRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceAclRequest
	GetPageSize() *int32
}

type ListInstanceAclRequest struct {
	// The condition that you specify to filter the ACLs. If you do not specify this parameter, all ACLs are queried.
	//
	// example:
	//
	// CID-TEST
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListInstanceAclRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAclRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceAclRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListInstanceAclRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceAclRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceAclRequest) SetFilter(v string) *ListInstanceAclRequest {
	s.Filter = &v
	return s
}

func (s *ListInstanceAclRequest) SetPageNumber(v int32) *ListInstanceAclRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceAclRequest) SetPageSize(v int32) *ListInstanceAclRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceAclRequest) Validate() error {
	return dara.Validate(s)
}
