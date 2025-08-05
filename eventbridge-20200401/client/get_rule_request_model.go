// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventBusName(v string) *GetRuleRequest
	GetEventBusName() *string
	SetRuleName(v string) *GetRuleRequest
	GetRuleName() *string
}

type GetRuleRequest struct {
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ram-changes
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRuleRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *GetRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *GetRuleRequest) SetEventBusName(v string) *GetRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *GetRuleRequest) SetRuleName(v string) *GetRuleRequest {
	s.RuleName = &v
	return s
}

func (s *GetRuleRequest) Validate() error {
	return dara.Validate(s)
}
