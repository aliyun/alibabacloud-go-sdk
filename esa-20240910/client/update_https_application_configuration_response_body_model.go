// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpsApplicationConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHttpsApplicationConfigurationResponseBody
	GetRequestId() *string
}

type UpdateHttpsApplicationConfigurationResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 35C66C7B-671H-4297-9187-2C4477247A78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHttpsApplicationConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpsApplicationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpsApplicationConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpsApplicationConfigurationResponseBody) SetRequestId(v string) *UpdateHttpsApplicationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
