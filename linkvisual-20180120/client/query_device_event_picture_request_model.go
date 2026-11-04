// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceEventPictureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryDeviceEventPictureRequest
	GetDeviceName() *string
	SetEventId(v string) *QueryDeviceEventPictureRequest
	GetEventId() *string
	SetIotId(v string) *QueryDeviceEventPictureRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryDeviceEventPictureRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryDeviceEventPictureRequest
	GetProductKey() *string
}

type QueryDeviceEventPictureRequest struct {
	// The name of the IPC device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the intelligent alert event. You can obtain the ID by calling the [QueryDeviceEvent](https://help.aliyun.com/document_detail/107916.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ae6b74f1dc7d402****32c21018****f1608036****71
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the IPC device. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- uniquely identifies the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
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
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-cn-n6w1y*****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
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

func (s QueryDeviceEventPictureRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceEventPictureRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventPictureRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryDeviceEventPictureRequest) GetEventId() *string {
	return s.EventId
}

func (s *QueryDeviceEventPictureRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryDeviceEventPictureRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryDeviceEventPictureRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryDeviceEventPictureRequest) SetDeviceName(v string) *QueryDeviceEventPictureRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetEventId(v string) *QueryDeviceEventPictureRequest {
	s.EventId = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetIotId(v string) *QueryDeviceEventPictureRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetIotInstanceId(v string) *QueryDeviceEventPictureRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) SetProductKey(v string) *QueryDeviceEventPictureRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDeviceEventPictureRequest) Validate() error {
	return dara.Validate(s)
}
