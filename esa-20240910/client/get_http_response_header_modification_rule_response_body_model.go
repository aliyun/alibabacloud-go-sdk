// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpResponseHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpResponseHeaderModificationRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetHttpResponseHeaderModificationRuleResponseBody
	GetConfigType() *string
	SetRequestId(v string) *GetHttpResponseHeaderModificationRuleResponseBody
	GetRequestId() *string
	SetResponseHeaderModification(v []*GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) *GetHttpResponseHeaderModificationRuleResponseBody
	GetResponseHeaderModification() []*GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification
	SetRule(v string) *GetHttpResponseHeaderModificationRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetHttpResponseHeaderModificationRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetHttpResponseHeaderModificationRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetHttpResponseHeaderModificationRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetHttpResponseHeaderModificationRuleResponseBody
	GetSiteVersion() *int32
}

type GetHttpResponseHeaderModificationRuleResponseBody struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, with the following values:
	//
	// - global: Global configuration.
	//
	// - rule: Rule-based configuration.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Modify response headers, supporting add, delete, and modify operations.
	ResponseHeaderModification []*GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Possible values are:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites that have enabled configuration version management, you can use this parameter to specify the effective version of the site configuration, defaulting to version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetHttpResponseHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpResponseHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) GetResponseHeaderModification() []*GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) SetConfigId(v int64) *GetHttpResponseHeaderModificationRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) SetConfigType(v string) *GetHttpResponseHeaderModificationRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) SetRequestId(v string) *GetHttpResponseHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) SetResponseHeaderModification(v []*GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) *GetHttpResponseHeaderModificationRuleResponseBody {
	s.ResponseHeaderModification = v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) SetRule(v string) *GetHttpResponseHeaderModificationRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) SetRuleEnable(v string) *GetHttpResponseHeaderModificationRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) SetRuleName(v string) *GetHttpResponseHeaderModificationRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) SetSequence(v int32) *GetHttpResponseHeaderModificationRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) SetSiteVersion(v int32) *GetHttpResponseHeaderModificationRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBody) Validate() error {
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

type GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification struct {
	// Response header name.
	//
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Operation method. Possible values are:
	//
	// - add: Add.
	//
	// - del: Delete
	//
	// - modify: Modify.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Response header value.
	//
	// example:
	//
	// headerValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GetType() *string {
	return s.Type
}

func (s *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) SetName(v string) *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) SetOperation(v string) *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) SetType(v string) *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	s.Type = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) SetValue(v string) *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponseBodyResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}
