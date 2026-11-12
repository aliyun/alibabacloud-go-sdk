// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *InvokeCommandResponseBody
	GetInvokeId() *string
	SetRequestId(v string) *InvokeCommandResponseBody
	GetRequestId() *string
}

type InvokeCommandResponseBody struct {
	// The execution ID of the command.
	//
	// example:
	//
	// t-bj02prjhw1n****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvokeCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeCommandResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeCommandResponseBody) GetInvokeId() *string {
	return s.InvokeId
}

func (s *InvokeCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeCommandResponseBody) SetInvokeId(v string) *InvokeCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *InvokeCommandResponseBody) SetRequestId(v string) *InvokeCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
