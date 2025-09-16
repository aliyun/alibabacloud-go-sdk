// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpIncomingRequestHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpIncomingRequestHeaderModificationRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody
	GetConfigType() *string
	SetRequestHeaderModification(v []*GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) *GetHttpIncomingRequestHeaderModificationRuleResponseBody
	GetRequestHeaderModification() []*GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification
	SetRequestId(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody
	GetRequestId() *string
	SetRule(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetHttpIncomingRequestHeaderModificationRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetHttpIncomingRequestHeaderModificationRuleResponseBody
	GetSiteVersion() *int32
}

type GetHttpIncomingRequestHeaderModificationRuleResponseBody struct {
	// example:
	//
	// 402411533580288
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// global
	ConfigType                *string                                                                              `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	RequestHeaderModification []*GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	// example:
	//
	// 5CC228B4-7A67-4016-9C9F-4A4133494A91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetHttpIncomingRequestHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpIncomingRequestHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) GetRequestHeaderModification() []*GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) SetConfigId(v int64) *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) SetConfigType(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) SetRequestHeaderModification(v []*GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.RequestHeaderModification = v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) SetRequestId(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) SetRule(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) SetRuleEnable(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) SetRuleName(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) SetSequence(v int32) *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) SetSiteVersion(v int32) *GetHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification struct {
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// headValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GetType() *string {
	return s.Type
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) SetName(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) SetOperation(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) SetType(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	s.Type = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) SetValue(v string) *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *GetHttpIncomingRequestHeaderModificationRuleResponseBodyRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}
