// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordInitializationConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPasswordInitializationConfigurationResponseBody
	GetRequestId() *string
}

type SetPasswordInitializationConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordInitializationConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordInitializationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordInitializationConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPasswordInitializationConfigurationResponseBody) SetRequestId(v string) *SetPasswordInitializationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPasswordInitializationConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
