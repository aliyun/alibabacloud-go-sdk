// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAccessStrategyAvailableConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest
	GetLang() *string
	SetStrategyMode(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest
	GetStrategyMode() *string
}

type DescribeDnsGtmAccessStrategyAvailableConfigRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language to return some response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the access policy. Valid values:
	//
	// 	- GEO: geographical location-based
	//
	// 	- LATENCY: latency-based
	//
	// This parameter is required.
	//
	// example:
	//
	// geo
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) GetStrategyMode() *string {
	return s.StrategyMode
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) SetInstanceId(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) SetLang(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) SetStrategyMode(v string) *DescribeDnsGtmAccessStrategyAvailableConfigRequest {
	s.StrategyMode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyAvailableConfigRequest) Validate() error {
	return dara.Validate(s)
}
