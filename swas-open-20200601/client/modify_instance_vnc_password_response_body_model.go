// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVncPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceVncPasswordResponseBody
	GetRequestId() *string
}

type ModifyInstanceVncPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceVncPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVncPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVncPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceVncPasswordResponseBody) SetRequestId(v string) *ModifyInstanceVncPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceVncPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
