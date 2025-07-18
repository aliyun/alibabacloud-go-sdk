// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteDocumentCollectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDocumentCollectionResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteDocumentCollectionResponseBody
	GetStatus() *string
}

type DeleteDocumentCollectionResponseBody struct {
	// Return message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status, with the following possible values:
	//
	// - **success**: Success.
	//
	// - **fail**: Failure.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteDocumentCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentCollectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDocumentCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentCollectionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteDocumentCollectionResponseBody) SetMessage(v string) *DeleteDocumentCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDocumentCollectionResponseBody) SetRequestId(v string) *DeleteDocumentCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentCollectionResponseBody) SetStatus(v string) *DeleteDocumentCollectionResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteDocumentCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
