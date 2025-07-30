// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeParametersRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *DescribeParametersRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeParametersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParametersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeParametersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeParametersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParametersRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeParametersRequest
	GetSecurityToken() *string
}

type DescribeParametersRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the node.
	//
	// > You can set this parameter to query the parameter settings of the specified node in a cluster instance.
	//
	// example:
	//
	// r-bp1xxxxxxxxxxxxx-db-0
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeParametersRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeParametersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParametersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParametersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParametersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParametersRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeParametersRequest) SetDBInstanceId(v string) *DescribeParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParametersRequest) SetNodeId(v string) *DescribeParametersRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeParametersRequest) SetOwnerAccount(v string) *DescribeParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParametersRequest) SetOwnerId(v int64) *DescribeParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParametersRequest) SetRegionId(v string) *DescribeParametersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParametersRequest) SetResourceOwnerAccount(v string) *DescribeParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParametersRequest) SetResourceOwnerId(v int64) *DescribeParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParametersRequest) SetSecurityToken(v string) *DescribeParametersRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParametersRequest) Validate() error {
	return dara.Validate(s)
}
