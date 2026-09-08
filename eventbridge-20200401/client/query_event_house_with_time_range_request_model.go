// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventHouseWithTimeRangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryEventHouseWithTimeRangeRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *QueryEventHouseWithTimeRangeRequest
	GetEndTime() *int64
	SetLimit(v int32) *QueryEventHouseWithTimeRangeRequest
	GetLimit() *int32
	SetQuery(v string) *QueryEventHouseWithTimeRangeRequest
	GetQuery() *string
}

type QueryEventHouseWithTimeRangeRequest struct {
	// The start time for querying internal EventHouse data. Specify a UNIX timestamp in seconds. The time range includes this point in time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1787587200
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time for querying internal EventHouse data. Specify a UNIX timestamp in seconds. The time range excludes this point in time. The value must be greater than BeginTime.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1787590800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of result rows that can be returned for this query.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The single read-only SQL statement to execute. You can query internal EventHouse data or perform federated queries with mounted external data sources.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM my_catalog.my_namespace.my_table LIMIT 100
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s QueryEventHouseWithTimeRangeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEventHouseWithTimeRangeRequest) GoString() string {
	return s.String()
}

func (s *QueryEventHouseWithTimeRangeRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryEventHouseWithTimeRangeRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryEventHouseWithTimeRangeRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryEventHouseWithTimeRangeRequest) GetQuery() *string {
	return s.Query
}

func (s *QueryEventHouseWithTimeRangeRequest) SetBeginTime(v int64) *QueryEventHouseWithTimeRangeRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeRequest) SetEndTime(v int64) *QueryEventHouseWithTimeRangeRequest {
	s.EndTime = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeRequest) SetLimit(v int32) *QueryEventHouseWithTimeRangeRequest {
	s.Limit = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeRequest) SetQuery(v string) *QueryEventHouseWithTimeRangeRequest {
	s.Query = &v
	return s
}

func (s *QueryEventHouseWithTimeRangeRequest) Validate() error {
	return dara.Validate(s)
}
