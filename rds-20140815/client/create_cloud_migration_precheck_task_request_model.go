// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMigrationPrecheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetDBInstanceName() *string
	SetResourceOwnerId(v int64) *CreateCloudMigrationPrecheckTaskRequest
	GetResourceOwnerId() *int64
	SetSourceAccount(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetSourceAccount() *string
	SetSourceCategory(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetSourceCategory() *string
	SetSourceIpAddress(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetSourceIpAddress() *string
	SetSourcePassword(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetSourcePassword() *string
	SetSourcePort(v int64) *CreateCloudMigrationPrecheckTaskRequest
	GetSourcePort() *int64
	SetTaskName(v string) *CreateCloudMigrationPrecheckTaskRequest
	GetTaskName() *string
}

type CreateCloudMigrationPrecheckTaskRequest struct {
	// This parameter is required.
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SourceAccount *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	// This parameter is required.
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// This parameter is required.
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// This parameter is required.
	SourcePassword *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	// This parameter is required.
	SourcePort *int64  `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateCloudMigrationPrecheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMigrationPrecheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourceAccount() *string {
	return s.SourceAccount
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourcePassword() *string {
	return s.SourcePassword
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *CreateCloudMigrationPrecheckTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetDBInstanceName(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetResourceOwnerId(v int64) *CreateCloudMigrationPrecheckTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourceAccount(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourceAccount = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourceCategory(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourceCategory = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourceIpAddress(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourcePassword(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourcePassword = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetSourcePort(v int64) *CreateCloudMigrationPrecheckTaskRequest {
	s.SourcePort = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) SetTaskName(v string) *CreateCloudMigrationPrecheckTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateCloudMigrationPrecheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
