// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnsSaleControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEnsSaleControlResponseBody
	GetRequestId() *string
}

type UpdateEnsSaleControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnsSaleControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnsSaleControlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnsSaleControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnsSaleControlResponseBody) SetRequestId(v string) *UpdateEnsSaleControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnsSaleControlResponseBody) Validate() error {
	return dara.Validate(s)
}
