// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationMaintenanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeActiveOperationMaintenanceConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActiveOperationMaintenanceConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeActiveOperationMaintenanceConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationMaintenanceConfigRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeActiveOperationMaintenanceConfigRequest
	GetSecurityToken() *string
}

type DescribeActiveOperationMaintenanceConfigRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeActiveOperationMaintenanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationMaintenanceConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) SetOwnerAccount(v string) *DescribeActiveOperationMaintenanceConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) SetOwnerId(v int64) *DescribeActiveOperationMaintenanceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationMaintenanceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationMaintenanceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) SetSecurityToken(v string) *DescribeActiveOperationMaintenanceConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationMaintenanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
