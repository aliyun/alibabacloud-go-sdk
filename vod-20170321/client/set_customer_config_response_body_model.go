// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCustomerConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCustomerConfigResponseBody
	GetRequestId() *string
}

type SetCustomerConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCustomerConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCustomerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetCustomerConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCustomerConfigResponseBody) SetRequestId(v string) *SetCustomerConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCustomerConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
