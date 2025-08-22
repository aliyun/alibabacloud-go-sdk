// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnBgpBpsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DescribeDcdnBgpBpsDataRequest
	GetDeviceName() *string
	SetDevicePort(v string) *DescribeDcdnBgpBpsDataRequest
	GetDevicePort() *string
	SetEndTime(v string) *DescribeDcdnBgpBpsDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDcdnBgpBpsDataRequest
	GetInterval() *string
	SetIsp(v string) *DescribeDcdnBgpBpsDataRequest
	GetIsp() *string
	SetStartTime(v string) *DescribeDcdnBgpBpsDataRequest
	GetStartTime() *string
}

type DescribeDcdnBgpBpsDataRequest struct {
	// The name of the device. If you specify this parameter, the data of the device is returned. Otherwise, the data of all devices is returned.
	//
	// example:
	//
	// devicename
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The port of the device. If you specify this parameter, the data of the port is returned. Otherwise, the data of all ports is returned. This parameter takes effect only when the **DeviceName*	- parameter is specified.
	//
	// example:
	//
	// 123
	DevicePort *string `json:"DevicePort,omitempty" xml:"DevicePort,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-11-30T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data collection interval. Unit: seconds. Valid values: 300 and 3600. Default value: 300. The default value of 300 seconds is equal to 5 minutes. The value of this parameter varies based on the time range from the specified start time to the specified end time.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The ISPs. If you need to specify multiple ISPs, separate them with commas (,). If you specify multiple ISPs, the data for the ISPs is aggregated. If you do not specify this parameter, the operation returns the data for all the ISPs.
	//
	// Valid values:
	//
	// 	- cu: China Unicom
	//
	// 	- cmi: China Mobile
	//
	// 	- ct: China Telecom
	//
	// example:
	//
	// cu
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The minimum data collection interval is an hour.
	//
	// If you do not set this parameter, data collected in the last 24 hours is queried.
	//
	// example:
	//
	// 2018-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnBgpBpsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBgpBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpBpsDataRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DescribeDcdnBgpBpsDataRequest) GetDevicePort() *string {
	return s.DevicePort
}

func (s *DescribeDcdnBgpBpsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnBgpBpsDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDcdnBgpBpsDataRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeDcdnBgpBpsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnBgpBpsDataRequest) SetDeviceName(v string) *DescribeDcdnBgpBpsDataRequest {
	s.DeviceName = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) SetDevicePort(v string) *DescribeDcdnBgpBpsDataRequest {
	s.DevicePort = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) SetEndTime(v string) *DescribeDcdnBgpBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) SetInterval(v string) *DescribeDcdnBgpBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) SetIsp(v string) *DescribeDcdnBgpBpsDataRequest {
	s.Isp = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) SetStartTime(v string) *DescribeDcdnBgpBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataRequest) Validate() error {
	return dara.Validate(s)
}
