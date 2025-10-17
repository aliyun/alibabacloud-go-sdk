// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterEndpointsZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterEndpointsZonalRequest
	GetDBClusterId() *string
	SetDBEndpointId(v string) *DescribeDBClusterEndpointsZonalRequest
	GetDBEndpointId() *string
	SetDescribeType(v string) *DescribeDBClusterEndpointsZonalRequest
	GetDescribeType() *string
	SetOwnerAccount(v string) *DescribeDBClusterEndpointsZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterEndpointsZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterEndpointsZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterEndpointsZonalRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterEndpointsZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// pe-*************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// example:
	//
	// AI
	DescribeType         *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterEndpointsZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsZonalRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterEndpointsZonalRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribeDBClusterEndpointsZonalRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeDBClusterEndpointsZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterEndpointsZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterEndpointsZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterEndpointsZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterEndpointsZonalRequest) SetDBClusterId(v string) *DescribeDBClusterEndpointsZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalRequest) SetDBEndpointId(v string) *DescribeDBClusterEndpointsZonalRequest {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalRequest) SetDescribeType(v string) *DescribeDBClusterEndpointsZonalRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalRequest) SetOwnerAccount(v string) *DescribeDBClusterEndpointsZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalRequest) SetOwnerId(v int64) *DescribeDBClusterEndpointsZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterEndpointsZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalRequest) SetResourceOwnerId(v int64) *DescribeDBClusterEndpointsZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalRequest) Validate() error {
	return dara.Validate(s)
}
