// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *RunCommandResponseBody
	GetInvokeId() *string
	SetRequestId(v string) *RunCommandResponseBody
	GetRequestId() *string
}

type RunCommandResponseBody struct {
	// The command execution ID.
	//
	// example:
	//
	// t-hz02p9545t6****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) GetInvokeId() *string {
	return s.InvokeId
}

func (s *RunCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
