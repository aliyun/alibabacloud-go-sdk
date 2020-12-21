// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ValidateVirtualClusterNameRequest struct {
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s ValidateVirtualClusterNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameRequest) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameRequest) SetVcName(v string) *ValidateVirtualClusterNameRequest {
	s.VcName = &v
	return s
}

type ValidateVirtualClusterNameResponse struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ValidateVirtualClusterNameResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ValidateVirtualClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameResponse) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameResponse) SetRequestId(v string) *ValidateVirtualClusterNameResponse {
	s.RequestId = &v
	return s
}

func (s *ValidateVirtualClusterNameResponse) SetData(v *ValidateVirtualClusterNameResponseData) *ValidateVirtualClusterNameResponse {
	s.Data = v
	return s
}

type ValidateVirtualClusterNameResponseData struct {
	Legal   *string `json:"Legal,omitempty" xml:"Legal,omitempty" require:"true"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s ValidateVirtualClusterNameResponseData) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameResponseData) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameResponseData) SetLegal(v string) *ValidateVirtualClusterNameResponseData {
	s.Legal = &v
	return s
}

func (s *ValidateVirtualClusterNameResponseData) SetMessage(v string) *ValidateVirtualClusterNameResponseData {
	s.Message = &v
	return s
}

type UpgradeInstanceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ChargeType   *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty" require:"true"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" require:"true"`
	Component    *string `json:"Component,omitempty" xml:"Component,omitempty" require:"true"`
}

func (s UpgradeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceRequest) SetRegionId(v string) *UpgradeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetInstanceId(v string) *UpgradeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetChargeType(v string) *UpgradeInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *UpgradeInstanceRequest) SetInstanceType(v string) *UpgradeInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *UpgradeInstanceRequest) SetComponent(v string) *UpgradeInstanceRequest {
	s.Component = &v
	return s
}

type UpgradeInstanceResponse struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" require:"true"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpgradeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponse) SetSuccess(v bool) *UpgradeInstanceResponse {
	s.Success = &v
	return s
}

func (s *UpgradeInstanceResponse) SetErrorCode(v string) *UpgradeInstanceResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeInstanceResponse) SetErrorInfo(v string) *UpgradeInstanceResponse {
	s.ErrorInfo = &v
	return s
}

func (s *UpgradeInstanceResponse) SetResult(v string) *UpgradeInstanceResponse {
	s.Result = &v
	return s
}

func (s *UpgradeInstanceResponse) SetRequestId(v string) *UpgradeInstanceResponse {
	s.RequestId = &v
	return s
}

type CreateInstanceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ChargeType   *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty" require:"true"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" require:"true"`
	Component    *string `json:"Component,omitempty" xml:"Component,omitempty" require:"true"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetComponent(v string) *CreateInstanceRequest {
	s.Component = &v
	return s
}

type CreateInstanceResponse struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" require:"true"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetSuccess(v bool) *CreateInstanceResponse {
	s.Success = &v
	return s
}

func (s *CreateInstanceResponse) SetErrorCode(v string) *CreateInstanceResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateInstanceResponse) SetErrorInfo(v string) *CreateInstanceResponse {
	s.ErrorInfo = &v
	return s
}

func (s *CreateInstanceResponse) SetResult(v string) *CreateInstanceResponse {
	s.Result = &v
	return s
}

func (s *CreateInstanceResponse) SetRequestId(v string) *CreateInstanceResponse {
	s.RequestId = &v
	return s
}

type ReleaseInstanceRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetRegionId(v string) *ReleaseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleaseInstanceResponse struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" require:"true"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ReleaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) SetSuccess(v bool) *ReleaseInstanceResponse {
	s.Success = &v
	return s
}

func (s *ReleaseInstanceResponse) SetErrorCode(v string) *ReleaseInstanceResponse {
	s.ErrorCode = &v
	return s
}

func (s *ReleaseInstanceResponse) SetErrorInfo(v string) *ReleaseInstanceResponse {
	s.ErrorInfo = &v
	return s
}

func (s *ReleaseInstanceResponse) SetResult(v string) *ReleaseInstanceResponse {
	s.Result = &v
	return s
}

func (s *ReleaseInstanceResponse) SetRequestId(v string) *ReleaseInstanceResponse {
	s.RequestId = &v
	return s
}

type GetJobDetailRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty" require:"true"`
}

func (s GetJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailRequest) GoString() string {
	return s.String()
}

func (s *GetJobDetailRequest) SetJobId(v string) *GetJobDetailRequest {
	s.JobId = &v
	return s
}

func (s *GetJobDetailRequest) SetVcName(v string) *GetJobDetailRequest {
	s.VcName = &v
	return s
}

type GetJobDetailResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	JobDetail *GetJobDetailResponseJobDetail `json:"JobDetail,omitempty" xml:"JobDetail,omitempty" require:"true" type:"Struct"`
}

func (s GetJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailResponse) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponse) SetRequestId(v string) *GetJobDetailResponse {
	s.RequestId = &v
	return s
}

func (s *GetJobDetailResponse) SetJobDetail(v *GetJobDetailResponseJobDetail) *GetJobDetailResponse {
	s.JobDetail = v
	return s
}

type GetJobDetailResponseJobDetail struct {
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty" require:"true"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Detail               *string `json:"Detail,omitempty" xml:"Detail,omitempty" require:"true"`
	SparkUI              *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty" require:"true"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	SubmitTime           *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty" require:"true"`
	CreateTimeValue      *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty" require:"true"`
	UpdateTimeValue      *string `json:"UpdateTimeValue,omitempty" xml:"UpdateTimeValue,omitempty" require:"true"`
	SubmitTimeValue      *string `json:"SubmitTimeValue,omitempty" xml:"SubmitTimeValue,omitempty" require:"true"`
	VcName               *string `json:"VcName,omitempty" xml:"VcName,omitempty" require:"true"`
	DriverResourceSpec   *string `json:"DriverResourceSpec,omitempty" xml:"DriverResourceSpec,omitempty" require:"true"`
	ExecutorResourceSpec *string `json:"ExecutorResourceSpec,omitempty" xml:"ExecutorResourceSpec,omitempty" require:"true"`
	ExecutorInstances    *string `json:"ExecutorInstances,omitempty" xml:"ExecutorInstances,omitempty" require:"true"`
}

