// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterEndpointsRequest
	GetDBClusterId() *string
	SetDBEndpointId(v string) *DescribeDBClusterEndpointsRequest
	GetDBEndpointId() *string
	SetDescribeType(v string) *DescribeDBClusterEndpointsRequest
	GetDescribeType() *string
	SetOwnerAccount(v string) *DescribeDBClusterEndpointsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterEndpointsRequest
	GetOwnerId() *int64
	SetPolarFsInstanceId(v string) *DescribeDBClusterEndpointsRequest
	GetPolarFsInstanceId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterEndpointsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterEndpointsRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterEndpointsRequest struct {
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of the clusters that belong to your Alibaba Cloud account, such as cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// pe-*************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// example:
	//
	// AI
	DescribeType *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// pfs-test*****
	PolarFsInstanceId    *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterEndpointsRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribeDBClusterEndpointsRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeDBClusterEndpointsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterEndpointsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterEndpointsRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribeDBClusterEndpointsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterEndpointsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterEndpointsRequest) SetDBClusterId(v string) *DescribeDBClusterEndpointsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetDBEndpointId(v string) *DescribeDBClusterEndpointsRequest {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetDescribeType(v string) *DescribeDBClusterEndpointsRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetOwnerAccount(v string) *DescribeDBClusterEndpointsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetOwnerId(v int64) *DescribeDBClusterEndpointsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetPolarFsInstanceId(v string) *DescribeDBClusterEndpointsRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterEndpointsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetResourceOwnerId(v int64) *DescribeDBClusterEndpointsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
