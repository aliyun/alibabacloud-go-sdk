// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcls(v *DescribeAccessControlListsResponseBodyAcls) *DescribeAccessControlListsResponseBody
	GetAcls() *DescribeAccessControlListsResponseBodyAcls
	SetPageNumber(v int32) *DescribeAccessControlListsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessControlListsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccessControlListsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAccessControlListsResponseBody
	GetTotalCount() *int32
}

type DescribeAccessControlListsResponseBody struct {
	// The ACLs.
	Acls *DescribeAccessControlListsResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Struct"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessControlListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBody) GetAcls() *DescribeAccessControlListsResponseBodyAcls {
	return s.Acls
}

func (s *DescribeAccessControlListsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessControlListsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessControlListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessControlListsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccessControlListsResponseBody) SetAcls(v *DescribeAccessControlListsResponseBodyAcls) *DescribeAccessControlListsResponseBody {
	s.Acls = v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetPageNumber(v int32) *DescribeAccessControlListsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetPageSize(v int32) *DescribeAccessControlListsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetRequestId(v string) *DescribeAccessControlListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) SetTotalCount(v int32) *DescribeAccessControlListsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessControlListsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessControlListsResponseBodyAcls struct {
	Acl []*DescribeAccessControlListsResponseBodyAclsAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListsResponseBodyAcls) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBodyAcls) GetAcl() []*DescribeAccessControlListsResponseBodyAclsAcl {
	return s.Acl
}

func (s *DescribeAccessControlListsResponseBodyAcls) SetAcl(v []*DescribeAccessControlListsResponseBodyAclsAcl) *DescribeAccessControlListsResponseBodyAcls {
	s.Acl = v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAcls) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessControlListsResponseBodyAclsAcl struct {
	// The ID of the access control policy.
	//
	// example:
	//
	// acl-3nsohdozz0ru8fi5onwz1
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// testAcl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// 访问控制策略组的IP版本。
	//
	// - **IPv4**。
	//
	// - **IPv6**。
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
}

func (s DescribeAccessControlListsResponseBodyAclsAcl) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsResponseBodyAclsAcl) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) GetAclId() *string {
	return s.AclId
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) GetAclName() *string {
	return s.AclName
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetAclId(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetAclName(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.AclName = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) SetAddressIPVersion(v string) *DescribeAccessControlListsResponseBodyAclsAcl {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAccessControlListsResponseBodyAclsAcl) Validate() error {
	return dara.Validate(s)
}
