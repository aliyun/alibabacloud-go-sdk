// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceWebTerminalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckInstanceWebTerminalResponseBody
	GetRequestId() *string
}

type CheckInstanceWebTerminalResponseBody struct {
	// example:
	//
	// F2D0392B-D749-5C48-A98A-3FAE5C9444A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckInstanceWebTerminalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceWebTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceWebTerminalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInstanceWebTerminalResponseBody) SetRequestId(v string) *CheckInstanceWebTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstanceWebTerminalResponseBody) Validate() error {
	return dara.Validate(s)
}
