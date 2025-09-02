// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityFollowerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMode(v int32) *UpdateQualityFollowerRequest
	GetAlarmMode() *int32
	SetFollower(v string) *UpdateQualityFollowerRequest
	GetFollower() *string
	SetFollowerId(v int64) *UpdateQualityFollowerRequest
	GetFollowerId() *int64
	SetProjectId(v int64) *UpdateQualityFollowerRequest
	GetProjectId() *int64
	SetProjectName(v string) *UpdateQualityFollowerRequest
	GetProjectName() *string
}

type UpdateQualityFollowerRequest struct {
	// The notification method. Valid values: 1, 2, 4, 5, 6, 7, 8, and 9. The value 1 indicates that the notification is sent by email. The value 2 indicates that the notification is sent by email and text message. The value 4 indicates that the notification is sent by a DingTalk chatbot. The value 5 indicates that the notification is sent by a DingTalk chatbot to all members in a DingTalk group. The value 6 indicates that the notification is sent by Lark. The value 7 indicates that the notification is sent by WeCom. The value 8 indicates that the notification is sent by webhook. The value 9 indicates that the notification is sent by phone call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	AlarmMode *int32 `json:"AlarmMode,omitempty" xml:"AlarmMode,omitempty"`
	// The name of the subscriber.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1912232****
	Follower *string `json:"Follower,omitempty" xml:"Follower,omitempty"`
	// The ID of the subscription relationship.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	FollowerId *int64 `json:"FollowerId,omitempty" xml:"FollowerId,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console to obtain the workspace ID.
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

func (s UpdateQualityFollowerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityFollowerRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityFollowerRequest) GetAlarmMode() *int32 {
	return s.AlarmMode
}

func (s *UpdateQualityFollowerRequest) GetFollower() *string {
	return s.Follower
}

func (s *UpdateQualityFollowerRequest) GetFollowerId() *int64 {
	return s.FollowerId
}

func (s *UpdateQualityFollowerRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateQualityFollowerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateQualityFollowerRequest) SetAlarmMode(v int32) *UpdateQualityFollowerRequest {
	s.AlarmMode = &v
	return s
}

func (s *UpdateQualityFollowerRequest) SetFollower(v string) *UpdateQualityFollowerRequest {
	s.Follower = &v
	return s
}

func (s *UpdateQualityFollowerRequest) SetFollowerId(v int64) *UpdateQualityFollowerRequest {
	s.FollowerId = &v
	return s
}

func (s *UpdateQualityFollowerRequest) SetProjectId(v int64) *UpdateQualityFollowerRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateQualityFollowerRequest) SetProjectName(v string) *UpdateQualityFollowerRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateQualityFollowerRequest) Validate() error {
	return dara.Validate(s)
}
