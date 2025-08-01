// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMseSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetMseSourceRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetMseSourceRequest
	GetGatewayUniqueId() *string
	SetType(v string) *GetMseSourceRequest
	GetType() *string
	SetVpcId(v string) *GetMseSourceRequest
	GetVpcId() *string
}

type GetMseSourceRequest struct {
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
	// gw-7ea3da97b96543e19f6c597cd4a9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The MSE engine type. Valid values:
	//
	// 	- NACOS
	//
	// 	- ZOOKEEPER
	//
	// example:
	//
	// NACOS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-bp1t50e045b5g7i3p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetMseSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMseSourceRequest) GoString() string {
	return s.String()
}

func (s *GetMseSourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetMseSourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetMseSourceRequest) GetType() *string {
	return s.Type
}

func (s *GetMseSourceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *GetMseSourceRequest) SetAcceptLanguage(v string) *GetMseSourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetMseSourceRequest) SetGatewayUniqueId(v string) *GetMseSourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetMseSourceRequest) SetType(v string) *GetMseSourceRequest {
	s.Type = &v
	return s
}

func (s *GetMseSourceRequest) SetVpcId(v string) *GetMseSourceRequest {
	s.VpcId = &v
	return s
}

func (s *GetMseSourceRequest) Validate() error {
	return dara.Validate(s)
}
