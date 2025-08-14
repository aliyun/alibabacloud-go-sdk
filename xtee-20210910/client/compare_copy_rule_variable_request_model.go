// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareCopyRuleVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateType(v string) *CompareCopyRuleVariableRequest
	GetCreateType() *string
	SetLang(v string) *CompareCopyRuleVariableRequest
	GetLang() *string
	SetRegId(v string) *CompareCopyRuleVariableRequest
	GetRegId() *string
	SetSourceRuleId(v string) *CompareCopyRuleVariableRequest
	GetSourceRuleId() *string
	SetSourceRuleIds(v string) *CompareCopyRuleVariableRequest
	GetSourceRuleIds() *string
	SetTargetEventCode(v string) *CompareCopyRuleVariableRequest
	GetTargetEventCode() *string
}

type CompareCopyRuleVariableRequest struct {
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Original policy ruleId.
	//
	// example:
	//
	// 102125
	SourceRuleId *string `json:"SourceRuleId,omitempty" xml:"SourceRuleId,omitempty"`
	// Original policy ruleIds.
	//
	// example:
	//
	// 102125,102129
	SourceRuleIds *string `json:"SourceRuleIds,omitempty" xml:"SourceRuleIds,omitempty"`
	// Target event eventCode.
	//
	// example:
	//
	// de_ajtshf1581
	TargetEventCode *string `json:"TargetEventCode,omitempty" xml:"TargetEventCode,omitempty"`
}

func (s CompareCopyRuleVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareCopyRuleVariableRequest) GoString() string {
	return s.String()
}

func (s *CompareCopyRuleVariableRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *CompareCopyRuleVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *CompareCopyRuleVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *CompareCopyRuleVariableRequest) GetSourceRuleId() *string {
	return s.SourceRuleId
}

func (s *CompareCopyRuleVariableRequest) GetSourceRuleIds() *string {
	return s.SourceRuleIds
}

func (s *CompareCopyRuleVariableRequest) GetTargetEventCode() *string {
	return s.TargetEventCode
}

func (s *CompareCopyRuleVariableRequest) SetCreateType(v string) *CompareCopyRuleVariableRequest {
	s.CreateType = &v
	return s
}

func (s *CompareCopyRuleVariableRequest) SetLang(v string) *CompareCopyRuleVariableRequest {
	s.Lang = &v
	return s
}

func (s *CompareCopyRuleVariableRequest) SetRegId(v string) *CompareCopyRuleVariableRequest {
	s.RegId = &v
	return s
}

func (s *CompareCopyRuleVariableRequest) SetSourceRuleId(v string) *CompareCopyRuleVariableRequest {
	s.SourceRuleId = &v
	return s
}

func (s *CompareCopyRuleVariableRequest) SetSourceRuleIds(v string) *CompareCopyRuleVariableRequest {
	s.SourceRuleIds = &v
	return s
}

func (s *CompareCopyRuleVariableRequest) SetTargetEventCode(v string) *CompareCopyRuleVariableRequest {
	s.TargetEventCode = &v
	return s
}

func (s *CompareCopyRuleVariableRequest) Validate() error {
	return dara.Validate(s)
}
