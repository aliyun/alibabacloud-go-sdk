// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoForChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetDBInstanceId() *string
	SetDBInstanceNetRWSplitType(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetDBInstanceNetRWSplitType() *string
	SetFlag(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetFlag() *string
	SetOwnerAccount(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceNetInfoForChannelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceNetInfoForChannelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoForChannelRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceNetInfoForChannelRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId             *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceNetRWSplitType *string `json:"DBInstanceNetRWSplitType,omitempty" xml:"DBInstanceNetRWSplitType,omitempty"`
	Flag                     *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceNetInfoForChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoForChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetDBInstanceNetRWSplitType() *string {
	return s.DBInstanceNetRWSplitType
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetFlag() *string {
	return s.Flag
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetClientToken(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetDBInstanceNetRWSplitType(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.DBInstanceNetRWSplitType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetFlag(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.Flag = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetOwnerAccount(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetOwnerId(v int64) *DescribeDBInstanceNetInfoForChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceNetInfoForChannelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoForChannelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoForChannelRequest) Validate() error {
	return dara.Validate(s)
}
