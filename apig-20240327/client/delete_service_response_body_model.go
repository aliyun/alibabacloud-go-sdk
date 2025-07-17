// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteServiceResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceResponseBody
	GetRequestId() *string
}

type DeleteServiceResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 3C3B9A12-3868-5EB9-8BEA-F99E03DD125C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceResponseBody) SetCode(v string) *DeleteServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceResponseBody) SetMessage(v string) *DeleteServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
