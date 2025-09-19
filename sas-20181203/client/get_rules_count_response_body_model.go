// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRulesCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRulesCountResponseBody
	GetRequestId() *string
	SetTotalSystemClientRuleCount(v int64) *GetRulesCountResponseBody
	GetTotalSystemClientRuleCount() *int64
	SetTotalUserDefineRuleCount(v int64) *GetRulesCountResponseBody
	GetTotalUserDefineRuleCount() *int64
}

type GetRulesCountResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578AB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of system defense rules.
	//
	// example:
	//
	// 10
	TotalSystemClientRuleCount *int64 `json:"TotalSystemClientRuleCount,omitempty" xml:"TotalSystemClientRuleCount,omitempty"`
	// The total number of custom defense rules.
	//
	// example:
	//
	// 10
	TotalUserDefineRuleCount *int64 `json:"TotalUserDefineRuleCount,omitempty" xml:"TotalUserDefineRuleCount,omitempty"`
}

func (s GetRulesCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRulesCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetRulesCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRulesCountResponseBody) GetTotalSystemClientRuleCount() *int64 {
	return s.TotalSystemClientRuleCount
}

func (s *GetRulesCountResponseBody) GetTotalUserDefineRuleCount() *int64 {
	return s.TotalUserDefineRuleCount
}

func (s *GetRulesCountResponseBody) SetRequestId(v string) *GetRulesCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRulesCountResponseBody) SetTotalSystemClientRuleCount(v int64) *GetRulesCountResponseBody {
	s.TotalSystemClientRuleCount = &v
	return s
}

func (s *GetRulesCountResponseBody) SetTotalUserDefineRuleCount(v int64) *GetRulesCountResponseBody {
	s.TotalUserDefineRuleCount = &v
	return s
}

func (s *GetRulesCountResponseBody) Validate() error {
	return dara.Validate(s)
}
