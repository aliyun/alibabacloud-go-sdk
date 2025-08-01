// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayServiceVersionRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *AddGatewayServiceVersionRequest
	GetGatewayUniqueId() *string
	SetServiceId(v int64) *AddGatewayServiceVersionRequest
	GetServiceId() *int64
	SetServiceVersion(v string) *AddGatewayServiceVersionRequest
	GetServiceVersion() *string
}

type AddGatewayServiceVersionRequest struct {
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
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-913a49bab6c5461187a3305fb8da****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 33
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// {\\"name\\":\\"v1\\",\\"labels\\":[{\\"key\\":\\"version\\",\\"value\\":\\"v1\\"}]}
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s AddGatewayServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayServiceVersionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayServiceVersionRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewayServiceVersionRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *AddGatewayServiceVersionRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *AddGatewayServiceVersionRequest) SetAcceptLanguage(v string) *AddGatewayServiceVersionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayServiceVersionRequest) SetGatewayUniqueId(v string) *AddGatewayServiceVersionRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewayServiceVersionRequest) SetServiceId(v int64) *AddGatewayServiceVersionRequest {
	s.ServiceId = &v
	return s
}

func (s *AddGatewayServiceVersionRequest) SetServiceVersion(v string) *AddGatewayServiceVersionRequest {
	s.ServiceVersion = &v
	return s
}

func (s *AddGatewayServiceVersionRequest) Validate() error {
	return dara.Validate(s)
}
