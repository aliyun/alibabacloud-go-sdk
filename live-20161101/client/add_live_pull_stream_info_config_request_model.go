// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLivePullStreamInfoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLivePullStreamInfoConfigRequest
	GetAppName() *string
	SetDomainName(v string) *AddLivePullStreamInfoConfigRequest
	GetDomainName() *string
	SetEndTime(v string) *AddLivePullStreamInfoConfigRequest
	GetEndTime() *string
	SetOwnerId(v int64) *AddLivePullStreamInfoConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddLivePullStreamInfoConfigRequest
	GetRegionId() *string
	SetSourceUrl(v string) *AddLivePullStreamInfoConfigRequest
	GetSourceUrl() *string
	SetStartTime(v string) *AddLivePullStreamInfoConfigRequest
	GetStartTime() *string
	SetStreamName(v string) *AddLivePullStreamInfoConfigRequest
	GetStreamName() *string
}

type AddLivePullStreamInfoConfigRequest struct {
	// The name of the application to which the live stream belongs. This parameter is determined by you.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The domain name used for stream pulling. It is the main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The time when stream pulling ends.
	//
	// The time range specified by the StartTime and EndTime parameters cannot exceed seven days. The time specified by the EndTime parameter must be later than the current time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The origin URL of the live stream. You can specify multiple URLs. Separate them with semicolons (;).
	//
	// >  You can pull only live streams in the Real-Time Messaging Protocol (RTMP), Flash Video (FLV), HTTP Live Streaming (HLS), and Secure Reliable Transport (SRT) formats.
	//
	// This parameter is required.
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	// The time when stream pulling starts.
	//
	// The time range specified by the StartTime and EndTime parameters cannot exceed seven days. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream. This parameter is determined by you.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s AddLivePullStreamInfoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLivePullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLivePullStreamInfoConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLivePullStreamInfoConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLivePullStreamInfoConfigRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *AddLivePullStreamInfoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLivePullStreamInfoConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLivePullStreamInfoConfigRequest) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *AddLivePullStreamInfoConfigRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddLivePullStreamInfoConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddLivePullStreamInfoConfigRequest) SetAppName(v string) *AddLivePullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetDomainName(v string) *AddLivePullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetEndTime(v string) *AddLivePullStreamInfoConfigRequest {
	s.EndTime = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetOwnerId(v int64) *AddLivePullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetRegionId(v string) *AddLivePullStreamInfoConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetSourceUrl(v string) *AddLivePullStreamInfoConfigRequest {
	s.SourceUrl = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetStartTime(v string) *AddLivePullStreamInfoConfigRequest {
	s.StartTime = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) SetStreamName(v string) *AddLivePullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLivePullStreamInfoConfigRequest) Validate() error {
	return dara.Validate(s)
}
