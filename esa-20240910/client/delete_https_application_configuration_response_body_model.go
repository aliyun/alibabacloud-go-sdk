// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpsApplicationConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHttpsApplicationConfigurationResponseBody
	GetRequestId() *string
}

type DeleteHttpsApplicationConfigurationResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 35C66C7B-671H-4297-9187-2C4477247A78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHttpsApplicationConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpsApplicationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpsApplicationConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpsApplicationConfigurationResponseBody) SetRequestId(v string) *DeleteHttpsApplicationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpsApplicationConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
