// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInactivateDeviceCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaSn(v string) *InactivateDeviceCertificateRequest
	GetCaSn() *string
	SetDeviceSn(v string) *InactivateDeviceCertificateRequest
	GetDeviceSn() *string
	SetMqttInstanceId(v string) *InactivateDeviceCertificateRequest
	GetMqttInstanceId() *string
}

type InactivateDeviceCertificateRequest struct {
	// The SN of the CA certificate to which the device certificate belongs. The SN uniquely identifies a CA certificate.
	//
	// The value can be up to 128 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 007269004887******
	CaSn *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	// The serial number (SN) of the device certificate that you want to inactivate. The SN uniquely identifies a device certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 356217374433******
	DeviceSn *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance to which the device certificate is attached.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s InactivateDeviceCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s InactivateDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *InactivateDeviceCertificateRequest) GetCaSn() *string {
	return s.CaSn
}

func (s *InactivateDeviceCertificateRequest) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *InactivateDeviceCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *InactivateDeviceCertificateRequest) SetCaSn(v string) *InactivateDeviceCertificateRequest {
	s.CaSn = &v
	return s
}

func (s *InactivateDeviceCertificateRequest) SetDeviceSn(v string) *InactivateDeviceCertificateRequest {
	s.DeviceSn = &v
	return s
}

func (s *InactivateDeviceCertificateRequest) SetMqttInstanceId(v string) *InactivateDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *InactivateDeviceCertificateRequest) Validate() error {
	return dara.Validate(s)
}
