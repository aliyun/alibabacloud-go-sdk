// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveMPUTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartLiveMPUTaskShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *StartLiveMPUTaskShrinkRequest
	GetChannelId() *string
	SetMaxIdleTime(v string) *StartLiveMPUTaskShrinkRequest
	GetMaxIdleTime() *string
	SetMixMode(v string) *StartLiveMPUTaskShrinkRequest
	GetMixMode() *string
	SetMultiStreamURLShrink(v string) *StartLiveMPUTaskShrinkRequest
	GetMultiStreamURLShrink() *string
	SetRegion(v string) *StartLiveMPUTaskShrinkRequest
	GetRegion() *string
	SetSeiParamsShrink(v string) *StartLiveMPUTaskShrinkRequest
	GetSeiParamsShrink() *string
	SetSingleSubParamsShrink(v string) *StartLiveMPUTaskShrinkRequest
	GetSingleSubParamsShrink() *string
	SetStreamURL(v string) *StartLiveMPUTaskShrinkRequest
	GetStreamURL() *string
	SetTaskId(v string) *StartLiveMPUTaskShrinkRequest
	GetTaskId() *string
	SetTranscodeParamsShrink(v string) *StartLiveMPUTaskShrinkRequest
	GetTranscodeParamsShrink() *string
}

type StartLiveMPUTaskShrinkRequest struct {
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
	// The timeout period of an idle connection. Unit: seconds. Valid values: [10,86400].
	//
	// >  If the task is idle for a period of time longer than the duration specified by the MaxIdleTime parameter, the task is automatically stopped. If the parameter is not specified, the task is stopped after the channel is closed.
	//
	// example:
	//
	// 10
	MaxIdleTime *string `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// The stream mixing mode. Valid values:
	//
	// 	- **0**: the single-stream relay mode. In this mode, the service only relays the original single stream, but does not transcode mixed streams. You do not need to set parameters for mixed-stream transcoding.
	//
	// 	- **1*	- (default): the mixed-stream relay mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MixMode *string `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	// The multiple ingest URLs to relay. This parameter allows you to specify multiple ingest URLs.
	//
	// >  The StreamURL and MultiStreamURL parameters are mutually exclusive. You must specify one of the two parameters.
	MultiStreamURLShrink *string `json:"MultiStreamURL,omitempty" xml:"MultiStreamURL,omitempty"`
	// The region in which the streams are mixed. Valid values:
	//
	// 	- **CN-Shanghai**
	//
	// 	- **AP-Singapore*	- (default)
	//
	// 	- **EMAA-Saudi**
	//
	// example:
	//
	// CN-Shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The supplemental enhancement information (SEI) parameters.
	SeiParamsShrink *string `json:"SeiParams,omitempty" xml:"SeiParams,omitempty"`
	// The single-stream relay parameters. These parameters are required if you set MixMode to 0. Leave these parameters empty in the mixed-stream relay mode.
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
	// The mixed-stream relay parameters. These parameters are required if you set MixMode to 1. Leave these parameters empty if you use the single-stream relay mode.
	TranscodeParamsShrink *string `json:"TranscodeParams,omitempty" xml:"TranscodeParams,omitempty"`
}

func (s StartLiveMPUTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartLiveMPUTaskShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartLiveMPUTaskShrinkRequest) GetMaxIdleTime() *string {
	return s.MaxIdleTime
}

func (s *StartLiveMPUTaskShrinkRequest) GetMixMode() *string {
	return s.MixMode
}

func (s *StartLiveMPUTaskShrinkRequest) GetMultiStreamURLShrink() *string {
	return s.MultiStreamURLShrink
}

func (s *StartLiveMPUTaskShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *StartLiveMPUTaskShrinkRequest) GetSeiParamsShrink() *string {
	return s.SeiParamsShrink
}

func (s *StartLiveMPUTaskShrinkRequest) GetSingleSubParamsShrink() *string {
	return s.SingleSubParamsShrink
}

func (s *StartLiveMPUTaskShrinkRequest) GetStreamURL() *string {
	return s.StreamURL
}

func (s *StartLiveMPUTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartLiveMPUTaskShrinkRequest) GetTranscodeParamsShrink() *string {
	return s.TranscodeParamsShrink
}

func (s *StartLiveMPUTaskShrinkRequest) SetAppId(v string) *StartLiveMPUTaskShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetChannelId(v string) *StartLiveMPUTaskShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetMaxIdleTime(v string) *StartLiveMPUTaskShrinkRequest {
	s.MaxIdleTime = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetMixMode(v string) *StartLiveMPUTaskShrinkRequest {
	s.MixMode = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetMultiStreamURLShrink(v string) *StartLiveMPUTaskShrinkRequest {
	s.MultiStreamURLShrink = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetRegion(v string) *StartLiveMPUTaskShrinkRequest {
	s.Region = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetSeiParamsShrink(v string) *StartLiveMPUTaskShrinkRequest {
	s.SeiParamsShrink = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetSingleSubParamsShrink(v string) *StartLiveMPUTaskShrinkRequest {
	s.SingleSubParamsShrink = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetStreamURL(v string) *StartLiveMPUTaskShrinkRequest {
	s.StreamURL = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetTaskId(v string) *StartLiveMPUTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) SetTranscodeParamsShrink(v string) *StartLiveMPUTaskShrinkRequest {
	s.TranscodeParamsShrink = &v
	return s
}

func (s *StartLiveMPUTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
