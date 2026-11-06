// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLExplorerPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSQLExplorerPolicyRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeSQLExplorerPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLExplorerPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeSQLExplorerPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLExplorerPolicyRequest
	GetResourceOwnerId() *int64
}

type DescribeSQLExplorerPolicyRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters under your account, such as the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSQLExplorerPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLExplorerPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLExplorerPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLExplorerPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLExplorerPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLExplorerPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLExplorerPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLExplorerPolicyRequest) SetDBInstanceId(v string) *DescribeSQLExplorerPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLExplorerPolicyRequest) SetOwnerAccount(v string) *DescribeSQLExplorerPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLExplorerPolicyRequest) SetOwnerId(v int64) *DescribeSQLExplorerPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLExplorerPolicyRequest) SetResourceOwnerAccount(v string) *DescribeSQLExplorerPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLExplorerPolicyRequest) SetResourceOwnerId(v int64) *DescribeSQLExplorerPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLExplorerPolicyRequest) Validate() error {
	return dara.Validate(s)
}
