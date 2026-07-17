// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatisticReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v string) *ListStatisticReportRequest
	GetColumns() *string
	SetEndTime(v int32) *ListStatisticReportRequest
	GetEndTime() *int32
	SetPageNumber(v int32) *ListStatisticReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStatisticReportRequest
	GetPageSize() *int32
	SetQuery(v string) *ListStatisticReportRequest
	GetQuery() *string
	SetStartTime(v int32) *ListStatisticReportRequest
	GetStartTime() *int32
}

type ListStatisticReportRequest struct {
	// The fields to query. Specify the fields in the columns="pv,uv,ipv" format. For more information, see [Metrics of statistical reports](https://help.aliyun.com/document_detail/187665.html).
	//
	// example:
	//
	// pv,uv
	Columns *string `json:"columns,omitempty" xml:"columns,omitempty"`
	// The end timestamp, in seconds. The default value is the current timestamp.
	//
	// example:
	//
	// 1582646399
	EndTime *int32 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The query conditions, in the k1:v1,k2:v2 format. Valid values:
	//
	// - experimentSerialNumber: the globally unique serial number of the experiment.
	//
	// - sceneTag: the tag of the scenario.
	//
	// - bizType: the business identity.
	//
	// - modelId: the ID of the algorithm model.
	//
	// example:
	//
	// bizType:test,sceneTag:myTag
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The start timestamp, in seconds.
	//
	// example:
	//
	// 1582214400
	StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListStatisticReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStatisticReportRequest) GoString() string {
	return s.String()
}

func (s *ListStatisticReportRequest) GetColumns() *string {
	return s.Columns
}

func (s *ListStatisticReportRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ListStatisticReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStatisticReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStatisticReportRequest) GetQuery() *string {
	return s.Query
}

func (s *ListStatisticReportRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ListStatisticReportRequest) SetColumns(v string) *ListStatisticReportRequest {
	s.Columns = &v
	return s
}

func (s *ListStatisticReportRequest) SetEndTime(v int32) *ListStatisticReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListStatisticReportRequest) SetPageNumber(v int32) *ListStatisticReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStatisticReportRequest) SetPageSize(v int32) *ListStatisticReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListStatisticReportRequest) SetQuery(v string) *ListStatisticReportRequest {
	s.Query = &v
	return s
}

func (s *ListStatisticReportRequest) SetStartTime(v int32) *ListStatisticReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListStatisticReportRequest) Validate() error {
	return dara.Validate(s)
}
