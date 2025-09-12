// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphRAGJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetGraphRAGJobResponseBodyJob) *GetGraphRAGJobResponseBody
	GetJob() *GetGraphRAGJobResponseBodyJob
	SetMessage(v string) *GetGraphRAGJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGraphRAGJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetGraphRAGJobResponseBody
	GetStatus() *string
	SetUsage(v *GetGraphRAGJobResponseBodyUsage) *GetGraphRAGJobResponseBody
	GetUsage() *GetGraphRAGJobResponseBodyUsage
}

type GetGraphRAGJobResponseBody struct {
	Job *GetGraphRAGJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Usage  *GetGraphRAGJobResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s GetGraphRAGJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGraphRAGJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetGraphRAGJobResponseBody) GetJob() *GetGraphRAGJobResponseBodyJob {
	return s.Job
}

func (s *GetGraphRAGJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGraphRAGJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGraphRAGJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetGraphRAGJobResponseBody) GetUsage() *GetGraphRAGJobResponseBodyUsage {
	return s.Usage
}

func (s *GetGraphRAGJobResponseBody) SetJob(v *GetGraphRAGJobResponseBodyJob) *GetGraphRAGJobResponseBody {
	s.Job = v
	return s
}

func (s *GetGraphRAGJobResponseBody) SetMessage(v string) *GetGraphRAGJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetGraphRAGJobResponseBody) SetRequestId(v string) *GetGraphRAGJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGraphRAGJobResponseBody) SetStatus(v string) *GetGraphRAGJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetGraphRAGJobResponseBody) SetUsage(v *GetGraphRAGJobResponseBodyUsage) *GetGraphRAGJobResponseBody {
	s.Usage = v
	return s
}

func (s *GetGraphRAGJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetGraphRAGJobResponseBodyJob struct {
	// example:
	//
	// false
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 2024-01-08 16:52:04.864664
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// Failed to connect database.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 20
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-01-08 16:53:04.864664
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetGraphRAGJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetGraphRAGJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetGraphRAGJobResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *GetGraphRAGJobResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetGraphRAGJobResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *GetGraphRAGJobResponseBodyJob) GetId() *string {
	return s.Id
}

func (s *GetGraphRAGJobResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *GetGraphRAGJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetGraphRAGJobResponseBodyJob) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetGraphRAGJobResponseBodyJob) SetCompleted(v bool) *GetGraphRAGJobResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyJob) SetCreateTime(v string) *GetGraphRAGJobResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyJob) SetError(v string) *GetGraphRAGJobResponseBodyJob {
	s.Error = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyJob) SetId(v string) *GetGraphRAGJobResponseBodyJob {
	s.Id = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyJob) SetProgress(v int32) *GetGraphRAGJobResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyJob) SetStatus(v string) *GetGraphRAGJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyJob) SetUpdateTime(v string) *GetGraphRAGJobResponseBodyJob {
	s.UpdateTime = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}

type GetGraphRAGJobResponseBodyUsage struct {
	// example:
	//
	// 475
	EmbeddingTokens *int32 `json:"EmbeddingTokens,omitempty" xml:"EmbeddingTokens,omitempty"`
	// example:
	//
	// 600
	LLMInputTokens *int32 `json:"LLMInputTokens,omitempty" xml:"LLMInputTokens,omitempty"`
	// example:
	//
	// 125
	LLMOutputTokens *int32 `json:"LLMOutputTokens,omitempty" xml:"LLMOutputTokens,omitempty"`
}

func (s GetGraphRAGJobResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetGraphRAGJobResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetGraphRAGJobResponseBodyUsage) GetEmbeddingTokens() *int32 {
	return s.EmbeddingTokens
}

func (s *GetGraphRAGJobResponseBodyUsage) GetLLMInputTokens() *int32 {
	return s.LLMInputTokens
}

func (s *GetGraphRAGJobResponseBodyUsage) GetLLMOutputTokens() *int32 {
	return s.LLMOutputTokens
}

func (s *GetGraphRAGJobResponseBodyUsage) SetEmbeddingTokens(v int32) *GetGraphRAGJobResponseBodyUsage {
	s.EmbeddingTokens = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyUsage) SetLLMInputTokens(v int32) *GetGraphRAGJobResponseBodyUsage {
	s.LLMInputTokens = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyUsage) SetLLMOutputTokens(v int32) *GetGraphRAGJobResponseBodyUsage {
	s.LLMOutputTokens = &v
	return s
}

func (s *GetGraphRAGJobResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
