// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateWafRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *BatchCreateWafRulesResponseBody
	GetIds() []*int64
	SetRequestId(v string) *BatchCreateWafRulesResponseBody
	GetRequestId() *string
	SetRulesetId(v int64) *BatchCreateWafRulesResponseBody
	GetRulesetId() *int64
}

type BatchCreateWafRulesResponseBody struct {
	// ID of the WAF rule, which can be obtained by calling the [ListWafRules](https://help.aliyun.com/document_detail/2878257.html) interface.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ID of the WAF ruleset, which can be obtained by calling the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) interface.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
}

func (s BatchCreateWafRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateWafRulesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateWafRulesResponseBody) GetIds() []*int64 {
	return s.Ids
}

func (s *BatchCreateWafRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateWafRulesResponseBody) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *BatchCreateWafRulesResponseBody) SetIds(v []*int64) *BatchCreateWafRulesResponseBody {
	s.Ids = v
	return s
}

func (s *BatchCreateWafRulesResponseBody) SetRequestId(v string) *BatchCreateWafRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateWafRulesResponseBody) SetRulesetId(v int64) *BatchCreateWafRulesResponseBody {
	s.RulesetId = &v
	return s
}

func (s *BatchCreateWafRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
