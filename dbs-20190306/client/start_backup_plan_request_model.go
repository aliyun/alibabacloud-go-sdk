// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *StartBackupPlanRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *StartBackupPlanRequest
	GetClientToken() *string
	SetOwnerId(v string) *StartBackupPlanRequest
	GetOwnerId() *string
}

type StartBackupPlanRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain it.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbsqdss5tmh****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StartBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s StartBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *StartBackupPlanRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *StartBackupPlanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartBackupPlanRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *StartBackupPlanRequest) SetBackupPlanId(v string) *StartBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *StartBackupPlanRequest) SetClientToken(v string) *StartBackupPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *StartBackupPlanRequest) SetOwnerId(v string) *StartBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *StartBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
