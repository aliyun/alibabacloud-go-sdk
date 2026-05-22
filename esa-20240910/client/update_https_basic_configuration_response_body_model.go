// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpsBasicConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHttpsBasicConfigurationResponseBody
	GetRequestId() *string
}

type UpdateHttpsBasicConfigurationResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHttpsBasicConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpsBasicConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpsBasicConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpsBasicConfigurationResponseBody) SetRequestId(v string) *UpdateHttpsBasicConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
