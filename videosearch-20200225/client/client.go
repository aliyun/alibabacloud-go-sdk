// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type AddDeletionAudioTaskRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AudioId     *string `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddDeletionAudioTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDeletionAudioTaskRequest) GoString() string {
	return s.String()
}

func (s *AddDeletionAudioTaskRequest) SetClientToken(v string) *AddDeletionAudioTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDeletionAudioTaskRequest) SetAudioId(v string) *AddDeletionAudioTaskRequest {
	s.AudioId = &v
	return s
}

func (s *AddDeletionAudioTaskRequest) SetInstanceId(v string) *AddDeletionAudioTaskRequest {
	s.InstanceId = &v
	return s
}

type AddDeletionAudioTaskResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s AddDeletionAudioTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDeletionAudioTaskResponse) GoString() string {
	return s.String()
}

func (s *AddDeletionAudioTaskResponse) SetRequestId(v string) *AddDeletionAudioTaskResponse {
	s.RequestId = &v
	return s
}

func (s *AddDeletionAudioTaskResponse) SetData(v bool) *AddDeletionAudioTaskResponse {
	s.Data = &v
	return s
}

type GetAudioTaskStatusRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s GetAudioTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAudioTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAudioTaskStatusRequest) SetClientToken(v string) *GetAudioTaskStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAudioTaskStatusRequest) SetTaskId(v string) *GetAudioTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetAudioTaskStatusRequest) SetInstanceId(v string) *GetAudioTaskStatusRequest {
	s.InstanceId = &v
	return s
}

type GetAudioTaskStatusResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *int                                `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	TaskInfo  *GetAudioTaskStatusResponseTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetAudioTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAudioTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAudioTaskStatusResponse) SetRequestId(v string) *GetAudioTaskStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetAudioTaskStatusResponse) SetData(v int) *GetAudioTaskStatusResponse {
	s.Data = &v
	return s
}

func (s *GetAudioTaskStatusResponse) SetTaskInfo(v *GetAudioTaskStatusResponseTaskInfo) *GetAudioTaskStatusResponse {
	s.TaskInfo = v
	return s
}

type GetAudioTaskStatusResponseTaskInfo struct {
	AnalysisUseTime         *int64   `json:"AnalysisUseTime,omitempty" xml:"AnalysisUseTime,omitempty" require:"true"`
	Duration                *float32 `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	ProcessResultOss        *string  `json:"ProcessResultOss,omitempty" xml:"ProcessResultOss,omitempty" require:"true"`
	Status                  *int     `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	SubmitTime              *int64   `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty" require:"true"`
	FinishTime              *int64   `json:"FinishTime,omitempty" xml:"FinishTime,omitempty" require:"true"`
	TaskId                  *int64   `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	ErrorInfo               *string  `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" require:"true"`
	StorageInfo             *int     `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty" require:"true"`
	Description             *string  `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AudioId                 *string  `json:"AudioId,omitempty" xml:"AudioId,omitempty" require:"true"`
	AudioTags               *string  `json:"AudioTags,omitempty" xml:"AudioTags,omitempty" require:"true"`
	AudioUrl                *string  `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty" require:"true"`
	QueryTags               *string  `json:"QueryTags,omitempty" xml:"QueryTags,omitempty" require:"true"`
	ResourceType            *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ReplaceStorageThreshold *string  `json:"ReplaceStorageThreshold,omitempty" xml:"ReplaceStorageThreshold,omitempty" require:"true"`
}

func (s GetAudioTaskStatusResponseTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAudioTaskStatusResponseTaskInfo) GoString() string {
	return s.String()
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetAnalysisUseTime(v int64) *GetAudioTaskStatusResponseTaskInfo {
	s.AnalysisUseTime = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetDuration(v float32) *GetAudioTaskStatusResponseTaskInfo {
	s.Duration = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetProcessResultOss(v string) *GetAudioTaskStatusResponseTaskInfo {
	s.ProcessResultOss = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetStatus(v int) *GetAudioTaskStatusResponseTaskInfo {
	s.Status = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetSubmitTime(v int64) *GetAudioTaskStatusResponseTaskInfo {
	s.SubmitTime = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetFinishTime(v int64) *GetAudioTaskStatusResponseTaskInfo {
	s.FinishTime = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetTaskId(v int64) *GetAudioTaskStatusResponseTaskInfo {
	s.TaskId = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetErrorInfo(v string) *GetAudioTaskStatusResponseTaskInfo {
	s.ErrorInfo = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetStorageInfo(v int) *GetAudioTaskStatusResponseTaskInfo {
	s.StorageInfo = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetDescription(v string) *GetAudioTaskStatusResponseTaskInfo {
	s.Description = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetAudioId(v string) *GetAudioTaskStatusResponseTaskInfo {
	s.AudioId = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetAudioTags(v string) *GetAudioTaskStatusResponseTaskInfo {
	s.AudioTags = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetAudioUrl(v string) *GetAudioTaskStatusResponseTaskInfo {
	s.AudioUrl = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetQueryTags(v string) *GetAudioTaskStatusResponseTaskInfo {
	s.QueryTags = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetResourceType(v string) *GetAudioTaskStatusResponseTaskInfo {
	s.ResourceType = &v
	return s
}

func (s *GetAudioTaskStatusResponseTaskInfo) SetReplaceStorageThreshold(v string) *GetAudioTaskStatusResponseTaskInfo {
	s.ReplaceStorageThreshold = &v
	return s
}

type CancelBatchTaskRequest struct {
	BatchTaskId *int64 `json:"BatchTaskId,omitempty" xml:"BatchTaskId,omitempty"`
}

func (s CancelBatchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelBatchTaskRequest) SetBatchTaskId(v int64) *CancelBatchTaskRequest {
	s.BatchTaskId = &v
	return s
}

type CancelBatchTaskResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s CancelBatchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelBatchTaskResponse) SetRequestId(v string) *CancelBatchTaskResponse {
	s.RequestId = &v
	return s
}

func (s *CancelBatchTaskResponse) SetData(v bool) *CancelBatchTaskResponse {
	s.Data = &v
	return s
}

type GetAudioStorageHistoryRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	AudioId     *string `json:"AudioId,omitempty" xml:"AudioId,omitempty" require:"true"`
	PageNumber  *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetAudioStorageHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAudioStorageHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetAudioStorageHistoryRequest) SetClientToken(v string) *GetAudioStorageHistoryRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAudioStorageHistoryRequest) SetInstanceId(v string) *GetAudioStorageHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAudioStorageHistoryRequest) SetAudioId(v string) *GetAudioStorageHistoryRequest {
	s.AudioId = &v
	return s
}

func (s *GetAudioStorageHistoryRequest) SetPageNumber(v int) *GetAudioStorageHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAudioStorageHistoryRequest) SetPageSize(v int) *GetAudioStorageHistoryRequest {
	s.PageSize = &v
	return s
}

type GetAudioStorageHistoryResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetAudioStorageHistoryResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetAudioStorageHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAudioStorageHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetAudioStorageHistoryResponse) SetRequestId(v string) *GetAudioStorageHistoryResponse {
	s.RequestId = &v
	return s
}

func (s *GetAudioStorageHistoryResponse) SetData(v *GetAudioStorageHistoryResponseData) *GetAudioStorageHistoryResponse {
	s.Data = v
	return s
}

type GetAudioStorageHistoryResponseData struct {
	PageNumber *int                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Count      *int64                                    `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	List       []*GetAudioStorageHistoryResponseDataList `json:"List,omitempty" xml:"List,omitempty" require:"true" type:"Repeated"`
}

func (s GetAudioStorageHistoryResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetAudioStorageHistoryResponseData) GoString() string {
	return s.String()
}

func (s *GetAudioStorageHistoryResponseData) SetPageNumber(v int) *GetAudioStorageHistoryResponseData {
	s.PageNumber = &v
	return s
}

func (s *GetAudioStorageHistoryResponseData) SetPageSize(v int) *GetAudioStorageHistoryResponseData {
	s.PageSize = &v
	return s
}

func (s *GetAudioStorageHistoryResponseData) SetCount(v int64) *GetAudioStorageHistoryResponseData {
	s.Count = &v
	return s
}

func (s *GetAudioStorageHistoryResponseData) SetList(v []*GetAudioStorageHistoryResponseDataList) *GetAudioStorageHistoryResponseData {
	s.List = v
	return s
}

type GetAudioStorageHistoryResponseDataList struct {
	AudioId      *string `json:"AudioId,omitempty" xml:"AudioId,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	StorageType  *int    `json:"StorageType,omitempty" xml:"StorageType,omitempty" require:"true"`
	StorageInfo  *int    `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty" require:"true"`
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty" require:"true"`
	AudioUrl     *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty" require:"true"`
}

func (s GetAudioStorageHistoryResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetAudioStorageHistoryResponseDataList) GoString() string {
	return s.String()
}

func (s *GetAudioStorageHistoryResponseDataList) SetAudioId(v string) *GetAudioStorageHistoryResponseDataList {
	s.AudioId = &v
	return s
}

func (s *GetAudioStorageHistoryResponseDataList) SetDescription(v string) *GetAudioStorageHistoryResponseDataList {
	s.Description = &v
	return s
}

func (s *GetAudioStorageHistoryResponseDataList) SetStorageType(v int) *GetAudioStorageHistoryResponseDataList {
	s.StorageType = &v
	return s
}

func (s *GetAudioStorageHistoryResponseDataList) SetStorageInfo(v int) *GetAudioStorageHistoryResponseDataList {
	s.StorageInfo = &v
	return s
}

func (s *GetAudioStorageHistoryResponseDataList) SetModifiedTime(v int64) *GetAudioStorageHistoryResponseDataList {
	s.ModifiedTime = &v
	return s
}

func (s *GetAudioStorageHistoryResponseDataList) SetAudioUrl(v string) *GetAudioStorageHistoryResponseDataList {
	s.AudioUrl = &v
	return s
}

type ModifyPriorityRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	Sort        *int    `json:"Sort,omitempty" xml:"Sort,omitempty" require:"true"`
}

func (s ModifyPriorityRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPriorityRequest) GoString() string {
	return s.String()
}

func (s *ModifyPriorityRequest) SetClientToken(v string) *ModifyPriorityRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyPriorityRequest) SetTaskId(v string) *ModifyPriorityRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyPriorityRequest) SetSort(v int) *ModifyPriorityRequest {
	s.Sort = &v
	return s
}

type ModifyPriorityResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s ModifyPriorityResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPriorityResponse) GoString() string {
	return s.String()
}

func (s *ModifyPriorityResponse) SetRequestId(v string) *ModifyPriorityResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyPriorityResponse) SetData(v bool) *ModifyPriorityResponse {
	s.Data = &v
	return s
}

type GetAudioInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s GetAudioInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAudioInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetAudioInstanceRequest) SetClientToken(v string) *GetAudioInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAudioInstanceRequest) SetInstanceId(v string) *GetAudioInstanceRequest {
	s.InstanceId = &v
	return s
}

type GetAudioInstanceResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetAudioInstanceResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetAudioInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAudioInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetAudioInstanceResponse) SetRequestId(v string) *GetAudioInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *GetAudioInstanceResponse) SetData(v *GetAudioInstanceResponseData) *GetAudioInstanceResponse {
	s.Data = v
	return s
}

