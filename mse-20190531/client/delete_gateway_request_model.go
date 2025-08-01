// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayRequest
	GetAcceptLanguage() *string
	SetDeleteSlb(v bool) *DeleteGatewayRequest
	GetDeleteSlb() *bool
	SetGatewayUniqueId(v string) *DeleteGatewayRequest
	GetGatewayUniqueId() *string
}

type DeleteGatewayRequest struct {
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
	// Specifies whether to delete the SLB instance purchased for the gateway when you delete the gateway.
	//
	// example:
	//
	// true
	DeleteSlb *bool `json:"DeleteSlb,omitempty" xml:"DeleteSlb,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-0fe488252dc44d55a9dd57875193a1d7
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s DeleteGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayRequest) GetDeleteSlb() *bool {
	return s.DeleteSlb
}

func (s *DeleteGatewayRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayRequest) SetAcceptLanguage(v string) *DeleteGatewayRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayRequest) SetDeleteSlb(v bool) *DeleteGatewayRequest {
	s.DeleteSlb = &v
	return s
}

func (s *DeleteGatewayRequest) SetGatewayUniqueId(v string) *DeleteGatewayRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayRequest) Validate() error {
	return dara.Validate(s)
}
