// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *QueryDeviceStatRequest
	GetAppKey() *int64
	SetDeviceType(v string) *QueryDeviceStatRequest
	GetDeviceType() *string
	SetEndTime(v string) *QueryDeviceStatRequest
	GetEndTime() *string
	SetQueryType(v string) *QueryDeviceStatRequest
	GetQueryType() *string
	SetStartTime(v string) *QueryDeviceStatRequest
	GetStartTime() *string
}

type QueryDeviceStatRequest struct {
	// AppKey information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 	Notice:
	//
	// This parameter is only valid for early Android and iOS dual-platform application types. If your application is a dual-platform application, specify this parameter as iOS or ANDROID to query the number of devices for each type. By default, it queries ALL types.
	//
	//
	//
	// The device type. Valid values:
	//
	// - **iOS**: iOS devices
	//
	// - **ANDROID**: Android devices
	//
	// - **ALL**: All device types
	//
	// example:
	//
	// iOS
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The end time of the query. The time format follows the ISO8601 standard and uses UTC time, in the format YYYY-MM-DDThh:mm:ssZ.
	//
	// > The statistics end date is the end time\\"s day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-07-29T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Query new devices or historical cumulative devices. Valid values:
	//
	// - **NEW**: New devices
	//
	// - **TOTAL**: Cumulative devices
	//
	// This parameter is required.
	//
	// example:
	//
	// TOTAL
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The start time of the query. The time format follows the ISO8601 standard and uses UTC time, in the format YYYY-MM-DDThh:mm:ssZ.
	//
	// > The statistics start date is 00:00 UTC+8 on the start time\\"s day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-07-28T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryDeviceStatRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryDeviceStatRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *QueryDeviceStatRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryDeviceStatRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryDeviceStatRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryDeviceStatRequest) SetAppKey(v int64) *QueryDeviceStatRequest {
	s.AppKey = &v
	return s
}

func (s *QueryDeviceStatRequest) SetDeviceType(v string) *QueryDeviceStatRequest {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceStatRequest) SetEndTime(v string) *QueryDeviceStatRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceStatRequest) SetQueryType(v string) *QueryDeviceStatRequest {
	s.QueryType = &v
	return s
}

func (s *QueryDeviceStatRequest) SetStartTime(v string) *QueryDeviceStatRequest {
	s.StartTime = &v
	return s
}

func (s *QueryDeviceStatRequest) Validate() error {
	return dara.Validate(s)
}
