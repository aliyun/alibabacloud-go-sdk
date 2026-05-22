// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWafBatchRuleShared interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *WafBatchRuleShared
	GetAction() *string
	SetActions(v *WafBatchRuleSharedActions) *WafBatchRuleShared
	GetActions() *WafBatchRuleSharedActions
	SetCrossSiteId(v int64) *WafBatchRuleShared
	GetCrossSiteId() *int64
	SetExpression(v string) *WafBatchRuleShared
	GetExpression() *string
	SetMatch(v *WafRuleMatch2) *WafBatchRuleShared
	GetMatch() *WafRuleMatch2
	SetMode(v string) *WafBatchRuleShared
	GetMode() *string
	SetName(v string) *WafBatchRuleShared
	GetName() *string
	SetTarget(v string) *WafBatchRuleShared
	GetTarget() *string
}

type WafBatchRuleShared struct {
	// The action that you want WAF to perform on requests that match the rule.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The extended action configurations.
	Actions *WafBatchRuleSharedActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Struct"`
	// Specifies the cross-domain website ID.
	//
	// example:
	//
	// 10000001
	CrossSiteId *int64 `json:"CrossSiteId,omitempty" xml:"CrossSiteId,omitempty"`
	// The expression.
	//
	// example:
	//
	// ip.src eq 1.1.1.1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The matching rule.
	Match *WafRuleMatch2 `json:"Match,omitempty" xml:"Match,omitempty"`
	// Web SDK integration method: automatic integration (automatic) or manual integration (manual).
	//
	// example:
	//
	// automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ruleset name.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Protection type: web or app.
	//
	// example:
	//
	// web
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s WafBatchRuleShared) String() string {
	return dara.Prettify(s)
}

func (s WafBatchRuleShared) GoString() string {
	return s.String()
}

func (s *WafBatchRuleShared) GetAction() *string {
	return s.Action
}

func (s *WafBatchRuleShared) GetActions() *WafBatchRuleSharedActions {
	return s.Actions
}

func (s *WafBatchRuleShared) GetCrossSiteId() *int64 {
	return s.CrossSiteId
}

func (s *WafBatchRuleShared) GetExpression() *string {
	return s.Expression
}

func (s *WafBatchRuleShared) GetMatch() *WafRuleMatch2 {
	return s.Match
}

func (s *WafBatchRuleShared) GetMode() *string {
	return s.Mode
}

func (s *WafBatchRuleShared) GetName() *string {
	return s.Name
}

func (s *WafBatchRuleShared) GetTarget() *string {
	return s.Target
}

func (s *WafBatchRuleShared) SetAction(v string) *WafBatchRuleShared {
	s.Action = &v
	return s
}

func (s *WafBatchRuleShared) SetActions(v *WafBatchRuleSharedActions) *WafBatchRuleShared {
	s.Actions = v
	return s
}

func (s *WafBatchRuleShared) SetCrossSiteId(v int64) *WafBatchRuleShared {
	s.CrossSiteId = &v
	return s
}

func (s *WafBatchRuleShared) SetExpression(v string) *WafBatchRuleShared {
	s.Expression = &v
	return s
}

func (s *WafBatchRuleShared) SetMatch(v *WafRuleMatch2) *WafBatchRuleShared {
	s.Match = v
	return s
}

func (s *WafBatchRuleShared) SetMode(v string) *WafBatchRuleShared {
	s.Mode = &v
	return s
}

func (s *WafBatchRuleShared) SetName(v string) *WafBatchRuleShared {
	s.Name = &v
	return s
}

func (s *WafBatchRuleShared) SetTarget(v string) *WafBatchRuleShared {
	s.Target = &v
	return s
}

func (s *WafBatchRuleShared) Validate() error {
	if s.Actions != nil {
		if err := s.Actions.Validate(); err != nil {
			return err
		}
	}
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WafBatchRuleSharedActions struct {
	// The custom response.
	Response *WafBatchRuleSharedActionsResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
}

func (s WafBatchRuleSharedActions) String() string {
	return dara.Prettify(s)
}

func (s WafBatchRuleSharedActions) GoString() string {
	return s.String()
}

func (s *WafBatchRuleSharedActions) GetResponse() *WafBatchRuleSharedActionsResponse {
	return s.Response
}

func (s *WafBatchRuleSharedActions) SetResponse(v *WafBatchRuleSharedActionsResponse) *WafBatchRuleSharedActions {
	s.Response = v
	return s
}

func (s *WafBatchRuleSharedActions) Validate() error {
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WafBatchRuleSharedActionsResponse struct {
	// The custom response code.
	//
	// example:
	//
	// 403
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the custom response page.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s WafBatchRuleSharedActionsResponse) String() string {
	return dara.Prettify(s)
}

func (s WafBatchRuleSharedActionsResponse) GoString() string {
	return s.String()
}

func (s *WafBatchRuleSharedActionsResponse) GetCode() *int32 {
	return s.Code
}

func (s *WafBatchRuleSharedActionsResponse) GetId() *int64 {
	return s.Id
}

func (s *WafBatchRuleSharedActionsResponse) SetCode(v int32) *WafBatchRuleSharedActionsResponse {
	s.Code = &v
	return s
}

func (s *WafBatchRuleSharedActionsResponse) SetId(v int64) *WafBatchRuleSharedActionsResponse {
	s.Id = &v
	return s
}

func (s *WafBatchRuleSharedActionsResponse) Validate() error {
	return dara.Validate(s)
}
