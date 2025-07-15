// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntrys(v *DescribeAccessControlListAttributeResponseBodyAclEntrys) *DescribeAccessControlListAttributeResponseBody
	GetAclEntrys() *DescribeAccessControlListAttributeResponseBodyAclEntrys
	SetAclId(v string) *DescribeAccessControlListAttributeResponseBody
	GetAclId() *string
	SetAclName(v string) *DescribeAccessControlListAttributeResponseBody
	GetAclName() *string
	SetAddressIPVersion(v string) *DescribeAccessControlListAttributeResponseBody
	GetAddressIPVersion() *string
	SetRequestId(v string) *DescribeAccessControlListAttributeResponseBody
	GetRequestId() *string
}

type DescribeAccessControlListAttributeResponseBody struct {
	// The information about the access control policy.
	AclEntrys *DescribeAccessControlListAttributeResponseBodyAclEntrys `json:"AclEntrys,omitempty" xml:"AclEntrys,omitempty" type:"Struct"`
	// The ID of the access control policy.
	//
	// example:
	//
	// acl-uf6fpfdg3b5muska7uqem
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// testAcl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The IP protocol version. Valid values: **ipv4*	- and **ipv6**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccessControlListAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBody) GetAclEntrys() *DescribeAccessControlListAttributeResponseBodyAclEntrys {
	return s.AclEntrys
}

func (s *DescribeAccessControlListAttributeResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *DescribeAccessControlListAttributeResponseBody) GetAclName() *string {
	return s.AclName
}

func (s *DescribeAccessControlListAttributeResponseBody) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeAccessControlListAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAclEntrys(v *DescribeAccessControlListAttributeResponseBodyAclEntrys) *DescribeAccessControlListAttributeResponseBody {
	s.AclEntrys = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAclId(v string) *DescribeAccessControlListAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAclName(v string) *DescribeAccessControlListAttributeResponseBody {
	s.AclName = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetAddressIPVersion(v string) *DescribeAccessControlListAttributeResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) SetRequestId(v string) *DescribeAccessControlListAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessControlListAttributeResponseBodyAclEntrys struct {
	AclEntry []*DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry `json:"AclEntry,omitempty" xml:"AclEntry,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrys) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrys) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrys) GetAclEntry() []*DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry {
	return s.AclEntry
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrys) SetAclEntry(v []*DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) *DescribeAccessControlListAttributeResponseBodyAclEntrys {
	s.AclEntry = v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrys) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry struct {
	// The entry description.
	//
	// example:
	//
	// default
	AclEntryComment *string `json:"AclEntryComment,omitempty" xml:"AclEntryComment,omitempty"`
	// The ACL entry.
	//
	// example:
	//
	// 192.168.1.0/24
	AclEntryIp *string `json:"AclEntryIp,omitempty" xml:"AclEntryIp,omitempty"`
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) GetAclEntryComment() *string {
	return s.AclEntryComment
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) GetAclEntryIp() *string {
	return s.AclEntryIp
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) SetAclEntryComment(v string) *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry {
	s.AclEntryComment = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) SetAclEntryIp(v string) *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry {
	s.AclEntryIp = &v
	return s
}

func (s *DescribeAccessControlListAttributeResponseBodyAclEntrysAclEntry) Validate() error {
	return dara.Validate(s)
}
