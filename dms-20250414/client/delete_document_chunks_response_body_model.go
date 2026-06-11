// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int32) *DeleteDocumentChunksResponseBody
	GetData() *int32
	SetErrorCode(v string) *DeleteDocumentChunksResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDocumentChunksResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDocumentChunksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDocumentChunksResponseBody
	GetSuccess() *bool
}

type DeleteDocumentChunksResponseBody struct {
	// The number of deleted chunks.
	//
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the call fails.
	//
	// example:
	//
	// KnowledgeBaseNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Resource not found kb-***
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique request ID for the call. If an error occurs, provide this request ID to support.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDocumentChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentChunksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentChunksResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteDocumentChunksResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDocumentChunksResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDocumentChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentChunksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDocumentChunksResponseBody) SetData(v int32) *DeleteDocumentChunksResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDocumentChunksResponseBody) SetErrorCode(v string) *DeleteDocumentChunksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDocumentChunksResponseBody) SetErrorMessage(v string) *DeleteDocumentChunksResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDocumentChunksResponseBody) SetRequestId(v string) *DeleteDocumentChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentChunksResponseBody) SetSuccess(v bool) *DeleteDocumentChunksResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDocumentChunksResponseBody) Validate() error {
	return dara.Validate(s)
}
