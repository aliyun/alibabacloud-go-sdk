// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPurchasedDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindPurchasedDeviceResponseBody
	GetRequestId() *string
}

type BindPurchasedDeviceResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindPurchasedDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindPurchasedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindPurchasedDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindPurchasedDeviceResponseBody) SetRequestId(v string) *BindPurchasedDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindPurchasedDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
