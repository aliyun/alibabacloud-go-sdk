// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTriggeredRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *StopTriggeredRecordRequest
	GetDeviceName() *string
	SetIotId(v string) *StopTriggeredRecordRequest
	GetIotId() *string
	SetIotInstanceId(v string) *StopTriggeredRecordRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *StopTriggeredRecordRequest
	GetProductKey() *string
	SetRecordId(v string) *StopTriggeredRecordRequest
	GetRecordId() *string
}

type StopTriggeredRecordRequest struct {
	// The DeviceName of the IP camera. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The device ID of the IP camera. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the device ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- uniquely identifies the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// The device ID of the IP camera. The ID is issued by IoT Platform and uniquely identifies the device.
	//
	// You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the device ID.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHRWCg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify this ID. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IP camera belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The recording ID. You can call the [TriggerRecord](https://help.aliyun.com/document_detail/107915.html) operation to query the recording ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fc82774fa749485bad7d719f9670****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s StopTriggeredRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTriggeredRecordRequest) GoString() string {
	return s.String()
}

func (s *StopTriggeredRecordRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *StopTriggeredRecordRequest) GetIotId() *string {
	return s.IotId
}

func (s *StopTriggeredRecordRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *StopTriggeredRecordRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *StopTriggeredRecordRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *StopTriggeredRecordRequest) SetDeviceName(v string) *StopTriggeredRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetIotId(v string) *StopTriggeredRecordRequest {
	s.IotId = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetIotInstanceId(v string) *StopTriggeredRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetProductKey(v string) *StopTriggeredRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *StopTriggeredRecordRequest) SetRecordId(v string) *StopTriggeredRecordRequest {
	s.RecordId = &v
	return s
}

func (s *StopTriggeredRecordRequest) Validate() error {
	return dara.Validate(s)
}
