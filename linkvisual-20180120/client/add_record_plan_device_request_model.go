// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordPlanDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *AddRecordPlanDeviceRequest
	GetDeviceName() *string
	SetIotId(v string) *AddRecordPlanDeviceRequest
	GetIotId() *string
	SetIotInstanceId(v string) *AddRecordPlanDeviceRequest
	GetIotInstanceId() *string
	SetPlanId(v string) *AddRecordPlanDeviceRequest
	GetPlanId() *string
	SetProductKey(v string) *AddRecordPlanDeviceRequest
	GetProductKey() *string
	SetStreamType(v int32) *AddRecordPlanDeviceRequest
	GetStreamType() *int32
}

type AddRecordPlanDeviceRequest struct {
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
	// The IotId of the IPC device. The IotId is the unique identifier that is issued by IoT Platform to the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If an ID is displayed, you must specify the ID. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ID of the recording plan. You can call the [QueryRecordPlans](https://help.aliyun.com/document_detail/107896.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b60bf0dd55944c19d7817442bc9****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
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
	// - **0*	- (default): main stream.
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

func (s AddRecordPlanDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddRecordPlanDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *AddRecordPlanDeviceRequest) GetIotId() *string {
	return s.IotId
}

func (s *AddRecordPlanDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *AddRecordPlanDeviceRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *AddRecordPlanDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *AddRecordPlanDeviceRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *AddRecordPlanDeviceRequest) SetDeviceName(v string) *AddRecordPlanDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetIotId(v string) *AddRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetIotInstanceId(v string) *AddRecordPlanDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetPlanId(v string) *AddRecordPlanDeviceRequest {
	s.PlanId = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetProductKey(v string) *AddRecordPlanDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) SetStreamType(v int32) *AddRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

func (s *AddRecordPlanDeviceRequest) Validate() error {
	return dara.Validate(s)
}
