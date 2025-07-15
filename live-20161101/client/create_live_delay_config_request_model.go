// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveDelayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *CreateLiveDelayConfigRequest
	GetApp() *string
	SetDelayTime(v int32) *CreateLiveDelayConfigRequest
	GetDelayTime() *int32
	SetDomain(v string) *CreateLiveDelayConfigRequest
	GetDomain() *string
	SetOwnerId(v int64) *CreateLiveDelayConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateLiveDelayConfigRequest
	GetRegionId() *string
	SetStream(v string) *CreateLiveDelayConfigRequest
	GetStream() *string
	SetTaskTriggerMode(v string) *CreateLiveDelayConfigRequest
	GetTaskTriggerMode() *string
}

type CreateLiveDelayConfigRequest struct {
	// The name of the application to which the live stream belongs. You can specify an asterisk (\\*) as the value to match all applications that belong to the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The duration for which the playback of the live stream is delayed. The value must be an integer. Valid values: 16 to 3600. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. You can use the wildcard (\\*) to specify all streams of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// The trigger mode. Valid values:
	//
	// 	- **PUBLISH_ONLY**: Stream delay can be triggered only by specifying the stream delay parameter in the ingest URL.
	//
	// 	- **CONFIG_ONLY**: Stream delay can be triggered only by the stream delay configuration.
	//
	// 	- **PUBLISH_CONFIG**: Stream delay can be triggered by the stream delay parameter in the ingest URL or the stream delay configuration. The stream delay parameter takes precedence over the stream delay configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// PUBLISH_ONLY
	TaskTriggerMode *string `json:"TaskTriggerMode,omitempty" xml:"TaskTriggerMode,omitempty"`
}

func (s CreateLiveDelayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveDelayConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveDelayConfigRequest) GetApp() *string {
	return s.App
}

func (s *CreateLiveDelayConfigRequest) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *CreateLiveDelayConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateLiveDelayConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLiveDelayConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLiveDelayConfigRequest) GetStream() *string {
	return s.Stream
}

func (s *CreateLiveDelayConfigRequest) GetTaskTriggerMode() *string {
	return s.TaskTriggerMode
}

func (s *CreateLiveDelayConfigRequest) SetApp(v string) *CreateLiveDelayConfigRequest {
	s.App = &v
	return s
}

func (s *CreateLiveDelayConfigRequest) SetDelayTime(v int32) *CreateLiveDelayConfigRequest {
	s.DelayTime = &v
	return s
}

func (s *CreateLiveDelayConfigRequest) SetDomain(v string) *CreateLiveDelayConfigRequest {
	s.Domain = &v
	return s
}

func (s *CreateLiveDelayConfigRequest) SetOwnerId(v int64) *CreateLiveDelayConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLiveDelayConfigRequest) SetRegionId(v string) *CreateLiveDelayConfigRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLiveDelayConfigRequest) SetStream(v string) *CreateLiveDelayConfigRequest {
	s.Stream = &v
	return s
}

func (s *CreateLiveDelayConfigRequest) SetTaskTriggerMode(v string) *CreateLiveDelayConfigRequest {
	s.TaskTriggerMode = &v
	return s
}

func (s *CreateLiveDelayConfigRequest) Validate() error {
	return dara.Validate(s)
}
