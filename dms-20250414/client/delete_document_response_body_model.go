// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteDocumentResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteDocumentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDocumentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDocumentResponseBody
	GetSuccess() *bool
}

type DeleteDocumentResponseBody struct {
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

func (s DeleteDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDocumentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDocumentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDocumentResponseBody) SetData(v bool) *DeleteDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetErrorCode(v string) *DeleteDocumentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetErrorMessage(v string) *DeleteDocumentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetRequestId(v string) *DeleteDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetSuccess(v bool) *DeleteDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
