// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AiStoreUserTask struct {
	// ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 创建时间
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 地域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 地域描述
	RegionDesc *string `json:"RegionDesc,omitempty" xml:"RegionDesc,omitempty"`
	// 任务名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 产品名称
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitempty" xml:"ProductDesc,omitempty"`
	// API名称
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// API描述
	ApiDesc *string `json:"ApiDesc,omitempty" xml:"ApiDesc,omitempty"`
	// API版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// 参数信息
	ParamInfo *string `json:"ParamInfo,omitempty" xml:"ParamInfo,omitempty"`
	// bucket名称
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// bucketKey前缀
	BucketKeyPrefix *string `json:"BucketKeyPrefix,omitempty" xml:"BucketKeyPrefix,omitempty"`
	// 备注
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 接收消息配置
	ReceiveConfig *string `json:"ReceiveConfig,omitempty" xml:"ReceiveConfig,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 启用时间
	EnableTime *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	// 停用时间
	DisableTime *string `json:"DisableTime,omitempty" xml:"DisableTime,omitempty"`
}

func (s AiStoreUserTask) String() string {
	return tea.Prettify(s)
}

func (s AiStoreUserTask) GoString() string {
	return s.String()
}

func (s *AiStoreUserTask) SetId(v int64) *AiStoreUserTask {
	s.Id = &v
	return s
}

func (s *AiStoreUserTask) SetGmtCreate(v string) *AiStoreUserTask {
	s.GmtCreate = &v
	return s
}

func (s *AiStoreUserTask) SetGmtModified(v string) *AiStoreUserTask {
	s.GmtModified = &v
	return s
}

func (s *AiStoreUserTask) SetRegion(v string) *AiStoreUserTask {
	s.Region = &v
	return s
}

func (s *AiStoreUserTask) SetRegionDesc(v string) *AiStoreUserTask {
	s.RegionDesc = &v
	return s
}

func (s *AiStoreUserTask) SetName(v string) *AiStoreUserTask {
	s.Name = &v
	return s
}

func (s *AiStoreUserTask) SetProduct(v string) *AiStoreUserTask {
	s.Product = &v
	return s
}

func (s *AiStoreUserTask) SetProductDesc(v string) *AiStoreUserTask {
	s.ProductDesc = &v
	return s
}

func (s *AiStoreUserTask) SetApiName(v string) *AiStoreUserTask {
	s.ApiName = &v
	return s
}

func (s *AiStoreUserTask) SetApiDesc(v string) *AiStoreUserTask {
	s.ApiDesc = &v
	return s
}

func (s *AiStoreUserTask) SetVersion(v string) *AiStoreUserTask {
	s.Version = &v
	return s
}

func (s *AiStoreUserTask) SetParamInfo(v string) *AiStoreUserTask {
	s.ParamInfo = &v
	return s
}

func (s *AiStoreUserTask) SetBucketName(v string) *AiStoreUserTask {
	s.BucketName = &v
	return s
}

func (s *AiStoreUserTask) SetBucketKeyPrefix(v string) *AiStoreUserTask {
	s.BucketKeyPrefix = &v
	return s
}

func (s *AiStoreUserTask) SetRemark(v string) *AiStoreUserTask {
	s.Remark = &v
	return s
}

func (s *AiStoreUserTask) SetReceiveConfig(v string) *AiStoreUserTask {
	s.ReceiveConfig = &v
	return s
}

func (s *AiStoreUserTask) SetStatus(v string) *AiStoreUserTask {
	s.Status = &v
	return s
}

func (s *AiStoreUserTask) SetEnableTime(v string) *AiStoreUserTask {
	s.EnableTime = &v
	return s
}

func (s *AiStoreUserTask) SetDisableTime(v string) *AiStoreUserTask {
	s.DisableTime = &v
	return s
}

type AiStoreReceiveConfig struct {
	// 事件总线
	EventBridge *AiStoreReceiveConfigEventBridge `json:"EventBridge,omitempty" xml:"EventBridge,omitempty" type:"Struct"`
	// MNS消息
	Mns *AiStoreReceiveConfigMns `json:"Mns,omitempty" xml:"Mns,omitempty" type:"Struct"`
}

func (s AiStoreReceiveConfig) String() string {
	return tea.Prettify(s)
}

func (s AiStoreReceiveConfig) GoString() string {
	return s.String()
}

func (s *AiStoreReceiveConfig) SetEventBridge(v *AiStoreReceiveConfigEventBridge) *AiStoreReceiveConfig {
	s.EventBridge = v
	return s
}

func (s *AiStoreReceiveConfig) SetMns(v *AiStoreReceiveConfigMns) *AiStoreReceiveConfig {
	s.Mns = v
	return s
}

type AiStoreReceiveConfigEventBridge struct {
	// 事件总线
	EventBus *string `json:"EventBus,omitempty" xml:"EventBus,omitempty"`
	// 事件规则
	EventRule *string `json:"EventRule,omitempty" xml:"EventRule,omitempty"`
}

func (s AiStoreReceiveConfigEventBridge) String() string {
	return tea.Prettify(s)
}

func (s AiStoreReceiveConfigEventBridge) GoString() string {
	return s.String()
}

func (s *AiStoreReceiveConfigEventBridge) SetEventBus(v string) *AiStoreReceiveConfigEventBridge {
	s.EventBus = &v
	return s
}

func (s *AiStoreReceiveConfigEventBridge) SetEventRule(v string) *AiStoreReceiveConfigEventBridge {
	s.EventRule = &v
	return s
}

type AiStoreReceiveConfigMns struct {
	// 消息队列
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s AiStoreReceiveConfigMns) String() string {
	return tea.Prettify(s)
}

func (s AiStoreReceiveConfigMns) GoString() string {
	return s.String()
}

func (s *AiStoreReceiveConfigMns) SetQueue(v string) *AiStoreReceiveConfigMns {
	s.Queue = &v
	return s
}

type AiStoreApiNode struct {
	// 产品名称
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitempty" xml:"ProductDesc,omitempty"`
	// API列表
	Apis []*AiStoreUserTask `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Repeated"`
}

