// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDesensitizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFuncType(v string) *ListDesensitizationRuleRequest
	GetFuncType() *string
	SetPageNumber(v int32) *ListDesensitizationRuleRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDesensitizationRuleRequest
	GetPageSize() *int32
	SetRuleId(v int32) *ListDesensitizationRuleRequest
	GetRuleId() *int32
	SetRuleName(v string) *ListDesensitizationRuleRequest
	GetRuleName() *string
	SetRuleType(v string) *ListDesensitizationRuleRequest
	GetRuleType() *string
	SetTid(v int64) *ListDesensitizationRuleRequest
	GetTid() *int64
}

type ListDesensitizationRuleRequest struct {
	// The type of the masking algorithm.
	//
	// example:
	//
	// MD5
	FuncType *string `json:"FuncType,omitempty" xml:"FuncType,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the masking rule.
	//
	// example:
	//
	// 1
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
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDesensitizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDesensitizationRuleRequest) GoString() string {
	return s.String()
}

func (s *ListDesensitizationRuleRequest) GetFuncType() *string {
	return s.FuncType
}

func (s *ListDesensitizationRuleRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDesensitizationRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDesensitizationRuleRequest) GetRuleId() *int32 {
	return s.RuleId
}

func (s *ListDesensitizationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListDesensitizationRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ListDesensitizationRuleRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDesensitizationRuleRequest) SetFuncType(v string) *ListDesensitizationRuleRequest {
	s.FuncType = &v
	return s
}

func (s *ListDesensitizationRuleRequest) SetPageNumber(v int32) *ListDesensitizationRuleRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDesensitizationRuleRequest) SetPageSize(v int32) *ListDesensitizationRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListDesensitizationRuleRequest) SetRuleId(v int32) *ListDesensitizationRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ListDesensitizationRuleRequest) SetRuleName(v string) *ListDesensitizationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ListDesensitizationRuleRequest) SetRuleType(v string) *ListDesensitizationRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ListDesensitizationRuleRequest) SetTid(v int64) *ListDesensitizationRuleRequest {
	s.Tid = &v
	return s
}

func (s *ListDesensitizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
