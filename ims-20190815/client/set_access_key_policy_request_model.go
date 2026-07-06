// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessKeyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyPolicy(v string) *SetAccessKeyPolicyRequest
	GetAccessKeyPolicy() *string
	SetUserAccessKeyId(v string) *SetAccessKeyPolicyRequest
	GetUserAccessKeyId() *string
	SetUserPrincipalName(v string) *SetAccessKeyPolicyRequest
	GetUserPrincipalName() *string
}

type SetAccessKeyPolicyRequest struct {
	// The network access restriction policy.
	//
	// A JSON-formatted string. For more information, see the AccessKeyPolicy structure description.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Status":"Inactive","Statements":[{"Value":"AllowAllVPC","Type":"VPCWhiteList","IPList":["::/0","0.0.0.0/0"]}]}
	AccessKeyPolicy *string `json:"AccessKeyPolicy,omitempty" xml:"AccessKeyPolicy,omitempty"`
	// The AccessKey ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// If this parameter is left empty, the network access restriction policy is set for the specified AccessKey pair of the current user by default.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s SetAccessKeyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAccessKeyPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetAccessKeyPolicyRequest) GetAccessKeyPolicy() *string {
	return s.AccessKeyPolicy
}

func (s *SetAccessKeyPolicyRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *SetAccessKeyPolicyRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *SetAccessKeyPolicyRequest) SetAccessKeyPolicy(v string) *SetAccessKeyPolicyRequest {
	s.AccessKeyPolicy = &v
	return s
}

func (s *SetAccessKeyPolicyRequest) SetUserAccessKeyId(v string) *SetAccessKeyPolicyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *SetAccessKeyPolicyRequest) SetUserPrincipalName(v string) *SetAccessKeyPolicyRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *SetAccessKeyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
