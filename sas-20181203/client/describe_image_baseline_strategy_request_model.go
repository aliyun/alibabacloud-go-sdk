// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeImageBaselineStrategyRequest
	GetLang() *string
	SetSource(v string) *DescribeImageBaselineStrategyRequest
	GetSource() *string
	SetStrategyId(v int64) *DescribeImageBaselineStrategyRequest
	GetStrategyId() *int64
}

type DescribeImageBaselineStrategyRequest struct {
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The data source. If this parameter is left empty, the image baseline policy is queried by default. Valid values:
	//
	// - **default**: image
	//
	// - **agentless**: agentless.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the baseline check policy.
	//
	// example:
	//
	// 8037
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeImageBaselineStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineStrategyRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageBaselineStrategyRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeImageBaselineStrategyRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeImageBaselineStrategyRequest) SetLang(v string) *DescribeImageBaselineStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageBaselineStrategyRequest) SetSource(v string) *DescribeImageBaselineStrategyRequest {
	s.Source = &v
	return s
}

func (s *DescribeImageBaselineStrategyRequest) SetStrategyId(v int64) *DescribeImageBaselineStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeImageBaselineStrategyRequest) Validate() error {
	return dara.Validate(s)
}
