// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPerformanceViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyPerformanceViewRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyPerformanceViewRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPerformanceViewRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyPerformanceViewRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyPerformanceViewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPerformanceViewRequest
	GetResourceOwnerId() *int64
	SetViewDetail(v *ModifyPerformanceViewRequestViewDetail) *ModifyPerformanceViewRequest
	GetViewDetail() *ModifyPerformanceViewRequestViewDetail
	SetViewName(v string) *ModifyPerformanceViewRequest
	GetViewName() *string
}

type ModifyPerformanceViewRequest struct {
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
	ViewDetail *ModifyPerformanceViewRequestViewDetail `json:"ViewDetail,omitempty" xml:"ViewDetail,omitempty" type:"Struct"`
	// The name of the monitoring view.
	//
	// This parameter is required.
	//
	// example:
	//
	// Basic
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s ModifyPerformanceViewRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPerformanceViewRequest) GoString() string {
	return s.String()
}

func (s *ModifyPerformanceViewRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyPerformanceViewRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPerformanceViewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPerformanceViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPerformanceViewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPerformanceViewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPerformanceViewRequest) GetViewDetail() *ModifyPerformanceViewRequestViewDetail {
	return s.ViewDetail
}

func (s *ModifyPerformanceViewRequest) GetViewName() *string {
	return s.ViewName
}

func (s *ModifyPerformanceViewRequest) SetDBClusterId(v string) *ModifyPerformanceViewRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyPerformanceViewRequest) SetOwnerAccount(v string) *ModifyPerformanceViewRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPerformanceViewRequest) SetOwnerId(v int64) *ModifyPerformanceViewRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPerformanceViewRequest) SetRegionId(v string) *ModifyPerformanceViewRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPerformanceViewRequest) SetResourceOwnerAccount(v string) *ModifyPerformanceViewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPerformanceViewRequest) SetResourceOwnerId(v int64) *ModifyPerformanceViewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPerformanceViewRequest) SetViewDetail(v *ModifyPerformanceViewRequestViewDetail) *ModifyPerformanceViewRequest {
	s.ViewDetail = v
	return s
}

func (s *ModifyPerformanceViewRequest) SetViewName(v string) *ModifyPerformanceViewRequest {
	s.ViewName = &v
	return s
}

func (s *ModifyPerformanceViewRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyPerformanceViewRequestViewDetail struct {
	// The metric categories.
	Categories []*ModifyPerformanceViewRequestViewDetailCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// Specifies whether to enable the filter interaction feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ChartLinked *bool `json:"ChartLinked,omitempty" xml:"ChartLinked,omitempty"`
	// The number of charts to display in each row.
	//
	// example:
	//
	// 3
	ChartsPerLine *int32 `json:"ChartsPerLine,omitempty" xml:"ChartsPerLine,omitempty"`
}

func (s ModifyPerformanceViewRequestViewDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyPerformanceViewRequestViewDetail) GoString() string {
	return s.String()
}

func (s *ModifyPerformanceViewRequestViewDetail) GetCategories() []*ModifyPerformanceViewRequestViewDetailCategories {
	return s.Categories
}

func (s *ModifyPerformanceViewRequestViewDetail) GetChartLinked() *bool {
	return s.ChartLinked
}

func (s *ModifyPerformanceViewRequestViewDetail) GetChartsPerLine() *int32 {
	return s.ChartsPerLine
}

func (s *ModifyPerformanceViewRequestViewDetail) SetCategories(v []*ModifyPerformanceViewRequestViewDetailCategories) *ModifyPerformanceViewRequestViewDetail {
	s.Categories = v
	return s
}

func (s *ModifyPerformanceViewRequestViewDetail) SetChartLinked(v bool) *ModifyPerformanceViewRequestViewDetail {
	s.ChartLinked = &v
	return s
}

func (s *ModifyPerformanceViewRequestViewDetail) SetChartsPerLine(v int32) *ModifyPerformanceViewRequestViewDetail {
	s.ChartsPerLine = &v
	return s
}

func (s *ModifyPerformanceViewRequestViewDetail) Validate() error {
	return dara.Validate(s)
}

type ModifyPerformanceViewRequestViewDetailCategories struct {
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
	Keys []*ModifyPerformanceViewRequestViewDetailCategoriesKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s ModifyPerformanceViewRequestViewDetailCategories) String() string {
	return dara.Prettify(s)
}

func (s ModifyPerformanceViewRequestViewDetailCategories) GoString() string {
	return s.String()
}

func (s *ModifyPerformanceViewRequestViewDetailCategories) GetCategory() *string {
	return s.Category
}

func (s *ModifyPerformanceViewRequestViewDetailCategories) GetKeys() []*ModifyPerformanceViewRequestViewDetailCategoriesKeys {
	return s.Keys
}

func (s *ModifyPerformanceViewRequestViewDetailCategories) SetCategory(v string) *ModifyPerformanceViewRequestViewDetailCategories {
	s.Category = &v
	return s
}

func (s *ModifyPerformanceViewRequestViewDetailCategories) SetKeys(v []*ModifyPerformanceViewRequestViewDetailCategoriesKeys) *ModifyPerformanceViewRequestViewDetailCategories {
	s.Keys = v
	return s
}

func (s *ModifyPerformanceViewRequestViewDetailCategories) Validate() error {
	return dara.Validate(s)
}

type ModifyPerformanceViewRequestViewDetailCategoriesKeys struct {
	// The name of the metric.
	//
	// example:
	//
	// AnalyticDB_CPU
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// Specifies whether to select the metric. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s ModifyPerformanceViewRequestViewDetailCategoriesKeys) String() string {
	return dara.Prettify(s)
}

func (s ModifyPerformanceViewRequestViewDetailCategoriesKeys) GoString() string {
	return s.String()
}

func (s *ModifyPerformanceViewRequestViewDetailCategoriesKeys) GetKeyName() *string {
	return s.KeyName
}

func (s *ModifyPerformanceViewRequestViewDetailCategoriesKeys) GetSelected() *bool {
	return s.Selected
}

func (s *ModifyPerformanceViewRequestViewDetailCategoriesKeys) SetKeyName(v string) *ModifyPerformanceViewRequestViewDetailCategoriesKeys {
	s.KeyName = &v
	return s
}

func (s *ModifyPerformanceViewRequestViewDetailCategoriesKeys) SetSelected(v bool) *ModifyPerformanceViewRequestViewDetailCategoriesKeys {
	s.Selected = &v
	return s
}

func (s *ModifyPerformanceViewRequestViewDetailCategoriesKeys) Validate() error {
	return dara.Validate(s)
}
