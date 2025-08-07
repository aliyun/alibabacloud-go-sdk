// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAuditLogCollectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterAuditLogCollectorRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterAuditLogCollectorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterAuditLogCollectorRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterAuditLogCollectorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterAuditLogCollectorRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterAuditLogCollectorRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterAuditLogCollectorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAuditLogCollectorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAuditLogCollectorRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterAuditLogCollectorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterAuditLogCollectorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterAuditLogCollectorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterAuditLogCollectorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetDBClusterId(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetOwnerAccount(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetOwnerId(v int64) *DescribeDBClusterAuditLogCollectorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) Validate() error {
	return dara.Validate(s)
}
