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
	// The client token that is used to ensure the idempotence of the request. You can generate a value for this parameter on the client side and ensure that the value is unique across different requests.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// **[Deprecated]*	- This field is deprecated.
	//
	// example:
	//
	// --
	DeliveryMode *string `json:"DeliveryMode,omitempty" xml:"DeliveryMode,omitempty"`
	// The endpoint that receives messages for this subscription.
	EndpointShrink *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// **[Deprecated]*	- This field is deprecated. Use the Endpoint parameter instead.
	EndpointsShrink *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// The list of event types. Valid values depend on the selected cloud service. For example, if OSS is selected, common values include ObjectCreated:PutObject and ObjectRemoved:DeleteObject. For more information, see the event notification documentation for each cloud service.
	//
	// This parameter is required.
	EventTypesShrink *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
	// The list of matching rules. A maximum of 10 rules can be specified at the outer level. The rules are connected by the OR operator.
	//
	// This parameter is required.
	MatchRulesShrink *string `json:"MatchRules,omitempty" xml:"MatchRules,omitempty"`
	// The name of the cloud service for event notification. Currently, only OSS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The name of the event notification rule. This parameter is required.
	//
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
