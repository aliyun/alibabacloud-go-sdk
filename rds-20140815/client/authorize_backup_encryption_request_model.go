// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeBackupEncryptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *AuthorizeBackupEncryptionRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *AuthorizeBackupEncryptionRequest
	GetResourceOwnerId() *int64
}

type AuthorizeBackupEncryptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AuthorizeBackupEncryptionRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeBackupEncryptionRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeBackupEncryptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeBackupEncryptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AuthorizeBackupEncryptionRequest) SetRegionId(v string) *AuthorizeBackupEncryptionRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeBackupEncryptionRequest) SetResourceOwnerId(v int64) *AuthorizeBackupEncryptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AuthorizeBackupEncryptionRequest) Validate() error {
	return dara.Validate(s)
}
