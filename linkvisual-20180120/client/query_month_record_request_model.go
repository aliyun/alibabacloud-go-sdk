// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMonthRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceName(v string) *QueryMonthRecordRequest
	GetDeviceName() *string
	SetIotId(v string) *QueryMonthRecordRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryMonthRecordRequest
	GetIotInstanceId() *string
	SetMonth(v string) *QueryMonthRecordRequest
	GetMonth() *string
	SetProductKey(v string) *QueryMonthRecordRequest
	GetProductKey() *string
}

type QueryMonthRecordRequest struct {
	// The DeviceName of the IP camera. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The IotId of the IP camera. The IotId is the unique identifier issued by IoT Platform for the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// VrJwPGjC1sJlqPjZA3cxg4****
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
	// iot-cn-n6w1y****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The month to query, in the YYYYMM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 201901
	Month *string `json:"Month,omitempty" xml:"Month,omitempty"`
	// The ProductKey of the product to which the IP camera belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query all products under the current account.	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryMonthRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMonthRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryMonthRecordRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryMonthRecordRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryMonthRecordRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryMonthRecordRequest) GetMonth() *string {
	return s.Month
}

func (s *QueryMonthRecordRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryMonthRecordRequest) SetDeviceName(v string) *QueryMonthRecordRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryMonthRecordRequest) SetIotId(v string) *QueryMonthRecordRequest {
	s.IotId = &v
	return s
}

func (s *QueryMonthRecordRequest) SetIotInstanceId(v string) *QueryMonthRecordRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryMonthRecordRequest) SetMonth(v string) *QueryMonthRecordRequest {
	s.Month = &v
	return s
}

func (s *QueryMonthRecordRequest) SetProductKey(v string) *QueryMonthRecordRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryMonthRecordRequest) Validate() error {
	return dara.Validate(s)
}
