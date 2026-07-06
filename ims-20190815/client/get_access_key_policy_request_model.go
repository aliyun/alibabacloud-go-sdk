// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserAccessKeyId(v string) *GetAccessKeyPolicyRequest
	GetUserAccessKeyId() *string
	SetUserPrincipalName(v string) *GetAccessKeyPolicyRequest
	GetUserPrincipalName() *string
}

type GetAccessKeyPolicyRequest struct {
	// The access key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// If this parameter is left empty, the network access restriction policy of the specified access key for the current user is returned by default.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetAccessKeyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyPolicyRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetAccessKeyPolicyRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *GetAccessKeyPolicyRequest) SetUserAccessKeyId(v string) *GetAccessKeyPolicyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetAccessKeyPolicyRequest) SetUserPrincipalName(v string) *GetAccessKeyPolicyRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetAccessKeyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
