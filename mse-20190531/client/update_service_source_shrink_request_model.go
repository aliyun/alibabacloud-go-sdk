// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateServiceSourceShrinkRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *UpdateServiceSourceShrinkRequest
	GetAddress() *string
	SetGatewayId(v int64) *UpdateServiceSourceShrinkRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateServiceSourceShrinkRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateServiceSourceShrinkRequest
	GetId() *int64
	SetIngressOptionsRequestShrink(v string) *UpdateServiceSourceShrinkRequest
	GetIngressOptionsRequestShrink() *string
	SetName(v string) *UpdateServiceSourceShrinkRequest
	GetName() *string
	SetPathListShrink(v string) *UpdateServiceSourceShrinkRequest
	GetPathListShrink() *string
	SetSource(v string) *UpdateServiceSourceShrinkRequest
	GetSource() *string
	SetType(v string) *UpdateServiceSourceShrinkRequest
	GetType() *string
}

type UpdateServiceSourceShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese.
	//
	// 	- en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The address.
	//
	// example:
	//
	// c9ad2a0717032427e920754e25b49e3b5
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 429
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-cdd1bb9bfb8341e9805f931a3ba1f4c6
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the service source.
	//
	// example:
	//
	// 63
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The configurations of Ingress resources.
	//
	// example:
	//
	// zh
	IngressOptionsRequestShrink *string `json:"IngressOptionsRequest,omitempty" xml:"IngressOptionsRequest,omitempty"`
	// The name.
	//
	// example:
	//
	// istio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array of service root paths.
	PathListShrink *string `json:"PathList,omitempty" xml:"PathList,omitempty"`
	// The service source. Valid values:
	//
	// 	- K8s: ACK cluster.
	//
	// 	- MSE: Nacos instance.
	//
	// example:
	//
	// K8s
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the service source. Valid values:
	//
	// 	- K8s: ACK cluster.
	//
	// 	- NACOS: Nacos instance.
	//
	// example:
	//
	// K8s
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateServiceSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateServiceSourceShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateServiceSourceShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateServiceSourceShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateServiceSourceShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateServiceSourceShrinkRequest) GetIngressOptionsRequestShrink() *string {
	return s.IngressOptionsRequestShrink
}

func (s *UpdateServiceSourceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateServiceSourceShrinkRequest) GetPathListShrink() *string {
	return s.PathListShrink
}

func (s *UpdateServiceSourceShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateServiceSourceShrinkRequest) GetType() *string {
	return s.Type
}

func (s *UpdateServiceSourceShrinkRequest) SetAcceptLanguage(v string) *UpdateServiceSourceShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) SetAddress(v string) *UpdateServiceSourceShrinkRequest {
	s.Address = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) SetGatewayId(v int64) *UpdateServiceSourceShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) SetGatewayUniqueId(v string) *UpdateServiceSourceShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) SetId(v int64) *UpdateServiceSourceShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) SetIngressOptionsRequestShrink(v string) *UpdateServiceSourceShrinkRequest {
	s.IngressOptionsRequestShrink = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) SetName(v string) *UpdateServiceSourceShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) SetPathListShrink(v string) *UpdateServiceSourceShrinkRequest {
	s.PathListShrink = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) SetSource(v string) *UpdateServiceSourceShrinkRequest {
	s.Source = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) SetType(v string) *UpdateServiceSourceShrinkRequest {
	s.Type = &v
	return s
}

func (s *UpdateServiceSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
