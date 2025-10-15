// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableApplicationTokenResponseBody
	GetRequestId() *string
}

type DisableApplicationTokenResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationTokenResponseBody) SetRequestId(v string) *DisableApplicationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
