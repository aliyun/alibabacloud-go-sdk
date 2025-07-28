// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowTopResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFlowTopResourceResponseBody
	GetRequestId() *string
	SetRuleHitsTopResource(v []*DescribeFlowTopResourceResponseBodyRuleHitsTopResource) *DescribeFlowTopResourceResponseBody
	GetRuleHitsTopResource() []*DescribeFlowTopResourceResponseBodyRuleHitsTopResource
}

type DescribeFlowTopResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8F0E0B9A-B518-5C6D-BEFC-A373DDE4F652
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 protected objects that receive requests.
	RuleHitsTopResource []*DescribeFlowTopResourceResponseBodyRuleHitsTopResource `json:"RuleHitsTopResource,omitempty" xml:"RuleHitsTopResource,omitempty" type:"Repeated"`
}

func (s DescribeFlowTopResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowTopResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowTopResourceResponseBody) GetRuleHitsTopResource() []*DescribeFlowTopResourceResponseBodyRuleHitsTopResource {
	return s.RuleHitsTopResource
}

func (s *DescribeFlowTopResourceResponseBody) SetRequestId(v string) *DescribeFlowTopResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowTopResourceResponseBody) SetRuleHitsTopResource(v []*DescribeFlowTopResourceResponseBodyRuleHitsTopResource) *DescribeFlowTopResourceResponseBody {
	s.RuleHitsTopResource = v
	return s
}

func (s *DescribeFlowTopResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowTopResourceResponseBodyRuleHitsTopResource struct {
	// The total number of requests received by the protected object in a specified time range.
	//
	// example:
	//
	// 181174784
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The protected object.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s DescribeFlowTopResourceResponseBodyRuleHitsTopResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowTopResourceResponseBodyRuleHitsTopResource) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceResponseBodyRuleHitsTopResource) GetCount() *int64 {
	return s.Count
}

func (s *DescribeFlowTopResourceResponseBodyRuleHitsTopResource) GetResource() *string {
	return s.Resource
}

func (s *DescribeFlowTopResourceResponseBodyRuleHitsTopResource) SetCount(v int64) *DescribeFlowTopResourceResponseBodyRuleHitsTopResource {
	s.Count = &v
	return s
}

func (s *DescribeFlowTopResourceResponseBodyRuleHitsTopResource) SetResource(v string) *DescribeFlowTopResourceResponseBodyRuleHitsTopResource {
	s.Resource = &v
	return s
}

func (s *DescribeFlowTopResourceResponseBodyRuleHitsTopResource) Validate() error {
	return dara.Validate(s)
}
