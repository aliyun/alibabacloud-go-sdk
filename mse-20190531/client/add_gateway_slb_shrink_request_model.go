// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewaySlbShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewaySlbShrinkRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *AddGatewaySlbShrinkRequest
	GetGatewayUniqueId() *string
	SetHttpPort(v int32) *AddGatewaySlbShrinkRequest
	GetHttpPort() *int32
	SetHttpsPort(v int32) *AddGatewaySlbShrinkRequest
	GetHttpsPort() *int32
	SetHttpsVServerGroupId(v string) *AddGatewaySlbShrinkRequest
	GetHttpsVServerGroupId() *string
	SetServiceWeight(v int32) *AddGatewaySlbShrinkRequest
	GetServiceWeight() *int32
	SetSlbId(v string) *AddGatewaySlbShrinkRequest
	GetSlbId() *string
	SetType(v string) *AddGatewaySlbShrinkRequest
	GetType() *string
	SetVServerGroupId(v string) *AddGatewaySlbShrinkRequest
	GetVServerGroupId() *string
	SetVServiceListShrink(v string) *AddGatewaySlbShrinkRequest
	GetVServiceListShrink() *string
}

type AddGatewaySlbShrinkRequest struct {
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
	// gw-9cdcf8e4f58144059e73ff4c5ef9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The HTTP port number (virtual service group).
	//
	// example:
	//
	// 80
	HttpPort *int32 `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// The HTTPS port number (virtual service group).
	//
	// example:
	//
	// 443
	HttpsPort *int32 `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	// The ID of the HTTPS virtual service group.
	//
	// example:
	//
	// 353
	HttpsVServerGroupId *string `json:"HttpsVServerGroupId,omitempty" xml:"HttpsVServerGroupId,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 80
	ServiceWeight *int32 `json:"ServiceWeight,omitempty" xml:"ServiceWeight,omitempty"`
	// The ID of the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp18t6jjskwxh6wy1****
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The type of the service source. Valid values:
	//
	// 	- PUB_NET: Internet
	//
	// 	- PRIVATE_NET: VPC
	//
	// example:
	//
	// PUB_NET
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the HTTP virtual service group.
	//
	// example:
	//
	// 353
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The SLB monitoring information.
	VServiceListShrink *string `json:"VServiceList,omitempty" xml:"VServiceList,omitempty"`
}

func (s AddGatewaySlbShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewaySlbShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddGatewaySlbShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewaySlbShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewaySlbShrinkRequest) GetHttpPort() *int32 {
	return s.HttpPort
}

func (s *AddGatewaySlbShrinkRequest) GetHttpsPort() *int32 {
	return s.HttpsPort
}

func (s *AddGatewaySlbShrinkRequest) GetHttpsVServerGroupId() *string {
	return s.HttpsVServerGroupId
}

func (s *AddGatewaySlbShrinkRequest) GetServiceWeight() *int32 {
	return s.ServiceWeight
}

func (s *AddGatewaySlbShrinkRequest) GetSlbId() *string {
	return s.SlbId
}

func (s *AddGatewaySlbShrinkRequest) GetType() *string {
	return s.Type
}

func (s *AddGatewaySlbShrinkRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *AddGatewaySlbShrinkRequest) GetVServiceListShrink() *string {
	return s.VServiceListShrink
}

func (s *AddGatewaySlbShrinkRequest) SetAcceptLanguage(v string) *AddGatewaySlbShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) SetGatewayUniqueId(v string) *AddGatewaySlbShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) SetHttpPort(v int32) *AddGatewaySlbShrinkRequest {
	s.HttpPort = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) SetHttpsPort(v int32) *AddGatewaySlbShrinkRequest {
	s.HttpsPort = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) SetHttpsVServerGroupId(v string) *AddGatewaySlbShrinkRequest {
	s.HttpsVServerGroupId = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) SetServiceWeight(v int32) *AddGatewaySlbShrinkRequest {
	s.ServiceWeight = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) SetSlbId(v string) *AddGatewaySlbShrinkRequest {
	s.SlbId = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) SetType(v string) *AddGatewaySlbShrinkRequest {
	s.Type = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) SetVServerGroupId(v string) *AddGatewaySlbShrinkRequest {
	s.VServerGroupId = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) SetVServiceListShrink(v string) *AddGatewaySlbShrinkRequest {
	s.VServiceListShrink = &v
	return s
}

func (s *AddGatewaySlbShrinkRequest) Validate() error {
	return dara.Validate(s)
}
