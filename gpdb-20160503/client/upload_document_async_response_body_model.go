// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocumentAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UploadDocumentAsyncResponseBody
	GetJobId() *string
	SetMessage(v string) *UploadDocumentAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadDocumentAsyncResponseBody
	GetRequestId() *string
	SetStatus(v string) *UploadDocumentAsyncResponseBody
	GetStatus() *string
}

type UploadDocumentAsyncResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API execution status, with the following values:
	//
	// - **success**: Execution succeeded.
	//
	// - **fail**: Execution failed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UploadDocumentAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDocumentAsyncResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UploadDocumentAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadDocumentAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDocumentAsyncResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UploadDocumentAsyncResponseBody) SetJobId(v string) *UploadDocumentAsyncResponseBody {
	s.JobId = &v
	return s
}

func (s *UploadDocumentAsyncResponseBody) SetMessage(v string) *UploadDocumentAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDocumentAsyncResponseBody) SetRequestId(v string) *UploadDocumentAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDocumentAsyncResponseBody) SetStatus(v string) *UploadDocumentAsyncResponseBody {
	s.Status = &v
	return s
}

func (s *UploadDocumentAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
