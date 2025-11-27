// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIPArrayListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceIPArrayListRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceIPArrayListRequest
	GetOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceIPArrayListRequest
	GetResourceOwnerId() *int64
	SetWhitelistNetworkType(v string) *DescribeDBInstanceIPArrayListRequest
	GetWhitelistNetworkType() *string
}

type DescribeDBInstanceIPArrayListRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The network type of the IP address whitelist. Valid values:
	//
	// 	- **Classic**: classic network in enhanced whitelist mode
	//
	// 	- **VPC**: virtual private cloud (VPC) in enhanced whitelist mode
	//
	// 	- **MIX**: standard whitelist mode
	//
	// By default, this operation returns IP address whitelists of all network types.
	//
	// example:
	//
	// VPC
	WhitelistNetworkType *string `json:"WhitelistNetworkType,omitempty" xml:"WhitelistNetworkType,omitempty"`
}

func (s DescribeDBInstanceIPArrayListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceIPArrayListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceIPArrayListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceIPArrayListRequest) GetWhitelistNetworkType() *string {
	return s.WhitelistNetworkType
}

func (s *DescribeDBInstanceIPArrayListRequest) SetDBInstanceId(v string) *DescribeDBInstanceIPArrayListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListRequest) SetOwnerAccount(v string) *DescribeDBInstanceIPArrayListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceIPArrayListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListRequest) SetWhitelistNetworkType(v string) *DescribeDBInstanceIPArrayListRequest {
	s.WhitelistNetworkType = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListRequest) Validate() error {
	return dara.Validate(s)
}
