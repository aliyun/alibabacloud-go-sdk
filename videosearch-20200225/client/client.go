// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateBatchTaskRequest struct {
	BatchTaskType *int    `json:"BatchTaskType,omitempty" xml:"BatchTaskType,omitempty"`
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssDataPath   *string `json:"OssDataPath,omitempty" xml:"OssDataPath,omitempty"`
	OssMetaFile   *string `json:"OssMetaFile,omitempty" xml:"OssMetaFile,omitempty"`
	FileUrl       *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	RoleArn       *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	TaskId         *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	Status         *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	TaskType       *int    `json:"TaskType,omitempty" xml:"TaskType,omitempty" require:"true"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	BucketName     *string `json:"BucketName,omitempty" xml:"BucketName,omitempty" require:"true"`
	DataPath       *string `json:"DataPath,omitempty" xml:"DataPath,omitempty" require:"true"`
	MetaFile       *string `json:"MetaFile,omitempty" xml:"MetaFile,omitempty" require:"true"`
	ModifiedTime   *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty" require:"true"`
	ProcessMessage *string `json:"ProcessMessage,omitempty" xml:"ProcessMessage,omitempty" require:"true"`
}

func (s ListBatchTaskResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s ListBatchTaskResponseDataList) GoString() string {
	return s.String()
}

func (s *ListBatchTaskResponseDataList) SetTaskId(v int64) *ListBatchTaskResponseDataList {
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

type ListInstancesRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNumber   *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Status       *int    `json:"Status,omitempty" xml:"Status,omitempty"`
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
	CreateTime     *int64                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
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

func (s *ListInstancesResponseDataList) SetCreateTime(v int64) *ListInstancesResponseDataList {
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

type ListSearchVideoTasksRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VideoName      *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	VideoId        *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	PageNumber     *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	StatusList     *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	SearchTypeList *string `json:"SearchTypeList,omitempty" xml:"SearchTypeList,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s *ListSearchVideoTasksRequest) SetVideoId(v string) *ListSearchVideoTasksRequest {
	s.VideoId = &v
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

type AddStorageVideoTaskRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VideoUrl    *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
	VideoId     *string `json:"VideoId,omitempty" xml:"VideoId,omitempty" require:"true"`
	VideoTags   *string `json:"VideoTags,omitempty" xml:"VideoTags,omitempty"`
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StorageInfo *int    `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

type AddSearchVideoTaskRequest struct {
	ClientToken             *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	VideoUrl                *string  `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" require:"true"`
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

func (client *Client) CreateBatchTaskWithOptions(request *CreateBatchTaskRequest, runtime *util.RuntimeOptions) (_result *CreateBatchTaskResponse, _err error) {
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

func (client *Client) CreateBatchTask(request *CreateBatchTaskRequest) (_result *CreateBatchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBatchTaskResponse{}
	_body, _err := client.CreateBatchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStorageHistoryWithOptions(request *GetStorageHistoryRequest, runtime *util.RuntimeOptions) (_result *GetStorageHistoryResponse, _err error) {
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

func (client *Client) GetStorageHistory(request *GetStorageHistoryRequest) (_result *GetStorageHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStorageHistoryResponse{}
	_body, _err := client.GetStorageHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBatchTaskWithOptions(request *ListBatchTaskRequest, runtime *util.RuntimeOptions) (_result *ListBatchTaskResponse, _err error) {
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

func (client *Client) ListBatchTask(request *ListBatchTaskRequest) (_result *ListBatchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBatchTaskResponse{}
	_body, _err := client.ListBatchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
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

func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStorageVideoTasksWithOptions(request *ListStorageVideoTasksRequest, runtime *util.RuntimeOptions) (_result *ListStorageVideoTasksResponse, _err error) {
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

func (client *Client) ListStorageVideoTasks(request *ListStorageVideoTasksRequest) (_result *ListStorageVideoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStorageVideoTasksResponse{}
	_body, _err := client.ListStorageVideoTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSearchVideoTasksWithOptions(request *ListSearchVideoTasksRequest, runtime *util.RuntimeOptions) (_result *ListSearchVideoTasksResponse, _err error) {
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

func (client *Client) ListSearchVideoTasks(request *ListSearchVideoTasksRequest) (_result *ListSearchVideoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSearchVideoTasksResponse{}
	_body, _err := client.ListSearchVideoTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddStorageVideoTaskWithOptions(request *AddStorageVideoTaskRequest, runtime *util.RuntimeOptions) (_result *AddStorageVideoTaskResponse, _err error) {
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

func (client *Client) AddStorageVideoTask(request *AddStorageVideoTaskRequest) (_result *AddStorageVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddStorageVideoTaskResponse{}
	_body, _err := client.AddStorageVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
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

func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDeletionVideoTaskWithOptions(request *AddDeletionVideoTaskRequest, runtime *util.RuntimeOptions) (_result *AddDeletionVideoTaskResponse, _err error) {
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

func (client *Client) AddDeletionVideoTask(request *AddDeletionVideoTaskRequest) (_result *AddDeletionVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDeletionVideoTaskResponse{}
	_body, _err := client.AddDeletionVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskStatusWithOptions(request *GetTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
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

func (client *Client) GetTaskStatus(request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSearchVideoTaskWithOptions(request *AddSearchVideoTaskRequest, runtime *util.RuntimeOptions) (_result *AddSearchVideoTaskResponse, _err error) {
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

func (client *Client) AddSearchVideoTask(request *AddSearchVideoTaskRequest) (_result *AddSearchVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSearchVideoTaskResponse{}
	_body, _err := client.AddSearchVideoTaskWithOptions(request, runtime)
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
