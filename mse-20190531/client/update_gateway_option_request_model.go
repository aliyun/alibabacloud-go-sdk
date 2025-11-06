// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayOptionRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayOptionRequest
	GetGatewayId() *int64
	SetGatewayOption(v *GatewayOption) *UpdateGatewayOptionRequest
	GetGatewayOption() *GatewayOption
	SetGatewayUniqueId(v string) *UpdateGatewayOptionRequest
	GetGatewayUniqueId() *string
}

type UpdateGatewayOptionRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 421
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The detailed configurations of the gateway.
	//
	// 	- **TraceDetails**: the sampling description of Managed Service for OpenTelemetry. Content: TraceEnabled indicates whether Managed Service for OpenTelemetry is activated. Sample indicates the sampling rate of Managed Service for OpenTelemetry.
	//
	// 	- **LogConfigDetails**: the description of Simple Log Service. Content: LogEnabled indicates whether Simple Log Service is activated. ProjectName indicates the Simple Log Service project to which logs are delivered. LogStoreName indicates the name of the Logstore.
	//
	// 	- **EnableHardwareAcceleration**: indicates whether hardware acceleration is enabled.
	//
	// 	- **DisableHttp2Alpn**: indicates whether the HTTP/2 protocol is disabled.
	//
	// 	- **EnableWaf**: indicates whether Web Application Firewall (WAF) is enabled.
	GatewayOption *GatewayOption `json:"GatewayOption,omitempty" xml:"GatewayOption,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-83b0ddb569434f82b9fe8e4c60c4****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s UpdateGatewayOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayOptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayOptionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayOptionRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayOptionRequest) GetGatewayOption() *GatewayOption {
	return s.GatewayOption
}

func (s *UpdateGatewayOptionRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayOptionRequest) SetAcceptLanguage(v string) *UpdateGatewayOptionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayOptionRequest) SetGatewayId(v int64) *UpdateGatewayOptionRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayOptionRequest) SetGatewayOption(v *GatewayOption) *UpdateGatewayOptionRequest {
	s.GatewayOption = v
	return s
}

func (s *UpdateGatewayOptionRequest) SetGatewayUniqueId(v string) *UpdateGatewayOptionRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayOptionRequest) Validate() error {
	if s.GatewayOption != nil {
		if err := s.GatewayOption.Validate(); err != nil {
			return err
		}
	}
	return nil
}
