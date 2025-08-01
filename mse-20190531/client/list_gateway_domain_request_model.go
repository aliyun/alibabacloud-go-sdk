// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayDomainRequest
	GetAcceptLanguage() *string
	SetDomainName(v string) *ListGatewayDomainRequest
	GetDomainName() *string
	SetGatewayUniqueId(v string) *ListGatewayDomainRequest
	GetGatewayUniqueId() *string
	SetType(v string) *ListGatewayDomainRequest
	GetType() *string
}

type ListGatewayDomainRequest struct {
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
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-c9bc5afd61014165bd58f621b491****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The type of the domain name.
	//
	// example:
	//
	// All
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ListGatewayDomainRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayDomainRequest) GetType() *string {
	return s.Type
}

func (s *ListGatewayDomainRequest) SetAcceptLanguage(v string) *ListGatewayDomainRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayDomainRequest) SetDomainName(v string) *ListGatewayDomainRequest {
	s.DomainName = &v
	return s
}

func (s *ListGatewayDomainRequest) SetGatewayUniqueId(v string) *ListGatewayDomainRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayDomainRequest) SetType(v string) *ListGatewayDomainRequest {
	s.Type = &v
	return s
}

func (s *ListGatewayDomainRequest) Validate() error {
	return dara.Validate(s)
}
