// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocationDateClustersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressShrink(v string) *QueryLocationDateClustersShrinkRequest
	GetAddressShrink() *string
	SetCreateTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest
	GetCreateTimeRangeShrink() *string
	SetCustomLabels(v string) *QueryLocationDateClustersShrinkRequest
	GetCustomLabels() *string
	SetDatasetName(v string) *QueryLocationDateClustersShrinkRequest
	GetDatasetName() *string
	SetLocationDateClusterEndTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest
	GetLocationDateClusterEndTimeRangeShrink() *string
	SetLocationDateClusterLevelsShrink(v string) *QueryLocationDateClustersShrinkRequest
	GetLocationDateClusterLevelsShrink() *string
	SetLocationDateClusterStartTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest
	GetLocationDateClusterStartTimeRangeShrink() *string
	SetMaxResults(v int32) *QueryLocationDateClustersShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryLocationDateClustersShrinkRequest
	GetNextToken() *string
	SetObjectId(v string) *QueryLocationDateClustersShrinkRequest
	GetObjectId() *string
	SetOrder(v string) *QueryLocationDateClustersShrinkRequest
	GetOrder() *string
	SetProjectName(v string) *QueryLocationDateClustersShrinkRequest
	GetProjectName() *string
	SetSort(v string) *QueryLocationDateClustersShrinkRequest
	GetSort() *string
	SetTitle(v string) *QueryLocationDateClustersShrinkRequest
	GetTitle() *string
	SetUpdateTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest
	GetUpdateTimeRangeShrink() *string
}

type QueryLocationDateClustersShrinkRequest struct {
	// The address information.
	AddressShrink *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The time range during which the spatiotemporal clusters were generated.
	CreateTimeRangeShrink *string `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	// The custom labels, which can be used as query conditions.
	//
	// example:
	//
	// key=value
	CustomLabels *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset. For more information, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The time range during which the latest photo in a cluster was taken.
	LocationDateClusterEndTimeRangeShrink *string `json:"LocationDateClusterEndTimeRange,omitempty" xml:"LocationDateClusterEndTimeRange,omitempty"`
	// The administrative level of the spatiotemporal clustering groups to be queried.
	LocationDateClusterLevelsShrink *string `json:"LocationDateClusterLevels,omitempty" xml:"LocationDateClusterLevels,omitempty"`
	// The time range during which the earliest photo in a cluster was taken.
	LocationDateClusterStartTimeRangeShrink *string `json:"LocationDateClusterStartTimeRange,omitempty" xml:"LocationDateClusterStartTimeRange,omitempty"`
	// The number of entries per page. Valid values: [1,100]. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// MzQNjmY2MzYxNhNjk2ZNjEu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the group that you want to query. Specify this parameter if you want to obtain the information about a specific spatiotemporal clustering group. Otherwise, leave this parameter empty and use other parameters to query the groups that meet the matching conditions.
	//
	// example:
	//
	// location-date-cluster-71dd4f32-9597-4085-a2ab-3a7b0fd0aff9
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The sorting order.
	//
	// Default value: asc. Valid values:
	//
	// 	- asc: ascending order.
	//
	// 	- desc: descending order.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The condition by which the results are sorted.
	//
	// Valid values:
	//
	// 	- LocationDateClusterEndTime: by the end time of the spatiotemporal clustering groups.
	//
	// 	- CreateTime: by the creation time of the spatiotemporal clustering groups.
	//
	// 	- UpdateTime: by the update time of the spatiotemporal clustering groups.
	//
	// 	- LocationDateClusterStartTime: by the start time of the spatiotemporal clustering groups. This is the default value.
	//
	// example:
	//
	// LocationDateClusterStartTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The title of spatiotemporal clustering. Fuzzy matching is performed.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time range during which the spatiotemporal clusters were updated.
	UpdateTimeRangeShrink *string `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
}

func (s QueryLocationDateClustersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLocationDateClustersShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryLocationDateClustersShrinkRequest) GetAddressShrink() *string {
	return s.AddressShrink
}

func (s *QueryLocationDateClustersShrinkRequest) GetCreateTimeRangeShrink() *string {
	return s.CreateTimeRangeShrink
}

func (s *QueryLocationDateClustersShrinkRequest) GetCustomLabels() *string {
	return s.CustomLabels
}

func (s *QueryLocationDateClustersShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *QueryLocationDateClustersShrinkRequest) GetLocationDateClusterEndTimeRangeShrink() *string {
	return s.LocationDateClusterEndTimeRangeShrink
}

func (s *QueryLocationDateClustersShrinkRequest) GetLocationDateClusterLevelsShrink() *string {
	return s.LocationDateClusterLevelsShrink
}

func (s *QueryLocationDateClustersShrinkRequest) GetLocationDateClusterStartTimeRangeShrink() *string {
	return s.LocationDateClusterStartTimeRangeShrink
}

func (s *QueryLocationDateClustersShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryLocationDateClustersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryLocationDateClustersShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *QueryLocationDateClustersShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *QueryLocationDateClustersShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryLocationDateClustersShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *QueryLocationDateClustersShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *QueryLocationDateClustersShrinkRequest) GetUpdateTimeRangeShrink() *string {
	return s.UpdateTimeRangeShrink
}

func (s *QueryLocationDateClustersShrinkRequest) SetAddressShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.AddressShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetCreateTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.CreateTimeRangeShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetCustomLabels(v string) *QueryLocationDateClustersShrinkRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetDatasetName(v string) *QueryLocationDateClustersShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetLocationDateClusterEndTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.LocationDateClusterEndTimeRangeShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetLocationDateClusterLevelsShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.LocationDateClusterLevelsShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetLocationDateClusterStartTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.LocationDateClusterStartTimeRangeShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetMaxResults(v int32) *QueryLocationDateClustersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetNextToken(v string) *QueryLocationDateClustersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetObjectId(v string) *QueryLocationDateClustersShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetOrder(v string) *QueryLocationDateClustersShrinkRequest {
	s.Order = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetProjectName(v string) *QueryLocationDateClustersShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetSort(v string) *QueryLocationDateClustersShrinkRequest {
	s.Sort = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetTitle(v string) *QueryLocationDateClustersShrinkRequest {
	s.Title = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) SetUpdateTimeRangeShrink(v string) *QueryLocationDateClustersShrinkRequest {
	s.UpdateTimeRangeShrink = &v
	return s
}

func (s *QueryLocationDateClustersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
