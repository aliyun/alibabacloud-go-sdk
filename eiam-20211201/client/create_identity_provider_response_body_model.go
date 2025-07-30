// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *CreateIdentityProviderResponseBody
	GetIdentityProviderId() *string
	SetRequestId(v string) *CreateIdentityProviderResponseBody
	GetRequestId() *string
}

type CreateIdentityProviderResponseBody struct {
	// Identity provider ID.
	//
	// example:
	//
	// idp_mwpcwnhrimlr2horxXXXX
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBody) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *CreateIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdentityProviderResponseBody) SetIdentityProviderId(v string) *CreateIdentityProviderResponseBody {
	s.IdentityProviderId = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetRequestId(v string) *CreateIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
