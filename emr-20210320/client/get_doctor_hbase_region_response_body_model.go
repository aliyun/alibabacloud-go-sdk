// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorHBaseRegionResponseBodyData) *GetDoctorHBaseRegionResponseBody
	GetData() *GetDoctorHBaseRegionResponseBodyData
	SetRequestId(v string) *GetDoctorHBaseRegionResponseBody
	GetRequestId() *string
}

type GetDoctorHBaseRegionResponseBody struct {
	// Returned data.
	Data *GetDoctorHBaseRegionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorHBaseRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionResponseBody) GetData() *GetDoctorHBaseRegionResponseBodyData {
	return s.Data
}

func (s *GetDoctorHBaseRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorHBaseRegionResponseBody) SetData(v *GetDoctorHBaseRegionResponseBodyData) *GetDoctorHBaseRegionResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorHBaseRegionResponseBody) SetRequestId(v string) *GetDoctorHBaseRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionResponseBodyData struct {
	// Metrics information.
	Metrics *GetDoctorHBaseRegionResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// Host of the RegionServer.
	//
	// example:
	//
	// emr-worker-2.cluster-20****
	RegionServerHost *string `json:"RegionServerHost,omitempty" xml:"RegionServerHost,omitempty"`
	// Table name.
	//
	// example:
	//
	// tb_item
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetDoctorHBaseRegionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionResponseBodyData) GetMetrics() *GetDoctorHBaseRegionResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorHBaseRegionResponseBodyData) GetRegionServerHost() *string {
	return s.RegionServerHost
}

func (s *GetDoctorHBaseRegionResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *GetDoctorHBaseRegionResponseBodyData) SetMetrics(v *GetDoctorHBaseRegionResponseBodyDataMetrics) *GetDoctorHBaseRegionResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyData) SetRegionServerHost(v string) *GetDoctorHBaseRegionResponseBodyData {
	s.RegionServerHost = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyData) SetTableName(v string) *GetDoctorHBaseRegionResponseBodyData {
	s.TableName = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionResponseBodyDataMetrics struct {
	// Number of read requests in a single day.
	DailyReadRequest *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest `json:"DailyReadRequest,omitempty" xml:"DailyReadRequest,omitempty" type:"Struct"`
	// Number of write requests in a single day.
	DailyWriteRequest *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest `json:"DailyWriteRequest,omitempty" xml:"DailyWriteRequest,omitempty" type:"Struct"`
	// Store file count.
	StoreFileCount *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount `json:"StoreFileCount,omitempty" xml:"StoreFileCount,omitempty" type:"Struct"`
	// Total read request count
	TotalReadRequest *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest `json:"TotalReadRequest,omitempty" xml:"TotalReadRequest,omitempty" type:"Struct"`
	// Total write request count
	TotalWriteRequest *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest `json:"TotalWriteRequest,omitempty" xml:"TotalWriteRequest,omitempty" type:"Struct"`
}

func (s GetDoctorHBaseRegionResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) GetDailyReadRequest() *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest {
	return s.DailyReadRequest
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) GetDailyWriteRequest() *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest {
	return s.DailyWriteRequest
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) GetStoreFileCount() *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount {
	return s.StoreFileCount
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) GetTotalReadRequest() *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest {
	return s.TotalReadRequest
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) GetTotalWriteRequest() *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest {
	return s.TotalWriteRequest
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) SetDailyReadRequest(v *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) *GetDoctorHBaseRegionResponseBodyDataMetrics {
	s.DailyReadRequest = v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) SetDailyWriteRequest(v *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) *GetDoctorHBaseRegionResponseBodyDataMetrics {
	s.DailyWriteRequest = v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) SetStoreFileCount(v *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) *GetDoctorHBaseRegionResponseBodyDataMetrics {
	s.StoreFileCount = v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) SetTotalReadRequest(v *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) *GetDoctorHBaseRegionResponseBodyDataMetrics {
	s.TotalReadRequest = v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) SetTotalWriteRequest(v *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) *GetDoctorHBaseRegionResponseBodyDataMetrics {
	s.TotalWriteRequest = v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest struct {
	// Description of the metric.
	//
	// example:
	//
	// Number of read requests per day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// dailyReadRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) SetDescription(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) SetName(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) SetUnit(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) SetValue(v int64) *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyReadRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest struct {
	// Description of the metric.
	//
	// example:
	//
	// Number of write requests per day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the metric.
	//
	// example:
	//
	// dailyWriteRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) SetDescription(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) SetName(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) SetUnit(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) SetValue(v int64) *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsDailyWriteRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount struct {
	// Description of the metric.
	//
	// example:
	//
	// Number of store file
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// storeFileCount
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 100
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) SetDescription(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) SetName(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) SetUnit(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) SetValue(v int64) *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsStoreFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest struct {
	// Metric description.
	//
	// example:
	//
	// Total read request
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// totalReadRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) SetDescription(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) SetName(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) SetUnit(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) SetValue(v int64) *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalReadRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest struct {
	// Metric description.
	//
	// example:
	//
	// Total Write Request
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Metric name.
	//
	// example:
	//
	// totalWriteRequest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Metric unit.
	//
	// example:
	//
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 1000
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) SetDescription(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) SetName(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) SetUnit(v string) *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) SetValue(v int64) *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseRegionResponseBodyDataMetricsTotalWriteRequest) Validate() error {
	return dara.Validate(s)
}
