// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeAndroidInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *DowngradeAndroidInstanceGroupResponseBody
	GetOrderId() *string
	SetRequestId(v string) *DowngradeAndroidInstanceGroupResponseBody
	GetRequestId() *string
}

type DowngradeAndroidInstanceGroupResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 22326560487****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3AF82CE1-2801-52CE-BF64-B491DD7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DowngradeAndroidInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DowngradeAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DowngradeAndroidInstanceGroupResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DowngradeAndroidInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DowngradeAndroidInstanceGroupResponseBody) SetOrderId(v string) *DowngradeAndroidInstanceGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *DowngradeAndroidInstanceGroupResponseBody) SetRequestId(v string) *DowngradeAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DowngradeAndroidInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
