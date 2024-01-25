// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AiStoreApiNode struct {
	Apis        []*AiStoreUserTask `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Repeated"`
	Product     *string            `json:"Product,omitempty" xml:"Product,omitempty"`
	ProductDesc *string            `json:"ProductDesc,omitempty" xml:"ProductDesc,omitempty"`
}

func (s AiStoreApiNode) String() string {
	return tea.Prettify(s)
}

func (s AiStoreApiNode) GoString() string {
	return s.String()
}

func (s *AiStoreApiNode) SetApis(v []*AiStoreUserTask) *AiStoreApiNode {
	s.Apis = v
	return s
}

func (s *AiStoreApiNode) SetProduct(v string) *AiStoreApiNode {
	s.Product = &v
	return s
}

func (s *AiStoreApiNode) SetProductDesc(v string) *AiStoreApiNode {
	s.ProductDesc = &v
	return s
}

type AiStoreReceiveConfig struct {
	Custom      *string                          `json:"Custom,omitempty" xml:"Custom,omitempty"`
	DingTalk    *AiStoreReceiveConfigDingTalk    `json:"DingTalk,omitempty" xml:"DingTalk,omitempty" type:"Struct"`
	EventBridge *AiStoreReceiveConfigEventBridge `json:"EventBridge,omitempty" xml:"EventBridge,omitempty" type:"Struct"`
	Http        *AiStoreReceiveConfigHttp        `json:"Http,omitempty" xml:"Http,omitempty" type:"Struct"`
	Https       *AiStoreReceiveConfigHttps       `json:"Https,omitempty" xml:"Https,omitempty" type:"Struct"`
	Mns         *AiStoreReceiveConfigMns         `json:"Mns,omitempty" xml:"Mns,omitempty" type:"Struct"`
	RocketMQ    *AiStoreReceiveConfigRocketMQ    `json:"RocketMQ,omitempty" xml:"RocketMQ,omitempty" type:"Struct"`
}

func (s AiStoreReceiveConfig) String() string {
	return tea.Prettify(s)
}

func (s AiStoreReceiveConfig) GoString() string {
	return s.String()
}

func (s *AiStoreReceiveConfig) SetCustom(v string) *AiStoreReceiveConfig {
	s.Custom = &v
	return s
}

func (s *AiStoreReceiveConfig) SetDingTalk(v *AiStoreReceiveConfigDingTalk) *AiStoreReceiveConfig {
	s.DingTalk = v
	return s
}

func (s *AiStoreReceiveConfig) SetEventBridge(v *AiStoreReceiveConfigEventBridge) *AiStoreReceiveConfig {
	s.EventBridge = v
	return s
}

func (s *AiStoreReceiveConfig) SetHttp(v *AiStoreReceiveConfigHttp) *AiStoreReceiveConfig {
	s.Http = v
	return s
}

func (s *AiStoreReceiveConfig) SetHttps(v *AiStoreReceiveConfigHttps) *AiStoreReceiveConfig {
	s.Https = v
	return s
}

func (s *AiStoreReceiveConfig) SetMns(v *AiStoreReceiveConfigMns) *AiStoreReceiveConfig {
	s.Mns = v
	return s
}

func (s *AiStoreReceiveConfig) SetRocketMQ(v *AiStoreReceiveConfigRocketMQ) *AiStoreReceiveConfig {
	s.RocketMQ = v
	return s
}

type AiStoreReceiveConfigDingTalk struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Secret  *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
}

func (s AiStoreReceiveConfigDingTalk) String() string {
	return tea.Prettify(s)
}

func (s AiStoreReceiveConfigDingTalk) GoString() string {
	return s.String()
}

func (s *AiStoreReceiveConfigDingTalk) SetAddress(v string) *AiStoreReceiveConfigDingTalk {
	s.Address = &v
	return s
}

func (s *AiStoreReceiveConfigDingTalk) SetSecret(v string) *AiStoreReceiveConfigDingTalk {
	s.Secret = &v
	return s
}

type AiStoreReceiveConfigEventBridge struct {
	EventBus  *string `json:"EventBus,omitempty" xml:"EventBus,omitempty"`
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

type AiStoreReceiveConfigHttp struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s AiStoreReceiveConfigHttp) String() string {
	return tea.Prettify(s)
}

