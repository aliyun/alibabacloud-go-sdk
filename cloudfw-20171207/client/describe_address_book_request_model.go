// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddressBookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainPort(v string) *DescribeAddressBookRequest
	GetContainPort() *string
	SetCurrentPage(v string) *DescribeAddressBookRequest
	GetCurrentPage() *string
	SetGroupType(v string) *DescribeAddressBookRequest
	GetGroupType() *string
	SetGroupUuid(v string) *DescribeAddressBookRequest
	GetGroupUuid() *string
	SetLang(v string) *DescribeAddressBookRequest
	GetLang() *string
	SetPageSize(v string) *DescribeAddressBookRequest
	GetPageSize() *string
	SetQuery(v string) *DescribeAddressBookRequest
	GetQuery() *string
}

type DescribeAddressBookRequest struct {
	// Filters the query to return only address books that contain the specified port. This parameter is valid only when **GroupType*	- is set to **port**.
	//
	// example:
	//
	// 80
	ContainPort *string `json:"ContainPort,omitempty" xml:"ContainPort,omitempty"`
	// The page number for a paginated query.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the address book.
	//
	// > If this parameter is not specified, the query returns both IPv4 and ECS tag address books.
	//
	// example:
	//
	// ip
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The unique identifier of the address book.
	//
	// example:
	//
	// f04ac7ce-628b-4cb7-be61-310222b7****
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The language of the content in the response.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of address books per page.
	//
	// Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword for address books.
	//
	// example:
	//
	// 192.0.XX.XX
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s DescribeAddressBookRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookRequest) GetContainPort() *string {
	return s.ContainPort
}

func (s *DescribeAddressBookRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeAddressBookRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeAddressBookRequest) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *DescribeAddressBookRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAddressBookRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAddressBookRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeAddressBookRequest) SetContainPort(v string) *DescribeAddressBookRequest {
	s.ContainPort = &v
	return s
}

func (s *DescribeAddressBookRequest) SetCurrentPage(v string) *DescribeAddressBookRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAddressBookRequest) SetGroupType(v string) *DescribeAddressBookRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeAddressBookRequest) SetGroupUuid(v string) *DescribeAddressBookRequest {
	s.GroupUuid = &v
	return s
}

func (s *DescribeAddressBookRequest) SetLang(v string) *DescribeAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAddressBookRequest) SetPageSize(v string) *DescribeAddressBookRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAddressBookRequest) SetQuery(v string) *DescribeAddressBookRequest {
	s.Query = &v
	return s
}

func (s *DescribeAddressBookRequest) Validate() error {
	return dara.Validate(s)
}
