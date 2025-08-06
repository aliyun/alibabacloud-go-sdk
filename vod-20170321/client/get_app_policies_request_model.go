// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetAppPoliciesRequest
	GetOwnerId() *int64
	SetPolicyNames(v string) *GetAppPoliciesRequest
	GetPolicyNames() *string
	SetResourceOwnerAccount(v string) *GetAppPoliciesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetAppPoliciesRequest
	GetResourceOwnerId() *int64
}

type GetAppPoliciesRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PolicyNames          *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAppPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppPoliciesRequest) GoString() string {
	return s.String()
}

func (s *GetAppPoliciesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetAppPoliciesRequest) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *GetAppPoliciesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAppPoliciesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAppPoliciesRequest) SetOwnerId(v int64) *GetAppPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAppPoliciesRequest) SetPolicyNames(v string) *GetAppPoliciesRequest {
	s.PolicyNames = &v
	return s
}

func (s *GetAppPoliciesRequest) SetResourceOwnerAccount(v string) *GetAppPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAppPoliciesRequest) SetResourceOwnerId(v int64) *GetAppPoliciesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAppPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
