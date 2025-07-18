// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUploadDocumentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CancelUploadDocumentJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelUploadDocumentJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *CancelUploadDocumentJobResponseBody
	GetStatus() *string
}

type CancelUploadDocumentJobResponseBody struct {
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
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CancelUploadDocumentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelUploadDocumentJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUploadDocumentJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelUploadDocumentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelUploadDocumentJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CancelUploadDocumentJobResponseBody) SetMessage(v string) *CancelUploadDocumentJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelUploadDocumentJobResponseBody) SetRequestId(v string) *CancelUploadDocumentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelUploadDocumentJobResponseBody) SetStatus(v string) *CancelUploadDocumentJobResponseBody {
	s.Status = &v
	return s
}

func (s *CancelUploadDocumentJobResponseBody) Validate() error {
	return dara.Validate(s)
}
