// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHitRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeHitRuleListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeHitRuleListResponseBodyResultObject) *DescribeHitRuleListResponseBody
	GetResultObject() []*DescribeHitRuleListResponseBodyResultObject
}

type DescribeHitRuleListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Response object
	ResultObject []*DescribeHitRuleListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeHitRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHitRuleListResponseBody) GetResultObject() []*DescribeHitRuleListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeHitRuleListResponseBody) SetRequestId(v string) *DescribeHitRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHitRuleListResponseBody) SetResultObject(v []*DescribeHitRuleListResponseBodyResultObject) *DescribeHitRuleListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeHitRuleListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHitRuleListResponseBodyResultObject struct {
	// Number of hits.
	//
	// example:
	//
	// 100
	HitCount *int32 `json:"hitCount,omitempty" xml:"hitCount,omitempty"`
	// Strategy name
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s DescribeHitRuleListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleListResponseBodyResultObject) GetHitCount() *int32 {
	return s.HitCount
}

func (s *DescribeHitRuleListResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeHitRuleListResponseBodyResultObject) SetHitCount(v int32) *DescribeHitRuleListResponseBodyResultObject {
	s.HitCount = &v
	return s
}

func (s *DescribeHitRuleListResponseBodyResultObject) SetRuleName(v string) *DescribeHitRuleListResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *DescribeHitRuleListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
