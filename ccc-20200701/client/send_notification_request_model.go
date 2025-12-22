// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SendNotificationRequest
	GetInstanceId() *string
	SetMessageBody(v string) *SendNotificationRequest
	GetMessageBody() *string
	SetNotificationTarget(v string) *SendNotificationRequest
	GetNotificationTarget() *string
	SetNotificationType(v string) *SendNotificationRequest
	GetNotificationType() *string
	SetShardingKey(v string) *SendNotificationRequest
	GetShardingKey() *string
}

type SendNotificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {   "eventType": "VnRecordingReady",   "eventTime": "2025-11-05T01:51:20.974Z",    "instanceId": "ccc-test",                    "scenarioId": "scenario-xxxx",               "contactId": "job-xxxxxx",                   "duration": 59,                              "startTime": "2025-11-05T01:50:04.732Z",    "endTime": "2025-11-05T01:51:02.822Z",      "fileName": "xxxxxxx.wav",                   "downloadURL": "https://ccc-v2-online.oss-cn-shanghai.aliyuncs.com/ccc-record-mixed/xxxxxxx.wav?Expires=1762393880" }
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOMER
	NotificationTarget *string `json:"NotificationTarget,omitempty" xml:"NotificationTarget,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VnRecordingReady
	NotificationType *string `json:"NotificationType,omitempty" xml:"NotificationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-xxxxxx
	ShardingKey *string `json:"ShardingKey,omitempty" xml:"ShardingKey,omitempty"`
}

func (s SendNotificationRequest) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendNotificationRequest) GetMessageBody() *string {
	return s.MessageBody
}

func (s *SendNotificationRequest) GetNotificationTarget() *string {
	return s.NotificationTarget
}

func (s *SendNotificationRequest) GetNotificationType() *string {
	return s.NotificationType
}

func (s *SendNotificationRequest) GetShardingKey() *string {
	return s.ShardingKey
}

func (s *SendNotificationRequest) SetInstanceId(v string) *SendNotificationRequest {
	s.InstanceId = &v
	return s
}

func (s *SendNotificationRequest) SetMessageBody(v string) *SendNotificationRequest {
	s.MessageBody = &v
	return s
}

func (s *SendNotificationRequest) SetNotificationTarget(v string) *SendNotificationRequest {
	s.NotificationTarget = &v
	return s
}

func (s *SendNotificationRequest) SetNotificationType(v string) *SendNotificationRequest {
	s.NotificationType = &v
	return s
}

func (s *SendNotificationRequest) SetShardingKey(v string) *SendNotificationRequest {
	s.ShardingKey = &v
	return s
}

func (s *SendNotificationRequest) Validate() error {
	return dara.Validate(s)
}
