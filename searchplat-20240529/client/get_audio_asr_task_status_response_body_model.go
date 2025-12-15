// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioAsrTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetAudioAsrTaskStatusResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetAudioAsrTaskStatusResponseBody
	GetRequestId() *string
	SetResult(v *GetAudioAsrTaskStatusResponseBodyResult) *GetAudioAsrTaskStatusResponseBody
	GetResult() *GetAudioAsrTaskStatusResponseBodyResult
	SetUsage(v *GetAudioAsrTaskStatusResponseBodyUsage) *GetAudioAsrTaskStatusResponseBody
	GetUsage() *GetAudioAsrTaskStatusResponseBodyUsage
}

type GetAudioAsrTaskStatusResponseBody struct {
	Latency   *int32                                   `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                  `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetAudioAsrTaskStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetAudioAsrTaskStatusResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetAudioAsrTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAudioAsrTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAudioAsrTaskStatusResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetAudioAsrTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAudioAsrTaskStatusResponseBody) GetResult() *GetAudioAsrTaskStatusResponseBodyResult {
	return s.Result
}

func (s *GetAudioAsrTaskStatusResponseBody) GetUsage() *GetAudioAsrTaskStatusResponseBodyUsage {
	return s.Usage
}

func (s *GetAudioAsrTaskStatusResponseBody) SetLatency(v int32) *GetAudioAsrTaskStatusResponseBody {
	s.Latency = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBody) SetRequestId(v string) *GetAudioAsrTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBody) SetResult(v *GetAudioAsrTaskStatusResponseBodyResult) *GetAudioAsrTaskStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBody) SetUsage(v *GetAudioAsrTaskStatusResponseBodyUsage) *GetAudioAsrTaskStatusResponseBody {
	s.Usage = v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBody) Validate() error {
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

type GetAudioAsrTaskStatusResponseBodyResult struct {
	Data   []*GetAudioAsrTaskStatusResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Error  *string                                        `json:"error,omitempty" xml:"error,omitempty"`
	Status *string                                        `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string                                        `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetAudioAsrTaskStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAudioAsrTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAudioAsrTaskStatusResponseBodyResult) GetData() []*GetAudioAsrTaskStatusResponseBodyResultData {
	return s.Data
}

func (s *GetAudioAsrTaskStatusResponseBodyResult) GetError() *string {
	return s.Error
}

func (s *GetAudioAsrTaskStatusResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetAudioAsrTaskStatusResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAudioAsrTaskStatusResponseBodyResult) SetData(v []*GetAudioAsrTaskStatusResponseBodyResultData) *GetAudioAsrTaskStatusResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBodyResult) SetError(v string) *GetAudioAsrTaskStatusResponseBodyResult {
	s.Error = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBodyResult) SetStatus(v string) *GetAudioAsrTaskStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBodyResult) SetTaskId(v string) *GetAudioAsrTaskStatusResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBodyResult) Validate() error {
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

type GetAudioAsrTaskStatusResponseBodyResultData struct {
	End   *float32 `json:"end,omitempty" xml:"end,omitempty"`
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
	Text  *string  `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetAudioAsrTaskStatusResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s GetAudioAsrTaskStatusResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetAudioAsrTaskStatusResponseBodyResultData) GetEnd() *float32 {
	return s.End
}

func (s *GetAudioAsrTaskStatusResponseBodyResultData) GetStart() *float32 {
	return s.Start
}

func (s *GetAudioAsrTaskStatusResponseBodyResultData) GetText() *string {
	return s.Text
}

func (s *GetAudioAsrTaskStatusResponseBodyResultData) SetEnd(v float32) *GetAudioAsrTaskStatusResponseBodyResultData {
	s.End = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBodyResultData) SetStart(v float32) *GetAudioAsrTaskStatusResponseBodyResultData {
	s.Start = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBodyResultData) SetText(v string) *GetAudioAsrTaskStatusResponseBodyResultData {
	s.Text = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}

type GetAudioAsrTaskStatusResponseBodyUsage struct {
	Duration *float32 `json:"duration,omitempty" xml:"duration,omitempty"`
}

func (s GetAudioAsrTaskStatusResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetAudioAsrTaskStatusResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetAudioAsrTaskStatusResponseBodyUsage) GetDuration() *float32 {
	return s.Duration
}

func (s *GetAudioAsrTaskStatusResponseBodyUsage) SetDuration(v float32) *GetAudioAsrTaskStatusResponseBodyUsage {
	s.Duration = &v
	return s
}

func (s *GetAudioAsrTaskStatusResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
