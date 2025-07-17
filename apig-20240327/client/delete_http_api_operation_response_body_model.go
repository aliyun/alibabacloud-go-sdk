// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpApiOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHttpApiOperationResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteHttpApiOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHttpApiOperationResponseBody
	GetRequestId() *string
}

type DeleteHttpApiOperationResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message,
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteHttpApiOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpApiOperationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHttpApiOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHttpApiOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpApiOperationResponseBody) SetCode(v string) *DeleteHttpApiOperationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHttpApiOperationResponseBody) SetMessage(v string) *DeleteHttpApiOperationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHttpApiOperationResponseBody) SetRequestId(v string) *DeleteHttpApiOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpApiOperationResponseBody) Validate() error {
	return dara.Validate(s)
}
