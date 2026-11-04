// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryDeviceEventRequest
	GetBeginTime() *int64
	SetCurrentPage(v int32) *QueryDeviceEventRequest
	GetCurrentPage() *int32
	SetDeviceName(v string) *QueryDeviceEventRequest
	GetDeviceName() *string
	SetEndTime(v int64) *QueryDeviceEventRequest
	GetEndTime() *int64
	SetEventType(v int32) *QueryDeviceEventRequest
	GetEventType() *int32
	SetIotId(v string) *QueryDeviceEventRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryDeviceEventRequest
	GetIotInstanceId() *string
	SetPageSize(v int32) *QueryDeviceEventRequest
	GetPageSize() *int32
	SetProductKey(v string) *QueryDeviceEventRequest
	GetProductKey() *string
}

type QueryDeviceEventRequest struct {
	// The start time of the query range. The time is a 13-digit timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1539507310000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The page number of the returned page. Pages start from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the IPC device. You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// camera1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The end time of the query range. The time is a 13-digit timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1539607310000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the intelligent alert event. Valid values:
	//
	// - 0: all intelligent alert events.
	//
	// - A specific sequence number: a single intelligent alert event, such as 10001.
	//
	// You can view all intelligent alert event types on the **Define Feature*	- tab of the **Product Details*	- page in the IoT Platform console. For more information, see [View intelligent alert event types](https://help.aliyun.com/document_detail/207701.html).
	//
	// example:
	//
	// 0
	EventType *int32 `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the IPC device. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- uniquely identifies the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// zLZyi6aOLyOSHa9hsPyD00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// <props="china">You can view the **ID*	- of the current instance on the **Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify the ID. Otherwise, the call fails.
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
	// The number of entries to return on each page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ProductKey of the product to which the IPC device belongs. You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryDeviceEventRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceEventRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryDeviceEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryDeviceEventRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryDeviceEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryDeviceEventRequest) GetEventType() *int32 {
	return s.EventType
}

func (s *QueryDeviceEventRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryDeviceEventRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryDeviceEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDeviceEventRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryDeviceEventRequest) SetBeginTime(v int64) *QueryDeviceEventRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryDeviceEventRequest) SetCurrentPage(v int32) *QueryDeviceEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryDeviceEventRequest) SetDeviceName(v string) *QueryDeviceEventRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceEventRequest) SetEndTime(v int64) *QueryDeviceEventRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDeviceEventRequest) SetEventType(v int32) *QueryDeviceEventRequest {
	s.EventType = &v
	return s
}

func (s *QueryDeviceEventRequest) SetIotId(v string) *QueryDeviceEventRequest {
	s.IotId = &v
	return s
}

func (s *QueryDeviceEventRequest) SetIotInstanceId(v string) *QueryDeviceEventRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryDeviceEventRequest) SetPageSize(v int32) *QueryDeviceEventRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDeviceEventRequest) SetProductKey(v string) *QueryDeviceEventRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryDeviceEventRequest) Validate() error {
	return dara.Validate(s)
}
