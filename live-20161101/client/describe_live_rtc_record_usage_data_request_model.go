// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRtcRecordUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeLiveRtcRecordUsageDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeLiveRtcRecordUsageDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveRtcRecordUsageDataRequest
	GetInterval() *string
	SetRecordMode(v string) *DescribeLiveRtcRecordUsageDataRequest
	GetRecordMode() *string
	SetStartTime(v string) *DescribeLiveRtcRecordUsageDataRequest
	GetStartTime() *string
}

type DescribeLiveRtcRecordUsageDataRequest struct {
	// The ID of the ApsaraVideo Real-time Communication application. You can view the ID in [ApsaraVideo Real-time Communication application management](https://help.aliyun.com/document_detail/2355593.html). Navigate to **ApsaraVideo Live > Live+ > Real-time Communication > Application Management*	- to view your application IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// intl7f92-a5a8*************7ce4eb44a6
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time for the query. The query granularity must be ≥ 5 minutes and ≤ 31 days. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for the query data. Unit: seconds. Valid values:
	//
	// - 3600 (default).
	//
	// - 86400.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The recording mode. Valid values:
	//
	// - 0: single-stream recording mode.
	//
	// - 1: stream mixing recording mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RecordMode *string `json:"RecordMode,omitempty" xml:"RecordMode,omitempty"`
	// The start time for the query. Format: yyyy-MM-ddTHH:mm:ssZ (UTC time).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveRtcRecordUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRtcRecordUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRtcRecordUsageDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLiveRtcRecordUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveRtcRecordUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveRtcRecordUsageDataRequest) GetRecordMode() *string {
	return s.RecordMode
}

func (s *DescribeLiveRtcRecordUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveRtcRecordUsageDataRequest) SetAppId(v string) *DescribeLiveRtcRecordUsageDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataRequest) SetEndTime(v string) *DescribeLiveRtcRecordUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataRequest) SetInterval(v string) *DescribeLiveRtcRecordUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataRequest) SetRecordMode(v string) *DescribeLiveRtcRecordUsageDataRequest {
	s.RecordMode = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataRequest) SetStartTime(v string) *DescribeLiveRtcRecordUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
