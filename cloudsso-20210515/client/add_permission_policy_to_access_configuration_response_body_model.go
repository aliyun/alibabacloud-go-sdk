// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPermissionPolicyToAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddPermissionPolicyToAccessConfigurationResponseBody
	GetRequestId() *string
}

type AddPermissionPolicyToAccessConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B7C6E839-FB65-59BE-B753-003AA8AF7DF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPermissionPolicyToAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionPolicyToAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *AddPermissionPolicyToAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPermissionPolicyToAccessConfigurationResponseBody) SetRequestId(v string) *AddPermissionPolicyToAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
