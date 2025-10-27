// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTaskInfoRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeTaskInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTaskInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTaskInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTaskInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTaskInfoRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int32) *DescribeTaskInfoRequest
	GetTaskId() *int32
}

type DescribeTaskInfoRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hongkong
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 225685759
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTaskInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTaskInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTaskInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTaskInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTaskInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTaskInfoRequest) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeTaskInfoRequest) SetDBClusterId(v string) *DescribeTaskInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetOwnerAccount(v string) *DescribeTaskInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetOwnerId(v int64) *DescribeTaskInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetRegionId(v string) *DescribeTaskInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetResourceOwnerAccount(v string) *DescribeTaskInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetResourceOwnerId(v int64) *DescribeTaskInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetTaskId(v int32) *DescribeTaskInfoRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
