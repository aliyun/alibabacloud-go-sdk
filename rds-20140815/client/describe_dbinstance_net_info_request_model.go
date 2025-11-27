// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBInstanceNetInfoRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest
	GetDBInstanceId() *string
	SetDBInstanceNetRWSplitType(v string) *DescribeDBInstanceNetInfoRequest
	GetDBInstanceNetRWSplitType() *string
	SetFlag(v int32) *DescribeDBInstanceNetInfoRequest
	GetFlag() *int32
	SetGeneralGroupName(v string) *DescribeDBInstanceNetInfoRequest
	GetGeneralGroupName() *string
	SetOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceNetInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceNetInfoRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Normal**: regular endpoint
	//
	// 	- **ReadWriteSplitting**: read/write splitting endpoint
	//
	// > By default, the system returns both types of endpoints.
	//
	// example:
	//
	// Normal
	DBInstanceNetRWSplitType *string `json:"DBInstanceNetRWSplitType,omitempty" xml:"DBInstanceNetRWSplitType,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// None
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The name of the dedicated cluster to which the instance belongs. This parameter takes effect only when the instance runs MySQL on RDS Standard Edition and is created in a dedicated cluster.
	//
	// example:
	//
	// rgc-2ze*****
	GeneralGroupName     *string `json:"GeneralGroupName,omitempty" xml:"GeneralGroupName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceNetInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstanceNetInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceNetInfoRequest) GetDBInstanceNetRWSplitType() *string {
	return s.DBInstanceNetRWSplitType
}

func (s *DescribeDBInstanceNetInfoRequest) GetFlag() *int32 {
	return s.Flag
}

func (s *DescribeDBInstanceNetInfoRequest) GetGeneralGroupName() *string {
	return s.GeneralGroupName
}

func (s *DescribeDBInstanceNetInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceNetInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceNetInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceNetInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceNetInfoRequest) SetClientToken(v string) *DescribeDBInstanceNetInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetDBInstanceNetRWSplitType(v string) *DescribeDBInstanceNetInfoRequest {
	s.DBInstanceNetRWSplitType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetFlag(v int32) *DescribeDBInstanceNetInfoRequest {
	s.Flag = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetGeneralGroupName(v string) *DescribeDBInstanceNetInfoRequest {
	s.GeneralGroupName = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetOwnerId(v int64) *DescribeDBInstanceNetInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceNetInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) Validate() error {
	return dara.Validate(s)
}
