// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadDocumentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChunkResult(v *GetUploadDocumentJobResponseBodyChunkResult) *GetUploadDocumentJobResponseBody
	GetChunkResult() *GetUploadDocumentJobResponseBodyChunkResult
	SetJob(v *GetUploadDocumentJobResponseBodyJob) *GetUploadDocumentJobResponseBody
	GetJob() *GetUploadDocumentJobResponseBodyJob
	SetMessage(v string) *GetUploadDocumentJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUploadDocumentJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetUploadDocumentJobResponseBody
	GetStatus() *string
	SetUsage(v *GetUploadDocumentJobResponseBodyUsage) *GetUploadDocumentJobResponseBody
	GetUsage() *GetUploadDocumentJobResponseBodyUsage
}

type GetUploadDocumentJobResponseBody struct {
	// The chunking result.
	ChunkResult *GetUploadDocumentJobResponseBodyChunkResult `json:"ChunkResult,omitempty" xml:"ChunkResult,omitempty" type:"Struct"`
	// The information about the document upload job.
	Job *GetUploadDocumentJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of tokens that are used for document understanding or embedding.
	Usage *GetUploadDocumentJobResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s GetUploadDocumentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDocumentJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadDocumentJobResponseBody) GetChunkResult() *GetUploadDocumentJobResponseBodyChunkResult {
	return s.ChunkResult
}

func (s *GetUploadDocumentJobResponseBody) GetJob() *GetUploadDocumentJobResponseBodyJob {
	return s.Job
}

func (s *GetUploadDocumentJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUploadDocumentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadDocumentJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetUploadDocumentJobResponseBody) GetUsage() *GetUploadDocumentJobResponseBodyUsage {
	return s.Usage
}

func (s *GetUploadDocumentJobResponseBody) SetChunkResult(v *GetUploadDocumentJobResponseBodyChunkResult) *GetUploadDocumentJobResponseBody {
	s.ChunkResult = v
	return s
}

func (s *GetUploadDocumentJobResponseBody) SetJob(v *GetUploadDocumentJobResponseBodyJob) *GetUploadDocumentJobResponseBody {
	s.Job = v
	return s
}

func (s *GetUploadDocumentJobResponseBody) SetMessage(v string) *GetUploadDocumentJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetUploadDocumentJobResponseBody) SetRequestId(v string) *GetUploadDocumentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadDocumentJobResponseBody) SetStatus(v string) *GetUploadDocumentJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetUploadDocumentJobResponseBody) SetUsage(v *GetUploadDocumentJobResponseBodyUsage) *GetUploadDocumentJobResponseBody {
	s.Usage = v
	return s
}

func (s *GetUploadDocumentJobResponseBody) Validate() error {
	if s.ChunkResult != nil {
		if err := s.ChunkResult.Validate(); err != nil {
			return err
		}
	}
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
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

type GetUploadDocumentJobResponseBodyChunkResult struct {
	// The URL of the file after chunking. The validity period of the URL is 2 hours. The file is in the JSONL format. Each line is in the `{"page_content":"*****", "metadata": {"**":"***","**":"***"}` format.
	//
	// example:
	//
	// http://xxx/test.jsonl
	ChunkFileUrl                *string `json:"ChunkFileUrl,omitempty" xml:"ChunkFileUrl,omitempty"`
	DocumentLoaderResultFileUrl *string `json:"DocumentLoaderResultFileUrl,omitempty" xml:"DocumentLoaderResultFileUrl,omitempty"`
	// The URL of the file that does not contain metadata after chunking. The validity period of the URL is 2 hours. The file is in the TXT format. Each line is a chunk. The file can be easily used for embedding.
	//
	// example:
	//
	// http://xxx/test.txt
	PlainChunkFileUrl *string `json:"PlainChunkFileUrl,omitempty" xml:"PlainChunkFileUrl,omitempty"`
}

func (s GetUploadDocumentJobResponseBodyChunkResult) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDocumentJobResponseBodyChunkResult) GoString() string {
	return s.String()
}

func (s *GetUploadDocumentJobResponseBodyChunkResult) GetChunkFileUrl() *string {
	return s.ChunkFileUrl
}

