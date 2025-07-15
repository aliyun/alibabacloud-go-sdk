// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeToutiaoLivePlayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DescribeToutiaoLivePlayRequest
	GetApp() *string
	SetDomain(v string) *DescribeToutiaoLivePlayRequest
	GetDomain() *string
	SetEndTime(v string) *DescribeToutiaoLivePlayRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeToutiaoLivePlayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeToutiaoLivePlayRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeToutiaoLivePlayRequest
	GetStartTime() *string
	SetStream(v string) *DescribeToutiaoLivePlayRequest
	GetStream() *string
}

type DescribeToutiaoLivePlayRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time span cannot be greater than 10 hours. If you do not configure StartTime and EndTime, the data within the previous hour is queried. Specify the time in the ISO 8601 standard. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T20:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s DescribeToutiaoLivePlayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeToutiaoLivePlayRequest) GoString() string {
	return s.String()
}

func (s *DescribeToutiaoLivePlayRequest) GetApp() *string {
	return s.App
}

func (s *DescribeToutiaoLivePlayRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeToutiaoLivePlayRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeToutiaoLivePlayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeToutiaoLivePlayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeToutiaoLivePlayRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeToutiaoLivePlayRequest) GetStream() *string {
	return s.Stream
}

func (s *DescribeToutiaoLivePlayRequest) SetApp(v string) *DescribeToutiaoLivePlayRequest {
	s.App = &v
	return s
}

func (s *DescribeToutiaoLivePlayRequest) SetDomain(v string) *DescribeToutiaoLivePlayRequest {
	s.Domain = &v
	return s
}

func (s *DescribeToutiaoLivePlayRequest) SetEndTime(v string) *DescribeToutiaoLivePlayRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeToutiaoLivePlayRequest) SetOwnerId(v int64) *DescribeToutiaoLivePlayRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeToutiaoLivePlayRequest) SetRegionId(v string) *DescribeToutiaoLivePlayRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeToutiaoLivePlayRequest) SetStartTime(v string) *DescribeToutiaoLivePlayRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeToutiaoLivePlayRequest) SetStream(v string) *DescribeToutiaoLivePlayRequest {
	s.Stream = &v
	return s
}

func (s *DescribeToutiaoLivePlayRequest) Validate() error {
	return dara.Validate(s)
}
