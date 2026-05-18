// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertDocumentChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpsertDocumentChunksResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpsertDocumentChunksResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpsertDocumentChunksResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpsertDocumentChunksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpsertDocumentChunksResponseBody
	GetSuccess() *bool
}

type UpsertDocumentChunksResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// KnowledgeBaseNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Resource not found kb-***
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpsertDocumentChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertDocumentChunksResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertDocumentChunksResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpsertDocumentChunksResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpsertDocumentChunksResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpsertDocumentChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertDocumentChunksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpsertDocumentChunksResponseBody) SetData(v bool) *UpsertDocumentChunksResponseBody {
	s.Data = &v
	return s
}

func (s *UpsertDocumentChunksResponseBody) SetErrorCode(v string) *UpsertDocumentChunksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpsertDocumentChunksResponseBody) SetErrorMessage(v string) *UpsertDocumentChunksResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpsertDocumentChunksResponseBody) SetRequestId(v string) *UpsertDocumentChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertDocumentChunksResponseBody) SetSuccess(v bool) *UpsertDocumentChunksResponseBody {
	s.Success = &v
	return s
}

func (s *UpsertDocumentChunksResponseBody) Validate() error {
	return dara.Validate(s)
}
