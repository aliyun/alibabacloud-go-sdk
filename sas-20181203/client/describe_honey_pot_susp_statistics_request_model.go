// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHoneyPotSuspStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeHoneyPotSuspStatisticsRequest
	GetFrom() *string
	SetLang(v string) *DescribeHoneyPotSuspStatisticsRequest
	GetLang() *string
	SetStatisticsDays(v int32) *DescribeHoneyPotSuspStatisticsRequest
	GetStatisticsDays() *int32
	SetStatisticsKeyType(v string) *DescribeHoneyPotSuspStatisticsRequest
	GetStatisticsKeyType() *string
}

type DescribeHoneyPotSuspStatisticsRequest struct {
	// The source of the request. Set the value to **honeypot**.
	//
	// This parameter is required.
	//
	// example:
	//
	// honeypot
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The time range of the data to query. Unit: days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	StatisticsDays *int32 `json:"StatisticsDays,omitempty" xml:"StatisticsDays,omitempty"`
	// The type of the asset to query. Valid values:
	//
	// 	- **vpcInstanceId**: VPC
	//
	// 	- **uuid**: server
	//
	// This parameter is required.
	//
	// example:
	//
	// vpcInstanceId
	StatisticsKeyType *string `json:"StatisticsKeyType,omitempty" xml:"StatisticsKeyType,omitempty"`
}

func (s DescribeHoneyPotSuspStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHoneyPotSuspStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotSuspStatisticsRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeHoneyPotSuspStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeHoneyPotSuspStatisticsRequest) GetStatisticsDays() *int32 {
	return s.StatisticsDays
}

func (s *DescribeHoneyPotSuspStatisticsRequest) GetStatisticsKeyType() *string {
	return s.StatisticsKeyType
}

func (s *DescribeHoneyPotSuspStatisticsRequest) SetFrom(v string) *DescribeHoneyPotSuspStatisticsRequest {
	s.From = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsRequest) SetLang(v string) *DescribeHoneyPotSuspStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsRequest) SetStatisticsDays(v int32) *DescribeHoneyPotSuspStatisticsRequest {
	s.StatisticsDays = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsRequest) SetStatisticsKeyType(v string) *DescribeHoneyPotSuspStatisticsRequest {
	s.StatisticsKeyType = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
