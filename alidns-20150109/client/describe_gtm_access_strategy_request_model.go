// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAccessStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGtmAccessStrategyRequest
	GetLang() *string
	SetStrategyId(v string) *DescribeGtmAccessStrategyRequest
	GetStrategyId() *string
}

type DescribeGtmAccessStrategyRequest struct {
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the access policy that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// hra0hs
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeGtmAccessStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmAccessStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeGtmAccessStrategyRequest) SetLang(v string) *DescribeGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmAccessStrategyRequest) SetStrategyId(v string) *DescribeGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeGtmAccessStrategyRequest) Validate() error {
	return dara.Validate(s)
}
