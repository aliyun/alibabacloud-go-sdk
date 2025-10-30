// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleListByEventCodesListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleListByEventCodesListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeRuleListByEventCodesListResponseBodyResultObject) *DescribeRuleListByEventCodesListResponseBody
	GetResultObject() []*DescribeRuleListByEventCodesListResponseBodyResultObject
}

type DescribeRuleListByEventCodesListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject []*DescribeRuleListByEventCodesListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeRuleListByEventCodesListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleListByEventCodesListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleListByEventCodesListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleListByEventCodesListResponseBody) GetResultObject() []*DescribeRuleListByEventCodesListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRuleListByEventCodesListResponseBody) SetRequestId(v string) *DescribeRuleListByEventCodesListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleListByEventCodesListResponseBody) SetResultObject(v []*DescribeRuleListByEventCodesListResponseBodyResultObject) *DescribeRuleListByEventCodesListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRuleListByEventCodesListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRuleListByEventCodesListResponseBodyResultObject struct {
	// Policy ID
	//
	// example:
	//
	// 4730
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s DescribeRuleListByEventCodesListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleListByEventCodesListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRuleListByEventCodesListResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleListByEventCodesListResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRuleListByEventCodesListResponseBodyResultObject) SetRuleId(v string) *DescribeRuleListByEventCodesListResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleListByEventCodesListResponseBodyResultObject) SetRuleName(v string) *DescribeRuleListByEventCodesListResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *DescribeRuleListByEventCodesListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
