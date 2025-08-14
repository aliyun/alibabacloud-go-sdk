// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentExtractResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentExtractResultResponseBody
	GetCode() *string
	SetCompleted(v bool) *GetDocumentExtractResultResponseBody
	GetCompleted() *bool
	SetData(v map[string]interface{}) *GetDocumentExtractResultResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetDocumentExtractResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentExtractResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetDocumentExtractResultResponseBody
	GetStatus() *string
}

type GetDocumentExtractResultResponseBody struct {
	// example:
	//
	// noPermission
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDocumentExtractResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentExtractResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentExtractResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentExtractResultResponseBody) GetCompleted() *bool {
	return s.Completed
}

func (s *GetDocumentExtractResultResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetDocumentExtractResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentExtractResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentExtractResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDocumentExtractResultResponseBody) SetCode(v string) *GetDocumentExtractResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetCompleted(v bool) *GetDocumentExtractResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetData(v map[string]interface{}) *GetDocumentExtractResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetMessage(v string) *GetDocumentExtractResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetRequestId(v string) *GetDocumentExtractResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentExtractResultResponseBody) SetStatus(v string) *GetDocumentExtractResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetDocumentExtractResultResponseBody) Validate() error {
	return dara.Validate(s)
}
