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
	ConfigId                  *int64                                                                       `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType                *string                                                                      `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	RequestHeaderModification []*GetHttpRequestHeaderModificationRuleResponseBodyRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	RequestId                 *string                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rule                      *string                                                                      `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                *string                                                                      `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                  *string                                                                      `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                  *int32                                                                       `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SiteVersion               *int32                                                                       `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
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
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
