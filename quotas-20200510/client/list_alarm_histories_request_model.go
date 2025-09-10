// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v string) *ListAlarmHistoriesRequest
	GetAlarmId() *string
	SetEndTime(v int64) *ListAlarmHistoriesRequest
	GetEndTime() *int64
	SetKeyword(v string) *ListAlarmHistoriesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListAlarmHistoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAlarmHistoriesRequest
	GetNextToken() *string
	SetProductCode(v string) *ListAlarmHistoriesRequest
	GetProductCode() *string
	SetStartTime(v int64) *ListAlarmHistoriesRequest
	GetStartTime() *int64
}

type ListAlarmHistoriesRequest struct {
	// The ID of the alert.
	//
	// example:
	//
	// 18b3be23-b7b0-4d45-91bc-d0c331aa****
	AlarmId *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 20201024
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that is used for the query.
	//
	// example:
	//
	// Quantity
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of records that can be returned for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// > For more information, see [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 20201020
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAlarmHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesRequest) GetAlarmId() *string {
	return s.AlarmId
}

func (s *ListAlarmHistoriesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAlarmHistoriesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAlarmHistoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlarmHistoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlarmHistoriesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListAlarmHistoriesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAlarmHistoriesRequest) SetAlarmId(v string) *ListAlarmHistoriesRequest {
	s.AlarmId = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetEndTime(v int64) *ListAlarmHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetKeyword(v string) *ListAlarmHistoriesRequest {
	s.Keyword = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetMaxResults(v int32) *ListAlarmHistoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetNextToken(v string) *ListAlarmHistoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetProductCode(v string) *ListAlarmHistoriesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetStartTime(v int64) *ListAlarmHistoriesRequest {
	s.StartTime = &v
	return s
}

func (s *ListAlarmHistoriesRequest) Validate() error {
	return dara.Validate(s)
}
