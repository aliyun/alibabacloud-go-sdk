// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDcdnWafRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchCreateDcdnWafRulesResponseBody
	GetRequestId() *string
	SetRuleIds(v *BatchCreateDcdnWafRulesResponseBodyRuleIds) *BatchCreateDcdnWafRulesResponseBody
	GetRuleIds() *BatchCreateDcdnWafRulesResponseBodyRuleIds
}

type BatchCreateDcdnWafRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of created rules.
	RuleIds *BatchCreateDcdnWafRulesResponseBodyRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
}

func (s BatchCreateDcdnWafRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDcdnWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateDcdnWafRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateDcdnWafRulesResponseBody) GetRuleIds() *BatchCreateDcdnWafRulesResponseBodyRuleIds {
	return s.RuleIds
}

func (s *BatchCreateDcdnWafRulesResponseBody) SetRequestId(v string) *BatchCreateDcdnWafRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateDcdnWafRulesResponseBody) SetRuleIds(v *BatchCreateDcdnWafRulesResponseBodyRuleIds) *BatchCreateDcdnWafRulesResponseBody {
	s.RuleIds = v
	return s
}

func (s *BatchCreateDcdnWafRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchCreateDcdnWafRulesResponseBodyRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s BatchCreateDcdnWafRulesResponseBodyRuleIds) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDcdnWafRulesResponseBodyRuleIds) GoString() string {
	return s.String()
}

func (s *BatchCreateDcdnWafRulesResponseBodyRuleIds) GetRuleId() []*string {
	return s.RuleId
}

func (s *BatchCreateDcdnWafRulesResponseBodyRuleIds) SetRuleId(v []*string) *BatchCreateDcdnWafRulesResponseBodyRuleIds {
	s.RuleId = v
	return s
}

func (s *BatchCreateDcdnWafRulesResponseBodyRuleIds) Validate() error {
	return dara.Validate(s)
}
