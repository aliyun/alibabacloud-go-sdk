// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteAppPolicyRequest
	GetOwnerId() *int64
	SetPolicyNames(v string) *DeleteAppPolicyRequest
	GetPolicyNames() *string
	SetResourceOwnerAccount(v string) *DeleteAppPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAppPolicyRequest
	GetResourceOwnerId() *int64
}

type DeleteAppPolicyRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PolicyNames          *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteAppPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAppPolicyRequest) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *DeleteAppPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAppPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAppPolicyRequest) SetOwnerId(v int64) *DeleteAppPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAppPolicyRequest) SetPolicyNames(v string) *DeleteAppPolicyRequest {
	s.PolicyNames = &v
	return s
}

func (s *DeleteAppPolicyRequest) SetResourceOwnerAccount(v string) *DeleteAppPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAppPolicyRequest) SetResourceOwnerId(v int64) *DeleteAppPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAppPolicyRequest) Validate() error {
	return dara.Validate(s)
}
