// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateAppPolicyRequest
	GetDescription() *string
	SetOwnerId(v int64) *CreateAppPolicyRequest
	GetOwnerId() *int64
	SetPolicyName(v string) *CreateAppPolicyRequest
	GetPolicyName() *string
	SetPolicyValue(v string) *CreateAppPolicyRequest
	GetPolicyValue() *string
	SetResourceOwnerAccount(v string) *CreateAppPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAppPolicyRequest
	GetResourceOwnerId() *int64
}

type CreateAppPolicyRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// This parameter is required.
	PolicyValue          *string `json:"PolicyValue,omitempty" xml:"PolicyValue,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateAppPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAppPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAppPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateAppPolicyRequest) GetPolicyValue() *string {
	return s.PolicyValue
}

func (s *CreateAppPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAppPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAppPolicyRequest) SetDescription(v string) *CreateAppPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateAppPolicyRequest) SetOwnerId(v int64) *CreateAppPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAppPolicyRequest) SetPolicyName(v string) *CreateAppPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateAppPolicyRequest) SetPolicyValue(v string) *CreateAppPolicyRequest {
	s.PolicyValue = &v
	return s
}

func (s *CreateAppPolicyRequest) SetResourceOwnerAccount(v string) *CreateAppPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAppPolicyRequest) SetResourceOwnerId(v int64) *CreateAppPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAppPolicyRequest) Validate() error {
	return dara.Validate(s)
}
