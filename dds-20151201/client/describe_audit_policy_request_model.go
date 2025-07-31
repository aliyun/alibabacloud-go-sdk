// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAuditPolicyRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeAuditPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAuditPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAuditPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAuditPolicyRequest
	GetResourceOwnerId() *int64
}

type DescribeAuditPolicyRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAuditPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAuditPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAuditPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAuditPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAuditPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAuditPolicyRequest) SetDBInstanceId(v string) *DescribeAuditPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAuditPolicyRequest) SetOwnerAccount(v string) *DescribeAuditPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditPolicyRequest) SetOwnerId(v int64) *DescribeAuditPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditPolicyRequest) SetResourceOwnerAccount(v string) *DescribeAuditPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditPolicyRequest) SetResourceOwnerId(v int64) *DescribeAuditPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditPolicyRequest) Validate() error {
	return dara.Validate(s)
}
