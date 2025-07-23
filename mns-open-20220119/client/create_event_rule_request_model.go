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
	ClientToken  *string                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeliveryMode *string                            `json:"DeliveryMode,omitempty" xml:"DeliveryMode,omitempty"`
	Endpoint     *CreateEventRuleRequestEndpoint    `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Struct"`
	Endpoints    []*CreateEventRuleRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// This parameter is required.
	EventTypes []*string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	MatchRules [][]*EventMatchRule `json:"MatchRules,omitempty" xml:"MatchRules,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
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
	return dara.Validate(s)
}

type CreateEventRuleRequestEndpoint struct {
	EndpointType  *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
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
	// example:
	//
	// http
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
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
