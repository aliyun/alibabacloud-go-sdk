// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveStreamingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *StopLiveStreamingRequest
	GetDeviceName() *string
	SetIotId(v string) *StopLiveStreamingRequest
	GetIotId() *string
	SetIotInstanceId(v string) *StopLiveStreamingRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *StopLiveStreamingRequest
	GetProductKey() *string
	SetStreamType(v int32) *StopLiveStreamingRequest
	GetStreamType() *int32
}

type StopLiveStreamingRequest struct {
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
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. **IotId*	- uniquely identifies the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// C47T6xwp6ms4bNlkHRWCg4****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// Instance ID of the instance. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If the instance has an ID, you must specify instance ID. Otherwise, the call fails.
	//
	// - If the instance does not have an <props="intl"><ph>**Instance Overview*	- page or </ph>ID, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
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
	// - **0*	- (default): main stream.
	//
	// - **1**: sub stream.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s StopLiveStreamingRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLiveStreamingRequest) GoString() string {
	return s.String()
}

func (s *StopLiveStreamingRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *StopLiveStreamingRequest) GetIotId() *string {
	return s.IotId
}

func (s *StopLiveStreamingRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *StopLiveStreamingRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *StopLiveStreamingRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *StopLiveStreamingRequest) SetDeviceName(v string) *StopLiveStreamingRequest {
	s.DeviceName = &v
	return s
}

func (s *StopLiveStreamingRequest) SetIotId(v string) *StopLiveStreamingRequest {
	s.IotId = &v
	return s
}

func (s *StopLiveStreamingRequest) SetIotInstanceId(v string) *StopLiveStreamingRequest {
	s.IotInstanceId = &v
	return s
}

func (s *StopLiveStreamingRequest) SetProductKey(v string) *StopLiveStreamingRequest {
	s.ProductKey = &v
	return s
}

func (s *StopLiveStreamingRequest) SetStreamType(v int32) *StopLiveStreamingRequest {
	s.StreamType = &v
	return s
}

func (s *StopLiveStreamingRequest) Validate() error {
	return dara.Validate(s)
}
