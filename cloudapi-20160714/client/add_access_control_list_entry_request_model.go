// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAccessControlListEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntrys(v string) *AddAccessControlListEntryRequest
	GetAclEntrys() *string
	SetAclId(v string) *AddAccessControlListEntryRequest
	GetAclId() *string
	SetSecurityToken(v string) *AddAccessControlListEntryRequest
	GetSecurityToken() *string
}

type AddAccessControlListEntryRequest struct {
	// The ACL settings.
	//
	// 	- entry: the entries that you want to add to the ACL. You can add CIDR blocks. Separate multiple CIDR blocks with commas (,).
	//
	// 	- comment: the description of the ACL.
	//
	// > You can add at most 50 IP addresses or CIDR blocks to an ACL in each call. If the IP address or CIDR block that you want to add to an ACL already exists, the IP address or CIDR block is not added. The entries that you add must be CIDR blocks.
	//
	// example:
	//
	// [{\\"entry\\": \\"192.168.1.0/24\\", \\"comment\\": \\"test\\"}]
	AclEntrys *string `json:"AclEntrys,omitempty" xml:"AclEntrys,omitempty"`
	// The ID of the access control list (ACL).
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-bp1ohqkonqybecf4llbrc
	AclId         *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AddAccessControlListEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAccessControlListEntryRequest) GoString() string {
	return s.String()
}

func (s *AddAccessControlListEntryRequest) GetAclEntrys() *string {
	return s.AclEntrys
}

func (s *AddAccessControlListEntryRequest) GetAclId() *string {
	return s.AclId
}

func (s *AddAccessControlListEntryRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddAccessControlListEntryRequest) SetAclEntrys(v string) *AddAccessControlListEntryRequest {
	s.AclEntrys = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetAclId(v string) *AddAccessControlListEntryRequest {
	s.AclId = &v
	return s
}

func (s *AddAccessControlListEntryRequest) SetSecurityToken(v string) *AddAccessControlListEntryRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddAccessControlListEntryRequest) Validate() error {
	return dara.Validate(s)
}
