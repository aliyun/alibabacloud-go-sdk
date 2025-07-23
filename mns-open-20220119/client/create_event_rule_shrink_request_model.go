// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEventRuleShrinkRequest
	GetClientToken() *string
	SetDeliveryMode(v string) *CreateEventRuleShrinkRequest
	GetDeliveryMode() *string
	SetEndpointShrink(v string) *CreateEventRuleShrinkRequest
	GetEndpointShrink() *string
	SetEndpointsShrink(v string) *CreateEventRuleShrinkRequest
	GetEndpointsShrink() *string
	SetEventTypesShrink(v string) *CreateEventRuleShrinkRequest
	GetEventTypesShrink() *string
	SetMatchRulesShrink(v string) *CreateEventRuleShrinkRequest
	GetMatchRulesShrink() *string
	SetProductName(v string) *CreateEventRuleShrinkRequest
	GetProductName() *string
	SetRuleName(v string) *CreateEventRuleShrinkRequest
	GetRuleName() *string
}

type CreateEventRuleShrinkRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeliveryMode    *string `json:"DeliveryMode,omitempty" xml:"DeliveryMode,omitempty"`
	EndpointShrink  *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EndpointsShrink *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// This parameter is required.
	EventTypesShrink *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
	// This parameter is required.
	MatchRulesShrink *string `json:"MatchRules,omitempty" xml:"MatchRules,omitempty"`
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

func (s CreateEventRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEventRuleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEventRuleShrinkRequest) GetDeliveryMode() *string {
	return s.DeliveryMode
}

func (s *CreateEventRuleShrinkRequest) GetEndpointShrink() *string {
	return s.EndpointShrink
}

func (s *CreateEventRuleShrinkRequest) GetEndpointsShrink() *string {
	return s.EndpointsShrink
}

func (s *CreateEventRuleShrinkRequest) GetEventTypesShrink() *string {
	return s.EventTypesShrink
}

func (s *CreateEventRuleShrinkRequest) GetMatchRulesShrink() *string {
	return s.MatchRulesShrink
}

func (s *CreateEventRuleShrinkRequest) GetProductName() *string {
	return s.ProductName
}

func (s *CreateEventRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateEventRuleShrinkRequest) SetClientToken(v string) *CreateEventRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetDeliveryMode(v string) *CreateEventRuleShrinkRequest {
	s.DeliveryMode = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetEndpointShrink(v string) *CreateEventRuleShrinkRequest {
	s.EndpointShrink = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetEndpointsShrink(v string) *CreateEventRuleShrinkRequest {
	s.EndpointsShrink = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetEventTypesShrink(v string) *CreateEventRuleShrinkRequest {
	s.EventTypesShrink = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetMatchRulesShrink(v string) *CreateEventRuleShrinkRequest {
	s.MatchRulesShrink = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetProductName(v string) *CreateEventRuleShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) SetRuleName(v string) *CreateEventRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateEventRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
