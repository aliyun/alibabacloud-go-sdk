// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarSQLCollectorPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePolarSQLCollectorPolicyRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribePolarSQLCollectorPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePolarSQLCollectorPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribePolarSQLCollectorPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePolarSQLCollectorPolicyRequest
	GetResourceOwnerId() *int64
}

type DescribePolarSQLCollectorPolicyRequest struct {
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all the clusters for your account, such as the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePolarSQLCollectorPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarSQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarSQLCollectorPolicyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePolarSQLCollectorPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePolarSQLCollectorPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePolarSQLCollectorPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePolarSQLCollectorPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetDBClusterId(v string) *DescribePolarSQLCollectorPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetOwnerAccount(v string) *DescribePolarSQLCollectorPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetOwnerId(v int64) *DescribePolarSQLCollectorPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetResourceOwnerAccount(v string) *DescribePolarSQLCollectorPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetResourceOwnerId(v int64) *DescribePolarSQLCollectorPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyRequest) Validate() error {
	return dara.Validate(s)
}
