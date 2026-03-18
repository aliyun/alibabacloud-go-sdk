// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialProviderId(v string) *GetCredentialProviderRequest
	GetCredentialProviderId() *string
	SetInstanceId(v string) *GetCredentialProviderRequest
	GetInstanceId() *string
}

type GetCredentialProviderRequest struct {
	// 认证令牌提供商ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"CredentialProviderId,omitempty" xml:"CredentialProviderId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *GetCredentialProviderRequest) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *GetCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCredentialProviderRequest) SetCredentialProviderId(v string) *GetCredentialProviderRequest {
	s.CredentialProviderId = &v
	return s
}

func (s *GetCredentialProviderRequest) SetInstanceId(v string) *GetCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