func (s GetJobDetailResponseJobDetail) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailResponseJobDetail) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseJobDetail) SetJobId(v string) *GetJobDetailResponseJobDetail {
	s.JobId = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetJobName(v string) *GetJobDetailResponseJobDetail {
	s.JobName = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetStatus(v string) *GetJobDetailResponseJobDetail {
	s.Status = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetDetail(v string) *GetJobDetailResponseJobDetail {
	s.Detail = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetSparkUI(v string) *GetJobDetailResponseJobDetail {
	s.SparkUI = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetCreateTime(v string) *GetJobDetailResponseJobDetail {
	s.CreateTime = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetUpdateTime(v string) *GetJobDetailResponseJobDetail {
	s.UpdateTime = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetSubmitTime(v string) *GetJobDetailResponseJobDetail {
	s.SubmitTime = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetCreateTimeValue(v string) *GetJobDetailResponseJobDetail {
	s.CreateTimeValue = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetUpdateTimeValue(v string) *GetJobDetailResponseJobDetail {
	s.UpdateTimeValue = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetSubmitTimeValue(v string) *GetJobDetailResponseJobDetail {
	s.SubmitTimeValue = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetVcName(v string) *GetJobDetailResponseJobDetail {
	s.VcName = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetDriverResourceSpec(v string) *GetJobDetailResponseJobDetail {
	s.DriverResourceSpec = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetExecutorResourceSpec(v string) *GetJobDetailResponseJobDetail {
	s.ExecutorResourceSpec = &v
	return s
}

func (s *GetJobDetailResponseJobDetail) SetExecutorInstances(v string) *GetJobDetailResponseJobDetail {
	s.ExecutorInstances = &v
	return s
}

type GetJobStatusRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty" require:"true"`
}

func (s GetJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetJobStatusRequest) SetJobId(v string) *GetJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetJobStatusRequest) SetVcName(v string) *GetJobStatusRequest {
	s.VcName = &v
	return s
}

type GetJobStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s GetJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponse) SetRequestId(v string) *GetJobStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetJobStatusResponse) SetStatus(v string) *GetJobStatusResponse {
	s.Status = &v
	return s
}

type KillSparkJobRequest struct {
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty" require:"true"`
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
}

func (s KillSparkJobRequest) String() string {
	return tea.Prettify(s)
}

func (s KillSparkJobRequest) GoString() string {
	return s.String()
}

func (s *KillSparkJobRequest) SetVcName(v string) *KillSparkJobRequest {
	s.VcName = &v
	return s
}

func (s *KillSparkJobRequest) SetJobId(v string) *KillSparkJobRequest {
	s.JobId = &v
	return s
}

type KillSparkJobResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s KillSparkJobResponse) String() string {
	return tea.Prettify(s)
}

func (s KillSparkJobResponse) GoString() string {
	return s.String()
}

func (s *KillSparkJobResponse) SetData(v string) *KillSparkJobResponse {
	s.Data = &v
	return s
}

func (s *KillSparkJobResponse) SetRequestId(v string) *KillSparkJobResponse {
	s.RequestId = &v
	return s
}

type SubmitSparkJobRequest struct {
	VcName     *string `json:"VcName,omitempty" xml:"VcName,omitempty" require:"true"`
	ConfigJson *string `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty" require:"true"`
}

func (s SubmitSparkJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSparkJobRequest) SetVcName(v string) *SubmitSparkJobRequest {
	s.VcName = &v
	return s
}

func (s *SubmitSparkJobRequest) SetConfigJson(v string) *SubmitSparkJobRequest {
	s.ConfigJson = &v
	return s
}

type SubmitSparkJobResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
}

func (s SubmitSparkJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSparkJobResponse) SetRequestId(v string) *SubmitSparkJobResponse {
	s.RequestId = &v
	return s
}

func (s *SubmitSparkJobResponse) SetJobId(v string) *SubmitSparkJobResponse {
	s.JobId = &v
	return s
}

type GetJobLogRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty" require:"true"`
}

func (s GetJobLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetJobLogRequest) SetJobId(v string) *GetJobLogRequest {
	s.JobId = &v
	return s
}

func (s *GetJobLogRequest) SetVcName(v string) *GetJobLogRequest {
	s.VcName = &v
	return s
}

type GetJobLogResponse struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s GetJobLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetJobLogResponse) SetData(v string) *GetJobLogResponse {
	s.Data = &v
	return s
}

func (s *GetJobLogResponse) SetRequestId(v string) *GetJobLogResponse {
	s.RequestId = &v
	return s
}

type ListSparkJobRequest struct {
	VcName     *string `json:"VcName,omitempty" xml:"VcName,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s ListSparkJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobRequest) GoString() string {
	return s.String()
}

func (s *ListSparkJobRequest) SetVcName(v string) *ListSparkJobRequest {
	s.VcName = &v
	return s
}

func (s *ListSparkJobRequest) SetPageNumber(v int) *ListSparkJobRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobRequest) SetPageSize(v int) *ListSparkJobRequest {
	s.PageSize = &v
	return s
}

type ListSparkJobResponse struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DataResult *ListSparkJobResponseDataResult `json:"DataResult,omitempty" xml:"DataResult,omitempty" require:"true" type:"Struct"`
}

func (s ListSparkJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponse) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponse) SetRequestId(v string) *ListSparkJobResponse {
	s.RequestId = &v
	return s
}

func (s *ListSparkJobResponse) SetDataResult(v *ListSparkJobResponseDataResult) *ListSparkJobResponse {
	s.DataResult = v
	return s
}

type ListSparkJobResponseDataResult struct {
	PageNumber *string                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *string                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *string                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	JobList    []*ListSparkJobResponseDataResultJobList `json:"JobList,omitempty" xml:"JobList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSparkJobResponseDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponseDataResult) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponseDataResult) SetPageNumber(v string) *ListSparkJobResponseDataResult {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobResponseDataResult) SetPageSize(v string) *ListSparkJobResponseDataResult {
	s.PageSize = &v
	return s
}

func (s *ListSparkJobResponseDataResult) SetTotalCount(v string) *ListSparkJobResponseDataResult {
	s.TotalCount = &v
	return s
}

func (s *ListSparkJobResponseDataResult) SetJobList(v []*ListSparkJobResponseDataResultJobList) *ListSparkJobResponseDataResult {
	s.JobList = v
	return s
}

type ListSparkJobResponseDataResultJobList struct {
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	CreateTimeValue      *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty" require:"true"`
	Detail               *string `json:"Detail,omitempty" xml:"Detail,omitempty" require:"true"`
	DriverResourceSpec   *string `json:"DriverResourceSpec,omitempty" xml:"DriverResourceSpec,omitempty" require:"true"`
	ExecutorInstances    *string `json:"ExecutorInstances,omitempty" xml:"ExecutorInstances,omitempty" require:"true"`
	ExecutorResourceSpec *string `json:"ExecutorResourceSpec,omitempty" xml:"ExecutorResourceSpec,omitempty" require:"true"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty" require:"true"`
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty" require:"true"`
	SparkUI              *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty" require:"true"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	SubmitTime           *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty" require:"true"`
	SubmitTimeValue      *string `json:"SubmitTimeValue,omitempty" xml:"SubmitTimeValue,omitempty" require:"true"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	UpdateTimeValue      *string `json:"UpdateTimeValue,omitempty" xml:"UpdateTimeValue,omitempty" require:"true"`
	VcName               *string `json:"VcName,omitempty" xml:"VcName,omitempty" require:"true"`
}

func (s ListSparkJobResponseDataResultJobList) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponseDataResultJobList) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponseDataResultJobList) SetCreateTime(v string) *ListSparkJobResponseDataResultJobList {
	s.CreateTime = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetCreateTimeValue(v string) *ListSparkJobResponseDataResultJobList {
	s.CreateTimeValue = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetDetail(v string) *ListSparkJobResponseDataResultJobList {
	s.Detail = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetDriverResourceSpec(v string) *ListSparkJobResponseDataResultJobList {
	s.DriverResourceSpec = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetExecutorInstances(v string) *ListSparkJobResponseDataResultJobList {
	s.ExecutorInstances = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetExecutorResourceSpec(v string) *ListSparkJobResponseDataResultJobList {
	s.ExecutorResourceSpec = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetJobId(v string) *ListSparkJobResponseDataResultJobList {
	s.JobId = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetJobName(v string) *ListSparkJobResponseDataResultJobList {
	s.JobName = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetSparkUI(v string) *ListSparkJobResponseDataResultJobList {
	s.SparkUI = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetStatus(v string) *ListSparkJobResponseDataResultJobList {
	s.Status = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetSubmitTime(v string) *ListSparkJobResponseDataResultJobList {
	s.SubmitTime = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetSubmitTimeValue(v string) *ListSparkJobResponseDataResultJobList {
	s.SubmitTimeValue = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetUpdateTime(v string) *ListSparkJobResponseDataResultJobList {
	s.UpdateTime = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetUpdateTimeValue(v string) *ListSparkJobResponseDataResultJobList {
	s.UpdateTimeValue = &v
	return s
}

func (s *ListSparkJobResponseDataResultJobList) SetVcName(v string) *ListSparkJobResponseDataResultJobList {
	s.VcName = &v
	return s
}

type GetAllowIPRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" require:"true"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetAllowIPRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllowIPRequest) GoString() string {
	return s.String()
}

func (s *GetAllowIPRequest) SetRegionId(v string) *GetAllowIPRequest {
	s.RegionId = &v
	return s
}

func (s *GetAllowIPRequest) SetNetworkType(v string) *GetAllowIPRequest {
	s.NetworkType = &v
	return s
}

func (s *GetAllowIPRequest) SetProduct(v string) *GetAllowIPRequest {
	s.Product = &v
	return s
}

type GetAllowIPResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	AllowIP   *string `json:"AllowIP,omitempty" xml:"AllowIP,omitempty" require:"true"`
}

func (s GetAllowIPResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllowIPResponse) GoString() string {
	return s.String()
}

func (s *GetAllowIPResponse) SetRequestId(v string) *GetAllowIPResponse {
	s.RequestId = &v
	return s
}

func (s *GetAllowIPResponse) SetRegionId(v string) *GetAllowIPResponse {
	s.RegionId = &v
	return s
}

func (s *GetAllowIPResponse) SetAllowIP(v string) *GetAllowIPResponse {
	s.AllowIP = &v
	return s
}

type SetAllowIPRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" require:"true"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	AllowIP     *string `json:"AllowIP,omitempty" xml:"AllowIP,omitempty" require:"true"`
	Append      *bool   `json:"Append,omitempty" xml:"Append,omitempty"`
}

func (s SetAllowIPRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAllowIPRequest) GoString() string {
	return s.String()
}

func (s *SetAllowIPRequest) SetRegionId(v string) *SetAllowIPRequest {
	s.RegionId = &v
	return s
}

func (s *SetAllowIPRequest) SetNetworkType(v string) *SetAllowIPRequest {
	s.NetworkType = &v
	return s
}

func (s *SetAllowIPRequest) SetProduct(v string) *SetAllowIPRequest {
	s.Product = &v
	return s
}

func (s *SetAllowIPRequest) SetAllowIP(v string) *SetAllowIPRequest {
	s.AllowIP = &v
	return s
}

func (s *SetAllowIPRequest) SetAppend(v bool) *SetAllowIPRequest {
	s.Append = &v
	return s
}

type SetAllowIPResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s SetAllowIPResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAllowIPResponse) GoString() string {
	return s.String()
}

func (s *SetAllowIPResponse) SetRequestId(v string) *SetAllowIPResponse {
	s.RequestId = &v
	return s
}

func (s *SetAllowIPResponse) SetRegionId(v string) *SetAllowIPResponse {
	s.RegionId = &v
	return s
}

type ResetMainPasswordRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UseRandomPassword    *bool   `json:"UseRandomPassword,omitempty" xml:"UseRandomPassword,omitempty" require:"true"`
	InitPassword         *string `json:"InitPassword,omitempty" xml:"InitPassword,omitempty"`
	EnableKMS            *bool   `json:"EnableKMS,omitempty" xml:"EnableKMS,omitempty"`
	ExternalUid          *string `json:"ExternalUid,omitempty" xml:"ExternalUid,omitempty"`
	ExternalAliyunUid    *string `json:"ExternalAliyunUid,omitempty" xml:"ExternalAliyunUid,omitempty"`
	ExternalBizAliyunUid *string `json:"ExternalBizAliyunUid,omitempty" xml:"ExternalBizAliyunUid,omitempty"`
}

func (s ResetMainPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetMainPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetMainPasswordRequest) SetRegionId(v string) *ResetMainPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetMainPasswordRequest) SetUseRandomPassword(v bool) *ResetMainPasswordRequest {
	s.UseRandomPassword = &v
	return s
}

func (s *ResetMainPasswordRequest) SetInitPassword(v string) *ResetMainPasswordRequest {
	s.InitPassword = &v
	return s
}

func (s *ResetMainPasswordRequest) SetEnableKMS(v bool) *ResetMainPasswordRequest {
	s.EnableKMS = &v
	return s
}

func (s *ResetMainPasswordRequest) SetExternalUid(v string) *ResetMainPasswordRequest {
	s.ExternalUid = &v
	return s
}

func (s *ResetMainPasswordRequest) SetExternalAliyunUid(v string) *ResetMainPasswordRequest {
	s.ExternalAliyunUid = &v
	return s
}

func (s *ResetMainPasswordRequest) SetExternalBizAliyunUid(v string) *ResetMainPasswordRequest {
	s.ExternalBizAliyunUid = &v
	return s
}

type ResetMainPasswordResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId  *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Account   *ResetMainPasswordResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s ResetMainPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetMainPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetMainPasswordResponse) SetRequestId(v string) *ResetMainPasswordResponse {
	s.RequestId = &v
	return s
}

func (s *ResetMainPasswordResponse) SetRegionId(v string) *ResetMainPasswordResponse {
	s.RegionId = &v
	return s
}

func (s *ResetMainPasswordResponse) SetAccount(v *ResetMainPasswordResponseAccount) *ResetMainPasswordResponse {
	s.Account = v
	return s
}

type ResetMainPasswordResponseAccount struct {
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty" require:"true"`
}

func (s ResetMainPasswordResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s ResetMainPasswordResponseAccount) GoString() string {
	return s.String()
}

func (s *ResetMainPasswordResponseAccount) SetUserName(v string) *ResetMainPasswordResponseAccount {
	s.UserName = &v
	return s
}

func (s *ResetMainPasswordResponseAccount) SetPassword(v string) *ResetMainPasswordResponseAccount {
	s.Password = &v
	return s
}

type UnSubscribeRegionRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s UnSubscribeRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s UnSubscribeRegionRequest) GoString() string {
	return s.String()
}

func (s *UnSubscribeRegionRequest) SetRegionId(v string) *UnSubscribeRegionRequest {
	s.RegionId = &v
	return s
}

type UnSubscribeRegionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s UnSubscribeRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s UnSubscribeRegionResponse) GoString() string {
	return s.String()
}

func (s *UnSubscribeRegionResponse) SetRequestId(v string) *UnSubscribeRegionResponse {
	s.RequestId = &v
	return s
}

func (s *UnSubscribeRegionResponse) SetRegionId(v string) *UnSubscribeRegionResponse {
	s.RegionId = &v
	return s
}

type QueryAccountListRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
}

func (s QueryAccountListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountListRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountListRequest) SetRegionId(v string) *QueryAccountListRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAccountListRequest) SetPageSize(v int) *QueryAccountListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAccountListRequest) SetPageNumber(v int) *QueryAccountListRequest {
	s.PageNumber = &v
	return s
}

type QueryAccountListResponse struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId   *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PageSize   *int                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNumber *int                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	TotalCount *int                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Data       []*QueryAccountListResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAccountListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountListResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountListResponse) SetRequestId(v string) *QueryAccountListResponse {
	s.RequestId = &v
	return s
}

func (s *QueryAccountListResponse) SetRegionId(v string) *QueryAccountListResponse {
	s.RegionId = &v
	return s
}

func (s *QueryAccountListResponse) SetPageSize(v int) *QueryAccountListResponse {
	s.PageSize = &v
	return s
}

func (s *QueryAccountListResponse) SetPageNumber(v int) *QueryAccountListResponse {
	s.PageNumber = &v
	return s
}

func (s *QueryAccountListResponse) SetTotalCount(v int) *QueryAccountListResponse {
	s.TotalCount = &v
	return s
}

func (s *QueryAccountListResponse) SetData(v []*QueryAccountListResponseData) *QueryAccountListResponse {
	s.Data = v
	return s
}

type QueryAccountListResponseData struct {
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty" require:"true"`
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty" require:"true"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty" require:"true"`
	RamUid    *string `json:"RamUid,omitempty" xml:"RamUid,omitempty" require:"true"`
}

func (s QueryAccountListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountListResponseData) GoString() string {
	return s.String()
}

func (s *QueryAccountListResponseData) SetUserName(v string) *QueryAccountListResponseData {
	s.UserName = &v
	return s
}

func (s *QueryAccountListResponseData) SetRole(v string) *QueryAccountListResponseData {
	s.Role = &v
	return s
}

func (s *QueryAccountListResponseData) SetShortName(v string) *QueryAccountListResponseData {
	s.ShortName = &v
	return s
}

func (s *QueryAccountListResponseData) SetRemark(v string) *QueryAccountListResponseData {
	s.Remark = &v
	return s
}

func (s *QueryAccountListResponseData) SetRamUid(v string) *QueryAccountListResponseData {
	s.RamUid = &v
	return s
}

type DeleteAccountRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	IsShort       *bool   `json:"IsShort,omitempty" xml:"IsShort,omitempty"`
	IsServiceUser *bool   `json:"IsServiceUser,omitempty" xml:"IsServiceUser,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetRegionId(v string) *DeleteAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetIsShort(v bool) *DeleteAccountRequest {
	s.IsShort = &v
	return s
}

func (s *DeleteAccountRequest) SetIsServiceUser(v bool) *DeleteAccountRequest {
	s.IsServiceUser = &v
	return s
}

type DeleteAccountResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetRequestId(v string) *DeleteAccountResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountResponse) SetRegionId(v string) *DeleteAccountResponse {
	s.RegionId = &v
	return s
}

type AddAccountRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
	IsShort           *bool   `json:"IsShort,omitempty" xml:"IsShort,omitempty"`
	EnableKMS         *bool   `json:"EnableKMS,omitempty" xml:"EnableKMS,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RamUid            *string `json:"RamUid,omitempty" xml:"RamUid,omitempty"`
	UseRandomPassword *bool   `json:"UseRandomPassword,omitempty" xml:"UseRandomPassword,omitempty"`
}

func (s AddAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAccountRequest) GoString() string {
	return s.String()
}

func (s *AddAccountRequest) SetRegionId(v string) *AddAccountRequest {
	s.RegionId = &v
	return s
}

func (s *AddAccountRequest) SetAccountName(v string) *AddAccountRequest {
	s.AccountName = &v
	return s
}

func (s *AddAccountRequest) SetPassword(v string) *AddAccountRequest {
	s.Password = &v
	return s
}

func (s *AddAccountRequest) SetIsShort(v bool) *AddAccountRequest {
	s.IsShort = &v
	return s
}

func (s *AddAccountRequest) SetEnableKMS(v bool) *AddAccountRequest {
	s.EnableKMS = &v
	return s
}

func (s *AddAccountRequest) SetRemark(v string) *AddAccountRequest {
	s.Remark = &v
	return s
}

func (s *AddAccountRequest) SetRamUid(v string) *AddAccountRequest {
	s.RamUid = &v
	return s
}

func (s *AddAccountRequest) SetUseRandomPassword(v bool) *AddAccountRequest {
	s.UseRandomPassword = &v
	return s
}

type AddAccountResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId  *string                    `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Account   *AddAccountResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s AddAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAccountResponse) GoString() string {
	return s.String()
}

func (s *AddAccountResponse) SetRequestId(v string) *AddAccountResponse {
	s.RequestId = &v
	return s
}

func (s *AddAccountResponse) SetRegionId(v string) *AddAccountResponse {
	s.RegionId = &v
	return s
}

func (s *AddAccountResponse) SetAccount(v *AddAccountResponseAccount) *AddAccountResponse {
	s.Account = v
	return s
}

type AddAccountResponseAccount struct {
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty" require:"true"`
}

func (s AddAccountResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s AddAccountResponseAccount) GoString() string {
	return s.String()
}

func (s *AddAccountResponseAccount) SetUserName(v string) *AddAccountResponseAccount {
	s.UserName = &v
	return s
}

func (s *AddAccountResponseAccount) SetPassword(v string) *AddAccountResponseAccount {
	s.Password = &v
	return s
}

type UpdateAccountPasswordRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	IsShort           *bool   `json:"IsShort,omitempty" xml:"IsShort,omitempty"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
	EnableKMS         *bool   `json:"EnableKMS,omitempty" xml:"EnableKMS,omitempty"`
	UseRandomPassword *bool   `json:"UseRandomPassword,omitempty" xml:"UseRandomPassword,omitempty"`
}

func (s UpdateAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountPasswordRequest) SetRegionId(v string) *UpdateAccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetAccountName(v string) *UpdateAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetIsShort(v bool) *UpdateAccountPasswordRequest {
	s.IsShort = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetPassword(v string) *UpdateAccountPasswordRequest {
	s.Password = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetEnableKMS(v bool) *UpdateAccountPasswordRequest {
	s.EnableKMS = &v
	return s
}

func (s *UpdateAccountPasswordRequest) SetUseRandomPassword(v bool) *UpdateAccountPasswordRequest {
	s.UseRandomPassword = &v
	return s
}

type UpdateAccountPasswordResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId  *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Account   *UpdateAccountPasswordResponseAccount `json:"Account,omitempty" xml:"Account,omitempty" require:"true" type:"Struct"`
}

func (s UpdateAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountPasswordResponse) SetRequestId(v string) *UpdateAccountPasswordResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateAccountPasswordResponse) SetRegionId(v string) *UpdateAccountPasswordResponse {
	s.RegionId = &v
	return s
}

func (s *UpdateAccountPasswordResponse) SetAccount(v *UpdateAccountPasswordResponseAccount) *UpdateAccountPasswordResponse {
	s.Account = v
	return s
}

type UpdateAccountPasswordResponseAccount struct {
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty" require:"true"`
}

