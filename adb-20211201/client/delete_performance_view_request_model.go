// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePerformanceViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeletePerformanceViewRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeletePerformanceViewRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeletePerformanceViewRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeletePerformanceViewRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeletePerformanceViewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePerformanceViewRequest
	GetResourceOwnerId() *int64
	SetViewName(v string) *DeletePerformanceViewRequest
	GetViewName() *string
}

type DeletePerformanceViewRequest struct {
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf6wjk5xxxxxxxxxx
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the supported regions and zones, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the monitoring view.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom-All metrics-2 columns-Linked
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s DeletePerformanceViewRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePerformanceViewRequest) GoString() string {
	return s.String()
}

func (s *DeletePerformanceViewRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeletePerformanceViewRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeletePerformanceViewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePerformanceViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePerformanceViewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePerformanceViewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePerformanceViewRequest) GetViewName() *string {
	return s.ViewName
}

func (s *DeletePerformanceViewRequest) SetDBClusterId(v string) *DeletePerformanceViewRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeletePerformanceViewRequest) SetOwnerAccount(v string) *DeletePerformanceViewRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePerformanceViewRequest) SetOwnerId(v int64) *DeletePerformanceViewRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePerformanceViewRequest) SetRegionId(v string) *DeletePerformanceViewRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePerformanceViewRequest) SetResourceOwnerAccount(v string) *DeletePerformanceViewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePerformanceViewRequest) SetResourceOwnerId(v int64) *DeletePerformanceViewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePerformanceViewRequest) SetViewName(v string) *DeletePerformanceViewRequest {
	s.ViewName = &v
	return s
}

func (s *DeletePerformanceViewRequest) Validate() error {
	return dara.Validate(s)
}
