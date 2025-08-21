// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwitdhByInternetChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeBandwitdhByInternetChargeTypeRequest
	GetEndTime() *string
	SetEnsRegionId(v string) *DescribeBandwitdhByInternetChargeTypeRequest
	GetEnsRegionId() *string
	SetIsp(v string) *DescribeBandwitdhByInternetChargeTypeRequest
	GetIsp() *string
	SetStartTime(v string) *DescribeBandwitdhByInternetChargeTypeRequest
	GetStartTime() *string
}

type DescribeBandwitdhByInternetChargeTypeRequest struct {
	// The end of the time range to query.
	//
	// 	- Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- If the value of the seconds place is not 00, the start time is automatically set to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-18T09:39:54Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-cbn-2
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The Internet service provider (ISP). Valid values:
	//
	// 	- cmcc: China Mobile
	//
	// 	- telecom: China Telecom
	//
	// 	- unicom: China Unicom
	//
	// 	- multiCarrier: multi-line ISP
	//
	// example:
	//
	// unicom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The beginning of the time range to query.
	//
	// 	- Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- If the value of the seconds place is not 00, the start time is automatically set to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-15T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBandwitdhByInternetChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwitdhByInternetChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetEndTime(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetEnsRegionId(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetIsp(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.Isp = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) SetStartTime(v string) *DescribeBandwitdhByInternetChargeTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBandwitdhByInternetChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
