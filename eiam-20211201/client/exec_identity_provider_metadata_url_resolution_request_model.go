// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecIdentityProviderMetadataUrlResolutionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetIdentityProviderId(v string) *ExecIdentityProviderMetadataUrlResolutionRequest
  GetIdentityProviderId() *string 
  SetInstanceId(v string) *ExecIdentityProviderMetadataUrlResolutionRequest
  GetInstanceId() *string 
  SetNetworkAccessEndpointId(v string) *ExecIdentityProviderMetadataUrlResolutionRequest
  GetNetworkAccessEndpointId() *string 
  SetOidcIssuer(v string) *ExecIdentityProviderMetadataUrlResolutionRequest
  GetOidcIssuer() *string 
}

type ExecIdentityProviderMetadataUrlResolutionRequest struct {
  // example:
  // 
  // idp_mwpcwnhrimlr2horxXXXX
  IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // inae_public
  NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
  // OIDC Issuer地址。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // https://login.example.com/.well-known/openid-configuration
  OidcIssuer *string `json:"OidcIssuer,omitempty" xml:"OidcIssuer,omitempty"`
}

func (s ExecIdentityProviderMetadataUrlResolutionRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecIdentityProviderMetadataUrlResolutionRequest) GoString() string {
  return s.String()
}

func (s *ExecIdentityProviderMetadataUrlResolutionRequest) GetIdentityProviderId() *string  {
  return s.IdentityProviderId
}

func (s *ExecIdentityProviderMetadataUrlResolutionRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExecIdentityProviderMetadataUrlResolutionRequest) GetNetworkAccessEndpointId() *string  {
  return s.NetworkAccessEndpointId
}

func (s *ExecIdentityProviderMetadataUrlResolutionRequest) GetOidcIssuer() *string  {
  return s.OidcIssuer
}

func (s *ExecIdentityProviderMetadataUrlResolutionRequest) SetIdentityProviderId(v string) *ExecIdentityProviderMetadataUrlResolutionRequest {
  s.IdentityProviderId = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionRequest) SetInstanceId(v string) *ExecIdentityProviderMetadataUrlResolutionRequest {
  s.InstanceId = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionRequest) SetNetworkAccessEndpointId(v string) *ExecIdentityProviderMetadataUrlResolutionRequest {
  s.NetworkAccessEndpointId = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionRequest) SetOidcIssuer(v string) *ExecIdentityProviderMetadataUrlResolutionRequest {
  s.OidcIssuer = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionRequest) Validate() error {
  return dara.Validate(s)
}

