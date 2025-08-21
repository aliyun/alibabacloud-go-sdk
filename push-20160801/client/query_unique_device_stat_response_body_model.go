// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUniqueDeviceStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppDeviceStats(v *QueryUniqueDeviceStatResponseBodyAppDeviceStats) *QueryUniqueDeviceStatResponseBody
	GetAppDeviceStats() *QueryUniqueDeviceStatResponseBodyAppDeviceStats
	SetRequestId(v string) *QueryUniqueDeviceStatResponseBody
	GetRequestId() *string
}

type QueryUniqueDeviceStatResponseBody struct {
	AppDeviceStats *QueryUniqueDeviceStatResponseBodyAppDeviceStats `json:"AppDeviceStats,omitempty" xml:"AppDeviceStats,omitempty" type:"Struct"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryUniqueDeviceStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUniqueDeviceStatResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatResponseBody) GetAppDeviceStats() *QueryUniqueDeviceStatResponseBodyAppDeviceStats {
	return s.AppDeviceStats
}

func (s *QueryUniqueDeviceStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUniqueDeviceStatResponseBody) SetAppDeviceStats(v *QueryUniqueDeviceStatResponseBodyAppDeviceStats) *QueryUniqueDeviceStatResponseBody {
	s.AppDeviceStats = v
	return s
}

func (s *QueryUniqueDeviceStatResponseBody) SetRequestId(v string) *QueryUniqueDeviceStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUniqueDeviceStatResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryUniqueDeviceStatResponseBodyAppDeviceStats struct {
	AppDeviceStat []*QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat `json:"AppDeviceStat,omitempty" xml:"AppDeviceStat,omitempty" type:"Repeated"`
}

func (s QueryUniqueDeviceStatResponseBodyAppDeviceStats) String() string {
	return dara.Prettify(s)
}

func (s QueryUniqueDeviceStatResponseBodyAppDeviceStats) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStats) GetAppDeviceStat() []*QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	return s.AppDeviceStat
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStats) SetAppDeviceStat(v []*QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) *QueryUniqueDeviceStatResponseBodyAppDeviceStats {
	s.AppDeviceStat = v
	return s
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStats) Validate() error {
	return dara.Validate(s)
}

type QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat struct {
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 2016-07-25T00:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) String() string {
	return dara.Prettify(s)
}

func (s QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) GoString() string {
	return s.String()
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) GetCount() *int64 {
	return s.Count
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) GetTime() *string {
	return s.Time
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetCount(v int64) *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.Count = &v
	return s
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) SetTime(v string) *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat {
	s.Time = &v
	return s
}

func (s *QueryUniqueDeviceStatResponseBodyAppDeviceStatsAppDeviceStat) Validate() error {
	return dara.Validate(s)
}
