// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateCloudMigrationTaskRequest
	GetDBInstanceName() *string
	SetResourceOwnerId(v int64) *CreateCloudMigrationTaskRequest
	GetResourceOwnerId() *int64
	SetSourceAccount(v string) *CreateCloudMigrationTaskRequest
	GetSourceAccount() *string
	SetSourceCategory(v string) *CreateCloudMigrationTaskRequest
	GetSourceCategory() *string
	SetSourceIpAddress(v string) *CreateCloudMigrationTaskRequest
	GetSourceIpAddress() *string
	SetSourcePassword(v string) *CreateCloudMigrationTaskRequest
	GetSourcePassword() *string
	SetSourcePort(v int64) *CreateCloudMigrationTaskRequest
	GetSourcePort() *int64
	SetTaskName(v string) *CreateCloudMigrationTaskRequest
	GetTaskName() *string
}

type CreateCloudMigrationTaskRequest struct {
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

func (s CreateCloudMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudMigrationTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateCloudMigrationTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCloudMigrationTaskRequest) GetSourceAccount() *string {
	return s.SourceAccount
}

func (s *CreateCloudMigrationTaskRequest) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *CreateCloudMigrationTaskRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *CreateCloudMigrationTaskRequest) GetSourcePassword() *string {
	return s.SourcePassword
}

func (s *CreateCloudMigrationTaskRequest) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *CreateCloudMigrationTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateCloudMigrationTaskRequest) SetDBInstanceName(v string) *CreateCloudMigrationTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetResourceOwnerId(v int64) *CreateCloudMigrationTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourceAccount(v string) *CreateCloudMigrationTaskRequest {
	s.SourceAccount = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourceCategory(v string) *CreateCloudMigrationTaskRequest {
	s.SourceCategory = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourceIpAddress(v string) *CreateCloudMigrationTaskRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourcePassword(v string) *CreateCloudMigrationTaskRequest {
	s.SourcePassword = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetSourcePort(v int64) *CreateCloudMigrationTaskRequest {
	s.SourcePort = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) SetTaskName(v string) *CreateCloudMigrationTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateCloudMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
