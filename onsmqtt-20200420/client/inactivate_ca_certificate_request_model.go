// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInactivateCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMqttInstanceId(v string) *InactivateCaCertificateRequest
	GetMqttInstanceId() *string
	SetSn(v string) *InactivateCaCertificateRequest
	GetSn() *string
}

type InactivateCaCertificateRequest struct {
	// The ID of the ApsaraMQ for MQTT instance to which the CA certificate is attached.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	// The serial number (SN) of the CA certificate that you want to unregister. The SN uniquely identifies a CA certificate.
	//
	// The value can be up to 128 bytes long.
	//
	// This parameter is required.
	//
	// example:
	//
	// 007269004887******
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s InactivateCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s InactivateCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *InactivateCaCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *InactivateCaCertificateRequest) GetSn() *string {
	return s.Sn
}

func (s *InactivateCaCertificateRequest) SetMqttInstanceId(v string) *InactivateCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *InactivateCaCertificateRequest) SetSn(v string) *InactivateCaCertificateRequest {
	s.Sn = &v
	return s
}

func (s *InactivateCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
