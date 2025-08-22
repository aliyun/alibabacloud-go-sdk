// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserSecDropRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeDcdnUserSecDropRequest
	GetData() *string
	SetMetric(v string) *DescribeDcdnUserSecDropRequest
	GetMetric() *string
	SetSecFunc(v string) *DescribeDcdnUserSecDropRequest
	GetSecFunc() *string
}

type DescribeDcdnUserSecDropRequest struct {
	// The date or month that you want to query.
	//
	// 	- If data is collected every day, set Data in the format of yyyymmdd, such as 20201203.
	//
	// 	- If data is collected every month, set Data in the format of yyyymm, such as 202012.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20201203
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The time interval at which data is collected.
	//
	// 	- If data is collected every day, the number of blocked packets on the specified day is calculated.
	//
	// 	- If data is collected every month, the number of blocked packets in the specified month is calculated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1day
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The security feature. Valid values:
	//
	// 	- waf: WAF
	//
	// 	- tmd: rate limiting
	//
	// 	- robot: bot traffic recognition
	//
	// 	- l4_dm_drop: domain name blocking at Layer 4
	//
	// This parameter is required.
	//
	// example:
	//
	// waf
	SecFunc *string `json:"SecFunc,omitempty" xml:"SecFunc,omitempty"`
}

func (s DescribeDcdnUserSecDropRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserSecDropRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropRequest) GetData() *string {
	return s.Data
}

func (s *DescribeDcdnUserSecDropRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeDcdnUserSecDropRequest) GetSecFunc() *string {
	return s.SecFunc
}

func (s *DescribeDcdnUserSecDropRequest) SetData(v string) *DescribeDcdnUserSecDropRequest {
	s.Data = &v
	return s
}

func (s *DescribeDcdnUserSecDropRequest) SetMetric(v string) *DescribeDcdnUserSecDropRequest {
	s.Metric = &v
	return s
}

func (s *DescribeDcdnUserSecDropRequest) SetSecFunc(v string) *DescribeDcdnUserSecDropRequest {
	s.SecFunc = &v
	return s
}

func (s *DescribeDcdnUserSecDropRequest) Validate() error {
	return dara.Validate(s)
}
