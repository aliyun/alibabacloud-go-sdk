// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerCapturePictureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *TriggerCapturePictureRequest
	GetDeviceName() *string
	SetIotId(v string) *TriggerCapturePictureRequest
	GetIotId() *string
	SetIotInstanceId(v string) *TriggerCapturePictureRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *TriggerCapturePictureRequest
	GetProductKey() *string
}

type TriggerCapturePictureRequest struct {
	// The DeviceName of the IPC device. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the IPC device. The ID is the unique identifier that is issued by IoT Platform to the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- parameter specifies a globally unique identifier (GUID) of the device, which corresponds to a combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// <props="china">You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify the ID. Otherwise, the call fails.
	//
	// <props="china">
	//
	// > If the public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// <props="intl">You do not need to specify this parameter.
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IPC device belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s TriggerCapturePictureRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerCapturePictureRequest) GoString() string {
	return s.String()
}

func (s *TriggerCapturePictureRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *TriggerCapturePictureRequest) GetIotId() *string {
	return s.IotId
}

func (s *TriggerCapturePictureRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *TriggerCapturePictureRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *TriggerCapturePictureRequest) SetDeviceName(v string) *TriggerCapturePictureRequest {
	s.DeviceName = &v
	return s
}

func (s *TriggerCapturePictureRequest) SetIotId(v string) *TriggerCapturePictureRequest {
	s.IotId = &v
	return s
}

func (s *TriggerCapturePictureRequest) SetIotInstanceId(v string) *TriggerCapturePictureRequest {
	s.IotInstanceId = &v
	return s
}

func (s *TriggerCapturePictureRequest) SetProductKey(v string) *TriggerCapturePictureRequest {
	s.ProductKey = &v
	return s
}

func (s *TriggerCapturePictureRequest) Validate() error {
	return dara.Validate(s)
}
