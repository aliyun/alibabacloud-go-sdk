// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamSnapshotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *DeleteStreamSnapshotJobRequest
	GetDeviceName() *string
	SetIotId(v string) *DeleteStreamSnapshotJobRequest
	GetIotId() *string
	SetIotInstanceId(v string) *DeleteStreamSnapshotJobRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *DeleteStreamSnapshotJobRequest
	GetProductKey() *string
	SetStreamType(v int32) *DeleteStreamSnapshotJobRequest
	GetStreamType() *int32
}

type DeleteStreamSnapshotJobRequest struct {
	// The DeviceName of the device. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the device. The ID is issued by IoT Platform and is the unique identifier of the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName*	- in a one-to-one mapping. If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
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
	// The **ProductKey*	- of the product to which the device belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information of all products under the current account.
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

func (s DeleteStreamSnapshotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamSnapshotJobRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *DeleteStreamSnapshotJobRequest) GetIotId() *string {
	return s.IotId
}

func (s *DeleteStreamSnapshotJobRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *DeleteStreamSnapshotJobRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *DeleteStreamSnapshotJobRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *DeleteStreamSnapshotJobRequest) SetDeviceName(v string) *DeleteStreamSnapshotJobRequest {
	s.DeviceName = &v
	return s
}

func (s *DeleteStreamSnapshotJobRequest) SetIotId(v string) *DeleteStreamSnapshotJobRequest {
	s.IotId = &v
	return s
}

func (s *DeleteStreamSnapshotJobRequest) SetIotInstanceId(v string) *DeleteStreamSnapshotJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteStreamSnapshotJobRequest) SetProductKey(v string) *DeleteStreamSnapshotJobRequest {
	s.ProductKey = &v
	return s
}

func (s *DeleteStreamSnapshotJobRequest) SetStreamType(v int32) *DeleteStreamSnapshotJobRequest {
	s.StreamType = &v
	return s
}

func (s *DeleteStreamSnapshotJobRequest) Validate() error {
	return dara.Validate(s)
}
