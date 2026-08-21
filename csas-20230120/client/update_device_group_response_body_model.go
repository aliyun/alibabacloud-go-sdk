// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDeviceGroupResponseBody
	GetRequestId() *string
}

type UpdateDeviceGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeviceGroupResponseBody) SetRequestId(v string) *UpdateDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
