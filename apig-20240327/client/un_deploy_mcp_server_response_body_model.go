// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnDeployMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnDeployMcpServerResponseBody
	GetCode() *string
	SetMessage(v string) *UnDeployMcpServerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnDeployMcpServerResponseBody
	GetRequestId() *string
}

type UnDeployMcpServerResponseBody struct {
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
	// CE857A85-251D-5018-8103-A38957D71E20
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UnDeployMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnDeployMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *UnDeployMcpServerResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnDeployMcpServerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnDeployMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnDeployMcpServerResponseBody) SetCode(v string) *UnDeployMcpServerResponseBody {
	s.Code = &v
	return s
}

func (s *UnDeployMcpServerResponseBody) SetMessage(v string) *UnDeployMcpServerResponseBody {
	s.Message = &v
	return s
}

func (s *UnDeployMcpServerResponseBody) SetRequestId(v string) *UnDeployMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnDeployMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
