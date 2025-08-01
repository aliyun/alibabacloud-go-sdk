// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteHeaderOpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteHeaderOpRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayRouteHeaderOpRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteHeaderOpRequest
	GetGatewayUniqueId() *string
	SetHeaderOpJSON(v string) *UpdateGatewayRouteHeaderOpRequest
	GetHeaderOpJSON() *string
	SetId(v int64) *UpdateGatewayRouteHeaderOpRequest
	GetId() *int64
}

type UpdateGatewayRouteHeaderOpRequest struct {
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
	// 324
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-6bdc977deda44bf589c49d063b4c2d1d
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The description of user header settings.
	//
	// example:
	//
	// {\\"status\\":\\"on\\",\\"headerOpItems\\":[{\\"directionType\\":\\"Request\\",\\"opType\\":\\"Update\\",\\"key\\":\\"hosts\\",\\"value\\":\\"test.com\\"}]}
	HeaderOpJSON *string `json:"HeaderOpJSON,omitempty" xml:"HeaderOpJSON,omitempty"`
	// The ID of the record.
	//
	// example:
	//
	// 411
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateGatewayRouteHeaderOpRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteHeaderOpRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteHeaderOpRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteHeaderOpRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteHeaderOpRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteHeaderOpRequest) GetHeaderOpJSON() *string {
	return s.HeaderOpJSON
}

func (s *UpdateGatewayRouteHeaderOpRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteHeaderOpRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteHeaderOpRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpRequest) SetGatewayId(v int64) *UpdateGatewayRouteHeaderOpRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteHeaderOpRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpRequest) SetHeaderOpJSON(v string) *UpdateGatewayRouteHeaderOpRequest {
	s.HeaderOpJSON = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpRequest) SetId(v int64) *UpdateGatewayRouteHeaderOpRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpRequest) Validate() error {
	return dara.Validate(s)
}
