// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKubernetesSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetKubernetesSourceRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetKubernetesSourceRequest
	GetGatewayUniqueId() *string
	SetIsAll(v bool) *GetKubernetesSourceRequest
	GetIsAll() *bool
	SetVpcId(v string) *GetKubernetesSourceRequest
	GetVpcId() *string
}

type GetKubernetesSourceRequest struct {
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
	// gw-c5d1aadb7df646cfb7065fbf75c1****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies whether to obtain the information about all Kubernetes clusters. If you set the value to false, only the information about unassociated clusters is obtained.
	IsAll *bool `json:"IsAll,omitempty" xml:"IsAll,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1t50e045b5g7i3p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetKubernetesSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKubernetesSourceRequest) GoString() string {
	return s.String()
}

func (s *GetKubernetesSourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetKubernetesSourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetKubernetesSourceRequest) GetIsAll() *bool {
	return s.IsAll
}

func (s *GetKubernetesSourceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *GetKubernetesSourceRequest) SetAcceptLanguage(v string) *GetKubernetesSourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetKubernetesSourceRequest) SetGatewayUniqueId(v string) *GetKubernetesSourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetKubernetesSourceRequest) SetIsAll(v bool) *GetKubernetesSourceRequest {
	s.IsAll = &v
	return s
}

func (s *GetKubernetesSourceRequest) SetVpcId(v string) *GetKubernetesSourceRequest {
	s.VpcId = &v
	return s
}

func (s *GetKubernetesSourceRequest) Validate() error {
	return dara.Validate(s)
}
