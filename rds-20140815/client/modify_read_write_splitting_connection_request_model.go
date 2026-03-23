// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReadWriteSplittingConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *ModifyReadWriteSplittingConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *ModifyReadWriteSplittingConnectionRequest
	GetDBInstanceId() *string
	SetDistributionType(v string) *ModifyReadWriteSplittingConnectionRequest
	GetDistributionType() *string
	SetMaxDelayTime(v string) *ModifyReadWriteSplittingConnectionRequest
	GetMaxDelayTime() *string
	SetOwnerAccount(v string) *ModifyReadWriteSplittingConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyReadWriteSplittingConnectionRequest
	GetOwnerId() *int64
	SetPort(v string) *ModifyReadWriteSplittingConnectionRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *ModifyReadWriteSplittingConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyReadWriteSplittingConnectionRequest
	GetResourceOwnerId() *int64
	SetWeight(v string) *ModifyReadWriteSplittingConnectionRequest
	GetWeight() *string
}

type ModifyReadWriteSplittingConnectionRequest struct {
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DistributionType     *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	MaxDelayTime         *string `json:"MaxDelayTime,omitempty" xml:"MaxDelayTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Weight               *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifyReadWriteSplittingConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyReadWriteSplittingConnectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetDistributionType() *string {
	return s.DistributionType
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetMaxDelayTime() *string {
	return s.MaxDelayTime
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetWeight() *string {
	return s.Weight
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetConnectionStringPrefix(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetDBInstanceId(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetDistributionType(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.DistributionType = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetMaxDelayTime(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.MaxDelayTime = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetOwnerAccount(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetOwnerId(v int64) *ModifyReadWriteSplittingConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetPort(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.Port = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetResourceOwnerAccount(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetResourceOwnerId(v int64) *ModifyReadWriteSplittingConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetWeight(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.Weight = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) Validate() error {
	return dara.Validate(s)
}
