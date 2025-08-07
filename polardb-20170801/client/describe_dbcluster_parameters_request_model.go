// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterParametersRequest
	GetDBClusterId() *string
	SetDescribeType(v string) *DescribeDBClusterParametersRequest
	GetDescribeType() *string
	SetOwnerAccount(v string) *DescribeDBClusterParametersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterParametersRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterParametersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterParametersRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterParametersRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**********
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The type of the parameter information to query. Valid values:
	//
	// 	- **Normal**: the information about the cluster parameters
	//
	// 	- **MigrationFromRDS**: a comparison of parameters between the source RDS instance and the destination PolarDB cluster
	//
	// example:
	//
	// Normal
	DescribeType         *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterParametersRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeDBClusterParametersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterParametersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterParametersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterParametersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterParametersRequest) SetDBClusterId(v string) *DescribeDBClusterParametersRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetDescribeType(v string) *DescribeDBClusterParametersRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetOwnerAccount(v string) *DescribeDBClusterParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetOwnerId(v int64) *DescribeDBClusterParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetResourceOwnerId(v int64) *DescribeDBClusterParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) Validate() error {
	return dara.Validate(s)
}
