// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDeviceGroupMatchDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveDeviceGroupMatchDevicesResponseBody
	GetRequestId() *string
}

type RemoveDeviceGroupMatchDevicesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDeviceGroupMatchDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDeviceGroupMatchDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDeviceGroupMatchDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDeviceGroupMatchDevicesResponseBody) SetRequestId(v string) *RemoveDeviceGroupMatchDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDeviceGroupMatchDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}
