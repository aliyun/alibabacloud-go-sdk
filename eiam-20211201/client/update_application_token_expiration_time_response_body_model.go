// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationTokenExpirationTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationTokenExpirationTimeResponseBody
	GetRequestId() *string
}

type UpdateApplicationTokenExpirationTimeResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationTokenExpirationTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTokenExpirationTimeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTokenExpirationTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationTokenExpirationTimeResponseBody) SetRequestId(v string) *UpdateApplicationTokenExpirationTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationTokenExpirationTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
