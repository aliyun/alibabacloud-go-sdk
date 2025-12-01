// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ReleaseBackupPlanRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *ReleaseBackupPlanRequest
	GetClientToken() *string
	SetOwnerId(v string) *ReleaseBackupPlanRequest
	GetOwnerId() *string
}

type ReleaseBackupPlanRequest struct {
	// The backup schedule ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi01****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// dbs
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ReleaseBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *ReleaseBackupPlanRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ReleaseBackupPlanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReleaseBackupPlanRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ReleaseBackupPlanRequest) SetBackupPlanId(v string) *ReleaseBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ReleaseBackupPlanRequest) SetClientToken(v string) *ReleaseBackupPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleaseBackupPlanRequest) SetOwnerId(v string) *ReleaseBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
