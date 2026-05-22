// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpIncomingResponseHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResponseHeaderModification(v []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetResponseHeaderModification() []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification
	SetRule(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetSiteVersion() *int32
}

type CreateHttpIncomingResponseHeaderModificationRuleRequest struct {
	// The configurations of modifying response headers. You can add, delete, or modify a response header.
	//
	// This parameter is required.
	ResponseHeaderModification []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
	// The content of the rule. A conditional expression is used to match a user request. You do not need to set this parameter when you add global configuration. Use cases:
	//
	// 	- true: Match all incoming requests.
	//
	// 	- Set the value to a custom expression, for example: (http.host eq "video.example.com"): Match the specified request
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. Valid values: You do not need to set this parameter when you add global configuration. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when you add global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The order in which the rule is executed. A smaller value gives priority to the rule.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 608665779308176
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the website configurations. You can use this parameter to specify a version of your website to apply the feature settings. By default, version 0 is used.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpIncomingResponseHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingResponseHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetResponseHeaderModification() []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetResponseHeaderModification(v []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.ResponseHeaderModification = v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetRule(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetRuleEnable(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetRuleName(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetSequence(v int32) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetSiteId(v int64) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetSiteVersion(v int32) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) Validate() error {
	if s.ResponseHeaderModification != nil {
		for _, item := range s.ResponseHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification struct {
	// The name of the response header.
	//
	// This parameter is required.
	//
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The action. Valid values:
	//
	// 	- add: adds a response header.
	//
	// 	- del: deletes a response header.
	//
	// 	- modify: modifies a response header.
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The type of the value. Valid values:
	//
	// 	- static
	//
	// 	- dynamic
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the response header.
	//
	// example:
	//
	// headerValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetType() *string {
	return s.Type
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetName(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetOperation(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetType(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Type = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetValue(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}
