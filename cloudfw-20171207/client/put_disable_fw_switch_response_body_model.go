// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDisableFwSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutDisableFwSwitchResponseBody
	GetRequestId() *string
}

type PutDisableFwSwitchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF7334F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDisableFwSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutDisableFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutDisableFwSwitchResponseBody) SetRequestId(v string) *PutDisableFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutDisableFwSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
