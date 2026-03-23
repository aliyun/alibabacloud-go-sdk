// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateReadWriteSplittingConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateReadWriteSplittingConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *AllocateReadWriteSplittingConnectionRequest
	GetDBInstanceId() *string
	SetDistributionType(v string) *AllocateReadWriteSplittingConnectionRequest
	GetDistributionType() *string
	SetMaxDelayTime(v string) *AllocateReadWriteSplittingConnectionRequest
	GetMaxDelayTime() *string
	SetNetType(v string) *AllocateReadWriteSplittingConnectionRequest
	GetNetType() *string
	SetOwnerAccount(v string) *AllocateReadWriteSplittingConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateReadWriteSplittingConnectionRequest
	GetOwnerId() *int64
	SetPort(v string) *AllocateReadWriteSplittingConnectionRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *AllocateReadWriteSplittingConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateReadWriteSplittingConnectionRequest
	GetResourceOwnerId() *int64
	SetWeight(v string) *AllocateReadWriteSplittingConnectionRequest
	GetWeight() *string
}

type AllocateReadWriteSplittingConnectionRequest struct {
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DistributionType     *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	MaxDelayTime         *string `json:"MaxDelayTime,omitempty" xml:"MaxDelayTime,omitempty"`
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Weight               *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AllocateReadWriteSplittingConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateReadWriteSplittingConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetDistributionType() *string {
	return s.DistributionType
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetMaxDelayTime() *string {
	return s.MaxDelayTime
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetNetType() *string {
	return s.NetType
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetWeight() *string {
	return s.Weight
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetConnectionStringPrefix(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetDBInstanceId(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetDistributionType(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.DistributionType = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetMaxDelayTime(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.MaxDelayTime = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetNetType(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.NetType = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetOwnerAccount(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetOwnerId(v int64) *AllocateReadWriteSplittingConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetPort(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetResourceOwnerAccount(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetResourceOwnerId(v int64) *AllocateReadWriteSplittingConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetWeight(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.Weight = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) Validate() error {
	return dara.Validate(s)
}
