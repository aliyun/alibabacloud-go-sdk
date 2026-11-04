// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventRecordPlanDeviceByDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryEventRecordPlanDeviceByDeviceRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryEventRecordPlanDeviceByDeviceRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryEventRecordPlanDeviceByDeviceRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryEventRecordPlanDeviceByDeviceRequest
	GetProductKey() *string
	SetStreamType(v int32) *QueryEventRecordPlanDeviceByDeviceRequest
	GetStreamType() *int32
}

type QueryEventRecordPlanDeviceByDeviceRequest struct {
	// The device name of the IPC device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The device ID of the IPC device. This is the unique identifier issued by IoT Platform for the device. You can call [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) of IoT Platform to query the device ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. **IotId*	- is a globally unique identifier of the device and corresponds to a combination of **ProductKey*	- and **DeviceName*	- on a one-to-one basis. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Overview*	- page in the IoT Platform console. If the ID value exists, you must specify this ID value. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID value, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IPC device belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console, or call [QueryProductList](https://help.aliyun.com/document_detail/69271.html) to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The stream type. Valid values:
	//
	// - **0*	- (default): main stream
	//
	// - **1**: sub stream
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryEventRecordPlanDeviceByDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEventRecordPlanDeviceByDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetDeviceName(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetIotId(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetIotInstanceId(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetProductKey(v string) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) SetStreamType(v int32) *QueryEventRecordPlanDeviceByDeviceRequest {
	s.StreamType = &v
	return s
}

func (s *QueryEventRecordPlanDeviceByDeviceRequest) Validate() error {
	return dara.Validate(s)
}
