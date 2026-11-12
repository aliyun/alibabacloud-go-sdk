// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *CreateCommandResponseBody
	GetCommandId() *string
	SetRequestId(v string) *CreateCommandResponseBody
	GetRequestId() *string
}

type CreateCommandResponseBody struct {
	// The command ID.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommandResponseBody) GetCommandId() *string {
	return s.CommandId
}

func (s *CreateCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCommandResponseBody) SetCommandId(v string) *CreateCommandResponseBody {
	s.CommandId = &v
	return s
}

func (s *CreateCommandResponseBody) SetRequestId(v string) *CreateCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
