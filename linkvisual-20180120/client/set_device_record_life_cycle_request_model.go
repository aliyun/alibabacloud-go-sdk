// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceRecordLifeCycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDay(v int32) *SetDeviceRecordLifeCycleRequest
	GetDay() *int32
	SetDeviceName(v string) *SetDeviceRecordLifeCycleRequest
	GetDeviceName() *string
	SetIotId(v string) *SetDeviceRecordLifeCycleRequest
	GetIotId() *string
	SetIotInstanceId(v string) *SetDeviceRecordLifeCycleRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *SetDeviceRecordLifeCycleRequest
	GetProductKey() *string
}

type SetDeviceRecordLifeCycleRequest struct {
	// The recording lifecycle, which is the number of days that recording files are stored in the cloud.
	//
	// > You must set the recording lifecycle to a value greater than 0 to use other recording features provided by the video service. The maximum recording lifecycle is 3650 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	Day *int32 `json:"Day,omitempty" xml:"Day,omitempty"`
	// The DeviceName of the IPC device. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The IotId of the IPC device. The IotId is the unique identifier issued by IoT Platform for the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// d7XmBoJhAr88C6PelXDF00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If an ID value is available, you must specify it. Otherwise, the call fails.
	//
	// - If the <props="intl"><ph>**Instance Overview*	- page or </ph>ID value is not available, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IPC device belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s SetDeviceRecordLifeCycleRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceRecordLifeCycleRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceRecordLifeCycleRequest) GetDay() *int32 {
	return s.Day
}

func (s *SetDeviceRecordLifeCycleRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *SetDeviceRecordLifeCycleRequest) GetIotId() *string {
	return s.IotId
}

func (s *SetDeviceRecordLifeCycleRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *SetDeviceRecordLifeCycleRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *SetDeviceRecordLifeCycleRequest) SetDay(v int32) *SetDeviceRecordLifeCycleRequest {
	s.Day = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetDeviceName(v string) *SetDeviceRecordLifeCycleRequest {
	s.DeviceName = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetIotId(v string) *SetDeviceRecordLifeCycleRequest {
	s.IotId = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetIotInstanceId(v string) *SetDeviceRecordLifeCycleRequest {
	s.IotInstanceId = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) SetProductKey(v string) *SetDeviceRecordLifeCycleRequest {
	s.ProductKey = &v
	return s
}

func (s *SetDeviceRecordLifeCycleRequest) Validate() error {
	return dara.Validate(s)
}
