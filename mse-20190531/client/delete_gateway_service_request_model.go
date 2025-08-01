// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayServiceRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *DeleteGatewayServiceRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *DeleteGatewayServiceRequest
	GetGatewayUniqueId() *string
	SetServiceId(v string) *DeleteGatewayServiceRequest
	GetServiceId() *string
}

type DeleteGatewayServiceRequest struct {
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
	// 60
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-1a4ab101d5924b6f92c5ec98a841761f
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 190
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteGatewayServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayServiceRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *DeleteGatewayServiceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeleteGatewayServiceRequest) SetAcceptLanguage(v string) *DeleteGatewayServiceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayServiceRequest) SetGatewayId(v int64) *DeleteGatewayServiceRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayServiceRequest) SetGatewayUniqueId(v string) *DeleteGatewayServiceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayServiceRequest) SetServiceId(v string) *DeleteGatewayServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteGatewayServiceRequest) Validate() error {
	return dara.Validate(s)
}
