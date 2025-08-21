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
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// iOS
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2016-07-29T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TOTAL
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
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
