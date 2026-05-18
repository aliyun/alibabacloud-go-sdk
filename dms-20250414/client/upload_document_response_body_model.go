// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OneMetaKnowledgeBaseDocument) *UploadDocumentResponseBody
	GetData() *OneMetaKnowledgeBaseDocument
	SetErrorCode(v string) *UploadDocumentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UploadDocumentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UploadDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadDocumentResponseBody
	GetSuccess() *bool
}

type UploadDocumentResponseBody struct {
	Data *OneMetaKnowledgeBaseDocument `json:"Data,omitempty" xml:"Data,omitempty"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDocumentResponseBody) GetData() *OneMetaKnowledgeBaseDocument {
	return s.Data
}

func (s *UploadDocumentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UploadDocumentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UploadDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadDocumentResponseBody) SetData(v *OneMetaKnowledgeBaseDocument) *UploadDocumentResponseBody {
	s.Data = v
	return s
}

func (s *UploadDocumentResponseBody) SetErrorCode(v string) *UploadDocumentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UploadDocumentResponseBody) SetErrorMessage(v string) *UploadDocumentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UploadDocumentResponseBody) SetRequestId(v string) *UploadDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDocumentResponseBody) SetSuccess(v bool) *UploadDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
