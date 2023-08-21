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

type DemoCategory struct {
	CategoryCode  *string         `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	CategoryName  *string         `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	Order         *int64          `json:"Order,omitempty" xml:"Order,omitempty"`
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

type CreateIdleInstanceCullerRequest struct {
	CpuPercentThreshold  *int32 `json:"CpuPercentThreshold,omitempty" xml:"CpuPercentThreshold,omitempty"`
	GpuPercentThreshold  *int32 `json:"GpuPercentThreshold,omitempty" xml:"GpuPercentThreshold,omitempty"`
	MaxIdleTimeInMinutes *int32 `json:"MaxIdleTimeInMinutes,omitempty" xml:"MaxIdleTimeInMinutes,omitempty"`
}

func (s CreateIdleInstanceCullerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIdleInstanceCullerRequest) GoString() string {
	return s.String()
}

func (s *CreateIdleInstanceCullerRequest) SetCpuPercentThreshold(v int32) *CreateIdleInstanceCullerRequest {
	s.CpuPercentThreshold = &v
	return s
}

func (s *CreateIdleInstanceCullerRequest) SetGpuPercentThreshold(v int32) *CreateIdleInstanceCullerRequest {
	s.GpuPercentThreshold = &v
	return s
}

func (s *CreateIdleInstanceCullerRequest) SetMaxIdleTimeInMinutes(v int32) *CreateIdleInstanceCullerRequest {
	s.MaxIdleTimeInMinutes = &v
	return s
}

type CreateIdleInstanceCullerResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateIdleInstanceCullerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIdleInstanceCullerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdleInstanceCullerResponseBody) SetCode(v string) *CreateIdleInstanceCullerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIdleInstanceCullerResponseBody) SetInstanceId(v string) *CreateIdleInstanceCullerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateIdleInstanceCullerResponseBody) SetMessage(v string) *CreateIdleInstanceCullerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIdleInstanceCullerResponseBody) SetRequestId(v string) *CreateIdleInstanceCullerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdleInstanceCullerResponseBody) SetSuccess(v bool) *CreateIdleInstanceCullerResponseBody {
	s.Success = &v
	return s
}

type CreateIdleInstanceCullerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIdleInstanceCullerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIdleInstanceCullerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIdleInstanceCullerResponse) GoString() string {
	return s.String()
}

func (s *CreateIdleInstanceCullerResponse) SetHeaders(v map[string]*string) *CreateIdleInstanceCullerResponse {
	s.Headers = v
	return s
}

func (s *CreateIdleInstanceCullerResponse) SetStatusCode(v int32) *CreateIdleInstanceCullerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdleInstanceCullerResponse) SetBody(v *CreateIdleInstanceCullerResponseBody) *CreateIdleInstanceCullerResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	Accessibility        *string                                 `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CloudDisks           []*CreateInstanceRequestCloudDisks      `json:"CloudDisks,omitempty" xml:"CloudDisks,omitempty" type:"Repeated"`
	Datasets             []*CreateInstanceRequestDatasets        `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	EcsSpec              *string                                 `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	EnvironmentVariables map[string]*string                      `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	ImageId              *string                                 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUrl             *string                                 `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	InstanceName         *string                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Labels               []*CreateInstanceRequestLabels          `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Priority             *int64                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RequestedResource    *CreateInstanceRequestRequestedResource `json:"RequestedResource,omitempty" xml:"RequestedResource,omitempty" type:"Struct"`
	ResourceId           *string                                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	UserId               *string                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserVpc              *CreateInstanceRequestUserVpc           `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	WorkspaceId          *string                                 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceSource      *string                                 `json:"WorkspaceSource,omitempty" xml:"WorkspaceSource,omitempty"`
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

func (s *CreateInstanceRequest) SetCloudDisks(v []*CreateInstanceRequestCloudDisks) *CreateInstanceRequest {
	s.CloudDisks = v
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

func (s *CreateInstanceRequest) SetLabels(v []*CreateInstanceRequestLabels) *CreateInstanceRequest {
	s.Labels = v
	return s
}

func (s *CreateInstanceRequest) SetPriority(v int64) *CreateInstanceRequest {
	s.Priority = &v
	return s
}

func (s *CreateInstanceRequest) SetRequestedResource(v *CreateInstanceRequestRequestedResource) *CreateInstanceRequest {
	s.RequestedResource = v
	return s
}

func (s *CreateInstanceRequest) SetResourceId(v string) *CreateInstanceRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateInstanceRequest) SetUserId(v string) *CreateInstanceRequest {
	s.UserId = &v
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

func (s *CreateInstanceRequest) SetWorkspaceSource(v string) *CreateInstanceRequest {
	s.WorkspaceSource = &v
	return s
}

type CreateInstanceRequestCloudDisks struct {
	Capacity  *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SubType   *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s CreateInstanceRequestCloudDisks) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestCloudDisks) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCloudDisks) SetCapacity(v string) *CreateInstanceRequestCloudDisks {
	s.Capacity = &v
	return s
}

func (s *CreateInstanceRequestCloudDisks) SetMountPath(v string) *CreateInstanceRequestCloudDisks {
	s.MountPath = &v
	return s
}

func (s *CreateInstanceRequestCloudDisks) SetPath(v string) *CreateInstanceRequestCloudDisks {
	s.Path = &v
	return s
}

func (s *CreateInstanceRequestCloudDisks) SetSubType(v string) *CreateInstanceRequestCloudDisks {
	s.SubType = &v
	return s
}

type CreateInstanceRequestDatasets struct {
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
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

type CreateInstanceRequestLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestLabels) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestLabels) SetKey(v string) *CreateInstanceRequestLabels {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestLabels) SetValue(v string) *CreateInstanceRequestLabels {
	s.Value = &v
	return s
}

type CreateInstanceRequestRequestedResource struct {
	CPU          *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU          *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType      *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s CreateInstanceRequestRequestedResource) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestRequestedResource) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestRequestedResource) SetCPU(v string) *CreateInstanceRequestRequestedResource {
	s.CPU = &v
	return s
}

func (s *CreateInstanceRequestRequestedResource) SetGPU(v string) *CreateInstanceRequestRequestedResource {
	s.GPU = &v
	return s
}

func (s *CreateInstanceRequestRequestedResource) SetGPUType(v string) *CreateInstanceRequestRequestedResource {
	s.GPUType = &v
	return s
}

func (s *CreateInstanceRequestRequestedResource) SetMemory(v string) *CreateInstanceRequestRequestedResource {
	s.Memory = &v
	return s
}

func (s *CreateInstanceRequestRequestedResource) SetSharedMemory(v string) *CreateInstanceRequestRequestedResource {
	s.SharedMemory = &v
	return s
}

type CreateInstanceRequestUserVpc struct {
	DefaultRoute    *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs   []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId       *string   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateInstanceRequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestUserVpc) SetDefaultRoute(v string) *CreateInstanceRequestUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *CreateInstanceRequestUserVpc) SetExtendedCIDRs(v []*string) *CreateInstanceRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
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
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateInstanceShutdownTimerRequest struct {
	DueTime           *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	RemainingTimeInMs *int64  `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
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
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceShutdownTimerResponse) SetStatusCode(v int32) *CreateInstanceShutdownTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponse) SetBody(v *CreateInstanceShutdownTimerResponseBody) *CreateInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

