// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentConnectionString(v string) *ReleaseInstanceConnectionRequest
	GetCurrentConnectionString() *string
	SetDBInstanceId(v string) *ReleaseInstanceConnectionRequest
	GetDBInstanceId() *string
	SetInstanceNetworkType(v string) *ReleaseInstanceConnectionRequest
	GetInstanceNetworkType() *string
	SetOwnerAccount(v string) *ReleaseInstanceConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseInstanceConnectionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReleaseInstanceConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseInstanceConnectionRequest
	GetResourceOwnerId() *int64
}

type ReleaseInstanceConnectionRequest struct {
	// This parameter is required.
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	InstanceNetworkType  *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseInstanceConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceConnectionRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ReleaseInstanceConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ReleaseInstanceConnectionRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *ReleaseInstanceConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseInstanceConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseInstanceConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseInstanceConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseInstanceConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstanceConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseInstanceConnectionRequest) SetDBInstanceId(v string) *ReleaseInstanceConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseInstanceConnectionRequest) SetInstanceNetworkType(v string) *ReleaseInstanceConnectionRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ReleaseInstanceConnectionRequest) SetOwnerAccount(v string) *ReleaseInstanceConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseInstanceConnectionRequest) SetOwnerId(v int64) *ReleaseInstanceConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseInstanceConnectionRequest) SetResourceOwnerAccount(v string) *ReleaseInstanceConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseInstanceConnectionRequest) SetResourceOwnerId(v int64) *ReleaseInstanceConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseInstanceConnectionRequest) Validate() error {
	return dara.Validate(s)
}
