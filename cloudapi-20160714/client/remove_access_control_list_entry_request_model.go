// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAccessControlListEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntrys(v string) *RemoveAccessControlListEntryRequest
	GetAclEntrys() *string
	SetAclId(v string) *RemoveAccessControlListEntryRequest
	GetAclId() *string
	SetSecurityToken(v string) *RemoveAccessControlListEntryRequest
	GetSecurityToken() *string
}

type RemoveAccessControlListEntryRequest struct {
	// example:
	//
	// [{\\"entry\\":\\"192.168.1.0/24\\",\\"comment\\":\\"WhiteIp\\"}]
	AclEntrys *string `json:"AclEntrys,omitempty" xml:"AclEntrys,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acl-bp12ag0xxcfhq1ll68wp9
	AclId         *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveAccessControlListEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAccessControlListEntryRequest) GoString() string {
	return s.String()
}

func (s *RemoveAccessControlListEntryRequest) GetAclEntrys() *string {
	return s.AclEntrys
}

func (s *RemoveAccessControlListEntryRequest) GetAclId() *string {
	return s.AclId
}

func (s *RemoveAccessControlListEntryRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveAccessControlListEntryRequest) SetAclEntrys(v string) *RemoveAccessControlListEntryRequest {
	s.AclEntrys = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetAclId(v string) *RemoveAccessControlListEntryRequest {
	s.AclId = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) SetSecurityToken(v string) *RemoveAccessControlListEntryRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveAccessControlListEntryRequest) Validate() error {
	return dara.Validate(s)
}
