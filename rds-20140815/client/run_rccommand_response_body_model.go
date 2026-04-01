// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunRCCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *RunRCCommandResponseBody
	GetCommandId() *string
	SetInvokeId(v string) *RunRCCommandResponseBody
	GetInvokeId() *string
	SetRequestId(v string) *RunRCCommandResponseBody
	GetRequestId() *string
}

type RunRCCommandResponseBody struct {
	// example:
	//
	// c-7d2a745b412b4601b2d47f6a768d****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// example:
	//
	// t-7d2a745b412b4601b2d47f6a768d****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// example:
	//
	// 06B220E2-EAC5-4DBE-A1FC-1B62DB6A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunRCCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunRCCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunRCCommandResponseBody) GetCommandId() *string {
	return s.CommandId
}

func (s *RunRCCommandResponseBody) GetInvokeId() *string {
	return s.InvokeId
}

func (s *RunRCCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunRCCommandResponseBody) SetCommandId(v string) *RunRCCommandResponseBody {
	s.CommandId = &v
	return s
}

func (s *RunRCCommandResponseBody) SetInvokeId(v string) *RunRCCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunRCCommandResponseBody) SetRequestId(v string) *RunRCCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunRCCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
