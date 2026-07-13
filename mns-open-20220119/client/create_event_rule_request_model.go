// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEventRuleRequest
	GetClientToken() *string
	SetDeliveryMode(v string) *CreateEventRuleRequest
	GetDeliveryMode() *string
	SetEndpoint(v *CreateEventRuleRequestEndpoint) *CreateEventRuleRequest
	GetEndpoint() *CreateEventRuleRequestEndpoint
	SetEndpoints(v []*CreateEventRuleRequestEndpoints) *CreateEventRuleRequest
	GetEndpoints() []*CreateEventRuleRequestEndpoints
	SetEventTypes(v []*string) *CreateEventRuleRequest
	GetEventTypes() []*string
	SetMatchRules(v [][]*EventMatchRule) *CreateEventRuleRequest
	GetMatchRules() [][]*EventMatchRule
	SetProductName(v string) *CreateEventRuleRequest
	GetProductName() *string
	SetRuleName(v string) *CreateEventRuleRequest
	GetRuleName() *string
}

type CreateEventRuleRequest struct {
	// A client token to ensure the idempotence of the request.
	//
	// Generate a unique value for this parameter from your client for each request.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// --
	DeliveryMode *string `json:"DeliveryMode,omitempty" xml:"DeliveryMode,omitempty"`
	// The endpoint that receives messages for this subscription.
	Endpoint *CreateEventRuleRequestEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Struct"`
	// This parameter is deprecated. Use Endpoint instead.
	Endpoints []*CreateEventRuleRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// A list of event types.
	//
	// This parameter is required.
	EventTypes []*string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// A list of matching rules. The logical relationship between the rules is OR.
	//
	// This parameter is required.
	MatchRules [][]*EventMatchRule `json:"MatchRules,omitempty" xml:"MatchRules,omitempty" type:"Repeated"`
	// The name of the Alibaba Cloud product for which you want to receive event notifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-xsXDW
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateEventRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateEventRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEventRuleRequest) GetDeliveryMode() *string {
	return s.DeliveryMode
}

func (s *CreateEventRuleRequest) GetEndpoint() *CreateEventRuleRequestEndpoint {
	return s.Endpoint
}

func (s *CreateEventRuleRequest) GetEndpoints() []*CreateEventRuleRequestEndpoints {
	return s.Endpoints
}

func (s *CreateEventRuleRequest) GetEventTypes() []*string {
	return s.EventTypes
}

func (s *CreateEventRuleRequest) GetMatchRules() [][]*EventMatchRule {
	return s.MatchRules
}

func (s *CreateEventRuleRequest) GetProductName() *string {
	return s.ProductName
}

func (s *CreateEventRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateEventRuleRequest) SetClientToken(v string) *CreateEventRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEventRuleRequest) SetDeliveryMode(v string) *CreateEventRuleRequest {
	s.DeliveryMode = &v
	return s
}

func (s *CreateEventRuleRequest) SetEndpoint(v *CreateEventRuleRequestEndpoint) *CreateEventRuleRequest {
	s.Endpoint = v
	return s
}

func (s *CreateEventRuleRequest) SetEndpoints(v []*CreateEventRuleRequestEndpoints) *CreateEventRuleRequest {
	s.Endpoints = v
	return s
}

func (s *CreateEventRuleRequest) SetEventTypes(v []*string) *CreateEventRuleRequest {
	s.EventTypes = v
	return s
}

func (s *CreateEventRuleRequest) SetMatchRules(v [][]*EventMatchRule) *CreateEventRuleRequest {
	s.MatchRules = v
	return s
}

func (s *CreateEventRuleRequest) SetProductName(v string) *CreateEventRuleRequest {
	s.ProductName = &v
	return s
}

func (s *CreateEventRuleRequest) SetRuleName(v string) *CreateEventRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateEventRuleRequest) Validate() error {
	if s.Endpoint != nil {
		if err := s.Endpoint.Validate(); err != nil {
			return err
		}
	}
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEventRuleRequestEndpoint struct {
	// The endpoint type. Valid values:
	//
	// - **topic**: The endpoint is a topic. A topic can deliver messages to multiple subscribers. You can add or remove subscribers later.
	//
	// - **queue**: The endpoint is a queue. Messages are delivered directly to the queue. This simplifies the delivery path, but you cannot add new subscribers later.
	//
	// example:
	//
	// topic
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The value of the endpoint.
	//
	// example:
	//
	// test-topic
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
}

func (s CreateEventRuleRequestEndpoint) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRuleRequestEndpoint) GoString() string {
	return s.String()
}

func (s *CreateEventRuleRequestEndpoint) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateEventRuleRequestEndpoint) GetEndpointValue() *string {
	return s.EndpointValue
}

func (s *CreateEventRuleRequestEndpoint) SetEndpointType(v string) *CreateEventRuleRequestEndpoint {
	s.EndpointType = &v
	return s
}

func (s *CreateEventRuleRequestEndpoint) SetEndpointValue(v string) *CreateEventRuleRequestEndpoint {
	s.EndpointValue = &v
	return s
}

func (s *CreateEventRuleRequestEndpoint) Validate() error {
	return dara.Validate(s)
}

type CreateEventRuleRequestEndpoints struct {
	// Deprecated. Use Endpoint.EndpointType instead.
	//
	// example:
	//
	// http
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// Deprecated. Use Endpoint.EndpointValue instead.
	//
	// example:
	//
	// test-xxx-queue
	EndpointValue *string `json:"EndpointValue,omitempty" xml:"EndpointValue,omitempty"`
}

func (s CreateEventRuleRequestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRuleRequestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateEventRuleRequestEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateEventRuleRequestEndpoints) GetEndpointValue() *string {
	return s.EndpointValue
}

func (s *CreateEventRuleRequestEndpoints) SetEndpointType(v string) *CreateEventRuleRequestEndpoints {
	s.EndpointType = &v
	return s
}

func (s *CreateEventRuleRequestEndpoints) SetEndpointValue(v string) *CreateEventRuleRequestEndpoints {
	s.EndpointValue = &v
	return s
}

func (s *CreateEventRuleRequestEndpoints) Validate() error {
	return dara.Validate(s)
}
