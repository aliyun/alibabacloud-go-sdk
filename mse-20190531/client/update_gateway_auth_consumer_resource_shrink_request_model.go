// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerResourceShrinkRequest
	GetAcceptLanguage() *string
	SetConsumerId(v int64) *UpdateGatewayAuthConsumerResourceShrinkRequest
	GetConsumerId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerResourceShrinkRequest
	GetGatewayUniqueId() *string
	SetResourceListShrink(v string) *UpdateGatewayAuthConsumerResourceShrinkRequest
	GetResourceListShrink() *string
}

type UpdateGatewayAuthConsumerResourceShrinkRequest struct {
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
	// The gateway authentication consumer ID.
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
	// gw-3f97e2989c344f35ab3fd62b19f1****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The gateway authentication consumer ID.
	ResourceListShrink *string `json:"ResourceList,omitempty" xml:"ResourceList,omitempty"`
}

func (s UpdateGatewayAuthConsumerResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResourceShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayAuthConsumerResourceShrinkRequest) GetConsumerId() *int64 {
	return s.ConsumerId
}

func (s *UpdateGatewayAuthConsumerResourceShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayAuthConsumerResourceShrinkRequest) GetResourceListShrink() *string {
	return s.ResourceListShrink
}

func (s *UpdateGatewayAuthConsumerResourceShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerResourceShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceShrinkRequest) SetConsumerId(v int64) *UpdateGatewayAuthConsumerResourceShrinkRequest {
	s.ConsumerId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerResourceShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceShrinkRequest) SetResourceListShrink(v string) *UpdateGatewayAuthConsumerResourceShrinkRequest {
	s.ResourceListShrink = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
