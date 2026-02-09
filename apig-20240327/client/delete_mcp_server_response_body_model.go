// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMcpServerResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMcpServerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMcpServerResponseBody
	GetRequestId() *string
}

type DeleteMcpServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status code.
	//
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The MCP server ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpServerResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMcpServerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcpServerResponseBody) SetCode(v string) *DeleteMcpServerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMcpServerResponseBody) SetMessage(v string) *DeleteMcpServerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMcpServerResponseBody) SetRequestId(v string) *DeleteMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
