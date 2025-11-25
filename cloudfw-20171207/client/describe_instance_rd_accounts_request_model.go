// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRdAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeInstanceRdAccountsRequest
	GetCurrentPage() *string
	SetLang(v string) *DescribeInstanceRdAccountsRequest
	GetLang() *string
	SetMemberDesc(v string) *DescribeInstanceRdAccountsRequest
	GetMemberDesc() *string
	SetMemberDisplayName(v string) *DescribeInstanceRdAccountsRequest
	GetMemberDisplayName() *string
	SetMemberUid(v string) *DescribeInstanceRdAccountsRequest
	GetMemberUid() *string
	SetPageSize(v string) *DescribeInstanceRdAccountsRequest
	GetPageSize() *string
	SetSourceIp(v string) *DescribeInstanceRdAccountsRequest
	GetSourceIp() *string
}

type DescribeInstanceRdAccountsRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// renewal
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// example:
	//
	// cloudfirewall_2
	MemberDisplayName *string `json:"MemberDisplayName,omitempty" xml:"MemberDisplayName,omitempty"`
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 47.100.170.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInstanceRdAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRdAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRdAccountsRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeInstanceRdAccountsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInstanceRdAccountsRequest) GetMemberDesc() *string {
	return s.MemberDesc
}

func (s *DescribeInstanceRdAccountsRequest) GetMemberDisplayName() *string {
	return s.MemberDisplayName
}

func (s *DescribeInstanceRdAccountsRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeInstanceRdAccountsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInstanceRdAccountsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInstanceRdAccountsRequest) SetCurrentPage(v string) *DescribeInstanceRdAccountsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceRdAccountsRequest) SetLang(v string) *DescribeInstanceRdAccountsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceRdAccountsRequest) SetMemberDesc(v string) *DescribeInstanceRdAccountsRequest {
	s.MemberDesc = &v
	return s
}

func (s *DescribeInstanceRdAccountsRequest) SetMemberDisplayName(v string) *DescribeInstanceRdAccountsRequest {
	s.MemberDisplayName = &v
	return s
}

func (s *DescribeInstanceRdAccountsRequest) SetMemberUid(v string) *DescribeInstanceRdAccountsRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeInstanceRdAccountsRequest) SetPageSize(v string) *DescribeInstanceRdAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceRdAccountsRequest) SetSourceIp(v string) *DescribeInstanceRdAccountsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceRdAccountsRequest) Validate() error {
	return dara.Validate(s)
}
