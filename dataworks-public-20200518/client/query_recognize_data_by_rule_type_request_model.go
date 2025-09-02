// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognizeDataByRuleTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecognizeRulesType(v string) *QueryRecognizeDataByRuleTypeRequest
	GetRecognizeRulesType() *string
	SetTenantId(v string) *QueryRecognizeDataByRuleTypeRequest
	GetTenantId() *string
}

type QueryRecognizeDataByRuleTypeRequest struct {
	// The type of a sensitive data identification rule. You can call the [QueryRecognizeRulesType](https://help.aliyun.com/document_detail/2746905.html) operation to obtain the type of the rule.
	//
	// 	- 1: regular expression
	//
	// 	- 2: built-in rule
	//
	// 	- 3: sample library
	//
	// 	- 4: self-generated data identification model
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RecognizeRulesType *string `json:"RecognizeRulesType,omitempty" xml:"RecognizeRulesType,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryRecognizeDataByRuleTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognizeDataByRuleTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryRecognizeDataByRuleTypeRequest) GetRecognizeRulesType() *string {
	return s.RecognizeRulesType
}

func (s *QueryRecognizeDataByRuleTypeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryRecognizeDataByRuleTypeRequest) SetRecognizeRulesType(v string) *QueryRecognizeDataByRuleTypeRequest {
	s.RecognizeRulesType = &v
	return s
}

func (s *QueryRecognizeDataByRuleTypeRequest) SetTenantId(v string) *QueryRecognizeDataByRuleTypeRequest {
	s.TenantId = &v
	return s
}

func (s *QueryRecognizeDataByRuleTypeRequest) Validate() error {
	return dara.Validate(s)
}
