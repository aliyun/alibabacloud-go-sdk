// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePerformanceViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateFromViewType(v string) *CreatePerformanceViewRequest
	GetCreateFromViewType() *string
	SetDBClusterId(v string) *CreatePerformanceViewRequest
	GetDBClusterId() *string
	SetFillOriginViewKeys(v bool) *CreatePerformanceViewRequest
	GetFillOriginViewKeys() *bool
	SetOwnerAccount(v string) *CreatePerformanceViewRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePerformanceViewRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreatePerformanceViewRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreatePerformanceViewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePerformanceViewRequest
	GetResourceOwnerId() *int64
	SetViewDetail(v *CreatePerformanceViewRequestViewDetail) *CreatePerformanceViewRequest
	GetViewDetail() *CreatePerformanceViewRequestViewDetail
	SetViewName(v string) *CreatePerformanceViewRequest
	GetViewName() *string
}

type CreatePerformanceViewRequest struct {
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
	ViewDetail *CreatePerformanceViewRequestViewDetail `json:"ViewDetail,omitempty" xml:"ViewDetail,omitempty" type:"Struct"`
	// The name of the view.
	//
	// This parameter is required.
	//
	// example:
	//
	// viewname
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s CreatePerformanceViewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePerformanceViewRequest) GoString() string {
	return s.String()
}

func (s *CreatePerformanceViewRequest) GetCreateFromViewType() *string {
	return s.CreateFromViewType
}

func (s *CreatePerformanceViewRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreatePerformanceViewRequest) GetFillOriginViewKeys() *bool {
	return s.FillOriginViewKeys
}

func (s *CreatePerformanceViewRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePerformanceViewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePerformanceViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePerformanceViewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePerformanceViewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePerformanceViewRequest) GetViewDetail() *CreatePerformanceViewRequestViewDetail {
	return s.ViewDetail
}

func (s *CreatePerformanceViewRequest) GetViewName() *string {
	return s.ViewName
}

func (s *CreatePerformanceViewRequest) SetCreateFromViewType(v string) *CreatePerformanceViewRequest {
	s.CreateFromViewType = &v
	return s
}

func (s *CreatePerformanceViewRequest) SetDBClusterId(v string) *CreatePerformanceViewRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreatePerformanceViewRequest) SetFillOriginViewKeys(v bool) *CreatePerformanceViewRequest {
	s.FillOriginViewKeys = &v
	return s
}

func (s *CreatePerformanceViewRequest) SetOwnerAccount(v string) *CreatePerformanceViewRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePerformanceViewRequest) SetOwnerId(v int64) *CreatePerformanceViewRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePerformanceViewRequest) SetRegionId(v string) *CreatePerformanceViewRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePerformanceViewRequest) SetResourceOwnerAccount(v string) *CreatePerformanceViewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePerformanceViewRequest) SetResourceOwnerId(v int64) *CreatePerformanceViewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePerformanceViewRequest) SetViewDetail(v *CreatePerformanceViewRequestViewDetail) *CreatePerformanceViewRequest {
	s.ViewDetail = v
	return s
}

func (s *CreatePerformanceViewRequest) SetViewName(v string) *CreatePerformanceViewRequest {
	s.ViewName = &v
	return s
}

func (s *CreatePerformanceViewRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePerformanceViewRequestViewDetail struct {
	// The metric categories.
	Categories []*CreatePerformanceViewRequestViewDetailCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// Specifies whether to enable the filter interaction feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ChartLinked *bool `json:"ChartLinked,omitempty" xml:"ChartLinked,omitempty"`
	// The number of charts to display in each row.
	//
	// example:
	//
	// 2
	ChartsPerLine *int32 `json:"ChartsPerLine,omitempty" xml:"ChartsPerLine,omitempty"`
}

func (s CreatePerformanceViewRequestViewDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePerformanceViewRequestViewDetail) GoString() string {
	return s.String()
}

func (s *CreatePerformanceViewRequestViewDetail) GetCategories() []*CreatePerformanceViewRequestViewDetailCategories {
	return s.Categories
}

func (s *CreatePerformanceViewRequestViewDetail) GetChartLinked() *bool {
	return s.ChartLinked
}

func (s *CreatePerformanceViewRequestViewDetail) GetChartsPerLine() *int32 {
	return s.ChartsPerLine
}

func (s *CreatePerformanceViewRequestViewDetail) SetCategories(v []*CreatePerformanceViewRequestViewDetailCategories) *CreatePerformanceViewRequestViewDetail {
	s.Categories = v
	return s
}

func (s *CreatePerformanceViewRequestViewDetail) SetChartLinked(v bool) *CreatePerformanceViewRequestViewDetail {
	s.ChartLinked = &v
	return s
}

func (s *CreatePerformanceViewRequestViewDetail) SetChartsPerLine(v int32) *CreatePerformanceViewRequestViewDetail {
	s.ChartsPerLine = &v
	return s
}

func (s *CreatePerformanceViewRequestViewDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePerformanceViewRequestViewDetailCategories struct {
	// The name of the metric category. Valid values:
	//
	// 	- **Node**
	//
	// 	- **DiskData**
	//
	// 	- **WorkLoad**
	//
	// 	- **ResourceGroup**
	//
	// example:
	//
	// Node
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The metrics.
	Keys []*CreatePerformanceViewRequestViewDetailCategoriesKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s CreatePerformanceViewRequestViewDetailCategories) String() string {
	return dara.Prettify(s)
}

func (s CreatePerformanceViewRequestViewDetailCategories) GoString() string {
	return s.String()
}

func (s *CreatePerformanceViewRequestViewDetailCategories) GetCategory() *string {
	return s.Category
}

func (s *CreatePerformanceViewRequestViewDetailCategories) GetKeys() []*CreatePerformanceViewRequestViewDetailCategoriesKeys {
	return s.Keys
}

func (s *CreatePerformanceViewRequestViewDetailCategories) SetCategory(v string) *CreatePerformanceViewRequestViewDetailCategories {
	s.Category = &v
	return s
}

func (s *CreatePerformanceViewRequestViewDetailCategories) SetKeys(v []*CreatePerformanceViewRequestViewDetailCategoriesKeys) *CreatePerformanceViewRequestViewDetailCategories {
	s.Keys = v
	return s
}

func (s *CreatePerformanceViewRequestViewDetailCategories) Validate() error {
	return dara.Validate(s)
}

type CreatePerformanceViewRequestViewDetailCategoriesKeys struct {
	// The name of the metric.
	//
	// example:
	//
	// AnalyticDB_CPU
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// Specifies whether to select the metric. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s CreatePerformanceViewRequestViewDetailCategoriesKeys) String() string {
	return dara.Prettify(s)
}

func (s CreatePerformanceViewRequestViewDetailCategoriesKeys) GoString() string {
	return s.String()
}

func (s *CreatePerformanceViewRequestViewDetailCategoriesKeys) GetKeyName() *string {
	return s.KeyName
}

func (s *CreatePerformanceViewRequestViewDetailCategoriesKeys) GetSelected() *bool {
	return s.Selected
}

func (s *CreatePerformanceViewRequestViewDetailCategoriesKeys) SetKeyName(v string) *CreatePerformanceViewRequestViewDetailCategoriesKeys {
	s.KeyName = &v
	return s
}

func (s *CreatePerformanceViewRequestViewDetailCategoriesKeys) SetSelected(v bool) *CreatePerformanceViewRequestViewDetailCategoriesKeys {
	s.Selected = &v
	return s
}

func (s *CreatePerformanceViewRequestViewDetailCategoriesKeys) Validate() error {
	return dara.Validate(s)
}
