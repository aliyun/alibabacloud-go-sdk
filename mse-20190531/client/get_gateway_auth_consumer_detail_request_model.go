// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayAuthConsumerDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetGatewayAuthConsumerDetailRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetGatewayAuthConsumerDetailRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *GetGatewayAuthConsumerDetailRequest
	GetId() *int64
}

type GetGatewayAuthConsumerDetailRequest struct {
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
	// gw-1a4ab101d5924b6f92c5ec98a84*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the consumer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetGatewayAuthConsumerDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthConsumerDetailRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthConsumerDetailRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetGatewayAuthConsumerDetailRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayAuthConsumerDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayAuthConsumerDetailRequest) SetAcceptLanguage(v string) *GetGatewayAuthConsumerDetailRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailRequest) SetGatewayUniqueId(v string) *GetGatewayAuthConsumerDetailRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailRequest) SetId(v int64) *GetGatewayAuthConsumerDetailRequest {
	s.Id = &v
	return s
}

func (s *GetGatewayAuthConsumerDetailRequest) Validate() error {
	return dara.Validate(s)
}