type GetAudioInstanceResponseData struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceStatus       *int    `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty" require:"true"`
	AudioNumber          *int    `json:"AudioNumber,omitempty" xml:"AudioNumber,omitempty" require:"true"`
	ExpireTime           *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	ConcurrencyNumber    *int    `json:"ConcurrencyNumber,omitempty" xml:"ConcurrencyNumber,omitempty" require:"true"`
	MaxAudioCapacity     *string `json:"MaxAudioCapacity,omitempty" xml:"MaxAudioCapacity,omitempty" require:"true"`
	CurrentAudioCapacity *string `json:"CurrentAudioCapacity,omitempty" xml:"CurrentAudioCapacity,omitempty" require:"true"`
	AudioInfoUpdateTime  *int64  `json:"AudioInfoUpdateTime,omitempty" xml:"AudioInfoUpdateTime,omitempty" require:"true"`
	BundlingType         *string `json:"BundlingType,omitempty" xml:"BundlingType,omitempty" require:"true"`
}

func (s GetAudioInstanceResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetAudioInstanceResponseData) GoString() string {
	return s.String()
}

func (s *GetAudioInstanceResponseData) SetInstanceId(v string) *GetAudioInstanceResponseData {
	s.InstanceId = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetInstanceName(v string) *GetAudioInstanceResponseData {
	s.InstanceName = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetCreateTime(v int64) *GetAudioInstanceResponseData {
	s.CreateTime = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetRegionId(v string) *GetAudioInstanceResponseData {
	s.RegionId = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetInstanceStatus(v int) *GetAudioInstanceResponseData {
	s.InstanceStatus = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetAudioNumber(v int) *GetAudioInstanceResponseData {
	s.AudioNumber = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetExpireTime(v int64) *GetAudioInstanceResponseData {
	s.ExpireTime = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetConcurrencyNumber(v int) *GetAudioInstanceResponseData {
	s.ConcurrencyNumber = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetMaxAudioCapacity(v string) *GetAudioInstanceResponseData {
	s.MaxAudioCapacity = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetCurrentAudioCapacity(v string) *GetAudioInstanceResponseData {
	s.CurrentAudioCapacity = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetAudioInfoUpdateTime(v int64) *GetAudioInstanceResponseData {
	s.AudioInfoUpdateTime = &v
	return s
}

func (s *GetAudioInstanceResponseData) SetBundlingType(v string) *GetAudioInstanceResponseData {
	s.BundlingType = &v
	return s
}

type GetBatchTaskRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	BatchTaskId *int64  `json:"BatchTaskId,omitempty" xml:"BatchTaskId,omitempty" require:"true"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s GetBatchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTaskRequest) SetClientToken(v string) *GetBatchTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBatchTaskRequest) SetBatchTaskId(v int64) *GetBatchTaskRequest {
	s.BatchTaskId = &v
	return s
}

func (s *GetBatchTaskRequest) SetInstanceId(v string) *GetBatchTaskRequest {
	s.InstanceId = &v
	return s
}

type GetBatchTaskResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BatchTaskId    *int64  `json:"BatchTaskId,omitempty" xml:"BatchTaskId,omitempty" require:"true"`
	Status         *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifiedTime   *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty" require:"true"`
	ProcessMessage *string `json:"ProcessMessage,omitempty" xml:"ProcessMessage,omitempty" require:"true"`
	SubTaskDetail  *string `json:"SubTaskDetail,omitempty" xml:"SubTaskDetail,omitempty" require:"true"`
}

func (s GetBatchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTaskResponse) SetRequestId(v string) *GetBatchTaskResponse {
	s.RequestId = &v
	return s
}

func (s *GetBatchTaskResponse) SetBatchTaskId(v int64) *GetBatchTaskResponse {
	s.BatchTaskId = &v
	return s
}

func (s *GetBatchTaskResponse) SetStatus(v int) *GetBatchTaskResponse {
	s.Status = &v
	return s
}

func (s *GetBatchTaskResponse) SetModifiedTime(v int64) *GetBatchTaskResponse {
	s.ModifiedTime = &v
	return s
}

func (s *GetBatchTaskResponse) SetProcessMessage(v string) *GetBatchTaskResponse {
	s.ProcessMessage = &v
	return s
}

func (s *GetBatchTaskResponse) SetSubTaskDetail(v string) *GetBatchTaskResponse {
	s.SubTaskDetail = &v
	return s
}

type AddSearchAudioTaskRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AudioUrl          *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	ReturnAudioNumber *int    `json:"ReturnAudioNumber,omitempty" xml:"ReturnAudioNumber,omitempty"`
	QueryTags         *string `json:"QueryTags,omitempty" xml:"QueryTags,omitempty"`
	CallbackUrl       *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ContentSource     *int    `json:"ContentSource,omitempty" xml:"ContentSource,omitempty"`
	AudioFile         *string `json:"AudioFile,omitempty" xml:"AudioFile,omitempty"`
	Sort              *int    `json:"Sort,omitempty" xml:"Sort,omitempty"`
	NeedFeatureFile   *int    `json:"NeedFeatureFile,omitempty" xml:"NeedFeatureFile,omitempty"`
	ResourceType      *int    `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s AddSearchAudioTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSearchAudioTaskRequest) GoString() string {
	return s.String()
}

func (s *AddSearchAudioTaskRequest) SetClientToken(v string) *AddSearchAudioTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetAudioUrl(v string) *AddSearchAudioTaskRequest {
	s.AudioUrl = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetReturnAudioNumber(v int) *AddSearchAudioTaskRequest {
	s.ReturnAudioNumber = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetQueryTags(v string) *AddSearchAudioTaskRequest {
	s.QueryTags = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetCallbackUrl(v string) *AddSearchAudioTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetInstanceId(v string) *AddSearchAudioTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetDescription(v string) *AddSearchAudioTaskRequest {
	s.Description = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetContentSource(v int) *AddSearchAudioTaskRequest {
	s.ContentSource = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetAudioFile(v string) *AddSearchAudioTaskRequest {
	s.AudioFile = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetSort(v int) *AddSearchAudioTaskRequest {
	s.Sort = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetNeedFeatureFile(v int) *AddSearchAudioTaskRequest {
	s.NeedFeatureFile = &v
	return s
}

func (s *AddSearchAudioTaskRequest) SetResourceType(v int) *AddSearchAudioTaskRequest {
	s.ResourceType = &v
	return s
}

type AddSearchAudioTaskResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *AddSearchAudioTaskResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AddSearchAudioTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSearchAudioTaskResponse) GoString() string {
	return s.String()
}

func (s *AddSearchAudioTaskResponse) SetRequestId(v string) *AddSearchAudioTaskResponse {
	s.RequestId = &v
	return s
}

func (s *AddSearchAudioTaskResponse) SetData(v *AddSearchAudioTaskResponseData) *AddSearchAudioTaskResponse {
	s.Data = v
	return s
}

type AddSearchAudioTaskResponseData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s AddSearchAudioTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddSearchAudioTaskResponseData) GoString() string {
	return s.String()
}

func (s *AddSearchAudioTaskResponseData) SetTaskId(v string) *AddSearchAudioTaskResponseData {
	s.TaskId = &v
	return s
}

type AddSearchAudioTaskAdvanceRequest struct {
	AudioFileObject   io.Reader `json:"AudioFileObject,omitempty" xml:"AudioFileObject,omitempty" require:"true"`
	ClientToken       *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AudioUrl          *string   `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	ReturnAudioNumber *int      `json:"ReturnAudioNumber,omitempty" xml:"ReturnAudioNumber,omitempty"`
	QueryTags         *string   `json:"QueryTags,omitempty" xml:"QueryTags,omitempty"`
	CallbackUrl       *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	InstanceId        *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Description       *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ContentSource     *int      `json:"ContentSource,omitempty" xml:"ContentSource,omitempty"`
	Sort              *int      `json:"Sort,omitempty" xml:"Sort,omitempty"`
	NeedFeatureFile   *int      `json:"NeedFeatureFile,omitempty" xml:"NeedFeatureFile,omitempty"`
	ResourceType      *int      `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s AddSearchAudioTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSearchAudioTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddSearchAudioTaskAdvanceRequest) SetAudioFileObject(v io.Reader) *AddSearchAudioTaskAdvanceRequest {
	s.AudioFileObject = v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetClientToken(v string) *AddSearchAudioTaskAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetAudioUrl(v string) *AddSearchAudioTaskAdvanceRequest {
	s.AudioUrl = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetReturnAudioNumber(v int) *AddSearchAudioTaskAdvanceRequest {
	s.ReturnAudioNumber = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetQueryTags(v string) *AddSearchAudioTaskAdvanceRequest {
	s.QueryTags = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetCallbackUrl(v string) *AddSearchAudioTaskAdvanceRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetInstanceId(v string) *AddSearchAudioTaskAdvanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetDescription(v string) *AddSearchAudioTaskAdvanceRequest {
	s.Description = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetContentSource(v int) *AddSearchAudioTaskAdvanceRequest {
	s.ContentSource = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetSort(v int) *AddSearchAudioTaskAdvanceRequest {
	s.Sort = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetNeedFeatureFile(v int) *AddSearchAudioTaskAdvanceRequest {
	s.NeedFeatureFile = &v
	return s
}

func (s *AddSearchAudioTaskAdvanceRequest) SetResourceType(v int) *AddSearchAudioTaskAdvanceRequest {
	s.ResourceType = &v
	return s
}

type AddStorageAudioTaskRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	AudioUrl      *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	AudioId       *string `json:"AudioId,omitempty" xml:"AudioId,omitempty" require:"true"`
	AudioTags     *string `json:"AudioTags,omitempty" xml:"AudioTags,omitempty"`
	CallbackUrl   *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContentSource *int    `json:"ContentSource,omitempty" xml:"ContentSource,omitempty"`
	AudioFile     *string `json:"AudioFile,omitempty" xml:"AudioFile,omitempty"`
	Sort          *int    `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s AddStorageAudioTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AddStorageAudioTaskRequest) GoString() string {
	return s.String()
}

func (s *AddStorageAudioTaskRequest) SetInstanceId(v string) *AddStorageAudioTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *AddStorageAudioTaskRequest) SetAudioUrl(v string) *AddStorageAudioTaskRequest {
	s.AudioUrl = &v
	return s
}

func (s *AddStorageAudioTaskRequest) SetAudioId(v string) *AddStorageAudioTaskRequest {
	s.AudioId = &v
	return s
}

func (s *AddStorageAudioTaskRequest) SetAudioTags(v string) *AddStorageAudioTaskRequest {
	s.AudioTags = &v
	return s
}

func (s *AddStorageAudioTaskRequest) SetCallbackUrl(v string) *AddStorageAudioTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddStorageAudioTaskRequest) SetDescription(v string) *AddStorageAudioTaskRequest {
	s.Description = &v
	return s
}

func (s *AddStorageAudioTaskRequest) SetClientToken(v string) *AddStorageAudioTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *AddStorageAudioTaskRequest) SetContentSource(v int) *AddStorageAudioTaskRequest {
	s.ContentSource = &v
	return s
}

func (s *AddStorageAudioTaskRequest) SetAudioFile(v string) *AddStorageAudioTaskRequest {
	s.AudioFile = &v
	return s
}

func (s *AddStorageAudioTaskRequest) SetSort(v int) *AddStorageAudioTaskRequest {
	s.Sort = &v
	return s
}

type AddStorageAudioTaskResponse struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *AddStorageAudioTaskResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AddStorageAudioTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AddStorageAudioTaskResponse) GoString() string {
	return s.String()
}

func (s *AddStorageAudioTaskResponse) SetRequestId(v string) *AddStorageAudioTaskResponse {
	s.RequestId = &v
	return s
}

func (s *AddStorageAudioTaskResponse) SetData(v *AddStorageAudioTaskResponseData) *AddStorageAudioTaskResponse {
	s.Data = v
	return s
}

type AddStorageAudioTaskResponseData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s AddStorageAudioTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddStorageAudioTaskResponseData) GoString() string {
	return s.String()
}

func (s *AddStorageAudioTaskResponseData) SetTaskId(v string) *AddStorageAudioTaskResponseData {
	s.TaskId = &v
	return s
}

type AddStorageAudioTaskAdvanceRequest struct {
	AudioFileObject io.Reader `json:"AudioFileObject,omitempty" xml:"AudioFileObject,omitempty" require:"true"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	AudioUrl        *string   `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty"`
	AudioId         *string   `json:"AudioId,omitempty" xml:"AudioId,omitempty" require:"true"`
	AudioTags       *string   `json:"AudioTags,omitempty" xml:"AudioTags,omitempty"`
	CallbackUrl     *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ClientToken     *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContentSource   *int      `json:"ContentSource,omitempty" xml:"ContentSource,omitempty"`
	Sort            *int      `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s AddStorageAudioTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddStorageAudioTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddStorageAudioTaskAdvanceRequest) SetAudioFileObject(v io.Reader) *AddStorageAudioTaskAdvanceRequest {
	s.AudioFileObject = v
	return s
}

func (s *AddStorageAudioTaskAdvanceRequest) SetInstanceId(v string) *AddStorageAudioTaskAdvanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AddStorageAudioTaskAdvanceRequest) SetAudioUrl(v string) *AddStorageAudioTaskAdvanceRequest {
	s.AudioUrl = &v
	return s
}

func (s *AddStorageAudioTaskAdvanceRequest) SetAudioId(v string) *AddStorageAudioTaskAdvanceRequest {
	s.AudioId = &v
	return s
}

func (s *AddStorageAudioTaskAdvanceRequest) SetAudioTags(v string) *AddStorageAudioTaskAdvanceRequest {
	s.AudioTags = &v
	return s
}

func (s *AddStorageAudioTaskAdvanceRequest) SetCallbackUrl(v string) *AddStorageAudioTaskAdvanceRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddStorageAudioTaskAdvanceRequest) SetDescription(v string) *AddStorageAudioTaskAdvanceRequest {
	s.Description = &v
	return s
}

func (s *AddStorageAudioTaskAdvanceRequest) SetClientToken(v string) *AddStorageAudioTaskAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *AddStorageAudioTaskAdvanceRequest) SetContentSource(v int) *AddStorageAudioTaskAdvanceRequest {
	s.ContentSource = &v
	return s
}

func (s *AddStorageAudioTaskAdvanceRequest) SetSort(v int) *AddStorageAudioTaskAdvanceRequest {
	s.Sort = &v
	return s
}

type ListStorageAudioTasksRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	AudioId         *string `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	PageNumber      *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	StatusList      *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StorageInfoList *string `json:"StorageInfoList,omitempty" xml:"StorageInfoList,omitempty"`
	SortList        *string `json:"SortList,omitempty" xml:"SortList,omitempty"`
}

func (s ListStorageAudioTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStorageAudioTasksRequest) GoString() string {
	return s.String()
}

func (s *ListStorageAudioTasksRequest) SetClientToken(v string) *ListStorageAudioTasksRequest {
	s.ClientToken = &v
	return s
}

func (s *ListStorageAudioTasksRequest) SetTaskId(v string) *ListStorageAudioTasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListStorageAudioTasksRequest) SetAudioId(v string) *ListStorageAudioTasksRequest {
	s.AudioId = &v
	return s
}

func (s *ListStorageAudioTasksRequest) SetPageNumber(v int) *ListStorageAudioTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStorageAudioTasksRequest) SetPageSize(v int) *ListStorageAudioTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListStorageAudioTasksRequest) SetInstanceId(v string) *ListStorageAudioTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListStorageAudioTasksRequest) SetStatusList(v string) *ListStorageAudioTasksRequest {
	s.StatusList = &v
	return s
}

