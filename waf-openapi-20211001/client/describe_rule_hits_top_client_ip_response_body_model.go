// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopClientIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleHitsTopClientIpResponseBody
	GetRequestId() *string
	SetRuleHitsTopClientIp(v []*DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) *DescribeRuleHitsTopClientIpResponseBody
	GetRuleHitsTopClientIp() []*DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp
}

type DescribeRuleHitsTopClientIpResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9F0F9AD6-62E2-50BB-A3E5-30FFB9410262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 IP addresses from which attacks are initiated.
	RuleHitsTopClientIp []*DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp `json:"RuleHitsTopClientIp,omitempty" xml:"RuleHitsTopClientIp,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopClientIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleHitsTopClientIpResponseBody) GetRuleHitsTopClientIp() []*DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp {
	return s.RuleHitsTopClientIp
}

func (s *DescribeRuleHitsTopClientIpResponseBody) SetRequestId(v string) *DescribeRuleHitsTopClientIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponseBody) SetRuleHitsTopClientIp(v []*DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) *DescribeRuleHitsTopClientIpResponseBody {
	s.RuleHitsTopClientIp = v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp struct {
	// The IP address of the service client.
	//
	// example:
	//
	// 3.3.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The number of attacks that are initiated from the IP address.
	//
	// example:
	//
	// 531
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) GetCount() *int64 {
	return s.Count
}

func (s *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) SetClientIp(v string) *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp {
	s.ClientIp = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) SetCount(v int64) *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) Validate() error {
	return dara.Validate(s)
}