func (s AiStoreReceiveConfigHttp) GoString() string {
	return s.String()
}

func (s *AiStoreReceiveConfigHttp) SetUrl(v string) *AiStoreReceiveConfigHttp {
	s.Url = &v
	return s
}

type AiStoreReceiveConfigHttps struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s AiStoreReceiveConfigHttps) String() string {
	return tea.Prettify(s)
}

func (s AiStoreReceiveConfigHttps) GoString() string {
	return s.String()
}

func (s *AiStoreReceiveConfigHttps) SetUrl(v string) *AiStoreReceiveConfigHttps {
	s.Url = &v
	return s
}

type AiStoreReceiveConfigMns struct {
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

type AiStoreReceiveConfigRocketMQ struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TopicName  *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s AiStoreReceiveConfigRocketMQ) String() string {
	return tea.Prettify(s)
}

func (s AiStoreReceiveConfigRocketMQ) GoString() string {
	return s.String()
}

func (s *AiStoreReceiveConfigRocketMQ) SetInstanceId(v string) *AiStoreReceiveConfigRocketMQ {
	s.InstanceId = &v
	return s
}

func (s *AiStoreReceiveConfigRocketMQ) SetTopicName(v string) *AiStoreReceiveConfigRocketMQ {
	s.TopicName = &v
	return s
}

type AiStoreTemplate struct {
	TemplateContext  *string `json:"TemplateContext,omitempty" xml:"TemplateContext,omitempty"`
	TemplateVariable *string `json:"TemplateVariable,omitempty" xml:"TemplateVariable,omitempty"`
}

func (s AiStoreTemplate) String() string {
	return tea.Prettify(s)
}

func (s AiStoreTemplate) GoString() string {
	return s.String()
}

func (s *AiStoreTemplate) SetTemplateContext(v string) *AiStoreTemplate {
	s.TemplateContext = &v
	return s
}

func (s *AiStoreTemplate) SetTemplateVariable(v string) *AiStoreTemplate {
	s.TemplateVariable = &v
	return s
}

type AiStoreUserTask struct {
	ApiDesc         *string `json:"ApiDesc,omitempty" xml:"ApiDesc,omitempty"`
	ApiName         *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BucketKeyPrefix *string `json:"BucketKeyPrefix,omitempty" xml:"BucketKeyPrefix,omitempty"`
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	DisableTime     *string `json:"DisableTime,omitempty" xml:"DisableTime,omitempty"`
	EnableTime      *string `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamInfo       *string `json:"ParamInfo,omitempty" xml:"ParamInfo,omitempty"`
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ProductDesc     *string `json:"ProductDesc,omitempty" xml:"ProductDesc,omitempty"`
	ReceiveConfig   *string `json:"ReceiveConfig,omitempty" xml:"ReceiveConfig,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionDesc      *string `json:"RegionDesc,omitempty" xml:"RegionDesc,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskVersion     *string `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s AiStoreUserTask) String() string {
	return tea.Prettify(s)
}

func (s AiStoreUserTask) GoString() string {
	return s.String()
}

func (s *AiStoreUserTask) SetApiDesc(v string) *AiStoreUserTask {
	s.ApiDesc = &v
	return s
}

func (s *AiStoreUserTask) SetApiName(v string) *AiStoreUserTask {
	s.ApiName = &v
	return s
}

func (s *AiStoreUserTask) SetBucketKeyPrefix(v string) *AiStoreUserTask {
	s.BucketKeyPrefix = &v
	return s
}

func (s *AiStoreUserTask) SetBucketName(v string) *AiStoreUserTask {
	s.BucketName = &v
	return s
}

func (s *AiStoreUserTask) SetDisableTime(v string) *AiStoreUserTask {
	s.DisableTime = &v
	return s
}

func (s *AiStoreUserTask) SetEnableTime(v string) *AiStoreUserTask {
	s.EnableTime = &v
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

func (s *AiStoreUserTask) SetId(v int64) *AiStoreUserTask {
	s.Id = &v
	return s
}

func (s *AiStoreUserTask) SetName(v string) *AiStoreUserTask {
	s.Name = &v
	return s
}

func (s *AiStoreUserTask) SetParamInfo(v string) *AiStoreUserTask {
	s.ParamInfo = &v
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

func (s *AiStoreUserTask) SetReceiveConfig(v string) *AiStoreUserTask {
	s.ReceiveConfig = &v
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

func (s *AiStoreUserTask) SetRemark(v string) *AiStoreUserTask {
	s.Remark = &v
	return s
}

func (s *AiStoreUserTask) SetStatus(v string) *AiStoreUserTask {
	s.Status = &v
	return s
}

func (s *AiStoreUserTask) SetTaskVersion(v string) *AiStoreUserTask {
	s.TaskVersion = &v
	return s
}

func (s *AiStoreUserTask) SetVersion(v string) *AiStoreUserTask {
	s.Version = &v
	return s
}

type CheckServiceLinkedRoleForDeletingRequest struct {
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
	RoleArn        *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	SPIRegionId    *string `json:"SPIRegionId,omitempty" xml:"SPIRegionId,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CheckServiceLinkedRoleForDeletingRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetAccountId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.AccountId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetDeletionTaskId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.DeletionTaskId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetRoleArn(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.RoleArn = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetSPIRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.SPIRegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetServiceName(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.ServiceName = &v
	return s
}

