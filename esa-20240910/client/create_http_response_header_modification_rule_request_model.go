// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpResponseHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResponseHeaderModification(v []*CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) *CreateHttpResponseHeaderModificationRuleRequest
	GetResponseHeaderModification() []*CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification
	SetRule(v string) *CreateHttpResponseHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpResponseHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpResponseHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpResponseHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpResponseHeaderModificationRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpResponseHeaderModificationRuleRequest
	GetSiteVersion() *int32
}

type CreateHttpResponseHeaderModificationRuleRequest struct {
	// The response header modifications. Three operation types are supported: add, delete, and modify.
	//
	// This parameter is required.
	ResponseHeaderModification []*CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
	// The rule content, which uses conditional expressions to match user requests. This parameter is not required when you add a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specified requests: Set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. This parameter is not required when you add a global configuration. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is not required when you add a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to query the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version on which the configuration takes effect. The default value is 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpResponseHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpResponseHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) GetResponseHeaderModification() []*CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) SetResponseHeaderModification(v []*CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) *CreateHttpResponseHeaderModificationRuleRequest {
	s.ResponseHeaderModification = v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) SetRule(v string) *CreateHttpResponseHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) SetRuleEnable(v string) *CreateHttpResponseHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) SetRuleName(v string) *CreateHttpResponseHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) SetSequence(v int32) *CreateHttpResponseHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) SetSiteId(v int64) *CreateHttpResponseHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) SetSiteVersion(v int32) *CreateHttpResponseHeaderModificationRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequest) Validate() error {
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

type CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification struct {
	// The response header name.
	//
	// This parameter is required.
	//
	// example:
	//
	// headername
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type. Valid values:
	//
	// - add: adds a response header.
	//
	// - del: deletes a response header.
	//
	// - modify: modifies a response header.
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The value type. Valid values:
	//
	// - static: static pattern.
	//
	// - dynamic: dynamic schema.
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The response header value.
	//
	// example:
	//
	// headervalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GetType() *string {
	return s.Type
}

func (s *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) SetName(v string) *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) SetOperation(v string) *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) SetType(v string) *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Type = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) SetValue(v string) *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}