type CreateInstanceSnapshotRequest struct {
	ImageUrl            *string                                `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Labels              []*CreateInstanceSnapshotRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	SnapshotDescription *string                                `json:"SnapshotDescription,omitempty" xml:"SnapshotDescription,omitempty"`
	SnapshotName        *string                                `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
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

func (s *CreateInstanceSnapshotRequest) SetLabels(v []*CreateInstanceSnapshotRequestLabels) *CreateInstanceSnapshotRequest {
	s.Labels = v
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

type CreateInstanceSnapshotRequestLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceSnapshotRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceSnapshotRequestLabels) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotRequestLabels) SetKey(v string) *CreateInstanceSnapshotRequestLabels {
	s.Key = &v
	return s
}

func (s *CreateInstanceSnapshotRequestLabels) SetValue(v string) *CreateInstanceSnapshotRequestLabels {
	s.Value = &v
	return s
}

type CreateInstanceSnapshotResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceSnapshotResponse) SetStatusCode(v int32) *CreateInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceSnapshotResponse) SetBody(v *CreateInstanceSnapshotResponseBody) *CreateInstanceSnapshotResponse {
	s.Body = v
	return s
}

type DeleteIdleInstanceCullerResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteIdleInstanceCullerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIdleInstanceCullerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdleInstanceCullerResponseBody) SetCode(v string) *DeleteIdleInstanceCullerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponseBody) SetInstanceId(v string) *DeleteIdleInstanceCullerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponseBody) SetMessage(v string) *DeleteIdleInstanceCullerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponseBody) SetRequestId(v string) *DeleteIdleInstanceCullerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponseBody) SetSuccess(v bool) *DeleteIdleInstanceCullerResponseBody {
	s.Success = &v
	return s
}

type DeleteIdleInstanceCullerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIdleInstanceCullerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIdleInstanceCullerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIdleInstanceCullerResponse) GoString() string {
	return s.String()
}

func (s *DeleteIdleInstanceCullerResponse) SetHeaders(v map[string]*string) *DeleteIdleInstanceCullerResponse {
	s.Headers = v
	return s
}

func (s *DeleteIdleInstanceCullerResponse) SetStatusCode(v int32) *DeleteIdleInstanceCullerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponse) SetBody(v *DeleteIdleInstanceCullerResponseBody) *DeleteIdleInstanceCullerResponse {
	s.Body = v
	return s
}

type DeleteInstanceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteInstanceShutdownTimerResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteInstanceShutdownTimerResponse) SetStatusCode(v int32) *DeleteInstanceShutdownTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponse) SetBody(v *DeleteInstanceShutdownTimerResponseBody) *DeleteInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

type DeleteInstanceSnapshotResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteInstanceSnapshotResponse) SetStatusCode(v int32) *DeleteInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceSnapshotResponse) SetBody(v *DeleteInstanceSnapshotResponseBody) *DeleteInstanceSnapshotResponse {
	s.Body = v
	return s
}

