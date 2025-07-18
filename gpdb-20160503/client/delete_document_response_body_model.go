// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDocumentResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteDocumentResponseBody
	GetStatus() *string
}

type DeleteDocumentResponseBody struct {
	// Detailed information returned by the API.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Creation status, value description: - **success**: Success - **fail**: Fail
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteDocumentResponseBody) SetMessage(v string) *DeleteDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetRequestId(v string) *DeleteDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetStatus(v string) *DeleteDocumentResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
