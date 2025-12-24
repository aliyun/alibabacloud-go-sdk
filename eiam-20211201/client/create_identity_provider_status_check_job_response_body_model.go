// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderStatusCheckJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderStatusCheckJobId(v string) *CreateIdentityProviderStatusCheckJobResponseBody
	GetIdentityProviderStatusCheckJobId() *string
	SetRequestId(v string) *CreateIdentityProviderStatusCheckJobResponseBody
	GetRequestId() *string
}

type CreateIdentityProviderStatusCheckJobResponseBody struct {
	// example:
	//
	// async_000xxxx
	IdentityProviderStatusCheckJobId *string `json:"IdentityProviderStatusCheckJobId,omitempty" xml:"IdentityProviderStatusCheckJobId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIdentityProviderStatusCheckJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderStatusCheckJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderStatusCheckJobResponseBody) GetIdentityProviderStatusCheckJobId() *string {
	return s.IdentityProviderStatusCheckJobId
}

func (s *CreateIdentityProviderStatusCheckJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdentityProviderStatusCheckJobResponseBody) SetIdentityProviderStatusCheckJobId(v string) *CreateIdentityProviderStatusCheckJobResponseBody {
	s.IdentityProviderStatusCheckJobId = &v
	return s
}

func (s *CreateIdentityProviderStatusCheckJobResponseBody) SetRequestId(v string) *CreateIdentityProviderStatusCheckJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdentityProviderStatusCheckJobResponseBody) Validate() error {
	return dara.Validate(s)
}
