// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeUserEncryptionKeyListRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeUserEncryptionKeyListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUserEncryptionKeyListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeUserEncryptionKeyListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUserEncryptionKeyListRequest
	GetResourceOwnerId() *int64
	SetRoleARN(v string) *DescribeUserEncryptionKeyListRequest
	GetRoleARN() *string
	SetTargetRegionId(v string) *DescribeUserEncryptionKeyListRequest
	GetTargetRegionId() *string
}

type DescribeUserEncryptionKeyListRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the specified role. Format: `acs:ram::$accountID:role/$roleName `. Default value: `acs:ram::$accountID:role/aliyunrdsinstanceencryptiondefaultrole`.
	//
	// > - `$accountID`: the Alibaba Cloud account ID. You can log on to the Alibaba Cloud Management Console, move the mouse over the profile picture in the upper-right corner, and click Security Settings to view it.
	//
	// > - `$roleName`: the name of the RAM role. You can log on to the Resource Access Management (RAM) console, click RAM Roles in the left-side navigation pane, and view the name in the RAM role name list.
	//
	// example:
	//
	// acs:ram::123456789012****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The target zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query available zones.
	//
	// example:
	//
	// cn-shanghai
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
}

func (s DescribeUserEncryptionKeyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeUserEncryptionKeyListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUserEncryptionKeyListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserEncryptionKeyListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUserEncryptionKeyListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUserEncryptionKeyListRequest) GetRoleARN() *string {
	return s.RoleARN
}

func (s *DescribeUserEncryptionKeyListRequest) GetTargetRegionId() *string {
	return s.TargetRegionId
}

func (s *DescribeUserEncryptionKeyListRequest) SetDBInstanceId(v string) *DescribeUserEncryptionKeyListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetOwnerAccount(v string) *DescribeUserEncryptionKeyListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetOwnerId(v int64) *DescribeUserEncryptionKeyListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetResourceOwnerAccount(v string) *DescribeUserEncryptionKeyListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetResourceOwnerId(v int64) *DescribeUserEncryptionKeyListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetRoleARN(v string) *DescribeUserEncryptionKeyListRequest {
	s.RoleARN = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetTargetRegionId(v string) *DescribeUserEncryptionKeyListRequest {
	s.TargetRegionId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) Validate() error {
	return dara.Validate(s)
}
