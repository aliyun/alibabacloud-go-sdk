// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialProviderId(v string) *DisableCredentialProviderRequest
	GetCredentialProviderId() *string
	SetInstanceId(v string) *DisableCredentialProviderRequest
	GetInstanceId() *string
}

type DisableCredentialProviderRequest struct {
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

func (s DisableCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *DisableCredentialProviderRequest) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *DisableCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableCredentialProviderRequest) SetCredentialProviderId(v string) *DisableCredentialProviderRequest {
	s.CredentialProviderId = &v
	return s
}

func (s *DisableCredentialProviderRequest) SetInstanceId(v string) *DisableCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
