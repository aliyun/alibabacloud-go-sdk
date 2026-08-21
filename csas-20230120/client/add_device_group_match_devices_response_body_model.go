// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceGroupMatchDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDeviceGroupMatchDevicesResponseBody
	GetRequestId() *string
}

type AddDeviceGroupMatchDevicesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 578B9ADD-FB3E-57E4-AB7D-77BC9D39B591
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDeviceGroupMatchDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceGroupMatchDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *AddDeviceGroupMatchDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDeviceGroupMatchDevicesResponseBody) SetRequestId(v string) *AddDeviceGroupMatchDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDeviceGroupMatchDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}
