// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStreamSnapshotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryStreamSnapshotJobRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryStreamSnapshotJobRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryStreamSnapshotJobRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryStreamSnapshotJobRequest
	GetProductKey() *string
}

type QueryStreamSnapshotJobRequest struct {
	// The DeviceName of the device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the device. The ID is issued by IoT Platform and is the unique identifier of the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  As the unique identifier of the device, **IotId*	- has a one-to-one mapping with the combination of **ProductKey*	- and **DeviceName**. If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
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
	// The ProductKey of the product to which the device belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryStreamSnapshotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *QueryStreamSnapshotJobRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryStreamSnapshotJobRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryStreamSnapshotJobRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryStreamSnapshotJobRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryStreamSnapshotJobRequest) SetDeviceName(v string) *QueryStreamSnapshotJobRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryStreamSnapshotJobRequest) SetIotId(v string) *QueryStreamSnapshotJobRequest {
	s.IotId = &v
	return s
}

func (s *QueryStreamSnapshotJobRequest) SetIotInstanceId(v string) *QueryStreamSnapshotJobRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryStreamSnapshotJobRequest) SetProductKey(v string) *QueryStreamSnapshotJobRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryStreamSnapshotJobRequest) Validate() error {
	return dara.Validate(s)
}
