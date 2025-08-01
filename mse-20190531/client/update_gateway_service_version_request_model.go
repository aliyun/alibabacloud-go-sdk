// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayServiceVersionRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *UpdateGatewayServiceVersionRequest
	GetGatewayUniqueId() *string
	SetServiceId(v int64) *UpdateGatewayServiceVersionRequest
	GetServiceId() *int64
	SetServiceVersion(v string) *UpdateGatewayServiceVersionRequest
	GetServiceVersion() *string
}

type UpdateGatewayServiceVersionRequest struct {
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
	// gw-eccf313e2224438ba53d95d039e5****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 575
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// {\\"name\\":\\"v3\\",\\"labels\\":[{\\"key\\":\\"version\\",\\"value\\":\\"v3\\"}]}
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s UpdateGatewayServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceVersionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayServiceVersionRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayServiceVersionRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *UpdateGatewayServiceVersionRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *UpdateGatewayServiceVersionRequest) SetAcceptLanguage(v string) *UpdateGatewayServiceVersionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayServiceVersionRequest) SetGatewayUniqueId(v string) *UpdateGatewayServiceVersionRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayServiceVersionRequest) SetServiceId(v int64) *UpdateGatewayServiceVersionRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateGatewayServiceVersionRequest) SetServiceVersion(v string) *UpdateGatewayServiceVersionRequest {
	s.ServiceVersion = &v
	return s
}

func (s *UpdateGatewayServiceVersionRequest) Validate() error {
	return dara.Validate(s)
}
