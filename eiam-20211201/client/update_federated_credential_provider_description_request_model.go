// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFederatedCredentialProviderDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateFederatedCredentialProviderDescriptionRequest
	GetDescription() *string
	SetFederatedCredentialProviderId(v string) *UpdateFederatedCredentialProviderDescriptionRequest
	GetFederatedCredentialProviderId() *string
	SetInstanceId(v string) *UpdateFederatedCredentialProviderDescriptionRequest
	GetInstanceId() *string
}

type UpdateFederatedCredentialProviderDescriptionRequest struct {
	// 联邦凭证提供方描述
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 联邦凭证提供方ID
	//
	// This parameter is required.
	//
	// example:
	//
	// fcp_mkv7rgt4d7i4u7zqtzev2mxxxx
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateFederatedCredentialProviderDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFederatedCredentialProviderDescriptionRequest) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *UpdateFederatedCredentialProviderDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateFederatedCredentialProviderDescriptionRequest) SetDescription(v string) *UpdateFederatedCredentialProviderDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateFederatedCredentialProviderDescriptionRequest) SetFederatedCredentialProviderId(v string) *UpdateFederatedCredentialProviderDescriptionRequest {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *UpdateFederatedCredentialProviderDescriptionRequest) SetInstanceId(v string) *UpdateFederatedCredentialProviderDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateFederatedCredentialProviderDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
