// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMqttInstanceId(v string) *GetCaCertificateRequest
	GetMqttInstanceId() *string
	SetSn(v string) *GetCaCertificateRequest
	GetSn() *string
}

type GetCaCertificateRequest struct {
	// The instance ID bound to the CA certificate, which is the instance ID of the MQTT version of the cloud message queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	// The SN serial number of the CA certificate to be queried, used to uniquely identify a CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 007269004887******
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s GetCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetCaCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *GetCaCertificateRequest) GetSn() *string {
	return s.Sn
}

func (s *GetCaCertificateRequest) SetMqttInstanceId(v string) *GetCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *GetCaCertificateRequest) SetSn(v string) *GetCaCertificateRequest {
	s.Sn = &v
	return s
}

func (s *GetCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
