// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectGatewaySlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *SelectGatewaySlbRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *SelectGatewaySlbRequest
	GetGatewayUniqueId() *string
	SetName(v string) *SelectGatewaySlbRequest
	GetName() *string
	SetType(v string) *SelectGatewaySlbRequest
	GetType() *string
}

type SelectGatewaySlbRequest struct {
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
	// gw-492af9b04bb4474cae9d645be850****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The name of the SLB instance.
	//
	// example:
	//
	// test-slb
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the SLB instance.
	//
	// example:
	//
	// PUB_NET
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SelectGatewaySlbRequest) String() string {
	return dara.Prettify(s)
}

func (s SelectGatewaySlbRequest) GoString() string {
	return s.String()
}

func (s *SelectGatewaySlbRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *SelectGatewaySlbRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *SelectGatewaySlbRequest) GetName() *string {
	return s.Name
}

func (s *SelectGatewaySlbRequest) GetType() *string {
	return s.Type
}

func (s *SelectGatewaySlbRequest) SetAcceptLanguage(v string) *SelectGatewaySlbRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *SelectGatewaySlbRequest) SetGatewayUniqueId(v string) *SelectGatewaySlbRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *SelectGatewaySlbRequest) SetName(v string) *SelectGatewaySlbRequest {
	s.Name = &v
	return s
}

func (s *SelectGatewaySlbRequest) SetType(v string) *SelectGatewaySlbRequest {
	s.Type = &v
	return s
}

func (s *SelectGatewaySlbRequest) Validate() error {
	return dara.Validate(s)
}
