// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePerformanceViewShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateFromViewType(v string) *CreatePerformanceViewShrinkRequest
	GetCreateFromViewType() *string
	SetDBClusterId(v string) *CreatePerformanceViewShrinkRequest
	GetDBClusterId() *string
	SetFillOriginViewKeys(v bool) *CreatePerformanceViewShrinkRequest
	GetFillOriginViewKeys() *bool
	SetOwnerAccount(v string) *CreatePerformanceViewShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePerformanceViewShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreatePerformanceViewShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreatePerformanceViewShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePerformanceViewShrinkRequest
	GetResourceOwnerId() *int64
	SetViewDetailShrink(v string) *CreatePerformanceViewShrinkRequest
	GetViewDetailShrink() *string
	SetViewName(v string) *CreatePerformanceViewShrinkRequest
	GetViewName() *string
}

type CreatePerformanceViewShrinkRequest struct {
	// The type of the view.
	//
	// example:
	//
	// Basic
	CreateFromViewType *string `json:"CreateFromViewType,omitempty" xml:"CreateFromViewType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to populate the names of the metrics in the original monitoring view when you view the monitoring view. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FillOriginViewKeys *bool   `json:"FillOriginViewKeys,omitempty" xml:"FillOriginViewKeys,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the monitoring view.
	//
	// This parameter is required.
	ViewDetailShrink *string `json:"ViewDetail,omitempty" xml:"ViewDetail,omitempty"`
	// The name of the view.
	//
	// This parameter is required.
	//
	// example:
	//
	// viewname
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s CreatePerformanceViewShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePerformanceViewShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePerformanceViewShrinkRequest) GetCreateFromViewType() *string {
	return s.CreateFromViewType
}

func (s *CreatePerformanceViewShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreatePerformanceViewShrinkRequest) GetFillOriginViewKeys() *bool {
	return s.FillOriginViewKeys
}

func (s *CreatePerformanceViewShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePerformanceViewShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePerformanceViewShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePerformanceViewShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePerformanceViewShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePerformanceViewShrinkRequest) GetViewDetailShrink() *string {
	return s.ViewDetailShrink
}

func (s *CreatePerformanceViewShrinkRequest) GetViewName() *string {
	return s.ViewName
}

func (s *CreatePerformanceViewShrinkRequest) SetCreateFromViewType(v string) *CreatePerformanceViewShrinkRequest {
	s.CreateFromViewType = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) SetDBClusterId(v string) *CreatePerformanceViewShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) SetFillOriginViewKeys(v bool) *CreatePerformanceViewShrinkRequest {
	s.FillOriginViewKeys = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) SetOwnerAccount(v string) *CreatePerformanceViewShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) SetOwnerId(v int64) *CreatePerformanceViewShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) SetRegionId(v string) *CreatePerformanceViewShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) SetResourceOwnerAccount(v string) *CreatePerformanceViewShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) SetResourceOwnerId(v int64) *CreatePerformanceViewShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) SetViewDetailShrink(v string) *CreatePerformanceViewShrinkRequest {
	s.ViewDetailShrink = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) SetViewName(v string) *CreatePerformanceViewShrinkRequest {
	s.ViewName = &v
	return s
}

func (s *CreatePerformanceViewShrinkRequest) Validate() error {
	return dara.Validate(s)
}
