// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFederatedCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFederatedCredentialProviderId(v string) *DisableFederatedCredentialProviderRequest
	GetFederatedCredentialProviderId() *string
	SetInstanceId(v string) *DisableFederatedCredentialProviderRequest
	GetInstanceId() *string
}

type DisableFederatedCredentialProviderRequest struct {
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

func (s DisableFederatedCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableFederatedCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *DisableFederatedCredentialProviderRequest) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *DisableFederatedCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableFederatedCredentialProviderRequest) SetFederatedCredentialProviderId(v string) *DisableFederatedCredentialProviderRequest {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *DisableFederatedCredentialProviderRequest) SetInstanceId(v string) *DisableFederatedCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableFederatedCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
