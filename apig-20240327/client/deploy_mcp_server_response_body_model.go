// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeployMcpServerResponseBody
	GetCode() *string
	SetMessage(v string) *DeployMcpServerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeployMcpServerResponseBody
	GetRequestId() *string
}

type DeployMcpServerResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeployMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeployMcpServerResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeployMcpServerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeployMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployMcpServerResponseBody) SetCode(v string) *DeployMcpServerResponseBody {
	s.Code = &v
	return s
}

func (s *DeployMcpServerResponseBody) SetMessage(v string) *DeployMcpServerResponseBody {
	s.Message = &v
	return s
}

func (s *DeployMcpServerResponseBody) SetRequestId(v string) *DeployMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
