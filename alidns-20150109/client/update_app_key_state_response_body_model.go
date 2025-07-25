// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppKeyStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAppKeyStateResponseBody
	GetRequestId() *string
}

type UpdateAppKeyStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppKeyStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppKeyStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppKeyStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppKeyStateResponseBody) SetRequestId(v string) *UpdateAppKeyStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppKeyStateResponseBody) Validate() error {
	return dara.Validate(s)
}
