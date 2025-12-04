// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortsForClickHouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreatePortsForClickHouseRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CreatePortsForClickHouseRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePortsForClickHouseRequest
	GetOwnerId() *int64
	SetPortType(v string) *CreatePortsForClickHouseRequest
	GetPortType() *string
	SetRegionId(v string) *CreatePortsForClickHouseRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreatePortsForClickHouseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePortsForClickHouseRequest
	GetResourceOwnerId() *int64
}

type CreatePortsForClickHouseRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port type. Set the value to mysql_port.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_port
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreatePortsForClickHouseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePortsForClickHouseRequest) GoString() string {
	return s.String()
}

func (s *CreatePortsForClickHouseRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreatePortsForClickHouseRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePortsForClickHouseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePortsForClickHouseRequest) GetPortType() *string {
	return s.PortType
}

func (s *CreatePortsForClickHouseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePortsForClickHouseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePortsForClickHouseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePortsForClickHouseRequest) SetDBClusterId(v string) *CreatePortsForClickHouseRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetOwnerAccount(v string) *CreatePortsForClickHouseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetOwnerId(v int64) *CreatePortsForClickHouseRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetPortType(v string) *CreatePortsForClickHouseRequest {
	s.PortType = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetRegionId(v string) *CreatePortsForClickHouseRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetResourceOwnerAccount(v string) *CreatePortsForClickHouseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetResourceOwnerId(v int64) *CreatePortsForClickHouseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) Validate() error {
	return dara.Validate(s)
}
