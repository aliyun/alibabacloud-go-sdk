// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordHistoryConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPasswordHistoryConfigurationResponseBody
	GetRequestId() *string
}

type SetPasswordHistoryConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordHistoryConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordHistoryConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordHistoryConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPasswordHistoryConfigurationResponseBody) SetRequestId(v string) *SetPasswordHistoryConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPasswordHistoryConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
