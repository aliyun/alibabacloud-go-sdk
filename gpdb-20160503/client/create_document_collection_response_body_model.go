// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocumentCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateDocumentCollectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDocumentCollectionResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateDocumentCollectionResponseBody
	GetStatus() *string
}

type CreateDocumentCollectionResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// - **success**
	//
	// - **fail**
	//
	// example:
	//
	// successs
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateDocumentCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocumentCollectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDocumentCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDocumentCollectionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateDocumentCollectionResponseBody) SetMessage(v string) *CreateDocumentCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDocumentCollectionResponseBody) SetRequestId(v string) *CreateDocumentCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocumentCollectionResponseBody) SetStatus(v string) *CreateDocumentCollectionResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDocumentCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
