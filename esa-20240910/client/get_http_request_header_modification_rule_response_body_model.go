// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpRequestHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpRequestHeaderModificationRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetHttpRequestHeaderModificationRuleResponseBody
	GetConfigType() *string
	SetRequestHeaderModification(v []*GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) *GetHttpRequestHeaderModificationRuleResponseBody
	GetRequestHeaderModification() []*GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification
	SetRequestId(v string) *GetHttpRequestHeaderModificationRuleResponseBody
	GetRequestId() *string
	SetRule(v string) *GetHttpRequestHeaderModificationRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetHttpRequestHeaderModificationRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetHttpRequestHeaderModificationRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetHttpRequestHeaderModificationRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetHttpRequestHeaderModificationRuleResponseBody
	GetSiteVersion() *int32
}

type GetHttpRequestHeaderModificationRuleResponseBody struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. Valid values:
	//
	// - global: Global configuration.
	//
	// - rule: Rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The request header modification settings. Add, delete, and modify operations are supported.
	RequestHeaderModification []*GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specific requests: Set the value to a custom expression, such as (http.host eq "video.example.com").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The name of the rule. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution sequence of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetHttpRequestHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpRequestHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) GetRequestHeaderModification() []*GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) SetConfigId(v int64) *GetHttpRequestHeaderModificationRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) SetConfigType(v string) *GetHttpRequestHeaderModificationRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) SetRequestHeaderModification(v []*GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) *GetHttpRequestHeaderModificationRuleResponseBody {
	s.RequestHeaderModification = v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) SetRequestId(v string) *GetHttpRequestHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) SetRule(v string) *GetHttpRequestHeaderModificationRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) SetRuleEnable(v string) *GetHttpRequestHeaderModificationRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) SetRuleName(v string) *GetHttpRequestHeaderModificationRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) SetSequence(v int32) *GetHttpRequestHeaderModificationRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) SetSiteVersion(v int32) *GetHttpRequestHeaderModificationRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBody) Validate() error {
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

type GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification struct {
	// The name of the request header.
	//
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type. Valid values:
	//
	// - add: Add.
	//
	// - del: Delete.
	//
	// - modify: Modify.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The value type. Valid values:
	//
	// - static: Static pattern.
	//
	// - dynamic: Dynamic schema.
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the request header.
	//
	// example:
	//
	// headValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GetType() *string {
	return s.Type
}

func (s *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) SetName(v string) *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) SetOperation(v string) *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) SetType(v string) *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	s.Type = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) SetValue(v string) *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}
