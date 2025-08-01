// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteOnAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayRouteOnAuthRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *ListGatewayRouteOnAuthRequest
	GetGatewayUniqueId() *string
	SetType(v string) *ListGatewayRouteOnAuthRequest
	GetType() *string
}

type ListGatewayRouteOnAuthRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// gw-c9bc5afd61014165bd58f621b491****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The authentication method. Valid values:
	//
	// 	- JWT
	//
	// This parameter is required.
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayRouteOnAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteOnAuthRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteOnAuthRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayRouteOnAuthRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayRouteOnAuthRequest) GetType() *string {
	return s.Type
}

func (s *ListGatewayRouteOnAuthRequest) SetAcceptLanguage(v string) *ListGatewayRouteOnAuthRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayRouteOnAuthRequest) SetGatewayUniqueId(v string) *ListGatewayRouteOnAuthRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayRouteOnAuthRequest) SetType(v string) *ListGatewayRouteOnAuthRequest {
	s.Type = &v
	return s
}

func (s *ListGatewayRouteOnAuthRequest) Validate() error {
	return dara.Validate(s)
}
