// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePortalResponseBody
	GetCode() *string
	SetMessage(v string) *DeletePortalResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePortalResponseBody
	GetRequestId() *string
}

type DeletePortalResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePortalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePortalResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePortalResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePortalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePortalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePortalResponseBody) SetCode(v string) *DeletePortalResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePortalResponseBody) SetMessage(v string) *DeletePortalResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePortalResponseBody) SetRequestId(v string) *DeletePortalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePortalResponseBody) Validate() error {
	return dara.Validate(s)
}
