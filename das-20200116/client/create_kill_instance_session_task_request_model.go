// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKillInstanceSessionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbUser(v string) *CreateKillInstanceSessionTaskRequest
	GetDbUser() *string
	SetDbUserPassword(v string) *CreateKillInstanceSessionTaskRequest
	GetDbUserPassword() *string
	SetIgnoredUsers(v string) *CreateKillInstanceSessionTaskRequest
	GetIgnoredUsers() *string
	SetInstanceId(v string) *CreateKillInstanceSessionTaskRequest
	GetInstanceId() *string
	SetKillAllSessions(v bool) *CreateKillInstanceSessionTaskRequest
	GetKillAllSessions() *bool
	SetNodeId(v string) *CreateKillInstanceSessionTaskRequest
	GetNodeId() *string
	SetSessionIds(v string) *CreateKillInstanceSessionTaskRequest
	GetSessionIds() *string
}

type CreateKillInstanceSessionTaskRequest struct {
	// The database account that has the permissions to terminate sessions.
	//
	// This parameter is required.
	//
	// example:
	//
	// testUser
	DbUser *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	// The password of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testPassword
	DbUserPassword *string `json:"DbUserPassword,omitempty" xml:"DbUserPassword,omitempty"`
	// The account whose sessions do not need to be terminated.
	//
	// >  Set this parameter to a JSON array. Separate database accounts with commas (,). Example: [\\"Database account 1\\",\\"Database account 2\\"].
	//
	// example:
	//
	// [\\"db_user1\\",\\"db_user2\\"]
	IgnoredUsers *string `json:"IgnoredUsers,omitempty" xml:"IgnoredUsers,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to terminate all sessions.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If you set this parameter to **true**, sessions of the accounts that are specified by **IgnoredUsers**, sessions of internal O\\&M accounts of Alibaba Cloud, and **Binlog Dump*	- sessions are not terminated.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	KillAllSessions *bool `json:"KillAllSessions,omitempty" xml:"KillAllSessions,omitempty"`
	// The node ID.
	//
	// >  This parameter must be specified if the database instance is a PolarDB for MySQL cluster. If you do not specify a node ID and set **KillAllSessions*	- to **true**, the system traverses all nodes in the PolarDB for MySQL cluster and terminates the active sessions on each node.
	//
	// example:
	//
	// pi-bp1v203xzzh0a****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The IDs of sessions that need to be terminated.
	//
	// >  Set this parameter to a JSON array. Separate session IDs with commas (,). Example: [\\"Session ID1\\",\\"Session ID2\\"]. If **KillAllSessions*	- is set to **true**, this parameter does not take effect.
	//
	// example:
	//
	// [10805639,10805623,10805645,10805553,10805566,10805616]
	SessionIds *string `json:"SessionIds,omitempty" xml:"SessionIds,omitempty"`
}

func (s CreateKillInstanceSessionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKillInstanceSessionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateKillInstanceSessionTaskRequest) GetDbUser() *string {
	return s.DbUser
}

func (s *CreateKillInstanceSessionTaskRequest) GetDbUserPassword() *string {
	return s.DbUserPassword
}

func (s *CreateKillInstanceSessionTaskRequest) GetIgnoredUsers() *string {
	return s.IgnoredUsers
}

func (s *CreateKillInstanceSessionTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateKillInstanceSessionTaskRequest) GetKillAllSessions() *bool {
	return s.KillAllSessions
}

func (s *CreateKillInstanceSessionTaskRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateKillInstanceSessionTaskRequest) GetSessionIds() *string {
	return s.SessionIds
}

func (s *CreateKillInstanceSessionTaskRequest) SetDbUser(v string) *CreateKillInstanceSessionTaskRequest {
	s.DbUser = &v
	return s
}

func (s *CreateKillInstanceSessionTaskRequest) SetDbUserPassword(v string) *CreateKillInstanceSessionTaskRequest {
	s.DbUserPassword = &v
	return s
}

func (s *CreateKillInstanceSessionTaskRequest) SetIgnoredUsers(v string) *CreateKillInstanceSessionTaskRequest {
	s.IgnoredUsers = &v
	return s
}

func (s *CreateKillInstanceSessionTaskRequest) SetInstanceId(v string) *CreateKillInstanceSessionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateKillInstanceSessionTaskRequest) SetKillAllSessions(v bool) *CreateKillInstanceSessionTaskRequest {
	s.KillAllSessions = &v
	return s
}

func (s *CreateKillInstanceSessionTaskRequest) SetNodeId(v string) *CreateKillInstanceSessionTaskRequest {
	s.NodeId = &v
	return s
}

func (s *CreateKillInstanceSessionTaskRequest) SetSessionIds(v string) *CreateKillInstanceSessionTaskRequest {
	s.SessionIds = &v
	return s
}

func (s *CreateKillInstanceSessionTaskRequest) Validate() error {
	return dara.Validate(s)
}
