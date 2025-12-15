// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoSnapshotTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetVideoSnapshotTaskStatusResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetVideoSnapshotTaskStatusResponseBody
	GetRequestId() *string
	SetResult(v *GetVideoSnapshotTaskStatusResponseBodyResult) *GetVideoSnapshotTaskStatusResponseBody
	GetResult() *GetVideoSnapshotTaskStatusResponseBodyResult
	SetUsage(v *GetVideoSnapshotTaskStatusResponseBodyUsage) *GetVideoSnapshotTaskStatusResponseBody
	GetUsage() *GetVideoSnapshotTaskStatusResponseBodyUsage
}

type GetVideoSnapshotTaskStatusResponseBody struct {
	Latency   *int32                                        `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                       `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetVideoSnapshotTaskStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetVideoSnapshotTaskStatusResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoSnapshotTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSnapshotTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoSnapshotTaskStatusResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetVideoSnapshotTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoSnapshotTaskStatusResponseBody) GetResult() *GetVideoSnapshotTaskStatusResponseBodyResult {
	return s.Result
}

func (s *GetVideoSnapshotTaskStatusResponseBody) GetUsage() *GetVideoSnapshotTaskStatusResponseBodyUsage {
	return s.Usage
}

func (s *GetVideoSnapshotTaskStatusResponseBody) SetLatency(v int32) *GetVideoSnapshotTaskStatusResponseBody {
	s.Latency = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBody) SetRequestId(v string) *GetVideoSnapshotTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBody) SetResult(v *GetVideoSnapshotTaskStatusResponseBodyResult) *GetVideoSnapshotTaskStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBody) SetUsage(v *GetVideoSnapshotTaskStatusResponseBodyUsage) *GetVideoSnapshotTaskStatusResponseBody {
	s.Usage = v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBody) Validate() error {
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

type GetVideoSnapshotTaskStatusResponseBodyResult struct {
	Data   []*GetVideoSnapshotTaskStatusResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Error  *string                                             `json:"error,omitempty" xml:"error,omitempty"`
	Status *string                                             `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string                                             `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetVideoSnapshotTaskStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSnapshotTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResult) GetData() []*GetVideoSnapshotTaskStatusResponseBodyResultData {
	return s.Data
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResult) GetError() *string {
	return s.Error
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResult) SetData(v []*GetVideoSnapshotTaskStatusResponseBodyResultData) *GetVideoSnapshotTaskStatusResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResult) SetError(v string) *GetVideoSnapshotTaskStatusResponseBodyResult {
	s.Error = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResult) SetStatus(v string) *GetVideoSnapshotTaskStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResult) SetTaskId(v string) *GetVideoSnapshotTaskStatusResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResult) Validate() error {
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

type GetVideoSnapshotTaskStatusResponseBodyResultData struct {
	FrameIndex *int64   `json:"frame_index,omitempty" xml:"frame_index,omitempty"`
	FrameTime  *float32 `json:"frame_time,omitempty" xml:"frame_time,omitempty"`
	Path       *string  `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetVideoSnapshotTaskStatusResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSnapshotTaskStatusResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResultData) GetFrameIndex() *int64 {
	return s.FrameIndex
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResultData) GetFrameTime() *float32 {
	return s.FrameTime
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResultData) GetPath() *string {
	return s.Path
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResultData) SetFrameIndex(v int64) *GetVideoSnapshotTaskStatusResponseBodyResultData {
	s.FrameIndex = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResultData) SetFrameTime(v float32) *GetVideoSnapshotTaskStatusResponseBodyResultData {
	s.FrameTime = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResultData) SetPath(v string) *GetVideoSnapshotTaskStatusResponseBodyResultData {
	s.Path = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}

type GetVideoSnapshotTaskStatusResponseBodyUsage struct {
	ImageCount *int64 `json:"image_count,omitempty" xml:"image_count,omitempty"`
}

func (s GetVideoSnapshotTaskStatusResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSnapshotTaskStatusResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetVideoSnapshotTaskStatusResponseBodyUsage) GetImageCount() *int64 {
	return s.ImageCount
}

func (s *GetVideoSnapshotTaskStatusResponseBodyUsage) SetImageCount(v int64) *GetVideoSnapshotTaskStatusResponseBodyUsage {
	s.ImageCount = &v
	return s
}

func (s *GetVideoSnapshotTaskStatusResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
