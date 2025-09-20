// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFigureClustersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeRangeShrink(v string) *QueryFigureClustersShrinkRequest
	GetCreateTimeRangeShrink() *string
	SetCustomLabels(v string) *QueryFigureClustersShrinkRequest
	GetCustomLabels() *string
	SetDatasetName(v string) *QueryFigureClustersShrinkRequest
	GetDatasetName() *string
	SetMaxResults(v int64) *QueryFigureClustersShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *QueryFigureClustersShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *QueryFigureClustersShrinkRequest
	GetOrder() *string
	SetProjectName(v string) *QueryFigureClustersShrinkRequest
	GetProjectName() *string
	SetSort(v string) *QueryFigureClustersShrinkRequest
	GetSort() *string
	SetUpdateTimeRangeShrink(v string) *QueryFigureClustersShrinkRequest
	GetUpdateTimeRangeShrink() *string
	SetWithTotalCount(v bool) *QueryFigureClustersShrinkRequest
	GetWithTotalCount() *bool
}

type QueryFigureClustersShrinkRequest struct {
	// The time range within which the face group was created.
	CreateTimeRangeShrink *string `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	// The custom labels, which can be used as query conditions.
	//
	// example:
	//
	// key=value
	CustomLabels *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset. You can obtain the name of the dataset from the response of the [CreateDataset](https://help.aliyun.com/document_detail/478160.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The maximum number of entries to return. Valid values: 0 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Default value: asc.
	//
	// Valid values:
	//
	// 	- asc: ascending order.
	//
	// 	- desc: descending order.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The sort field. If you leave this parameter empty, the group ID is used as the sort field.
	//
	// Valid values:
	//
	// 	- ImageCount: the number of images.
	//
	// 	- VideoCount: the number of videos.
	//
	// 	- ProjectName: the name of the project.
	//
	// 	- DatasetName: the name of the dataset.
	//
	// 	- CreateTime: the point in time when the group is created.
	//
	// 	- UpdateTime: the most recent point in time when the group is updated.
	//
	// 	- Gender: the gender.
	//
	// 	- FaceCount: the number of faces.
	//
	// 	- GroupName: the name of the group.
	//
	// example:
	//
	// ImageCount
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The time range within which the face group was last updated.
	UpdateTimeRangeShrink *string `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
	// Specifies whether to return the total number of face groups that match the current query conditions. Default value: false.
	//
	// example:
	//
	// false
	WithTotalCount *bool `json:"WithTotalCount,omitempty" xml:"WithTotalCount,omitempty"`
}

func (s QueryFigureClustersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFigureClustersShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryFigureClustersShrinkRequest) GetCreateTimeRangeShrink() *string {
	return s.CreateTimeRangeShrink
}

func (s *QueryFigureClustersShrinkRequest) GetCustomLabels() *string {
	return s.CustomLabels
}

func (s *QueryFigureClustersShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *QueryFigureClustersShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *QueryFigureClustersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryFigureClustersShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *QueryFigureClustersShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryFigureClustersShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *QueryFigureClustersShrinkRequest) GetUpdateTimeRangeShrink() *string {
	return s.UpdateTimeRangeShrink
}

func (s *QueryFigureClustersShrinkRequest) GetWithTotalCount() *bool {
	return s.WithTotalCount
}

func (s *QueryFigureClustersShrinkRequest) SetCreateTimeRangeShrink(v string) *QueryFigureClustersShrinkRequest {
	s.CreateTimeRangeShrink = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetCustomLabels(v string) *QueryFigureClustersShrinkRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetDatasetName(v string) *QueryFigureClustersShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetMaxResults(v int64) *QueryFigureClustersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetNextToken(v string) *QueryFigureClustersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetOrder(v string) *QueryFigureClustersShrinkRequest {
	s.Order = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetProjectName(v string) *QueryFigureClustersShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetSort(v string) *QueryFigureClustersShrinkRequest {
	s.Sort = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetUpdateTimeRangeShrink(v string) *QueryFigureClustersShrinkRequest {
	s.UpdateTimeRangeShrink = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) SetWithTotalCount(v bool) *QueryFigureClustersShrinkRequest {
	s.WithTotalCount = &v
	return s
}

func (s *QueryFigureClustersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
