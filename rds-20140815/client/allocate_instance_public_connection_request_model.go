// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBabelfishPort(v string) *AllocateInstancePublicConnectionRequest
	GetBabelfishPort() *string
	SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest
	GetDBInstanceId() *string
	SetGeneralGroupName(v string) *AllocateInstancePublicConnectionRequest
	GetGeneralGroupName() *string
	SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetOwnerId() *int64
	SetPGBouncerPort(v string) *AllocateInstancePublicConnectionRequest
	GetPGBouncerPort() *string
	SetPort(v string) *AllocateInstancePublicConnectionRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest
	GetResourceOwnerId() *int64
}

type AllocateInstancePublicConnectionRequest struct {
	BabelfishPort *string `json:"BabelfishPort,omitempty" xml:"BabelfishPort,omitempty"`
	// This parameter is required.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// This parameter is required.
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	GeneralGroupName *string `json:"GeneralGroupName,omitempty" xml:"GeneralGroupName,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PGBouncerPort    *string `json:"PGBouncerPort,omitempty" xml:"PGBouncerPort,omitempty"`
	// This parameter is required.
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) GetBabelfishPort() *string {
	return s.BabelfishPort
}

func (s *AllocateInstancePublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateInstancePublicConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AllocateInstancePublicConnectionRequest) GetGeneralGroupName() *string {
	return s.GeneralGroupName
}

func (s *AllocateInstancePublicConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateInstancePublicConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateInstancePublicConnectionRequest) GetPGBouncerPort() *string {
	return s.PGBouncerPort
}

func (s *AllocateInstancePublicConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateInstancePublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateInstancePublicConnectionRequest) SetBabelfishPort(v string) *AllocateInstancePublicConnectionRequest {
	s.BabelfishPort = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetGeneralGroupName(v string) *AllocateInstancePublicConnectionRequest {
	s.GeneralGroupName = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPGBouncerPort(v string) *AllocateInstancePublicConnectionRequest {
	s.PGBouncerPort = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
