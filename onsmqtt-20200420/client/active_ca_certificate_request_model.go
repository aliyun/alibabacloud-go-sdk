// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMqttInstanceId(v string) *ActiveCaCertificateRequest
	GetMqttInstanceId() *string
	SetSn(v string) *ActiveCaCertificateRequest
	GetSn() *string
}

type ActiveCaCertificateRequest struct {
	// The ID of the instance to which the CA certificate is bound, which is the instance ID of ApsaraMQ for MQTT.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	// The serial number (SN) of the CA certificate to be activated. The SN uniquely identifies a CA certificate. The value can be up to 128 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 007269004887******
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s ActiveCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *ActiveCaCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *ActiveCaCertificateRequest) GetSn() *string {
	return s.Sn
}

func (s *ActiveCaCertificateRequest) SetMqttInstanceId(v string) *ActiveCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *ActiveCaCertificateRequest) SetSn(v string) *ActiveCaCertificateRequest {
	s.Sn = &v
	return s
}

func (s *ActiveCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
