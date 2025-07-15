// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorAclListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6TranslatorAcls(v *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls) *DescribeIPv6TranslatorAclListsResponseBody
	GetIpv6TranslatorAcls() *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls
	SetPageNumber(v int32) *DescribeIPv6TranslatorAclListsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIPv6TranslatorAclListsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIPv6TranslatorAclListsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIPv6TranslatorAclListsResponseBody
	GetTotalCount() *int32
}

type DescribeIPv6TranslatorAclListsResponseBody struct {
	// The list of network ACLs.
	Ipv6TranslatorAcls *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls `json:"Ipv6TranslatorAcls,omitempty" xml:"Ipv6TranslatorAcls,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIPv6TranslatorAclListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) GetIpv6TranslatorAcls() *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls {
	return s.Ipv6TranslatorAcls
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) SetIpv6TranslatorAcls(v *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls) *DescribeIPv6TranslatorAclListsResponseBody {
	s.Ipv6TranslatorAcls = v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) SetPageNumber(v int32) *DescribeIPv6TranslatorAclListsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) SetPageSize(v int32) *DescribeIPv6TranslatorAclListsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) SetRequestId(v string) *DescribeIPv6TranslatorAclListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) SetTotalCount(v int32) *DescribeIPv6TranslatorAclListsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls struct {
	IPv6TranslatorAcl []*DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl `json:"IPv6TranslatorAcl,omitempty" xml:"IPv6TranslatorAcl,omitempty" type:"Repeated"`
}

func (s DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls) GetIPv6TranslatorAcl() []*DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl {
	return s.IPv6TranslatorAcl
}

func (s *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls) SetIPv6TranslatorAcl(v []*DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl) *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls {
	s.IPv6TranslatorAcl = v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAcls) Validate() error {
	return dara.Validate(s)
}

type DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl struct {
	// The ACL ID.
	//
	// example:
	//
	// ipv6transacl-bp1de2****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ACL name.
	//
	// example:
	//
	// acl1
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
}

func (s DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl) GetAclId() *string {
	return s.AclId
}

func (s *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl) GetAclName() *string {
	return s.AclName
}

func (s *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl) SetAclId(v string) *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl {
	s.AclId = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl) SetAclName(v string) *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl {
	s.AclName = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponseBodyIpv6TranslatorAclsIPv6TranslatorAcl) Validate() error {
	return dara.Validate(s)
}
