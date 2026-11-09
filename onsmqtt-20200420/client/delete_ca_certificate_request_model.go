// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMqttInstanceId(v string) *DeleteCaCertificateRequest
	GetMqttInstanceId() *string
	SetSn(v string) *DeleteCaCertificateRequest
	GetSn() *string
}

type DeleteCaCertificateRequest struct {
	// The ID of the instance to which the CA certificate is bound, which is the instance ID of ApsaraMQ for MQTT.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	// The serial number (SN) of the CA certificate to delete. The SN uniquely identifies a CA certificate. The value cannot exceed 128 bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 007269004887******
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s DeleteCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCaCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *DeleteCaCertificateRequest) GetSn() *string {
	return s.Sn
}

func (s *DeleteCaCertificateRequest) SetMqttInstanceId(v string) *DeleteCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *DeleteCaCertificateRequest) SetSn(v string) *DeleteCaCertificateRequest {
	s.Sn = &v
	return s
}

func (s *DeleteCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
