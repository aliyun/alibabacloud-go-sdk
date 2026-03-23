// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiSparkApplicationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetApplicationId() *string
	SetApplicationName(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetApplicationName() *string
	SetEndTimeShrink(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetEndTimeShrink() *string
	SetLatestSqlStatementStatuses(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetLatestSqlStatementStatuses() *string
	SetMaxResults(v int32) *ListKyuubiSparkApplicationsShrinkRequest
	GetMaxResults() *int32
	SetMinDuration(v int64) *ListKyuubiSparkApplicationsShrinkRequest
	GetMinDuration() *int64
	SetNextToken(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetNextToken() *string
	SetOrderByShrink(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetOrderByShrink() *string
	SetResourceQueueId(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetResourceQueueId() *string
	SetSort(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetSort() *string
	SetStartTimeShrink(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetStartTimeShrink() *string
	SetStates(v string) *ListKyuubiSparkApplicationsShrinkRequest
	GetStates() *string
}

type ListKyuubiSparkApplicationsShrinkRequest struct {
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
	EndTimeShrink   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// ["SUCCESS"]
	LatestSqlStatementStatuses *string `json:"latestSqlStatementStatuses,omitempty" xml:"latestSqlStatementStatuses,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 60000
	MinDuration *int64 `json:"minDuration,omitempty" xml:"minDuration,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken     *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrderByShrink *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// root_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// example:
	//
	// asc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// The range of start time.
	StartTimeShrink *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// ["ERROR"]
	States *string `json:"states,omitempty" xml:"states,omitempty"`
}

func (s ListKyuubiSparkApplicationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiSparkApplicationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetEndTimeShrink() *string {
	return s.EndTimeShrink
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetLatestSqlStatementStatuses() *string {
	return s.LatestSqlStatementStatuses
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetMinDuration() *int64 {
	return s.MinDuration
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetOrderByShrink() *string {
	return s.OrderByShrink
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetResourceQueueId() *string {
	return s.ResourceQueueId
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetStartTimeShrink() *string {
	return s.StartTimeShrink
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) GetStates() *string {
	return s.States
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetApplicationId(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetApplicationName(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetEndTimeShrink(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.EndTimeShrink = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetLatestSqlStatementStatuses(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.LatestSqlStatementStatuses = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetMaxResults(v int32) *ListKyuubiSparkApplicationsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetMinDuration(v int64) *ListKyuubiSparkApplicationsShrinkRequest {
	s.MinDuration = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetNextToken(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetOrderByShrink(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.OrderByShrink = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetResourceQueueId(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetSort(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetStartTimeShrink(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.StartTimeShrink = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) SetStates(v string) *ListKyuubiSparkApplicationsShrinkRequest {
	s.States = &v
	return s
}

func (s *ListKyuubiSparkApplicationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
