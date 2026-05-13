// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairSkvDdbTableSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeTairSkvDdbTableSchemaRequest
	GetBackupId() *string
	SetInstanceId(v string) *DescribeTairSkvDdbTableSchemaRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeTairSkvDdbTableSchemaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTairSkvDdbTableSchemaRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeTairSkvDdbTableSchemaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTairSkvDdbTableSchemaRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeTairSkvDdbTableSchemaRequest
	GetSecurityToken() *string
}

type DescribeTairSkvDdbTableSchemaRequest struct {
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tt-bp1zxszhcgatnx\\\\*\\\\*\\\\*\\\\*
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeTairSkvDdbTableSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableSchemaRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableSchemaRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeTairSkvDdbTableSchemaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairSkvDdbTableSchemaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTairSkvDdbTableSchemaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTairSkvDdbTableSchemaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTairSkvDdbTableSchemaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTairSkvDdbTableSchemaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairSkvDdbTableSchemaRequest) SetBackupId(v string) *DescribeTairSkvDdbTableSchemaRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaRequest) SetInstanceId(v string) *DescribeTairSkvDdbTableSchemaRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaRequest) SetOwnerAccount(v string) *DescribeTairSkvDdbTableSchemaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaRequest) SetOwnerId(v int64) *DescribeTairSkvDdbTableSchemaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaRequest) SetResourceOwnerAccount(v string) *DescribeTairSkvDdbTableSchemaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaRequest) SetResourceOwnerId(v int64) *DescribeTairSkvDdbTableSchemaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaRequest) SetSecurityToken(v string) *DescribeTairSkvDdbTableSchemaRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaRequest) Validate() error {
	return dara.Validate(s)
}
