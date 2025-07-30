// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordExpirationConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPasswordExpirationConfigurationResponseBody
	GetRequestId() *string
}

type SetPasswordExpirationConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordExpirationConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordExpirationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordExpirationConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPasswordExpirationConfigurationResponseBody) SetRequestId(v string) *SetPasswordExpirationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPasswordExpirationConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