type CheckServiceLinkedRoleForDeletingResponseBody struct {
	Deletable  *bool                                                      `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleUsages []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages `json:"RoleUsages,omitempty" xml:"RoleUsages,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetDeletable(v bool) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.Deletable = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceLinkedRoleForDeletingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CheckServiceLinkedRoleForDeletingResponse) SetStatusCode(v int32) *CheckServiceLinkedRoleForDeletingResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponse) SetBody(v *CheckServiceLinkedRoleForDeletingResponseBody) *CheckServiceLinkedRoleForDeletingResponse {
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
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAiStoreBucketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreBucketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiStoreBucketResponseBody) SetData(v string) *CreateAiStoreBucketResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAiStoreBucketResponseBody) SetRequestId(v string) *CreateAiStoreBucketResponseBody {
	s.RequestId = &v
	return s
}

type CreateAiStoreBucketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiStoreBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAiStoreBucketResponse) SetStatusCode(v int32) *CreateAiStoreBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiStoreBucketResponse) SetBody(v *CreateAiStoreBucketResponseBody) *CreateAiStoreBucketResponse {
	s.Body = v
	return s
}

type CreateAiStoreReceiveConfigRequest struct {
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s CreateAiStoreReceiveConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreReceiveConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAiStoreReceiveConfigRequest) SetApiName(v string) *CreateAiStoreReceiveConfigRequest {
	s.ApiName = &v
	return s
}

func (s *CreateAiStoreReceiveConfigRequest) SetProduct(v string) *CreateAiStoreReceiveConfigRequest {
	s.Product = &v
	return s
}

type CreateAiStoreReceiveConfigResponseBody struct {
	Data      *AiStoreReceiveConfig `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAiStoreReceiveConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreReceiveConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiStoreReceiveConfigResponseBody) SetData(v *AiStoreReceiveConfig) *CreateAiStoreReceiveConfigResponseBody {
	s.Data = v
	return s
}

func (s *CreateAiStoreReceiveConfigResponseBody) SetRequestId(v string) *CreateAiStoreReceiveConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateAiStoreReceiveConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiStoreReceiveConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAiStoreReceiveConfigResponse) SetStatusCode(v int32) *CreateAiStoreReceiveConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiStoreReceiveConfigResponse) SetBody(v *CreateAiStoreReceiveConfigResponseBody) *CreateAiStoreReceiveConfigResponse {
	s.Body = v
	return s
}

type CreateAiStoreUserTaskRequest struct {
	ApiName         *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BucketKeyPrefix *string `json:"BucketKeyPrefix,omitempty" xml:"BucketKeyPrefix,omitempty"`
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	CreateConfig    *string `json:"CreateConfig,omitempty" xml:"CreateConfig,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamInfo       *string `json:"ParamInfo,omitempty" xml:"ParamInfo,omitempty"`
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ReceiveConfig   *string `json:"ReceiveConfig,omitempty" xml:"ReceiveConfig,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAiStoreUserTaskRequest) SetApiName(v string) *CreateAiStoreUserTaskRequest {
	s.ApiName = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetBucketKeyPrefix(v string) *CreateAiStoreUserTaskRequest {
	s.BucketKeyPrefix = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetBucketName(v string) *CreateAiStoreUserTaskRequest {
	s.BucketName = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetCreateConfig(v string) *CreateAiStoreUserTaskRequest {
	s.CreateConfig = &v
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

func (s *CreateAiStoreUserTaskRequest) SetProduct(v string) *CreateAiStoreUserTaskRequest {
	s.Product = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetReceiveConfig(v string) *CreateAiStoreUserTaskRequest {
	s.ReceiveConfig = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetRemark(v string) *CreateAiStoreUserTaskRequest {
	s.Remark = &v
	return s
}

func (s *CreateAiStoreUserTaskRequest) SetStatus(v string) *CreateAiStoreUserTaskRequest {
	s.Status = &v
	return s
}

type CreateAiStoreUserTaskResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAiStoreUserTaskResponseBody) SetData(v int64) *CreateAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAiStoreUserTaskResponseBody) SetRequestId(v string) *CreateAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateAiStoreUserTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAiStoreUserTaskResponse) SetStatusCode(v int32) *CreateAiStoreUserTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiStoreUserTaskResponse) SetBody(v *CreateAiStoreUserTaskResponseBody) *CreateAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type DeleteAiStoreUserTaskRequest struct {
	AistoreVersion *string `json:"AistoreVersion,omitempty" xml:"AistoreVersion,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteAiStoreUserTaskRequest) SetAistoreVersion(v string) *DeleteAiStoreUserTaskRequest {
	s.AistoreVersion = &v
	return s
}

