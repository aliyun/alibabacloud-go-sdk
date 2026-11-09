// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryMqttTraceDeviceResponseBody
	GetCurrentPage() *int32
	SetDeviceInfoList(v []*QueryMqttTraceDeviceResponseBodyDeviceInfoList) *QueryMqttTraceDeviceResponseBody
	GetDeviceInfoList() []*QueryMqttTraceDeviceResponseBodyDeviceInfoList
	SetPageSize(v int32) *QueryMqttTraceDeviceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryMqttTraceDeviceResponseBody
	GetRequestId() *string
	SetTotal(v int64) *QueryMqttTraceDeviceResponseBody
	GetTotal() *int64
}

type QueryMqttTraceDeviceResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The details of the device activity entries.
	DeviceInfoList []*QueryMqttTraceDeviceResponseBodyDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Repeated"`
	// The maximum number of device activity entries displayed per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A common parameter. Each request has a unique ID that is active for troubleshooting and locating issues.
	//
	// example:
	//
	// 317076B7-F946-46BC-A98F-4CF9777C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of device activity entries returned.
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMqttTraceDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMqttTraceDeviceResponseBody) GetDeviceInfoList() []*QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	return s.DeviceInfoList
}

func (s *QueryMqttTraceDeviceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMqttTraceDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMqttTraceDeviceResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *QueryMqttTraceDeviceResponseBody) SetCurrentPage(v int32) *QueryMqttTraceDeviceResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetDeviceInfoList(v []*QueryMqttTraceDeviceResponseBodyDeviceInfoList) *QueryMqttTraceDeviceResponseBody {
	s.DeviceInfoList = v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetPageSize(v int32) *QueryMqttTraceDeviceResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetRequestId(v string) *QueryMqttTraceDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetTotal(v int64) *QueryMqttTraceDeviceResponseBody {
	s.Total = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) Validate() error {
	if s.DeviceInfoList != nil {
		for _, item := range s.DeviceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMqttTraceDeviceResponseBodyDeviceInfoList struct {
	// The device action. Valid values:
	//
	// - **connect**: The client requests a connection to the server.
	//
	// - **close**: The TCP connection is closed.
	//
	// - **disconnect**: The client requests to disconnect from the server.
	//
	// example:
	//
	// connect
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The return code of the device action. Valid values:
	//
	// - **mqtt.trace.action.connect**: returned when the device action is **connect**.
	//
	// - **mqtt.trace.action.close**: returned when the device action is **close**.
	//
	// - **mqtt.trace.action.disconnect**: returned when the device action is **disconnect**.
	//
	// example:
	//
	// mqtt.trace.action.connect
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The description of the device action. Valid values:
	//
	// - **accepted**: The server accepts the connection request from the client.
	//
	// - **not authorized**: The TCP connection is closed because the permission verification of the instance accessed by the client fails.
	//
	// - **clientId conflict**: The TCP connection is closed because of a client ID conflict.
	//
	// - **resource auth failed**: The TCP connection is closed because the permission verification of the topic or group ID resource accessed by the client fails.
	//
	// - **no heart**: The TCP connection is closed because no heartbeat is detected from the client.
	//
	// - **closed by client**: The TCP connection is closed because of a client exception.
	//
	// - **disconnected by client**: The client requests to disconnect.
	//
	// - **invalid param**: The TCP connection is closed because of invalid request parameters.
	//
	// - **Socket IOException**: The TCP connection is closed because of network jitter or packet loss.
	//
	// example:
	//
	// accept
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The connection identifier of the device.
	//
	// example:
	//
	// c69fe839209547fa9d073781b9cd****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The time when the device action occurred.
	//
	// example:
	//
	// 2021-05-21 15:51:54.867
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryMqttTraceDeviceResponseBodyDeviceInfoList) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceDeviceResponseBodyDeviceInfoList) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) GetAction() *string {
	return s.Action
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) GetActionCode() *string {
	return s.ActionCode
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) GetActionInfo() *string {
	return s.ActionInfo
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) GetTime() *string {
	return s.Time
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetAction(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.Action = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetActionCode(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.ActionCode = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetActionInfo(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.ActionInfo = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetChannelId(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.ChannelId = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetTime(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.Time = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) Validate() error {
	return dara.Validate(s)
}
