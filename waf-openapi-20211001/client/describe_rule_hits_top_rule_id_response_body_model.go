// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopRuleIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleHitsTopRuleIdResponseBody
	GetRequestId() *string
	SetRuleHitsTopRuleId(v []*DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) *DescribeRuleHitsTopRuleIdResponseBody
	GetRuleHitsTopRuleId() []*DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId
}

type DescribeRuleHitsTopRuleIdResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6334274-8870-5D2F-A1AD-D6EF885AC1ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The top 10 rule IDs by rule hits.
	RuleHitsTopRuleId []*DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId `json:"RuleHitsTopRuleId,omitempty" xml:"RuleHitsTopRuleId,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopRuleIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleHitsTopRuleIdResponseBody) GetRuleHitsTopRuleId() []*DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	return s.RuleHitsTopRuleId
}

func (s *DescribeRuleHitsTopRuleIdResponseBody) SetRequestId(v string) *DescribeRuleHitsTopRuleIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBody) SetRuleHitsTopRuleId(v []*DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) *DescribeRuleHitsTopRuleIdResponseBody {
	s.RuleHitsTopRuleId = v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBody) Validate() error {
	if s.RuleHitsTopRuleId != nil {
		for _, item := range s.RuleHitsTopRuleId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId struct {
	// The number of requests that hit the rule.
	//
	// example:
	//
	// 181174784
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The protected object. This parameter is returned when the IsGroupResource request parameter is set to false.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 5465465
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) GetCount() *int64 {
	return s.Count
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) GetResource() *string {
	return s.Resource
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) SetCount(v int64) *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) SetResource(v string) *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) SetRuleId(v string) *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) Validate() error {
	return dara.Validate(s)
}
