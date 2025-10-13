// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupType(v string) *ListBackupDataRequest
	GetBackupType() *string
	SetInstanceId(v string) *ListBackupDataRequest
	GetInstanceId() *string
}

type ListBackupDataRequest struct {
	// The backup type. Specific backup data is filtered based on the type. If you leave this parameter empty, all backup data is returned.
	//
	// Valid values:
	//
	// 	- redundant_remote
	//
	// 	- remote
	//
	// 	- redundant
	//
	// 	- full_remote
	//
	// 	- local
	//
	// 	- full
	//
	// example:
	//
	// redundant
	BackupType *string `json:"backupType,omitempty" xml:"backupType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgprecn-cn-wwoxxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListBackupDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBackupDataRequest) GoString() string {
	return s.String()
}

func (s *ListBackupDataRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *ListBackupDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBackupDataRequest) SetBackupType(v string) *ListBackupDataRequest {
	s.BackupType = &v
	return s
}

func (s *ListBackupDataRequest) SetInstanceId(v string) *ListBackupDataRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBackupDataRequest) Validate() error {
	return dara.Validate(s)
}
