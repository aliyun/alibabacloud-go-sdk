// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPurchasedDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindPurchasedDeviceResponseBody
	GetRequestId() *string
}

type UnbindPurchasedDeviceResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindPurchasedDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindPurchasedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindPurchasedDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindPurchasedDeviceResponseBody) SetRequestId(v string) *UnbindPurchasedDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindPurchasedDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
