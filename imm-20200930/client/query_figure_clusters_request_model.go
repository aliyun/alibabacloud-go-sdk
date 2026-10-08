// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFigureClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeRange(v *TimeRange) *QueryFigureClustersRequest
	GetCreateTimeRange() *TimeRange
	SetCustomLabels(v string) *QueryFigureClustersRequest
	GetCustomLabels() *string
	SetDatasetName(v string) *QueryFigureClustersRequest
	GetDatasetName() *string
	SetMaxResults(v int64) *QueryFigureClustersRequest
	GetMaxResults() *int64
	SetNextToken(v string) *QueryFigureClustersRequest
	GetNextToken() *string
	SetOrder(v string) *QueryFigureClustersRequest
	GetOrder() *string
	SetProjectName(v string) *QueryFigureClustersRequest
	GetProjectName() *string
	SetSort(v string) *QueryFigureClustersRequest
	GetSort() *string
	SetUpdateTimeRange(v *TimeRange) *QueryFigureClustersRequest
	GetUpdateTimeRange() *TimeRange
	SetWithTotalCount(v bool) *QueryFigureClustersRequest
	GetWithTotalCount() *bool
}

type QueryFigureClustersRequest struct {
	// The time range during which the face clusters were created.
	CreateTimeRange *TimeRange `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
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
	UpdateTimeRange *TimeRange `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
	// Specifies whether to return the total number of face clusters that meet the current query conditions. Default value: false, which indicates that the total number of clusters is not returned.
	//
	// example:
	//
	// false
	WithTotalCount *bool `json:"WithTotalCount,omitempty" xml:"WithTotalCount,omitempty"`
}

func (s QueryFigureClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFigureClustersRequest) GoString() string {
	return s.String()
}

func (s *QueryFigureClustersRequest) GetCreateTimeRange() *TimeRange {
	return s.CreateTimeRange
}

func (s *QueryFigureClustersRequest) GetCustomLabels() *string {
	return s.CustomLabels
}

func (s *QueryFigureClustersRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *QueryFigureClustersRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *QueryFigureClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryFigureClustersRequest) GetOrder() *string {
	return s.Order
}

func (s *QueryFigureClustersRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryFigureClustersRequest) GetSort() *string {
	return s.Sort
}

func (s *QueryFigureClustersRequest) GetUpdateTimeRange() *TimeRange {
	return s.UpdateTimeRange
}

func (s *QueryFigureClustersRequest) GetWithTotalCount() *bool {
	return s.WithTotalCount
}

func (s *QueryFigureClustersRequest) SetCreateTimeRange(v *TimeRange) *QueryFigureClustersRequest {
	s.CreateTimeRange = v
	return s
}

func (s *QueryFigureClustersRequest) SetCustomLabels(v string) *QueryFigureClustersRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryFigureClustersRequest) SetDatasetName(v string) *QueryFigureClustersRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryFigureClustersRequest) SetMaxResults(v int64) *QueryFigureClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryFigureClustersRequest) SetNextToken(v string) *QueryFigureClustersRequest {
	s.NextToken = &v
	return s
}

func (s *QueryFigureClustersRequest) SetOrder(v string) *QueryFigureClustersRequest {
	s.Order = &v
	return s
}

func (s *QueryFigureClustersRequest) SetProjectName(v string) *QueryFigureClustersRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryFigureClustersRequest) SetSort(v string) *QueryFigureClustersRequest {
	s.Sort = &v
	return s
}

func (s *QueryFigureClustersRequest) SetUpdateTimeRange(v *TimeRange) *QueryFigureClustersRequest {
	s.UpdateTimeRange = v
	return s
}

func (s *QueryFigureClustersRequest) SetWithTotalCount(v bool) *QueryFigureClustersRequest {
	s.WithTotalCount = &v
	return s
}

func (s *QueryFigureClustersRequest) Validate() error {
	if s.CreateTimeRange != nil {
		if err := s.CreateTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.UpdateTimeRange != nil {
		if err := s.UpdateTimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}
