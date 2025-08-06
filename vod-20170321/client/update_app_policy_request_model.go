// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateAppPolicyRequest
	GetDescription() *string
	SetOwnerId(v int64) *UpdateAppPolicyRequest
	GetOwnerId() *int64
	SetPolicyName(v string) *UpdateAppPolicyRequest
	GetPolicyName() *string
	SetPolicyValue(v string) *UpdateAppPolicyRequest
	GetPolicyValue() *string
	SetResourceOwnerAccount(v string) *UpdateAppPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAppPolicyRequest
	GetResourceOwnerId() *int64
}

type UpdateAppPolicyRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PolicyName           *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyValue          *string `json:"PolicyValue,omitempty" xml:"PolicyValue,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateAppPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAppPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAppPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdateAppPolicyRequest) GetPolicyValue() *string {
	return s.PolicyValue
}

func (s *UpdateAppPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAppPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAppPolicyRequest) SetDescription(v string) *UpdateAppPolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppPolicyRequest) SetOwnerId(v int64) *UpdateAppPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAppPolicyRequest) SetPolicyName(v string) *UpdateAppPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *UpdateAppPolicyRequest) SetPolicyValue(v string) *UpdateAppPolicyRequest {
	s.PolicyValue = &v
	return s
}

func (s *UpdateAppPolicyRequest) SetResourceOwnerAccount(v string) *UpdateAppPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAppPolicyRequest) SetResourceOwnerId(v int64) *UpdateAppPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAppPolicyRequest) Validate() error {
	return dara.Validate(s)
}
