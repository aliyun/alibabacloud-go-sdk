// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceEventRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryDeviceEventRecordRequest
	GetDeviceName() *string
	SetEventId(v string) *QueryDeviceEventRecordRequest
	GetEventId() *string
	SetIotId(v string) *QueryDeviceEventRecordRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryDeviceEventRecordRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryDeviceEventRecordRequest
	GetProductKey() *string
}

type QueryDeviceEventRecordRequest struct {
	// The DeviceName of the IP camera (IPC) device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page of the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the intelligent alert event. You can call the [QueryDeviceEvent](https://help.aliyun.com/document_detail/107916.html) operation to obtain the event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 907084996e1d44788d0*****43906_160803*****37
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The IotId of the IP camera (IPC) device. The IotId is the globally unique identifier of the device that is issued by IoT Platform. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the globally unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The ID of the instance.
	//
	// <props="china">You can view the **ID*	- of the instance on the **Instance Overview*	- page of the IoT Platform console. If your instance has an ID, you must specify the ID. Otherwise, the call fails.
	//
	// <props="china">
	//
	// > If your public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// <props="intl">You do not need to specify this parameter.
	//
	// example:
	//
	// iot-cn-n6w1****z0t
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IP camera (IPC) device belongs.
	//
	// You can view the ProductKey on the **Products*	- page of the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information of all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryDeviceEventRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceEventRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRecordRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryDeviceEventRecordRequest) GetEventId() *string {
	return s.EventId
}

func (s *QueryDeviceEventRecordRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryDeviceEventRecordRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryDeviceEventRecordRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryDeviceEventRecordRequest) SetDeviceName(v string) *QueryDeviceEventRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetEventId(v string) *QueryDeviceEventRecordRequest {
	s.EventId = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetIotId(v string) *QueryDeviceEventRecordRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetIotInstanceId(v string) *QueryDeviceEventRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) SetProductKey(v string) *QueryDeviceEventRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDeviceEventRecordRequest) Validate() error {
	return dara.Validate(s)
}
