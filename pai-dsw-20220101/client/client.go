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

type DemoCategory struct {
	// 目录码
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// 目录名
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// 序号
	Order *int64 `json:"Order,omitempty" xml:"Order,omitempty"`
	// 目录列表
	SubCategories []*DemoCategory `json:"SubCategories,omitempty" xml:"SubCategories,omitempty" type:"Repeated"`
}

func (s DemoCategory) String() string {
	return tea.Prettify(s)
}

func (s DemoCategory) GoString() string {
	return s.String()
}

func (s *DemoCategory) SetCategoryCode(v string) *DemoCategory {
	s.CategoryCode = &v
	return s
}

func (s *DemoCategory) SetCategoryName(v string) *DemoCategory {
	s.CategoryName = &v
	return s
}

func (s *DemoCategory) SetOrder(v int64) *DemoCategory {
	s.Order = &v
	return s
}

func (s *DemoCategory) SetSubCategories(v []*DemoCategory) *DemoCategory {
	s.SubCategories = v
	return s
}

type CreateInstanceRequest struct {
	// 工作空间内是否他人可见
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 数据集集合
	Datasets []*CreateInstanceRequestDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// 实例对应的Ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 环境变量
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// user vpc配置
	UserVpc *CreateInstanceRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAccessibility(v string) *CreateInstanceRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateInstanceRequest) SetDatasets(v []*CreateInstanceRequestDatasets) *CreateInstanceRequest {
	s.Datasets = v
	return s
}

func (s *CreateInstanceRequest) SetEcsSpec(v string) *CreateInstanceRequest {
	s.EcsSpec = &v
	return s
}

func (s *CreateInstanceRequest) SetEnvironmentVariables(v map[string]*string) *CreateInstanceRequest {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetImageUrl(v string) *CreateInstanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetUserVpc(v *CreateInstanceRequestUserVpc) *CreateInstanceRequest {
	s.UserVpc = v
	return s
}

func (s *CreateInstanceRequest) SetWorkspaceId(v string) *CreateInstanceRequest {
	s.WorkspaceId = &v
	return s
}

type CreateInstanceRequestDatasets struct {
	// 数据集Id
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 容器内挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s CreateInstanceRequestDatasets) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestDatasets) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestDatasets) SetDatasetId(v string) *CreateInstanceRequestDatasets {
	s.DatasetId = &v
	return s
}

func (s *CreateInstanceRequestDatasets) SetMountPath(v string) *CreateInstanceRequestDatasets {
	s.MountPath = &v
	return s
}

type CreateInstanceRequestUserVpc struct {
	// Security Group Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// VSwitch Id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateInstanceRequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestUserVpc) SetSecurityGroupId(v string) *CreateInstanceRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateInstanceRequestUserVpc) SetVSwitchId(v string) *CreateInstanceRequestUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequestUserVpc) SetVpcId(v string) *CreateInstanceRequestUserVpc {
	s.VpcId = &v
	return s
}

type CreateInstanceResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetHttpStatusCode(v int32) *CreateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
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

type CreateInstanceShutdownTimerRequest struct {
	// 定时关机设定时间
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// 距离定时关机时间段
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s CreateInstanceShutdownTimerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerRequest) SetDueTime(v string) *CreateInstanceShutdownTimerRequest {
	s.DueTime = &v
	return s
}

func (s *CreateInstanceShutdownTimerRequest) SetRemainingTimeInMs(v int64) *CreateInstanceShutdownTimerRequest {
	s.RemainingTimeInMs = &v
	return s
}

type CreateInstanceShutdownTimerResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceShutdownTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerResponseBody) SetCode(v string) *CreateInstanceShutdownTimerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetHttpStatusCode(v int32) *CreateInstanceShutdownTimerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetInstanceId(v string) *CreateInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetMessage(v string) *CreateInstanceShutdownTimerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetRequestId(v string) *CreateInstanceShutdownTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetSuccess(v bool) *CreateInstanceShutdownTimerResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceShutdownTimerResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceShutdownTimerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShutdownTimerResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerResponse) SetHeaders(v map[string]*string) *CreateInstanceShutdownTimerResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceShutdownTimerResponse) SetBody(v *CreateInstanceShutdownTimerResponseBody) *CreateInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

