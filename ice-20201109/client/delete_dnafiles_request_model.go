// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDNAFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBId(v string) *DeleteDNAFilesRequest
	GetDBId() *string
	SetOwnerAccount(v string) *DeleteDNAFilesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDNAFilesRequest
	GetOwnerId() *int64
	SetPrimaryKeys(v string) *DeleteDNAFilesRequest
	GetPrimaryKeys() *string
	SetResourceOwnerAccount(v string) *DeleteDNAFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDNAFilesRequest
	GetResourceOwnerId() *int64
}

type DeleteDNAFilesRequest struct {
	// The ID of the media fingerprint library from which you want to delete files.
	//
	// This parameter is required.
	//
	// example:
	//
	// fb712a6890464059b1b2ea7c8647****
	DBId         *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The primary key values of the files that you want to delete. Separate multiple values with commas (,). You can delete up to 50 files at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 41e6536e4f2250e2e9bf26cdea19****
	PrimaryKeys          *string `json:"PrimaryKeys,omitempty" xml:"PrimaryKeys,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDNAFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDNAFilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDNAFilesRequest) GetDBId() *string {
	return s.DBId
}

func (s *DeleteDNAFilesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDNAFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDNAFilesRequest) GetPrimaryKeys() *string {
	return s.PrimaryKeys
}

func (s *DeleteDNAFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDNAFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDNAFilesRequest) SetDBId(v string) *DeleteDNAFilesRequest {
	s.DBId = &v
	return s
}

func (s *DeleteDNAFilesRequest) SetOwnerAccount(v string) *DeleteDNAFilesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDNAFilesRequest) SetOwnerId(v int64) *DeleteDNAFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDNAFilesRequest) SetPrimaryKeys(v string) *DeleteDNAFilesRequest {
	s.PrimaryKeys = &v
	return s
}

func (s *DeleteDNAFilesRequest) SetResourceOwnerAccount(v string) *DeleteDNAFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDNAFilesRequest) SetResourceOwnerId(v int64) *DeleteDNAFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDNAFilesRequest) Validate() error {
	return dara.Validate(s)
}
