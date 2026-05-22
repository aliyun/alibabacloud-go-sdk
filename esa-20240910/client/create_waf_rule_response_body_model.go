// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWafRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateWafRuleResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateWafRuleResponseBody
	GetRequestId() *string
	SetRulesetId(v int64) *CreateWafRuleResponseBody
	GetRulesetId() *int64
}

type CreateWafRuleResponseBody struct {
	// The ID of the WAF rule, which can be obtained by calling the [ListWafRules](https://help.aliyun.com/document_detail/2878257.html) API.
	//
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s CreateWafRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWafRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateWafRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWafRuleResponseBody) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *CreateWafRuleResponseBody) SetId(v int64) *CreateWafRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateWafRuleResponseBody) SetRequestId(v string) *CreateWafRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWafRuleResponseBody) SetRulesetId(v int64) *CreateWafRuleResponseBody {
	s.RulesetId = &v
	return s
}

func (s *CreateWafRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
