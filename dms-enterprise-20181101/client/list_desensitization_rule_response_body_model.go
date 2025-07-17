// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDesensitizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesensitizationRuleList(v []*ListDesensitizationRuleResponseBodyDesensitizationRuleList) *ListDesensitizationRuleResponseBody
	GetDesensitizationRuleList() []*ListDesensitizationRuleResponseBodyDesensitizationRuleList
	SetErrorCode(v string) *ListDesensitizationRuleResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDesensitizationRuleResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDesensitizationRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDesensitizationRuleResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDesensitizationRuleResponseBody
	GetTotalCount() *int32
}

type ListDesensitizationRuleResponseBody struct {
	// The list of masking rules.
	DesensitizationRuleList []*ListDesensitizationRuleResponseBodyDesensitizationRuleList `json:"DesensitizationRuleList,omitempty" xml:"DesensitizationRuleList,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// This parameter is required.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E76DD2E7-EBAC-5724-B163-19AAC233F8F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned. By default, this parameter is not returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDesensitizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDesensitizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListDesensitizationRuleResponseBody) GetDesensitizationRuleList() []*ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	return s.DesensitizationRuleList
}

func (s *ListDesensitizationRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDesensitizationRuleResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDesensitizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDesensitizationRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDesensitizationRuleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDesensitizationRuleResponseBody) SetDesensitizationRuleList(v []*ListDesensitizationRuleResponseBodyDesensitizationRuleList) *ListDesensitizationRuleResponseBody {
	s.DesensitizationRuleList = v
	return s
}

func (s *ListDesensitizationRuleResponseBody) SetErrorCode(v string) *ListDesensitizationRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDesensitizationRuleResponseBody) SetErrorMessage(v string) *ListDesensitizationRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDesensitizationRuleResponseBody) SetRequestId(v string) *ListDesensitizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDesensitizationRuleResponseBody) SetSuccess(v bool) *ListDesensitizationRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ListDesensitizationRuleResponseBody) SetTotalCount(v int32) *ListDesensitizationRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDesensitizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDesensitizationRuleResponseBodyDesensitizationRuleList struct {
	// The parameter.
	//
	// example:
	//
	// {paramName: salt, paramValue: 1}
	FuncParams *string `json:"FuncParams,omitempty" xml:"FuncParams,omitempty"`
	// The example.
	//
	// example:
	//
	// [{paramName: testStr, paramValue: 1}]
	FuncSample *string `json:"FuncSample,omitempty" xml:"FuncSample,omitempty"`
	// The algorithm type.
	//
	// example:
	//
	// MD5
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The ID of the user who last modified the masking rule.
	//
	// example:
	//
	// 2
	LastModifierId *string `json:"LastModifierId,omitempty" xml:"LastModifierId,omitempty"`
	// The name of the user who last modified the masking rule.
	//
	// example:
	//
	// test user
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// The number of times that the masking was used.
	//
	// example:
	//
	// 1
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// hash desensitization algorithm
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The ID of the masking rule.
	//
	// example:
	//
	// 23
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the masking rule.
	//
	// example:
	//
	// default desensitization rule test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The algorithm used for masking.
	//
	// example:
	//
	// HASH
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListDesensitizationRuleResponseBodyDesensitizationRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListDesensitizationRuleResponseBodyDesensitizationRuleList) GoString() string {
	return s.String()
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetFuncParams() *string {
	return s.FuncParams
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetFuncSample() *string {
	return s.FuncSample
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetLastModifierId() *string {
	return s.LastModifierId
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetReferenceCount() *int32 {
	return s.ReferenceCount
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetRuleId() *int32 {
	return s.RuleId
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetFuncParams(v string) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.FuncParams = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetFuncSample(v string) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.FuncSample = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetFunctionType(v string) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.FunctionType = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetLastModifierId(v string) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.LastModifierId = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetLastModifierName(v string) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.LastModifierName = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetReferenceCount(v int32) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.ReferenceCount = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetRuleDesc(v string) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.RuleDesc = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetRuleId(v int32) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.RuleId = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetRuleName(v string) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.RuleName = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) SetRuleType(v string) *ListDesensitizationRuleResponseBodyDesensitizationRuleList {
	s.RuleType = &v
	return s
}

func (s *ListDesensitizationRuleResponseBodyDesensitizationRuleList) Validate() error {
	return dara.Validate(s)
}
