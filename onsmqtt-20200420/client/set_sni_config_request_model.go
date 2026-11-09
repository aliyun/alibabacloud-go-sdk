// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSniConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultCertificate(v string) *SetSniConfigRequest
	GetDefaultCertificate() *string
	SetMqttInstanceId(v string) *SetSniConfigRequest
	GetMqttInstanceId() *string
	SetSniConfig(v string) *SetSniConfigRequest
	GetSniConfig() *string
}

type SetSniConfigRequest struct {
	// The default certificate. This certificate is returned if no matching domain name is found in the server certificate list.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1533xxxx-cn-hangzhou
	DefaultCertificate *string `json:"DefaultCertificate,omitempty" xml:"DefaultCertificate,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-5yd3xxx
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	// The Server Name Indication (SNI) configuration that maps domain names to certificates. Format: DomainName#CertificateID#Password(if any);DomainName#CertificateID#Password(if any)
	//
	// example:
	//
	// *.mqtt.aliyuncs.com#15xxxxx-cn-hangzhou;mqtt-test.mqtt.aliyuncs.com#15xxxx9-cn-hangzhou
	SniConfig *string `json:"SniConfig,omitempty" xml:"SniConfig,omitempty"`
}

func (s SetSniConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSniConfigRequest) GoString() string {
	return s.String()
}

func (s *SetSniConfigRequest) GetDefaultCertificate() *string {
	return s.DefaultCertificate
}

func (s *SetSniConfigRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *SetSniConfigRequest) GetSniConfig() *string {
	return s.SniConfig
}

func (s *SetSniConfigRequest) SetDefaultCertificate(v string) *SetSniConfigRequest {
	s.DefaultCertificate = &v
	return s
}

func (s *SetSniConfigRequest) SetMqttInstanceId(v string) *SetSniConfigRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *SetSniConfigRequest) SetSniConfig(v string) *SetSniConfigRequest {
	s.SniConfig = &v
	return s
}

func (s *SetSniConfigRequest) Validate() error {
	return dara.Validate(s)
}
