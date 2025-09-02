// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityFollowerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMode(v int32) *CreateQualityFollowerRequest
	GetAlarmMode() *int32
	SetEntityId(v int64) *CreateQualityFollowerRequest
	GetEntityId() *int64
	SetFollower(v string) *CreateQualityFollowerRequest
	GetFollower() *string
	SetProjectId(v int64) *CreateQualityFollowerRequest
	GetProjectId() *int64
	SetProjectName(v string) *CreateQualityFollowerRequest
	GetProjectName() *string
}

type CreateQualityFollowerRequest struct {
	// The notification method. Valid values: 1, 2, 4, 5, 6, 7, 8, and 9. The value 1 indicates that the notification is sent by email. The value 2 indicates that the notification is sent by email and text message. The value 4 indicates that the notification is sent by a DingTalk chatbot. The value 5 indicates that the notification is sent by a DingTalk chatbot to all members in a DingTalk group. The value 6 indicates that the notification is sent by Lark. The value 7 indicates that the notification is sent by WeCom. The value 8 indicates that the notification is sent by webhook. The value 9 indicates that the notification is sent by phone call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AlarmMode *int32 `json:"AlarmMode,omitempty" xml:"AlarmMode,omitempty"`
	// The ID of the partition filter expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The account ID of the subscriber.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Follower *string `json:"Follower,omitempty" xml:"Follower,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console to query the ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine or data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateQualityFollowerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityFollowerRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityFollowerRequest) GetAlarmMode() *int32 {
	return s.AlarmMode
}

func (s *CreateQualityFollowerRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *CreateQualityFollowerRequest) GetFollower() *string {
	return s.Follower
}

func (s *CreateQualityFollowerRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateQualityFollowerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateQualityFollowerRequest) SetAlarmMode(v int32) *CreateQualityFollowerRequest {
	s.AlarmMode = &v
	return s
}

func (s *CreateQualityFollowerRequest) SetEntityId(v int64) *CreateQualityFollowerRequest {
	s.EntityId = &v
	return s
}

func (s *CreateQualityFollowerRequest) SetFollower(v string) *CreateQualityFollowerRequest {
	s.Follower = &v
	return s
}

func (s *CreateQualityFollowerRequest) SetProjectId(v int64) *CreateQualityFollowerRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateQualityFollowerRequest) SetProjectName(v string) *CreateQualityFollowerRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateQualityFollowerRequest) Validate() error {
	return dara.Validate(s)
}
