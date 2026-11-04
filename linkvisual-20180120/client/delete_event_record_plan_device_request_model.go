// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRecordPlanDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DeleteEventRecordPlanDeviceRequest
	GetDeviceName() *string
	SetIotId(v string) *DeleteEventRecordPlanDeviceRequest
	GetIotId() *string
	SetIotInstanceId(v string) *DeleteEventRecordPlanDeviceRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *DeleteEventRecordPlanDeviceRequest
	GetProductKey() *string
	SetStreamType(v int32) *DeleteEventRecordPlanDeviceRequest
	GetStreamType() *int32
}

type DeleteEventRecordPlanDeviceRequest struct {
	// The DeviceName of the IP camera.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The IotId of the IP camera. The IotId is issued by IoT Platform and serves as the unique identifier of the device.
	//
	// You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify this ID. Otherwise, the call fails.
	//
	// > If a public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IP camera belongs.
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
	// - **0*	- (default): main stream.
	//
	// - **1**: sub stream.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DeleteEventRecordPlanDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeleteEventRecordPlanDeviceRequest) GetIotId() *string {
	return s.IotId
}

func (s *DeleteEventRecordPlanDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *DeleteEventRecordPlanDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *DeleteEventRecordPlanDeviceRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *DeleteEventRecordPlanDeviceRequest) SetDeviceName(v string) *DeleteEventRecordPlanDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetIotId(v string) *DeleteEventRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetIotInstanceId(v string) *DeleteEventRecordPlanDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetProductKey(v string) *DeleteEventRecordPlanDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) SetStreamType(v int32) *DeleteEventRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceRequest) Validate() error {
	return dara.Validate(s)
}
