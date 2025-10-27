// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAuditLogConfigRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAuditLogConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAuditLogConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAuditLogConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAuditLogConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAuditLogConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeAuditLogConfigRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-t4nj8619bz2w3****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAuditLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAuditLogConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAuditLogConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAuditLogConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAuditLogConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAuditLogConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAuditLogConfigRequest) SetDBClusterId(v string) *DescribeAuditLogConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetOwnerAccount(v string) *DescribeAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetOwnerId(v int64) *DescribeAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetRegionId(v string) *DescribeAuditLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetResourceOwnerId(v int64) *DescribeAuditLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
