// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthConsumerResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayAuthConsumerResourceRequest
	GetAcceptLanguage() *string
	SetConsumerId(v int64) *ListGatewayAuthConsumerResourceRequest
	GetConsumerId() *int64
	SetGatewayUniqueId(v string) *ListGatewayAuthConsumerResourceRequest
	GetGatewayUniqueId() *string
	SetPageNum(v string) *ListGatewayAuthConsumerResourceRequest
	GetPageNum() *string
	SetPageSize(v string) *ListGatewayAuthConsumerResourceRequest
	GetPageSize() *string
	SetResourceStatus(v bool) *ListGatewayAuthConsumerResourceRequest
	GetResourceStatus() *bool
	SetRouteName(v string) *ListGatewayAuthConsumerResourceRequest
	GetRouteName() *string
}

type ListGatewayAuthConsumerResourceRequest struct {
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
	// gw-3f97e2989c344f35ab3fd62b19f1****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource authorization status. Valid values:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	ResourceStatus *bool `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// test
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
}

func (s ListGatewayAuthConsumerResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerResourceRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerResourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayAuthConsumerResourceRequest) GetConsumerId() *int64 {
	return s.ConsumerId
}

func (s *ListGatewayAuthConsumerResourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayAuthConsumerResourceRequest) GetPageNum() *string {
	return s.PageNum
}

func (s *ListGatewayAuthConsumerResourceRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListGatewayAuthConsumerResourceRequest) GetResourceStatus() *bool {
	return s.ResourceStatus
}

func (s *ListGatewayAuthConsumerResourceRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *ListGatewayAuthConsumerResourceRequest) SetAcceptLanguage(v string) *ListGatewayAuthConsumerResourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceRequest) SetConsumerId(v int64) *ListGatewayAuthConsumerResourceRequest {
	s.ConsumerId = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceRequest) SetGatewayUniqueId(v string) *ListGatewayAuthConsumerResourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceRequest) SetPageNum(v string) *ListGatewayAuthConsumerResourceRequest {
	s.PageNum = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceRequest) SetPageSize(v string) *ListGatewayAuthConsumerResourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceRequest) SetResourceStatus(v bool) *ListGatewayAuthConsumerResourceRequest {
	s.ResourceStatus = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceRequest) SetRouteName(v string) *ListGatewayAuthConsumerResourceRequest {
	s.RouteName = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceRequest) Validate() error {
	return dara.Validate(s)
}
