// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceMessageSubscribeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryMqttTraceMessageSubscribeResponseBody
	GetCurrentPage() *int32
	SetMessageTraceLists(v []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) *QueryMqttTraceMessageSubscribeResponseBody
	GetMessageTraceLists() []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists
	SetPageSize(v int32) *QueryMqttTraceMessageSubscribeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryMqttTraceMessageSubscribeResponseBody
	GetRequestId() *string
	SetTotal(v int64) *QueryMqttTraceMessageSubscribeResponseBody
	GetTotal() *int64
}

type QueryMqttTraceMessageSubscribeResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The delivery trace of the queried message.
	MessageTraceLists []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists `json:"MessageTraceLists,omitempty" xml:"MessageTraceLists,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique ID that the system generates for the request. You can use the ID to troubleshoot issues. This parameter is a common parameter.
	//
	// example:
	//
	// 4E685844-ADAF-4D85-9EAC-F9471E8C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned delivery traces.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMqttTraceMessageSubscribeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) GetMessageTraceLists() []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	return s.MessageTraceLists
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetCurrentPage(v int32) *QueryMqttTraceMessageSubscribeResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetMessageTraceLists(v []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) *QueryMqttTraceMessageSubscribeResponseBody {
	s.MessageTraceLists = v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetPageSize(v int32) *QueryMqttTraceMessageSubscribeResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetRequestId(v string) *QueryMqttTraceMessageSubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetTotal(v int64) *QueryMqttTraceMessageSubscribeResponseBody {
	s.Total = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) Validate() error {
	if s.MessageTraceLists != nil {
		for _, item := range s.MessageTraceLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists struct {
	// The action on the message. Valid values:
	//
	// - **sub**: The ApsaraMQ for MQTT client subscribes to the message.
	//
	// - **push_offline**: The ApsaraMQ for MQTT broker pushes the offline message to the ApsaraMQ for MQTT client.
	//
	// example:
	//
	// sub
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The code returned for the action on the message. Valid values:
	//
	// - **mqtt.trace.action.msg.sub**: The value that is returned if the value of Action is **sub**.
	//
	// - **mqtt.trace.action.msg.push.offline**: The value that is returned if the value of Action is **push_offline**.
	//
	// example:
	//
	// mqtt.trace.action.msg.sub
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The returned information for the action on the message. Valid values:
	//
	// - **Push To Mqtt Client**: The value that is returned if the value of Action is **sub**.
	//
	// - **Push Offline Msg To Mqtt Client**: The value that is returned if the value of Action is **push_offline**.
	//
	// example:
	//
	// Push To Mqtt Client
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The ID of the client that subscribes to the message.
	//
	// example:
	//
	// GID_test@@@consumer
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The message ID.
	//
	// example:
	//
	// AC1EC1B33D5978308DB17F3245E4****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The time when the message was delivered.
	//
	// example:
	//
	// 2021-05-25 16:46:41.274
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) GetAction() *string {
	return s.Action
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) GetActionCode() *string {
	return s.ActionCode
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) GetActionInfo() *string {
	return s.ActionInfo
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) GetClientId() *string {
	return s.ClientId
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) GetMsgId() *string {
	return s.MsgId
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) GetTime() *string {
	return s.Time
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetAction(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.Action = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetActionCode(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.ActionCode = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetActionInfo(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.ActionInfo = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetClientId(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetMsgId(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetTime(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.Time = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) Validate() error {
	return dara.Validate(s)
}
