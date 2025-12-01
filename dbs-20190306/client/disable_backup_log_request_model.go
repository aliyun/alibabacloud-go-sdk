// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBackupLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DisableBackupLogRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *DisableBackupLogRequest
	GetClientToken() *string
	SetOwnerId(v string) *DisableBackupLogRequest
	GetOwnerId() *string
}

type DisableBackupLogRequest struct {
	// The backup schedule ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01xxxx
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// dbs
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DisableBackupLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableBackupLogRequest) GoString() string {
	return s.String()
}

func (s *DisableBackupLogRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DisableBackupLogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableBackupLogRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DisableBackupLogRequest) SetBackupPlanId(v string) *DisableBackupLogRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DisableBackupLogRequest) SetClientToken(v string) *DisableBackupLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableBackupLogRequest) SetOwnerId(v string) *DisableBackupLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableBackupLogRequest) Validate() error {
	return dara.Validate(s)
}
