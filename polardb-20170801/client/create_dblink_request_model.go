// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBLinkRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateDBLinkRequest
	GetDBClusterId() *string
	SetDBLinkName(v string) *CreateDBLinkRequest
	GetDBLinkName() *string
	SetOwnerAccount(v string) *CreateDBLinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBLinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateDBLinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBLinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBLinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBLinkRequest
	GetResourceOwnerId() *int64
	SetSourceDBName(v string) *CreateDBLinkRequest
	GetSourceDBName() *string
	SetTargetDBAccount(v string) *CreateDBLinkRequest
	GetTargetDBAccount() *string
	SetTargetDBInstanceName(v string) *CreateDBLinkRequest
	GetTargetDBInstanceName() *string
	SetTargetDBName(v string) *CreateDBLinkRequest
	GetTargetDBName() *string
	SetTargetDBPasswd(v string) *CreateDBLinkRequest
	GetTargetDBPasswd() *string
	SetTargetIp(v string) *CreateDBLinkRequest
	GetTargetIp() *string
	SetTargetPort(v string) *CreateDBLinkRequest
	GetTargetPort() *string
	SetVpcId(v string) *CreateDBLinkRequest
	GetVpcId() *string
}

type CreateDBLinkRequest struct {
	// A client token that is used to ensure the idempotence of the request. The client generates the token, but you must make sure that the token is unique among different requests. The token is case-sensitive and must not exceed 64 ASCII characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the source cluster for the DBLink.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/173433.html) operation to query the list of PolarDB clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-a************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the DBLink.
	//
	// - It must contain lowercase letters, and can also contain digits and underscores (_).
	//
	// - It must start with a letter and end with a letter or a digit.
	//
	// - It must be no more than 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// dblink_test
	DBLinkName   *string `json:"DBLinkName,omitempty" xml:"DBLinkName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the details of regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the source database.
	//
	// > You can call the [DescribeDatabases](https://help.aliyun.com/document_detail/173558.html) operation to query information about databases in a PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb1
	SourceDBName *string `json:"SourceDBName,omitempty" xml:"SourceDBName,omitempty"`
	// The account of the destination database.
	//
	// > You can call the [DescribeAccounts](https://help.aliyun.com/document_detail/173549.html) operation to query the database accounts of a PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	TargetDBAccount *string `json:"TargetDBAccount,omitempty" xml:"TargetDBAccount,omitempty"`
	// The ID of the destination cluster for the DBLink.
	//
	// > - If the destination is a self-managed Oracle database that runs on an ECS instance, set this parameter to `null`.
	//
	// >
	//
	// > - You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/173433.html) operation to query the list of PolarDB clusters.
	//
	// example:
	//
	// pc-b************
	TargetDBInstanceName *string `json:"TargetDBInstanceName,omitempty" xml:"TargetDBInstanceName,omitempty"`
	// The name of the destination database.
	//
	// > You can call the [DescribeDatabases](https://help.aliyun.com/document_detail/173558.html) operation to query information about databases in a PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb2
	TargetDBName *string `json:"TargetDBName,omitempty" xml:"TargetDBName,omitempty"`
	// The password for the destination database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test1111
	TargetDBPasswd *string `json:"TargetDBPasswd,omitempty" xml:"TargetDBPasswd,omitempty"`
	// The IP address of the self-managed Oracle database that runs on an ECS instance.
	//
	// example:
	//
	// 192.**.**.46
	TargetIp *string `json:"TargetIp,omitempty" xml:"TargetIp,omitempty"`
	// The port number of the self-managed Oracle database that runs on an ECS instance.
	//
	// example:
	//
	// 1521
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// > You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to query the details of VPCs.
	//
	// example:
	//
	// vpc-bp1qpo0kug3a20qqe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateDBLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBLinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBLinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBLinkRequest) GetDBLinkName() *string {
	return s.DBLinkName
}

func (s *CreateDBLinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBLinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBLinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBLinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBLinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBLinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBLinkRequest) GetSourceDBName() *string {
	return s.SourceDBName
}

func (s *CreateDBLinkRequest) GetTargetDBAccount() *string {
	return s.TargetDBAccount
}

func (s *CreateDBLinkRequest) GetTargetDBInstanceName() *string {
	return s.TargetDBInstanceName
}

func (s *CreateDBLinkRequest) GetTargetDBName() *string {
	return s.TargetDBName
}

func (s *CreateDBLinkRequest) GetTargetDBPasswd() *string {
	return s.TargetDBPasswd
}

func (s *CreateDBLinkRequest) GetTargetIp() *string {
	return s.TargetIp
}

func (s *CreateDBLinkRequest) GetTargetPort() *string {
	return s.TargetPort
}

func (s *CreateDBLinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDBLinkRequest) SetClientToken(v string) *CreateDBLinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBLinkRequest) SetDBClusterId(v string) *CreateDBLinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBLinkRequest) SetDBLinkName(v string) *CreateDBLinkRequest {
	s.DBLinkName = &v
	return s
}

func (s *CreateDBLinkRequest) SetOwnerAccount(v string) *CreateDBLinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBLinkRequest) SetOwnerId(v int64) *CreateDBLinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBLinkRequest) SetRegionId(v string) *CreateDBLinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBLinkRequest) SetResourceGroupId(v string) *CreateDBLinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBLinkRequest) SetResourceOwnerAccount(v string) *CreateDBLinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBLinkRequest) SetResourceOwnerId(v int64) *CreateDBLinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBLinkRequest) SetSourceDBName(v string) *CreateDBLinkRequest {
	s.SourceDBName = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetDBAccount(v string) *CreateDBLinkRequest {
	s.TargetDBAccount = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetDBInstanceName(v string) *CreateDBLinkRequest {
	s.TargetDBInstanceName = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetDBName(v string) *CreateDBLinkRequest {
	s.TargetDBName = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetDBPasswd(v string) *CreateDBLinkRequest {
	s.TargetDBPasswd = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetIp(v string) *CreateDBLinkRequest {
	s.TargetIp = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetPort(v string) *CreateDBLinkRequest {
	s.TargetPort = &v
	return s
}

func (s *CreateDBLinkRequest) SetVpcId(v string) *CreateDBLinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBLinkRequest) Validate() error {
	return dara.Validate(s)
}
