// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCopyRuleVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateType(v string) *CheckCopyRuleVariableRequest
	GetCreateType() *string
	SetLang(v string) *CheckCopyRuleVariableRequest
	GetLang() *string
	SetRegId(v string) *CheckCopyRuleVariableRequest
	GetRegId() *string
	SetSourceRuleId(v string) *CheckCopyRuleVariableRequest
	GetSourceRuleId() *string
	SetSourceRuleIds(v string) *CheckCopyRuleVariableRequest
	GetSourceRuleIds() *string
	SetTargetEventCode(v string) *CheckCopyRuleVariableRequest
	GetTargetEventCode() *string
}

type CheckCopyRuleVariableRequest struct {
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
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Source policy ID
	//
	// example:
	//
	// 102125
	SourceRuleId *string `json:"SourceRuleId,omitempty" xml:"SourceRuleId,omitempty"`
	// Source policy IDs
	//
	// example:
	//
	// 02125,102129
	SourceRuleIds *string `json:"SourceRuleIds,omitempty" xml:"SourceRuleIds,omitempty"`
	// Target event
	//
	// example:
	//
	// de_ajtshf1581
	TargetEventCode *string `json:"TargetEventCode,omitempty" xml:"TargetEventCode,omitempty"`
}

func (s CheckCopyRuleVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCopyRuleVariableRequest) GoString() string {
	return s.String()
}

func (s *CheckCopyRuleVariableRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *CheckCopyRuleVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckCopyRuleVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *CheckCopyRuleVariableRequest) GetSourceRuleId() *string {
	return s.SourceRuleId
}

func (s *CheckCopyRuleVariableRequest) GetSourceRuleIds() *string {
	return s.SourceRuleIds
}

func (s *CheckCopyRuleVariableRequest) GetTargetEventCode() *string {
	return s.TargetEventCode
}

func (s *CheckCopyRuleVariableRequest) SetCreateType(v string) *CheckCopyRuleVariableRequest {
	s.CreateType = &v
	return s
}

func (s *CheckCopyRuleVariableRequest) SetLang(v string) *CheckCopyRuleVariableRequest {
	s.Lang = &v
	return s
}

func (s *CheckCopyRuleVariableRequest) SetRegId(v string) *CheckCopyRuleVariableRequest {
	s.RegId = &v
	return s
}

func (s *CheckCopyRuleVariableRequest) SetSourceRuleId(v string) *CheckCopyRuleVariableRequest {
	s.SourceRuleId = &v
	return s
}

func (s *CheckCopyRuleVariableRequest) SetSourceRuleIds(v string) *CheckCopyRuleVariableRequest {
	s.SourceRuleIds = &v
	return s
}

func (s *CheckCopyRuleVariableRequest) SetTargetEventCode(v string) *CheckCopyRuleVariableRequest {
	s.TargetEventCode = &v
	return s
}

func (s *CheckCopyRuleVariableRequest) Validate() error {
	return dara.Validate(s)
}
