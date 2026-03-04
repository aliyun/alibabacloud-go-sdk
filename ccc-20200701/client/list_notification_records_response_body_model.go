// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNotificationRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNotificationRecordsResponseBody
	GetCode() *string
	SetData(v []*ListNotificationRecordsResponseBodyData) *ListNotificationRecordsResponseBody
	GetData() []*ListNotificationRecordsResponseBodyData
	SetHttpStatusCode(v int32) *ListNotificationRecordsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListNotificationRecordsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListNotificationRecordsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListNotificationRecordsResponseBody
	GetRequestId() *string
}

type ListNotificationRecordsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListNotificationRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 0630E5DF-CEB0-445B-8626-D5C7481181C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNotificationRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNotificationRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNotificationRecordsResponseBody) GetData() []*ListNotificationRecordsResponseBodyData {
	return s.Data
}

func (s *ListNotificationRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNotificationRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNotificationRecordsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListNotificationRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNotificationRecordsResponseBody) SetCode(v string) *ListNotificationRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListNotificationRecordsResponseBody) SetData(v []*ListNotificationRecordsResponseBodyData) *ListNotificationRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListNotificationRecordsResponseBody) SetHttpStatusCode(v int32) *ListNotificationRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNotificationRecordsResponseBody) SetMessage(v string) *ListNotificationRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListNotificationRecordsResponseBody) SetParams(v []*string) *ListNotificationRecordsResponseBody {
	s.Params = v
	return s
}

func (s *ListNotificationRecordsResponseBody) SetRequestId(v string) *ListNotificationRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNotificationRecordsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNotificationRecordsResponseBodyData struct {
	// example:
	//
	// {
	//
	// 	"agentId": "agent@ccc-test",
	//
	// 	"callType": "OUTBOUND",
	//
	// 	"callee": "13****00",
	//
	// 	"caller": "05****81",
	//
	// 	"channelId": "ch-user-13****00-05****81-1772619731285-job-*****",
	//
	// 	"contactId": "job-*****",
	//
	// 	"eventTime": "2026-03-04T10:22:11.309Z",
	//
	// 	"eventType": "Dialing",
	//
	// 	"instanceId": "ccc-test",
	//
	// 	"mediaType": "AUDIO",
	//
	// 	"skillGroupId": "skill@ccc-test"
	//
	// }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-468a63a2-****-****-****-b1ecf726d4be
	NotificationKey *string `json:"NotificationKey,omitempty" xml:"NotificationKey,omitempty"`
	// example:
	//
	// Dialing
	NotificationType *string `json:"NotificationType,omitempty" xml:"NotificationType,omitempty"`
}

func (s ListNotificationRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNotificationRecordsResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListNotificationRecordsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNotificationRecordsResponseBodyData) GetNotificationKey() *string {
	return s.NotificationKey
}

func (s *ListNotificationRecordsResponseBodyData) GetNotificationType() *string {
	return s.NotificationType
}

func (s *ListNotificationRecordsResponseBodyData) SetContent(v string) *ListNotificationRecordsResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListNotificationRecordsResponseBodyData) SetInstanceId(v string) *ListNotificationRecordsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListNotificationRecordsResponseBodyData) SetNotificationKey(v string) *ListNotificationRecordsResponseBodyData {
	s.NotificationKey = &v
	return s
}

func (s *ListNotificationRecordsResponseBodyData) SetNotificationType(v string) *ListNotificationRecordsResponseBodyData {
	s.NotificationType = &v
	return s
}

func (s *ListNotificationRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
