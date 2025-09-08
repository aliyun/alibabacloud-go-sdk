// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventDisposeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEventDisposeResponseBody
	GetCode() *int32
	SetData(v *DescribeEventDisposeResponseBodyData) *DescribeEventDisposeResponseBody
	GetData() *DescribeEventDisposeResponseBodyData
	SetMessage(v string) *DescribeEventDisposeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEventDisposeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEventDisposeResponseBody
	GetSuccess() *bool
}

type DescribeEventDisposeResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeEventDisposeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEventDisposeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDisposeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEventDisposeResponseBody) GetData() *DescribeEventDisposeResponseBodyData {
	return s.Data
}

func (s *DescribeEventDisposeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventDisposeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventDisposeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventDisposeResponseBody) SetCode(v int32) *DescribeEventDisposeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetData(v *DescribeEventDisposeResponseBodyData) *DescribeEventDisposeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetMessage(v string) *DescribeEventDisposeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetRequestId(v string) *DescribeEventDisposeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetSuccess(v bool) *DescribeEventDisposeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventDisposeResponseBodyData struct {
	// An array consisting of JSON objects that are configured for event handling.
	//
	// example:
	//
	// { playbookName: "使用安全组封禁入方向IP", sophonTaskId: "400442a5-4f98-45ed-97db-5ab117eb0b8f", … }
	EventDispose []interface{} `json:"EventDispose,omitempty" xml:"EventDispose,omitempty" type:"Repeated"`
	// The JSON object that is configured for an alert recipient.
	ReceiverInfo *DescribeEventDisposeResponseBodyDataReceiverInfo `json:"ReceiverInfo,omitempty" xml:"ReceiverInfo,omitempty" type:"Struct"`
	// The description of the event.
	//
	// example:
	//
	// dealed
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: not handled
	//
	// 	- 1: handing
	//
	// 	- 5: handling failed
	//
	// 	- 10: handled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventDisposeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDisposeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponseBodyData) GetEventDispose() []interface{} {
	return s.EventDispose
}

func (s *DescribeEventDisposeResponseBodyData) GetReceiverInfo() *DescribeEventDisposeResponseBodyDataReceiverInfo {
	return s.ReceiverInfo
}

func (s *DescribeEventDisposeResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *DescribeEventDisposeResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeEventDisposeResponseBodyData) SetEventDispose(v []interface{}) *DescribeEventDisposeResponseBodyData {
	s.EventDispose = v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) SetReceiverInfo(v *DescribeEventDisposeResponseBodyDataReceiverInfo) *DescribeEventDisposeResponseBodyData {
	s.ReceiverInfo = v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) SetRemark(v string) *DescribeEventDisposeResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) SetStatus(v int32) *DescribeEventDisposeResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeEventDisposeResponseBodyDataReceiverInfo struct {
	// The channel of the contact information. Valid values:
	//
	// 	- message
	//
	// 	- mail
	//
	// example:
	//
	// message
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the recipient who receives the event handling result.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The message title.
	//
	// example:
	//
	// siem event dealed message
	MessageTitle *string `json:"MessageTitle,omitempty" xml:"MessageTitle,omitempty"`
	// The contact information of the recipient.
	//
	// example:
	//
	// 138xxxxxx
	Receiver *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
	// Indicates whether the message is sent. Valid values:
	//
	// 	- 0: not sent
	//
	// 	- 1: sent
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventDisposeResponseBodyDataReceiverInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDisposeResponseBodyDataReceiverInfo) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) GetChannel() *string {
	return s.Channel
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) GetMessageTitle() *string {
	return s.MessageTitle
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) GetReceiver() *string {
	return s.Receiver
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetChannel(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Channel = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetGmtCreate(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetGmtModified(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.GmtModified = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetId(v int64) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Id = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetIncidentUuid(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetMessageTitle(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.MessageTitle = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetReceiver(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Receiver = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetStatus(v int32) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Status = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) Validate() error {
	return dara.Validate(s)
}
