// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedStrategyTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCustomizedStrategyTargetsRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeCustomizedStrategyTargetsRequest
	GetSourceIp() *string
}

type DescribeCustomizedStrategyTargetsRequest struct {
	// The language type for requests and responses. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 39.170.43.**
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeCustomizedStrategyTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedStrategyTargetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedStrategyTargetsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustomizedStrategyTargetsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCustomizedStrategyTargetsRequest) SetLang(v string) *DescribeCustomizedStrategyTargetsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustomizedStrategyTargetsRequest) SetSourceIp(v string) *DescribeCustomizedStrategyTargetsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCustomizedStrategyTargetsRequest) Validate() error {
	return dara.Validate(s)
}
