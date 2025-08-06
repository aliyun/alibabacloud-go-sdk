// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStreamInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateStreamInfoResponseBody
	GetRequestId() *string
}

type UpdateStreamInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStreamInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStreamInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStreamInfoResponseBody) SetRequestId(v string) *UpdateStreamInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStreamInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
