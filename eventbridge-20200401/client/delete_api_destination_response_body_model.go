// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteApiDestinationResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteApiDestinationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApiDestinationResponseBody
	GetRequestId() *string
}

type DeleteApiDestinationResponseBody struct {
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
	// 382E6272-8E9C-5681-AC96-A8AF0BFAC1A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiDestinationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteApiDestinationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApiDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiDestinationResponseBody) SetCode(v string) *DeleteApiDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApiDestinationResponseBody) SetMessage(v string) *DeleteApiDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApiDestinationResponseBody) SetRequestId(v string) *DeleteApiDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiDestinationResponseBody) Validate() error {
	return dara.Validate(s)
}
