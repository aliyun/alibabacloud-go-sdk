// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableApplicationResponseBody
	GetRequestId() *string
}

type DisableApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationResponseBody) SetRequestId(v string) *DisableApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
