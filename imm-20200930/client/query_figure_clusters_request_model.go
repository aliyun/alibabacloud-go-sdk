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
	// The time range within which the face group was created.
	CreateTimeRange *TimeRange `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
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
	UpdateTimeRange *TimeRange `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
	// Specifies whether to return the total number of face groups that match the current query conditions. Default value: false.
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
