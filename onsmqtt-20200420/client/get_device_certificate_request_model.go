// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaSn(v string) *GetDeviceCertificateRequest
	GetCaSn() *string
	SetDeviceSn(v string) *GetDeviceCertificateRequest
	GetDeviceSn() *string
	SetMqttInstanceId(v string) *GetDeviceCertificateRequest
	GetMqttInstanceId() *string
}

type GetDeviceCertificateRequest struct {
	// The SN serial number of the CA certificate to which the device certificate to be queried belongs, used to uniquely identify a CA certificate. Value range: no more than 128 bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 007269004887******
	CaSn *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	// The SN serial number of the device certificate to be queried, used to uniquely identify a device certificate. Value range: no more than 128 bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 356217374433******
	DeviceSn *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	// The instance ID to which the device certificate is bound, i.e., the instance ID of the Cloud Message Queue MQTT version.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s GetDeviceCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceCertificateRequest) GetCaSn() *string {
	return s.CaSn
}

func (s *GetDeviceCertificateRequest) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *GetDeviceCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *GetDeviceCertificateRequest) SetCaSn(v string) *GetDeviceCertificateRequest {
	s.CaSn = &v
	return s
}

func (s *GetDeviceCertificateRequest) SetDeviceSn(v string) *GetDeviceCertificateRequest {
	s.DeviceSn = &v
	return s
}

func (s *GetDeviceCertificateRequest) SetMqttInstanceId(v string) *GetDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *GetDeviceCertificateRequest) Validate() error {
	return dara.Validate(s)
}
