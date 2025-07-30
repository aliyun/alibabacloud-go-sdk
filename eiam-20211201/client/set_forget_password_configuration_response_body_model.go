// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetForgetPasswordConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetForgetPasswordConfigurationResponseBody
	GetRequestId() *string
}

type SetForgetPasswordConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetForgetPasswordConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetForgetPasswordConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetForgetPasswordConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetForgetPasswordConfigurationResponseBody) SetRequestId(v string) *SetForgetPasswordConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetForgetPasswordConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
