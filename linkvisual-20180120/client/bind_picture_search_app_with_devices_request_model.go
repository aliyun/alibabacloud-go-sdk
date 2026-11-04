// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPictureSearchAppWithDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *BindPictureSearchAppWithDevicesRequest
	GetAppInstanceId() *string
	SetDeviceIotIds(v []*string) *BindPictureSearchAppWithDevicesRequest
	GetDeviceIotIds() []*string
	SetIotInstanceId(v string) *BindPictureSearchAppWithDevicesRequest
	GetIotInstanceId() *string
}

type BindPictureSearchAppWithDevicesRequest struct {
	// Instance ID of the image search application instance. You can call the [QueryPictureSearchApps](https://help.aliyun.com/document_detail/206608.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a502d3fbab8410e8fd4be9037c7****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The list of device IoT IDs. You can specify the IoT IDs of V-series edge all-in-one devices or cameras.
	//
	// An IoT ID is a unique identifier issued by IoT Platform for a device.
	//
	// You can call the IoT Platform [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation to query the IoT ID.
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

func (s BindPictureSearchAppWithDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s BindPictureSearchAppWithDevicesRequest) GoString() string {
	return s.String()
}

func (s *BindPictureSearchAppWithDevicesRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *BindPictureSearchAppWithDevicesRequest) GetDeviceIotIds() []*string {
	return s.DeviceIotIds
}

func (s *BindPictureSearchAppWithDevicesRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *BindPictureSearchAppWithDevicesRequest) SetAppInstanceId(v string) *BindPictureSearchAppWithDevicesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesRequest) SetDeviceIotIds(v []*string) *BindPictureSearchAppWithDevicesRequest {
	s.DeviceIotIds = v
	return s
}

func (s *BindPictureSearchAppWithDevicesRequest) SetIotInstanceId(v string) *BindPictureSearchAppWithDevicesRequest {
	s.IotInstanceId = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesRequest) Validate() error {
	return dara.Validate(s)
}
