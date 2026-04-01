// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBackupEncryptionAuthorizedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CheckBackupEncryptionAuthorizedRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *CheckBackupEncryptionAuthorizedRequest
	GetResourceOwnerId() *int64
}

type CheckBackupEncryptionAuthorizedRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckBackupEncryptionAuthorizedRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckBackupEncryptionAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *CheckBackupEncryptionAuthorizedRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckBackupEncryptionAuthorizedRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckBackupEncryptionAuthorizedRequest) SetRegionId(v string) *CheckBackupEncryptionAuthorizedRequest {
	s.RegionId = &v
	return s
}

func (s *CheckBackupEncryptionAuthorizedRequest) SetResourceOwnerId(v int64) *CheckBackupEncryptionAuthorizedRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckBackupEncryptionAuthorizedRequest) Validate() error {
	return dara.Validate(s)
}
