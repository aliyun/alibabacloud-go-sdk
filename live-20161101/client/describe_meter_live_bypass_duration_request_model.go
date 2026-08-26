// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterLiveBypassDurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeMeterLiveBypassDurationRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeMeterLiveBypassDurationRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeMeterLiveBypassDurationRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeMeterLiveBypassDurationRequest
	GetStartTime() *string
}

type DescribeMeterLiveBypassDurationRequest struct {
	// The application ID. You can view the application ID on the [Application Management](https://help.aliyun.com/document_detail/2355593.html) page of ApsaraVideo Real-time Communication.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4346289a-a790-4869-9e23-22766d5e****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query. The end time must be later than the start time. The query granularity must be ≥ 5 minutes and ≤ 31 days. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
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
	// If this parameter is not specified or set to an unsupported value, the default value 3600 is used.
	//
	// example:
	//
	// 86400
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMeterLiveBypassDurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterLiveBypassDurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterLiveBypassDurationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeMeterLiveBypassDurationRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMeterLiveBypassDurationRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeMeterLiveBypassDurationRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMeterLiveBypassDurationRequest) SetAppId(v string) *DescribeMeterLiveBypassDurationRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationRequest) SetEndTime(v string) *DescribeMeterLiveBypassDurationRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationRequest) SetInterval(v string) *DescribeMeterLiveBypassDurationRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationRequest) SetStartTime(v string) *DescribeMeterLiveBypassDurationRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationRequest) Validate() error {
	return dara.Validate(s)
}
