// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserMFAAuthenticationSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserMFAAuthenticationSettingsResponseBody
	GetRequestId() *string
}

type UpdateUserMFAAuthenticationSettingsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5E6C6049-E9B0-5F6F-A104-6150E3B1F4D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserMFAAuthenticationSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserMFAAuthenticationSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserMFAAuthenticationSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserMFAAuthenticationSettingsResponseBody) SetRequestId(v string) *UpdateUserMFAAuthenticationSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserMFAAuthenticationSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
