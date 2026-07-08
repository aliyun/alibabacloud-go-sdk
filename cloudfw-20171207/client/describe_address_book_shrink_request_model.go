// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddressBookShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetMemberUidsShrink(v string) *DescribeAddressBookShrinkRequest
	GetAssetMemberUidsShrink() *string
	SetContainPort(v string) *DescribeAddressBookShrinkRequest
	GetContainPort() *string
	SetCurrentPage(v string) *DescribeAddressBookShrinkRequest
	GetCurrentPage() *string
	SetGroupType(v string) *DescribeAddressBookShrinkRequest
	GetGroupType() *string
	SetGroupUuid(v string) *DescribeAddressBookShrinkRequest
	GetGroupUuid() *string
	SetLang(v string) *DescribeAddressBookShrinkRequest
	GetLang() *string
	SetPageSize(v string) *DescribeAddressBookShrinkRequest
	GetPageSize() *string
	SetQuery(v string) *DescribeAddressBookShrinkRequest
	GetQuery() *string
}

type DescribeAddressBookShrinkRequest struct {
	// The list of member accounts for the asset address book.
	AssetMemberUidsShrink *string `json:"AssetMemberUids,omitempty" xml:"AssetMemberUids,omitempty"`
	// Queries address books that contain the specified port. This parameter takes effect only when the **GroupType*	- parameter is set to **port**.
	//
	// example:
	//
	// 80
	ContainPort *string `json:"ContainPort,omitempty" xml:"ContainPort,omitempty"`
	// The page number in a paged query.
	//
	// Default value: 1, which indicates that the first page of data is returned.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the address book.
	//
	// > If you do not set this parameter, IP address books and ECS tag-based address books are queried.
	//
	// example:
	//
	// ip
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The unique ID of the address book.
	//
	// example:
	//
	// f04ac7ce-628b-4cb7-be61-310222b7****
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The language type for the address book description. Valid values:
	//
	// - **en**: English.
	//
	// - **zh**: Chinese (default).
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of address books on each page in a paged query.
	//
	// Default value: 10, which indicates that each page contains 10 results. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search condition. Enter the address book information that you want to query.
	//
	// example:
	//
	// 192.0.XX.XX
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s DescribeAddressBookShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressBookShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookShrinkRequest) GetAssetMemberUidsShrink() *string {
	return s.AssetMemberUidsShrink
}

func (s *DescribeAddressBookShrinkRequest) GetContainPort() *string {
	return s.ContainPort
}

func (s *DescribeAddressBookShrinkRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeAddressBookShrinkRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeAddressBookShrinkRequest) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *DescribeAddressBookShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAddressBookShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAddressBookShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeAddressBookShrinkRequest) SetAssetMemberUidsShrink(v string) *DescribeAddressBookShrinkRequest {
	s.AssetMemberUidsShrink = &v
	return s
}

func (s *DescribeAddressBookShrinkRequest) SetContainPort(v string) *DescribeAddressBookShrinkRequest {
	s.ContainPort = &v
	return s
}

func (s *DescribeAddressBookShrinkRequest) SetCurrentPage(v string) *DescribeAddressBookShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAddressBookShrinkRequest) SetGroupType(v string) *DescribeAddressBookShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeAddressBookShrinkRequest) SetGroupUuid(v string) *DescribeAddressBookShrinkRequest {
	s.GroupUuid = &v
	return s
}

func (s *DescribeAddressBookShrinkRequest) SetLang(v string) *DescribeAddressBookShrinkRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAddressBookShrinkRequest) SetPageSize(v string) *DescribeAddressBookShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAddressBookShrinkRequest) SetQuery(v string) *DescribeAddressBookShrinkRequest {
	s.Query = &v
	return s
}

func (s *DescribeAddressBookShrinkRequest) Validate() error {
	return dara.Validate(s)
}
