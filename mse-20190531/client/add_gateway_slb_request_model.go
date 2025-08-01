// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewaySlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewaySlbRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *AddGatewaySlbRequest
	GetGatewayUniqueId() *string
	SetHttpPort(v int32) *AddGatewaySlbRequest
	GetHttpPort() *int32
	SetHttpsPort(v int32) *AddGatewaySlbRequest
	GetHttpsPort() *int32
	SetHttpsVServerGroupId(v string) *AddGatewaySlbRequest
	GetHttpsVServerGroupId() *string
	SetServiceWeight(v int32) *AddGatewaySlbRequest
	GetServiceWeight() *int32
	SetSlbId(v string) *AddGatewaySlbRequest
	GetSlbId() *string
	SetType(v string) *AddGatewaySlbRequest
	GetType() *string
	SetVServerGroupId(v string) *AddGatewaySlbRequest
	GetVServerGroupId() *string
	SetVServiceList(v []*AddGatewaySlbRequestVServiceList) *AddGatewaySlbRequest
	GetVServiceList() []*AddGatewaySlbRequestVServiceList
}

type AddGatewaySlbRequest struct {
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
	VServiceList []*AddGatewaySlbRequestVServiceList `json:"VServiceList,omitempty" xml:"VServiceList,omitempty" type:"Repeated"`
}

func (s AddGatewaySlbRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewaySlbRequest) GoString() string {
	return s.String()
}

func (s *AddGatewaySlbRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewaySlbRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewaySlbRequest) GetHttpPort() *int32 {
	return s.HttpPort
}

func (s *AddGatewaySlbRequest) GetHttpsPort() *int32 {
	return s.HttpsPort
}

func (s *AddGatewaySlbRequest) GetHttpsVServerGroupId() *string {
	return s.HttpsVServerGroupId
}

func (s *AddGatewaySlbRequest) GetServiceWeight() *int32 {
	return s.ServiceWeight
}

func (s *AddGatewaySlbRequest) GetSlbId() *string {
	return s.SlbId
}

func (s *AddGatewaySlbRequest) GetType() *string {
	return s.Type
}

func (s *AddGatewaySlbRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *AddGatewaySlbRequest) GetVServiceList() []*AddGatewaySlbRequestVServiceList {
	return s.VServiceList
}

func (s *AddGatewaySlbRequest) SetAcceptLanguage(v string) *AddGatewaySlbRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewaySlbRequest) SetGatewayUniqueId(v string) *AddGatewaySlbRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewaySlbRequest) SetHttpPort(v int32) *AddGatewaySlbRequest {
	s.HttpPort = &v
	return s
}

func (s *AddGatewaySlbRequest) SetHttpsPort(v int32) *AddGatewaySlbRequest {
	s.HttpsPort = &v
	return s
}

func (s *AddGatewaySlbRequest) SetHttpsVServerGroupId(v string) *AddGatewaySlbRequest {
	s.HttpsVServerGroupId = &v
	return s
}

func (s *AddGatewaySlbRequest) SetServiceWeight(v int32) *AddGatewaySlbRequest {
	s.ServiceWeight = &v
	return s
}

func (s *AddGatewaySlbRequest) SetSlbId(v string) *AddGatewaySlbRequest {
	s.SlbId = &v
	return s
}

func (s *AddGatewaySlbRequest) SetType(v string) *AddGatewaySlbRequest {
	s.Type = &v
	return s
}

func (s *AddGatewaySlbRequest) SetVServerGroupId(v string) *AddGatewaySlbRequest {
	s.VServerGroupId = &v
	return s
}

func (s *AddGatewaySlbRequest) SetVServiceList(v []*AddGatewaySlbRequestVServiceList) *AddGatewaySlbRequest {
	s.VServiceList = v
	return s
}

func (s *AddGatewaySlbRequest) Validate() error {
	return dara.Validate(s)
}

type AddGatewaySlbRequestVServiceList struct {
	// The port number.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTPS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the virtual server group.
	//
	// example:
	//
	// rsp-bp1j**t0fyl**
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The name of the virtual server group.
	VServerGroupName *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s AddGatewaySlbRequestVServiceList) String() string {
	return dara.Prettify(s)
}

func (s AddGatewaySlbRequestVServiceList) GoString() string {
	return s.String()
}

func (s *AddGatewaySlbRequestVServiceList) GetPort() *int32 {
	return s.Port
}

func (s *AddGatewaySlbRequestVServiceList) GetProtocol() *string {
	return s.Protocol
}

func (s *AddGatewaySlbRequestVServiceList) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *AddGatewaySlbRequestVServiceList) GetVServerGroupName() *string {
	return s.VServerGroupName
}

func (s *AddGatewaySlbRequestVServiceList) SetPort(v int32) *AddGatewaySlbRequestVServiceList {
	s.Port = &v
	return s
}

func (s *AddGatewaySlbRequestVServiceList) SetProtocol(v string) *AddGatewaySlbRequestVServiceList {
	s.Protocol = &v
	return s
}

func (s *AddGatewaySlbRequestVServiceList) SetVServerGroupId(v string) *AddGatewaySlbRequestVServiceList {
	s.VServerGroupId = &v
	return s
}

func (s *AddGatewaySlbRequestVServiceList) SetVServerGroupName(v string) *AddGatewaySlbRequestVServiceList {
	s.VServerGroupName = &v
	return s
}

func (s *AddGatewaySlbRequestVServiceList) Validate() error {
	return dara.Validate(s)
}
