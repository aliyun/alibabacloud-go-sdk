// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialProviderId(v string) *DeleteCredentialProviderRequest
	GetCredentialProviderId() *string
	SetInstanceId(v string) *DeleteCredentialProviderRequest
	GetInstanceId() *string
}

type DeleteCredentialProviderRequest struct {
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

func (s DeleteCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteCredentialProviderRequest) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *DeleteCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCredentialProviderRequest) SetCredentialProviderId(v string) *DeleteCredentialProviderRequest {
	s.CredentialProviderId = &v
	return s
}

func (s *DeleteCredentialProviderRequest) SetInstanceId(v string) *DeleteCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
