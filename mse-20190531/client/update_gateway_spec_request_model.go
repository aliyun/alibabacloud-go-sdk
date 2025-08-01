// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewaySpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewaySpecRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *UpdateGatewaySpecRequest
	GetGatewayUniqueId() *string
	SetReplica(v int32) *UpdateGatewaySpecRequest
	GetReplica() *int32
	SetSpec(v string) *UpdateGatewaySpecRequest
	GetSpec() *string
}

type UpdateGatewaySpecRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// gw-c70622ff52fe49beb29bea9a6f52****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The number of nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Replica *int32 `json:"Replica,omitempty" xml:"Replica,omitempty"`
	// The node specifications of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// MSE_GTW_4_8_200_c
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateGatewaySpecRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewaySpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewaySpecRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewaySpecRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewaySpecRequest) GetReplica() *int32 {
	return s.Replica
}

func (s *UpdateGatewaySpecRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateGatewaySpecRequest) SetAcceptLanguage(v string) *UpdateGatewaySpecRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewaySpecRequest) SetGatewayUniqueId(v string) *UpdateGatewaySpecRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewaySpecRequest) SetReplica(v int32) *UpdateGatewaySpecRequest {
	s.Replica = &v
	return s
}

func (s *UpdateGatewaySpecRequest) SetSpec(v string) *UpdateGatewaySpecRequest {
	s.Spec = &v
	return s
}

func (s *UpdateGatewaySpecRequest) Validate() error {
	return dara.Validate(s)
}
