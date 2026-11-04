// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *TriggerRecordRequest
	GetDeviceName() *string
	SetIotId(v string) *TriggerRecordRequest
	GetIotId() *string
	SetIotInstanceId(v string) *TriggerRecordRequest
	GetIotInstanceId() *string
	SetPreRecordDuration(v int32) *TriggerRecordRequest
	GetPreRecordDuration() *int32
	SetProductKey(v string) *TriggerRecordRequest
	GetProductKey() *string
	SetRecordDuration(v int32) *TriggerRecordRequest
	GetRecordDuration() *int32
	SetStreamType(v int32) *TriggerRecordRequest
	GetStreamType() *int32
}

type TriggerRecordRequest struct {
	// The DeviceName of the IPC device. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
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
	// zLZyi6aOLyOSHa****yD000100
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
	// iot-cn-n6w1y5****0t
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The pre-recording duration. Valid values: 0 to 10. Unit: seconds.
	//
	// >The pre-recording duration is generated only when the device supports the pre-recording feature. Otherwise, the pre-recording duration is 0.
	//
	// example:
	//
	// 5
	PreRecordDuration *int32 `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	// The ProductKey of the product to which the IPC device belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The recording duration, including the pre-recording duration. Valid values: 10 to 7200. Unit: seconds.
	//
	// > The maximum length of a single recording file is 30 minutes. If the recording duration exceeds 30 minutes, the recording is automatically split into multiple recording files.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	RecordDuration *int32 `json:"RecordDuration,omitempty" xml:"RecordDuration,omitempty"`
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

func (s TriggerRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerRecordRequest) GoString() string {
	return s.String()
}

func (s *TriggerRecordRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *TriggerRecordRequest) GetIotId() *string {
	return s.IotId
}

func (s *TriggerRecordRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *TriggerRecordRequest) GetPreRecordDuration() *int32 {
	return s.PreRecordDuration
}

func (s *TriggerRecordRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *TriggerRecordRequest) GetRecordDuration() *int32 {
	return s.RecordDuration
}

func (s *TriggerRecordRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *TriggerRecordRequest) SetDeviceName(v string) *TriggerRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *TriggerRecordRequest) SetIotId(v string) *TriggerRecordRequest {
	s.IotId = &v
	return s
}

func (s *TriggerRecordRequest) SetIotInstanceId(v string) *TriggerRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *TriggerRecordRequest) SetPreRecordDuration(v int32) *TriggerRecordRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *TriggerRecordRequest) SetProductKey(v string) *TriggerRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *TriggerRecordRequest) SetRecordDuration(v int32) *TriggerRecordRequest {
	s.RecordDuration = &v
	return s
}

func (s *TriggerRecordRequest) SetStreamType(v int32) *TriggerRecordRequest {
	s.StreamType = &v
	return s
}

func (s *TriggerRecordRequest) Validate() error {
	return dara.Validate(s)
}
