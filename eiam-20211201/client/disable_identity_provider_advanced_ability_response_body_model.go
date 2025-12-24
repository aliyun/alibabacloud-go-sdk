// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableIdentityProviderAdvancedAbilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableIdentityProviderAdvancedAbilityResponseBody
	GetRequestId() *string
}

type DisableIdentityProviderAdvancedAbilityResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableIdentityProviderAdvancedAbilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableIdentityProviderAdvancedAbilityResponseBody) GoString() string {
	return s.String()
}

func (s *DisableIdentityProviderAdvancedAbilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableIdentityProviderAdvancedAbilityResponseBody) SetRequestId(v string) *DisableIdentityProviderAdvancedAbilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableIdentityProviderAdvancedAbilityResponseBody) Validate() error {
	return dara.Validate(s)
}
