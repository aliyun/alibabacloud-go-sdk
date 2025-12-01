// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckProgressListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribePreCheckProgressListRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *DescribePreCheckProgressListRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribePreCheckProgressListRequest
	GetOwnerId() *string
	SetRestoreTaskId(v string) *DescribePreCheckProgressListRequest
	GetRestoreTaskId() *string
}

type DescribePreCheckProgressListRequest struct {
	// The backup schedule ID.
	//
	// >  You must specify one of BackupPlanId and RestoreTaskId.
	//
	// example:
	//
	// dbstooi01XXXX
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The restoration task ID.
	//
	// example:
	//
	// dbasdsaXXXX
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
}

func (s DescribePreCheckProgressListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckProgressListRequest) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribePreCheckProgressListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePreCheckProgressListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribePreCheckProgressListRequest) GetRestoreTaskId() *string {
	return s.RestoreTaskId
}

func (s *DescribePreCheckProgressListRequest) SetBackupPlanId(v string) *DescribePreCheckProgressListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribePreCheckProgressListRequest) SetClientToken(v string) *DescribePreCheckProgressListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePreCheckProgressListRequest) SetOwnerId(v string) *DescribePreCheckProgressListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePreCheckProgressListRequest) SetRestoreTaskId(v string) *DescribePreCheckProgressListRequest {
	s.RestoreTaskId = &v
	return s
}

func (s *DescribePreCheckProgressListRequest) Validate() error {
	return dara.Validate(s)
}
