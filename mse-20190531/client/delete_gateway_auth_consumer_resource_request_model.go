// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayAuthConsumerResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayAuthConsumerResourceRequest
	GetAcceptLanguage() *string
	SetConsumerId(v int64) *DeleteGatewayAuthConsumerResourceRequest
	GetConsumerId() *int64
	SetGatewayUniqueId(v string) *DeleteGatewayAuthConsumerResourceRequest
	GetGatewayUniqueId() *string
	SetIdList(v string) *DeleteGatewayAuthConsumerResourceRequest
	GetIdList() *string
}

type DeleteGatewayAuthConsumerResourceRequest struct {
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
	// The consumer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ConsumerId *int64 `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-ubuwqygbq4783gqb2y3f87q****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The IDs of the authorized resources that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3,4
	IdList *string `json:"IdList,omitempty" xml:"IdList,omitempty"`
}

func (s DeleteGatewayAuthConsumerResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayAuthConsumerResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayAuthConsumerResourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayAuthConsumerResourceRequest) GetConsumerId() *int64 {
	return s.ConsumerId
}

func (s *DeleteGatewayAuthConsumerResourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayAuthConsumerResourceRequest) GetIdList() *string {
	return s.IdList
}

func (s *DeleteGatewayAuthConsumerResourceRequest) SetAcceptLanguage(v string) *DeleteGatewayAuthConsumerResourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceRequest) SetConsumerId(v int64) *DeleteGatewayAuthConsumerResourceRequest {
	s.ConsumerId = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceRequest) SetGatewayUniqueId(v string) *DeleteGatewayAuthConsumerResourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceRequest) SetIdList(v string) *DeleteGatewayAuthConsumerResourceRequest {
	s.IdList = &v
	return s
}

func (s *DeleteGatewayAuthConsumerResourceRequest) Validate() error {
	return dara.Validate(s)
}
