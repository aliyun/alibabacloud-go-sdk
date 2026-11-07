// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListServiceSourceRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *ListServiceSourceRequest
	GetGatewayUniqueId() *string
	SetSource(v string) *ListServiceSourceRequest
	GetSource() *string
}

type ListServiceSourceRequest struct {
	// The language of the returned information. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-c51a15c7ee934a4fb945ccf35fe1****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies the service source type. If this parameter is left empty, all types are returned. Valid values:
	//
	// - K8s: Container Service.
	//
	// - MSE: MSE Nacos.
	//
	// - MSE_ZK: MSE ZooKeeper.
	//
	// - SAE: SAE registry.
	//
	// - EDAS: EDAS registry.
	//
	// example:
	//
	// MSE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListServiceSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSourceRequest) GoString() string {
	return s.String()
}

func (s *ListServiceSourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListServiceSourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListServiceSourceRequest) GetSource() *string {
	return s.Source
}

func (s *ListServiceSourceRequest) SetAcceptLanguage(v string) *ListServiceSourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListServiceSourceRequest) SetGatewayUniqueId(v string) *ListServiceSourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListServiceSourceRequest) SetSource(v string) *ListServiceSourceRequest {
	s.Source = &v
	return s
}

func (s *ListServiceSourceRequest) Validate() error {
	return dara.Validate(s)
}
