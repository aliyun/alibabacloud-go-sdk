// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEmbeddingTokens(v string) *UpsertChunksResponseBody
	GetEmbeddingTokens() *string
	SetJobId(v string) *UpsertChunksResponseBody
	GetJobId() *string
	SetMessage(v string) *UpsertChunksResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpsertChunksResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpsertChunksResponseBody
	GetStatus() *string
}

type UpsertChunksResponseBody struct {
	// The number of tokens used for vectorization.
	//
	// > A token is the smallest unit into which input text is divided. A token can be a word, a phrase, a punctuation mark, or a character.
	//
	// example:
	//
	// 100
	EmbeddingTokens *string `json:"EmbeddingTokens,omitempty" xml:"EmbeddingTokens,omitempty"`
	// The job ID. You can use this ID to call the GetGraphRAGJob operation to query the job status later.
	//
	// > This parameter is returned only when knowledge graph construction is enabled for the knowledge base.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution status of the API operation. Valid values:
	//
	// - **success**: The operation is successful.
	//
	// - **fail**: The operation fails.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpsertChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertChunksResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertChunksResponseBody) GetEmbeddingTokens() *string {
	return s.EmbeddingTokens
}

func (s *UpsertChunksResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpsertChunksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpsertChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertChunksResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpsertChunksResponseBody) SetEmbeddingTokens(v string) *UpsertChunksResponseBody {
	s.EmbeddingTokens = &v
	return s
}

func (s *UpsertChunksResponseBody) SetJobId(v string) *UpsertChunksResponseBody {
	s.JobId = &v
	return s
}

func (s *UpsertChunksResponseBody) SetMessage(v string) *UpsertChunksResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertChunksResponseBody) SetRequestId(v string) *UpsertChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertChunksResponseBody) SetStatus(v string) *UpsertChunksResponseBody {
	s.Status = &v
	return s
}

func (s *UpsertChunksResponseBody) Validate() error {
	return dara.Validate(s)
}
