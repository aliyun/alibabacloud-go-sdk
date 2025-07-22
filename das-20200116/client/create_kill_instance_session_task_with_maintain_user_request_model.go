// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKillInstanceSessionTaskWithMaintainUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoredUsers(v string) *CreateKillInstanceSessionTaskWithMaintainUserRequest
	GetIgnoredUsers() *string
	SetInstanceId(v string) *CreateKillInstanceSessionTaskWithMaintainUserRequest
	GetInstanceId() *string
	SetKillAllSessions(v bool) *CreateKillInstanceSessionTaskWithMaintainUserRequest
	GetKillAllSessions() *bool
	SetNodeId(v string) *CreateKillInstanceSessionTaskWithMaintainUserRequest
	GetNodeId() *string
	SetSessionIds(v string) *CreateKillInstanceSessionTaskWithMaintainUserRequest
	GetSessionIds() *string
}

type CreateKillInstanceSessionTaskWithMaintainUserRequest struct {
	IgnoredUsers *string `json:"IgnoredUsers,omitempty" xml:"IgnoredUsers,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	KillAllSessions *bool   `json:"KillAllSessions,omitempty" xml:"KillAllSessions,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	SessionIds      *string `json:"SessionIds,omitempty" xml:"SessionIds,omitempty"`
}

func (s CreateKillInstanceSessionTaskWithMaintainUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKillInstanceSessionTaskWithMaintainUserRequest) GoString() string {
	return s.String()
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) GetIgnoredUsers() *string {
	return s.IgnoredUsers
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) GetKillAllSessions() *bool {
	return s.KillAllSessions
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) GetSessionIds() *string {
	return s.SessionIds
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) SetIgnoredUsers(v string) *CreateKillInstanceSessionTaskWithMaintainUserRequest {
	s.IgnoredUsers = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) SetInstanceId(v string) *CreateKillInstanceSessionTaskWithMaintainUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) SetKillAllSessions(v bool) *CreateKillInstanceSessionTaskWithMaintainUserRequest {
	s.KillAllSessions = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) SetNodeId(v string) *CreateKillInstanceSessionTaskWithMaintainUserRequest {
	s.NodeId = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) SetSessionIds(v string) *CreateKillInstanceSessionTaskWithMaintainUserRequest {
	s.SessionIds = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserRequest) Validate() error {
	return dara.Validate(s)
}
