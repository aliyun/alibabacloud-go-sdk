// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowTopUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFlowTopUrlResponseBody
	GetRequestId() *string
	SetRuleHitsTopUrl(v []*DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) *DescribeFlowTopUrlResponseBody
	GetRuleHitsTopUrl() []*DescribeFlowTopUrlResponseBodyRuleHitsTopUrl
}

type DescribeFlowTopUrlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 520D4E4C-B8EC-5602-ACB6-4D378ACBA28D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 URLs that are used to initiate requests.
	RuleHitsTopUrl []*DescribeFlowTopUrlResponseBodyRuleHitsTopUrl `json:"RuleHitsTopUrl,omitempty" xml:"RuleHitsTopUrl,omitempty" type:"Repeated"`
}

func (s DescribeFlowTopUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowTopUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowTopUrlResponseBody) GetRuleHitsTopUrl() []*DescribeFlowTopUrlResponseBodyRuleHitsTopUrl {
	return s.RuleHitsTopUrl
}

func (s *DescribeFlowTopUrlResponseBody) SetRequestId(v string) *DescribeFlowTopUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowTopUrlResponseBody) SetRuleHitsTopUrl(v []*DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) *DescribeFlowTopUrlResponseBody {
	s.RuleHitsTopUrl = v
	return s
}

func (s *DescribeFlowTopUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowTopUrlResponseBodyRuleHitsTopUrl struct {
	// The total number of requests that are initiated by using the URL.
	//
	// example:
	//
	// 181174784
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The URL that is used to initiate requests.
	//
	// example:
	//
	// www.aliyundoc.com/path1
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) GetCount() *int64 {
	return s.Count
}

func (s *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) GetUrl() *string {
	return s.Url
}

func (s *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) SetCount(v int64) *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl {
	s.Count = &v
	return s
}

func (s *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) SetUrl(v string) *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl {
	s.Url = &v
	return s
}

func (s *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) Validate() error {
	return dara.Validate(s)
}
