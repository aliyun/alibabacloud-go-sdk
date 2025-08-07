// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInitializeVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBInitializeVariableRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBInitializeVariableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInitializeVariableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInitializeVariableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInitializeVariableRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInitializeVariableRequest struct {
	// The ID of cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters that are deployed in a specified region, such as the cluster ID.
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

func (s DescribeDBInitializeVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInitializeVariableRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBInitializeVariableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInitializeVariableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInitializeVariableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInitializeVariableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInitializeVariableRequest) SetDBClusterId(v string) *DescribeDBInitializeVariableRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBInitializeVariableRequest) SetOwnerAccount(v string) *DescribeDBInitializeVariableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInitializeVariableRequest) SetOwnerId(v int64) *DescribeDBInitializeVariableRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInitializeVariableRequest) SetResourceOwnerAccount(v string) *DescribeDBInitializeVariableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInitializeVariableRequest) SetResourceOwnerId(v int64) *DescribeDBInitializeVariableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInitializeVariableRequest) Validate() error {
	return dara.Validate(s)
}
