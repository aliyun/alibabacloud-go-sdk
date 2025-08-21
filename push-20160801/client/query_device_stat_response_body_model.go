// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppDeviceStats(v *QueryDeviceStatResponseBodyAppDeviceStats) *QueryDeviceStatResponseBody
	GetAppDeviceStats() *QueryDeviceStatResponseBodyAppDeviceStats
	SetRequestId(v string) *QueryDeviceStatResponseBody
	GetRequestId() *string
}

type QueryDeviceStatResponseBody struct {
	AppDeviceStats *QueryDeviceStatResponseBodyAppDeviceStats `json:"AppDeviceStats,omitempty" xml:"AppDeviceStats,omitempty" type:"Struct"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDeviceStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatResponseBody) GetAppDeviceStats() *QueryDeviceStatResponseBodyAppDeviceStats {
	return s.AppDeviceStats
}

func (s *QueryDeviceStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDeviceStatResponseBody) SetAppDeviceStats(v *QueryDeviceStatResponseBodyAppDeviceStats) *QueryDeviceStatResponseBody {
	s.AppDeviceStats = v
	return s
}

func (s *QueryDeviceStatResponseBody) SetRequestId(v string) *QueryDeviceStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceStatResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDeviceStatResponseBodyAppDeviceStats struct {
	AppDeviceStat []*QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat `json:"AppDeviceStat,omitempty" xml:"AppDeviceStat,omitempty" type:"Repeated"`
}

func (s QueryDeviceStatResponseBodyAppDeviceStats) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatResponseBodyAppDeviceStats) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatResponseBodyAppDeviceStats) GetAppDeviceStat() []*QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	return s.AppDeviceStat
}

func (s *QueryDeviceStatResponseBodyAppDeviceStats) SetAppDeviceStat(v []*QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) *QueryDeviceStatResponseBodyAppDeviceStats {
	s.AppDeviceStat = v
	return s
}

func (s *QueryDeviceStatResponseBodyAppDeviceStats) Validate() error {
	return dara.Validate(s)
}

type QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat struct {
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// iOS
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 2016-07-28T16:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) GetCount() *int64 {
	return s.Count
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) GetDeviceType() *string {
	return s.DeviceType
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) GetTime() *string {
	return s.Time
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetCount(v int64) *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.Count = &v
	return s
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetDeviceType(v string) *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetTime(v string) *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.Time = &v
	return s
}

func (s *QueryDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) Validate() error {
	return dara.Validate(s)
}
