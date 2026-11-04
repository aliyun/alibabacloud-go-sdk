// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamSnapshotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *CreateStreamSnapshotJobRequest
	GetDeviceName() *string
	SetIotId(v string) *CreateStreamSnapshotJobRequest
	GetIotId() *string
	SetIotInstanceId(v string) *CreateStreamSnapshotJobRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *CreateStreamSnapshotJobRequest
	GetProductKey() *string
	SetSnapshotInterval(v int32) *CreateStreamSnapshotJobRequest
	GetSnapshotInterval() *int32
	SetStreamType(v int32) *CreateStreamSnapshotJobRequest
	GetStreamType() *int32
}

type CreateStreamSnapshotJobRequest struct {
	// The **DeviceName*	- of the device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the device. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  **IotId*	- uniquely identifies the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// d7XmBoJhAr88C6PelXDF00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID. You can view the ID of the current instance on the **Instance Overview*	- tab in the IoT Platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// iot-cn-n6w1y59****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The **ProductKey*	- of the product to which the device belongs.
	//
	// You can view the **ProductKey*	- on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information of all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The snapshot interval. The value ranges from 60 to 1,800 seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 600
	SnapshotInterval *int32 `json:"SnapshotInterval,omitempty" xml:"SnapshotInterval,omitempty"`
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

func (s CreateStreamSnapshotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotJobRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateStreamSnapshotJobRequest) GetIotId() *string {
	return s.IotId
}

func (s *CreateStreamSnapshotJobRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *CreateStreamSnapshotJobRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *CreateStreamSnapshotJobRequest) GetSnapshotInterval() *int32 {
	return s.SnapshotInterval
}

func (s *CreateStreamSnapshotJobRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *CreateStreamSnapshotJobRequest) SetDeviceName(v string) *CreateStreamSnapshotJobRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetIotId(v string) *CreateStreamSnapshotJobRequest {
	s.IotId = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetIotInstanceId(v string) *CreateStreamSnapshotJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetProductKey(v string) *CreateStreamSnapshotJobRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetSnapshotInterval(v int32) *CreateStreamSnapshotJobRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) SetStreamType(v int32) *CreateStreamSnapshotJobRequest {
	s.StreamType = &v
	return s
}

func (s *CreateStreamSnapshotJobRequest) Validate() error {
	return dara.Validate(s)
}
