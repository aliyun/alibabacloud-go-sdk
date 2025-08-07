// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *CreateRestoreTaskRequest
	GetBackupId() *string
	SetServiceInstanceId(v string) *CreateRestoreTaskRequest
	GetServiceInstanceId() *string
}

type CreateRestoreTaskRequest struct {
	// The backup ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// backup-385d35fb088f453c8fa1
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the service instance.
	//
	// You can call [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) to obtain the ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-49793f3bfa034ec6a990
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s CreateRestoreTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRestoreTaskRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateRestoreTaskRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *CreateRestoreTaskRequest) SetBackupId(v string) *CreateRestoreTaskRequest {
	s.BackupId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetServiceInstanceId(v string) *CreateRestoreTaskRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateRestoreTaskRequest) Validate() error {
	return dara.Validate(s)
}
