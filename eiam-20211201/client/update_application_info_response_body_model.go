// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationInfoResponseBody
	GetRequestId() *string
}

type UpdateApplicationInfoResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationInfoResponseBody) SetRequestId(v string) *UpdateApplicationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
