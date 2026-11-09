// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveDeviceCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaSn(v string) *ActiveDeviceCertificateRequest
	GetCaSn() *string
	SetDeviceSn(v string) *ActiveDeviceCertificateRequest
	GetDeviceSn() *string
	SetMqttInstanceId(v string) *ActiveDeviceCertificateRequest
	GetMqttInstanceId() *string
}

type ActiveDeviceCertificateRequest struct {
	// The SN of the CA certificate to which the device certificate to be activated belongs. The SN uniquely identifies a CA certificate.
	//
	// Valid values: up to 128 bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 007269004887******
	CaSn *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	// The serial number (SN) of the device certificate to be activated. The SN uniquely identifies a device certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 356217374433******
	DeviceSn *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	// The instance ID to which the device certificate is attached, which is the instance ID of MSMQ for MQTT.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s ActiveDeviceCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *ActiveDeviceCertificateRequest) GetCaSn() *string {
	return s.CaSn
}

func (s *ActiveDeviceCertificateRequest) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *ActiveDeviceCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *ActiveDeviceCertificateRequest) SetCaSn(v string) *ActiveDeviceCertificateRequest {
	s.CaSn = &v
	return s
}

func (s *ActiveDeviceCertificateRequest) SetDeviceSn(v string) *ActiveDeviceCertificateRequest {
	s.DeviceSn = &v
	return s
}

func (s *ActiveDeviceCertificateRequest) SetMqttInstanceId(v string) *ActiveDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *ActiveDeviceCertificateRequest) Validate() error {
	return dara.Validate(s)
}
