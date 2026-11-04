// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCarProcessEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *QueryCarProcessEventsRequest
	GetActionType() *int32
	SetAreaIndex(v int32) *QueryCarProcessEventsRequest
	GetAreaIndex() *int32
	SetBeginTime(v int64) *QueryCarProcessEventsRequest
	GetBeginTime() *int64
	SetCurrentPage(v int32) *QueryCarProcessEventsRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *QueryCarProcessEventsRequest
	GetEndTime() *int64
	SetIotInstanceId(v string) *QueryCarProcessEventsRequest
	GetIotInstanceId() *string
	SetPageSize(v int32) *QueryCarProcessEventsRequest
	GetPageSize() *int32
	SetPlateNo(v string) *QueryCarProcessEventsRequest
	GetPlateNo() *string
	SetSubDeviceName(v string) *QueryCarProcessEventsRequest
	GetSubDeviceName() *string
	SetSubIotId(v string) *QueryCarProcessEventsRequest
	GetSubIotId() *string
	SetSubProductKey(v string) *QueryCarProcessEventsRequest
	GetSubProductKey() *string
}

type QueryCarProcessEventsRequest struct {
	// The action type of the vehicle. Valid values:
	//
	// - **1**: Enter.
	//
	// - **2**: Stay.
	//
	// - **3**: Leave.
	//
	// example:
	//
	// 2
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The detection box of the IPC device. You can call the [InvokeThingService](https://help.aliyun.com/document_detail/69584.html) operation to query the number of detection boxes. Set **identifier*	- to **QueryAlgoTaskDetails**.
	//
	// example:
	//
	// 1
	AreaIndex *int32 `json:"AreaIndex,omitempty" xml:"AreaIndex,omitempty"`
	// The start time of the query range. The time is a 13-digit timestamp in milliseconds.
	//
	// example:
	//
	// 1646288487368
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The page number of the returned page. Pages start from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query range. The time is a 13-digit timestamp in milliseconds.
	//
	// example:
	//
	// 1646288488368
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID. You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console.
	//
	// 	Notice:
	//
	// - If the instance has an ID, you must specify this ID. Otherwise, the call fails.
	//
	// - If the instance does not have an ID<props="intl"><ph> or the **Instance Overview*	- page is not displayed</ph>, you do not need to specify this parameter.
	//
	//
	//
	// For more information about instances, see [Instance overview](https://help.aliyun.com/document_detail/356505.html).
	//
	// example:
	//
	// iotx-oxsshare****
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 20. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The license plate number. The value can be up to 15 characters in length.
	//
	// example:
	//
	// ZheAM****
	PlateNo *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	// The name of the IPC device.
	//
	// You can view the **DeviceName*	- on the **Devices*	- page in the IoT Platform console.
	//
	// 	Notice:  If you specify this parameter, you must also specify **ProductKey**.
	//
	// example:
	//
	// car_test
	SubDeviceName *string `json:"SubDeviceName,omitempty" xml:"SubDeviceName,omitempty"`
	// The ID of the IP camera (IPC) device. The ID is issued by IoT Platform and uniquely identifies the device. You can call the [QueryDeviceDetail](https://help.aliyun.com/document_detail/69594.html) operation of IoT Platform to query the ID.
	//
	// 	Notice:  If you specify this parameter, you do not need to specify **ProductKey*	- or **DeviceName**. The **IotId*	- uniquely identifies the device and corresponds to the combination of **ProductKey*	- and **DeviceName**. If you specify both **IotId*	- and the combination of **ProductKey*	- and **DeviceName**, **IotId*	- takes precedence.
	//
	// example:
	//
	// ImP6ECd1aQ****CAfWTO000101
	SubIotId *string `json:"SubIotId,omitempty" xml:"SubIotId,omitempty"`
	// The ProductKey of the product to which the IPC device belongs.
	//
	// You can view the ProductKey on the **Products*	- page in the IoT Platform console or call the [QueryProductList](https://help.aliyun.com/document_detail/69271.html) operation to query the information about all products under the current account.
	//
	// 	Notice:  If you specify this parameter, you must also specify **DeviceName**.
	//
	// example:
	//
	// a1Duh0F****
	SubProductKey *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
}

func (s QueryCarProcessEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCarProcessEventsRequest) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *QueryCarProcessEventsRequest) GetAreaIndex() *int32 {
	return s.AreaIndex
}

func (s *QueryCarProcessEventsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryCarProcessEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryCarProcessEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryCarProcessEventsRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *QueryCarProcessEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCarProcessEventsRequest) GetPlateNo() *string {
	return s.PlateNo
}

func (s *QueryCarProcessEventsRequest) GetSubDeviceName() *string {
	return s.SubDeviceName
}

func (s *QueryCarProcessEventsRequest) GetSubIotId() *string {
	return s.SubIotId
}

func (s *QueryCarProcessEventsRequest) GetSubProductKey() *string {
	return s.SubProductKey
}

func (s *QueryCarProcessEventsRequest) SetActionType(v int32) *QueryCarProcessEventsRequest {
	s.ActionType = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetAreaIndex(v int32) *QueryCarProcessEventsRequest {
	s.AreaIndex = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetBeginTime(v int64) *QueryCarProcessEventsRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetCurrentPage(v int32) *QueryCarProcessEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetEndTime(v int64) *QueryCarProcessEventsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetIotInstanceId(v string) *QueryCarProcessEventsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetPageSize(v int32) *QueryCarProcessEventsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetPlateNo(v string) *QueryCarProcessEventsRequest {
	s.PlateNo = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetSubDeviceName(v string) *QueryCarProcessEventsRequest {
	s.SubDeviceName = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetSubIotId(v string) *QueryCarProcessEventsRequest {
	s.SubIotId = &v
	return s
}

func (s *QueryCarProcessEventsRequest) SetSubProductKey(v string) *QueryCarProcessEventsRequest {
	s.SubProductKey = &v
	return s
}

func (s *QueryCarProcessEventsRequest) Validate() error {
	return dara.Validate(s)
}