func (s AiStoreApiNode) String() string {
	return tea.Prettify(s)
}

func (s AiStoreApiNode) GoString() string {
	return s.String()
}

func (s *AiStoreApiNode) SetProduct(v string) *AiStoreApiNode {
	s.Product = &v
	return s
}

func (s *AiStoreApiNode) SetProductDesc(v string) *AiStoreApiNode {
	s.ProductDesc = &v
	return s
}

func (s *AiStoreApiNode) SetApis(v []*AiStoreUserTask) *AiStoreApiNode {
	s.Apis = v
	return s
}

type GetAiStoreUserTaskRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAiStoreUserTaskRequest) SetId(v int64) *GetAiStoreUserTaskRequest {
	s.Id = &v
	return s
}

type GetAiStoreUserTaskResponseBody struct {
	// Id of the request
	RequestId *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AiStoreUserTask `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiStoreUserTaskResponseBody) SetRequestId(v string) *GetAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiStoreUserTaskResponseBody) SetData(v *AiStoreUserTask) *GetAiStoreUserTaskResponseBody {
	s.Data = v
	return s
}

type GetAiStoreUserTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAiStoreUserTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreUserTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAiStoreUserTaskResponse) SetHeaders(v map[string]*string) *GetAiStoreUserTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAiStoreUserTaskResponse) SetBody(v *GetAiStoreUserTaskResponseBody) *GetAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type QueryAiStoreUserTaskPageRequest struct {
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ApiName    *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s QueryAiStoreUserTaskPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreUserTaskPageRequest) GoString() string {
	return s.String()
}

func (s *QueryAiStoreUserTaskPageRequest) SetProduct(v string) *QueryAiStoreUserTaskPageRequest {
	s.Product = &v
	return s
}

func (s *QueryAiStoreUserTaskPageRequest) SetApiName(v string) *QueryAiStoreUserTaskPageRequest {
	s.ApiName = &v
	return s
}

func (s *QueryAiStoreUserTaskPageRequest) SetStatus(v string) *QueryAiStoreUserTaskPageRequest {
	s.Status = &v
	return s
}

func (s *QueryAiStoreUserTaskPageRequest) SetPageNo(v int32) *QueryAiStoreUserTaskPageRequest {
	s.PageNo = &v
	return s
}

func (s *QueryAiStoreUserTaskPageRequest) SetPageSize(v int32) *QueryAiStoreUserTaskPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAiStoreUserTaskPageRequest) SetName(v string) *QueryAiStoreUserTaskPageRequest {
	s.Name = &v
	return s
}

func (s *QueryAiStoreUserTaskPageRequest) SetBucketName(v string) *QueryAiStoreUserTaskPageRequest {
	s.BucketName = &v
	return s
}

type QueryAiStoreUserTaskPageResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryAiStoreUserTaskPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryAiStoreUserTaskPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreUserTaskPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiStoreUserTaskPageResponseBody) SetRequestId(v string) *QueryAiStoreUserTaskPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAiStoreUserTaskPageResponseBody) SetData(v *QueryAiStoreUserTaskPageResponseBodyData) *QueryAiStoreUserTaskPageResponseBody {
	s.Data = v
	return s
}

type QueryAiStoreUserTaskPageResponseBodyData struct {
	TaskList   []*AiStoreUserTask `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	TotalCount *int32             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryAiStoreUserTaskPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreUserTaskPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAiStoreUserTaskPageResponseBodyData) SetTaskList(v []*AiStoreUserTask) *QueryAiStoreUserTaskPageResponseBodyData {
	s.TaskList = v
	return s
}