func (s *ListStorageAudioTasksRequest) SetDescription(v string) *ListStorageAudioTasksRequest {
	s.Description = &v
	return s
}

func (s *ListStorageAudioTasksRequest) SetStorageInfoList(v string) *ListStorageAudioTasksRequest {
	s.StorageInfoList = &v
	return s
}

func (s *ListStorageAudioTasksRequest) SetSortList(v string) *ListStorageAudioTasksRequest {
	s.SortList = &v
	return s
}

type ListStorageAudioTasksResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListStorageAudioTasksResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListStorageAudioTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStorageAudioTasksResponse) GoString() string {
	return s.String()
}

func (s *ListStorageAudioTasksResponse) SetRequestId(v string) *ListStorageAudioTasksResponse {
	s.RequestId = &v
	return s
}

func (s *ListStorageAudioTasksResponse) SetData(v *ListStorageAudioTasksResponseData) *ListStorageAudioTasksResponse {
	s.Data = v
	return s
}

type ListStorageAudioTasksResponseData struct {
	PageNumber *int                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Count      *int64                                       `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	TaskList   []*ListStorageAudioTasksResponseDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" require:"true" type:"Repeated"`
}

func (s ListStorageAudioTasksResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListStorageAudioTasksResponseData) GoString() string {
	return s.String()
}

func (s *ListStorageAudioTasksResponseData) SetPageNumber(v int) *ListStorageAudioTasksResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListStorageAudioTasksResponseData) SetPageSize(v int) *ListStorageAudioTasksResponseData {
	s.PageSize = &v
	return s
}

func (s *ListStorageAudioTasksResponseData) SetCount(v int64) *ListStorageAudioTasksResponseData {
	s.Count = &v
	return s
}

func (s *ListStorageAudioTasksResponseData) SetTaskList(v []*ListStorageAudioTasksResponseDataTaskList) *ListStorageAudioTasksResponseData {
	s.TaskList = v
	return s
}

type ListStorageAudioTasksResponseDataTaskList struct {
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	AudioId      *string `json:"AudioId,omitempty" xml:"AudioId,omitempty" require:"true"`
	ProcessTime  *int64  `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty" require:"true"`
	StorageInfo  *int    `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty" require:"true"`
	UpdateTime   *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ErrorDetail  *string `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" require:"true"`
	RemoteTaskId *string `json:"RemoteTaskId,omitempty" xml:"RemoteTaskId,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AudioUrl     *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty" require:"true"`
	Sort         *int    `json:"Sort,omitempty" xml:"Sort,omitempty" require:"true"`
}

func (s ListStorageAudioTasksResponseDataTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListStorageAudioTasksResponseDataTaskList) GoString() string {
	return s.String()
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetTaskId(v int64) *ListStorageAudioTasksResponseDataTaskList {
	s.TaskId = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetAudioId(v string) *ListStorageAudioTasksResponseDataTaskList {
	s.AudioId = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetProcessTime(v int64) *ListStorageAudioTasksResponseDataTaskList {
	s.ProcessTime = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetStorageInfo(v int) *ListStorageAudioTasksResponseDataTaskList {
	s.StorageInfo = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetUpdateTime(v int64) *ListStorageAudioTasksResponseDataTaskList {
	s.UpdateTime = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetStatus(v string) *ListStorageAudioTasksResponseDataTaskList {
	s.Status = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetErrorDetail(v string) *ListStorageAudioTasksResponseDataTaskList {
	s.ErrorDetail = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetRemoteTaskId(v string) *ListStorageAudioTasksResponseDataTaskList {
	s.RemoteTaskId = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetDescription(v string) *ListStorageAudioTasksResponseDataTaskList {
	s.Description = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetAudioUrl(v string) *ListStorageAudioTasksResponseDataTaskList {
	s.AudioUrl = &v
	return s
}

func (s *ListStorageAudioTasksResponseDataTaskList) SetSort(v int) *ListStorageAudioTasksResponseDataTaskList {
	s.Sort = &v
	return s
}

type ListSearchAudioTasksRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	PageNumber  *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	StatusList  *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SortList    *string `json:"SortList,omitempty" xml:"SortList,omitempty"`
}

func (s ListSearchAudioTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSearchAudioTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSearchAudioTasksRequest) SetClientToken(v string) *ListSearchAudioTasksRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSearchAudioTasksRequest) SetTaskId(v string) *ListSearchAudioTasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListSearchAudioTasksRequest) SetPageNumber(v int) *ListSearchAudioTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSearchAudioTasksRequest) SetPageSize(v int) *ListSearchAudioTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchAudioTasksRequest) SetInstanceId(v string) *ListSearchAudioTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSearchAudioTasksRequest) SetStatusList(v string) *ListSearchAudioTasksRequest {
	s.StatusList = &v
	return s
}

func (s *ListSearchAudioTasksRequest) SetDescription(v string) *ListSearchAudioTasksRequest {
	s.Description = &v
	return s
}

func (s *ListSearchAudioTasksRequest) SetSortList(v string) *ListSearchAudioTasksRequest {
	s.SortList = &v
	return s
}

type ListSearchAudioTasksResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListSearchAudioTasksResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListSearchAudioTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSearchAudioTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSearchAudioTasksResponse) SetRequestId(v string) *ListSearchAudioTasksResponse {
	s.RequestId = &v
	return s
}

func (s *ListSearchAudioTasksResponse) SetData(v *ListSearchAudioTasksResponseData) *ListSearchAudioTasksResponse {
	s.Data = v
	return s
}

type ListSearchAudioTasksResponseData struct {
	Count      *int64                                      `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	PageNumber *int                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TaskList   []*ListSearchAudioTasksResponseDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSearchAudioTasksResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListSearchAudioTasksResponseData) GoString() string {
	return s.String()
}

func (s *ListSearchAudioTasksResponseData) SetCount(v int64) *ListSearchAudioTasksResponseData {
	s.Count = &v
	return s
}

func (s *ListSearchAudioTasksResponseData) SetPageNumber(v int) *ListSearchAudioTasksResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListSearchAudioTasksResponseData) SetPageSize(v int) *ListSearchAudioTasksResponseData {
	s.PageSize = &v
	return s
}

func (s *ListSearchAudioTasksResponseData) SetTaskList(v []*ListSearchAudioTasksResponseDataTaskList) *ListSearchAudioTasksResponseData {
	s.TaskList = v
	return s
}

type ListSearchAudioTasksResponseDataTaskList struct {
	TaskId           *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	AudioId          *string `json:"AudioId,omitempty" xml:"AudioId,omitempty" require:"true"`
	ProcessTime      *int64  `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty" require:"true"`
	Status           *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifiedTime     *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty" require:"true"`
	ProcessResultUrl *string `json:"ProcessResultUrl,omitempty" xml:"ProcessResultUrl,omitempty" require:"true"`
	StorageInfo      *int    `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty" require:"true"`
	AudioUrl         *string `json:"AudioUrl,omitempty" xml:"AudioUrl,omitempty" require:"true"`
	ErrorDetail      *string `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" require:"true"`
	RemoteTaskId     *string `json:"RemoteTaskId,omitempty" xml:"RemoteTaskId,omitempty" require:"true"`
	AudioTags        *string `json:"AudioTags,omitempty" xml:"AudioTags,omitempty" require:"true"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Sort             *int    `json:"Sort,omitempty" xml:"Sort,omitempty" require:"true"`
}

func (s ListSearchAudioTasksResponseDataTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListSearchAudioTasksResponseDataTaskList) GoString() string {
	return s.String()
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetTaskId(v string) *ListSearchAudioTasksResponseDataTaskList {
	s.TaskId = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetAudioId(v string) *ListSearchAudioTasksResponseDataTaskList {
	s.AudioId = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetProcessTime(v int64) *ListSearchAudioTasksResponseDataTaskList {
	s.ProcessTime = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetStatus(v int) *ListSearchAudioTasksResponseDataTaskList {
	s.Status = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetModifiedTime(v int64) *ListSearchAudioTasksResponseDataTaskList {
	s.ModifiedTime = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetProcessResultUrl(v string) *ListSearchAudioTasksResponseDataTaskList {
	s.ProcessResultUrl = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetStorageInfo(v int) *ListSearchAudioTasksResponseDataTaskList {
	s.StorageInfo = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetAudioUrl(v string) *ListSearchAudioTasksResponseDataTaskList {
	s.AudioUrl = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetErrorDetail(v string) *ListSearchAudioTasksResponseDataTaskList {
	s.ErrorDetail = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetRemoteTaskId(v string) *ListSearchAudioTasksResponseDataTaskList {
	s.RemoteTaskId = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetAudioTags(v string) *ListSearchAudioTasksResponseDataTaskList {
	s.AudioTags = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetDescription(v string) *ListSearchAudioTasksResponseDataTaskList {
	s.Description = &v
	return s
}

func (s *ListSearchAudioTasksResponseDataTaskList) SetSort(v int) *ListSearchAudioTasksResponseDataTaskList {
	s.Sort = &v
	return s
}

type CreateBatchTaskRequest struct {
	BatchTaskType *int    `json:"BatchTaskType,omitempty" xml:"BatchTaskType,omitempty"`
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssDataPath   *string `json:"OssDataPath,omitempty" xml:"OssDataPath,omitempty"`
	OssMetaFile   *string `json:"OssMetaFile,omitempty" xml:"OssMetaFile,omitempty"`
	FileUrl       *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	RoleArn       *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CallbackUrl   *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
}

func (s CreateBatchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskRequest) SetBatchTaskType(v int) *CreateBatchTaskRequest {
	s.BatchTaskType = &v
	return s
}

func (s *CreateBatchTaskRequest) SetOssBucketName(v string) *CreateBatchTaskRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateBatchTaskRequest) SetOssDataPath(v string) *CreateBatchTaskRequest {
	s.OssDataPath = &v
	return s
}

func (s *CreateBatchTaskRequest) SetOssMetaFile(v string) *CreateBatchTaskRequest {
	s.OssMetaFile = &v
	return s
}

func (s *CreateBatchTaskRequest) SetFileUrl(v string) *CreateBatchTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateBatchTaskRequest) SetRoleArn(v string) *CreateBatchTaskRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateBatchTaskRequest) SetInstanceId(v string) *CreateBatchTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBatchTaskRequest) SetClientToken(v string) *CreateBatchTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBatchTaskRequest) SetCallbackUrl(v string) *CreateBatchTaskRequest {
	s.CallbackUrl = &v
	return s
}

type CreateBatchTaskResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s CreateBatchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskResponse) SetRequestId(v string) *CreateBatchTaskResponse {
	s.RequestId = &v
	return s
}

func (s *CreateBatchTaskResponse) SetTaskId(v int64) *CreateBatchTaskResponse {
	s.TaskId = &v
	return s
}

type CreateBatchTaskAdvanceRequest struct {
	FileUrlObject io.Reader `json:"FileUrlObject,omitempty" xml:"FileUrlObject,omitempty" require:"true"`
	BatchTaskType *int      `json:"BatchTaskType,omitempty" xml:"BatchTaskType,omitempty"`
	OssBucketName *string   `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssDataPath   *string   `json:"OssDataPath,omitempty" xml:"OssDataPath,omitempty"`
	OssMetaFile   *string   `json:"OssMetaFile,omitempty" xml:"OssMetaFile,omitempty"`
	RoleArn       *string   `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ClientToken   *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CallbackUrl   *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
}

func (s CreateBatchTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBatchTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskAdvanceRequest) SetFileUrlObject(v io.Reader) *CreateBatchTaskAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *CreateBatchTaskAdvanceRequest) SetBatchTaskType(v int) *CreateBatchTaskAdvanceRequest {
	s.BatchTaskType = &v
	return s
}

func (s *CreateBatchTaskAdvanceRequest) SetOssBucketName(v string) *CreateBatchTaskAdvanceRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateBatchTaskAdvanceRequest) SetOssDataPath(v string) *CreateBatchTaskAdvanceRequest {
	s.OssDataPath = &v
	return s
}

func (s *CreateBatchTaskAdvanceRequest) SetOssMetaFile(v string) *CreateBatchTaskAdvanceRequest {
	s.OssMetaFile = &v
	return s
}

func (s *CreateBatchTaskAdvanceRequest) SetRoleArn(v string) *CreateBatchTaskAdvanceRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateBatchTaskAdvanceRequest) SetInstanceId(v string) *CreateBatchTaskAdvanceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBatchTaskAdvanceRequest) SetClientToken(v string) *CreateBatchTaskAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBatchTaskAdvanceRequest) SetCallbackUrl(v string) *CreateBatchTaskAdvanceRequest {
	s.CallbackUrl = &v
	return s
}

type GetStorageHistoryRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VideoId     *string `json:"VideoId,omitempty" xml:"VideoId,omitempty" require:"true"`
	PageNumber  *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetStorageHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStorageHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetStorageHistoryRequest) SetClientToken(v string) *GetStorageHistoryRequest {
	s.ClientToken = &v
	return s
}

func (s *GetStorageHistoryRequest) SetInstanceId(v string) *GetStorageHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetStorageHistoryRequest) SetVideoId(v string) *GetStorageHistoryRequest {
	s.VideoId = &v
	return s
}

func (s *GetStorageHistoryRequest) SetPageNumber(v int) *GetStorageHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *GetStorageHistoryRequest) SetPageSize(v int) *GetStorageHistoryRequest {
	s.PageSize = &v
	return s
}

type GetStorageHistoryResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetStorageHistoryResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetStorageHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStorageHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetStorageHistoryResponse) SetRequestId(v string) *GetStorageHistoryResponse {
	s.RequestId = &v
	return s
}

func (s *GetStorageHistoryResponse) SetData(v *GetStorageHistoryResponseData) *GetStorageHistoryResponse {
	s.Data = v
	return s
}

type GetStorageHistoryResponseData struct {
	PageNumber *int                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Count      *int64                               `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	List       []*GetStorageHistoryResponseDataList `json:"List,omitempty" xml:"List,omitempty" require:"true" type:"Repeated"`
}

func (s GetStorageHistoryResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetStorageHistoryResponseData) GoString() string {
	return s.String()
}

func (s *GetStorageHistoryResponseData) SetPageNumber(v int) *GetStorageHistoryResponseData {
	s.PageNumber = &v
	return s
}

func (s *GetStorageHistoryResponseData) SetPageSize(v int) *GetStorageHistoryResponseData {
	s.PageSize = &v
	return s
}

func (s *GetStorageHistoryResponseData) SetCount(v int64) *GetStorageHistoryResponseData {
	s.Count = &v
	return s
}

func (s *GetStorageHistoryResponseData) SetList(v []*GetStorageHistoryResponseDataList) *GetStorageHistoryResponseData {
	s.List = v
	return s
}

type GetStorageHistoryResponseDataList struct {
	VideoId      *string `json:"VideoId,omitempty" xml:"VideoId,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	StorageType  *int    `json:"StorageType,omitempty" xml:"StorageType,omitempty" require:"true"`
	StorageInfo  *int    `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty" require:"true"`
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty" require:"true"`
	VideoUrl     *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
}

func (s GetStorageHistoryResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetStorageHistoryResponseDataList) GoString() string {
	return s.String()
}

func (s *GetStorageHistoryResponseDataList) SetVideoId(v string) *GetStorageHistoryResponseDataList {
	s.VideoId = &v
	return s
}

func (s *GetStorageHistoryResponseDataList) SetDescription(v string) *GetStorageHistoryResponseDataList {
	s.Description = &v
	return s
}

func (s *GetStorageHistoryResponseDataList) SetStorageType(v int) *GetStorageHistoryResponseDataList {
	s.StorageType = &v
	return s
}

func (s *GetStorageHistoryResponseDataList) SetStorageInfo(v int) *GetStorageHistoryResponseDataList {
	s.StorageInfo = &v
	return s
}

func (s *GetStorageHistoryResponseDataList) SetModifiedTime(v int64) *GetStorageHistoryResponseDataList {
	s.ModifiedTime = &v
	return s
}

func (s *GetStorageHistoryResponseDataList) SetVideoUrl(v string) *GetStorageHistoryResponseDataList {
	s.VideoUrl = &v
	return s
}

type ListBatchTaskRequest struct {
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	BatchTaskType *string `json:"BatchTaskType,omitempty" xml:"BatchTaskType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BucketName    *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	DataPath      *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
}

func (s ListBatchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *ListBatchTaskRequest) SetPageNumber(v int64) *ListBatchTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBatchTaskRequest) SetPageSize(v int64) *ListBatchTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListBatchTaskRequest) SetClientToken(v string) *ListBatchTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *ListBatchTaskRequest) SetBatchTaskType(v string) *ListBatchTaskRequest {
	s.BatchTaskType = &v
	return s
}

func (s *ListBatchTaskRequest) SetStatus(v string) *ListBatchTaskRequest {
	s.Status = &v
	return s
}

func (s *ListBatchTaskRequest) SetInstanceId(v string) *ListBatchTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBatchTaskRequest) SetBucketName(v string) *ListBatchTaskRequest {
	s.BucketName = &v
	return s
}

func (s *ListBatchTaskRequest) SetDataPath(v string) *ListBatchTaskRequest {
	s.DataPath = &v
	return s
}

type ListBatchTaskResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListBatchTaskResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListBatchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *ListBatchTaskResponse) SetRequestId(v string) *ListBatchTaskResponse {
	s.RequestId = &v
	return s
}

func (s *ListBatchTaskResponse) SetData(v *ListBatchTaskResponseData) *ListBatchTaskResponse {
	s.Data = v
	return s
}

type ListBatchTaskResponseData struct {
	Count      *int64                           `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	PageNumber *int64                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64                           `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	List       []*ListBatchTaskResponseDataList `json:"List,omitempty" xml:"List,omitempty" require:"true" type:"Repeated"`
}

func (s ListBatchTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTaskResponseData) GoString() string {
	return s.String()
}

func (s *ListBatchTaskResponseData) SetCount(v int64) *ListBatchTaskResponseData {
	s.Count = &v
	return s
}

func (s *ListBatchTaskResponseData) SetPageNumber(v int64) *ListBatchTaskResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListBatchTaskResponseData) SetPageSize(v int64) *ListBatchTaskResponseData {
	s.PageSize = &v
	return s
}

func (s *ListBatchTaskResponseData) SetList(v []*ListBatchTaskResponseDataList) *ListBatchTaskResponseData {
	s.List = v
	return s
}

type ListBatchTaskResponseDataList struct {
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	Status         *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	TaskType       *int    `json:"TaskType,omitempty" xml:"TaskType,omitempty" require:"true"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	BucketName     *string `json:"BucketName,omitempty" xml:"BucketName,omitempty" require:"true"`
	DataPath       *string `json:"DataPath,omitempty" xml:"DataPath,omitempty" require:"true"`
	MetaFile       *string `json:"MetaFile,omitempty" xml:"MetaFile,omitempty" require:"true"`
	ModifiedTime   *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty" require:"true"`
	ProcessMessage *string `json:"ProcessMessage,omitempty" xml:"ProcessMessage,omitempty" require:"true"`
	SubTaskDetail  *string `json:"SubTaskDetail,omitempty" xml:"SubTaskDetail,omitempty" require:"true"`
	Arn            *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
}

func (s ListBatchTaskResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTaskResponseDataList) GoString() string {
	return s.String()
}

func (s *ListBatchTaskResponseDataList) SetTaskId(v string) *ListBatchTaskResponseDataList {
	s.TaskId = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetStatus(v int) *ListBatchTaskResponseDataList {
	s.Status = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetTaskType(v int) *ListBatchTaskResponseDataList {
	s.TaskType = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetRegionId(v string) *ListBatchTaskResponseDataList {
	s.RegionId = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetBucketName(v string) *ListBatchTaskResponseDataList {
	s.BucketName = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetDataPath(v string) *ListBatchTaskResponseDataList {
	s.DataPath = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetMetaFile(v string) *ListBatchTaskResponseDataList {
	s.MetaFile = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetModifiedTime(v int64) *ListBatchTaskResponseDataList {
	s.ModifiedTime = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetProcessMessage(v string) *ListBatchTaskResponseDataList {
	s.ProcessMessage = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetSubTaskDetail(v string) *ListBatchTaskResponseDataList {
	s.SubTaskDetail = &v
	return s
}

func (s *ListBatchTaskResponseDataList) SetArn(v string) *ListBatchTaskResponseDataList {
	s.Arn = &v
	return s
}

type ListInstancesRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNumber   *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Status       *int    `json:"Status,omitempty" xml:"Status,omitempty"`
	InstanceType *int    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetClientToken(v string) *ListInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceName(v string) *ListInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetTags(v string) *ListInstancesRequest {
	s.Tags = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v int) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceType(v int) *ListInstancesRequest {
	s.InstanceType = &v
	return s
}

type ListInstancesResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListInstancesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetRequestId(v string) *ListInstancesResponse {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponse) SetData(v *ListInstancesResponseData) *ListInstancesResponse {
	s.Data = v
	return s
}

type ListInstancesResponseData struct {
	PageNumber *int                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                             `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Count      *int64                           `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	List       []*ListInstancesResponseDataList `json:"List,omitempty" xml:"List,omitempty" require:"true" type:"Repeated"`
}

func (s ListInstancesResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseData) SetPageNumber(v int) *ListInstancesResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseData) SetPageSize(v int) *ListInstancesResponseData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseData) SetCount(v int64) *ListInstancesResponseData {
	s.Count = &v
	return s
}

func (s *ListInstancesResponseData) SetList(v []*ListInstancesResponseDataList) *ListInstancesResponseData {
	s.List = v
	return s
}

type ListInstancesResponseDataList struct {
	InstanceId     *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceName   *string                              `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	InstanceStatus *int                                 `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty" require:"true"`
	CreateTime     *string                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ExpiredTime    *string                              `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" require:"true"`
	Tags           []*ListInstancesResponseDataListTags `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Repeated"`
}

func (s ListInstancesResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseDataList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseDataList) SetInstanceId(v string) *ListInstancesResponseDataList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseDataList) SetInstanceName(v string) *ListInstancesResponseDataList {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseDataList) SetInstanceStatus(v int) *ListInstancesResponseDataList {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstancesResponseDataList) SetCreateTime(v string) *ListInstancesResponseDataList {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseDataList) SetExpiredTime(v string) *ListInstancesResponseDataList {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseDataList) SetTags(v []*ListInstancesResponseDataListTags) *ListInstancesResponseDataList {
	s.Tags = v
	return s
}

type ListInstancesResponseDataListTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s ListInstancesResponseDataListTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseDataListTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseDataListTags) SetKey(v string) *ListInstancesResponseDataListTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseDataListTags) SetValue(v string) *ListInstancesResponseDataListTags {
	s.Value = &v
	return s
}

type ListStorageVideoTasksRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VideoName       *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	VideoId         *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	PageNumber      *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	StatusList      *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StorageInfoList *string `json:"StorageInfoList,omitempty" xml:"StorageInfoList,omitempty"`
	SortList        *string `json:"SortList,omitempty" xml:"SortList,omitempty"`
}

func (s ListStorageVideoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStorageVideoTasksRequest) GoString() string {
	return s.String()
}

func (s *ListStorageVideoTasksRequest) SetClientToken(v string) *ListStorageVideoTasksRequest {
	s.ClientToken = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetTaskId(v string) *ListStorageVideoTasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetVideoName(v string) *ListStorageVideoTasksRequest {
	s.VideoName = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetVideoId(v string) *ListStorageVideoTasksRequest {
	s.VideoId = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetPageNumber(v int) *ListStorageVideoTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetPageSize(v int) *ListStorageVideoTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetInstanceId(v string) *ListStorageVideoTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetStatusList(v string) *ListStorageVideoTasksRequest {
	s.StatusList = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetDescription(v string) *ListStorageVideoTasksRequest {
	s.Description = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetStorageInfoList(v string) *ListStorageVideoTasksRequest {
	s.StorageInfoList = &v
	return s
}

func (s *ListStorageVideoTasksRequest) SetSortList(v string) *ListStorageVideoTasksRequest {
	s.SortList = &v
	return s
}

type ListStorageVideoTasksResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListStorageVideoTasksResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListStorageVideoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStorageVideoTasksResponse) GoString() string {
	return s.String()
}

func (s *ListStorageVideoTasksResponse) SetRequestId(v string) *ListStorageVideoTasksResponse {
	s.RequestId = &v
	return s
}

func (s *ListStorageVideoTasksResponse) SetData(v *ListStorageVideoTasksResponseData) *ListStorageVideoTasksResponse {
	s.Data = v
	return s
}

type ListStorageVideoTasksResponseData struct {
	PageNumber *int                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Count      *int64                                       `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	TaskList   []*ListStorageVideoTasksResponseDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" require:"true" type:"Repeated"`
}

func (s ListStorageVideoTasksResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListStorageVideoTasksResponseData) GoString() string {
	return s.String()
}

func (s *ListStorageVideoTasksResponseData) SetPageNumber(v int) *ListStorageVideoTasksResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListStorageVideoTasksResponseData) SetPageSize(v int) *ListStorageVideoTasksResponseData {
	s.PageSize = &v
	return s
}

func (s *ListStorageVideoTasksResponseData) SetCount(v int64) *ListStorageVideoTasksResponseData {
	s.Count = &v
	return s
}

func (s *ListStorageVideoTasksResponseData) SetTaskList(v []*ListStorageVideoTasksResponseDataTaskList) *ListStorageVideoTasksResponseData {
	s.TaskList = v
	return s
}

type ListStorageVideoTasksResponseDataTaskList struct {
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	VideoId      *string `json:"VideoId,omitempty" xml:"VideoId,omitempty" require:"true"`
	VideoName    *string `json:"VideoName,omitempty" xml:"VideoName,omitempty" require:"true"`
	ProcessTime  *int64  `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty" require:"true"`
	StorageInfo  *int    `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty" require:"true"`
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ErrorDetail  *string `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" require:"true"`
	RemoteTaskId *string `json:"RemoteTaskId,omitempty" xml:"RemoteTaskId,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	VideoUrl     *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	Sort         *int    `json:"Sort,omitempty" xml:"Sort,omitempty" require:"true"`
}

func (s ListStorageVideoTasksResponseDataTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListStorageVideoTasksResponseDataTaskList) GoString() string {
	return s.String()
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetTaskId(v int64) *ListStorageVideoTasksResponseDataTaskList {
	s.TaskId = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetVideoId(v string) *ListStorageVideoTasksResponseDataTaskList {
	s.VideoId = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetVideoName(v string) *ListStorageVideoTasksResponseDataTaskList {
	s.VideoName = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetProcessTime(v int64) *ListStorageVideoTasksResponseDataTaskList {
	s.ProcessTime = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetStorageInfo(v int) *ListStorageVideoTasksResponseDataTaskList {
	s.StorageInfo = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetModifiedTime(v int64) *ListStorageVideoTasksResponseDataTaskList {
	s.ModifiedTime = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetStatus(v string) *ListStorageVideoTasksResponseDataTaskList {
	s.Status = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetErrorDetail(v string) *ListStorageVideoTasksResponseDataTaskList {
	s.ErrorDetail = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetRemoteTaskId(v string) *ListStorageVideoTasksResponseDataTaskList {
	s.RemoteTaskId = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetDescription(v string) *ListStorageVideoTasksResponseDataTaskList {
	s.Description = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetVideoUrl(v string) *ListStorageVideoTasksResponseDataTaskList {
	s.VideoUrl = &v
	return s
}

func (s *ListStorageVideoTasksResponseDataTaskList) SetSort(v int) *ListStorageVideoTasksResponseDataTaskList {
	s.Sort = &v
	return s
}

type ListSearchVideoTasksRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VideoName      *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	PageNumber     *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	StatusList     *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	SearchTypeList *string `json:"SearchTypeList,omitempty" xml:"SearchTypeList,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SortList       *string `json:"SortList,omitempty" xml:"SortList,omitempty"`
	VideoId        *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListSearchVideoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSearchVideoTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSearchVideoTasksRequest) SetClientToken(v string) *ListSearchVideoTasksRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetTaskId(v string) *ListSearchVideoTasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetVideoName(v string) *ListSearchVideoTasksRequest {
	s.VideoName = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetPageNumber(v int) *ListSearchVideoTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetPageSize(v int) *ListSearchVideoTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetInstanceId(v string) *ListSearchVideoTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetStatusList(v string) *ListSearchVideoTasksRequest {
	s.StatusList = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetSearchTypeList(v string) *ListSearchVideoTasksRequest {
	s.SearchTypeList = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetDescription(v string) *ListSearchVideoTasksRequest {
	s.Description = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetSortList(v string) *ListSearchVideoTasksRequest {
	s.SortList = &v
	return s
}

func (s *ListSearchVideoTasksRequest) SetVideoId(v string) *ListSearchVideoTasksRequest {
	s.VideoId = &v
	return s
}

type ListSearchVideoTasksResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ListSearchVideoTasksResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListSearchVideoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSearchVideoTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSearchVideoTasksResponse) SetRequestId(v string) *ListSearchVideoTasksResponse {
	s.RequestId = &v
	return s
}

func (s *ListSearchVideoTasksResponse) SetData(v *ListSearchVideoTasksResponseData) *ListSearchVideoTasksResponse {
	s.Data = v
	return s
}

type ListSearchVideoTasksResponseData struct {
	Count      *int64                                      `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	PageNumber *int                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TaskList   []*ListSearchVideoTasksResponseDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSearchVideoTasksResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListSearchVideoTasksResponseData) GoString() string {
	return s.String()
}

func (s *ListSearchVideoTasksResponseData) SetCount(v int64) *ListSearchVideoTasksResponseData {
	s.Count = &v
	return s
}

func (s *ListSearchVideoTasksResponseData) SetPageNumber(v int) *ListSearchVideoTasksResponseData {
	s.PageNumber = &v
	return s
}

func (s *ListSearchVideoTasksResponseData) SetPageSize(v int) *ListSearchVideoTasksResponseData {
	s.PageSize = &v
	return s
}

func (s *ListSearchVideoTasksResponseData) SetTaskList(v []*ListSearchVideoTasksResponseDataTaskList) *ListSearchVideoTasksResponseData {
	s.TaskList = v
	return s
}

type ListSearchVideoTasksResponseDataTaskList struct {
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	VideoId                 *string `json:"VideoId,omitempty" xml:"VideoId,omitempty" require:"true"`
	VideoName               *string `json:"VideoName,omitempty" xml:"VideoName,omitempty" require:"true"`
	ProcessTime             *int64  `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty" require:"true"`
	Status                  *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ModifiedTime            *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty" require:"true"`
	ProcessResultUrl        *string `json:"ProcessResultUrl,omitempty" xml:"ProcessResultUrl,omitempty" require:"true"`
	StorageType             *int    `json:"StorageType,omitempty" xml:"StorageType,omitempty" require:"true"`
	StorageInfo             *int    `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty" require:"true"`
	VideoUrl                *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	ErrorDetail             *string `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" require:"true"`
	ReplaceStorageThreshold *string `json:"ReplaceStorageThreshold,omitempty" xml:"ReplaceStorageThreshold,omitempty" require:"true"`
	QueryTags               *string `json:"QueryTags,omitempty" xml:"QueryTags,omitempty" require:"true"`
	RemoteTaskId            *string `json:"RemoteTaskId,omitempty" xml:"RemoteTaskId,omitempty" require:"true"`
	VideoTags               *string `json:"VideoTags,omitempty" xml:"VideoTags,omitempty" require:"true"`
	SearchType              *int    `json:"SearchType,omitempty" xml:"SearchType,omitempty" require:"true"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Sort                    *int    `json:"Sort,omitempty" xml:"Sort,omitempty" require:"true"`
}

func (s ListSearchVideoTasksResponseDataTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListSearchVideoTasksResponseDataTaskList) GoString() string {
	return s.String()
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetTaskId(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.TaskId = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetVideoId(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.VideoId = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetVideoName(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.VideoName = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetProcessTime(v int64) *ListSearchVideoTasksResponseDataTaskList {
	s.ProcessTime = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetStatus(v int) *ListSearchVideoTasksResponseDataTaskList {
	s.Status = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetModifiedTime(v int64) *ListSearchVideoTasksResponseDataTaskList {
	s.ModifiedTime = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetProcessResultUrl(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.ProcessResultUrl = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetStorageType(v int) *ListSearchVideoTasksResponseDataTaskList {
	s.StorageType = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetStorageInfo(v int) *ListSearchVideoTasksResponseDataTaskList {
	s.StorageInfo = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetVideoUrl(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.VideoUrl = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetErrorDetail(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.ErrorDetail = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetReplaceStorageThreshold(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.ReplaceStorageThreshold = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetQueryTags(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.QueryTags = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetRemoteTaskId(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.RemoteTaskId = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetVideoTags(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.VideoTags = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetSearchType(v int) *ListSearchVideoTasksResponseDataTaskList {
	s.SearchType = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetDescription(v string) *ListSearchVideoTasksResponseDataTaskList {
	s.Description = &v
	return s
}

func (s *ListSearchVideoTasksResponseDataTaskList) SetSort(v int) *ListSearchVideoTasksResponseDataTaskList {
	s.Sort = &v
	return s
}

type AddStorageVideoTaskRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VideoUrl    *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	VideoId     *string `json:"VideoId,omitempty" xml:"VideoId,omitempty" require:"true"`
	VideoTags   *string `json:"VideoTags,omitempty" xml:"VideoTags,omitempty"`
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StorageInfo *int    `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Sort        *int    `json:"Sort,omitempty" xml:"Sort,omitempty"`
	VideoFile   *string `json:"VideoFile,omitempty" xml:"VideoFile,omitempty"`
}

func (s AddStorageVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AddStorageVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *AddStorageVideoTaskRequest) SetInstanceId(v string) *AddStorageVideoTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *AddStorageVideoTaskRequest) SetVideoUrl(v string) *AddStorageVideoTaskRequest {
	s.VideoUrl = &v
	return s
}

func (s *AddStorageVideoTaskRequest) SetVideoId(v string) *AddStorageVideoTaskRequest {
	s.VideoId = &v
	return s
}

func (s *AddStorageVideoTaskRequest) SetVideoTags(v string) *AddStorageVideoTaskRequest {
	s.VideoTags = &v
	return s
}

func (s *AddStorageVideoTaskRequest) SetCallbackUrl(v string) *AddStorageVideoTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddStorageVideoTaskRequest) SetDescription(v string) *AddStorageVideoTaskRequest {
	s.Description = &v
	return s
}

func (s *AddStorageVideoTaskRequest) SetStorageInfo(v int) *AddStorageVideoTaskRequest {
	s.StorageInfo = &v
	return s
}

func (s *AddStorageVideoTaskRequest) SetClientToken(v string) *AddStorageVideoTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *AddStorageVideoTaskRequest) SetSort(v int) *AddStorageVideoTaskRequest {
	s.Sort = &v
	return s
}

func (s *AddStorageVideoTaskRequest) SetVideoFile(v string) *AddStorageVideoTaskRequest {
	s.VideoFile = &v
	return s
}

type AddStorageVideoTaskResponse struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *AddStorageVideoTaskResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AddStorageVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AddStorageVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *AddStorageVideoTaskResponse) SetRequestId(v string) *AddStorageVideoTaskResponse {
	s.RequestId = &v
	return s
}

func (s *AddStorageVideoTaskResponse) SetData(v *AddStorageVideoTaskResponseData) *AddStorageVideoTaskResponse {
	s.Data = v
	return s
}

type AddStorageVideoTaskResponseData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s AddStorageVideoTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddStorageVideoTaskResponseData) GoString() string {
	return s.String()
}

func (s *AddStorageVideoTaskResponseData) SetTaskId(v string) *AddStorageVideoTaskResponseData {
	s.TaskId = &v
	return s
}

type AddStorageVideoTaskAdvanceRequest struct {
	VideoFileObject io.Reader `json:"VideoFileObject,omitempty" xml:"VideoFileObject,omitempty" require:"true"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VideoUrl        *string   `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	VideoId         *string   `json:"VideoId,omitempty" xml:"VideoId,omitempty" require:"true"`
	VideoTags       *string   `json:"VideoTags,omitempty" xml:"VideoTags,omitempty"`
	CallbackUrl     *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	StorageInfo     *int      `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty"`
	ClientToken     *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Sort            *int      `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s AddStorageVideoTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddStorageVideoTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddStorageVideoTaskAdvanceRequest) SetVideoFileObject(v io.Reader) *AddStorageVideoTaskAdvanceRequest {
	s.VideoFileObject = v
	return s
}

func (s *AddStorageVideoTaskAdvanceRequest) SetInstanceId(v string) *AddStorageVideoTaskAdvanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AddStorageVideoTaskAdvanceRequest) SetVideoUrl(v string) *AddStorageVideoTaskAdvanceRequest {
	s.VideoUrl = &v
	return s
}

func (s *AddStorageVideoTaskAdvanceRequest) SetVideoId(v string) *AddStorageVideoTaskAdvanceRequest {
	s.VideoId = &v
	return s
}

func (s *AddStorageVideoTaskAdvanceRequest) SetVideoTags(v string) *AddStorageVideoTaskAdvanceRequest {
	s.VideoTags = &v
	return s
}

func (s *AddStorageVideoTaskAdvanceRequest) SetCallbackUrl(v string) *AddStorageVideoTaskAdvanceRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddStorageVideoTaskAdvanceRequest) SetDescription(v string) *AddStorageVideoTaskAdvanceRequest {
	s.Description = &v
	return s
}

func (s *AddStorageVideoTaskAdvanceRequest) SetStorageInfo(v int) *AddStorageVideoTaskAdvanceRequest {
	s.StorageInfo = &v
	return s
}

func (s *AddStorageVideoTaskAdvanceRequest) SetClientToken(v string) *AddStorageVideoTaskAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *AddStorageVideoTaskAdvanceRequest) SetSort(v int) *AddStorageVideoTaskAdvanceRequest {
	s.Sort = &v
	return s
}

type GetInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetClientToken(v string) *GetInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *GetInstanceRequest) SetInstanceId(v string) *GetInstanceRequest {
	s.InstanceId = &v
	return s
}

type GetInstanceResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetInstanceResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetRequestId(v string) *GetInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponse) SetData(v *GetInstanceResponseData) *GetInstanceResponse {
	s.Data = v
	return s
}

type GetInstanceResponseData struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceStatus       *int    `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty" require:"true"`
	VideoNumber          *int    `json:"VideoNumber,omitempty" xml:"VideoNumber,omitempty" require:"true"`
	ExpireTime           *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	ConcurrencyNumber    *int    `json:"ConcurrencyNumber,omitempty" xml:"ConcurrencyNumber,omitempty" require:"true"`
	MaxVideoCapacity     *string `json:"MaxVideoCapacity,omitempty" xml:"MaxVideoCapacity,omitempty" require:"true"`
	CurrentVideoCapacity *string `json:"CurrentVideoCapacity,omitempty" xml:"CurrentVideoCapacity,omitempty" require:"true"`
	VideoInfoUpdateTime  *int64  `json:"VideoInfoUpdateTime,omitempty" xml:"VideoInfoUpdateTime,omitempty" require:"true"`
	BundlingType         *string `json:"BundlingType,omitempty" xml:"BundlingType,omitempty" require:"true"`
}

func (s GetInstanceResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseData) SetInstanceId(v string) *GetInstanceResponseData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseData) SetInstanceName(v string) *GetInstanceResponseData {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseData) SetCreateTime(v int64) *GetInstanceResponseData {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseData) SetRegionId(v string) *GetInstanceResponseData {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseData) SetInstanceStatus(v int) *GetInstanceResponseData {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceResponseData) SetVideoNumber(v int) *GetInstanceResponseData {
	s.VideoNumber = &v
	return s
}

func (s *GetInstanceResponseData) SetExpireTime(v int64) *GetInstanceResponseData {
	s.ExpireTime = &v
	return s
}

func (s *GetInstanceResponseData) SetConcurrencyNumber(v int) *GetInstanceResponseData {
	s.ConcurrencyNumber = &v
	return s
}

func (s *GetInstanceResponseData) SetMaxVideoCapacity(v string) *GetInstanceResponseData {
	s.MaxVideoCapacity = &v
	return s
}

func (s *GetInstanceResponseData) SetCurrentVideoCapacity(v string) *GetInstanceResponseData {
	s.CurrentVideoCapacity = &v
	return s
}

func (s *GetInstanceResponseData) SetVideoInfoUpdateTime(v int64) *GetInstanceResponseData {
	s.VideoInfoUpdateTime = &v
	return s
}

func (s *GetInstanceResponseData) SetBundlingType(v string) *GetInstanceResponseData {
	s.BundlingType = &v
	return s
}

type AddDeletionVideoTaskRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	VideoId     *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddDeletionVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDeletionVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *AddDeletionVideoTaskRequest) SetClientToken(v string) *AddDeletionVideoTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDeletionVideoTaskRequest) SetVideoId(v string) *AddDeletionVideoTaskRequest {
	s.VideoId = &v
	return s
}

func (s *AddDeletionVideoTaskRequest) SetInstanceId(v string) *AddDeletionVideoTaskRequest {
	s.InstanceId = &v
	return s
}

type AddDeletionVideoTaskResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s AddDeletionVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDeletionVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *AddDeletionVideoTaskResponse) SetRequestId(v string) *AddDeletionVideoTaskResponse {
	s.RequestId = &v
	return s
}

func (s *AddDeletionVideoTaskResponse) SetData(v bool) *AddDeletionVideoTaskResponse {
	s.Data = &v
	return s
}

type GetTaskStatusRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s GetTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) SetClientToken(v string) *GetTaskStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTaskStatusRequest) SetTaskId(v string) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusRequest) SetInstanceId(v string) *GetTaskStatusRequest {
	s.InstanceId = &v
	return s
}

type GetTaskStatusResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *int                           `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	TaskInfo  *GetTaskStatusResponseTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponse) SetRequestId(v string) *GetTaskStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponse) SetData(v int) *GetTaskStatusResponse {
	s.Data = &v
	return s
}

func (s *GetTaskStatusResponse) SetTaskInfo(v *GetTaskStatusResponseTaskInfo) *GetTaskStatusResponse {
	s.TaskInfo = v
	return s
}

type GetTaskStatusResponseTaskInfo struct {
	AnalysisUseTime         *int64   `json:"AnalysisUseTime,omitempty" xml:"AnalysisUseTime,omitempty" require:"true"`
	Duration                *float32 `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	ProcessResultOss        *string  `json:"ProcessResultOss,omitempty" xml:"ProcessResultOss,omitempty" require:"true"`
	Resolution              *string  `json:"Resolution,omitempty" xml:"Resolution,omitempty" require:"true"`
	Status                  *int     `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	SubmitTime              *int64   `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty" require:"true"`
	FinishTime              *int64   `json:"FinishTime,omitempty" xml:"FinishTime,omitempty" require:"true"`
	TaskId                  *int64   `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	ErrorInfo               *string  `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" require:"true"`
	StorageInfo             *int     `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty" require:"true"`
	Description             *string  `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	VideoId                 *string  `json:"VideoId,omitempty" xml:"VideoId,omitempty" require:"true"`
	VideoTags               *string  `json:"VideoTags,omitempty" xml:"VideoTags,omitempty" require:"true"`
	VideoUrl                *string  `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	QueryTags               *string  `json:"QueryTags,omitempty" xml:"QueryTags,omitempty" require:"true"`
	ResourceType            *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ReplaceStorageThreshold *string  `json:"ReplaceStorageThreshold,omitempty" xml:"ReplaceStorageThreshold,omitempty" require:"true"`
	StorageType             *int     `json:"StorageType,omitempty" xml:"StorageType,omitempty" require:"true"`
}

