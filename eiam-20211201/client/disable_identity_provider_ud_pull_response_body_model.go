// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableIdentityProviderUdPullResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableIdentityProviderUdPullResponseBody
	GetRequestId() *string
}

type DisableIdentityProviderUdPullResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableIdentityProviderUdPullResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableIdentityProviderUdPullResponseBody) GoString() string {
	return s.String()
}

func (s *DisableIdentityProviderUdPullResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableIdentityProviderUdPullResponseBody) SetRequestId(v string) *DisableIdentityProviderUdPullResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableIdentityProviderUdPullResponseBody) Validate() error {
	return dara.Validate(s)
}
