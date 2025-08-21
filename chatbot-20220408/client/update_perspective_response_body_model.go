// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePerspectiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePerspectiveResponseBody
	GetRequestId() *string
}

type UpdatePerspectiveResponseBody struct {
	// example:
	//
	// FC384CE1-8D42-1900-84E1-F33F990F2B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePerspectiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePerspectiveResponseBody) SetRequestId(v string) *UpdatePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePerspectiveResponseBody) Validate() error {
	return dara.Validate(s)
}
