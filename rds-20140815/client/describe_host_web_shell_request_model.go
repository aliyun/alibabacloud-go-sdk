// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHostWebShellRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeHostWebShellRequest
	GetAccountName() *string
	SetAccountPassword(v string) *DescribeHostWebShellRequest
	GetAccountPassword() *string
	SetDBInstanceId(v string) *DescribeHostWebShellRequest
	GetDBInstanceId() *string
	SetHostName(v string) *DescribeHostWebShellRequest
	GetHostName() *string
	SetOwnerId(v int64) *DescribeHostWebShellRequest
	GetOwnerId() *int64
	SetRegionID(v string) *DescribeHostWebShellRequest
	GetRegionID() *string
	SetResourceOwnerAccount(v string) *DescribeHostWebShellRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHostWebShellRequest
	GetResourceOwnerId() *int64
}

type DescribeHostWebShellRequest struct {
	// The username of the account that is used to log on to the host of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// testOsAccount1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the host account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ***
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance hostname. You can call the DescribeDBInstanceIpHostname operation to query the hostname.
	//
	// This parameter is required.
	//
	// example:
	//
	// testHost1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the DescribeDBInstanceAttribute operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionID             *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHostWebShellRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostWebShellRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostWebShellRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeHostWebShellRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *DescribeHostWebShellRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeHostWebShellRequest) GetHostName() *string {
	return s.HostName
}

func (s *DescribeHostWebShellRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHostWebShellRequest) GetRegionID() *string {
	return s.RegionID
}

func (s *DescribeHostWebShellRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHostWebShellRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHostWebShellRequest) SetAccountName(v string) *DescribeHostWebShellRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetAccountPassword(v string) *DescribeHostWebShellRequest {
	s.AccountPassword = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetDBInstanceId(v string) *DescribeHostWebShellRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetHostName(v string) *DescribeHostWebShellRequest {
	s.HostName = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetOwnerId(v int64) *DescribeHostWebShellRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetRegionID(v string) *DescribeHostWebShellRequest {
	s.RegionID = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetResourceOwnerAccount(v string) *DescribeHostWebShellRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetResourceOwnerId(v int64) *DescribeHostWebShellRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHostWebShellRequest) Validate() error {
	return dara.Validate(s)
}
