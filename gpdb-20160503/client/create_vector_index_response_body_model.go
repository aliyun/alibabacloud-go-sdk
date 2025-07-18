// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVectorIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateVectorIndexResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVectorIndexResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateVectorIndexResponseBody
	GetStatus() *string
}

type CreateVectorIndexResponseBody struct {
	// Detailed information returned by the API.
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
	// API execution status. Value description:
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

func (s CreateVectorIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVectorIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVectorIndexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVectorIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVectorIndexResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateVectorIndexResponseBody) SetMessage(v string) *CreateVectorIndexResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVectorIndexResponseBody) SetRequestId(v string) *CreateVectorIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVectorIndexResponseBody) SetStatus(v string) *CreateVectorIndexResponseBody {
	s.Status = &v
	return s
}

func (s *CreateVectorIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
