// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHttpApiResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteHttpApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHttpApiResponseBody
	GetRequestId() *string
}

type DeleteHttpApiResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5B626361-070A-56A7-B127-ADAC8F3655DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteHttpApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHttpApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHttpApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpApiResponseBody) SetCode(v string) *DeleteHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHttpApiResponseBody) SetMessage(v string) *DeleteHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHttpApiResponseBody) SetRequestId(v string) *DeleteHttpApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpApiResponseBody) Validate() error {
	return dara.Validate(s)
}
