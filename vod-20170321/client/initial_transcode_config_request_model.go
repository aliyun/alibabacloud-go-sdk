// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitialTranscodeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *InitialTranscodeConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *InitialTranscodeConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *InitialTranscodeConfigRequest
	GetResourceOwnerId() *int64
	SetStorageLocation(v string) *InitialTranscodeConfigRequest
	GetStorageLocation() *string
}

type InitialTranscodeConfigRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s InitialTranscodeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s InitialTranscodeConfigRequest) GoString() string {
	return s.String()
}

func (s *InitialTranscodeConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *InitialTranscodeConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *InitialTranscodeConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *InitialTranscodeConfigRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *InitialTranscodeConfigRequest) SetOwnerId(v int64) *InitialTranscodeConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *InitialTranscodeConfigRequest) SetResourceOwnerAccount(v string) *InitialTranscodeConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InitialTranscodeConfigRequest) SetResourceOwnerId(v int64) *InitialTranscodeConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InitialTranscodeConfigRequest) SetStorageLocation(v string) *InitialTranscodeConfigRequest {
	s.StorageLocation = &v
	return s
}

func (s *InitialTranscodeConfigRequest) Validate() error {
	return dara.Validate(s)
}
