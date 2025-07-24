// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorReportsResponseBodyData) *ListDoctorReportsResponseBody
	GetData() []*ListDoctorReportsResponseBodyData
	SetMaxResults(v int32) *ListDoctorReportsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorReportsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorReportsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorReportsResponseBody
	GetTotalCount() *int32
}

type ListDoctorReportsResponseBody struct {
	// The reports.
	Data []*ListDoctorReportsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDoctorReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorReportsResponseBody) GetData() []*ListDoctorReportsResponseBodyData {
	return s.Data
}

func (s *ListDoctorReportsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorReportsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorReportsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorReportsResponseBody) SetData(v []*ListDoctorReportsResponseBodyData) *ListDoctorReportsResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorReportsResponseBody) SetMaxResults(v int32) *ListDoctorReportsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorReportsResponseBody) SetNextToken(v string) *ListDoctorReportsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorReportsResponseBody) SetRequestId(v string) *ListDoctorReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorReportsResponseBody) SetTotalCount(v int32) *ListDoctorReportsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorReportsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDoctorReportsResponseBodyData struct {
	// The component types.
	//
	// Valid values:
	//
	// 	- compute
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- hive
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- hdfs
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- yarn
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- oss
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- hbase
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// null
	ComponentTypes []*string `json:"ComponentTypes,omitempty" xml:"ComponentTypes,omitempty" type:"Repeated"`
	// The date on which the report was generated.
	//
	// example:
	//
	// 2023-06-29
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// The summary of the report.
	SummaryReport *ListDoctorReportsResponseBodyDataSummaryReport `json:"SummaryReport,omitempty" xml:"SummaryReport,omitempty" type:"Struct"`
}

func (s ListDoctorReportsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorReportsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorReportsResponseBodyData) GetComponentTypes() []*string {
	return s.ComponentTypes
}

func (s *ListDoctorReportsResponseBodyData) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorReportsResponseBodyData) GetSummaryReport() *ListDoctorReportsResponseBodyDataSummaryReport {
	return s.SummaryReport
}

func (s *ListDoctorReportsResponseBodyData) SetComponentTypes(v []*string) *ListDoctorReportsResponseBodyData {
	s.ComponentTypes = v
	return s
}

func (s *ListDoctorReportsResponseBodyData) SetDateTime(v string) *ListDoctorReportsResponseBodyData {
	s.DateTime = &v
	return s
}

func (s *ListDoctorReportsResponseBodyData) SetSummaryReport(v *ListDoctorReportsResponseBodyDataSummaryReport) *ListDoctorReportsResponseBodyData {
	s.SummaryReport = v
	return s
}

func (s *ListDoctorReportsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDoctorReportsResponseBodyDataSummaryReport struct {
	// The score.
	//
	// example:
	//
	// 88
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The optimization suggestion.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The summary of the report.
	//
	// example:
	//
	// eastbuy-mse-plugin-auth
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListDoctorReportsResponseBodyDataSummaryReport) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorReportsResponseBodyDataSummaryReport) GoString() string {
	return s.String()
}

func (s *ListDoctorReportsResponseBodyDataSummaryReport) GetScore() *int32 {
	return s.Score
}

func (s *ListDoctorReportsResponseBodyDataSummaryReport) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ListDoctorReportsResponseBodyDataSummaryReport) GetSummary() *string {
	return s.Summary
}

func (s *ListDoctorReportsResponseBodyDataSummaryReport) SetScore(v int32) *ListDoctorReportsResponseBodyDataSummaryReport {
	s.Score = &v
	return s
}

func (s *ListDoctorReportsResponseBodyDataSummaryReport) SetSuggestion(v string) *ListDoctorReportsResponseBodyDataSummaryReport {
	s.Suggestion = &v
	return s
}

func (s *ListDoctorReportsResponseBodyDataSummaryReport) SetSummary(v string) *ListDoctorReportsResponseBodyDataSummaryReport {
	s.Summary = &v
	return s
}

func (s *ListDoctorReportsResponseBodyDataSummaryReport) Validate() error {
	return dara.Validate(s)
}
