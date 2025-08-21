// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUniqueDeviceStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *QueryUniqueDeviceStatRequest
	GetAppKey() *int64
	SetEndTime(v string) *QueryUniqueDeviceStatRequest
	GetEndTime() *string
	SetGranularity(v string) *QueryUniqueDeviceStatRequest
	GetGranularity() *string
	SetStartTime(v string) *QueryUniqueDeviceStatRequest
	GetStartTime() *string
}

type QueryUniqueDeviceStatRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2016-07-26T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DAY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2016-07-25T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryUniqueDeviceStatRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUniqueDeviceStatRequest) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryUniqueDeviceStatRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryUniqueDeviceStatRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *QueryUniqueDeviceStatRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryUniqueDeviceStatRequest) SetAppKey(v int64) *QueryUniqueDeviceStatRequest {
	s.AppKey = &v
	return s
}

func (s *QueryUniqueDeviceStatRequest) SetEndTime(v string) *QueryUniqueDeviceStatRequest {
	s.EndTime = &v
	return s
}

func (s *QueryUniqueDeviceStatRequest) SetGranularity(v string) *QueryUniqueDeviceStatRequest {
	s.Granularity = &v
	return s
}

func (s *QueryUniqueDeviceStatRequest) SetStartTime(v string) *QueryUniqueDeviceStatRequest {
	s.StartTime = &v
	return s
}

func (s *QueryUniqueDeviceStatRequest) Validate() error {
	return dara.Validate(s)
}
