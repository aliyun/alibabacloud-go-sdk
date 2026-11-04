// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEventRecordPlanDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *AddEventRecordPlanDeviceRequest
	GetDeviceName() *string
	SetIotId(v string) *AddEventRecordPlanDeviceRequest
	GetIotId() *string
	SetIotInstanceId(v string) *AddEventRecordPlanDeviceRequest
	GetIotInstanceId() *string
	SetPlanId(v string) *AddEventRecordPlanDeviceRequest
	GetPlanId() *string
	SetProductKey(v string) *AddEventRecordPlanDeviceRequest
	GetProductKey() *string
	SetStreamType(v int32) *AddEventRecordPlanDeviceRequest
	GetStreamType() *int32
}

type AddEventRecordPlanDeviceRequest struct {
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
	// The IotId of the IP camera. The IotId is the unique identifier that is issued by IoT Platform for the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If an ID value is available, you must specify this ID. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID value, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ID of the intelligent alert event recording plan. You can call the [QueryEventRecordPlans](https://help.aliyun.com/document_detail/107907.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b60bf0dd55944c19d7817442bc9****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
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

func (s AddEventRecordPlanDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEventRecordPlanDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddEventRecordPlanDeviceRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *AddEventRecordPlanDeviceRequest) GetIotId() *string {
	return s.IotId
}

func (s *AddEventRecordPlanDeviceRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *AddEventRecordPlanDeviceRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *AddEventRecordPlanDeviceRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *AddEventRecordPlanDeviceRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *AddEventRecordPlanDeviceRequest) SetDeviceName(v string) *AddEventRecordPlanDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetIotId(v string) *AddEventRecordPlanDeviceRequest {
	s.IotId = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetIotInstanceId(v string) *AddEventRecordPlanDeviceRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetPlanId(v string) *AddEventRecordPlanDeviceRequest {
	s.PlanId = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetProductKey(v string) *AddEventRecordPlanDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) SetStreamType(v int32) *AddEventRecordPlanDeviceRequest {
	s.StreamType = &v
	return s
}

func (s *AddEventRecordPlanDeviceRequest) Validate() error {
	return dara.Validate(s)
}
