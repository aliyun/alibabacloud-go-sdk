// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendConsoleMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendConsoleMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendConsoleMessageResponseBody
	GetSuccess() *bool
}

type SendConsoleMessageResponseBody struct {
	// The ID of the request. You can use the ID to locate and troubleshoot issues.
	//
	// example:
	//
	// 9F1D7E42-06A8-4B50-9A21-6E0C8D3F5B72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is accepted. A value of true indicates success, and a value of false indicates failure. For asynchronous pushes, a successful acceptance does not mean that all targets have received the message.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendConsoleMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendConsoleMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendConsoleMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendConsoleMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendConsoleMessageResponseBody) SetRequestId(v string) *SendConsoleMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendConsoleMessageResponseBody) SetSuccess(v bool) *SendConsoleMessageResponseBody {
	s.Success = &v
	return s
}

func (s *SendConsoleMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
