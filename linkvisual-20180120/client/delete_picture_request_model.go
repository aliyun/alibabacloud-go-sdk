// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePictureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DeletePictureRequest
	GetDeviceName() *string
	SetIotId(v string) *DeletePictureRequest
	GetIotId() *string
	SetIotInstanceId(v string) *DeletePictureRequest
	GetIotInstanceId() *string
	SetPictureIdList(v []*string) *DeletePictureRequest
	GetPictureIdList() []*string
	SetProductKey(v string) *DeletePictureRequest
	GetProductKey() *string
}

type DeletePictureRequest struct {
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
	// The IotId of the IPC device. The IotId is a unique identifier issued by IoT Platform for the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is a unique identifier for the device and has a one-to-one mapping with the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// 1CJ0****e7qiq8VkQi1000000
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If the instance has an ID, you must specify the ID. Otherwise, the call fails.
	//
	// - If the instance does not have an <props="intl"><ph>**Instance Overview*	- page or an </ph>ID, you do not need to specify this parameter.
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-cn-n6w1y******
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The list of picture IDs. Picture IDs are generated when you call the [TriggerCapturePicture](https://help.aliyun.com/document_detail/107914.html) operation.
	//
	// > You can specify up to 50 picture IDs in a single call.
	//
	// This parameter is required.
	//
	// example:
	//
	// cFc5MzZ****npYT1***TDFzX1hrVzA0Z****
	PictureIdList []*string `json:"PictureIdList,omitempty" xml:"PictureIdList,omitempty" type:"Repeated"`
	// The ProductKey of the product to which the IPC device belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeletePictureRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePictureRequest) GoString() string {
	return s.String()
}

func (s *DeletePictureRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeletePictureRequest) GetIotId() *string {
	return s.IotId
}

func (s *DeletePictureRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *DeletePictureRequest) GetPictureIdList() []*string {
	return s.PictureIdList
}

func (s *DeletePictureRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *DeletePictureRequest) SetDeviceName(v string) *DeletePictureRequest {
	s.DeviceName = &v
	return s
}

func (s *DeletePictureRequest) SetIotId(v string) *DeletePictureRequest {
	s.IotId = &v
	return s
}

func (s *DeletePictureRequest) SetIotInstanceId(v string) *DeletePictureRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeletePictureRequest) SetPictureIdList(v []*string) *DeletePictureRequest {
	s.PictureIdList = v
	return s
}

func (s *DeletePictureRequest) SetProductKey(v string) *DeletePictureRequest {
	s.ProductKey = &v
	return s
}

func (s *DeletePictureRequest) Validate() error {
	return dara.Validate(s)
}
