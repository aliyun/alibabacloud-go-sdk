// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePerformanceViewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribePerformanceViewAttributeResponseBody
	GetAccessDeniedDetail() *string
	SetCreateFromViewType(v string) *DescribePerformanceViewAttributeResponseBody
	GetCreateFromViewType() *string
	SetDBClusterId(v string) *DescribePerformanceViewAttributeResponseBody
	GetDBClusterId() *string
	SetFillOriginViewKeys(v bool) *DescribePerformanceViewAttributeResponseBody
	GetFillOriginViewKeys() *bool
	SetRequestId(v string) *DescribePerformanceViewAttributeResponseBody
	GetRequestId() *string
	SetViewDetail(v *DescribePerformanceViewAttributeResponseBodyViewDetail) *DescribePerformanceViewAttributeResponseBody
	GetViewDetail() *DescribePerformanceViewAttributeResponseBodyViewDetail
	SetViewName(v string) *DescribePerformanceViewAttributeResponseBody
	GetViewName() *string
}

type DescribePerformanceViewAttributeResponseBody struct {
	// The details about the access denial.
	//
	// >  This parameter is returned only if Resource Access Management (RAM) permission verification failed.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1*****************7",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcKwyhk62IeYly4hQ+5IpXjkh1GQXuDRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "2***************9",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The type of the view.
	//
	// example:
	//
	// Basic
	CreateFromViewType *string `json:"CreateFromViewType,omitempty" xml:"CreateFromViewType,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// example:
	//
	// amv-bp198m028ih55xxxx
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
	FillOriginViewKeys *bool `json:"FillOriginViewKeys,omitempty" xml:"FillOriginViewKeys,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E031AABF-BD56-5966-A063-4283EF18DB45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the monitoring view.
	ViewDetail *DescribePerformanceViewAttributeResponseBodyViewDetail `json:"ViewDetail,omitempty" xml:"ViewDetail,omitempty" type:"Struct"`
	// The name of the view.
	//
	// example:
	//
	// Basic
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s DescribePerformanceViewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewAttributeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribePerformanceViewAttributeResponseBody) GetCreateFromViewType() *string {
	return s.CreateFromViewType
}

func (s *DescribePerformanceViewAttributeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePerformanceViewAttributeResponseBody) GetFillOriginViewKeys() *bool {
	return s.FillOriginViewKeys
}

func (s *DescribePerformanceViewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePerformanceViewAttributeResponseBody) GetViewDetail() *DescribePerformanceViewAttributeResponseBodyViewDetail {
	return s.ViewDetail
}

func (s *DescribePerformanceViewAttributeResponseBody) GetViewName() *string {
	return s.ViewName
}

func (s *DescribePerformanceViewAttributeResponseBody) SetAccessDeniedDetail(v string) *DescribePerformanceViewAttributeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBody) SetCreateFromViewType(v string) *DescribePerformanceViewAttributeResponseBody {
	s.CreateFromViewType = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBody) SetDBClusterId(v string) *DescribePerformanceViewAttributeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBody) SetFillOriginViewKeys(v bool) *DescribePerformanceViewAttributeResponseBody {
	s.FillOriginViewKeys = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBody) SetRequestId(v string) *DescribePerformanceViewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBody) SetViewDetail(v *DescribePerformanceViewAttributeResponseBodyViewDetail) *DescribePerformanceViewAttributeResponseBody {
	s.ViewDetail = v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBody) SetViewName(v string) *DescribePerformanceViewAttributeResponseBody {
	s.ViewName = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePerformanceViewAttributeResponseBodyViewDetail struct {
	// The metric category.
	Categories []*DescribePerformanceViewAttributeResponseBodyViewDetailCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
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

func (s DescribePerformanceViewAttributeResponseBodyViewDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewAttributeResponseBodyViewDetail) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetail) GetCategories() []*DescribePerformanceViewAttributeResponseBodyViewDetailCategories {
	return s.Categories
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetail) GetChartLinked() *bool {
	return s.ChartLinked
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetail) GetChartsPerLine() *int32 {
	return s.ChartsPerLine
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetail) SetCategories(v []*DescribePerformanceViewAttributeResponseBodyViewDetailCategories) *DescribePerformanceViewAttributeResponseBodyViewDetail {
	s.Categories = v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetail) SetChartLinked(v bool) *DescribePerformanceViewAttributeResponseBodyViewDetail {
	s.ChartLinked = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetail) SetChartsPerLine(v int32) *DescribePerformanceViewAttributeResponseBodyViewDetail {
	s.ChartsPerLine = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePerformanceViewAttributeResponseBodyViewDetailCategories struct {
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
	Keys []*DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s DescribePerformanceViewAttributeResponseBodyViewDetailCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewAttributeResponseBodyViewDetailCategories) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategories) GetCategory() *string {
	return s.Category
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategories) GetKeys() []*DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys {
	return s.Keys
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategories) SetCategory(v string) *DescribePerformanceViewAttributeResponseBodyViewDetailCategories {
	s.Category = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategories) SetKeys(v []*DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) *DescribePerformanceViewAttributeResponseBodyViewDetailCategories {
	s.Keys = v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategories) Validate() error {
	return dara.Validate(s)
}

type DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys struct {
	// Indicates whether the multi-cluster feature is enabled. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	EnableAutoMc *bool `json:"EnableAutoMc,omitempty" xml:"EnableAutoMc,omitempty"`
	// The database engine of the cluster. Valid values:
	//
	// 	- AnalyticDB
	Engine []*string `json:"Engine,omitempty" xml:"Engine,omitempty" type:"Repeated"`
	// The type of the resource group. Valid values:
	//
	// 	- **Interactive**
	//
	// 	- **Job**
	//
	// >  For more information about resource groups, see [Resource group overview](https://help.aliyun.com/document_detail/428610.html).
	GroupType []*string `json:"GroupType,omitempty" xml:"GroupType,omitempty" type:"Repeated"`
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

func (s DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) GoString() string {
	return s.String()
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) GetEnableAutoMc() *bool {
	return s.EnableAutoMc
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) GetEngine() []*string {
	return s.Engine
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) GetGroupType() []*string {
	return s.GroupType
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) GetKeyName() *string {
	return s.KeyName
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) GetSelected() *bool {
	return s.Selected
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) SetEnableAutoMc(v bool) *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys {
	s.EnableAutoMc = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) SetEngine(v []*string) *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys {
	s.Engine = v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) SetGroupType(v []*string) *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys {
	s.GroupType = v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) SetKeyName(v string) *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys {
	s.KeyName = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) SetSelected(v bool) *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys {
	s.Selected = &v
	return s
}

func (s *DescribePerformanceViewAttributeResponseBodyViewDetailCategoriesKeys) Validate() error {
	return dara.Validate(s)
}
