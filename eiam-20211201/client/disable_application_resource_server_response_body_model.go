// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationResourceServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableApplicationResourceServerResponseBody
	GetRequestId() *string
}

type DisableApplicationResourceServerResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationResourceServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationResourceServerResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationResourceServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationResourceServerResponseBody) SetRequestId(v string) *DisableApplicationResourceServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationResourceServerResponseBody) Validate() error {
	return dara.Validate(s)
}
