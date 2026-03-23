// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReplicationLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateReplicationLinkRequest
	GetDBInstanceId() *string
	SetDryRun(v bool) *CreateReplicationLinkRequest
	GetDryRun() *bool
	SetReplicatorAccount(v string) *CreateReplicationLinkRequest
	GetReplicatorAccount() *string
	SetReplicatorPassword(v string) *CreateReplicationLinkRequest
	GetReplicatorPassword() *string
	SetSourceAddress(v string) *CreateReplicationLinkRequest
	GetSourceAddress() *string
	SetSourceCategory(v string) *CreateReplicationLinkRequest
	GetSourceCategory() *string
	SetSourceInstanceName(v string) *CreateReplicationLinkRequest
	GetSourceInstanceName() *string
	SetSourceInstanceRegionId(v string) *CreateReplicationLinkRequest
	GetSourceInstanceRegionId() *string
	SetSourcePort(v int64) *CreateReplicationLinkRequest
	GetSourcePort() *int64
	SetTargetAddress(v string) *CreateReplicationLinkRequest
	GetTargetAddress() *string
	SetTaskId(v int64) *CreateReplicationLinkRequest
	GetTaskId() *int64
	SetTaskName(v string) *CreateReplicationLinkRequest
	GetTaskName() *string
}

type CreateReplicationLinkRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DryRun                 *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ReplicatorAccount      *string `json:"ReplicatorAccount,omitempty" xml:"ReplicatorAccount,omitempty"`
	ReplicatorPassword     *string `json:"ReplicatorPassword,omitempty" xml:"ReplicatorPassword,omitempty"`
	SourceAddress          *string `json:"SourceAddress,omitempty" xml:"SourceAddress,omitempty"`
	SourceCategory         *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	SourceInstanceName     *string `json:"SourceInstanceName,omitempty" xml:"SourceInstanceName,omitempty"`
	SourceInstanceRegionId *string `json:"SourceInstanceRegionId,omitempty" xml:"SourceInstanceRegionId,omitempty"`
	SourcePort             *int64  `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	TargetAddress          *string `json:"TargetAddress,omitempty" xml:"TargetAddress,omitempty"`
	TaskId                 *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName               *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateReplicationLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateReplicationLinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateReplicationLinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateReplicationLinkRequest) GetReplicatorAccount() *string {
	return s.ReplicatorAccount
}

func (s *CreateReplicationLinkRequest) GetReplicatorPassword() *string {
	return s.ReplicatorPassword
}

func (s *CreateReplicationLinkRequest) GetSourceAddress() *string {
	return s.SourceAddress
}

func (s *CreateReplicationLinkRequest) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *CreateReplicationLinkRequest) GetSourceInstanceName() *string {
	return s.SourceInstanceName
}

func (s *CreateReplicationLinkRequest) GetSourceInstanceRegionId() *string {
	return s.SourceInstanceRegionId
}

func (s *CreateReplicationLinkRequest) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *CreateReplicationLinkRequest) GetTargetAddress() *string {
	return s.TargetAddress
}

func (s *CreateReplicationLinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateReplicationLinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateReplicationLinkRequest) SetDBInstanceId(v string) *CreateReplicationLinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetDryRun(v bool) *CreateReplicationLinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetReplicatorAccount(v string) *CreateReplicationLinkRequest {
	s.ReplicatorAccount = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetReplicatorPassword(v string) *CreateReplicationLinkRequest {
	s.ReplicatorPassword = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetSourceAddress(v string) *CreateReplicationLinkRequest {
	s.SourceAddress = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetSourceCategory(v string) *CreateReplicationLinkRequest {
	s.SourceCategory = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetSourceInstanceName(v string) *CreateReplicationLinkRequest {
	s.SourceInstanceName = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetSourceInstanceRegionId(v string) *CreateReplicationLinkRequest {
	s.SourceInstanceRegionId = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetSourcePort(v int64) *CreateReplicationLinkRequest {
	s.SourcePort = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetTargetAddress(v string) *CreateReplicationLinkRequest {
	s.TargetAddress = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetTaskId(v int64) *CreateReplicationLinkRequest {
	s.TaskId = &v
	return s
}

func (s *CreateReplicationLinkRequest) SetTaskName(v string) *CreateReplicationLinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateReplicationLinkRequest) Validate() error {
	return dara.Validate(s)
}