type CreateInstanceSnapshotRequest struct {
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例快照描述
	SnapshotDescription *string `json:"SnapshotDescription,omitempty" xml:"SnapshotDescription,omitempty"`
	// 实例快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateInstanceSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotRequest) SetImageUrl(v string) *CreateInstanceSnapshotRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetSnapshotDescription(v string) *CreateInstanceSnapshotRequest {
	s.SnapshotDescription = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetSnapshotName(v string) *CreateInstanceSnapshotRequest {
	s.SnapshotName = &v
	return s
}

type CreateInstanceSnapshotResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例快照Id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotResponseBody) SetCode(v string) *CreateInstanceSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetHttpStatusCode(v int32) *CreateInstanceSnapshotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetInstanceId(v string) *CreateInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetMessage(v string) *CreateInstanceSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetRequestId(v string) *CreateInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetSnapshotId(v string) *CreateInstanceSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetSuccess(v bool) *CreateInstanceSnapshotResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceSnapshotResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotResponse) SetHeaders(v map[string]*string) *CreateInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceSnapshotResponse) SetBody(v *CreateInstanceSnapshotResponseBody) *CreateInstanceSnapshotResponse {
	s.Body = v
	return s
}

type DeleteInstanceResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetCode(v string) *DeleteInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetInstanceId(v string) *DeleteInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetMessage(v string) *DeleteInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
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

type DeleteInstanceShutdownTimerResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceShutdownTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetCode(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceShutdownTimerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetInstanceId(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetMessage(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetRequestId(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetSuccess(v bool) *DeleteInstanceShutdownTimerResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceShutdownTimerResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceShutdownTimerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceShutdownTimerResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShutdownTimerResponse) SetHeaders(v map[string]*string) *DeleteInstanceShutdownTimerResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceShutdownTimerResponse) SetBody(v *DeleteInstanceShutdownTimerResponseBody) *DeleteInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

type DeleteInstanceSnapshotResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例快照Id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceSnapshotResponseBody) SetCode(v string) *DeleteInstanceSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceSnapshotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetInstanceId(v string) *DeleteInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetMessage(v string) *DeleteInstanceSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetRequestId(v string) *DeleteInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetSnapshotId(v string) *DeleteInstanceSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetSuccess(v bool) *DeleteInstanceSnapshotResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceSnapshotResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceSnapshotResponse) SetHeaders(v map[string]*string) *DeleteInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceSnapshotResponse) SetBody(v *DeleteInstanceSnapshotResponseBody) *DeleteInstanceSnapshotResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	// 实例计算类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 工作空间内是否他人可见
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 累计运行时间（ms）
	AccumulatedRunningTimeInMs *int64 `json:"AccumulatedRunningTimeInMs,omitempty" xml:"AccumulatedRunningTimeInMs,omitempty"`
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 数据集集合
	Datasets []*GetInstanceResponseBodyDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// 实例对应的Ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 环境变量
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// 实例创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 定时关机任务
	InstanceShutdownTimer *GetInstanceResponseBodyInstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty" type:"Struct"`
	// 实例Url
	InstanceUrl *string `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	// Jupyterlab Url
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// 最新保存的用户镜像
	LatestSnapshot *GetInstanceResponseBodyLatestSnapshot `json:"LatestSnapshot,omitempty" xml:"LatestSnapshot,omitempty" type:"Struct"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 支付类型
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 实例错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例错误原因
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 终端url
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	// 用户Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户名称
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// user vpc配置
	UserVpc *GetInstanceResponseBodyUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// Web IDE url
	WebIDEUrl *string `json:"WebIDEUrl,omitempty" xml:"WebIDEUrl,omitempty"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetAcceleratorType(v string) *GetInstanceResponseBody {
	s.AcceleratorType = &v
	return s
}

func (s *GetInstanceResponseBody) SetAccessibility(v string) *GetInstanceResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetInstanceResponseBody) SetAccumulatedRunningTimeInMs(v int64) *GetInstanceResponseBody {
	s.AccumulatedRunningTimeInMs = &v
	return s
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetDatasets(v []*GetInstanceResponseBodyDatasets) *GetInstanceResponseBody {
	s.Datasets = v
	return s
}

func (s *GetInstanceResponseBody) SetEcsSpec(v string) *GetInstanceResponseBody {
	s.EcsSpec = &v
	return s
}

func (s *GetInstanceResponseBody) SetEnvironmentVariables(v map[string]*string) *GetInstanceResponseBody {
	s.EnvironmentVariables = v
	return s
}

func (s *GetInstanceResponseBody) SetGmtCreateTime(v string) *GetInstanceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetGmtModifiedTime(v string) *GetInstanceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v int32) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetImageId(v string) *GetInstanceResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetInstanceResponseBody) SetImageName(v string) *GetInstanceResponseBody {
	s.ImageName = &v
	return s
}

func (s *GetInstanceResponseBody) SetImageUrl(v string) *GetInstanceResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceId(v string) *GetInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceName(v string) *GetInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceShutdownTimer(v *GetInstanceResponseBodyInstanceShutdownTimer) *GetInstanceResponseBody {
	s.InstanceShutdownTimer = v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceUrl(v string) *GetInstanceResponseBody {
	s.InstanceUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetJupyterlabUrl(v string) *GetInstanceResponseBody {
	s.JupyterlabUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetLatestSnapshot(v *GetInstanceResponseBodyLatestSnapshot) *GetInstanceResponseBody {
	s.LatestSnapshot = v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetPaymentType(v string) *GetInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *GetInstanceResponseBody) SetReasonCode(v string) *GetInstanceResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetReasonMessage(v string) *GetInstanceResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetStatus(v string) *GetInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceResponseBody) SetTerminalUrl(v string) *GetInstanceResponseBody {
	s.TerminalUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserId(v string) *GetInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserName(v string) *GetInstanceResponseBody {
	s.UserName = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserVpc(v *GetInstanceResponseBodyUserVpc) *GetInstanceResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetInstanceResponseBody) SetWebIDEUrl(v string) *GetInstanceResponseBody {
	s.WebIDEUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetWorkspaceId(v string) *GetInstanceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetWorkspaceName(v string) *GetInstanceResponseBody {
	s.WorkspaceName = &v
	return s
}

type GetInstanceResponseBodyDatasets struct {
	// 数据集Id
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 容器内挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s GetInstanceResponseBodyDatasets) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDatasets) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDatasets) SetDatasetId(v string) *GetInstanceResponseBodyDatasets {
	s.DatasetId = &v
	return s
}

func (s *GetInstanceResponseBodyDatasets) SetMountPath(v string) *GetInstanceResponseBodyDatasets {
	s.MountPath = &v
	return s
}

type GetInstanceResponseBodyInstanceShutdownTimer struct {
	// 设定关机时间
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 剩余关机时间（ms）
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s GetInstanceResponseBodyInstanceShutdownTimer) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceShutdownTimer) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetDueTime(v string) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.DueTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetGmtCreateTime(v string) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetGmtModifiedTime(v string) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetInstanceId(v string) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetRemainingTimeInMs(v int64) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.RemainingTimeInMs = &v
	return s
}

type GetInstanceResponseBodyLatestSnapshot struct {
	// 快照创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 快照修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像Url
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 镜像仓库Url
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
}

func (s GetInstanceResponseBodyLatestSnapshot) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyLatestSnapshot) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetGmtCreateTime(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetGmtModifiedTime(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetImageId(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ImageId = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetImageName(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ImageName = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetImageUrl(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetRepositoryUrl(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.RepositoryUrl = &v
	return s
}

type GetInstanceResponseBodyUserVpc struct {
	// Security Group Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// VSwitch Id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetInstanceResponseBodyUserVpc) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyUserVpc) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyUserVpc) SetSecurityGroupId(v string) *GetInstanceResponseBodyUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) SetVSwitchId(v string) *GetInstanceResponseBodyUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) SetVpcId(v string) *GetInstanceResponseBodyUserVpc {
	s.VpcId = &v
	return s
}

type GetInstanceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetInstanceShutdownTimerResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 设定关机时间
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 剩余关机时间（ms）
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceShutdownTimerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceShutdownTimerResponseBody) SetCode(v string) *GetInstanceShutdownTimerResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetDueTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.DueTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetGmtCreateTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetGmtModifiedTime(v string) *GetInstanceShutdownTimerResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetHttpStatusCode(v int32) *GetInstanceShutdownTimerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetInstanceId(v string) *GetInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetMessage(v string) *GetInstanceShutdownTimerResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetRemainingTimeInMs(v int64) *GetInstanceShutdownTimerResponseBody {
	s.RemainingTimeInMs = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetRequestId(v string) *GetInstanceShutdownTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceShutdownTimerResponseBody) SetSuccess(v bool) *GetInstanceShutdownTimerResponseBody {
	s.Success = &v
	return s
}

type GetInstanceShutdownTimerResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceShutdownTimerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceShutdownTimerResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceShutdownTimerResponse) SetHeaders(v map[string]*string) *GetInstanceShutdownTimerResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceShutdownTimerResponse) SetBody(v *GetInstanceShutdownTimerResponseBody) *GetInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

type GetInstanceSnapshotResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 实例快照创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例快照修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例快照的镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 实例快照的镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 实例快照错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例快照错误消息
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 实例快照Id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// 实例快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// 实例快照状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotResponseBody) SetCode(v string) *GetInstanceSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetGmtCreateTime(v string) *GetInstanceSnapshotResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetGmtModifiedTime(v string) *GetInstanceSnapshotResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetHttpStatusCode(v int32) *GetInstanceSnapshotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetImageId(v string) *GetInstanceSnapshotResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetImageUrl(v string) *GetInstanceSnapshotResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetInstanceId(v string) *GetInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetMessage(v string) *GetInstanceSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetReasonCode(v string) *GetInstanceSnapshotResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetReasonMessage(v string) *GetInstanceSnapshotResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetRequestId(v string) *GetInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetSnapshotId(v string) *GetInstanceSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetSnapshotName(v string) *GetInstanceSnapshotResponseBody {
	s.SnapshotName = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetStatus(v string) *GetInstanceSnapshotResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceSnapshotResponseBody) SetSuccess(v bool) *GetInstanceSnapshotResponseBody {
	s.Success = &v
	return s
}

type GetInstanceSnapshotResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceSnapshotResponse) SetHeaders(v map[string]*string) *GetInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceSnapshotResponse) SetBody(v *GetInstanceSnapshotResponseBody) *GetInstanceSnapshotResponse {
	s.Body = v
	return s
}

type GetUserConfigResponseBody struct {
	// 用户账号金额是否充足
	AccountSufficient *bool `json:"AccountSufficient,omitempty" xml:"AccountSufficient,omitempty"`
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponseBody) SetAccountSufficient(v bool) *GetUserConfigResponseBody {
	s.AccountSufficient = &v
	return s
}

func (s *GetUserConfigResponseBody) SetCode(v string) *GetUserConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserConfigResponseBody) SetHttpStatusCode(v int32) *GetUserConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserConfigResponseBody) SetMessage(v string) *GetUserConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserConfigResponseBody) SetRequestId(v string) *GetUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserConfigResponseBody) SetSuccess(v bool) *GetUserConfigResponseBody {
	s.Success = &v
	return s
}

type GetUserConfigResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserConfigResponse) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponse) SetHeaders(v map[string]*string) *GetUserConfigResponse {
	s.Headers = v
	return s
}

func (s *GetUserConfigResponse) SetBody(v *GetUserConfigResponseBody) *GetUserConfigResponse {
	s.Body = v
	return s
}

type ListDemoCategoriesResponseBody struct {
	// 样例列表
	Categories []*DemoCategory `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDemoCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDemoCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDemoCategoriesResponseBody) SetCategories(v []*DemoCategory) *ListDemoCategoriesResponseBody {
	s.Categories = v
	return s
}