func (s UpdateAccountPasswordResponseAccount) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountPasswordResponseAccount) GoString() string {
	return s.String()
}

func (s *UpdateAccountPasswordResponseAccount) SetUserName(v string) *UpdateAccountPasswordResponseAccount {
	s.UserName = &v
	return s
}

func (s *UpdateAccountPasswordResponseAccount) SetPassword(v string) *UpdateAccountPasswordResponseAccount {
	s.Password = &v
	return s
}

type RemoveEndPointRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	EndPointID *string `json:"EndPointID,omitempty" xml:"EndPointID,omitempty" require:"true"`
}

func (s RemoveEndPointRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEndPointRequest) GoString() string {
	return s.String()
}

func (s *RemoveEndPointRequest) SetRegionId(v string) *RemoveEndPointRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveEndPointRequest) SetEndPointID(v string) *RemoveEndPointRequest {
	s.EndPointID = &v
	return s
}

type RemoveEndPointResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s RemoveEndPointResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveEndPointResponse) GoString() string {
	return s.String()
}

func (s *RemoveEndPointResponse) SetRequestId(v string) *RemoveEndPointResponse {
	s.RequestId = &v
	return s
}

func (s *RemoveEndPointResponse) SetRegionId(v string) *RemoveEndPointResponse {
	s.RegionId = &v
	return s
}

type QueryEndPointListRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s QueryEndPointListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEndPointListRequest) GoString() string {
	return s.String()
}

func (s *QueryEndPointListRequest) SetRegionId(v string) *QueryEndPointListRequest {
	s.RegionId = &v
	return s
}

type QueryEndPointListResponse struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId     *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	EndPointList []*QueryEndPointListResponseEndPointList `json:"EndPointList,omitempty" xml:"EndPointList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEndPointListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEndPointListResponse) GoString() string {
	return s.String()
}

func (s *QueryEndPointListResponse) SetRequestId(v string) *QueryEndPointListResponse {
	s.RequestId = &v
	return s
}

func (s *QueryEndPointListResponse) SetRegionId(v string) *QueryEndPointListResponse {
	s.RegionId = &v
	return s
}

func (s *QueryEndPointListResponse) SetEndPointList(v []*QueryEndPointListResponseEndPointList) *QueryEndPointListResponse {
	s.EndPointList = v
	return s
}

type QueryEndPointListResponseEndPointList struct {
	EndPointID  *string `json:"endPointID,omitempty" xml:"endPointID,omitempty" require:"true"`
	Zone        *string `json:"zone,omitempty" xml:"zone,omitempty" require:"true"`
	VSwitch     *string `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" require:"true"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	VpcID       *string `json:"vpcID,omitempty" xml:"vpcID,omitempty" require:"true"`
	Host        *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
	DomainURL   *string `json:"domainURL,omitempty" xml:"domainURL,omitempty" require:"true"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty" require:"true"`
	AllowIP     *string `json:"allowIP,omitempty" xml:"allowIP,omitempty" require:"true"`
	Port        *string `json:"port,omitempty" xml:"port,omitempty" require:"true"`
	Product     *string `json:"product,omitempty" xml:"product,omitempty" require:"true"`
}

