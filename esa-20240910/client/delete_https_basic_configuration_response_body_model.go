// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpsBasicConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHttpsBasicConfigurationResponseBody
	GetRequestId() *string
}

type DeleteHttpsBasicConfigurationResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHttpsBasicConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpsBasicConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpsBasicConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpsBasicConfigurationResponseBody) SetRequestId(v string) *DeleteHttpsBasicConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpsBasicConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
