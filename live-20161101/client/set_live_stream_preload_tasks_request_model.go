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
	// The acceleration region where you want to prefetch the live content. Valid values:
	//
	// 	- domestic: regions in the Chinese mainland.
	//
	// 	- overseas: regions outside the Chinese mainland.
	//
	// 	- global: regions in and outside the Chinese mainland.
	//
	// If you do not specify this parameter, the acceleration region configured for the domain name is used.
	//
	// example:
	//
	// domestic
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The streaming domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The streaming URL. You can specify up to 100 streaming URLs in a request. Separate multiple streaming URLs with commas (,).
	//
	// This parameter is required.
	PlayUrl *string `json:"PlayUrl,omitempty" xml:"PlayUrl,omitempty"`
	// The end time of the prefetch task. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2016-06-30T19:00:00Z. The interval between the start time and end time cannot exceed 6 hours.
	//
	// example:
	//
	// 2016-06-30T19:00:00Z
	PreloadedEndTime *string `json:"PreloadedEndTime,omitempty" xml:"PreloadedEndTime,omitempty"`
	// The start time of the prefetch task. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Example: 2016-06-29T19:00:00Z. If you do not specify this parameter, the prefetch task runs for 1 hour by default.
	//
	// example:
	//
	// 2016-06-29T19:00:00Z
	PreloadedStartTime *string `json:"PreloadedStartTime,omitempty" xml:"PreloadedStartTime,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
