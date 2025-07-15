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
	// The ID of the application. Separate multiple application IDs with commas (,). You can specify up to 30 application IDs. By default, the aggregated data of all applications is returned.
	//
	// example:
	//
	// 4346289a-a790-4869-9e23-22766d5e****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the billable region. Valid values:
	//
	// 	- CN: Chinese mainland
	//
	// 	- OverSeas: countries and regions outside the Chinese mainland
	//
	// 	- AP1: Asia Pacific 1, including Hong Kong (China), Macao (China), Taiwan (China), Japan, and other Southeast Asia countries and regions except Vietnam and Indonesia
	//
	// 	- AP2: Asia Pacific 2, including Indonesia, South Korea, and Vietnam
	//
	// 	- AP3: Asia Pacific 3, including Australia and New Zealand
	//
	// 	- NA: North America, including US and Canada
	//
	// 	- SA: South America, specifically meaning Brazil
	//
	// 	- EU: Europe, including Ukraine, UK, France, Netherlands, Spain, Italy, Sweden, and Germany
	//
	// 	- MEAA: Middle East and Africa, including South Africa, Oman, UAE, and Kuwait
	//
	// If you do not specify this parameter, data of all regions is aggregated and returned by default.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The end time must be later than the start time. The time range that can be specified is greater than or equal to 5 minutes and less than or equal to 31 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- 300
	//
	// 	- 3600
	//
	// 	- 86400
	//
	// If you specify an invalid value or do not specify this parameter, the default value 3600 is used.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
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
