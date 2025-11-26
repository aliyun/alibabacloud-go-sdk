// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorHBaseClusterResponseBodyData) *GetDoctorHBaseClusterResponseBody
	GetData() *GetDoctorHBaseClusterResponseBodyData
	SetRequestId(v string) *GetDoctorHBaseClusterResponseBody
	GetRequestId() *string
}

type GetDoctorHBaseClusterResponseBody struct {
	// The returned data.
	Data *GetDoctorHBaseClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBody) GetData() *GetDoctorHBaseClusterResponseBodyData {
	return s.Data
}

func (s *GetDoctorHBaseClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorHBaseClusterResponseBody) SetData(v *GetDoctorHBaseClusterResponseBodyData) *GetDoctorHBaseClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBody) SetRequestId(v string) *GetDoctorHBaseClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDoctorHBaseClusterResponseBodyData struct {
	// The analysis result.
	Analysis *GetDoctorHBaseClusterResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The metric information.
	Metrics *GetDoctorHBaseClusterResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s GetDoctorHBaseClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyData) GetAnalysis() *GetDoctorHBaseClusterResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *GetDoctorHBaseClusterResponseBodyData) GetMetrics() *GetDoctorHBaseClusterResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorHBaseClusterResponseBodyData) SetAnalysis(v *GetDoctorHBaseClusterResponseBodyDataAnalysis) *GetDoctorHBaseClusterResponseBodyData {
	s.Analysis = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyData) SetMetrics(v *GetDoctorHBaseClusterResponseBodyDataMetrics) *GetDoctorHBaseClusterResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyData) Validate() error {
	if s.Analysis != nil {
		if err := s.Analysis.Validate(); err != nil {
			return err
		}
	}
	if s.Metrics != nil {
		if err := s.Metrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDoctorHBaseClusterResponseBodyDataAnalysis struct {
	// The overall score of the HBase cluster.
	//
	// example:
	//
	// 85
	HbaseScore *int32 `json:"HbaseScore,omitempty" xml:"HbaseScore,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataAnalysis) GetHbaseScore() *int32 {
	return s.HbaseScore
}

func (s *GetDoctorHBaseClusterResponseBodyDataAnalysis) SetHbaseScore(v int32) *GetDoctorHBaseClusterResponseBodyDataAnalysis {
	s.HbaseScore = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetrics struct {
	// The average load.
	AvgLoad *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad `json:"AvgLoad,omitempty" xml:"AvgLoad,omitempty" type:"Struct"`
	// The number of read requests in a day.
	DailyReadRequest *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest `json:"DailyReadRequest,omitempty" xml:"DailyReadRequest,omitempty" type:"Struct"`
	// The number of write requests in a day.
	DailyWriteRequest *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest `json:"DailyWriteRequest,omitempty" xml:"DailyWriteRequest,omitempty" type:"Struct"`
	// The memory size.
	MemHeap *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap `json:"MemHeap,omitempty" xml:"MemHeap,omitempty" type:"Struct"`
	// The normal average load.
	NormalAvgLoad *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad `json:"NormalAvgLoad,omitempty" xml:"NormalAvgLoad,omitempty" type:"Struct"`
	// The region balance degree.
	RegionBalance *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance `json:"RegionBalance,omitempty" xml:"RegionBalance,omitempty" type:"Struct"`
	// The number of regions.
	RegionCount *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount `json:"RegionCount,omitempty" xml:"RegionCount,omitempty" type:"Struct"`
	// The number of region servers.
	RegionServerCount *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount `json:"RegionServerCount,omitempty" xml:"RegionServerCount,omitempty" type:"Struct"`
	// The number of StoreFiles.
	StoreFileCount *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount `json:"StoreFileCount,omitempty" xml:"StoreFileCount,omitempty" type:"Struct"`
	// The number of tables.
	TableCount *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount `json:"TableCount,omitempty" xml:"TableCount,omitempty" type:"Struct"`
	// The size of the cluster.
	TotalDataSize *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty" type:"Struct"`
	// The total number of read requests.
	TotalReadRequest *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest `json:"TotalReadRequest,omitempty" xml:"TotalReadRequest,omitempty" type:"Struct"`
	// The total number of requests in the cluster.
	TotalRequest *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest `json:"TotalRequest,omitempty" xml:"TotalRequest,omitempty" type:"Struct"`
	// The total number of write requests.
	TotalWriteRequest *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest `json:"TotalWriteRequest,omitempty" xml:"TotalWriteRequest,omitempty" type:"Struct"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetAvgLoad() *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad {
	return s.AvgLoad
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetDailyReadRequest() *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest {
	return s.DailyReadRequest
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetDailyWriteRequest() *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest {
	return s.DailyWriteRequest
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetMemHeap() *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap {
	return s.MemHeap
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetNormalAvgLoad() *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad {
	return s.NormalAvgLoad
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetRegionBalance() *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance {
	return s.RegionBalance
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetRegionCount() *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount {
	return s.RegionCount
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetRegionServerCount() *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount {
	return s.RegionServerCount
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetStoreFileCount() *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount {
	return s.StoreFileCount
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetTableCount() *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount {
	return s.TableCount
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetTotalDataSize() *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize {
	return s.TotalDataSize
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetTotalReadRequest() *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest {
	return s.TotalReadRequest
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetTotalRequest() *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest {
	return s.TotalRequest
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) GetTotalWriteRequest() *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest {
	return s.TotalWriteRequest
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetAvgLoad(v *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.AvgLoad = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetDailyReadRequest(v *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.DailyReadRequest = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetDailyWriteRequest(v *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.DailyWriteRequest = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetMemHeap(v *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.MemHeap = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetNormalAvgLoad(v *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.NormalAvgLoad = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetRegionBalance(v *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.RegionBalance = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetRegionCount(v *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.RegionCount = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetRegionServerCount(v *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.RegionServerCount = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetStoreFileCount(v *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.StoreFileCount = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetTableCount(v *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.TableCount = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetTotalDataSize(v *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.TotalDataSize = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetTotalReadRequest(v *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.TotalReadRequest = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetTotalRequest(v *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.TotalRequest = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) SetTotalWriteRequest(v *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) *GetDoctorHBaseClusterResponseBodyDataMetrics {
	s.TotalWriteRequest = v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetrics) Validate() error {
	if s.AvgLoad != nil {
		if err := s.AvgLoad.Validate(); err != nil {
			return err
		}
	}
	if s.DailyReadRequest != nil {
		if err := s.DailyReadRequest.Validate(); err != nil {
			return err
		}
	}
	if s.DailyWriteRequest != nil {
		if err := s.DailyWriteRequest.Validate(); err != nil {
			return err
		}
	}
	if s.MemHeap != nil {
		if err := s.MemHeap.Validate(); err != nil {
			return err
		}
	}
	if s.NormalAvgLoad != nil {
		if err := s.NormalAvgLoad.Validate(); err != nil {
			return err
		}
	}
	if s.RegionBalance != nil {
		if err := s.RegionBalance.Validate(); err != nil {
			return err
		}
	}
	if s.RegionCount != nil {
		if err := s.RegionCount.Validate(); err != nil {
			return err
		}
	}
	if s.RegionServerCount != nil {
		if err := s.RegionServerCount.Validate(); err != nil {
			return err
		}
	}
	if s.StoreFileCount != nil {
		if err := s.StoreFileCount.Validate(); err != nil {
			return err
		}
	}
	if s.TableCount != nil {
		if err := s.TableCount.Validate(); err != nil {
			return err
		}
	}
	if s.TotalDataSize != nil {
		if err := s.TotalDataSize.Validate(); err != nil {
			return err
		}
	}
	if s.TotalReadRequest != nil {
		if err := s.TotalReadRequest.Validate(); err != nil {
			return err
		}
	}
	if s.TotalRequest != nil {
		if err := s.TotalRequest.Validate(); err != nil {
			return err
		}
	}
	if s.TotalWriteRequest != nil {
		if err := s.TotalWriteRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad struct {
	// The description of the metric.
	//
	// example:
	//
	// The average load under normal working conditions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// avgLoad
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
	// 36.0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) SetValue(v float32) *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsAvgLoad) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of read requests per day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
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
	// 430
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyReadRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of write requests per day
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
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
	// 128
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsDailyWriteRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap struct {
	// The description of the metric.
	//
	// example:
	//
	// Memory heap usage in megabytes (MB)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// memHeap
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
	// 240
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsMemHeap) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad struct {
	// The description of the metric.
	//
	// example:
	//
	// The average load under normal working conditions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// normalAvgLoad
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
	// 526.4
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) SetValue(v float32) *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsNormalAvgLoad) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance struct {
	// The description of the metric.
	//
	// example:
	//
	// The ability to evenly distribute Regions on different RegionServer nodes
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// regionBalance
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
	// 1.0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) SetValue(v float32) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionBalance) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of regions count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// regionCount
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
	// 161
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of region servers count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// regionServerCount
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
	// 6
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsRegionServerCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of store files
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
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
	// 298
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsStoreFileCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsTableCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Number of tables
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// tableCount
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
	// 10
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTableCount) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize struct {
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
	// 256
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalDataSize) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Total number of read requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalReadRequest
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
	// 430
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalReadRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Total number of requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalRequest
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
	// 576
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalRequest) Validate() error {
	return dara.Validate(s)
}

type GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest struct {
	// The description of the metric.
	//
	// example:
	//
	// Total number of write requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// totalWriteRequest
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
	// 520
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) GetName() *string {
	return s.Name
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) SetDescription(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest {
	s.Description = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) SetName(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest {
	s.Name = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) SetUnit(v string) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest {
	s.Unit = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) SetValue(v int64) *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest {
	s.Value = &v
	return s
}

func (s *GetDoctorHBaseClusterResponseBodyDataMetricsTotalWriteRequest) Validate() error {
	return dara.Validate(s)
}
