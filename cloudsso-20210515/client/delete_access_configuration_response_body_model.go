// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccessConfigurationResponseBody
	GetRequestId() *string
}

type DeleteAccessConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B13E4EE-3853-5852-9165-597C32AD8FB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessConfigurationResponseBody) SetRequestId(v string) *DeleteAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