func (s *DeleteAiStoreUserTaskRequest) SetId(v int64) *DeleteAiStoreUserTaskRequest {
	s.Id = &v
	return s
}

type DeleteAiStoreUserTaskResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAiStoreUserTaskResponseBody) SetData(v bool) *DeleteAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAiStoreUserTaskResponseBody) SetRequestId(v string) *DeleteAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAiStoreUserTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteAiStoreUserTaskResponse) SetStatusCode(v int32) *DeleteAiStoreUserTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAiStoreUserTaskResponse) SetBody(v *DeleteAiStoreUserTaskResponseBody) *DeleteAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type DisableAiStoreUserTaskRequest struct {
	AistoreVersion *string `json:"AistoreVersion,omitempty" xml:"AistoreVersion,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DisableAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *DisableAiStoreUserTaskRequest) SetAistoreVersion(v string) *DisableAiStoreUserTaskRequest {
	s.AistoreVersion = &v
	return s
}

func (s *DisableAiStoreUserTaskRequest) SetId(v int64) *DisableAiStoreUserTaskRequest {
	s.Id = &v
	return s
}

type DisableAiStoreUserTaskResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAiStoreUserTaskResponseBody) SetData(v bool) *DisableAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DisableAiStoreUserTaskResponseBody) SetRequestId(v string) *DisableAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

type DisableAiStoreUserTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DisableAiStoreUserTaskResponse) SetStatusCode(v int32) *DisableAiStoreUserTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAiStoreUserTaskResponse) SetBody(v *DisableAiStoreUserTaskResponseBody) *DisableAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type EnableAiStoreUserTaskRequest struct {
	AistoreVersion *string `json:"AistoreVersion,omitempty" xml:"AistoreVersion,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s EnableAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *EnableAiStoreUserTaskRequest) SetAistoreVersion(v string) *EnableAiStoreUserTaskRequest {
	s.AistoreVersion = &v
	return s
}

func (s *EnableAiStoreUserTaskRequest) SetId(v int64) *EnableAiStoreUserTaskRequest {
	s.Id = &v
	return s
}

type EnableAiStoreUserTaskResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAiStoreUserTaskResponseBody) SetData(v string) *EnableAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

func (s *EnableAiStoreUserTaskResponseBody) SetRequestId(v string) *EnableAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

type EnableAiStoreUserTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EnableAiStoreUserTaskResponse) SetStatusCode(v int32) *EnableAiStoreUserTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAiStoreUserTaskResponse) SetBody(v *EnableAiStoreUserTaskResponseBody) *EnableAiStoreUserTaskResponse {
	s.Body = v
	return s
}

type GetAiStoreReceiveConfigRequest struct {
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetAiStoreReceiveConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreReceiveConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAiStoreReceiveConfigRequest) SetApiName(v string) *GetAiStoreReceiveConfigRequest {
	s.ApiName = &v
	return s
}

