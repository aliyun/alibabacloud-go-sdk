// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordPlanDeviceByDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryRecordPlanDeviceByDeviceRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryRecordPlanDeviceByDeviceRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryRecordPlanDeviceByDeviceRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryRecordPlanDeviceByDeviceRequest
	GetProductKey() *string
	SetStreamType(v int32) *QueryRecordPlanDeviceByDeviceRequest
	GetStreamType() *int32
}

type QueryRecordPlanDeviceByDeviceRequest struct {
	// The DeviceName of the IP camera (IPC) device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The device ID of the IP camera (IPC) device. This ID is issued by IoT Platform and serves as the unique identifier of the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. As the unique identifier of the device, **IotId*	- has a one-to-one mapping with the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If an ID value exists, you must specify this value. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID value, refer to [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IP camera (IPC) device belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The stream type. Valid values:
	//
	// - **0**: main stream.
	//
	// - **1**: sub stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryRecordPlanDeviceByDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlanDeviceByDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDeviceByDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRecordPlanDeviceByDeviceRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryRecordPlanDeviceByDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryRecordPlanDeviceByDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryRecordPlanDeviceByDeviceRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetDeviceName(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetIotId(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetIotInstanceId(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetProductKey(v string) *QueryRecordPlanDeviceByDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) SetStreamType(v int32) *QueryRecordPlanDeviceByDeviceRequest {
	s.StreamType = &v
	return s
}

func (s *QueryRecordPlanDeviceByDeviceRequest) Validate() error {
	return dara.Validate(s)
}
