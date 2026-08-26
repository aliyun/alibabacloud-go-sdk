// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDelayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DescribeLiveDelayConfigResponseBody
	GetApp() *string
	SetDelayTime(v string) *DescribeLiveDelayConfigResponseBody
	GetDelayTime() *string
	SetDomain(v string) *DescribeLiveDelayConfigResponseBody
	GetDomain() *string
	SetRequestId(v string) *DescribeLiveDelayConfigResponseBody
	GetRequestId() *string
	SetStream(v string) *DescribeLiveDelayConfigResponseBody
	GetStream() *string
	SetTaskTriggerMode(v string) *DescribeLiveDelayConfigResponseBody
	GetTaskTriggerMode() *string
}

type DescribeLiveDelayConfigResponseBody struct {
	// The application name.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The playback latency of the stream.
	//
	// example:
	//
	// 60
	DelayTime *string `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3136B58-5876-4168-83CA-B562781981A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stream name.
	//
	// example:
	//
	// liveStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// The trigger mode for the task. Valid values:
	//
	// - **PUBLISH_ONLY**: The task is triggered only when stream ingest parameters for delayed playback are specified.
	//
	// - **CONFIG_ONLY**: The task is triggered only by the configuration. Stream ingest parameters are ignored.
	//
	// - **PUBLISH_CONFIG**: The task can be triggered by both stream ingest parameters and the configuration. Stream ingest parameters have a higher priority than the configuration.
	//
	// example:
	//
	// PUBLISH_ONLY
	TaskTriggerMode *string `json:"TaskTriggerMode,omitempty" xml:"TaskTriggerMode,omitempty"`
}

func (s DescribeLiveDelayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDelayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDelayConfigResponseBody) GetApp() *string {
	return s.App
}

func (s *DescribeLiveDelayConfigResponseBody) GetDelayTime() *string {
	return s.DelayTime
}

func (s *DescribeLiveDelayConfigResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveDelayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDelayConfigResponseBody) GetStream() *string {
	return s.Stream
}

func (s *DescribeLiveDelayConfigResponseBody) GetTaskTriggerMode() *string {
	return s.TaskTriggerMode
}

func (s *DescribeLiveDelayConfigResponseBody) SetApp(v string) *DescribeLiveDelayConfigResponseBody {
	s.App = &v
	return s
}

func (s *DescribeLiveDelayConfigResponseBody) SetDelayTime(v string) *DescribeLiveDelayConfigResponseBody {
	s.DelayTime = &v
	return s
}

func (s *DescribeLiveDelayConfigResponseBody) SetDomain(v string) *DescribeLiveDelayConfigResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeLiveDelayConfigResponseBody) SetRequestId(v string) *DescribeLiveDelayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDelayConfigResponseBody) SetStream(v string) *DescribeLiveDelayConfigResponseBody {
	s.Stream = &v
	return s
}

func (s *DescribeLiveDelayConfigResponseBody) SetTaskTriggerMode(v string) *DescribeLiveDelayConfigResponseBody {
	s.TaskTriggerMode = &v
	return s
}

func (s *DescribeLiveDelayConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