func (s *GetAiStoreReceiveConfigRequest) SetProduct(v string) *GetAiStoreReceiveConfigRequest {
	s.Product = &v
	return s
}

type GetAiStoreReceiveConfigResponseBody struct {
	Data      *AiStoreReceiveConfig `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAiStoreReceiveConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreReceiveConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiStoreReceiveConfigResponseBody) SetData(v *AiStoreReceiveConfig) *GetAiStoreReceiveConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAiStoreReceiveConfigResponseBody) SetRequestId(v string) *GetAiStoreReceiveConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetAiStoreReceiveConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiStoreReceiveConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetAiStoreReceiveConfigResponse) SetStatusCode(v int32) *GetAiStoreReceiveConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiStoreReceiveConfigResponse) SetBody(v *GetAiStoreReceiveConfigResponseBody) *GetAiStoreReceiveConfigResponse {
	s.Body = v
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
	Data      *AiStoreUserTask `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiStoreUserTaskResponseBody) SetData(v *AiStoreUserTask) *GetAiStoreUserTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetAiStoreUserTaskResponseBody) SetRequestId(v string) *GetAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetAiStoreUserTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetAiStoreUserTaskResponse) SetStatusCode(v int32) *GetAiStoreUserTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiStoreUserTaskResponse) SetBody(v *GetAiStoreUserTaskResponseBody) *GetAiStoreUserTaskResponse {
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
	Data      *AiStoreUserTask `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAiStoreUserTaskByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAiStoreUserTaskByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiStoreUserTaskByNameResponseBody) SetData(v *AiStoreUserTask) *GetAiStoreUserTaskByNameResponseBody {
	s.Data = v
	return s
}

func (s *GetAiStoreUserTaskByNameResponseBody) SetRequestId(v string) *GetAiStoreUserTaskByNameResponseBody {
	s.RequestId = &v
	return s
}

type GetAiStoreUserTaskByNameResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiStoreUserTaskByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetAiStoreUserTaskByNameResponse) SetStatusCode(v int32) *GetAiStoreUserTaskByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiStoreUserTaskByNameResponse) SetBody(v *GetAiStoreUserTaskByNameResponseBody) *GetAiStoreUserTaskByNameResponse {
	s.Body = v
	return s
}

type ListAiStoreBucketsRequest struct {
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s ListAiStoreBucketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAiStoreBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListAiStoreBucketsRequest) SetApiName(v string) *ListAiStoreBucketsRequest {
	s.ApiName = &v
	return s
}

func (s *ListAiStoreBucketsRequest) SetProduct(v string) *ListAiStoreBucketsRequest {
	s.Product = &v
	return s
}

type ListAiStoreBucketsResponseBody struct {
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAiStoreBucketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAiStoreBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiStoreBucketsResponseBody) SetData(v []*string) *ListAiStoreBucketsResponseBody {
	s.Data = v
	return s
}

func (s *ListAiStoreBucketsResponseBody) SetRequestId(v string) *ListAiStoreBucketsResponseBody {
	s.RequestId = &v
	return s
}

type ListAiStoreBucketsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiStoreBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListAiStoreBucketsResponse) SetStatusCode(v int32) *ListAiStoreBucketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiStoreBucketsResponse) SetBody(v *ListAiStoreBucketsResponseBody) *ListAiStoreBucketsResponse {
	s.Body = v
	return s
}

type QueryAiStoreApiTreeResponseBody struct {
	Data      []*AiStoreApiNode `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAiStoreApiTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreApiTreeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiStoreApiTreeResponseBody) SetData(v []*AiStoreApiNode) *QueryAiStoreApiTreeResponseBody {
	s.Data = v
	return s
}

func (s *QueryAiStoreApiTreeResponseBody) SetRequestId(v string) *QueryAiStoreApiTreeResponseBody {
	s.RequestId = &v
	return s
}

type QueryAiStoreApiTreeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAiStoreApiTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryAiStoreApiTreeResponse) SetStatusCode(v int32) *QueryAiStoreApiTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAiStoreApiTreeResponse) SetBody(v *QueryAiStoreApiTreeResponseBody) *QueryAiStoreApiTreeResponse {
	s.Body = v
	return s
}

