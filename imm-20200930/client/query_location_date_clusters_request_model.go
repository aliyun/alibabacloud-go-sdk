// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocationDateClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v *Address) *QueryLocationDateClustersRequest
	GetAddress() *Address
	SetCreateTimeRange(v *TimeRange) *QueryLocationDateClustersRequest
	GetCreateTimeRange() *TimeRange
	SetCustomLabels(v string) *QueryLocationDateClustersRequest
	GetCustomLabels() *string
	SetDatasetName(v string) *QueryLocationDateClustersRequest
	GetDatasetName() *string
	SetLocationDateClusterEndTimeRange(v *TimeRange) *QueryLocationDateClustersRequest
	GetLocationDateClusterEndTimeRange() *TimeRange
	SetLocationDateClusterLevels(v []*string) *QueryLocationDateClustersRequest
	GetLocationDateClusterLevels() []*string
	SetLocationDateClusterStartTimeRange(v *TimeRange) *QueryLocationDateClustersRequest
	GetLocationDateClusterStartTimeRange() *TimeRange
	SetMaxResults(v int32) *QueryLocationDateClustersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryLocationDateClustersRequest
	GetNextToken() *string
	SetObjectId(v string) *QueryLocationDateClustersRequest
	GetObjectId() *string
	SetOrder(v string) *QueryLocationDateClustersRequest
	GetOrder() *string
	SetProjectName(v string) *QueryLocationDateClustersRequest
	GetProjectName() *string
	SetSort(v string) *QueryLocationDateClustersRequest
	GetSort() *string
	SetTitle(v string) *QueryLocationDateClustersRequest
	GetTitle() *string
	SetUpdateTimeRange(v *TimeRange) *QueryLocationDateClustersRequest
	GetUpdateTimeRange() *TimeRange
}

