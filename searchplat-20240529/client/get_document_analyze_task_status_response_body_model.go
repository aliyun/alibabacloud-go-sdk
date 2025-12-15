// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentAnalyzeTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetDocumentAnalyzeTaskStatusResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetDocumentAnalyzeTaskStatusResponseBody
	GetRequestId() *string
	SetResult(v *GetDocumentAnalyzeTaskStatusResponseBodyResult) *GetDocumentAnalyzeTaskStatusResponseBody
	GetResult() *GetDocumentAnalyzeTaskStatusResponseBodyResult
	SetUsage(v *GetDocumentAnalyzeTaskStatusResponseBodyUsage) *GetDocumentAnalyzeTaskStatusResponseBody
	GetUsage() *GetDocumentAnalyzeTaskStatusResponseBodyUsage
}

type GetDocumentAnalyzeTaskStatusResponseBody struct {
	Latency   *int32                                          `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetDocumentAnalyzeTaskStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetDocumentAnalyzeTaskStatusResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetDocumentAnalyzeTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) GetResult() *GetDocumentAnalyzeTaskStatusResponseBodyResult {
	return s.Result
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) GetUsage() *GetDocumentAnalyzeTaskStatusResponseBodyUsage {
	return s.Usage
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) SetLatency(v int32) *GetDocumentAnalyzeTaskStatusResponseBody {
	s.Latency = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) SetRequestId(v string) *GetDocumentAnalyzeTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) SetResult(v *GetDocumentAnalyzeTaskStatusResponseBodyResult) *GetDocumentAnalyzeTaskStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) SetUsage(v *GetDocumentAnalyzeTaskStatusResponseBodyUsage) *GetDocumentAnalyzeTaskStatusResponseBody {
	s.Usage = v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) Validate() error {
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

type GetDocumentAnalyzeTaskStatusResponseBodyResult struct {
	Data   *GetDocumentAnalyzeTaskStatusResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Error  *string                                             `json:"error,omitempty" xml:"error,omitempty"`
	Status *string                                             `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string                                             `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) GetData() *GetDocumentAnalyzeTaskStatusResponseBodyResultData {
	return s.Data
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) GetError() *string {
	return s.Error
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) SetData(v *GetDocumentAnalyzeTaskStatusResponseBodyResultData) *GetDocumentAnalyzeTaskStatusResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) SetError(v string) *GetDocumentAnalyzeTaskStatusResponseBodyResult {
	s.Error = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) SetStatus(v string) *GetDocumentAnalyzeTaskStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) SetTaskId(v string) *GetDocumentAnalyzeTaskStatusResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocumentAnalyzeTaskStatusResponseBodyResultData struct {
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	PageNum     *int32  `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) GetContent() *string {
	return s.Content
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) GetContentType() *string {
	return s.ContentType
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) SetContent(v string) *GetDocumentAnalyzeTaskStatusResponseBodyResultData {
	s.Content = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) SetContentType(v string) *GetDocumentAnalyzeTaskStatusResponseBodyResultData {
	s.ContentType = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) SetPageNum(v int32) *GetDocumentAnalyzeTaskStatusResponseBodyResultData {
	s.PageNum = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}

type GetDocumentAnalyzeTaskStatusResponseBodyUsage struct {
	ImageCount         *int64 `json:"image_count,omitempty" xml:"image_count,omitempty"`
	SemanticTokenCount *int64 `json:"semantic_token_count,omitempty" xml:"semantic_token_count,omitempty"`
	TableCount         *int64 `json:"table_count,omitempty" xml:"table_count,omitempty"`
	TokenCount         *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) GetImageCount() *int64 {
	return s.ImageCount
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) GetSemanticTokenCount() *int64 {
	return s.SemanticTokenCount
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) GetTableCount() *int64 {
	return s.TableCount
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) GetTokenCount() *int64 {
	return s.TokenCount
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) SetImageCount(v int64) *GetDocumentAnalyzeTaskStatusResponseBodyUsage {
	s.ImageCount = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) SetSemanticTokenCount(v int64) *GetDocumentAnalyzeTaskStatusResponseBodyUsage {
	s.SemanticTokenCount = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) SetTableCount(v int64) *GetDocumentAnalyzeTaskStatusResponseBodyUsage {
	s.TableCount = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) SetTokenCount(v int64) *GetDocumentAnalyzeTaskStatusResponseBodyUsage {
	s.TokenCount = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
