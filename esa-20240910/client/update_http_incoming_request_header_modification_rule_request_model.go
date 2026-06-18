// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpIncomingRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetConfigId() *int64
	SetRequestHeaderModification(v []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetRequestHeaderModification() []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification
	SetRule(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetSiteId() *int64
}

type UpdateHttpIncomingRequestHeaderModificationRuleRequest struct {
	// The ID of the configuration. To obtain this ID, call the ListHttpIncomingRequestHeaderModificationRules API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 419717024278528
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// An array of objects that defines the request header modifications. Supported operations include `add`, `del`, and `modify`.
	RequestHeaderModification []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	// The rule expression, a conditional expression that matches user requests. This parameter is not required for a global configuration. You can use this parameter in two ways:
	//
	// - To match all incoming requests, set this value to `true`.
	//
	// - To match specific requests, provide a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The status of the rule. This parameter is not required for a global configuration. Valid values:
	//
	// - `on`: The rule is enabled.
	//
	// - `off`: The rule is disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. This parameter is not required for a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies the rule\\"s priority. Rules with a lower value are executed first.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The ID of the site. To obtain this ID, call the [ListSites](~~ListSites~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 568181195163328
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetRequestHeaderModification() []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetConfigId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetRequestHeaderModification(v []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.RequestHeaderModification = v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetRule(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetRuleEnable(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetRuleName(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetSequence(v int32) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetSiteId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) Validate() error {
	if s.RequestHeaderModification != nil {
		for _, item := range s.RequestHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification struct {
	// The name of the request header.
	//
	// This parameter is required.
	//
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation to perform on the request header. Valid values:
	//
	// - `add`: Adds a request header.
	//
	// - `del`: Deletes a request header.
	//
	// - `modify`: Modifies an existing request header.
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The type of value. Valid values:
	//
	// - `static`: The value is a fixed, literal string.
	//
	// - `dynamic`: The value is generated dynamically at runtime.
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the request header.
	//
	// example:
	//
	// headerValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetType() *string {
	return s.Type
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetName(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetOperation(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetType(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Type = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetValue(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}
