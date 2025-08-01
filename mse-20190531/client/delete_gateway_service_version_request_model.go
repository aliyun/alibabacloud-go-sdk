// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayServiceVersionRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteGatewayServiceVersionRequest
	GetGatewayUniqueId() *string
	SetServiceId(v int64) *DeleteGatewayServiceVersionRequest
	GetServiceId() *int64
	SetServiceVersion(v string) *DeleteGatewayServiceVersionRequest
	GetServiceVersion() *string
}

type DeleteGatewayServiceVersionRequest struct {
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
	// gw-b6988bd16920479d9104e1729f97****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 777
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// {\\"name\\":\\"app\\",\\"labels\\":[{\\"key\\":\\"app\\",\\"value\\":\\"demo-server\\"}]}
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s DeleteGatewayServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceVersionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayServiceVersionRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayServiceVersionRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *DeleteGatewayServiceVersionRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *DeleteGatewayServiceVersionRequest) SetAcceptLanguage(v string) *DeleteGatewayServiceVersionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayServiceVersionRequest) SetGatewayUniqueId(v string) *DeleteGatewayServiceVersionRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayServiceVersionRequest) SetServiceId(v int64) *DeleteGatewayServiceVersionRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteGatewayServiceVersionRequest) SetServiceVersion(v string) *DeleteGatewayServiceVersionRequest {
	s.ServiceVersion = &v
	return s
}

func (s *DeleteGatewayServiceVersionRequest) Validate() error {
	return dara.Validate(s)
}
