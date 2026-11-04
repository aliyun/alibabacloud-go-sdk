// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordUrlByTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int32) *QueryRecordUrlByTimeRequest
	GetBeginTime() *int32
	SetDeviceName(v string) *QueryRecordUrlByTimeRequest
	GetDeviceName() *string
	SetEndTime(v int32) *QueryRecordUrlByTimeRequest
	GetEndTime() *int32
	SetIotId(v string) *QueryRecordUrlByTimeRequest
	GetIotId() *string
	SetIotInstanceId(v string) *QueryRecordUrlByTimeRequest
	GetIotInstanceId() *string
	SetProductKey(v string) *QueryRecordUrlByTimeRequest
	GetProductKey() *string
	SetStreamType(v int32) *QueryRecordUrlByTimeRequest
	GetStreamType() *int32
	SetUrlValidDuration(v int32) *QueryRecordUrlByTimeRequest
	GetUrlValidDuration() *int32
}

type QueryRecordUrlByTimeRequest struct {
	// The start time of the query range. The value is a 10-digit timestamp in seconds.
	//
	// The minimum value is 0, and the maximum value is 2147483647.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638675381
	BeginTime *int32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
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
	// The end time of the query range. The value is a 10-digit timestamp in seconds.
	//
	// The minimum value is 0, and the maximum value is 2147483647.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638866820
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IotId of the IP camera. The IotId is the unique identifier issued by IoT Platform for the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the IotId.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- is the unique identifier of the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// d7XmBoJhAr88C6PelXDF00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page of the IoT Platform console. If the instance has an ID, you must specify this ID. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot-060a****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The ProductKey of the product to which the IP camera belongs.
	//
	// You can view the ProductKey on the **Products*	- page of the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1BwAGV****
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	// The type of the video stream. Valid values:
	//
	// - **0*	- (default): main stream
	//
	// - **1**: sub stream
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The validity period of the URL. Valid values: 10 to 3600. Unit: seconds.
	//
	// example:
	//
	// 200
	UrlValidDuration *int32 `json:"UrlValidDuration,omitempty" xml:"UrlValidDuration,omitempty"`
}

func (s QueryRecordUrlByTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordUrlByTimeRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordUrlByTimeRequest) GetBeginTime() *int32 {
	return s.BeginTime
}

func (s *QueryRecordUrlByTimeRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryRecordUrlByTimeRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *QueryRecordUrlByTimeRequest) GetIotId() *string {
	return s.IotId
}

func (s *QueryRecordUrlByTimeRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryRecordUrlByTimeRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryRecordUrlByTimeRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *QueryRecordUrlByTimeRequest) GetUrlValidDuration() *int32 {
	return s.UrlValidDuration
}

func (s *QueryRecordUrlByTimeRequest) SetBeginTime(v int32) *QueryRecordUrlByTimeRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetDeviceName(v string) *QueryRecordUrlByTimeRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetEndTime(v int32) *QueryRecordUrlByTimeRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetIotId(v string) *QueryRecordUrlByTimeRequest {
	s.IotId = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetIotInstanceId(v string) *QueryRecordUrlByTimeRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetProductKey(v string) *QueryRecordUrlByTimeRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetStreamType(v int32) *QueryRecordUrlByTimeRequest {
	s.StreamType = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) SetUrlValidDuration(v int32) *QueryRecordUrlByTimeRequest {
	s.UrlValidDuration = &v
	return s
}

func (s *QueryRecordUrlByTimeRequest) Validate() error {
	return dara.Validate(s)
}
