// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepCopyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateType(v string) *DeepCopyRuleRequest
	GetCreateType() *string
	SetCustInsertInfo(v string) *DeepCopyRuleRequest
	GetCustInsertInfo() *string
	SetCustWriteInfo(v string) *DeepCopyRuleRequest
	GetCustWriteInfo() *string
	SetExpressionVariableInfo(v string) *DeepCopyRuleRequest
	GetExpressionVariableInfo() *string
	SetLang(v string) *DeepCopyRuleRequest
	GetLang() *string
	SetQueryExpressionVariableInfo(v string) *DeepCopyRuleRequest
	GetQueryExpressionVariableInfo() *string
	SetRegId(v string) *DeepCopyRuleRequest
	GetRegId() *string
	SetSourceRuleId(v string) *DeepCopyRuleRequest
	GetSourceRuleId() *string
	SetSourceRuleIds(v string) *DeepCopyRuleRequest
	GetSourceRuleIds() *string
	SetTargetEventCode(v string) *DeepCopyRuleRequest
	GetTargetEventCode() *string
	SetTargetEventName(v string) *DeepCopyRuleRequest
	GetTargetEventName() *string
}

type DeepCopyRuleRequest struct {
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	// Newly added cumulative variable
	//
	// example:
	//
	// [{"id":"1288","title":"新标题"}]
	CustInsertInfo *string `json:"CustInsertInfo,omitempty" xml:"CustInsertInfo,omitempty"`
	// Read cumulative variable
	//
	// example:
	//
	// [1234，2345]
	CustWriteInfo *string `json:"CustWriteInfo,omitempty" xml:"CustWriteInfo,omitempty"`
	// Custom variables to be added
	//
	// example:
	//
	// [{"id":"ex_2wxZPcxg3a03","title":"新标题"}]
	ExpressionVariableInfo *string `json:"ExpressionVariableInfo,omitempty" xml:"ExpressionVariableInfo,omitempty"`
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
	// Custom query variables to be added
	//
	// example:
	//
	// [{"id":"ex_2wxZPcxg3a03","title":"新标题"}]
	QueryExpressionVariableInfo *string `json:"QueryExpressionVariableInfo,omitempty" xml:"QueryExpressionVariableInfo,omitempty"`
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
	// Target policy ID
	//
	// example:
	//
	// 102125,102129
	SourceRuleIds *string `json:"SourceRuleIds,omitempty" xml:"SourceRuleIds,omitempty"`
	// Target event
	//
	// example:
	//
	// de_ajtshf1581
	TargetEventCode *string `json:"TargetEventCode,omitempty" xml:"TargetEventCode,omitempty"`
	// Target event name
	//
	// example:
	//
	// 目标事件名称
	TargetEventName *string `json:"TargetEventName,omitempty" xml:"TargetEventName,omitempty"`
}

func (s DeepCopyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeepCopyRuleRequest) GoString() string {
	return s.String()
}

func (s *DeepCopyRuleRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DeepCopyRuleRequest) GetCustInsertInfo() *string {
	return s.CustInsertInfo
}

func (s *DeepCopyRuleRequest) GetCustWriteInfo() *string {
	return s.CustWriteInfo
}

func (s *DeepCopyRuleRequest) GetExpressionVariableInfo() *string {
	return s.ExpressionVariableInfo
}

func (s *DeepCopyRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeepCopyRuleRequest) GetQueryExpressionVariableInfo() *string {
	return s.QueryExpressionVariableInfo
}

func (s *DeepCopyRuleRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeepCopyRuleRequest) GetSourceRuleId() *string {
	return s.SourceRuleId
}

func (s *DeepCopyRuleRequest) GetSourceRuleIds() *string {
	return s.SourceRuleIds
}

func (s *DeepCopyRuleRequest) GetTargetEventCode() *string {
	return s.TargetEventCode
}

func (s *DeepCopyRuleRequest) GetTargetEventName() *string {
	return s.TargetEventName
}

func (s *DeepCopyRuleRequest) SetCreateType(v string) *DeepCopyRuleRequest {
	s.CreateType = &v
	return s
}

func (s *DeepCopyRuleRequest) SetCustInsertInfo(v string) *DeepCopyRuleRequest {
	s.CustInsertInfo = &v
	return s
}

func (s *DeepCopyRuleRequest) SetCustWriteInfo(v string) *DeepCopyRuleRequest {
	s.CustWriteInfo = &v
	return s
}

func (s *DeepCopyRuleRequest) SetExpressionVariableInfo(v string) *DeepCopyRuleRequest {
	s.ExpressionVariableInfo = &v
	return s
}

func (s *DeepCopyRuleRequest) SetLang(v string) *DeepCopyRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeepCopyRuleRequest) SetQueryExpressionVariableInfo(v string) *DeepCopyRuleRequest {
	s.QueryExpressionVariableInfo = &v
	return s
}

func (s *DeepCopyRuleRequest) SetRegId(v string) *DeepCopyRuleRequest {
	s.RegId = &v
	return s
}

func (s *DeepCopyRuleRequest) SetSourceRuleId(v string) *DeepCopyRuleRequest {
	s.SourceRuleId = &v
	return s
}

func (s *DeepCopyRuleRequest) SetSourceRuleIds(v string) *DeepCopyRuleRequest {
	s.SourceRuleIds = &v
	return s
}

func (s *DeepCopyRuleRequest) SetTargetEventCode(v string) *DeepCopyRuleRequest {
	s.TargetEventCode = &v
	return s
}

func (s *DeepCopyRuleRequest) SetTargetEventName(v string) *DeepCopyRuleRequest {
	s.TargetEventName = &v
	return s
}

func (s *DeepCopyRuleRequest) Validate() error {
	return dara.Validate(s)
}