func (s *ListDemoCategoriesResponseBody) SetRequestId(v string) *ListDemoCategoriesResponseBody {
	s.RequestId = &v
	return s
}

type ListDemoCategoriesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDemoCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDemoCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDemoCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListDemoCategoriesResponse) SetHeaders(v map[string]*string) *ListDemoCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListDemoCategoriesResponse) SetBody(v *ListDemoCategoriesResponseBody) *ListDemoCategoriesResponse {
	s.Body = v
	return s
}

type ListDemosRequest struct {
	// 目录
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DemoName   *string `json:"DemoName,omitempty" xml:"DemoName,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDemosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDemosRequest) GoString() string {
	return s.String()
}

func (s *ListDemosRequest) SetCategory(v string) *ListDemosRequest {
	s.Category = &v
	return s
}

func (s *ListDemosRequest) SetDemoName(v string) *ListDemosRequest {
	s.DemoName = &v
	return s
}

func (s *ListDemosRequest) SetPageNumber(v int64) *ListDemosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDemosRequest) SetPageSize(v int64) *ListDemosRequest {
	s.PageSize = &v
	return s
}

type ListDemosResponseBody struct {
	// 样例列表
	Demos []*ListDemosResponseBodyDemos `json:"Demos,omitempty" xml:"Demos,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDemosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDemosResponseBody) GoString() string {
	return s.String()
}

func (s *ListDemosResponseBody) SetDemos(v []*ListDemosResponseBodyDemos) *ListDemosResponseBody {
	s.Demos = v
	return s
}

func (s *ListDemosResponseBody) SetRequestId(v string) *ListDemosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDemosResponseBody) SetTotalCount(v int64) *ListDemosResponseBody {
	s.TotalCount = &v
	return s
}

type ListDemosResponseBodyDemos struct {
	// 所在目录
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// 样例描述
	DemoDescription *string `json:"DemoDescription,omitempty" xml:"DemoDescription,omitempty"`
	// 样例名称
	DemoName *string `json:"DemoName,omitempty" xml:"DemoName,omitempty"`
	// 样例地址
	DemoUrl *string `json:"DemoUrl,omitempty" xml:"DemoUrl,omitempty"`
	// 序号
	Order *int64 `json:"Order,omitempty" xml:"Order,omitempty"`
	// 大小
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListDemosResponseBodyDemos) String() string {
	return tea.Prettify(s)
}

func (s ListDemosResponseBodyDemos) GoString() string {
	return s.String()
}

func (s *ListDemosResponseBodyDemos) SetCategories(v []*string) *ListDemosResponseBodyDemos {
	s.Categories = v
	return s
}

