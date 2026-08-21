// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupId(v string) *CreateDeviceGroupResponseBody
	GetDeviceGroupId() *string
	SetRequestId(v string) *CreateDeviceGroupResponseBody
	GetRequestId() *string
}

type CreateDeviceGroupResponseBody struct {
	// The device label ID.
	//
	// example:
	//
	// device-group-5191cf830a5e****
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceGroupResponseBody) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *CreateDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeviceGroupResponseBody) SetDeviceGroupId(v string) *CreateDeviceGroupResponseBody {
	s.DeviceGroupId = &v
	return s
}

func (s *CreateDeviceGroupResponseBody) SetRequestId(v string) *CreateDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
