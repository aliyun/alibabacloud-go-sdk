// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeToutiaoLivePublishRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DescribeToutiaoLivePublishRequest
	GetApp() *string
	SetDomain(v string) *DescribeToutiaoLivePublishRequest
	GetDomain() *string
	SetEndTime(v string) *DescribeToutiaoLivePublishRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeToutiaoLivePublishRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeToutiaoLivePublishRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeToutiaoLivePublishRequest
	GetStartTime() *string
	SetStream(v string) *DescribeToutiaoLivePublishRequest
	GetStream() *string
}

type DescribeToutiaoLivePublishRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The end time must be later than the start time. The maximum time range that can be specified is 10 hours. If you specify neither StartTime nor EndTime, the data of the last hour is queried by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T21:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-10-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the ingested stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s DescribeToutiaoLivePublishRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeToutiaoLivePublishRequest) GoString() string {
	return s.String()
}

func (s *DescribeToutiaoLivePublishRequest) GetApp() *string {
	return s.App
}

func (s *DescribeToutiaoLivePublishRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeToutiaoLivePublishRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeToutiaoLivePublishRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeToutiaoLivePublishRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeToutiaoLivePublishRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeToutiaoLivePublishRequest) GetStream() *string {
	return s.Stream
}

func (s *DescribeToutiaoLivePublishRequest) SetApp(v string) *DescribeToutiaoLivePublishRequest {
	s.App = &v
	return s
}

func (s *DescribeToutiaoLivePublishRequest) SetDomain(v string) *DescribeToutiaoLivePublishRequest {
	s.Domain = &v
	return s
}

func (s *DescribeToutiaoLivePublishRequest) SetEndTime(v string) *DescribeToutiaoLivePublishRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeToutiaoLivePublishRequest) SetOwnerId(v int64) *DescribeToutiaoLivePublishRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeToutiaoLivePublishRequest) SetRegionId(v string) *DescribeToutiaoLivePublishRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeToutiaoLivePublishRequest) SetStartTime(v string) *DescribeToutiaoLivePublishRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeToutiaoLivePublishRequest) SetStream(v string) *DescribeToutiaoLivePublishRequest {
	s.Stream = &v
	return s
}

func (s *DescribeToutiaoLivePublishRequest) Validate() error {
	return dara.Validate(s)
}
