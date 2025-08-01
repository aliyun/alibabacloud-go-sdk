// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayOptionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayOptionShrinkRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayOptionShrinkRequest
	GetGatewayId() *int64
	SetGatewayOptionShrink(v string) *UpdateGatewayOptionShrinkRequest
	GetGatewayOptionShrink() *string
	SetGatewayUniqueId(v string) *UpdateGatewayOptionShrinkRequest
	GetGatewayUniqueId() *string
}

type UpdateGatewayOptionShrinkRequest struct {
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
	GatewayOptionShrink *string `json:"GatewayOption,omitempty" xml:"GatewayOption,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-83b0ddb569434f82b9fe8e4c60c4****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s UpdateGatewayOptionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayOptionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayOptionShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayOptionShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayOptionShrinkRequest) GetGatewayOptionShrink() *string {
	return s.GatewayOptionShrink
}

func (s *UpdateGatewayOptionShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayOptionShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayOptionShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayOptionShrinkRequest) SetGatewayId(v int64) *UpdateGatewayOptionShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayOptionShrinkRequest) SetGatewayOptionShrink(v string) *UpdateGatewayOptionShrinkRequest {
	s.GatewayOptionShrink = &v
	return s
}

func (s *UpdateGatewayOptionShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayOptionShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayOptionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
