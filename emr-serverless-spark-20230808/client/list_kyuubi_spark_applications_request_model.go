// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiSparkApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListKyuubiSparkApplicationsRequest
	GetApplicationId() *string
	SetApplicationName(v string) *ListKyuubiSparkApplicationsRequest
	GetApplicationName() *string
	SetMaxResults(v int32) *ListKyuubiSparkApplicationsRequest
	GetMaxResults() *int32
	SetMinDuration(v int64) *ListKyuubiSparkApplicationsRequest
	GetMinDuration() *int64
	SetNextToken(v string) *ListKyuubiSparkApplicationsRequest
	GetNextToken() *string
	SetOrderBy(v []*string) *ListKyuubiSparkApplicationsRequest
	GetOrderBy() []*string
	SetResourceQueueId(v string) *ListKyuubiSparkApplicationsRequest
	GetResourceQueueId() *string
	SetSort(v string) *ListKyuubiSparkApplicationsRequest
	GetSort() *string
	SetStartTime(v *ListKyuubiSparkApplicationsRequestStartTime) *ListKyuubiSparkApplicationsRequest
	GetStartTime() *ListKyuubiSparkApplicationsRequestStartTime
}

type ListKyuubiSparkApplicationsRequest struct {
	// The ID of the application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// spark-339f844005b6404c95f9f7c7a13b****
	ApplicationId *string `json:"applicationId,omitempty" xml:"applicationId,omitempty"`
	// The name of the Spark application that is submitted by using a Kyuubi gateway.
	//
	// example:
	//
	// kyuubi-connection-spark-sql-anonymous-fa9a5e73-b4b1-474a-b****
	ApplicationName *string `json:"applicationName,omitempty" xml:"applicationName,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults  *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	MinDuration *int64 `json:"minDuration,omitempty" xml:"minDuration,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken       *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrderBy         []*string `json:"orderBy,omitempty" xml:"orderBy,omitempty" type:"Repeated"`
	ResourceQueueId *string   `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	Sort            *string   `json:"sort,omitempty" xml:"sort,omitempty"`
	// The range of start time.
	StartTime *ListKyuubiSparkApplicationsRequestStartTime `json:"startTime,omitempty" xml:"startTime,omitempty" type:"Struct"`
}

func (s ListKyuubiSparkApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiSparkApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListKyuubiSparkApplicationsRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListKyuubiSparkApplicationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKyuubiSparkApplicationsRequest) GetMinDuration() *int64 {
	return s.MinDuration
}

func (s *ListKyuubiSparkApplicationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKyuubiSparkApplicationsRequest) GetOrderBy() []*string {
	return s.OrderBy
}

func (s *ListKyuubiSparkApplicationsRequest) GetResourceQueueId() *string {
	return s.ResourceQueueId
}

func (s *ListKyuubiSparkApplicationsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListKyuubiSparkApplicationsRequest) GetStartTime() *ListKyuubiSparkApplicationsRequestStartTime {
	return s.StartTime
}

func (s *ListKyuubiSparkApplicationsRequest) SetApplicationId(v string) *ListKyuubiSparkApplicationsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetApplicationName(v string) *ListKyuubiSparkApplicationsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetMaxResults(v int32) *ListKyuubiSparkApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetMinDuration(v int64) *ListKyuubiSparkApplicationsRequest {
	s.MinDuration = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetNextToken(v string) *ListKyuubiSparkApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetOrderBy(v []*string) *ListKyuubiSparkApplicationsRequest {
	s.OrderBy = v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetResourceQueueId(v string) *ListKyuubiSparkApplicationsRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetSort(v string) *ListKyuubiSparkApplicationsRequest {
	s.Sort = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) SetStartTime(v *ListKyuubiSparkApplicationsRequestStartTime) *ListKyuubiSparkApplicationsRequest {
	s.StartTime = v
	return s
}

func (s *ListKyuubiSparkApplicationsRequest) Validate() error {
	return dara.Validate(s)
}

type ListKyuubiSparkApplicationsRequestStartTime struct {
	// The end of the start time range.
	//
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The beginning of the start time range.
	//
	// example:
	//
	// 1709740800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListKyuubiSparkApplicationsRequestStartTime) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiSparkApplicationsRequestStartTime) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsRequestStartTime) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListKyuubiSparkApplicationsRequestStartTime) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListKyuubiSparkApplicationsRequestStartTime) SetEndTime(v int64) *ListKyuubiSparkApplicationsRequestStartTime {
	s.EndTime = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequestStartTime) SetStartTime(v int64) *ListKyuubiSparkApplicationsRequestStartTime {
	s.StartTime = &v
	return s
}

func (s *ListKyuubiSparkApplicationsRequestStartTime) Validate() error {
	return dara.Validate(s)
}
