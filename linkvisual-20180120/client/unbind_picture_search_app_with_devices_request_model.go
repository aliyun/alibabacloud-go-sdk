// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPictureSearchAppWithDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *UnbindPictureSearchAppWithDevicesRequest
	GetAppInstanceId() *string
	SetDeviceIotIds(v []*string) *UnbindPictureSearchAppWithDevicesRequest
	GetDeviceIotIds() []*string
	SetIotInstanceId(v string) *UnbindPictureSearchAppWithDevicesRequest
	GetIotInstanceId() *string
}

type UnbindPictureSearchAppWithDevicesRequest struct {
	// Instance ID of the image search application instance. You can call the [QueryPictureSearchApps](https://help.aliyun.com/document_detail/206608.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a502d3fbab8410e8fd4be9037c7****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The list of device IoT IDs. You can specify the IoT IDs of V-series edge all-in-one machines or cameras.
	//
	// An IoT ID is the unique identifier issued by IoT Platform for a device. You can call the IoT Platform [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation to query the IoT ID.
	//
	// You can specify up to 20 device IoT IDs in a single call.
	//
	// This parameter is required.
	//
	// example:
	//
	// VrJwPGjC1sJlqPjZA3cxg4****
	DeviceIotIds []*string `json:"DeviceIotIds,omitempty" xml:"DeviceIotIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// - Enterprise Enterprise instance: You must specify this parameter. You can view the ID of your Enterprise Enterprise instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// - Public instance: You do not need to specify this parameter.
	//
	// example:
	//
	// iot-060a****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s UnbindPictureSearchAppWithDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindPictureSearchAppWithDevicesRequest) GoString() string {
	return s.String()
}

func (s *UnbindPictureSearchAppWithDevicesRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *UnbindPictureSearchAppWithDevicesRequest) GetDeviceIotIds() []*string {
	return s.DeviceIotIds
}

func (s *UnbindPictureSearchAppWithDevicesRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetAppInstanceId(v string) *UnbindPictureSearchAppWithDevicesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetDeviceIotIds(v []*string) *UnbindPictureSearchAppWithDevicesRequest {
	s.DeviceIotIds = v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesRequest) SetIotInstanceId(v string) *UnbindPictureSearchAppWithDevicesRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesRequest) Validate() error {
	return dara.Validate(s)
}
