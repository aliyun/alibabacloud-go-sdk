// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveMPUTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateLiveMPUTaskShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *UpdateLiveMPUTaskShrinkRequest
	GetChannelId() *string
	SetMixMode(v string) *UpdateLiveMPUTaskShrinkRequest
	GetMixMode() *string
	SetMultiStreamURLShrink(v string) *UpdateLiveMPUTaskShrinkRequest
	GetMultiStreamURLShrink() *string
	SetSeiParamsShrink(v string) *UpdateLiveMPUTaskShrinkRequest
	GetSeiParamsShrink() *string
	SetSingleSubParamsShrink(v string) *UpdateLiveMPUTaskShrinkRequest
	GetSingleSubParamsShrink() *string
	SetStreamURL(v string) *UpdateLiveMPUTaskShrinkRequest
	GetStreamURL() *string
	SetTaskId(v string) *UpdateLiveMPUTaskShrinkRequest
	GetTaskId() *string
	SetTranscodeParamsShrink(v string) *UpdateLiveMPUTaskShrinkRequest
	GetTranscodeParamsShrink() *string
}

type UpdateLiveMPUTaskShrinkRequest struct {
	// The application ID. You can specify only one application ID. The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel ID. You can specify only one channel ID. The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The stream mixing mode. Valid values:
	//
	// 	- **0**: the single-stream relay mode. In this mode, the service only relays the original single stream, but does not transcode mixed streams. You do not need to set parameters for mixed-stream transcoding.
	//
	// 	- **1*	- (default): the mixed-stream relay mode.
	//
	// example:
	//
	// 0
	MixMode *string `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	// The multiple ingest URLs to relay. This parameter allows you to specify multiple ingest URLs.
	MultiStreamURLShrink *string `json:"MultiStreamURL,omitempty" xml:"MultiStreamURL,omitempty"`
	// The supplemental enhancement information (SEI) parameters.
	SeiParamsShrink *string `json:"SeiParams,omitempty" xml:"SeiParams,omitempty"`
	// The single-stream relay parameters. These parameters are required if you set MixMode to 0.
	SingleSubParamsShrink *string `json:"SingleSubParams,omitempty" xml:"SingleSubParams,omitempty"`
	// The ingest URL. You can specify only one ingest URL in the Real-Time Messaging Protocol (RTMP) format. The URL can be up to 2,048 characters in length. For information about the generation rules of ingest URLs, see [Ingest and streaming URLs](https://help.aliyun.com/document_detail/199339.html).
	//
	// >
	//
	// 	- If the ingest URL is under a domain name for which hotlink protection is enabled, you must include an access token in the URL.
	//
	// 	- You cannot use the same ingest URL in different tasks.
	//
	// 	- You cannot use the same ingest URL within 10 seconds after a task is stopped.
	//
	// example:
	//
	// rtmp://example.com/live/stream
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// The task ID. You can specify only one task ID. The ID can be up to 55 characters in length and can contain letters, digits, underscores (_), and hyphens (-). The ID must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The mixed-stream relay parameters. These parameters are required if you set MixMode to 1.
	TranscodeParamsShrink *string `json:"TranscodeParams,omitempty" xml:"TranscodeParams,omitempty"`
}

func (s UpdateLiveMPUTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateLiveMPUTaskShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateLiveMPUTaskShrinkRequest) GetMixMode() *string {
	return s.MixMode
}

func (s *UpdateLiveMPUTaskShrinkRequest) GetMultiStreamURLShrink() *string {
	return s.MultiStreamURLShrink
}

func (s *UpdateLiveMPUTaskShrinkRequest) GetSeiParamsShrink() *string {
	return s.SeiParamsShrink
}

func (s *UpdateLiveMPUTaskShrinkRequest) GetSingleSubParamsShrink() *string {
	return s.SingleSubParamsShrink
}

func (s *UpdateLiveMPUTaskShrinkRequest) GetStreamURL() *string {
	return s.StreamURL
}

func (s *UpdateLiveMPUTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateLiveMPUTaskShrinkRequest) GetTranscodeParamsShrink() *string {
	return s.TranscodeParamsShrink
}

func (s *UpdateLiveMPUTaskShrinkRequest) SetAppId(v string) *UpdateLiveMPUTaskShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateLiveMPUTaskShrinkRequest) SetChannelId(v string) *UpdateLiveMPUTaskShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateLiveMPUTaskShrinkRequest) SetMixMode(v string) *UpdateLiveMPUTaskShrinkRequest {
	s.MixMode = &v
	return s
}

func (s *UpdateLiveMPUTaskShrinkRequest) SetMultiStreamURLShrink(v string) *UpdateLiveMPUTaskShrinkRequest {
	s.MultiStreamURLShrink = &v
	return s
}

func (s *UpdateLiveMPUTaskShrinkRequest) SetSeiParamsShrink(v string) *UpdateLiveMPUTaskShrinkRequest {
	s.SeiParamsShrink = &v
	return s
}

func (s *UpdateLiveMPUTaskShrinkRequest) SetSingleSubParamsShrink(v string) *UpdateLiveMPUTaskShrinkRequest {
	s.SingleSubParamsShrink = &v
	return s
}

func (s *UpdateLiveMPUTaskShrinkRequest) SetStreamURL(v string) *UpdateLiveMPUTaskShrinkRequest {
	s.StreamURL = &v
	return s
}

func (s *UpdateLiveMPUTaskShrinkRequest) SetTaskId(v string) *UpdateLiveMPUTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateLiveMPUTaskShrinkRequest) SetTranscodeParamsShrink(v string) *UpdateLiveMPUTaskShrinkRequest {
	s.TranscodeParamsShrink = &v
	return s
}

func (s *UpdateLiveMPUTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
