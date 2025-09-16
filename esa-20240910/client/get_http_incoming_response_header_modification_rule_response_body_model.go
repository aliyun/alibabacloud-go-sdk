// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpIncomingResponseHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpIncomingResponseHeaderModificationRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody
	GetConfigType() *string
	SetRequestId(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody
	GetRequestId() *string
	SetResponseHeaderModification(v []*GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) *GetHttpIncomingResponseHeaderModificationRuleResponseBody
	GetResponseHeaderModification() []*GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification
	SetRule(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetHttpIncomingResponseHeaderModificationRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetHttpIncomingResponseHeaderModificationRuleResponseBody
	GetSiteVersion() *int32
}

type GetHttpIncomingResponseHeaderModificationRuleResponseBody struct {
	// example:
	//
	// 432637955352576
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId                  *string                                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseHeaderModification []*GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
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

func (s GetHttpIncomingResponseHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpIncomingResponseHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) GetResponseHeaderModification() []*GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) SetConfigId(v int64) *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) SetConfigType(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) SetRequestId(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) SetResponseHeaderModification(v []*GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.ResponseHeaderModification = v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) SetRule(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) SetRuleEnable(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) SetRuleName(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) SetSequence(v int32) *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) SetSiteVersion(v int32) *GetHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification struct {
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
	// headerValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GetType() *string {
	return s.Type
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) SetName(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) SetOperation(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) SetType(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	s.Type = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) SetValue(v string) *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *GetHttpIncomingResponseHeaderModificationRuleResponseBodyResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}
