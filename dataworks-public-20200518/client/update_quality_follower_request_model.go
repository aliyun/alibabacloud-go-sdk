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
	// The notification method. Valid values: 1 (email), 2 (email and SMS), 4 (DingTalk group chatbot), 5 (DingTalk group chatbot with an @all reminder), 6 (Lark), 7 (WeCom), 8 (webhook), and 9 (phone call).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	AlarmMode *int32 `json:"AlarmMode,omitempty" xml:"AlarmMode,omitempty"`
	// The subscriber.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1912232****
	Follower *string `json:"Follower,omitempty" xml:"Follower,omitempty"`
	// The ID of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	FollowerId *int64 `json:"FollowerId,omitempty" xml:"FollowerId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the DataWorks console to obtain the ID.
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