func (s GetTaskStatusResponseTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseTaskInfo) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseTaskInfo) SetAnalysisUseTime(v int64) *GetTaskStatusResponseTaskInfo {
	s.AnalysisUseTime = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetDuration(v float32) *GetTaskStatusResponseTaskInfo {
	s.Duration = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetProcessResultOss(v string) *GetTaskStatusResponseTaskInfo {
	s.ProcessResultOss = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetResolution(v string) *GetTaskStatusResponseTaskInfo {
	s.Resolution = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetStatus(v int) *GetTaskStatusResponseTaskInfo {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetSubmitTime(v int64) *GetTaskStatusResponseTaskInfo {
	s.SubmitTime = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetFinishTime(v int64) *GetTaskStatusResponseTaskInfo {
	s.FinishTime = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetTaskId(v int64) *GetTaskStatusResponseTaskInfo {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetErrorInfo(v string) *GetTaskStatusResponseTaskInfo {
	s.ErrorInfo = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetStorageInfo(v int) *GetTaskStatusResponseTaskInfo {
	s.StorageInfo = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetDescription(v string) *GetTaskStatusResponseTaskInfo {
	s.Description = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetVideoId(v string) *GetTaskStatusResponseTaskInfo {
	s.VideoId = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetVideoTags(v string) *GetTaskStatusResponseTaskInfo {
	s.VideoTags = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetVideoUrl(v string) *GetTaskStatusResponseTaskInfo {
	s.VideoUrl = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetQueryTags(v string) *GetTaskStatusResponseTaskInfo {
	s.QueryTags = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetResourceType(v string) *GetTaskStatusResponseTaskInfo {
	s.ResourceType = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetReplaceStorageThreshold(v string) *GetTaskStatusResponseTaskInfo {
	s.ReplaceStorageThreshold = &v
	return s
}

func (s *GetTaskStatusResponseTaskInfo) SetStorageType(v int) *GetTaskStatusResponseTaskInfo {
	s.StorageType = &v
	return s
}

type AddSearchVideoTaskRequest struct {
	ClientToken             *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	VideoUrl                *string  `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	VideoId                 *string  `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	VideoTags               *string  `json:"VideoTags,omitempty" xml:"VideoTags,omitempty"`
	ReturnVideoNumber       *int     `json:"ReturnVideoNumber,omitempty" xml:"ReturnVideoNumber,omitempty"`
	QueryTags               *string  `json:"QueryTags,omitempty" xml:"QueryTags,omitempty"`
	StorageType             *int     `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	CallbackUrl             *string  `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ReplaceStorageThreshold *float32 `json:"ReplaceStorageThreshold,omitempty" xml:"ReplaceStorageThreshold,omitempty"`
	InstanceId              *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Description             *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	SearchType              *int     `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	VideoFile               *string  `json:"VideoFile,omitempty" xml:"VideoFile,omitempty"`
	Sort                    *int     `json:"Sort,omitempty" xml:"Sort,omitempty"`
	NeedFeatureFile         *int     `json:"NeedFeatureFile,omitempty" xml:"NeedFeatureFile,omitempty"`
}

func (s AddSearchVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSearchVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *AddSearchVideoTaskRequest) SetClientToken(v string) *AddSearchVideoTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetVideoUrl(v string) *AddSearchVideoTaskRequest {
	s.VideoUrl = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetVideoId(v string) *AddSearchVideoTaskRequest {
	s.VideoId = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetVideoTags(v string) *AddSearchVideoTaskRequest {
	s.VideoTags = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetReturnVideoNumber(v int) *AddSearchVideoTaskRequest {
	s.ReturnVideoNumber = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetQueryTags(v string) *AddSearchVideoTaskRequest {
	s.QueryTags = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetStorageType(v int) *AddSearchVideoTaskRequest {
	s.StorageType = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetCallbackUrl(v string) *AddSearchVideoTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetReplaceStorageThreshold(v float32) *AddSearchVideoTaskRequest {
	s.ReplaceStorageThreshold = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetInstanceId(v string) *AddSearchVideoTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetDescription(v string) *AddSearchVideoTaskRequest {
	s.Description = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetSearchType(v int) *AddSearchVideoTaskRequest {
	s.SearchType = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetVideoFile(v string) *AddSearchVideoTaskRequest {
	s.VideoFile = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetSort(v int) *AddSearchVideoTaskRequest {
	s.Sort = &v
	return s
}

func (s *AddSearchVideoTaskRequest) SetNeedFeatureFile(v int) *AddSearchVideoTaskRequest {
	s.NeedFeatureFile = &v
	return s
}

type AddSearchVideoTaskResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *AddSearchVideoTaskResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AddSearchVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSearchVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *AddSearchVideoTaskResponse) SetRequestId(v string) *AddSearchVideoTaskResponse {
	s.RequestId = &v
	return s
}

func (s *AddSearchVideoTaskResponse) SetData(v *AddSearchVideoTaskResponseData) *AddSearchVideoTaskResponse {
	s.Data = v
	return s
}

type AddSearchVideoTaskResponseData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s AddSearchVideoTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddSearchVideoTaskResponseData) GoString() string {
	return s.String()
}

func (s *AddSearchVideoTaskResponseData) SetTaskId(v string) *AddSearchVideoTaskResponseData {
	s.TaskId = &v
	return s
}

type AddSearchVideoTaskAdvanceRequest struct {
	VideoFileObject         io.Reader `json:"VideoFileObject,omitempty" xml:"VideoFileObject,omitempty" require:"true"`
	ClientToken             *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	VideoUrl                *string   `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	VideoId                 *string   `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	VideoTags               *string   `json:"VideoTags,omitempty" xml:"VideoTags,omitempty"`
	ReturnVideoNumber       *int      `json:"ReturnVideoNumber,omitempty" xml:"ReturnVideoNumber,omitempty"`
	QueryTags               *string   `json:"QueryTags,omitempty" xml:"QueryTags,omitempty"`
	StorageType             *int      `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	CallbackUrl             *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ReplaceStorageThreshold *float32  `json:"ReplaceStorageThreshold,omitempty" xml:"ReplaceStorageThreshold,omitempty"`
	InstanceId              *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Description             *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	SearchType              *int      `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	Sort                    *int      `json:"Sort,omitempty" xml:"Sort,omitempty"`
	NeedFeatureFile         *int      `json:"NeedFeatureFile,omitempty" xml:"NeedFeatureFile,omitempty"`
}

func (s AddSearchVideoTaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSearchVideoTaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddSearchVideoTaskAdvanceRequest) SetVideoFileObject(v io.Reader) *AddSearchVideoTaskAdvanceRequest {
	s.VideoFileObject = v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetClientToken(v string) *AddSearchVideoTaskAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetVideoUrl(v string) *AddSearchVideoTaskAdvanceRequest {
	s.VideoUrl = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetVideoId(v string) *AddSearchVideoTaskAdvanceRequest {
	s.VideoId = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetVideoTags(v string) *AddSearchVideoTaskAdvanceRequest {
	s.VideoTags = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetReturnVideoNumber(v int) *AddSearchVideoTaskAdvanceRequest {
	s.ReturnVideoNumber = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetQueryTags(v string) *AddSearchVideoTaskAdvanceRequest {
	s.QueryTags = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetStorageType(v int) *AddSearchVideoTaskAdvanceRequest {
	s.StorageType = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetCallbackUrl(v string) *AddSearchVideoTaskAdvanceRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetReplaceStorageThreshold(v float32) *AddSearchVideoTaskAdvanceRequest {
	s.ReplaceStorageThreshold = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetInstanceId(v string) *AddSearchVideoTaskAdvanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetDescription(v string) *AddSearchVideoTaskAdvanceRequest {
	s.Description = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetSearchType(v int) *AddSearchVideoTaskAdvanceRequest {
	s.SearchType = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetSort(v int) *AddSearchVideoTaskAdvanceRequest {
	s.Sort = &v
	return s
}

func (s *AddSearchVideoTaskAdvanceRequest) SetNeedFeatureFile(v int) *AddSearchVideoTaskAdvanceRequest {
	s.NeedFeatureFile = &v
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":  tea.String("multisearch.cn-beijing.aliyuncs.com"),
		"cn-hangzhou": tea.String("multisearch.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("videosearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) AddDeletionAudioTask(request *AddDeletionAudioTaskRequest, runtime *util.RuntimeOptions) (_result *AddDeletionAudioTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddDeletionAudioTaskResponse{}
	_body, _err := client.DoRequest(tea.String("AddDeletionAudioTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDeletionAudioTaskSimply(request *AddDeletionAudioTaskRequest) (_result *AddDeletionAudioTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDeletionAudioTaskResponse{}
	_body, _err := client.AddDeletionAudioTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAudioTaskStatus(request *GetAudioTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetAudioTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAudioTaskStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetAudioTaskStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAudioTaskStatusSimply(request *GetAudioTaskStatusRequest) (_result *GetAudioTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAudioTaskStatusResponse{}
	_body, _err := client.GetAudioTaskStatus(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelBatchTask(request *CancelBatchTaskRequest, runtime *util.RuntimeOptions) (_result *CancelBatchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelBatchTaskResponse{}
	_body, _err := client.DoRequest(tea.String("CancelBatchTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelBatchTaskSimply(request *CancelBatchTaskRequest) (_result *CancelBatchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelBatchTaskResponse{}
	_body, _err := client.CancelBatchTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAudioStorageHistory(request *GetAudioStorageHistoryRequest, runtime *util.RuntimeOptions) (_result *GetAudioStorageHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAudioStorageHistoryResponse{}
	_body, _err := client.DoRequest(tea.String("GetAudioStorageHistory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAudioStorageHistorySimply(request *GetAudioStorageHistoryRequest) (_result *GetAudioStorageHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAudioStorageHistoryResponse{}
	_body, _err := client.GetAudioStorageHistory(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPriority(request *ModifyPriorityRequest, runtime *util.RuntimeOptions) (_result *ModifyPriorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyPriorityResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyPriority"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPrioritySimply(request *ModifyPriorityRequest) (_result *ModifyPriorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPriorityResponse{}
	_body, _err := client.ModifyPriority(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAudioInstance(request *GetAudioInstanceRequest, runtime *util.RuntimeOptions) (_result *GetAudioInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAudioInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("GetAudioInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAudioInstanceSimply(request *GetAudioInstanceRequest) (_result *GetAudioInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAudioInstanceResponse{}
	_body, _err := client.GetAudioInstance(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBatchTask(request *GetBatchTaskRequest, runtime *util.RuntimeOptions) (_result *GetBatchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetBatchTaskResponse{}
	_body, _err := client.DoRequest(tea.String("GetBatchTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBatchTaskSimply(request *GetBatchTaskRequest) (_result *GetBatchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBatchTaskResponse{}
	_body, _err := client.GetBatchTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSearchAudioTask(request *AddSearchAudioTaskRequest, runtime *util.RuntimeOptions) (_result *AddSearchAudioTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddSearchAudioTaskResponse{}
	_body, _err := client.DoRequest(tea.String("AddSearchAudioTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSearchAudioTaskSimply(request *AddSearchAudioTaskRequest) (_result *AddSearchAudioTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSearchAudioTaskResponse{}
	_body, _err := client.AddSearchAudioTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSearchAudioTaskAdvance(request *AddSearchAudioTaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddSearchAudioTaskResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videosearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	addSearchAudioTaskReq := &AddSearchAudioTaskRequest{}
	rpcutil.Convert(request, addSearchAudioTaskReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.AudioFileObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	addSearchAudioTaskReq.AudioFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	addSearchAudioTaskResp, _err := client.AddSearchAudioTask(addSearchAudioTaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addSearchAudioTaskResp
	return _result, _err
}

func (client *Client) AddStorageAudioTask(request *AddStorageAudioTaskRequest, runtime *util.RuntimeOptions) (_result *AddStorageAudioTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddStorageAudioTaskResponse{}
	_body, _err := client.DoRequest(tea.String("AddStorageAudioTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddStorageAudioTaskSimply(request *AddStorageAudioTaskRequest) (_result *AddStorageAudioTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddStorageAudioTaskResponse{}
	_body, _err := client.AddStorageAudioTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddStorageAudioTaskAdvance(request *AddStorageAudioTaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddStorageAudioTaskResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videosearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	addStorageAudioTaskReq := &AddStorageAudioTaskRequest{}
	rpcutil.Convert(request, addStorageAudioTaskReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.AudioFileObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	addStorageAudioTaskReq.AudioFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	addStorageAudioTaskResp, _err := client.AddStorageAudioTask(addStorageAudioTaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addStorageAudioTaskResp
	return _result, _err
}

func (client *Client) ListStorageAudioTasks(request *ListStorageAudioTasksRequest, runtime *util.RuntimeOptions) (_result *ListStorageAudioTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListStorageAudioTasksResponse{}
	_body, _err := client.DoRequest(tea.String("ListStorageAudioTasks"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStorageAudioTasksSimply(request *ListStorageAudioTasksRequest) (_result *ListStorageAudioTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStorageAudioTasksResponse{}
	_body, _err := client.ListStorageAudioTasks(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSearchAudioTasks(request *ListSearchAudioTasksRequest, runtime *util.RuntimeOptions) (_result *ListSearchAudioTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListSearchAudioTasksResponse{}
	_body, _err := client.DoRequest(tea.String("ListSearchAudioTasks"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSearchAudioTasksSimply(request *ListSearchAudioTasksRequest) (_result *ListSearchAudioTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSearchAudioTasksResponse{}
	_body, _err := client.ListSearchAudioTasks(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBatchTask(request *CreateBatchTaskRequest, runtime *util.RuntimeOptions) (_result *CreateBatchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateBatchTaskResponse{}
	_body, _err := client.DoRequest(tea.String("CreateBatchTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBatchTaskSimply(request *CreateBatchTaskRequest) (_result *CreateBatchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBatchTaskResponse{}
	_body, _err := client.CreateBatchTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBatchTaskAdvance(request *CreateBatchTaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *CreateBatchTaskResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videosearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	createBatchTaskReq := &CreateBatchTaskRequest{}
	rpcutil.Convert(request, createBatchTaskReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.FileUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	createBatchTaskReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	createBatchTaskResp, _err := client.CreateBatchTask(createBatchTaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createBatchTaskResp
	return _result, _err
}

func (client *Client) GetStorageHistory(request *GetStorageHistoryRequest, runtime *util.RuntimeOptions) (_result *GetStorageHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetStorageHistoryResponse{}
	_body, _err := client.DoRequest(tea.String("GetStorageHistory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStorageHistorySimply(request *GetStorageHistoryRequest) (_result *GetStorageHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStorageHistoryResponse{}
	_body, _err := client.GetStorageHistory(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBatchTask(request *ListBatchTaskRequest, runtime *util.RuntimeOptions) (_result *ListBatchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListBatchTaskResponse{}
	_body, _err := client.DoRequest(tea.String("ListBatchTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBatchTaskSimply(request *ListBatchTaskRequest) (_result *ListBatchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBatchTaskResponse{}
	_body, _err := client.ListBatchTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstances(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("ListInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstancesSimply(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstances(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStorageVideoTasks(request *ListStorageVideoTasksRequest, runtime *util.RuntimeOptions) (_result *ListStorageVideoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListStorageVideoTasksResponse{}
	_body, _err := client.DoRequest(tea.String("ListStorageVideoTasks"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStorageVideoTasksSimply(request *ListStorageVideoTasksRequest) (_result *ListStorageVideoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStorageVideoTasksResponse{}
	_body, _err := client.ListStorageVideoTasks(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSearchVideoTasks(request *ListSearchVideoTasksRequest, runtime *util.RuntimeOptions) (_result *ListSearchVideoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListSearchVideoTasksResponse{}
	_body, _err := client.DoRequest(tea.String("ListSearchVideoTasks"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSearchVideoTasksSimply(request *ListSearchVideoTasksRequest) (_result *ListSearchVideoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSearchVideoTasksResponse{}
	_body, _err := client.ListSearchVideoTasks(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddStorageVideoTask(request *AddStorageVideoTaskRequest, runtime *util.RuntimeOptions) (_result *AddStorageVideoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddStorageVideoTaskResponse{}
	_body, _err := client.DoRequest(tea.String("AddStorageVideoTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddStorageVideoTaskSimply(request *AddStorageVideoTaskRequest) (_result *AddStorageVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddStorageVideoTaskResponse{}
	_body, _err := client.AddStorageVideoTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddStorageVideoTaskAdvance(request *AddStorageVideoTaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddStorageVideoTaskResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videosearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	addStorageVideoTaskReq := &AddStorageVideoTaskRequest{}
	rpcutil.Convert(request, addStorageVideoTaskReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoFileObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	addStorageVideoTaskReq.VideoFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	addStorageVideoTaskResp, _err := client.AddStorageVideoTask(addStorageVideoTaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addStorageVideoTaskResp
	return _result, _err
}

func (client *Client) GetInstance(request *GetInstanceRequest, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("GetInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceSimply(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstance(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDeletionVideoTask(request *AddDeletionVideoTaskRequest, runtime *util.RuntimeOptions) (_result *AddDeletionVideoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddDeletionVideoTaskResponse{}
	_body, _err := client.DoRequest(tea.String("AddDeletionVideoTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDeletionVideoTaskSimply(request *AddDeletionVideoTaskRequest) (_result *AddDeletionVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDeletionVideoTaskResponse{}
	_body, _err := client.AddDeletionVideoTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskStatus(request *GetTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetTaskStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskStatusSimply(request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatus(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSearchVideoTask(request *AddSearchVideoTaskRequest, runtime *util.RuntimeOptions) (_result *AddSearchVideoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddSearchVideoTaskResponse{}
	_body, _err := client.DoRequest(tea.String("AddSearchVideoTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-02-25"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSearchVideoTaskSimply(request *AddSearchVideoTaskRequest) (_result *AddSearchVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSearchVideoTaskResponse{}
	_body, _err := client.AddSearchVideoTask(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSearchVideoTaskAdvance(request *AddSearchVideoTaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddSearchVideoTaskResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("videosearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	rpcutil.Convert(runtime, ossRuntime)
	addSearchVideoTaskReq := &AddSearchVideoTaskRequest{}
	rpcutil.Convert(request, addSearchVideoTaskReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.VideoFileObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	addSearchVideoTaskReq.VideoFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	addSearchVideoTaskResp, _err := client.AddSearchVideoTask(addSearchVideoTaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addSearchVideoTaskResp
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
