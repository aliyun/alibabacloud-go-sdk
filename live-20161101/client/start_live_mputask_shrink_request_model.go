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
	// The application ID. Only one ID is supported. It can contain uppercase letters, lowercase letters, digits, underscores (_), and hyphens (-). The maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel ID. Only one ID is supported. It can contain uppercase letters, lowercase letters, digits, underscores (_), and hyphens (-). The maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The idle timeout period. Unit: seconds. The value must be in the range of [10, 86400].
	//
	// > If you set this parameter, the task is automatically stopped when it has been idle for a period longer than MaxIdleTime. If you do not set this parameter, the task is stopped immediately after the channel is closed.
	//
	// example:
	//
	// 10
	MaxIdleTime *string `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// The stream mixing mode. Valid values:
	//
	// - **0**: Single-stream ingest. The original single stream is ingested without stream mixing or transcoding. You do not need to configure stream mixing and transcoding parameters.
	//
	// - **1*	- (default): Stream mixing and transcoding.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MixMode *string `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	// The parameters for ingesting to multiple URLs. You can specify multiple live ingest URLs.
	//
	// > When you set the ingest URL for a task, you must configure either the StreamURL parameter or the MultiStreamURL parameter, but not both.
	MultiStreamURLShrink *string `json:"MultiStreamURL,omitempty" xml:"MultiStreamURL,omitempty"`
	// The region where the stream mixing service is located. Valid values:
	//
	// - **CN-Shanghai<props="china">(default)**: Shanghai.
	//
	// - **AP-Singapore<props="intl">(default)**: Singapore.
	//
	// - **EMAA-Saudi**: Saudi Arabia.
	//
	// example:
	//
	// CN-Shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The SEI configuration parameters.
	SeiParamsShrink *string `json:"SeiParams,omitempty" xml:"SeiParams,omitempty"`
	// The parameters for single-stream ingest. This parameter is required when MixMode is set to 0. Do not set this parameter for stream mixing and transcoding.
	SingleSubParamsShrink *string `json:"SingleSubParams,omitempty" xml:"SingleSubParams,omitempty"`
	// The live ingest URL. Only the RTMP protocol is supported. Only one URL is supported. The maximum length is 2048 characters. For information about how to generate the URL, see [Ingest URLs and playback URLs](https://help.aliyun.com/document_detail/199339.html).
	//
	// > - For domain names with hotlink protection enabled, the ingest URL must include an access token.
	//
	// - Do not use the same StreamURL in different tasks at the same time.
	//
	// - Do not use the same StreamURL within 10 seconds after a task stops.
	//
	// example:
	//
	// rtmp://example.com/live/stream
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// The task ID. Only one ID is supported. It can contain uppercase letters, lowercase letters, digits, underscores (_), and hyphens (-). The maximum length is 55 characters. This ID is the unique identifier for the bypass ingest task.
	//
	// If a task with the same ID still exists and has not been cleared when you start a new task, \\`InvalidParam\\` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The parameters for stream mixing and transcoding. This parameter is required when MixMode is set to 1. Do not set this parameter for single-stream ingest.
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
