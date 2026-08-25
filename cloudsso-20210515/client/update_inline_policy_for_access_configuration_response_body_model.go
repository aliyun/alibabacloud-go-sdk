// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInlinePolicyForAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInlinePolicyForAccessConfigurationResponseBody
	GetRequestId() *string
}

type UpdateInlinePolicyForAccessConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B13E4EE-3853-5852-9165-597C32AD8FB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInlinePolicyForAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInlinePolicyForAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInlinePolicyForAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInlinePolicyForAccessConfigurationResponseBody) SetRequestId(v string) *UpdateInlinePolicyForAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInlinePolicyForAccessConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
