// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddAddressResponseBody
	GetRequestId() *string
}

type AddAddressResponseBody struct {
	// example:
	//
	// 2EFCFE18-78F8-5079-B312-07***48B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AddAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAddressResponseBody) SetRequestId(v string) *AddAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
