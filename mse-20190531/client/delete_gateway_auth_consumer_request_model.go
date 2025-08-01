// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayAuthConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayAuthConsumerRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteGatewayAuthConsumerRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *DeleteGatewayAuthConsumerRequest
	GetId() *int64
}

type DeleteGatewayAuthConsumerRequest struct {
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
	// gw-c70622ff52fe49beb29bea9a6f52****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the consumer on which the gateway performs authentication operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33ff74b6-d21e-4f9b-91a8-bc1ea4ef****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteGatewayAuthConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthConsumerRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthConsumerRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayAuthConsumerRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayAuthConsumerRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteGatewayAuthConsumerRequest) SetAcceptLanguage(v string) *DeleteGatewayAuthConsumerRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayAuthConsumerRequest) SetGatewayUniqueId(v string) *DeleteGatewayAuthConsumerRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayAuthConsumerRequest) SetId(v int64) *DeleteGatewayAuthConsumerRequest {
	s.Id = &v
	return s
}

func (s *DeleteGatewayAuthConsumerRequest) Validate() error {
	return dara.Validate(s)
}
