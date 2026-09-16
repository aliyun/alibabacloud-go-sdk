// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLumaDocumentResponseBody
	GetCode() *string
	SetData(v *KnowledgeBaseDocument) *GetLumaDocumentResponseBody
	GetData() *KnowledgeBaseDocument
	SetMessage(v string) *GetLumaDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLumaDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLumaDocumentResponseBody
	GetSuccess() *bool
}

type GetLumaDocumentResponseBody struct {
	// The response code returned by the operation. A value of Success indicates that the call succeeds. Otherwise, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The document details, including the processing status and metadata.
	Data *KnowledgeBaseDocument `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, used for troubleshooting and ticket submission.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call succeeds.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLumaDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLumaDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetLumaDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLumaDocumentResponseBody) GetData() *KnowledgeBaseDocument {
	return s.Data
}

func (s *GetLumaDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLumaDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLumaDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLumaDocumentResponseBody) SetCode(v string) *GetLumaDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *GetLumaDocumentResponseBody) SetData(v *KnowledgeBaseDocument) *GetLumaDocumentResponseBody {
	s.Data = v
	return s
}

func (s *GetLumaDocumentResponseBody) SetMessage(v string) *GetLumaDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *GetLumaDocumentResponseBody) SetRequestId(v string) *GetLumaDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLumaDocumentResponseBody) SetSuccess(v bool) *GetLumaDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *GetLumaDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
