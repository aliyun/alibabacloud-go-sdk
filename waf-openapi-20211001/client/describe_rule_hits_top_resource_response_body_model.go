// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleHitsTopResourceResponseBody
	GetRequestId() *string
	SetRuleHitsTopResource(v []*DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) *DescribeRuleHitsTopResourceResponseBody
	GetRuleHitsTopResource() []*DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource
}

type DescribeRuleHitsTopResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ADA11BC7-AA95-5C31-9095-5802C02ED1DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 protected objects that trigger protection rules.
	RuleHitsTopResource []*DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource `json:"RuleHitsTopResource,omitempty" xml:"RuleHitsTopResource,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleHitsTopResourceResponseBody) GetRuleHitsTopResource() []*DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource {
	return s.RuleHitsTopResource
}

func (s *DescribeRuleHitsTopResourceResponseBody) SetRequestId(v string) *DescribeRuleHitsTopResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponseBody) SetRuleHitsTopResource(v []*DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) *DescribeRuleHitsTopResourceResponseBody {
	s.RuleHitsTopResource = v
	return s
}

func (s *DescribeRuleHitsTopResourceResponseBody) Validate() error {
	if s.RuleHitsTopResource != nil {
		for _, item := range s.RuleHitsTopResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource struct {
	// The number of requests that match protection rules.
	//
	// example:
	//
	// 14219
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The protected object.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) GetCount() *int64 {
	return s.Count
}

func (s *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) GetResource() *string {
	return s.Resource
}

func (s *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) SetCount(v int64) *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) SetResource(v string) *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) Validate() error {
	return dara.Validate(s)
}
