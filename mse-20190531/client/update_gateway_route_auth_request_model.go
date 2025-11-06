// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteAuthRequest
	GetAcceptLanguage() *string
	SetAuthJSON(v *UpdateGatewayRouteAuthRequestAuthJSON) *UpdateGatewayRouteAuthRequest
	GetAuthJSON() *UpdateGatewayRouteAuthRequestAuthJSON
	SetGatewayId(v int64) *UpdateGatewayRouteAuthRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteAuthRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteAuthRequest
	GetId() *int64
}

type UpdateGatewayRouteAuthRequest struct {
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
	// The authentication configurations.
	//
	// This parameter is required.
	AuthJSON *UpdateGatewayRouteAuthRequestAuthJSON `json:"AuthJSON,omitempty" xml:"AuthJSON,omitempty" type:"Struct"`
	// The gateway ID.
	//
	// example:
	//
	// 102
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-0adf3ad751284cc69fcf9669fba*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The route ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 109
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateGatewayRouteAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteAuthRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteAuthRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteAuthRequest) GetAuthJSON() *UpdateGatewayRouteAuthRequestAuthJSON {
	return s.AuthJSON
}

func (s *UpdateGatewayRouteAuthRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteAuthRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteAuthRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteAuthRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteAuthRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteAuthRequest) SetAuthJSON(v *UpdateGatewayRouteAuthRequestAuthJSON) *UpdateGatewayRouteAuthRequest {
	s.AuthJSON = v
	return s
}

func (s *UpdateGatewayRouteAuthRequest) SetGatewayId(v int64) *UpdateGatewayRouteAuthRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteAuthRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteAuthRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteAuthRequest) SetId(v int64) *UpdateGatewayRouteAuthRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteAuthRequest) Validate() error {
	if s.AuthJSON != nil {
		if err := s.AuthJSON.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayRouteAuthRequestAuthJSON struct {
	// The authentication type. If an empty string is passed, no authentication type is available. Valid values:
	//
	// 	- JWT
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateGatewayRouteAuthRequestAuthJSON) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteAuthRequestAuthJSON) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteAuthRequestAuthJSON) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayRouteAuthRequestAuthJSON) SetType(v string) *UpdateGatewayRouteAuthRequestAuthJSON {
	s.Type = &v
	return s
}

func (s *UpdateGatewayRouteAuthRequestAuthJSON) Validate() error {
	return dara.Validate(s)
}
