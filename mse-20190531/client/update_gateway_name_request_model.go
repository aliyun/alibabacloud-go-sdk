// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayNameRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *UpdateGatewayNameRequest
	GetGatewayUniqueId() *string
	SetName(v string) *UpdateGatewayNameRequest
	GetName() *string
}

type UpdateGatewayNameRequest struct {
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
	// gw-1cef54brvecdb419fb264d4f9b8c
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The name of the gateway.
	//
	// example:
	//
	// demo-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateGatewayNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayNameRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayNameRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayNameRequest) SetAcceptLanguage(v string) *UpdateGatewayNameRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayNameRequest) SetGatewayUniqueId(v string) *UpdateGatewayNameRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayNameRequest) SetName(v string) *UpdateGatewayNameRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayNameRequest) Validate() error {
	return dara.Validate(s)
}
