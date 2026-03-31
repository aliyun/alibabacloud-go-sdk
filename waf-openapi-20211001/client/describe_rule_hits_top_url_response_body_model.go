// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleHitsTopUrlResponseBody
	GetRequestId() *string
	SetRuleHitsTopUrl(v []*DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) *DescribeRuleHitsTopUrlResponseBody
	GetRuleHitsTopUrl() []*DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl
}

type DescribeRuleHitsTopUrlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D8AF43B-08EB-51CE-B33A-93AA****9B0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The top 10 URLs that match protection rules.
	RuleHitsTopUrl []*DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl `json:"RuleHitsTopUrl,omitempty" xml:"RuleHitsTopUrl,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleHitsTopUrlResponseBody) GetRuleHitsTopUrl() []*DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl {
	return s.RuleHitsTopUrl
}

func (s *DescribeRuleHitsTopUrlResponseBody) SetRequestId(v string) *DescribeRuleHitsTopUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponseBody) SetRuleHitsTopUrl(v []*DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) *DescribeRuleHitsTopUrlResponseBody {
	s.RuleHitsTopUrl = v
	return s
}

func (s *DescribeRuleHitsTopUrlResponseBody) Validate() error {
	if s.RuleHitsTopUrl != nil {
		for _, item := range s.RuleHitsTopUrl {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl struct {
	// The number of requests that match protection rules.
	//
	// example:
	//
	// 21862
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The request URL.
	//
	// >  The value is Base64-encoded.
	//
	// example:
	//
	// d3d3LmFsaXl1bmRvYy5jb20vcGF0aDM=
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) GetCount() *int64 {
	return s.Count
}

func (s *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) GetUrl() *string {
	return s.Url
}

func (s *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) SetCount(v int64) *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) SetUrl(v string) *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl {
	s.Url = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) Validate() error {
	return dara.Validate(s)
}
