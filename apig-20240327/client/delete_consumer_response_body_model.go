// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteConsumerResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteConsumerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConsumerResponseBody
	GetRequestId() *string
}

type DeleteConsumerResponseBody struct {
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
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40C8A4FF-7AF1-5B52-A662-02EEE24C6908
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteConsumerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerResponseBody) SetCode(v string) *DeleteConsumerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConsumerResponseBody) SetMessage(v string) *DeleteConsumerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConsumerResponseBody) SetRequestId(v string) *DeleteConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerResponseBody) Validate() error {
	return dara.Validate(s)
}
