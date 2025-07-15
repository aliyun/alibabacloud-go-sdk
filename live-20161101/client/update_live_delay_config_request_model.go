// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveDelayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *UpdateLiveDelayConfigRequest
	GetApp() *string
	SetDelayTime(v int32) *UpdateLiveDelayConfigRequest
	GetDelayTime() *int32
	SetDomain(v string) *UpdateLiveDelayConfigRequest
	GetDomain() *string
	SetOwnerId(v int64) *UpdateLiveDelayConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveDelayConfigRequest
	GetRegionId() *string
	SetStream(v string) *UpdateLiveDelayConfigRequest
	GetStream() *string
	SetTaskTriggerMode(v string) *UpdateLiveDelayConfigRequest
	GetTaskTriggerMode() *string
}

type UpdateLiveDelayConfigRequest struct {
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
	// example:
	//
	// PUBLISH_ONLY
	TaskTriggerMode *string `json:"TaskTriggerMode,omitempty" xml:"TaskTriggerMode,omitempty"`
}

func (s UpdateLiveDelayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveDelayConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveDelayConfigRequest) GetApp() *string {
	return s.App
}

func (s *UpdateLiveDelayConfigRequest) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *UpdateLiveDelayConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateLiveDelayConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveDelayConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveDelayConfigRequest) GetStream() *string {
	return s.Stream
}

func (s *UpdateLiveDelayConfigRequest) GetTaskTriggerMode() *string {
	return s.TaskTriggerMode
}

func (s *UpdateLiveDelayConfigRequest) SetApp(v string) *UpdateLiveDelayConfigRequest {
	s.App = &v
	return s
}

func (s *UpdateLiveDelayConfigRequest) SetDelayTime(v int32) *UpdateLiveDelayConfigRequest {
	s.DelayTime = &v
	return s
}

func (s *UpdateLiveDelayConfigRequest) SetDomain(v string) *UpdateLiveDelayConfigRequest {
	s.Domain = &v
	return s
}

func (s *UpdateLiveDelayConfigRequest) SetOwnerId(v int64) *UpdateLiveDelayConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveDelayConfigRequest) SetRegionId(v string) *UpdateLiveDelayConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveDelayConfigRequest) SetStream(v string) *UpdateLiveDelayConfigRequest {
	s.Stream = &v
	return s
}

func (s *UpdateLiveDelayConfigRequest) SetTaskTriggerMode(v string) *UpdateLiveDelayConfigRequest {
	s.TaskTriggerMode = &v
	return s
}

func (s *UpdateLiveDelayConfigRequest) Validate() error {
	return dara.Validate(s)
}
