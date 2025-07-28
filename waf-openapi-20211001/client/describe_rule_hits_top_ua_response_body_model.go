// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopUaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleHitsTopUaResponseBody
	GetRequestId() *string
	SetRuleHitsTopUa(v []*DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) *DescribeRuleHitsTopUaResponseBody
	GetRuleHitsTopUa() []*DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa
}

type DescribeRuleHitsTopUaResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8E5C7ED7-503A-5986-A005-36F2511EB89F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 user agents that are used to initiate attacks.
	RuleHitsTopUa []*DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa `json:"RuleHitsTopUa,omitempty" xml:"RuleHitsTopUa,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopUaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopUaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleHitsTopUaResponseBody) GetRuleHitsTopUa() []*DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa {
	return s.RuleHitsTopUa
}

func (s *DescribeRuleHitsTopUaResponseBody) SetRequestId(v string) *DescribeRuleHitsTopUaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponseBody) SetRuleHitsTopUa(v []*DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) *DescribeRuleHitsTopUaResponseBody {
	s.RuleHitsTopUa = v
	return s
}

func (s *DescribeRuleHitsTopUaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa struct {
	// The number of attacks that are initiated from the IP address.
	//
	// example:
	//
	// 531
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The user agent.
	//
	// example:
	//
	// android
	Ua *string `json:"Ua,omitempty" xml:"Ua,omitempty"`
}

func (s DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) GetCount() *int64 {
	return s.Count
}

func (s *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) GetUa() *string {
	return s.Ua
}

func (s *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) SetCount(v int64) *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) SetUa(v string) *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa {
	s.Ua = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) Validate() error {
	return dara.Validate(s)
}
