// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateCollectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCollectionResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateCollectionResponseBody
	GetStatus() *string
}

type CreateCollectionResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// create successfully
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

func (s CreateCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCollectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCollectionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateCollectionResponseBody) SetMessage(v string) *CreateCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCollectionResponseBody) SetRequestId(v string) *CreateCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCollectionResponseBody) SetStatus(v string) *CreateCollectionResponseBody {
	s.Status = &v
	return s
}

func (s *CreateCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
