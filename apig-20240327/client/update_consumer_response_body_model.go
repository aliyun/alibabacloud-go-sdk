// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConsumerResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateConsumerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConsumerResponseBody
	GetRequestId() *string
}

type UpdateConsumerResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B917B12C-030A-597A-AF2B-6C4353FC9F10
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConsumerResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConsumerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConsumerResponseBody) SetCode(v string) *UpdateConsumerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConsumerResponseBody) SetMessage(v string) *UpdateConsumerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConsumerResponseBody) SetRequestId(v string) *UpdateConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConsumerResponseBody) Validate() error {
	return dara.Validate(s)
}
