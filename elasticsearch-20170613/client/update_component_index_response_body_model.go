// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateComponentIndexResponseBody
	GetRequestId() *string
}

type UpdateComponentIndexResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComponentIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentIndexResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComponentIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComponentIndexResponseBody) SetRequestId(v string) *UpdateComponentIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComponentIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
