// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoSummarizationTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetVideoSummarizationTaskStatusResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetVideoSummarizationTaskStatusResponseBody
	GetRequestId() *string
	SetResult(v *GetVideoSummarizationTaskStatusResponseBodyResult) *GetVideoSummarizationTaskStatusResponseBody
	GetResult() *GetVideoSummarizationTaskStatusResponseBodyResult
	SetUsage(v *GetVideoSummarizationTaskStatusResponseBodyUsage) *GetVideoSummarizationTaskStatusResponseBody
	GetUsage() *GetVideoSummarizationTaskStatusResponseBodyUsage
}

type GetVideoSummarizationTaskStatusResponseBody struct {
	Latency   *int32                                             `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                            `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetVideoSummarizationTaskStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetVideoSummarizationTaskStatusResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoSummarizationTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSummarizationTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoSummarizationTaskStatusResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetVideoSummarizationTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoSummarizationTaskStatusResponseBody) GetResult() *GetVideoSummarizationTaskStatusResponseBodyResult {
	return s.Result
}

func (s *GetVideoSummarizationTaskStatusResponseBody) GetUsage() *GetVideoSummarizationTaskStatusResponseBodyUsage {
	return s.Usage
}

func (s *GetVideoSummarizationTaskStatusResponseBody) SetLatency(v int32) *GetVideoSummarizationTaskStatusResponseBody {
	s.Latency = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBody) SetRequestId(v string) *GetVideoSummarizationTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBody) SetResult(v *GetVideoSummarizationTaskStatusResponseBodyResult) *GetVideoSummarizationTaskStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBody) SetUsage(v *GetVideoSummarizationTaskStatusResponseBodyUsage) *GetVideoSummarizationTaskStatusResponseBody {
	s.Usage = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoSummarizationTaskStatusResponseBodyResult struct {
	Data   *GetVideoSummarizationTaskStatusResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Error  *string                                                `json:"error,omitempty" xml:"error,omitempty"`
	Status *string                                                `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string                                                `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetVideoSummarizationTaskStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSummarizationTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResult) GetData() *GetVideoSummarizationTaskStatusResponseBodyResultData {
	return s.Data
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResult) GetError() *string {
	return s.Error
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResult) SetData(v *GetVideoSummarizationTaskStatusResponseBodyResultData) *GetVideoSummarizationTaskStatusResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResult) SetError(v string) *GetVideoSummarizationTaskStatusResponseBodyResult {
	s.Error = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResult) SetStatus(v string) *GetVideoSummarizationTaskStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResult) SetTaskId(v string) *GetVideoSummarizationTaskStatusResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoSummarizationTaskStatusResponseBodyResultData struct {
	Chunks        []*GetVideoSummarizationTaskStatusResponseBodyResultDataChunks      `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
	VideoMetadata *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata `json:"video_metadata,omitempty" xml:"video_metadata,omitempty" type:"Struct"`
}

func (s GetVideoSummarizationTaskStatusResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSummarizationTaskStatusResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultData) GetChunks() []*GetVideoSummarizationTaskStatusResponseBodyResultDataChunks {
	return s.Chunks
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultData) GetVideoMetadata() *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata {
	return s.VideoMetadata
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultData) SetChunks(v []*GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) *GetVideoSummarizationTaskStatusResponseBodyResultData {
	s.Chunks = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultData) SetVideoMetadata(v *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) *GetVideoSummarizationTaskStatusResponseBodyResultData {
	s.VideoMetadata = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultData) Validate() error {
	if s.Chunks != nil {
		for _, item := range s.Chunks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoMetadata != nil {
		if err := s.VideoMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoSummarizationTaskStatusResponseBodyResultDataChunks struct {
	EnhancedTranscript *string                                                              `json:"enhanced_transcript,omitempty" xml:"enhanced_transcript,omitempty"`
	Index              *int32                                                               `json:"index,omitempty" xml:"index,omitempty"`
	Metadata           *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
}

func (s GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) GoString() string {
	return s.String()
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) GetEnhancedTranscript() *string {
	return s.EnhancedTranscript
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) GetIndex() *int32 {
	return s.Index
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) GetMetadata() *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata {
	return s.Metadata
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) SetEnhancedTranscript(v string) *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks {
	s.EnhancedTranscript = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) SetIndex(v int32) *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks {
	s.Index = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) SetMetadata(v *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks {
	s.Metadata = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunks) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata struct {
	Summary *string   `json:"summary,omitempty" xml:"summary,omitempty"`
	Tags    []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Title   *string   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) GoString() string {
	return s.String()
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) GetSummary() *string {
	return s.Summary
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) GetTags() []*string {
	return s.Tags
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) GetTitle() *string {
	return s.Title
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) SetSummary(v string) *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata {
	s.Summary = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) SetTags(v []*string) *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata {
	s.Tags = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) SetTitle(v string) *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata {
	s.Title = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataChunksMetadata) Validate() error {
	return dara.Validate(s)
}

type GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata struct {
	Summary *string   `json:"summary,omitempty" xml:"summary,omitempty"`
	Tags    []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Title   *string   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) GoString() string {
	return s.String()
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) GetSummary() *string {
	return s.Summary
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) GetTags() []*string {
	return s.Tags
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) GetTitle() *string {
	return s.Title
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) SetSummary(v string) *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata {
	s.Summary = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) SetTags(v []*string) *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata {
	s.Tags = v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) SetTitle(v string) *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata {
	s.Title = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyResultDataVideoMetadata) Validate() error {
	return dara.Validate(s)
}

type GetVideoSummarizationTaskStatusResponseBodyUsage struct {
	AudioToken  *int64 `json:"audio_token,omitempty" xml:"audio_token,omitempty"`
	ImageToken  *int64 `json:"image_token,omitempty" xml:"image_token,omitempty"`
	InputToken  *int64 `json:"input_token,omitempty" xml:"input_token,omitempty"`
	OutputToken *int64 `json:"output_token,omitempty" xml:"output_token,omitempty"`
}

func (s GetVideoSummarizationTaskStatusResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSummarizationTaskStatusResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetVideoSummarizationTaskStatusResponseBodyUsage) GetAudioToken() *int64 {
	return s.AudioToken
}

func (s *GetVideoSummarizationTaskStatusResponseBodyUsage) GetImageToken() *int64 {
	return s.ImageToken
}

func (s *GetVideoSummarizationTaskStatusResponseBodyUsage) GetInputToken() *int64 {
	return s.InputToken
}

func (s *GetVideoSummarizationTaskStatusResponseBodyUsage) GetOutputToken() *int64 {
	return s.OutputToken
}

func (s *GetVideoSummarizationTaskStatusResponseBodyUsage) SetAudioToken(v int64) *GetVideoSummarizationTaskStatusResponseBodyUsage {
	s.AudioToken = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyUsage) SetImageToken(v int64) *GetVideoSummarizationTaskStatusResponseBodyUsage {
	s.ImageToken = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyUsage) SetInputToken(v int64) *GetVideoSummarizationTaskStatusResponseBodyUsage {
	s.InputToken = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyUsage) SetOutputToken(v int64) *GetVideoSummarizationTaskStatusResponseBodyUsage {
	s.OutputToken = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
