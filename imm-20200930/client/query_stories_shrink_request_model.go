// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStoriesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeRangeShrink(v string) *QueryStoriesShrinkRequest
	GetCreateTimeRangeShrink() *string
	SetCustomLabels(v string) *QueryStoriesShrinkRequest
	GetCustomLabels() *string
	SetDatasetName(v string) *QueryStoriesShrinkRequest
	GetDatasetName() *string
	SetFigureClusterIdsShrink(v string) *QueryStoriesShrinkRequest
	GetFigureClusterIdsShrink() *string
	SetMaxResults(v int64) *QueryStoriesShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *QueryStoriesShrinkRequest
	GetNextToken() *string
	SetObjectId(v string) *QueryStoriesShrinkRequest
	GetObjectId() *string
	SetOrder(v string) *QueryStoriesShrinkRequest
	GetOrder() *string
	SetProjectName(v string) *QueryStoriesShrinkRequest
	GetProjectName() *string
	SetSort(v string) *QueryStoriesShrinkRequest
	GetSort() *string
	SetStoryEndTimeRangeShrink(v string) *QueryStoriesShrinkRequest
	GetStoryEndTimeRangeShrink() *string
	SetStoryName(v string) *QueryStoriesShrinkRequest
	GetStoryName() *string
	SetStoryStartTimeRangeShrink(v string) *QueryStoriesShrinkRequest
	GetStoryStartTimeRangeShrink() *string
	SetStorySubType(v string) *QueryStoriesShrinkRequest
	GetStorySubType() *string
	SetStoryType(v string) *QueryStoriesShrinkRequest
	GetStoryType() *string
	SetWithEmptyStories(v bool) *QueryStoriesShrinkRequest
	GetWithEmptyStories() *bool
}

type QueryStoriesShrinkRequest struct {
	// The time range in which stories were created.
	CreateTimeRangeShrink *string `json:"CreateTimeRange,omitempty" xml:"CreateTimeRange,omitempty"`
	// The custom labels in key-value pairs.
	//
	// example:
	//
	// key=value
	CustomLabels *string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The IDs of the face clusters.
	FigureClusterIdsShrink *string `json:"FigureClusterIds,omitempty" xml:"FigureClusterIds,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If you do not specify this token in the next request, results are returned from the beginning.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpw****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the story.
	//
	// example:
	//
	// id1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The sort order. Valid values:
	//
	// 	- asc: in ascending order.
	//
	// 	- desc: in descending order.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The sort field. Valid values:
	//
	// 	- CreateTime: sorts by story creation time.
	//
	// 	- StoryName: sorts by story name.
	//
	// 	- StoryStartTime: sorts by story start time.
	//
	// 	- StoryEndTime: sorts by story end time.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The time range for the creation time of the last photo or video in the story.
	StoryEndTimeRangeShrink *string `json:"StoryEndTimeRange,omitempty" xml:"StoryEndTimeRange,omitempty"`
	// The name of the story.
	//
	// example:
	//
	// name1
	StoryName *string `json:"StoryName,omitempty" xml:"StoryName,omitempty"`
	// The time range for the creation time of the first photo or video in the story.
	StoryStartTimeRangeShrink *string `json:"StoryStartTimeRange,omitempty" xml:"StoryStartTimeRange,omitempty"`
	// The subtype of the story. For a list of valid values, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// example:
	//
	// SeasonHighlights
	StorySubType *string `json:"StorySubType,omitempty" xml:"StorySubType,omitempty"`
	// The type of the story. For a list of valid values, see [Story types and subtypes](https://help.aliyun.com/document_detail/2743998.html).
	//
	// example:
	//
	// TimeMemory
	StoryType *string `json:"StoryType,omitempty" xml:"StoryType,omitempty"`
	// Specifies whether to return empty stories. Valid values:
	//
	// 	- true (The default value)
	//
	// 	- false
	//
	// example:
	//
	// true
	WithEmptyStories *bool `json:"WithEmptyStories,omitempty" xml:"WithEmptyStories,omitempty"`
}

func (s QueryStoriesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryStoriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryStoriesShrinkRequest) GetCreateTimeRangeShrink() *string {
	return s.CreateTimeRangeShrink
}

func (s *QueryStoriesShrinkRequest) GetCustomLabels() *string {
	return s.CustomLabels
}

func (s *QueryStoriesShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *QueryStoriesShrinkRequest) GetFigureClusterIdsShrink() *string {
	return s.FigureClusterIdsShrink
}

func (s *QueryStoriesShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *QueryStoriesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryStoriesShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *QueryStoriesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *QueryStoriesShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryStoriesShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *QueryStoriesShrinkRequest) GetStoryEndTimeRangeShrink() *string {
	return s.StoryEndTimeRangeShrink
}

func (s *QueryStoriesShrinkRequest) GetStoryName() *string {
	return s.StoryName
}

func (s *QueryStoriesShrinkRequest) GetStoryStartTimeRangeShrink() *string {
	return s.StoryStartTimeRangeShrink
}

func (s *QueryStoriesShrinkRequest) GetStorySubType() *string {
	return s.StorySubType
}

func (s *QueryStoriesShrinkRequest) GetStoryType() *string {
	return s.StoryType
}

func (s *QueryStoriesShrinkRequest) GetWithEmptyStories() *bool {
	return s.WithEmptyStories
}

func (s *QueryStoriesShrinkRequest) SetCreateTimeRangeShrink(v string) *QueryStoriesShrinkRequest {
	s.CreateTimeRangeShrink = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetCustomLabels(v string) *QueryStoriesShrinkRequest {
	s.CustomLabels = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetDatasetName(v string) *QueryStoriesShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetFigureClusterIdsShrink(v string) *QueryStoriesShrinkRequest {
	s.FigureClusterIdsShrink = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetMaxResults(v int64) *QueryStoriesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetNextToken(v string) *QueryStoriesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetObjectId(v string) *QueryStoriesShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetOrder(v string) *QueryStoriesShrinkRequest {
	s.Order = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetProjectName(v string) *QueryStoriesShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetSort(v string) *QueryStoriesShrinkRequest {
	s.Sort = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStoryEndTimeRangeShrink(v string) *QueryStoriesShrinkRequest {
	s.StoryEndTimeRangeShrink = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStoryName(v string) *QueryStoriesShrinkRequest {
	s.StoryName = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStoryStartTimeRangeShrink(v string) *QueryStoriesShrinkRequest {
	s.StoryStartTimeRangeShrink = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStorySubType(v string) *QueryStoriesShrinkRequest {
	s.StorySubType = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetStoryType(v string) *QueryStoriesShrinkRequest {
	s.StoryType = &v
	return s
}

func (s *QueryStoriesShrinkRequest) SetWithEmptyStories(v bool) *QueryStoriesShrinkRequest {
	s.WithEmptyStories = &v
	return s
}

func (s *QueryStoriesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