type GetIdleInstanceCullerResponseBody struct {
	Code                 *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CpuPercentThreshold  *int32  `json:"CpuPercentThreshold,omitempty" xml:"CpuPercentThreshold,omitempty"`
	GpuPercentThreshold  *int32  `json:"GpuPercentThreshold,omitempty" xml:"GpuPercentThreshold,omitempty"`
	IdleTimeInMinutes    *int32  `json:"IdleTimeInMinutes,omitempty" xml:"IdleTimeInMinutes,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxIdleTimeInMinutes *int32  `json:"MaxIdleTimeInMinutes,omitempty" xml:"MaxIdleTimeInMinutes,omitempty"`
	Message              *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success              *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIdleInstanceCullerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIdleInstanceCullerResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdleInstanceCullerResponseBody) SetCode(v string) *GetIdleInstanceCullerResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetCpuPercentThreshold(v int32) *GetIdleInstanceCullerResponseBody {
	s.CpuPercentThreshold = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetGpuPercentThreshold(v int32) *GetIdleInstanceCullerResponseBody {
	s.GpuPercentThreshold = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetIdleTimeInMinutes(v int32) *GetIdleInstanceCullerResponseBody {
	s.IdleTimeInMinutes = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetInstanceId(v string) *GetIdleInstanceCullerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetMaxIdleTimeInMinutes(v int32) *GetIdleInstanceCullerResponseBody {
	s.MaxIdleTimeInMinutes = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetMessage(v string) *GetIdleInstanceCullerResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetRequestId(v string) *GetIdleInstanceCullerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdleInstanceCullerResponseBody) SetSuccess(v bool) *GetIdleInstanceCullerResponseBody {
	s.Success = &v
	return s
}

type GetIdleInstanceCullerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIdleInstanceCullerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIdleInstanceCullerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIdleInstanceCullerResponse) GoString() string {
	return s.String()
}

func (s *GetIdleInstanceCullerResponse) SetHeaders(v map[string]*string) *GetIdleInstanceCullerResponse {
	s.Headers = v
	return s
}

func (s *GetIdleInstanceCullerResponse) SetStatusCode(v int32) *GetIdleInstanceCullerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdleInstanceCullerResponse) SetBody(v *GetIdleInstanceCullerResponseBody) *GetIdleInstanceCullerResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	AcceleratorType            *string                                        `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	Accessibility              *string                                        `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	AccumulatedRunningTimeInMs *int64                                         `json:"AccumulatedRunningTimeInMs,omitempty" xml:"AccumulatedRunningTimeInMs,omitempty"`
	CloudDisks                 []*GetInstanceResponseBodyCloudDisks           `json:"CloudDisks,omitempty" xml:"CloudDisks,omitempty" type:"Repeated"`
	Code                       *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Datasets                   []*GetInstanceResponseBodyDatasets             `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	EcsSpec                    *string                                        `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	EnvironmentVariables       map[string]*string                             `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	GmtCreateTime              *string                                        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime            *string                                        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HttpStatusCode             *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	IdleInstanceCuller         *GetInstanceResponseBodyIdleInstanceCuller     `json:"IdleInstanceCuller,omitempty" xml:"IdleInstanceCuller,omitempty" type:"Struct"`
	ImageId                    *string                                        `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                  *string                                        `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageUrl                   *string                                        `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	InstanceId                 *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName               *string                                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceShutdownTimer      *GetInstanceResponseBodyInstanceShutdownTimer  `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty" type:"Struct"`
	InstanceSnapshotList       []*GetInstanceResponseBodyInstanceSnapshotList `json:"InstanceSnapshotList,omitempty" xml:"InstanceSnapshotList,omitempty" type:"Repeated"`
	InstanceUrl                *string                                        `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	// Jupyterlab Url。
	JupyterlabUrl     *string                                   `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	Labels            []*GetInstanceResponseBodyLabels          `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestSnapshot    *GetInstanceResponseBodyLatestSnapshot    `json:"LatestSnapshot,omitempty" xml:"LatestSnapshot,omitempty" type:"Struct"`
	Message           *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PaymentType       *string                                   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	Priority          *int64                                    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReasonCode        *string                                   `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage     *string                                   `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RequestId         *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestedResource *GetInstanceResponseBodyRequestedResource `json:"RequestedResource,omitempty" xml:"RequestedResource,omitempty" type:"Struct"`
	ResourceId        *string                                   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName      *string                                   `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Status            *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Success           *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TerminalUrl       *string                                   `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	UserId            *string                                   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName          *string                                   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserVpc           *GetInstanceResponseBodyUserVpc           `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// Web IDE url。
	WebIDEUrl       *string `json:"WebIDEUrl,omitempty" xml:"WebIDEUrl,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName   *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	WorkspaceSource *string `json:"WorkspaceSource,omitempty" xml:"WorkspaceSource,omitempty"`
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

func (s *GetInstanceResponseBody) SetCloudDisks(v []*GetInstanceResponseBodyCloudDisks) *GetInstanceResponseBody {
	s.CloudDisks = v
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

func (s *GetInstanceResponseBody) SetIdleInstanceCuller(v *GetInstanceResponseBodyIdleInstanceCuller) *GetInstanceResponseBody {
	s.IdleInstanceCuller = v
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

func (s *GetInstanceResponseBody) SetInstanceSnapshotList(v []*GetInstanceResponseBodyInstanceSnapshotList) *GetInstanceResponseBody {
	s.InstanceSnapshotList = v
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

func (s *GetInstanceResponseBody) SetLabels(v []*GetInstanceResponseBodyLabels) *GetInstanceResponseBody {
	s.Labels = v
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

func (s *GetInstanceResponseBody) SetPriority(v int64) *GetInstanceResponseBody {
	s.Priority = &v
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

func (s *GetInstanceResponseBody) SetRequestedResource(v *GetInstanceResponseBodyRequestedResource) *GetInstanceResponseBody {
	s.RequestedResource = v
	return s
}

func (s *GetInstanceResponseBody) SetResourceId(v string) *GetInstanceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResourceName(v string) *GetInstanceResponseBody {
	s.ResourceName = &v
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

func (s *GetInstanceResponseBody) SetWorkspaceSource(v string) *GetInstanceResponseBody {
	s.WorkspaceSource = &v
	return s
}

type GetInstanceResponseBodyCloudDisks struct {
	Capacity  *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SubType   *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s GetInstanceResponseBodyCloudDisks) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyCloudDisks) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyCloudDisks) SetCapacity(v string) *GetInstanceResponseBodyCloudDisks {
	s.Capacity = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisks) SetMountPath(v string) *GetInstanceResponseBodyCloudDisks {
	s.MountPath = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisks) SetPath(v string) *GetInstanceResponseBodyCloudDisks {
	s.Path = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisks) SetSubType(v string) *GetInstanceResponseBodyCloudDisks {
	s.SubType = &v
	return s
}

type GetInstanceResponseBodyDatasets struct {
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
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

type GetInstanceResponseBodyIdleInstanceCuller struct {
	CpuPercentThreshold  *int32  `json:"CpuPercentThreshold,omitempty" xml:"CpuPercentThreshold,omitempty"`
	GpuPercentThreshold  *int32  `json:"GpuPercentThreshold,omitempty" xml:"GpuPercentThreshold,omitempty"`
	IdleTimeInMinutes    *int32  `json:"IdleTimeInMinutes,omitempty" xml:"IdleTimeInMinutes,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxIdleTimeInMinutes *int32  `json:"MaxIdleTimeInMinutes,omitempty" xml:"MaxIdleTimeInMinutes,omitempty"`
}