func (s *GetUploadDocumentJobResponseBodyChunkResult) GetDocumentLoaderResultFileUrl() *string {
	return s.DocumentLoaderResultFileUrl
}

func (s *GetUploadDocumentJobResponseBodyChunkResult) GetPlainChunkFileUrl() *string {
	return s.PlainChunkFileUrl
}

func (s *GetUploadDocumentJobResponseBodyChunkResult) SetChunkFileUrl(v string) *GetUploadDocumentJobResponseBodyChunkResult {
	s.ChunkFileUrl = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyChunkResult) SetDocumentLoaderResultFileUrl(v string) *GetUploadDocumentJobResponseBodyChunkResult {
	s.DocumentLoaderResultFileUrl = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyChunkResult) SetPlainChunkFileUrl(v string) *GetUploadDocumentJobResponseBodyChunkResult {
	s.PlainChunkFileUrl = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyChunkResult) Validate() error {
	return dara.Validate(s)
}

type GetUploadDocumentJobResponseBodyJob struct {
	// Indicates whether the operation is complete.
	//
	// example:
	//
	// false
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2024-01-08 16:52:04.864664
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// Failed to connect database.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The error code.
	//
	// example:
	//
	// InternalError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The progress of the document upload job. Unit: %. A value of 100 indicates that the job is complete.
	//
	// example:
	//
	// 20
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the job. Valid values: Valid values:
	//
	// 	- Success
	//
	// 	- Failed (See the Error parameter for failure reasons.)
	//
	// 	- Cancelling
	//
	// 	- Cancelled
	//
	// 	- Start
	//
	// 	- Running
	//
	// 	- Pending
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the job was updated.
	//
	// example:
	//
	// 2024-01-08 16:53:04.864664
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetUploadDocumentJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDocumentJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetUploadDocumentJobResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *GetUploadDocumentJobResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUploadDocumentJobResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *GetUploadDocumentJobResponseBodyJob) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUploadDocumentJobResponseBodyJob) GetId() *string {
	return s.Id
}

func (s *GetUploadDocumentJobResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *GetUploadDocumentJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetUploadDocumentJobResponseBodyJob) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUploadDocumentJobResponseBodyJob) SetCompleted(v bool) *GetUploadDocumentJobResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyJob) SetCreateTime(v string) *GetUploadDocumentJobResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyJob) SetError(v string) *GetUploadDocumentJobResponseBodyJob {
	s.Error = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyJob) SetErrorCode(v string) *GetUploadDocumentJobResponseBodyJob {
	s.ErrorCode = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyJob) SetId(v string) *GetUploadDocumentJobResponseBodyJob {
	s.Id = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyJob) SetProgress(v int32) *GetUploadDocumentJobResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyJob) SetStatus(v string) *GetUploadDocumentJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyJob) SetUpdateTime(v string) *GetUploadDocumentJobResponseBodyJob {
	s.UpdateTime = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}

type GetUploadDocumentJobResponseBodyUsage struct {
	// The count of embedding entries.
	//
	// example:
	//
	// 10
	EmbeddingEntries *int32 `json:"EmbeddingEntries,omitempty" xml:"EmbeddingEntries,omitempty"`
	// The number of tokens that are used for vectorization.
	//
	// >  A token is the minimum unit for splitting text. A token can be a word, phrase, punctuation, or character.
	//
	// example:
	//
	// 475
	EmbeddingTokens *int32 `json:"EmbeddingTokens,omitempty" xml:"EmbeddingTokens,omitempty"`
}

func (s GetUploadDocumentJobResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDocumentJobResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetUploadDocumentJobResponseBodyUsage) GetEmbeddingEntries() *int32 {
	return s.EmbeddingEntries
}

func (s *GetUploadDocumentJobResponseBodyUsage) GetEmbeddingTokens() *int32 {
	return s.EmbeddingTokens
}

func (s *GetUploadDocumentJobResponseBodyUsage) SetEmbeddingEntries(v int32) *GetUploadDocumentJobResponseBodyUsage {
	s.EmbeddingEntries = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyUsage) SetEmbeddingTokens(v int32) *GetUploadDocumentJobResponseBodyUsage {
	s.EmbeddingTokens = &v
	return s
}

func (s *GetUploadDocumentJobResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
