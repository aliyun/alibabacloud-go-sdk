// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewaySecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *CreateGatewaySecretRequest
	GetCert() *string
	SetIstioGatewayName(v string) *CreateGatewaySecretRequest
	GetIstioGatewayName() *string
	SetKey(v string) *CreateGatewaySecretRequest
	GetKey() *string
	SetSecretName(v string) *CreateGatewaySecretRequest
	GetSecretName() *string
	SetServiceMeshId(v string) *CreateGatewaySecretRequest
	GetServiceMeshId() *string
}

type CreateGatewaySecretRequest struct {
	// The content of the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIC2DCCAcACA-----END CERTIF****-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The private key of the certificate.
	//
	// example:
	//
	// MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC2ag/Bzcgm****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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

func (s CreateGatewaySecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewaySecretRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewaySecretRequest) GetCert() *string {
	return s.Cert
}

func (s *CreateGatewaySecretRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *CreateGatewaySecretRequest) GetKey() *string {
	return s.Key
}

func (s *CreateGatewaySecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *CreateGatewaySecretRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *CreateGatewaySecretRequest) SetCert(v string) *CreateGatewaySecretRequest {
	s.Cert = &v
	return s
}

func (s *CreateGatewaySecretRequest) SetIstioGatewayName(v string) *CreateGatewaySecretRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateGatewaySecretRequest) SetKey(v string) *CreateGatewaySecretRequest {
	s.Key = &v
	return s
}

func (s *CreateGatewaySecretRequest) SetSecretName(v string) *CreateGatewaySecretRequest {
	s.SecretName = &v
	return s
}

func (s *CreateGatewaySecretRequest) SetServiceMeshId(v string) *CreateGatewaySecretRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateGatewaySecretRequest) Validate() error {
	return dara.Validate(s)
}