func (s GetInstanceResponseBodyIdleInstanceCuller) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyIdleInstanceCuller) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetCpuPercentThreshold(v int32) *GetInstanceResponseBodyIdleInstanceCuller {
	s.CpuPercentThreshold = &v
	return s
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetGpuPercentThreshold(v int32) *GetInstanceResponseBodyIdleInstanceCuller {
	s.GpuPercentThreshold = &v
	return s
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetIdleTimeInMinutes(v int32) *GetInstanceResponseBodyIdleInstanceCuller {
	s.IdleTimeInMinutes = &v
	return s
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetInstanceId(v string) *GetInstanceResponseBodyIdleInstanceCuller {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetMaxIdleTimeInMinutes(v int32) *GetInstanceResponseBodyIdleInstanceCuller {
	s.MaxIdleTimeInMinutes = &v
	return s
}

type GetInstanceResponseBodyInstanceShutdownTimer struct {
	DueTime           *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	GmtCreateTime     *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RemainingTimeInMs *int64  `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
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

type GetInstanceResponseBodyInstanceSnapshotList struct {
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
	// 实例快照错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例快照错误消息
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 镜像仓库Url
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
	// 实例快照状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyInstanceSnapshotList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceSnapshotList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetGmtCreateTime(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetGmtModifiedTime(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetImageId(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ImageId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetImageName(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ImageName = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetImageUrl(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetReasonCode(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetReasonMessage(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetRepositoryUrl(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.RepositoryUrl = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetStatus(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.Status = &v
	return s
}

type GetInstanceResponseBodyLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyLabels) SetKey(v string) *GetInstanceResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyLabels) SetValue(v string) *GetInstanceResponseBodyLabels {
	s.Value = &v
	return s
}

type GetInstanceResponseBodyLatestSnapshot struct {
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName       *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 实例快照错误代码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 实例快照错误消息
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
	// 实例快照状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *GetInstanceResponseBodyLatestSnapshot) SetReasonCode(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetReasonMessage(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetRepositoryUrl(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.RepositoryUrl = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetStatus(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.Status = &v
	return s
}

type GetInstanceResponseBodyRequestedResource struct {
	CPU          *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU          *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType      *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s GetInstanceResponseBodyRequestedResource) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyRequestedResource) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyRequestedResource) SetCPU(v string) *GetInstanceResponseBodyRequestedResource {
	s.CPU = &v
	return s
}

func (s *GetInstanceResponseBodyRequestedResource) SetGPU(v string) *GetInstanceResponseBodyRequestedResource {
	s.GPU = &v
	return s
}

func (s *GetInstanceResponseBodyRequestedResource) SetGPUType(v string) *GetInstanceResponseBodyRequestedResource {
	s.GPUType = &v
	return s
}

func (s *GetInstanceResponseBodyRequestedResource) SetMemory(v string) *GetInstanceResponseBodyRequestedResource {
	s.Memory = &v
	return s
}

func (s *GetInstanceResponseBodyRequestedResource) SetSharedMemory(v string) *GetInstanceResponseBodyRequestedResource {
	s.SharedMemory = &v
	return s
}

type GetInstanceResponseBodyUserVpc struct {
	DefaultRoute    *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs   []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId       *string   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Vpc Id。
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetInstanceResponseBodyUserVpc) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyUserVpc) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyUserVpc) SetDefaultRoute(v string) *GetInstanceResponseBodyUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) SetExtendedCIDRs(v []*string) *GetInstanceResponseBodyUserVpc {
	s.ExtendedCIDRs = v
	return s
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetInstanceMetricsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeStep   *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s GetInstanceMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsRequest) SetEndTime(v string) *GetInstanceMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetMetricType(v string) *GetInstanceMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetStartTime(v string) *GetInstanceMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetTimeStep(v string) *GetInstanceMetricsRequest {
	s.TimeStep = &v
	return s
}

type GetInstanceMetricsResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PodMetrics     []*GetInstanceMetricsResponseBodyPodMetrics `json:"PodMetrics,omitempty" xml:"PodMetrics,omitempty" type:"Repeated"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBody) SetCode(v string) *GetInstanceMetricsResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetHttpStatusCode(v int32) *GetInstanceMetricsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetInstanceId(v string) *GetInstanceMetricsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetMessage(v string) *GetInstanceMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetPodMetrics(v []*GetInstanceMetricsResponseBodyPodMetrics) *GetInstanceMetricsResponseBody {
	s.PodMetrics = v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetRequestId(v string) *GetInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetSuccess(v bool) *GetInstanceMetricsResponseBody {
	s.Success = &v
	return s
}

type GetInstanceMetricsResponseBodyPodMetrics struct {
	Metrics []*GetInstanceMetricsResponseBodyPodMetricsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	PodId   *string                                            `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s GetInstanceMetricsResponseBodyPodMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsResponseBodyPodMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) SetMetrics(v []*GetInstanceMetricsResponseBodyPodMetricsMetrics) *GetInstanceMetricsResponseBodyPodMetrics {
	s.Metrics = v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) SetPodId(v string) *GetInstanceMetricsResponseBodyPodMetrics {
	s.PodId = &v
	return s
}

type GetInstanceMetricsResponseBodyPodMetricsMetrics struct {
	Time  *int64   `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceMetricsResponseBodyPodMetricsMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsResponseBodyPodMetricsMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) SetTime(v int64) *GetInstanceMetricsResponseBodyPodMetricsMetrics {
	s.Time = &v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) SetValue(v float32) *GetInstanceMetricsResponseBodyPodMetricsMetrics {
	s.Value = &v
	return s
}

type GetInstanceMetricsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponse) SetHeaders(v map[string]*string) *GetInstanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceMetricsResponse) SetStatusCode(v int32) *GetInstanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceMetricsResponse) SetBody(v *GetInstanceMetricsResponseBody) *GetInstanceMetricsResponse {
	s.Body = v
	return s
}

type GetInstanceShutdownTimerResponseBody struct {
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DueTime           *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	GmtCreateTime     *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HttpStatusCode    *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RemainingTimeInMs *int64  `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetInstanceShutdownTimerResponse) SetStatusCode(v int32) *GetInstanceShutdownTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceShutdownTimerResponse) SetBody(v *GetInstanceShutdownTimerResponseBody) *GetInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

type GetInstanceSnapshotResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HttpStatusCode  *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReasonCode      *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage   *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotId      *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName    *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetInstanceSnapshotResponse) SetStatusCode(v int32) *GetInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceSnapshotResponse) SetBody(v *GetInstanceSnapshotResponseBody) *GetInstanceSnapshotResponse {
	s.Body = v
	return s
}

type GetLifecycleRequest struct {
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit         *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	SessionNumber *int32  `json:"SessionNumber,omitempty" xml:"SessionNumber,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetLifecycleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLifecycleRequest) GoString() string {
	return s.String()
}

func (s *GetLifecycleRequest) SetEndTime(v string) *GetLifecycleRequest {
	s.EndTime = &v
	return s
}

func (s *GetLifecycleRequest) SetLimit(v int32) *GetLifecycleRequest {
	s.Limit = &v
	return s
}

func (s *GetLifecycleRequest) SetOrder(v string) *GetLifecycleRequest {
	s.Order = &v
	return s
}

func (s *GetLifecycleRequest) SetSessionNumber(v int32) *GetLifecycleRequest {
	s.SessionNumber = &v
	return s
}

func (s *GetLifecycleRequest) SetStartTime(v string) *GetLifecycleRequest {
	s.StartTime = &v
	return s
}

type GetLifecycleResponseBody struct {
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Lifecycle  [][]*GetLifecycleResponseBodyLifecycle `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty" type:"Repeated"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetLifecycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLifecycleResponseBody) GoString() string {
	return s.String()
}

func (s *GetLifecycleResponseBody) SetCode(v string) *GetLifecycleResponseBody {
	s.Code = &v
	return s
}

func (s *GetLifecycleResponseBody) SetLifecycle(v [][]*GetLifecycleResponseBodyLifecycle) *GetLifecycleResponseBody {
	s.Lifecycle = v
	return s
}

func (s *GetLifecycleResponseBody) SetMessage(v string) *GetLifecycleResponseBody {
	s.Message = &v
	return s
}

func (s *GetLifecycleResponseBody) SetRequestId(v string) *GetLifecycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLifecycleResponseBody) SetSuccess(v bool) *GetLifecycleResponseBody {
	s.Success = &v
	return s
}

func (s *GetLifecycleResponseBody) SetTotalCount(v int32) *GetLifecycleResponseBody {
	s.TotalCount = &v
	return s
}

type GetLifecycleResponseBodyLifecycle struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ReasonCode    *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
}

func (s GetLifecycleResponseBodyLifecycle) String() string {
	return tea.Prettify(s)
}

func (s GetLifecycleResponseBodyLifecycle) GoString() string {
	return s.String()
}

func (s *GetLifecycleResponseBodyLifecycle) SetStatus(v string) *GetLifecycleResponseBodyLifecycle {
	s.Status = &v
	return s
}

func (s *GetLifecycleResponseBodyLifecycle) SetReasonCode(v string) *GetLifecycleResponseBodyLifecycle {
	s.ReasonCode = &v
	return s
}

func (s *GetLifecycleResponseBodyLifecycle) SetReasonMessage(v string) *GetLifecycleResponseBodyLifecycle {
	s.ReasonMessage = &v
	return s
}

func (s *GetLifecycleResponseBodyLifecycle) SetGmtCreateTime(v string) *GetLifecycleResponseBodyLifecycle {
	s.GmtCreateTime = &v
	return s
}

type GetLifecycleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLifecycleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLifecycleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLifecycleResponse) GoString() string {
	return s.String()
}

func (s *GetLifecycleResponse) SetHeaders(v map[string]*string) *GetLifecycleResponse {
	s.Headers = v
	return s
}

func (s *GetLifecycleResponse) SetStatusCode(v int32) *GetLifecycleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLifecycleResponse) SetBody(v *GetLifecycleResponseBody) *GetLifecycleResponse {
	s.Body = v
	return s
}

type GetResourceGroupStatisticsRequest struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s GetResourceGroupStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupStatisticsRequest) SetEndTime(v string) *GetResourceGroupStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceGroupStatisticsRequest) SetResourceId(v string) *GetResourceGroupStatisticsRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceGroupStatisticsRequest) SetStartTime(v string) *GetResourceGroupStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *GetResourceGroupStatisticsRequest) SetWorkspaceIds(v string) *GetResourceGroupStatisticsRequest {
	s.WorkspaceIds = &v
	return s
}

type GetResourceGroupStatisticsResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics     map[string]map[string]interface{} `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResourceGroupStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupStatisticsResponseBody) SetCode(v string) *GetResourceGroupStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetHttpStatusCode(v int32) *GetResourceGroupStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetMessage(v string) *GetResourceGroupStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetRequestId(v string) *GetResourceGroupStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetStatistics(v map[string]map[string]interface{}) *GetResourceGroupStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *GetResourceGroupStatisticsResponseBody) SetSuccess(v bool) *GetResourceGroupStatisticsResponseBody {
	s.Success = &v
	return s
}

type GetResourceGroupStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceGroupStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceGroupStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupStatisticsResponse) SetHeaders(v map[string]*string) *GetResourceGroupStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupStatisticsResponse) SetStatusCode(v int32) *GetResourceGroupStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupStatisticsResponse) SetBody(v *GetResourceGroupStatisticsResponseBody) *GetResourceGroupStatisticsResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	ExpireTime *int32  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetExpireTime(v int32) *GetTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenRequest) SetInstanceId(v string) *GetTokenRequest {
	s.InstanceId = &v
	return s
}

type GetTokenResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetCode(v string) *GetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenResponseBody) SetMessage(v string) *GetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetSuccess(v bool) *GetTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetTokenResponseBody) SetToken(v string) *GetTokenResponseBody {
	s.Token = &v
	return s
}

type GetTokenResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type GetUserConfigResponseBody struct {
	AccountSufficient     *bool                              `json:"AccountSufficient,omitempty" xml:"AccountSufficient,omitempty"`
	Code                  *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	EnableEciDisk         *bool                              `json:"EnableEciDisk,omitempty" xml:"EnableEciDisk,omitempty"`
	FreeTier              *GetUserConfigResponseBodyFreeTier `json:"FreeTier,omitempty" xml:"FreeTier,omitempty" type:"Struct"`
	FreeTierSpecAvailable *bool                              `json:"FreeTierSpecAvailable,omitempty" xml:"FreeTierSpecAvailable,omitempty"`
	HttpStatusCode        *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message               *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId             *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success               *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *GetUserConfigResponseBody) SetEnableEciDisk(v bool) *GetUserConfigResponseBody {
	s.EnableEciDisk = &v
	return s
}

func (s *GetUserConfigResponseBody) SetFreeTier(v *GetUserConfigResponseBodyFreeTier) *GetUserConfigResponseBody {
	s.FreeTier = v
	return s
}

