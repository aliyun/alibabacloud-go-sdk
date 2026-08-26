// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamPreloadTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *SetLiveStreamPreloadTasksRequest
	GetArea() *string
	SetDomainName(v string) *SetLiveStreamPreloadTasksRequest
	GetDomainName() *string
	SetOwnerId(v int64) *SetLiveStreamPreloadTasksRequest
	GetOwnerId() *int64
	SetPlayUrl(v string) *SetLiveStreamPreloadTasksRequest
	GetPlayUrl() *string
	SetPreloadedEndTime(v string) *SetLiveStreamPreloadTasksRequest
	GetPreloadedEndTime() *string
	SetPreloadedStartTime(v string) *SetLiveStreamPreloadTasksRequest
	GetPreloadedStartTime() *string
	SetRegionId(v string) *SetLiveStreamPreloadTasksRequest
	GetRegionId() *string
}

type SetLiveStreamPreloadTasksRequest struct {
	// The prefetch area. Valid values:
	//
	// - domestic: the Chinese mainland.
	//
	// - overseas: outside the Chinese mainland, including Hong Kong (China), Macao (China), and Taiwan (China).
	//
	// - global: global acceleration.
	//
	//
	//
	// If you do not specify this parameter, the default prefetch area is the acceleration region configured for your domain name.
	//
	// example:
	//
	// domestic
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The live stream URLs. You can specify multiple URLs separated by commas (,). A maximum of 100 URLs can be specified.
	//
	// This parameter is required.
	PlayUrl *string `json:"PlayUrl,omitempty" xml:"PlayUrl,omitempty"`
	// The end time of the prefetch task in UTC. Example: 2016-06-30T19:00:00Z. The interval between EndTime and StartTime cannot exceed 6 hours.
	//
	// example:
	//
	// 2016-06-30T19:00:00Z
	PreloadedEndTime *string `json:"PreloadedEndTime,omitempty" xml:"PreloadedEndTime,omitempty"`
	// The start time of the prefetch task in UTC. Example: 2016-06-29T19:00:00Z. If you do not specify this parameter, the default prefetch duration is 1 hour.
	//
	// example:
	//
	// 2016-06-29T19:00:00Z
	PreloadedStartTime *string `json:"PreloadedStartTime,omitempty" xml:"PreloadedStartTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetLiveStreamPreloadTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamPreloadTasksRequest) GoString() string {
	return s.String()
}

func (s *SetLiveStreamPreloadTasksRequest) GetArea() *string {
	return s.Area
}

func (s *SetLiveStreamPreloadTasksRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetLiveStreamPreloadTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveStreamPreloadTasksRequest) GetPlayUrl() *string {
	return s.PlayUrl
}

func (s *SetLiveStreamPreloadTasksRequest) GetPreloadedEndTime() *string {
	return s.PreloadedEndTime
}

func (s *SetLiveStreamPreloadTasksRequest) GetPreloadedStartTime() *string {
	return s.PreloadedStartTime
}

func (s *SetLiveStreamPreloadTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLiveStreamPreloadTasksRequest) SetArea(v string) *SetLiveStreamPreloadTasksRequest {
	s.Area = &v
	return s
}

func (s *SetLiveStreamPreloadTasksRequest) SetDomainName(v string) *SetLiveStreamPreloadTasksRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveStreamPreloadTasksRequest) SetOwnerId(v int64) *SetLiveStreamPreloadTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveStreamPreloadTasksRequest) SetPlayUrl(v string) *SetLiveStreamPreloadTasksRequest {
	s.PlayUrl = &v
	return s
}

func (s *SetLiveStreamPreloadTasksRequest) SetPreloadedEndTime(v string) *SetLiveStreamPreloadTasksRequest {
	s.PreloadedEndTime = &v
	return s
}

func (s *SetLiveStreamPreloadTasksRequest) SetPreloadedStartTime(v string) *SetLiveStreamPreloadTasksRequest {
	s.PreloadedStartTime = &v
	return s
}

func (s *SetLiveStreamPreloadTasksRequest) SetRegionId(v string) *SetLiveStreamPreloadTasksRequest {
	s.RegionId = &v
	return s
}

func (s *SetLiveStreamPreloadTasksRequest) Validate() error {
	return dara.Validate(s)
}
