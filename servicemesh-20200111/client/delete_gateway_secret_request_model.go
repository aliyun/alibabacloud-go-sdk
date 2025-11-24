// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewaySecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIstioGatewayName(v string) *DeleteGatewaySecretRequest
	GetIstioGatewayName() *string
	SetSecretName(v string) *DeleteGatewaySecretRequest
	GetSecretName() *string
	SetServiceMeshId(v string) *DeleteGatewaySecretRequest
	GetServiceMeshId() *string
}

type DeleteGatewaySecretRequest struct {
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// bookinfo-secret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteGatewaySecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewaySecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecretRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *DeleteGatewaySecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *DeleteGatewaySecretRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DeleteGatewaySecretRequest) SetIstioGatewayName(v string) *DeleteGatewaySecretRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DeleteGatewaySecretRequest) SetSecretName(v string) *DeleteGatewaySecretRequest {
	s.SecretName = &v
	return s
}

func (s *DeleteGatewaySecretRequest) SetServiceMeshId(v string) *DeleteGatewaySecretRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DeleteGatewaySecretRequest) Validate() error {
	return dara.Validate(s)
}
