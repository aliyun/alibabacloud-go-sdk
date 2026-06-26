// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoginPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLoginPreferenceResponseBody
	GetRequestId() *string
}

type UpdateLoginPreferenceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoginPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoginPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLoginPreferenceResponseBody) SetRequestId(v string) *UpdateLoginPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoginPreferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