func (s *QueryAiStoreUserTaskPageResponseBodyData) SetTotalCount(v int32) *QueryAiStoreUserTaskPageResponseBodyData {
	s.TotalCount = &v
	return s
}

type QueryAiStoreUserTaskPageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAiStoreUserTaskPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAiStoreUserTaskPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreUserTaskPageResponse) GoString() string {
	return s.String()
}

func (s *QueryAiStoreUserTaskPageResponse) SetHeaders(v map[string]*string) *QueryAiStoreUserTaskPageResponse {
	s.Headers = v
	return s
}

func (s *QueryAiStoreUserTaskPageResponse) SetBody(v *QueryAiStoreUserTaskPageResponseBody) *QueryAiStoreUserTaskPageResponse {
	s.Body = v
	return s
}

type QueryAiStoreRegionsResponseBody struct {
	// Id of the request
	RequestId *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*AiStoreUserTask `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QueryAiStoreRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiStoreRegionsResponseBody) SetRequestId(v string) *QueryAiStoreRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAiStoreRegionsResponseBody) SetData(v []*AiStoreUserTask) *QueryAiStoreRegionsResponseBody {
	s.Data = v
	return s
}

type QueryAiStoreRegionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAiStoreRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAiStoreRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreRegionsResponse) GoString() string {
	return s.String()
}

func (s *QueryAiStoreRegionsResponse) SetHeaders(v map[string]*string) *QueryAiStoreRegionsResponse {
	s.Headers = v
	return s
}

func (s *QueryAiStoreRegionsResponse) SetBody(v *QueryAiStoreRegionsResponseBody) *QueryAiStoreRegionsResponse {
	s.Body = v
	return s
}

type ListAiStoreBucketsRequest struct {
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
}

func (s ListAiStoreBucketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAiStoreBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListAiStoreBucketsRequest) SetProduct(v string) *ListAiStoreBucketsRequest {
	s.Product = &v
	return s
}

func (s *ListAiStoreBucketsRequest) SetApiName(v string) *ListAiStoreBucketsRequest {
	s.ApiName = &v
	return s
}

type ListAiStoreBucketsResponseBody struct {
	// Id of the request
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListAiStoreBucketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAiStoreBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiStoreBucketsResponseBody) SetRequestId(v string) *ListAiStoreBucketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAiStoreBucketsResponseBody) SetData(v []*string) *ListAiStoreBucketsResponseBody {
	s.Data = v
	return s
}

type ListAiStoreBucketsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAiStoreBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAiStoreBucketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAiStoreBucketsResponse) GoString() string {
	return s.String()
}

func (s *ListAiStoreBucketsResponse) SetHeaders(v map[string]*string) *ListAiStoreBucketsResponse {
	s.Headers = v
	return s
}

func (s *ListAiStoreBucketsResponse) SetBody(v *ListAiStoreBucketsResponseBody) *ListAiStoreBucketsResponse {
	s.Body = v
	return s
}

type GetAiStoreUserTaskByNameRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAiStoreUserTaskByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreUserTaskByNameRequest) GoString() string {
	return s.String()
}

func (s *GetAiStoreUserTaskByNameRequest) SetName(v string) *GetAiStoreUserTaskByNameRequest {
	s.Name = &v
	return s
}

type GetAiStoreUserTaskByNameResponseBody struct {
	// Id of the request
	RequestId *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AiStoreUserTask `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetAiStoreUserTaskByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreUserTaskByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiStoreUserTaskByNameResponseBody) SetRequestId(v string) *GetAiStoreUserTaskByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiStoreUserTaskByNameResponseBody) SetData(v *AiStoreUserTask) *GetAiStoreUserTaskByNameResponseBody {
	s.Data = v
	return s
}

type GetAiStoreUserTaskByNameResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAiStoreUserTaskByNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAiStoreUserTaskByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreUserTaskByNameResponse) GoString() string {
	return s.String()
}

func (s *GetAiStoreUserTaskByNameResponse) SetHeaders(v map[string]*string) *GetAiStoreUserTaskByNameResponse {
	s.Headers = v
	return s
}

func (s *GetAiStoreUserTaskByNameResponse) SetBody(v *GetAiStoreUserTaskByNameResponseBody) *GetAiStoreUserTaskByNameResponse {
	s.Body = v
	return s
}

type UpdateAiStoreUserTaskRequest struct {
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ApiName         *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BucketKeyPrefix *string `json:"BucketKeyPrefix,omitempty" xml:"BucketKeyPrefix,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamInfo       *string `json:"ParamInfo,omitempty" xml:"ParamInfo,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ReceiveConfig   *string `json:"ReceiveConfig,omitempty" xml:"ReceiveConfig,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiStoreUserTaskRequest) SetProduct(v string) *UpdateAiStoreUserTaskRequest {
	s.Product = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetApiName(v string) *UpdateAiStoreUserTaskRequest {
	s.ApiName = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetBucketName(v string) *UpdateAiStoreUserTaskRequest {
	s.BucketName = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetBucketKeyPrefix(v string) *UpdateAiStoreUserTaskRequest {
	s.BucketKeyPrefix = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetName(v string) *UpdateAiStoreUserTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetParamInfo(v string) *UpdateAiStoreUserTaskRequest {
	s.ParamInfo = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetRemark(v string) *UpdateAiStoreUserTaskRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetReceiveConfig(v string) *UpdateAiStoreUserTaskRequest {
	s.ReceiveConfig = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetStatus(v string) *UpdateAiStoreUserTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetId(v int64) *UpdateAiStoreUserTaskRequest {
	s.Id = &v
	return s
}

type UpdateAiStoreUserTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s UpdateAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAiStoreUserTaskResponseBody) SetRequestId(v string) *UpdateAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAiStoreUserTaskResponseBody) SetData(v bool) *UpdateAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

type UpdateAiStoreUserTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAiStoreUserTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAiStoreUserTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateAiStoreUserTaskResponse) SetHeaders(v map[string]*string) *UpdateAiStoreUserTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateAiStoreUserTaskResponse) SetBody(v *UpdateAiStoreUserTaskResponseBody) *UpdateAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type DisableAiStoreUserTaskRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DisableAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *DisableAiStoreUserTaskRequest) SetId(v int64) *DisableAiStoreUserTaskRequest {
	s.Id = &v
	return s
}

type DisableAiStoreUserTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s DisableAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAiStoreUserTaskResponseBody) SetRequestId(v string) *DisableAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAiStoreUserTaskResponseBody) SetData(v bool) *DisableAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

type DisableAiStoreUserTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAiStoreUserTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAiStoreUserTaskResponse) GoString() string {
	return s.String()
}

func (s *DisableAiStoreUserTaskResponse) SetHeaders(v map[string]*string) *DisableAiStoreUserTaskResponse {
	s.Headers = v
	return s
}

func (s *DisableAiStoreUserTaskResponse) SetBody(v *DisableAiStoreUserTaskResponseBody) *DisableAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type QueryAiStoreApiTreeResponseBody struct {
	// Id of the request
	RequestId *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*AiStoreApiNode `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QueryAiStoreApiTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreApiTreeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiStoreApiTreeResponseBody) SetRequestId(v string) *QueryAiStoreApiTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAiStoreApiTreeResponseBody) SetData(v []*AiStoreApiNode) *QueryAiStoreApiTreeResponseBody {
	s.Data = v
	return s
}

type QueryAiStoreApiTreeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAiStoreApiTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAiStoreApiTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreApiTreeResponse) GoString() string {
	return s.String()
}

func (s *QueryAiStoreApiTreeResponse) SetHeaders(v map[string]*string) *QueryAiStoreApiTreeResponse {
	s.Headers = v
	return s
}

func (s *QueryAiStoreApiTreeResponse) SetBody(v *QueryAiStoreApiTreeResponseBody) *QueryAiStoreApiTreeResponse {
	s.Body = v
	return s
}

type DeleteAiStoreUserTaskRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteAiStoreUserTaskRequest) SetId(v int64) *DeleteAiStoreUserTaskRequest {
	s.Id = &v
	return s
}

type DeleteAiStoreUserTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s DeleteAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAiStoreUserTaskResponseBody) SetRequestId(v string) *DeleteAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAiStoreUserTaskResponseBody) SetData(v bool) *DeleteAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

type DeleteAiStoreUserTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAiStoreUserTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAiStoreUserTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteAiStoreUserTaskResponse) SetHeaders(v map[string]*string) *DeleteAiStoreUserTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteAiStoreUserTaskResponse) SetBody(v *DeleteAiStoreUserTaskResponseBody) *DeleteAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type CreateAiStoreUserTaskRequest struct {
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ApiName         *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BucketKeyPrefix *string `json:"BucketKeyPrefix,omitempty" xml:"BucketKeyPrefix,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamInfo       *string `json:"ParamInfo,omitempty" xml:"ParamInfo,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ReceiveConfig   *string `json:"ReceiveConfig,omitempty" xml:"ReceiveConfig,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAiStoreUserTaskRequest) SetProduct(v string) *CreateAiStoreUserTaskRequest {
	s.Product = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetApiName(v string) *CreateAiStoreUserTaskRequest {
	s.ApiName = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetBucketName(v string) *CreateAiStoreUserTaskRequest {
	s.BucketName = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetBucketKeyPrefix(v string) *CreateAiStoreUserTaskRequest {
	s.BucketKeyPrefix = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetName(v string) *CreateAiStoreUserTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetParamInfo(v string) *CreateAiStoreUserTaskRequest {
	s.ParamInfo = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetRemark(v string) *CreateAiStoreUserTaskRequest {
	s.Remark = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetReceiveConfig(v string) *CreateAiStoreUserTaskRequest {
	s.ReceiveConfig = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetStatus(v string) *CreateAiStoreUserTaskRequest {
	s.Status = &v
	return s
}

type CreateAiStoreUserTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s CreateAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiStoreUserTaskResponseBody) SetRequestId(v string) *CreateAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAiStoreUserTaskResponseBody) SetData(v int64) *CreateAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

type CreateAiStoreUserTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAiStoreUserTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreUserTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAiStoreUserTaskResponse) SetHeaders(v map[string]*string) *CreateAiStoreUserTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAiStoreUserTaskResponse) SetBody(v *CreateAiStoreUserTaskResponseBody) *CreateAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type CreateAiStoreReceiveConfigRequest struct {
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
}

func (s CreateAiStoreReceiveConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreReceiveConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAiStoreReceiveConfigRequest) SetProduct(v string) *CreateAiStoreReceiveConfigRequest {
	s.Product = &v
	return s
}

func (s *CreateAiStoreReceiveConfigRequest) SetApiName(v string) *CreateAiStoreReceiveConfigRequest {
	s.ApiName = &v
	return s
}

type CreateAiStoreReceiveConfigResponseBody struct {
	// Id of the request
	RequestId *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AiStoreReceiveConfig `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s CreateAiStoreReceiveConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreReceiveConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiStoreReceiveConfigResponseBody) SetRequestId(v string) *CreateAiStoreReceiveConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAiStoreReceiveConfigResponseBody) SetData(v *AiStoreReceiveConfig) *CreateAiStoreReceiveConfigResponseBody {
	s.Data = v
	return s
}

type CreateAiStoreReceiveConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAiStoreReceiveConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAiStoreReceiveConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreReceiveConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAiStoreReceiveConfigResponse) SetHeaders(v map[string]*string) *CreateAiStoreReceiveConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAiStoreReceiveConfigResponse) SetBody(v *CreateAiStoreReceiveConfigResponseBody) *CreateAiStoreReceiveConfigResponse {
	s.Body = v
	return s
}

type GetAiStoreReceiveConfigRequest struct {
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
}

func (s GetAiStoreReceiveConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreReceiveConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAiStoreReceiveConfigRequest) SetProduct(v string) *GetAiStoreReceiveConfigRequest {
	s.Product = &v
	return s
}

func (s *GetAiStoreReceiveConfigRequest) SetApiName(v string) *GetAiStoreReceiveConfigRequest {
	s.ApiName = &v
	return s
}

type GetAiStoreReceiveConfigResponseBody struct {
	// Id of the request
	RequestId *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AiStoreReceiveConfig `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetAiStoreReceiveConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreReceiveConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiStoreReceiveConfigResponseBody) SetRequestId(v string) *GetAiStoreReceiveConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiStoreReceiveConfigResponseBody) SetData(v *AiStoreReceiveConfig) *GetAiStoreReceiveConfigResponseBody {
	s.Data = v
	return s
}

type GetAiStoreReceiveConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAiStoreReceiveConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAiStoreReceiveConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreReceiveConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAiStoreReceiveConfigResponse) SetHeaders(v map[string]*string) *GetAiStoreReceiveConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAiStoreReceiveConfigResponse) SetBody(v *GetAiStoreReceiveConfigResponseBody) *GetAiStoreReceiveConfigResponse {
	s.Body = v
	return s
}

type EnableAiStoreUserTaskRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s EnableAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *EnableAiStoreUserTaskRequest) SetId(v int64) *EnableAiStoreUserTaskRequest {
	s.Id = &v
	return s
}

type EnableAiStoreUserTaskResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s EnableAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAiStoreUserTaskResponseBody) SetRequestId(v string) *EnableAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableAiStoreUserTaskResponseBody) SetData(v string) *EnableAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

type EnableAiStoreUserTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableAiStoreUserTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAiStoreUserTaskResponse) GoString() string {
	return s.String()
}

func (s *EnableAiStoreUserTaskResponse) SetHeaders(v map[string]*string) *EnableAiStoreUserTaskResponse {
	s.Headers = v
	return s
}

func (s *EnableAiStoreUserTaskResponse) SetBody(v *EnableAiStoreUserTaskResponseBody) *EnableAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type CreateAiStoreBucketRequest struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s CreateAiStoreBucketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreBucketRequest) GoString() string {
	return s.String()
}

func (s *CreateAiStoreBucketRequest) SetBucketName(v string) *CreateAiStoreBucketRequest {
	s.BucketName = &v
	return s
}

type CreateAiStoreBucketResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s CreateAiStoreBucketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreBucketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiStoreBucketResponseBody) SetRequestId(v string) *CreateAiStoreBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAiStoreBucketResponseBody) SetData(v string) *CreateAiStoreBucketResponseBody {
	s.Data = &v
	return s
}

type CreateAiStoreBucketResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAiStoreBucketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAiStoreBucketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreBucketResponse) GoString() string {
	return s.String()
}

func (s *CreateAiStoreBucketResponse) SetHeaders(v map[string]*string) *CreateAiStoreBucketResponse {
	s.Headers = v
	return s
}

func (s *CreateAiStoreBucketResponse) SetBody(v *CreateAiStoreBucketResponseBody) *CreateAiStoreBucketResponse {
	s.Body = v
	return s
}

type CheckServiceLinkedRoleForDeletingRequest struct {
	RoleArn        *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SPIRegionId    *string `json:"SPIRegionId,omitempty" xml:"SPIRegionId,omitempty"`
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CheckServiceLinkedRoleForDeletingRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetRoleArn(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.RoleArn = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetServiceName(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.ServiceName = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetSPIRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.SPIRegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetDeletionTaskId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.DeletionTaskId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetAccountId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.AccountId = &v
	return s
}

type CheckServiceLinkedRoleForDeletingResponseBody struct {
	// Id of the request
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Deletable  *bool                                                      `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	RoleUsages []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages `json:"RoleUsages,omitempty" xml:"RoleUsages,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetDeletable(v bool) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.Deletable = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetRoleUsages(v []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.RoleUsages = v
	return s
}

type CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages struct {
	Region    *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	Resources [][]byte `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) SetRegion(v string) *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages {
	s.Region = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) SetResources(v [][]byte) *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages {
	s.Resources = v
	return s
}

type CheckServiceLinkedRoleForDeletingResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckServiceLinkedRoleForDeletingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckServiceLinkedRoleForDeletingResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleForDeletingResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponse) SetBody(v *CheckServiceLinkedRoleForDeletingResponseBody) *CheckServiceLinkedRoleForDeletingResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("viapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetAiStoreUserTaskWithOptions(request *GetAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *GetAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAiStoreUserTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAiStoreUserTask"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAiStoreUserTask(request *GetAiStoreUserTaskRequest) (_result *GetAiStoreUserTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAiStoreUserTaskResponse{}
	_body, _err := client.GetAiStoreUserTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAiStoreUserTaskPageWithOptions(request *QueryAiStoreUserTaskPageRequest, runtime *util.RuntimeOptions) (_result *QueryAiStoreUserTaskPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAiStoreUserTaskPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAiStoreUserTaskPage"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAiStoreUserTaskPage(request *QueryAiStoreUserTaskPageRequest) (_result *QueryAiStoreUserTaskPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAiStoreUserTaskPageResponse{}
	_body, _err := client.QueryAiStoreUserTaskPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAiStoreRegionsWithOptions(runtime *util.RuntimeOptions) (_result *QueryAiStoreRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &QueryAiStoreRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAiStoreRegions"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAiStoreRegions() (_result *QueryAiStoreRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAiStoreRegionsResponse{}
	_body, _err := client.QueryAiStoreRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAiStoreBucketsWithOptions(request *ListAiStoreBucketsRequest, runtime *util.RuntimeOptions) (_result *ListAiStoreBucketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAiStoreBucketsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAiStoreBuckets"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAiStoreBuckets(request *ListAiStoreBucketsRequest) (_result *ListAiStoreBucketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAiStoreBucketsResponse{}
	_body, _err := client.ListAiStoreBucketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAiStoreUserTaskByNameWithOptions(request *GetAiStoreUserTaskByNameRequest, runtime *util.RuntimeOptions) (_result *GetAiStoreUserTaskByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAiStoreUserTaskByNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAiStoreUserTaskByName"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAiStoreUserTaskByName(request *GetAiStoreUserTaskByNameRequest) (_result *GetAiStoreUserTaskByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAiStoreUserTaskByNameResponse{}
	_body, _err := client.GetAiStoreUserTaskByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAiStoreUserTaskWithOptions(request *UpdateAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAiStoreUserTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAiStoreUserTask"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAiStoreUserTask(request *UpdateAiStoreUserTaskRequest) (_result *UpdateAiStoreUserTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAiStoreUserTaskResponse{}
	_body, _err := client.UpdateAiStoreUserTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAiStoreUserTaskWithOptions(request *DisableAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *DisableAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableAiStoreUserTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableAiStoreUserTask"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAiStoreUserTask(request *DisableAiStoreUserTaskRequest) (_result *DisableAiStoreUserTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAiStoreUserTaskResponse{}
	_body, _err := client.DisableAiStoreUserTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAiStoreApiTreeWithOptions(runtime *util.RuntimeOptions) (_result *QueryAiStoreApiTreeResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &QueryAiStoreApiTreeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAiStoreApiTree"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAiStoreApiTree() (_result *QueryAiStoreApiTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAiStoreApiTreeResponse{}
	_body, _err := client.QueryAiStoreApiTreeWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAiStoreUserTaskWithOptions(request *DeleteAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAiStoreUserTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAiStoreUserTask"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAiStoreUserTask(request *DeleteAiStoreUserTaskRequest) (_result *DeleteAiStoreUserTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAiStoreUserTaskResponse{}
	_body, _err := client.DeleteAiStoreUserTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAiStoreUserTaskWithOptions(request *CreateAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *CreateAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAiStoreUserTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAiStoreUserTask"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAiStoreUserTask(request *CreateAiStoreUserTaskRequest) (_result *CreateAiStoreUserTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAiStoreUserTaskResponse{}
	_body, _err := client.CreateAiStoreUserTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAiStoreReceiveConfigWithOptions(request *CreateAiStoreReceiveConfigRequest, runtime *util.RuntimeOptions) (_result *CreateAiStoreReceiveConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAiStoreReceiveConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAiStoreReceiveConfig"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAiStoreReceiveConfig(request *CreateAiStoreReceiveConfigRequest) (_result *CreateAiStoreReceiveConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAiStoreReceiveConfigResponse{}
	_body, _err := client.CreateAiStoreReceiveConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAiStoreReceiveConfigWithOptions(request *GetAiStoreReceiveConfigRequest, runtime *util.RuntimeOptions) (_result *GetAiStoreReceiveConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAiStoreReceiveConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAiStoreReceiveConfig"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAiStoreReceiveConfig(request *GetAiStoreReceiveConfigRequest) (_result *GetAiStoreReceiveConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAiStoreReceiveConfigResponse{}
	_body, _err := client.GetAiStoreReceiveConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableAiStoreUserTaskWithOptions(request *EnableAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *EnableAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableAiStoreUserTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableAiStoreUserTask"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableAiStoreUserTask(request *EnableAiStoreUserTaskRequest) (_result *EnableAiStoreUserTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAiStoreUserTaskResponse{}
	_body, _err := client.EnableAiStoreUserTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAiStoreBucketWithOptions(request *CreateAiStoreBucketRequest, runtime *util.RuntimeOptions) (_result *CreateAiStoreBucketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAiStoreBucketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAiStoreBucket"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAiStoreBucket(request *CreateAiStoreBucketRequest) (_result *CreateAiStoreBucketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAiStoreBucketResponse{}
	_body, _err := client.CreateAiStoreBucketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckServiceLinkedRoleForDeletingWithOptions(request *CheckServiceLinkedRoleForDeletingRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleForDeletingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckServiceLinkedRoleForDeletingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckServiceLinkedRoleForDeleting"), tea.String("2021-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckServiceLinkedRoleForDeleting(request *CheckServiceLinkedRoleForDeletingRequest) (_result *CheckServiceLinkedRoleForDeletingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleForDeletingResponse{}
	_body, _err := client.CheckServiceLinkedRoleForDeletingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
