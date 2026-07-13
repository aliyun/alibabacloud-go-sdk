// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAccessStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsGtmAccessStrategiesRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsGtmAccessStrategiesRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeDnsGtmAccessStrategiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnsGtmAccessStrategiesRequest
	GetPageSize() *int32
	SetStrategyMode(v string) *DescribeDnsGtmAccessStrategiesRequest
	GetStrategyMode() *string
}

type DescribeDnsGtmAccessStrategiesRequest struct {
	// The instance ID.<props="china"> Call [DescribeDnsGtmInstances](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describednsgtminstances?spm=a2c4g.11186623.help-menu-29697.d_0_5_1_3_8_8.2aea3618RlSR9K) to obtain the instance ID.
	//
	// <props="intl">Call [DescribeDnsGtmInstances](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describednsgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0) to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the access policy.
	//
	// - GEO: Geographic location-based access policy
	//
	// - LATENCY: Latency-based access policy
	//
	// This parameter is required.
	//
	// example:
	//
	// GEO
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
}

func (s DescribeDnsGtmAccessStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategiesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmAccessStrategiesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmAccessStrategiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnsGtmAccessStrategiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnsGtmAccessStrategiesRequest) GetStrategyMode() *string {
	return s.StrategyMode
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetInstanceId(v string) *DescribeDnsGtmAccessStrategiesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetLang(v string) *DescribeDnsGtmAccessStrategiesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetPageNumber(v int32) *DescribeDnsGtmAccessStrategiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetPageSize(v int32) *DescribeDnsGtmAccessStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) SetStrategyMode(v string) *DescribeDnsGtmAccessStrategiesRequest {
	s.StrategyMode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
