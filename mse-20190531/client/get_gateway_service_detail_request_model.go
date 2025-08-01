// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayServiceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetGatewayServiceDetailRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetGatewayServiceDetailRequest
	GetGatewayUniqueId() *string
	SetServiceId(v int64) *GetGatewayServiceDetailRequest
	GetServiceId() *int64
}

type GetGatewayServiceDetailRequest struct {
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
	// gw-f4c960ad071a48a790b36324394c****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 35
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s GetGatewayServiceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetGatewayServiceDetailRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayServiceDetailRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *GetGatewayServiceDetailRequest) SetAcceptLanguage(v string) *GetGatewayServiceDetailRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetGatewayServiceDetailRequest) SetGatewayUniqueId(v string) *GetGatewayServiceDetailRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayServiceDetailRequest) SetServiceId(v int64) *GetGatewayServiceDetailRequest {
	s.ServiceId = &v
	return s
}

func (s *GetGatewayServiceDetailRequest) Validate() error {
	return dara.Validate(s)
}
