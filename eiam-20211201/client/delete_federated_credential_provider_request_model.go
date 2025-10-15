// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFederatedCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFederatedCredentialProviderId(v string) *DeleteFederatedCredentialProviderRequest
	GetFederatedCredentialProviderId() *string
	SetInstanceId(v string) *DeleteFederatedCredentialProviderRequest
	GetInstanceId() *string
}

type DeleteFederatedCredentialProviderRequest struct {
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

func (s DeleteFederatedCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFederatedCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteFederatedCredentialProviderRequest) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *DeleteFederatedCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteFederatedCredentialProviderRequest) SetFederatedCredentialProviderId(v string) *DeleteFederatedCredentialProviderRequest {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *DeleteFederatedCredentialProviderRequest) SetInstanceId(v string) *DeleteFederatedCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFederatedCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
