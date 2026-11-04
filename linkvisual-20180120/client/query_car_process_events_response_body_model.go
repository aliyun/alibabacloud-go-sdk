// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCarProcessEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCarProcessEventsResponseBody
	GetCode() *string
	SetData(v *QueryCarProcessEventsResponseBodyData) *QueryCarProcessEventsResponseBody
	GetData() *QueryCarProcessEventsResponseBodyData
	SetErrorMessage(v string) *QueryCarProcessEventsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryCarProcessEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCarProcessEventsResponseBody
	GetSuccess() *bool
}

type QueryCarProcessEventsResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: An error occurred. For more information, see [Error codes](https://help.aliyun.com/document_detail/288314.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned if the call is successful.
	Data *QueryCarProcessEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Specified access key is not found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8B8E77F0-7713-0139-A3D6-93482A025D45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCarProcessEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCarProcessEventsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCarProcessEventsResponseBody) GetData() *QueryCarProcessEventsResponseBodyData {
	return s.Data
}

func (s *QueryCarProcessEventsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryCarProcessEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCarProcessEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCarProcessEventsResponseBody) SetCode(v string) *QueryCarProcessEventsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCarProcessEventsResponseBody) SetData(v *QueryCarProcessEventsResponseBodyData) *QueryCarProcessEventsResponseBody {
	s.Data = v
	return s
}

func (s *QueryCarProcessEventsResponseBody) SetErrorMessage(v string) *QueryCarProcessEventsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryCarProcessEventsResponseBody) SetRequestId(v string) *QueryCarProcessEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBody) SetSuccess(v bool) *QueryCarProcessEventsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCarProcessEventsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCarProcessEventsResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The list of returned intelligent alert events.
	PageData []*QueryCarProcessEventsResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of vehicle recognition events.
	//
	// example:
	//
	// 50
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryCarProcessEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCarProcessEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryCarProcessEventsResponseBodyData) GetPageCount() *int32 {
	return s.PageCount
}

func (s *QueryCarProcessEventsResponseBodyData) GetPageData() []*QueryCarProcessEventsResponseBodyDataPageData {
	return s.PageData
}

func (s *QueryCarProcessEventsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCarProcessEventsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryCarProcessEventsResponseBodyData) SetCurrentPage(v int32) *QueryCarProcessEventsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyData) SetPageCount(v int32) *QueryCarProcessEventsResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyData) SetPageData(v []*QueryCarProcessEventsResponseBodyDataPageData) *QueryCarProcessEventsResponseBodyData {
	s.PageData = v
	return s
}

func (s *QueryCarProcessEventsResponseBodyData) SetPageSize(v int32) *QueryCarProcessEventsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyData) SetTotal(v int32) *QueryCarProcessEventsResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCarProcessEventsResponseBodyDataPageData struct {
	// The action type of the vehicle.
	//
	// example:
	//
	// 1
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The number of detection boxes recognized by the IPC device.
	//
	// example:
	//
	// 0
	AreaIndex *int32 `json:"AreaIndex,omitempty" xml:"AreaIndex,omitempty"`
	// The confidence level of the vehicle recognition.
	//
	// example:
	//
	// 63
	Confidence *int32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The ID of the intelligent alert event.
	//
	// example:
	//
	// dailyb3d6658d9****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the intelligent alert event picture.
	//
	// example:
	//
	// SU5LT2pXaUM4MV****Nzg
	EventPicId *string `json:"EventPicId,omitempty" xml:"EventPicId,omitempty"`
	// The URL of the intelligent alert event picture.
	//
	// example:
	//
	// https://example.com/test.jpg
	EventPicUrl *string `json:"EventPicUrl,omitempty" xml:"EventPicUrl,omitempty"`
	// The timestamp when the event occurred, in milliseconds.
	//
	// example:
	//
	// 1646288428361
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The type of the intelligent alert event. Valid values:
	//
	// - **0**: All intelligent alert events.
	//
	// - Multiple intelligent alert events: Enter the sequence numbers of the intelligent alert events and separate them with commas (,), such as `10001,10002`.
	//
	// > You can specify up to 10 intelligent alert events at a time.
	//
	// To view all intelligent alert event types, go to the **Define Feature*	- tab of the **Product Details*	- page in the IoT Platform console. For more information, see [View intelligent alert event types](https://help.aliyun.com/document_detail/207701.html).
	//
	// example:
	//
	// 10005
	EventType *int32 `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The IotId of the V-series edge appliance.
	//
	// example:
	//
	// gl9kLgp0UsPSCcIM8EAY00****
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// The license plate number.
	//
	// example:
	//
	// ZheA6****
	PlateNo *string `json:"PlateNo,omitempty" xml:"PlateNo,omitempty"`
	// The DeviceName of the IPC device.
	//
	// example:
	//
	// car_test
	SubDeviceName *string `json:"SubDeviceName,omitempty" xml:"SubDeviceName,omitempty"`
	// The DeviceNickName of the IPC device.
	//
	// example:
	//
	// TestVehicle
	SubDeviceNickName *string `json:"SubDeviceNickName,omitempty" xml:"SubDeviceNickName,omitempty"`
	// The ID of the IPC device.
	//
	// example:
	//
	// ImP6ECd1aQxio9CAfWTO****
	SubIotId *string `json:"SubIotId,omitempty" xml:"SubIotId,omitempty"`
	// The ProductKey of the IPC device.
	//
	// example:
	//
	// a1Duh0F****
	SubProductKey *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
	// The task ID of the vehicle detection event.
	//
	// example:
	//
	// daily55617e7653354490938d****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryCarProcessEventsResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s QueryCarProcessEventsResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetActionType() *int32 {
	return s.ActionType
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetAreaIndex() *int32 {
	return s.AreaIndex
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetConfidence() *int32 {
	return s.Confidence
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetEventId() *string {
	return s.EventId
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetEventPicId() *string {
	return s.EventPicId
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetEventPicUrl() *string {
	return s.EventPicUrl
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetEventTime() *int64 {
	return s.EventTime
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetEventType() *int32 {
	return s.EventType
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetIotId() *string {
	return s.IotId
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetPlateNo() *string {
	return s.PlateNo
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetSubDeviceName() *string {
	return s.SubDeviceName
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetSubDeviceNickName() *string {
	return s.SubDeviceNickName
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetSubIotId() *string {
	return s.SubIotId
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetSubProductKey() *string {
	return s.SubProductKey
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetActionType(v int32) *QueryCarProcessEventsResponseBodyDataPageData {
	s.ActionType = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetAreaIndex(v int32) *QueryCarProcessEventsResponseBodyDataPageData {
	s.AreaIndex = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetConfidence(v int32) *QueryCarProcessEventsResponseBodyDataPageData {
	s.Confidence = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventPicId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventPicId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventPicUrl(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventPicUrl = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventTime(v int64) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventTime = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetEventType(v int32) *QueryCarProcessEventsResponseBodyDataPageData {
	s.EventType = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetIotId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.IotId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetPlateNo(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.PlateNo = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetSubDeviceName(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.SubDeviceName = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetSubDeviceNickName(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.SubDeviceNickName = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetSubIotId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.SubIotId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetSubProductKey(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.SubProductKey = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) SetTaskId(v string) *QueryCarProcessEventsResponseBodyDataPageData {
	s.TaskId = &v
	return s
}

func (s *QueryCarProcessEventsResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
