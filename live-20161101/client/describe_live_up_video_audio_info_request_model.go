// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUpVideoAudioInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeLiveUpVideoAudioInfoRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveUpVideoAudioInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveUpVideoAudioInfoRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveUpVideoAudioInfoRequest
	GetStartTime() *string
	SetStream(v string) *DescribeLiveUpVideoAudioInfoRequest
	GetStream() *string
}

type DescribeLiveUpVideoAudioInfoRequest struct {
	// The end time. The format is \\`yyyy-MM-ddTHH:mm:ssZ\\` in UTC.
	//
	// example:
	//
	// 2015-12-10T15:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. The format is \\`yyyy-MM-ddTHH:mm:ssZ\\` in UTC.
	//
	// example:
	//
	// 2015-12-10T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The full name of the stream to query. The format is `rtmp://<ingest domain>/<AppName>/<StreamName>`
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://example.com/AppName/StreamName
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveUpVideoAudioInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveUpVideoAudioInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveUpVideoAudioInfoRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveUpVideoAudioInfoRequest) GetStream() *string {
	return s.Stream
}

func (s *DescribeLiveUpVideoAudioInfoRequest) SetEndTime(v string) *DescribeLiveUpVideoAudioInfoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoRequest) SetOwnerId(v int64) *DescribeLiveUpVideoAudioInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoRequest) SetRegionId(v string) *DescribeLiveUpVideoAudioInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoRequest) SetStartTime(v string) *DescribeLiveUpVideoAudioInfoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoRequest) SetStream(v string) *DescribeLiveUpVideoAudioInfoRequest {
	s.Stream = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoRequest) Validate() error {
	return dara.Validate(s)
}
