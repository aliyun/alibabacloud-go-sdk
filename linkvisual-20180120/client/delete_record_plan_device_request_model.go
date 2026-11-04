// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordPlanDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DeleteRecordPlanDeviceRequest
	GetDeviceName() *string
	SetIotId(v string) *DeleteRecordPlanDeviceRequest
	GetIotId() *string
	SetIotInstanceId(v string) *DeleteRecordPlanDeviceRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *DeleteRecordPlanDeviceRequest
	GetProductKey() *string
	SetStreamType(v int32) *DeleteRecordPlanDeviceRequest
	GetStreamType() *int32
}

type DeleteRecordPlanDeviceRequest struct {
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
	// The ID of the IPC device. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- uniquely identifies the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify this ID. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
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
	// The stream type. Valid values:
	//
	// - **0**: main stream
	//
	// - **1**: sub stream
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DeleteRecordPlanDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeleteRecordPlanDeviceRequest) GetIotId() *string {
	return s.IotId
}

func (s *DeleteRecordPlanDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *DeleteRecordPlanDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *DeleteRecordPlanDeviceRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *DeleteRecordPlanDeviceRequest) SetDeviceName(v string) *DeleteRecordPlanDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetIotId(v string) *DeleteRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetIotInstanceId(v string) *DeleteRecordPlanDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetProductKey(v string) *DeleteRecordPlanDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) SetStreamType(v int32) *DeleteRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

func (s *DeleteRecordPlanDeviceRequest) Validate() error {
	return dara.Validate(s)
}