func (s *GetUserConfigResponseBody) SetFreeTierSpecAvailable(v bool) *GetUserConfigResponseBody {
	s.FreeTierSpecAvailable = &v
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

type GetUserConfigResponseBodyFreeTier struct {
	EndTime         *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InitBaseUnit    *string  `json:"InitBaseUnit,omitempty" xml:"InitBaseUnit,omitempty"`
	InitBaseValue   *float64 `json:"InitBaseValue,omitempty" xml:"InitBaseValue,omitempty"`
	InitShowUnit    *string  `json:"InitShowUnit,omitempty" xml:"InitShowUnit,omitempty"`
	InitShowValue   *string  `json:"InitShowValue,omitempty" xml:"InitShowValue,omitempty"`
	IsFreeTierUser  *bool    `json:"IsFreeTierUser,omitempty" xml:"IsFreeTierUser,omitempty"`
	PeriodBaseUnit  *string  `json:"PeriodBaseUnit,omitempty" xml:"PeriodBaseUnit,omitempty"`
	PeriodBaseValue *float64 `json:"PeriodBaseValue,omitempty" xml:"PeriodBaseValue,omitempty"`
	PeriodShowUnit  *string  `json:"PeriodShowUnit,omitempty" xml:"PeriodShowUnit,omitempty"`
	PeriodShowValue *string  `json:"PeriodShowValue,omitempty" xml:"PeriodShowValue,omitempty"`
	StartTime       *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUserConfigResponseBodyFreeTier) String() string {
	return tea.Prettify(s)
}

func (s GetUserConfigResponseBodyFreeTier) GoString() string {
	return s.String()
}

func (s *GetUserConfigResponseBodyFreeTier) SetEndTime(v string) *GetUserConfigResponseBodyFreeTier {
	s.EndTime = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetInitBaseUnit(v string) *GetUserConfigResponseBodyFreeTier {
	s.InitBaseUnit = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetInitBaseValue(v float64) *GetUserConfigResponseBodyFreeTier {
	s.InitBaseValue = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetInitShowUnit(v string) *GetUserConfigResponseBodyFreeTier {
	s.InitShowUnit = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetInitShowValue(v string) *GetUserConfigResponseBodyFreeTier {
	s.InitShowValue = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetIsFreeTierUser(v bool) *GetUserConfigResponseBodyFreeTier {
	s.IsFreeTierUser = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetPeriodBaseUnit(v string) *GetUserConfigResponseBodyFreeTier {
	s.PeriodBaseUnit = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetPeriodBaseValue(v float64) *GetUserConfigResponseBodyFreeTier {
	s.PeriodBaseValue = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetPeriodShowUnit(v string) *GetUserConfigResponseBodyFreeTier {
	s.PeriodShowUnit = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetPeriodShowValue(v string) *GetUserConfigResponseBodyFreeTier {
	s.PeriodShowValue = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetStartTime(v string) *GetUserConfigResponseBodyFreeTier {
	s.StartTime = &v
	return s
}

func (s *GetUserConfigResponseBodyFreeTier) SetStatus(v string) *GetUserConfigResponseBodyFreeTier {
	s.Status = &v
	return s
}

type GetUserConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetUserConfigResponse) SetStatusCode(v int32) *GetUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserConfigResponse) SetBody(v *GetUserConfigResponseBody) *GetUserConfigResponse {
	s.Body = v
	return s
}

type ListEcsSpecsRequest struct {
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	Order           *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber      *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy          *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	EcsSpecs       []*ListEcsSpecsResponseBodyEcsSpecs `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AcceleratorType     *string                                   `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	CPU                 *int64                                    `json:"CPU,omitempty" xml:"CPU,omitempty"`
	Currency            *string                                   `json:"Currency,omitempty" xml:"Currency,omitempty"`
	GPU                 *int64                                    `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType             *string                                   `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	InstanceBandwidthRx *int64                                    `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	InstanceType        *string                                   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IsAvailable         *bool                                     `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	Labels              []*ListEcsSpecsResponseBodyEcsSpecsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Memory              *float32                                  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Price               *float64                                  `json:"Price,omitempty" xml:"Price,omitempty"`
	SystemDiskCapacity  *int64                                    `json:"SystemDiskCapacity,omitempty" xml:"SystemDiskCapacity,omitempty"`
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

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetIsAvailable(v bool) *ListEcsSpecsResponseBodyEcsSpecs {
	s.IsAvailable = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecs) SetLabels(v []*ListEcsSpecsResponseBodyEcsSpecsLabels) *ListEcsSpecsResponseBodyEcsSpecs {
	s.Labels = v
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

type ListEcsSpecsResponseBodyEcsSpecsLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEcsSpecsResponseBodyEcsSpecsLabels) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponseBodyEcsSpecsLabels) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBodyEcsSpecsLabels) SetKey(v string) *ListEcsSpecsResponseBodyEcsSpecsLabels {
	s.Key = &v
	return s
}

func (s *ListEcsSpecsResponseBodyEcsSpecsLabels) SetValue(v string) *ListEcsSpecsResponseBodyEcsSpecsLabels {
	s.Value = &v
	return s
}

type ListEcsSpecsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEcsSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListEcsSpecsResponse) SetStatusCode(v int32) *ListEcsSpecsResponse {
	s.StatusCode = &v
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
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots      []*ListInstanceSnapshotResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	GmtCreateTime   *string                                            `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                                            `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ImageId         *string                                            `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUrl        *string                                            `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	InstanceId      *string                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Labels          []*ListInstanceSnapshotResponseBodySnapshotsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	ReasonCode      *string                                            `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage   *string                                            `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	SnapshotId      *string                                            `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName    *string                                            `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	Status          *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListInstanceSnapshotResponseBodySnapshots) SetLabels(v []*ListInstanceSnapshotResponseBodySnapshotsLabels) *ListInstanceSnapshotResponseBodySnapshots {
	s.Labels = v
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

type ListInstanceSnapshotResponseBodySnapshotsLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstanceSnapshotResponseBodySnapshotsLabels) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceSnapshotResponseBodySnapshotsLabels) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotResponseBodySnapshotsLabels) SetKey(v string) *ListInstanceSnapshotResponseBodySnapshotsLabels {
	s.Key = &v
	return s
}

func (s *ListInstanceSnapshotResponseBodySnapshotsLabels) SetValue(v string) *ListInstanceSnapshotResponseBodySnapshotsLabels {
	s.Value = &v
	return s
}

type ListInstanceSnapshotResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstanceSnapshotResponse) SetStatusCode(v int32) *ListInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceSnapshotResponse) SetBody(v *ListInstanceSnapshotResponseBody) *ListInstanceSnapshotResponse {
	s.Body = v
	return s
}

type ListInstanceStatisticsRequest struct {
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
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics     map[string]map[string]interface{} `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstanceStatisticsResponse) SetStatusCode(v int32) *ListInstanceStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceStatisticsResponse) SetBody(v *ListInstanceStatisticsResponseBody) *ListInstanceStatisticsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Order           *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber      *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PaymentType     *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	SortBy          *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *ListInstancesRequest) SetInstanceId(v string) *ListInstancesRequest {
	s.InstanceId = &v
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

func (s *ListInstancesRequest) SetResourceId(v string) *ListInstancesRequest {
	s.ResourceId = &v
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
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Instances      []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AcceleratorType            *string                                                   `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	Accessibility              *string                                                   `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	AccumulatedRunningTimeInMs *int64                                                    `json:"AccumulatedRunningTimeInMs,omitempty" xml:"AccumulatedRunningTimeInMs,omitempty"`
	CloudDisks                 []*ListInstancesResponseBodyInstancesCloudDisks           `json:"CloudDisks,omitempty" xml:"CloudDisks,omitempty" type:"Repeated"`
	Datasets                   []*ListInstancesResponseBodyInstancesDatasets             `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	EcsSpec                    *string                                                   `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	EnvironmentVariables       map[string]*string                                        `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	GmtCreateTime              *string                                                   `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime            *string                                                   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IdleInstanceCuller         *ListInstancesResponseBodyInstancesIdleInstanceCuller     `json:"IdleInstanceCuller,omitempty" xml:"IdleInstanceCuller,omitempty" type:"Struct"`
	ImageId                    *string                                                   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName                  *string                                                   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageUrl                   *string                                                   `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	InstanceId                 *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName               *string                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceShutdownTimer      *ListInstancesResponseBodyInstancesInstanceShutdownTimer  `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty" type:"Struct"`
	InstanceSnapshotList       []*ListInstancesResponseBodyInstancesInstanceSnapshotList `json:"InstanceSnapshotList,omitempty" xml:"InstanceSnapshotList,omitempty" type:"Repeated"`
	InstanceUrl                *string                                                   `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	// Jupyterlab Url。
	JupyterlabUrl     *string                                              `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	Labels            []*ListInstancesResponseBodyInstancesLabels          `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestSnapshot    *ListInstancesResponseBodyInstancesLatestSnapshot    `json:"LatestSnapshot,omitempty" xml:"LatestSnapshot,omitempty" type:"Struct"`
	PaymentType       *string                                              `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	Priority          *int64                                               `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReasonCode        *string                                              `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage     *string                                              `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RequestedResource *ListInstancesResponseBodyInstancesRequestedResource `json:"RequestedResource,omitempty" xml:"RequestedResource,omitempty" type:"Struct"`
	ResourceId        *string                                              `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName      *string                                              `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Status            *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	TerminalUrl       *string                                              `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	UserId            *string                                              `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName          *string                                              `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserVpc           *ListInstancesResponseBodyInstancesUserVpc           `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// Web IDE url。
	WebIDEUrl       *string `json:"WebIDEUrl,omitempty" xml:"WebIDEUrl,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName   *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	WorkspaceSource *string `json:"WorkspaceSource,omitempty" xml:"WorkspaceSource,omitempty"`
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

func (s *ListInstancesResponseBodyInstances) SetCloudDisks(v []*ListInstancesResponseBodyInstancesCloudDisks) *ListInstancesResponseBodyInstances {
	s.CloudDisks = v
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

func (s *ListInstancesResponseBodyInstances) SetIdleInstanceCuller(v *ListInstancesResponseBodyInstancesIdleInstanceCuller) *ListInstancesResponseBodyInstances {
	s.IdleInstanceCuller = v
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

func (s *ListInstancesResponseBodyInstances) SetInstanceSnapshotList(v []*ListInstancesResponseBodyInstancesInstanceSnapshotList) *ListInstancesResponseBodyInstances {
	s.InstanceSnapshotList = v
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

func (s *ListInstancesResponseBodyInstances) SetLabels(v []*ListInstancesResponseBodyInstancesLabels) *ListInstancesResponseBodyInstances {
	s.Labels = v
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

func (s *ListInstancesResponseBodyInstances) SetPriority(v int64) *ListInstancesResponseBodyInstances {
	s.Priority = &v
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

func (s *ListInstancesResponseBodyInstances) SetRequestedResource(v *ListInstancesResponseBodyInstancesRequestedResource) *ListInstancesResponseBodyInstances {
	s.RequestedResource = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceId(v string) *ListInstancesResponseBodyInstances {
	s.ResourceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceName(v string) *ListInstancesResponseBodyInstances {
	s.ResourceName = &v
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

func (s *ListInstancesResponseBodyInstances) SetWorkspaceSource(v string) *ListInstancesResponseBodyInstances {
	s.WorkspaceSource = &v
	return s
}

type ListInstancesResponseBodyInstancesCloudDisks struct {
	Capacity  *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SubType   *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s ListInstancesResponseBodyInstancesCloudDisks) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesCloudDisks) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) SetCapacity(v string) *ListInstancesResponseBodyInstancesCloudDisks {
	s.Capacity = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) SetMountPath(v string) *ListInstancesResponseBodyInstancesCloudDisks {
	s.MountPath = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) SetPath(v string) *ListInstancesResponseBodyInstancesCloudDisks {
	s.Path = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) SetSubType(v string) *ListInstancesResponseBodyInstancesCloudDisks {
	s.SubType = &v
	return s
}

type ListInstancesResponseBodyInstancesDatasets struct {
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
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

type ListInstancesResponseBodyInstancesIdleInstanceCuller struct {
	CpuPercentThreshold  *int32  `json:"CpuPercentThreshold,omitempty" xml:"CpuPercentThreshold,omitempty"`
	GpuPercentThreshold  *int32  `json:"GpuPercentThreshold,omitempty" xml:"GpuPercentThreshold,omitempty"`
	IdleTimeInMinutes    *int32  `json:"IdleTimeInMinutes,omitempty" xml:"IdleTimeInMinutes,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxIdleTimeInMinutes *int32  `json:"MaxIdleTimeInMinutes,omitempty" xml:"MaxIdleTimeInMinutes,omitempty"`
}

func (s ListInstancesResponseBodyInstancesIdleInstanceCuller) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesIdleInstanceCuller) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetCpuPercentThreshold(v int32) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.CpuPercentThreshold = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetGpuPercentThreshold(v int32) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.GpuPercentThreshold = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetIdleTimeInMinutes(v int32) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.IdleTimeInMinutes = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetInstanceId(v string) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetMaxIdleTimeInMinutes(v int32) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.MaxIdleTimeInMinutes = &v
	return s
}

type ListInstancesResponseBodyInstancesInstanceShutdownTimer struct {
	DueTime           *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	GmtCreateTime     *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RemainingTimeInMs *int64  `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
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

type ListInstancesResponseBodyInstancesInstanceSnapshotList struct {
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName       *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ReasonCode      *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage   *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RepositoryUrl   *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesInstanceSnapshotList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesInstanceSnapshotList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetImageId(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetImageName(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ImageName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetImageUrl(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ImageUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetReasonCode(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ReasonCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetReasonMessage(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetRepositoryUrl(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.RepositoryUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetStatus(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.Status = &v
	return s
}

type ListInstancesResponseBodyInstancesLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesLabels) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesLabels) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesLabels) SetKey(v string) *ListInstancesResponseBodyInstancesLabels {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLabels) SetValue(v string) *ListInstancesResponseBodyInstancesLabels {
	s.Value = &v
	return s
}

type ListInstancesResponseBodyInstancesLatestSnapshot struct {
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName       *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ReasonCode      *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage   *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RepositoryUrl   *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetReasonCode(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ReasonCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetReasonMessage(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetRepositoryUrl(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.RepositoryUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetStatus(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.Status = &v
	return s
}

type ListInstancesResponseBodyInstancesRequestedResource struct {
	CPU          *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU          *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType      *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s ListInstancesResponseBodyInstancesRequestedResource) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesRequestedResource) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetCPU(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.CPU = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetGPU(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.GPU = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetGPUType(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.GPUType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetMemory(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.Memory = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetSharedMemory(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.SharedMemory = &v
	return s
}

type ListInstancesResponseBodyInstancesUserVpc struct {
	DefaultRoute    *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs   []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId       *string   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListInstancesResponseBodyInstancesUserVpc) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesUserVpc) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetDefaultRoute(v string) *ListInstancesResponseBodyInstancesUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetExtendedCIDRs(v []*string) *ListInstancesResponseBodyInstancesUserVpc {
	s.ExtendedCIDRs = v
	return s
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type StartInstanceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartInstanceResponse) SetStatusCode(v int32) *StartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
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
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	Accessibility        *string                                 `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Datasets             []*UpdateInstanceRequestDatasets        `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	DisassociateDatasets *bool                                   `json:"DisassociateDatasets,omitempty" xml:"DisassociateDatasets,omitempty"`
	DisassociateVpc      *bool                                   `json:"DisassociateVpc,omitempty" xml:"DisassociateVpc,omitempty"`
	EcsSpec              *string                                 `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	ImageId              *string                                 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUrl             *string                                 `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	InstanceName         *string                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RequestedResource    *UpdateInstanceRequestRequestedResource `json:"RequestedResource,omitempty" xml:"RequestedResource,omitempty" type:"Struct"`
	UserId               *string                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserVpc              *UpdateInstanceRequestUserVpc           `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	WorkspaceSource      *string                                 `json:"WorkspaceSource,omitempty" xml:"WorkspaceSource,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetAccessibility(v string) *UpdateInstanceRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateInstanceRequest) SetDatasets(v []*UpdateInstanceRequestDatasets) *UpdateInstanceRequest {
	s.Datasets = v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateDatasets(v bool) *UpdateInstanceRequest {
	s.DisassociateDatasets = &v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateVpc(v bool) *UpdateInstanceRequest {
	s.DisassociateVpc = &v
	return s
}

func (s *UpdateInstanceRequest) SetEcsSpec(v string) *UpdateInstanceRequest {
	s.EcsSpec = &v
	return s
}

func (s *UpdateInstanceRequest) SetImageId(v string) *UpdateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateInstanceRequest) SetImageUrl(v string) *UpdateInstanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetRequestedResource(v *UpdateInstanceRequestRequestedResource) *UpdateInstanceRequest {
	s.RequestedResource = v
	return s
}

func (s *UpdateInstanceRequest) SetUserId(v string) *UpdateInstanceRequest {
	s.UserId = &v
	return s
}

func (s *UpdateInstanceRequest) SetUserVpc(v *UpdateInstanceRequestUserVpc) *UpdateInstanceRequest {
	s.UserVpc = v
	return s
}

func (s *UpdateInstanceRequest) SetWorkspaceSource(v string) *UpdateInstanceRequest {
	s.WorkspaceSource = &v
	return s
}

type UpdateInstanceRequestDatasets struct {
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s UpdateInstanceRequestDatasets) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestDatasets) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestDatasets) SetDatasetId(v string) *UpdateInstanceRequestDatasets {
	s.DatasetId = &v
	return s
}

func (s *UpdateInstanceRequestDatasets) SetMountPath(v string) *UpdateInstanceRequestDatasets {
	s.MountPath = &v
	return s
}

type UpdateInstanceRequestRequestedResource struct {
	CPU          *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU          *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType      *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s UpdateInstanceRequestRequestedResource) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestRequestedResource) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestRequestedResource) SetCPU(v string) *UpdateInstanceRequestRequestedResource {
	s.CPU = &v
	return s
}

func (s *UpdateInstanceRequestRequestedResource) SetGPU(v string) *UpdateInstanceRequestRequestedResource {
	s.GPU = &v
	return s
}

func (s *UpdateInstanceRequestRequestedResource) SetGPUType(v string) *UpdateInstanceRequestRequestedResource {
	s.GPUType = &v
	return s
}

func (s *UpdateInstanceRequestRequestedResource) SetMemory(v string) *UpdateInstanceRequestRequestedResource {
	s.Memory = &v
	return s
}

func (s *UpdateInstanceRequestRequestedResource) SetSharedMemory(v string) *UpdateInstanceRequestRequestedResource {
	s.SharedMemory = &v
	return s
}

type UpdateInstanceRequestUserVpc struct {
	DefaultRoute    *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs   []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId       *string   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateInstanceRequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestUserVpc) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestUserVpc) SetDefaultRoute(v string) *UpdateInstanceRequestUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetExtendedCIDRs(v []*string) *UpdateInstanceRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetSecurityGroupId(v string) *UpdateInstanceRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetVSwitchId(v string) *UpdateInstanceRequestUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetVpcId(v string) *UpdateInstanceRequestUserVpc {
	s.VpcId = &v
	return s
}

type UpdateInstanceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
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

func (client *Client) CreateIdleInstanceCullerWithOptions(InstanceId *string, request *CreateIdleInstanceCullerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIdleInstanceCullerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CpuPercentThreshold)) {
		body["CpuPercentThreshold"] = request.CpuPercentThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.GpuPercentThreshold)) {
		body["GpuPercentThreshold"] = request.GpuPercentThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.MaxIdleTimeInMinutes)) {
		body["MaxIdleTimeInMinutes"] = request.MaxIdleTimeInMinutes
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIdleInstanceCuller"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/idleinstanceculler"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIdleInstanceCullerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIdleInstanceCuller(InstanceId *string, request *CreateIdleInstanceCullerRequest) (_result *CreateIdleInstanceCullerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIdleInstanceCullerResponse{}
	_body, _err := client.CreateIdleInstanceCullerWithOptions(InstanceId, request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.CloudDisks)) {
		body["CloudDisks"] = request.CloudDisks
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

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RequestedResource)) {
		body["RequestedResource"] = request.RequestedResource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserVpc)) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceSource)) {
		body["WorkspaceSource"] = request.WorkspaceSource
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

func (client *Client) CreateInstanceShutdownTimerWithOptions(InstanceId *string, request *CreateInstanceShutdownTimerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceShutdownTimerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/shutdowntimer"),
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

func (client *Client) CreateInstanceSnapshotWithOptions(InstanceId *string, request *CreateInstanceSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
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
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/snapshots"),
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

func (client *Client) DeleteIdleInstanceCullerWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIdleInstanceCullerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIdleInstanceCuller"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/idleinstanceculler"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIdleInstanceCullerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIdleInstanceCuller(InstanceId *string) (_result *DeleteIdleInstanceCullerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIdleInstanceCullerResponse{}
	_body, _err := client.DeleteIdleInstanceCullerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId))),
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

func (client *Client) DeleteInstanceShutdownTimerWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceShutdownTimerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceShutdownTimer"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/shutdowntimer"),
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

func (client *Client) DeleteInstanceSnapshotWithOptions(InstanceId *string, SnapshotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceSnapshotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceSnapshot"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/snapshots/" + tea.StringValue(openapiutil.GetEncodeParam(SnapshotId))),
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

func (client *Client) GetIdleInstanceCullerWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIdleInstanceCullerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIdleInstanceCuller"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/idleinstanceculler"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIdleInstanceCullerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIdleInstanceCuller(InstanceId *string) (_result *GetIdleInstanceCullerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIdleInstanceCullerResponse{}
	_body, _err := client.GetIdleInstanceCullerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId))),
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

func (client *Client) GetInstanceMetricsWithOptions(InstanceId *string, request *GetInstanceMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceMetrics"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/metrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceMetrics(InstanceId *string, request *GetInstanceMetricsRequest) (_result *GetInstanceMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceMetricsResponse{}
	_body, _err := client.GetInstanceMetricsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceShutdownTimerWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceShutdownTimerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceShutdownTimer"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/shutdowntimer"),
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

func (client *Client) GetInstanceSnapshotWithOptions(InstanceId *string, SnapshotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceSnapshotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceSnapshot"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/snapshots/" + tea.StringValue(openapiutil.GetEncodeParam(SnapshotId))),
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

func (client *Client) GetLifecycleWithOptions(InstanceId *string, request *GetLifecycleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLifecycleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.SessionNumber)) {
		query["SessionNumber"] = request.SessionNumber
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLifecycle"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/lifecycle"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLifecycleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLifecycle(InstanceId *string, request *GetLifecycleRequest) (_result *GetLifecycleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLifecycleResponse{}
	_body, _err := client.GetLifecycleWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceGroupStatisticsWithOptions(request *GetResourceGroupStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceGroupStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceGroupStatistics"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resourcegroupstatistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceGroupStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceGroupStatistics(request *GetResourceGroupStatisticsRequest) (_result *GetResourceGroupStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupStatisticsResponse{}
	_body, _err := client.GetResourceGroupStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/tokens"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
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

func (client *Client) ListInstanceSnapshotWithOptions(InstanceId *string, request *ListInstanceSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/snapshots"),
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

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
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

func (client *Client) StartInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/start"),
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

func (client *Client) StopInstanceWithOptions(InstanceId *string, request *StopInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/stop"),
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

func (client *Client) UpdateInstanceWithOptions(InstanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.DisassociateDatasets)) {
		body["DisassociateDatasets"] = request.DisassociateDatasets
	}

	if !tea.BoolValue(util.IsUnset(request.DisassociateVpc)) {
		body["DisassociateVpc"] = request.DisassociateVpc
	}

	if !tea.BoolValue(util.IsUnset(request.EcsSpec)) {
		body["EcsSpec"] = request.EcsSpec
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

	if !tea.BoolValue(util.IsUnset(request.RequestedResource)) {
		body["RequestedResource"] = request.RequestedResource
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserVpc)) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceSource)) {
		body["WorkspaceSource"] = request.WorkspaceSource
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId))),
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
