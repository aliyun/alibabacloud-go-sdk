// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordComplexityConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPasswordComplexityConfigurationResponseBody
	GetRequestId() *string
}

type SetPasswordComplexityConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordComplexityConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordComplexityConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordComplexityConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPasswordComplexityConfigurationResponseBody) SetRequestId(v string) *SetPasswordComplexityConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPasswordComplexityConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
