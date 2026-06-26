// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIdentityProviderResponseBody
	GetRequestId() *string
}

type UpdateIdentityProviderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIdentityProviderResponseBody) SetRequestId(v string) *UpdateIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
