// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDevicePictureLifeCycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDay(v int32) *SetDevicePictureLifeCycleRequest
	GetDay() *int32
	SetDeviceName(v string) *SetDevicePictureLifeCycleRequest
	GetDeviceName() *string
	SetIotId(v string) *SetDevicePictureLifeCycleRequest
	GetIotId() *string
	SetIotInstanceId(v string) *SetDevicePictureLifeCycleRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *SetDevicePictureLifeCycleRequest
	GetProductKey() *string
}

type SetDevicePictureLifeCycleRequest struct {
	// The storage period of images in the cloud. Unit: days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// The DeviceName of the IPC device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The IotId of the IPC device. The IotId is the unique identifier that is issued by IoT Platform for the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHR****c300
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify the ID. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IPC device belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s SetDevicePictureLifeCycleRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDevicePictureLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *SetDevicePictureLifeCycleRequest) GetDay() *int32 {
	return s.Day
}

func (s *SetDevicePictureLifeCycleRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *SetDevicePictureLifeCycleRequest) GetIotId() *string {
	return s.IotId
}

func (s *SetDevicePictureLifeCycleRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *SetDevicePictureLifeCycleRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *SetDevicePictureLifeCycleRequest) SetDay(v int32) *SetDevicePictureLifeCycleRequest {
	s.Day = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetDeviceName(v string) *SetDevicePictureLifeCycleRequest {
	s.DeviceName = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetIotId(v string) *SetDevicePictureLifeCycleRequest {
	s.IotId = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetIotInstanceId(v string) *SetDevicePictureLifeCycleRequest {
	s.IotInstanceId = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) SetProductKey(v string) *SetDevicePictureLifeCycleRequest {
	s.ProductKey = &v
	return s
}

func (s *SetDevicePictureLifeCycleRequest) Validate() error {
	return dara.Validate(s)
}
