// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateZonesRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ActivateZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateZonesRequest) GoString() string {
	return s.String()
}

func (s *ActivateZonesRequest) SetClientToken(v string) *ActivateZonesRequest {
	s.ClientToken = &v
	return s
}

type ActivateZonesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ActivateZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateZonesResponseBody) SetRequestId(v string) *ActivateZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateZonesResponseBody) SetResult(v bool) *ActivateZonesResponseBody {
	s.Result = &v
	return s
}

type ActivateZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ActivateZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivateZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateZonesResponse) GoString() string {
	return s.String()
}

func (s *ActivateZonesResponse) SetHeaders(v map[string]*string) *ActivateZonesResponse {
	s.Headers = v
	return s
}

func (s *ActivateZonesResponse) SetBody(v *ActivateZonesResponseBody) *ActivateZonesResponse {
	s.Body = v
	return s
}

type AddConnectableClusterRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s AddConnectableClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s AddConnectableClusterRequest) GoString() string {
	return s.String()
}

func (s *AddConnectableClusterRequest) SetClientToken(v string) *AddConnectableClusterRequest {
	s.ClientToken = &v
	return s
}

type AddConnectableClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AddConnectableClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddConnectableClusterResponseBody) GoString() string {
	return s.String()
}

func (s *AddConnectableClusterResponseBody) SetRequestId(v string) *AddConnectableClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddConnectableClusterResponseBody) SetResult(v bool) *AddConnectableClusterResponseBody {
	s.Result = &v
	return s
}

type AddConnectableClusterResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddConnectableClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddConnectableClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s AddConnectableClusterResponse) GoString() string {
	return s.String()
}

func (s *AddConnectableClusterResponse) SetHeaders(v map[string]*string) *AddConnectableClusterResponse {
	s.Headers = v
	return s
}

func (s *AddConnectableClusterResponse) SetBody(v *AddConnectableClusterResponseBody) *AddConnectableClusterResponse {
	s.Body = v
	return s
}

type AddSnapshotRepoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AddSnapshotRepoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSnapshotRepoResponseBody) GoString() string {
	return s.String()
}

func (s *AddSnapshotRepoResponseBody) SetRequestId(v string) *AddSnapshotRepoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSnapshotRepoResponseBody) SetResult(v bool) *AddSnapshotRepoResponseBody {
	s.Result = &v
	return s
}

type AddSnapshotRepoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSnapshotRepoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSnapshotRepoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSnapshotRepoResponse) GoString() string {
	return s.String()
}

func (s *AddSnapshotRepoResponse) SetHeaders(v map[string]*string) *AddSnapshotRepoResponse {
	s.Headers = v
	return s
}

func (s *AddSnapshotRepoResponse) SetBody(v *AddSnapshotRepoResponseBody) *AddSnapshotRepoResponse {
	s.Body = v
	return s
}

type CancelDeletionRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CancelDeletionRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDeletionRequest) GoString() string {
	return s.String()
}

func (s *CancelDeletionRequest) SetClientToken(v string) *CancelDeletionRequest {
	s.ClientToken = &v
	return s
}

type CancelDeletionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelDeletionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDeletionResponseBody) SetRequestId(v string) *CancelDeletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDeletionResponseBody) SetResult(v bool) *CancelDeletionResponseBody {
	s.Result = &v
	return s
}

type CancelDeletionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelDeletionResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDeletionResponse) GoString() string {
	return s.String()
}

func (s *CancelDeletionResponse) SetHeaders(v map[string]*string) *CancelDeletionResponse {
	s.Headers = v
	return s
}

func (s *CancelDeletionResponse) SetBody(v *CancelDeletionResponseBody) *CancelDeletionResponse {
	s.Body = v
	return s
}

type CancelLogstashDeletionRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CancelLogstashDeletionRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelLogstashDeletionRequest) GoString() string {
	return s.String()
}

func (s *CancelLogstashDeletionRequest) SetClientToken(v string) *CancelLogstashDeletionRequest {
	s.ClientToken = &v
	return s
}

type CancelLogstashDeletionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelLogstashDeletionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelLogstashDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelLogstashDeletionResponseBody) SetRequestId(v string) *CancelLogstashDeletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelLogstashDeletionResponseBody) SetResult(v bool) *CancelLogstashDeletionResponseBody {
	s.Result = &v
	return s
}

type CancelLogstashDeletionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelLogstashDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelLogstashDeletionResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelLogstashDeletionResponse) GoString() string {
	return s.String()
}

func (s *CancelLogstashDeletionResponse) SetHeaders(v map[string]*string) *CancelLogstashDeletionResponse {
	s.Headers = v
	return s
}

func (s *CancelLogstashDeletionResponse) SetBody(v *CancelLogstashDeletionResponseBody) *CancelLogstashDeletionResponse {
	s.Body = v
	return s
}

type CancelTaskRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	TaskType    *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CancelTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTaskRequest) SetClientToken(v string) *CancelTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelTaskRequest) SetTaskType(v string) *CancelTaskRequest {
	s.TaskType = &v
	return s
}

type CancelTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTaskResponseBody) SetResult(v bool) *CancelTaskResponseBody {
	s.Result = &v
	return s
}

type CancelTaskResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelTaskResponse) SetHeaders(v map[string]*string) *CancelTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelTaskResponse) SetBody(v *CancelTaskResponseBody) *CancelTaskResponse {
	s.Body = v
	return s
}

type CapacityPlanRequest struct {
	ComplexQueryAvailable *bool                          `json:"complexQueryAvailable,omitempty" xml:"complexQueryAvailable,omitempty"`
	DataInfo              []*CapacityPlanRequestDataInfo `json:"dataInfo,omitempty" xml:"dataInfo,omitempty" type:"Repeated"`
	Metric                []*CapacityPlanRequestMetric   `json:"metric,omitempty" xml:"metric,omitempty" type:"Repeated"`
	UsageScenario         *string                        `json:"usageScenario,omitempty" xml:"usageScenario,omitempty"`
}

func (s CapacityPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacityPlanRequest) GoString() string {
	return s.String()
}

func (s *CapacityPlanRequest) SetComplexQueryAvailable(v bool) *CapacityPlanRequest {
	s.ComplexQueryAvailable = &v
	return s
}

func (s *CapacityPlanRequest) SetDataInfo(v []*CapacityPlanRequestDataInfo) *CapacityPlanRequest {
	s.DataInfo = v
	return s
}

func (s *CapacityPlanRequest) SetMetric(v []*CapacityPlanRequestMetric) *CapacityPlanRequest {
	s.Metric = v
	return s
}

func (s *CapacityPlanRequest) SetUsageScenario(v string) *CapacityPlanRequest {
	s.UsageScenario = &v
	return s
}

type CapacityPlanRequestDataInfo struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	Size       *int64  `json:"size,omitempty" xml:"size,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	Unit       *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s CapacityPlanRequestDataInfo) String() string {
	return tea.Prettify(s)
}

func (s CapacityPlanRequestDataInfo) GoString() string {
	return s.String()
}

func (s *CapacityPlanRequestDataInfo) SetCode(v string) *CapacityPlanRequestDataInfo {
	s.Code = &v
	return s
}

func (s *CapacityPlanRequestDataInfo) SetSize(v int64) *CapacityPlanRequestDataInfo {
	s.Size = &v
	return s
}

func (s *CapacityPlanRequestDataInfo) SetTotalCount(v int32) *CapacityPlanRequestDataInfo {
	s.TotalCount = &v
	return s
}

func (s *CapacityPlanRequestDataInfo) SetType(v string) *CapacityPlanRequestDataInfo {
	s.Type = &v
	return s
}

func (s *CapacityPlanRequestDataInfo) SetUnit(v string) *CapacityPlanRequestDataInfo {
	s.Unit = &v
	return s
}

type CapacityPlanRequestMetric struct {
	AverageQps   *int32  `json:"averageQps,omitempty" xml:"averageQps,omitempty"`
	Code         *string `json:"code,omitempty" xml:"code,omitempty"`
	Concurrent   *int64  `json:"concurrent,omitempty" xml:"concurrent,omitempty"`
	PeakQps      *int32  `json:"peakQps,omitempty" xml:"peakQps,omitempty"`
	ResponseTime *int32  `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
	Throughput   *int64  `json:"throughput,omitempty" xml:"throughput,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CapacityPlanRequestMetric) String() string {
	return tea.Prettify(s)
}

func (s CapacityPlanRequestMetric) GoString() string {
	return s.String()
}

func (s *CapacityPlanRequestMetric) SetAverageQps(v int32) *CapacityPlanRequestMetric {
	s.AverageQps = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetCode(v string) *CapacityPlanRequestMetric {
	s.Code = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetConcurrent(v int64) *CapacityPlanRequestMetric {
	s.Concurrent = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetPeakQps(v int32) *CapacityPlanRequestMetric {
	s.PeakQps = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetResponseTime(v int32) *CapacityPlanRequestMetric {
	s.ResponseTime = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetThroughput(v int64) *CapacityPlanRequestMetric {
	s.Throughput = &v
	return s
}

func (s *CapacityPlanRequestMetric) SetType(v string) *CapacityPlanRequestMetric {
	s.Type = &v
	return s
}

type CapacityPlanResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CapacityPlanResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CapacityPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CapacityPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponseBody) SetRequestId(v string) *CapacityPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CapacityPlanResponseBody) SetResult(v *CapacityPlanResponseBodyResult) *CapacityPlanResponseBody {
	s.Result = v
	return s
}

type CapacityPlanResponseBodyResult struct {
	ExtendConfigs      []*CapacityPlanResponseBodyResultExtendConfigs      `json:"ExtendConfigs,omitempty" xml:"ExtendConfigs,omitempty" type:"Repeated"`
	InstanceCategory   *string                                             `json:"InstanceCategory,omitempty" xml:"InstanceCategory,omitempty"`
	NodeConfigurations []*CapacityPlanResponseBodyResultNodeConfigurations `json:"NodeConfigurations,omitempty" xml:"NodeConfigurations,omitempty" type:"Repeated"`
	OversizedCluster   *bool                                               `json:"OversizedCluster,omitempty" xml:"OversizedCluster,omitempty"`
}

func (s CapacityPlanResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CapacityPlanResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponseBodyResult) SetExtendConfigs(v []*CapacityPlanResponseBodyResultExtendConfigs) *CapacityPlanResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *CapacityPlanResponseBodyResult) SetInstanceCategory(v string) *CapacityPlanResponseBodyResult {
	s.InstanceCategory = &v
	return s
}

func (s *CapacityPlanResponseBodyResult) SetNodeConfigurations(v []*CapacityPlanResponseBodyResultNodeConfigurations) *CapacityPlanResponseBodyResult {
	s.NodeConfigurations = v
	return s
}

func (s *CapacityPlanResponseBodyResult) SetOversizedCluster(v bool) *CapacityPlanResponseBodyResult {
	s.OversizedCluster = &v
	return s
}

type CapacityPlanResponseBodyResultExtendConfigs struct {
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	Disk       *int64  `json:"Disk,omitempty" xml:"Disk,omitempty"`
	DiskType   *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s CapacityPlanResponseBodyResultExtendConfigs) String() string {
	return tea.Prettify(s)
}

func (s CapacityPlanResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) SetConfigType(v string) *CapacityPlanResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) SetDisk(v int64) *CapacityPlanResponseBodyResultExtendConfigs {
	s.Disk = &v
	return s
}

func (s *CapacityPlanResponseBodyResultExtendConfigs) SetDiskType(v string) *CapacityPlanResponseBodyResultExtendConfigs {
	s.DiskType = &v
	return s
}

type CapacityPlanResponseBodyResultNodeConfigurations struct {
	Amount   *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Cpu      *int64  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Disk     *int64  `json:"Disk,omitempty" xml:"Disk,omitempty"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	Memory   *int64  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s CapacityPlanResponseBodyResultNodeConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CapacityPlanResponseBodyResultNodeConfigurations) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetAmount(v int64) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.Amount = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetCpu(v int64) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.Cpu = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetDisk(v int64) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.Disk = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetDiskType(v string) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.DiskType = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetMemory(v int64) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.Memory = &v
	return s
}

func (s *CapacityPlanResponseBodyResultNodeConfigurations) SetNodeType(v string) *CapacityPlanResponseBodyResultNodeConfigurations {
	s.NodeType = &v
	return s
}

type CapacityPlanResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CapacityPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CapacityPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacityPlanResponse) GoString() string {
	return s.String()
}

func (s *CapacityPlanResponse) SetHeaders(v map[string]*string) *CapacityPlanResponse {
	s.Headers = v
	return s
}

func (s *CapacityPlanResponse) SetBody(v *CapacityPlanResponseBody) *CapacityPlanResponse {
	s.Body = v
	return s
}

type CloseDiagnosisRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang        *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s CloseDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *CloseDiagnosisRequest) SetClientToken(v string) *CloseDiagnosisRequest {
	s.ClientToken = &v
	return s
}

func (s *CloseDiagnosisRequest) SetLang(v string) *CloseDiagnosisRequest {
	s.Lang = &v
	return s
}

type CloseDiagnosisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloseDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDiagnosisResponseBody) SetRequestId(v string) *CloseDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseDiagnosisResponseBody) SetResult(v bool) *CloseDiagnosisResponseBody {
	s.Result = &v
	return s
}

type CloseDiagnosisResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *CloseDiagnosisResponse) SetHeaders(v map[string]*string) *CloseDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *CloseDiagnosisResponse) SetBody(v *CloseDiagnosisResponseBody) *CloseDiagnosisResponse {
	s.Body = v
	return s
}

type CloseHttpsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CloseHttpsRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseHttpsRequest) GoString() string {
	return s.String()
}

func (s *CloseHttpsRequest) SetClientToken(v string) *CloseHttpsRequest {
	s.ClientToken = &v
	return s
}

type CloseHttpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloseHttpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseHttpsResponseBody) GoString() string {
	return s.String()
}

func (s *CloseHttpsResponseBody) SetRequestId(v string) *CloseHttpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseHttpsResponseBody) SetResult(v bool) *CloseHttpsResponseBody {
	s.Result = &v
	return s
}

type CloseHttpsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseHttpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseHttpsResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseHttpsResponse) GoString() string {
	return s.String()
}

func (s *CloseHttpsResponse) SetHeaders(v map[string]*string) *CloseHttpsResponse {
	s.Headers = v
	return s
}

func (s *CloseHttpsResponse) SetBody(v *CloseHttpsResponseBody) *CloseHttpsResponse {
	s.Body = v
	return s
}

type CloseManagedIndexRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CloseManagedIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseManagedIndexRequest) GoString() string {
	return s.String()
}

func (s *CloseManagedIndexRequest) SetClientToken(v string) *CloseManagedIndexRequest {
	s.ClientToken = &v
	return s
}

type CloseManagedIndexResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloseManagedIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseManagedIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CloseManagedIndexResponseBody) SetRequestId(v string) *CloseManagedIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseManagedIndexResponseBody) SetResult(v bool) *CloseManagedIndexResponseBody {
	s.Result = &v
	return s
}

type CloseManagedIndexResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseManagedIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseManagedIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseManagedIndexResponse) GoString() string {
	return s.String()
}

func (s *CloseManagedIndexResponse) SetHeaders(v map[string]*string) *CloseManagedIndexResponse {
	s.Headers = v
	return s
}

func (s *CloseManagedIndexResponse) SetBody(v *CloseManagedIndexResponseBody) *CloseManagedIndexResponse {
	s.Body = v
	return s
}

type CreateCollectorRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectorRequest) GoString() string {
	return s.String()
}

func (s *CreateCollectorRequest) SetClientToken(v string) *CreateCollectorRequest {
	s.ClientToken = &v
	return s
}

type CreateCollectorResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateCollectorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCollectorResponseBody) SetRequestId(v string) *CreateCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCollectorResponseBody) SetResult(v *CreateCollectorResponseBodyResult) *CreateCollectorResponseBody {
	s.Result = v
	return s
}

type CreateCollectorResponseBodyResult struct {
	ResId *string `json:"resId,omitempty" xml:"resId,omitempty"`
}

func (s CreateCollectorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCollectorResponseBodyResult) SetResId(v string) *CreateCollectorResponseBodyResult {
	s.ResId = &v
	return s
}

type CreateCollectorResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCollectorResponse) GoString() string {
	return s.String()
}

func (s *CreateCollectorResponse) SetHeaders(v map[string]*string) *CreateCollectorResponse {
	s.Headers = v
	return s
}

func (s *CreateCollectorResponse) SetBody(v *CreateCollectorResponseBody) *CreateCollectorResponse {
	s.Body = v
	return s
}

type CreateComponentIndexRequest struct {
	Meta     map[string]interface{}               `json:"_meta,omitempty" xml:"_meta,omitempty"`
	Template *CreateComponentIndexRequestTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s CreateComponentIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateComponentIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateComponentIndexRequest) SetMeta(v map[string]interface{}) *CreateComponentIndexRequest {
	s.Meta = v
	return s
}

func (s *CreateComponentIndexRequest) SetTemplate(v *CreateComponentIndexRequestTemplate) *CreateComponentIndexRequest {
	s.Template = v
	return s
}

type CreateComponentIndexRequestTemplate struct {
	Aliases  map[string]interface{} `json:"aliases,omitempty" xml:"aliases,omitempty"`
	Mappings map[string]interface{} `json:"mappings,omitempty" xml:"mappings,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s CreateComponentIndexRequestTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreateComponentIndexRequestTemplate) GoString() string {
	return s.String()
}

func (s *CreateComponentIndexRequestTemplate) SetAliases(v map[string]interface{}) *CreateComponentIndexRequestTemplate {
	s.Aliases = v
	return s
}

func (s *CreateComponentIndexRequestTemplate) SetMappings(v map[string]interface{}) *CreateComponentIndexRequestTemplate {
	s.Mappings = v
	return s
}

func (s *CreateComponentIndexRequestTemplate) SetSettings(v map[string]interface{}) *CreateComponentIndexRequestTemplate {
	s.Settings = v
	return s
}

type CreateComponentIndexResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateComponentIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateComponentIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComponentIndexResponseBody) SetRequestId(v string) *CreateComponentIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComponentIndexResponseBody) SetResult(v bool) *CreateComponentIndexResponseBody {
	s.Result = &v
	return s
}

type CreateComponentIndexResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateComponentIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateComponentIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateComponentIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateComponentIndexResponse) SetHeaders(v map[string]*string) *CreateComponentIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateComponentIndexResponse) SetBody(v *CreateComponentIndexResponseBody) *CreateComponentIndexResponse {
	s.Body = v
	return s
}

type CreateDataStreamRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateDataStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataStreamRequest) GoString() string {
	return s.String()
}

func (s *CreateDataStreamRequest) SetClientToken(v string) *CreateDataStreamRequest {
	s.ClientToken = &v
	return s
}

type CreateDataStreamResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateDataStreamResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateDataStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataStreamResponseBody) SetRequestId(v string) *CreateDataStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataStreamResponseBody) SetResult(v *CreateDataStreamResponseBodyResult) *CreateDataStreamResponseBody {
	s.Result = v
	return s
}

type CreateDataStreamResponseBodyResult struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateDataStreamResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateDataStreamResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDataStreamResponseBodyResult) SetName(v string) *CreateDataStreamResponseBodyResult {
	s.Name = &v
	return s
}

type CreateDataStreamResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataStreamResponse) GoString() string {
	return s.String()
}

func (s *CreateDataStreamResponse) SetHeaders(v map[string]*string) *CreateDataStreamResponse {
	s.Headers = v
	return s
}

func (s *CreateDataStreamResponse) SetBody(v *CreateDataStreamResponseBody) *CreateDataStreamResponse {
	s.Body = v
	return s
}

type CreateDataTasksRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateDataTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataTasksRequest) GoString() string {
	return s.String()
}

func (s *CreateDataTasksRequest) SetClientToken(v string) *CreateDataTasksRequest {
	s.ClientToken = &v
	return s
}

type CreateDataTasksResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*CreateDataTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s CreateDataTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataTasksResponseBody) SetRequestId(v string) *CreateDataTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataTasksResponseBody) SetResult(v []*CreateDataTasksResponseBodyResult) *CreateDataTasksResponseBody {
	s.Result = v
	return s
}

type CreateDataTasksResponseBodyResult struct {
	SinkCluster   *CreateDataTasksResponseBodyResultSinkCluster   `json:"sinkCluster,omitempty" xml:"sinkCluster,omitempty" type:"Struct"`
	SourceCluster *CreateDataTasksResponseBodyResultSourceCluster `json:"sourceCluster,omitempty" xml:"sourceCluster,omitempty" type:"Struct"`
}

func (s CreateDataTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateDataTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDataTasksResponseBodyResult) SetSinkCluster(v *CreateDataTasksResponseBodyResultSinkCluster) *CreateDataTasksResponseBodyResult {
	s.SinkCluster = v
	return s
}

func (s *CreateDataTasksResponseBodyResult) SetSourceCluster(v *CreateDataTasksResponseBodyResultSourceCluster) *CreateDataTasksResponseBodyResult {
	s.SourceCluster = v
	return s
}

type CreateDataTasksResponseBodyResultSinkCluster struct {
	DataSourceType  *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	Index           *string `json:"index,omitempty" xml:"index,omitempty"`
	Mapping         *string `json:"mapping,omitempty" xml:"mapping,omitempty"`
	Password        *string `json:"password,omitempty" xml:"password,omitempty"`
	Routing         *string `json:"routing,omitempty" xml:"routing,omitempty"`
	Settings        *string `json:"settings,omitempty" xml:"settings,omitempty"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	Username        *string `json:"username,omitempty" xml:"username,omitempty"`
	VpcId           *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VpcInstanceId   *string `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
	VpcInstancePort *string `json:"vpcInstancePort,omitempty" xml:"vpcInstancePort,omitempty"`
}

func (s CreateDataTasksResponseBodyResultSinkCluster) String() string {
	return tea.Prettify(s)
}

func (s CreateDataTasksResponseBodyResultSinkCluster) GoString() string {
	return s.String()
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetDataSourceType(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.DataSourceType = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetIndex(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.Index = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetMapping(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.Mapping = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetPassword(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.Password = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetRouting(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.Routing = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetSettings(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.Settings = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetType(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.Type = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetUsername(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.Username = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetVpcId(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.VpcId = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetVpcInstanceId(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.VpcInstanceId = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSinkCluster) SetVpcInstancePort(v string) *CreateDataTasksResponseBodyResultSinkCluster {
	s.VpcInstancePort = &v
	return s
}

type CreateDataTasksResponseBodyResultSourceCluster struct {
	DataSourceType  *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Index           *string `json:"index,omitempty" xml:"index,omitempty"`
	Password        *string `json:"password,omitempty" xml:"password,omitempty"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	Username        *string `json:"username,omitempty" xml:"username,omitempty"`
	VpcId           *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VpcInstanceId   *string `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
	VpcInstancePort *int32  `json:"vpcInstancePort,omitempty" xml:"vpcInstancePort,omitempty"`
}

func (s CreateDataTasksResponseBodyResultSourceCluster) String() string {
	return tea.Prettify(s)
}

func (s CreateDataTasksResponseBodyResultSourceCluster) GoString() string {
	return s.String()
}

func (s *CreateDataTasksResponseBodyResultSourceCluster) SetDataSourceType(v string) *CreateDataTasksResponseBodyResultSourceCluster {
	s.DataSourceType = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSourceCluster) SetEndpoint(v string) *CreateDataTasksResponseBodyResultSourceCluster {
	s.Endpoint = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSourceCluster) SetIndex(v string) *CreateDataTasksResponseBodyResultSourceCluster {
	s.Index = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSourceCluster) SetPassword(v string) *CreateDataTasksResponseBodyResultSourceCluster {
	s.Password = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSourceCluster) SetType(v string) *CreateDataTasksResponseBodyResultSourceCluster {
	s.Type = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSourceCluster) SetUsername(v string) *CreateDataTasksResponseBodyResultSourceCluster {
	s.Username = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSourceCluster) SetVpcId(v string) *CreateDataTasksResponseBodyResultSourceCluster {
	s.VpcId = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSourceCluster) SetVpcInstanceId(v string) *CreateDataTasksResponseBodyResultSourceCluster {
	s.VpcInstanceId = &v
	return s
}

func (s *CreateDataTasksResponseBodyResultSourceCluster) SetVpcInstancePort(v int32) *CreateDataTasksResponseBodyResultSourceCluster {
	s.VpcInstancePort = &v
	return s
}

type CreateDataTasksResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataTasksResponse) GoString() string {
	return s.String()
}

func (s *CreateDataTasksResponse) SetHeaders(v map[string]*string) *CreateDataTasksResponse {
	s.Headers = v
	return s
}

func (s *CreateDataTasksResponse) SetBody(v *CreateDataTasksResponseBody) *CreateDataTasksResponse {
	s.Body = v
	return s
}

type CreateILMPolicyRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateILMPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateILMPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateILMPolicyRequest) SetClientToken(v string) *CreateILMPolicyRequest {
	s.ClientToken = &v
	return s
}

type CreateILMPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateILMPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateILMPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateILMPolicyResponseBody) SetRequestId(v string) *CreateILMPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateILMPolicyResponseBody) SetResult(v string) *CreateILMPolicyResponseBody {
	s.Result = &v
	return s
}

type CreateILMPolicyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateILMPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateILMPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateILMPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateILMPolicyResponse) SetHeaders(v map[string]*string) *CreateILMPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateILMPolicyResponse) SetBody(v *CreateILMPolicyResponseBody) *CreateILMPolicyResponse {
	s.Body = v
	return s
}

type CreateIndexTemplateRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateIndexTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexTemplateRequest) SetClientToken(v string) *CreateIndexTemplateRequest {
	s.ClientToken = &v
	return s
}

type CreateIndexTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateIndexTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndexTemplateResponseBody) SetRequestId(v string) *CreateIndexTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndexTemplateResponseBody) SetResult(v string) *CreateIndexTemplateResponseBody {
	s.Result = &v
	return s
}

type CreateIndexTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIndexTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIndexTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateIndexTemplateResponse) SetHeaders(v map[string]*string) *CreateIndexTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateIndexTemplateResponse) SetBody(v *CreateIndexTemplateResponseBody) *CreateIndexTemplateResponse {
	s.Body = v
	return s
}

type CreateLogstashRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateLogstashRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLogstashRequest) GoString() string {
	return s.String()
}

func (s *CreateLogstashRequest) SetClientToken(v string) *CreateLogstashRequest {
	s.ClientToken = &v
	return s
}

type CreateLogstashResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateLogstashResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateLogstashResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogstashResponseBody) SetRequestId(v string) *CreateLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogstashResponseBody) SetResult(v *CreateLogstashResponseBodyResult) *CreateLogstashResponseBody {
	s.Result = v
	return s
}

type CreateLogstashResponseBodyResult struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s CreateLogstashResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateLogstashResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateLogstashResponseBodyResult) SetInstanceId(v string) *CreateLogstashResponseBodyResult {
	s.InstanceId = &v
	return s
}

type CreateLogstashResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLogstashResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLogstashResponse) GoString() string {
	return s.String()
}

func (s *CreateLogstashResponse) SetHeaders(v map[string]*string) *CreateLogstashResponse {
	s.Headers = v
	return s
}

func (s *CreateLogstashResponse) SetBody(v *CreateLogstashResponseBody) *CreateLogstashResponse {
	s.Body = v
	return s
}

type CreatePipelinesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Trigger     *bool   `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s CreatePipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelinesRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelinesRequest) SetClientToken(v string) *CreatePipelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePipelinesRequest) SetTrigger(v bool) *CreatePipelinesRequest {
	s.Trigger = &v
	return s
}

type CreatePipelinesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreatePipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelinesResponseBody) SetRequestId(v string) *CreatePipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelinesResponseBody) SetResult(v bool) *CreatePipelinesResponseBody {
	s.Result = &v
	return s
}

type CreatePipelinesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelinesResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelinesResponse) SetHeaders(v map[string]*string) *CreatePipelinesResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelinesResponse) SetBody(v *CreatePipelinesResponseBody) *CreatePipelinesResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetClientToken(v string) *CreateSnapshotRequest {
	s.ClientToken = &v
	return s
}

type CreateSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetResult(v bool) *CreateSnapshotResponseBody {
	s.Result = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type CreateVpcEndpointRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 终端节点服务ID
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// 可用区ID
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
	DryRun *bool   `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointRequest) SetClientToken(v string) *CreateVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetServiceId(v string) *CreateVpcEndpointRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetZoneId(v string) *CreateVpcEndpointRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetDryRun(v bool) *CreateVpcEndpointRequest {
	s.DryRun = &v
	return s
}

type CreateVpcEndpointResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateVpcEndpointResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponseBody) SetRequestId(v string) *CreateVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetResult(v *CreateVpcEndpointResponseBodyResult) *CreateVpcEndpointResponseBody {
	s.Result = v
	return s
}

type CreateVpcEndpointResponseBodyResult struct {
	EndpointDomain *string `json:"endpointDomain,omitempty" xml:"endpointDomain,omitempty"`
	EndpointId     *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	EndpointName   *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
	ServiceId      *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s CreateVpcEndpointResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponseBodyResult) SetEndpointDomain(v string) *CreateVpcEndpointResponseBodyResult {
	s.EndpointDomain = &v
	return s
}

func (s *CreateVpcEndpointResponseBodyResult) SetEndpointId(v string) *CreateVpcEndpointResponseBodyResult {
	s.EndpointId = &v
	return s
}

func (s *CreateVpcEndpointResponseBodyResult) SetEndpointName(v string) *CreateVpcEndpointResponseBodyResult {
	s.EndpointName = &v
	return s
}

func (s *CreateVpcEndpointResponseBodyResult) SetServiceId(v string) *CreateVpcEndpointResponseBodyResult {
	s.ServiceId = &v
	return s
}

type CreateVpcEndpointResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponse) SetHeaders(v map[string]*string) *CreateVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcEndpointResponse) SetBody(v *CreateVpcEndpointResponseBody) *CreateVpcEndpointResponse {
	s.Body = v
	return s
}

type DeactivateZonesRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeactivateZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeactivateZonesRequest) GoString() string {
	return s.String()
}

func (s *DeactivateZonesRequest) SetClientToken(v string) *DeactivateZonesRequest {
	s.ClientToken = &v
	return s
}

type DeactivateZonesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeactivateZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeactivateZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateZonesResponseBody) SetRequestId(v string) *DeactivateZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactivateZonesResponseBody) SetResult(v bool) *DeactivateZonesResponseBody {
	s.Result = &v
	return s
}

type DeactivateZonesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeactivateZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeactivateZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeactivateZonesResponse) GoString() string {
	return s.String()
}

func (s *DeactivateZonesResponse) SetHeaders(v map[string]*string) *DeactivateZonesResponse {
	s.Headers = v
	return s
}

func (s *DeactivateZonesResponse) SetBody(v *DeactivateZonesResponseBody) *DeactivateZonesResponse {
	s.Body = v
	return s
}

type DeleteCollectorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectorRequest) GoString() string {
	return s.String()
}

func (s *DeleteCollectorRequest) SetClientToken(v string) *DeleteCollectorRequest {
	s.ClientToken = &v
	return s
}

type DeleteCollectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCollectorResponseBody) SetRequestId(v string) *DeleteCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCollectorResponseBody) SetResult(v bool) *DeleteCollectorResponseBody {
	s.Result = &v
	return s
}

type DeleteCollectorResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCollectorResponse) GoString() string {
	return s.String()
}

func (s *DeleteCollectorResponse) SetHeaders(v map[string]*string) *DeleteCollectorResponse {
	s.Headers = v
	return s
}

func (s *DeleteCollectorResponse) SetBody(v *DeleteCollectorResponseBody) *DeleteCollectorResponse {
	s.Body = v
	return s
}

type DeleteComponentIndexResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteComponentIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComponentIndexResponseBody) SetRequestId(v string) *DeleteComponentIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComponentIndexResponseBody) SetResult(v bool) *DeleteComponentIndexResponseBody {
	s.Result = &v
	return s
}

type DeleteComponentIndexResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteComponentIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteComponentIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteComponentIndexResponse) SetHeaders(v map[string]*string) *DeleteComponentIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteComponentIndexResponse) SetBody(v *DeleteComponentIndexResponseBody) *DeleteComponentIndexResponse {
	s.Body = v
	return s
}

type DeleteConnectedClusterRequest struct {
	ClientToken         *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ConnectedInstanceId *string `json:"connectedInstanceId,omitempty" xml:"connectedInstanceId,omitempty"`
}

func (s DeleteConnectedClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectedClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectedClusterRequest) SetClientToken(v string) *DeleteConnectedClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteConnectedClusterRequest) SetConnectedInstanceId(v string) *DeleteConnectedClusterRequest {
	s.ConnectedInstanceId = &v
	return s
}

type DeleteConnectedClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteConnectedClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectedClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectedClusterResponseBody) SetRequestId(v string) *DeleteConnectedClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConnectedClusterResponseBody) SetResult(v bool) *DeleteConnectedClusterResponseBody {
	s.Result = &v
	return s
}

type DeleteConnectedClusterResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConnectedClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConnectedClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectedClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnectedClusterResponse) SetHeaders(v map[string]*string) *DeleteConnectedClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnectedClusterResponse) SetBody(v *DeleteConnectedClusterResponseBody) *DeleteConnectedClusterResponse {
	s.Body = v
	return s
}

type DeleteDataStreamRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteDataStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataStreamRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataStreamRequest) SetClientToken(v string) *DeleteDataStreamRequest {
	s.ClientToken = &v
	return s
}

type DeleteDataStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteDataStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataStreamResponseBody) SetRequestId(v string) *DeleteDataStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataStreamResponseBody) SetResult(v bool) *DeleteDataStreamResponseBody {
	s.Result = &v
	return s
}

type DeleteDataStreamResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataStreamResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataStreamResponse) SetHeaders(v map[string]*string) *DeleteDataStreamResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataStreamResponse) SetBody(v *DeleteDataStreamResponseBody) *DeleteDataStreamResponse {
	s.Body = v
	return s
}

type DeleteDataTaskRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TaskId      *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteDataTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataTaskRequest) SetClientToken(v string) *DeleteDataTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDataTaskRequest) SetTaskId(v string) *DeleteDataTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteDataTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteDataTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataTaskResponseBody) SetRequestId(v string) *DeleteDataTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataTaskResponseBody) SetResult(v bool) *DeleteDataTaskResponseBody {
	s.Result = &v
	return s
}

type DeleteDataTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDataTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataTaskResponse) SetHeaders(v map[string]*string) *DeleteDataTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataTaskResponse) SetBody(v *DeleteDataTaskResponseBody) *DeleteDataTaskResponse {
	s.Body = v
	return s
}

type DeleteDeprecatedTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteDeprecatedTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeprecatedTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeprecatedTemplateResponseBody) SetRequestId(v string) *DeleteDeprecatedTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeprecatedTemplateResponseBody) SetResult(v bool) *DeleteDeprecatedTemplateResponseBody {
	s.Result = &v
	return s
}

type DeleteDeprecatedTemplateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeprecatedTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeprecatedTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeprecatedTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeprecatedTemplateResponse) SetHeaders(v map[string]*string) *DeleteDeprecatedTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeprecatedTemplateResponse) SetBody(v *DeleteDeprecatedTemplateResponseBody) *DeleteDeprecatedTemplateResponse {
	s.Body = v
	return s
}

type DeleteILMPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteILMPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteILMPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteILMPolicyResponseBody) SetRequestId(v string) *DeleteILMPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteILMPolicyResponseBody) SetResult(v bool) *DeleteILMPolicyResponseBody {
	s.Result = &v
	return s
}

type DeleteILMPolicyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteILMPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteILMPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteILMPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteILMPolicyResponse) SetHeaders(v map[string]*string) *DeleteILMPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteILMPolicyResponse) SetBody(v *DeleteILMPolicyResponseBody) *DeleteILMPolicyResponse {
	s.Body = v
	return s
}

type DeleteIndexTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteIndexTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexTemplateResponseBody) SetRequestId(v string) *DeleteIndexTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexTemplateResponseBody) SetResult(v bool) *DeleteIndexTemplateResponseBody {
	s.Result = &v
	return s
}

type DeleteIndexTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIndexTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIndexTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexTemplateResponse) SetHeaders(v map[string]*string) *DeleteIndexTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexTemplateResponse) SetBody(v *DeleteIndexTemplateResponseBody) *DeleteIndexTemplateResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	DeleteType  *string `json:"deleteType,omitempty" xml:"deleteType,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetClientToken(v string) *DeleteInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteInstanceRequest) SetDeleteType(v string) *DeleteInstanceRequest {
	s.DeleteType = &v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteLogstashRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	DeleteType  *string `json:"deleteType,omitempty" xml:"deleteType,omitempty"`
}

func (s DeleteLogstashRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogstashRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogstashRequest) SetClientToken(v string) *DeleteLogstashRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteLogstashRequest) SetDeleteType(v string) *DeleteLogstashRequest {
	s.DeleteType = &v
	return s
}

type DeleteLogstashResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLogstashResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogstashResponseBody) SetRequestId(v string) *DeleteLogstashResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLogstashResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLogstashResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogstashResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogstashResponse) SetHeaders(v map[string]*string) *DeleteLogstashResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogstashResponse) SetBody(v *DeleteLogstashResponseBody) *DeleteLogstashResponse {
	s.Body = v
	return s
}

type DeletePipelinesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	PipelineIds *string `json:"pipelineIds,omitempty" xml:"pipelineIds,omitempty"`
}

func (s DeletePipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelinesRequest) GoString() string {
	return s.String()
}

func (s *DeletePipelinesRequest) SetClientToken(v string) *DeletePipelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeletePipelinesRequest) SetPipelineIds(v string) *DeletePipelinesRequest {
	s.PipelineIds = &v
	return s
}

type DeletePipelinesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeletePipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelinesResponseBody) SetRequestId(v string) *DeletePipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelinesResponseBody) SetResult(v bool) *DeletePipelinesResponseBody {
	s.Result = &v
	return s
}

type DeletePipelinesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelinesResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelinesResponse) SetHeaders(v map[string]*string) *DeletePipelinesResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelinesResponse) SetBody(v *DeletePipelinesResponseBody) *DeletePipelinesResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRepoRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	RepoPath    *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
}

func (s DeleteSnapshotRepoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRepoRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRepoRequest) SetClientToken(v string) *DeleteSnapshotRepoRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSnapshotRepoRequest) SetRepoPath(v string) *DeleteSnapshotRepoRequest {
	s.RepoPath = &v
	return s
}

type DeleteSnapshotRepoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteSnapshotRepoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRepoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRepoResponseBody) SetRequestId(v string) *DeleteSnapshotRepoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotRepoResponseBody) SetResult(v bool) *DeleteSnapshotRepoResponseBody {
	s.Result = &v
	return s
}

type DeleteSnapshotRepoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSnapshotRepoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSnapshotRepoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRepoResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRepoResponse) SetHeaders(v map[string]*string) *DeleteSnapshotRepoResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotRepoResponse) SetBody(v *DeleteSnapshotRepoResponseBody) *DeleteSnapshotRepoResponse {
	s.Body = v
	return s
}

type DeleteVpcEndpointRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointRequest) SetClientToken(v string) *DeleteVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

type DeleteVpcEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointResponseBody) SetRequestId(v string) *DeleteVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcEndpointResponseBody) SetResult(v bool) *DeleteVpcEndpointResponseBody {
	s.Result = &v
	return s
}

type DeleteVpcEndpointResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointResponse) SetHeaders(v map[string]*string) *DeleteVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcEndpointResponse) SetBody(v *DeleteVpcEndpointResponseBody) *DeleteVpcEndpointResponse {
	s.Body = v
	return s
}

type DescribeAckOperatorResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAckOperatorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAckOperatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAckOperatorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAckOperatorResponseBody) SetRequestId(v string) *DescribeAckOperatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAckOperatorResponseBody) SetResult(v *DescribeAckOperatorResponseBodyResult) *DescribeAckOperatorResponseBody {
	s.Result = v
	return s
}

type DescribeAckOperatorResponseBodyResult struct {
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeAckOperatorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAckOperatorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAckOperatorResponseBodyResult) SetStatus(v string) *DescribeAckOperatorResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAckOperatorResponseBodyResult) SetVersion(v string) *DescribeAckOperatorResponseBodyResult {
	s.Version = &v
	return s
}

type DescribeAckOperatorResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAckOperatorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAckOperatorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAckOperatorResponse) GoString() string {
	return s.String()
}

func (s *DescribeAckOperatorResponse) SetHeaders(v map[string]*string) *DescribeAckOperatorResponse {
	s.Headers = v
	return s
}

func (s *DescribeAckOperatorResponse) SetBody(v *DescribeAckOperatorResponseBody) *DescribeAckOperatorResponse {
	s.Body = v
	return s
}

type DescribeApmResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeApmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeApmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApmResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApmResponseBody) SetRequestId(v string) *DescribeApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApmResponseBody) SetResult(v *DescribeApmResponseBodyResult) *DescribeApmResponseBody {
	s.Result = v
	return s
}

type DescribeApmResponseBodyResult struct {
	ApmServerDomain  *string `json:"apmServerDomain,omitempty" xml:"apmServerDomain,omitempty"`
	CreatedAt        *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	DeployedReplica  *int64  `json:"deployedReplica,omitempty" xml:"deployedReplica,omitempty"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	EndTime          *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId       *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	NodeAmount       *int64  `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	OutputES         *string `json:"outputES,omitempty" xml:"outputES,omitempty"`
	OutputESUserName *string `json:"outputESUserName,omitempty" xml:"outputESUserName,omitempty"`
	OwnerId          *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	PaymentType      *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	Region           *string `json:"region,omitempty" xml:"region,omitempty"`
	Replica          *int32  `json:"replica,omitempty" xml:"replica,omitempty"`
	ResourceSpec     *string `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
	Version          *string `json:"version,omitempty" xml:"version,omitempty"`
	VpcId            *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea           *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId        *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s DescribeApmResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeApmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeApmResponseBodyResult) SetApmServerDomain(v string) *DescribeApmResponseBodyResult {
	s.ApmServerDomain = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetCreatedAt(v string) *DescribeApmResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetDeployedReplica(v int64) *DescribeApmResponseBodyResult {
	s.DeployedReplica = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetDescription(v string) *DescribeApmResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetEndTime(v int64) *DescribeApmResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetInstanceId(v string) *DescribeApmResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetNodeAmount(v int64) *DescribeApmResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetOutputES(v string) *DescribeApmResponseBodyResult {
	s.OutputES = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetOutputESUserName(v string) *DescribeApmResponseBodyResult {
	s.OutputESUserName = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetOwnerId(v string) *DescribeApmResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetPaymentType(v string) *DescribeApmResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetRegion(v string) *DescribeApmResponseBodyResult {
	s.Region = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetReplica(v int32) *DescribeApmResponseBodyResult {
	s.Replica = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetResourceSpec(v string) *DescribeApmResponseBodyResult {
	s.ResourceSpec = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetStatus(v string) *DescribeApmResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetVersion(v string) *DescribeApmResponseBodyResult {
	s.Version = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetVpcId(v string) *DescribeApmResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetVsArea(v string) *DescribeApmResponseBodyResult {
	s.VsArea = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetVswitchId(v string) *DescribeApmResponseBodyResult {
	s.VswitchId = &v
	return s
}

type DescribeApmResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApmResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApmResponse) GoString() string {
	return s.String()
}

func (s *DescribeApmResponse) SetHeaders(v map[string]*string) *DescribeApmResponse {
	s.Headers = v
	return s
}

func (s *DescribeApmResponse) SetBody(v *DescribeApmResponseBody) *DescribeApmResponse {
	s.Body = v
	return s
}

type DescribeCollectorResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeCollectorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBody) SetRequestId(v string) *DescribeCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCollectorResponseBody) SetResult(v *DescribeCollectorResponseBodyResult) *DescribeCollectorResponseBody {
	s.Result = v
	return s
}

type DescribeCollectorResponseBodyResult struct {
	CollectorPaths []*string                                           `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	Configs        []*DescribeCollectorResponseBodyResultConfigs       `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	DryRun         *bool                                               `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	ExtendConfigs  []*DescribeCollectorResponseBodyResultExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	GmtCreatedTime *string                                             `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	GmtUpdateTime  *string                                             `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	Name           *string                                             `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId        *string                                             `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	ResId          *string                                             `json:"resId,omitempty" xml:"resId,omitempty"`
	ResType        *string                                             `json:"resType,omitempty" xml:"resType,omitempty"`
	ResVersion     *string                                             `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	Status         *string                                             `json:"status,omitempty" xml:"status,omitempty"`
	VpcId          *string                                             `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DescribeCollectorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeCollectorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBodyResult) SetCollectorPaths(v []*string) *DescribeCollectorResponseBodyResult {
	s.CollectorPaths = v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetConfigs(v []*DescribeCollectorResponseBodyResultConfigs) *DescribeCollectorResponseBodyResult {
	s.Configs = v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetDryRun(v bool) *DescribeCollectorResponseBodyResult {
	s.DryRun = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetExtendConfigs(v []*DescribeCollectorResponseBodyResultExtendConfigs) *DescribeCollectorResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetGmtCreatedTime(v string) *DescribeCollectorResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetGmtUpdateTime(v string) *DescribeCollectorResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetName(v string) *DescribeCollectorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetOwnerId(v string) *DescribeCollectorResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetResId(v string) *DescribeCollectorResponseBodyResult {
	s.ResId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetResType(v string) *DescribeCollectorResponseBodyResult {
	s.ResType = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetResVersion(v string) *DescribeCollectorResponseBodyResult {
	s.ResVersion = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetStatus(v string) *DescribeCollectorResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetVpcId(v string) *DescribeCollectorResponseBodyResult {
	s.VpcId = &v
	return s
}

type DescribeCollectorResponseBodyResultConfigs struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s DescribeCollectorResponseBodyResultConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeCollectorResponseBodyResultConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBodyResultConfigs) SetContent(v string) *DescribeCollectorResponseBodyResultConfigs {
	s.Content = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultConfigs) SetFileName(v string) *DescribeCollectorResponseBodyResultConfigs {
	s.FileName = &v
	return s
}

type DescribeCollectorResponseBodyResultExtendConfigs struct {
	ConfigType       *string                                                     `json:"configType,omitempty" xml:"configType,omitempty"`
	EnableMonitoring *bool                                                       `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	GroupId          *string                                                     `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Host             *string                                                     `json:"host,omitempty" xml:"host,omitempty"`
	Hosts            []*string                                                   `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	InstanceId       *string                                                     `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType     *string                                                     `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	KibanaHost       *string                                                     `json:"kibanaHost,omitempty" xml:"kibanaHost,omitempty"`
	Machines         []*DescribeCollectorResponseBodyResultExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	Protocol         *string                                                     `json:"protocol,omitempty" xml:"protocol,omitempty"`
	SuccessPodsCount *string                                                     `json:"successPodsCount,omitempty" xml:"successPodsCount,omitempty"`
	TotalPodsCount   *string                                                     `json:"totalPodsCount,omitempty" xml:"totalPodsCount,omitempty"`
	Type             *string                                                     `json:"type,omitempty" xml:"type,omitempty"`
	UserName         *string                                                     `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DescribeCollectorResponseBodyResultExtendConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeCollectorResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetConfigType(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetEnableMonitoring(v bool) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetGroupId(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetHost(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Host = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetHosts(v []*string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Hosts = v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetInstanceId(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetInstanceType(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetKibanaHost(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.KibanaHost = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetMachines(v []*DescribeCollectorResponseBodyResultExtendConfigsMachines) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Machines = v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetProtocol(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetSuccessPodsCount(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.SuccessPodsCount = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetTotalPodsCount(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.TotalPodsCount = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetType(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Type = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetUserName(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.UserName = &v
	return s
}

type DescribeCollectorResponseBodyResultExtendConfigsMachines struct {
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DescribeCollectorResponseBodyResultExtendConfigsMachines) String() string {
	return tea.Prettify(s)
}

func (s DescribeCollectorResponseBodyResultExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBodyResultExtendConfigsMachines) SetAgentStatus(v string) *DescribeCollectorResponseBodyResultExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigsMachines) SetInstanceId(v string) *DescribeCollectorResponseBodyResultExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

type DescribeCollectorResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCollectorResponse) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponse) SetHeaders(v map[string]*string) *DescribeCollectorResponse {
	s.Headers = v
	return s
}

func (s *DescribeCollectorResponse) SetBody(v *DescribeCollectorResponseBody) *DescribeCollectorResponse {
	s.Body = v
	return s
}

type DescribeComponentIndexResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeComponentIndexResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeComponentIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentIndexResponseBody) SetRequestId(v string) *DescribeComponentIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentIndexResponseBody) SetResult(v *DescribeComponentIndexResponseBodyResult) *DescribeComponentIndexResponseBody {
	s.Result = v
	return s
}

type DescribeComponentIndexResponseBodyResult struct {
	Meta     map[string]interface{}                            `json:"_meta,omitempty" xml:"_meta,omitempty"`
	Template *DescribeComponentIndexResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s DescribeComponentIndexResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentIndexResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeComponentIndexResponseBodyResult) SetMeta(v map[string]interface{}) *DescribeComponentIndexResponseBodyResult {
	s.Meta = v
	return s
}

func (s *DescribeComponentIndexResponseBodyResult) SetTemplate(v *DescribeComponentIndexResponseBodyResultTemplate) *DescribeComponentIndexResponseBodyResult {
	s.Template = v
	return s
}

type DescribeComponentIndexResponseBodyResultTemplate struct {
	Aliases  map[string]interface{} `json:"aliases,omitempty" xml:"aliases,omitempty"`
	Mappings map[string]interface{} `json:"mappings,omitempty" xml:"mappings,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s DescribeComponentIndexResponseBodyResultTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentIndexResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) SetAliases(v map[string]interface{}) *DescribeComponentIndexResponseBodyResultTemplate {
	s.Aliases = v
	return s
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) SetMappings(v map[string]interface{}) *DescribeComponentIndexResponseBodyResultTemplate {
	s.Mappings = v
	return s
}

func (s *DescribeComponentIndexResponseBodyResultTemplate) SetSettings(v map[string]interface{}) *DescribeComponentIndexResponseBodyResultTemplate {
	s.Settings = v
	return s
}

type DescribeComponentIndexResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeComponentIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComponentIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentIndexResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentIndexResponse) SetHeaders(v map[string]*string) *DescribeComponentIndexResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentIndexResponse) SetBody(v *DescribeComponentIndexResponseBody) *DescribeComponentIndexResponse {
	s.Body = v
	return s
}

type DescribeConnectableClustersRequest struct {
	AlreadySetItems *bool `json:"alreadySetItems,omitempty" xml:"alreadySetItems,omitempty"`
}

func (s DescribeConnectableClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectableClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeConnectableClustersRequest) SetAlreadySetItems(v bool) *DescribeConnectableClustersRequest {
	s.AlreadySetItems = &v
	return s
}

type DescribeConnectableClustersResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeConnectableClustersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeConnectableClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectableClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectableClustersResponseBody) SetRequestId(v string) *DescribeConnectableClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectableClustersResponseBody) SetResult(v []*DescribeConnectableClustersResponseBodyResult) *DescribeConnectableClustersResponseBody {
	s.Result = v
	return s
}

type DescribeConnectableClustersResponseBodyResult struct {
	Instances   *string `json:"instances,omitempty" xml:"instances,omitempty"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s DescribeConnectableClustersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectableClustersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeConnectableClustersResponseBodyResult) SetInstances(v string) *DescribeConnectableClustersResponseBodyResult {
	s.Instances = &v
	return s
}

func (s *DescribeConnectableClustersResponseBodyResult) SetNetworkType(v string) *DescribeConnectableClustersResponseBodyResult {
	s.NetworkType = &v
	return s
}

type DescribeConnectableClustersResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConnectableClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConnectableClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectableClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeConnectableClustersResponse) SetHeaders(v map[string]*string) *DescribeConnectableClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeConnectableClustersResponse) SetBody(v *DescribeConnectableClustersResponseBody) *DescribeConnectableClustersResponse {
	s.Body = v
	return s
}

type DescribeDeprecatedTemplateResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDeprecatedTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDeprecatedTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeprecatedTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeprecatedTemplateResponseBody) SetRequestId(v string) *DescribeDeprecatedTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBody) SetResult(v *DescribeDeprecatedTemplateResponseBodyResult) *DescribeDeprecatedTemplateResponseBody {
	s.Result = v
	return s
}

type DescribeDeprecatedTemplateResponseBodyResult struct {
	DataStream    *bool                                                 `json:"dataStream,omitempty" xml:"dataStream,omitempty"`
	IndexPatterns []*string                                             `json:"indexPatterns,omitempty" xml:"indexPatterns,omitempty" type:"Repeated"`
	IndexTemplate *string                                               `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
	Order         *int64                                                `json:"order,omitempty" xml:"order,omitempty"`
	Template      *DescribeDeprecatedTemplateResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	Version       *string                                               `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeDeprecatedTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeprecatedTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetDataStream(v bool) *DescribeDeprecatedTemplateResponseBodyResult {
	s.DataStream = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetIndexPatterns(v []*string) *DescribeDeprecatedTemplateResponseBodyResult {
	s.IndexPatterns = v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetIndexTemplate(v string) *DescribeDeprecatedTemplateResponseBodyResult {
	s.IndexTemplate = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetOrder(v int64) *DescribeDeprecatedTemplateResponseBodyResult {
	s.Order = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetTemplate(v *DescribeDeprecatedTemplateResponseBodyResultTemplate) *DescribeDeprecatedTemplateResponseBodyResult {
	s.Template = v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResult) SetVersion(v string) *DescribeDeprecatedTemplateResponseBodyResult {
	s.Version = &v
	return s
}

type DescribeDeprecatedTemplateResponseBodyResultTemplate struct {
	Aliases  *string `json:"aliases,omitempty" xml:"aliases,omitempty"`
	Mappings *string `json:"mappings,omitempty" xml:"mappings,omitempty"`
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s DescribeDeprecatedTemplateResponseBodyResultTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeprecatedTemplateResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) SetAliases(v string) *DescribeDeprecatedTemplateResponseBodyResultTemplate {
	s.Aliases = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) SetMappings(v string) *DescribeDeprecatedTemplateResponseBodyResultTemplate {
	s.Mappings = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponseBodyResultTemplate) SetSettings(v string) *DescribeDeprecatedTemplateResponseBodyResultTemplate {
	s.Settings = &v
	return s
}

type DescribeDeprecatedTemplateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeprecatedTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeprecatedTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeprecatedTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeprecatedTemplateResponse) SetHeaders(v map[string]*string) *DescribeDeprecatedTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeprecatedTemplateResponse) SetBody(v *DescribeDeprecatedTemplateResponseBody) *DescribeDeprecatedTemplateResponse {
	s.Body = v
	return s
}

type DescribeDiagnoseReportRequest struct {
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s DescribeDiagnoseReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportRequest) SetLang(v string) *DescribeDiagnoseReportRequest {
	s.Lang = &v
	return s
}

type DescribeDiagnoseReportResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDiagnoseReportResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDiagnoseReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBody) SetRequestId(v string) *DescribeDiagnoseReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBody) SetResult(v *DescribeDiagnoseReportResponseBodyResult) *DescribeDiagnoseReportResponseBody {
	s.Result = v
	return s
}

type DescribeDiagnoseReportResponseBodyResult struct {
	CreateTime    *int64                                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiagnoseItems []*DescribeDiagnoseReportResponseBodyResultDiagnoseItems `json:"diagnoseItems,omitempty" xml:"diagnoseItems,omitempty" type:"Repeated"`
	Health        *string                                                  `json:"health,omitempty" xml:"health,omitempty"`
	InstanceId    *string                                                  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ReportId      *string                                                  `json:"reportId,omitempty" xml:"reportId,omitempty"`
	State         *string                                                  `json:"state,omitempty" xml:"state,omitempty"`
	Trigger       *string                                                  `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s DescribeDiagnoseReportResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetCreateTime(v int64) *DescribeDiagnoseReportResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetDiagnoseItems(v []*DescribeDiagnoseReportResponseBodyResultDiagnoseItems) *DescribeDiagnoseReportResponseBodyResult {
	s.DiagnoseItems = v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetHealth(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.Health = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetInstanceId(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetReportId(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetState(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResult) SetTrigger(v string) *DescribeDiagnoseReportResponseBodyResult {
	s.Trigger = &v
	return s
}

type DescribeDiagnoseReportResponseBodyResultDiagnoseItems struct {
	Detail *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
	Health *string                                                      `json:"health,omitempty" xml:"health,omitempty"`
	Item   *string                                                      `json:"item,omitempty" xml:"item,omitempty"`
}

func (s DescribeDiagnoseReportResponseBodyResultDiagnoseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBodyResultDiagnoseItems) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) SetDetail(v *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) *DescribeDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Detail = v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) SetHealth(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Health = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItems) SetItem(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Item = &v
	return s
}

type DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail struct {
	Desc    *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Suggest *string `json:"suggest,omitempty" xml:"suggest,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetDesc(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Desc = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetName(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Name = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetResult(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Result = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetSuggest(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Suggest = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetType(v string) *DescribeDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Type = &v
	return s
}

type DescribeDiagnoseReportResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDiagnoseReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnoseReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnoseReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponse) SetHeaders(v map[string]*string) *DescribeDiagnoseReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnoseReportResponse) SetBody(v *DescribeDiagnoseReportResponseBody) *DescribeDiagnoseReportResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisSettingsRequest struct {
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s DescribeDiagnosisSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSettingsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsRequest) SetLang(v string) *DescribeDiagnosisSettingsRequest {
	s.Lang = &v
	return s
}

type DescribeDiagnosisSettingsResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDiagnosisSettingsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDiagnosisSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsResponseBody) SetRequestId(v string) *DescribeDiagnosisSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBody) SetResult(v *DescribeDiagnosisSettingsResponseBodyResult) *DescribeDiagnosisSettingsResponseBody {
	s.Result = v
	return s
}

type DescribeDiagnosisSettingsResponseBodyResult struct {
	Scene      *string `json:"scene,omitempty" xml:"scene,omitempty"`
	UpdateTime *int64  `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s DescribeDiagnosisSettingsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSettingsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetScene(v string) *DescribeDiagnosisSettingsResponseBodyResult {
	s.Scene = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetUpdateTime(v int64) *DescribeDiagnosisSettingsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

type DescribeDiagnosisSettingsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDiagnosisSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisSettingsResponse) SetBody(v *DescribeDiagnosisSettingsResponseBody) *DescribeDiagnosisSettingsResponse {
	s.Body = v
	return s
}

type DescribeElasticsearchHealthResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeElasticsearchHealthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticsearchHealthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticsearchHealthResponseBody) SetCode(v string) *DescribeElasticsearchHealthResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeElasticsearchHealthResponseBody) SetMessage(v string) *DescribeElasticsearchHealthResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeElasticsearchHealthResponseBody) SetRequestId(v string) *DescribeElasticsearchHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticsearchHealthResponseBody) SetResult(v string) *DescribeElasticsearchHealthResponseBody {
	s.Result = &v
	return s
}

type DescribeElasticsearchHealthResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeElasticsearchHealthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeElasticsearchHealthResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticsearchHealthResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticsearchHealthResponse) SetHeaders(v map[string]*string) *DescribeElasticsearchHealthResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticsearchHealthResponse) SetBody(v *DescribeElasticsearchHealthResponseBody) *DescribeElasticsearchHealthResponse {
	s.Body = v
	return s
}

type DescribeILMPolicyResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeILMPolicyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeILMPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeILMPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeILMPolicyResponseBody) SetRequestId(v string) *DescribeILMPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeILMPolicyResponseBody) SetResult(v *DescribeILMPolicyResponseBodyResult) *DescribeILMPolicyResponseBody {
	s.Result = v
	return s
}

type DescribeILMPolicyResponseBodyResult struct {
	Name   *string                `json:"name,omitempty" xml:"name,omitempty"`
	Phases map[string]interface{} `json:"phases,omitempty" xml:"phases,omitempty"`
}

func (s DescribeILMPolicyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeILMPolicyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeILMPolicyResponseBodyResult) SetName(v string) *DescribeILMPolicyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeILMPolicyResponseBodyResult) SetPhases(v map[string]interface{}) *DescribeILMPolicyResponseBodyResult {
	s.Phases = v
	return s
}

type DescribeILMPolicyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeILMPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeILMPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeILMPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeILMPolicyResponse) SetHeaders(v map[string]*string) *DescribeILMPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeILMPolicyResponse) SetBody(v *DescribeILMPolicyResponseBody) *DescribeILMPolicyResponse {
	s.Body = v
	return s
}

type DescribeIndexTemplateResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeIndexTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeIndexTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIndexTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIndexTemplateResponseBody) SetRequestId(v string) *DescribeIndexTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIndexTemplateResponseBody) SetResult(v *DescribeIndexTemplateResponseBodyResult) *DescribeIndexTemplateResponseBody {
	s.Result = v
	return s
}

type DescribeIndexTemplateResponseBodyResult struct {
	DataStream    *bool                                            `json:"dataStream,omitempty" xml:"dataStream,omitempty"`
	IlmPolicy     *string                                          `json:"ilmPolicy,omitempty" xml:"ilmPolicy,omitempty"`
	IndexPatterns []*string                                        `json:"indexPatterns,omitempty" xml:"indexPatterns,omitempty" type:"Repeated"`
	IndexTemplate *string                                          `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
	Priority      *int32                                           `json:"priority,omitempty" xml:"priority,omitempty"`
	Template      *DescribeIndexTemplateResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s DescribeIndexTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeIndexTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeIndexTemplateResponseBodyResult) SetDataStream(v bool) *DescribeIndexTemplateResponseBodyResult {
	s.DataStream = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetIlmPolicy(v string) *DescribeIndexTemplateResponseBodyResult {
	s.IlmPolicy = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetIndexPatterns(v []*string) *DescribeIndexTemplateResponseBodyResult {
	s.IndexPatterns = v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetIndexTemplate(v string) *DescribeIndexTemplateResponseBodyResult {
	s.IndexTemplate = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetPriority(v int32) *DescribeIndexTemplateResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResult) SetTemplate(v *DescribeIndexTemplateResponseBodyResultTemplate) *DescribeIndexTemplateResponseBodyResult {
	s.Template = v
	return s
}

type DescribeIndexTemplateResponseBodyResultTemplate struct {
	Aliases  *string `json:"aliases,omitempty" xml:"aliases,omitempty"`
	Mappings *string `json:"mappings,omitempty" xml:"mappings,omitempty"`
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s DescribeIndexTemplateResponseBodyResultTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeIndexTemplateResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) SetAliases(v string) *DescribeIndexTemplateResponseBodyResultTemplate {
	s.Aliases = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) SetMappings(v string) *DescribeIndexTemplateResponseBodyResultTemplate {
	s.Mappings = &v
	return s
}

func (s *DescribeIndexTemplateResponseBodyResultTemplate) SetSettings(v string) *DescribeIndexTemplateResponseBodyResultTemplate {
	s.Settings = &v
	return s
}

type DescribeIndexTemplateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIndexTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIndexTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIndexTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeIndexTemplateResponse) SetHeaders(v map[string]*string) *DescribeIndexTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeIndexTemplateResponse) SetBody(v *DescribeIndexTemplateResponseBody) *DescribeIndexTemplateResponse {
	s.Body = v
	return s
}

type DescribeInstanceResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetResult(v *DescribeInstanceResponseBodyResult) *DescribeInstanceResponseBody {
	s.Result = v
	return s
}

type DescribeInstanceResponseBodyResult struct {
	AdvancedDedicateMaster       *bool                                                           `json:"advancedDedicateMaster,omitempty" xml:"advancedDedicateMaster,omitempty"`
	AdvancedSetting              *DescribeInstanceResponseBodyResultAdvancedSetting              `json:"advancedSetting,omitempty" xml:"advancedSetting,omitempty" type:"Struct"`
	AliwsDicts                   []*DescribeInstanceResponseBodyResultAliwsDicts                 `json:"aliwsDicts,omitempty" xml:"aliwsDicts,omitempty" type:"Repeated"`
	ClientNodeConfiguration      *DescribeInstanceResponseBodyResultClientNodeConfiguration      `json:"clientNodeConfiguration,omitempty" xml:"clientNodeConfiguration,omitempty" type:"Struct"`
	CreatedAt                    *string                                                         `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	DedicateMaster               *bool                                                           `json:"dedicateMaster,omitempty" xml:"dedicateMaster,omitempty"`
	Description                  *string                                                         `json:"description,omitempty" xml:"description,omitempty"`
	DictList                     []*DescribeInstanceResponseBodyResultDictList                   `json:"dictList,omitempty" xml:"dictList,omitempty" type:"Repeated"`
	Domain                       *string                                                         `json:"domain,omitempty" xml:"domain,omitempty"`
	ElasticDataNodeConfiguration *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration `json:"elasticDataNodeConfiguration,omitempty" xml:"elasticDataNodeConfiguration,omitempty" type:"Struct"`
	EnableKibanaPrivateNetwork   *bool                                                           `json:"enableKibanaPrivateNetwork,omitempty" xml:"enableKibanaPrivateNetwork,omitempty"`
	EnableKibanaPublicNetwork    *bool                                                           `json:"enableKibanaPublicNetwork,omitempty" xml:"enableKibanaPublicNetwork,omitempty"`
	EnablePublic                 *bool                                                           `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
	EsConfig                     map[string]interface{}                                          `json:"esConfig,omitempty" xml:"esConfig,omitempty"`
	EsIPBlacklist                []*string                                                       `json:"esIPBlacklist,omitempty" xml:"esIPBlacklist,omitempty" type:"Repeated"`
	EsIPWhitelist                []*string                                                       `json:"esIPWhitelist,omitempty" xml:"esIPWhitelist,omitempty" type:"Repeated"`
	EsVersion                    *string                                                         `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	ExtendConfigs                []map[string]interface{}                                        `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	HaveClientNode               *bool                                                           `json:"haveClientNode,omitempty" xml:"haveClientNode,omitempty"`
	HaveKibana                   *bool                                                           `json:"haveKibana,omitempty" xml:"haveKibana,omitempty"`
	InstanceId                   *string                                                         `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsNewDeployment              *bool                                                           `json:"isNewDeployment,omitempty" xml:"isNewDeployment,omitempty"`
	KibanaConfiguration          *DescribeInstanceResponseBodyResultKibanaConfiguration          `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	KibanaDomain                 *string                                                         `json:"kibanaDomain,omitempty" xml:"kibanaDomain,omitempty"`
	KibanaIPWhitelist            []*string                                                       `json:"kibanaIPWhitelist,omitempty" xml:"kibanaIPWhitelist,omitempty" type:"Repeated"`
	KibanaPort                   *int32                                                          `json:"kibanaPort,omitempty" xml:"kibanaPort,omitempty"`
	KibanaPrivateIPWhitelist     []*string                                                       `json:"kibanaPrivateIPWhitelist,omitempty" xml:"kibanaPrivateIPWhitelist,omitempty" type:"Repeated"`
	MasterConfiguration          *DescribeInstanceResponseBodyResultMasterConfiguration          `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	NetworkConfig                *DescribeInstanceResponseBodyResultNetworkConfig                `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	NodeAmount                   *int32                                                          `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec                     *DescribeInstanceResponseBodyResultNodeSpec                     `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	PaymentType                  *string                                                         `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	Port                         *int32                                                          `json:"port,omitempty" xml:"port,omitempty"`
	PostpaidServiceStatus        *string                                                         `json:"postpaidServiceStatus,omitempty" xml:"postpaidServiceStatus,omitempty"`
	PrivateNetworkIpWhiteList    []*string                                                       `json:"privateNetworkIpWhiteList,omitempty" xml:"privateNetworkIpWhiteList,omitempty" type:"Repeated"`
	Protocol                     *string                                                         `json:"protocol,omitempty" xml:"protocol,omitempty"`
	PublicDomain                 *string                                                         `json:"publicDomain,omitempty" xml:"publicDomain,omitempty"`
	PublicIpWhitelist            []*string                                                       `json:"publicIpWhitelist,omitempty" xml:"publicIpWhitelist,omitempty" type:"Repeated"`
	PublicPort                   *int32                                                          `json:"publicPort,omitempty" xml:"publicPort,omitempty"`
	ResourceGroupId              *string                                                         `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ServiceVpc                   *bool                                                           `json:"serviceVpc,omitempty" xml:"serviceVpc,omitempty"`
	Status                       *string                                                         `json:"status,omitempty" xml:"status,omitempty"`
	SynonymsDicts                []*DescribeInstanceResponseBodyResultSynonymsDicts              `json:"synonymsDicts,omitempty" xml:"synonymsDicts,omitempty" type:"Repeated"`
	Tags                         []*DescribeInstanceResponseBodyResultTags                       `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	UpdatedAt                    *string                                                         `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	VpcInstanceId                *string                                                         `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
	WarmNode                     *bool                                                           `json:"warmNode,omitempty" xml:"warmNode,omitempty"`
	WarmNodeConfiguration        *DescribeInstanceResponseBodyResultWarmNodeConfiguration        `json:"warmNodeConfiguration,omitempty" xml:"warmNodeConfiguration,omitempty" type:"Struct"`
	ZoneCount                    *int32                                                          `json:"zoneCount,omitempty" xml:"zoneCount,omitempty"`
	ZoneInfos                    []*DescribeInstanceResponseBodyResultZoneInfos                  `json:"zoneInfos,omitempty" xml:"zoneInfos,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResult) SetAdvancedDedicateMaster(v bool) *DescribeInstanceResponseBodyResult {
	s.AdvancedDedicateMaster = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetAdvancedSetting(v *DescribeInstanceResponseBodyResultAdvancedSetting) *DescribeInstanceResponseBodyResult {
	s.AdvancedSetting = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetAliwsDicts(v []*DescribeInstanceResponseBodyResultAliwsDicts) *DescribeInstanceResponseBodyResult {
	s.AliwsDicts = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetClientNodeConfiguration(v *DescribeInstanceResponseBodyResultClientNodeConfiguration) *DescribeInstanceResponseBodyResult {
	s.ClientNodeConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetCreatedAt(v string) *DescribeInstanceResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetDedicateMaster(v bool) *DescribeInstanceResponseBodyResult {
	s.DedicateMaster = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetDescription(v string) *DescribeInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetDictList(v []*DescribeInstanceResponseBodyResultDictList) *DescribeInstanceResponseBodyResult {
	s.DictList = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetDomain(v string) *DescribeInstanceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetElasticDataNodeConfiguration(v *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) *DescribeInstanceResponseBodyResult {
	s.ElasticDataNodeConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEnableKibanaPrivateNetwork(v bool) *DescribeInstanceResponseBodyResult {
	s.EnableKibanaPrivateNetwork = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEnableKibanaPublicNetwork(v bool) *DescribeInstanceResponseBodyResult {
	s.EnableKibanaPublicNetwork = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEnablePublic(v bool) *DescribeInstanceResponseBodyResult {
	s.EnablePublic = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEsConfig(v map[string]interface{}) *DescribeInstanceResponseBodyResult {
	s.EsConfig = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEsIPBlacklist(v []*string) *DescribeInstanceResponseBodyResult {
	s.EsIPBlacklist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEsIPWhitelist(v []*string) *DescribeInstanceResponseBodyResult {
	s.EsIPWhitelist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEsVersion(v string) *DescribeInstanceResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetExtendConfigs(v []map[string]interface{}) *DescribeInstanceResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetHaveClientNode(v bool) *DescribeInstanceResponseBodyResult {
	s.HaveClientNode = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetHaveKibana(v bool) *DescribeInstanceResponseBodyResult {
	s.HaveKibana = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetInstanceId(v string) *DescribeInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetIsNewDeployment(v bool) *DescribeInstanceResponseBodyResult {
	s.IsNewDeployment = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaConfiguration(v *DescribeInstanceResponseBodyResultKibanaConfiguration) *DescribeInstanceResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaDomain(v string) *DescribeInstanceResponseBodyResult {
	s.KibanaDomain = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaIPWhitelist(v []*string) *DescribeInstanceResponseBodyResult {
	s.KibanaIPWhitelist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaPort(v int32) *DescribeInstanceResponseBodyResult {
	s.KibanaPort = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaPrivateIPWhitelist(v []*string) *DescribeInstanceResponseBodyResult {
	s.KibanaPrivateIPWhitelist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetMasterConfiguration(v *DescribeInstanceResponseBodyResultMasterConfiguration) *DescribeInstanceResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetNetworkConfig(v *DescribeInstanceResponseBodyResultNetworkConfig) *DescribeInstanceResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetNodeAmount(v int32) *DescribeInstanceResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetNodeSpec(v *DescribeInstanceResponseBodyResultNodeSpec) *DescribeInstanceResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPaymentType(v string) *DescribeInstanceResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPort(v int32) *DescribeInstanceResponseBodyResult {
	s.Port = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPostpaidServiceStatus(v string) *DescribeInstanceResponseBodyResult {
	s.PostpaidServiceStatus = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPrivateNetworkIpWhiteList(v []*string) *DescribeInstanceResponseBodyResult {
	s.PrivateNetworkIpWhiteList = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetProtocol(v string) *DescribeInstanceResponseBodyResult {
	s.Protocol = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPublicDomain(v string) *DescribeInstanceResponseBodyResult {
	s.PublicDomain = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPublicIpWhitelist(v []*string) *DescribeInstanceResponseBodyResult {
	s.PublicIpWhitelist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPublicPort(v int32) *DescribeInstanceResponseBodyResult {
	s.PublicPort = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetResourceGroupId(v string) *DescribeInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetServiceVpc(v bool) *DescribeInstanceResponseBodyResult {
	s.ServiceVpc = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetStatus(v string) *DescribeInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetSynonymsDicts(v []*DescribeInstanceResponseBodyResultSynonymsDicts) *DescribeInstanceResponseBodyResult {
	s.SynonymsDicts = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetTags(v []*DescribeInstanceResponseBodyResultTags) *DescribeInstanceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetUpdatedAt(v string) *DescribeInstanceResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetVpcInstanceId(v string) *DescribeInstanceResponseBodyResult {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetWarmNode(v bool) *DescribeInstanceResponseBodyResult {
	s.WarmNode = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetWarmNodeConfiguration(v *DescribeInstanceResponseBodyResultWarmNodeConfiguration) *DescribeInstanceResponseBodyResult {
	s.WarmNodeConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetZoneCount(v int32) *DescribeInstanceResponseBodyResult {
	s.ZoneCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetZoneInfos(v []*DescribeInstanceResponseBodyResultZoneInfos) *DescribeInstanceResponseBodyResult {
	s.ZoneInfos = v
	return s
}

type DescribeInstanceResponseBodyResultAdvancedSetting struct {
	GcName *string `json:"gcName,omitempty" xml:"gcName,omitempty"`
}

func (s DescribeInstanceResponseBodyResultAdvancedSetting) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultAdvancedSetting) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultAdvancedSetting) SetGcName(v string) *DescribeInstanceResponseBodyResultAdvancedSetting {
	s.GcName = &v
	return s
}

type DescribeInstanceResponseBodyResultAliwsDicts struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeInstanceResponseBodyResultAliwsDicts) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultAliwsDicts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) SetFileSize(v int64) *DescribeInstanceResponseBodyResultAliwsDicts {
	s.FileSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) SetName(v string) *DescribeInstanceResponseBodyResultAliwsDicts {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) SetSourceType(v string) *DescribeInstanceResponseBodyResultAliwsDicts {
	s.SourceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) SetType(v string) *DescribeInstanceResponseBodyResultAliwsDicts {
	s.Type = &v
	return s
}

type DescribeInstanceResponseBodyResultClientNodeConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeInstanceResponseBodyResultClientNodeConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultClientNodeConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) SetDisk(v int32) *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) SetDiskType(v string) *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	s.Spec = &v
	return s
}

type DescribeInstanceResponseBodyResultDictList struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeInstanceResponseBodyResultDictList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultDictList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultDictList) SetFileSize(v int64) *DescribeInstanceResponseBodyResultDictList {
	s.FileSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultDictList) SetName(v string) *DescribeInstanceResponseBodyResultDictList {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultDictList) SetSourceType(v string) *DescribeInstanceResponseBodyResultDictList {
	s.SourceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultDictList) SetType(v string) *DescribeInstanceResponseBodyResultDictList {
	s.Type = &v
	return s
}

type DescribeInstanceResponseBodyResultElasticDataNodeConfiguration struct {
	Amount         *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk           *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskEncryption *bool   `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType       *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec           *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetDisk(v int32) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetDiskEncryption(v bool) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetDiskType(v string) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Spec = &v
	return s
}

type DescribeInstanceResponseBodyResultKibanaConfiguration struct {
	Amount *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Spec   *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeInstanceResponseBodyResultKibanaConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

type DescribeInstanceResponseBodyResultMasterConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeInstanceResponseBodyResultMasterConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) SetDisk(v int32) *DescribeInstanceResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) SetDiskType(v string) *DescribeInstanceResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

type DescribeInstanceResponseBodyResultNetworkConfig struct {
	Type             *string                                                            `json:"type,omitempty" xml:"type,omitempty"`
	VpcId            *string                                                            `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea           *string                                                            `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId        *string                                                            `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	WhiteIpGroupList []*DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList `json:"whiteIpGroupList,omitempty" xml:"whiteIpGroupList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyResultNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetType(v string) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetVpcId(v string) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetVsArea(v string) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetVswitchId(v string) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetWhiteIpGroupList(v []*DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.WhiteIpGroupList = v
	return s
}

type DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList struct {
	GroupName   *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips         []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	WhiteIpType *string   `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) SetGroupName(v string) *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.GroupName = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) SetIps(v []*string) *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.Ips = v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) SetWhiteIpType(v string) *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.WhiteIpType = &v
	return s
}

type DescribeInstanceResponseBodyResultNodeSpec struct {
	Disk           *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskEncryption *bool   `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType       *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec           *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeInstanceResponseBodyResultNodeSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetDisk(v int32) *DescribeInstanceResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetDiskEncryption(v bool) *DescribeInstanceResponseBodyResultNodeSpec {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetDiskType(v string) *DescribeInstanceResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetSpec(v string) *DescribeInstanceResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

type DescribeInstanceResponseBodyResultSynonymsDicts struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeInstanceResponseBodyResultSynonymsDicts) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultSynonymsDicts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) SetFileSize(v int64) *DescribeInstanceResponseBodyResultSynonymsDicts {
	s.FileSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) SetName(v string) *DescribeInstanceResponseBodyResultSynonymsDicts {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) SetSourceType(v string) *DescribeInstanceResponseBodyResultSynonymsDicts {
	s.SourceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) SetType(v string) *DescribeInstanceResponseBodyResultSynonymsDicts {
	s.Type = &v
	return s
}

type DescribeInstanceResponseBodyResultTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s DescribeInstanceResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultTags) SetTagKey(v string) *DescribeInstanceResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultTags) SetTagValue(v string) *DescribeInstanceResponseBodyResultTags {
	s.TagValue = &v
	return s
}

type DescribeInstanceResponseBodyResultWarmNodeConfiguration struct {
	Amount         *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk           *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskEncryption *bool   `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType       *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec           *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeInstanceResponseBodyResultWarmNodeConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultWarmNodeConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetDisk(v int32) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetDiskEncryption(v bool) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetDiskType(v string) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.Spec = &v
	return s
}

type DescribeInstanceResponseBodyResultZoneInfos struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s DescribeInstanceResponseBodyResultZoneInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultZoneInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultZoneInfos) SetStatus(v string) *DescribeInstanceResponseBodyResultZoneInfos {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultZoneInfos) SetZoneId(v string) *DescribeInstanceResponseBodyResultZoneInfos {
	s.ZoneId = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeKibanaSettingsResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeKibanaSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKibanaSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKibanaSettingsResponseBody) SetRequestId(v string) *DescribeKibanaSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKibanaSettingsResponseBody) SetResult(v map[string]interface{}) *DescribeKibanaSettingsResponseBody {
	s.Result = v
	return s
}

type DescribeKibanaSettingsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeKibanaSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeKibanaSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKibanaSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeKibanaSettingsResponse) SetHeaders(v map[string]*string) *DescribeKibanaSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeKibanaSettingsResponse) SetBody(v *DescribeKibanaSettingsResponseBody) *DescribeKibanaSettingsResponse {
	s.Body = v
	return s
}

type DescribeLogstashResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeLogstashResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeLogstashResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBody) SetRequestId(v string) *DescribeLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogstashResponseBody) SetResult(v *DescribeLogstashResponseBodyResult) *DescribeLogstashResponseBody {
	s.Result = v
	return s
}

type DescribeLogstashResponseBodyResult struct {
	ExtendConfigs   []map[string]interface{}                          `json:"ExtendConfigs,omitempty" xml:"ExtendConfigs,omitempty" type:"Repeated"`
	ResourceGroupId *string                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*DescribeLogstashResponseBodyResultTags         `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ZoneInfos       []*DescribeLogstashResponseBodyResultZoneInfos    `json:"ZoneInfos,omitempty" xml:"ZoneInfos,omitempty" type:"Repeated"`
	Config          map[string]interface{}                            `json:"config,omitempty" xml:"config,omitempty"`
	CreatedAt       *string                                           `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description     *string                                           `json:"description,omitempty" xml:"description,omitempty"`
	EndpointList    []*DescribeLogstashResponseBodyResultEndpointList `json:"endpointList,omitempty" xml:"endpointList,omitempty" type:"Repeated"`
	InstanceId      *string                                           `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	NetworkConfig   *DescribeLogstashResponseBodyResultNetworkConfig  `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	NodeAmount      *int32                                            `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec        *DescribeLogstashResponseBodyResultNodeSpec       `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	PaymentType     *string                                           `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	Status          *string                                           `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt       *string                                           `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Version         *string                                           `json:"version,omitempty" xml:"version,omitempty"`
	VpcInstanceId   *string                                           `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
}

func (s DescribeLogstashResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogstashResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResult) SetExtendConfigs(v []map[string]interface{}) *DescribeLogstashResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetResourceGroupId(v string) *DescribeLogstashResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetTags(v []*DescribeLogstashResponseBodyResultTags) *DescribeLogstashResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetZoneInfos(v []*DescribeLogstashResponseBodyResultZoneInfos) *DescribeLogstashResponseBodyResult {
	s.ZoneInfos = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetConfig(v map[string]interface{}) *DescribeLogstashResponseBodyResult {
	s.Config = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetCreatedAt(v string) *DescribeLogstashResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetDescription(v string) *DescribeLogstashResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetEndpointList(v []*DescribeLogstashResponseBodyResultEndpointList) *DescribeLogstashResponseBodyResult {
	s.EndpointList = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetInstanceId(v string) *DescribeLogstashResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetNetworkConfig(v *DescribeLogstashResponseBodyResultNetworkConfig) *DescribeLogstashResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetNodeAmount(v int32) *DescribeLogstashResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetNodeSpec(v *DescribeLogstashResponseBodyResultNodeSpec) *DescribeLogstashResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetPaymentType(v string) *DescribeLogstashResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetStatus(v string) *DescribeLogstashResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetUpdatedAt(v string) *DescribeLogstashResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetVersion(v string) *DescribeLogstashResponseBodyResult {
	s.Version = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetVpcInstanceId(v string) *DescribeLogstashResponseBodyResult {
	s.VpcInstanceId = &v
	return s
}

type DescribeLogstashResponseBodyResultTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s DescribeLogstashResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultTags) SetTagKey(v string) *DescribeLogstashResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultTags) SetTagValue(v string) *DescribeLogstashResponseBodyResultTags {
	s.TagValue = &v
	return s
}

type DescribeLogstashResponseBodyResultZoneInfos struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s DescribeLogstashResponseBodyResultZoneInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultZoneInfos) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultZoneInfos) SetStatus(v string) *DescribeLogstashResponseBodyResultZoneInfos {
	s.Status = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultZoneInfos) SetZoneId(v string) *DescribeLogstashResponseBodyResultZoneInfos {
	s.ZoneId = &v
	return s
}

type DescribeLogstashResponseBodyResultEndpointList struct {
	Host   *string `json:"host,omitempty" xml:"host,omitempty"`
	Port   *string `json:"port,omitempty" xml:"port,omitempty"`
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s DescribeLogstashResponseBodyResultEndpointList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultEndpointList) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultEndpointList) SetHost(v string) *DescribeLogstashResponseBodyResultEndpointList {
	s.Host = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultEndpointList) SetPort(v string) *DescribeLogstashResponseBodyResultEndpointList {
	s.Port = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultEndpointList) SetZoneId(v string) *DescribeLogstashResponseBodyResultEndpointList {
	s.ZoneId = &v
	return s
}

type DescribeLogstashResponseBodyResultNetworkConfig struct {
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	VpcId     *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea    *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s DescribeLogstashResponseBodyResultNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) SetType(v string) *DescribeLogstashResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) SetVpcId(v string) *DescribeLogstashResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) SetVsArea(v string) *DescribeLogstashResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) SetVswitchId(v string) *DescribeLogstashResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

type DescribeLogstashResponseBodyResultNodeSpec struct {
	Disk           *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskEncryption *bool   `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType       *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec           *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeLogstashResponseBodyResultNodeSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) SetDisk(v int32) *DescribeLogstashResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) SetDiskEncryption(v bool) *DescribeLogstashResponseBodyResultNodeSpec {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) SetDiskType(v string) *DescribeLogstashResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) SetSpec(v string) *DescribeLogstashResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

type DescribeLogstashResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogstashResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogstashResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponse) SetHeaders(v map[string]*string) *DescribeLogstashResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogstashResponse) SetBody(v *DescribeLogstashResponseBody) *DescribeLogstashResponse {
	s.Body = v
	return s
}

type DescribePipelineResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribePipelineResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBody) SetRequestId(v string) *DescribePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePipelineResponseBody) SetResult(v *DescribePipelineResponseBodyResult) *DescribePipelineResponseBody {
	s.Result = v
	return s
}

type DescribePipelineResponseBodyResult struct {
	BatchDelay            *int32  `json:"batchDelay,omitempty" xml:"batchDelay,omitempty"`
	BatchSize             *int32  `json:"batchSize,omitempty" xml:"batchSize,omitempty"`
	Config                *string `json:"config,omitempty" xml:"config,omitempty"`
	Description           *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreatedTime        *string `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	GmtUpdateTime         *string `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	PipelineId            *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PipelineStatus        *string `json:"pipelineStatus,omitempty" xml:"pipelineStatus,omitempty"`
	QueueCheckPointWrites *int32  `json:"queueCheckPointWrites,omitempty" xml:"queueCheckPointWrites,omitempty"`
	QueueMaxBytes         *int32  `json:"queueMaxBytes,omitempty" xml:"queueMaxBytes,omitempty"`
	QueueType             *string `json:"queueType,omitempty" xml:"queueType,omitempty"`
	Workers               *int32  `json:"workers,omitempty" xml:"workers,omitempty"`
}

func (s DescribePipelineResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBodyResult) SetBatchDelay(v int32) *DescribePipelineResponseBodyResult {
	s.BatchDelay = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetBatchSize(v int32) *DescribePipelineResponseBodyResult {
	s.BatchSize = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetConfig(v string) *DescribePipelineResponseBodyResult {
	s.Config = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetDescription(v string) *DescribePipelineResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetGmtCreatedTime(v string) *DescribePipelineResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetGmtUpdateTime(v string) *DescribePipelineResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetPipelineId(v string) *DescribePipelineResponseBodyResult {
	s.PipelineId = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetPipelineStatus(v string) *DescribePipelineResponseBodyResult {
	s.PipelineStatus = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetQueueCheckPointWrites(v int32) *DescribePipelineResponseBodyResult {
	s.QueueCheckPointWrites = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetQueueMaxBytes(v int32) *DescribePipelineResponseBodyResult {
	s.QueueMaxBytes = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetQueueType(v string) *DescribePipelineResponseBodyResult {
	s.QueueType = &v
	return s
}

func (s *DescribePipelineResponseBodyResult) SetWorkers(v int32) *DescribePipelineResponseBodyResult {
	s.Workers = &v
	return s
}

type DescribePipelineResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineResponse) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponse) SetHeaders(v map[string]*string) *DescribePipelineResponse {
	s.Headers = v
	return s
}

func (s *DescribePipelineResponse) SetBody(v *DescribePipelineResponseBody) *DescribePipelineResponse {
	s.Body = v
	return s
}

type DescribePipelineManagementConfigRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DescribePipelineManagementConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineManagementConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribePipelineManagementConfigRequest) SetClientToken(v string) *DescribePipelineManagementConfigRequest {
	s.ClientToken = &v
	return s
}

type DescribePipelineManagementConfigResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribePipelineManagementConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribePipelineManagementConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineManagementConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePipelineManagementConfigResponseBody) SetRequestId(v string) *DescribePipelineManagementConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePipelineManagementConfigResponseBody) SetResult(v *DescribePipelineManagementConfigResponseBodyResult) *DescribePipelineManagementConfigResponseBody {
	s.Result = v
	return s
}

type DescribePipelineManagementConfigResponseBodyResult struct {
	Endpoints              *string   `json:"endpoints,omitempty" xml:"endpoints,omitempty"`
	EsInstanceId           *string   `json:"esInstanceId,omitempty" xml:"esInstanceId,omitempty"`
	PipelineIds            []*string `json:"pipelineIds,omitempty" xml:"pipelineIds,omitempty" type:"Repeated"`
	PipelineManagementType *string   `json:"pipelineManagementType,omitempty" xml:"pipelineManagementType,omitempty"`
	UserName               *string   `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DescribePipelineManagementConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineManagementConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetEndpoints(v string) *DescribePipelineManagementConfigResponseBodyResult {
	s.Endpoints = &v
	return s
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetEsInstanceId(v string) *DescribePipelineManagementConfigResponseBodyResult {
	s.EsInstanceId = &v
	return s
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetPipelineIds(v []*string) *DescribePipelineManagementConfigResponseBodyResult {
	s.PipelineIds = v
	return s
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetPipelineManagementType(v string) *DescribePipelineManagementConfigResponseBodyResult {
	s.PipelineManagementType = &v
	return s
}

func (s *DescribePipelineManagementConfigResponseBodyResult) SetUserName(v string) *DescribePipelineManagementConfigResponseBodyResult {
	s.UserName = &v
	return s
}

type DescribePipelineManagementConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePipelineManagementConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePipelineManagementConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineManagementConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribePipelineManagementConfigResponse) SetHeaders(v map[string]*string) *DescribePipelineManagementConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribePipelineManagementConfigResponse) SetBody(v *DescribePipelineManagementConfigResponseBody) *DescribePipelineManagementConfigResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeRegionsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetResult(v []*DescribeRegionsResponseBodyResult) *DescribeRegionsResponseBody {
	s.Result = v
	return s
}

type DescribeRegionsResponseBodyResult struct {
	ConsoleEndpoint *string `json:"consoleEndpoint,omitempty" xml:"consoleEndpoint,omitempty"`
	LocalName       *string `json:"localName,omitempty" xml:"localName,omitempty"`
	RegionEndpoint  *string `json:"regionEndpoint,omitempty" xml:"regionEndpoint,omitempty"`
	RegionId        *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeRegionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyResult) SetConsoleEndpoint(v string) *DescribeRegionsResponseBodyResult {
	s.ConsoleEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetLocalName(v string) *DescribeRegionsResponseBodyResult {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyResult {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetRegionId(v string) *DescribeRegionsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetStatus(v string) *DescribeRegionsResponseBodyResult {
	s.Status = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSnapshotSettingResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeSnapshotSettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeSnapshotSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotSettingResponseBody) SetRequestId(v string) *DescribeSnapshotSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotSettingResponseBody) SetResult(v *DescribeSnapshotSettingResponseBodyResult) *DescribeSnapshotSettingResponseBody {
	s.Result = v
	return s
}

type DescribeSnapshotSettingResponseBodyResult struct {
	Enable      *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	QuartzRegex *string `json:"QuartzRegex,omitempty" xml:"QuartzRegex,omitempty"`
}

func (s DescribeSnapshotSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotSettingResponseBodyResult) SetEnable(v bool) *DescribeSnapshotSettingResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *DescribeSnapshotSettingResponseBodyResult) SetQuartzRegex(v string) *DescribeSnapshotSettingResponseBodyResult {
	s.QuartzRegex = &v
	return s
}

type DescribeSnapshotSettingResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSnapshotSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSnapshotSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotSettingResponse) SetHeaders(v map[string]*string) *DescribeSnapshotSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotSettingResponse) SetBody(v *DescribeSnapshotSettingResponseBody) *DescribeSnapshotSettingResponse {
	s.Body = v
	return s
}

type DescribeTemplatesResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeTemplatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBody) SetRequestId(v string) *DescribeTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplatesResponseBody) SetResult(v []*DescribeTemplatesResponseBodyResult) *DescribeTemplatesResponseBody {
	s.Result = v
	return s
}

type DescribeTemplatesResponseBodyResult struct {
	Content      *string `json:"content,omitempty" xml:"content,omitempty"`
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s DescribeTemplatesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponseBodyResult) SetContent(v string) *DescribeTemplatesResponseBodyResult {
	s.Content = &v
	return s
}

func (s *DescribeTemplatesResponseBodyResult) SetTemplateName(v string) *DescribeTemplatesResponseBodyResult {
	s.TemplateName = &v
	return s
}

type DescribeTemplatesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponse) SetHeaders(v map[string]*string) *DescribeTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplatesResponse) SetBody(v *DescribeTemplatesResponseBody) *DescribeTemplatesResponse {
	s.Body = v
	return s
}

type DescribeXpackMonitorConfigResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeXpackMonitorConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeXpackMonitorConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeXpackMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeXpackMonitorConfigResponseBody) SetRequestId(v string) *DescribeXpackMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBody) SetResult(v *DescribeXpackMonitorConfigResponseBodyResult) *DescribeXpackMonitorConfigResponseBody {
	s.Result = v
	return s
}

type DescribeXpackMonitorConfigResponseBodyResult struct {
	Enable       *bool     `json:"enable,omitempty" xml:"enable,omitempty"`
	Endpoints    []*string `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	EsInstanceId *string   `json:"esInstanceId,omitempty" xml:"esInstanceId,omitempty"`
	PipelineIds  []*string `json:"pipelineIds,omitempty" xml:"pipelineIds,omitempty" type:"Repeated"`
	UserName     *string   `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DescribeXpackMonitorConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeXpackMonitorConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetEnable(v bool) *DescribeXpackMonitorConfigResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetEndpoints(v []*string) *DescribeXpackMonitorConfigResponseBodyResult {
	s.Endpoints = v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetEsInstanceId(v string) *DescribeXpackMonitorConfigResponseBodyResult {
	s.EsInstanceId = &v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetPipelineIds(v []*string) *DescribeXpackMonitorConfigResponseBodyResult {
	s.PipelineIds = v
	return s
}

func (s *DescribeXpackMonitorConfigResponseBodyResult) SetUserName(v string) *DescribeXpackMonitorConfigResponseBodyResult {
	s.UserName = &v
	return s
}

type DescribeXpackMonitorConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeXpackMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeXpackMonitorConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeXpackMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeXpackMonitorConfigResponse) SetHeaders(v map[string]*string) *DescribeXpackMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeXpackMonitorConfigResponse) SetBody(v *DescribeXpackMonitorConfigResponseBody) *DescribeXpackMonitorConfigResponse {
	s.Body = v
	return s
}

type DiagnoseInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang        *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s DiagnoseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DiagnoseInstanceRequest) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceRequest) SetClientToken(v string) *DiagnoseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DiagnoseInstanceRequest) SetLang(v string) *DiagnoseInstanceRequest {
	s.Lang = &v
	return s
}

type DiagnoseInstanceResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DiagnoseInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DiagnoseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DiagnoseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponseBody) SetRequestId(v string) *DiagnoseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiagnoseInstanceResponseBody) SetResult(v *DiagnoseInstanceResponseBodyResult) *DiagnoseInstanceResponseBody {
	s.Result = v
	return s
}

type DiagnoseInstanceResponseBodyResult struct {
	CreateTime    *int64                                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiagnoseItems []*DiagnoseInstanceResponseBodyResultDiagnoseItems `json:"diagnoseItems,omitempty" xml:"diagnoseItems,omitempty" type:"Repeated"`
	InstanceId    *string                                            `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ReportId      *string                                            `json:"reportId,omitempty" xml:"reportId,omitempty"`
	State         *string                                            `json:"state,omitempty" xml:"state,omitempty"`
}

func (s DiagnoseInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DiagnoseInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponseBodyResult) SetCreateTime(v int64) *DiagnoseInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetDiagnoseItems(v []*DiagnoseInstanceResponseBodyResultDiagnoseItems) *DiagnoseInstanceResponseBodyResult {
	s.DiagnoseItems = v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetInstanceId(v string) *DiagnoseInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetReportId(v string) *DiagnoseInstanceResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *DiagnoseInstanceResponseBodyResult) SetState(v string) *DiagnoseInstanceResponseBodyResult {
	s.State = &v
	return s
}

type DiagnoseInstanceResponseBodyResultDiagnoseItems struct {
	Item *string `json:"item,omitempty" xml:"item,omitempty"`
}

func (s DiagnoseInstanceResponseBodyResultDiagnoseItems) String() string {
	return tea.Prettify(s)
}

func (s DiagnoseInstanceResponseBodyResultDiagnoseItems) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponseBodyResultDiagnoseItems) SetItem(v string) *DiagnoseInstanceResponseBodyResultDiagnoseItems {
	s.Item = &v
	return s
}

type DiagnoseInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DiagnoseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DiagnoseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DiagnoseInstanceResponse) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponse) SetHeaders(v map[string]*string) *DiagnoseInstanceResponse {
	s.Headers = v
	return s
}

func (s *DiagnoseInstanceResponse) SetBody(v *DiagnoseInstanceResponseBody) *DiagnoseInstanceResponse {
	s.Body = v
	return s
}

type EstimatedLogstashRestartTimeRequest struct {
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s EstimatedLogstashRestartTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s EstimatedLogstashRestartTimeRequest) GoString() string {
	return s.String()
}

func (s *EstimatedLogstashRestartTimeRequest) SetForce(v bool) *EstimatedLogstashRestartTimeRequest {
	s.Force = &v
	return s
}

type EstimatedLogstashRestartTimeResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EstimatedLogstashRestartTimeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EstimatedLogstashRestartTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EstimatedLogstashRestartTimeResponseBody) GoString() string {
	return s.String()
}

func (s *EstimatedLogstashRestartTimeResponseBody) SetRequestId(v string) *EstimatedLogstashRestartTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *EstimatedLogstashRestartTimeResponseBody) SetResult(v *EstimatedLogstashRestartTimeResponseBodyResult) *EstimatedLogstashRestartTimeResponseBody {
	s.Result = v
	return s
}

type EstimatedLogstashRestartTimeResponseBodyResult struct {
	Unit  *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Value *int64  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s EstimatedLogstashRestartTimeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EstimatedLogstashRestartTimeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EstimatedLogstashRestartTimeResponseBodyResult) SetUnit(v string) *EstimatedLogstashRestartTimeResponseBodyResult {
	s.Unit = &v
	return s
}

func (s *EstimatedLogstashRestartTimeResponseBodyResult) SetValue(v int64) *EstimatedLogstashRestartTimeResponseBodyResult {
	s.Value = &v
	return s
}

type EstimatedLogstashRestartTimeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EstimatedLogstashRestartTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EstimatedLogstashRestartTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s EstimatedLogstashRestartTimeResponse) GoString() string {
	return s.String()
}

func (s *EstimatedLogstashRestartTimeResponse) SetHeaders(v map[string]*string) *EstimatedLogstashRestartTimeResponse {
	s.Headers = v
	return s
}

func (s *EstimatedLogstashRestartTimeResponse) SetBody(v *EstimatedLogstashRestartTimeResponseBody) *EstimatedLogstashRestartTimeResponse {
	s.Body = v
	return s
}

type EstimatedRestartTimeRequest struct {
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s EstimatedRestartTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s EstimatedRestartTimeRequest) GoString() string {
	return s.String()
}

func (s *EstimatedRestartTimeRequest) SetForce(v bool) *EstimatedRestartTimeRequest {
	s.Force = &v
	return s
}

type EstimatedRestartTimeResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EstimatedRestartTimeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EstimatedRestartTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EstimatedRestartTimeResponseBody) GoString() string {
	return s.String()
}

func (s *EstimatedRestartTimeResponseBody) SetRequestId(v string) *EstimatedRestartTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *EstimatedRestartTimeResponseBody) SetResult(v *EstimatedRestartTimeResponseBodyResult) *EstimatedRestartTimeResponseBody {
	s.Result = v
	return s
}

type EstimatedRestartTimeResponseBodyResult struct {
	Unit  *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Value *int64  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s EstimatedRestartTimeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EstimatedRestartTimeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EstimatedRestartTimeResponseBodyResult) SetUnit(v string) *EstimatedRestartTimeResponseBodyResult {
	s.Unit = &v
	return s
}

func (s *EstimatedRestartTimeResponseBodyResult) SetValue(v int64) *EstimatedRestartTimeResponseBodyResult {
	s.Value = &v
	return s
}

type EstimatedRestartTimeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EstimatedRestartTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EstimatedRestartTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s EstimatedRestartTimeResponse) GoString() string {
	return s.String()
}

func (s *EstimatedRestartTimeResponse) SetHeaders(v map[string]*string) *EstimatedRestartTimeResponse {
	s.Headers = v
	return s
}

func (s *EstimatedRestartTimeResponse) SetBody(v *EstimatedRestartTimeResponseBody) *EstimatedRestartTimeResponse {
	s.Body = v
	return s
}

type GetClusterDataInformationResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetClusterDataInformationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetClusterDataInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterDataInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterDataInformationResponseBody) SetRequestId(v string) *GetClusterDataInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterDataInformationResponseBody) SetResult(v *GetClusterDataInformationResponseBodyResult) *GetClusterDataInformationResponseBody {
	s.Result = v
	return s
}

type GetClusterDataInformationResponseBodyResult struct {
	Connectable *bool                                                `json:"connectable,omitempty" xml:"connectable,omitempty"`
	MetaInfo    *GetClusterDataInformationResponseBodyResultMetaInfo `json:"metaInfo,omitempty" xml:"metaInfo,omitempty" type:"Struct"`
}

func (s GetClusterDataInformationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClusterDataInformationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClusterDataInformationResponseBodyResult) SetConnectable(v bool) *GetClusterDataInformationResponseBodyResult {
	s.Connectable = &v
	return s
}

func (s *GetClusterDataInformationResponseBodyResult) SetMetaInfo(v *GetClusterDataInformationResponseBodyResultMetaInfo) *GetClusterDataInformationResponseBodyResult {
	s.MetaInfo = v
	return s
}

type GetClusterDataInformationResponseBodyResultMetaInfo struct {
	Fields   []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Indices  []*string `json:"indices,omitempty" xml:"indices,omitempty" type:"Repeated"`
	Mapping  *string   `json:"mapping,omitempty" xml:"mapping,omitempty"`
	Settings *string   `json:"settings,omitempty" xml:"settings,omitempty"`
	TypeName []*string `json:"typeName,omitempty" xml:"typeName,omitempty" type:"Repeated"`
}

func (s GetClusterDataInformationResponseBodyResultMetaInfo) String() string {
	return tea.Prettify(s)
}

func (s GetClusterDataInformationResponseBodyResultMetaInfo) GoString() string {
	return s.String()
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetFields(v []*string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.Fields = v
	return s
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetIndices(v []*string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.Indices = v
	return s
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetMapping(v string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.Mapping = &v
	return s
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetSettings(v string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.Settings = &v
	return s
}

func (s *GetClusterDataInformationResponseBodyResultMetaInfo) SetTypeName(v []*string) *GetClusterDataInformationResponseBodyResultMetaInfo {
	s.TypeName = v
	return s
}

type GetClusterDataInformationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClusterDataInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterDataInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterDataInformationResponse) GoString() string {
	return s.String()
}

func (s *GetClusterDataInformationResponse) SetHeaders(v map[string]*string) *GetClusterDataInformationResponse {
	s.Headers = v
	return s
}

func (s *GetClusterDataInformationResponse) SetBody(v *GetClusterDataInformationResponseBody) *GetClusterDataInformationResponse {
	s.Body = v
	return s
}

type GetElastictaskResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetElastictaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetElastictaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetElastictaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponseBody) SetRequestId(v string) *GetElastictaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetElastictaskResponseBody) SetResult(v *GetElastictaskResponseBodyResult) *GetElastictaskResponseBody {
	s.Result = v
	return s
}

type GetElastictaskResponseBodyResult struct {
	ElasticExpansionTask *GetElastictaskResponseBodyResultElasticExpansionTask `json:"elasticExpansionTask,omitempty" xml:"elasticExpansionTask,omitempty" type:"Struct"`
	ElasticShrinkTask    *GetElastictaskResponseBodyResultElasticShrinkTask    `json:"elasticShrinkTask,omitempty" xml:"elasticShrinkTask,omitempty" type:"Struct"`
}

func (s GetElastictaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetElastictaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponseBodyResult) SetElasticExpansionTask(v *GetElastictaskResponseBodyResultElasticExpansionTask) *GetElastictaskResponseBodyResult {
	s.ElasticExpansionTask = v
	return s
}

func (s *GetElastictaskResponseBodyResult) SetElasticShrinkTask(v *GetElastictaskResponseBodyResultElasticShrinkTask) *GetElastictaskResponseBodyResult {
	s.ElasticShrinkTask = v
	return s
}

type GetElastictaskResponseBodyResultElasticExpansionTask struct {
	CronExpression   *string   `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	ElasticNodeCount *int32    `json:"elasticNodeCount,omitempty" xml:"elasticNodeCount,omitempty"`
	ReplicaCount     *int32    `json:"replicaCount,omitempty" xml:"replicaCount,omitempty"`
	TargetIndices    []*string `json:"targetIndices,omitempty" xml:"targetIndices,omitempty" type:"Repeated"`
	TriggerType      *string   `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetElastictaskResponseBodyResultElasticExpansionTask) String() string {
	return tea.Prettify(s)
}

func (s GetElastictaskResponseBodyResultElasticExpansionTask) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetCronExpression(v string) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.CronExpression = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetElasticNodeCount(v int32) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.ElasticNodeCount = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetReplicaCount(v int32) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.ReplicaCount = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetTargetIndices(v []*string) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.TargetIndices = v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticExpansionTask) SetTriggerType(v string) *GetElastictaskResponseBodyResultElasticExpansionTask {
	s.TriggerType = &v
	return s
}

type GetElastictaskResponseBodyResultElasticShrinkTask struct {
	CronExpression   *string   `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	ElasticNodeCount *int32    `json:"elasticNodeCount,omitempty" xml:"elasticNodeCount,omitempty"`
	ReplicaCount     *int32    `json:"replicaCount,omitempty" xml:"replicaCount,omitempty"`
	TargetIndices    []*string `json:"targetIndices,omitempty" xml:"targetIndices,omitempty" type:"Repeated"`
	TriggerType      *string   `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetElastictaskResponseBodyResultElasticShrinkTask) String() string {
	return tea.Prettify(s)
}

func (s GetElastictaskResponseBodyResultElasticShrinkTask) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetCronExpression(v string) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.CronExpression = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetElasticNodeCount(v int32) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.ElasticNodeCount = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetReplicaCount(v int32) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.ReplicaCount = &v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetTargetIndices(v []*string) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.TargetIndices = v
	return s
}

func (s *GetElastictaskResponseBodyResultElasticShrinkTask) SetTriggerType(v string) *GetElastictaskResponseBodyResultElasticShrinkTask {
	s.TriggerType = &v
	return s
}

type GetElastictaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetElastictaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetElastictaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetElastictaskResponse) GoString() string {
	return s.String()
}

func (s *GetElastictaskResponse) SetHeaders(v map[string]*string) *GetElastictaskResponse {
	s.Headers = v
	return s
}

func (s *GetElastictaskResponse) SetBody(v *GetElastictaskResponseBody) *GetElastictaskResponse {
	s.Body = v
	return s
}

type GetEmonGrafanaAlertsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEmonGrafanaAlertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmonGrafanaAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaAlertsResponseBody) SetCode(v string) *GetEmonGrafanaAlertsResponseBody {
	s.Code = &v
	return s
}

func (s *GetEmonGrafanaAlertsResponseBody) SetMessage(v string) *GetEmonGrafanaAlertsResponseBody {
	s.Message = &v
	return s
}

func (s *GetEmonGrafanaAlertsResponseBody) SetRequestId(v string) *GetEmonGrafanaAlertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmonGrafanaAlertsResponseBody) SetSuccess(v bool) *GetEmonGrafanaAlertsResponseBody {
	s.Success = &v
	return s
}

type GetEmonGrafanaAlertsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEmonGrafanaAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEmonGrafanaAlertsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmonGrafanaAlertsResponse) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaAlertsResponse) SetHeaders(v map[string]*string) *GetEmonGrafanaAlertsResponse {
	s.Headers = v
	return s
}

func (s *GetEmonGrafanaAlertsResponse) SetBody(v *GetEmonGrafanaAlertsResponseBody) *GetEmonGrafanaAlertsResponse {
	s.Body = v
	return s
}

type GetEmonGrafanaDashboardsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEmonGrafanaDashboardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmonGrafanaDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaDashboardsResponseBody) SetCode(v string) *GetEmonGrafanaDashboardsResponseBody {
	s.Code = &v
	return s
}

func (s *GetEmonGrafanaDashboardsResponseBody) SetMessage(v string) *GetEmonGrafanaDashboardsResponseBody {
	s.Message = &v
	return s
}

func (s *GetEmonGrafanaDashboardsResponseBody) SetRequestId(v string) *GetEmonGrafanaDashboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmonGrafanaDashboardsResponseBody) SetSuccess(v bool) *GetEmonGrafanaDashboardsResponseBody {
	s.Success = &v
	return s
}

type GetEmonGrafanaDashboardsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEmonGrafanaDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEmonGrafanaDashboardsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmonGrafanaDashboardsResponse) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaDashboardsResponse) SetHeaders(v map[string]*string) *GetEmonGrafanaDashboardsResponse {
	s.Headers = v
	return s
}

func (s *GetEmonGrafanaDashboardsResponse) SetBody(v *GetEmonGrafanaDashboardsResponseBody) *GetEmonGrafanaDashboardsResponse {
	s.Body = v
	return s
}

type GetEmonMonitorDataResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetEmonMonitorDataResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEmonMonitorDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmonMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmonMonitorDataResponseBody) SetCode(v string) *GetEmonMonitorDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetEmonMonitorDataResponseBody) SetMessage(v string) *GetEmonMonitorDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetEmonMonitorDataResponseBody) SetRequestId(v string) *GetEmonMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmonMonitorDataResponseBody) SetResult(v []*GetEmonMonitorDataResponseBodyResult) *GetEmonMonitorDataResponseBody {
	s.Result = v
	return s
}

func (s *GetEmonMonitorDataResponseBody) SetSuccess(v bool) *GetEmonMonitorDataResponseBody {
	s.Success = &v
	return s
}

type GetEmonMonitorDataResponseBodyResult struct {
	Dps              map[string]interface{} `json:"dps,omitempty" xml:"dps,omitempty"`
	Integrity        *float32               `json:"integrity,omitempty" xml:"integrity,omitempty"`
	MessageWatermark *int64                 `json:"messageWatermark,omitempty" xml:"messageWatermark,omitempty"`
	Metric           *string                `json:"metric,omitempty" xml:"metric,omitempty"`
	Summary          *float32               `json:"summary,omitempty" xml:"summary,omitempty"`
	Tags             map[string]interface{} `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s GetEmonMonitorDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetEmonMonitorDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetEmonMonitorDataResponseBodyResult) SetDps(v map[string]interface{}) *GetEmonMonitorDataResponseBodyResult {
	s.Dps = v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetIntegrity(v float32) *GetEmonMonitorDataResponseBodyResult {
	s.Integrity = &v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetMessageWatermark(v int64) *GetEmonMonitorDataResponseBodyResult {
	s.MessageWatermark = &v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetMetric(v string) *GetEmonMonitorDataResponseBodyResult {
	s.Metric = &v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetSummary(v float32) *GetEmonMonitorDataResponseBodyResult {
	s.Summary = &v
	return s
}

func (s *GetEmonMonitorDataResponseBodyResult) SetTags(v map[string]interface{}) *GetEmonMonitorDataResponseBodyResult {
	s.Tags = v
	return s
}

type GetEmonMonitorDataResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEmonMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEmonMonitorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmonMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *GetEmonMonitorDataResponse) SetHeaders(v map[string]*string) *GetEmonMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *GetEmonMonitorDataResponse) SetBody(v *GetEmonMonitorDataResponseBody) *GetEmonMonitorDataResponse {
	s.Body = v
	return s
}

type GetOpenStoreUsageResponseBody struct {
	// Id of the request
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetOpenStoreUsageResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetOpenStoreUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenStoreUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenStoreUsageResponseBody) SetRequestId(v string) *GetOpenStoreUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenStoreUsageResponseBody) SetResult(v *GetOpenStoreUsageResponseBodyResult) *GetOpenStoreUsageResponseBody {
	s.Result = v
	return s
}

type GetOpenStoreUsageResponseBodyResult struct {
	// 当前使用量
	CurrentUsage *int64 `json:"currentUsage,omitempty" xml:"currentUsage,omitempty"`
	// 昨日使用容量
	LastDayUsage *int64 `json:"lastDayUsage,omitempty" xml:"lastDayUsage,omitempty"`
}

func (s GetOpenStoreUsageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOpenStoreUsageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOpenStoreUsageResponseBodyResult) SetCurrentUsage(v int64) *GetOpenStoreUsageResponseBodyResult {
	s.CurrentUsage = &v
	return s
}

func (s *GetOpenStoreUsageResponseBodyResult) SetLastDayUsage(v int64) *GetOpenStoreUsageResponseBodyResult {
	s.LastDayUsage = &v
	return s
}

type GetOpenStoreUsageResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOpenStoreUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpenStoreUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenStoreUsageResponse) GoString() string {
	return s.String()
}

func (s *GetOpenStoreUsageResponse) SetHeaders(v map[string]*string) *GetOpenStoreUsageResponse {
	s.Headers = v
	return s
}

func (s *GetOpenStoreUsageResponse) SetBody(v *GetOpenStoreUsageResponseBody) *GetOpenStoreUsageResponse {
	s.Body = v
	return s
}

type GetRegionConfigurationRequest struct {
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s GetRegionConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationRequest) SetZoneId(v string) *GetRegionConfigurationRequest {
	s.ZoneId = &v
	return s
}

type GetRegionConfigurationResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetRegionConfigurationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRegionConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBody) SetRequestId(v string) *GetRegionConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegionConfigurationResponseBody) SetResult(v *GetRegionConfigurationResponseBodyResult) *GetRegionConfigurationResponseBody {
	s.Result = v
	return s
}

type GetRegionConfigurationResponseBodyResult struct {
	ClientNodeAmountRange *GetRegionConfigurationResponseBodyResultClientNodeAmountRange  `json:"clientNodeAmountRange,omitempty" xml:"clientNodeAmountRange,omitempty" type:"Struct"`
	ClientNodeDiskList    []*GetRegionConfigurationResponseBodyResultClientNodeDiskList   `json:"clientNodeDiskList,omitempty" xml:"clientNodeDiskList,omitempty" type:"Repeated"`
	ClientNodeSpec        []*string                                                       `json:"clientNodeSpec,omitempty" xml:"clientNodeSpec,omitempty" type:"Repeated"`
	CreateUrl             *string                                                         `json:"createUrl,omitempty" xml:"createUrl,omitempty"`
	DataDiskList          []*GetRegionConfigurationResponseBodyResultDataDiskList         `json:"dataDiskList,omitempty" xml:"dataDiskList,omitempty" type:"Repeated"`
	ElasticNodeProperties *GetRegionConfigurationResponseBodyResultElasticNodeProperties  `json:"elasticNodeProperties,omitempty" xml:"elasticNodeProperties,omitempty" type:"Struct"`
	Env                   *string                                                         `json:"env,omitempty" xml:"env,omitempty"`
	EsVersions            []*string                                                       `json:"esVersions,omitempty" xml:"esVersions,omitempty" type:"Repeated"`
	EsVersionsLatestList  []*GetRegionConfigurationResponseBodyResultEsVersionsLatestList `json:"esVersionsLatestList,omitempty" xml:"esVersionsLatestList,omitempty" type:"Repeated"`
	InstanceSupportNodes  []*string                                                       `json:"instanceSupportNodes,omitempty" xml:"instanceSupportNodes,omitempty" type:"Repeated"`
	JvmConfine            *GetRegionConfigurationResponseBodyResultJvmConfine             `json:"jvmConfine,omitempty" xml:"jvmConfine,omitempty" type:"Struct"`
	KibanaNodeProperties  *GetRegionConfigurationResponseBodyResultKibanaNodeProperties   `json:"kibanaNodeProperties,omitempty" xml:"kibanaNodeProperties,omitempty" type:"Struct"`
	MasterDiskList        []*GetRegionConfigurationResponseBodyResultMasterDiskList       `json:"masterDiskList,omitempty" xml:"masterDiskList,omitempty" type:"Repeated"`
	MasterSpec            []*string                                                       `json:"masterSpec,omitempty" xml:"masterSpec,omitempty" type:"Repeated"`
	Node                  *GetRegionConfigurationResponseBodyResultNode                   `json:"node,omitempty" xml:"node,omitempty" type:"Struct"`
	NodeSpecList          []*GetRegionConfigurationResponseBodyResultNodeSpecList         `json:"nodeSpecList,omitempty" xml:"nodeSpecList,omitempty" type:"Repeated"`
	RegionId              *string                                                         `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SupportVersions       []*GetRegionConfigurationResponseBodyResultSupportVersions      `json:"supportVersions,omitempty" xml:"supportVersions,omitempty" type:"Repeated"`
	WarmNodeProperties    *GetRegionConfigurationResponseBodyResultWarmNodeProperties     `json:"warmNodeProperties,omitempty" xml:"warmNodeProperties,omitempty" type:"Struct"`
	Zones                 []*string                                                       `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResult) SetClientNodeAmountRange(v *GetRegionConfigurationResponseBodyResultClientNodeAmountRange) *GetRegionConfigurationResponseBodyResult {
	s.ClientNodeAmountRange = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetClientNodeDiskList(v []*GetRegionConfigurationResponseBodyResultClientNodeDiskList) *GetRegionConfigurationResponseBodyResult {
	s.ClientNodeDiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetClientNodeSpec(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.ClientNodeSpec = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetCreateUrl(v string) *GetRegionConfigurationResponseBodyResult {
	s.CreateUrl = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetDataDiskList(v []*GetRegionConfigurationResponseBodyResultDataDiskList) *GetRegionConfigurationResponseBodyResult {
	s.DataDiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetElasticNodeProperties(v *GetRegionConfigurationResponseBodyResultElasticNodeProperties) *GetRegionConfigurationResponseBodyResult {
	s.ElasticNodeProperties = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetEnv(v string) *GetRegionConfigurationResponseBodyResult {
	s.Env = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetEsVersions(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.EsVersions = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetEsVersionsLatestList(v []*GetRegionConfigurationResponseBodyResultEsVersionsLatestList) *GetRegionConfigurationResponseBodyResult {
	s.EsVersionsLatestList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetInstanceSupportNodes(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.InstanceSupportNodes = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetJvmConfine(v *GetRegionConfigurationResponseBodyResultJvmConfine) *GetRegionConfigurationResponseBodyResult {
	s.JvmConfine = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetKibanaNodeProperties(v *GetRegionConfigurationResponseBodyResultKibanaNodeProperties) *GetRegionConfigurationResponseBodyResult {
	s.KibanaNodeProperties = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetMasterDiskList(v []*GetRegionConfigurationResponseBodyResultMasterDiskList) *GetRegionConfigurationResponseBodyResult {
	s.MasterDiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetMasterSpec(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.MasterSpec = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetNode(v *GetRegionConfigurationResponseBodyResultNode) *GetRegionConfigurationResponseBodyResult {
	s.Node = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetNodeSpecList(v []*GetRegionConfigurationResponseBodyResultNodeSpecList) *GetRegionConfigurationResponseBodyResult {
	s.NodeSpecList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetRegionId(v string) *GetRegionConfigurationResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetSupportVersions(v []*GetRegionConfigurationResponseBodyResultSupportVersions) *GetRegionConfigurationResponseBodyResult {
	s.SupportVersions = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetWarmNodeProperties(v *GetRegionConfigurationResponseBodyResultWarmNodeProperties) *GetRegionConfigurationResponseBodyResult {
	s.WarmNodeProperties = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetZones(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.Zones = v
	return s
}

type GetRegionConfigurationResponseBodyResultClientNodeAmountRange struct {
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultClientNodeAmountRange) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultClientNodeAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeAmountRange) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultClientNodeAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeAmountRange) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultClientNodeAmountRange {
	s.MinAmount = &v
	return s
}

type GetRegionConfigurationResponseBodyResultClientNodeDiskList struct {
	DiskType   *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	MaxSize    *int32  `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize    *int32  `json:"minSize,omitempty" xml:"minSize,omitempty"`
	ScaleLimit *int32  `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultClientNodeDiskList) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultClientNodeDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultClientNodeDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultClientNodeDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultClientNodeDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultClientNodeDiskList {
	s.ScaleLimit = &v
	return s
}

type GetRegionConfigurationResponseBodyResultDataDiskList struct {
	DiskType      *string   `json:"diskType,omitempty" xml:"diskType,omitempty"`
	MaxSize       *int32    `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize       *int32    `json:"minSize,omitempty" xml:"minSize,omitempty"`
	ScaleLimit    *int32    `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
	ValueLimitSet []*string `json:"valueLimitSet,omitempty" xml:"valueLimitSet,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultDataDiskList) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultDataDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetValueLimitSet(v []*string) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.ValueLimitSet = v
	return s
}

type GetRegionConfigurationResponseBodyResultElasticNodeProperties struct {
	AmountRange *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange `json:"amountRange,omitempty" xml:"amountRange,omitempty" type:"Struct"`
	DiskList    []*GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList  `json:"diskList,omitempty" xml:"diskList,omitempty" type:"Repeated"`
	Spec        []*string                                                                 `json:"spec,omitempty" xml:"spec,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultElasticNodeProperties) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultElasticNodeProperties) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) SetAmountRange(v *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) *GetRegionConfigurationResponseBodyResultElasticNodeProperties {
	s.AmountRange = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) SetDiskList(v []*GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) *GetRegionConfigurationResponseBodyResultElasticNodeProperties {
	s.DiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) SetSpec(v []*string) *GetRegionConfigurationResponseBodyResultElasticNodeProperties {
	s.Spec = v
	return s
}

type GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange struct {
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange {
	s.MinAmount = &v
	return s
}

type GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList struct {
	DiskEncryption *bool     `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType       *string   `json:"diskType,omitempty" xml:"diskType,omitempty"`
	MaxSize        *int32    `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize        *int32    `json:"minSize,omitempty" xml:"minSize,omitempty"`
	ScaleLimit     *int32    `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
	ValueLimitSet  []*string `json:"valueLimitSet,omitempty" xml:"valueLimitSet,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetDiskEncryption(v bool) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.DiskEncryption = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetValueLimitSet(v []*string) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.ValueLimitSet = v
	return s
}

type GetRegionConfigurationResponseBodyResultEsVersionsLatestList struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultEsVersionsLatestList) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultEsVersionsLatestList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultEsVersionsLatestList) SetKey(v string) *GetRegionConfigurationResponseBodyResultEsVersionsLatestList {
	s.Key = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultEsVersionsLatestList) SetValue(v string) *GetRegionConfigurationResponseBodyResultEsVersionsLatestList {
	s.Value = &v
	return s
}

type GetRegionConfigurationResponseBodyResultJvmConfine struct {
	Memory            *int32    `json:"memory,omitempty" xml:"memory,omitempty"`
	SupportEsVersions []*string `json:"supportEsVersions,omitempty" xml:"supportEsVersions,omitempty" type:"Repeated"`
	SupportGcs        []*string `json:"supportGcs,omitempty" xml:"supportGcs,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultJvmConfine) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultJvmConfine) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) SetMemory(v int32) *GetRegionConfigurationResponseBodyResultJvmConfine {
	s.Memory = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) SetSupportEsVersions(v []*string) *GetRegionConfigurationResponseBodyResultJvmConfine {
	s.SupportEsVersions = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) SetSupportGcs(v []*string) *GetRegionConfigurationResponseBodyResultJvmConfine {
	s.SupportGcs = v
	return s
}

type GetRegionConfigurationResponseBodyResultKibanaNodeProperties struct {
	AmountRange *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange `json:"amountRange,omitempty" xml:"amountRange,omitempty" type:"Struct"`
	Spec        []*string                                                                `json:"spec,omitempty" xml:"spec,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultKibanaNodeProperties) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultKibanaNodeProperties) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodeProperties) SetAmountRange(v *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) *GetRegionConfigurationResponseBodyResultKibanaNodeProperties {
	s.AmountRange = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodeProperties) SetSpec(v []*string) *GetRegionConfigurationResponseBodyResultKibanaNodeProperties {
	s.Spec = v
	return s
}

type GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange struct {
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange {
	s.MinAmount = &v
	return s
}

type GetRegionConfigurationResponseBodyResultMasterDiskList struct {
	DiskType   *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	MaxSize    *int32  `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize    *int32  `json:"minSize,omitempty" xml:"minSize,omitempty"`
	ScaleLimit *int32  `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultMasterDiskList) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultMasterDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultMasterDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultMasterDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultMasterDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultMasterDiskList {
	s.ScaleLimit = &v
	return s
}

type GetRegionConfigurationResponseBodyResultNode struct {
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultNode) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultNode) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultNode) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultNode {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNode) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultNode {
	s.MinAmount = &v
	return s
}

type GetRegionConfigurationResponseBodyResultNodeSpecList struct {
	CpuCount      *int32  `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
	Disk          *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType      *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Enable        *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	MemorySize    *int32  `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	Spec          *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecGroupType *string `json:"specGroupType,omitempty" xml:"specGroupType,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultNodeSpecList) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultNodeSpecList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetCpuCount(v int32) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.CpuCount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetDisk(v int32) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.Disk = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetEnable(v bool) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.Enable = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetMemorySize(v int32) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.MemorySize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetSpec(v string) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.Spec = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetSpecGroupType(v string) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.SpecGroupType = &v
	return s
}

type GetRegionConfigurationResponseBodyResultSupportVersions struct {
	InstanceCategory   *string                                                                      `json:"instanceCategory,omitempty" xml:"instanceCategory,omitempty"`
	SupportVersionList []*GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList `json:"supportVersionList,omitempty" xml:"supportVersionList,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultSupportVersions) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultSupportVersions) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersions) SetInstanceCategory(v string) *GetRegionConfigurationResponseBodyResultSupportVersions {
	s.InstanceCategory = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersions) SetSupportVersionList(v []*GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) *GetRegionConfigurationResponseBodyResultSupportVersions {
	s.SupportVersionList = v
	return s
}

type GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) SetKey(v string) *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList {
	s.Key = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) SetValue(v string) *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList {
	s.Value = &v
	return s
}

type GetRegionConfigurationResponseBodyResultWarmNodeProperties struct {
	AmountRange *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange `json:"amountRange,omitempty" xml:"amountRange,omitempty" type:"Struct"`
	DiskList    []*GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList  `json:"diskList,omitempty" xml:"diskList,omitempty" type:"Repeated"`
	Spec        []*string                                                              `json:"spec,omitempty" xml:"spec,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultWarmNodeProperties) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultWarmNodeProperties) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) SetAmountRange(v *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) *GetRegionConfigurationResponseBodyResultWarmNodeProperties {
	s.AmountRange = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) SetDiskList(v []*GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) *GetRegionConfigurationResponseBodyResultWarmNodeProperties {
	s.DiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) SetSpec(v []*string) *GetRegionConfigurationResponseBodyResultWarmNodeProperties {
	s.Spec = v
	return s
}

type GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange struct {
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange {
	s.MinAmount = &v
	return s
}

type GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList struct {
	DiskEncryption *bool     `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType       *string   `json:"diskType,omitempty" xml:"diskType,omitempty"`
	MaxSize        *int32    `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	MinSize        *int32    `json:"minSize,omitempty" xml:"minSize,omitempty"`
	ScaleLimit     *int32    `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
	ValueLimitSet  []*string `json:"valueLimitSet,omitempty" xml:"valueLimitSet,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetDiskEncryption(v bool) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.DiskEncryption = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetValueLimitSet(v []*string) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.ValueLimitSet = v
	return s
}

type GetRegionConfigurationResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRegionConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRegionConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegionConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponse) SetHeaders(v map[string]*string) *GetRegionConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetRegionConfigurationResponse) SetBody(v *GetRegionConfigurationResponseBody) *GetRegionConfigurationResponse {
	s.Body = v
	return s
}

type GetSuggestShrinkableNodesRequest struct {
	Count        *int32  `json:"count,omitempty" xml:"count,omitempty"`
	IgnoreStatus *bool   `json:"ignoreStatus,omitempty" xml:"ignoreStatus,omitempty"`
	NodeType     *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s GetSuggestShrinkableNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSuggestShrinkableNodesRequest) GoString() string {
	return s.String()
}

func (s *GetSuggestShrinkableNodesRequest) SetCount(v int32) *GetSuggestShrinkableNodesRequest {
	s.Count = &v
	return s
}

func (s *GetSuggestShrinkableNodesRequest) SetIgnoreStatus(v bool) *GetSuggestShrinkableNodesRequest {
	s.IgnoreStatus = &v
	return s
}

func (s *GetSuggestShrinkableNodesRequest) SetNodeType(v string) *GetSuggestShrinkableNodesRequest {
	s.NodeType = &v
	return s
}

type GetSuggestShrinkableNodesResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetSuggestShrinkableNodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetSuggestShrinkableNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuggestShrinkableNodesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuggestShrinkableNodesResponseBody) SetRequestId(v string) *GetSuggestShrinkableNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuggestShrinkableNodesResponseBody) SetResult(v []*GetSuggestShrinkableNodesResponseBodyResult) *GetSuggestShrinkableNodesResponseBody {
	s.Result = v
	return s
}

type GetSuggestShrinkableNodesResponseBodyResult struct {
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	Port *int32  `json:"port,omitempty" xml:"port,omitempty"`
}

func (s GetSuggestShrinkableNodesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSuggestShrinkableNodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSuggestShrinkableNodesResponseBodyResult) SetHost(v string) *GetSuggestShrinkableNodesResponseBodyResult {
	s.Host = &v
	return s
}

func (s *GetSuggestShrinkableNodesResponseBodyResult) SetPort(v int32) *GetSuggestShrinkableNodesResponseBodyResult {
	s.Port = &v
	return s
}

type GetSuggestShrinkableNodesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSuggestShrinkableNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSuggestShrinkableNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuggestShrinkableNodesResponse) GoString() string {
	return s.String()
}

func (s *GetSuggestShrinkableNodesResponse) SetHeaders(v map[string]*string) *GetSuggestShrinkableNodesResponse {
	s.Headers = v
	return s
}

func (s *GetSuggestShrinkableNodesResponse) SetBody(v *GetSuggestShrinkableNodesResponseBody) *GetSuggestShrinkableNodesResponse {
	s.Body = v
	return s
}

type GetTransferableNodesRequest struct {
	Count    *int32  `json:"count,omitempty" xml:"count,omitempty"`
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s GetTransferableNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTransferableNodesRequest) GoString() string {
	return s.String()
}

func (s *GetTransferableNodesRequest) SetCount(v int32) *GetTransferableNodesRequest {
	s.Count = &v
	return s
}

func (s *GetTransferableNodesRequest) SetNodeType(v string) *GetTransferableNodesRequest {
	s.NodeType = &v
	return s
}

type GetTransferableNodesResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetTransferableNodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetTransferableNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTransferableNodesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTransferableNodesResponseBody) SetRequestId(v string) *GetTransferableNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTransferableNodesResponseBody) SetResult(v []*GetTransferableNodesResponseBodyResult) *GetTransferableNodesResponseBody {
	s.Result = v
	return s
}

type GetTransferableNodesResponseBodyResult struct {
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	Port *int32  `json:"port,omitempty" xml:"port,omitempty"`
}

func (s GetTransferableNodesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTransferableNodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTransferableNodesResponseBodyResult) SetHost(v string) *GetTransferableNodesResponseBodyResult {
	s.Host = &v
	return s
}

func (s *GetTransferableNodesResponseBodyResult) SetPort(v int32) *GetTransferableNodesResponseBodyResult {
	s.Port = &v
	return s
}

type GetTransferableNodesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTransferableNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTransferableNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTransferableNodesResponse) GoString() string {
	return s.String()
}

func (s *GetTransferableNodesResponse) SetHeaders(v map[string]*string) *GetTransferableNodesResponse {
	s.Headers = v
	return s
}

func (s *GetTransferableNodesResponse) SetBody(v *GetTransferableNodesResponseBody) *GetTransferableNodesResponse {
	s.Body = v
	return s
}

type InitializeOperationRoleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s InitializeOperationRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeOperationRoleRequest) GoString() string {
	return s.String()
}

func (s *InitializeOperationRoleRequest) SetClientToken(v string) *InitializeOperationRoleRequest {
	s.ClientToken = &v
	return s
}

type InitializeOperationRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InitializeOperationRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeOperationRoleResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeOperationRoleResponseBody) SetRequestId(v string) *InitializeOperationRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeOperationRoleResponseBody) SetResult(v bool) *InitializeOperationRoleResponseBody {
	s.Result = &v
	return s
}

type InitializeOperationRoleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitializeOperationRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitializeOperationRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeOperationRoleResponse) GoString() string {
	return s.String()
}

func (s *InitializeOperationRoleResponse) SetHeaders(v map[string]*string) *InitializeOperationRoleResponse {
	s.Headers = v
	return s
}

func (s *InitializeOperationRoleResponse) SetBody(v *InitializeOperationRoleResponseBody) *InitializeOperationRoleResponse {
	s.Body = v
	return s
}

type InstallAckOperatorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s InstallAckOperatorRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallAckOperatorRequest) GoString() string {
	return s.String()
}

func (s *InstallAckOperatorRequest) SetClientToken(v string) *InstallAckOperatorRequest {
	s.ClientToken = &v
	return s
}

type InstallAckOperatorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InstallAckOperatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallAckOperatorResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAckOperatorResponseBody) SetRequestId(v string) *InstallAckOperatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAckOperatorResponseBody) SetResult(v bool) *InstallAckOperatorResponseBody {
	s.Result = &v
	return s
}

type InstallAckOperatorResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallAckOperatorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallAckOperatorResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallAckOperatorResponse) GoString() string {
	return s.String()
}

func (s *InstallAckOperatorResponse) SetHeaders(v map[string]*string) *InstallAckOperatorResponse {
	s.Headers = v
	return s
}

func (s *InstallAckOperatorResponse) SetBody(v *InstallAckOperatorResponseBody) *InstallAckOperatorResponse {
	s.Body = v
	return s
}

type InstallKibanaSystemPluginRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s InstallKibanaSystemPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallKibanaSystemPluginRequest) GoString() string {
	return s.String()
}

func (s *InstallKibanaSystemPluginRequest) SetClientToken(v string) *InstallKibanaSystemPluginRequest {
	s.ClientToken = &v
	return s
}

type InstallKibanaSystemPluginResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InstallKibanaSystemPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallKibanaSystemPluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallKibanaSystemPluginResponseBody) SetRequestId(v string) *InstallKibanaSystemPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallKibanaSystemPluginResponseBody) SetResult(v []*string) *InstallKibanaSystemPluginResponseBody {
	s.Result = v
	return s
}

type InstallKibanaSystemPluginResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallKibanaSystemPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallKibanaSystemPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallKibanaSystemPluginResponse) GoString() string {
	return s.String()
}

func (s *InstallKibanaSystemPluginResponse) SetHeaders(v map[string]*string) *InstallKibanaSystemPluginResponse {
	s.Headers = v
	return s
}

func (s *InstallKibanaSystemPluginResponse) SetBody(v *InstallKibanaSystemPluginResponseBody) *InstallKibanaSystemPluginResponse {
	s.Body = v
	return s
}

type InstallLogstashSystemPluginRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s InstallLogstashSystemPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallLogstashSystemPluginRequest) GoString() string {
	return s.String()
}

func (s *InstallLogstashSystemPluginRequest) SetClientToken(v string) *InstallLogstashSystemPluginRequest {
	s.ClientToken = &v
	return s
}

type InstallLogstashSystemPluginResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InstallLogstashSystemPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallLogstashSystemPluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallLogstashSystemPluginResponseBody) SetRequestId(v string) *InstallLogstashSystemPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallLogstashSystemPluginResponseBody) SetResult(v []*string) *InstallLogstashSystemPluginResponseBody {
	s.Result = v
	return s
}

type InstallLogstashSystemPluginResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallLogstashSystemPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallLogstashSystemPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallLogstashSystemPluginResponse) GoString() string {
	return s.String()
}

func (s *InstallLogstashSystemPluginResponse) SetHeaders(v map[string]*string) *InstallLogstashSystemPluginResponse {
	s.Headers = v
	return s
}

func (s *InstallLogstashSystemPluginResponse) SetBody(v *InstallLogstashSystemPluginResponseBody) *InstallLogstashSystemPluginResponse {
	s.Body = v
	return s
}

type InstallSystemPluginRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s InstallSystemPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallSystemPluginRequest) GoString() string {
	return s.String()
}

func (s *InstallSystemPluginRequest) SetClientToken(v string) *InstallSystemPluginRequest {
	s.ClientToken = &v
	return s
}

type InstallSystemPluginResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InstallSystemPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallSystemPluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallSystemPluginResponseBody) SetRequestId(v string) *InstallSystemPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallSystemPluginResponseBody) SetResult(v []*string) *InstallSystemPluginResponseBody {
	s.Result = v
	return s
}

type InstallSystemPluginResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallSystemPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallSystemPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallSystemPluginResponse) GoString() string {
	return s.String()
}

func (s *InstallSystemPluginResponse) SetHeaders(v map[string]*string) *InstallSystemPluginResponse {
	s.Headers = v
	return s
}

func (s *InstallSystemPluginResponse) SetBody(v *InstallSystemPluginResponseBody) *InstallSystemPluginResponse {
	s.Body = v
	return s
}

type InstallUserPluginsResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InstallUserPluginsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallUserPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *InstallUserPluginsResponseBody) SetRequestId(v string) *InstallUserPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallUserPluginsResponseBody) SetResult(v []*string) *InstallUserPluginsResponseBody {
	s.Result = v
	return s
}

type InstallUserPluginsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallUserPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallUserPluginsResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallUserPluginsResponse) GoString() string {
	return s.String()
}

func (s *InstallUserPluginsResponse) SetHeaders(v map[string]*string) *InstallUserPluginsResponse {
	s.Headers = v
	return s
}

func (s *InstallUserPluginsResponse) SetBody(v *InstallUserPluginsResponseBody) *InstallUserPluginsResponse {
	s.Body = v
	return s
}

type InterruptElasticsearchTaskRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s InterruptElasticsearchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s InterruptElasticsearchTaskRequest) GoString() string {
	return s.String()
}

func (s *InterruptElasticsearchTaskRequest) SetClientToken(v string) *InterruptElasticsearchTaskRequest {
	s.ClientToken = &v
	return s
}

type InterruptElasticsearchTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InterruptElasticsearchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InterruptElasticsearchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *InterruptElasticsearchTaskResponseBody) SetCode(v string) *InterruptElasticsearchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *InterruptElasticsearchTaskResponseBody) SetMessage(v string) *InterruptElasticsearchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *InterruptElasticsearchTaskResponseBody) SetRequestId(v string) *InterruptElasticsearchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *InterruptElasticsearchTaskResponseBody) SetResult(v bool) *InterruptElasticsearchTaskResponseBody {
	s.Result = &v
	return s
}

type InterruptElasticsearchTaskResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InterruptElasticsearchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InterruptElasticsearchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s InterruptElasticsearchTaskResponse) GoString() string {
	return s.String()
}

func (s *InterruptElasticsearchTaskResponse) SetHeaders(v map[string]*string) *InterruptElasticsearchTaskResponse {
	s.Headers = v
	return s
}

func (s *InterruptElasticsearchTaskResponse) SetBody(v *InterruptElasticsearchTaskResponseBody) *InterruptElasticsearchTaskResponse {
	s.Body = v
	return s
}

type InterruptLogstashTaskRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s InterruptLogstashTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s InterruptLogstashTaskRequest) GoString() string {
	return s.String()
}

func (s *InterruptLogstashTaskRequest) SetClientToken(v string) *InterruptLogstashTaskRequest {
	s.ClientToken = &v
	return s
}

type InterruptLogstashTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InterruptLogstashTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InterruptLogstashTaskResponseBody) GoString() string {
	return s.String()
}

func (s *InterruptLogstashTaskResponseBody) SetCode(v string) *InterruptLogstashTaskResponseBody {
	s.Code = &v
	return s
}

func (s *InterruptLogstashTaskResponseBody) SetMessage(v string) *InterruptLogstashTaskResponseBody {
	s.Message = &v
	return s
}

func (s *InterruptLogstashTaskResponseBody) SetRequestId(v string) *InterruptLogstashTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *InterruptLogstashTaskResponseBody) SetResult(v bool) *InterruptLogstashTaskResponseBody {
	s.Result = &v
	return s
}

type InterruptLogstashTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InterruptLogstashTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InterruptLogstashTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s InterruptLogstashTaskResponse) GoString() string {
	return s.String()
}

func (s *InterruptLogstashTaskResponse) SetHeaders(v map[string]*string) *InterruptLogstashTaskResponse {
	s.Headers = v
	return s
}

func (s *InterruptLogstashTaskResponse) SetBody(v *InterruptLogstashTaskResponseBody) *InterruptLogstashTaskResponse {
	s.Body = v
	return s
}

type ListAckClustersRequest struct {
	Page  *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size  *int32  `json:"size,omitempty" xml:"size,omitempty"`
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListAckClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAckClustersRequest) GoString() string {
	return s.String()
}

func (s *ListAckClustersRequest) SetPage(v int32) *ListAckClustersRequest {
	s.Page = &v
	return s
}

func (s *ListAckClustersRequest) SetSize(v int32) *ListAckClustersRequest {
	s.Size = &v
	return s
}

func (s *ListAckClustersRequest) SetVpcId(v string) *ListAckClustersRequest {
	s.VpcId = &v
	return s
}

type ListAckClustersResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListAckClustersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAckClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAckClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAckClustersResponseBody) SetRequestId(v string) *ListAckClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAckClustersResponseBody) SetResult(v []*ListAckClustersResponseBodyResult) *ListAckClustersResponseBody {
	s.Result = v
	return s
}

type ListAckClustersResponseBodyResult struct {
	ClusterId   *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	VpcId       *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListAckClustersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAckClustersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAckClustersResponseBodyResult) SetClusterId(v string) *ListAckClustersResponseBodyResult {
	s.ClusterId = &v
	return s
}

func (s *ListAckClustersResponseBodyResult) SetClusterType(v string) *ListAckClustersResponseBodyResult {
	s.ClusterType = &v
	return s
}

func (s *ListAckClustersResponseBodyResult) SetName(v string) *ListAckClustersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAckClustersResponseBodyResult) SetVpcId(v string) *ListAckClustersResponseBodyResult {
	s.VpcId = &v
	return s
}

type ListAckClustersResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAckClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAckClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAckClustersResponse) GoString() string {
	return s.String()
}

func (s *ListAckClustersResponse) SetHeaders(v map[string]*string) *ListAckClustersResponse {
	s.Headers = v
	return s
}

func (s *ListAckClustersResponse) SetBody(v *ListAckClustersResponseBody) *ListAckClustersResponse {
	s.Body = v
	return s
}

type ListAckNamespacesRequest struct {
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAckNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAckNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListAckNamespacesRequest) SetPage(v int32) *ListAckNamespacesRequest {
	s.Page = &v
	return s
}

func (s *ListAckNamespacesRequest) SetSize(v int32) *ListAckNamespacesRequest {
	s.Size = &v
	return s
}

type ListAckNamespacesResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListAckNamespacesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAckNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAckNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAckNamespacesResponseBody) SetRequestId(v string) *ListAckNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAckNamespacesResponseBody) SetResult(v []*ListAckNamespacesResponseBodyResult) *ListAckNamespacesResponseBody {
	s.Result = v
	return s
}

type ListAckNamespacesResponseBodyResult struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAckNamespacesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAckNamespacesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAckNamespacesResponseBodyResult) SetNamespace(v string) *ListAckNamespacesResponseBodyResult {
	s.Namespace = &v
	return s
}

func (s *ListAckNamespacesResponseBodyResult) SetStatus(v string) *ListAckNamespacesResponseBodyResult {
	s.Status = &v
	return s
}

type ListAckNamespacesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAckNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAckNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAckNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListAckNamespacesResponse) SetHeaders(v map[string]*string) *ListAckNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListAckNamespacesResponse) SetBody(v *ListAckNamespacesResponseBody) *ListAckNamespacesResponse {
	s.Body = v
	return s
}

type ListAllNodeRequest struct {
	Extended *bool `json:"extended,omitempty" xml:"extended,omitempty"`
}

func (s ListAllNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAllNodeRequest) GoString() string {
	return s.String()
}

func (s *ListAllNodeRequest) SetExtended(v bool) *ListAllNodeRequest {
	s.Extended = &v
	return s
}

type ListAllNodeResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListAllNodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAllNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllNodeResponseBody) SetRequestId(v string) *ListAllNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllNodeResponseBody) SetResult(v []*ListAllNodeResponseBodyResult) *ListAllNodeResponseBody {
	s.Result = v
	return s
}

type ListAllNodeResponseBodyResult struct {
	CpuPercent      *string `json:"cpuPercent,omitempty" xml:"cpuPercent,omitempty"`
	DiskUsedPercent *string `json:"diskUsedPercent,omitempty" xml:"diskUsedPercent,omitempty"`
	Health          *string `json:"health,omitempty" xml:"health,omitempty"`
	HeapPercent     *string `json:"heapPercent,omitempty" xml:"heapPercent,omitempty"`
	Host            *string `json:"host,omitempty" xml:"host,omitempty"`
	LoadOneM        *string `json:"loadOneM,omitempty" xml:"loadOneM,omitempty"`
	NodeType        *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	Port            *int32  `json:"port,omitempty" xml:"port,omitempty"`
	ZoneId          *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListAllNodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAllNodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAllNodeResponseBodyResult) SetCpuPercent(v string) *ListAllNodeResponseBodyResult {
	s.CpuPercent = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetDiskUsedPercent(v string) *ListAllNodeResponseBodyResult {
	s.DiskUsedPercent = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetHealth(v string) *ListAllNodeResponseBodyResult {
	s.Health = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetHeapPercent(v string) *ListAllNodeResponseBodyResult {
	s.HeapPercent = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetHost(v string) *ListAllNodeResponseBodyResult {
	s.Host = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetLoadOneM(v string) *ListAllNodeResponseBodyResult {
	s.LoadOneM = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetNodeType(v string) *ListAllNodeResponseBodyResult {
	s.NodeType = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetPort(v int32) *ListAllNodeResponseBodyResult {
	s.Port = &v
	return s
}

func (s *ListAllNodeResponseBodyResult) SetZoneId(v string) *ListAllNodeResponseBodyResult {
	s.ZoneId = &v
	return s
}

type ListAllNodeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAllNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAllNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllNodeResponse) GoString() string {
	return s.String()
}

func (s *ListAllNodeResponse) SetHeaders(v map[string]*string) *ListAllNodeResponse {
	s.Headers = v
	return s
}

func (s *ListAllNodeResponse) SetBody(v *ListAllNodeResponseBody) *ListAllNodeResponse {
	s.Body = v
	return s
}

type ListAlternativeSnapshotReposRequest struct {
	AlreadySetItems *bool `json:"alreadySetItems,omitempty" xml:"alreadySetItems,omitempty"`
}

func (s ListAlternativeSnapshotReposRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlternativeSnapshotReposRequest) GoString() string {
	return s.String()
}

func (s *ListAlternativeSnapshotReposRequest) SetAlreadySetItems(v bool) *ListAlternativeSnapshotReposRequest {
	s.AlreadySetItems = &v
	return s
}

type ListAlternativeSnapshotReposResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListAlternativeSnapshotReposResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAlternativeSnapshotReposResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlternativeSnapshotReposResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlternativeSnapshotReposResponseBody) SetRequestId(v string) *ListAlternativeSnapshotReposResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlternativeSnapshotReposResponseBody) SetResult(v []*ListAlternativeSnapshotReposResponseBodyResult) *ListAlternativeSnapshotReposResponseBody {
	s.Result = v
	return s
}

type ListAlternativeSnapshotReposResponseBodyResult struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	RepoPath   *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
}

func (s ListAlternativeSnapshotReposResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAlternativeSnapshotReposResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAlternativeSnapshotReposResponseBodyResult) SetInstanceId(v string) *ListAlternativeSnapshotReposResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListAlternativeSnapshotReposResponseBodyResult) SetRepoPath(v string) *ListAlternativeSnapshotReposResponseBodyResult {
	s.RepoPath = &v
	return s
}

type ListAlternativeSnapshotReposResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlternativeSnapshotReposResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlternativeSnapshotReposResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlternativeSnapshotReposResponse) GoString() string {
	return s.String()
}

func (s *ListAlternativeSnapshotReposResponse) SetHeaders(v map[string]*string) *ListAlternativeSnapshotReposResponse {
	s.Headers = v
	return s
}

func (s *ListAlternativeSnapshotReposResponse) SetBody(v *ListAlternativeSnapshotReposResponseBody) *ListAlternativeSnapshotReposResponse {
	s.Body = v
	return s
}

type ListApmRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Output      *string `json:"output,omitempty" xml:"output,omitempty"`
	Page        *int64  `json:"page,omitempty" xml:"page,omitempty"`
	Size        *int64  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListApmRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApmRequest) GoString() string {
	return s.String()
}

func (s *ListApmRequest) SetDescription(v string) *ListApmRequest {
	s.Description = &v
	return s
}

func (s *ListApmRequest) SetInstanceId(v string) *ListApmRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApmRequest) SetOutput(v string) *ListApmRequest {
	s.Output = &v
	return s
}

func (s *ListApmRequest) SetPage(v int64) *ListApmRequest {
	s.Page = &v
	return s
}

func (s *ListApmRequest) SetSize(v int64) *ListApmRequest {
	s.Size = &v
	return s
}

type ListApmResponseBody struct {
	Headers   *ListApmResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListApmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListApmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApmResponseBody) GoString() string {
	return s.String()
}

func (s *ListApmResponseBody) SetHeaders(v *ListApmResponseBodyHeaders) *ListApmResponseBody {
	s.Headers = v
	return s
}

func (s *ListApmResponseBody) SetRequestId(v string) *ListApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApmResponseBody) SetResult(v []*ListApmResponseBodyResult) *ListApmResponseBody {
	s.Result = v
	return s
}

type ListApmResponseBodyHeaders struct {
	XTotalCount *int64 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListApmResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListApmResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListApmResponseBodyHeaders) SetXTotalCount(v int64) *ListApmResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListApmResponseBodyResult struct {
	CreatedAt        *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	DeployedReplica  *int64  `json:"deployedReplica,omitempty" xml:"deployedReplica,omitempty"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId       *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	NodeAmount       *int64  `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	OutputES         *string `json:"outputES,omitempty" xml:"outputES,omitempty"`
	OutputESUserName *string `json:"outputESUserName,omitempty" xml:"outputESUserName,omitempty"`
	OwnerId          *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	PaymentType      *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	Region           *string `json:"region,omitempty" xml:"region,omitempty"`
	Replica          *int64  `json:"replica,omitempty" xml:"replica,omitempty"`
	ResourceSpec     *string `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
	Version          *string `json:"version,omitempty" xml:"version,omitempty"`
	VpcId            *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea           *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId        *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s ListApmResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListApmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListApmResponseBodyResult) SetCreatedAt(v string) *ListApmResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListApmResponseBodyResult) SetDeployedReplica(v int64) *ListApmResponseBodyResult {
	s.DeployedReplica = &v
	return s
}

func (s *ListApmResponseBodyResult) SetDescription(v string) *ListApmResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListApmResponseBodyResult) SetInstanceId(v string) *ListApmResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListApmResponseBodyResult) SetNodeAmount(v int64) *ListApmResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *ListApmResponseBodyResult) SetOutputES(v string) *ListApmResponseBodyResult {
	s.OutputES = &v
	return s
}

func (s *ListApmResponseBodyResult) SetOutputESUserName(v string) *ListApmResponseBodyResult {
	s.OutputESUserName = &v
	return s
}

func (s *ListApmResponseBodyResult) SetOwnerId(v string) *ListApmResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListApmResponseBodyResult) SetPaymentType(v string) *ListApmResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *ListApmResponseBodyResult) SetRegion(v string) *ListApmResponseBodyResult {
	s.Region = &v
	return s
}

func (s *ListApmResponseBodyResult) SetReplica(v int64) *ListApmResponseBodyResult {
	s.Replica = &v
	return s
}

func (s *ListApmResponseBodyResult) SetResourceSpec(v string) *ListApmResponseBodyResult {
	s.ResourceSpec = &v
	return s
}

func (s *ListApmResponseBodyResult) SetStatus(v string) *ListApmResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListApmResponseBodyResult) SetVersion(v string) *ListApmResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ListApmResponseBodyResult) SetVpcId(v string) *ListApmResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *ListApmResponseBodyResult) SetVsArea(v string) *ListApmResponseBodyResult {
	s.VsArea = &v
	return s
}

func (s *ListApmResponseBodyResult) SetVswitchId(v string) *ListApmResponseBodyResult {
	s.VswitchId = &v
	return s
}

type ListApmResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApmResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApmResponse) GoString() string {
	return s.String()
}

func (s *ListApmResponse) SetHeaders(v map[string]*string) *ListApmResponse {
	s.Headers = v
	return s
}

func (s *ListApmResponse) SetBody(v *ListApmResponseBody) *ListApmResponse {
	s.Body = v
	return s
}

type ListAvailableEsInstanceIdsResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListAvailableEsInstanceIdsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAvailableEsInstanceIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEsInstanceIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableEsInstanceIdsResponseBody) SetRequestId(v string) *ListAvailableEsInstanceIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBody) SetResult(v []*ListAvailableEsInstanceIdsResponseBodyResult) *ListAvailableEsInstanceIdsResponseBody {
	s.Result = v
	return s
}

type ListAvailableEsInstanceIdsResponseBodyResult struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Endpoint       *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	EsInstanceId   *string `json:"esInstanceId,omitempty" xml:"esInstanceId,omitempty"`
	KibanaEndpoint *string `json:"kibanaEndpoint,omitempty" xml:"kibanaEndpoint,omitempty"`
}

func (s ListAvailableEsInstanceIdsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEsInstanceIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) SetDescription(v string) *ListAvailableEsInstanceIdsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) SetEndpoint(v string) *ListAvailableEsInstanceIdsResponseBodyResult {
	s.Endpoint = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) SetEsInstanceId(v string) *ListAvailableEsInstanceIdsResponseBodyResult {
	s.EsInstanceId = &v
	return s
}

func (s *ListAvailableEsInstanceIdsResponseBodyResult) SetKibanaEndpoint(v string) *ListAvailableEsInstanceIdsResponseBodyResult {
	s.KibanaEndpoint = &v
	return s
}

type ListAvailableEsInstanceIdsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAvailableEsInstanceIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableEsInstanceIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEsInstanceIdsResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableEsInstanceIdsResponse) SetHeaders(v map[string]*string) *ListAvailableEsInstanceIdsResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableEsInstanceIdsResponse) SetBody(v *ListAvailableEsInstanceIdsResponseBody) *ListAvailableEsInstanceIdsResponse {
	s.Body = v
	return s
}

type ListCollectorsRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Page       *int32  `json:"page,omitempty" xml:"page,omitempty"`
	ResId      *string `json:"resId,omitempty" xml:"resId,omitempty"`
	Size       *int32  `json:"size,omitempty" xml:"size,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListCollectorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCollectorsRequest) GoString() string {
	return s.String()
}

func (s *ListCollectorsRequest) SetInstanceId(v string) *ListCollectorsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCollectorsRequest) SetName(v string) *ListCollectorsRequest {
	s.Name = &v
	return s
}

func (s *ListCollectorsRequest) SetPage(v int32) *ListCollectorsRequest {
	s.Page = &v
	return s
}

func (s *ListCollectorsRequest) SetResId(v string) *ListCollectorsRequest {
	s.ResId = &v
	return s
}

func (s *ListCollectorsRequest) SetSize(v int32) *ListCollectorsRequest {
	s.Size = &v
	return s
}

func (s *ListCollectorsRequest) SetSourceType(v string) *ListCollectorsRequest {
	s.SourceType = &v
	return s
}

type ListCollectorsResponseBody struct {
	Headers   *ListCollectorsResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCollectorsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListCollectorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCollectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBody) SetHeaders(v *ListCollectorsResponseBodyHeaders) *ListCollectorsResponseBody {
	s.Headers = v
	return s
}

func (s *ListCollectorsResponseBody) SetRequestId(v string) *ListCollectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollectorsResponseBody) SetResult(v []*ListCollectorsResponseBodyResult) *ListCollectorsResponseBody {
	s.Result = v
	return s
}

type ListCollectorsResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListCollectorsResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCollectorsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyHeaders) SetXTotalCount(v int32) *ListCollectorsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListCollectorsResponseBodyResult struct {
	CollectorPaths []*string                                        `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	Configs        []*ListCollectorsResponseBodyResultConfigs       `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	DryRun         *bool                                            `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	ExtendConfigs  []*ListCollectorsResponseBodyResultExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	GmtCreatedTime *string                                          `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	GmtUpdateTime  *string                                          `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	Name           *string                                          `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId        *string                                          `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	ResId          *string                                          `json:"resId,omitempty" xml:"resId,omitempty"`
	ResType        *string                                          `json:"resType,omitempty" xml:"resType,omitempty"`
	ResVersion     *string                                          `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	Status         *string                                          `json:"status,omitempty" xml:"status,omitempty"`
	VpcId          *string                                          `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListCollectorsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCollectorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyResult) SetCollectorPaths(v []*string) *ListCollectorsResponseBodyResult {
	s.CollectorPaths = v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetConfigs(v []*ListCollectorsResponseBodyResultConfigs) *ListCollectorsResponseBodyResult {
	s.Configs = v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetDryRun(v bool) *ListCollectorsResponseBodyResult {
	s.DryRun = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetExtendConfigs(v []*ListCollectorsResponseBodyResultExtendConfigs) *ListCollectorsResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetGmtCreatedTime(v string) *ListCollectorsResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetGmtUpdateTime(v string) *ListCollectorsResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetName(v string) *ListCollectorsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetOwnerId(v string) *ListCollectorsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetResId(v string) *ListCollectorsResponseBodyResult {
	s.ResId = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetResType(v string) *ListCollectorsResponseBodyResult {
	s.ResType = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetResVersion(v string) *ListCollectorsResponseBodyResult {
	s.ResVersion = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetStatus(v string) *ListCollectorsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetVpcId(v string) *ListCollectorsResponseBodyResult {
	s.VpcId = &v
	return s
}

type ListCollectorsResponseBodyResultConfigs struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ListCollectorsResponseBodyResultConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListCollectorsResponseBodyResultConfigs) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyResultConfigs) SetContent(v string) *ListCollectorsResponseBodyResultConfigs {
	s.Content = &v
	return s
}

func (s *ListCollectorsResponseBodyResultConfigs) SetFileName(v string) *ListCollectorsResponseBodyResultConfigs {
	s.FileName = &v
	return s
}

type ListCollectorsResponseBodyResultExtendConfigs struct {
	ConfigType       *string                                                  `json:"configType,omitempty" xml:"configType,omitempty"`
	EnableMonitoring *bool                                                    `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	GroupId          *string                                                  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Host             *string                                                  `json:"host,omitempty" xml:"host,omitempty"`
	Hosts            []*string                                                `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	InstanceId       *string                                                  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType     *string                                                  `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	KibanaHost       *string                                                  `json:"kibanaHost,omitempty" xml:"kibanaHost,omitempty"`
	Machines         []*ListCollectorsResponseBodyResultExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	Protocol         *string                                                  `json:"protocol,omitempty" xml:"protocol,omitempty"`
	SuccessPodsCount *string                                                  `json:"successPodsCount,omitempty" xml:"successPodsCount,omitempty"`
	TotalPodsCount   *string                                                  `json:"totalPodsCount,omitempty" xml:"totalPodsCount,omitempty"`
	Type             *string                                                  `json:"type,omitempty" xml:"type,omitempty"`
	UserName         *string                                                  `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListCollectorsResponseBodyResultExtendConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListCollectorsResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetConfigType(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetEnableMonitoring(v bool) *ListCollectorsResponseBodyResultExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetGroupId(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetHost(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Host = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetHosts(v []*string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Hosts = v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetInstanceId(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetInstanceType(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetKibanaHost(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.KibanaHost = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetMachines(v []*ListCollectorsResponseBodyResultExtendConfigsMachines) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Machines = v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetProtocol(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetSuccessPodsCount(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.SuccessPodsCount = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetTotalPodsCount(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.TotalPodsCount = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetType(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Type = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetUserName(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.UserName = &v
	return s
}

type ListCollectorsResponseBodyResultExtendConfigsMachines struct {
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListCollectorsResponseBodyResultExtendConfigsMachines) String() string {
	return tea.Prettify(s)
}

func (s ListCollectorsResponseBodyResultExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyResultExtendConfigsMachines) SetAgentStatus(v string) *ListCollectorsResponseBodyResultExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigsMachines) SetInstanceId(v string) *ListCollectorsResponseBodyResultExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

type ListCollectorsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCollectorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCollectorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCollectorsResponse) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponse) SetHeaders(v map[string]*string) *ListCollectorsResponse {
	s.Headers = v
	return s
}

func (s *ListCollectorsResponse) SetBody(v *ListCollectorsResponseBody) *ListCollectorsResponse {
	s.Body = v
	return s
}

type ListComponentIndicesRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Page *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListComponentIndicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesRequest) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesRequest) SetName(v string) *ListComponentIndicesRequest {
	s.Name = &v
	return s
}

func (s *ListComponentIndicesRequest) SetPage(v int32) *ListComponentIndicesRequest {
	s.Page = &v
	return s
}

func (s *ListComponentIndicesRequest) SetSize(v int32) *ListComponentIndicesRequest {
	s.Size = &v
	return s
}

type ListComponentIndicesResponseBody struct {
	Headers   *ListComponentIndicesResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListComponentIndicesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListComponentIndicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBody) SetHeaders(v *ListComponentIndicesResponseBodyHeaders) *ListComponentIndicesResponseBody {
	s.Headers = v
	return s
}

func (s *ListComponentIndicesResponseBody) SetRequestId(v string) *ListComponentIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentIndicesResponseBody) SetResult(v []*ListComponentIndicesResponseBodyResult) *ListComponentIndicesResponseBody {
	s.Result = v
	return s
}

type ListComponentIndicesResponseBodyHeaders struct {
	XTotalCount *int64 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListComponentIndicesResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyHeaders) SetXTotalCount(v int64) *ListComponentIndicesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListComponentIndicesResponseBodyResult struct {
	Composed []*string                                      `json:"composed,omitempty" xml:"composed,omitempty" type:"Repeated"`
	Content  *ListComponentIndicesResponseBodyResultContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	Name     *string                                        `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListComponentIndicesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResult) SetComposed(v []*string) *ListComponentIndicesResponseBodyResult {
	s.Composed = v
	return s
}

func (s *ListComponentIndicesResponseBodyResult) SetContent(v *ListComponentIndicesResponseBodyResultContent) *ListComponentIndicesResponseBodyResult {
	s.Content = v
	return s
}

func (s *ListComponentIndicesResponseBodyResult) SetName(v string) *ListComponentIndicesResponseBodyResult {
	s.Name = &v
	return s
}

type ListComponentIndicesResponseBodyResultContent struct {
	Meta     *ListComponentIndicesResponseBodyResultContentMeta     `json:"_meta,omitempty" xml:"_meta,omitempty" type:"Struct"`
	Template *ListComponentIndicesResponseBodyResultContentTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	Version  *int64                                                 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListComponentIndicesResponseBodyResultContent) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContent) SetMeta(v *ListComponentIndicesResponseBodyResultContentMeta) *ListComponentIndicesResponseBodyResultContent {
	s.Meta = v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContent) SetTemplate(v *ListComponentIndicesResponseBodyResultContentTemplate) *ListComponentIndicesResponseBodyResultContent {
	s.Template = v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContent) SetVersion(v int64) *ListComponentIndicesResponseBodyResultContent {
	s.Version = &v
	return s
}

type ListComponentIndicesResponseBodyResultContentMeta struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Managed     *bool   `json:"managed,omitempty" xml:"managed,omitempty"`
}

func (s ListComponentIndicesResponseBodyResultContentMeta) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContentMeta) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContentMeta) SetDescription(v string) *ListComponentIndicesResponseBodyResultContentMeta {
	s.Description = &v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContentMeta) SetManaged(v bool) *ListComponentIndicesResponseBodyResultContentMeta {
	s.Managed = &v
	return s
}

type ListComponentIndicesResponseBodyResultContentTemplate struct {
	Settings *ListComponentIndicesResponseBodyResultContentTemplateSettings `json:"settings,omitempty" xml:"settings,omitempty" type:"Struct"`
}

func (s ListComponentIndicesResponseBodyResultContentTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContentTemplate) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContentTemplate) SetSettings(v *ListComponentIndicesResponseBodyResultContentTemplateSettings) *ListComponentIndicesResponseBodyResultContentTemplate {
	s.Settings = v
	return s
}

type ListComponentIndicesResponseBodyResultContentTemplateSettings struct {
	Index *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex `json:"index,omitempty" xml:"index,omitempty" type:"Struct"`
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettings) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettings) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettings) SetIndex(v *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) *ListComponentIndicesResponseBodyResultContentTemplateSettings {
	s.Index = v
	return s
}

type ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex struct {
	Codec     *string                                                                      `json:"codec,omitempty" xml:"codec,omitempty"`
	Lifecycle *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle `json:"lifecycle,omitempty" xml:"lifecycle,omitempty" type:"Struct"`
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) SetCodec(v string) *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex {
	s.Codec = &v
	return s
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex) SetLifecycle(v *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndex {
	s.Lifecycle = v
	return s
}

type ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle) SetName(v string) *ListComponentIndicesResponseBodyResultContentTemplateSettingsIndexLifecycle {
	s.Name = &v
	return s
}

type ListComponentIndicesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListComponentIndicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListComponentIndicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListComponentIndicesResponse) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesResponse) SetHeaders(v map[string]*string) *ListComponentIndicesResponse {
	s.Headers = v
	return s
}

func (s *ListComponentIndicesResponse) SetBody(v *ListComponentIndicesResponseBody) *ListComponentIndicesResponse {
	s.Body = v
	return s
}

type ListConnectedClustersResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListConnectedClustersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListConnectedClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectedClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectedClustersResponseBody) SetRequestId(v string) *ListConnectedClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectedClustersResponseBody) SetResult(v *ListConnectedClustersResponseBodyResult) *ListConnectedClustersResponseBody {
	s.Result = v
	return s
}

type ListConnectedClustersResponseBodyResult struct {
	Result []*ListConnectedClustersResponseBodyResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListConnectedClustersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListConnectedClustersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListConnectedClustersResponseBodyResult) SetResult(v []*ListConnectedClustersResponseBodyResultResult) *ListConnectedClustersResponseBodyResult {
	s.Result = v
	return s
}

type ListConnectedClustersResponseBodyResultResult struct {
	Instances   *string `json:"instances,omitempty" xml:"instances,omitempty"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s ListConnectedClustersResponseBodyResultResult) String() string {
	return tea.Prettify(s)
}

func (s ListConnectedClustersResponseBodyResultResult) GoString() string {
	return s.String()
}

func (s *ListConnectedClustersResponseBodyResultResult) SetInstances(v string) *ListConnectedClustersResponseBodyResultResult {
	s.Instances = &v
	return s
}

func (s *ListConnectedClustersResponseBodyResultResult) SetNetworkType(v string) *ListConnectedClustersResponseBodyResultResult {
	s.NetworkType = &v
	return s
}

type ListConnectedClustersResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConnectedClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConnectedClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectedClustersResponse) GoString() string {
	return s.String()
}

func (s *ListConnectedClustersResponse) SetHeaders(v map[string]*string) *ListConnectedClustersResponse {
	s.Headers = v
	return s
}

func (s *ListConnectedClustersResponse) SetBody(v *ListConnectedClustersResponseBody) *ListConnectedClustersResponse {
	s.Body = v
	return s
}

type ListDataStreamsRequest struct {
	IsManaged *bool   `json:"isManaged,omitempty" xml:"isManaged,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListDataStreamsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsRequest) GoString() string {
	return s.String()
}

func (s *ListDataStreamsRequest) SetIsManaged(v bool) *ListDataStreamsRequest {
	s.IsManaged = &v
	return s
}

func (s *ListDataStreamsRequest) SetName(v string) *ListDataStreamsRequest {
	s.Name = &v
	return s
}

type ListDataStreamsResponseBody struct {
	Headers   *ListDataStreamsResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDataStreamsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDataStreamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBody) SetHeaders(v *ListDataStreamsResponseBodyHeaders) *ListDataStreamsResponseBody {
	s.Headers = v
	return s
}

func (s *ListDataStreamsResponseBody) SetRequestId(v string) *ListDataStreamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataStreamsResponseBody) SetResult(v []*ListDataStreamsResponseBodyResult) *ListDataStreamsResponseBody {
	s.Result = v
	return s
}

type ListDataStreamsResponseBodyHeaders struct {
	XManagedCount       *int32 `json:"X-Managed-Count,omitempty" xml:"X-Managed-Count,omitempty"`
	XManagedStorageSize *int64 `json:"X-Managed-StorageSize,omitempty" xml:"X-Managed-StorageSize,omitempty"`
}

func (s ListDataStreamsResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBodyHeaders) SetXManagedCount(v int32) *ListDataStreamsResponseBodyHeaders {
	s.XManagedCount = &v
	return s
}

func (s *ListDataStreamsResponseBodyHeaders) SetXManagedStorageSize(v int64) *ListDataStreamsResponseBodyHeaders {
	s.XManagedStorageSize = &v
	return s
}

type ListDataStreamsResponseBodyResult struct {
	Health             *string                                     `json:"health,omitempty" xml:"health,omitempty"`
	IlmPolicyName      *string                                     `json:"ilmPolicyName,omitempty" xml:"ilmPolicyName,omitempty"`
	IndexTemplateName  *string                                     `json:"indexTemplateName,omitempty" xml:"indexTemplateName,omitempty"`
	Indices            []*ListDataStreamsResponseBodyResultIndices `json:"indices,omitempty" xml:"indices,omitempty" type:"Repeated"`
	ManagedStorageSize *int64                                      `json:"managedStorageSize,omitempty" xml:"managedStorageSize,omitempty"`
	Name               *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	TotalStorageSize   *int64                                      `json:"totalStorageSize,omitempty" xml:"totalStorageSize,omitempty"`
}

func (s ListDataStreamsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBodyResult) SetHealth(v string) *ListDataStreamsResponseBodyResult {
	s.Health = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetIlmPolicyName(v string) *ListDataStreamsResponseBodyResult {
	s.IlmPolicyName = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetIndexTemplateName(v string) *ListDataStreamsResponseBodyResult {
	s.IndexTemplateName = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetIndices(v []*ListDataStreamsResponseBodyResultIndices) *ListDataStreamsResponseBodyResult {
	s.Indices = v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetManagedStorageSize(v int64) *ListDataStreamsResponseBodyResult {
	s.ManagedStorageSize = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetName(v string) *ListDataStreamsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataStreamsResponseBodyResult) SetTotalStorageSize(v int64) *ListDataStreamsResponseBodyResult {
	s.TotalStorageSize = &v
	return s
}

type ListDataStreamsResponseBodyResultIndices struct {
	CreateTime    *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Health        *string `json:"health,omitempty" xml:"health,omitempty"`
	IsManaged     *bool   `json:"isManaged,omitempty" xml:"isManaged,omitempty"`
	ManagedStatus *string `json:"managedStatus,omitempty" xml:"managedStatus,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Size          *int64  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListDataStreamsResponseBodyResultIndices) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsResponseBodyResultIndices) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponseBodyResultIndices) SetCreateTime(v string) *ListDataStreamsResponseBodyResultIndices {
	s.CreateTime = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetHealth(v string) *ListDataStreamsResponseBodyResultIndices {
	s.Health = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetIsManaged(v bool) *ListDataStreamsResponseBodyResultIndices {
	s.IsManaged = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetManagedStatus(v string) *ListDataStreamsResponseBodyResultIndices {
	s.ManagedStatus = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetName(v string) *ListDataStreamsResponseBodyResultIndices {
	s.Name = &v
	return s
}

func (s *ListDataStreamsResponseBodyResultIndices) SetSize(v int64) *ListDataStreamsResponseBodyResultIndices {
	s.Size = &v
	return s
}

type ListDataStreamsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataStreamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataStreamsResponse) GoString() string {
	return s.String()
}

func (s *ListDataStreamsResponse) SetHeaders(v map[string]*string) *ListDataStreamsResponse {
	s.Headers = v
	return s
}

func (s *ListDataStreamsResponse) SetBody(v *ListDataStreamsResponseBody) *ListDataStreamsResponse {
	s.Body = v
	return s
}

type ListDataTasksResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDataTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDataTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponseBody) SetRequestId(v string) *ListDataTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataTasksResponseBody) SetResult(v []*ListDataTasksResponseBodyResult) *ListDataTasksResponseBody {
	s.Result = v
	return s
}

type ListDataTasksResponseBodyResult struct {
	CreateTime    *string                                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	SinkCluster   *ListDataTasksResponseBodyResultSinkCluster   `json:"sinkCluster,omitempty" xml:"sinkCluster,omitempty" type:"Struct"`
	SourceCluster *ListDataTasksResponseBodyResultSourceCluster `json:"sourceCluster,omitempty" xml:"sourceCluster,omitempty" type:"Struct"`
	Status        *string                                       `json:"status,omitempty" xml:"status,omitempty"`
	TaskId        *string                                       `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListDataTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponseBodyResult) SetCreateTime(v string) *ListDataTasksResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListDataTasksResponseBodyResult) SetSinkCluster(v *ListDataTasksResponseBodyResultSinkCluster) *ListDataTasksResponseBodyResult {
	s.SinkCluster = v
	return s
}

func (s *ListDataTasksResponseBodyResult) SetSourceCluster(v *ListDataTasksResponseBodyResultSourceCluster) *ListDataTasksResponseBodyResult {
	s.SourceCluster = v
	return s
}

func (s *ListDataTasksResponseBodyResult) SetStatus(v string) *ListDataTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataTasksResponseBodyResult) SetTaskId(v string) *ListDataTasksResponseBodyResult {
	s.TaskId = &v
	return s
}

type ListDataTasksResponseBodyResultSinkCluster struct {
	DataSourceType  *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Index           *string `json:"index,omitempty" xml:"index,omitempty"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	VpcId           *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VpcInstanceId   *string `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
	VpcInstancePort *string `json:"vpcInstancePort,omitempty" xml:"vpcInstancePort,omitempty"`
}

func (s ListDataTasksResponseBodyResultSinkCluster) String() string {
	return tea.Prettify(s)
}

func (s ListDataTasksResponseBodyResultSinkCluster) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetDataSourceType(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.DataSourceType = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetEndpoint(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.Endpoint = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetIndex(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.Index = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetType(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.Type = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetVpcId(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.VpcId = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetVpcInstanceId(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.VpcInstanceId = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSinkCluster) SetVpcInstancePort(v string) *ListDataTasksResponseBodyResultSinkCluster {
	s.VpcInstancePort = &v
	return s
}

type ListDataTasksResponseBodyResultSourceCluster struct {
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	Index          *string `json:"index,omitempty" xml:"index,omitempty"`
	Mapping        *string `json:"mapping,omitempty" xml:"mapping,omitempty"`
	Routing        *string `json:"routing,omitempty" xml:"routing,omitempty"`
	Settings       *string `json:"settings,omitempty" xml:"settings,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDataTasksResponseBodyResultSourceCluster) String() string {
	return tea.Prettify(s)
}

func (s ListDataTasksResponseBodyResultSourceCluster) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetDataSourceType(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.DataSourceType = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetIndex(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Index = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetMapping(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Mapping = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetRouting(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Routing = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetSettings(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Settings = &v
	return s
}

func (s *ListDataTasksResponseBodyResultSourceCluster) SetType(v string) *ListDataTasksResponseBodyResultSourceCluster {
	s.Type = &v
	return s
}

type ListDataTasksResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDataTasksResponse) SetHeaders(v map[string]*string) *ListDataTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDataTasksResponse) SetBody(v *ListDataTasksResponseBody) *ListDataTasksResponse {
	s.Body = v
	return s
}

type ListDefaultCollectorConfigurationsRequest struct {
	ResType    *string `json:"resType,omitempty" xml:"resType,omitempty"`
	ResVersion *string `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListDefaultCollectorConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultCollectorConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListDefaultCollectorConfigurationsRequest) SetResType(v string) *ListDefaultCollectorConfigurationsRequest {
	s.ResType = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsRequest) SetResVersion(v string) *ListDefaultCollectorConfigurationsRequest {
	s.ResVersion = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsRequest) SetSourceType(v string) *ListDefaultCollectorConfigurationsRequest {
	s.SourceType = &v
	return s
}

type ListDefaultCollectorConfigurationsResponseBody struct {
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDefaultCollectorConfigurationsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDefaultCollectorConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultCollectorConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDefaultCollectorConfigurationsResponseBody) SetRequestId(v string) *ListDefaultCollectorConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponseBody) SetResult(v []*ListDefaultCollectorConfigurationsResponseBodyResult) *ListDefaultCollectorConfigurationsResponseBody {
	s.Result = v
	return s
}

type ListDefaultCollectorConfigurationsResponseBodyResult struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ListDefaultCollectorConfigurationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultCollectorConfigurationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDefaultCollectorConfigurationsResponseBodyResult) SetContent(v string) *ListDefaultCollectorConfigurationsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponseBodyResult) SetFileName(v string) *ListDefaultCollectorConfigurationsResponseBodyResult {
	s.FileName = &v
	return s
}

type ListDefaultCollectorConfigurationsResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDefaultCollectorConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDefaultCollectorConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultCollectorConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListDefaultCollectorConfigurationsResponse) SetHeaders(v map[string]*string) *ListDefaultCollectorConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListDefaultCollectorConfigurationsResponse) SetBody(v *ListDefaultCollectorConfigurationsResponseBody) *ListDefaultCollectorConfigurationsResponse {
	s.Body = v
	return s
}

type ListDeprecatedTemplatesRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Page *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListDeprecatedTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeprecatedTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesRequest) SetName(v string) *ListDeprecatedTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListDeprecatedTemplatesRequest) SetPage(v int32) *ListDeprecatedTemplatesRequest {
	s.Page = &v
	return s
}

func (s *ListDeprecatedTemplatesRequest) SetSize(v int32) *ListDeprecatedTemplatesRequest {
	s.Size = &v
	return s
}

type ListDeprecatedTemplatesResponseBody struct {
	Headers   *ListDeprecatedTemplatesResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDeprecatedTemplatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDeprecatedTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeprecatedTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponseBody) SetHeaders(v *ListDeprecatedTemplatesResponseBodyHeaders) *ListDeprecatedTemplatesResponseBody {
	s.Headers = v
	return s
}

func (s *ListDeprecatedTemplatesResponseBody) SetRequestId(v string) *ListDeprecatedTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBody) SetResult(v []*ListDeprecatedTemplatesResponseBodyResult) *ListDeprecatedTemplatesResponseBody {
	s.Result = v
	return s
}

type ListDeprecatedTemplatesResponseBodyHeaders struct {
	XTotalCount *int64 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListDeprecatedTemplatesResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeprecatedTemplatesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponseBodyHeaders) SetXTotalCount(v int64) *ListDeprecatedTemplatesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListDeprecatedTemplatesResponseBodyResult struct {
	DataStream    *bool                                              `json:"dataStream,omitempty" xml:"dataStream,omitempty"`
	IndexPatterns []*string                                          `json:"indexPatterns,omitempty" xml:"indexPatterns,omitempty" type:"Repeated"`
	IndexTemplate *string                                            `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
	Order         *int64                                             `json:"order,omitempty" xml:"order,omitempty"`
	Template      *ListDeprecatedTemplatesResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	Version       *string                                            `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListDeprecatedTemplatesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDeprecatedTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetDataStream(v bool) *ListDeprecatedTemplatesResponseBodyResult {
	s.DataStream = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetIndexPatterns(v []*string) *ListDeprecatedTemplatesResponseBodyResult {
	s.IndexPatterns = v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetIndexTemplate(v string) *ListDeprecatedTemplatesResponseBodyResult {
	s.IndexTemplate = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetOrder(v int64) *ListDeprecatedTemplatesResponseBodyResult {
	s.Order = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetTemplate(v *ListDeprecatedTemplatesResponseBodyResultTemplate) *ListDeprecatedTemplatesResponseBodyResult {
	s.Template = v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResult) SetVersion(v string) *ListDeprecatedTemplatesResponseBodyResult {
	s.Version = &v
	return s
}

type ListDeprecatedTemplatesResponseBodyResultTemplate struct {
	Aliases  *string `json:"aliases,omitempty" xml:"aliases,omitempty"`
	Mappings *string `json:"mappings,omitempty" xml:"mappings,omitempty"`
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s ListDeprecatedTemplatesResponseBodyResultTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListDeprecatedTemplatesResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) SetAliases(v string) *ListDeprecatedTemplatesResponseBodyResultTemplate {
	s.Aliases = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) SetMappings(v string) *ListDeprecatedTemplatesResponseBodyResultTemplate {
	s.Mappings = &v
	return s
}

func (s *ListDeprecatedTemplatesResponseBodyResultTemplate) SetSettings(v string) *ListDeprecatedTemplatesResponseBodyResultTemplate {
	s.Settings = &v
	return s
}

type ListDeprecatedTemplatesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeprecatedTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeprecatedTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeprecatedTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesResponse) SetHeaders(v map[string]*string) *ListDeprecatedTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListDeprecatedTemplatesResponse) SetBody(v *ListDeprecatedTemplatesResponseBody) *ListDeprecatedTemplatesResponse {
	s.Body = v
	return s
}

type ListDiagnoseIndicesRequest struct {
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s ListDiagnoseIndicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseIndicesRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnoseIndicesRequest) SetLang(v string) *ListDiagnoseIndicesRequest {
	s.Lang = &v
	return s
}

type ListDiagnoseIndicesResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDiagnoseIndicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnoseIndicesResponseBody) SetRequestId(v string) *ListDiagnoseIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnoseIndicesResponseBody) SetResult(v []*string) *ListDiagnoseIndicesResponseBody {
	s.Result = v
	return s
}

type ListDiagnoseIndicesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDiagnoseIndicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDiagnoseIndicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseIndicesResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnoseIndicesResponse) SetHeaders(v map[string]*string) *ListDiagnoseIndicesResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnoseIndicesResponse) SetBody(v *ListDiagnoseIndicesResponseBody) *ListDiagnoseIndicesResponse {
	s.Body = v
	return s
}

type ListDiagnoseReportRequest struct {
	Detail    *bool   `json:"detail,omitempty" xml:"detail,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Lang      *string `json:"lang,omitempty" xml:"lang,omitempty"`
	Page      *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size      *int32  `json:"size,omitempty" xml:"size,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Trigger   *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s ListDiagnoseReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportRequest) SetDetail(v bool) *ListDiagnoseReportRequest {
	s.Detail = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetEndTime(v int64) *ListDiagnoseReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetLang(v string) *ListDiagnoseReportRequest {
	s.Lang = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetPage(v int32) *ListDiagnoseReportRequest {
	s.Page = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetSize(v int32) *ListDiagnoseReportRequest {
	s.Size = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetStartTime(v int64) *ListDiagnoseReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListDiagnoseReportRequest) SetTrigger(v string) *ListDiagnoseReportRequest {
	s.Trigger = &v
	return s
}

type ListDiagnoseReportResponseBody struct {
	Headers   *ListDiagnoseReportResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDiagnoseReportResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDiagnoseReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBody) SetHeaders(v *ListDiagnoseReportResponseBodyHeaders) *ListDiagnoseReportResponseBody {
	s.Headers = v
	return s
}

func (s *ListDiagnoseReportResponseBody) SetRequestId(v string) *ListDiagnoseReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnoseReportResponseBody) SetResult(v []*ListDiagnoseReportResponseBodyResult) *ListDiagnoseReportResponseBody {
	s.Result = v
	return s
}

type ListDiagnoseReportResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListDiagnoseReportResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBodyHeaders) SetXTotalCount(v int32) *ListDiagnoseReportResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListDiagnoseReportResponseBodyResult struct {
	CreateTime    *int64                                               `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DiagnoseItems []*ListDiagnoseReportResponseBodyResultDiagnoseItems `json:"diagnoseItems,omitempty" xml:"diagnoseItems,omitempty" type:"Repeated"`
	Health        *string                                              `json:"health,omitempty" xml:"health,omitempty"`
	InstanceId    *string                                              `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ReportId      *string                                              `json:"reportId,omitempty" xml:"reportId,omitempty"`
	State         *string                                              `json:"state,omitempty" xml:"state,omitempty"`
	Trigger       *string                                              `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s ListDiagnoseReportResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBodyResult) SetCreateTime(v int64) *ListDiagnoseReportResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetDiagnoseItems(v []*ListDiagnoseReportResponseBodyResultDiagnoseItems) *ListDiagnoseReportResponseBodyResult {
	s.DiagnoseItems = v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetHealth(v string) *ListDiagnoseReportResponseBodyResult {
	s.Health = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetInstanceId(v string) *ListDiagnoseReportResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetReportId(v string) *ListDiagnoseReportResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetState(v string) *ListDiagnoseReportResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResult) SetTrigger(v string) *ListDiagnoseReportResponseBodyResult {
	s.Trigger = &v
	return s
}

type ListDiagnoseReportResponseBodyResultDiagnoseItems struct {
	Detail *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
	Health *string                                                  `json:"health,omitempty" xml:"health,omitempty"`
	Item   *string                                                  `json:"item,omitempty" xml:"item,omitempty"`
}

func (s ListDiagnoseReportResponseBodyResultDiagnoseItems) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportResponseBodyResultDiagnoseItems) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) SetDetail(v *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) *ListDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Detail = v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) SetHealth(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Health = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItems) SetItem(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItems {
	s.Item = &v
	return s
}

type ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail struct {
	Desc    *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Suggest *string `json:"suggest,omitempty" xml:"suggest,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetDesc(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Desc = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetName(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Name = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetResult(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Result = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetSuggest(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Suggest = &v
	return s
}

func (s *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail) SetType(v string) *ListDiagnoseReportResponseBodyResultDiagnoseItemsDetail {
	s.Type = &v
	return s
}

type ListDiagnoseReportResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDiagnoseReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDiagnoseReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportResponse) SetHeaders(v map[string]*string) *ListDiagnoseReportResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnoseReportResponse) SetBody(v *ListDiagnoseReportResponseBody) *ListDiagnoseReportResponse {
	s.Body = v
	return s
}

type ListDiagnoseReportIdsRequest struct {
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Lang      *string `json:"lang,omitempty" xml:"lang,omitempty"`
	Page      *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size      *int32  `json:"size,omitempty" xml:"size,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Trigger   *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s ListDiagnoseReportIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportIdsRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportIdsRequest) SetEndTime(v int64) *ListDiagnoseReportIdsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetLang(v string) *ListDiagnoseReportIdsRequest {
	s.Lang = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetPage(v int32) *ListDiagnoseReportIdsRequest {
	s.Page = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetSize(v int32) *ListDiagnoseReportIdsRequest {
	s.Size = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetStartTime(v int64) *ListDiagnoseReportIdsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDiagnoseReportIdsRequest) SetTrigger(v string) *ListDiagnoseReportIdsRequest {
	s.Trigger = &v
	return s
}

type ListDiagnoseReportIdsResponseBody struct {
	Headers   *ListDiagnoseReportIdsResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string                                 `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDiagnoseReportIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportIdsResponseBody) SetHeaders(v *ListDiagnoseReportIdsResponseBodyHeaders) *ListDiagnoseReportIdsResponseBody {
	s.Headers = v
	return s
}

func (s *ListDiagnoseReportIdsResponseBody) SetRequestId(v string) *ListDiagnoseReportIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnoseReportIdsResponseBody) SetResult(v []*string) *ListDiagnoseReportIdsResponseBody {
	s.Result = v
	return s
}

type ListDiagnoseReportIdsResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListDiagnoseReportIdsResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportIdsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportIdsResponseBodyHeaders) SetXTotalCount(v int32) *ListDiagnoseReportIdsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListDiagnoseReportIdsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDiagnoseReportIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDiagnoseReportIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnoseReportIdsResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnoseReportIdsResponse) SetHeaders(v map[string]*string) *ListDiagnoseReportIdsResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnoseReportIdsResponse) SetBody(v *ListDiagnoseReportIdsResponseBody) *ListDiagnoseReportIdsResponse {
	s.Body = v
	return s
}

type ListDictInformationRequest struct {
	AnalyzerType *string `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	BucketName   *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	Key          *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListDictInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDictInformationRequest) GoString() string {
	return s.String()
}

func (s *ListDictInformationRequest) SetAnalyzerType(v string) *ListDictInformationRequest {
	s.AnalyzerType = &v
	return s
}

func (s *ListDictInformationRequest) SetBucketName(v string) *ListDictInformationRequest {
	s.BucketName = &v
	return s
}

func (s *ListDictInformationRequest) SetKey(v string) *ListDictInformationRequest {
	s.Key = &v
	return s
}

type ListDictInformationResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListDictInformationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListDictInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDictInformationResponseBody) GoString() string {
	return s.String()
}

func (s *ListDictInformationResponseBody) SetRequestId(v string) *ListDictInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDictInformationResponseBody) SetResult(v *ListDictInformationResponseBodyResult) *ListDictInformationResponseBody {
	s.Result = v
	return s
}

type ListDictInformationResponseBodyResult struct {
	FileSize  *int64                                          `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	OssObject *ListDictInformationResponseBodyResultOssObject `json:"ossObject,omitempty" xml:"ossObject,omitempty" type:"Struct"`
	Type      *string                                         `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDictInformationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDictInformationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDictInformationResponseBodyResult) SetFileSize(v int64) *ListDictInformationResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *ListDictInformationResponseBodyResult) SetOssObject(v *ListDictInformationResponseBodyResultOssObject) *ListDictInformationResponseBodyResult {
	s.OssObject = v
	return s
}

func (s *ListDictInformationResponseBodyResult) SetType(v string) *ListDictInformationResponseBodyResult {
	s.Type = &v
	return s
}

type ListDictInformationResponseBodyResultOssObject struct {
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	Etag       *string `json:"etag,omitempty" xml:"etag,omitempty"`
	Key        *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListDictInformationResponseBodyResultOssObject) String() string {
	return tea.Prettify(s)
}

func (s ListDictInformationResponseBodyResultOssObject) GoString() string {
	return s.String()
}

func (s *ListDictInformationResponseBodyResultOssObject) SetBucketName(v string) *ListDictInformationResponseBodyResultOssObject {
	s.BucketName = &v
	return s
}

func (s *ListDictInformationResponseBodyResultOssObject) SetEtag(v string) *ListDictInformationResponseBodyResultOssObject {
	s.Etag = &v
	return s
}

func (s *ListDictInformationResponseBodyResultOssObject) SetKey(v string) *ListDictInformationResponseBodyResultOssObject {
	s.Key = &v
	return s
}

type ListDictInformationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDictInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDictInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDictInformationResponse) GoString() string {
	return s.String()
}

func (s *ListDictInformationResponse) SetHeaders(v map[string]*string) *ListDictInformationResponse {
	s.Headers = v
	return s
}

func (s *ListDictInformationResponse) SetBody(v *ListDictInformationResponseBody) *ListDictInformationResponse {
	s.Body = v
	return s
}

type ListDictsRequest struct {
	AnalyzerType *string `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListDictsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDictsRequest) GoString() string {
	return s.String()
}

func (s *ListDictsRequest) SetAnalyzerType(v string) *ListDictsRequest {
	s.AnalyzerType = &v
	return s
}

func (s *ListDictsRequest) SetName(v string) *ListDictsRequest {
	s.Name = &v
	return s
}

type ListDictsResponseBody struct {
	Headers   *ListDictsResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDictsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDictsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDictsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDictsResponseBody) SetHeaders(v *ListDictsResponseBodyHeaders) *ListDictsResponseBody {
	s.Headers = v
	return s
}

func (s *ListDictsResponseBody) SetRequestId(v string) *ListDictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDictsResponseBody) SetResult(v []*ListDictsResponseBodyResult) *ListDictsResponseBody {
	s.Result = v
	return s
}

type ListDictsResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListDictsResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDictsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDictsResponseBodyHeaders) SetXTotalCount(v int32) *ListDictsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListDictsResponseBodyResult struct {
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	FileSize    *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType  *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDictsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDictsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDictsResponseBodyResult) SetDownloadUrl(v string) *ListDictsResponseBodyResult {
	s.DownloadUrl = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetFileSize(v int64) *ListDictsResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetName(v string) *ListDictsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetSourceType(v string) *ListDictsResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetType(v string) *ListDictsResponseBodyResult {
	s.Type = &v
	return s
}

type ListDictsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDictsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDictsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDictsResponse) GoString() string {
	return s.String()
}

func (s *ListDictsResponse) SetHeaders(v map[string]*string) *ListDictsResponse {
	s.Headers = v
	return s
}

func (s *ListDictsResponse) SetBody(v *ListDictsResponseBody) *ListDictsResponse {
	s.Body = v
	return s
}

type ListEcsInstancesRequest struct {
	EcsInstanceIds  *string `json:"ecsInstanceIds,omitempty" xml:"ecsInstanceIds,omitempty"`
	EcsInstanceName *string `json:"ecsInstanceName,omitempty" xml:"ecsInstanceName,omitempty"`
	Page            *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size            *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Tags            *string `json:"tags,omitempty" xml:"tags,omitempty"`
	VpcId           *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListEcsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesRequest) SetEcsInstanceIds(v string) *ListEcsInstancesRequest {
	s.EcsInstanceIds = &v
	return s
}

func (s *ListEcsInstancesRequest) SetEcsInstanceName(v string) *ListEcsInstancesRequest {
	s.EcsInstanceName = &v
	return s
}

func (s *ListEcsInstancesRequest) SetPage(v int32) *ListEcsInstancesRequest {
	s.Page = &v
	return s
}

func (s *ListEcsInstancesRequest) SetSize(v int32) *ListEcsInstancesRequest {
	s.Size = &v
	return s
}

func (s *ListEcsInstancesRequest) SetTags(v string) *ListEcsInstancesRequest {
	s.Tags = &v
	return s
}

func (s *ListEcsInstancesRequest) SetVpcId(v string) *ListEcsInstancesRequest {
	s.VpcId = &v
	return s
}

type ListEcsInstancesResponseBody struct {
	Headers   *ListEcsInstancesResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListEcsInstancesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListEcsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBody) SetHeaders(v *ListEcsInstancesResponseBodyHeaders) *ListEcsInstancesResponseBody {
	s.Headers = v
	return s
}

func (s *ListEcsInstancesResponseBody) SetRequestId(v string) *ListEcsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsInstancesResponseBody) SetResult(v []*ListEcsInstancesResponseBodyResult) *ListEcsInstancesResponseBody {
	s.Result = v
	return s
}

type ListEcsInstancesResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListEcsInstancesResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyHeaders) SetXTotalCount(v int32) *ListEcsInstancesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListEcsInstancesResponseBodyResult struct {
	CloudAssistantStatus *string                                         `json:"cloudAssistantStatus,omitempty" xml:"cloudAssistantStatus,omitempty"`
	Collectors           []*ListEcsInstancesResponseBodyResultCollectors `json:"collectors,omitempty" xml:"collectors,omitempty" type:"Repeated"`
	EcsInstanceId        *string                                         `json:"ecsInstanceId,omitempty" xml:"ecsInstanceId,omitempty"`
	EcsInstanceName      *string                                         `json:"ecsInstanceName,omitempty" xml:"ecsInstanceName,omitempty"`
	IpAddress            []*ListEcsInstancesResponseBodyResultIpAddress  `json:"ipAddress,omitempty" xml:"ipAddress,omitempty" type:"Repeated"`
	OsType               *string                                         `json:"osType,omitempty" xml:"osType,omitempty"`
	Status               *string                                         `json:"status,omitempty" xml:"status,omitempty"`
	Tags                 *string                                         `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListEcsInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResult) SetCloudAssistantStatus(v string) *ListEcsInstancesResponseBodyResult {
	s.CloudAssistantStatus = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetCollectors(v []*ListEcsInstancesResponseBodyResultCollectors) *ListEcsInstancesResponseBodyResult {
	s.Collectors = v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetEcsInstanceId(v string) *ListEcsInstancesResponseBodyResult {
	s.EcsInstanceId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetEcsInstanceName(v string) *ListEcsInstancesResponseBodyResult {
	s.EcsInstanceName = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetIpAddress(v []*ListEcsInstancesResponseBodyResultIpAddress) *ListEcsInstancesResponseBodyResult {
	s.IpAddress = v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetOsType(v string) *ListEcsInstancesResponseBodyResult {
	s.OsType = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetStatus(v string) *ListEcsInstancesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetTags(v string) *ListEcsInstancesResponseBodyResult {
	s.Tags = &v
	return s
}

type ListEcsInstancesResponseBodyResultCollectors struct {
	CollectorPaths []*string                                                    `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	Configs        []*ListEcsInstancesResponseBodyResultCollectorsConfigs       `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	DryRun         *bool                                                        `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	ExtendConfigs  []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	GmtCreatedTime *string                                                      `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	GmtUpdateTime  *string                                                      `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	Name           *string                                                      `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId        *string                                                      `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	ResId          *string                                                      `json:"resId,omitempty" xml:"resId,omitempty"`
	ResType        *string                                                      `json:"resType,omitempty" xml:"resType,omitempty"`
	ResVersion     *string                                                      `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	Status         *string                                                      `json:"status,omitempty" xml:"status,omitempty"`
	VpcId          *string                                                      `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultCollectors) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultCollectors) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetCollectorPaths(v []*string) *ListEcsInstancesResponseBodyResultCollectors {
	s.CollectorPaths = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetConfigs(v []*ListEcsInstancesResponseBodyResultCollectorsConfigs) *ListEcsInstancesResponseBodyResultCollectors {
	s.Configs = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetDryRun(v bool) *ListEcsInstancesResponseBodyResultCollectors {
	s.DryRun = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetExtendConfigs(v []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) *ListEcsInstancesResponseBodyResultCollectors {
	s.ExtendConfigs = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetGmtCreatedTime(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.GmtCreatedTime = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetGmtUpdateTime(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.GmtUpdateTime = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetName(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.Name = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetOwnerId(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.OwnerId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetResId(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.ResId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetResType(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.ResType = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetResVersion(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.ResVersion = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetStatus(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.Status = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetVpcId(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.VpcId = &v
	return s
}

type ListEcsInstancesResponseBodyResultCollectorsConfigs struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultCollectorsConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultCollectorsConfigs) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultCollectorsConfigs) SetContent(v string) *ListEcsInstancesResponseBodyResultCollectorsConfigs {
	s.Content = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsConfigs) SetFileName(v string) *ListEcsInstancesResponseBodyResultCollectorsConfigs {
	s.FileName = &v
	return s
}

type ListEcsInstancesResponseBodyResultCollectorsExtendConfigs struct {
	ConfigType       *string                                                              `json:"configType,omitempty" xml:"configType,omitempty"`
	EnableMonitoring *bool                                                                `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	GroupId          *string                                                              `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Hosts            []*string                                                            `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	InstanceId       *string                                                              `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType     *string                                                              `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	Machines         []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	Protocol         *string                                                              `json:"protocol,omitempty" xml:"protocol,omitempty"`
	Type             *string                                                              `json:"type,omitempty" xml:"type,omitempty"`
	UserName         *string                                                              `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetConfigType(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetEnableMonitoring(v bool) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetGroupId(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetHosts(v []*string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.Hosts = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetInstanceId(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetInstanceType(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetMachines(v []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.Machines = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetProtocol(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetType(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.Type = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetUserName(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.UserName = &v
	return s
}

type ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines struct {
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) SetAgentStatus(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) SetInstanceId(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

type ListEcsInstancesResponseBodyResultIpAddress struct {
	Host   *string `json:"host,omitempty" xml:"host,omitempty"`
	IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultIpAddress) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultIpAddress) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultIpAddress) SetHost(v string) *ListEcsInstancesResponseBodyResultIpAddress {
	s.Host = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultIpAddress) SetIpType(v string) *ListEcsInstancesResponseBodyResultIpAddress {
	s.IpType = &v
	return s
}

type ListEcsInstancesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEcsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEcsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEcsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponse) SetHeaders(v map[string]*string) *ListEcsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListEcsInstancesResponse) SetBody(v *ListEcsInstancesResponseBody) *ListEcsInstancesResponse {
	s.Body = v
	return s
}

type ListExtendfilesResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListExtendfilesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListExtendfilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExtendfilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExtendfilesResponseBody) SetRequestId(v string) *ListExtendfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExtendfilesResponseBody) SetResult(v []*ListExtendfilesResponseBodyResult) *ListExtendfilesResponseBody {
	s.Result = v
	return s
}

type ListExtendfilesResponseBodyResult struct {
	FilePath   *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListExtendfilesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListExtendfilesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListExtendfilesResponseBodyResult) SetFilePath(v string) *ListExtendfilesResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *ListExtendfilesResponseBodyResult) SetFileSize(v int64) *ListExtendfilesResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *ListExtendfilesResponseBodyResult) SetName(v string) *ListExtendfilesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListExtendfilesResponseBodyResult) SetSourceType(v string) *ListExtendfilesResponseBodyResult {
	s.SourceType = &v
	return s
}

type ListExtendfilesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExtendfilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExtendfilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExtendfilesResponse) GoString() string {
	return s.String()
}

func (s *ListExtendfilesResponse) SetHeaders(v map[string]*string) *ListExtendfilesResponse {
	s.Headers = v
	return s
}

func (s *ListExtendfilesResponse) SetBody(v *ListExtendfilesResponseBody) *ListExtendfilesResponse {
	s.Body = v
	return s
}

type ListILMPoliciesRequest struct {
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
}

func (s ListILMPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListILMPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListILMPoliciesRequest) SetPolicyName(v string) *ListILMPoliciesRequest {
	s.PolicyName = &v
	return s
}

type ListILMPoliciesResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListILMPoliciesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListILMPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListILMPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListILMPoliciesResponseBody) SetRequestId(v string) *ListILMPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListILMPoliciesResponseBody) SetResult(v []*ListILMPoliciesResponseBodyResult) *ListILMPoliciesResponseBody {
	s.Result = v
	return s
}

type ListILMPoliciesResponseBodyResult struct {
	Name   *string                `json:"name,omitempty" xml:"name,omitempty"`
	Phases map[string]interface{} `json:"phases,omitempty" xml:"phases,omitempty"`
}

func (s ListILMPoliciesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListILMPoliciesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListILMPoliciesResponseBodyResult) SetName(v string) *ListILMPoliciesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListILMPoliciesResponseBodyResult) SetPhases(v map[string]interface{}) *ListILMPoliciesResponseBodyResult {
	s.Phases = v
	return s
}

type ListILMPoliciesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListILMPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListILMPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListILMPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListILMPoliciesResponse) SetHeaders(v map[string]*string) *ListILMPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListILMPoliciesResponse) SetBody(v *ListILMPoliciesResponseBody) *ListILMPoliciesResponse {
	s.Body = v
	return s
}

type ListIndexTemplatesRequest struct {
	IndexTemplate *string `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
}

func (s ListIndexTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIndexTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesRequest) SetIndexTemplate(v string) *ListIndexTemplatesRequest {
	s.IndexTemplate = &v
	return s
}

type ListIndexTemplatesResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListIndexTemplatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListIndexTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIndexTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesResponseBody) SetRequestId(v string) *ListIndexTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexTemplatesResponseBody) SetResult(v []*ListIndexTemplatesResponseBodyResult) *ListIndexTemplatesResponseBody {
	s.Result = v
	return s
}

type ListIndexTemplatesResponseBodyResult struct {
	DataStream    *bool                                         `json:"dataStream,omitempty" xml:"dataStream,omitempty"`
	IlmPolicy     *string                                       `json:"ilmPolicy,omitempty" xml:"ilmPolicy,omitempty"`
	IndexPatterns []*string                                     `json:"indexPatterns,omitempty" xml:"indexPatterns,omitempty" type:"Repeated"`
	IndexTemplate *string                                       `json:"indexTemplate,omitempty" xml:"indexTemplate,omitempty"`
	Priority      *int32                                        `json:"priority,omitempty" xml:"priority,omitempty"`
	Template      *ListIndexTemplatesResponseBodyResultTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s ListIndexTemplatesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListIndexTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesResponseBodyResult) SetDataStream(v bool) *ListIndexTemplatesResponseBodyResult {
	s.DataStream = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetIlmPolicy(v string) *ListIndexTemplatesResponseBodyResult {
	s.IlmPolicy = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetIndexPatterns(v []*string) *ListIndexTemplatesResponseBodyResult {
	s.IndexPatterns = v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetIndexTemplate(v string) *ListIndexTemplatesResponseBodyResult {
	s.IndexTemplate = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetPriority(v int32) *ListIndexTemplatesResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResult) SetTemplate(v *ListIndexTemplatesResponseBodyResultTemplate) *ListIndexTemplatesResponseBodyResult {
	s.Template = v
	return s
}

type ListIndexTemplatesResponseBodyResultTemplate struct {
	Aliases  *string `json:"aliases,omitempty" xml:"aliases,omitempty"`
	Mappings *string `json:"mappings,omitempty" xml:"mappings,omitempty"`
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s ListIndexTemplatesResponseBodyResultTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListIndexTemplatesResponseBodyResultTemplate) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) SetAliases(v string) *ListIndexTemplatesResponseBodyResultTemplate {
	s.Aliases = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) SetMappings(v string) *ListIndexTemplatesResponseBodyResultTemplate {
	s.Mappings = &v
	return s
}

func (s *ListIndexTemplatesResponseBodyResultTemplate) SetSettings(v string) *ListIndexTemplatesResponseBodyResultTemplate {
	s.Settings = &v
	return s
}

type ListIndexTemplatesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIndexTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIndexTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIndexTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListIndexTemplatesResponse) SetHeaders(v map[string]*string) *ListIndexTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListIndexTemplatesResponse) SetBody(v *ListIndexTemplatesResponseBody) *ListIndexTemplatesResponse {
	s.Body = v
	return s
}

type ListInstanceRequest struct {
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	EsVersion        *string `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	InstanceCategory *string `json:"instanceCategory,omitempty" xml:"instanceCategory,omitempty"`
	InstanceId       *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Page             *int32  `json:"page,omitempty" xml:"page,omitempty"`
	PaymentType      *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	ResourceGroupId  *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Size             *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Tags             *string `json:"tags,omitempty" xml:"tags,omitempty"`
	VpcId            *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	ZoneId           *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) SetDescription(v string) *ListInstanceRequest {
	s.Description = &v
	return s
}

func (s *ListInstanceRequest) SetEsVersion(v string) *ListInstanceRequest {
	s.EsVersion = &v
	return s
}

func (s *ListInstanceRequest) SetInstanceCategory(v string) *ListInstanceRequest {
	s.InstanceCategory = &v
	return s
}

func (s *ListInstanceRequest) SetInstanceId(v string) *ListInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRequest) SetPage(v int32) *ListInstanceRequest {
	s.Page = &v
	return s
}

func (s *ListInstanceRequest) SetPaymentType(v string) *ListInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *ListInstanceRequest) SetResourceGroupId(v string) *ListInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstanceRequest) SetSize(v int32) *ListInstanceRequest {
	s.Size = &v
	return s
}

func (s *ListInstanceRequest) SetTags(v string) *ListInstanceRequest {
	s.Tags = &v
	return s
}

func (s *ListInstanceRequest) SetVpcId(v string) *ListInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *ListInstanceRequest) SetZoneId(v string) *ListInstanceRequest {
	s.ZoneId = &v
	return s
}

type ListInstanceResponseBody struct {
	Headers   *ListInstanceResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) SetHeaders(v *ListInstanceResponseBodyHeaders) *ListInstanceResponseBody {
	s.Headers = v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetResult(v []*ListInstanceResponseBodyResult) *ListInstanceResponseBody {
	s.Result = v
	return s
}

type ListInstanceResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListInstanceResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyHeaders) SetXTotalCount(v int32) *ListInstanceResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListInstanceResponseBodyResult struct {
	AdvancedDedicateMaster       *bool                                                       `json:"advancedDedicateMaster,omitempty" xml:"advancedDedicateMaster,omitempty"`
	ClientNodeConfiguration      *ListInstanceResponseBodyResultClientNodeConfiguration      `json:"clientNodeConfiguration,omitempty" xml:"clientNodeConfiguration,omitempty" type:"Struct"`
	CreatedAt                    *string                                                     `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	DedicateMaster               *bool                                                       `json:"dedicateMaster,omitempty" xml:"dedicateMaster,omitempty"`
	Description                  *string                                                     `json:"description,omitempty" xml:"description,omitempty"`
	ElasticDataNodeConfiguration *ListInstanceResponseBodyResultElasticDataNodeConfiguration `json:"elasticDataNodeConfiguration,omitempty" xml:"elasticDataNodeConfiguration,omitempty" type:"Struct"`
	EsVersion                    *string                                                     `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	ExtendConfigs                []map[string]interface{}                                    `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	InstanceId                   *string                                                     `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsNewDeployment              *string                                                     `json:"isNewDeployment,omitempty" xml:"isNewDeployment,omitempty"`
	KibanaConfiguration          *ListInstanceResponseBodyResultKibanaConfiguration          `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	MasterConfiguration          *ListInstanceResponseBodyResultMasterConfiguration          `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	NetworkConfig                *ListInstanceResponseBodyResultNetworkConfig                `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	NodeAmount                   *int32                                                      `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec                     *ListInstanceResponseBodyResultNodeSpec                     `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	PaymentType                  *string                                                     `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	PostpaidServiceStatus        *string                                                     `json:"postpaidServiceStatus,omitempty" xml:"postpaidServiceStatus,omitempty"`
	ResourceGroupId              *string                                                     `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ServiceVpc                   *bool                                                       `json:"serviceVpc,omitempty" xml:"serviceVpc,omitempty"`
	Status                       *string                                                     `json:"status,omitempty" xml:"status,omitempty"`
	Tags                         []*ListInstanceResponseBodyResultTags                       `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	UpdatedAt                    *string                                                     `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResult) SetAdvancedDedicateMaster(v bool) *ListInstanceResponseBodyResult {
	s.AdvancedDedicateMaster = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetClientNodeConfiguration(v *ListInstanceResponseBodyResultClientNodeConfiguration) *ListInstanceResponseBodyResult {
	s.ClientNodeConfiguration = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetCreatedAt(v string) *ListInstanceResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetDedicateMaster(v bool) *ListInstanceResponseBodyResult {
	s.DedicateMaster = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetDescription(v string) *ListInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetElasticDataNodeConfiguration(v *ListInstanceResponseBodyResultElasticDataNodeConfiguration) *ListInstanceResponseBodyResult {
	s.ElasticDataNodeConfiguration = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetEsVersion(v string) *ListInstanceResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetExtendConfigs(v []map[string]interface{}) *ListInstanceResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetInstanceId(v string) *ListInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetIsNewDeployment(v string) *ListInstanceResponseBodyResult {
	s.IsNewDeployment = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetKibanaConfiguration(v *ListInstanceResponseBodyResultKibanaConfiguration) *ListInstanceResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetMasterConfiguration(v *ListInstanceResponseBodyResultMasterConfiguration) *ListInstanceResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetNetworkConfig(v *ListInstanceResponseBodyResultNetworkConfig) *ListInstanceResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetNodeAmount(v int32) *ListInstanceResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetNodeSpec(v *ListInstanceResponseBodyResultNodeSpec) *ListInstanceResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetPaymentType(v string) *ListInstanceResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetPostpaidServiceStatus(v string) *ListInstanceResponseBodyResult {
	s.PostpaidServiceStatus = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetResourceGroupId(v string) *ListInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetServiceVpc(v bool) *ListInstanceResponseBodyResult {
	s.ServiceVpc = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetStatus(v string) *ListInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListInstanceResponseBodyResult) SetTags(v []*ListInstanceResponseBodyResultTags) *ListInstanceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListInstanceResponseBodyResult) SetUpdatedAt(v string) *ListInstanceResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

type ListInstanceResponseBodyResultClientNodeConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListInstanceResponseBodyResultClientNodeConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResultClientNodeConfiguration) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) SetAmount(v int32) *ListInstanceResponseBodyResultClientNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) SetDisk(v int32) *ListInstanceResponseBodyResultClientNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) SetDiskType(v string) *ListInstanceResponseBodyResultClientNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultClientNodeConfiguration) SetSpec(v string) *ListInstanceResponseBodyResultClientNodeConfiguration {
	s.Spec = &v
	return s
}

type ListInstanceResponseBodyResultElasticDataNodeConfiguration struct {
	Amount         *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk           *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskEncryption *bool   `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType       *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec           *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListInstanceResponseBodyResultElasticDataNodeConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResultElasticDataNodeConfiguration) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetAmount(v int32) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetDisk(v int32) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetDiskEncryption(v bool) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.DiskEncryption = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetDiskType(v string) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultElasticDataNodeConfiguration) SetSpec(v string) *ListInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Spec = &v
	return s
}

type ListInstanceResponseBodyResultKibanaConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListInstanceResponseBodyResultKibanaConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) SetAmount(v int32) *ListInstanceResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) SetDisk(v int32) *ListInstanceResponseBodyResultKibanaConfiguration {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) SetDiskType(v string) *ListInstanceResponseBodyResultKibanaConfiguration {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultKibanaConfiguration) SetSpec(v string) *ListInstanceResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

type ListInstanceResponseBodyResultMasterConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListInstanceResponseBodyResultMasterConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) SetAmount(v int32) *ListInstanceResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) SetDisk(v int32) *ListInstanceResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) SetDiskType(v string) *ListInstanceResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultMasterConfiguration) SetSpec(v string) *ListInstanceResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

type ListInstanceResponseBodyResultNetworkConfig struct {
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	VpcId     *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea    *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s ListInstanceResponseBodyResultNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultNetworkConfig) SetType(v string) *ListInstanceResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfig) SetVpcId(v string) *ListInstanceResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfig) SetVsArea(v string) *ListInstanceResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *ListInstanceResponseBodyResultNetworkConfig) SetVswitchId(v string) *ListInstanceResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

type ListInstanceResponseBodyResultNodeSpec struct {
	Disk           *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskEncryption *bool   `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType       *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec           *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListInstanceResponseBodyResultNodeSpec) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetDisk(v int32) *ListInstanceResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetDiskEncryption(v bool) *ListInstanceResponseBodyResultNodeSpec {
	s.DiskEncryption = &v
	return s
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetDiskType(v string) *ListInstanceResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *ListInstanceResponseBodyResultNodeSpec) SetSpec(v string) *ListInstanceResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

type ListInstanceResponseBodyResultTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListInstanceResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyResultTags) SetTagKey(v string) *ListInstanceResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *ListInstanceResponseBodyResultTags) SetTagValue(v string) *ListInstanceResponseBodyResultTags {
	s.TagValue = &v
	return s
}

type ListInstanceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResponse) SetHeaders(v map[string]*string) *ListInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResponse) SetBody(v *ListInstanceResponseBody) *ListInstanceResponse {
	s.Body = v
	return s
}

type ListInstanceIndicesRequest struct {
	All         *bool   `json:"all,omitempty" xml:"all,omitempty"`
	IsManaged   *bool   `json:"isManaged,omitempty" xml:"isManaged,omitempty"`
	IsOpenstore *bool   `json:"isOpenstore,omitempty" xml:"isOpenstore,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Page        *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size        *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListInstanceIndicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceIndicesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesRequest) SetAll(v bool) *ListInstanceIndicesRequest {
	s.All = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetIsManaged(v bool) *ListInstanceIndicesRequest {
	s.IsManaged = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetIsOpenstore(v bool) *ListInstanceIndicesRequest {
	s.IsOpenstore = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetName(v string) *ListInstanceIndicesRequest {
	s.Name = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetPage(v int32) *ListInstanceIndicesRequest {
	s.Page = &v
	return s
}

func (s *ListInstanceIndicesRequest) SetSize(v int32) *ListInstanceIndicesRequest {
	s.Size = &v
	return s
}

type ListInstanceIndicesResponseBody struct {
	Headers   *ListInstanceIndicesResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListInstanceIndicesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListInstanceIndicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesResponseBody) SetHeaders(v *ListInstanceIndicesResponseBodyHeaders) *ListInstanceIndicesResponseBody {
	s.Headers = v
	return s
}

func (s *ListInstanceIndicesResponseBody) SetRequestId(v string) *ListInstanceIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceIndicesResponseBody) SetResult(v []*ListInstanceIndicesResponseBodyResult) *ListInstanceIndicesResponseBody {
	s.Result = v
	return s
}

type ListInstanceIndicesResponseBodyHeaders struct {
	XManagedCount       *int32 `json:"X-Managed-Count,omitempty" xml:"X-Managed-Count,omitempty"`
	XManagedStorageSize *int64 `json:"X-Managed-StorageSize,omitempty" xml:"X-Managed-StorageSize,omitempty"`
	XOSSCount           *int32 `json:"X-OSS-Count,omitempty" xml:"X-OSS-Count,omitempty"`
	XOSSStorageSize     *int64 `json:"X-OSS-StorageSize,omitempty" xml:"X-OSS-StorageSize,omitempty"`
}

func (s ListInstanceIndicesResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceIndicesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesResponseBodyHeaders) SetXManagedCount(v int32) *ListInstanceIndicesResponseBodyHeaders {
	s.XManagedCount = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyHeaders) SetXManagedStorageSize(v int64) *ListInstanceIndicesResponseBodyHeaders {
	s.XManagedStorageSize = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyHeaders) SetXOSSCount(v int32) *ListInstanceIndicesResponseBodyHeaders {
	s.XOSSCount = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyHeaders) SetXOSSStorageSize(v int64) *ListInstanceIndicesResponseBodyHeaders {
	s.XOSSStorageSize = &v
	return s
}

type ListInstanceIndicesResponseBodyResult struct {
	CreateTime    *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Health        *string `json:"health,omitempty" xml:"health,omitempty"`
	IlmExplain    *string `json:"ilmExplain,omitempty" xml:"ilmExplain,omitempty"`
	IsManaged     *string `json:"isManaged,omitempty" xml:"isManaged,omitempty"`
	ManagedStatus *string `json:"managedStatus,omitempty" xml:"managedStatus,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Phase         *string `json:"phase,omitempty" xml:"phase,omitempty"`
	Size          *int64  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListInstanceIndicesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceIndicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesResponseBodyResult) SetCreateTime(v string) *ListInstanceIndicesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetHealth(v string) *ListInstanceIndicesResponseBodyResult {
	s.Health = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetIlmExplain(v string) *ListInstanceIndicesResponseBodyResult {
	s.IlmExplain = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetIsManaged(v string) *ListInstanceIndicesResponseBodyResult {
	s.IsManaged = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetManagedStatus(v string) *ListInstanceIndicesResponseBodyResult {
	s.ManagedStatus = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetName(v string) *ListInstanceIndicesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetPhase(v string) *ListInstanceIndicesResponseBodyResult {
	s.Phase = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetSize(v int64) *ListInstanceIndicesResponseBodyResult {
	s.Size = &v
	return s
}

type ListInstanceIndicesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceIndicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceIndicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceIndicesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesResponse) SetHeaders(v map[string]*string) *ListInstanceIndicesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceIndicesResponse) SetBody(v *ListInstanceIndicesResponseBody) *ListInstanceIndicesResponse {
	s.Body = v
	return s
}

type ListKibanaPluginsRequest struct {
	Page *string `json:"page,omitempty" xml:"page,omitempty"`
	Size *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListKibanaPluginsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKibanaPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsRequest) SetPage(v string) *ListKibanaPluginsRequest {
	s.Page = &v
	return s
}

func (s *ListKibanaPluginsRequest) SetSize(v int32) *ListKibanaPluginsRequest {
	s.Size = &v
	return s
}

type ListKibanaPluginsResponseBody struct {
	Headers   *ListKibanaPluginsResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListKibanaPluginsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListKibanaPluginsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKibanaPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsResponseBody) SetHeaders(v *ListKibanaPluginsResponseBodyHeaders) *ListKibanaPluginsResponseBody {
	s.Headers = v
	return s
}

func (s *ListKibanaPluginsResponseBody) SetRequestId(v string) *ListKibanaPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKibanaPluginsResponseBody) SetResult(v []*ListKibanaPluginsResponseBodyResult) *ListKibanaPluginsResponseBody {
	s.Result = v
	return s
}

type ListKibanaPluginsResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListKibanaPluginsResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListKibanaPluginsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsResponseBodyHeaders) SetXTotalCount(v int32) *ListKibanaPluginsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListKibanaPluginsResponseBodyResult struct {
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	Source           *string `json:"source,omitempty" xml:"source,omitempty"`
	SpecificationUrl *string `json:"specificationUrl,omitempty" xml:"specificationUrl,omitempty"`
	State            *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListKibanaPluginsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListKibanaPluginsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsResponseBodyResult) SetDescription(v string) *ListKibanaPluginsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyResult) SetName(v string) *ListKibanaPluginsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyResult) SetSource(v string) *ListKibanaPluginsResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyResult) SetSpecificationUrl(v string) *ListKibanaPluginsResponseBodyResult {
	s.SpecificationUrl = &v
	return s
}

func (s *ListKibanaPluginsResponseBodyResult) SetState(v string) *ListKibanaPluginsResponseBodyResult {
	s.State = &v
	return s
}

type ListKibanaPluginsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListKibanaPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListKibanaPluginsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKibanaPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsResponse) SetHeaders(v map[string]*string) *ListKibanaPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListKibanaPluginsResponse) SetBody(v *ListKibanaPluginsResponseBody) *ListKibanaPluginsResponse {
	s.Body = v
	return s
}

type ListLogstashRequest struct {
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OwnerId         *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	Page            *int32  `json:"page,omitempty" xml:"page,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Size            *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Version         *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListLogstashRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashRequest) GoString() string {
	return s.String()
}

func (s *ListLogstashRequest) SetDescription(v string) *ListLogstashRequest {
	s.Description = &v
	return s
}

func (s *ListLogstashRequest) SetInstanceId(v string) *ListLogstashRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLogstashRequest) SetOwnerId(v string) *ListLogstashRequest {
	s.OwnerId = &v
	return s
}

func (s *ListLogstashRequest) SetPage(v int32) *ListLogstashRequest {
	s.Page = &v
	return s
}

func (s *ListLogstashRequest) SetResourceGroupId(v string) *ListLogstashRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLogstashRequest) SetSize(v int32) *ListLogstashRequest {
	s.Size = &v
	return s
}

func (s *ListLogstashRequest) SetVersion(v string) *ListLogstashRequest {
	s.Version = &v
	return s
}

type ListLogstashResponseBody struct {
	Headers   *ListLogstashResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListLogstashResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListLogstashResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBody) SetHeaders(v *ListLogstashResponseBodyHeaders) *ListLogstashResponseBody {
	s.Headers = v
	return s
}

func (s *ListLogstashResponseBody) SetRequestId(v string) *ListLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogstashResponseBody) SetResult(v []*ListLogstashResponseBodyResult) *ListLogstashResponseBody {
	s.Result = v
	return s
}

type ListLogstashResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListLogstashResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyHeaders) SetXTotalCount(v int32) *ListLogstashResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListLogstashResponseBodyResult struct {
	Tags          []*ListLogstashResponseBodyResultTags        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	CreatedAt     *string                                      `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description   *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId    *string                                      `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	NetworkConfig *ListLogstashResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	NodeAmount    *int32                                       `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec      *ListLogstashResponseBodyResultNodeSpec      `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	PaymentType   *string                                      `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	Status        *string                                      `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt     *string                                      `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Version       *string                                      `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListLogstashResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyResult) SetTags(v []*ListLogstashResponseBodyResultTags) *ListLogstashResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListLogstashResponseBodyResult) SetCreatedAt(v string) *ListLogstashResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetDescription(v string) *ListLogstashResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetInstanceId(v string) *ListLogstashResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetNetworkConfig(v *ListLogstashResponseBodyResultNetworkConfig) *ListLogstashResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *ListLogstashResponseBodyResult) SetNodeAmount(v int32) *ListLogstashResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetNodeSpec(v *ListLogstashResponseBodyResultNodeSpec) *ListLogstashResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *ListLogstashResponseBodyResult) SetPaymentType(v string) *ListLogstashResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetStatus(v string) *ListLogstashResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetUpdatedAt(v string) *ListLogstashResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetVersion(v string) *ListLogstashResponseBodyResult {
	s.Version = &v
	return s
}

type ListLogstashResponseBodyResultTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListLogstashResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyResultTags) SetTagKey(v string) *ListLogstashResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *ListLogstashResponseBodyResultTags) SetTagValue(v string) *ListLogstashResponseBodyResultTags {
	s.TagValue = &v
	return s
}

type ListLogstashResponseBodyResultNetworkConfig struct {
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	VpcId     *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea    *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s ListLogstashResponseBodyResultNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyResultNetworkConfig) SetType(v string) *ListLogstashResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *ListLogstashResponseBodyResultNetworkConfig) SetVpcId(v string) *ListLogstashResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *ListLogstashResponseBodyResultNetworkConfig) SetVsArea(v string) *ListLogstashResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *ListLogstashResponseBodyResultNetworkConfig) SetVswitchId(v string) *ListLogstashResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

type ListLogstashResponseBodyResultNodeSpec struct {
	Disk           *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskEncryption *bool   `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	DiskType       *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec           *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListLogstashResponseBodyResultNodeSpec) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyResultNodeSpec) SetDisk(v int32) *ListLogstashResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *ListLogstashResponseBodyResultNodeSpec) SetDiskEncryption(v bool) *ListLogstashResponseBodyResultNodeSpec {
	s.DiskEncryption = &v
	return s
}

func (s *ListLogstashResponseBodyResultNodeSpec) SetDiskType(v string) *ListLogstashResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *ListLogstashResponseBodyResultNodeSpec) SetSpec(v string) *ListLogstashResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

type ListLogstashResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogstashResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashResponse) GoString() string {
	return s.String()
}

func (s *ListLogstashResponse) SetHeaders(v map[string]*string) *ListLogstashResponse {
	s.Headers = v
	return s
}

func (s *ListLogstashResponse) SetBody(v *ListLogstashResponseBody) *ListLogstashResponse {
	s.Body = v
	return s
}

type ListLogstashLogRequest struct {
	BeginTime *int64  `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Page      *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
	Size      *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListLogstashLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashLogRequest) GoString() string {
	return s.String()
}

func (s *ListLogstashLogRequest) SetBeginTime(v int64) *ListLogstashLogRequest {
	s.BeginTime = &v
	return s
}

func (s *ListLogstashLogRequest) SetEndTime(v int64) *ListLogstashLogRequest {
	s.EndTime = &v
	return s
}

func (s *ListLogstashLogRequest) SetPage(v int32) *ListLogstashLogRequest {
	s.Page = &v
	return s
}

func (s *ListLogstashLogRequest) SetQuery(v string) *ListLogstashLogRequest {
	s.Query = &v
	return s
}

func (s *ListLogstashLogRequest) SetSize(v int32) *ListLogstashLogRequest {
	s.Size = &v
	return s
}

func (s *ListLogstashLogRequest) SetType(v string) *ListLogstashLogRequest {
	s.Type = &v
	return s
}

type ListLogstashLogResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListLogstashLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListLogstashLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogstashLogResponseBody) SetRequestId(v string) *ListLogstashLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogstashLogResponseBody) SetResult(v []*ListLogstashLogResponseBodyResult) *ListLogstashLogResponseBody {
	s.Result = v
	return s
}

type ListLogstashLogResponseBodyResult struct {
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
	Host       *string `json:"host,omitempty" xml:"host,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Level      *string `json:"level,omitempty" xml:"level,omitempty"`
	Timestamp  *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ListLogstashLogResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLogstashLogResponseBodyResult) SetContent(v string) *ListLogstashLogResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListLogstashLogResponseBodyResult) SetHost(v string) *ListLogstashLogResponseBodyResult {
	s.Host = &v
	return s
}

func (s *ListLogstashLogResponseBodyResult) SetInstanceId(v string) *ListLogstashLogResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListLogstashLogResponseBodyResult) SetLevel(v string) *ListLogstashLogResponseBodyResult {
	s.Level = &v
	return s
}

func (s *ListLogstashLogResponseBodyResult) SetTimestamp(v int64) *ListLogstashLogResponseBodyResult {
	s.Timestamp = &v
	return s
}

type ListLogstashLogResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLogstashLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogstashLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashLogResponse) GoString() string {
	return s.String()
}

func (s *ListLogstashLogResponse) SetHeaders(v map[string]*string) *ListLogstashLogResponse {
	s.Headers = v
	return s
}

func (s *ListLogstashLogResponse) SetBody(v *ListLogstashLogResponseBody) *ListLogstashLogResponse {
	s.Body = v
	return s
}

type ListLogstashPluginsRequest struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Page   *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size   *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListLogstashPluginsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListLogstashPluginsRequest) SetName(v string) *ListLogstashPluginsRequest {
	s.Name = &v
	return s
}

func (s *ListLogstashPluginsRequest) SetPage(v int32) *ListLogstashPluginsRequest {
	s.Page = &v
	return s
}

func (s *ListLogstashPluginsRequest) SetSize(v int32) *ListLogstashPluginsRequest {
	s.Size = &v
	return s
}

func (s *ListLogstashPluginsRequest) SetSource(v string) *ListLogstashPluginsRequest {
	s.Source = &v
	return s
}

type ListLogstashPluginsResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListLogstashPluginsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListLogstashPluginsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogstashPluginsResponseBody) SetRequestId(v string) *ListLogstashPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogstashPluginsResponseBody) SetResult(v []*ListLogstashPluginsResponseBodyResult) *ListLogstashPluginsResponseBody {
	s.Result = v
	return s
}

type ListLogstashPluginsResponseBodyResult struct {
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	Source           *string `json:"source,omitempty" xml:"source,omitempty"`
	SpecificationUrl *string `json:"specificationUrl,omitempty" xml:"specificationUrl,omitempty"`
	State            *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListLogstashPluginsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashPluginsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLogstashPluginsResponseBodyResult) SetDescription(v string) *ListLogstashPluginsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListLogstashPluginsResponseBodyResult) SetName(v string) *ListLogstashPluginsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListLogstashPluginsResponseBodyResult) SetSource(v string) *ListLogstashPluginsResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListLogstashPluginsResponseBodyResult) SetSpecificationUrl(v string) *ListLogstashPluginsResponseBodyResult {
	s.SpecificationUrl = &v
	return s
}

func (s *ListLogstashPluginsResponseBodyResult) SetState(v string) *ListLogstashPluginsResponseBodyResult {
	s.State = &v
	return s
}

type ListLogstashPluginsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLogstashPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogstashPluginsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogstashPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListLogstashPluginsResponse) SetHeaders(v map[string]*string) *ListLogstashPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListLogstashPluginsResponse) SetBody(v *ListLogstashPluginsResponseBody) *ListLogstashPluginsResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	EcsInstanceIds  *string `json:"ecsInstanceIds,omitempty" xml:"ecsInstanceIds,omitempty"`
	EcsInstanceName *string `json:"ecsInstanceName,omitempty" xml:"ecsInstanceName,omitempty"`
	Page            *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Size            *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Tags            *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetEcsInstanceIds(v string) *ListNodesRequest {
	s.EcsInstanceIds = &v
	return s
}

func (s *ListNodesRequest) SetEcsInstanceName(v string) *ListNodesRequest {
	s.EcsInstanceName = &v
	return s
}

func (s *ListNodesRequest) SetPage(v int32) *ListNodesRequest {
	s.Page = &v
	return s
}

func (s *ListNodesRequest) SetSize(v int32) *ListNodesRequest {
	s.Size = &v
	return s
}

func (s *ListNodesRequest) SetTags(v string) *ListNodesRequest {
	s.Tags = &v
	return s
}

type ListNodesResponseBody struct {
	Headers   *ListNodesResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListNodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetHeaders(v *ListNodesResponseBodyHeaders) *ListNodesResponseBody {
	s.Headers = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetResult(v []*ListNodesResponseBodyResult) *ListNodesResponseBody {
	s.Result = v
	return s
}

type ListNodesResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListNodesResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyHeaders) SetXTotalCount(v int32) *ListNodesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListNodesResponseBodyResult struct {
	AgentStatus          *string                                 `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	CloudAssistantStatus *string                                 `json:"cloudAssistantStatus,omitempty" xml:"cloudAssistantStatus,omitempty"`
	EcsInstanceId        *string                                 `json:"ecsInstanceId,omitempty" xml:"ecsInstanceId,omitempty"`
	EcsInstanceName      *string                                 `json:"ecsInstanceName,omitempty" xml:"ecsInstanceName,omitempty"`
	IpAddress            []*ListNodesResponseBodyResultIpAddress `json:"ipAddress,omitempty" xml:"ipAddress,omitempty" type:"Repeated"`
	OsType               *string                                 `json:"osType,omitempty" xml:"osType,omitempty"`
	Status               *string                                 `json:"status,omitempty" xml:"status,omitempty"`
	Tags                 []*ListNodesResponseBodyResultTags      `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyResult) SetAgentStatus(v string) *ListNodesResponseBodyResult {
	s.AgentStatus = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetCloudAssistantStatus(v string) *ListNodesResponseBodyResult {
	s.CloudAssistantStatus = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetEcsInstanceId(v string) *ListNodesResponseBodyResult {
	s.EcsInstanceId = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetEcsInstanceName(v string) *ListNodesResponseBodyResult {
	s.EcsInstanceName = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetIpAddress(v []*ListNodesResponseBodyResultIpAddress) *ListNodesResponseBodyResult {
	s.IpAddress = v
	return s
}

func (s *ListNodesResponseBodyResult) SetOsType(v string) *ListNodesResponseBodyResult {
	s.OsType = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetStatus(v string) *ListNodesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListNodesResponseBodyResult) SetTags(v []*ListNodesResponseBodyResultTags) *ListNodesResponseBodyResult {
	s.Tags = v
	return s
}

type ListNodesResponseBodyResultIpAddress struct {
	Host   *string `json:"host,omitempty" xml:"host,omitempty"`
	IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty"`
}

func (s ListNodesResponseBodyResultIpAddress) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyResultIpAddress) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyResultIpAddress) SetHost(v string) *ListNodesResponseBodyResultIpAddress {
	s.Host = &v
	return s
}

func (s *ListNodesResponseBodyResultIpAddress) SetIpType(v string) *ListNodesResponseBodyResultIpAddress {
	s.IpType = &v
	return s
}

type ListNodesResponseBodyResultTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListNodesResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyResultTags) SetTagKey(v string) *ListNodesResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *ListNodesResponseBodyResultTags) SetTagValue(v string) *ListNodesResponseBodyResultTags {
	s.TagValue = &v
	return s
}

type ListNodesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListPipelineRequest struct {
	Page       *int32  `json:"page,omitempty" xml:"page,omitempty"`
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	Size       *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListPipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRequest) SetPage(v int32) *ListPipelineRequest {
	s.Page = &v
	return s
}

func (s *ListPipelineRequest) SetPipelineId(v string) *ListPipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineRequest) SetSize(v int32) *ListPipelineRequest {
	s.Size = &v
	return s
}

type ListPipelineResponseBody struct {
	Headers   *ListPipelineResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListPipelineResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListPipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineResponseBody) SetHeaders(v *ListPipelineResponseBodyHeaders) *ListPipelineResponseBody {
	s.Headers = v
	return s
}

func (s *ListPipelineResponseBody) SetRequestId(v string) *ListPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineResponseBody) SetResult(v []*ListPipelineResponseBodyResult) *ListPipelineResponseBody {
	s.Result = v
	return s
}

type ListPipelineResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListPipelineResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListPipelineResponseBodyHeaders) SetXTotalCount(v int32) *ListPipelineResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListPipelineResponseBodyResult struct {
	GmtCreatedTime *string `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	GmtUpdateTime  *string `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	PipelineId     *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PipelineStatus *string `json:"pipelineStatus,omitempty" xml:"pipelineStatus,omitempty"`
}

func (s ListPipelineResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPipelineResponseBodyResult) SetGmtCreatedTime(v string) *ListPipelineResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *ListPipelineResponseBodyResult) SetGmtUpdateTime(v string) *ListPipelineResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *ListPipelineResponseBodyResult) SetPipelineId(v string) *ListPipelineResponseBodyResult {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineResponseBodyResult) SetPipelineStatus(v string) *ListPipelineResponseBodyResult {
	s.PipelineStatus = &v
	return s
}

type ListPipelineResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineResponse) SetHeaders(v map[string]*string) *ListPipelineResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineResponse) SetBody(v *ListPipelineResponseBody) *ListPipelineResponse {
	s.Body = v
	return s
}

type ListPipelineIdsResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListPipelineIdsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListPipelineIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineIdsResponseBody) SetRequestId(v string) *ListPipelineIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineIdsResponseBody) SetResult(v []*ListPipelineIdsResponseBodyResult) *ListPipelineIdsResponseBody {
	s.Result = v
	return s
}

type ListPipelineIdsResponseBodyResult struct {
	Available  *bool   `json:"available,omitempty" xml:"available,omitempty"`
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	Message    *string `json:"message,omitempty" xml:"message,omitempty"`
	PipelineId *string `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
}

func (s ListPipelineIdsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPipelineIdsResponseBodyResult) SetAvailable(v bool) *ListPipelineIdsResponseBodyResult {
	s.Available = &v
	return s
}

func (s *ListPipelineIdsResponseBodyResult) SetCode(v string) *ListPipelineIdsResponseBodyResult {
	s.Code = &v
	return s
}

func (s *ListPipelineIdsResponseBodyResult) SetMessage(v string) *ListPipelineIdsResponseBodyResult {
	s.Message = &v
	return s
}

func (s *ListPipelineIdsResponseBodyResult) SetPipelineId(v string) *ListPipelineIdsResponseBodyResult {
	s.PipelineId = &v
	return s
}

type ListPipelineIdsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPipelineIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineIdsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineIdsResponse) SetHeaders(v map[string]*string) *ListPipelineIdsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineIdsResponse) SetBody(v *ListPipelineIdsResponseBody) *ListPipelineIdsResponse {
	s.Body = v
	return s
}

type ListPluginsRequest struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Page   *string `json:"page,omitempty" xml:"page,omitempty"`
	Size   *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListPluginsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListPluginsRequest) SetName(v string) *ListPluginsRequest {
	s.Name = &v
	return s
}

func (s *ListPluginsRequest) SetPage(v string) *ListPluginsRequest {
	s.Page = &v
	return s
}

func (s *ListPluginsRequest) SetSize(v int32) *ListPluginsRequest {
	s.Size = &v
	return s
}

func (s *ListPluginsRequest) SetSource(v string) *ListPluginsRequest {
	s.Source = &v
	return s
}

type ListPluginsResponseBody struct {
	Headers   *ListPluginsResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListPluginsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListPluginsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBody) SetHeaders(v *ListPluginsResponseBodyHeaders) *ListPluginsResponseBody {
	s.Headers = v
	return s
}

func (s *ListPluginsResponseBody) SetRequestId(v string) *ListPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginsResponseBody) SetResult(v []*ListPluginsResponseBodyResult) *ListPluginsResponseBody {
	s.Result = v
	return s
}

type ListPluginsResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListPluginsResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListPluginsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBodyHeaders) SetXTotalCount(v int32) *ListPluginsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListPluginsResponseBodyResult struct {
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	Source           *string `json:"source,omitempty" xml:"source,omitempty"`
	SpecificationUrl *string `json:"specificationUrl,omitempty" xml:"specificationUrl,omitempty"`
	State            *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListPluginsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListPluginsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPluginsResponseBodyResult) SetDescription(v string) *ListPluginsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListPluginsResponseBodyResult) SetName(v string) *ListPluginsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListPluginsResponseBodyResult) SetSource(v string) *ListPluginsResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListPluginsResponseBodyResult) SetSpecificationUrl(v string) *ListPluginsResponseBodyResult {
	s.SpecificationUrl = &v
	return s
}

func (s *ListPluginsResponseBodyResult) SetState(v string) *ListPluginsResponseBodyResult {
	s.State = &v
	return s
}

type ListPluginsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPluginsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListPluginsResponse) SetHeaders(v map[string]*string) *ListPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListPluginsResponse) SetBody(v *ListPluginsResponseBody) *ListPluginsResponse {
	s.Body = v
	return s
}

type ListSearchLogRequest struct {
	BeginTime *int64  `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Page      *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
	Size      *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSearchLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSearchLogRequest) GoString() string {
	return s.String()
}

func (s *ListSearchLogRequest) SetBeginTime(v int64) *ListSearchLogRequest {
	s.BeginTime = &v
	return s
}

func (s *ListSearchLogRequest) SetEndTime(v int64) *ListSearchLogRequest {
	s.EndTime = &v
	return s
}

func (s *ListSearchLogRequest) SetPage(v int32) *ListSearchLogRequest {
	s.Page = &v
	return s
}

func (s *ListSearchLogRequest) SetQuery(v string) *ListSearchLogRequest {
	s.Query = &v
	return s
}

func (s *ListSearchLogRequest) SetSize(v int32) *ListSearchLogRequest {
	s.Size = &v
	return s
}

func (s *ListSearchLogRequest) SetType(v string) *ListSearchLogRequest {
	s.Type = &v
	return s
}

type ListSearchLogResponseBody struct {
	Headers   *ListSearchLogResponseBodyHeaders  `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListSearchLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListSearchLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSearchLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchLogResponseBody) SetHeaders(v *ListSearchLogResponseBodyHeaders) *ListSearchLogResponseBody {
	s.Headers = v
	return s
}

func (s *ListSearchLogResponseBody) SetRequestId(v string) *ListSearchLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchLogResponseBody) SetResult(v []*ListSearchLogResponseBodyResult) *ListSearchLogResponseBody {
	s.Result = v
	return s
}

type ListSearchLogResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListSearchLogResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSearchLogResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListSearchLogResponseBodyHeaders) SetXTotalCount(v int32) *ListSearchLogResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListSearchLogResponseBodyResult struct {
	Content           *string                `json:"content,omitempty" xml:"content,omitempty"`
	ContentCollection map[string]interface{} `json:"contentCollection,omitempty" xml:"contentCollection,omitempty"`
	Host              *string                `json:"host,omitempty" xml:"host,omitempty"`
	InstanceId        *string                `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Level             *string                `json:"level,omitempty" xml:"level,omitempty"`
	Timestamp         *int64                 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s ListSearchLogResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSearchLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSearchLogResponseBodyResult) SetContent(v string) *ListSearchLogResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetContentCollection(v map[string]interface{}) *ListSearchLogResponseBodyResult {
	s.ContentCollection = v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetHost(v string) *ListSearchLogResponseBodyResult {
	s.Host = &v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetInstanceId(v string) *ListSearchLogResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetLevel(v string) *ListSearchLogResponseBodyResult {
	s.Level = &v
	return s
}

func (s *ListSearchLogResponseBodyResult) SetTimestamp(v int64) *ListSearchLogResponseBodyResult {
	s.Timestamp = &v
	return s
}

type ListSearchLogResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSearchLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSearchLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSearchLogResponse) GoString() string {
	return s.String()
}

func (s *ListSearchLogResponse) SetHeaders(v map[string]*string) *ListSearchLogResponse {
	s.Headers = v
	return s
}

func (s *ListSearchLogResponse) SetBody(v *ListSearchLogResponseBody) *ListSearchLogResponse {
	s.Body = v
	return s
}

type ListShardRecoveriesRequest struct {
	ActiveOnly *bool `json:"activeOnly,omitempty" xml:"activeOnly,omitempty"`
}

func (s ListShardRecoveriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListShardRecoveriesRequest) GoString() string {
	return s.String()
}

func (s *ListShardRecoveriesRequest) SetActiveOnly(v bool) *ListShardRecoveriesRequest {
	s.ActiveOnly = &v
	return s
}

type ListShardRecoveriesResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListShardRecoveriesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListShardRecoveriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListShardRecoveriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListShardRecoveriesResponseBody) SetRequestId(v string) *ListShardRecoveriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListShardRecoveriesResponseBody) SetResult(v []*ListShardRecoveriesResponseBodyResult) *ListShardRecoveriesResponseBody {
	s.Result = v
	return s
}

type ListShardRecoveriesResponseBodyResult struct {
	BytesPercent       *string `json:"bytesPercent,omitempty" xml:"bytesPercent,omitempty"`
	BytesTotal         *int64  `json:"bytesTotal,omitempty" xml:"bytesTotal,omitempty"`
	FilesPercent       *string `json:"filesPercent,omitempty" xml:"filesPercent,omitempty"`
	FilesTotal         *int64  `json:"filesTotal,omitempty" xml:"filesTotal,omitempty"`
	Index              *string `json:"index,omitempty" xml:"index,omitempty"`
	SourceHost         *string `json:"sourceHost,omitempty" xml:"sourceHost,omitempty"`
	SourceNode         *string `json:"sourceNode,omitempty" xml:"sourceNode,omitempty"`
	Stage              *string `json:"stage,omitempty" xml:"stage,omitempty"`
	TargetHost         *string `json:"targetHost,omitempty" xml:"targetHost,omitempty"`
	TargetNode         *string `json:"targetNode,omitempty" xml:"targetNode,omitempty"`
	TranslogOps        *int64  `json:"translogOps,omitempty" xml:"translogOps,omitempty"`
	TranslogOpsPercent *string `json:"translogOpsPercent,omitempty" xml:"translogOpsPercent,omitempty"`
}

func (s ListShardRecoveriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListShardRecoveriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListShardRecoveriesResponseBodyResult) SetBytesPercent(v string) *ListShardRecoveriesResponseBodyResult {
	s.BytesPercent = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetBytesTotal(v int64) *ListShardRecoveriesResponseBodyResult {
	s.BytesTotal = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetFilesPercent(v string) *ListShardRecoveriesResponseBodyResult {
	s.FilesPercent = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetFilesTotal(v int64) *ListShardRecoveriesResponseBodyResult {
	s.FilesTotal = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetIndex(v string) *ListShardRecoveriesResponseBodyResult {
	s.Index = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetSourceHost(v string) *ListShardRecoveriesResponseBodyResult {
	s.SourceHost = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetSourceNode(v string) *ListShardRecoveriesResponseBodyResult {
	s.SourceNode = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetStage(v string) *ListShardRecoveriesResponseBodyResult {
	s.Stage = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetTargetHost(v string) *ListShardRecoveriesResponseBodyResult {
	s.TargetHost = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetTargetNode(v string) *ListShardRecoveriesResponseBodyResult {
	s.TargetNode = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetTranslogOps(v int64) *ListShardRecoveriesResponseBodyResult {
	s.TranslogOps = &v
	return s
}

func (s *ListShardRecoveriesResponseBodyResult) SetTranslogOpsPercent(v string) *ListShardRecoveriesResponseBodyResult {
	s.TranslogOpsPercent = &v
	return s
}

type ListShardRecoveriesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListShardRecoveriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListShardRecoveriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListShardRecoveriesResponse) GoString() string {
	return s.String()
}

func (s *ListShardRecoveriesResponse) SetHeaders(v map[string]*string) *ListShardRecoveriesResponse {
	s.Headers = v
	return s
}

func (s *ListShardRecoveriesResponse) SetBody(v *ListShardRecoveriesResponseBody) *ListShardRecoveriesResponse {
	s.Body = v
	return s
}

type ListSnapshotReposByInstanceIdResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListSnapshotReposByInstanceIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListSnapshotReposByInstanceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotReposByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotReposByInstanceIdResponseBody) SetRequestId(v string) *ListSnapshotReposByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBody) SetResult(v []*ListSnapshotReposByInstanceIdResponseBodyResult) *ListSnapshotReposByInstanceIdResponseBody {
	s.Result = v
	return s
}

type ListSnapshotReposByInstanceIdResponseBodyResult struct {
	InstanceId    *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	RepoPath      *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
	SnapWarehouse *string `json:"snapWarehouse,omitempty" xml:"snapWarehouse,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListSnapshotReposByInstanceIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotReposByInstanceIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) SetInstanceId(v string) *ListSnapshotReposByInstanceIdResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) SetRepoPath(v string) *ListSnapshotReposByInstanceIdResponseBodyResult {
	s.RepoPath = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) SetSnapWarehouse(v string) *ListSnapshotReposByInstanceIdResponseBodyResult {
	s.SnapWarehouse = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) SetStatus(v string) *ListSnapshotReposByInstanceIdResponseBodyResult {
	s.Status = &v
	return s
}

type ListSnapshotReposByInstanceIdResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSnapshotReposByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSnapshotReposByInstanceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotReposByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotReposByInstanceIdResponse) SetHeaders(v map[string]*string) *ListSnapshotReposByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponse) SetBody(v *ListSnapshotReposByInstanceIdResponseBody) *ListSnapshotReposByInstanceIdResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Page         *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	ResourceIds  *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetPage(v int32) *ListTagResourcesRequest {
	s.Page = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceIds(v string) *ListTagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetSize(v int32) *ListTagResourcesRequest {
	s.Size = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v string) *ListTagResourcesRequest {
	s.Tags = &v
	return s
}

type ListTagResourcesResponseBody struct {
	Headers      *ListTagResourcesResponseBodyHeaders      `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	PageSize     *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetHeaders(v *ListTagResourcesResponseBodyHeaders) *ListTagResourcesResponseBody {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponseBody) SetPageSize(v int32) *ListTagResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyHeaders struct {
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListTagResourcesResponseBodyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyHeaders) SetXTotalCount(v int32) *ListTagResourcesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagsRequest struct {
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) SetPageSize(v int32) *ListTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagsRequest) SetResourceType(v string) *ListTagsRequest {
	s.ResourceType = &v
	return s
}

type ListTagsResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListTagsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetResult(v []*ListTagsResponseBodyResult) *ListTagsResponseBody {
	s.Result = v
	return s
}

type ListTagsResponseBodyResult struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyResult) SetTagKey(v string) *ListTagsResponseBodyResult {
	s.TagKey = &v
	return s
}

func (s *ListTagsResponseBodyResult) SetTagValue(v string) *ListTagsResponseBodyResult {
	s.TagValue = &v
	return s
}

type ListTagsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponse) GoString() string {
	return s.String()
}

func (s *ListTagsResponse) SetHeaders(v map[string]*string) *ListTagsResponse {
	s.Headers = v
	return s
}

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

type ListVpcEndpointsRequest struct {
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListVpcEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsRequest) SetPage(v int32) *ListVpcEndpointsRequest {
	s.Page = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetSize(v int32) *ListVpcEndpointsRequest {
	s.Size = &v
	return s
}

type ListVpcEndpointsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListVpcEndpointsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBody) SetRequestId(v string) *ListVpcEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetResult(v []*ListVpcEndpointsResponseBodyResult) *ListVpcEndpointsResponseBody {
	s.Result = v
	return s
}

type ListVpcEndpointsResponseBodyResult struct {
	ConnectionStatus       *string `json:"connectionStatus,omitempty" xml:"connectionStatus,omitempty"`
	CreateTime             *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	EndpointBusinessStatus *string `json:"endpointBusinessStatus,omitempty" xml:"endpointBusinessStatus,omitempty"`
	EndpointDomain         *string `json:"endpointDomain,omitempty" xml:"endpointDomain,omitempty"`
	EndpointId             *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	EndpointName           *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
	EndpointStatus         *string `json:"endpointStatus,omitempty" xml:"endpointStatus,omitempty"`
	ServiceId              *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	ServiceName            *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListVpcEndpointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBodyResult) SetConnectionStatus(v string) *ListVpcEndpointsResponseBodyResult {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetCreateTime(v string) *ListVpcEndpointsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointBusinessStatus(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointBusinessStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointDomain(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointDomain = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointId(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointName(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointName = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetEndpointStatus(v string) *ListVpcEndpointsResponseBodyResult {
	s.EndpointStatus = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetServiceId(v string) *ListVpcEndpointsResponseBodyResult {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyResult) SetServiceName(v string) *ListVpcEndpointsResponseBodyResult {
	s.ServiceName = &v
	return s
}

type ListVpcEndpointsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponse) SetHeaders(v map[string]*string) *ListVpcEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointsResponse) SetBody(v *ListVpcEndpointsResponseBody) *ListVpcEndpointsResponse {
	s.Body = v
	return s
}

type MigrateToOtherZoneRequest struct {
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s MigrateToOtherZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateToOtherZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneRequest) SetDryRun(v bool) *MigrateToOtherZoneRequest {
	s.DryRun = &v
	return s
}

type MigrateToOtherZoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s MigrateToOtherZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateToOtherZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneResponseBody) SetRequestId(v string) *MigrateToOtherZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateToOtherZoneResponseBody) SetResult(v bool) *MigrateToOtherZoneResponseBody {
	s.Result = &v
	return s
}

type MigrateToOtherZoneResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MigrateToOtherZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateToOtherZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateToOtherZoneResponse) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneResponse) SetHeaders(v map[string]*string) *MigrateToOtherZoneResponse {
	s.Headers = v
	return s
}

func (s *MigrateToOtherZoneResponse) SetBody(v *MigrateToOtherZoneResponseBody) *MigrateToOtherZoneResponse {
	s.Body = v
	return s
}

type ModifyDeployMachineRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyDeployMachineRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeployMachineRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeployMachineRequest) SetClientToken(v string) *ModifyDeployMachineRequest {
	s.ClientToken = &v
	return s
}

type ModifyDeployMachineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyDeployMachineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeployMachineResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeployMachineResponseBody) SetRequestId(v string) *ModifyDeployMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeployMachineResponseBody) SetResult(v bool) *ModifyDeployMachineResponseBody {
	s.Result = &v
	return s
}

type ModifyDeployMachineResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDeployMachineResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDeployMachineResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeployMachineResponse) SetHeaders(v map[string]*string) *ModifyDeployMachineResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeployMachineResponse) SetBody(v *ModifyDeployMachineResponseBody) *ModifyDeployMachineResponse {
	s.Body = v
	return s
}

type ModifyElastictaskResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ModifyElastictaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifyElastictaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyElastictaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponseBody) SetRequestId(v string) *ModifyElastictaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElastictaskResponseBody) SetResult(v *ModifyElastictaskResponseBodyResult) *ModifyElastictaskResponseBody {
	s.Result = v
	return s
}

type ModifyElastictaskResponseBodyResult struct {
	ElasticExpansionTask *ModifyElastictaskResponseBodyResultElasticExpansionTask `json:"elasticExpansionTask,omitempty" xml:"elasticExpansionTask,omitempty" type:"Struct"`
	ElasticShrinkTask    *ModifyElastictaskResponseBodyResultElasticShrinkTask    `json:"elasticShrinkTask,omitempty" xml:"elasticShrinkTask,omitempty" type:"Struct"`
}

func (s ModifyElastictaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyElastictaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponseBodyResult) SetElasticExpansionTask(v *ModifyElastictaskResponseBodyResultElasticExpansionTask) *ModifyElastictaskResponseBodyResult {
	s.ElasticExpansionTask = v
	return s
}

func (s *ModifyElastictaskResponseBodyResult) SetElasticShrinkTask(v *ModifyElastictaskResponseBodyResultElasticShrinkTask) *ModifyElastictaskResponseBodyResult {
	s.ElasticShrinkTask = v
	return s
}

type ModifyElastictaskResponseBodyResultElasticExpansionTask struct {
	CronExpression   *string   `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	ElasticNodeCount *int32    `json:"elasticNodeCount,omitempty" xml:"elasticNodeCount,omitempty"`
	ReplicaCount     *int32    `json:"replicaCount,omitempty" xml:"replicaCount,omitempty"`
	TargetIndices    []*string `json:"targetIndices,omitempty" xml:"targetIndices,omitempty" type:"Repeated"`
	TriggerType      *string   `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s ModifyElastictaskResponseBodyResultElasticExpansionTask) String() string {
	return tea.Prettify(s)
}

func (s ModifyElastictaskResponseBodyResultElasticExpansionTask) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetCronExpression(v string) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.CronExpression = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetElasticNodeCount(v int32) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.ElasticNodeCount = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetReplicaCount(v int32) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.ReplicaCount = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetTargetIndices(v []*string) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.TargetIndices = v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticExpansionTask) SetTriggerType(v string) *ModifyElastictaskResponseBodyResultElasticExpansionTask {
	s.TriggerType = &v
	return s
}

type ModifyElastictaskResponseBodyResultElasticShrinkTask struct {
	CronExpression   *string   `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	ElasticNodeCount *int32    `json:"elasticNodeCount,omitempty" xml:"elasticNodeCount,omitempty"`
	ReplicaCount     *int32    `json:"replicaCount,omitempty" xml:"replicaCount,omitempty"`
	TargetIndices    []*string `json:"targetIndices,omitempty" xml:"targetIndices,omitempty" type:"Repeated"`
	TriggerType      *string   `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s ModifyElastictaskResponseBodyResultElasticShrinkTask) String() string {
	return tea.Prettify(s)
}

func (s ModifyElastictaskResponseBodyResultElasticShrinkTask) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetCronExpression(v string) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.CronExpression = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetElasticNodeCount(v int32) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.ElasticNodeCount = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetReplicaCount(v int32) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.ReplicaCount = &v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetTargetIndices(v []*string) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.TargetIndices = v
	return s
}

func (s *ModifyElastictaskResponseBodyResultElasticShrinkTask) SetTriggerType(v string) *ModifyElastictaskResponseBodyResultElasticShrinkTask {
	s.TriggerType = &v
	return s
}

type ModifyElastictaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyElastictaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyElastictaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyElastictaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyElastictaskResponse) SetHeaders(v map[string]*string) *ModifyElastictaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyElastictaskResponse) SetBody(v *ModifyElastictaskResponseBody) *ModifyElastictaskResponse {
	s.Body = v
	return s
}

type ModifyInstanceMaintainTimeRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyInstanceMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeRequest) SetClientToken(v string) *ModifyInstanceMaintainTimeRequest {
	s.ClientToken = &v
	return s
}

type ModifyInstanceMaintainTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyInstanceMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeResponseBody) SetResult(v bool) *ModifyInstanceMaintainTimeResponseBody {
	s.Result = &v
	return s
}

type ModifyInstanceMaintainTimeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyInstanceMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMaintainTimeResponse) SetBody(v *ModifyInstanceMaintainTimeResponseBody) *ModifyInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyWhiteIpsRequest struct {
	ClientToken  *string                            `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ModifyMode   *string                            `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
	NetworkType  *string                            `json:"networkType,omitempty" xml:"networkType,omitempty"`
	NodeType     *string                            `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	WhiteIpGroup *ModifyWhiteIpsRequestWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Struct"`
	WhiteIpList  []*string                          `json:"whiteIpList,omitempty" xml:"whiteIpList,omitempty" type:"Repeated"`
}

func (s ModifyWhiteIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyWhiteIpsRequest) SetClientToken(v string) *ModifyWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyWhiteIpsRequest) SetModifyMode(v string) *ModifyWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyWhiteIpsRequest) SetNetworkType(v string) *ModifyWhiteIpsRequest {
	s.NetworkType = &v
	return s
}

func (s *ModifyWhiteIpsRequest) SetNodeType(v string) *ModifyWhiteIpsRequest {
	s.NodeType = &v
	return s
}

func (s *ModifyWhiteIpsRequest) SetWhiteIpGroup(v *ModifyWhiteIpsRequestWhiteIpGroup) *ModifyWhiteIpsRequest {
	s.WhiteIpGroup = v
	return s
}

func (s *ModifyWhiteIpsRequest) SetWhiteIpList(v []*string) *ModifyWhiteIpsRequest {
	s.WhiteIpList = v
	return s
}

type ModifyWhiteIpsRequestWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s ModifyWhiteIpsRequestWhiteIpGroup) String() string {
	return tea.Prettify(s)
}

func (s ModifyWhiteIpsRequestWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *ModifyWhiteIpsRequestWhiteIpGroup) SetGroupName(v string) *ModifyWhiteIpsRequestWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *ModifyWhiteIpsRequestWhiteIpGroup) SetIps(v []*string) *ModifyWhiteIpsRequestWhiteIpGroup {
	s.Ips = v
	return s
}

type ModifyWhiteIpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyWhiteIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWhiteIpsResponseBody) SetRequestId(v string) *ModifyWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWhiteIpsResponseBody) SetResult(v bool) *ModifyWhiteIpsResponseBody {
	s.Result = &v
	return s
}

type ModifyWhiteIpsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWhiteIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyWhiteIpsResponse) SetHeaders(v map[string]*string) *ModifyWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyWhiteIpsResponse) SetBody(v *ModifyWhiteIpsResponseBody) *ModifyWhiteIpsResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetClientToken(v string) *MoveResourceGroupRequest {
	s.ClientToken = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *MoveResourceGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetResult(v *MoveResourceGroupResponseBodyResult) *MoveResourceGroupResponseBody {
	s.Result = v
	return s
}

type MoveResourceGroupResponseBodyResult struct {
	CreatedAt           *string                                                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description         *string                                                 `json:"description,omitempty" xml:"description,omitempty"`
	DictList            []*MoveResourceGroupResponseBodyResultDictList          `json:"dictList,omitempty" xml:"dictList,omitempty" type:"Repeated"`
	Domain              *string                                                 `json:"domain,omitempty" xml:"domain,omitempty"`
	EsVersion           *string                                                 `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	InstanceId          *string                                                 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	KibanaConfiguration *MoveResourceGroupResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	KibanaDomain        *string                                                 `json:"kibanaDomain,omitempty" xml:"kibanaDomain,omitempty"`
	KibanaPort          *int32                                                  `json:"kibanaPort,omitempty" xml:"kibanaPort,omitempty"`
	MasterConfiguration *MoveResourceGroupResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	NetworkConfig       *MoveResourceGroupResponseBodyResultNetworkConfig       `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	NodeAmount          *int32                                                  `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec            *MoveResourceGroupResponseBodyResultNodeSpec            `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	PaymentType         *string                                                 `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	PublicDomain        *string                                                 `json:"publicDomain,omitempty" xml:"publicDomain,omitempty"`
	PublicPort          *int32                                                  `json:"publicPort,omitempty" xml:"publicPort,omitempty"`
	Status              *string                                                 `json:"status,omitempty" xml:"status,omitempty"`
	SynonymsDicts       []*MoveResourceGroupResponseBodyResultSynonymsDicts     `json:"synonymsDicts,omitempty" xml:"synonymsDicts,omitempty" type:"Repeated"`
	UpdatedAt           *string                                                 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s MoveResourceGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResult) SetCreatedAt(v string) *MoveResourceGroupResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetDescription(v string) *MoveResourceGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetDictList(v []*MoveResourceGroupResponseBodyResultDictList) *MoveResourceGroupResponseBodyResult {
	s.DictList = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetDomain(v string) *MoveResourceGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetEsVersion(v string) *MoveResourceGroupResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetInstanceId(v string) *MoveResourceGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetKibanaConfiguration(v *MoveResourceGroupResponseBodyResultKibanaConfiguration) *MoveResourceGroupResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetKibanaDomain(v string) *MoveResourceGroupResponseBodyResult {
	s.KibanaDomain = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetKibanaPort(v int32) *MoveResourceGroupResponseBodyResult {
	s.KibanaPort = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetMasterConfiguration(v *MoveResourceGroupResponseBodyResultMasterConfiguration) *MoveResourceGroupResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetNetworkConfig(v *MoveResourceGroupResponseBodyResultNetworkConfig) *MoveResourceGroupResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetNodeAmount(v int32) *MoveResourceGroupResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetNodeSpec(v *MoveResourceGroupResponseBodyResultNodeSpec) *MoveResourceGroupResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetPaymentType(v string) *MoveResourceGroupResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetPublicDomain(v string) *MoveResourceGroupResponseBodyResult {
	s.PublicDomain = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetPublicPort(v int32) *MoveResourceGroupResponseBodyResult {
	s.PublicPort = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetStatus(v string) *MoveResourceGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetSynonymsDicts(v []*MoveResourceGroupResponseBodyResultSynonymsDicts) *MoveResourceGroupResponseBodyResult {
	s.SynonymsDicts = v
	return s
}

func (s *MoveResourceGroupResponseBodyResult) SetUpdatedAt(v string) *MoveResourceGroupResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

type MoveResourceGroupResponseBodyResultDictList struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MoveResourceGroupResponseBodyResultDictList) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultDictList) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultDictList) SetFileSize(v int64) *MoveResourceGroupResponseBodyResultDictList {
	s.FileSize = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultDictList) SetName(v string) *MoveResourceGroupResponseBodyResultDictList {
	s.Name = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultDictList) SetSourceType(v string) *MoveResourceGroupResponseBodyResultDictList {
	s.SourceType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultDictList) SetType(v string) *MoveResourceGroupResponseBodyResultDictList {
	s.Type = &v
	return s
}

type MoveResourceGroupResponseBodyResultKibanaConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s MoveResourceGroupResponseBodyResultKibanaConfiguration) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) SetAmount(v int32) *MoveResourceGroupResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) SetDisk(v int32) *MoveResourceGroupResponseBodyResultKibanaConfiguration {
	s.Disk = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) SetDiskType(v string) *MoveResourceGroupResponseBodyResultKibanaConfiguration {
	s.DiskType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultKibanaConfiguration) SetSpec(v string) *MoveResourceGroupResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

type MoveResourceGroupResponseBodyResultMasterConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s MoveResourceGroupResponseBodyResultMasterConfiguration) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) SetAmount(v int32) *MoveResourceGroupResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) SetDisk(v int32) *MoveResourceGroupResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) SetDiskType(v string) *MoveResourceGroupResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultMasterConfiguration) SetSpec(v string) *MoveResourceGroupResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

type MoveResourceGroupResponseBodyResultNetworkConfig struct {
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	VpcId     *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea    *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s MoveResourceGroupResponseBodyResultNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) SetType(v string) *MoveResourceGroupResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) SetVpcId(v string) *MoveResourceGroupResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) SetVsArea(v string) *MoveResourceGroupResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNetworkConfig) SetVswitchId(v string) *MoveResourceGroupResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

type MoveResourceGroupResponseBodyResultNodeSpec struct {
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s MoveResourceGroupResponseBodyResultNodeSpec) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) SetDisk(v int32) *MoveResourceGroupResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) SetDiskType(v string) *MoveResourceGroupResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultNodeSpec) SetSpec(v string) *MoveResourceGroupResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

type MoveResourceGroupResponseBodyResultSynonymsDicts struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MoveResourceGroupResponseBodyResultSynonymsDicts) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBodyResultSynonymsDicts) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) SetFileSize(v int64) *MoveResourceGroupResponseBodyResultSynonymsDicts {
	s.FileSize = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) SetName(v string) *MoveResourceGroupResponseBodyResultSynonymsDicts {
	s.Name = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) SetSourceType(v string) *MoveResourceGroupResponseBodyResultSynonymsDicts {
	s.SourceType = &v
	return s
}

func (s *MoveResourceGroupResponseBodyResultSynonymsDicts) SetType(v string) *MoveResourceGroupResponseBodyResultSynonymsDicts {
	s.Type = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type OpenDiagnosisRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang        *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s OpenDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *OpenDiagnosisRequest) SetClientToken(v string) *OpenDiagnosisRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenDiagnosisRequest) SetLang(v string) *OpenDiagnosisRequest {
	s.Lang = &v
	return s
}

type OpenDiagnosisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s OpenDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDiagnosisResponseBody) SetRequestId(v string) *OpenDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenDiagnosisResponseBody) SetResult(v bool) *OpenDiagnosisResponseBody {
	s.Result = &v
	return s
}

type OpenDiagnosisResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *OpenDiagnosisResponse) SetHeaders(v map[string]*string) *OpenDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *OpenDiagnosisResponse) SetBody(v *OpenDiagnosisResponseBody) *OpenDiagnosisResponse {
	s.Body = v
	return s
}

type OpenHttpsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s OpenHttpsRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenHttpsRequest) GoString() string {
	return s.String()
}

func (s *OpenHttpsRequest) SetClientToken(v string) *OpenHttpsRequest {
	s.ClientToken = &v
	return s
}

type OpenHttpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s OpenHttpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenHttpsResponseBody) GoString() string {
	return s.String()
}

func (s *OpenHttpsResponseBody) SetRequestId(v string) *OpenHttpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenHttpsResponseBody) SetResult(v bool) *OpenHttpsResponseBody {
	s.Result = &v
	return s
}

type OpenHttpsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenHttpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenHttpsResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenHttpsResponse) GoString() string {
	return s.String()
}

func (s *OpenHttpsResponse) SetHeaders(v map[string]*string) *OpenHttpsResponse {
	s.Headers = v
	return s
}

func (s *OpenHttpsResponse) SetBody(v *OpenHttpsResponseBody) *OpenHttpsResponse {
	s.Body = v
	return s
}

type PostEmonTryAlarmRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostEmonTryAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostEmonTryAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PostEmonTryAlarmRuleResponseBody) SetCode(v string) *PostEmonTryAlarmRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PostEmonTryAlarmRuleResponseBody) SetMessage(v string) *PostEmonTryAlarmRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PostEmonTryAlarmRuleResponseBody) SetRequestId(v string) *PostEmonTryAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostEmonTryAlarmRuleResponseBody) SetSuccess(v bool) *PostEmonTryAlarmRuleResponseBody {
	s.Success = &v
	return s
}

type PostEmonTryAlarmRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PostEmonTryAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PostEmonTryAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PostEmonTryAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *PostEmonTryAlarmRuleResponse) SetHeaders(v map[string]*string) *PostEmonTryAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *PostEmonTryAlarmRuleResponse) SetBody(v *PostEmonTryAlarmRuleResponseBody) *PostEmonTryAlarmRuleResponse {
	s.Body = v
	return s
}

type RecommendTemplatesRequest struct {
	UsageScenario *string `json:"usageScenario,omitempty" xml:"usageScenario,omitempty"`
}

func (s RecommendTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s RecommendTemplatesRequest) GoString() string {
	return s.String()
}

func (s *RecommendTemplatesRequest) SetUsageScenario(v string) *RecommendTemplatesRequest {
	s.UsageScenario = &v
	return s
}

type RecommendTemplatesResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*RecommendTemplatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s RecommendTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecommendTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *RecommendTemplatesResponseBody) SetRequestId(v string) *RecommendTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecommendTemplatesResponseBody) SetResult(v []*RecommendTemplatesResponseBodyResult) *RecommendTemplatesResponseBody {
	s.Result = v
	return s
}

type RecommendTemplatesResponseBodyResult struct {
	Content      *string `json:"content,omitempty" xml:"content,omitempty"`
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s RecommendTemplatesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RecommendTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RecommendTemplatesResponseBodyResult) SetContent(v string) *RecommendTemplatesResponseBodyResult {
	s.Content = &v
	return s
}

func (s *RecommendTemplatesResponseBodyResult) SetTemplateName(v string) *RecommendTemplatesResponseBodyResult {
	s.TemplateName = &v
	return s
}

type RecommendTemplatesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecommendTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecommendTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s RecommendTemplatesResponse) GoString() string {
	return s.String()
}

func (s *RecommendTemplatesResponse) SetHeaders(v map[string]*string) *RecommendTemplatesResponse {
	s.Headers = v
	return s
}

func (s *RecommendTemplatesResponse) SetBody(v *RecommendTemplatesResponseBody) *RecommendTemplatesResponse {
	s.Body = v
	return s
}

type ReinstallCollectorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ReinstallCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s ReinstallCollectorRequest) GoString() string {
	return s.String()
}

func (s *ReinstallCollectorRequest) SetClientToken(v string) *ReinstallCollectorRequest {
	s.ClientToken = &v
	return s
}

type ReinstallCollectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ReinstallCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReinstallCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *ReinstallCollectorResponseBody) SetRequestId(v string) *ReinstallCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReinstallCollectorResponseBody) SetResult(v bool) *ReinstallCollectorResponseBody {
	s.Result = &v
	return s
}

type ReinstallCollectorResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReinstallCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReinstallCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s ReinstallCollectorResponse) GoString() string {
	return s.String()
}

func (s *ReinstallCollectorResponse) SetHeaders(v map[string]*string) *ReinstallCollectorResponse {
	s.Headers = v
	return s
}

func (s *ReinstallCollectorResponse) SetBody(v *ReinstallCollectorResponseBody) *ReinstallCollectorResponse {
	s.Body = v
	return s
}

type RemoveApmResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RemoveApmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveApmResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApmResponseBody) SetRequestId(v string) *RemoveApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveApmResponseBody) SetResult(v bool) *RemoveApmResponseBody {
	s.Result = &v
	return s
}

type RemoveApmResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveApmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveApmResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveApmResponse) GoString() string {
	return s.String()
}

func (s *RemoveApmResponse) SetHeaders(v map[string]*string) *RemoveApmResponse {
	s.Headers = v
	return s
}

func (s *RemoveApmResponse) SetBody(v *RemoveApmResponseBody) *RemoveApmResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

type RenewInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetResult(v bool) *RenewInstanceResponseBody {
	s.Result = &v
	return s
}

type RenewInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type RenewLogstashRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s RenewLogstashRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewLogstashRequest) GoString() string {
	return s.String()
}

func (s *RenewLogstashRequest) SetClientToken(v string) *RenewLogstashRequest {
	s.ClientToken = &v
	return s
}

type RenewLogstashResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RenewLogstashResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *RenewLogstashResponseBody) SetRequestId(v string) *RenewLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewLogstashResponseBody) SetResult(v bool) *RenewLogstashResponseBody {
	s.Result = &v
	return s
}

type RenewLogstashResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewLogstashResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewLogstashResponse) GoString() string {
	return s.String()
}

func (s *RenewLogstashResponse) SetHeaders(v map[string]*string) *RenewLogstashResponse {
	s.Headers = v
	return s
}

func (s *RenewLogstashResponse) SetBody(v *RenewLogstashResponseBody) *RenewLogstashResponse {
	s.Body = v
	return s
}

type RestartCollectorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RestartCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartCollectorRequest) GoString() string {
	return s.String()
}

func (s *RestartCollectorRequest) SetClientToken(v string) *RestartCollectorRequest {
	s.ClientToken = &v
	return s
}

type RestartCollectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RestartCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *RestartCollectorResponseBody) SetRequestId(v string) *RestartCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartCollectorResponseBody) SetResult(v bool) *RestartCollectorResponseBody {
	s.Result = &v
	return s
}

type RestartCollectorResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartCollectorResponse) GoString() string {
	return s.String()
}

func (s *RestartCollectorResponse) SetHeaders(v map[string]*string) *RestartCollectorResponse {
	s.Headers = v
	return s
}

func (s *RestartCollectorResponse) SetBody(v *RestartCollectorResponseBody) *RestartCollectorResponse {
	s.Body = v
	return s
}

type RestartInstanceRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Force       *bool   `json:"force,omitempty" xml:"force,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) SetClientToken(v string) *RestartInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartInstanceRequest) SetForce(v bool) *RestartInstanceRequest {
	s.Force = &v
	return s
}

type RestartInstanceResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *RestartInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s RestartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetResult(v *RestartInstanceResponseBodyResult) *RestartInstanceResponseBody {
	s.Result = v
	return s
}

type RestartInstanceResponseBodyResult struct {
	CreatedAt           *string                                               `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description         *string                                               `json:"description,omitempty" xml:"description,omitempty"`
	DictList            []*RestartInstanceResponseBodyResultDictList          `json:"dictList,omitempty" xml:"dictList,omitempty" type:"Repeated"`
	Domain              *string                                               `json:"domain,omitempty" xml:"domain,omitempty"`
	EsVersion           *string                                               `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	InstanceId          *string                                               `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	KibanaConfiguration *RestartInstanceResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	KibanaDomain        *string                                               `json:"kibanaDomain,omitempty" xml:"kibanaDomain,omitempty"`
	KibanaPort          *int32                                                `json:"kibanaPort,omitempty" xml:"kibanaPort,omitempty"`
	MasterConfiguration *RestartInstanceResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	NetworkConfig       *RestartInstanceResponseBodyResultNetworkConfig       `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	NodeAmount          *int32                                                `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec            *RestartInstanceResponseBodyResultNodeSpec            `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	PaymentType         *string                                               `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	PublicDomain        *string                                               `json:"publicDomain,omitempty" xml:"publicDomain,omitempty"`
	PublicPort          *int32                                                `json:"publicPort,omitempty" xml:"publicPort,omitempty"`
	Status              *string                                               `json:"status,omitempty" xml:"status,omitempty"`
	SynonymsDicts       []*RestartInstanceResponseBodyResultSynonymsDicts     `json:"synonymsDicts,omitempty" xml:"synonymsDicts,omitempty" type:"Repeated"`
	UpdatedAt           *string                                               `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s RestartInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResult) SetCreatedAt(v string) *RestartInstanceResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetDescription(v string) *RestartInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetDictList(v []*RestartInstanceResponseBodyResultDictList) *RestartInstanceResponseBodyResult {
	s.DictList = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetDomain(v string) *RestartInstanceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetEsVersion(v string) *RestartInstanceResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetInstanceId(v string) *RestartInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetKibanaConfiguration(v *RestartInstanceResponseBodyResultKibanaConfiguration) *RestartInstanceResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetKibanaDomain(v string) *RestartInstanceResponseBodyResult {
	s.KibanaDomain = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetKibanaPort(v int32) *RestartInstanceResponseBodyResult {
	s.KibanaPort = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetMasterConfiguration(v *RestartInstanceResponseBodyResultMasterConfiguration) *RestartInstanceResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetNetworkConfig(v *RestartInstanceResponseBodyResultNetworkConfig) *RestartInstanceResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetNodeAmount(v int32) *RestartInstanceResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetNodeSpec(v *RestartInstanceResponseBodyResultNodeSpec) *RestartInstanceResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetPaymentType(v string) *RestartInstanceResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetPublicDomain(v string) *RestartInstanceResponseBodyResult {
	s.PublicDomain = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetPublicPort(v int32) *RestartInstanceResponseBodyResult {
	s.PublicPort = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetStatus(v string) *RestartInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetSynonymsDicts(v []*RestartInstanceResponseBodyResultSynonymsDicts) *RestartInstanceResponseBodyResult {
	s.SynonymsDicts = v
	return s
}

func (s *RestartInstanceResponseBodyResult) SetUpdatedAt(v string) *RestartInstanceResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

type RestartInstanceResponseBodyResultDictList struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RestartInstanceResponseBodyResultDictList) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBodyResultDictList) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultDictList) SetFileSize(v int64) *RestartInstanceResponseBodyResultDictList {
	s.FileSize = &v
	return s
}

func (s *RestartInstanceResponseBodyResultDictList) SetName(v string) *RestartInstanceResponseBodyResultDictList {
	s.Name = &v
	return s
}

func (s *RestartInstanceResponseBodyResultDictList) SetSourceType(v string) *RestartInstanceResponseBodyResultDictList {
	s.SourceType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultDictList) SetType(v string) *RestartInstanceResponseBodyResultDictList {
	s.Type = &v
	return s
}

type RestartInstanceResponseBodyResultKibanaConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s RestartInstanceResponseBodyResultKibanaConfiguration) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) SetAmount(v int32) *RestartInstanceResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) SetDisk(v int32) *RestartInstanceResponseBodyResultKibanaConfiguration {
	s.Disk = &v
	return s
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) SetDiskType(v string) *RestartInstanceResponseBodyResultKibanaConfiguration {
	s.DiskType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultKibanaConfiguration) SetSpec(v string) *RestartInstanceResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

type RestartInstanceResponseBodyResultMasterConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s RestartInstanceResponseBodyResultMasterConfiguration) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) SetAmount(v int32) *RestartInstanceResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) SetDisk(v int32) *RestartInstanceResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) SetDiskType(v string) *RestartInstanceResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultMasterConfiguration) SetSpec(v string) *RestartInstanceResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

type RestartInstanceResponseBodyResultNetworkConfig struct {
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	VpcId     *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea    *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s RestartInstanceResponseBodyResultNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) SetType(v string) *RestartInstanceResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) SetVpcId(v string) *RestartInstanceResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) SetVsArea(v string) *RestartInstanceResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNetworkConfig) SetVswitchId(v string) *RestartInstanceResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

type RestartInstanceResponseBodyResultNodeSpec struct {
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s RestartInstanceResponseBodyResultNodeSpec) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultNodeSpec) SetDisk(v int32) *RestartInstanceResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNodeSpec) SetDiskType(v string) *RestartInstanceResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultNodeSpec) SetSpec(v string) *RestartInstanceResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

type RestartInstanceResponseBodyResultSynonymsDicts struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RestartInstanceResponseBodyResultSynonymsDicts) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBodyResultSynonymsDicts) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) SetFileSize(v int64) *RestartInstanceResponseBodyResultSynonymsDicts {
	s.FileSize = &v
	return s
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) SetName(v string) *RestartInstanceResponseBodyResultSynonymsDicts {
	s.Name = &v
	return s
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) SetSourceType(v string) *RestartInstanceResponseBodyResultSynonymsDicts {
	s.SourceType = &v
	return s
}

func (s *RestartInstanceResponseBodyResultSynonymsDicts) SetType(v string) *RestartInstanceResponseBodyResultSynonymsDicts {
	s.Type = &v
	return s
}

type RestartInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) SetHeaders(v map[string]*string) *RestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

type RestartLogstashRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Force       *bool   `json:"force,omitempty" xml:"force,omitempty"`
}

func (s RestartLogstashRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartLogstashRequest) GoString() string {
	return s.String()
}

func (s *RestartLogstashRequest) SetClientToken(v string) *RestartLogstashRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartLogstashRequest) SetForce(v bool) *RestartLogstashRequest {
	s.Force = &v
	return s
}

type RestartLogstashResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartLogstashResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *RestartLogstashResponseBody) SetRequestId(v string) *RestartLogstashResponseBody {
	s.RequestId = &v
	return s
}

type RestartLogstashResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartLogstashResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartLogstashResponse) GoString() string {
	return s.String()
}

func (s *RestartLogstashResponse) SetHeaders(v map[string]*string) *RestartLogstashResponse {
	s.Headers = v
	return s
}

func (s *RestartLogstashResponse) SetBody(v *RestartLogstashResponseBody) *RestartLogstashResponse {
	s.Body = v
	return s
}

type ResumeElasticsearchTaskRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ResumeElasticsearchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeElasticsearchTaskRequest) GoString() string {
	return s.String()
}

func (s *ResumeElasticsearchTaskRequest) SetClientToken(v string) *ResumeElasticsearchTaskRequest {
	s.ClientToken = &v
	return s
}

type ResumeElasticsearchTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ResumeElasticsearchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeElasticsearchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeElasticsearchTaskResponseBody) SetRequestId(v string) *ResumeElasticsearchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeElasticsearchTaskResponseBody) SetResult(v bool) *ResumeElasticsearchTaskResponseBody {
	s.Result = &v
	return s
}

type ResumeElasticsearchTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeElasticsearchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeElasticsearchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeElasticsearchTaskResponse) GoString() string {
	return s.String()
}

func (s *ResumeElasticsearchTaskResponse) SetHeaders(v map[string]*string) *ResumeElasticsearchTaskResponse {
	s.Headers = v
	return s
}

func (s *ResumeElasticsearchTaskResponse) SetBody(v *ResumeElasticsearchTaskResponseBody) *ResumeElasticsearchTaskResponse {
	s.Body = v
	return s
}

type ResumeLogstashTaskRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ResumeLogstashTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeLogstashTaskRequest) GoString() string {
	return s.String()
}

func (s *ResumeLogstashTaskRequest) SetClientToken(v string) *ResumeLogstashTaskRequest {
	s.ClientToken = &v
	return s
}

type ResumeLogstashTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ResumeLogstashTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeLogstashTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeLogstashTaskResponseBody) SetCode(v string) *ResumeLogstashTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeLogstashTaskResponseBody) SetMessage(v string) *ResumeLogstashTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeLogstashTaskResponseBody) SetRequestId(v string) *ResumeLogstashTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeLogstashTaskResponseBody) SetResult(v bool) *ResumeLogstashTaskResponseBody {
	s.Result = &v
	return s
}

type ResumeLogstashTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeLogstashTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeLogstashTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeLogstashTaskResponse) GoString() string {
	return s.String()
}

func (s *ResumeLogstashTaskResponse) SetHeaders(v map[string]*string) *ResumeLogstashTaskResponse {
	s.Headers = v
	return s
}

func (s *ResumeLogstashTaskResponse) SetBody(v *ResumeLogstashTaskResponseBody) *ResumeLogstashTaskResponse {
	s.Body = v
	return s
}

type RolloverDataStreamRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RolloverDataStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s RolloverDataStreamRequest) GoString() string {
	return s.String()
}

func (s *RolloverDataStreamRequest) SetClientToken(v string) *RolloverDataStreamRequest {
	s.ClientToken = &v
	return s
}

type RolloverDataStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RolloverDataStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RolloverDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *RolloverDataStreamResponseBody) SetRequestId(v string) *RolloverDataStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *RolloverDataStreamResponseBody) SetResult(v bool) *RolloverDataStreamResponseBody {
	s.Result = &v
	return s
}

type RolloverDataStreamResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RolloverDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RolloverDataStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s RolloverDataStreamResponse) GoString() string {
	return s.String()
}

func (s *RolloverDataStreamResponse) SetHeaders(v map[string]*string) *RolloverDataStreamResponse {
	s.Headers = v
	return s
}

func (s *RolloverDataStreamResponse) SetBody(v *RolloverDataStreamResponseBody) *RolloverDataStreamResponse {
	s.Body = v
	return s
}

type RunPipelinesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RunPipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s RunPipelinesRequest) GoString() string {
	return s.String()
}

func (s *RunPipelinesRequest) SetClientToken(v string) *RunPipelinesRequest {
	s.ClientToken = &v
	return s
}

type RunPipelinesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RunPipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *RunPipelinesResponseBody) SetRequestId(v string) *RunPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPipelinesResponseBody) SetResult(v bool) *RunPipelinesResponseBody {
	s.Result = &v
	return s
}

type RunPipelinesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunPipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s RunPipelinesResponse) GoString() string {
	return s.String()
}

func (s *RunPipelinesResponse) SetHeaders(v map[string]*string) *RunPipelinesResponse {
	s.Headers = v
	return s
}

func (s *RunPipelinesResponse) SetBody(v *RunPipelinesResponseBody) *RunPipelinesResponse {
	s.Body = v
	return s
}

type ShrinkNodeRequest struct {
	ClientToken  *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IgnoreStatus *bool   `json:"ignoreStatus,omitempty" xml:"ignoreStatus,omitempty"`
	NodeType     *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s ShrinkNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ShrinkNodeRequest) GoString() string {
	return s.String()
}

func (s *ShrinkNodeRequest) SetClientToken(v string) *ShrinkNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *ShrinkNodeRequest) SetIgnoreStatus(v bool) *ShrinkNodeRequest {
	s.IgnoreStatus = &v
	return s
}

func (s *ShrinkNodeRequest) SetNodeType(v string) *ShrinkNodeRequest {
	s.NodeType = &v
	return s
}

type ShrinkNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ShrinkNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ShrinkNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ShrinkNodeResponseBody) SetRequestId(v string) *ShrinkNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShrinkNodeResponseBody) SetResult(v bool) *ShrinkNodeResponseBody {
	s.Result = &v
	return s
}

type ShrinkNodeResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ShrinkNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ShrinkNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ShrinkNodeResponse) GoString() string {
	return s.String()
}

func (s *ShrinkNodeResponse) SetHeaders(v map[string]*string) *ShrinkNodeResponse {
	s.Headers = v
	return s
}

func (s *ShrinkNodeResponse) SetBody(v *ShrinkNodeResponseBody) *ShrinkNodeResponse {
	s.Body = v
	return s
}

type StartApmResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StartApmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartApmResponseBody) GoString() string {
	return s.String()
}

func (s *StartApmResponseBody) SetRequestId(v string) *StartApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartApmResponseBody) SetResult(v bool) *StartApmResponseBody {
	s.Result = &v
	return s
}

type StartApmResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartApmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartApmResponse) String() string {
	return tea.Prettify(s)
}

func (s StartApmResponse) GoString() string {
	return s.String()
}

func (s *StartApmResponse) SetHeaders(v map[string]*string) *StartApmResponse {
	s.Headers = v
	return s
}

func (s *StartApmResponse) SetBody(v *StartApmResponseBody) *StartApmResponse {
	s.Body = v
	return s
}

type StartCollectorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s StartCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCollectorRequest) GoString() string {
	return s.String()
}

func (s *StartCollectorRequest) SetClientToken(v string) *StartCollectorRequest {
	s.ClientToken = &v
	return s
}

type StartCollectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StartCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *StartCollectorResponseBody) SetRequestId(v string) *StartCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCollectorResponseBody) SetResult(v bool) *StartCollectorResponseBody {
	s.Result = &v
	return s
}

type StartCollectorResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCollectorResponse) GoString() string {
	return s.String()
}

func (s *StartCollectorResponse) SetHeaders(v map[string]*string) *StartCollectorResponse {
	s.Headers = v
	return s
}

func (s *StartCollectorResponse) SetBody(v *StartCollectorResponseBody) *StartCollectorResponse {
	s.Body = v
	return s
}

type StopApmResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StopApmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopApmResponseBody) GoString() string {
	return s.String()
}

func (s *StopApmResponseBody) SetRequestId(v string) *StopApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopApmResponseBody) SetResult(v bool) *StopApmResponseBody {
	s.Result = &v
	return s
}

type StopApmResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopApmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopApmResponse) String() string {
	return tea.Prettify(s)
}

func (s StopApmResponse) GoString() string {
	return s.String()
}

func (s *StopApmResponse) SetHeaders(v map[string]*string) *StopApmResponse {
	s.Headers = v
	return s
}

func (s *StopApmResponse) SetBody(v *StopApmResponseBody) *StopApmResponse {
	s.Body = v
	return s
}

type StopCollectorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s StopCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCollectorRequest) GoString() string {
	return s.String()
}

func (s *StopCollectorRequest) SetClientToken(v string) *StopCollectorRequest {
	s.ClientToken = &v
	return s
}

type StopCollectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StopCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *StopCollectorResponseBody) SetRequestId(v string) *StopCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCollectorResponseBody) SetResult(v bool) *StopCollectorResponseBody {
	s.Result = &v
	return s
}

type StopCollectorResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCollectorResponse) GoString() string {
	return s.String()
}

func (s *StopCollectorResponse) SetHeaders(v map[string]*string) *StopCollectorResponse {
	s.Headers = v
	return s
}

func (s *StopCollectorResponse) SetBody(v *StopCollectorResponseBody) *StopCollectorResponse {
	s.Body = v
	return s
}

type StopPipelinesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s StopPipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopPipelinesRequest) GoString() string {
	return s.String()
}

func (s *StopPipelinesRequest) SetClientToken(v string) *StopPipelinesRequest {
	s.ClientToken = &v
	return s
}

type StopPipelinesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StopPipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *StopPipelinesResponseBody) SetRequestId(v string) *StopPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPipelinesResponseBody) SetResult(v bool) *StopPipelinesResponseBody {
	s.Result = &v
	return s
}

type StopPipelinesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopPipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopPipelinesResponse) GoString() string {
	return s.String()
}

func (s *StopPipelinesResponse) SetHeaders(v map[string]*string) *StopPipelinesResponse {
	s.Headers = v
	return s
}

func (s *StopPipelinesResponse) SetBody(v *StopPipelinesResponseBody) *StopPipelinesResponse {
	s.Body = v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type TransferNodeRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	NodeType    *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s TransferNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferNodeRequest) GoString() string {
	return s.String()
}

func (s *TransferNodeRequest) SetClientToken(v string) *TransferNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferNodeRequest) SetNodeType(v string) *TransferNodeRequest {
	s.NodeType = &v
	return s
}

type TransferNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s TransferNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferNodeResponseBody) GoString() string {
	return s.String()
}

func (s *TransferNodeResponseBody) SetRequestId(v string) *TransferNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferNodeResponseBody) SetResult(v bool) *TransferNodeResponseBody {
	s.Result = &v
	return s
}

type TransferNodeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferNodeResponse) GoString() string {
	return s.String()
}

func (s *TransferNodeResponse) SetHeaders(v map[string]*string) *TransferNodeResponse {
	s.Headers = v
	return s
}

func (s *TransferNodeResponse) SetBody(v *TransferNodeResponseBody) *TransferNodeResponse {
	s.Body = v
	return s
}

type TriggerNetworkRequest struct {
	ActionType  *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	NodeType    *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s TriggerNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerNetworkRequest) GoString() string {
	return s.String()
}

func (s *TriggerNetworkRequest) SetActionType(v string) *TriggerNetworkRequest {
	s.ActionType = &v
	return s
}

func (s *TriggerNetworkRequest) SetClientToken(v string) *TriggerNetworkRequest {
	s.ClientToken = &v
	return s
}

func (s *TriggerNetworkRequest) SetNetworkType(v string) *TriggerNetworkRequest {
	s.NetworkType = &v
	return s
}

func (s *TriggerNetworkRequest) SetNodeType(v string) *TriggerNetworkRequest {
	s.NodeType = &v
	return s
}

type TriggerNetworkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s TriggerNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerNetworkResponseBody) SetRequestId(v string) *TriggerNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerNetworkResponseBody) SetResult(v bool) *TriggerNetworkResponseBody {
	s.Result = &v
	return s
}

type TriggerNetworkResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TriggerNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TriggerNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerNetworkResponse) GoString() string {
	return s.String()
}

func (s *TriggerNetworkResponse) SetHeaders(v map[string]*string) *TriggerNetworkResponse {
	s.Headers = v
	return s
}

func (s *TriggerNetworkResponse) SetBody(v *TriggerNetworkResponseBody) *TriggerNetworkResponse {
	s.Body = v
	return s
}

type UninstallKibanaPluginRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UninstallKibanaPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallKibanaPluginRequest) GoString() string {
	return s.String()
}

func (s *UninstallKibanaPluginRequest) SetClientToken(v string) *UninstallKibanaPluginRequest {
	s.ClientToken = &v
	return s
}

type UninstallKibanaPluginResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UninstallKibanaPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallKibanaPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallKibanaPluginResponseBody) SetRequestId(v string) *UninstallKibanaPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallKibanaPluginResponseBody) SetResult(v []*string) *UninstallKibanaPluginResponseBody {
	s.Result = v
	return s
}

type UninstallKibanaPluginResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UninstallKibanaPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallKibanaPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallKibanaPluginResponse) GoString() string {
	return s.String()
}

func (s *UninstallKibanaPluginResponse) SetHeaders(v map[string]*string) *UninstallKibanaPluginResponse {
	s.Headers = v
	return s
}

func (s *UninstallKibanaPluginResponse) SetBody(v *UninstallKibanaPluginResponseBody) *UninstallKibanaPluginResponse {
	s.Body = v
	return s
}

type UninstallLogstashPluginRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UninstallLogstashPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallLogstashPluginRequest) GoString() string {
	return s.String()
}

func (s *UninstallLogstashPluginRequest) SetClientToken(v string) *UninstallLogstashPluginRequest {
	s.ClientToken = &v
	return s
}

type UninstallLogstashPluginResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UninstallLogstashPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallLogstashPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallLogstashPluginResponseBody) SetRequestId(v string) *UninstallLogstashPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallLogstashPluginResponseBody) SetResult(v []*string) *UninstallLogstashPluginResponseBody {
	s.Result = v
	return s
}

type UninstallLogstashPluginResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UninstallLogstashPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallLogstashPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallLogstashPluginResponse) GoString() string {
	return s.String()
}

func (s *UninstallLogstashPluginResponse) SetHeaders(v map[string]*string) *UninstallLogstashPluginResponse {
	s.Headers = v
	return s
}

func (s *UninstallLogstashPluginResponse) SetBody(v *UninstallLogstashPluginResponseBody) *UninstallLogstashPluginResponse {
	s.Body = v
	return s
}

type UninstallPluginRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UninstallPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallPluginRequest) GoString() string {
	return s.String()
}

func (s *UninstallPluginRequest) SetClientToken(v string) *UninstallPluginRequest {
	s.ClientToken = &v
	return s
}

type UninstallPluginResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UninstallPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallPluginResponseBody) SetRequestId(v string) *UninstallPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallPluginResponseBody) SetResult(v []*string) *UninstallPluginResponseBody {
	s.Result = v
	return s
}

type UninstallPluginResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UninstallPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallPluginResponse) GoString() string {
	return s.String()
}

func (s *UninstallPluginResponse) SetHeaders(v map[string]*string) *UninstallPluginResponse {
	s.Headers = v
	return s
}

func (s *UninstallPluginResponse) SetBody(v *UninstallPluginResponseBody) *UninstallPluginResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	ResourceIds  *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKeys      *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v string) *UntagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v string) *UntagResourcesRequest {
	s.TagKeys = &v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetResult(v bool) *UntagResourcesResponseBody {
	s.Result = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAdminPasswordRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAdminPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdminPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdminPasswordRequest) SetClientToken(v string) *UpdateAdminPasswordRequest {
	s.ClientToken = &v
	return s
}

type UpdateAdminPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAdminPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdminPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdminPasswordResponseBody) SetRequestId(v string) *UpdateAdminPasswordResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAdminPasswordResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAdminPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAdminPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdminPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdminPasswordResponse) SetHeaders(v map[string]*string) *UpdateAdminPasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdminPasswordResponse) SetBody(v *UpdateAdminPasswordResponseBody) *UpdateAdminPasswordResponse {
	s.Body = v
	return s
}

type UpdateAdvancedSettingRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAdvancedSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdvancedSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdvancedSettingRequest) SetClientToken(v string) *UpdateAdvancedSettingRequest {
	s.ClientToken = &v
	return s
}

type UpdateAdvancedSettingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateAdvancedSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdvancedSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdvancedSettingResponseBody) SetRequestId(v string) *UpdateAdvancedSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdvancedSettingResponseBody) SetResult(v bool) *UpdateAdvancedSettingResponseBody {
	s.Result = &v
	return s
}

type UpdateAdvancedSettingResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAdvancedSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAdvancedSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdvancedSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdvancedSettingResponse) SetHeaders(v map[string]*string) *UpdateAdvancedSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdvancedSettingResponse) SetBody(v *UpdateAdvancedSettingResponseBody) *UpdateAdvancedSettingResponse {
	s.Body = v
	return s
}

type UpdateAliwsDictRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAliwsDictRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliwsDictRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliwsDictRequest) SetClientToken(v string) *UpdateAliwsDictRequest {
	s.ClientToken = &v
	return s
}

type UpdateAliwsDictResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*UpdateAliwsDictResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateAliwsDictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliwsDictResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAliwsDictResponseBody) SetRequestId(v string) *UpdateAliwsDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAliwsDictResponseBody) SetResult(v []*UpdateAliwsDictResponseBodyResult) *UpdateAliwsDictResponseBody {
	s.Result = v
	return s
}

type UpdateAliwsDictResponseBodyResult struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAliwsDictResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliwsDictResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateAliwsDictResponseBodyResult) SetFileSize(v int64) *UpdateAliwsDictResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateAliwsDictResponseBodyResult) SetName(v string) *UpdateAliwsDictResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateAliwsDictResponseBodyResult) SetSourceType(v string) *UpdateAliwsDictResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateAliwsDictResponseBodyResult) SetType(v string) *UpdateAliwsDictResponseBodyResult {
	s.Type = &v
	return s
}

type UpdateAliwsDictResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAliwsDictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAliwsDictResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliwsDictResponse) GoString() string {
	return s.String()
}

func (s *UpdateAliwsDictResponse) SetHeaders(v map[string]*string) *UpdateAliwsDictResponse {
	s.Headers = v
	return s
}

func (s *UpdateAliwsDictResponse) SetBody(v *UpdateAliwsDictResponseBody) *UpdateAliwsDictResponse {
	s.Body = v
	return s
}

type UpdateApmRequest struct {
	// apm实例名
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// es实例id
	OutputES *string `json:"outputES,omitempty" xml:"outputES,omitempty"`
	// es实例密码
	OutputESPassword *string `json:"outputESPassword,omitempty" xml:"outputESPassword,omitempty"`
	// es实例用户名
	OutputESUserName *string `json:"outputESUserName,omitempty" xml:"outputESUserName,omitempty"`
	// apm server密码
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s UpdateApmRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApmRequest) GoString() string {
	return s.String()
}

func (s *UpdateApmRequest) SetDescription(v string) *UpdateApmRequest {
	s.Description = &v
	return s
}

func (s *UpdateApmRequest) SetOutputES(v string) *UpdateApmRequest {
	s.OutputES = &v
	return s
}

func (s *UpdateApmRequest) SetOutputESPassword(v string) *UpdateApmRequest {
	s.OutputESPassword = &v
	return s
}

func (s *UpdateApmRequest) SetOutputESUserName(v string) *UpdateApmRequest {
	s.OutputESUserName = &v
	return s
}

func (s *UpdateApmRequest) SetToken(v string) *UpdateApmRequest {
	s.Token = &v
	return s
}

type UpdateApmResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateApmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApmResponseBody) SetRequestId(v string) *UpdateApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApmResponseBody) SetResult(v bool) *UpdateApmResponseBody {
	s.Result = &v
	return s
}

type UpdateApmResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApmResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApmResponse) GoString() string {
	return s.String()
}

func (s *UpdateApmResponse) SetHeaders(v map[string]*string) *UpdateApmResponse {
	s.Headers = v
	return s
}

func (s *UpdateApmResponse) SetBody(v *UpdateApmResponseBody) *UpdateApmResponse {
	s.Body = v
	return s
}

type UpdateBlackIpsRequest struct {
	ClientToken   *string   `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	EsIPBlacklist []*string `json:"esIPBlacklist,omitempty" xml:"esIPBlacklist,omitempty" type:"Repeated"`
}

func (s UpdateBlackIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBlackIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdateBlackIpsRequest) SetClientToken(v string) *UpdateBlackIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBlackIpsRequest) SetEsIPBlacklist(v []*string) *UpdateBlackIpsRequest {
	s.EsIPBlacklist = v
	return s
}

type UpdateBlackIpsResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateBlackIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateBlackIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBlackIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBlackIpsResponseBody) SetRequestId(v string) *UpdateBlackIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBlackIpsResponseBody) SetResult(v *UpdateBlackIpsResponseBodyResult) *UpdateBlackIpsResponseBody {
	s.Result = v
	return s
}

type UpdateBlackIpsResponseBodyResult struct {
	EsIPBlacklist []*string `json:"esIPBlacklist,omitempty" xml:"esIPBlacklist,omitempty" type:"Repeated"`
}

func (s UpdateBlackIpsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateBlackIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateBlackIpsResponseBodyResult) SetEsIPBlacklist(v []*string) *UpdateBlackIpsResponseBodyResult {
	s.EsIPBlacklist = v
	return s
}

type UpdateBlackIpsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateBlackIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBlackIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBlackIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdateBlackIpsResponse) SetHeaders(v map[string]*string) *UpdateBlackIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdateBlackIpsResponse) SetBody(v *UpdateBlackIpsResponseBody) *UpdateBlackIpsResponse {
	s.Body = v
	return s
}

type UpdateCollectorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorRequest) GoString() string {
	return s.String()
}

func (s *UpdateCollectorRequest) SetClientToken(v string) *UpdateCollectorRequest {
	s.ClientToken = &v
	return s
}

type UpdateCollectorResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateCollectorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBody) SetRequestId(v string) *UpdateCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCollectorResponseBody) SetResult(v *UpdateCollectorResponseBodyResult) *UpdateCollectorResponseBody {
	s.Result = v
	return s
}

type UpdateCollectorResponseBodyResult struct {
	CollectorPaths []*string                                         `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	Configs        []*UpdateCollectorResponseBodyResultConfigs       `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	DryRun         *bool                                             `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	ExtendConfigs  []*UpdateCollectorResponseBodyResultExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	GmtCreatedTime *string                                           `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	GmtUpdateTime  *string                                           `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	Name           *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId        *string                                           `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	ResId          *string                                           `json:"resId,omitempty" xml:"resId,omitempty"`
	ResType        *string                                           `json:"resType,omitempty" xml:"resType,omitempty"`
	ResVersion     *string                                           `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	Status         *string                                           `json:"status,omitempty" xml:"status,omitempty"`
	VpcId          *string                                           `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s UpdateCollectorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBodyResult) SetCollectorPaths(v []*string) *UpdateCollectorResponseBodyResult {
	s.CollectorPaths = v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetConfigs(v []*UpdateCollectorResponseBodyResultConfigs) *UpdateCollectorResponseBodyResult {
	s.Configs = v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetDryRun(v bool) *UpdateCollectorResponseBodyResult {
	s.DryRun = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetExtendConfigs(v []*UpdateCollectorResponseBodyResultExtendConfigs) *UpdateCollectorResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetGmtCreatedTime(v string) *UpdateCollectorResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetGmtUpdateTime(v string) *UpdateCollectorResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetName(v string) *UpdateCollectorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetOwnerId(v string) *UpdateCollectorResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetResId(v string) *UpdateCollectorResponseBodyResult {
	s.ResId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetResType(v string) *UpdateCollectorResponseBodyResult {
	s.ResType = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetResVersion(v string) *UpdateCollectorResponseBodyResult {
	s.ResVersion = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetStatus(v string) *UpdateCollectorResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetVpcId(v string) *UpdateCollectorResponseBodyResult {
	s.VpcId = &v
	return s
}

type UpdateCollectorResponseBodyResultConfigs struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s UpdateCollectorResponseBodyResultConfigs) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorResponseBodyResultConfigs) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBodyResultConfigs) SetContent(v string) *UpdateCollectorResponseBodyResultConfigs {
	s.Content = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultConfigs) SetFileName(v string) *UpdateCollectorResponseBodyResultConfigs {
	s.FileName = &v
	return s
}

type UpdateCollectorResponseBodyResultExtendConfigs struct {
	ConfigType       *string                                                   `json:"configType,omitempty" xml:"configType,omitempty"`
	EnableMonitoring *bool                                                     `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	GroupId          *string                                                   `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Host             *string                                                   `json:"host,omitempty" xml:"host,omitempty"`
	Hosts            []*string                                                 `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	InstanceId       *string                                                   `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType     *string                                                   `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	KibanaHost       *string                                                   `json:"kibanaHost,omitempty" xml:"kibanaHost,omitempty"`
	Machines         []*UpdateCollectorResponseBodyResultExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	Protocol         *string                                                   `json:"protocol,omitempty" xml:"protocol,omitempty"`
	SuccessPodsCount *string                                                   `json:"successPodsCount,omitempty" xml:"successPodsCount,omitempty"`
	TotalPodsCount   *string                                                   `json:"totalPodsCount,omitempty" xml:"totalPodsCount,omitempty"`
	Type             *string                                                   `json:"type,omitempty" xml:"type,omitempty"`
	UserName         *string                                                   `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s UpdateCollectorResponseBodyResultExtendConfigs) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetConfigType(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetEnableMonitoring(v bool) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetGroupId(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetHost(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Host = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetHosts(v []*string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Hosts = v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetInstanceId(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetInstanceType(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetKibanaHost(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.KibanaHost = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetMachines(v []*UpdateCollectorResponseBodyResultExtendConfigsMachines) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Machines = v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetProtocol(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetSuccessPodsCount(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.SuccessPodsCount = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetTotalPodsCount(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.TotalPodsCount = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetType(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Type = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetUserName(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.UserName = &v
	return s
}

type UpdateCollectorResponseBodyResultExtendConfigsMachines struct {
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateCollectorResponseBodyResultExtendConfigsMachines) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorResponseBodyResultExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBodyResultExtendConfigsMachines) SetAgentStatus(v string) *UpdateCollectorResponseBodyResultExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigsMachines) SetInstanceId(v string) *UpdateCollectorResponseBodyResultExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

type UpdateCollectorResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorResponse) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponse) SetHeaders(v map[string]*string) *UpdateCollectorResponse {
	s.Headers = v
	return s
}

func (s *UpdateCollectorResponse) SetBody(v *UpdateCollectorResponseBody) *UpdateCollectorResponse {
	s.Body = v
	return s
}

type UpdateCollectorNameRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateCollectorNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameRequest) SetClientToken(v string) *UpdateCollectorNameRequest {
	s.ClientToken = &v
	return s
}

type UpdateCollectorNameResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateCollectorNameResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateCollectorNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBody) SetRequestId(v string) *UpdateCollectorNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCollectorNameResponseBody) SetResult(v *UpdateCollectorNameResponseBodyResult) *UpdateCollectorNameResponseBody {
	s.Result = v
	return s
}

type UpdateCollectorNameResponseBodyResult struct {
	CollectorPaths []*string                                             `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	Configs        []*UpdateCollectorNameResponseBodyResultConfigs       `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	DryRun         *bool                                                 `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	ExtendConfigs  []*UpdateCollectorNameResponseBodyResultExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	GmtCreatedTime *string                                               `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	GmtUpdateTime  *string                                               `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	Name           *string                                               `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId        *string                                               `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	ResId          *string                                               `json:"resId,omitempty" xml:"resId,omitempty"`
	ResType        *string                                               `json:"resType,omitempty" xml:"resType,omitempty"`
	ResVersion     *string                                               `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	Status         *string                                               `json:"status,omitempty" xml:"status,omitempty"`
	VpcId          *string                                               `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s UpdateCollectorNameResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorNameResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBodyResult) SetCollectorPaths(v []*string) *UpdateCollectorNameResponseBodyResult {
	s.CollectorPaths = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetConfigs(v []*UpdateCollectorNameResponseBodyResultConfigs) *UpdateCollectorNameResponseBodyResult {
	s.Configs = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetDryRun(v bool) *UpdateCollectorNameResponseBodyResult {
	s.DryRun = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetExtendConfigs(v []*UpdateCollectorNameResponseBodyResultExtendConfigs) *UpdateCollectorNameResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetGmtCreatedTime(v string) *UpdateCollectorNameResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetGmtUpdateTime(v string) *UpdateCollectorNameResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetName(v string) *UpdateCollectorNameResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetOwnerId(v string) *UpdateCollectorNameResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetResId(v string) *UpdateCollectorNameResponseBodyResult {
	s.ResId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetResType(v string) *UpdateCollectorNameResponseBodyResult {
	s.ResType = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetResVersion(v string) *UpdateCollectorNameResponseBodyResult {
	s.ResVersion = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetStatus(v string) *UpdateCollectorNameResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetVpcId(v string) *UpdateCollectorNameResponseBodyResult {
	s.VpcId = &v
	return s
}

type UpdateCollectorNameResponseBodyResultConfigs struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s UpdateCollectorNameResponseBodyResultConfigs) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorNameResponseBodyResultConfigs) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBodyResultConfigs) SetContent(v string) *UpdateCollectorNameResponseBodyResultConfigs {
	s.Content = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultConfigs) SetFileName(v string) *UpdateCollectorNameResponseBodyResultConfigs {
	s.FileName = &v
	return s
}

type UpdateCollectorNameResponseBodyResultExtendConfigs struct {
	ConfigType       *string                                                       `json:"configType,omitempty" xml:"configType,omitempty"`
	EnableMonitoring *bool                                                         `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	GroupId          *string                                                       `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Host             *string                                                       `json:"host,omitempty" xml:"host,omitempty"`
	Hosts            []*string                                                     `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	InstanceId       *string                                                       `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType     *string                                                       `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	KibanaHost       *string                                                       `json:"kibanaHost,omitempty" xml:"kibanaHost,omitempty"`
	Machines         []*UpdateCollectorNameResponseBodyResultExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	Protocol         *string                                                       `json:"protocol,omitempty" xml:"protocol,omitempty"`
	SuccessPodsCount *string                                                       `json:"successPodsCount,omitempty" xml:"successPodsCount,omitempty"`
	TotalPodsCount   *string                                                       `json:"totalPodsCount,omitempty" xml:"totalPodsCount,omitempty"`
	Type             *string                                                       `json:"type,omitempty" xml:"type,omitempty"`
	UserName         *string                                                       `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s UpdateCollectorNameResponseBodyResultExtendConfigs) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorNameResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetConfigType(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetEnableMonitoring(v bool) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetGroupId(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetHost(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Host = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetHosts(v []*string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Hosts = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetInstanceId(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetInstanceType(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetKibanaHost(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.KibanaHost = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetMachines(v []*UpdateCollectorNameResponseBodyResultExtendConfigsMachines) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Machines = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetProtocol(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetSuccessPodsCount(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.SuccessPodsCount = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetTotalPodsCount(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.TotalPodsCount = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetType(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Type = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetUserName(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.UserName = &v
	return s
}

type UpdateCollectorNameResponseBodyResultExtendConfigsMachines struct {
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateCollectorNameResponseBodyResultExtendConfigsMachines) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorNameResponseBodyResultExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigsMachines) SetAgentStatus(v string) *UpdateCollectorNameResponseBodyResultExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigsMachines) SetInstanceId(v string) *UpdateCollectorNameResponseBodyResultExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

type UpdateCollectorNameResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCollectorNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCollectorNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCollectorNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponse) SetHeaders(v map[string]*string) *UpdateCollectorNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateCollectorNameResponse) SetBody(v *UpdateCollectorNameResponseBody) *UpdateCollectorNameResponse {
	s.Body = v
	return s
}

type UpdateComponentIndexRequest struct {
	Meta     map[string]interface{}               `json:"_meta,omitempty" xml:"_meta,omitempty"`
	Template *UpdateComponentIndexRequestTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s UpdateComponentIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateComponentIndexRequest) GoString() string {
	return s.String()
}

func (s *UpdateComponentIndexRequest) SetMeta(v map[string]interface{}) *UpdateComponentIndexRequest {
	s.Meta = v
	return s
}

func (s *UpdateComponentIndexRequest) SetTemplate(v *UpdateComponentIndexRequestTemplate) *UpdateComponentIndexRequest {
	s.Template = v
	return s
}

type UpdateComponentIndexRequestTemplate struct {
	Aliases  map[string]interface{} `json:"aliases,omitempty" xml:"aliases,omitempty"`
	Mappings map[string]interface{} `json:"mappings,omitempty" xml:"mappings,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty" xml:"settings,omitempty"`
}

func (s UpdateComponentIndexRequestTemplate) String() string {
	return tea.Prettify(s)
}

func (s UpdateComponentIndexRequestTemplate) GoString() string {
	return s.String()
}

func (s *UpdateComponentIndexRequestTemplate) SetAliases(v map[string]interface{}) *UpdateComponentIndexRequestTemplate {
	s.Aliases = v
	return s
}

func (s *UpdateComponentIndexRequestTemplate) SetMappings(v map[string]interface{}) *UpdateComponentIndexRequestTemplate {
	s.Mappings = v
	return s
}

func (s *UpdateComponentIndexRequestTemplate) SetSettings(v map[string]interface{}) *UpdateComponentIndexRequestTemplate {
	s.Settings = v
	return s
}

type UpdateComponentIndexResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComponentIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateComponentIndexResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComponentIndexResponseBody) SetRequestId(v string) *UpdateComponentIndexResponseBody {
	s.RequestId = &v
	return s
}

type UpdateComponentIndexResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateComponentIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateComponentIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateComponentIndexResponse) GoString() string {
	return s.String()
}

func (s *UpdateComponentIndexResponse) SetHeaders(v map[string]*string) *UpdateComponentIndexResponse {
	s.Headers = v
	return s
}

func (s *UpdateComponentIndexResponse) SetBody(v *UpdateComponentIndexResponseBody) *UpdateComponentIndexResponse {
	s.Body = v
	return s
}

type UpdateDescriptionRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateDescriptionRequest) SetDescription(v string) *UpdateDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateDescriptionRequest) SetClientToken(v string) *UpdateDescriptionRequest {
	s.ClientToken = &v
	return s
}

type UpdateDescriptionResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateDescriptionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDescriptionResponseBody) SetRequestId(v string) *UpdateDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDescriptionResponseBody) SetResult(v *UpdateDescriptionResponseBodyResult) *UpdateDescriptionResponseBody {
	s.Result = v
	return s
}

type UpdateDescriptionResponseBodyResult struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateDescriptionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateDescriptionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateDescriptionResponseBodyResult) SetDescription(v string) *UpdateDescriptionResponseBodyResult {
	s.Description = &v
	return s
}

type UpdateDescriptionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateDescriptionResponse) SetHeaders(v map[string]*string) *UpdateDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateDescriptionResponse) SetBody(v *UpdateDescriptionResponseBody) *UpdateDescriptionResponse {
	s.Body = v
	return s
}

type UpdateDiagnosisSettingsRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang        *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s UpdateDiagnosisSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDiagnosisSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateDiagnosisSettingsRequest) SetClientToken(v string) *UpdateDiagnosisSettingsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDiagnosisSettingsRequest) SetLang(v string) *UpdateDiagnosisSettingsRequest {
	s.Lang = &v
	return s
}

type UpdateDiagnosisSettingsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateDiagnosisSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDiagnosisSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDiagnosisSettingsResponseBody) SetRequestId(v string) *UpdateDiagnosisSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDiagnosisSettingsResponseBody) SetResult(v bool) *UpdateDiagnosisSettingsResponseBody {
	s.Result = &v
	return s
}

type UpdateDiagnosisSettingsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDiagnosisSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDiagnosisSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDiagnosisSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateDiagnosisSettingsResponse) SetHeaders(v map[string]*string) *UpdateDiagnosisSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateDiagnosisSettingsResponse) SetBody(v *UpdateDiagnosisSettingsResponseBody) *UpdateDiagnosisSettingsResponse {
	s.Body = v
	return s
}

type UpdateDictRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateDictRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDictRequest) GoString() string {
	return s.String()
}

func (s *UpdateDictRequest) SetClientToken(v string) *UpdateDictRequest {
	s.ClientToken = &v
	return s
}

type UpdateDictResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*UpdateDictResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateDictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDictResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDictResponseBody) SetRequestId(v string) *UpdateDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDictResponseBody) SetResult(v []*UpdateDictResponseBodyResult) *UpdateDictResponseBody {
	s.Result = v
	return s
}

type UpdateDictResponseBodyResult struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateDictResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateDictResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateDictResponseBodyResult) SetFileSize(v int64) *UpdateDictResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateDictResponseBodyResult) SetName(v string) *UpdateDictResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateDictResponseBodyResult) SetSourceType(v string) *UpdateDictResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateDictResponseBodyResult) SetType(v string) *UpdateDictResponseBodyResult {
	s.Type = &v
	return s
}

type UpdateDictResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDictResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDictResponse) GoString() string {
	return s.String()
}

func (s *UpdateDictResponse) SetHeaders(v map[string]*string) *UpdateDictResponse {
	s.Headers = v
	return s
}

func (s *UpdateDictResponse) SetBody(v *UpdateDictResponseBody) *UpdateDictResponse {
	s.Body = v
	return s
}

type UpdateExtendConfigRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateExtendConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtendConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateExtendConfigRequest) SetClientToken(v string) *UpdateExtendConfigRequest {
	s.ClientToken = &v
	return s
}

type UpdateExtendConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateExtendConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtendConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExtendConfigResponseBody) SetRequestId(v string) *UpdateExtendConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExtendConfigResponseBody) SetResult(v bool) *UpdateExtendConfigResponseBody {
	s.Result = &v
	return s
}

type UpdateExtendConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateExtendConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateExtendConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtendConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateExtendConfigResponse) SetHeaders(v map[string]*string) *UpdateExtendConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateExtendConfigResponse) SetBody(v *UpdateExtendConfigResponseBody) *UpdateExtendConfigResponse {
	s.Body = v
	return s
}

type UpdateExtendfilesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateExtendfilesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtendfilesRequest) GoString() string {
	return s.String()
}

func (s *UpdateExtendfilesRequest) SetClientToken(v string) *UpdateExtendfilesRequest {
	s.ClientToken = &v
	return s
}

type UpdateExtendfilesResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*UpdateExtendfilesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateExtendfilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtendfilesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExtendfilesResponseBody) SetRequestId(v string) *UpdateExtendfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExtendfilesResponseBody) SetResult(v []*UpdateExtendfilesResponseBodyResult) *UpdateExtendfilesResponseBody {
	s.Result = v
	return s
}

type UpdateExtendfilesResponseBodyResult struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s UpdateExtendfilesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtendfilesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateExtendfilesResponseBodyResult) SetFileSize(v int64) *UpdateExtendfilesResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateExtendfilesResponseBodyResult) SetName(v string) *UpdateExtendfilesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateExtendfilesResponseBodyResult) SetSourceType(v string) *UpdateExtendfilesResponseBodyResult {
	s.SourceType = &v
	return s
}

type UpdateExtendfilesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateExtendfilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateExtendfilesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtendfilesResponse) GoString() string {
	return s.String()
}

func (s *UpdateExtendfilesResponse) SetHeaders(v map[string]*string) *UpdateExtendfilesResponse {
	s.Headers = v
	return s
}

func (s *UpdateExtendfilesResponse) SetBody(v *UpdateExtendfilesResponseBody) *UpdateExtendfilesResponse {
	s.Body = v
	return s
}

type UpdateHotIkDictsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateHotIkDictsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotIkDictsRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotIkDictsRequest) SetClientToken(v string) *UpdateHotIkDictsRequest {
	s.ClientToken = &v
	return s
}

type UpdateHotIkDictsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*UpdateHotIkDictsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateHotIkDictsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotIkDictsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotIkDictsResponseBody) SetRequestId(v string) *UpdateHotIkDictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotIkDictsResponseBody) SetResult(v []*UpdateHotIkDictsResponseBodyResult) *UpdateHotIkDictsResponseBody {
	s.Result = v
	return s
}

type UpdateHotIkDictsResponseBodyResult struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHotIkDictsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotIkDictsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateHotIkDictsResponseBodyResult) SetFileSize(v int64) *UpdateHotIkDictsResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateHotIkDictsResponseBodyResult) SetName(v string) *UpdateHotIkDictsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateHotIkDictsResponseBodyResult) SetSourceType(v string) *UpdateHotIkDictsResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateHotIkDictsResponseBodyResult) SetType(v string) *UpdateHotIkDictsResponseBodyResult {
	s.Type = &v
	return s
}

type UpdateHotIkDictsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateHotIkDictsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHotIkDictsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHotIkDictsResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotIkDictsResponse) SetHeaders(v map[string]*string) *UpdateHotIkDictsResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotIkDictsResponse) SetBody(v *UpdateHotIkDictsResponseBody) *UpdateHotIkDictsResponse {
	s.Body = v
	return s
}

type UpdateILMPolicyRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateILMPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateILMPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateILMPolicyRequest) SetClientToken(v string) *UpdateILMPolicyRequest {
	s.ClientToken = &v
	return s
}

type UpdateILMPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateILMPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateILMPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateILMPolicyResponseBody) SetRequestId(v string) *UpdateILMPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateILMPolicyResponseBody) SetResult(v string) *UpdateILMPolicyResponseBody {
	s.Result = &v
	return s
}

type UpdateILMPolicyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateILMPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateILMPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateILMPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateILMPolicyResponse) SetHeaders(v map[string]*string) *UpdateILMPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateILMPolicyResponse) SetBody(v *UpdateILMPolicyResponseBody) *UpdateILMPolicyResponse {
	s.Body = v
	return s
}

type UpdateIndexTemplateRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateIndexTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIndexTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateIndexTemplateRequest) SetClientToken(v string) *UpdateIndexTemplateRequest {
	s.ClientToken = &v
	return s
}

type UpdateIndexTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateIndexTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIndexTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIndexTemplateResponseBody) SetRequestId(v string) *UpdateIndexTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIndexTemplateResponseBody) SetResult(v string) *UpdateIndexTemplateResponseBody {
	s.Result = &v
	return s
}

type UpdateIndexTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateIndexTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIndexTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIndexTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateIndexTemplateResponse) SetHeaders(v map[string]*string) *UpdateIndexTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateIndexTemplateResponse) SetBody(v *UpdateIndexTemplateResponseBody) *UpdateIndexTemplateResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	ClientToken     *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	IgnoreStatus    *bool   `json:"ignoreStatus,omitempty" xml:"ignoreStatus,omitempty"`
	OrderActionType *string `json:"orderActionType,omitempty" xml:"orderActionType,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetClientToken(v string) *UpdateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceRequest) SetIgnoreStatus(v bool) *UpdateInstanceRequest {
	s.IgnoreStatus = &v
	return s
}

func (s *UpdateInstanceRequest) SetOrderActionType(v string) *UpdateInstanceRequest {
	s.OrderActionType = &v
	return s
}

type UpdateInstanceResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetResult(v *UpdateInstanceResponseBodyResult) *UpdateInstanceResponseBody {
	s.Result = v
	return s
}

type UpdateInstanceResponseBodyResult struct {
	CreatedAt           *string                                              `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description         *string                                              `json:"description,omitempty" xml:"description,omitempty"`
	DictList            []*UpdateInstanceResponseBodyResultDictList          `json:"dictList,omitempty" xml:"dictList,omitempty" type:"Repeated"`
	Domain              *string                                              `json:"domain,omitempty" xml:"domain,omitempty"`
	EsVersion           *string                                              `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	InstanceId          *string                                              `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	KibanaConfiguration *UpdateInstanceResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	KibanaDomain        *string                                              `json:"kibanaDomain,omitempty" xml:"kibanaDomain,omitempty"`
	KibanaPort          *int32                                               `json:"kibanaPort,omitempty" xml:"kibanaPort,omitempty"`
	MasterConfiguration *UpdateInstanceResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	NetworkConfig       *UpdateInstanceResponseBodyResultNetworkConfig       `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	NodeAmount          *int32                                               `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec            *UpdateInstanceResponseBodyResultNodeSpec            `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	PaymentType         *string                                              `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	PublicDomain        *string                                              `json:"publicDomain,omitempty" xml:"publicDomain,omitempty"`
	PublicPort          *int32                                               `json:"publicPort,omitempty" xml:"publicPort,omitempty"`
	Status              *string                                              `json:"status,omitempty" xml:"status,omitempty"`
	SynonymsDicts       []*UpdateInstanceResponseBodyResultSynonymsDicts     `json:"synonymsDicts,omitempty" xml:"synonymsDicts,omitempty" type:"Repeated"`
	UpdatedAt           *string                                              `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResult) SetCreatedAt(v string) *UpdateInstanceResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetDescription(v string) *UpdateInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetDictList(v []*UpdateInstanceResponseBodyResultDictList) *UpdateInstanceResponseBodyResult {
	s.DictList = v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetDomain(v string) *UpdateInstanceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetEsVersion(v string) *UpdateInstanceResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetInstanceId(v string) *UpdateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetKibanaConfiguration(v *UpdateInstanceResponseBodyResultKibanaConfiguration) *UpdateInstanceResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetKibanaDomain(v string) *UpdateInstanceResponseBodyResult {
	s.KibanaDomain = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetKibanaPort(v int32) *UpdateInstanceResponseBodyResult {
	s.KibanaPort = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetMasterConfiguration(v *UpdateInstanceResponseBodyResultMasterConfiguration) *UpdateInstanceResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetNetworkConfig(v *UpdateInstanceResponseBodyResultNetworkConfig) *UpdateInstanceResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetNodeAmount(v int32) *UpdateInstanceResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetNodeSpec(v *UpdateInstanceResponseBodyResultNodeSpec) *UpdateInstanceResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetPaymentType(v string) *UpdateInstanceResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetPublicDomain(v string) *UpdateInstanceResponseBodyResult {
	s.PublicDomain = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetPublicPort(v int32) *UpdateInstanceResponseBodyResult {
	s.PublicPort = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetStatus(v string) *UpdateInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetSynonymsDicts(v []*UpdateInstanceResponseBodyResultSynonymsDicts) *UpdateInstanceResponseBodyResult {
	s.SynonymsDicts = v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetUpdatedAt(v string) *UpdateInstanceResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

type UpdateInstanceResponseBodyResultDictList struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateInstanceResponseBodyResultDictList) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBodyResultDictList) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResultDictList) SetFileSize(v int64) *UpdateInstanceResponseBodyResultDictList {
	s.FileSize = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultDictList) SetName(v string) *UpdateInstanceResponseBodyResultDictList {
	s.Name = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultDictList) SetSourceType(v string) *UpdateInstanceResponseBodyResultDictList {
	s.SourceType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultDictList) SetType(v string) *UpdateInstanceResponseBodyResultDictList {
	s.Type = &v
	return s
}

type UpdateInstanceResponseBodyResultKibanaConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s UpdateInstanceResponseBodyResultKibanaConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) SetAmount(v int32) *UpdateInstanceResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) SetDisk(v int32) *UpdateInstanceResponseBodyResultKibanaConfiguration {
	s.Disk = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) SetDiskType(v string) *UpdateInstanceResponseBodyResultKibanaConfiguration {
	s.DiskType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) SetSpec(v string) *UpdateInstanceResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

type UpdateInstanceResponseBodyResultMasterConfiguration struct {
	Amount   *int32  `json:"amount,omitempty" xml:"amount,omitempty"`
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s UpdateInstanceResponseBodyResultMasterConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) SetAmount(v int32) *UpdateInstanceResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) SetDisk(v int32) *UpdateInstanceResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) SetDiskType(v string) *UpdateInstanceResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) SetSpec(v string) *UpdateInstanceResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

type UpdateInstanceResponseBodyResultNetworkConfig struct {
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	VpcId     *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VsArea    *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s UpdateInstanceResponseBodyResultNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResultNetworkConfig) SetType(v string) *UpdateInstanceResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultNetworkConfig) SetVpcId(v string) *UpdateInstanceResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultNetworkConfig) SetVsArea(v string) *UpdateInstanceResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultNetworkConfig) SetVswitchId(v string) *UpdateInstanceResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

type UpdateInstanceResponseBodyResultNodeSpec struct {
	Disk     *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s UpdateInstanceResponseBodyResultNodeSpec) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) SetDisk(v int32) *UpdateInstanceResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) SetDiskType(v string) *UpdateInstanceResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) SetSpec(v string) *UpdateInstanceResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

type UpdateInstanceResponseBodyResultSynonymsDicts struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateInstanceResponseBodyResultSynonymsDicts) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBodyResultSynonymsDicts) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResultSynonymsDicts) SetFileSize(v int64) *UpdateInstanceResponseBodyResultSynonymsDicts {
	s.FileSize = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultSynonymsDicts) SetName(v string) *UpdateInstanceResponseBodyResultSynonymsDicts {
	s.Name = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultSynonymsDicts) SetSourceType(v string) *UpdateInstanceResponseBodyResultSynonymsDicts {
	s.SourceType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultSynonymsDicts) SetType(v string) *UpdateInstanceResponseBodyResultSynonymsDicts {
	s.Type = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceChargeTypeRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateInstanceChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceChargeTypeRequest) SetClientToken(v string) *UpdateInstanceChargeTypeRequest {
	s.ClientToken = &v
	return s
}

type UpdateInstanceChargeTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateInstanceChargeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceChargeTypeResponseBody) SetRequestId(v string) *UpdateInstanceChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceChargeTypeResponseBody) SetResult(v bool) *UpdateInstanceChargeTypeResponseBody {
	s.Result = &v
	return s
}

type UpdateInstanceChargeTypeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceChargeTypeResponse) SetHeaders(v map[string]*string) *UpdateInstanceChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceChargeTypeResponse) SetBody(v *UpdateInstanceChargeTypeResponseBody) *UpdateInstanceChargeTypeResponse {
	s.Body = v
	return s
}

type UpdateInstanceSettingsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateInstanceSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSettingsRequest) SetClientToken(v string) *UpdateInstanceSettingsRequest {
	s.ClientToken = &v
	return s
}

type UpdateInstanceSettingsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSettingsResponseBody) SetRequestId(v string) *UpdateInstanceSettingsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceSettingsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSettingsResponse) SetHeaders(v map[string]*string) *UpdateInstanceSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceSettingsResponse) SetBody(v *UpdateInstanceSettingsResponseBody) *UpdateInstanceSettingsResponse {
	s.Body = v
	return s
}

type UpdateKibanaSettingsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateKibanaSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKibanaSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateKibanaSettingsRequest) SetClientToken(v string) *UpdateKibanaSettingsRequest {
	s.ClientToken = &v
	return s
}

type UpdateKibanaSettingsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateKibanaSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKibanaSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKibanaSettingsResponseBody) SetRequestId(v string) *UpdateKibanaSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKibanaSettingsResponseBody) SetResult(v bool) *UpdateKibanaSettingsResponseBody {
	s.Result = &v
	return s
}

type UpdateKibanaSettingsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateKibanaSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateKibanaSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKibanaSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateKibanaSettingsResponse) SetHeaders(v map[string]*string) *UpdateKibanaSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateKibanaSettingsResponse) SetBody(v *UpdateKibanaSettingsResponseBody) *UpdateKibanaSettingsResponse {
	s.Body = v
	return s
}

type UpdateKibanaWhiteIpsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ModifyMode  *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdateKibanaWhiteIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKibanaWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsRequest) SetClientToken(v string) *UpdateKibanaWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateKibanaWhiteIpsRequest) SetModifyMode(v string) *UpdateKibanaWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

type UpdateKibanaWhiteIpsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateKibanaWhiteIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateKibanaWhiteIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKibanaWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsResponseBody) SetRequestId(v string) *UpdateKibanaWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBody) SetResult(v *UpdateKibanaWhiteIpsResponseBodyResult) *UpdateKibanaWhiteIpsResponseBody {
	s.Result = v
	return s
}

type UpdateKibanaWhiteIpsResponseBodyResult struct {
	KibanaIPWhitelist        []*string `json:"kibanaIPWhitelist,omitempty" xml:"kibanaIPWhitelist,omitempty" type:"Repeated"`
	KibanaPrivateIPWhitelist []*string `json:"kibanaPrivateIPWhitelist,omitempty" xml:"kibanaPrivateIPWhitelist,omitempty" type:"Repeated"`
}

func (s UpdateKibanaWhiteIpsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateKibanaWhiteIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsResponseBodyResult) SetKibanaIPWhitelist(v []*string) *UpdateKibanaWhiteIpsResponseBodyResult {
	s.KibanaIPWhitelist = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResult) SetKibanaPrivateIPWhitelist(v []*string) *UpdateKibanaWhiteIpsResponseBodyResult {
	s.KibanaPrivateIPWhitelist = v
	return s
}

type UpdateKibanaWhiteIpsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateKibanaWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateKibanaWhiteIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKibanaWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsResponse) SetHeaders(v map[string]*string) *UpdateKibanaWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponse) SetBody(v *UpdateKibanaWhiteIpsResponseBody) *UpdateKibanaWhiteIpsResponse {
	s.Body = v
	return s
}

type UpdateLogstashRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateLogstashRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogstashRequest) SetClientToken(v string) *UpdateLogstashRequest {
	s.ClientToken = &v
	return s
}

type UpdateLogstashResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateLogstashResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLogstashResponseBody) SetRequestId(v string) *UpdateLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLogstashResponseBody) SetResult(v bool) *UpdateLogstashResponseBody {
	s.Result = &v
	return s
}

type UpdateLogstashResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLogstashResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogstashResponse) SetHeaders(v map[string]*string) *UpdateLogstashResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogstashResponse) SetBody(v *UpdateLogstashResponseBody) *UpdateLogstashResponse {
	s.Body = v
	return s
}

type UpdateLogstashChargeTypeRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateLogstashChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogstashChargeTypeRequest) SetClientToken(v string) *UpdateLogstashChargeTypeRequest {
	s.ClientToken = &v
	return s
}

type UpdateLogstashChargeTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateLogstashChargeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLogstashChargeTypeResponseBody) SetRequestId(v string) *UpdateLogstashChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLogstashChargeTypeResponseBody) SetResult(v bool) *UpdateLogstashChargeTypeResponseBody {
	s.Result = &v
	return s
}

type UpdateLogstashChargeTypeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLogstashChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLogstashChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogstashChargeTypeResponse) SetHeaders(v map[string]*string) *UpdateLogstashChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogstashChargeTypeResponse) SetBody(v *UpdateLogstashChargeTypeResponseBody) *UpdateLogstashChargeTypeResponse {
	s.Body = v
	return s
}

type UpdateLogstashDescriptionRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateLogstashDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogstashDescriptionRequest) SetClientToken(v string) *UpdateLogstashDescriptionRequest {
	s.ClientToken = &v
	return s
}

type UpdateLogstashDescriptionResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateLogstashDescriptionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateLogstashDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLogstashDescriptionResponseBody) SetRequestId(v string) *UpdateLogstashDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLogstashDescriptionResponseBody) SetResult(v *UpdateLogstashDescriptionResponseBodyResult) *UpdateLogstashDescriptionResponseBody {
	s.Result = v
	return s
}

type UpdateLogstashDescriptionResponseBodyResult struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateLogstashDescriptionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashDescriptionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateLogstashDescriptionResponseBodyResult) SetDescription(v string) *UpdateLogstashDescriptionResponseBodyResult {
	s.Description = &v
	return s
}

type UpdateLogstashDescriptionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLogstashDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLogstashDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogstashDescriptionResponse) SetHeaders(v map[string]*string) *UpdateLogstashDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogstashDescriptionResponse) SetBody(v *UpdateLogstashDescriptionResponseBody) *UpdateLogstashDescriptionResponse {
	s.Body = v
	return s
}

type UpdateLogstashSettingsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateLogstashSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogstashSettingsRequest) SetClientToken(v string) *UpdateLogstashSettingsRequest {
	s.ClientToken = &v
	return s
}

type UpdateLogstashSettingsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLogstashSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLogstashSettingsResponseBody) SetRequestId(v string) *UpdateLogstashSettingsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLogstashSettingsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLogstashSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLogstashSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogstashSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogstashSettingsResponse) SetHeaders(v map[string]*string) *UpdateLogstashSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogstashSettingsResponse) SetBody(v *UpdateLogstashSettingsResponseBody) *UpdateLogstashSettingsResponse {
	s.Body = v
	return s
}

type UpdatePipelineManagementConfigRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdatePipelineManagementConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineManagementConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineManagementConfigRequest) SetClientToken(v string) *UpdatePipelineManagementConfigRequest {
	s.ClientToken = &v
	return s
}

type UpdatePipelineManagementConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdatePipelineManagementConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineManagementConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineManagementConfigResponseBody) SetRequestId(v string) *UpdatePipelineManagementConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineManagementConfigResponseBody) SetResult(v bool) *UpdatePipelineManagementConfigResponseBody {
	s.Result = &v
	return s
}

type UpdatePipelineManagementConfigResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePipelineManagementConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePipelineManagementConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineManagementConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineManagementConfigResponse) SetHeaders(v map[string]*string) *UpdatePipelineManagementConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineManagementConfigResponse) SetBody(v *UpdatePipelineManagementConfigResponseBody) *UpdatePipelineManagementConfigResponse {
	s.Body = v
	return s
}

type UpdatePipelinesRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Trigger     *bool   `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (s UpdatePipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelinesRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelinesRequest) SetClientToken(v string) *UpdatePipelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePipelinesRequest) SetTrigger(v bool) *UpdatePipelinesRequest {
	s.Trigger = &v
	return s
}

type UpdatePipelinesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdatePipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelinesResponseBody) SetRequestId(v string) *UpdatePipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelinesResponseBody) SetResult(v bool) *UpdatePipelinesResponseBody {
	s.Result = &v
	return s
}

type UpdatePipelinesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelinesResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelinesResponse) SetHeaders(v map[string]*string) *UpdatePipelinesResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelinesResponse) SetBody(v *UpdatePipelinesResponseBody) *UpdatePipelinesResponse {
	s.Body = v
	return s
}

type UpdatePrivateNetworkWhiteIpsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ModifyMode  *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdatePrivateNetworkWhiteIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateNetworkWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetworkWhiteIpsRequest) SetClientToken(v string) *UpdatePrivateNetworkWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsRequest) SetModifyMode(v string) *UpdatePrivateNetworkWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

type UpdatePrivateNetworkWhiteIpsResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdatePrivateNetworkWhiteIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdatePrivateNetworkWhiteIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateNetworkWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBody) SetRequestId(v string) *UpdatePrivateNetworkWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBody) SetResult(v *UpdatePrivateNetworkWhiteIpsResponseBodyResult) *UpdatePrivateNetworkWhiteIpsResponseBody {
	s.Result = v
	return s
}

type UpdatePrivateNetworkWhiteIpsResponseBodyResult struct {
	PrivateNetworkIpWhiteList []*string `json:"privateNetworkIpWhiteList,omitempty" xml:"privateNetworkIpWhiteList,omitempty" type:"Repeated"`
}

func (s UpdatePrivateNetworkWhiteIpsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateNetworkWhiteIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetworkWhiteIpsResponseBodyResult) SetPrivateNetworkIpWhiteList(v []*string) *UpdatePrivateNetworkWhiteIpsResponseBodyResult {
	s.PrivateNetworkIpWhiteList = v
	return s
}

type UpdatePrivateNetworkWhiteIpsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePrivateNetworkWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePrivateNetworkWhiteIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateNetworkWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetworkWhiteIpsResponse) SetHeaders(v map[string]*string) *UpdatePrivateNetworkWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsResponse) SetBody(v *UpdatePrivateNetworkWhiteIpsResponseBody) *UpdatePrivateNetworkWhiteIpsResponse {
	s.Body = v
	return s
}

type UpdatePublicNetworkRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdatePublicNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicNetworkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkRequest) SetClientToken(v string) *UpdatePublicNetworkRequest {
	s.ClientToken = &v
	return s
}

type UpdatePublicNetworkResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdatePublicNetworkResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdatePublicNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkResponseBody) SetRequestId(v string) *UpdatePublicNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublicNetworkResponseBody) SetResult(v *UpdatePublicNetworkResponseBodyResult) *UpdatePublicNetworkResponseBody {
	s.Result = v
	return s
}

type UpdatePublicNetworkResponseBodyResult struct {
	EnablePublic *bool `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
}

func (s UpdatePublicNetworkResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicNetworkResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkResponseBodyResult) SetEnablePublic(v bool) *UpdatePublicNetworkResponseBodyResult {
	s.EnablePublic = &v
	return s
}

type UpdatePublicNetworkResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePublicNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePublicNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicNetworkResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkResponse) SetHeaders(v map[string]*string) *UpdatePublicNetworkResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicNetworkResponse) SetBody(v *UpdatePublicNetworkResponseBody) *UpdatePublicNetworkResponse {
	s.Body = v
	return s
}

type UpdatePublicWhiteIpsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ModifyMode  *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdatePublicWhiteIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicWhiteIpsRequest) SetClientToken(v string) *UpdatePublicWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePublicWhiteIpsRequest) SetModifyMode(v string) *UpdatePublicWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

type UpdatePublicWhiteIpsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdatePublicWhiteIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdatePublicWhiteIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicWhiteIpsResponseBody) SetRequestId(v string) *UpdatePublicWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublicWhiteIpsResponseBody) SetResult(v *UpdatePublicWhiteIpsResponseBodyResult) *UpdatePublicWhiteIpsResponseBody {
	s.Result = v
	return s
}

type UpdatePublicWhiteIpsResponseBodyResult struct {
	PublicIpWhitelist []*string `json:"publicIpWhitelist,omitempty" xml:"publicIpWhitelist,omitempty" type:"Repeated"`
}

func (s UpdatePublicWhiteIpsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicWhiteIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePublicWhiteIpsResponseBodyResult) SetPublicIpWhitelist(v []*string) *UpdatePublicWhiteIpsResponseBodyResult {
	s.PublicIpWhitelist = v
	return s
}

type UpdatePublicWhiteIpsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePublicWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePublicWhiteIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicWhiteIpsResponse) SetHeaders(v map[string]*string) *UpdatePublicWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicWhiteIpsResponse) SetBody(v *UpdatePublicWhiteIpsResponseBody) *UpdatePublicWhiteIpsResponse {
	s.Body = v
	return s
}

type UpdateReadWritePolicyRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateReadWritePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateReadWritePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateReadWritePolicyRequest) SetClientToken(v string) *UpdateReadWritePolicyRequest {
	s.ClientToken = &v
	return s
}

type UpdateReadWritePolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateReadWritePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateReadWritePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateReadWritePolicyResponseBody) SetRequestId(v string) *UpdateReadWritePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateReadWritePolicyResponseBody) SetResult(v bool) *UpdateReadWritePolicyResponseBody {
	s.Result = &v
	return s
}

type UpdateReadWritePolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateReadWritePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateReadWritePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateReadWritePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateReadWritePolicyResponse) SetHeaders(v map[string]*string) *UpdateReadWritePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateReadWritePolicyResponse) SetBody(v *UpdateReadWritePolicyResponseBody) *UpdateReadWritePolicyResponse {
	s.Body = v
	return s
}

type UpdateSnapshotSettingResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateSnapshotSettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateSnapshotSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingResponseBody) SetRequestId(v string) *UpdateSnapshotSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSnapshotSettingResponseBody) SetResult(v *UpdateSnapshotSettingResponseBodyResult) *UpdateSnapshotSettingResponseBody {
	s.Result = v
	return s
}

type UpdateSnapshotSettingResponseBodyResult struct {
	Enable      *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	QuartzRegex *string `json:"quartzRegex,omitempty" xml:"quartzRegex,omitempty"`
}

func (s UpdateSnapshotSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingResponseBodyResult) SetEnable(v bool) *UpdateSnapshotSettingResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *UpdateSnapshotSettingResponseBodyResult) SetQuartzRegex(v string) *UpdateSnapshotSettingResponseBodyResult {
	s.QuartzRegex = &v
	return s
}

type UpdateSnapshotSettingResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSnapshotSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSnapshotSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingResponse) SetHeaders(v map[string]*string) *UpdateSnapshotSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateSnapshotSettingResponse) SetBody(v *UpdateSnapshotSettingResponseBody) *UpdateSnapshotSettingResponse {
	s.Body = v
	return s
}

type UpdateSynonymsDictsRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateSynonymsDictsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSynonymsDictsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsDictsRequest) SetClientToken(v string) *UpdateSynonymsDictsRequest {
	s.ClientToken = &v
	return s
}

type UpdateSynonymsDictsResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*UpdateSynonymsDictsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpdateSynonymsDictsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSynonymsDictsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsDictsResponseBody) SetRequestId(v string) *UpdateSynonymsDictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSynonymsDictsResponseBody) SetResult(v []*UpdateSynonymsDictsResponseBodyResult) *UpdateSynonymsDictsResponseBody {
	s.Result = v
	return s
}

type UpdateSynonymsDictsResponseBodyResult struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateSynonymsDictsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateSynonymsDictsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsDictsResponseBodyResult) SetFileSize(v int64) *UpdateSynonymsDictsResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *UpdateSynonymsDictsResponseBodyResult) SetName(v string) *UpdateSynonymsDictsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateSynonymsDictsResponseBodyResult) SetSourceType(v string) *UpdateSynonymsDictsResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateSynonymsDictsResponseBodyResult) SetType(v string) *UpdateSynonymsDictsResponseBodyResult {
	s.Type = &v
	return s
}

type UpdateSynonymsDictsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSynonymsDictsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSynonymsDictsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSynonymsDictsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsDictsResponse) SetHeaders(v map[string]*string) *UpdateSynonymsDictsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSynonymsDictsResponse) SetBody(v *UpdateSynonymsDictsResponseBody) *UpdateSynonymsDictsResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetClientToken(v string) *UpdateTemplateRequest {
	s.ClientToken = &v
	return s
}

type UpdateTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetResult(v bool) *UpdateTemplateResponseBody {
	s.Result = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

type UpdateWhiteIpsRequest struct {
	EsIPWhitelist []*string                          `json:"esIPWhitelist,omitempty" xml:"esIPWhitelist,omitempty" type:"Repeated"`
	WhiteIpGroup  *UpdateWhiteIpsRequestWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Struct"`
	ClientToken   *string                            `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ModifyMode    *string                            `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdateWhiteIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsRequest) SetEsIPWhitelist(v []*string) *UpdateWhiteIpsRequest {
	s.EsIPWhitelist = v
	return s
}

func (s *UpdateWhiteIpsRequest) SetWhiteIpGroup(v *UpdateWhiteIpsRequestWhiteIpGroup) *UpdateWhiteIpsRequest {
	s.WhiteIpGroup = v
	return s
}

func (s *UpdateWhiteIpsRequest) SetClientToken(v string) *UpdateWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWhiteIpsRequest) SetModifyMode(v string) *UpdateWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

type UpdateWhiteIpsRequestWhiteIpGroup struct {
	GroupName   *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips         []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	WhiteIpType *string   `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s UpdateWhiteIpsRequestWhiteIpGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteIpsRequestWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) SetGroupName(v string) *UpdateWhiteIpsRequestWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) SetIps(v []*string) *UpdateWhiteIpsRequestWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) SetWhiteIpType(v string) *UpdateWhiteIpsRequestWhiteIpGroup {
	s.WhiteIpType = &v
	return s
}

type UpdateWhiteIpsResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateWhiteIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateWhiteIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponseBody) SetRequestId(v string) *UpdateWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWhiteIpsResponseBody) SetResult(v *UpdateWhiteIpsResponseBodyResult) *UpdateWhiteIpsResponseBody {
	s.Result = v
	return s
}

type UpdateWhiteIpsResponseBodyResult struct {
	EsIPWhitelist []*string                                      `json:"esIPWhitelist,omitempty" xml:"esIPWhitelist,omitempty" type:"Repeated"`
	NetworkConfig *UpdateWhiteIpsResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
}

func (s UpdateWhiteIpsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponseBodyResult) SetEsIPWhitelist(v []*string) *UpdateWhiteIpsResponseBodyResult {
	s.EsIPWhitelist = v
	return s
}

func (s *UpdateWhiteIpsResponseBodyResult) SetNetworkConfig(v *UpdateWhiteIpsResponseBodyResultNetworkConfig) *UpdateWhiteIpsResponseBodyResult {
	s.NetworkConfig = v
	return s
}

type UpdateWhiteIpsResponseBodyResultNetworkConfig struct {
	WhiteIpGroupList []*UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList `json:"whiteIpGroupList,omitempty" xml:"whiteIpGroupList,omitempty" type:"Repeated"`
}

func (s UpdateWhiteIpsResponseBodyResultNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteIpsResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfig) SetWhiteIpGroupList(v []*UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) *UpdateWhiteIpsResponseBodyResultNetworkConfig {
	s.WhiteIpGroupList = v
	return s
}

type UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList struct {
	GroupName   *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips         []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	WhiteIpType *string   `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) SetGroupName(v string) *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.GroupName = &v
	return s
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) SetIps(v []*string) *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.Ips = v
	return s
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) SetWhiteIpType(v string) *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.WhiteIpType = &v
	return s
}

type UpdateWhiteIpsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWhiteIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponse) SetHeaders(v map[string]*string) *UpdateWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdateWhiteIpsResponse) SetBody(v *UpdateWhiteIpsResponseBody) *UpdateWhiteIpsResponse {
	s.Body = v
	return s
}

type UpdateXpackMonitorConfigRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateXpackMonitorConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateXpackMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateXpackMonitorConfigRequest) SetClientToken(v string) *UpdateXpackMonitorConfigRequest {
	s.ClientToken = &v
	return s
}

type UpdateXpackMonitorConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateXpackMonitorConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateXpackMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateXpackMonitorConfigResponseBody) SetRequestId(v string) *UpdateXpackMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateXpackMonitorConfigResponseBody) SetResult(v bool) *UpdateXpackMonitorConfigResponseBody {
	s.Result = &v
	return s
}

type UpdateXpackMonitorConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateXpackMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateXpackMonitorConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateXpackMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateXpackMonitorConfigResponse) SetHeaders(v map[string]*string) *UpdateXpackMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateXpackMonitorConfigResponse) SetBody(v *UpdateXpackMonitorConfigResponseBody) *UpdateXpackMonitorConfigResponse {
	s.Body = v
	return s
}

type UpgradeEngineVersionRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	DryRun      *bool   `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpgradeEngineVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeEngineVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionRequest) SetClientToken(v string) *UpgradeEngineVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeEngineVersionRequest) SetDryRun(v bool) *UpgradeEngineVersionRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeEngineVersionRequest) SetType(v string) *UpgradeEngineVersionRequest {
	s.Type = &v
	return s
}

func (s *UpgradeEngineVersionRequest) SetVersion(v string) *UpgradeEngineVersionRequest {
	s.Version = &v
	return s
}

type UpgradeEngineVersionResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpgradeEngineVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpgradeEngineVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionResponseBody) SetRequestId(v string) *UpgradeEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeEngineVersionResponseBody) SetResult(v *UpgradeEngineVersionResponseBodyResult) *UpgradeEngineVersionResponseBody {
	s.Result = v
	return s
}

type UpgradeEngineVersionResponseBodyResult struct {
	Status         *string                                               `json:"status,omitempty" xml:"status,omitempty"`
	ValidateResult *UpgradeEngineVersionResponseBodyResultValidateResult `json:"validateResult,omitempty" xml:"validateResult,omitempty" type:"Struct"`
	ValidateType   *string                                               `json:"validateType,omitempty" xml:"validateType,omitempty"`
}

func (s UpgradeEngineVersionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpgradeEngineVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionResponseBodyResult) SetStatus(v string) *UpgradeEngineVersionResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResult) SetValidateResult(v *UpgradeEngineVersionResponseBodyResultValidateResult) *UpgradeEngineVersionResponseBodyResult {
	s.ValidateResult = v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResult) SetValidateType(v string) *UpgradeEngineVersionResponseBodyResult {
	s.ValidateType = &v
	return s
}

type UpgradeEngineVersionResponseBodyResultValidateResult struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	ErrorType *string `json:"errorType,omitempty" xml:"errorType,omitempty"`
}

func (s UpgradeEngineVersionResponseBodyResultValidateResult) String() string {
	return tea.Prettify(s)
}

func (s UpgradeEngineVersionResponseBodyResultValidateResult) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) SetErrorCode(v string) *UpgradeEngineVersionResponseBodyResultValidateResult {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) SetErrorMsg(v string) *UpgradeEngineVersionResponseBodyResultValidateResult {
	s.ErrorMsg = &v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) SetErrorType(v string) *UpgradeEngineVersionResponseBodyResultValidateResult {
	s.ErrorType = &v
	return s
}

type UpgradeEngineVersionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeEngineVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeEngineVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeEngineVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionResponse) SetHeaders(v map[string]*string) *UpgradeEngineVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeEngineVersionResponse) SetBody(v *UpgradeEngineVersionResponseBody) *UpgradeEngineVersionResponse {
	s.Body = v
	return s
}

type ValidateConnectionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ValidateConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateConnectionRequest) GoString() string {
	return s.String()
}

func (s *ValidateConnectionRequest) SetClientToken(v string) *ValidateConnectionRequest {
	s.ClientToken = &v
	return s
}

type ValidateConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateConnectionResponseBody) SetRequestId(v string) *ValidateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateConnectionResponseBody) SetResult(v bool) *ValidateConnectionResponseBody {
	s.Result = &v
	return s
}

type ValidateConnectionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateConnectionResponse) GoString() string {
	return s.String()
}

func (s *ValidateConnectionResponse) SetHeaders(v map[string]*string) *ValidateConnectionResponse {
	s.Headers = v
	return s
}

func (s *ValidateConnectionResponse) SetBody(v *ValidateConnectionResponseBody) *ValidateConnectionResponse {
	s.Body = v
	return s
}

type ValidateShrinkNodesRequest struct {
	IgnoreStatus *bool   `json:"ignoreStatus,omitempty" xml:"ignoreStatus,omitempty"`
	NodeType     *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s ValidateShrinkNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateShrinkNodesRequest) GoString() string {
	return s.String()
}

func (s *ValidateShrinkNodesRequest) SetIgnoreStatus(v bool) *ValidateShrinkNodesRequest {
	s.IgnoreStatus = &v
	return s
}

func (s *ValidateShrinkNodesRequest) SetNodeType(v string) *ValidateShrinkNodesRequest {
	s.NodeType = &v
	return s
}

type ValidateShrinkNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateShrinkNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateShrinkNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateShrinkNodesResponseBody) SetRequestId(v string) *ValidateShrinkNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateShrinkNodesResponseBody) SetResult(v bool) *ValidateShrinkNodesResponseBody {
	s.Result = &v
	return s
}

type ValidateShrinkNodesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateShrinkNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateShrinkNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateShrinkNodesResponse) GoString() string {
	return s.String()
}

func (s *ValidateShrinkNodesResponse) SetHeaders(v map[string]*string) *ValidateShrinkNodesResponse {
	s.Headers = v
	return s
}

func (s *ValidateShrinkNodesResponse) SetBody(v *ValidateShrinkNodesResponseBody) *ValidateShrinkNodesResponse {
	s.Body = v
	return s
}

type ValidateSlrPermissionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Rolename    *string `json:"rolename,omitempty" xml:"rolename,omitempty"`
}

func (s ValidateSlrPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateSlrPermissionRequest) GoString() string {
	return s.String()
}

func (s *ValidateSlrPermissionRequest) SetClientToken(v string) *ValidateSlrPermissionRequest {
	s.ClientToken = &v
	return s
}

func (s *ValidateSlrPermissionRequest) SetRolename(v string) *ValidateSlrPermissionRequest {
	s.Rolename = &v
	return s
}

type ValidateSlrPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateSlrPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateSlrPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateSlrPermissionResponseBody) SetRequestId(v string) *ValidateSlrPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateSlrPermissionResponseBody) SetResult(v bool) *ValidateSlrPermissionResponseBody {
	s.Result = &v
	return s
}

type ValidateSlrPermissionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateSlrPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateSlrPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateSlrPermissionResponse) GoString() string {
	return s.String()
}

func (s *ValidateSlrPermissionResponse) SetHeaders(v map[string]*string) *ValidateSlrPermissionResponse {
	s.Headers = v
	return s
}

func (s *ValidateSlrPermissionResponse) SetBody(v *ValidateSlrPermissionResponseBody) *ValidateSlrPermissionResponse {
	s.Body = v
	return s
}

type ValidateTransferableNodesRequest struct {
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
}

func (s ValidateTransferableNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateTransferableNodesRequest) GoString() string {
	return s.String()
}

func (s *ValidateTransferableNodesRequest) SetNodeType(v string) *ValidateTransferableNodesRequest {
	s.NodeType = &v
	return s
}

type ValidateTransferableNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateTransferableNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateTransferableNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTransferableNodesResponseBody) SetRequestId(v string) *ValidateTransferableNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateTransferableNodesResponseBody) SetResult(v bool) *ValidateTransferableNodesResponseBody {
	s.Result = &v
	return s
}

type ValidateTransferableNodesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateTransferableNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateTransferableNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateTransferableNodesResponse) GoString() string {
	return s.String()
}

func (s *ValidateTransferableNodesResponse) SetHeaders(v map[string]*string) *ValidateTransferableNodesResponse {
	s.Headers = v
	return s
}

func (s *ValidateTransferableNodesResponse) SetBody(v *ValidateTransferableNodesResponseBody) *ValidateTransferableNodesResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

type CreateInstanceResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetResult(v *CreateInstanceResponseBodyResult) *CreateInstanceResponseBody {
	s.Result = v
	return s
}

type CreateInstanceResponseBodyResult struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s CreateInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyResult) SetInstanceId(v string) *CreateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("elasticsearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActivateZones(InstanceId *string, request *ActivateZonesRequest) (_result *ActivateZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActivateZonesResponse{}
	_body, _err := client.ActivateZonesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActivateZonesWithOptions(InstanceId *string, request *ActivateZonesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ActivateZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateZones"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/recover-zones"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddConnectableCluster(InstanceId *string, request *AddConnectableClusterRequest) (_result *AddConnectableClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddConnectableClusterResponse{}
	_body, _err := client.AddConnectableClusterWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddConnectableClusterWithOptions(InstanceId *string, request *AddConnectableClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddConnectableClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddConnectableCluster"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/connected-clusters"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddConnectableClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSnapshotRepo(InstanceId *string) (_result *AddSnapshotRepoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddSnapshotRepoResponse{}
	_body, _err := client.AddSnapshotRepoWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSnapshotRepoWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddSnapshotRepoResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AddSnapshotRepo"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/snapshot-repos"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSnapshotRepoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelDeletion(InstanceId *string, request *CancelDeletionRequest) (_result *CancelDeletionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelDeletionResponse{}
	_body, _err := client.CancelDeletionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelDeletionWithOptions(InstanceId *string, request *CancelDeletionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelDeletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDeletion"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/cancel-deletion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDeletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelLogstashDeletion(InstanceId *string, request *CancelLogstashDeletionRequest) (_result *CancelLogstashDeletionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelLogstashDeletionResponse{}
	_body, _err := client.CancelLogstashDeletionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelLogstashDeletionWithOptions(InstanceId *string, request *CancelLogstashDeletionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelLogstashDeletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelLogstashDeletion"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/actions/cancel-deletion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelLogstashDeletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelTask(InstanceId *string, request *CancelTaskRequest) (_result *CancelTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelTaskWithOptions(InstanceId *string, request *CancelTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["taskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelTask"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/cancel-task"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CapacityPlan(request *CapacityPlanRequest) (_result *CapacityPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CapacityPlanResponse{}
	_body, _err := client.CapacityPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CapacityPlanWithOptions(request *CapacityPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CapacityPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComplexQueryAvailable)) {
		body["complexQueryAvailable"] = request.ComplexQueryAvailable
	}

	if !tea.BoolValue(util.IsUnset(request.DataInfo)) {
		body["dataInfo"] = request.DataInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		body["metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.UsageScenario)) {
		body["usageScenario"] = request.UsageScenario
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CapacityPlan"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/assist/actions/capacity-plan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CapacityPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseDiagnosis(InstanceId *string, request *CloseDiagnosisRequest) (_result *CloseDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseDiagnosisResponse{}
	_body, _err := client.CloseDiagnosisWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseDiagnosisWithOptions(InstanceId *string, request *CloseDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloseDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseDiagnosis"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/diagnosis/instances/" + tea.StringValue(InstanceId) + "/actions/close-diagnosis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseHttps(InstanceId *string, request *CloseHttpsRequest) (_result *CloseHttpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseHttpsResponse{}
	_body, _err := client.CloseHttpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseHttpsWithOptions(InstanceId *string, request *CloseHttpsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloseHttpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseHttps"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/close-https"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseHttpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseManagedIndex(InstanceId *string, Index *string, request *CloseManagedIndexRequest) (_result *CloseManagedIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseManagedIndexResponse{}
	_body, _err := client.CloseManagedIndexWithOptions(InstanceId, Index, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseManagedIndexWithOptions(InstanceId *string, Index *string, request *CloseManagedIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloseManagedIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	Index = openapiutil.GetEncodeParam(Index)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseManagedIndex"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/indices/" + tea.StringValue(Index) + "/close-managed"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseManagedIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCollector(request *CreateCollectorRequest) (_result *CreateCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCollectorResponse{}
	_body, _err := client.CreateCollectorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCollectorWithOptions(request *CreateCollectorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCollector"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateComponentIndex(InstanceId *string, name *string, request *CreateComponentIndexRequest) (_result *CreateComponentIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateComponentIndexResponse{}
	_body, _err := client.CreateComponentIndexWithOptions(InstanceId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateComponentIndexWithOptions(InstanceId *string, name *string, request *CreateComponentIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateComponentIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	name = openapiutil.GetEncodeParam(name)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Meta)) {
		body["_meta"] = request.Meta
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Template))) {
		body["template"] = request.Template
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateComponentIndex"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/component-index/" + tea.StringValue(name)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateComponentIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataStream(InstanceId *string, request *CreateDataStreamRequest) (_result *CreateDataStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataStreamResponse{}
	_body, _err := client.CreateDataStreamWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataStreamWithOptions(InstanceId *string, request *CreateDataStreamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDataStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataStream"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/data-streams"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataTasks(InstanceId *string, request *CreateDataTasksRequest) (_result *CreateDataTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataTasksResponse{}
	_body, _err := client.CreateDataTasksWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataTasksWithOptions(InstanceId *string, request *CreateDataTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDataTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataTasks"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/data-task"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateILMPolicy(InstanceId *string, request *CreateILMPolicyRequest) (_result *CreateILMPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateILMPolicyResponse{}
	_body, _err := client.CreateILMPolicyWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateILMPolicyWithOptions(InstanceId *string, request *CreateILMPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateILMPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateILMPolicy"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/ilm-policies"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateILMPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIndexTemplate(InstanceId *string, request *CreateIndexTemplateRequest) (_result *CreateIndexTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndexTemplateResponse{}
	_body, _err := client.CreateIndexTemplateWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIndexTemplateWithOptions(InstanceId *string, request *CreateIndexTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIndexTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIndexTemplate"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/index-templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIndexTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLogstash(request *CreateLogstashRequest) (_result *CreateLogstashResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLogstashResponse{}
	_body, _err := client.CreateLogstashWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLogstashWithOptions(request *CreateLogstashRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLogstashResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogstash"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePipelines(InstanceId *string, request *CreatePipelinesRequest) (_result *CreatePipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePipelinesResponse{}
	_body, _err := client.CreatePipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePipelinesWithOptions(InstanceId *string, request *CreatePipelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Trigger)) {
		query["trigger"] = request.Trigger
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePipelines"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/pipelines"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(InstanceId *string, request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(InstanceId *string, request *CreateSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSnapshot"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/snapshots"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVpcEndpoint(InstanceId *string, request *CreateVpcEndpointRequest) (_result *CreateVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.CreateVpcEndpointWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVpcEndpointWithOptions(InstanceId *string, request *CreateVpcEndpointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["serviceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["zoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcEndpoint"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/vpc-endpoints"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeactivateZones(InstanceId *string, request *DeactivateZonesRequest) (_result *DeactivateZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeactivateZonesResponse{}
	_body, _err := client.DeactivateZonesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeactivateZonesWithOptions(InstanceId *string, request *DeactivateZonesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeactivateZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeactivateZones"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/down-zones"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeactivateZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCollector(ResId *string, request *DeleteCollectorRequest) (_result *DeleteCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCollectorResponse{}
	_body, _err := client.DeleteCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCollectorWithOptions(ResId *string, request *DeleteCollectorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ResId = openapiutil.GetEncodeParam(ResId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCollector"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteComponentIndex(InstanceId *string, name *string) (_result *DeleteComponentIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteComponentIndexResponse{}
	_body, _err := client.DeleteComponentIndexWithOptions(InstanceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteComponentIndexWithOptions(InstanceId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteComponentIndexResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteComponentIndex"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/component-index/" + tea.StringValue(name)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteComponentIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConnectedCluster(InstanceId *string, request *DeleteConnectedClusterRequest) (_result *DeleteConnectedClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConnectedClusterResponse{}
	_body, _err := client.DeleteConnectedClusterWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConnectedClusterWithOptions(InstanceId *string, request *DeleteConnectedClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConnectedClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectedInstanceId)) {
		query["connectedInstanceId"] = request.ConnectedInstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConnectedCluster"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/connected-clusters"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConnectedClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataStream(InstanceId *string, DataStream *string, request *DeleteDataStreamRequest) (_result *DeleteDataStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataStreamResponse{}
	_body, _err := client.DeleteDataStreamWithOptions(InstanceId, DataStream, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataStreamWithOptions(InstanceId *string, DataStream *string, request *DeleteDataStreamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDataStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	DataStream = openapiutil.GetEncodeParam(DataStream)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataStream"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/data-streams/" + tea.StringValue(DataStream)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataTask(InstanceId *string, request *DeleteDataTaskRequest) (_result *DeleteDataTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataTaskResponse{}
	_body, _err := client.DeleteDataTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataTaskWithOptions(InstanceId *string, request *DeleteDataTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDataTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataTask"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/data-task"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeprecatedTemplate(InstanceId *string, name *string) (_result *DeleteDeprecatedTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDeprecatedTemplateResponse{}
	_body, _err := client.DeleteDeprecatedTemplateWithOptions(InstanceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeprecatedTemplateWithOptions(InstanceId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDeprecatedTemplateResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDeprecatedTemplate"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/deprecated-templates/" + tea.StringValue(name)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeprecatedTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteILMPolicy(InstanceId *string, PolicyName *string) (_result *DeleteILMPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteILMPolicyResponse{}
	_body, _err := client.DeleteILMPolicyWithOptions(InstanceId, PolicyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteILMPolicyWithOptions(InstanceId *string, PolicyName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteILMPolicyResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	PolicyName = openapiutil.GetEncodeParam(PolicyName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteILMPolicy"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/ilm-policies/" + tea.StringValue(PolicyName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteILMPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIndexTemplate(InstanceId *string, IndexTemplate *string) (_result *DeleteIndexTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexTemplateResponse{}
	_body, _err := client.DeleteIndexTemplateWithOptions(InstanceId, IndexTemplate, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIndexTemplateWithOptions(InstanceId *string, IndexTemplate *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndexTemplateResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	IndexTemplate = openapiutil.GetEncodeParam(IndexTemplate)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndexTemplate"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/index-templates/" + tea.StringValue(IndexTemplate)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIndexTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(InstanceId *string, request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(InstanceId *string, request *DeleteInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteType)) {
		query["deleteType"] = request.DeleteType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLogstash(InstanceId *string, request *DeleteLogstashRequest) (_result *DeleteLogstashResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLogstashResponse{}
	_body, _err := client.DeleteLogstashWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLogstashWithOptions(InstanceId *string, request *DeleteLogstashRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLogstashResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteType)) {
		query["deleteType"] = request.DeleteType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLogstash"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePipelines(InstanceId *string, request *DeletePipelinesRequest) (_result *DeletePipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelinesResponse{}
	_body, _err := client.DeletePipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePipelinesWithOptions(InstanceId *string, request *DeletePipelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineIds)) {
		query["pipelineIds"] = request.PipelineIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipelines"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/pipelines"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshotRepo(InstanceId *string, request *DeleteSnapshotRepoRequest) (_result *DeleteSnapshotRepoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSnapshotRepoResponse{}
	_body, _err := client.DeleteSnapshotRepoWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotRepoWithOptions(InstanceId *string, request *DeleteSnapshotRepoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSnapshotRepoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RepoPath)) {
		query["repoPath"] = request.RepoPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshotRepo"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/snapshot-repos"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSnapshotRepoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVpcEndpoint(InstanceId *string, EndpointId *string, request *DeleteVpcEndpointRequest) (_result *DeleteVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.DeleteVpcEndpointWithOptions(InstanceId, EndpointId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVpcEndpointWithOptions(InstanceId *string, EndpointId *string, request *DeleteVpcEndpointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	EndpointId = openapiutil.GetEncodeParam(EndpointId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcEndpoint"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/vpc-endpoints/" + tea.StringValue(EndpointId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAckOperator(ClusterId *string) (_result *DescribeAckOperatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAckOperatorResponse{}
	_body, _err := client.DescribeAckOperatorWithOptions(ClusterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAckOperatorWithOptions(ClusterId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAckOperatorResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAckOperator"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ack-clusters/" + tea.StringValue(ClusterId) + "/operator"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAckOperatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApm(instanceId *string) (_result *DescribeApmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApmResponse{}
	_body, _err := client.DescribeApmWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApmWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApmResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApm"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/apm/" + tea.StringValue(instanceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCollector(ResId *string) (_result *DescribeCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCollectorResponse{}
	_body, _err := client.DescribeCollectorWithOptions(ResId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCollectorWithOptions(ResId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeCollectorResponse, _err error) {
	ResId = openapiutil.GetEncodeParam(ResId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCollector"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComponentIndex(InstanceId *string, name *string) (_result *DescribeComponentIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeComponentIndexResponse{}
	_body, _err := client.DescribeComponentIndexWithOptions(InstanceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComponentIndexWithOptions(InstanceId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeComponentIndexResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentIndex"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/component-index/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComponentIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConnectableClusters(InstanceId *string, request *DescribeConnectableClustersRequest) (_result *DescribeConnectableClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConnectableClustersResponse{}
	_body, _err := client.DescribeConnectableClustersWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConnectableClustersWithOptions(InstanceId *string, request *DescribeConnectableClustersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeConnectableClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlreadySetItems)) {
		query["alreadySetItems"] = request.AlreadySetItems
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConnectableClusters"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/connectable-clusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConnectableClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeprecatedTemplate(InstanceId *string, name *string) (_result *DescribeDeprecatedTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDeprecatedTemplateResponse{}
	_body, _err := client.DescribeDeprecatedTemplateWithOptions(InstanceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeprecatedTemplateWithOptions(InstanceId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeDeprecatedTemplateResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	name = openapiutil.GetEncodeParam(name)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeprecatedTemplate"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/deprecated-templates/" + tea.StringValue(name)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeprecatedTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnoseReport(InstanceId *string, ReportId *string, request *DescribeDiagnoseReportRequest) (_result *DescribeDiagnoseReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDiagnoseReportResponse{}
	_body, _err := client.DescribeDiagnoseReportWithOptions(InstanceId, ReportId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnoseReportWithOptions(InstanceId *string, ReportId *string, request *DescribeDiagnoseReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeDiagnoseReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	ReportId = openapiutil.GetEncodeParam(ReportId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnoseReport"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/diagnosis/instances/" + tea.StringValue(InstanceId) + "/reports/" + tea.StringValue(ReportId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnoseReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisSettings(InstanceId *string, request *DescribeDiagnosisSettingsRequest) (_result *DescribeDiagnosisSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDiagnosisSettingsResponse{}
	_body, _err := client.DescribeDiagnosisSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisSettingsWithOptions(InstanceId *string, request *DescribeDiagnosisSettingsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisSettings"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/diagnosis/instances/" + tea.StringValue(InstanceId) + "/settings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticsearchHealth(InstanceId *string) (_result *DescribeElasticsearchHealthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeElasticsearchHealthResponse{}
	_body, _err := client.DescribeElasticsearchHealthWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticsearchHealthWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeElasticsearchHealthResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticsearchHealth"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/elasticsearch-health"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeElasticsearchHealthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeILMPolicy(InstanceId *string, PolicyName *string) (_result *DescribeILMPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeILMPolicyResponse{}
	_body, _err := client.DescribeILMPolicyWithOptions(InstanceId, PolicyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeILMPolicyWithOptions(InstanceId *string, PolicyName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeILMPolicyResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	PolicyName = openapiutil.GetEncodeParam(PolicyName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeILMPolicy"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/ilm-policies/" + tea.StringValue(PolicyName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeILMPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIndexTemplate(InstanceId *string, IndexTemplate *string) (_result *DescribeIndexTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeIndexTemplateResponse{}
	_body, _err := client.DescribeIndexTemplateWithOptions(InstanceId, IndexTemplate, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIndexTemplateWithOptions(InstanceId *string, IndexTemplate *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeIndexTemplateResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	IndexTemplate = openapiutil.GetEncodeParam(IndexTemplate)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIndexTemplate"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/index-templates/" + tea.StringValue(IndexTemplate)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIndexTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(InstanceId *string) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKibanaSettings(InstanceId *string) (_result *DescribeKibanaSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeKibanaSettingsResponse{}
	_body, _err := client.DescribeKibanaSettingsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKibanaSettingsWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeKibanaSettingsResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeKibanaSettings"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/kibana-settings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeKibanaSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogstash(InstanceId *string) (_result *DescribeLogstashResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeLogstashResponse{}
	_body, _err := client.DescribeLogstashWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogstashWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeLogstashResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogstash"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePipeline(InstanceId *string, PipelineId *string) (_result *DescribePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePipelineResponse{}
	_body, _err := client.DescribePipelineWithOptions(InstanceId, PipelineId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePipelineWithOptions(InstanceId *string, PipelineId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribePipelineResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	PipelineId = openapiutil.GetEncodeParam(PipelineId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePipeline"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/pipelines/" + tea.StringValue(PipelineId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePipelineManagementConfig(InstanceId *string, request *DescribePipelineManagementConfigRequest) (_result *DescribePipelineManagementConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePipelineManagementConfigResponse{}
	_body, _err := client.DescribePipelineManagementConfigWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePipelineManagementConfigWithOptions(InstanceId *string, request *DescribePipelineManagementConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribePipelineManagementConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePipelineManagementConfig"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/pipeline-management-config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePipelineManagementConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/regions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSnapshotSetting(InstanceId *string) (_result *DescribeSnapshotSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSnapshotSettingResponse{}
	_body, _err := client.DescribeSnapshotSettingWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSnapshotSettingWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSnapshotSettingResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSnapshotSetting"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/snapshot-setting"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSnapshotSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplates(InstanceId *string) (_result *DescribeTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DescribeTemplatesWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplatesWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTemplates"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeXpackMonitorConfig(InstanceId *string) (_result *DescribeXpackMonitorConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeXpackMonitorConfigResponse{}
	_body, _err := client.DescribeXpackMonitorConfigWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeXpackMonitorConfigWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeXpackMonitorConfigResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeXpackMonitorConfig"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/xpack-monitor-config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeXpackMonitorConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DiagnoseInstance(InstanceId *string, request *DiagnoseInstanceRequest) (_result *DiagnoseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DiagnoseInstanceResponse{}
	_body, _err := client.DiagnoseInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DiagnoseInstanceWithOptions(InstanceId *string, request *DiagnoseInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DiagnoseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DiagnoseInstance"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/diagnosis/instances/" + tea.StringValue(InstanceId) + "/actions/diagnose"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DiagnoseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EstimatedLogstashRestartTime(InstanceId *string, request *EstimatedLogstashRestartTimeRequest) (_result *EstimatedLogstashRestartTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EstimatedLogstashRestartTimeResponse{}
	_body, _err := client.EstimatedLogstashRestartTimeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EstimatedLogstashRestartTimeWithOptions(InstanceId *string, request *EstimatedLogstashRestartTimeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EstimatedLogstashRestartTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EstimatedLogstashRestartTime"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/estimated-time/restart-time"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EstimatedLogstashRestartTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EstimatedRestartTime(InstanceId *string, request *EstimatedRestartTimeRequest) (_result *EstimatedRestartTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EstimatedRestartTimeResponse{}
	_body, _err := client.EstimatedRestartTimeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EstimatedRestartTimeWithOptions(InstanceId *string, request *EstimatedRestartTimeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EstimatedRestartTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EstimatedRestartTime"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/estimated-time/restart-time"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EstimatedRestartTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterDataInformation() (_result *GetClusterDataInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterDataInformationResponse{}
	_body, _err := client.GetClusterDataInformationWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterDataInformationWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterDataInformationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterDataInformation"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/cluster/data-information"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterDataInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetElastictask(InstanceId *string) (_result *GetElastictaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetElastictaskResponse{}
	_body, _err := client.GetElastictaskWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetElastictaskWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetElastictaskResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetElastictask"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/elastic-task"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetElastictaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEmonGrafanaAlerts(ProjectId *string) (_result *GetEmonGrafanaAlertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmonGrafanaAlertsResponse{}
	_body, _err := client.GetEmonGrafanaAlertsWithOptions(ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEmonGrafanaAlertsWithOptions(ProjectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEmonGrafanaAlertsResponse, _err error) {
	ProjectId = openapiutil.GetEncodeParam(ProjectId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEmonGrafanaAlerts"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/emon/projects/" + tea.StringValue(ProjectId) + "/grafana/proxy/api/alerts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmonGrafanaAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEmonGrafanaDashboards(ProjectId *string) (_result *GetEmonGrafanaDashboardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmonGrafanaDashboardsResponse{}
	_body, _err := client.GetEmonGrafanaDashboardsWithOptions(ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEmonGrafanaDashboardsWithOptions(ProjectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEmonGrafanaDashboardsResponse, _err error) {
	ProjectId = openapiutil.GetEncodeParam(ProjectId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEmonGrafanaDashboards"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/emon/projects/" + tea.StringValue(ProjectId) + "/grafana/proxy/api/search"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmonGrafanaDashboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEmonMonitorData(ProjectId *string) (_result *GetEmonMonitorDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmonMonitorDataResponse{}
	_body, _err := client.GetEmonMonitorDataWithOptions(ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEmonMonitorDataWithOptions(ProjectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEmonMonitorDataResponse, _err error) {
	ProjectId = openapiutil.GetEncodeParam(ProjectId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEmonMonitorData"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/emon/projects/" + tea.StringValue(ProjectId) + "/metrics/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEmonMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenStoreUsage(InstanceId *string) (_result *GetOpenStoreUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOpenStoreUsageResponse{}
	_body, _err := client.GetOpenStoreUsageWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenStoreUsageWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOpenStoreUsageResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOpenStoreUsage"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/openstore/usage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOpenStoreUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRegionConfiguration(request *GetRegionConfigurationRequest) (_result *GetRegionConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegionConfigurationResponse{}
	_body, _err := client.GetRegionConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRegionConfigurationWithOptions(request *GetRegionConfigurationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRegionConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["zoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegionConfiguration"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/region"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegionConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSuggestShrinkableNodes(InstanceId *string, request *GetSuggestShrinkableNodesRequest) (_result *GetSuggestShrinkableNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSuggestShrinkableNodesResponse{}
	_body, _err := client.GetSuggestShrinkableNodesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSuggestShrinkableNodesWithOptions(InstanceId *string, request *GetSuggestShrinkableNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSuggestShrinkableNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreStatus)) {
		query["ignoreStatus"] = request.IgnoreStatus
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["nodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSuggestShrinkableNodes"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/suggest-shrinkable-nodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSuggestShrinkableNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTransferableNodes(InstanceId *string, request *GetTransferableNodesRequest) (_result *GetTransferableNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTransferableNodesResponse{}
	_body, _err := client.GetTransferableNodesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTransferableNodesWithOptions(InstanceId *string, request *GetTransferableNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTransferableNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["nodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTransferableNodes"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/transferable-nodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTransferableNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitializeOperationRole(request *InitializeOperationRoleRequest) (_result *InitializeOperationRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitializeOperationRoleResponse{}
	_body, _err := client.InitializeOperationRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitializeOperationRoleWithOptions(request *InitializeOperationRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InitializeOperationRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InitializeOperationRole"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/user/slr"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeOperationRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallAckOperator(ClusterId *string, request *InstallAckOperatorRequest) (_result *InstallAckOperatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallAckOperatorResponse{}
	_body, _err := client.InstallAckOperatorWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallAckOperatorWithOptions(ClusterId *string, request *InstallAckOperatorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InstallAckOperatorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallAckOperator"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ack-clusters/" + tea.StringValue(ClusterId) + "/operator"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallAckOperatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallKibanaSystemPlugin(InstanceId *string, request *InstallKibanaSystemPluginRequest) (_result *InstallKibanaSystemPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallKibanaSystemPluginResponse{}
	_body, _err := client.InstallKibanaSystemPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallKibanaSystemPluginWithOptions(InstanceId *string, request *InstallKibanaSystemPluginRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InstallKibanaSystemPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallKibanaSystemPlugin"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/kibana-plugins/system/actions/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallKibanaSystemPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallLogstashSystemPlugin(InstanceId *string, request *InstallLogstashSystemPluginRequest) (_result *InstallLogstashSystemPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallLogstashSystemPluginResponse{}
	_body, _err := client.InstallLogstashSystemPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallLogstashSystemPluginWithOptions(InstanceId *string, request *InstallLogstashSystemPluginRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InstallLogstashSystemPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallLogstashSystemPlugin"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/plugins/system/actions/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallLogstashSystemPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallSystemPlugin(InstanceId *string, request *InstallSystemPluginRequest) (_result *InstallSystemPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallSystemPluginResponse{}
	_body, _err := client.InstallSystemPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallSystemPluginWithOptions(InstanceId *string, request *InstallSystemPluginRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InstallSystemPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallSystemPlugin"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/plugins/system/actions/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallSystemPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallUserPlugins(InstanceId *string) (_result *InstallUserPluginsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallUserPluginsResponse{}
	_body, _err := client.InstallUserPluginsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallUserPluginsWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InstallUserPluginsResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("InstallUserPlugins"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/plugins/user/actions/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallUserPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InterruptElasticsearchTask(InstanceId *string, request *InterruptElasticsearchTaskRequest) (_result *InterruptElasticsearchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InterruptElasticsearchTaskResponse{}
	_body, _err := client.InterruptElasticsearchTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InterruptElasticsearchTaskWithOptions(InstanceId *string, request *InterruptElasticsearchTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InterruptElasticsearchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InterruptElasticsearchTask"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/interrupt"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InterruptElasticsearchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InterruptLogstashTask(InstanceId *string, request *InterruptLogstashTaskRequest) (_result *InterruptLogstashTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InterruptLogstashTaskResponse{}
	_body, _err := client.InterruptLogstashTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InterruptLogstashTaskWithOptions(InstanceId *string, request *InterruptLogstashTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InterruptLogstashTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InterruptLogstashTask"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/actions/interrupt"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InterruptLogstashTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAckClusters(request *ListAckClustersRequest) (_result *ListAckClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAckClustersResponse{}
	_body, _err := client.ListAckClustersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAckClustersWithOptions(request *ListAckClustersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAckClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["vpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAckClusters"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ack-clusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAckClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAckNamespaces(ClusterId *string, request *ListAckNamespacesRequest) (_result *ListAckNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAckNamespacesResponse{}
	_body, _err := client.ListAckNamespacesWithOptions(ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAckNamespacesWithOptions(ClusterId *string, request *ListAckNamespacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAckNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAckNamespaces"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ack-clusters/" + tea.StringValue(ClusterId) + "/namespaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAckNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAllNode(InstanceId *string, request *ListAllNodeRequest) (_result *ListAllNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAllNodeResponse{}
	_body, _err := client.ListAllNodeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAllNodeWithOptions(InstanceId *string, request *ListAllNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAllNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extended)) {
		query["extended"] = request.Extended
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAllNode"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/nodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlternativeSnapshotRepos(InstanceId *string, request *ListAlternativeSnapshotReposRequest) (_result *ListAlternativeSnapshotReposResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlternativeSnapshotReposResponse{}
	_body, _err := client.ListAlternativeSnapshotReposWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlternativeSnapshotReposWithOptions(InstanceId *string, request *ListAlternativeSnapshotReposRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlternativeSnapshotReposResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlreadySetItems)) {
		query["alreadySetItems"] = request.AlreadySetItems
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlternativeSnapshotRepos"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/alternative-snapshot-repos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlternativeSnapshotReposResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApm(request *ListApmRequest) (_result *ListApmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApmResponse{}
	_body, _err := client.ListApmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApmWithOptions(request *ListApmRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListApmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		query["output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApm"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/apm"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableEsInstanceIds(InstanceId *string) (_result *ListAvailableEsInstanceIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAvailableEsInstanceIdsResponse{}
	_body, _err := client.ListAvailableEsInstanceIdsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableEsInstanceIdsWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAvailableEsInstanceIdsResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableEsInstanceIds"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/available-elasticsearch-for-centralized-management"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableEsInstanceIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCollectors(request *ListCollectorsRequest) (_result *ListCollectorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCollectorsResponse{}
	_body, _err := client.ListCollectorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCollectorsWithOptions(request *ListCollectorsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCollectorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.ResId)) {
		query["resId"] = request.ResId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCollectors"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCollectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListComponentIndices(InstanceId *string, request *ListComponentIndicesRequest) (_result *ListComponentIndicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComponentIndicesResponse{}
	_body, _err := client.ListComponentIndicesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListComponentIndicesWithOptions(InstanceId *string, request *ListComponentIndicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListComponentIndicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListComponentIndices"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/component-index"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListComponentIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConnectedClusters(InstanceId *string) (_result *ListConnectedClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConnectedClustersResponse{}
	_body, _err := client.ListConnectedClustersWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConnectedClustersWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConnectedClustersResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListConnectedClusters"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/connected-clusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConnectedClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataStreams(InstanceId *string, request *ListDataStreamsRequest) (_result *ListDataStreamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataStreamsResponse{}
	_body, _err := client.ListDataStreamsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataStreamsWithOptions(InstanceId *string, request *ListDataStreamsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataStreamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsManaged)) {
		query["isManaged"] = request.IsManaged
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataStreams"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/data-streams"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataStreamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataTasks(InstanceId *string) (_result *ListDataTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataTasksResponse{}
	_body, _err := client.ListDataTasksWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataTasksWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataTasksResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataTasks"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/data-task"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDefaultCollectorConfigurations(request *ListDefaultCollectorConfigurationsRequest) (_result *ListDefaultCollectorConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDefaultCollectorConfigurationsResponse{}
	_body, _err := client.ListDefaultCollectorConfigurationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDefaultCollectorConfigurationsWithOptions(request *ListDefaultCollectorConfigurationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDefaultCollectorConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["resType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResVersion)) {
		query["resVersion"] = request.ResVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDefaultCollectorConfigurations"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/beats/default-configurations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDefaultCollectorConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeprecatedTemplates(InstanceId *string, request *ListDeprecatedTemplatesRequest) (_result *ListDeprecatedTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDeprecatedTemplatesResponse{}
	_body, _err := client.ListDeprecatedTemplatesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeprecatedTemplatesWithOptions(InstanceId *string, request *ListDeprecatedTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDeprecatedTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeprecatedTemplates"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/deprecated-templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeprecatedTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDiagnoseIndices(InstanceId *string, request *ListDiagnoseIndicesRequest) (_result *ListDiagnoseIndicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiagnoseIndicesResponse{}
	_body, _err := client.ListDiagnoseIndicesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDiagnoseIndicesWithOptions(InstanceId *string, request *ListDiagnoseIndicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDiagnoseIndicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiagnoseIndices"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/diagnosis/instances/" + tea.StringValue(InstanceId) + "/indices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDiagnoseIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDiagnoseReport(InstanceId *string, request *ListDiagnoseReportRequest) (_result *ListDiagnoseReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiagnoseReportResponse{}
	_body, _err := client.ListDiagnoseReportWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDiagnoseReportWithOptions(InstanceId *string, request *ListDiagnoseReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDiagnoseReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		query["detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Trigger)) {
		query["trigger"] = request.Trigger
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiagnoseReport"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/diagnosis/instances/" + tea.StringValue(InstanceId) + "/reports"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDiagnoseReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDiagnoseReportIds(InstanceId *string, request *ListDiagnoseReportIdsRequest) (_result *ListDiagnoseReportIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiagnoseReportIdsResponse{}
	_body, _err := client.ListDiagnoseReportIdsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDiagnoseReportIdsWithOptions(InstanceId *string, request *ListDiagnoseReportIdsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDiagnoseReportIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Trigger)) {
		query["trigger"] = request.Trigger
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiagnoseReportIds"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/diagnosis/instances/" + tea.StringValue(InstanceId) + "/report-ids"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDiagnoseReportIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDictInformation(InstanceId *string, request *ListDictInformationRequest) (_result *ListDictInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDictInformationResponse{}
	_body, _err := client.ListDictInformationWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDictInformationWithOptions(InstanceId *string, request *ListDictInformationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDictInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnalyzerType)) {
		query["analyzerType"] = request.AnalyzerType
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["bucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["key"] = request.Key
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDictInformation"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/dict/_info"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDictInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDicts(InstanceId *string, request *ListDictsRequest) (_result *ListDictsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDictsResponse{}
	_body, _err := client.ListDictsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDictsWithOptions(InstanceId *string, request *ListDictsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDictsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnalyzerType)) {
		query["analyzerType"] = request.AnalyzerType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDicts"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/dicts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDictsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEcsInstances(request *ListEcsInstancesRequest) (_result *ListEcsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsInstancesResponse{}
	_body, _err := client.ListEcsInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEcsInstancesWithOptions(request *ListEcsInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEcsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcsInstanceIds)) {
		query["ecsInstanceIds"] = request.EcsInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceName)) {
		query["ecsInstanceName"] = request.EcsInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["vpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEcsInstances"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/ecs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEcsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExtendfiles(InstanceId *string) (_result *ListExtendfilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExtendfilesResponse{}
	_body, _err := client.ListExtendfilesWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExtendfilesWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListExtendfilesResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListExtendfiles"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/extendfiles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExtendfilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListILMPolicies(InstanceId *string, request *ListILMPoliciesRequest) (_result *ListILMPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListILMPoliciesResponse{}
	_body, _err := client.ListILMPoliciesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListILMPoliciesWithOptions(InstanceId *string, request *ListILMPoliciesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListILMPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["policyName"] = request.PolicyName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListILMPolicies"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/ilm-policies"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListILMPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIndexTemplates(InstanceId *string, request *ListIndexTemplatesRequest) (_result *ListIndexTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndexTemplatesResponse{}
	_body, _err := client.ListIndexTemplatesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIndexTemplatesWithOptions(InstanceId *string, request *ListIndexTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIndexTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndexTemplate)) {
		query["indexTemplate"] = request.IndexTemplate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIndexTemplates"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/index-templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIndexTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EsVersion)) {
		query["esVersion"] = request.EsVersion
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceCategory)) {
		query["instanceCategory"] = request.InstanceCategory
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["paymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["vpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["zoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstance"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceIndices(InstanceId *string, request *ListInstanceIndicesRequest) (_result *ListInstanceIndicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceIndicesResponse{}
	_body, _err := client.ListInstanceIndicesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceIndicesWithOptions(InstanceId *string, request *ListInstanceIndicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceIndicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["all"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.IsManaged)) {
		query["isManaged"] = request.IsManaged
	}

	if !tea.BoolValue(util.IsUnset(request.IsOpenstore)) {
		query["isOpenstore"] = request.IsOpenstore
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceIndices"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/indices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListKibanaPlugins(InstanceId *string, request *ListKibanaPluginsRequest) (_result *ListKibanaPluginsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKibanaPluginsResponse{}
	_body, _err := client.ListKibanaPluginsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListKibanaPluginsWithOptions(InstanceId *string, request *ListKibanaPluginsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListKibanaPluginsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKibanaPlugins"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/kibana-plugins"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKibanaPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogstash(request *ListLogstashRequest) (_result *ListLogstashResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogstashResponse{}
	_body, _err := client.ListLogstashWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogstashWithOptions(request *ListLogstashRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogstashResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["ownerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLogstash"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogstashLog(InstanceId *string, request *ListLogstashLogRequest) (_result *ListLogstashLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogstashLogResponse{}
	_body, _err := client.ListLogstashLogWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogstashLogWithOptions(InstanceId *string, request *ListLogstashLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogstashLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["beginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLogstashLog"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/search-log"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLogstashLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogstashPlugins(InstanceId *string, request *ListLogstashPluginsRequest) (_result *ListLogstashPluginsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogstashPluginsResponse{}
	_body, _err := client.ListLogstashPluginsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogstashPluginsWithOptions(InstanceId *string, request *ListLogstashPluginsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogstashPluginsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLogstashPlugins"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/plugins"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLogstashPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodes(ResId *string, request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesWithOptions(ResId *string, request *ListNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ResId = openapiutil.GetEncodeParam(ResId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcsInstanceIds)) {
		query["ecsInstanceIds"] = request.EcsInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceName)) {
		query["ecsInstanceName"] = request.EcsInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId) + "/nodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipeline(InstanceId *string, request *ListPipelineRequest) (_result *ListPipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineResponse{}
	_body, _err := client.ListPipelineWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineWithOptions(InstanceId *string, request *ListPipelineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["pipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipeline"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/pipelines"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineIds(InstanceId *string) (_result *ListPipelineIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineIdsResponse{}
	_body, _err := client.ListPipelineIdsWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineIdsWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineIdsResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineIds"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/pipeline-ids"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPlugins(InstanceId *string, request *ListPluginsRequest) (_result *ListPluginsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPluginsResponse{}
	_body, _err := client.ListPluginsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPluginsWithOptions(InstanceId *string, request *ListPluginsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPluginsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPlugins"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/plugins"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSearchLog(InstanceId *string, request *ListSearchLogRequest) (_result *ListSearchLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSearchLogResponse{}
	_body, _err := client.ListSearchLogWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSearchLogWithOptions(InstanceId *string, request *ListSearchLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSearchLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["beginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSearchLog"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/search-log"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSearchLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListShardRecoveries(InstanceId *string, request *ListShardRecoveriesRequest) (_result *ListShardRecoveriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListShardRecoveriesResponse{}
	_body, _err := client.ListShardRecoveriesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListShardRecoveriesWithOptions(InstanceId *string, request *ListShardRecoveriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListShardRecoveriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActiveOnly)) {
		query["activeOnly"] = request.ActiveOnly
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListShardRecoveries"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/cat-recovery"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListShardRecoveriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSnapshotReposByInstanceId(InstanceId *string) (_result *ListSnapshotReposByInstanceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSnapshotReposByInstanceIdResponse{}
	_body, _err := client.ListSnapshotReposByInstanceIdWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSnapshotReposByInstanceIdWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSnapshotReposByInstanceIdResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSnapshotReposByInstanceId"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/snapshot-repos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSnapshotReposByInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagsWithOptions(request *ListTagsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTags"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tags/all-tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpoints(InstanceId *string, request *ListVpcEndpointsRequest) (_result *ListVpcEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.ListVpcEndpointsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointsWithOptions(InstanceId *string, request *ListVpcEndpointsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListVpcEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcEndpoints"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/vpc-endpoints"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateToOtherZone(InstanceId *string, request *MigrateToOtherZoneRequest) (_result *MigrateToOtherZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MigrateToOtherZoneResponse{}
	_body, _err := client.MigrateToOtherZoneWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateToOtherZoneWithOptions(InstanceId *string, request *MigrateToOtherZoneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MigrateToOtherZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MigrateToOtherZone"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/migrate-zones"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &MigrateToOtherZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDeployMachine(ResId *string, request *ModifyDeployMachineRequest) (_result *ModifyDeployMachineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDeployMachineResponse{}
	_body, _err := client.ModifyDeployMachineWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDeployMachineWithOptions(ResId *string, request *ModifyDeployMachineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyDeployMachineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ResId = openapiutil.GetEncodeParam(ResId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDeployMachine"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId) + "/actions/modify-deploy-machines"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDeployMachineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyElastictask(InstanceId *string) (_result *ModifyElastictaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyElastictaskResponse{}
	_body, _err := client.ModifyElastictaskWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyElastictaskWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyElastictaskResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyElastictask"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/elastic-task"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyElastictaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceMaintainTime(InstanceId *string, request *ModifyInstanceMaintainTimeRequest) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyInstanceMaintainTimeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceMaintainTimeWithOptions(InstanceId *string, request *ModifyInstanceMaintainTimeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceMaintainTime"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/modify-maintaintime"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWhiteIps(InstanceId *string, request *ModifyWhiteIpsRequest) (_result *ModifyWhiteIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyWhiteIpsResponse{}
	_body, _err := client.ModifyWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWhiteIpsWithOptions(InstanceId *string, request *ModifyWhiteIpsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyWhiteIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		body["modifyMode"] = request.ModifyMode
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		body["networkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		body["nodeType"] = request.NodeType
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.WhiteIpGroup))) {
		body["whiteIpGroup"] = request.WhiteIpGroup
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteIpList)) {
		body["whiteIpList"] = request.WhiteIpList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyWhiteIps"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/modify-white-ips"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(InstanceId *string, request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(InstanceId *string, request *MoveResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/resourcegroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenDiagnosis(InstanceId *string, request *OpenDiagnosisRequest) (_result *OpenDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenDiagnosisResponse{}
	_body, _err := client.OpenDiagnosisWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenDiagnosisWithOptions(InstanceId *string, request *OpenDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OpenDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenDiagnosis"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/diagnosis/instances/" + tea.StringValue(InstanceId) + "/actions/open-diagnosis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenHttps(InstanceId *string, request *OpenHttpsRequest) (_result *OpenHttpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenHttpsResponse{}
	_body, _err := client.OpenHttpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenHttpsWithOptions(InstanceId *string, request *OpenHttpsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OpenHttpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenHttps"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/open-https"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenHttpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostEmonTryAlarmRule(ProjectId *string, AlarmGroupId *string) (_result *PostEmonTryAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PostEmonTryAlarmRuleResponse{}
	_body, _err := client.PostEmonTryAlarmRuleWithOptions(ProjectId, AlarmGroupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostEmonTryAlarmRuleWithOptions(ProjectId *string, AlarmGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PostEmonTryAlarmRuleResponse, _err error) {
	ProjectId = openapiutil.GetEncodeParam(ProjectId)
	AlarmGroupId = openapiutil.GetEncodeParam(AlarmGroupId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PostEmonTryAlarmRule"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/emon/projects/" + tea.StringValue(ProjectId) + "/alarm-groups/" + tea.StringValue(AlarmGroupId) + "/alarm-rules/_test"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PostEmonTryAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecommendTemplates(InstanceId *string, request *RecommendTemplatesRequest) (_result *RecommendTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecommendTemplatesResponse{}
	_body, _err := client.RecommendTemplatesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecommendTemplatesWithOptions(InstanceId *string, request *RecommendTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecommendTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UsageScenario)) {
		query["usageScenario"] = request.UsageScenario
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecommendTemplates"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/recommended-templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RecommendTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReinstallCollector(ResId *string, request *ReinstallCollectorRequest) (_result *ReinstallCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReinstallCollectorResponse{}
	_body, _err := client.ReinstallCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReinstallCollectorWithOptions(ResId *string, request *ReinstallCollectorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReinstallCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ResId = openapiutil.GetEncodeParam(ResId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReinstallCollector"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId) + "/actions/reinstall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReinstallCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveApm(instanceId *string) (_result *RemoveApmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveApmResponse{}
	_body, _err := client.RemoveApmWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveApmWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveApmResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveApm"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/apm/" + tea.StringValue(instanceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveApmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstance(InstanceId *string, request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewInstanceWithOptions(InstanceId *string, request *RenewInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstance"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/renew"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewLogstash(InstanceId *string, request *RenewLogstashRequest) (_result *RenewLogstashResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewLogstashResponse{}
	_body, _err := client.RenewLogstashWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewLogstashWithOptions(InstanceId *string, request *RenewLogstashRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenewLogstashResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewLogstash"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/actions/renew"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartCollector(ResId *string, request *RestartCollectorRequest) (_result *RestartCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartCollectorResponse{}
	_body, _err := client.RestartCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartCollectorWithOptions(ResId *string, request *RestartCollectorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ResId = openapiutil.GetEncodeParam(ResId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartCollector"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId) + "/actions/restart"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartInstance(InstanceId *string, request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartInstanceWithOptions(InstanceId *string, request *RestartInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartInstance"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/restart"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartLogstash(InstanceId *string, request *RestartLogstashRequest) (_result *RestartLogstashResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartLogstashResponse{}
	_body, _err := client.RestartLogstashWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartLogstashWithOptions(InstanceId *string, request *RestartLogstashRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartLogstashResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartLogstash"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/actions/restart"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeElasticsearchTask(InstanceId *string, request *ResumeElasticsearchTaskRequest) (_result *ResumeElasticsearchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeElasticsearchTaskResponse{}
	_body, _err := client.ResumeElasticsearchTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeElasticsearchTaskWithOptions(InstanceId *string, request *ResumeElasticsearchTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeElasticsearchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeElasticsearchTask"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/resume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeElasticsearchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeLogstashTask(InstanceId *string, request *ResumeLogstashTaskRequest) (_result *ResumeLogstashTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeLogstashTaskResponse{}
	_body, _err := client.ResumeLogstashTaskWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeLogstashTaskWithOptions(InstanceId *string, request *ResumeLogstashTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeLogstashTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeLogstashTask"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/actions/resume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeLogstashTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RolloverDataStream(InstanceId *string, DataStream *string, request *RolloverDataStreamRequest) (_result *RolloverDataStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RolloverDataStreamResponse{}
	_body, _err := client.RolloverDataStreamWithOptions(InstanceId, DataStream, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RolloverDataStreamWithOptions(InstanceId *string, DataStream *string, request *RolloverDataStreamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RolloverDataStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	DataStream = openapiutil.GetEncodeParam(DataStream)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RolloverDataStream"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/data-streams/" + tea.StringValue(DataStream) + "/rollover"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RolloverDataStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunPipelines(InstanceId *string, request *RunPipelinesRequest) (_result *RunPipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunPipelinesResponse{}
	_body, _err := client.RunPipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunPipelinesWithOptions(InstanceId *string, request *RunPipelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunPipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunPipelines"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/pipelines/action/run"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RunPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ShrinkNode(InstanceId *string, request *ShrinkNodeRequest) (_result *ShrinkNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ShrinkNodeResponse{}
	_body, _err := client.ShrinkNodeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ShrinkNodeWithOptions(InstanceId *string, request *ShrinkNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ShrinkNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreStatus)) {
		query["ignoreStatus"] = request.IgnoreStatus
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["nodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ShrinkNode"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/shrink"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ShrinkNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartApm(instanceId *string) (_result *StartApmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartApmResponse{}
	_body, _err := client.StartApmWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartApmWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartApmResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartApm"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/apm/" + tea.StringValue(instanceId) + "/actions/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartApmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCollector(ResId *string, request *StartCollectorRequest) (_result *StartCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartCollectorResponse{}
	_body, _err := client.StartCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCollectorWithOptions(ResId *string, request *StartCollectorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ResId = openapiutil.GetEncodeParam(ResId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartCollector"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId) + "/actions/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopApm(instanceId *string) (_result *StopApmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopApmResponse{}
	_body, _err := client.StopApmWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopApmWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopApmResponse, _err error) {
	instanceId = openapiutil.GetEncodeParam(instanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopApm"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/apm/" + tea.StringValue(instanceId) + "/actions/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopApmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCollector(ResId *string, request *StopCollectorRequest) (_result *StopCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopCollectorResponse{}
	_body, _err := client.StopCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCollectorWithOptions(ResId *string, request *StopCollectorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ResId = openapiutil.GetEncodeParam(ResId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopCollector"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId) + "/actions/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopPipelines(InstanceId *string, request *StopPipelinesRequest) (_result *StopPipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopPipelinesResponse{}
	_body, _err := client.StopPipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopPipelinesWithOptions(InstanceId *string, request *StopPipelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopPipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopPipelines"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/pipelines/action/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources() (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferNode(InstanceId *string, request *TransferNodeRequest) (_result *TransferNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TransferNodeResponse{}
	_body, _err := client.TransferNodeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferNodeWithOptions(InstanceId *string, request *TransferNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TransferNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["nodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferNode"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/transfer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerNetwork(InstanceId *string, request *TriggerNetworkRequest) (_result *TriggerNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TriggerNetworkResponse{}
	_body, _err := client.TriggerNetworkWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerNetworkWithOptions(InstanceId *string, request *TriggerNetworkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TriggerNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["actionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		body["networkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		body["nodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerNetwork"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/network-trigger"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallKibanaPlugin(InstanceId *string, request *UninstallKibanaPluginRequest) (_result *UninstallKibanaPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallKibanaPluginResponse{}
	_body, _err := client.UninstallKibanaPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallKibanaPluginWithOptions(InstanceId *string, request *UninstallKibanaPluginRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UninstallKibanaPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallKibanaPlugin"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/kibana-plugins/actions/uninstall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallKibanaPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallLogstashPlugin(InstanceId *string, request *UninstallLogstashPluginRequest) (_result *UninstallLogstashPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallLogstashPluginResponse{}
	_body, _err := client.UninstallLogstashPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallLogstashPluginWithOptions(InstanceId *string, request *UninstallLogstashPluginRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UninstallLogstashPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallLogstashPlugin"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/plugins/actions/uninstall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallLogstashPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallPlugin(InstanceId *string, request *UninstallPluginRequest) (_result *UninstallPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallPluginResponse{}
	_body, _err := client.UninstallPluginWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallPluginWithOptions(InstanceId *string, request *UninstallPluginRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UninstallPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallPlugin"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/plugins/actions/uninstall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeys)) {
		query["TagKeys"] = request.TagKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tags"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAdminPassword(InstanceId *string, request *UpdateAdminPasswordRequest) (_result *UpdateAdminPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAdminPasswordResponse{}
	_body, _err := client.UpdateAdminPasswordWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAdminPasswordWithOptions(InstanceId *string, request *UpdateAdminPasswordRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAdminPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAdminPassword"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/admin-pwd"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAdminPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAdvancedSetting(InstanceId *string, request *UpdateAdvancedSettingRequest) (_result *UpdateAdvancedSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAdvancedSettingResponse{}
	_body, _err := client.UpdateAdvancedSettingWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAdvancedSettingWithOptions(InstanceId *string, request *UpdateAdvancedSettingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAdvancedSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAdvancedSetting"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/update-advanced-setting"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAdvancedSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAliwsDict(InstanceId *string, request *UpdateAliwsDictRequest) (_result *UpdateAliwsDictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAliwsDictResponse{}
	_body, _err := client.UpdateAliwsDictWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAliwsDictWithOptions(InstanceId *string, request *UpdateAliwsDictRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAliwsDictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAliwsDict"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/aliws-dict"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAliwsDictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApm(instanceId *string, request *UpdateApmRequest) (_result *UpdateApmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApmResponse{}
	_body, _err := client.UpdateApmWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApmWithOptions(instanceId *string, request *UpdateApmRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateApmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OutputES)) {
		body["outputES"] = request.OutputES
	}

	if !tea.BoolValue(util.IsUnset(request.OutputESPassword)) {
		body["outputESPassword"] = request.OutputESPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OutputESUserName)) {
		body["outputESUserName"] = request.OutputESUserName
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApm"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/apm/" + tea.StringValue(instanceId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBlackIps(InstanceId *string, request *UpdateBlackIpsRequest) (_result *UpdateBlackIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBlackIpsResponse{}
	_body, _err := client.UpdateBlackIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBlackIpsWithOptions(InstanceId *string, request *UpdateBlackIpsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateBlackIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EsIPBlacklist)) {
		body["esIPBlacklist"] = request.EsIPBlacklist
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBlackIps"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/black-ips"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBlackIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCollector(ResId *string, request *UpdateCollectorRequest) (_result *UpdateCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCollectorResponse{}
	_body, _err := client.UpdateCollectorWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCollectorWithOptions(ResId *string, request *UpdateCollectorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ResId = openapiutil.GetEncodeParam(ResId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCollector"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCollectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCollectorName(ResId *string, request *UpdateCollectorNameRequest) (_result *UpdateCollectorNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCollectorNameResponse{}
	_body, _err := client.UpdateCollectorNameWithOptions(ResId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCollectorNameWithOptions(ResId *string, request *UpdateCollectorNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCollectorNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ResId = openapiutil.GetEncodeParam(ResId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCollectorName"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/collectors/" + tea.StringValue(ResId) + "/actions/rename"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCollectorNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateComponentIndex(InstanceId *string, name *string, request *UpdateComponentIndexRequest) (_result *UpdateComponentIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateComponentIndexResponse{}
	_body, _err := client.UpdateComponentIndexWithOptions(InstanceId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateComponentIndexWithOptions(InstanceId *string, name *string, request *UpdateComponentIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateComponentIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	name = openapiutil.GetEncodeParam(name)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Meta)) {
		body["_meta"] = request.Meta
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Template))) {
		body["template"] = request.Template
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateComponentIndex"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/component-index/" + tea.StringValue(name)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateComponentIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDescription(InstanceId *string, request *UpdateDescriptionRequest) (_result *UpdateDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDescriptionResponse{}
	_body, _err := client.UpdateDescriptionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDescriptionWithOptions(InstanceId *string, request *UpdateDescriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDescription"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/description"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDiagnosisSettings(InstanceId *string, request *UpdateDiagnosisSettingsRequest) (_result *UpdateDiagnosisSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDiagnosisSettingsResponse{}
	_body, _err := client.UpdateDiagnosisSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDiagnosisSettingsWithOptions(InstanceId *string, request *UpdateDiagnosisSettingsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDiagnosisSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDiagnosisSettings"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/diagnosis/instances/" + tea.StringValue(InstanceId) + "/settings"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDiagnosisSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDict(InstanceId *string, request *UpdateDictRequest) (_result *UpdateDictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDictResponse{}
	_body, _err := client.UpdateDictWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDictWithOptions(InstanceId *string, request *UpdateDictRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDict"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/dict"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateExtendConfig(InstanceId *string, request *UpdateExtendConfigRequest) (_result *UpdateExtendConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExtendConfigResponse{}
	_body, _err := client.UpdateExtendConfigWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateExtendConfigWithOptions(InstanceId *string, request *UpdateExtendConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExtendConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExtendConfig"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/extend-configs/actions/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExtendConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateExtendfiles(InstanceId *string, request *UpdateExtendfilesRequest) (_result *UpdateExtendfilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExtendfilesResponse{}
	_body, _err := client.UpdateExtendfilesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateExtendfilesWithOptions(InstanceId *string, request *UpdateExtendfilesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExtendfilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExtendfiles"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/extendfiles"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExtendfilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHotIkDicts(InstanceId *string, request *UpdateHotIkDictsRequest) (_result *UpdateHotIkDictsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHotIkDictsResponse{}
	_body, _err := client.UpdateHotIkDictsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHotIkDictsWithOptions(InstanceId *string, request *UpdateHotIkDictsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateHotIkDictsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHotIkDicts"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/ik-hot-dict"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHotIkDictsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateILMPolicy(InstanceId *string, PolicyName *string, request *UpdateILMPolicyRequest) (_result *UpdateILMPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateILMPolicyResponse{}
	_body, _err := client.UpdateILMPolicyWithOptions(InstanceId, PolicyName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateILMPolicyWithOptions(InstanceId *string, PolicyName *string, request *UpdateILMPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateILMPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	PolicyName = openapiutil.GetEncodeParam(PolicyName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateILMPolicy"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/ilm-policies/" + tea.StringValue(PolicyName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateILMPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIndexTemplate(InstanceId *string, IndexTemplate *string, request *UpdateIndexTemplateRequest) (_result *UpdateIndexTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIndexTemplateResponse{}
	_body, _err := client.UpdateIndexTemplateWithOptions(InstanceId, IndexTemplate, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIndexTemplateWithOptions(InstanceId *string, IndexTemplate *string, request *UpdateIndexTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateIndexTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	IndexTemplate = openapiutil.GetEncodeParam(IndexTemplate)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIndexTemplate"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/index-templates/" + tea.StringValue(IndexTemplate)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIndexTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstance(InstanceId *string, request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceWithOptions(InstanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreStatus)) {
		query["ignoreStatus"] = request.IgnoreStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OrderActionType)) {
		query["orderActionType"] = request.OrderActionType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceChargeType(InstanceId *string, request *UpdateInstanceChargeTypeRequest) (_result *UpdateInstanceChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceChargeTypeResponse{}
	_body, _err := client.UpdateInstanceChargeTypeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceChargeTypeWithOptions(InstanceId *string, request *UpdateInstanceChargeTypeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceChargeType"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/convert-pay-type"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceSettings(InstanceId *string, request *UpdateInstanceSettingsRequest) (_result *UpdateInstanceSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceSettingsResponse{}
	_body, _err := client.UpdateInstanceSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceSettingsWithOptions(InstanceId *string, request *UpdateInstanceSettingsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceSettings"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/instance-settings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateKibanaSettings(InstanceId *string, request *UpdateKibanaSettingsRequest) (_result *UpdateKibanaSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKibanaSettingsResponse{}
	_body, _err := client.UpdateKibanaSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateKibanaSettingsWithOptions(InstanceId *string, request *UpdateKibanaSettingsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateKibanaSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateKibanaSettings"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/update-kibana-settings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateKibanaSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateKibanaWhiteIps(InstanceId *string, request *UpdateKibanaWhiteIpsRequest) (_result *UpdateKibanaWhiteIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKibanaWhiteIpsResponse{}
	_body, _err := client.UpdateKibanaWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateKibanaWhiteIpsWithOptions(InstanceId *string, request *UpdateKibanaWhiteIpsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateKibanaWhiteIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["modifyMode"] = request.ModifyMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateKibanaWhiteIps"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/kibana-white-ips"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateKibanaWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLogstash(InstanceId *string, request *UpdateLogstashRequest) (_result *UpdateLogstashResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogstashResponse{}
	_body, _err := client.UpdateLogstashWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLogstashWithOptions(InstanceId *string, request *UpdateLogstashRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLogstashResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogstash"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLogstashResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLogstashChargeType(InstanceId *string, request *UpdateLogstashChargeTypeRequest) (_result *UpdateLogstashChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogstashChargeTypeResponse{}
	_body, _err := client.UpdateLogstashChargeTypeWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLogstashChargeTypeWithOptions(InstanceId *string, request *UpdateLogstashChargeTypeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLogstashChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogstashChargeType"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/actions/convert-pay-type"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLogstashChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLogstashDescription(InstanceId *string, request *UpdateLogstashDescriptionRequest) (_result *UpdateLogstashDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogstashDescriptionResponse{}
	_body, _err := client.UpdateLogstashDescriptionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLogstashDescriptionWithOptions(InstanceId *string, request *UpdateLogstashDescriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLogstashDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogstashDescription"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/description"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLogstashDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLogstashSettings(InstanceId *string, request *UpdateLogstashSettingsRequest) (_result *UpdateLogstashSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogstashSettingsResponse{}
	_body, _err := client.UpdateLogstashSettingsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLogstashSettingsWithOptions(InstanceId *string, request *UpdateLogstashSettingsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLogstashSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogstashSettings"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/instance-settings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLogstashSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePipelineManagementConfig(InstanceId *string, request *UpdatePipelineManagementConfigRequest) (_result *UpdatePipelineManagementConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelineManagementConfigResponse{}
	_body, _err := client.UpdatePipelineManagementConfigWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePipelineManagementConfigWithOptions(InstanceId *string, request *UpdatePipelineManagementConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePipelineManagementConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePipelineManagementConfig"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/pipeline-management-config"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePipelineManagementConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePipelines(InstanceId *string, request *UpdatePipelinesRequest) (_result *UpdatePipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelinesResponse{}
	_body, _err := client.UpdatePipelinesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePipelinesWithOptions(InstanceId *string, request *UpdatePipelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Trigger)) {
		query["trigger"] = request.Trigger
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePipelines"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/pipelines"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePrivateNetworkWhiteIps(InstanceId *string, request *UpdatePrivateNetworkWhiteIpsRequest) (_result *UpdatePrivateNetworkWhiteIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePrivateNetworkWhiteIpsResponse{}
	_body, _err := client.UpdatePrivateNetworkWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePrivateNetworkWhiteIpsWithOptions(InstanceId *string, request *UpdatePrivateNetworkWhiteIpsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePrivateNetworkWhiteIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["modifyMode"] = request.ModifyMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePrivateNetworkWhiteIps"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/private-network-white-ips"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePrivateNetworkWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePublicNetwork(InstanceId *string, request *UpdatePublicNetworkRequest) (_result *UpdatePublicNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePublicNetworkResponse{}
	_body, _err := client.UpdatePublicNetworkWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePublicNetworkWithOptions(InstanceId *string, request *UpdatePublicNetworkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePublicNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePublicNetwork"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/public-network"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePublicNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePublicWhiteIps(InstanceId *string, request *UpdatePublicWhiteIpsRequest) (_result *UpdatePublicWhiteIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePublicWhiteIpsResponse{}
	_body, _err := client.UpdatePublicWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePublicWhiteIpsWithOptions(InstanceId *string, request *UpdatePublicWhiteIpsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePublicWhiteIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["modifyMode"] = request.ModifyMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePublicWhiteIps"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/public-white-ips"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePublicWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateReadWritePolicy(InstanceId *string, request *UpdateReadWritePolicyRequest) (_result *UpdateReadWritePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateReadWritePolicyResponse{}
	_body, _err := client.UpdateReadWritePolicyWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateReadWritePolicyWithOptions(InstanceId *string, request *UpdateReadWritePolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateReadWritePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateReadWritePolicy"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/update-read-write-policy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateReadWritePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSnapshotSetting(InstanceId *string) (_result *UpdateSnapshotSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSnapshotSettingResponse{}
	_body, _err := client.UpdateSnapshotSettingWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSnapshotSettingWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSnapshotSettingResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSnapshotSetting"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/snapshot-setting"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSnapshotSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSynonymsDicts(InstanceId *string, request *UpdateSynonymsDictsRequest) (_result *UpdateSynonymsDictsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSynonymsDictsResponse{}
	_body, _err := client.UpdateSynonymsDictsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSynonymsDictsWithOptions(InstanceId *string, request *UpdateSynonymsDictsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSynonymsDictsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSynonymsDicts"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/synonymsDict"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSynonymsDictsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTemplate(InstanceId *string, TemplateName *string, request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(InstanceId, TemplateName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTemplateWithOptions(InstanceId *string, TemplateName *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	TemplateName = openapiutil.GetEncodeParam(TemplateName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTemplate"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/templates/" + tea.StringValue(TemplateName)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWhiteIps(InstanceId *string, request *UpdateWhiteIpsRequest) (_result *UpdateWhiteIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWhiteIpsResponse{}
	_body, _err := client.UpdateWhiteIpsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWhiteIpsWithOptions(InstanceId *string, request *UpdateWhiteIpsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWhiteIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["modifyMode"] = request.ModifyMode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EsIPWhitelist)) {
		body["esIPWhitelist"] = request.EsIPWhitelist
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.WhiteIpGroup))) {
		body["whiteIpGroup"] = request.WhiteIpGroup
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWhiteIps"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/white-ips"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWhiteIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateXpackMonitorConfig(InstanceId *string, request *UpdateXpackMonitorConfigRequest) (_result *UpdateXpackMonitorConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateXpackMonitorConfigResponse{}
	_body, _err := client.UpdateXpackMonitorConfigWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateXpackMonitorConfigWithOptions(InstanceId *string, request *UpdateXpackMonitorConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateXpackMonitorConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateXpackMonitorConfig"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/xpack-monitor-config"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateXpackMonitorConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeEngineVersion(InstanceId *string, request *UpgradeEngineVersionRequest) (_result *UpgradeEngineVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeEngineVersionResponse{}
	_body, _err := client.UpgradeEngineVersionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeEngineVersionWithOptions(InstanceId *string, request *UpgradeEngineVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpgradeEngineVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeEngineVersion"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/actions/upgrade-version"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeEngineVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateConnection(InstanceId *string, request *ValidateConnectionRequest) (_result *ValidateConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateConnectionResponse{}
	_body, _err := client.ValidateConnectionWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateConnectionWithOptions(InstanceId *string, request *ValidateConnectionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateConnection"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/logstashes/" + tea.StringValue(InstanceId) + "/validate-connection"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateShrinkNodes(InstanceId *string, request *ValidateShrinkNodesRequest) (_result *ValidateShrinkNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateShrinkNodesResponse{}
	_body, _err := client.ValidateShrinkNodesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateShrinkNodesWithOptions(InstanceId *string, request *ValidateShrinkNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateShrinkNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IgnoreStatus)) {
		query["ignoreStatus"] = request.IgnoreStatus
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["nodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateShrinkNodes"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/validate-shrink-nodes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateShrinkNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateSlrPermission(request *ValidateSlrPermissionRequest) (_result *ValidateSlrPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateSlrPermissionResponse{}
	_body, _err := client.ValidateSlrPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateSlrPermissionWithOptions(request *ValidateSlrPermissionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateSlrPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Rolename)) {
		query["rolename"] = request.Rolename
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateSlrPermission"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/user/servicerolepermission"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateSlrPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateTransferableNodes(InstanceId *string, request *ValidateTransferableNodesRequest) (_result *ValidateTransferableNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateTransferableNodesResponse{}
	_body, _err := client.ValidateTransferableNodesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateTransferableNodesWithOptions(InstanceId *string, request *ValidateTransferableNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateTransferableNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		query["nodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateTransferableNodes"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances/" + tea.StringValue(InstanceId) + "/validate-transfer-nodes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateTransferableNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("createInstance"),
		Version:     tea.String("2017-06-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
