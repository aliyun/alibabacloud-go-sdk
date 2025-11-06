// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerResourceRequest
	GetAcceptLanguage() *string
	SetConsumerId(v int64) *UpdateGatewayAuthConsumerResourceRequest
	GetConsumerId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerResourceRequest
	GetGatewayUniqueId() *string
	SetResourceList(v []*UpdateGatewayAuthConsumerResourceRequestResourceList) *UpdateGatewayAuthConsumerResourceRequest
	GetResourceList() []*UpdateGatewayAuthConsumerResourceRequestResourceList
}

type UpdateGatewayAuthConsumerResourceRequest struct {
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
	ResourceList []*UpdateGatewayAuthConsumerResourceRequestResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
}

func (s UpdateGatewayAuthConsumerResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayAuthConsumerResourceRequest) GetConsumerId() *int64 {
	return s.ConsumerId
}

func (s *UpdateGatewayAuthConsumerResourceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayAuthConsumerResourceRequest) GetResourceList() []*UpdateGatewayAuthConsumerResourceRequestResourceList {
	return s.ResourceList
}

func (s *UpdateGatewayAuthConsumerResourceRequest) SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerResourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceRequest) SetConsumerId(v int64) *UpdateGatewayAuthConsumerResourceRequest {
	s.ConsumerId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceRequest) SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerResourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceRequest) SetResourceList(v []*UpdateGatewayAuthConsumerResourceRequestResourceList) *UpdateGatewayAuthConsumerResourceRequest {
	s.ResourceList = v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceRequest) Validate() error {
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateGatewayAuthConsumerResourceRequestResourceList struct {
	// The route ID.
	//
	// example:
	//
	// 7816
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// helo
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
}

func (s UpdateGatewayAuthConsumerResourceRequestResourceList) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResourceRequestResourceList) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResourceRequestResourceList) GetRouteId() *int64 {
	return s.RouteId
}

func (s *UpdateGatewayAuthConsumerResourceRequestResourceList) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateGatewayAuthConsumerResourceRequestResourceList) SetRouteId(v int64) *UpdateGatewayAuthConsumerResourceRequestResourceList {
	s.RouteId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceRequestResourceList) SetRouteName(v string) *UpdateGatewayAuthConsumerResourceRequestResourceList {
	s.RouteName = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceRequestResourceList) Validate() error {
	return dara.Validate(s)
}
