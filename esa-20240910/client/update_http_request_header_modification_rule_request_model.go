// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateHttpRequestHeaderModificationRuleRequest
	GetConfigId() *int64
	SetRequestHeaderModification(v []*UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) *UpdateHttpRequestHeaderModificationRuleRequest
	GetRequestHeaderModification() []*UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification
	SetRule(v string) *UpdateHttpRequestHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpRequestHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpRequestHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpRequestHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpRequestHeaderModificationRuleRequest
	GetSiteId() *int64
}

type UpdateHttpRequestHeaderModificationRuleRequest struct {
	// The configuration ID. Call the [ListHttpRequestHeaderModificationRules](https://help.aliyun.com/document_detail/2867483.html) operation to obtain it.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies the modifications for the request header. Supported operations include `add`, `del`, and `modify`.
	RequestHeaderModification []*UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	// The Conditional Expression used to match User Requests. This parameter is not required for a Global Configuration. Use cases:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, use a custom expression, for example, `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the Rule is enabled. This parameter is not required for a Global Configuration. Valid values:
	//
	// - `on`: Enable
	//
	// - `off`: Disable
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the Rule. This parameter is not required for a Global Configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution priority of the Rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site ID. Call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain it.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) GetRequestHeaderModification() []*UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) SetConfigId(v int64) *UpdateHttpRequestHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) SetRequestHeaderModification(v []*UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) *UpdateHttpRequestHeaderModificationRuleRequest {
	s.RequestHeaderModification = v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) SetRule(v string) *UpdateHttpRequestHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) SetRuleEnable(v string) *UpdateHttpRequestHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) SetRuleName(v string) *UpdateHttpRequestHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) SetSequence(v int32) *UpdateHttpRequestHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) SetSiteId(v int64) *UpdateHttpRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequest) Validate() error {
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

type UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification struct {
	// The name of the Request Header.
	//
	// This parameter is required.
	//
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of Operation to perform. Valid values:
	//
	// - `add`: Add
	//
	// - `del`: Delete
	//
	// - `modify`: Modify
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The type of the header value. Valid values:
	//
	// - `static`: Static Mode
	//
	// - `dynamic`: Dynamic Mode
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the Request Header.
	//
	// example:
	//
	// headerValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GetType() *string {
	return s.Type
}

func (s *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) SetName(v string) *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) SetOperation(v string) *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) SetType(v string) *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Type = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) SetValue(v string) *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}
