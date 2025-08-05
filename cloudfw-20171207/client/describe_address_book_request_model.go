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
	SetLang(v string) *DescribeAddressBookRequest
	GetLang() *string
	SetPageSize(v string) *DescribeAddressBookRequest
	GetPageSize() *string
	SetQuery(v string) *DescribeAddressBookRequest
	GetQuery() *string
}

type DescribeAddressBookRequest struct {
	// The port that is included in the address book. This parameter takes effect only when the **GroupType*	- parameter is set to **port**.
	//
	// example:
	//
	// 80
	ContainPort *string `json:"ContainPort,omitempty" xml:"ContainPort,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the address book. Valid values:
	//
	// 	- **ip**: IP address book
	//
	// 	- **domain**: domain address book
	//
	// 	- **port**: port address book
	//
	// 	- **tag**: Elastic Compute Service (ECS) tag-based address book
	//
	// 	- **allCloud**: cloud service address book
	//
	// 	- **threat**: threat intelligence address book
	//
	// 	- **ipv6**: IPv6 address book
	//
	// >  If you do not specify a type, the domain address books and ECS tag-based address books are queried.
	//
	// example:
	//
	// ip
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The language of the content within the request. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query condition that is used to search for the address book.
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
