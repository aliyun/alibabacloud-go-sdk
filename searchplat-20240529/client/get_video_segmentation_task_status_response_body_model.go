// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoSegmentationTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetVideoSegmentationTaskStatusResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetVideoSegmentationTaskStatusResponseBody
	GetRequestId() *string
	SetResult(v *GetVideoSegmentationTaskStatusResponseBodyResult) *GetVideoSegmentationTaskStatusResponseBody
	GetResult() *GetVideoSegmentationTaskStatusResponseBodyResult
	SetUsage(v *GetVideoSegmentationTaskStatusResponseBodyUsage) *GetVideoSegmentationTaskStatusResponseBody
	GetUsage() *GetVideoSegmentationTaskStatusResponseBodyUsage
}

type GetVideoSegmentationTaskStatusResponseBody struct {
	Latency   *int32                                            `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                           `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetVideoSegmentationTaskStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetVideoSegmentationTaskStatusResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoSegmentationTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSegmentationTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoSegmentationTaskStatusResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetVideoSegmentationTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoSegmentationTaskStatusResponseBody) GetResult() *GetVideoSegmentationTaskStatusResponseBodyResult {
	return s.Result
}

func (s *GetVideoSegmentationTaskStatusResponseBody) GetUsage() *GetVideoSegmentationTaskStatusResponseBodyUsage {
	return s.Usage
}

func (s *GetVideoSegmentationTaskStatusResponseBody) SetLatency(v int32) *GetVideoSegmentationTaskStatusResponseBody {
	s.Latency = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBody) SetRequestId(v string) *GetVideoSegmentationTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBody) SetResult(v *GetVideoSegmentationTaskStatusResponseBodyResult) *GetVideoSegmentationTaskStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBody) SetUsage(v *GetVideoSegmentationTaskStatusResponseBodyUsage) *GetVideoSegmentationTaskStatusResponseBody {
	s.Usage = v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBody) Validate() error {
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

type GetVideoSegmentationTaskStatusResponseBodyResult struct {
	Data   []*GetVideoSegmentationTaskStatusResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Error  *string                                                 `json:"error,omitempty" xml:"error,omitempty"`
	Status *string                                                 `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string                                                 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetVideoSegmentationTaskStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSegmentationTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResult) GetData() []*GetVideoSegmentationTaskStatusResponseBodyResultData {
	return s.Data
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResult) GetError() *string {
	return s.Error
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResult) SetData(v []*GetVideoSegmentationTaskStatusResponseBodyResultData) *GetVideoSegmentationTaskStatusResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResult) SetError(v string) *GetVideoSegmentationTaskStatusResponseBodyResult {
	s.Error = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResult) SetStatus(v string) *GetVideoSegmentationTaskStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResult) SetTaskId(v string) *GetVideoSegmentationTaskStatusResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoSegmentationTaskStatusResponseBodyResultData struct {
	ChunkIndex *int32                                                           `json:"chunk_index,omitempty" xml:"chunk_index,omitempty"`
	EndTime    *float32                                                         `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Snapshots  []*GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots `json:"snapshots,omitempty" xml:"snapshots,omitempty" type:"Repeated"`
	StartTime  *float32                                                         `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Transcript *string                                                          `json:"transcript,omitempty" xml:"transcript,omitempty"`
}

func (s GetVideoSegmentationTaskStatusResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSegmentationTaskStatusResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) GetChunkIndex() *int32 {
	return s.ChunkIndex
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) GetEndTime() *float32 {
	return s.EndTime
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) GetSnapshots() []*GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots {
	return s.Snapshots
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) GetStartTime() *float32 {
	return s.StartTime
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) GetTranscript() *string {
	return s.Transcript
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) SetChunkIndex(v int32) *GetVideoSegmentationTaskStatusResponseBodyResultData {
	s.ChunkIndex = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) SetEndTime(v float32) *GetVideoSegmentationTaskStatusResponseBodyResultData {
	s.EndTime = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) SetSnapshots(v []*GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) *GetVideoSegmentationTaskStatusResponseBodyResultData {
	s.Snapshots = v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) SetStartTime(v float32) *GetVideoSegmentationTaskStatusResponseBodyResultData {
	s.StartTime = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) SetTranscript(v string) *GetVideoSegmentationTaskStatusResponseBodyResultData {
	s.Transcript = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultData) Validate() error {
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots struct {
	FrameIndex *int32   `json:"frame_index,omitempty" xml:"frame_index,omitempty"`
	FrameTime  *float32 `json:"frame_time,omitempty" xml:"frame_time,omitempty"`
	Path       *string  `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) GoString() string {
	return s.String()
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) GetFrameIndex() *int32 {
	return s.FrameIndex
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) GetFrameTime() *float32 {
	return s.FrameTime
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) GetPath() *string {
	return s.Path
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) SetFrameIndex(v int32) *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots {
	s.FrameIndex = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) SetFrameTime(v float32) *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots {
	s.FrameTime = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) SetPath(v string) *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots {
	s.Path = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyResultDataSnapshots) Validate() error {
	return dara.Validate(s)
}

type GetVideoSegmentationTaskStatusResponseBodyUsage struct {
	AudioToken *int64 `json:"audio_token,omitempty" xml:"audio_token,omitempty"`
	ImageToken *int64 `json:"image_token,omitempty" xml:"image_token,omitempty"`
}

func (s GetVideoSegmentationTaskStatusResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSegmentationTaskStatusResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetVideoSegmentationTaskStatusResponseBodyUsage) GetAudioToken() *int64 {
	return s.AudioToken
}

func (s *GetVideoSegmentationTaskStatusResponseBodyUsage) GetImageToken() *int64 {
	return s.ImageToken
}

func (s *GetVideoSegmentationTaskStatusResponseBodyUsage) SetAudioToken(v int64) *GetVideoSegmentationTaskStatusResponseBodyUsage {
	s.AudioToken = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyUsage) SetImageToken(v int64) *GetVideoSegmentationTaskStatusResponseBodyUsage {
	s.ImageToken = &v
	return s
}

func (s *GetVideoSegmentationTaskStatusResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
