// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutCustomEventResponseBody
	GetCode() *string
	SetMessage(v string) *PutCustomEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutCustomEventResponseBody
	GetRequestId() *string
}

type PutCustomEventResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 131DD9C8-9A32-4428-AD2E-4E3013B6E3A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutCustomEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutCustomEventResponseBody) GoString() string {
	return s.String()
}

func (s *PutCustomEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutCustomEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutCustomEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutCustomEventResponseBody) SetCode(v string) *PutCustomEventResponseBody {
	s.Code = &v
	return s
}

func (s *PutCustomEventResponseBody) SetMessage(v string) *PutCustomEventResponseBody {
	s.Message = &v
	return s
}

func (s *PutCustomEventResponseBody) SetRequestId(v string) *PutCustomEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutCustomEventResponseBody) Validate() error {
	return dara.Validate(s)
}
