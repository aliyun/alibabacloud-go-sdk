// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordByRecordIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryRecordByRecordIdRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryRecordByRecordIdRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryRecordByRecordIdRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryRecordByRecordIdRequest
	GetProductKey() *string
	SetRecordId(v string) *QueryRecordByRecordIdRequest
	GetRecordId() *string
}

type QueryRecordByRecordIdRequest struct {
	// The DeviceName of the IP camera.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page of the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The ID of the IP camera. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- uniquely identifies the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// <props="china">You can view the **ID*	- of the current instance on the **Instance Overview*	- page of the IoT Platform console. If the instance has an ID, you must specify the ID. Otherwise, the call fails.
	//
	// <props="china">
	//
	// > If the public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// <props="intl">You do not need to specify this parameter.
	//
	// example:
	//
	// iot-060a01****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IP camera belongs.
	//
	// You can view the ProductKey on the **Products*	- page of the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query all products under your account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The ID of the recording file. You can call the [QueryRecord](https://help.aliyun.com/document_detail/107922.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2c5bf5b6a8704bd3a6a1517fd125****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s QueryRecordByRecordIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordByRecordIdRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordByRecordIdRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRecordByRecordIdRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryRecordByRecordIdRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryRecordByRecordIdRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryRecordByRecordIdRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *QueryRecordByRecordIdRequest) SetDeviceName(v string) *QueryRecordByRecordIdRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetIotId(v string) *QueryRecordByRecordIdRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetIotInstanceId(v string) *QueryRecordByRecordIdRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetProductKey(v string) *QueryRecordByRecordIdRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) SetRecordId(v string) *QueryRecordByRecordIdRequest {
	s.RecordId = &v
	return s
}

func (s *QueryRecordByRecordIdRequest) Validate() error {
	return dara.Validate(s)
}
