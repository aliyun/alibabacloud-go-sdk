// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcDurationDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDurationDataPerInterval(v *DescribeRtcDurationDataResponseBodyDurationDataPerInterval) *DescribeRtcDurationDataResponseBody
	GetDurationDataPerInterval() *DescribeRtcDurationDataResponseBodyDurationDataPerInterval
	SetRequestId(v string) *DescribeRtcDurationDataResponseBody
	GetRequestId() *string
}

type DescribeRtcDurationDataResponseBody struct {
	DurationDataPerInterval *DescribeRtcDurationDataResponseBodyDurationDataPerInterval `json:"DurationDataPerInterval,omitempty" xml:"DurationDataPerInterval,omitempty" type:"Struct"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRtcDurationDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcDurationDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponseBody) GetDurationDataPerInterval() *DescribeRtcDurationDataResponseBodyDurationDataPerInterval {
	return s.DurationDataPerInterval
}

func (s *DescribeRtcDurationDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRtcDurationDataResponseBody) SetDurationDataPerInterval(v *DescribeRtcDurationDataResponseBodyDurationDataPerInterval) *DescribeRtcDurationDataResponseBody {
	s.DurationDataPerInterval = v
	return s
}

func (s *DescribeRtcDurationDataResponseBody) SetRequestId(v string) *DescribeRtcDurationDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBody) Validate() error {
	if s.DurationDataPerInterval != nil {
		if err := s.DurationDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRtcDurationDataResponseBodyDurationDataPerInterval struct {
	DurationModule []*DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule `json:"DurationModule,omitempty" xml:"DurationModule,omitempty" type:"Repeated"`
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerInterval) GetDurationModule() []*DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	return s.DurationModule
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerInterval) SetDurationModule(v []*DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) *DescribeRtcDurationDataResponseBodyDurationDataPerInterval {
	s.DurationModule = v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerInterval) Validate() error {
	if s.DurationModule != nil {
		for _, item := range s.DurationModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule struct {
	// example:
	//
	// 200
	AudioDuration *int64 `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	// example:
	//
	// 200
	ContentDuration *int64 `json:"ContentDuration,omitempty" xml:"ContentDuration,omitempty"`
	// example:
	//
	// 2020-02-04T05:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 1000
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// example:
	//
	// 300
	V1080Duration *int64 `json:"V1080Duration,omitempty" xml:"V1080Duration,omitempty"`
	// example:
	//
	// 300
	V360Duration *int64 `json:"V360Duration,omitempty" xml:"V360Duration,omitempty"`
	// example:
	//
	// 200
	V720Duration *int64 `json:"V720Duration,omitempty" xml:"V720Duration,omitempty"`
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GetAudioDuration() *int64 {
	return s.AudioDuration
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GetContentDuration() *int64 {
	return s.ContentDuration
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GetV1080Duration() *int64 {
	return s.V1080Duration
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GetV360Duration() *int64 {
	return s.V360Duration
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GetV720Duration() *int64 {
	return s.V720Duration
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetAudioDuration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.AudioDuration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetContentDuration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.ContentDuration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetTimeStamp(v string) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetTotalDuration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.TotalDuration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetV1080Duration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.V1080Duration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetV360Duration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.V360Duration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetV720Duration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.V720Duration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) Validate() error {
	return dara.Validate(s)
}