func (s *ListDemosResponseBodyDemos) SetDemoDescription(v string) *ListDemosResponseBodyDemos {
	s.DemoDescription = &v
	return s
}

func (s *ListDemosResponseBodyDemos) SetDemoName(v string) *ListDemosResponseBodyDemos {
	s.DemoName = &v
	return s
}

func (s *ListDemosResponseBodyDemos) SetDemoUrl(v string) *ListDemosResponseBodyDemos {
	s.DemoUrl = &v
	return s
}

func (s *ListDemosResponseBodyDemos) SetOrder(v int64) *ListDemosResponseBodyDemos {
	s.Order = &v
	return s
}

func (s *ListDemosResponseBodyDemos) SetSize(v int64) *ListDemosResponseBodyDemos {
	s.Size = &v
	return s
}

type ListDemosResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDemosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDemosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDemosResponse) GoString() string {
	return s.String()
}

func (s *ListDemosResponse) SetHeaders(v map[string]*string) *ListDemosResponse {
	s.Headers = v
	return s
}

func (s *ListDemosResponse) SetBody(v *ListDemosResponseBody) *ListDemosResponse {
	s.Body = v
	return s
}

type ListEcsSpecsRequest struct {
	// 加速类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 页数
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 排序字段
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListEcsSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsRequest) SetAcceleratorType(v string) *ListEcsSpecsRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListEcsSpecsRequest) SetOrder(v string) *ListEcsSpecsRequest {
	s.Order = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageNumber(v int64) *ListEcsSpecsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageSize(v int64) *ListEcsSpecsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEcsSpecsRequest) SetSortBy(v string) *ListEcsSpecsRequest {
	s.SortBy = &v
	return s
}

type ListEcsSpecsResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 本分页中请求的实例列表
	EcsSpecs []*ListEcsSpecsResponseBodyEcsSpecs `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEcsSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBody) SetCode(v string) *ListEcsSpecsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetEcsSpecs(v []*ListEcsSpecsResponseBodyEcsSpecs) *ListEcsSpecsResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListEcsSpecsResponseBody) SetHttpStatusCode(v int32) *ListEcsSpecsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetMessage(v string) *ListEcsSpecsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetRequestId(v string) *ListEcsSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetSuccess(v bool) *ListEcsSpecsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetTotalCount(v int64) *ListEcsSpecsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEcsSpecsResponseBodyEcsSpecs struct {
	// 资源类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// CPU核数
	CPU *int64 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// 货币单位
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// GPU卡数
	GPU *int64 `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// 显卡类型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// 实例接收带宽
	InstanceBandwidthRx *int64 `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	// 实例规格
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 内存大小(GB)
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 价格
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// 系统盘大小(GB)
	SystemDiskCapacity *int64 `json:"SystemDiskCapacity,omitempty" xml:"SystemDiskCapacity,omitempty"`
}

func (s ListEcsSpecsResponseBodyEcsSpecs) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponseBodyEcsSpecs) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetAcceleratorType(v string) *ListEcsSpecsResponseBodyEcsSpecs {
	s.AcceleratorType = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetCPU(v int64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.CPU = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetCurrency(v string) *ListEcsSpecsResponseBodyEcsSpecs {
	s.Currency = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetGPU(v int64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.GPU = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetGPUType(v string) *ListEcsSpecsResponseBodyEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetInstanceBandwidthRx(v int64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetInstanceType(v string) *ListEcsSpecsResponseBodyEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetMemory(v float32) *ListEcsSpecsResponseBodyEcsSpecs {
	s.Memory = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetPrice(v float64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.Price = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetSystemDiskCapacity(v int64) *ListEcsSpecsResponseBodyEcsSpecs {
	s.SystemDiskCapacity = &v
	return s
}

type ListEcsSpecsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEcsSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEcsSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponse) SetHeaders(v map[string]*string) *ListEcsSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListEcsSpecsResponse) SetBody(v *ListEcsSpecsResponseBody) *ListEcsSpecsResponse {
	s.Body = v
	return s
}

type ListInstanceSnapshotRequest struct {
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListInstanceSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotRequest) SetOrder(v string) *ListInstanceSnapshotRequest {
	s.Order = &v
	return s
}

func (s *ListInstanceSnapshotRequest) SetPageNumber(v int64) *ListInstanceSnapshotRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceSnapshotRequest) SetPageSize(v int64) *ListInstanceSnapshotRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceSnapshotRequest) SetSortBy(v string) *ListInstanceSnapshotRequest {
	s.SortBy = &v
	return s
}

type ListInstanceSnapshotResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 本分页中请求的实例镜像列表
	Snapshots []*ListInstanceSnapshotResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotResponseBody) SetCode(v string) *ListInstanceSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetHttpStatusCode(v int32) *ListInstanceSnapshotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetMessage(v string) *ListInstanceSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetRequestId(v string) *ListInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetSnapshots(v []*ListInstanceSnapshotResponseBodySnapshots) *ListInstanceSnapshotResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetSuccess(v bool) *ListInstanceSnapshotResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceSnapshotResponseBody) SetTotalCount(v int64) *ListInstanceSnapshotResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstanceSnapshotResponseBodySnapshots struct {
	// 实例快照创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例快照修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例快照的镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 实例快照的镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例快照错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例快照错误消息
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 实例快照Id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// 实例快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// 实例快照状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceSnapshotResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetGmtCreateTime(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetGmtModifiedTime(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetImageId(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.ImageId = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetImageUrl(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.ImageUrl = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetInstanceId(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetReasonCode(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.ReasonCode = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetReasonMessage(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetSnapshotId(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetSnapshotName(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshots) SetStatus(v string) *ListInstanceSnapshotResponseBodySnapshots {
	s.Status = &v
	return s
}

type ListInstanceSnapshotResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotResponse) SetHeaders(v map[string]*string) *ListInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceSnapshotResponse) SetBody(v *ListInstanceSnapshotResponseBody) *ListInstanceSnapshotResponse {
	s.Body = v
	return s
}

type ListInstanceStatisticsRequest struct {
	// 工作空间列表
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s ListInstanceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceStatisticsRequest) SetWorkspaceIds(v string) *ListInstanceStatisticsRequest {
	s.WorkspaceIds = &v
	return s
}

type ListInstanceStatisticsResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 统计信息
	Statistics map[string]map[string]interface{} `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstanceStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceStatisticsResponseBody) SetCode(v string) *ListInstanceStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetHttpStatusCode(v int32) *ListInstanceStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetMessage(v string) *ListInstanceStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetRequestId(v string) *ListInstanceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetStatistics(v map[string]map[string]interface{}) *ListInstanceStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *ListInstanceStatisticsResponseBody) SetSuccess(v bool) *ListInstanceStatisticsResponseBody {
	s.Success = &v
	return s
}

type ListInstanceStatisticsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceStatisticsResponse) SetHeaders(v map[string]*string) *ListInstanceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceStatisticsResponse) SetBody(v *ListInstanceStatisticsResponseBody) *ListInstanceStatisticsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 排列顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页数量大小
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 排序字段
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 实例状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetAcceleratorType(v string) *ListInstancesRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListInstancesRequest) SetAccessibility(v string) *ListInstancesRequest {
	s.Accessibility = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceName(v string) *ListInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesRequest) SetOrder(v string) *ListInstancesRequest {
	s.Order = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int64) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int64) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetPaymentType(v string) *ListInstancesRequest {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesRequest) SetSortBy(v string) *ListInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetWorkspaceId(v string) *ListInstancesRequest {
	s.WorkspaceId = &v
	return s
}

type ListInstancesResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 本分页中请求的实例列表
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	// 实例计算类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 工作空间内是否他人可见
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 累计运行时间（ms）
	AccumulatedRunningTimeInMs *int64 `json:"AccumulatedRunningTimeInMs,omitempty" xml:"AccumulatedRunningTimeInMs,omitempty"`
	// 数据集集合
	Datasets []*ListInstancesResponseBodyInstancesDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// 实例对应的Ecs规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 环境变量
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// 实例创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 实例修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 定时关机任务
	InstanceShutdownTimer *ListInstancesResponseBodyInstancesInstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty" type:"Struct"`
	// 实例Url
	InstanceUrl *string `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	// Jupyterlab Url
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// 最新保存的用户镜像
	LatestSnapshot *ListInstancesResponseBodyInstancesLatestSnapshot `json:"LatestSnapshot,omitempty" xml:"LatestSnapshot,omitempty" type:"Struct"`
	// 支付类型
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 实例错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例错误原因
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 实例状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 终端url
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	// 用户Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户名称
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// user vpc配置
	UserVpc *ListInstancesResponseBodyInstancesUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// Web IDE url
	WebIDEUrl *string `json:"WebIDEUrl,omitempty" xml:"WebIDEUrl,omitempty"`
	// 工作空间Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetAcceleratorType(v string) *ListInstancesResponseBodyInstances {
	s.AcceleratorType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetAccessibility(v string) *ListInstancesResponseBodyInstances {
	s.Accessibility = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetAccumulatedRunningTimeInMs(v int64) *ListInstancesResponseBodyInstances {
	s.AccumulatedRunningTimeInMs = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDatasets(v []*ListInstancesResponseBodyInstancesDatasets) *ListInstancesResponseBodyInstances {
	s.Datasets = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetEcsSpec(v string) *ListInstancesResponseBodyInstances {
	s.EcsSpec = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetEnvironmentVariables(v map[string]*string) *ListInstancesResponseBodyInstances {
	s.EnvironmentVariables = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageId(v string) *ListInstancesResponseBodyInstances {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageName(v string) *ListInstancesResponseBodyInstances {
	s.ImageName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageUrl(v string) *ListInstancesResponseBodyInstances {
	s.ImageUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceShutdownTimer(v *ListInstancesResponseBodyInstancesInstanceShutdownTimer) *ListInstancesResponseBodyInstances {
	s.InstanceShutdownTimer = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceUrl(v string) *ListInstancesResponseBodyInstances {
	s.InstanceUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetJupyterlabUrl(v string) *ListInstancesResponseBodyInstances {
	s.JupyterlabUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetLatestSnapshot(v *ListInstancesResponseBodyInstancesLatestSnapshot) *ListInstancesResponseBodyInstances {
	s.LatestSnapshot = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPaymentType(v string) *ListInstancesResponseBodyInstances {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetReasonCode(v string) *ListInstancesResponseBodyInstances {
	s.ReasonCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetReasonMessage(v string) *ListInstancesResponseBodyInstances {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetTerminalUrl(v string) *ListInstancesResponseBodyInstances {
	s.TerminalUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserId(v string) *ListInstancesResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserName(v string) *ListInstancesResponseBodyInstances {
	s.UserName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserVpc(v *ListInstancesResponseBodyInstancesUserVpc) *ListInstancesResponseBodyInstances {
	s.UserVpc = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetWebIDEUrl(v string) *ListInstancesResponseBodyInstances {
	s.WebIDEUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetWorkspaceId(v string) *ListInstancesResponseBodyInstances {
	s.WorkspaceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetWorkspaceName(v string) *ListInstancesResponseBodyInstances {
	s.WorkspaceName = &v
	return s
}

type ListInstancesResponseBodyInstancesDatasets struct {
	// 数据集Id
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 容器内挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s ListInstancesResponseBodyInstancesDatasets) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesDatasets) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetDatasetId(v string) *ListInstancesResponseBodyInstancesDatasets {
	s.DatasetId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetMountPath(v string) *ListInstancesResponseBodyInstancesDatasets {
	s.MountPath = &v
	return s
}

type ListInstancesResponseBodyInstancesInstanceShutdownTimer struct {
	// 设定关机时间
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 剩余关机时间（ms）
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s ListInstancesResponseBodyInstancesInstanceShutdownTimer) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesInstanceShutdownTimer) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetDueTime(v string) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.DueTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetInstanceId(v string) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetRemainingTimeInMs(v int64) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.RemainingTimeInMs = &v
	return s
}

type ListInstancesResponseBodyInstancesLatestSnapshot struct {
	// 快照创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 快照修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像Id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像名称
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// 镜像Url
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 镜像仓库Url
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
}

func (s ListInstancesResponseBodyInstancesLatestSnapshot) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesLatestSnapshot) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetImageId(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetImageName(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ImageName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetImageUrl(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ImageUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetRepositoryUrl(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.RepositoryUrl = &v
	return s
}

type ListInstancesResponseBodyInstancesUserVpc struct {
	// Security Group Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// VSwitch Id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListInstancesResponseBodyInstancesUserVpc) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesUserVpc) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetSecurityGroupId(v string) *ListInstancesResponseBodyInstancesUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetVSwitchId(v string) *ListInstancesResponseBodyInstancesUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetVpcId(v string) *ListInstancesResponseBodyInstancesUserVpc {
	s.VpcId = &v
	return s
}

type ListInstancesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type StartInstanceResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetCode(v string) *StartInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StartInstanceResponseBody) SetHttpStatusCode(v int32) *StartInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartInstanceResponseBody) SetInstanceId(v string) *StartInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceResponseBody) SetMessage(v string) *StartInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBody) SetSuccess(v bool) *StartInstanceResponseBody {
	s.Success = &v
	return s
}

type StartInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
	// 是否保存环境后再关闭实例
	SaveImage *bool `json:"SaveImage,omitempty" xml:"SaveImage,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetSaveImage(v bool) *StopInstanceRequest {
	s.SaveImage = &v
	return s
}

type StopInstanceResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetCode(v string) *StopInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StopInstanceResponseBody) SetHttpStatusCode(v int32) *StopInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopInstanceResponseBody) SetInstanceId(v string) *StopInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceResponseBody) SetMessage(v string) *StopInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) SetSuccess(v bool) *StopInstanceResponseBody {
	s.Success = &v
	return s
}

type StopInstanceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

type UpdateInstanceResponseBody struct {
	// 状态码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 说明
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetCode(v string) *UpdateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetInstanceId(v string) *UpdateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetMessage(v string) *UpdateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetSuccess(v bool) *UpdateInstanceResponseBody {
	s.Success = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("pai-dsw"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Datasets)) {
		body["Datasets"] = request.Datasets
	}

	if !tea.BoolValue(util.IsUnset(request.EcsSpec)) {
		body["EcsSpec"] = request.EcsSpec
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentVariables)) {
		body["EnvironmentVariables"] = request.EnvironmentVariables
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.UserVpc))) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances"),
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

func (client *Client) CreateInstanceShutdownTimer(InstanceId *string, request *CreateInstanceShutdownTimerRequest) (_result *CreateInstanceShutdownTimerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceShutdownTimerResponse{}
	_body, _err := client.CreateInstanceShutdownTimerWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceShutdownTimerWithOptions(InstanceId *string, request *CreateInstanceShutdownTimerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceShutdownTimerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["DueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.RemainingTimeInMs)) {
		body["RemainingTimeInMs"] = request.RemainingTimeInMs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceShutdownTimer"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/shutdowntimer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceShutdownTimerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceSnapshot(InstanceId *string, request *CreateInstanceSnapshotRequest) (_result *CreateInstanceSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceSnapshotResponse{}
	_body, _err := client.CreateInstanceSnapshotWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceSnapshotWithOptions(InstanceId *string, request *CreateInstanceSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotDescription)) {
		body["SnapshotDescription"] = request.SnapshotDescription
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		body["SnapshotName"] = request.SnapshotName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceSnapshot"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/snapshots"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(InstanceId *string) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId)),
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

func (client *Client) DeleteInstanceShutdownTimer(InstanceId *string) (_result *DeleteInstanceShutdownTimerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceShutdownTimerResponse{}
	_body, _err := client.DeleteInstanceShutdownTimerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceShutdownTimerWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceShutdownTimerResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceShutdownTimer"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/shutdowntimer"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceShutdownTimerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstanceSnapshot(InstanceId *string, SnapshotId *string) (_result *DeleteInstanceSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceSnapshotResponse{}
	_body, _err := client.DeleteInstanceSnapshotWithOptions(InstanceId, SnapshotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceSnapshotWithOptions(InstanceId *string, SnapshotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceSnapshotResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	SnapshotId = openapiutil.GetEncodeParam(SnapshotId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceSnapshot"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/snapshots/" + tea.StringValue(SnapshotId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(InstanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceShutdownTimer(InstanceId *string) (_result *GetInstanceShutdownTimerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceShutdownTimerResponse{}
	_body, _err := client.GetInstanceShutdownTimerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceShutdownTimerWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceShutdownTimerResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceShutdownTimer"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/shutdowntimer"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceShutdownTimerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceSnapshot(InstanceId *string, SnapshotId *string) (_result *GetInstanceSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceSnapshotResponse{}
	_body, _err := client.GetInstanceSnapshotWithOptions(InstanceId, SnapshotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceSnapshotWithOptions(InstanceId *string, SnapshotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceSnapshotResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	SnapshotId = openapiutil.GetEncodeParam(SnapshotId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceSnapshot"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/snapshots/" + tea.StringValue(SnapshotId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserConfig() (_result *GetUserConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserConfigResponse{}
	_body, _err := client.GetUserConfigWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserConfigWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserConfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/userconfig"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDemoCategories() (_result *ListDemoCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDemoCategoriesResponse{}
	_body, _err := client.ListDemoCategoriesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDemoCategoriesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDemoCategoriesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListDemoCategories"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/democategories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDemoCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDemos(request *ListDemosRequest) (_result *ListDemosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDemosResponse{}
	_body, _err := client.ListDemosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDemosWithOptions(request *ListDemosRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDemosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DemoName)) {
		query["DemoName"] = request.DemoName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDemos"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/demos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDemosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEcsSpecs(request *ListEcsSpecsRequest) (_result *ListEcsSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.ListEcsSpecsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEcsSpecsWithOptions(request *ListEcsSpecsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEcsSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEcsSpecs"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/ecsspecs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceSnapshot(InstanceId *string, request *ListInstanceSnapshotRequest) (_result *ListInstanceSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceSnapshotResponse{}
	_body, _err := client.ListInstanceSnapshotWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceSnapshotWithOptions(InstanceId *string, request *ListInstanceSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceSnapshot"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/snapshots"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceStatistics(request *ListInstanceStatisticsRequest) (_result *ListInstanceStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceStatisticsResponse{}
	_body, _err := client.ListInstanceStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceStatisticsWithOptions(request *ListInstanceStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceStatistics"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instancestatistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(InstanceId *string) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(InstanceId *string, request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(InstanceId *string, request *StopInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	InstanceId = openapiutil.GetEncodeParam(InstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SaveImage)) {
		query["SaveImage"] = request.SaveImage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstanceResponse{}
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(InstanceId)),
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
