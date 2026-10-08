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
	// The instance ID of the disaster recovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2zeytekus0r******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to perform a dry run for creating the synchronization link of the disaster recovery instance. Valid values:
	//
	// - **true**: Executes a dry run without creating the instance. The system checks items such as request parameters, request format, business limits, and inventory.
	//
	// - **false*	- (default): Sends a normal request and creates the instance after the check is passed.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The database account used for data synchronization.
	//
	// example:
	//
	// testdbuser
	ReplicatorAccount *string `json:"ReplicatorAccount,omitempty" xml:"ReplicatorAccount,omitempty"`
	// The password of the synchronization account.
	//
	// example:
	//
	// testpassword
	ReplicatorPassword *string `json:"ReplicatorPassword,omitempty" xml:"ReplicatorPassword,omitempty"`
	// The endpoint of the PostgreSQL source instance or the IP address of the SQL Server source instance.
	//
	// example:
	//
	// PostgreSQL：pgm-****.pg.rds.aliyuncs.com
	//
	// SQL Server：10.XX.XXX.XXX
	SourceAddress *string `json:"SourceAddress,omitempty" xml:"SourceAddress,omitempty"`
	// The category of the source instance. Valid values:
	//
	// - **other**: Other. (**Not supported for SQL Server.**)
	//
	// - **aliyunRDS**: ApsaraDB RDS instance.
	//
	// example:
	//
	// aliyunRDS
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// The name of the source instance. This parameter is required when **SourceCategory*	- is set to **aliyunRDS**.
	//
	// example:
	//
	// rm-2zeaaz62s18******
	SourceInstanceName *string `json:"SourceInstanceName,omitempty" xml:"SourceInstanceName,omitempty"`
	// The region ID of the source instance. This parameter is required when **SourceCategory*	- is set to **aliyunRDS**.
	//
	// example:
	//
	// cn-hangzhou
	SourceInstanceRegionId *string `json:"SourceInstanceRegionId,omitempty" xml:"SourceInstanceRegionId,omitempty"`
	// The port of the source instance.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The IP address of the SQL Server disaster recovery instance.
	//
	// example:
	//
	// 192.XXX.XX.XXX
	TargetAddress *string `json:"TargetAddress,omitempty" xml:"TargetAddress,omitempty"`
	// The ID of a successful dry run task.
	//
	// example:
	//
	// 43994****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the dry run task. You can specify a custom name. If you do not specify this parameter, the system automatically generates a name.
	//
	// example:
	//
	// zbtest
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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
