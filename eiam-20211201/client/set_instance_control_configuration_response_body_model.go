// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceControlConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetInstanceControlConfigurationResponseBody
	GetRequestId() *string
}

type SetInstanceControlConfigurationResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstanceControlConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceControlConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceControlConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetInstanceControlConfigurationResponseBody) SetRequestId(v string) *SetInstanceControlConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetInstanceControlConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
