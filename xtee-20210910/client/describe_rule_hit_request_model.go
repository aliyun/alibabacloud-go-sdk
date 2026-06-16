// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRuleHitRequest
	GetLang() *string
	SetRegId(v string) *DescribeRuleHitRequest
	GetRegId() *string
	SetRequestTime(v int64) *DescribeRuleHitRequest
	GetRequestTime() *int64
	SetRuleId(v string) *DescribeRuleHitRequest
	GetRuleId() *string
	SetRuleSnapshotId(v string) *DescribeRuleHitRequest
	GetRuleSnapshotId() *string
	SetSRequestId(v string) *DescribeRuleHitRequest
	GetSRequestId() *string
}

type DescribeRuleHitRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The execution time.
	//
	// example:
	//
	// 1752571330000
	RequestTime *int64 `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 102059
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// 27
	RuleSnapshotId *string `json:"ruleSnapshotId,omitempty" xml:"ruleSnapshotId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 60C97040-D5D5-4906-9522-B9B413730CAA
	SRequestId *string `json:"sRequestId,omitempty" xml:"sRequestId,omitempty"`
}

func (s DescribeRuleHitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRuleHitRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRuleHitRequest) GetRequestTime() *int64 {
	return s.RequestTime
}

func (s *DescribeRuleHitRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleHitRequest) GetRuleSnapshotId() *string {
	return s.RuleSnapshotId
}

func (s *DescribeRuleHitRequest) GetSRequestId() *string {
	return s.SRequestId
}

func (s *DescribeRuleHitRequest) SetLang(v string) *DescribeRuleHitRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRuleHitRequest) SetRegId(v string) *DescribeRuleHitRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRuleHitRequest) SetRequestTime(v int64) *DescribeRuleHitRequest {
	s.RequestTime = &v
	return s
}

func (s *DescribeRuleHitRequest) SetRuleId(v string) *DescribeRuleHitRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleHitRequest) SetRuleSnapshotId(v string) *DescribeRuleHitRequest {
	s.RuleSnapshotId = &v
	return s
}

func (s *DescribeRuleHitRequest) SetSRequestId(v string) *DescribeRuleHitRequest {
	s.SRequestId = &v
	return s
}

func (s *DescribeRuleHitRequest) Validate() error {
	return dara.Validate(s)
}
