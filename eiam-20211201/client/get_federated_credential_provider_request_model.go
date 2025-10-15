// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFederatedCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFederatedCredentialProviderId(v string) *GetFederatedCredentialProviderRequest
	GetFederatedCredentialProviderId() *string
	SetInstanceId(v string) *GetFederatedCredentialProviderRequest
	GetInstanceId() *string
}

type GetFederatedCredentialProviderRequest struct {
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

func (s GetFederatedCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderRequest) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *GetFederatedCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetFederatedCredentialProviderRequest) SetFederatedCredentialProviderId(v string) *GetFederatedCredentialProviderRequest {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *GetFederatedCredentialProviderRequest) SetInstanceId(v string) *GetFederatedCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFederatedCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
