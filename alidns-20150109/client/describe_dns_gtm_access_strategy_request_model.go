// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAccessStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDnsGtmAccessStrategyRequest
	GetLang() *string
	SetStrategyId(v string) *DescribeDnsGtmAccessStrategyRequest
	GetStrategyId() *string
}

type DescribeDnsGtmAccessStrategyRequest struct {
	// The language to return some response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// strategyId1
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmAccessStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeDnsGtmAccessStrategyRequest) SetLang(v string) *DescribeDnsGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyRequest) SetStrategyId(v string) *DescribeDnsGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyRequest) Validate() error {
	return dara.Validate(s)
}
