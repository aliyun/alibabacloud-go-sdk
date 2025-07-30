// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationApiInvokeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableApplicationApiInvokeResponseBody
	GetRequestId() *string
}

type DisableApplicationApiInvokeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationApiInvokeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationApiInvokeResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationApiInvokeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationApiInvokeResponseBody) SetRequestId(v string) *DisableApplicationApiInvokeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationApiInvokeResponseBody) Validate() error {
	return dara.Validate(s)
}
