// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeviceCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaSn(v string) *DeleteDeviceCertificateRequest
	GetCaSn() *string
	SetDeviceSn(v string) *DeleteDeviceCertificateRequest
	GetDeviceSn() *string
	SetMqttInstanceId(v string) *DeleteDeviceCertificateRequest
	GetMqttInstanceId() *string
}

type DeleteDeviceCertificateRequest struct {
	// The serial number of the CA certificate to which the device certificate belongs. The serial number uniquely identifies a CA certificate. The CA certificate is used to verify the validity of the device certificate. The value can be up to 128 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 007269004887******
	CaSn *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	// The serial number (SN) of the device certificate to delete. The SN uniquely identifies a device certificate. The value can be up to 128 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 356217374433****
	DeviceSn *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	// The ID of the instance to which the device certificate is bound, which is the instance ID of ApsaraMQ for MQTT.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s DeleteDeviceCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCertificateRequest) GetCaSn() *string {
	return s.CaSn
}

func (s *DeleteDeviceCertificateRequest) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *DeleteDeviceCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *DeleteDeviceCertificateRequest) SetCaSn(v string) *DeleteDeviceCertificateRequest {
	s.CaSn = &v
	return s
}

func (s *DeleteDeviceCertificateRequest) SetDeviceSn(v string) *DeleteDeviceCertificateRequest {
	s.DeviceSn = &v
	return s
}

func (s *DeleteDeviceCertificateRequest) SetMqttInstanceId(v string) *DeleteDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *DeleteDeviceCertificateRequest) Validate() error {
	return dara.Validate(s)
}
