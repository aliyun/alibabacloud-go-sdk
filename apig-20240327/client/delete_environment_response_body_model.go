// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEnvironmentResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteEnvironmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEnvironmentResponseBody
	GetRequestId() *string
}

type DeleteEnvironmentResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the request chain.
	//
	// example:
	//
	// C61E30D3-579A-5B43-994E-31E02EDC9129
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEnvironmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnvironmentResponseBody) SetCode(v string) *DeleteEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetMessage(v string) *DeleteEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetRequestId(v string) *DeleteEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) Validate() error {
	return dara.Validate(s)
}
