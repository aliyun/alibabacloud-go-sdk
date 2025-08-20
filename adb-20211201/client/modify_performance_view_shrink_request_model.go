// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPerformanceViewShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyPerformanceViewShrinkRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyPerformanceViewShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPerformanceViewShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyPerformanceViewShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyPerformanceViewShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPerformanceViewShrinkRequest
	GetResourceOwnerId() *int64
	SetViewDetailShrink(v string) *ModifyPerformanceViewShrinkRequest
	GetViewDetailShrink() *string
	SetViewName(v string) *ModifyPerformanceViewShrinkRequest
	GetViewName() *string
}

type ModifyPerformanceViewShrinkRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1ub9grke1****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new information about the monitoring view.
	//
	// This parameter is required.
	ViewDetailShrink *string `json:"ViewDetail,omitempty" xml:"ViewDetail,omitempty"`
	// The name of the monitoring view.
	//
	// This parameter is required.
	//
	// example:
	//
	// Basic
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s ModifyPerformanceViewShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPerformanceViewShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPerformanceViewShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyPerformanceViewShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPerformanceViewShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPerformanceViewShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPerformanceViewShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPerformanceViewShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPerformanceViewShrinkRequest) GetViewDetailShrink() *string {
	return s.ViewDetailShrink
}

func (s *ModifyPerformanceViewShrinkRequest) GetViewName() *string {
	return s.ViewName
}

func (s *ModifyPerformanceViewShrinkRequest) SetDBClusterId(v string) *ModifyPerformanceViewShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyPerformanceViewShrinkRequest) SetOwnerAccount(v string) *ModifyPerformanceViewShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPerformanceViewShrinkRequest) SetOwnerId(v int64) *ModifyPerformanceViewShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPerformanceViewShrinkRequest) SetRegionId(v string) *ModifyPerformanceViewShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPerformanceViewShrinkRequest) SetResourceOwnerAccount(v string) *ModifyPerformanceViewShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPerformanceViewShrinkRequest) SetResourceOwnerId(v int64) *ModifyPerformanceViewShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPerformanceViewShrinkRequest) SetViewDetailShrink(v string) *ModifyPerformanceViewShrinkRequest {
	s.ViewDetailShrink = &v
	return s
}

func (s *ModifyPerformanceViewShrinkRequest) SetViewName(v string) *ModifyPerformanceViewShrinkRequest {
	s.ViewName = &v
	return s
}

func (s *ModifyPerformanceViewShrinkRequest) Validate() error {
	return dara.Validate(s)
}
