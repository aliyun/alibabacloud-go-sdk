// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeRange(v *TimeRange) *QueryStoriesRequest
	GetCreateTimeRange() *TimeRange
	SetCustomLabels(v string) *QueryStoriesRequest
	GetCustomLabels() *string
	SetDatasetName(v string) *QueryStoriesRequest
	GetDatasetName() *string
	SetFigureClusterIds(v []*string) *QueryStoriesRequest
	GetFigureClusterIds() []*string
	SetMaxResults(v int64) *QueryStoriesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *QueryStoriesRequest
	GetNextToken() *string
	SetObjectId(v string) *QueryStoriesRequest
	GetObjectId() *string
	SetOrder(v string) *QueryStoriesRequest
	GetOrder() *string
	SetProjectName(v string) *QueryStoriesRequest
	GetProjectName() *string
	SetSort(v string) *QueryStoriesRequest
	GetSort() *string
	SetStoryEndTimeRange(v *TimeRange) *QueryStoriesRequest
	GetStoryEndTimeRange() *TimeRange
	SetStoryName(v string) *QueryStoriesRequest
	GetStoryName() *string
	SetStoryStartTimeRange(v *TimeRange) *QueryStoriesRequest
	GetStoryStartTimeRange() *TimeRange
	SetStorySubType(v string) *QueryStoriesRequest
	GetStorySubType() *string
	SetStoryType(v string) *QueryStoriesRequest
	GetStoryType() *string
	SetWithEmptyStories(v bool) *QueryStoriesRequest
	GetWithEmptyStories() *bool
}

type QueryStoriesRequest struct {
	// The creation time range of the story.
	CreateTimeRange *TimeRange `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	// The custom label key-value pairs. Only stories that match the specified label pairs are returned.
	//
	// example:
	//
	// key=value
	CustomLabels *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset. For more information about how to obtain the name, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The IDs of the figure clusters.
	FigureClusterIds []*string `json:"FigureClusterIds,omitempty" xml:"FigureClusterIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return in a single call. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. If this parameter is left empty, the query starts from the beginning. To query the next page, set this parameter to the NextToken value returned in the previous call.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpw****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the story object.
	//
	// example:
	//
	// id1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The sorting order. Valid values:
	//
	// - asc: Ascending order.
	//
	// - desc: Descending order.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project. For more information about how to obtain the name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The field used for sorting. Valid values:
	//
	// - CreateTime: Sorts by story creation time.
	//
	// - StoryName: Sorts by story name.
	//
	// - StoryStartTime: Sorts by story start time.
	//
	// - StoryEndTime: Sorts by story end time.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The end time range of the photos or videos in the story.
	StoryEndTimeRange *TimeRange `json:"StoryEndTimeRange,omitempty" xml:"StoryEndTimeRange,omitempty"`
	// The name of the story.
	//
	// example:
	//
	// name1
	StoryName *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	// The start time range of the photos or videos in the story.
	StoryStartTimeRange *TimeRange `json:"StoryStartTimeRange,omitempty" xml:"StoryStartTimeRange,omitempty"`
	// The subtype of the story. For valid values, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// example:
	//
	// SeasonHighlights
	StorySubType *string `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	// The type of the story. For valid values, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// example:
	//
	// TimeMemory
	StoryType *string `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	// Specifies whether to return empty stories. Valid values:
	//
	// - true: Returns empty stories. This is the default value.
	//
	// - false: Does not return empty stories.
	//
	// example:
	//
	// true
	WithEmptyStories *bool `json:"WithEmptyStories,omitempty" xml:"WithEmptyStories,omitempty"`
}

func (s QueryStoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryStoriesRequest) GoString() string {
	return s.String()
}

func (s *QueryStoriesRequest) GetCreateTimeRange() *TimeRange {
	return s.CreateTimeRange
}

func (s *QueryStoriesRequest) GetCustomLabels() *string {
	return s.CustomLabels
}

func (s *QueryStoriesRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *QueryStoriesRequest) GetFigureClusterIds() []*string {
	return s.FigureClusterIds
}

func (s *QueryStoriesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *QueryStoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryStoriesRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *QueryStoriesRequest) GetOrder() *string {
	return s.Order
}

func (s *QueryStoriesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryStoriesRequest) GetSort() *string {
	return s.Sort
}

func (s *QueryStoriesRequest) GetStoryEndTimeRange() *TimeRange {
	return s.StoryEndTimeRange
}

func (s *QueryStoriesRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *QueryStoriesRequest) GetStoryStartTimeRange() *TimeRange {
	return s.StoryStartTimeRange
}

func (s *QueryStoriesRequest) GetStorySubType() *string {
	return s.StorySubType
}

func (s *QueryStoriesRequest) GetStoryType() *string {
	return s.StoryType
}

func (s *QueryStoriesRequest) GetWithEmptyStories() *bool {
	return s.WithEmptyStories
}

func (s *QueryStoriesRequest) SetCreateTimeRange(v *TimeRange) *QueryStoriesRequest {
	s.CreateTimeRange = v
	return s
}

func (s *QueryStoriesRequest) SetCustomLabels(v string) *QueryStoriesRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryStoriesRequest) SetDatasetName(v string) *QueryStoriesRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryStoriesRequest) SetFigureClusterIds(v []*string) *QueryStoriesRequest {
	s.FigureClusterIds = v
	return s
}

func (s *QueryStoriesRequest) SetMaxResults(v int64) *QueryStoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryStoriesRequest) SetNextToken(v string) *QueryStoriesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryStoriesRequest) SetObjectId(v string) *QueryStoriesRequest {
	s.ObjectId = &v
	return s
}

func (s *QueryStoriesRequest) SetOrder(v string) *QueryStoriesRequest {
	s.Order = &v
	return s
}

func (s *QueryStoriesRequest) SetProjectName(v string) *QueryStoriesRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryStoriesRequest) SetSort(v string) *QueryStoriesRequest {
	s.Sort = &v
	return s
}

func (s *QueryStoriesRequest) SetStoryEndTimeRange(v *TimeRange) *QueryStoriesRequest {
	s.StoryEndTimeRange = v
	return s
}

func (s *QueryStoriesRequest) SetStoryName(v string) *QueryStoriesRequest {
	s.StoryName = &v
	return s
}

func (s *QueryStoriesRequest) SetStoryStartTimeRange(v *TimeRange) *QueryStoriesRequest {
	s.StoryStartTimeRange = v
	return s
}

func (s *QueryStoriesRequest) SetStorySubType(v string) *QueryStoriesRequest {
	s.StorySubType = &v
	return s
}

func (s *QueryStoriesRequest) SetStoryType(v string) *QueryStoriesRequest {
	s.StoryType = &v
	return s
}

func (s *QueryStoriesRequest) SetWithEmptyStories(v bool) *QueryStoriesRequest {
	s.WithEmptyStories = &v
	return s
}

func (s *QueryStoriesRequest) Validate() error {
	if s.CreateTimeRange != nil {
		if err := s.CreateTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.StoryEndTimeRange != nil {
		if err := s.StoryEndTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.StoryStartTimeRange != nil {
		if err := s.StoryStartTimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}
