// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAPIKeyCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAPIKeyCredentialProviderResponseBody
	GetRequestId() *string
}

type UpdateAPIKeyCredentialProviderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAPIKeyCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAPIKeyCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAPIKeyCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAPIKeyCredentialProviderResponseBody) SetRequestId(v string) *UpdateAPIKeyCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAPIKeyCredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