type QueryAiStoreRegionsResponseBody struct {
	Data      []*AiStoreUserTask `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAiStoreRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiStoreRegionsResponseBody) SetData(v []*AiStoreUserTask) *QueryAiStoreRegionsResponseBody {
	s.Data = v
	return s
}

func (s *QueryAiStoreRegionsResponseBody) SetRequestId(v string) *QueryAiStoreRegionsResponseBody {
	s.RequestId = &v
	return s
}

type QueryAiStoreRegionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAiStoreRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryAiStoreRegionsResponse) SetStatusCode(v int32) *QueryAiStoreRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAiStoreRegionsResponse) SetBody(v *QueryAiStoreRegionsResponseBody) *QueryAiStoreRegionsResponse {
	s.Body = v
	return s
}

type QueryAiStoreUserTaskPageRequest struct {
	ApiName    *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryAiStoreUserTaskPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreUserTaskPageRequest) GoString() string {
	return s.String()
}

func (s *QueryAiStoreUserTaskPageRequest) SetApiName(v string) *QueryAiStoreUserTaskPageRequest {
	s.ApiName = &v
	return s
}

func (s *QueryAiStoreUserTaskPageRequest) SetBucketName(v string) *QueryAiStoreUserTaskPageRequest {
	s.BucketName = &v
	return s
}

func (s *QueryAiStoreUserTaskPageRequest) SetName(v string) *QueryAiStoreUserTaskPageRequest {
	s.Name = &v
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

func (s *QueryAiStoreUserTaskPageRequest) SetProduct(v string) *QueryAiStoreUserTaskPageRequest {
	s.Product = &v
	return s
}

func (s *QueryAiStoreUserTaskPageRequest) SetStatus(v string) *QueryAiStoreUserTaskPageRequest {
	s.Status = &v
	return s
}

type QueryAiStoreUserTaskPageResponseBody struct {
	Data      *QueryAiStoreUserTaskPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAiStoreUserTaskPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAiStoreUserTaskPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiStoreUserTaskPageResponseBody) SetData(v *QueryAiStoreUserTaskPageResponseBodyData) *QueryAiStoreUserTaskPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryAiStoreUserTaskPageResponseBody) SetRequestId(v string) *QueryAiStoreUserTaskPageResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAiStoreUserTaskPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryAiStoreUserTaskPageResponse) SetStatusCode(v int32) *QueryAiStoreUserTaskPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAiStoreUserTaskPageResponse) SetBody(v *QueryAiStoreUserTaskPageResponseBody) *QueryAiStoreUserTaskPageResponse {
	s.Body = v
	return s
}

type UpdateAiStoreUserTaskRequest struct {
	ApiName         *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BucketKeyPrefix *string `json:"BucketKeyPrefix,omitempty" xml:"BucketKeyPrefix,omitempty"`
	BucketName      *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamInfo       *string `json:"ParamInfo,omitempty" xml:"ParamInfo,omitempty"`
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ReceiveConfig   *string `json:"ReceiveConfig,omitempty" xml:"ReceiveConfig,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAiStoreUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAiStoreUserTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiStoreUserTaskRequest) SetApiName(v string) *UpdateAiStoreUserTaskRequest {
	s.ApiName = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetBucketKeyPrefix(v string) *UpdateAiStoreUserTaskRequest {
	s.BucketKeyPrefix = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetBucketName(v string) *UpdateAiStoreUserTaskRequest {
	s.BucketName = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetId(v int64) *UpdateAiStoreUserTaskRequest {
	s.Id = &v
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

func (s *UpdateAiStoreUserTaskRequest) SetProduct(v string) *UpdateAiStoreUserTaskRequest {
	s.Product = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetReceiveConfig(v string) *UpdateAiStoreUserTaskRequest {
	s.ReceiveConfig = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetRemark(v string) *UpdateAiStoreUserTaskRequest {
	s.Remark = &v
	return s
}

func (s *UpdateAiStoreUserTaskRequest) SetStatus(v string) *UpdateAiStoreUserTaskRequest {
	s.Status = &v
	return s
}

type UpdateAiStoreUserTaskResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAiStoreUserTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAiStoreUserTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAiStoreUserTaskResponseBody) SetData(v bool) *UpdateAiStoreUserTaskResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAiStoreUserTaskResponseBody) SetRequestId(v string) *UpdateAiStoreUserTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAiStoreUserTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAiStoreUserTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateAiStoreUserTaskResponse) SetStatusCode(v int32) *UpdateAiStoreUserTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAiStoreUserTaskResponse) SetBody(v *UpdateAiStoreUserTaskResponseBody) *UpdateAiStoreUserTaskResponse {
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

func (client *Client) CheckServiceLinkedRoleForDeletingWithOptions(request *CheckServiceLinkedRoleForDeletingRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleForDeletingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.DeletionTaskId)) {
		query["DeletionTaskId"] = request.DeletionTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.SPIRegionId)) {
		query["SPIRegionId"] = request.SPIRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceLinkedRoleForDeleting"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckServiceLinkedRoleForDeletingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateAiStoreBucketWithOptions(request *CreateAiStoreBucketRequest, runtime *util.RuntimeOptions) (_result *CreateAiStoreBucketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		body["BucketName"] = request.BucketName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAiStoreBucket"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAiStoreBucketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateAiStoreReceiveConfigWithOptions(request *CreateAiStoreReceiveConfigRequest, runtime *util.RuntimeOptions) (_result *CreateAiStoreReceiveConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		body["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAiStoreReceiveConfig"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAiStoreReceiveConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateAiStoreUserTaskWithOptions(request *CreateAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *CreateAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		body["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.BucketKeyPrefix)) {
		body["BucketKeyPrefix"] = request.BucketKeyPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		body["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.CreateConfig)) {
		body["CreateConfig"] = request.CreateConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParamInfo)) {
		body["ParamInfo"] = request.ParamInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveConfig)) {
		body["ReceiveConfig"] = request.ReceiveConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAiStoreUserTask"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAiStoreUserTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteAiStoreUserTaskWithOptions(request *DeleteAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AistoreVersion)) {
		body["AistoreVersion"] = request.AistoreVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAiStoreUserTask"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAiStoreUserTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DisableAiStoreUserTaskWithOptions(request *DisableAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *DisableAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AistoreVersion)) {
		body["AistoreVersion"] = request.AistoreVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAiStoreUserTask"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAiStoreUserTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) EnableAiStoreUserTaskWithOptions(request *EnableAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *EnableAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AistoreVersion)) {
		body["AistoreVersion"] = request.AistoreVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAiStoreUserTask"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableAiStoreUserTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetAiStoreReceiveConfigWithOptions(request *GetAiStoreReceiveConfigRequest, runtime *util.RuntimeOptions) (_result *GetAiStoreReceiveConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		body["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAiStoreReceiveConfig"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAiStoreReceiveConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetAiStoreUserTaskWithOptions(request *GetAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *GetAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAiStoreUserTask"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAiStoreUserTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetAiStoreUserTaskByNameWithOptions(request *GetAiStoreUserTaskByNameRequest, runtime *util.RuntimeOptions) (_result *GetAiStoreUserTaskByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAiStoreUserTaskByName"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAiStoreUserTaskByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListAiStoreBucketsWithOptions(request *ListAiStoreBucketsRequest, runtime *util.RuntimeOptions) (_result *ListAiStoreBucketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		body["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAiStoreBuckets"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAiStoreBucketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QueryAiStoreApiTreeWithOptions(runtime *util.RuntimeOptions) (_result *QueryAiStoreApiTreeResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryAiStoreApiTree"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAiStoreApiTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QueryAiStoreRegionsWithOptions(runtime *util.RuntimeOptions) (_result *QueryAiStoreRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryAiStoreRegions"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAiStoreRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QueryAiStoreUserTaskPageWithOptions(request *QueryAiStoreUserTaskPageRequest, runtime *util.RuntimeOptions) (_result *QueryAiStoreUserTaskPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAiStoreUserTaskPage"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAiStoreUserTaskPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateAiStoreUserTaskWithOptions(request *UpdateAiStoreUserTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateAiStoreUserTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		body["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.BucketKeyPrefix)) {
		body["BucketKeyPrefix"] = request.BucketKeyPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		body["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParamInfo)) {
		body["ParamInfo"] = request.ParamInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveConfig)) {
		body["ReceiveConfig"] = request.ReceiveConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAiStoreUserTask"),
		Version:     tea.String("2021-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAiStoreUserTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
