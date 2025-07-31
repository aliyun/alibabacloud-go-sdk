// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleZoneInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeRoleZoneInfoRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeRoleZoneInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRoleZoneInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeRoleZoneInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRoleZoneInfoRequest
	GetResourceOwnerId() *int64
}

type DescribeRoleZoneInfoRequest struct {
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
}

func (s DescribeRoleZoneInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeRoleZoneInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRoleZoneInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRoleZoneInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRoleZoneInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRoleZoneInfoRequest) SetDBInstanceId(v string) *DescribeRoleZoneInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetOwnerAccount(v string) *DescribeRoleZoneInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetOwnerId(v int64) *DescribeRoleZoneInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetResourceOwnerAccount(v string) *DescribeRoleZoneInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetResourceOwnerId(v int64) *DescribeRoleZoneInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) Validate() error {
	return dara.Validate(s)
}