type QueryLocationDateClustersRequest struct {
	// The address information.
	Address *Address `json:"Address,omitempty" xml:"Address,omitempty"`
	// The time range during which the spatiotemporal clusters were generated.
	CreateTimeRange *TimeRange `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	// The custom labels.
	//
	// example:
	//
	// key=value
	CustomLabels *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset. For information about how to create a dataset, see [CreateDataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The time range during which the latest photo in a cluster was taken.
	LocationDateClusterEndTimeRange *TimeRange `json:"LocationDateClusterEndTimeRange,omitempty" xml:"LocationDateClusterEndTimeRange,omitempty"`
	// The container for the administrative division level of the spatiotemporal clusters to be queried.
	LocationDateClusterLevels []*string `json:"LocationDateClusterLevels,omitempty" xml:"LocationDateClusterLevels,omitempty" type:"Repeated"`
	// The time range during which the earliest photo in a cluster was taken.
	LocationDateClusterStartTimeRange *TimeRange `json:"LocationDateClusterStartTimeRange,omitempty" xml:"LocationDateClusterStartTimeRange,omitempty"`
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
	// The ID of the cluster that you want to query. Specify this parameter if you want to query a specific spatiotemporal cluster. Otherwise, leave this parameter empty to query spatiotemporal clusters that meet the specified conditions.
	//
	// example:
	//
	// location-date-cluster-71dd4f32-9597-4085-a2ab-3a7b0fd0aff9
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The order that you use to sort the query results.
	//
	// Valid values:
	//
	// 	- asc: ascending order. This is the default value.
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
	// The field that you use to sort the query results.
	//
	// Valid values:
	//
	// 	- LocationDateClusterEndTime: by the time at which the latest photo in a cluster was taken.
	//
	// 	- CreateTime: by the creation time of a spatiotemporal cluster.
	//
	// 	- UpdateTime: by the update time of a spatiotemporal cluster.
	//
	// 	- LocationDateClusterStartTime: by the time at which the earliest photo in a cluster was taken. This is the default value.
	//
	// example:
	//
	// LocationDateClusterStartTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The characters that are included in the titles of spatiotemporal clusters to be queried. Matches are found by using fuzzy matching.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time range during which the spatiotemporal clusters were updated.
	UpdateTimeRange *TimeRange `json:"UpdateTimeRange,omitempty" xml:"UpdateTimeRange,omitempty"`
}

func (s QueryLocationDateClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLocationDateClustersRequest) GoString() string {
	return s.String()
}

func (s *QueryLocationDateClustersRequest) GetAddress() *Address {
	return s.Address
}

func (s *QueryLocationDateClustersRequest) GetCreateTimeRange() *TimeRange {
	return s.CreateTimeRange
}

func (s *QueryLocationDateClustersRequest) GetCustomLabels() *string {
	return s.CustomLabels
}

func (s *QueryLocationDateClustersRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *QueryLocationDateClustersRequest) GetLocationDateClusterEndTimeRange() *TimeRange {
	return s.LocationDateClusterEndTimeRange
}

func (s *QueryLocationDateClustersRequest) GetLocationDateClusterLevels() []*string {
	return s.LocationDateClusterLevels
}

func (s *QueryLocationDateClustersRequest) GetLocationDateClusterStartTimeRange() *TimeRange {
	return s.LocationDateClusterStartTimeRange
}

func (s *QueryLocationDateClustersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryLocationDateClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryLocationDateClustersRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *QueryLocationDateClustersRequest) GetOrder() *string {
	return s.Order
}

func (s *QueryLocationDateClustersRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryLocationDateClustersRequest) GetSort() *string {
	return s.Sort
}

func (s *QueryLocationDateClustersRequest) GetTitle() *string {
	return s.Title
}

func (s *QueryLocationDateClustersRequest) GetUpdateTimeRange() *TimeRange {
	return s.UpdateTimeRange
}

func (s *QueryLocationDateClustersRequest) SetAddress(v *Address) *QueryLocationDateClustersRequest {
	s.Address = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetCreateTimeRange(v *TimeRange) *QueryLocationDateClustersRequest {
	s.CreateTimeRange = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetCustomLabels(v string) *QueryLocationDateClustersRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetDatasetName(v string) *QueryLocationDateClustersRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetLocationDateClusterEndTimeRange(v *TimeRange) *QueryLocationDateClustersRequest {
	s.LocationDateClusterEndTimeRange = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetLocationDateClusterLevels(v []*string) *QueryLocationDateClustersRequest {
	s.LocationDateClusterLevels = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetLocationDateClusterStartTimeRange(v *TimeRange) *QueryLocationDateClustersRequest {
	s.LocationDateClusterStartTimeRange = v
	return s
}

func (s *QueryLocationDateClustersRequest) SetMaxResults(v int32) *QueryLocationDateClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetNextToken(v string) *QueryLocationDateClustersRequest {
	s.NextToken = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetObjectId(v string) *QueryLocationDateClustersRequest {
	s.ObjectId = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetOrder(v string) *QueryLocationDateClustersRequest {
	s.Order = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetProjectName(v string) *QueryLocationDateClustersRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetSort(v string) *QueryLocationDateClustersRequest {
	s.Sort = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetTitle(v string) *QueryLocationDateClustersRequest {
	s.Title = &v
	return s
}

func (s *QueryLocationDateClustersRequest) SetUpdateTimeRange(v *TimeRange) *QueryLocationDateClustersRequest {
	s.UpdateTimeRange = v
	return s
}

func (s *QueryLocationDateClustersRequest) Validate() error {
	if s.Address != nil {
		if err := s.Address.Validate(); err != nil {
			return err
		}
	}
	if s.CreateTimeRange != nil {
		if err := s.CreateTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.LocationDateClusterEndTimeRange != nil {
		if err := s.LocationDateClusterEndTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.LocationDateClusterStartTimeRange != nil {
		if err := s.LocationDateClusterStartTimeRange.Validate(); err != nil {
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
