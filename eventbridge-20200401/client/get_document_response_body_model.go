// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentResponseBody
	GetCode() *string
	SetData(v *KnowledgeBaseDocument) *GetDocumentResponseBody
	GetData() *KnowledgeBaseDocument
	SetMessage(v string) *GetDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocumentResponseBody
	GetSuccess() *bool
}

type GetDocumentResponseBody struct {
	// The response code. A value of Success indicates a successful call. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The document details, including the processing status, number of chunks, and metadata.
	Data *KnowledgeBaseDocument `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message. A value of Operation success is returned if the call succeeds. A specific error description is returned if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentResponseBody) GetData() *KnowledgeBaseDocument {
	return s.Data
}

func (s *GetDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocumentResponseBody) SetCode(v string) *GetDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentResponseBody) SetData(v *KnowledgeBaseDocument) *GetDocumentResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentResponseBody) SetMessage(v string) *GetDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentResponseBody) SetRequestId(v string) *GetDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentResponseBody) SetSuccess(v bool) *GetDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
