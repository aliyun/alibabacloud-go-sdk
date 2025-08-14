// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentCompareResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentCompareResultResponseBody
	GetCode() *string
	SetCompleted(v bool) *GetDocumentCompareResultResponseBody
	GetCompleted() *bool
	SetData(v interface{}) *GetDocumentCompareResultResponseBody
	GetData() interface{}
	SetMessage(v string) *GetDocumentCompareResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentCompareResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetDocumentCompareResultResponseBody
	GetStatus() *string
}

type GetDocumentCompareResultResponseBody struct {
	// example:
	//
	// noPermission
	Code      *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool       `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s GetDocumentCompareResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentCompareResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentCompareResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentCompareResultResponseBody) GetCompleted() *bool {
	return s.Completed
}

func (s *GetDocumentCompareResultResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetDocumentCompareResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentCompareResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentCompareResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDocumentCompareResultResponseBody) SetCode(v string) *GetDocumentCompareResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetCompleted(v bool) *GetDocumentCompareResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetData(v interface{}) *GetDocumentCompareResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetMessage(v string) *GetDocumentCompareResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetRequestId(v string) *GetDocumentCompareResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentCompareResultResponseBody) SetStatus(v string) *GetDocumentCompareResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetDocumentCompareResultResponseBody) Validate() error {
	return dara.Validate(s)
}
