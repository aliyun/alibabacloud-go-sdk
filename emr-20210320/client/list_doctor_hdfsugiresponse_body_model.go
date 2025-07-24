// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHDFSUGIResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorHDFSUGIResponseBodyData) *ListDoctorHDFSUGIResponseBody
	GetData() []*ListDoctorHDFSUGIResponseBodyData
	SetMaxResults(v int32) *ListDoctorHDFSUGIResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHDFSUGIResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorHDFSUGIResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorHDFSUGIResponseBody
	GetTotalCount() *int32
}

type ListDoctorHDFSUGIResponseBody struct {
	// The results of batch HDFS analysis.
	Data []*ListDoctorHDFSUGIResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries that are returned.
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

func (s ListDoctorHDFSUGIResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSUGIResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSUGIResponseBody) GetData() []*ListDoctorHDFSUGIResponseBodyData {
	return s.Data
}

func (s *ListDoctorHDFSUGIResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHDFSUGIResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHDFSUGIResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorHDFSUGIResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorHDFSUGIResponseBody) SetData(v []*ListDoctorHDFSUGIResponseBodyData) *ListDoctorHDFSUGIResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorHDFSUGIResponseBody) SetMaxResults(v int32) *ListDoctorHDFSUGIResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBody) SetNextToken(v string) *ListDoctorHDFSUGIResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBody) SetRequestId(v string) *ListDoctorHDFSUGIResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBody) SetTotalCount(v int32) *ListDoctorHDFSUGIResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSUGIResponseBodyData struct {
	// The metric information.
	Metrics *ListDoctorHDFSUGIResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The actual name of the owner or group returned based on the value of Type.
	//
	// example:
	//
	// DW
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDoctorHDFSUGIResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSUGIResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSUGIResponseBodyData) GetMetrics() *ListDoctorHDFSUGIResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListDoctorHDFSUGIResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSUGIResponseBodyData) SetMetrics(v *ListDoctorHDFSUGIResponseBodyDataMetrics) *ListDoctorHDFSUGIResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyData) SetName(v string) *ListDoctorHDFSUGIResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSUGIResponseBodyDataMetrics struct {
	// The total data size.
	TotalDataSize *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	// The total number of directories.
	TotalDirCount *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount `json:"TotalDirCount,omitempty" xml:"TotalDirCount,omitempty" type:"Struct"`
	// The total number of files.
	TotalFileCount *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty" type:"Struct"`
}

func (s ListDoctorHDFSUGIResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSUGIResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetrics) GetTotalDataSize() *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetrics) GetTotalDirCount() *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount {
	return s.TotalDirCount
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetrics) GetTotalFileCount() *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount {
	return s.TotalFileCount
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetrics) SetTotalDataSize(v *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) *ListDoctorHDFSUGIResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetrics) SetTotalDirCount(v *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) *ListDoctorHDFSUGIResponseBodyDataMetrics {
	s.TotalDirCount = v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetrics) SetTotalFileCount(v *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) *ListDoctorHDFSUGIResponseBodyDataMetrics {
	s.TotalFileCount = v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize struct {
	// The description of the metric.
	//
	// example:
	//
	// Total data size in megabytes (MB)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalDataSize
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 40440503
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) SetName(v string) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of total dirs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalDirCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) SetDescription(v string) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) SetName(v string) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) SetUnit(v string) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) SetValue(v int64) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalDirCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of total files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ”“
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 34
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) SetDescription(v string) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount {
	s.Description = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) SetName(v string) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount {
	s.Name = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) SetUnit(v string) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) SetValue(v int64) *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount {
	s.Value = &v
	return s
}

func (s *ListDoctorHDFSUGIResponseBodyDataMetricsTotalFileCount) Validate() error {
	return dara.Validate(s)
}
