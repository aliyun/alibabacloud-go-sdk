// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRegionSupportBackupEncryptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceID(v string) *CheckRegionSupportBackupEncryptionRequest
	GetDBInstanceID() *string
	SetRegionId(v string) *CheckRegionSupportBackupEncryptionRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *CheckRegionSupportBackupEncryptionRequest
	GetResourceOwnerId() *int64
}

type CheckRegionSupportBackupEncryptionRequest struct {
	// example:
	//
	// rm-wz91q53f9*******
	DBInstanceID *string `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckRegionSupportBackupEncryptionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckRegionSupportBackupEncryptionRequest) GoString() string {
	return s.String()
}

func (s *CheckRegionSupportBackupEncryptionRequest) GetDBInstanceID() *string {
	return s.DBInstanceID
}

func (s *CheckRegionSupportBackupEncryptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckRegionSupportBackupEncryptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckRegionSupportBackupEncryptionRequest) SetDBInstanceID(v string) *CheckRegionSupportBackupEncryptionRequest {
	s.DBInstanceID = &v
	return s
}

func (s *CheckRegionSupportBackupEncryptionRequest) SetRegionId(v string) *CheckRegionSupportBackupEncryptionRequest {
	s.RegionId = &v
	return s
}

func (s *CheckRegionSupportBackupEncryptionRequest) SetResourceOwnerId(v int64) *CheckRegionSupportBackupEncryptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckRegionSupportBackupEncryptionRequest) Validate() error {
	return dara.Validate(s)
}
