// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OneMetaKnowledgeBaseDocument) *UpdateDocumentResponseBody
	GetData() *OneMetaKnowledgeBaseDocument
	SetErrorCode(v string) *UpdateDocumentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDocumentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDocumentResponseBody
	GetSuccess() *bool
}

type UpdateDocumentResponseBody struct {
	// The response object.
	Data *OneMetaKnowledgeBaseDocument `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code for a failed request.
	//
	// example:
	//
	// KnowledgeBaseNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message for a failed request.
	//
	// example:
	//
	// Resource not found kb-***
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique ID of the request. Use this ID for troubleshooting.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request succeeded.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocumentResponseBody) GetData() *OneMetaKnowledgeBaseDocument {
	return s.Data
}

func (s *UpdateDocumentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDocumentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDocumentResponseBody) SetData(v *OneMetaKnowledgeBaseDocument) *UpdateDocumentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDocumentResponseBody) SetErrorCode(v string) *UpdateDocumentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetErrorMessage(v string) *UpdateDocumentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetRequestId(v string) *UpdateDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetSuccess(v bool) *UpdateDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