func (s QueryEndPointListResponseEndPointList) String() string {
	return tea.Prettify(s)
}

func (s QueryEndPointListResponseEndPointList) GoString() string {
	return s.String()
}

func (s *QueryEndPointListResponseEndPointList) SetEndPointID(v string) *QueryEndPointListResponseEndPointList {
	s.EndPointID = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetZone(v string) *QueryEndPointListResponseEndPointList {
	s.Zone = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetVSwitch(v string) *QueryEndPointListResponseEndPointList {
	s.VSwitch = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetStatus(v string) *QueryEndPointListResponseEndPointList {
	s.Status = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetVpcID(v string) *QueryEndPointListResponseEndPointList {
	s.VpcID = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetHost(v string) *QueryEndPointListResponseEndPointList {
	s.Host = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetDomainURL(v string) *QueryEndPointListResponseEndPointList {
	s.DomainURL = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetNetworkType(v string) *QueryEndPointListResponseEndPointList {
	s.NetworkType = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetAllowIP(v string) *QueryEndPointListResponseEndPointList {
	s.AllowIP = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetPort(v string) *QueryEndPointListResponseEndPointList {
	s.Port = &v
	return s
}

func (s *QueryEndPointListResponseEndPointList) SetProduct(v string) *QueryEndPointListResponseEndPointList {
	s.Product = &v
	return s
}

type AddEndPointRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" require:"true"`
	VpcID       *string `json:"VpcID,omitempty" xml:"VpcID,omitempty" require:"true"`
	Vswitch     *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" require:"true"`
	Zone        *string `json:"Zone,omitempty" xml:"Zone,omitempty" require:"true"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s AddEndPointRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEndPointRequest) GoString() string {
	return s.String()
}

func (s *AddEndPointRequest) SetRegionId(v string) *AddEndPointRequest {
	s.RegionId = &v
	return s
}

func (s *AddEndPointRequest) SetNetworkType(v string) *AddEndPointRequest {
	s.NetworkType = &v
	return s
}

func (s *AddEndPointRequest) SetVpcID(v string) *AddEndPointRequest {
	s.VpcID = &v
	return s
}

func (s *AddEndPointRequest) SetVswitch(v string) *AddEndPointRequest {
	s.Vswitch = &v
	return s
}

func (s *AddEndPointRequest) SetZone(v string) *AddEndPointRequest {
	s.Zone = &v
	return s
}

func (s *AddEndPointRequest) SetProduct(v string) *AddEndPointRequest {
	s.Product = &v
	return s
}

type AddEndPointResponse struct {
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId     *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	EndPointInfo *AddEndPointResponseEndPointInfo `json:"EndPointInfo,omitempty" xml:"EndPointInfo,omitempty" require:"true" type:"Struct"`
}

func (s AddEndPointResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEndPointResponse) GoString() string {
	return s.String()
}

func (s *AddEndPointResponse) SetRequestId(v string) *AddEndPointResponse {
	s.RequestId = &v
	return s
}

func (s *AddEndPointResponse) SetRegionId(v string) *AddEndPointResponse {
	s.RegionId = &v
	return s
}

func (s *AddEndPointResponse) SetEndPointInfo(v *AddEndPointResponseEndPointInfo) *AddEndPointResponse {
	s.EndPointInfo = v
	return s
}

type AddEndPointResponseEndPointInfo struct {
	DomainURL   *string `json:"domainURL,omitempty" xml:"domainURL,omitempty" require:"true"`
	Host        *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
	Port        *string `json:"port,omitempty" xml:"port,omitempty" require:"true"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty" require:"true"`
	VpcID       *string `json:"vpcID,omitempty" xml:"vpcID,omitempty" require:"true"`
	VSwitch     *string `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" require:"true"`
	Zone        *string `json:"zone,omitempty" xml:"zone,omitempty" require:"true"`
	AllowIP     *string `json:"allowIP,omitempty" xml:"allowIP,omitempty" require:"true"`
	EndPointID  *string `json:"endPointID,omitempty" xml:"endPointID,omitempty" require:"true"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	Product     *string `json:"product,omitempty" xml:"product,omitempty" require:"true"`
}

func (s AddEndPointResponseEndPointInfo) String() string {
	return tea.Prettify(s)
}

func (s AddEndPointResponseEndPointInfo) GoString() string {
	return s.String()
}

func (s *AddEndPointResponseEndPointInfo) SetDomainURL(v string) *AddEndPointResponseEndPointInfo {
	s.DomainURL = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetHost(v string) *AddEndPointResponseEndPointInfo {
	s.Host = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetPort(v string) *AddEndPointResponseEndPointInfo {
	s.Port = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetNetworkType(v string) *AddEndPointResponseEndPointInfo {
	s.NetworkType = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetVpcID(v string) *AddEndPointResponseEndPointInfo {
	s.VpcID = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetVSwitch(v string) *AddEndPointResponseEndPointInfo {
	s.VSwitch = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetZone(v string) *AddEndPointResponseEndPointInfo {
	s.Zone = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetAllowIP(v string) *AddEndPointResponseEndPointInfo {
	s.AllowIP = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetEndPointID(v string) *AddEndPointResponseEndPointInfo {
	s.EndPointID = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetStatus(v string) *AddEndPointResponseEndPointInfo {
	s.Status = &v
	return s
}

func (s *AddEndPointResponseEndPointInfo) SetProduct(v string) *AddEndPointResponseEndPointInfo {
	s.Product = &v
	return s
}

type GetEndPointRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	EndPointID *string `json:"EndPointID,omitempty" xml:"EndPointID,omitempty" require:"true"`
}

func (s GetEndPointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointRequest) GoString() string {
	return s.String()
}

func (s *GetEndPointRequest) SetRegionId(v string) *GetEndPointRequest {
	s.RegionId = &v
	return s
}

func (s *GetEndPointRequest) SetEndPointID(v string) *GetEndPointRequest {
	s.EndPointID = &v
	return s
}

type GetEndPointResponse struct {
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId     *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	EndPointInfo *GetEndPointResponseEndPointInfo `json:"EndPointInfo,omitempty" xml:"EndPointInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetEndPointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointResponse) GoString() string {
	return s.String()
}

func (s *GetEndPointResponse) SetRequestId(v string) *GetEndPointResponse {
	s.RequestId = &v
	return s
}

func (s *GetEndPointResponse) SetRegionId(v string) *GetEndPointResponse {
	s.RegionId = &v
	return s
}

func (s *GetEndPointResponse) SetEndPointInfo(v *GetEndPointResponseEndPointInfo) *GetEndPointResponse {
	s.EndPointInfo = v
	return s
}

type GetEndPointResponseEndPointInfo struct {
	EndPointID  *string `json:"endPointID,omitempty" xml:"endPointID,omitempty" require:"true"`
	DomainURL   *string `json:"domainURL,omitempty" xml:"domainURL,omitempty" require:"true"`
	Host        *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
	Port        *string `json:"port,omitempty" xml:"port,omitempty" require:"true"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty" require:"true"`
	VpcID       *string `json:"vpcID,omitempty" xml:"vpcID,omitempty" require:"true"`
	VSwitch     *string `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" require:"true"`
	Zone        *string `json:"zone,omitempty" xml:"zone,omitempty" require:"true"`
	AllowIP     *string `json:"allowIP,omitempty" xml:"allowIP,omitempty" require:"true"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	Product     *string `json:"product,omitempty" xml:"product,omitempty" require:"true"`
}

func (s GetEndPointResponseEndPointInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointResponseEndPointInfo) GoString() string {
	return s.String()
}

func (s *GetEndPointResponseEndPointInfo) SetEndPointID(v string) *GetEndPointResponseEndPointInfo {
	s.EndPointID = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetDomainURL(v string) *GetEndPointResponseEndPointInfo {
	s.DomainURL = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetHost(v string) *GetEndPointResponseEndPointInfo {
	s.Host = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetPort(v string) *GetEndPointResponseEndPointInfo {
	s.Port = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetNetworkType(v string) *GetEndPointResponseEndPointInfo {
	s.NetworkType = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetVpcID(v string) *GetEndPointResponseEndPointInfo {
	s.VpcID = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetVSwitch(v string) *GetEndPointResponseEndPointInfo {
	s.VSwitch = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetZone(v string) *GetEndPointResponseEndPointInfo {
	s.Zone = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetAllowIP(v string) *GetEndPointResponseEndPointInfo {
	s.AllowIP = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetStatus(v string) *GetEndPointResponseEndPointInfo {
	s.Status = &v
	return s
}

func (s *GetEndPointResponseEndPointInfo) SetProduct(v string) *GetEndPointResponseEndPointInfo {
	s.Product = &v
	return s
}

type GetEndPointByDomainRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DomainURL *string `json:"DomainURL,omitempty" xml:"DomainURL,omitempty" require:"true"`
	RegionID  *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
}

func (s GetEndPointByDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointByDomainRequest) GoString() string {
	return s.String()
}

func (s *GetEndPointByDomainRequest) SetRegionId(v string) *GetEndPointByDomainRequest {
	s.RegionId = &v
	return s
}

func (s *GetEndPointByDomainRequest) SetDomainURL(v string) *GetEndPointByDomainRequest {
	s.DomainURL = &v
	return s
}

func (s *GetEndPointByDomainRequest) SetRegionID(v string) *GetEndPointByDomainRequest {
	s.RegionID = &v
	return s
}

type GetEndPointByDomainResponse struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId     *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	EndPointInfo *GetEndPointByDomainResponseEndPointInfo `json:"EndPointInfo,omitempty" xml:"EndPointInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetEndPointByDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointByDomainResponse) GoString() string {
	return s.String()
}

func (s *GetEndPointByDomainResponse) SetRequestId(v string) *GetEndPointByDomainResponse {
	s.RequestId = &v
	return s
}

func (s *GetEndPointByDomainResponse) SetRegionId(v string) *GetEndPointByDomainResponse {
	s.RegionId = &v
	return s
}

func (s *GetEndPointByDomainResponse) SetEndPointInfo(v *GetEndPointByDomainResponseEndPointInfo) *GetEndPointByDomainResponse {
	s.EndPointInfo = v
	return s
}

type GetEndPointByDomainResponseEndPointInfo struct {
	EndPointID      *string `json:"endPointID,omitempty" xml:"endPointID,omitempty" require:"true"`
	DomainURL       *string `json:"domainURL,omitempty" xml:"domainURL,omitempty" require:"true"`
	Host            *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
	Port            *string `json:"port,omitempty" xml:"port,omitempty" require:"true"`
	NetworkType     *string `json:"networkType,omitempty" xml:"networkType,omitempty" require:"true"`
	VpcID           *string `json:"vpcID,omitempty" xml:"vpcID,omitempty" require:"true"`
	VSwitch         *string `json:"vSwitch,omitempty" xml:"vSwitch,omitempty" require:"true"`
	Zone            *string `json:"zone,omitempty" xml:"zone,omitempty" require:"true"`
	AllowIP         *string `json:"allowIP,omitempty" xml:"allowIP,omitempty" require:"true"`
	CloudInstanceID *string `json:"cloudInstanceID,omitempty" xml:"cloudInstanceID,omitempty" require:"true"`
	Product         *string `json:"product,omitempty" xml:"product,omitempty" require:"true"`
}

func (s GetEndPointByDomainResponseEndPointInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEndPointByDomainResponseEndPointInfo) GoString() string {
	return s.String()
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetEndPointID(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.EndPointID = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetDomainURL(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.DomainURL = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetHost(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.Host = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetPort(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.Port = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetNetworkType(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.NetworkType = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetVpcID(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.VpcID = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetVSwitch(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.VSwitch = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetZone(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.Zone = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetAllowIP(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.AllowIP = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetCloudInstanceID(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.CloudInstanceID = &v
	return s
}

func (s *GetEndPointByDomainResponseEndPointInfo) SetProduct(v string) *GetEndPointByDomainResponseEndPointInfo {
	s.Product = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"cn-beijing":                  tea.String("openanalytics.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("openanalytics.cn-zhangjiakou.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("openanalytics.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":                 tea.String("openanalytics.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("openanalytics.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":                 tea.String("openanalytics.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1":              tea.String("openanalytics.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("datalakeanalytics.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":              tea.String("openanalytics.ap-southeast-3.aliyuncs.com"),
		"ap-northeast-1":              tea.String("datalakeanalytics.ap-northeast-1.aliyuncs.com"),
		"eu-west-1":                   tea.String("openanalytics.eu-west-1.aliyuncs.com"),
		"us-west-1":                   tea.String("openanalytics.us-west-1.aliyuncs.com"),
		"us-east-1":                   tea.String("datalakeanalytics.us-east-1.aliyuncs.com"),
		"eu-central-1":                tea.String("datalakeanalytics.eu-central-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("openanalytics.ap-south-1.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("openanalytics.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-5":              tea.String("openanalytics.ap-southeast-5.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("openanalytics.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("openanalytics.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("openanalytics.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-chengdu":                  tea.String("openanalytics.aliyuncs.com"),
		"cn-edge-1":                   tea.String("openanalytics.aliyuncs.com"),
		"cn-fujian":                   tea.String("openanalytics.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("openanalytics.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("openanalytics.aliyuncs.com"),
		"cn-huhehaote":                tea.String("openanalytics.cn-huhehaote.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("openanalytics.aliyuncs.com"),
		"cn-qingdao":                  tea.String("openanalytics.cn-qingdao.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("openanalytics.aliyuncs.com"),
		"cn-wuhan":                    tea.String("openanalytics.aliyuncs.com"),
		"cn-yushanfang":               tea.String("openanalytics.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("openanalytics.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("openanalytics.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("openanalytics.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("openanalytics.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("openanalytics.me-east-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("openanalytics.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("openanalytics-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) ValidateVirtualClusterNameWithOptions(request *ValidateVirtualClusterNameRequest, runtime *util.RuntimeOptions) (_result *ValidateVirtualClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ValidateVirtualClusterNameResponse{}
	_body, _err := client.DoRequest(tea.String("ValidateVirtualClusterName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateVirtualClusterName(request *ValidateVirtualClusterNameRequest) (_result *ValidateVirtualClusterNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateVirtualClusterNameResponse{}
	_body, _err := client.ValidateVirtualClusterNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeInstanceWithOptions(request *UpgradeInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("UpgradeInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeInstance(request *UpgradeInstanceRequest) (_result *UpgradeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.UpgradeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("CreateInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("ReleaseInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobDetailWithOptions(request *GetJobDetailRequest, runtime *util.RuntimeOptions) (_result *GetJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetJobDetailResponse{}
	_body, _err := client.DoRequest(tea.String("GetJobDetail"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobDetail(request *GetJobDetailRequest) (_result *GetJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobDetailResponse{}
	_body, _err := client.GetJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobStatusWithOptions(request *GetJobStatusRequest, runtime *util.RuntimeOptions) (_result *GetJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetJobStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetJobStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobStatus(request *GetJobStatusRequest) (_result *GetJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobStatusResponse{}
	_body, _err := client.GetJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillSparkJobWithOptions(request *KillSparkJobRequest, runtime *util.RuntimeOptions) (_result *KillSparkJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &KillSparkJobResponse{}
	_body, _err := client.DoRequest(tea.String("KillSparkJob"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillSparkJob(request *KillSparkJobRequest) (_result *KillSparkJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillSparkJobResponse{}
	_body, _err := client.KillSparkJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSparkJobWithOptions(request *SubmitSparkJobRequest, runtime *util.RuntimeOptions) (_result *SubmitSparkJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SubmitSparkJobResponse{}
	_body, _err := client.DoRequest(tea.String("SubmitSparkJob"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSparkJob(request *SubmitSparkJobRequest) (_result *SubmitSparkJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSparkJobResponse{}
	_body, _err := client.SubmitSparkJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobLogWithOptions(request *GetJobLogRequest, runtime *util.RuntimeOptions) (_result *GetJobLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetJobLogResponse{}
	_body, _err := client.DoRequest(tea.String("GetJobLog"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobLog(request *GetJobLogRequest) (_result *GetJobLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobLogResponse{}
	_body, _err := client.GetJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSparkJobWithOptions(request *ListSparkJobRequest, runtime *util.RuntimeOptions) (_result *ListSparkJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListSparkJobResponse{}
	_body, _err := client.DoRequest(tea.String("ListSparkJob"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSparkJob(request *ListSparkJobRequest) (_result *ListSparkJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSparkJobResponse{}
	_body, _err := client.ListSparkJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllowIPWithOptions(request *GetAllowIPRequest, runtime *util.RuntimeOptions) (_result *GetAllowIPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAllowIPResponse{}
	_body, _err := client.DoRequest(tea.String("GetAllowIP"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllowIP(request *GetAllowIPRequest) (_result *GetAllowIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllowIPResponse{}
	_body, _err := client.GetAllowIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAllowIPWithOptions(request *SetAllowIPRequest, runtime *util.RuntimeOptions) (_result *SetAllowIPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetAllowIPResponse{}
	_body, _err := client.DoRequest(tea.String("SetAllowIP"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAllowIP(request *SetAllowIPRequest) (_result *SetAllowIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAllowIPResponse{}
	_body, _err := client.SetAllowIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetMainPasswordWithOptions(request *ResetMainPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetMainPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResetMainPasswordResponse{}
	_body, _err := client.DoRequest(tea.String("ResetMainPassword"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetMainPassword(request *ResetMainPasswordRequest) (_result *ResetMainPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetMainPasswordResponse{}
	_body, _err := client.ResetMainPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnSubscribeRegionWithOptions(request *UnSubscribeRegionRequest, runtime *util.RuntimeOptions) (_result *UnSubscribeRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnSubscribeRegionResponse{}
	_body, _err := client.DoRequest(tea.String("UnSubscribeRegion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnSubscribeRegion(request *UnSubscribeRegionRequest) (_result *UnSubscribeRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnSubscribeRegionResponse{}
	_body, _err := client.UnSubscribeRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAccountListWithOptions(request *QueryAccountListRequest, runtime *util.RuntimeOptions) (_result *QueryAccountListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryAccountListResponse{}
	_body, _err := client.DoRequest(tea.String("QueryAccountList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAccountList(request *QueryAccountListRequest) (_result *QueryAccountListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccountListResponse{}
	_body, _err := client.QueryAccountListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAccountWithOptions(request *AddAccountRequest, runtime *util.RuntimeOptions) (_result *AddAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddAccountResponse{}
	_body, _err := client.DoRequest(tea.String("AddAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAccount(request *AddAccountRequest) (_result *AddAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAccountResponse{}
	_body, _err := client.AddAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAccountPasswordWithOptions(request *UpdateAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *UpdateAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateAccountPasswordResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateAccountPassword"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAccountPassword(request *UpdateAccountPasswordRequest) (_result *UpdateAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccountPasswordResponse{}
	_body, _err := client.UpdateAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveEndPointWithOptions(request *RemoveEndPointRequest, runtime *util.RuntimeOptions) (_result *RemoveEndPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RemoveEndPointResponse{}
	_body, _err := client.DoRequest(tea.String("RemoveEndPoint"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveEndPoint(request *RemoveEndPointRequest) (_result *RemoveEndPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveEndPointResponse{}
	_body, _err := client.RemoveEndPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEndPointListWithOptions(request *QueryEndPointListRequest, runtime *util.RuntimeOptions) (_result *QueryEndPointListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryEndPointListResponse{}
	_body, _err := client.DoRequest(tea.String("QueryEndPointList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEndPointList(request *QueryEndPointListRequest) (_result *QueryEndPointListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEndPointListResponse{}
	_body, _err := client.QueryEndPointListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEndPointWithOptions(request *AddEndPointRequest, runtime *util.RuntimeOptions) (_result *AddEndPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddEndPointResponse{}
	_body, _err := client.DoRequest(tea.String("AddEndPoint"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEndPoint(request *AddEndPointRequest) (_result *AddEndPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEndPointResponse{}
	_body, _err := client.AddEndPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndPointWithOptions(request *GetEndPointRequest, runtime *util.RuntimeOptions) (_result *GetEndPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetEndPointResponse{}
	_body, _err := client.DoRequest(tea.String("GetEndPoint"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEndPoint(request *GetEndPointRequest) (_result *GetEndPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEndPointResponse{}
	_body, _err := client.GetEndPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndPointByDomainWithOptions(request *GetEndPointByDomainRequest, runtime *util.RuntimeOptions) (_result *GetEndPointByDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetEndPointByDomainResponse{}
	_body, _err := client.DoRequest(tea.String("GetEndPointByDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-06-19"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEndPointByDomain(request *GetEndPointByDomainRequest) (_result *GetEndPointByDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEndPointByDomainResponse{}
	_body, _err := client.GetEndPointByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
