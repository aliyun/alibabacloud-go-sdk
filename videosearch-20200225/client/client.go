// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
