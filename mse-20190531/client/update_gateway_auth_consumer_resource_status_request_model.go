// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerResourceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerResourceStatusRequest
	GetAcceptLanguage() *string
	SetConsumerId(v int64) *UpdateGatewayAuthConsumerResourceStatusRequest
	GetConsumerId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerResourceStatusRequest
	GetGatewayUniqueId() *string
	SetIdList(v string) *UpdateGatewayAuthConsumerResourceStatusRequest
	GetIdList() *string
	SetResourceStatus(v bool) *UpdateGatewayAuthConsumerResourceStatusRequest
	GetResourceStatus() *bool
}

type UpdateGatewayAuthConsumerResourceStatusRequest struct {
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
	// The ID of the consumer.
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
	// gw-2a99625886d54722be94d92e9a69****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The list of IDs of the authorized resources that a user wants to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3,4
	IdList *string `json:"IdList,omitempty" xml:"IdList,omitempty"`
	// The resource authorization status. Valid values:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ResourceStatus *bool `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
}

func (s UpdateGatewayAuthConsumerResourceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResourceStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) GetConsumerId() *int64 {
	return s.ConsumerId
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) GetIdList() *string {
	return s.IdList
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) GetResourceStatus() *bool {
	return s.ResourceStatus
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerResourceStatusRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) SetConsumerId(v int64) *UpdateGatewayAuthConsumerResourceStatusRequest {
	s.ConsumerId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerResourceStatusRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) SetIdList(v string) *UpdateGatewayAuthConsumerResourceStatusRequest {
	s.IdList = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) SetResourceStatus(v bool) *UpdateGatewayAuthConsumerResourceStatusRequest {
	s.ResourceStatus = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusRequest) Validate() error {
	return dara.Validate(s)
}
