// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDesensitizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionParams(v []map[string]*string) *AddDesensitizationRuleRequest
	GetFunctionParams() []map[string]*string
	SetFunctionType(v string) *AddDesensitizationRuleRequest
	GetFunctionType() *string
	SetRuleDescription(v string) *AddDesensitizationRuleRequest
	GetRuleDescription() *string
	SetRuleName(v string) *AddDesensitizationRuleRequest
	GetRuleName() *string
	SetRuleType(v string) *AddDesensitizationRuleRequest
	GetRuleType() *string
	SetTid(v int64) *AddDesensitizationRuleRequest
	GetTid() *int64
}

type AddDesensitizationRuleRequest struct {
	// The parameters of the algorithm.
	FunctionParams []map[string]*string `json:"FunctionParams,omitempty" xml:"FunctionParams,omitempty" type:"Repeated"`
	// The type of the masking algorithm.
	//
	// Valid values:
	//
	// 	- FIX_POS : masks characters in the specified position.
	//
	// 	- DATE_ROUNDING: rounds the date.
	//
	// 	- PLAINTEXT: does not mask data.
	//
	// 	- SHA1: masks characters by using the secure hash algorithm 1 (SHA-1)
	//
	// 	- HMAC: masks characters by using the hash-based message authentication code (HMAC).
	//
	// 	- STRING_TRANSFORM: shift characters.
	//
	// 	- NUMBER_ROUNDING: rounds numbers.
	//
	// 	- AES: masks characters by using the advanced encryption standard (AES) algorithm.
	//
	// 	- SHA256: masks characters by using SHA-256 algorithm.
	//
	// 	- DES: masks characters by using the data encryption standard (DES) algorithm.
	//
	// 	- MAP_REPLACE: masks the mapped data.
	//
	// 	- FIX_CHAR: masks fixed characters.
	//
	// 	- DEFAULT: masks all characters.
	//
	// 	- RANDOM_REPLACE: randomly replaces characters.
	//
	// 	- MD5: masks characters by using the MD5 algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// MD5
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// hash algorithm
	RuleDescription *string `json:"RuleDescription,omitempty" xml:"RuleDescription,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// desensitization algorithm test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The masking algorithm.
	//
	// Valid values:
	//
	// 	- PLAINTEXT
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- TRANSFORM
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ENCRYPT
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- REPLACE
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- HASH
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- MASK
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// HASH
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The tenant ID.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s AddDesensitizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDesensitizationRuleRequest) GoString() string {
	return s.String()
}

func (s *AddDesensitizationRuleRequest) GetFunctionParams() []map[string]*string {
	return s.FunctionParams
}

func (s *AddDesensitizationRuleRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *AddDesensitizationRuleRequest) GetRuleDescription() *string {
	return s.RuleDescription
}

func (s *AddDesensitizationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *AddDesensitizationRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *AddDesensitizationRuleRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddDesensitizationRuleRequest) SetFunctionParams(v []map[string]*string) *AddDesensitizationRuleRequest {
	s.FunctionParams = v
	return s
}

func (s *AddDesensitizationRuleRequest) SetFunctionType(v string) *AddDesensitizationRuleRequest {
	s.FunctionType = &v
	return s
}

func (s *AddDesensitizationRuleRequest) SetRuleDescription(v string) *AddDesensitizationRuleRequest {
	s.RuleDescription = &v
	return s
}

func (s *AddDesensitizationRuleRequest) SetRuleName(v string) *AddDesensitizationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *AddDesensitizationRuleRequest) SetRuleType(v string) *AddDesensitizationRuleRequest {
	s.RuleType = &v
	return s
}

func (s *AddDesensitizationRuleRequest) SetTid(v int64) *AddDesensitizationRuleRequest {
	s.Tid = &v
	return s
}

func (s *AddDesensitizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
