// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteConnectionResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConnectionResponseBody
	GetRequestId() *string
}

type DeleteConnectionResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request is successful, success is returned. If the request failed, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8EF25E37-1750-5D7A-BA56-F8AE081A69C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConnectionResponseBody) SetCode(v string) *DeleteConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConnectionResponseBody) SetMessage(v string) *DeleteConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConnectionResponseBody) SetRequestId(v string) *DeleteConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
