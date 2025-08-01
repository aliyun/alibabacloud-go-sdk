// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddServiceSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddServiceSourceShrinkRequest
	GetAcceptLanguage() *string
	SetAddress(v string) *AddServiceSourceShrinkRequest
	GetAddress() *string
	SetGatewayUniqueId(v string) *AddServiceSourceShrinkRequest
	GetGatewayUniqueId() *string
	SetGroupListShrink(v string) *AddServiceSourceShrinkRequest
	GetGroupListShrink() *string
	SetIngressOptionsRequestShrink(v string) *AddServiceSourceShrinkRequest
	GetIngressOptionsRequestShrink() *string
	SetName(v string) *AddServiceSourceShrinkRequest
	GetName() *string
	SetPathListShrink(v string) *AddServiceSourceShrinkRequest
	GetPathListShrink() *string
	SetSource(v string) *AddServiceSourceShrinkRequest
	GetSource() *string
	SetToAuthorizeSecurityGroupsShrink(v string) *AddServiceSourceShrinkRequest
	GetToAuthorizeSecurityGroupsShrink() *string
	SetType(v string) *AddServiceSourceShrinkRequest
	GetType() *string
}

type AddServiceSourceShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN (default): Chinese
	//
	// 	- en-US: English
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Specifies whether to monitor Ingress classes.
	//
	// example:
	//
	// c9ad2a0717032427e920754e25b49e3b5
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Specifies whether to update the Ingress status.
	//
	// example:
	//
	// gw-c70622ff52fe49beb29bea9a6f52****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The data structure.
	GroupListShrink *string `json:"GroupList,omitempty" xml:"GroupList,omitempty"`
	// The list of service groups.
	IngressOptionsRequestShrink *string `json:"IngressOptionsRequest,omitempty" xml:"IngressOptionsRequest,omitempty"`
	// The namespace whose resources you want to monitor.
	//
	// example:
	//
	// istio
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The HTTP status code returned.
	PathListShrink *string `json:"PathList,omitempty" xml:"PathList,omitempty"`
	// The service source.
	//
	// 	- K8s: ACK cluster
	//
	// 	- NACOS: MSE Nacos instance
	//
	// example:
	//
	// K8s,MSE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of security groups to be authorized. You can specify security groups to allow backend services to access data sources that you create.
	ToAuthorizeSecurityGroupsShrink *string `json:"ToAuthorizeSecurityGroups,omitempty" xml:"ToAuthorizeSecurityGroups,omitempty"`
	// The type of the service source.
	//
	// 	- K8s: Container Service for Kubernetes (ACK) cluster
	//
	// 	- NACOS: Nacos instance
	//
	// example:
	//
	// The Ingress configuration.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddServiceSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddServiceSourceShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddServiceSourceShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *AddServiceSourceShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddServiceSourceShrinkRequest) GetGroupListShrink() *string {
	return s.GroupListShrink
}

func (s *AddServiceSourceShrinkRequest) GetIngressOptionsRequestShrink() *string {
	return s.IngressOptionsRequestShrink
}

func (s *AddServiceSourceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddServiceSourceShrinkRequest) GetPathListShrink() *string {
	return s.PathListShrink
}

func (s *AddServiceSourceShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *AddServiceSourceShrinkRequest) GetToAuthorizeSecurityGroupsShrink() *string {
	return s.ToAuthorizeSecurityGroupsShrink
}

func (s *AddServiceSourceShrinkRequest) GetType() *string {
	return s.Type
}

func (s *AddServiceSourceShrinkRequest) SetAcceptLanguage(v string) *AddServiceSourceShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) SetAddress(v string) *AddServiceSourceShrinkRequest {
	s.Address = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) SetGatewayUniqueId(v string) *AddServiceSourceShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) SetGroupListShrink(v string) *AddServiceSourceShrinkRequest {
	s.GroupListShrink = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) SetIngressOptionsRequestShrink(v string) *AddServiceSourceShrinkRequest {
	s.IngressOptionsRequestShrink = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) SetName(v string) *AddServiceSourceShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) SetPathListShrink(v string) *AddServiceSourceShrinkRequest {
	s.PathListShrink = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) SetSource(v string) *AddServiceSourceShrinkRequest {
	s.Source = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) SetToAuthorizeSecurityGroupsShrink(v string) *AddServiceSourceShrinkRequest {
	s.ToAuthorizeSecurityGroupsShrink = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) SetType(v string) *AddServiceSourceShrinkRequest {
	s.Type = &v
	return s
}

func (s *AddServiceSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
