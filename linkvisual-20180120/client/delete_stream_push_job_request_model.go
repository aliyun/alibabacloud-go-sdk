// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamPushJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DeleteStreamPushJobRequest
	GetDeviceName() *string
	SetIotId(v string) *DeleteStreamPushJobRequest
	GetIotId() *string
	SetIotInstanceId(v string) *DeleteStreamPushJobRequest
	GetIotInstanceId() *string
	SetJobId(v string) *DeleteStreamPushJobRequest
	GetJobId() *string
	SetProductKey(v string) *DeleteStreamPushJobRequest
	GetProductKey() *string
}

type DeleteStreamPushJobRequest struct {
	// The DeviceName of the device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page of the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the device. This is the unique identifier issued by IoT Platform for the device. You can call [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) of IoT Platform to query the device ID.
	//
	// 	Notice:  **IotId*	- is the globally unique identifier of a device and corresponds to a combination of **ProductKey*	- and **DeviceName*	- on a one-to-one basis. If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. If you specify **IotId*	- together with **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// d7XmBoJhAr88C6PelXDF00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the ID of the current instance on the **Overview*	- tab of the IoT Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ProductKey of the product to which the device belongs.
	//
	// You can view the ProductKey on the **Products*	- page of the IoT Platform console or call [QueryProductList](https://help.aliyun.com/document_detail/69271.html) to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s DeleteStreamPushJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamPushJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamPushJobRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeleteStreamPushJobRequest) GetIotId() *string {
	return s.IotId
}

func (s *DeleteStreamPushJobRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *DeleteStreamPushJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteStreamPushJobRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *DeleteStreamPushJobRequest) SetDeviceName(v string) *DeleteStreamPushJobRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteStreamPushJobRequest) SetIotId(v string) *DeleteStreamPushJobRequest {
	s.IotId = &v
	return s
}

func (s *DeleteStreamPushJobRequest) SetIotInstanceId(v string) *DeleteStreamPushJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteStreamPushJobRequest) SetJobId(v string) *DeleteStreamPushJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteStreamPushJobRequest) SetProductKey(v string) *DeleteStreamPushJobRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteStreamPushJobRequest) Validate() error {
	return dara.Validate(s)
}
