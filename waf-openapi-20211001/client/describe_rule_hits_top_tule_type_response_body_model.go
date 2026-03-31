// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopTuleTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleHitsTopTuleTypeResponseBody
	GetRequestId() *string
	SetRuleHitsTopTuleType(v []*DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) *DescribeRuleHitsTopTuleTypeResponseBody
	GetRuleHitsTopTuleType() []*DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType
}

type DescribeRuleHitsTopTuleTypeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 45E377CE-0B04-578E-B653-EEA63CFE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The top 10 protection modules that are matched.
	RuleHitsTopTuleType []*DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType `json:"RuleHitsTopTuleType,omitempty" xml:"RuleHitsTopTuleType,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopTuleTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleHitsTopTuleTypeResponseBody) GetRuleHitsTopTuleType() []*DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType {
	return s.RuleHitsTopTuleType
}

func (s *DescribeRuleHitsTopTuleTypeResponseBody) SetRequestId(v string) *DescribeRuleHitsTopTuleTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponseBody) SetRuleHitsTopTuleType(v []*DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) *DescribeRuleHitsTopTuleTypeResponseBody {
	s.RuleHitsTopTuleType = v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponseBody) Validate() error {
	if s.RuleHitsTopTuleType != nil {
		for _, item := range s.RuleHitsTopTuleType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType struct {
	// The number of requests that match protection rules.
	//
	// example:
	//
	// 698455
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The type of rule that is matched. By default, this parameter is not returned. This indicates that all types of rules that are matched are returned.
	//
	// 	- **waf:*	- basic protection rules.
	//
	// 	- **blacklist:*	- IP address blacklist rules.
	//
	// 	- **custom:*	- custom rules.
	//
	// 	- **antiscan:*	- scan protection rules.
	//
	// 	- **cc_system:*	- HTTP flood protection rules.
	//
	// 	- **region_block:*	- region blacklist rules.
	//
	// 	- **scene:*	- bot management rules.
	//
	// 	- **dlp:*	- data leakage prevention rules.
	//
	// example:
	//
	// cc_system
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) GetCount() *int64 {
	return s.Count
}

func (s *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) SetCount(v int64) *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) SetRuleType(v string) *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) Validate() error {
	return dara.Validate(s)
}
