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
	// The time range during which the face clusters were created.
	CreateTimeRangeShrink *string `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	// The query conditions for custom labels.
	//
	// example:
	//
	// key=value
	CustomLabels *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset. For more information about how to obtain the dataset name, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The maximum number of data records to return in this call. Valid values: 0 to 100. If this parameter is not specified or is set to 0, the default value 100 is used.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. If this parameter is left empty or set to None, the query starts from the beginning.
	//
	// example:
	//
	// CAESEgoQCg4KCkltYWdlQ291bnQQARgBIr0ECgkABAAAAAAAAAAKrwQDKgIAADFTMzEzMDMyMzMzMjMxMzAzMDMyMzQzNjM3MzczOTMzMzQzYTY5NmQ2ZDJkNjk2ZDYxNjc2NTJkNzQ2NTczNzQyZDY4N2E2NDY1NzYyZDMyMzUzMjM0MzIzOTMzMzczMTJkMzY1NDZhNzk3MzU2Njk3MjM0M2E2OTZkNmQyZDc0NjU3Mzc0MmQ3MzY1NzQyZDYzMzYzNjY0MzY2NjYxMzQyZDM1MzMzODM3MmQzMTMxNjU2NjJkNjI2NTM5MzYyZDM5MzgzMDMzMzk2MjMwMzE2NDYzNjMzMjNhNjY2OTY3NzU3MjY1MmQ2MzZjNzU3Mzc0NjU3MjNhNDM2Yzc1NzM3NDY1NzIyZDYxNjUzOTY0MzQzMzMxNjEyZDM3MzQ2NTY2MmQzNDM5Mzc2MjJkMzg2MjMxMzUyZDM0MzUzOTM1MzYzNzYxMzQ2NDM2MzE2Ni5TMzEzMDMyMzMzMjMxMzAzMDMyMzQzNjM3MzczOTMzMzQzYTY5NmQ2ZDJkNjk2ZDYxNjc2NTJkNzQ2NTczNzQyZDY4N2E2NDY1NzYyZDMyMzUzMjM0MzIzOTMzMzczMTJkMzY1NDZhNzk3MzU2Njk3MjM0M2E2OTZkNmQyZDc0NjU3Mzc0MmQ3MzY1NzQyZDYzMzYzNjY0MzY2NjYxMzQyZDM1MzMzODM3MmQzMTMxNjU2NjJkNjI2NTM5MzYyZDM5MzgzMDMzMzk2MjMwMzE2NDYzNjM*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Default value: asc.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The field used for sorting. By default, this parameter is left empty, which indicates that the results are sorted by cluster ID.
	//
	// example:
	//
	// ImageCount
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The time range during which the face clusters were updated.
	UpdateTimeRangeShrink *string `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
	// Specifies whether to return the total number of face clusters that meet the current query conditions. Default value: false, which indicates that the total number of clusters is not returned.
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
