// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayDomainRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteGatewayDomainRequest
	GetGatewayUniqueId() *string
	SetId(v string) *DeleteGatewayDomainRequest
	GetId() *string
}

type DeleteGatewayDomainRequest struct {
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
	// gw-90392d768a3847a7b804c505254d****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// 109
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteGatewayDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayDomainRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayDomainRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayDomainRequest) GetId() *string {
	return s.Id
}

func (s *DeleteGatewayDomainRequest) SetAcceptLanguage(v string) *DeleteGatewayDomainRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayDomainRequest) SetGatewayUniqueId(v string) *DeleteGatewayDomainRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayDomainRequest) SetId(v string) *DeleteGatewayDomainRequest {
	s.Id = &v
	return s
}

func (s *DeleteGatewayDomainRequest) Validate() error {
	return dara.Validate(s)
}
