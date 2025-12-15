// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageAnalyzeTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetImageAnalyzeTaskStatusResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetImageAnalyzeTaskStatusResponseBody
	GetRequestId() *string
	SetResult(v *GetImageAnalyzeTaskStatusResponseBodyResult) *GetImageAnalyzeTaskStatusResponseBody
	GetResult() *GetImageAnalyzeTaskStatusResponseBodyResult
	SetUsage(v *GetImageAnalyzeTaskStatusResponseBodyUsage) *GetImageAnalyzeTaskStatusResponseBody
	GetUsage() *GetImageAnalyzeTaskStatusResponseBodyUsage
}

type GetImageAnalyzeTaskStatusResponseBody struct {
	Latency   *int32                                       `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                      `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetImageAnalyzeTaskStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetImageAnalyzeTaskStatusResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetImageAnalyzeTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageAnalyzeTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageAnalyzeTaskStatusResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetImageAnalyzeTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageAnalyzeTaskStatusResponseBody) GetResult() *GetImageAnalyzeTaskStatusResponseBodyResult {
	return s.Result
}

func (s *GetImageAnalyzeTaskStatusResponseBody) GetUsage() *GetImageAnalyzeTaskStatusResponseBodyUsage {
	return s.Usage
}

func (s *GetImageAnalyzeTaskStatusResponseBody) SetLatency(v int32) *GetImageAnalyzeTaskStatusResponseBody {
	s.Latency = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBody) SetRequestId(v string) *GetImageAnalyzeTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBody) SetResult(v *GetImageAnalyzeTaskStatusResponseBodyResult) *GetImageAnalyzeTaskStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBody) SetUsage(v *GetImageAnalyzeTaskStatusResponseBodyUsage) *GetImageAnalyzeTaskStatusResponseBody {
	s.Usage = v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBody) Validate() error {
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

type GetImageAnalyzeTaskStatusResponseBodyResult struct {
	Data   *GetImageAnalyzeTaskStatusResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Error  *string                                          `json:"error,omitempty" xml:"error,omitempty"`
	Status *string                                          `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string                                          `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetImageAnalyzeTaskStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetImageAnalyzeTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResult) GetData() *GetImageAnalyzeTaskStatusResponseBodyResultData {
	return s.Data
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResult) GetError() *string {
	return s.Error
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResult) SetData(v *GetImageAnalyzeTaskStatusResponseBodyResultData) *GetImageAnalyzeTaskStatusResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResult) SetError(v string) *GetImageAnalyzeTaskStatusResponseBodyResult {
	s.Error = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResult) SetStatus(v string) *GetImageAnalyzeTaskStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResult) SetTaskId(v string) *GetImageAnalyzeTaskStatusResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageAnalyzeTaskStatusResponseBodyResultData struct {
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	PageNum     *int32  `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

func (s GetImageAnalyzeTaskStatusResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s GetImageAnalyzeTaskStatusResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResultData) GetContent() *string {
	return s.Content
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResultData) GetContentType() *string {
	return s.ContentType
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResultData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResultData) SetContent(v string) *GetImageAnalyzeTaskStatusResponseBodyResultData {
	s.Content = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResultData) SetContentType(v string) *GetImageAnalyzeTaskStatusResponseBodyResultData {
	s.ContentType = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResultData) SetPageNum(v int32) *GetImageAnalyzeTaskStatusResponseBodyResultData {
	s.PageNum = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}

type GetImageAnalyzeTaskStatusResponseBodyUsage struct {
	PvCount    *int64 `json:"pv_count,omitempty" xml:"pv_count,omitempty"`
	TokenCount *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetImageAnalyzeTaskStatusResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetImageAnalyzeTaskStatusResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetImageAnalyzeTaskStatusResponseBodyUsage) GetPvCount() *int64 {
	return s.PvCount
}

func (s *GetImageAnalyzeTaskStatusResponseBodyUsage) GetTokenCount() *int64 {
	return s.TokenCount
}

func (s *GetImageAnalyzeTaskStatusResponseBodyUsage) SetPvCount(v int64) *GetImageAnalyzeTaskStatusResponseBodyUsage {
	s.PvCount = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBodyUsage) SetTokenCount(v int64) *GetImageAnalyzeTaskStatusResponseBodyUsage {
	s.TokenCount = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
