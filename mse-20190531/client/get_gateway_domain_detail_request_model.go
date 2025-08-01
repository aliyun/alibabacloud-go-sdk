// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetGatewayDomainDetailRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetGatewayDomainDetailRequest
	GetGatewayUniqueId() *string
	SetId(v string) *GetGatewayDomainDetailRequest
	GetId() *string
}

type GetGatewayDomainDetailRequest struct {
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
	// gw-6d0b23e1d39e41658a968d79a635****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// 29
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetGatewayDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayDomainDetailRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetGatewayDomainDetailRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayDomainDetailRequest) GetId() *string {
	return s.Id
}

func (s *GetGatewayDomainDetailRequest) SetAcceptLanguage(v string) *GetGatewayDomainDetailRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetGatewayDomainDetailRequest) SetGatewayUniqueId(v string) *GetGatewayDomainDetailRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayDomainDetailRequest) SetId(v string) *GetGatewayDomainDetailRequest {
	s.Id = &v
	return s
}

func (s *GetGatewayDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
