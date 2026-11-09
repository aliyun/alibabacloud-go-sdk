// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaContent(v string) *RegisterCaCertificateRequest
	GetCaContent() *string
	SetCaName(v string) *RegisterCaCertificateRequest
	GetCaName() *string
	SetMqttInstanceId(v string) *RegisterCaCertificateRequest
	GetMqttInstanceId() *string
	SetVerificationContent(v string) *RegisterCaCertificateRequest
	GetVerificationContent() *string
}

type RegisterCaCertificateRequest struct {
	// The content of the CA certificate that you want to register with an ApsaraMQ for MQTT broker.
	//
	// > In the example,
	//
	//  indicates a line feed.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIIDuzCCAqdGVzdC5jbi1xaW5n******\\n-----END CERTIFICATE-----
	CaContent *string `json:"CaContent,omitempty" xml:"CaContent,omitempty"`
	// The name of the CA certificate that you want to register with an ApsaraMQ for MQTT broker.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt_ca
	CaName *string `json:"CaName,omitempty" xml:"CaName,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance to which you want to bind the CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	// The content of the validation certificate issued by the CA certificate that you want to register with an ApsaraMQ for MQTT broker. The validation certificate must be used together with the registration code of the CA certificate to verify the private key of the CA certificate.
	//
	// > In the example,
	//
	//  indicates a line feed.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIID/DCCAu+Y5sRMpp9tnd+4s******\\n-----END CERTIFICATE-----
	VerificationContent *string `json:"VerificationContent,omitempty" xml:"VerificationContent,omitempty"`
}

func (s RegisterCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *RegisterCaCertificateRequest) GetCaContent() *string {
	return s.CaContent
}

func (s *RegisterCaCertificateRequest) GetCaName() *string {
	return s.CaName
}

func (s *RegisterCaCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *RegisterCaCertificateRequest) GetVerificationContent() *string {
	return s.VerificationContent
}

func (s *RegisterCaCertificateRequest) SetCaContent(v string) *RegisterCaCertificateRequest {
	s.CaContent = &v
	return s
}

func (s *RegisterCaCertificateRequest) SetCaName(v string) *RegisterCaCertificateRequest {
	s.CaName = &v
	return s
}

func (s *RegisterCaCertificateRequest) SetMqttInstanceId(v string) *RegisterCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *RegisterCaCertificateRequest) SetVerificationContent(v string) *RegisterCaCertificateRequest {
	s.VerificationContent = &v
	return s
}

func (s *RegisterCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
