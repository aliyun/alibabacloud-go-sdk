// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerStatusRequest
	GetAcceptLanguage() *string
	SetConsumerStatus(v bool) *UpdateGatewayAuthConsumerStatusRequest
	GetConsumerStatus() *bool
	SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerStatusRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayAuthConsumerStatusRequest
	GetId() *int64
}

type UpdateGatewayAuthConsumerStatusRequest struct {
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
	// The status of the consumer. Valid values:
	//
	// 	- true: The consumer is enabled.
	//
	// 	- false: The consumer is disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ConsumerStatus *bool `json:"ConsumerStatus,omitempty" xml:"ConsumerStatus,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-685f661467b54f48b7b7a76605ce****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The consumer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateGatewayAuthConsumerStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerStatusRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayAuthConsumerStatusRequest) GetConsumerStatus() *bool {
	return s.ConsumerStatus
}

func (s *UpdateGatewayAuthConsumerStatusRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayAuthConsumerStatusRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayAuthConsumerStatusRequest) SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerStatusRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusRequest) SetConsumerStatus(v bool) *UpdateGatewayAuthConsumerStatusRequest {
	s.ConsumerStatus = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusRequest) SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerStatusRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusRequest) SetId(v int64) *UpdateGatewayAuthConsumerStatusRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusRequest) Validate() error {
	return dara.Validate(s)
}
