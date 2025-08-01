// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetGatewayRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetGatewayRequest
	GetGatewayUniqueId() *string
}

type GetGatewayRequest struct {
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
	// gw-685f661467b54f48b7b7a76605ce****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s GetGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetGatewayRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayRequest) SetAcceptLanguage(v string) *GetGatewayRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetGatewayRequest) SetGatewayUniqueId(v string) *GetGatewayRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayRequest) Validate() error {
	return dara.Validate(s)
}
