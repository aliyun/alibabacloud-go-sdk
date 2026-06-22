// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyUserItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePropertyUserItemRequest
	GetCurrentPage() *int32
	SetForceFlush(v bool) *DescribePropertyUserItemRequest
	GetForceFlush() *bool
	SetPageSize(v int32) *DescribePropertyUserItemRequest
	GetPageSize() *int32
	SetUser(v string) *DescribePropertyUserItemRequest
	GetUser() *string
}

type DescribePropertyUserItemRequest struct {
	// The page number of the current page to display in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to forcefully refresh the data to be queried. Valid values:
	//
	// - **true**: Forcefully refresh.
	//
	// - **false**: Do not forcefully refresh.
	//
	// example:
	//
	// true
	ForceFlush *bool `json:"ForceFlush,omitempty" xml:"ForceFlush,omitempty"`
	// The maximum number of entries to display on each page in a paged query.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The account information of Asset Fingerprints.
	//
	// example:
	//
	// adm
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribePropertyUserItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyUserItemRequest) GetForceFlush() *bool {
	return s.ForceFlush
}

func (s *DescribePropertyUserItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyUserItemRequest) GetUser() *string {
	return s.User
}

func (s *DescribePropertyUserItemRequest) SetCurrentPage(v int32) *DescribePropertyUserItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyUserItemRequest) SetForceFlush(v bool) *DescribePropertyUserItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertyUserItemRequest) SetPageSize(v int32) *DescribePropertyUserItemRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyUserItemRequest) SetUser(v string) *DescribePropertyUserItemRequest {
	s.User = &v
	return s
}

func (s *DescribePropertyUserItemRequest) Validate() error {
	return dara.Validate(s)
}
