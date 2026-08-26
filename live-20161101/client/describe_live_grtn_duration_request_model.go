// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveGrtnDurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeLiveGrtnDurationRequest
	GetAppId() *string
	SetArea(v string) *DescribeLiveGrtnDurationRequest
	GetArea() *string
	SetEndTime(v string) *DescribeLiveGrtnDurationRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveGrtnDurationRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeLiveGrtnDurationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveGrtnDurationRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveGrtnDurationRequest
	GetStartTime() *string
}

type DescribeLiveGrtnDurationRequest struct {
	// Application ID. You can query multiple application IDs separated by commas (half-width). A maximum of 30 IDs can be queried. By default, aggregated data for all applications is returned.
	//
	// example:
	//
	// 4346289a-a790-4869-9e23-22766d5e****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The area code. Valid values:
	//
	// - CN: Chinese mainland.
	//
	// - OverSeas: Overseas regions.
	//
	// - AP1: Asia Pacific 1, including Hong Kong (China), Macao (China), Taiwan (China), Japan, and Southeast Asian countries except Vietnam and Indonesia.
	//
	// - AP2: Asia Pacific 2, including Indonesia, South Korea, and Vietnam.
	//
	// - AP3: Asia Pacific 3, including Australia and New Zealand.
	//
	// - NA: North America, including the United States and Canada.
	//
	// - SA: South America, specifically Brazil.
	//
	// - EU: Europe, including Ukraine, the United Kingdom, France, the Netherlands, Spain, Italy, Sweden, and Germany.
	//
	// - MEAA: Middle East and Africa, including South Africa, Oman, the United Arab Emirates, and Kuwait.
	//
	// If not specified, aggregated data for all areas is returned by default.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The end time must be later than the start time. The query granularity must be ≥ 5 minutes and ≤ 31 days. The date format follows the ISO 8601 notation and uses UTC time in the format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for querying data. Unit: seconds. Valid values:
	//
	// - 300
	//
	// - 3600
	//
	// - 86400
	//
	// If not specified or an unsupported value is passed, the default value of 3600 seconds is used.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time for data retrieval. The date format follows the ISO 8601 notation and uses UTC time in the format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveGrtnDurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveGrtnDurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveGrtnDurationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLiveGrtnDurationRequest) GetArea() *string {
	return s.Area
}

func (s *DescribeLiveGrtnDurationRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveGrtnDurationRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveGrtnDurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveGrtnDurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveGrtnDurationRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveGrtnDurationRequest) SetAppId(v string) *DescribeLiveGrtnDurationRequest {
	s.AppId = &v
	return s
}

func (s *DescribeLiveGrtnDurationRequest) SetArea(v string) *DescribeLiveGrtnDurationRequest {
	s.Area = &v
	return s
}

func (s *DescribeLiveGrtnDurationRequest) SetEndTime(v string) *DescribeLiveGrtnDurationRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveGrtnDurationRequest) SetInterval(v string) *DescribeLiveGrtnDurationRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveGrtnDurationRequest) SetOwnerId(v int64) *DescribeLiveGrtnDurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveGrtnDurationRequest) SetRegionId(v string) *DescribeLiveGrtnDurationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveGrtnDurationRequest) SetStartTime(v string) *DescribeLiveGrtnDurationRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveGrtnDurationRequest) Validate() error {
	return dara.Validate(s)
}
