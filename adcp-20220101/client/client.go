// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Tag struct {
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Tag) String() string {
	return tea.Prettify(s)
}

func (s Tag) GoString() string {
	return s.String()
}

func (s *Tag) SetKey(v string) *Tag {
	s.Key = &v
	return s
}

func (s *Tag) SetValue(v string) *Tag {
	s.Value = &v
	return s
}

type AttachClusterToHubRequest struct {
	// The operation that you want to perform. Set the value to **AttachClusterToHub**.
	//
	// example:
	//
	// true
	AttachToMesh *bool `json:"AttachToMesh,omitempty" xml:"AttachToMesh,omitempty"`
	// The ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// cd08d62e6506a4fa5a8c44c19d0fc****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["cdea10134be464ba4acb36cc831a6****"]
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
}

func (s AttachClusterToHubRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachClusterToHubRequest) GoString() string {
	return s.String()
}

func (s *AttachClusterToHubRequest) SetAttachToMesh(v bool) *AttachClusterToHubRequest {
	s.AttachToMesh = &v
	return s
}

func (s *AttachClusterToHubRequest) SetClusterId(v string) *AttachClusterToHubRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachClusterToHubRequest) SetClusterIds(v string) *AttachClusterToHubRequest {
	s.ClusterIds = &v
	return s
}

type AttachClusterToHubResponseBody struct {
	// You can call the AttachClusterToHub operation to associate an Container Service for Kubernetes (ACK) cluster with a master instance of ACK One.
	//
	// example:
	//
	// c8e28143817db4b039b8548d7c899****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Zhishi
	ManagedClusterIds []*string `json:"ManagedClusterIds,omitempty" xml:"ManagedClusterIds,omitempty" type:"Repeated"`
	// Example 1
	//
	// example:
	//
	// EA06613B-37A3-549E-BAE0-E4AD8A6E93D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// T-623a96b7bbeaac074b00****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AttachClusterToHubResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachClusterToHubResponseBody) GoString() string {
	return s.String()
}

func (s *AttachClusterToHubResponseBody) SetClusterId(v string) *AttachClusterToHubResponseBody {
	s.ClusterId = &v
	return s
}

func (s *AttachClusterToHubResponseBody) SetManagedClusterIds(v []*string) *AttachClusterToHubResponseBody {
	s.ManagedClusterIds = v
	return s
}

func (s *AttachClusterToHubResponseBody) SetRequestId(v string) *AttachClusterToHubResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachClusterToHubResponseBody) SetTaskId(v string) *AttachClusterToHubResponseBody {
	s.TaskId = &v
	return s
}

type AttachClusterToHubResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachClusterToHubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachClusterToHubResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachClusterToHubResponse) GoString() string {
	return s.String()
}

func (s *AttachClusterToHubResponse) SetHeaders(v map[string]*string) *AttachClusterToHubResponse {
	s.Headers = v
	return s
}

func (s *AttachClusterToHubResponse) SetStatusCode(v int32) *AttachClusterToHubResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachClusterToHubResponse) SetBody(v *AttachClusterToHubResponseBody) *AttachClusterToHubResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The ID of the new resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzlvgbhaca***
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The resource ID. If ResourceType=cluster, the resource ID is ClusterId.
	//
	// This parameter is required.
	//
	// example:
	//
	// c9603ee23a84a41d6a1424619cb80****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Only cluster are supported. Set the value to cluster.
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5BE4C329-DCC2-5B61-8F66-112B7D7FC712
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CreateHubClusterRequest struct {
	// Specifies whether to expose the API server to the Internet. Valid values:
	//
	// 	- true: exposes the API server to the Internet.
	//
	// 	- false: exposes the API server to the internal network.
	//
	// example:
	//
	// true
	ApiServerPublicEip *bool `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	// Specifies whether to enable the workflow instance UI. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ArgoServerEnabled *bool `json:"ArgoServerEnabled,omitempty" xml:"ArgoServerEnabled,omitempty"`
	// Specifies whether to enable the audit log feature. Valid values:
	//
	// 	- true: enables the audit log feature.
	//
	// 	- false: disables the audit log feature.
	//
	// example:
	//
	// false
	AuditLogEnabled *bool `json:"AuditLogEnabled,omitempty" xml:"AuditLogEnabled,omitempty"`
	// Specifies whether to use an advanced security group.
	//
	// example:
	//
	// false
	IsEnterpriseSecurityGroup *bool `json:"IsEnterpriseSecurityGroup,omitempty" xml:"IsEnterpriseSecurityGroup,omitempty"`
	// The name of the master instance.
	//
	// example:
	//
	// ack-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The limit on the prices of containers in the workflow. This parameter takes effect only if the WorkflowScheduleMode parameter is set to cost-optimized.
	//
	// example:
	//
	// 0.08
	PriceLimit *string `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	// The type of scenario for which the master instance is suitable. Valid values:
	//
	// 	- `Default`: The master instance is suitable for standard scenarios.
	//
	// 	- `XFlow`: The master instance is suitable for workflow scenarios.
	//
	// Default value: `Default`.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the region. You can call the DescribeRegions operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Resource Group ID.
	//
	// example:
	//
	// rg-1exm6tg2h48***
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	Tag             []*Tag  `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["vsw-2zeaijsas4zkzz81xm***"]
	VSwitches *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the master instance belongs. You can call the DescribeVpcs operation to query available VPCs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-f8zin0jscsr84s96tg***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The scheduling mode of the workflow. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- cost-optimized: cost-prioritized scheduling mode.
	//
	// 	- stock-optimized: inventory-prioritized scheduling mode.
	//
	// example:
	//
	// cost-optimized
	WorkflowScheduleMode *string `json:"WorkflowScheduleMode,omitempty" xml:"WorkflowScheduleMode,omitempty"`
}

func (s CreateHubClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHubClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateHubClusterRequest) SetApiServerPublicEip(v bool) *CreateHubClusterRequest {
	s.ApiServerPublicEip = &v
	return s
}

func (s *CreateHubClusterRequest) SetArgoServerEnabled(v bool) *CreateHubClusterRequest {
	s.ArgoServerEnabled = &v
	return s
}

func (s *CreateHubClusterRequest) SetAuditLogEnabled(v bool) *CreateHubClusterRequest {
	s.AuditLogEnabled = &v
	return s
}

func (s *CreateHubClusterRequest) SetIsEnterpriseSecurityGroup(v bool) *CreateHubClusterRequest {
	s.IsEnterpriseSecurityGroup = &v
	return s
}

func (s *CreateHubClusterRequest) SetName(v string) *CreateHubClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateHubClusterRequest) SetPriceLimit(v string) *CreateHubClusterRequest {
	s.PriceLimit = &v
	return s
}

func (s *CreateHubClusterRequest) SetProfile(v string) *CreateHubClusterRequest {
	s.Profile = &v
	return s
}

func (s *CreateHubClusterRequest) SetRegionId(v string) *CreateHubClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHubClusterRequest) SetResourceGroupID(v string) *CreateHubClusterRequest {
	s.ResourceGroupID = &v
	return s
}

func (s *CreateHubClusterRequest) SetTag(v []*Tag) *CreateHubClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateHubClusterRequest) SetVSwitches(v string) *CreateHubClusterRequest {
	s.VSwitches = &v
	return s
}

func (s *CreateHubClusterRequest) SetVpcId(v string) *CreateHubClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateHubClusterRequest) SetWorkflowScheduleMode(v string) *CreateHubClusterRequest {
	s.WorkflowScheduleMode = &v
	return s
}

type CreateHubClusterShrinkRequest struct {
	// Specifies whether to expose the API server to the Internet. Valid values:
	//
	// 	- true: exposes the API server to the Internet.
	//
	// 	- false: exposes the API server to the internal network.
	//
	// example:
	//
	// true
	ApiServerPublicEip *bool `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	// Specifies whether to enable the workflow instance UI. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ArgoServerEnabled *bool `json:"ArgoServerEnabled,omitempty" xml:"ArgoServerEnabled,omitempty"`
	// Specifies whether to enable the audit log feature. Valid values:
	//
	// 	- true: enables the audit log feature.
	//
	// 	- false: disables the audit log feature.
	//
	// example:
	//
	// false
	AuditLogEnabled *bool `json:"AuditLogEnabled,omitempty" xml:"AuditLogEnabled,omitempty"`
	// Specifies whether to use an advanced security group.
	//
	// example:
	//
	// false
	IsEnterpriseSecurityGroup *bool `json:"IsEnterpriseSecurityGroup,omitempty" xml:"IsEnterpriseSecurityGroup,omitempty"`
	// The name of the master instance.
	//
	// example:
	//
	// ack-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The limit on the prices of containers in the workflow. This parameter takes effect only if the WorkflowScheduleMode parameter is set to cost-optimized.
	//
	// example:
	//
	// 0.08
	PriceLimit *string `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	// The type of scenario for which the master instance is suitable. Valid values:
	//
	// 	- `Default`: The master instance is suitable for standard scenarios.
	//
	// 	- `XFlow`: The master instance is suitable for workflow scenarios.
	//
	// Default value: `Default`.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the region. You can call the DescribeRegions operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Resource Group ID.
	//
	// example:
	//
	// rg-1exm6tg2h48***
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	TagShrink       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["vsw-2zeaijsas4zkzz81xm***"]
	VSwitches *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the master instance belongs. You can call the DescribeVpcs operation to query available VPCs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-f8zin0jscsr84s96tg***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The scheduling mode of the workflow. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- cost-optimized: cost-prioritized scheduling mode.
	//
	// 	- stock-optimized: inventory-prioritized scheduling mode.
	//
	// example:
	//
	// cost-optimized
	WorkflowScheduleMode *string `json:"WorkflowScheduleMode,omitempty" xml:"WorkflowScheduleMode,omitempty"`
}

func (s CreateHubClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHubClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHubClusterShrinkRequest) SetApiServerPublicEip(v bool) *CreateHubClusterShrinkRequest {
	s.ApiServerPublicEip = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetArgoServerEnabled(v bool) *CreateHubClusterShrinkRequest {
	s.ArgoServerEnabled = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetAuditLogEnabled(v bool) *CreateHubClusterShrinkRequest {
	s.AuditLogEnabled = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetIsEnterpriseSecurityGroup(v bool) *CreateHubClusterShrinkRequest {
	s.IsEnterpriseSecurityGroup = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetName(v string) *CreateHubClusterShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetPriceLimit(v string) *CreateHubClusterShrinkRequest {
	s.PriceLimit = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetProfile(v string) *CreateHubClusterShrinkRequest {
	s.Profile = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetRegionId(v string) *CreateHubClusterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetResourceGroupID(v string) *CreateHubClusterShrinkRequest {
	s.ResourceGroupID = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetTagShrink(v string) *CreateHubClusterShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetVSwitches(v string) *CreateHubClusterShrinkRequest {
	s.VSwitches = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetVpcId(v string) *CreateHubClusterShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateHubClusterShrinkRequest) SetWorkflowScheduleMode(v string) *CreateHubClusterShrinkRequest {
	s.WorkflowScheduleMode = &v
	return s
}

type CreateHubClusterResponseBody struct {
	// The ID of the master instance.
	//
	// example:
	//
	// c09946603cd764dac96135f51d1ba****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 62F5AA2B-A2C9-5135-BCE2-C2167099****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// T-62523eda841eca071400****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateHubClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHubClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHubClusterResponseBody) SetClusterId(v string) *CreateHubClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateHubClusterResponseBody) SetRequestId(v string) *CreateHubClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHubClusterResponseBody) SetTaskId(v string) *CreateHubClusterResponseBody {
	s.TaskId = &v
	return s
}

type CreateHubClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHubClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHubClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHubClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateHubClusterResponse) SetHeaders(v map[string]*string) *CreateHubClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateHubClusterResponse) SetStatusCode(v int32) *CreateHubClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHubClusterResponse) SetBody(v *CreateHubClusterResponseBody) *CreateHubClusterResponse {
	s.Body = v
	return s
}

type DeleteHubClusterRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cd90dd24a86fd42f895a1b77df620****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to forcefully delete the master instance in ACK One. Valid values:
	//
	// 	- true: forcefully deletes the master instance in ACK One.
	//
	// 	- false: does not forcibly delete the master instance in ACK One.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The list of resources to retain.
	RetainResources []*string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty" type:"Repeated"`
}

func (s DeleteHubClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHubClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteHubClusterRequest) SetClusterId(v string) *DeleteHubClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHubClusterRequest) SetForce(v bool) *DeleteHubClusterRequest {
	s.Force = &v
	return s
}

func (s *DeleteHubClusterRequest) SetRetainResources(v []*string) *DeleteHubClusterRequest {
	s.RetainResources = v
	return s
}

type DeleteHubClusterShrinkRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cd90dd24a86fd42f895a1b77df620****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to forcefully delete the master instance in ACK One. Valid values:
	//
	// 	- true: forcefully deletes the master instance in ACK One.
	//
	// 	- false: does not forcibly delete the master instance in ACK One.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The list of resources to retain.
	RetainResourcesShrink *string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty"`
}

func (s DeleteHubClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHubClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteHubClusterShrinkRequest) SetClusterId(v string) *DeleteHubClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHubClusterShrinkRequest) SetForce(v bool) *DeleteHubClusterShrinkRequest {
	s.Force = &v
	return s
}

func (s *DeleteHubClusterShrinkRequest) SetRetainResourcesShrink(v string) *DeleteHubClusterShrinkRequest {
	s.RetainResourcesShrink = &v
	return s
}

type DeleteHubClusterResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// cb09fda0dc2f94a8397c76638c7ec****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7A827E32-6D24-5757-B3FD-D9396495FBDC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// T-623a96b7bbeaac074b00****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteHubClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHubClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHubClusterResponseBody) SetClusterId(v string) *DeleteHubClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DeleteHubClusterResponseBody) SetRequestId(v string) *DeleteHubClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHubClusterResponseBody) SetTaskId(v string) *DeleteHubClusterResponseBody {
	s.TaskId = &v
	return s
}

type DeleteHubClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHubClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHubClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHubClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteHubClusterResponse) SetHeaders(v map[string]*string) *DeleteHubClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteHubClusterResponse) SetStatusCode(v int32) *DeleteHubClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHubClusterResponse) SetBody(v *DeleteHubClusterResponseBody) *DeleteHubClusterResponse {
	s.Body = v
	return s
}

type DeletePolicyInstanceRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c09946603cd764dac96135f51d1ba****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// A array of JSON strings. The JSON strings in the array indicate the IDs of the associated clusters for which the policy is deleted.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACKNoEnvVarSecrets
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DeletePolicyInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyInstanceRequest) SetClusterId(v string) *DeletePolicyInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeletePolicyInstanceRequest) SetClusterIds(v []*string) *DeletePolicyInstanceRequest {
	s.ClusterIds = v
	return s
}

func (s *DeletePolicyInstanceRequest) SetPolicyName(v string) *DeletePolicyInstanceRequest {
	s.PolicyName = &v
	return s
}

type DeletePolicyInstanceShrinkRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c09946603cd764dac96135f51d1ba****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// A array of JSON strings. The JSON strings in the array indicate the IDs of the associated clusters for which the policy is deleted.
	ClusterIdsShrink *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACKNoEnvVarSecrets
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DeletePolicyInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyInstanceShrinkRequest) SetClusterId(v string) *DeletePolicyInstanceShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeletePolicyInstanceShrinkRequest) SetClusterIdsShrink(v string) *DeletePolicyInstanceShrinkRequest {
	s.ClusterIdsShrink = &v
	return s
}

func (s *DeletePolicyInstanceShrinkRequest) SetPolicyName(v string) *DeletePolicyInstanceShrinkRequest {
	s.PolicyName = &v
	return s
}

type DeletePolicyInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EA06613B-37A3-549E-BAE0-E4AD8A6E93D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyInstanceResponseBody) SetRequestId(v string) *DeletePolicyInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyInstanceResponse) SetHeaders(v map[string]*string) *DeletePolicyInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyInstanceResponse) SetStatusCode(v int32) *DeletePolicyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyInstanceResponse) SetBody(v *DeletePolicyInstanceResponseBody) *DeletePolicyInstanceResponse {
	s.Body = v
	return s
}

type DeleteUserPermissionRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c09946603cd764dac96135f51d1ba****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2176****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPermissionRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserPermissionRequest) SetClusterId(v string) *DeleteUserPermissionRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUserPermissionRequest) SetUserId(v string) *DeleteUserPermissionRequest {
	s.UserId = &v
	return s
}

type DeleteUserPermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2D676EFC-8C04-5CCE-A08E-BB97D24B47E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserPermissionResponseBody) SetRequestId(v string) *DeleteUserPermissionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserPermissionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserPermissionResponse) SetHeaders(v map[string]*string) *DeleteUserPermissionResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserPermissionResponse) SetStatusCode(v int32) *DeleteUserPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserPermissionResponse) SetBody(v *DeleteUserPermissionResponseBody) *DeleteUserPermissionResponse {
	s.Body = v
	return s
}

type DeployPolicyInstanceRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c102fe5f1ee5d4c87a68121a77d8b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// An array of JSON strings. The JSON strings in the array indicate the IDs of the associated clusters in which the policy instance is deployed.
	//
	// This parameter is required.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
	// A list of namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The action of the policy. Valid values:
	//
	// 	- deny: blocks deployments that match the policy.
	//
	// 	- warn: generates alerts for deployments that match the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// warn
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACKNoEnvVarSecrets
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DeployPolicyInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployPolicyInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeployPolicyInstanceRequest) SetClusterId(v string) *DeployPolicyInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeployPolicyInstanceRequest) SetClusterIds(v []*string) *DeployPolicyInstanceRequest {
	s.ClusterIds = v
	return s
}

func (s *DeployPolicyInstanceRequest) SetNamespaces(v []*string) *DeployPolicyInstanceRequest {
	s.Namespaces = v
	return s
}

func (s *DeployPolicyInstanceRequest) SetPolicyAction(v string) *DeployPolicyInstanceRequest {
	s.PolicyAction = &v
	return s
}

func (s *DeployPolicyInstanceRequest) SetPolicyName(v string) *DeployPolicyInstanceRequest {
	s.PolicyName = &v
	return s
}

type DeployPolicyInstanceShrinkRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c102fe5f1ee5d4c87a68121a77d8b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// An array of JSON strings. The JSON strings in the array indicate the IDs of the associated clusters in which the policy instance is deployed.
	//
	// This parameter is required.
	ClusterIdsShrink *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// A list of namespaces.
	NamespacesShrink *string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty"`
	// The action of the policy. Valid values:
	//
	// 	- deny: blocks deployments that match the policy.
	//
	// 	- warn: generates alerts for deployments that match the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// warn
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACKNoEnvVarSecrets
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DeployPolicyInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployPolicyInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeployPolicyInstanceShrinkRequest) SetClusterId(v string) *DeployPolicyInstanceShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeployPolicyInstanceShrinkRequest) SetClusterIdsShrink(v string) *DeployPolicyInstanceShrinkRequest {
	s.ClusterIdsShrink = &v
	return s
}

func (s *DeployPolicyInstanceShrinkRequest) SetNamespacesShrink(v string) *DeployPolicyInstanceShrinkRequest {
	s.NamespacesShrink = &v
	return s
}

func (s *DeployPolicyInstanceShrinkRequest) SetPolicyAction(v string) *DeployPolicyInstanceShrinkRequest {
	s.PolicyAction = &v
	return s
}

func (s *DeployPolicyInstanceShrinkRequest) SetPolicyName(v string) *DeployPolicyInstanceShrinkRequest {
	s.PolicyName = &v
	return s
}

type DeployPolicyInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D676EFC-8C04-5CCE-A08E-BB97D24B47E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployPolicyInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployPolicyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeployPolicyInstanceResponseBody) SetRequestId(v string) *DeployPolicyInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeployPolicyInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployPolicyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployPolicyInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployPolicyInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeployPolicyInstanceResponse) SetHeaders(v map[string]*string) *DeployPolicyInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeployPolicyInstanceResponse) SetStatusCode(v int32) *DeployPolicyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployPolicyInstanceResponse) SetBody(v *DeployPolicyInstanceResponseBody) *DeployPolicyInstanceResponse {
	s.Body = v
	return s
}

type DescribeHubClusterDetailsRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c676decda7b8148d69a9aba751877****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeHubClusterDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsRequest) SetClusterId(v string) *DescribeHubClusterDetailsRequest {
	s.ClusterId = &v
	return s
}

type DescribeHubClusterDetailsResponseBody struct {
	// The details of the master instance.
	Cluster *DescribeHubClusterDetailsResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 52C1B7DF-96C1-5214-97B6-1B0859FEAFE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBody) SetCluster(v *DescribeHubClusterDetailsResponseBodyCluster) *DescribeHubClusterDetailsResponseBody {
	s.Cluster = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBody) SetRequestId(v string) *DescribeHubClusterDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyCluster struct {
	// The details of the API server of the master instance.
	ApiServer *DescribeHubClusterDetailsResponseBodyClusterApiServer `json:"ApiServer,omitempty" xml:"ApiServer,omitempty" type:"Struct"`
	// The details of the master instance.
	ClusterInfo *DescribeHubClusterDetailsResponseBodyClusterClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// The deletion conditions of the master instance.
	Conditions []*DescribeHubClusterDetailsResponseBodyClusterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The endpoint of the master instance.
	Endpoints *DescribeHubClusterDetailsResponseBodyClusterEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The logging configuration.
	LogConfig *DescribeHubClusterDetailsResponseBodyClusterLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// The configurations of Alibaba Cloud Service Mesh (ASM).
	MeshConfig *DescribeHubClusterDetailsResponseBodyClusterMeshConfig `json:"MeshConfig,omitempty" xml:"MeshConfig,omitempty" type:"Struct"`
	// The network configurations of the master instance.
	Network *DescribeHubClusterDetailsResponseBodyClusterNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The Argo workflow configuration.
	WorkflowConfig *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig `json:"WorkflowConfig,omitempty" xml:"WorkflowConfig,omitempty" type:"Struct"`
}

func (s DescribeHubClusterDetailsResponseBodyCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyCluster) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetApiServer(v *DescribeHubClusterDetailsResponseBodyClusterApiServer) *DescribeHubClusterDetailsResponseBodyCluster {
	s.ApiServer = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetClusterInfo(v *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) *DescribeHubClusterDetailsResponseBodyCluster {
	s.ClusterInfo = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetConditions(v []*DescribeHubClusterDetailsResponseBodyClusterConditions) *DescribeHubClusterDetailsResponseBodyCluster {
	s.Conditions = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetEndpoints(v *DescribeHubClusterDetailsResponseBodyClusterEndpoints) *DescribeHubClusterDetailsResponseBodyCluster {
	s.Endpoints = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetLogConfig(v *DescribeHubClusterDetailsResponseBodyClusterLogConfig) *DescribeHubClusterDetailsResponseBodyCluster {
	s.LogConfig = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetMeshConfig(v *DescribeHubClusterDetailsResponseBodyClusterMeshConfig) *DescribeHubClusterDetailsResponseBodyCluster {
	s.MeshConfig = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetNetwork(v *DescribeHubClusterDetailsResponseBodyClusterNetwork) *DescribeHubClusterDetailsResponseBodyCluster {
	s.Network = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetWorkflowConfig(v *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig) *DescribeHubClusterDetailsResponseBodyCluster {
	s.WorkflowConfig = v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterApiServer struct {
	// The ID of the elastic IP address (EIP).
	//
	// example:
	//
	// eip-abc****
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// Indicates whether the API server is accessible over the Internet. Valid values:
	//
	// 	- true: The API server is accessible over the Internet.
	//
	// 	- false: The API server is inaccessible over the Internet.
	//
	// example:
	//
	// true
	EnabledPublic *bool `json:"EnabledPublic,omitempty" xml:"EnabledPublic,omitempty"`
	// The ID of the Server Load Balancer (SLB) instance.
	//
	// example:
	//
	// lb-hp3ioqbfeq37h13rwe***
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterApiServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterApiServer) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterApiServer) SetApiServerEipId(v string) *DescribeHubClusterDetailsResponseBodyClusterApiServer {
	s.ApiServerEipId = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterApiServer) SetEnabledPublic(v bool) *DescribeHubClusterDetailsResponseBodyClusterApiServer {
	s.EnabledPublic = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterApiServer) SetLoadBalancerId(v string) *DescribeHubClusterDetailsResponseBodyClusterApiServer {
	s.LoadBalancerId = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterClusterInfo struct {
	// The ID of the master instance.
	//
	// example:
	//
	// cb09fda0dc2f94a8397c76638c7ecf***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The specification of the master instance. Valid value:
	//
	// 	- ack.pro.small: ACK Pro cluster
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The time when the master instance was created.
	//
	// example:
	//
	// 2022-03-23T06:22:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message returned when the master instance failed to be created.
	//
	// example:
	//
	// The specified product does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The cluster metadata.
	MetaData *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaData `json:"MetaData,omitempty" xml:"MetaData,omitempty" type:"Struct"`
	// The name of the master instance.
	//
	// example:
	//
	// ackone-heyuan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of the master instance.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the region in which the master instance resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of Resource Group.
	//
	// example:
	//
	// rg-1w4vdvo6p51***
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	// The status of the master instance. Valid values:
	//
	// 	- initial: The master instance is being initialized.
	//
	// 	- failed: The master instance failed to be created.
	//
	// 	- running: The master instance is running
	//
	// 	- inactive: The master instance is pending.
	//
	// 	- deleting: The master instance is being deleted.
	//
	// 	- delete_failed: The master instance failed to be deleted.
	//
	// 	- deleted: The master instance is deleted.
	//
	// example:
	//
	// running
	State *string                                                        `json:"State,omitempty" xml:"State,omitempty"`
	Tags  []*DescribeHubClusterDetailsResponseBodyClusterClusterInfoTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the master instance was updated.
	//
	// example:
	//
	// 2022-03-21T02:51:35.542Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version of the master instance.
	//
	// example:
	//
	// 1.22.3-aliyun.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetClusterId(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetClusterSpec(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetCreationTime(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.CreationTime = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetErrorMessage(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetMetaData(v *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaData) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.MetaData = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetName(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.Name = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetProfile(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.Profile = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetRegionId(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetResourceGroupID(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.ResourceGroupID = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetState(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.State = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetTags(v []*DescribeHubClusterDetailsResponseBodyClusterClusterInfoTags) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.Tags = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetUpdateTime(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetVersion(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.Version = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaData struct {
	// The cluster metadata.
	ACKOne *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOne `json:"ACKOne,omitempty" xml:"ACKOne,omitempty" type:"Struct"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaData) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaData) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaData) SetACKOne(v *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOne) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaData {
	s.ACKOne = v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOne struct {
	// The GitOps metadata.
	GitOps *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps `json:"GitOps,omitempty" xml:"GitOps,omitempty" type:"Struct"`
	// The workflow metadata.
	WorkFlow *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlow `json:"WorkFlow,omitempty" xml:"WorkFlow,omitempty" type:"Struct"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOne) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOne) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOne) SetGitOps(v *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOne {
	s.GitOps = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOne) SetWorkFlow(v *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlow) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOne {
	s.WorkFlow = v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps struct {
	// The Internet access control list (ACL). This parameter takes effect only if PublicAccessEnabled is set to true.
	AccessControlList []*string `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty" type:"Repeated"`
	// Indicates whether GitOps is enabled. Valid values:
	//
	// 	- true: GitOps is enabled.
	//
	// 	- false: GitOps is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether GitOps High Availability is enabled. Valid values:
	//
	// 	- true:  GitOps High Availability is enabled.
	//
	// 	- false:  GitOps High Availability is disabled.
	//
	// example:
	//
	// true
	HAEnabled *bool `json:"HAEnabled,omitempty" xml:"HAEnabled,omitempty"`
	// Specifies whether to enable public domain name resolution in the Argo CD or Argo Workflow console. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" xml:"PublicAccessEnabled,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps) SetAccessControlList(v []*string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps {
	s.AccessControlList = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps) SetEnabled(v bool) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps {
	s.Enabled = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps) SetHAEnabled(v bool) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps {
	s.HAEnabled = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps) SetPublicAccessEnabled(v bool) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneGitOps {
	s.PublicAccessEnabled = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlow struct {
	// The Argo workflow metadata.
	ArgoWorkflow *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow `json:"ArgoWorkflow,omitempty" xml:"ArgoWorkflow,omitempty" type:"Struct"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlow) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlow) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlow) SetArgoWorkflow(v *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlow {
	s.ArgoWorkflow = v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow struct {
	// The Internet access control list (ACL). This parameter takes effect only if PublicAccessEnabled is set to true.
	AccessControlList []*string `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty" type:"Repeated"`
	// Specifies whether to enable the argo workflow. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Specifies whether to enable public domain name resolution in the Argo CD or Argo Workflow console. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" xml:"PublicAccessEnabled,omitempty"`
	// Specifies whether to enable the argo workflow. UI Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// true
	ServerEnabled *string `json:"ServerEnabled,omitempty" xml:"ServerEnabled,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow) SetAccessControlList(v []*string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow {
	s.AccessControlList = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow) SetEnabled(v bool) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow {
	s.Enabled = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow) SetPublicAccessEnabled(v bool) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow {
	s.PublicAccessEnabled = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow) SetServerEnabled(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoMetaDataACKOneWorkFlowArgoWorkflow {
	s.ServerEnabled = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterClusterInfoTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterClusterInfoTags) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoTags) SetKey(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoTags {
	s.Key = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfoTags) SetValue(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfoTags {
	s.Value = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterConditions struct {
	// The error message returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the deletion condition.
	//
	// example:
	//
	// Successful
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the master instance that the deletion condition indicates. Valid values:
	//
	// 	- True: The master instance cannot be deleted.
	//
	// 	- False: The master instance can be deleted.
	//
	// 	- Unknow: Whether the master instance can be deleted is unknown.
	//
	// example:
	//
	// True
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of deletion condition.
	//
	// example:
	//
	// DeletionProtection
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterConditions) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterConditions) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterConditions) SetMessage(v string) *DescribeHubClusterDetailsResponseBodyClusterConditions {
	s.Message = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterConditions) SetReason(v string) *DescribeHubClusterDetailsResponseBodyClusterConditions {
	s.Reason = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterConditions) SetStatus(v string) *DescribeHubClusterDetailsResponseBodyClusterConditions {
	s.Status = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterConditions) SetType(v string) *DescribeHubClusterDetailsResponseBodyClusterConditions {
	s.Type = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterEndpoints struct {
	// The endpoint that is used to access the API server over the internal network.
	//
	// example:
	//
	// https://172.16.6.**:6443
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	// The endpoint that is used to access the API server over the Internet.
	//
	// example:
	//
	// https://123.57.21.***:6443
	PublicApiServerEndpoint *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterEndpoints) SetIntranetApiServerEndpoint(v string) *DescribeHubClusterDetailsResponseBodyClusterEndpoints {
	s.IntranetApiServerEndpoint = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterEndpoints) SetPublicApiServerEndpoint(v string) *DescribeHubClusterDetailsResponseBodyClusterEndpoints {
	s.PublicApiServerEndpoint = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterLogConfig struct {
	// Indicates whether the audit logging feature is enabled. Valid values:
	//
	// 	- true: Audit logging is enabled.
	//
	// 	- false: Audit logging is disabled.
	//
	// example:
	//
	// false
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// The name of the project of Log Service.
	//
	// example:
	//
	// k8s-log-abc
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The number of days that logs are retained by Log Service.
	//
	// example:
	//
	// 7
	LogStoreTTL *string `json:"LogStoreTTL,omitempty" xml:"LogStoreTTL,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterLogConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterLogConfig) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterLogConfig) SetEnableLog(v bool) *DescribeHubClusterDetailsResponseBodyClusterLogConfig {
	s.EnableLog = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterLogConfig) SetLogProject(v string) *DescribeHubClusterDetailsResponseBodyClusterLogConfig {
	s.LogProject = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterLogConfig) SetLogStoreTTL(v string) *DescribeHubClusterDetailsResponseBodyClusterLogConfig {
	s.LogStoreTTL = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterMeshConfig struct {
	// Indicates whether ASM is enabled. Valid values:
	//
	// 	- true: ASM is enabled.
	//
	// 	- false: ASM is disabled.
	//
	// example:
	//
	// false
	EnableMesh *bool `json:"EnableMesh,omitempty" xml:"EnableMesh,omitempty"`
	// service mesh (ASM) instance ID
	//
	// example:
	//
	// cb09fda0dc2f94a8397c76638c7ecf***
	MeshId *string `json:"MeshId,omitempty" xml:"MeshId,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterMeshConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterMeshConfig) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterMeshConfig) SetEnableMesh(v bool) *DescribeHubClusterDetailsResponseBodyClusterMeshConfig {
	s.EnableMesh = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterMeshConfig) SetMeshId(v string) *DescribeHubClusterDetailsResponseBodyClusterMeshConfig {
	s.MeshId = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterNetwork struct {
	// The domain name of the master instance.
	//
	// example:
	//
	// cluster.local
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The IP version that is supported by the master instance. Valid values:
	//
	// 	- ipv4: IPv4.
	//
	// 	- ipv6: IPv6.
	//
	// 	- dual: IPv4 and IPv6.
	//
	// example:
	//
	// ipv4
	IPStack *string `json:"IPStack,omitempty" xml:"IPStack,omitempty"`
	// The IDs of the associated security groups.
	SecurityGroupIDs []*string `json:"SecurityGroupIDs,omitempty" xml:"SecurityGroupIDs,omitempty" type:"Repeated"`
	// The details of the vSwitches.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) in which the master instance resides.
	//
	// example:
	//
	// vpc-f8ziib1019r9o0hdv2***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterNetwork) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterNetwork) SetClusterDomain(v string) *DescribeHubClusterDetailsResponseBodyClusterNetwork {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterNetwork) SetIPStack(v string) *DescribeHubClusterDetailsResponseBodyClusterNetwork {
	s.IPStack = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterNetwork) SetSecurityGroupIDs(v []*string) *DescribeHubClusterDetailsResponseBodyClusterNetwork {
	s.SecurityGroupIDs = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterNetwork) SetVSwitches(v []*string) *DescribeHubClusterDetailsResponseBodyClusterNetwork {
	s.VSwitches = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterNetwork) SetVpcId(v string) *DescribeHubClusterDetailsResponseBodyClusterNetwork {
	s.VpcId = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig struct {
	// Specifies whether to enable the workflow instance UI. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ArgoServerEnabled *bool `json:"ArgoServerEnabled,omitempty" xml:"ArgoServerEnabled,omitempty"`
	// The limit on the prices of containers in the workflow. This parameter takes effect only if the WorkflowScheduleMode parameter is set to cost-optimized.
	//
	// example:
	//
	// 0.08
	PriceLimit *string `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	// The scheduling mode of the workflow. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- cost-optimized: cost-prioritized scheduling mode.
	//
	// 	- stock-optimized: inventory-prioritized scheduling mode.
	//
	// example:
	//
	// cost-optimized
	WorkflowScheduleMode *string `json:"WorkflowScheduleMode,omitempty" xml:"WorkflowScheduleMode,omitempty"`
	// The Argo workflow regions  configuration.
	WorkflowUnits []*DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits `json:"WorkflowUnits,omitempty" xml:"WorkflowUnits,omitempty" type:"Repeated"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig) SetArgoServerEnabled(v bool) *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig {
	s.ArgoServerEnabled = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig) SetPriceLimit(v string) *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig {
	s.PriceLimit = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig) SetWorkflowScheduleMode(v string) *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig {
	s.WorkflowScheduleMode = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig) SetWorkflowUnits(v []*DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits) *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfig {
	s.WorkflowUnits = v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits struct {
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitches.
	VSwitches []*DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnitsVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-f8zukabbkv5aw7zkm****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits) SetRegionId(v string) *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits {
	s.RegionId = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits) SetVSwitches(v []*DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnitsVSwitches) *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits {
	s.VSwitches = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits) SetVpcId(v string) *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnits {
	s.VpcId = &v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnitsVSwitches struct {
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-wz9sf0hsuizl7bxnj****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnitsVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnitsVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnitsVSwitches) SetVswitchId(v string) *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnitsVSwitches {
	s.VswitchId = &v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnitsVSwitches) SetZoneId(v string) *DescribeHubClusterDetailsResponseBodyClusterWorkflowConfigWorkflowUnitsVSwitches {
	s.ZoneId = &v
	return s
}

type DescribeHubClusterDetailsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHubClusterDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHubClusterDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterDetailsResponse) SetHeaders(v map[string]*string) *DescribeHubClusterDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHubClusterDetailsResponse) SetStatusCode(v int32) *DescribeHubClusterDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHubClusterDetailsResponse) SetBody(v *DescribeHubClusterDetailsResponseBody) *DescribeHubClusterDetailsResponse {
	s.Body = v
	return s
}

type DescribeHubClusterKubeconfigRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c102fe5f1ee5d4c87a68121a77d8b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to obtain the kubeconfig file that is used to connect to the cluster over the internal network. Valid values:
	//
	// 	- `true`: obtains the kubeconfig file that is used to connect to the master instance over the internal network.
	//
	// 	- `false`: obtains the kubeconfig file that is used to connect to the master instance over the Internet.
	//
	// Default value: `false`
	//
	// example:
	//
	// false
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeHubClusterKubeconfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterKubeconfigRequest) SetClusterId(v string) *DescribeHubClusterKubeconfigRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHubClusterKubeconfigRequest) SetPrivateIpAddress(v bool) *DescribeHubClusterKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

type DescribeHubClusterKubeconfigResponseBody struct {
	// The content of the kubeconfig file.
	//
	// example:
	//
	// "\\napiVersion: v1\\nclusters:\\n- cluster:\\n    server: https://172.16.11.***:6443\\n    certificate-authority-data: LS0tLS1CRU=...\\n  name: kubernetes\\ncontexts:\\n- context:\\n    cluster: kubernetes\\n    user: \\"kubernetes-a****\\"\\n  name: kubernetes-admin-cc2cbf5802bec4bfa9fae4014d8c****\\ncurrent-context: kubernetes-admin-cc2cbf5802bec4bfa9fae4014d8c9****\\nkind: Config\\npreferences: {}\\nusers:\\n- name: \\"kubernetes-admin\\"\\n  user:\\n    client-certificate-data: LS0tLS1CRU...\\n    client-key-data: LS0tCg==...\\n"
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5BE4C329-DCC2-5B61-8F66-112B7D7FC712
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHubClusterKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterKubeconfigResponseBody) SetKubeconfig(v string) *DescribeHubClusterKubeconfigResponseBody {
	s.Kubeconfig = &v
	return s
}

func (s *DescribeHubClusterKubeconfigResponseBody) SetRequestId(v string) *DescribeHubClusterKubeconfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHubClusterKubeconfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHubClusterKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHubClusterKubeconfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterKubeconfigResponse) SetHeaders(v map[string]*string) *DescribeHubClusterKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeHubClusterKubeconfigResponse) SetStatusCode(v int32) *DescribeHubClusterKubeconfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHubClusterKubeconfigResponse) SetBody(v *DescribeHubClusterKubeconfigResponseBody) *DescribeHubClusterKubeconfigResponse {
	s.Body = v
	return s
}

type DescribeHubClusterLogsRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c9cb1d933b2ab47ff9cd25571477dc8f2
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeHubClusterLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterLogsRequest) SetClusterId(v string) *DescribeHubClusterLogsRequest {
	s.ClusterId = &v
	return s
}

type DescribeHubClusterLogsResponseBody struct {
	// Brief information about operation logs.
	Logs []*DescribeHubClusterLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 661192D7-25A6-54C2-B643-1129CB7D2768
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHubClusterLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterLogsResponseBody) SetLogs(v []*DescribeHubClusterLogsResponseBodyLogs) *DescribeHubClusterLogsResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeHubClusterLogsResponseBody) SetRequestId(v string) *DescribeHubClusterLogsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHubClusterLogsResponseBodyLogs struct {
	// The ID of the master instance.
	//
	// example:
	//
	// c102fe5f1ee5d4c87a68121a77d8b0f38
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// A log of the master instance.
	//
	// example:
	//
	// Cluster Created
	ClusterLog *string `json:"ClusterLog,omitempty" xml:"ClusterLog,omitempty"`
	// The time when the log was created. Format: <i>yyyy-mm-dd</i>t<i>hh:mm:ss</i>z (UTC time).
	//
	// example:
	//
	// 2021-12-02T11:48:15+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The severity level of the log. Valid values: - error: errors. - warn: warnings. - info: information.
	//
	// example:
	//
	// INFO
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
}

func (s DescribeHubClusterLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterLogsResponseBodyLogs) SetClusterId(v string) *DescribeHubClusterLogsResponseBodyLogs {
	s.ClusterId = &v
	return s
}

func (s *DescribeHubClusterLogsResponseBodyLogs) SetClusterLog(v string) *DescribeHubClusterLogsResponseBodyLogs {
	s.ClusterLog = &v
	return s
}

func (s *DescribeHubClusterLogsResponseBodyLogs) SetCreationTime(v string) *DescribeHubClusterLogsResponseBodyLogs {
	s.CreationTime = &v
	return s
}

func (s *DescribeHubClusterLogsResponseBodyLogs) SetLogLevel(v string) *DescribeHubClusterLogsResponseBodyLogs {
	s.LogLevel = &v
	return s
}

type DescribeHubClusterLogsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHubClusterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHubClusterLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHubClusterLogsResponse) SetHeaders(v map[string]*string) *DescribeHubClusterLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHubClusterLogsResponse) SetStatusCode(v int32) *DescribeHubClusterLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHubClusterLogsResponse) SetBody(v *DescribeHubClusterLogsResponseBody) *DescribeHubClusterLogsResponse {
	s.Body = v
	return s
}

type DescribeHubClustersRequest struct {
	// The configurations of the cluster.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz77rhypeu***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*Tag  `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeHubClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersRequest) SetProfile(v string) *DescribeHubClustersRequest {
	s.Profile = &v
	return s
}

func (s *DescribeHubClustersRequest) SetResourceGroupId(v string) *DescribeHubClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHubClustersRequest) SetTag(v []*Tag) *DescribeHubClustersRequest {
	s.Tag = v
	return s
}

type DescribeHubClustersShrinkRequest struct {
	// The configurations of the cluster.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz77rhypeu***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TagShrink       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeHubClustersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersShrinkRequest) SetProfile(v string) *DescribeHubClustersShrinkRequest {
	s.Profile = &v
	return s
}

func (s *DescribeHubClustersShrinkRequest) SetResourceGroupId(v string) *DescribeHubClustersShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHubClustersShrinkRequest) SetTagShrink(v string) *DescribeHubClustersShrinkRequest {
	s.TagShrink = &v
	return s
}

type DescribeHubClustersResponseBody struct {
	// The information about clusters.
	Clusters []*DescribeHubClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2D676EFC-8C04-5CCE-A08E-BB97D24B47E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHubClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBody) SetClusters(v []*DescribeHubClustersResponseBodyClusters) *DescribeHubClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeHubClustersResponseBody) SetRequestId(v string) *DescribeHubClustersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHubClustersResponseBodyClusters struct {
	// The information about the API server.
	ApiServer *DescribeHubClustersResponseBodyClustersApiServer `json:"ApiServer,omitempty" xml:"ApiServer,omitempty" type:"Struct"`
	// The details of the cluster.
	ClusterInfo *DescribeHubClustersResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// The deletion conditions of the cluster.
	Conditions []*DescribeHubClustersResponseBodyClustersConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The endpoint of the cluster.
	Endpoints *DescribeHubClustersResponseBodyClustersEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The logging configurations.
	LogConfig *DescribeHubClustersResponseBodyClustersLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// The configurations of Alibaba Cloud Service Mesh (ASM).
	MeshConfig *DescribeHubClustersResponseBodyClustersMeshConfig `json:"MeshConfig,omitempty" xml:"MeshConfig,omitempty" type:"Struct"`
	// The network configurations of the cluster.
	Network *DescribeHubClustersResponseBodyClustersNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
}

func (s DescribeHubClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBodyClusters) SetApiServer(v *DescribeHubClustersResponseBodyClustersApiServer) *DescribeHubClustersResponseBodyClusters {
	s.ApiServer = v
	return s
}

func (s *DescribeHubClustersResponseBodyClusters) SetClusterInfo(v *DescribeHubClustersResponseBodyClustersClusterInfo) *DescribeHubClustersResponseBodyClusters {
	s.ClusterInfo = v
	return s
}

func (s *DescribeHubClustersResponseBodyClusters) SetConditions(v []*DescribeHubClustersResponseBodyClustersConditions) *DescribeHubClustersResponseBodyClusters {
	s.Conditions = v
	return s
}

func (s *DescribeHubClustersResponseBodyClusters) SetEndpoints(v *DescribeHubClustersResponseBodyClustersEndpoints) *DescribeHubClustersResponseBodyClusters {
	s.Endpoints = v
	return s
}

func (s *DescribeHubClustersResponseBodyClusters) SetLogConfig(v *DescribeHubClustersResponseBodyClustersLogConfig) *DescribeHubClustersResponseBodyClusters {
	s.LogConfig = v
	return s
}

func (s *DescribeHubClustersResponseBodyClusters) SetMeshConfig(v *DescribeHubClustersResponseBodyClustersMeshConfig) *DescribeHubClustersResponseBodyClusters {
	s.MeshConfig = v
	return s
}

func (s *DescribeHubClustersResponseBodyClusters) SetNetwork(v *DescribeHubClustersResponseBodyClustersNetwork) *DescribeHubClustersResponseBodyClusters {
	s.Network = v
	return s
}

type DescribeHubClustersResponseBodyClustersApiServer struct {
	// The elastic IP address (EIP) ID.
	//
	// example:
	//
	// eip-xxx
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// Indicates whether public endpoint is enabled for the API server. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnabledPublic *bool `json:"EnabledPublic,omitempty" xml:"EnabledPublic,omitempty"`
	// The ID of the Server Load Balancer (SLB) instance that is associated with the cluster.
	//
	// example:
	//
	// lb-bp1qyp4l6bscqxw69****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DescribeHubClustersResponseBodyClustersApiServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClustersApiServer) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBodyClustersApiServer) SetApiServerEipId(v string) *DescribeHubClustersResponseBodyClustersApiServer {
	s.ApiServerEipId = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersApiServer) SetEnabledPublic(v bool) *DescribeHubClustersResponseBodyClustersApiServer {
	s.EnabledPublic = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersApiServer) SetLoadBalancerId(v string) *DescribeHubClustersResponseBodyClustersApiServer {
	s.LoadBalancerId = &v
	return s
}

type DescribeHubClustersResponseBodyClustersClusterInfo struct {
	// The cluster ID.
	//
	// example:
	//
	// c2d3e0121ea214b438010502a8019****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The specification of the cluster.
	//
	// 	- Only ack.pro.small is returned.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2021-11-05T10:25:48Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message that is returned if the cluster failed to be created.
	//
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ackone-heyuan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of the cluster.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of Resource Group.
	//
	// example:
	//
	// rg-qh2zgjsdv52***
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- initial: The cluster is being initialized.
	//
	// 	- failed: The cluster failed to be created.
	//
	// 	- running: The cluster is running
	//
	// 	- inactive: The cluster is pending.
	//
	// 	- deleting: The cluster is being deleted.
	//
	// 	- delete_failed: The cluster failed to be deleted.
	//
	// 	- deleted: The cluster is deleted.
	//
	// example:
	//
	// running
	State *string                                                   `json:"State,omitempty" xml:"State,omitempty"`
	Tags  []*DescribeHubClustersResponseBodyClustersClusterInfoTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the cluster was last updated.
	//
	// example:
	//
	// 2021-09-02T13:39:50Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The Kubernetes version of the cluster.
	//
	// example:
	//
	// 1.22.3-aliyun.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeHubClustersResponseBodyClustersClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClustersClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetClusterId(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetClusterSpec(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetCreationTime(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.CreationTime = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetErrorMessage(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetName(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.Name = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetProfile(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.Profile = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetRegionId(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetResourceGroupID(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.ResourceGroupID = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetState(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.State = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetTags(v []*DescribeHubClustersResponseBodyClustersClusterInfoTags) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.Tags = v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetUpdateTime(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetVersion(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.Version = &v
	return s
}

type DescribeHubClustersResponseBodyClustersClusterInfoTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHubClustersResponseBodyClustersClusterInfoTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClustersClusterInfoTags) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfoTags) SetKey(v string) *DescribeHubClustersResponseBodyClustersClusterInfoTags {
	s.Key = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersClusterInfoTags) SetValue(v string) *DescribeHubClustersResponseBodyClustersClusterInfoTags {
	s.Value = &v
	return s
}

type DescribeHubClustersResponseBodyClustersConditions struct {
	// The error message that is returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the deletion condition.
	//
	// example:
	//
	// Successful
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the cluster that the deletion condition indicates. Valid values:
	//
	// 	- True: The cluster cannot be deleted.
	//
	// 	- False: The cluster can be deleted.
	//
	// 	- Unknow: Whether the cluster can be deleted is unknown.
	//
	// example:
	//
	// True
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of deletion condition.
	//
	// example:
	//
	// DeletionProtection
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeHubClustersResponseBodyClustersConditions) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClustersConditions) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBodyClustersConditions) SetMessage(v string) *DescribeHubClustersResponseBodyClustersConditions {
	s.Message = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersConditions) SetReason(v string) *DescribeHubClustersResponseBodyClustersConditions {
	s.Reason = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersConditions) SetStatus(v string) *DescribeHubClustersResponseBodyClustersConditions {
	s.Status = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersConditions) SetType(v string) *DescribeHubClustersResponseBodyClustersConditions {
	s.Type = &v
	return s
}

type DescribeHubClustersResponseBodyClustersEndpoints struct {
	// The internal endpoint of the API server.
	//
	// example:
	//
	// https://172.16.6.**:6443
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	// The public endpoint of the API server.
	//
	// example:
	//
	// https://123.57.21.***:6443
	PublicApiServerEndpoint *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
}

func (s DescribeHubClustersResponseBodyClustersEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClustersEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBodyClustersEndpoints) SetIntranetApiServerEndpoint(v string) *DescribeHubClustersResponseBodyClustersEndpoints {
	s.IntranetApiServerEndpoint = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersEndpoints) SetPublicApiServerEndpoint(v string) *DescribeHubClustersResponseBodyClustersEndpoints {
	s.PublicApiServerEndpoint = &v
	return s
}

type DescribeHubClustersResponseBodyClustersLogConfig struct {
	// Indicates whether the audit logging feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// The name of the project in Simple Log Service.
	//
	// example:
	//
	// audit-log-abc
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The number of days that logs are retained by Simple Log Service.
	//
	// example:
	//
	// 7
	LogStoreTTL *string `json:"LogStoreTTL,omitempty" xml:"LogStoreTTL,omitempty"`
}

func (s DescribeHubClustersResponseBodyClustersLogConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClustersLogConfig) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBodyClustersLogConfig) SetEnableLog(v bool) *DescribeHubClustersResponseBodyClustersLogConfig {
	s.EnableLog = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersLogConfig) SetLogProject(v string) *DescribeHubClustersResponseBodyClustersLogConfig {
	s.LogProject = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersLogConfig) SetLogStoreTTL(v string) *DescribeHubClustersResponseBodyClustersLogConfig {
	s.LogStoreTTL = &v
	return s
}

type DescribeHubClustersResponseBodyClustersMeshConfig struct {
	// Indicates whether ASM is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableMesh *bool `json:"EnableMesh,omitempty" xml:"EnableMesh,omitempty"`
	// The ASM instance ID.
	//
	// example:
	//
	// c2d3e0121ea214b438010502a8019****
	MeshId *string `json:"MeshId,omitempty" xml:"MeshId,omitempty"`
}

func (s DescribeHubClustersResponseBodyClustersMeshConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClustersMeshConfig) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBodyClustersMeshConfig) SetEnableMesh(v bool) *DescribeHubClustersResponseBodyClustersMeshConfig {
	s.EnableMesh = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersMeshConfig) SetMeshId(v string) *DescribeHubClustersResponseBodyClustersMeshConfig {
	s.MeshId = &v
	return s
}

type DescribeHubClustersResponseBodyClustersNetwork struct {
	// The domain name of the cluster.
	//
	// example:
	//
	// cluster.local
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The security group IDs.
	SecurityGroupIDs []*string `json:"SecurityGroupIDs,omitempty" xml:"SecurityGroupIDs,omitempty" type:"Repeated"`
	// The IDs of vSwitches to which the cluster belongs.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the cluster belongs.
	//
	// example:
	//
	// vpc-2zeusrwi7c2mlww4a****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeHubClustersResponseBodyClustersNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClustersNetwork) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponseBodyClustersNetwork) SetClusterDomain(v string) *DescribeHubClustersResponseBodyClustersNetwork {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersNetwork) SetSecurityGroupIDs(v []*string) *DescribeHubClustersResponseBodyClustersNetwork {
	s.SecurityGroupIDs = v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersNetwork) SetVSwitches(v []*string) *DescribeHubClustersResponseBodyClustersNetwork {
	s.VSwitches = v
	return s
}

func (s *DescribeHubClustersResponseBodyClustersNetwork) SetVpcId(v string) *DescribeHubClustersResponseBodyClustersNetwork {
	s.VpcId = &v
	return s
}

type DescribeHubClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHubClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHubClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeHubClustersResponse) SetHeaders(v map[string]*string) *DescribeHubClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeHubClustersResponse) SetStatusCode(v int32) *DescribeHubClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHubClustersResponse) SetBody(v *DescribeHubClustersResponseBody) *DescribeHubClustersResponse {
	s.Body = v
	return s
}

type DescribeManagedClustersRequest struct {
	// The status of the association between the clusters and Service Mesh (ASM).
	//
	// This parameter is required.
	//
	// example:
	//
	// c2f41fd4599454a9c9ad8b3daafe873ad
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeManagedClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeManagedClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeManagedClustersRequest) SetClusterId(v string) *DescribeManagedClustersRequest {
	s.ClusterId = &v
	return s
}

type DescribeManagedClustersResponseBody struct {
	// The status of the associated clusters. Valid values: - initial: The associated clusters are being initialized. - failed: The associated clustersfailed to be created. - running: The associated clusters are running. - inactive: The associated clusters are inactive. - deleting: The associated clusters are being deleted. - deleted: The associated clusters are deleted.
	Clusters []*DescribeManagedClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// VPC ID
	//
	// example:
	//
	// BDA85C7A-FC81-56C4-9BC2-9112EE970059
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeManagedClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeManagedClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeManagedClustersResponseBody) SetClusters(v []*DescribeManagedClustersResponseBodyClusters) *DescribeManagedClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeManagedClustersResponseBody) SetRequestId(v string) *DescribeManagedClustersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeManagedClustersResponseBodyClusters struct {
	// The name of the master instance.
	Cluster *DescribeManagedClustersResponseBodyClustersCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	// Zhishi
	MeshStatus *DescribeManagedClustersResponseBodyClustersMeshStatus `json:"MeshStatus,omitempty" xml:"MeshStatus,omitempty" type:"Struct"`
	// Example 1
	Status *DescribeManagedClustersResponseBodyClustersStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s DescribeManagedClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeManagedClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeManagedClustersResponseBodyClusters) SetCluster(v *DescribeManagedClustersResponseBodyClustersCluster) *DescribeManagedClustersResponseBodyClusters {
	s.Cluster = v
	return s
}

func (s *DescribeManagedClustersResponseBodyClusters) SetMeshStatus(v *DescribeManagedClustersResponseBodyClustersMeshStatus) *DescribeManagedClustersResponseBodyClusters {
	s.MeshStatus = v
	return s
}

func (s *DescribeManagedClustersResponseBodyClusters) SetStatus(v *DescribeManagedClustersResponseBodyClustersStatus) *DescribeManagedClustersResponseBodyClusters {
	s.Status = v
	return s
}

type DescribeManagedClustersResponseBodyClustersCluster struct {
	// Information about the master instance.
	//
	// example:
	//
	// c2f41fd4599454a9c9ad8b3daafe873ad
	ClusterID *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	// The ID of the master instance.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The time when the master instance was created.
	//
	// example:
	//
	// One
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The ID of the master instance.
	//
	// example:
	//
	// 2022-03-23T06:22:28Z
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The name of the master instance.
	//
	// example:
	//
	// 1.22.3-aliyun.1
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The specification of the master instance. Valid values: - ack.pro.small: ACK Pro.
	//
	// example:
	//
	// 1.20.4-aliyun.1
	InitVersion *string `json:"InitVersion,omitempty" xml:"InitVersion,omitempty"`
	// The status information.
	//
	// example:
	//
	// ackone-heyuan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the master instance.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the master instance.
	//
	// example:
	//
	// rg-acfmx7o7ewyqcby
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The current Kubernetes version of the master instance.
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// 2022-03-23T06:22:28Z
	Updated *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	// The original Kubernetes version of the master instance.
	//
	// example:
	//
	// vsw-m5e0pbkgmhvzecf7enfym
	VSwitchID *string `json:"VSwitchID,omitempty" xml:"VSwitchID,omitempty"`
	// The status of the association between the clusters and the master instance. Valid values: - Installing: The clusters are being associated with the master instance. - Successed: The clusters are associated with the master instance. - Failed: The clusters failed to be associated with the master instance. - Deleting: The clusters are being disassociated from the master instance. - Deleted: The clusters are disassociated from the master instance.
	//
	// example:
	//
	// vpc-bp1fhizs9fxuvit06zeb9
	VpcID *string `json:"VpcID,omitempty" xml:"VpcID,omitempty"`
}

func (s DescribeManagedClustersResponseBodyClustersCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeManagedClustersResponseBodyClustersCluster) GoString() string {
	return s.String()
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetClusterID(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.ClusterID = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetClusterSpec(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetClusterType(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.ClusterType = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetCreated(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.Created = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetCurrentVersion(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetInitVersion(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.InitVersion = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetName(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.Name = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetProfile(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.Profile = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetRegion(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.Region = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetResourceGroupId(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetState(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.State = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetUpdated(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.Updated = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetVSwitchID(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.VSwitchID = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersCluster) SetVpcID(v string) *DescribeManagedClustersResponseBodyClustersCluster {
	s.VpcID = &v
	return s
}

type DescribeManagedClustersResponseBodyClustersMeshStatus struct {
	// example:
	//
	// true
	InMesh *bool `json:"InMesh,omitempty" xml:"InMesh,omitempty"`
}

func (s DescribeManagedClustersResponseBodyClustersMeshStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeManagedClustersResponseBodyClustersMeshStatus) GoString() string {
	return s.String()
}

func (s *DescribeManagedClustersResponseBodyClustersMeshStatus) SetInMesh(v bool) *DescribeManagedClustersResponseBodyClustersMeshStatus {
	s.InMesh = &v
	return s
}

type DescribeManagedClustersResponseBodyClustersStatus struct {
	// Query the clusters that are associated with a master instance.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// You can call the DescribeManagedClusters operation to query the clusters that are associated with a master instance.
	//
	// example:
	//
	// Successed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeManagedClustersResponseBodyClustersStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeManagedClustersResponseBodyClustersStatus) GoString() string {
	return s.String()
}

func (s *DescribeManagedClustersResponseBodyClustersStatus) SetMessage(v string) *DescribeManagedClustersResponseBodyClustersStatus {
	s.Message = &v
	return s
}

func (s *DescribeManagedClustersResponseBodyClustersStatus) SetState(v string) *DescribeManagedClustersResponseBodyClustersStatus {
	s.State = &v
	return s
}

type DescribeManagedClustersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeManagedClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeManagedClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeManagedClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeManagedClustersResponse) SetHeaders(v map[string]*string) *DescribeManagedClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeManagedClustersResponse) SetStatusCode(v int32) *DescribeManagedClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeManagedClustersResponse) SetBody(v *DescribeManagedClustersResponseBody) *DescribeManagedClustersResponse {
	s.Body = v
	return s
}

type DescribePoliciesResponseBody struct {
	// A list of policies.
	Policies []*DescribePoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9439169C-64C1-5849-9F7C-E3E60BDDEF7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePoliciesResponseBody) SetPolicies(v []*DescribePoliciesResponseBodyPolicies) *DescribePoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *DescribePoliciesResponseBody) SetRequestId(v string) *DescribePoliciesResponseBody {
	s.RequestId = &v
	return s
}

type DescribePoliciesResponseBodyPolicies struct {
	// The policy type.
	//
	// example:
	//
	// psp
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The names of the policies of each policy type.
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
}

func (s DescribePoliciesResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribePoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribePoliciesResponseBodyPolicies) SetCategory(v string) *DescribePoliciesResponseBodyPolicies {
	s.Category = &v
	return s
}

func (s *DescribePoliciesResponseBodyPolicies) SetNames(v []*string) *DescribePoliciesResponseBodyPolicies {
	s.Names = v
	return s
}

type DescribePoliciesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribePoliciesResponse) SetHeaders(v map[string]*string) *DescribePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribePoliciesResponse) SetStatusCode(v int32) *DescribePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePoliciesResponse) SetBody(v *DescribePoliciesResponseBody) *DescribePoliciesResponse {
	s.Body = v
	return s
}

type DescribePolicyDetailsRequest struct {
	// The policy name.
	//
	// example:
	//
	// ACKAllowedRepos
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DescribePolicyDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyDetailsRequest) SetPolicyName(v string) *DescribePolicyDetailsRequest {
	s.PolicyName = &v
	return s
}

type DescribePolicyDetailsResponseBody struct {
	// The policies.
	Policy *DescribePolicyDetailsResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2D676EFC-8C04-5CCE-A08E-BB97D24B47E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolicyDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyDetailsResponseBody) SetPolicy(v *DescribePolicyDetailsResponseBodyPolicy) *DescribePolicyDetailsResponseBody {
	s.Policy = v
	return s
}

func (s *DescribePolicyDetailsResponseBody) SetRequestId(v string) *DescribePolicyDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePolicyDetailsResponseBodyPolicy struct {
	// The action of the policy. Valid values:
	//
	// 	- enforce: blocks deployments that match the policy.
	//
	// 	- inform: generates alerts for deployments that match the policy.
	//
	// example:
	//
	// enforce
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// k8s-general
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the policy was created.
	//
	// example:
	//
	// 2021-11-18T10:52:17+08:00
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// Requires container images to begin with a repo string from a specified list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// ACKAllowedRepos
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether parameters are required. Valid values:
	//
	// 	- 0: Parameters are required.
	//
	// 	- 1: Parameters are not required.
	//
	// example:
	//
	// 0
	NoConfig *int32 `json:"NoConfig,omitempty" xml:"NoConfig,omitempty"`
	// The severity level of the policy.
	//
	// example:
	//
	// high
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The content of the policy.
	//
	// example:
	//
	// The content of the policy.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The time when the policy was last updated.
	//
	// example:
	//
	// 2021-11-18T10:52:17+08:00
	Updated *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s DescribePolicyDetailsResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyDetailsResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *DescribePolicyDetailsResponseBodyPolicy) SetAction(v string) *DescribePolicyDetailsResponseBodyPolicy {
	s.Action = &v
	return s
}

func (s *DescribePolicyDetailsResponseBodyPolicy) SetCategory(v string) *DescribePolicyDetailsResponseBodyPolicy {
	s.Category = &v
	return s
}

func (s *DescribePolicyDetailsResponseBodyPolicy) SetCreated(v string) *DescribePolicyDetailsResponseBodyPolicy {
	s.Created = &v
	return s
}

func (s *DescribePolicyDetailsResponseBodyPolicy) SetDescription(v string) *DescribePolicyDetailsResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *DescribePolicyDetailsResponseBodyPolicy) SetName(v string) *DescribePolicyDetailsResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *DescribePolicyDetailsResponseBodyPolicy) SetNoConfig(v int32) *DescribePolicyDetailsResponseBodyPolicy {
	s.NoConfig = &v
	return s
}

func (s *DescribePolicyDetailsResponseBodyPolicy) SetSeverity(v string) *DescribePolicyDetailsResponseBodyPolicy {
	s.Severity = &v
	return s
}

func (s *DescribePolicyDetailsResponseBodyPolicy) SetTemplate(v string) *DescribePolicyDetailsResponseBodyPolicy {
	s.Template = &v
	return s
}

func (s *DescribePolicyDetailsResponseBodyPolicy) SetUpdated(v string) *DescribePolicyDetailsResponseBodyPolicy {
	s.Updated = &v
	return s
}

type DescribePolicyDetailsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyDetailsResponse) SetHeaders(v map[string]*string) *DescribePolicyDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyDetailsResponse) SetStatusCode(v int32) *DescribePolicyDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyDetailsResponse) SetBody(v *DescribePolicyDetailsResponseBody) *DescribePolicyDetailsResponse {
	s.Body = v
	return s
}

type DescribePolicyGovernanceInClusterRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c09946603cd764dac96135f51d1ba****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribePolicyGovernanceInClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterRequest) SetClusterId(v string) *DescribePolicyGovernanceInClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribePolicyGovernanceInClusterResponseBody struct {
	// The detailed information about the policies.
	PolicyGovernances []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernances `json:"PolicyGovernances,omitempty" xml:"PolicyGovernances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 96C6A284-0EC3-5486-9A97-E8E9EE27E9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBody) SetPolicyGovernances(v []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernances) *DescribePolicyGovernanceInClusterResponseBody {
	s.PolicyGovernances = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBody) SetRequestId(v string) *DescribePolicyGovernanceInClusterResponseBody {
	s.RequestId = &v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernances struct {
	// The information about the associated clusters in which the policies are deployed.
	Cluster *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	// The detailed policy information.
	PolicyGovernance *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance `json:"PolicyGovernance,omitempty" xml:"PolicyGovernance,omitempty" type:"Struct"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernances) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernances) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernances) SetCluster(v *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernances {
	s.Cluster = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernances) SetPolicyGovernance(v *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernances {
	s.PolicyGovernance = v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster struct {
	// The ID of the associated cluster.
	//
	// example:
	//
	// ca5cf1b5edb5c4736a6ea0dfb4061****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The specifications of the associated cluster.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The type of the associated cluster.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The name of the associated cluster.
	//
	// example:
	//
	// ack-001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The identifier of the associated cluster.
	//
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The region ID of the associated cluster.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the associated cluster.
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) SetClusterId(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster {
	s.ClusterId = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) SetClusterSpec(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster {
	s.ClusterSpec = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) SetClusterType(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster {
	s.ClusterType = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) SetName(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster {
	s.Name = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) SetProfile(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster {
	s.Profile = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) SetRegionId(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster {
	s.RegionId = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster) SetState(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesCluster {
	s.State = &v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance struct {
	// The audit log generated by the associated cluster.
	AdmitLog *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog `json:"AdmitLog,omitempty" xml:"AdmitLog,omitempty" type:"Struct"`
	// The number of policies of each severity level enabled in the associated cluster.
	OnState []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState `json:"OnState,omitempty" xml:"OnState,omitempty" type:"Repeated"`
	// The number of deployments that match the policies in the associated cluster, including deployments that are blocked and deployments that have triggered alerting.
	Violation *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolation `json:"Violation,omitempty" xml:"Violation,omitempty" type:"Struct"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance) SetAdmitLog(v *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance {
	s.AdmitLog = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance) SetOnState(v []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance {
	s.OnState = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance) SetViolation(v *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolation) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernance {
	s.Violation = v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog struct {
	// The number of log entries in the query result.
	//
	// example:
	//
	// 100
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the Log Service project.
	//
	// example:
	//
	// demo
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// demo
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The content of the audit log.
	Logs []map[string]*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The status of the query. Valid values:
	//
	// 	- Complete: The query is successful, and the complete result is returned.
	//
	// 	- Incomplete: The query is successful, but the query result is incomplete. To obtain the complete result, you must call the operation again.
	//
	// example:
	//
	// Complete
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog) SetCount(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog {
	s.Count = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog) SetLogProject(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog {
	s.LogProject = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog) SetLogStore(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog {
	s.LogStore = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog) SetLogs(v []map[string]*string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog {
	s.Logs = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog) SetProgress(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceAdmitLog {
	s.Progress = &v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState struct {
	// The types of policies that are enabled in the associated cluster.
	//
	// example:
	//
	// 2
	EnabledCount *int64 `json:"EnabledCount,omitempty" xml:"EnabledCount,omitempty"`
	// The severity level.
	//
	// example:
	//
	// low
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The types of policies of each severity level.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState) SetEnabledCount(v int64) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState {
	s.EnabledCount = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState) SetTotalCount(v int64) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceOnState {
	s.TotalCount = &v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolation struct {
	// The number of deployments that match the policies in the associated cluster, including deployments that are blocked and deployments that have triggered alerting. The deployments are classified by severity level.
	TotalViolations *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolations `json:"TotalViolations,omitempty" xml:"TotalViolations,omitempty" type:"Struct"`
	// The number of deployments that match the policies in the associated cluster, including deployments that are blocked and deployments that have triggered alerting. The deployments are classified by policy type.
	Violations *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolations `json:"Violations,omitempty" xml:"Violations,omitempty" type:"Struct"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolation) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolation) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolation) SetTotalViolations(v *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolations) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolation {
	s.TotalViolations = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolation) SetViolations(v *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolations) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolation {
	s.Violations = v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolations struct {
	// The information about the deployments that are blocked.
	Deny []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsDeny `json:"Deny,omitempty" xml:"Deny,omitempty" type:"Repeated"`
	// The information about the deployments that have triggered alerting.
	Warn []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsWarn `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Repeated"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolations) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolations) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolations) SetDeny(v []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsDeny) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolations {
	s.Deny = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolations) SetWarn(v []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsWarn) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolations {
	s.Warn = v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsDeny struct {
	// The severity level.
	//
	// example:
	//
	// low
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The number of deployments that are blocked.
	//
	// example:
	//
	// 2
	Violations *int64 `json:"Violations,omitempty" xml:"Violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsDeny) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsDeny) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsDeny) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsDeny {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsDeny) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsDeny {
	s.Violations = &v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsWarn struct {
	// The severity level.
	//
	// example:
	//
	// low
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The number of deployments that have triggered alerting.
	//
	// example:
	//
	// 2
	Violations *string `json:"Violations,omitempty" xml:"Violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsWarn) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsWarn) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsWarn) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsWarn {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsWarn) SetViolations(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationTotalViolationsWarn {
	s.Violations = &v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolations struct {
	// The information about the deployments that are blocked.
	Deny []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny `json:"Deny,omitempty" xml:"Deny,omitempty" type:"Repeated"`
	// The information about the deployments that have triggered alerting.
	Warn []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Repeated"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolations) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolations) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolations) SetDeny(v []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolations {
	s.Deny = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolations) SetWarn(v []*DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolations {
	s.Warn = v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny struct {
	// The description of the policy.
	//
	// example:
	//
	// Restricts secrets used in pod envs
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// ACKPSPCapabilities
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The severity level of the policy.
	//
	// example:
	//
	// low
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The number of times that the policy blocks deployments.
	//
	// example:
	//
	// 2
	Violations *int64 `json:"Violations,omitempty" xml:"Violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny) SetPolicyDescription(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny) SetPolicyName(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsDeny {
	s.Violations = &v
	return s
}

type DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn struct {
	// The description of the policy.
	//
	// example:
	//
	// Restricts secrets used in pod envs
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// ACKPSPCapabilities
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The severity level of the policy.
	//
	// example:
	//
	// low
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The number of times that the policy generates alerts.
	//
	// example:
	//
	// 2
	Violations *int64 `json:"Violations,omitempty" xml:"Violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn) SetPolicyDescription(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn) SetPolicyName(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyPolicyGovernancesPolicyGovernanceViolationViolationsWarn {
	s.Violations = &v
	return s
}

type DescribePolicyGovernanceInClusterResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyGovernanceInClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponse) SetHeaders(v map[string]*string) *DescribePolicyGovernanceInClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponse) SetStatusCode(v int32) *DescribePolicyGovernanceInClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponse) SetBody(v *DescribePolicyGovernanceInClusterResponseBody) *DescribePolicyGovernanceInClusterResponse {
	s.Body = v
	return s
}

type DescribePolicyInstancesRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c102fe5f1ee5d4c87a68121a77d8b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACKNoEnvVarSecrets
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DescribePolicyInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesRequest) SetClusterId(v string) *DescribePolicyInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribePolicyInstancesRequest) SetPolicyName(v string) *DescribePolicyInstancesRequest {
	s.PolicyName = &v
	return s
}

type DescribePolicyInstancesResponseBody struct {
	// A list of policy instances.
	Policies []*DescribePolicyInstancesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5BE4C329-DCC2-5B61-8F66-112B7D7FC712
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolicyInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesResponseBody) SetPolicies(v []*DescribePolicyInstancesResponseBodyPolicies) *DescribePolicyInstancesResponseBody {
	s.Policies = v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetRequestId(v string) *DescribePolicyInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DescribePolicyInstancesResponseBodyPolicies struct {
	// The ID of the associated cluster.
	//
	// example:
	//
	// cd0e6882394f7496589837cac3585****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the policy instance.
	//
	// example:
	//
	// no-env-var-secrets-****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The action of the policy. Valid values:
	//
	// 	- deny: blocks deployments that match the policy.
	//
	// 	- warn: generates alerts for deployments that match the policy.
	//
	// example:
	//
	// warn
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// k8s-general
	PolicyCategory *string `json:"PolicyCategory,omitempty" xml:"PolicyCategory,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// Restricts secrets used in pod envs
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// ACKPSPCapabilities
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The parameters of the policy instance.
	PolicyParameters map[string]*string `json:"PolicyParameters,omitempty" xml:"PolicyParameters,omitempty"`
	// The applicable scope of the policy instance.
	//
	// A value of \\	- indicates all namespaces. This is the default value.
	//
	// Multiple namespaces are separated by commas (,).
	//
	// example:
	//
	// *
	PolicyScope *string `json:"PolicyScope,omitempty" xml:"PolicyScope,omitempty"`
	// The severity level of the policy.
	//
	// example:
	//
	// low
	PolicySeverity *string `json:"PolicySeverity,omitempty" xml:"PolicySeverity,omitempty"`
	// The total number of deployments that match the policy in the associated cluster, including the deployments that are blocked and the deployments that have triggered alerting.
	//
	// example:
	//
	// 2
	TotalViolations *int64 `json:"TotalViolations,omitempty" xml:"TotalViolations,omitempty"`
}

func (s DescribePolicyInstancesResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetClusterId(v string) *DescribePolicyInstancesResponseBodyPolicies {
	s.ClusterId = &v
	return s
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetInstanceName(v string) *DescribePolicyInstancesResponseBodyPolicies {
	s.InstanceName = &v
	return s
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetPolicyAction(v string) *DescribePolicyInstancesResponseBodyPolicies {
	s.PolicyAction = &v
	return s
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetPolicyCategory(v string) *DescribePolicyInstancesResponseBodyPolicies {
	s.PolicyCategory = &v
	return s
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetPolicyDescription(v string) *DescribePolicyInstancesResponseBodyPolicies {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetPolicyName(v string) *DescribePolicyInstancesResponseBodyPolicies {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetPolicyParameters(v map[string]*string) *DescribePolicyInstancesResponseBodyPolicies {
	s.PolicyParameters = v
	return s
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetPolicyScope(v string) *DescribePolicyInstancesResponseBodyPolicies {
	s.PolicyScope = &v
	return s
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetPolicySeverity(v string) *DescribePolicyInstancesResponseBodyPolicies {
	s.PolicySeverity = &v
	return s
}

func (s *DescribePolicyInstancesResponseBodyPolicies) SetTotalViolations(v int64) *DescribePolicyInstancesResponseBodyPolicies {
	s.TotalViolations = &v
	return s
}

type DescribePolicyInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesResponse) SetHeaders(v map[string]*string) *DescribePolicyInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyInstancesResponse) SetStatusCode(v int32) *DescribePolicyInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyInstancesResponse) SetBody(v *DescribePolicyInstancesResponseBody) *DescribePolicyInstancesResponse {
	s.Body = v
	return s
}

type DescribePolicyInstancesStatusRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c676decda7b8148d69a9aba751877****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribePolicyInstancesStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusRequest) SetClusterId(v string) *DescribePolicyInstancesStatusRequest {
	s.ClusterId = &v
	return s
}

type DescribePolicyInstancesStatusResponseBody struct {
	// The number of policy instances of each policy type.
	Policies *DescribePolicyInstancesStatusResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5D89C59A-A7EB-5BF8-B094-6479175346CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolicyInstancesStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusResponseBody) SetPolicies(v *DescribePolicyInstancesStatusResponseBodyPolicies) *DescribePolicyInstancesStatusResponseBody {
	s.Policies = v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBody) SetRequestId(v string) *DescribePolicyInstancesStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribePolicyInstancesStatusResponseBodyPolicies struct {
	// The number of policy instances of each policy type.
	PolicyInstances []*DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances `json:"PolicyInstances,omitempty" xml:"PolicyInstances,omitempty" type:"Repeated"`
	// The number of policy instances that are deployed in the cluster.
	SeverityInfo []*DescribePolicyInstancesStatusResponseBodyPoliciesSeverityInfo `json:"SeverityInfo,omitempty" xml:"SeverityInfo,omitempty" type:"Repeated"`
}

func (s DescribePolicyInstancesStatusResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesStatusResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicies) SetPolicyInstances(v []*DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances) *DescribePolicyInstancesStatusResponseBodyPolicies {
	s.PolicyInstances = v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicies) SetSeverityInfo(v []*DescribePolicyInstancesStatusResponseBodyPoliciesSeverityInfo) *DescribePolicyInstancesStatusResponseBodyPolicies {
	s.SeverityInfo = v
	return s
}

type DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances struct {
	// The type of the policy.
	//
	// example:
	//
	// compliance
	PolicyCategory *string `json:"PolicyCategory,omitempty" xml:"PolicyCategory,omitempty"`
	// The associated clusters in which the policy instances are deployed.
	PolicyClusters []*DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstancesPolicyClusters `json:"PolicyClusters,omitempty" xml:"PolicyClusters,omitempty" type:"Repeated"`
	// The description of the policy.
	//
	// example:
	//
	// Restricts use of the cluster-admin role.
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The number of policy instances that are deployed. If this parameter is empty, no policy instance is deployed.
	//
	// example:
	//
	// 2
	PolicyInstancesCount *int64 `json:"PolicyInstancesCount,omitempty" xml:"PolicyInstancesCount,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// ACKRestrictRoleBindings
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The severity level of the policy.
	//
	// example:
	//
	// low
	PolicySeverity *string `json:"PolicySeverity,omitempty" xml:"PolicySeverity,omitempty"`
}

func (s DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances) SetPolicyCategory(v string) *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances {
	s.PolicyCategory = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances) SetPolicyClusters(v []*DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstancesPolicyClusters) *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances {
	s.PolicyClusters = v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances) SetPolicyDescription(v string) *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances) SetPolicyInstancesCount(v int64) *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances {
	s.PolicyInstancesCount = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances) SetPolicyName(v string) *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances) SetPolicySeverity(v string) *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstances {
	s.PolicySeverity = &v
	return s
}

type DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstancesPolicyClusters struct {
	// The ID of the associated cluster.
	//
	// example:
	//
	// c639e5310e73e4a29ab18d6330a633****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The status of the deployment.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstancesPolicyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstancesPolicyClusters) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstancesPolicyClusters) SetClusterId(v string) *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstancesPolicyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstancesPolicyClusters) SetStatus(v string) *DescribePolicyInstancesStatusResponseBodyPoliciesPolicyInstancesPolicyClusters {
	s.Status = &v
	return s
}

type DescribePolicyInstancesStatusResponseBodyPoliciesSeverityInfo struct {
	// The number of policy instances.
	//
	// example:
	//
	// 2
	SeverityCount *string `json:"SeverityCount,omitempty" xml:"SeverityCount,omitempty"`
	// The severity level.
	//
	// example:
	//
	// low
	SeverityType *string `json:"SeverityType,omitempty" xml:"SeverityType,omitempty"`
}

func (s DescribePolicyInstancesStatusResponseBodyPoliciesSeverityInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesStatusResponseBodyPoliciesSeverityInfo) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesSeverityInfo) SetSeverityCount(v string) *DescribePolicyInstancesStatusResponseBodyPoliciesSeverityInfo {
	s.SeverityCount = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPoliciesSeverityInfo) SetSeverityType(v string) *DescribePolicyInstancesStatusResponseBodyPoliciesSeverityInfo {
	s.SeverityType = &v
	return s
}

type DescribePolicyInstancesStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyInstancesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyInstancesStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyInstancesStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusResponse) SetHeaders(v map[string]*string) *DescribePolicyInstancesStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyInstancesStatusResponse) SetStatusCode(v int32) *DescribePolicyInstancesStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponse) SetBody(v *DescribePolicyInstancesStatusResponseBody) *DescribePolicyInstancesStatusResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The language. Valid values: zh and en.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetLanguage(v string) *DescribeRegionsRequest {
	s.Language = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// A list of available regions that are returned.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C0EE05F4-6C1D-5993-B028-B569F9ED6B51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The name of the region.
	//
	// example:
	//
	// China (Beijing)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeUserPermissionsRequest struct {
	// The ID of the RAM user that you want to query.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 21175****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeUserPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsRequest) SetUserId(v string) *DescribeUserPermissionsRequest {
	s.UserId = &v
	return s
}

type DescribeUserPermissionsResponseBody struct {
	// The details about the permissions of the RAM user.
	Permissions []*DescribeUserPermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// EA06613B-37A3-549E-BAE0-E4AD8A6E93D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsResponseBody) SetPermissions(v []*DescribeUserPermissionsResponseBodyPermissions) *DescribeUserPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *DescribeUserPermissionsResponseBody) SetRequestId(v string) *DescribeUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserPermissionsResponseBodyPermissions struct {
	// The authorization setting. Valid values:
	//
	// 	- {cluster_id} is returned if the permissions are scoped to a cluster.
	//
	// 	- {cluster_id}/{namespace} is returned if the permissions are scoped to a namespace of a cluster.
	//
	// 	- all-clusters is returned if the permissions are scoped to all clusters.
	//
	// example:
	//
	// cffef3c9c7ba145b083292942a2c3****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The authorization type. Valid values:
	//
	// 	- cluster: indicates that the permissions are scoped to a cluster.
	//
	// 	- namespace: indicates that the permissions are scoped to a namespace of a cluster.
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The name of the custom role. If a custom role is assigned, the value is the name of the assigned custom role.
	//
	// example:
	//
	// view
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The type of predefined role. Valid values:
	//
	// 	- admin: administrator
	//
	// 	- dev: developer
	//
	// example:
	//
	// dev
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeUserPermissionsResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetResourceId(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.ResourceId = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetResourceType(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.ResourceType = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetRoleName(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.RoleName = &v
	return s
}

func (s *DescribeUserPermissionsResponseBodyPermissions) SetRoleType(v string) *DescribeUserPermissionsResponseBodyPermissions {
	s.RoleType = &v
	return s
}

type DescribeUserPermissionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsResponse) SetHeaders(v map[string]*string) *DescribeUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserPermissionsResponse) SetStatusCode(v int32) *DescribeUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserPermissionsResponse) SetBody(v *DescribeUserPermissionsResponseBody) *DescribeUserPermissionsResponse {
	s.Body = v
	return s
}

type DetachClusterFromHubRequest struct {
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb09fda0dc2f94a8397c76638c7ec****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The operation that you want to perform. Set the value to **DetachClusterFromHub**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["c1c731554c1ec4a1ca9bb690ff9ed****"]
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// Example 1
	//
	// example:
	//
	// false
	DetachFromMesh *bool `json:"DetachFromMesh,omitempty" xml:"DetachFromMesh,omitempty"`
}

func (s DetachClusterFromHubRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachClusterFromHubRequest) GoString() string {
	return s.String()
}

func (s *DetachClusterFromHubRequest) SetClusterId(v string) *DetachClusterFromHubRequest {
	s.ClusterId = &v
	return s
}

func (s *DetachClusterFromHubRequest) SetClusterIds(v string) *DetachClusterFromHubRequest {
	s.ClusterIds = &v
	return s
}

func (s *DetachClusterFromHubRequest) SetDetachFromMesh(v bool) *DetachClusterFromHubRequest {
	s.DetachFromMesh = &v
	return s
}

type DetachClusterFromHubResponseBody struct {
	// Zhishi
	//
	// example:
	//
	// cc490b1e67ccc43a784727f29f33e****
	ClusterId         *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ManagedClusterIds []*string `json:"ManagedClusterIds,omitempty" xml:"ManagedClusterIds,omitempty" type:"Repeated"`
	// You can call the DetachClusterFromHub operation to disassociate clusters from a master instance.
	//
	// example:
	//
	// 4412F213-DBCD-5D1B-A9A1-F6C26C6C19D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// T-623a96b7bbeaac074b00****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetachClusterFromHubResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachClusterFromHubResponseBody) GoString() string {
	return s.String()
}

func (s *DetachClusterFromHubResponseBody) SetClusterId(v string) *DetachClusterFromHubResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DetachClusterFromHubResponseBody) SetManagedClusterIds(v []*string) *DetachClusterFromHubResponseBody {
	s.ManagedClusterIds = v
	return s
}

func (s *DetachClusterFromHubResponseBody) SetRequestId(v string) *DetachClusterFromHubResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachClusterFromHubResponseBody) SetTaskId(v string) *DetachClusterFromHubResponseBody {
	s.TaskId = &v
	return s
}

type DetachClusterFromHubResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachClusterFromHubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachClusterFromHubResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachClusterFromHubResponse) GoString() string {
	return s.String()
}

func (s *DetachClusterFromHubResponse) SetHeaders(v map[string]*string) *DetachClusterFromHubResponse {
	s.Headers = v
	return s
}

func (s *DetachClusterFromHubResponse) SetStatusCode(v int32) *DetachClusterFromHubResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachClusterFromHubResponse) SetBody(v *DetachClusterFromHubResponseBody) *DetachClusterFromHubResponse {
	s.Body = v
	return s
}

type GrantUserPermissionRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c102fe5f1ee5d4c87a68121a77d8b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The entity to which the permissions are granted. A value of `true` indicates that the permissions are granted to a RAM user. A value of `false` indicates that the permissions are granted to a RAM role.
	//
	// example:
	//
	// false
	IsRamRole *bool `json:"IsRamRole,omitempty" xml:"IsRamRole,omitempty"`
	// The namespace to which the permissions are scoped. By default, this parameter is empty when you set RoleType to cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The predefined role that you want to assign. Valid values:
	//
	// 	- admin: the administrator role.
	//
	// 	- dev: the developer role.
	//
	// This parameter is required.
	//
	// example:
	//
	// admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The authorization type. Valid values:
	//
	// 	- cluster: specifies that the permissions are scoped to a master instance.
	//
	// 	- namespace: specifies that the permissions are scoped to a namespace of a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The ID of the RAM user or RAM role.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2176****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GrantUserPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionRequest) SetClusterId(v string) *GrantUserPermissionRequest {
	s.ClusterId = &v
	return s
}

func (s *GrantUserPermissionRequest) SetIsRamRole(v bool) *GrantUserPermissionRequest {
	s.IsRamRole = &v
	return s
}

func (s *GrantUserPermissionRequest) SetNamespace(v string) *GrantUserPermissionRequest {
	s.Namespace = &v
	return s
}

func (s *GrantUserPermissionRequest) SetRoleName(v string) *GrantUserPermissionRequest {
	s.RoleName = &v
	return s
}

func (s *GrantUserPermissionRequest) SetRoleType(v string) *GrantUserPermissionRequest {
	s.RoleType = &v
	return s
}

func (s *GrantUserPermissionRequest) SetUserId(v string) *GrantUserPermissionRequest {
	s.UserId = &v
	return s
}

type GrantUserPermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2D676EFC-8C04-5CCE-A08E-BB97D24B47E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantUserPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionResponseBody) SetRequestId(v string) *GrantUserPermissionResponseBody {
	s.RequestId = &v
	return s
}

type GrantUserPermissionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantUserPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantUserPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionResponse) SetHeaders(v map[string]*string) *GrantUserPermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantUserPermissionResponse) SetStatusCode(v int32) *GrantUserPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantUserPermissionResponse) SetBody(v *GrantUserPermissionResponseBody) *GrantUserPermissionResponse {
	s.Body = v
	return s
}

type GrantUserPermissionsRequest struct {
	// The list of permissions that you want to grant to the RAM user.
	Permissions []*GrantUserPermissionsRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The ID of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2367****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GrantUserPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsRequest) SetPermissions(v []*GrantUserPermissionsRequestPermissions) *GrantUserPermissionsRequest {
	s.Permissions = v
	return s
}

func (s *GrantUserPermissionsRequest) SetUserId(v string) *GrantUserPermissionsRequest {
	s.UserId = &v
	return s
}

type GrantUserPermissionsRequestPermissions struct {
	// The master instance ID.
	//
	// 	- When the role_type parameter is set to all-clusters, set the parameter to an empty string.
	//
	// example:
	//
	// cf67bdb0ffcb349ecabc1ca70da78****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The entity to which the permissions are granted. A value of `true` indicates that the permissions are granted to a RAM user. A value of `false` indicates that the permissions are granted to a RAM role.
	//
	// example:
	//
	// true
	IsRamRole *bool `json:"IsRamRole,omitempty" xml:"IsRamRole,omitempty"`
	// The namespace to which the permissions are scoped. By default, this parameter is empty when you set RoleType to cluster.
	//
	// example:
	//
	// test
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The predefined role that you want to assign. Valid values:
	//
	// 	- admin: the administrator role.
	//
	// 	- dev: the developer role.
	//
	// This parameter is required.
	//
	// example:
	//
	// dev
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The authorization type. Valid values:
	//
	// 	- cluster: specifies that the permissions are scoped to a master instance.
	//
	// 	- namespace: specifies that the permissions are scoped to a namespace of a cluster.
	//
	// 	- all-clusters: specifies that the permissions are scoped to all master instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s GrantUserPermissionsRequestPermissions) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionsRequestPermissions) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsRequestPermissions) SetClusterId(v string) *GrantUserPermissionsRequestPermissions {
	s.ClusterId = &v
	return s
}

func (s *GrantUserPermissionsRequestPermissions) SetIsRamRole(v bool) *GrantUserPermissionsRequestPermissions {
	s.IsRamRole = &v
	return s
}

func (s *GrantUserPermissionsRequestPermissions) SetNamespace(v string) *GrantUserPermissionsRequestPermissions {
	s.Namespace = &v
	return s
}

func (s *GrantUserPermissionsRequestPermissions) SetRoleName(v string) *GrantUserPermissionsRequestPermissions {
	s.RoleName = &v
	return s
}

func (s *GrantUserPermissionsRequestPermissions) SetRoleType(v string) *GrantUserPermissionsRequestPermissions {
	s.RoleType = &v
	return s
}

type GrantUserPermissionsShrinkRequest struct {
	// The list of permissions that you want to grant to the RAM user.
	PermissionsShrink *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The ID of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2367****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GrantUserPermissionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsShrinkRequest) SetPermissionsShrink(v string) *GrantUserPermissionsShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *GrantUserPermissionsShrinkRequest) SetUserId(v string) *GrantUserPermissionsShrinkRequest {
	s.UserId = &v
	return s
}

type GrantUserPermissionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4412F213-DBCD-5D1B-A9A1-F6C26C6C19D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantUserPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsResponseBody) SetRequestId(v string) *GrantUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

type GrantUserPermissionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantUserPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsResponse) SetHeaders(v map[string]*string) *GrantUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *GrantUserPermissionsResponse) SetStatusCode(v int32) *GrantUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantUserPermissionsResponse) SetBody(v *GrantUserPermissionsResponseBody) *GrantUserPermissionsResponse {
	s.Body = v
	return s
}

type UpdateHubClusterFeatureRequest struct {
	// The Internet access control list (ACL). This parameter takes effect only if PublicAccessEnabled is set to true.
	AccessControlList []*string `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty" type:"Repeated"`
	// The ID of the EIP.
	//
	// example:
	//
	// eip-xxx
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// Specifies whether to enable Argo CD. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ArgoCDEnabled *bool `json:"ArgoCDEnabled,omitempty" xml:"ArgoCDEnabled,omitempty"`
	// Specifies whether to enable high availability for Argo CD. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ArgoCDHAEnabled *bool `json:"ArgoCDHAEnabled,omitempty" xml:"ArgoCDHAEnabled,omitempty"`
	// Specifies whether to enable ArgoEvents. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	ArgoEventsEnabled *bool `json:"ArgoEventsEnabled,omitempty" xml:"ArgoEventsEnabled,omitempty"`
	// Specifies whether to enable the workflow instance UI. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ArgoServerEnabled *bool `json:"ArgoServerEnabled,omitempty" xml:"ArgoServerEnabled,omitempty"`
	// Specifies whether to enable the audit logging feature. Valid values:
	//
	// 	- true: enables the audit logging feature.
	//
	// 	- false: disables the audit logging feature.
	//
	// example:
	//
	// true
	AuditLogEnabled *bool `json:"AuditLogEnabled,omitempty" xml:"AuditLogEnabled,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c46979b1075f04d99b5f2b710393e5****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable the deletion protection feature for the cluster. After you enable the deletion protection feature for the cluster, you cannot delete the cluster in the console or by calling the DeleteHubCluster operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to enable Service Mesh (ASM). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableMesh *bool `json:"EnableMesh,omitempty" xml:"EnableMesh,omitempty"`
	// Specifies whether to enable Gateway. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	GatewayEnabled *bool `json:"GatewayEnabled,omitempty" xml:"GatewayEnabled,omitempty"`
	// Specifies whether to enable the monitoring dashboard feature for the workflow instance. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	MonitorEnabled *bool `json:"MonitorEnabled,omitempty" xml:"MonitorEnabled,omitempty"`
	// The name of the master instance. The name must be 1 to 63 characters in length. It must start with a letter, and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// ack-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The limit on the prices of containers in the workflow. This parameter takes effect only if the WorkflowScheduleMode parameter is set to cost-optimized.
	//
	// example:
	//
	// 0.08
	PriceLimit *string `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	// Specifies whether to enable public domain name resolution in the Argo CD or Argo Workflow console. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" xml:"PublicAccessEnabled,omitempty"`
	// Specifies whether to associate an elastic IP address (EIP) with the API server. Valid values:
	//
	// 	- true: associates an EIP with the API server. You can specify ApiServerEipId. If you do not specify ApiServerEipId, the system automatically creates an EIP.
	//
	// 	- false: disassociates an EIP from the API server.
	//
	// example:
	//
	// true
	PublicApiServerEnabled *bool `json:"PublicApiServerEnabled,omitempty" xml:"PublicApiServerEnabled,omitempty"`
	// The vSwitches.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The scheduling mode of the workflow. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- cost-optimized: cost-prioritized scheduling mode.
	//
	// 	- stock-optimized: inventory-prioritized scheduling mode.
	//
	// example:
	//
	// cost-optimized
	WorkflowScheduleMode *string `json:"WorkflowScheduleMode,omitempty" xml:"WorkflowScheduleMode,omitempty"`
}

func (s UpdateHubClusterFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHubClusterFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateHubClusterFeatureRequest) SetAccessControlList(v []*string) *UpdateHubClusterFeatureRequest {
	s.AccessControlList = v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetApiServerEipId(v string) *UpdateHubClusterFeatureRequest {
	s.ApiServerEipId = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetArgoCDEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.ArgoCDEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetArgoCDHAEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.ArgoCDHAEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetArgoEventsEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.ArgoEventsEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetArgoServerEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.ArgoServerEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetAuditLogEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.AuditLogEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetClusterId(v string) *UpdateHubClusterFeatureRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetDeletionProtection(v bool) *UpdateHubClusterFeatureRequest {
	s.DeletionProtection = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetEnableMesh(v bool) *UpdateHubClusterFeatureRequest {
	s.EnableMesh = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetGatewayEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.GatewayEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetMonitorEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.MonitorEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetName(v string) *UpdateHubClusterFeatureRequest {
	s.Name = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetPriceLimit(v string) *UpdateHubClusterFeatureRequest {
	s.PriceLimit = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetPublicAccessEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.PublicAccessEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetPublicApiServerEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.PublicApiServerEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetVSwitches(v []*string) *UpdateHubClusterFeatureRequest {
	s.VSwitches = v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetWorkflowScheduleMode(v string) *UpdateHubClusterFeatureRequest {
	s.WorkflowScheduleMode = &v
	return s
}

type UpdateHubClusterFeatureShrinkRequest struct {
	// The Internet access control list (ACL). This parameter takes effect only if PublicAccessEnabled is set to true.
	AccessControlListShrink *string `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty"`
	// The ID of the EIP.
	//
	// example:
	//
	// eip-xxx
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// Specifies whether to enable Argo CD. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ArgoCDEnabled *bool `json:"ArgoCDEnabled,omitempty" xml:"ArgoCDEnabled,omitempty"`
	// Specifies whether to enable high availability for Argo CD. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ArgoCDHAEnabled *bool `json:"ArgoCDHAEnabled,omitempty" xml:"ArgoCDHAEnabled,omitempty"`
	// Specifies whether to enable ArgoEvents. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	ArgoEventsEnabled *bool `json:"ArgoEventsEnabled,omitempty" xml:"ArgoEventsEnabled,omitempty"`
	// Specifies whether to enable the workflow instance UI. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ArgoServerEnabled *bool `json:"ArgoServerEnabled,omitempty" xml:"ArgoServerEnabled,omitempty"`
	// Specifies whether to enable the audit logging feature. Valid values:
	//
	// 	- true: enables the audit logging feature.
	//
	// 	- false: disables the audit logging feature.
	//
	// example:
	//
	// true
	AuditLogEnabled *bool `json:"AuditLogEnabled,omitempty" xml:"AuditLogEnabled,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c46979b1075f04d99b5f2b710393e5****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable the deletion protection feature for the cluster. After you enable the deletion protection feature for the cluster, you cannot delete the cluster in the console or by calling the DeleteHubCluster operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to enable Service Mesh (ASM). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableMesh *bool `json:"EnableMesh,omitempty" xml:"EnableMesh,omitempty"`
	// Specifies whether to enable Gateway. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	GatewayEnabled *bool `json:"GatewayEnabled,omitempty" xml:"GatewayEnabled,omitempty"`
	// Specifies whether to enable the monitoring dashboard feature for the workflow instance. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	MonitorEnabled *bool `json:"MonitorEnabled,omitempty" xml:"MonitorEnabled,omitempty"`
	// The name of the master instance. The name must be 1 to 63 characters in length. It must start with a letter, and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// ack-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The limit on the prices of containers in the workflow. This parameter takes effect only if the WorkflowScheduleMode parameter is set to cost-optimized.
	//
	// example:
	//
	// 0.08
	PriceLimit *string `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	// Specifies whether to enable public domain name resolution in the Argo CD or Argo Workflow console. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" xml:"PublicAccessEnabled,omitempty"`
	// Specifies whether to associate an elastic IP address (EIP) with the API server. Valid values:
	//
	// 	- true: associates an EIP with the API server. You can specify ApiServerEipId. If you do not specify ApiServerEipId, the system automatically creates an EIP.
	//
	// 	- false: disassociates an EIP from the API server.
	//
	// example:
	//
	// true
	PublicApiServerEnabled *bool `json:"PublicApiServerEnabled,omitempty" xml:"PublicApiServerEnabled,omitempty"`
	// The vSwitches.
	VSwitchesShrink *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	// The scheduling mode of the workflow. This parameter takes effect only if Profile is set to XFlow. Valid values:
	//
	// 	- cost-optimized: cost-prioritized scheduling mode.
	//
	// 	- stock-optimized: inventory-prioritized scheduling mode.
	//
	// example:
	//
	// cost-optimized
	WorkflowScheduleMode *string `json:"WorkflowScheduleMode,omitempty" xml:"WorkflowScheduleMode,omitempty"`
}

func (s UpdateHubClusterFeatureShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHubClusterFeatureShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetAccessControlListShrink(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.AccessControlListShrink = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetApiServerEipId(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.ApiServerEipId = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetArgoCDEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.ArgoCDEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetArgoCDHAEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.ArgoCDHAEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetArgoEventsEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.ArgoEventsEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetArgoServerEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.ArgoServerEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetAuditLogEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.AuditLogEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetClusterId(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetDeletionProtection(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetEnableMesh(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.EnableMesh = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetGatewayEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.GatewayEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetMonitorEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.MonitorEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetName(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetPriceLimit(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.PriceLimit = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetPublicAccessEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.PublicAccessEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetPublicApiServerEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.PublicApiServerEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetVSwitchesShrink(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.VSwitchesShrink = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetWorkflowScheduleMode(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.WorkflowScheduleMode = &v
	return s
}

type UpdateHubClusterFeatureResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 52C1B7DF-96C1-5214-97B6-1B0859FEAFE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHubClusterFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHubClusterFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHubClusterFeatureResponseBody) SetRequestId(v string) *UpdateHubClusterFeatureResponseBody {
	s.RequestId = &v
	return s
}

type UpdateHubClusterFeatureResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHubClusterFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHubClusterFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHubClusterFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateHubClusterFeatureResponse) SetHeaders(v map[string]*string) *UpdateHubClusterFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateHubClusterFeatureResponse) SetStatusCode(v int32) *UpdateHubClusterFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHubClusterFeatureResponse) SetBody(v *UpdateHubClusterFeatureResponseBody) *UpdateHubClusterFeatureResponse {
	s.Body = v
	return s
}

type UpdateUserPermissionRequest struct {
	// The ID of the master instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c8e28143817db4b039b8548d7c899****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespace to which the permissions are scoped. By default, this parameter is empty when you set RoleType to cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Specifies the predefined role that you want to assign. Valid values:
	//
	// 	- admin: the administrator role.
	//
	// 	- dev: the developer role.
	//
	// This parameter is required.
	//
	// example:
	//
	// admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The authorization type. Valid values:
	//
	// 	- cluster: specifies that the permissions are scoped to a master instance.
	//
	// 	- namespace: specifies that the permissions are scoped to a namespace of a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The ID of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2176****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserPermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserPermissionRequest) SetClusterId(v string) *UpdateUserPermissionRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateUserPermissionRequest) SetNamespace(v string) *UpdateUserPermissionRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateUserPermissionRequest) SetRoleName(v string) *UpdateUserPermissionRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateUserPermissionRequest) SetRoleType(v string) *UpdateUserPermissionRequest {
	s.RoleType = &v
	return s
}

func (s *UpdateUserPermissionRequest) SetUserId(v string) *UpdateUserPermissionRequest {
	s.UserId = &v
	return s
}

type UpdateUserPermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 62F5AA2B-A2C9-5135-BCE2-C2167099****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserPermissionResponseBody) SetRequestId(v string) *UpdateUserPermissionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserPermissionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserPermissionResponse) SetHeaders(v map[string]*string) *UpdateUserPermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserPermissionResponse) SetStatusCode(v int32) *UpdateUserPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserPermissionResponse) SetBody(v *UpdateUserPermissionResponseBody) *UpdateUserPermissionResponse {
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
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"cn-beijing":            tea.String("adcp.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("adcp.cn-zhangjiakou.aliyuncs.com"),
		"cn-hangzhou":           tea.String("adcp.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("adcp.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           tea.String("adcp.cn-shenzhen.aliyuncs.com"),
		"cn-heyuan":             tea.String("adcp.cn-heyuan.aliyuncs.com"),
		"cn-hongkong":           tea.String("adcp.cn-hongkong.aliyuncs.com"),
		"ap-northeast-1":        tea.String("adcp.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-1":        tea.String("adcp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("adcp.ap-southeast-5.aliyuncs.com"),
		"ap-south-1":            tea.String("adcp.ap-south-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("adcp.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        tea.String("adcp.ap-southeast-3.aliyuncs.com"),
		"cn-chengdu":            tea.String("adcp-vpc.cn-chengdu.aliyuncs.com"),
		"cn-huhehaote":          tea.String("adcp.cn-huhehaote.aliyuncs.com"),
		"cn-qingdao":            tea.String("adcp.cn-qingdao.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("adcp-vpc.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-wulanchabu":         tea.String("adcp.cn-wulanchabu.aliyuncs.com"),
		"eu-central-1":          tea.String("adcp.eu-central-1.aliyuncs.com"),
		"eu-west-1":             tea.String("adcp-vpc.eu-west-1.aliyuncs.com"),
		"me-east-1":             tea.String("adcp.me-east-1.aliyuncs.com"),
		"us-east-1":             tea.String("adcp.us-east-1.aliyuncs.com"),
		"us-west-1":             tea.String("adcp.us-west-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("adcp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// You can search for API operations, call and debug API operations online, and dynamically generate executable sample code for SDKs.
//
// @param request - AttachClusterToHubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachClusterToHubResponse
func (client *Client) AttachClusterToHubWithOptions(request *AttachClusterToHubRequest, runtime *util.RuntimeOptions) (_result *AttachClusterToHubResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachToMesh)) {
		query["AttachToMesh"] = request.AttachToMesh
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterIds)) {
		body["ClusterIds"] = request.ClusterIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachClusterToHub"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachClusterToHubResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can search for API operations, call and debug API operations online, and dynamically generate executable sample code for SDKs.
//
// @param request - AttachClusterToHubRequest
//
// @return AttachClusterToHubResponse
func (client *Client) AttachClusterToHub(request *AttachClusterToHubRequest) (_result *AttachClusterToHubResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachClusterToHubResponse{}
	_body, _err := client.AttachClusterToHubWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新资源组
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源组
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a master instance in Alibaba Cloud Distributed Cloud Container Platform (ACK One).
//
// @param tmpReq - CreateHubClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHubClusterResponse
func (client *Client) CreateHubClusterWithOptions(tmpReq *CreateHubClusterRequest, runtime *util.RuntimeOptions) (_result *CreateHubClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateHubClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiServerPublicEip)) {
		body["ApiServerPublicEip"] = request.ApiServerPublicEip
	}

	if !tea.BoolValue(util.IsUnset(request.ArgoServerEnabled)) {
		body["ArgoServerEnabled"] = request.ArgoServerEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AuditLogEnabled)) {
		body["AuditLogEnabled"] = request.AuditLogEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnterpriseSecurityGroup)) {
		body["IsEnterpriseSecurityGroup"] = request.IsEnterpriseSecurityGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PriceLimit)) {
		body["PriceLimit"] = request.PriceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Profile)) {
		body["Profile"] = request.Profile
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupID)) {
		body["ResourceGroupID"] = request.ResourceGroupID
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitches)) {
		body["VSwitches"] = request.VSwitches
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowScheduleMode)) {
		body["WorkflowScheduleMode"] = request.WorkflowScheduleMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHubCluster"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHubClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a master instance in Alibaba Cloud Distributed Cloud Container Platform (ACK One).
//
// @param request - CreateHubClusterRequest
//
// @return CreateHubClusterResponse
func (client *Client) CreateHubCluster(request *CreateHubClusterRequest) (_result *CreateHubClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHubClusterResponse{}
	_body, _err := client.CreateHubClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a master cluster in Alibaba Cloud Distributed Cloud Container Platform (ACK One).
//
// @param tmpReq - DeleteHubClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHubClusterResponse
func (client *Client) DeleteHubClusterWithOptions(tmpReq *DeleteHubClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteHubClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteHubClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RetainResources)) {
		request.RetainResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RetainResources, tea.String("RetainResources"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.RetainResourcesShrink)) {
		query["RetainResources"] = request.RetainResourcesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHubCluster"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHubClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a master cluster in Alibaba Cloud Distributed Cloud Container Platform (ACK One).
//
// @param request - DeleteHubClusterRequest
//
// @return DeleteHubClusterResponse
func (client *Client) DeleteHubCluster(request *DeleteHubClusterRequest) (_result *DeleteHubClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHubClusterResponse{}
	_body, _err := client.DeleteHubClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a policy for associated clusters.
//
// @param tmpReq - DeletePolicyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyInstanceResponse
func (client *Client) DeletePolicyInstanceWithOptions(tmpReq *DeletePolicyInstanceRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeletePolicyInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClusterIds)) {
		request.ClusterIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterIds, tea.String("ClusterIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterIdsShrink)) {
		query["ClusterIds"] = request.ClusterIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicyInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePolicyInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a policy for associated clusters.
//
// @param request - DeletePolicyInstanceRequest
//
// @return DeletePolicyInstanceResponse
func (client *Client) DeletePolicyInstance(request *DeletePolicyInstanceRequest) (_result *DeletePolicyInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyInstanceResponse{}
	_body, _err := client.DeletePolicyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the role-based access control (RBAC) permissions of a RAM user.
//
// @param request - DeleteUserPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserPermissionResponse
func (client *Client) DeleteUserPermissionWithOptions(request *DeleteUserPermissionRequest, runtime *util.RuntimeOptions) (_result *DeleteUserPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserPermission"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the role-based access control (RBAC) permissions of a RAM user.
//
// @param request - DeleteUserPermissionRequest
//
// @return DeleteUserPermissionResponse
func (client *Client) DeleteUserPermission(request *DeleteUserPermissionRequest) (_result *DeleteUserPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserPermissionResponse{}
	_body, _err := client.DeleteUserPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys a policy instance in the clusters that are associated with a master instance.
//
// @param tmpReq - DeployPolicyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployPolicyInstanceResponse
func (client *Client) DeployPolicyInstanceWithOptions(tmpReq *DeployPolicyInstanceRequest, runtime *util.RuntimeOptions) (_result *DeployPolicyInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeployPolicyInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClusterIds)) {
		request.ClusterIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterIds, tea.String("ClusterIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Namespaces)) {
		request.NamespacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Namespaces, tea.String("Namespaces"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterIdsShrink)) {
		query["ClusterIds"] = request.ClusterIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NamespacesShrink)) {
		query["Namespaces"] = request.NamespacesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyAction)) {
		query["PolicyAction"] = request.PolicyAction
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployPolicyInstance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployPolicyInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys a policy instance in the clusters that are associated with a master instance.
//
// @param request - DeployPolicyInstanceRequest
//
// @return DeployPolicyInstanceResponse
func (client *Client) DeployPolicyInstance(request *DeployPolicyInstanceRequest) (_result *DeployPolicyInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployPolicyInstanceResponse{}
	_body, _err := client.DeployPolicyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a master instance in Alibaba Cloud Distributed Cloud Container Platform (ACK One).
//
// @param request - DescribeHubClusterDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHubClusterDetailsResponse
func (client *Client) DescribeHubClusterDetailsWithOptions(request *DescribeHubClusterDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeHubClusterDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHubClusterDetails"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHubClusterDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a master instance in Alibaba Cloud Distributed Cloud Container Platform (ACK One).
//
// @param request - DescribeHubClusterDetailsRequest
//
// @return DescribeHubClusterDetailsResponse
func (client *Client) DescribeHubClusterDetails(request *DescribeHubClusterDetailsRequest) (_result *DescribeHubClusterDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHubClusterDetailsResponse{}
	_body, _err := client.DescribeHubClusterDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the kubeconfig file of a master instance.
//
// @param request - DescribeHubClusterKubeconfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHubClusterKubeconfigResponse
func (client *Client) DescribeHubClusterKubeconfigWithOptions(request *DescribeHubClusterKubeconfigRequest, runtime *util.RuntimeOptions) (_result *DescribeHubClusterKubeconfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHubClusterKubeconfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHubClusterKubeconfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the kubeconfig file of a master instance.
//
// @param request - DescribeHubClusterKubeconfigRequest
//
// @return DescribeHubClusterKubeconfigResponse
func (client *Client) DescribeHubClusterKubeconfig(request *DescribeHubClusterKubeconfigRequest) (_result *DescribeHubClusterKubeconfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHubClusterKubeconfigResponse{}
	_body, _err := client.DescribeHubClusterKubeconfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查查HUB集群日志
//
// @param request - DescribeHubClusterLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHubClusterLogsResponse
func (client *Client) DescribeHubClusterLogsWithOptions(request *DescribeHubClusterLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeHubClusterLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHubClusterLogs"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHubClusterLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查查HUB集群日志
//
// @param request - DescribeHubClusterLogsRequest
//
// @return DescribeHubClusterLogsResponse
func (client *Client) DescribeHubClusterLogs(request *DescribeHubClusterLogsRequest) (_result *DescribeHubClusterLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHubClusterLogsResponse{}
	_body, _err := client.DescribeHubClusterLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Distributed Cloud Container Platform for Kubernetes (ACK One) clusters that are created by the current user.
//
// @param tmpReq - DescribeHubClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHubClustersResponse
func (client *Client) DescribeHubClustersWithOptions(tmpReq *DescribeHubClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeHubClustersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeHubClustersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Profile)) {
		query["Profile"] = request.Profile
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHubClusters"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHubClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Distributed Cloud Container Platform for Kubernetes (ACK One) clusters that are created by the current user.
//
// @param request - DescribeHubClustersRequest
//
// @return DescribeHubClustersResponse
func (client *Client) DescribeHubClusters(request *DescribeHubClustersRequest) (_result *DescribeHubClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHubClustersResponse{}
	_body, _err := client.DescribeHubClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Alibaba Cloud CLI allows you to search for API operations, call and debug API operations online, and dynamically generate executable sample code for SDKs.
//
// @param request - DescribeManagedClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeManagedClustersResponse
func (client *Client) DescribeManagedClustersWithOptions(request *DescribeManagedClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeManagedClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeManagedClusters"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeManagedClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Alibaba Cloud CLI allows you to search for API operations, call and debug API operations online, and dynamically generate executable sample code for SDKs.
//
// @param request - DescribeManagedClustersRequest
//
// @return DescribeManagedClustersResponse
func (client *Client) DescribeManagedClusters(request *DescribeManagedClustersRequest) (_result *DescribeManagedClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeManagedClustersResponse{}
	_body, _err := client.DescribeManagedClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of policies.
//
// @param request - DescribePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePoliciesResponse
func (client *Client) DescribePoliciesWithOptions(runtime *util.RuntimeOptions) (_result *DescribePoliciesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicies"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of policies.
//
// @return DescribePoliciesResponse
func (client *Client) DescribePolicies() (_result *DescribePoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePoliciesResponse{}
	_body, _err := client.DescribePoliciesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries detailed information about a policy.
//
// @param request - DescribePolicyDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyDetailsResponse
func (client *Client) DescribePolicyDetailsWithOptions(request *DescribePolicyDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicyDetails"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePolicyDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed information about a policy.
//
// @param request - DescribePolicyDetailsRequest
//
// @return DescribePolicyDetailsResponse
func (client *Client) DescribePolicyDetails(request *DescribePolicyDetailsRequest) (_result *DescribePolicyDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyDetailsResponse{}
	_body, _err := client.DescribePolicyDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries detailed information about the policies used by the clusters that are associated with a master instance.
//
// @param request - DescribePolicyGovernanceInClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyGovernanceInClusterResponse
func (client *Client) DescribePolicyGovernanceInClusterWithOptions(request *DescribePolicyGovernanceInClusterRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyGovernanceInClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicyGovernanceInCluster"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePolicyGovernanceInClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed information about the policies used by the clusters that are associated with a master instance.
//
// @param request - DescribePolicyGovernanceInClusterRequest
//
// @return DescribePolicyGovernanceInClusterResponse
func (client *Client) DescribePolicyGovernanceInCluster(request *DescribePolicyGovernanceInClusterRequest) (_result *DescribePolicyGovernanceInClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyGovernanceInClusterResponse{}
	_body, _err := client.DescribePolicyGovernanceInClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries policy instances that are deployed in the clusters associated with a master instance.
//
// @param request - DescribePolicyInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyInstancesResponse
func (client *Client) DescribePolicyInstancesWithOptions(request *DescribePolicyInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicyInstances"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePolicyInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries policy instances that are deployed in the clusters associated with a master instance.
//
// @param request - DescribePolicyInstancesRequest
//
// @return DescribePolicyInstancesResponse
func (client *Client) DescribePolicyInstances(request *DescribePolicyInstancesRequest) (_result *DescribePolicyInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyInstancesResponse{}
	_body, _err := client.DescribePolicyInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries detailed information about policy instances that are deployed in the clusters associated with a master instance.
//
// @param request - DescribePolicyInstancesStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyInstancesStatusResponse
func (client *Client) DescribePolicyInstancesStatusWithOptions(request *DescribePolicyInstancesStatusRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyInstancesStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicyInstancesStatus"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePolicyInstancesStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed information about policy instances that are deployed in the clusters associated with a master instance.
//
// @param request - DescribePolicyInstancesStatusRequest
//
// @return DescribePolicyInstancesStatusResponse
func (client *Client) DescribePolicyInstancesStatus(request *DescribePolicyInstancesStatusRequest) (_result *DescribePolicyInstancesStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyInstancesStatusResponse{}
	_body, _err := client.DescribePolicyInstancesStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询地域列表
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// 查询地域列表
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the permissions of a Resource Access Management (RAM) user.
//
// @param request - DescribeUserPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserPermissionsResponse
func (client *Client) DescribeUserPermissionsWithOptions(request *DescribeUserPermissionsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserPermissions"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the permissions of a Resource Access Management (RAM) user.
//
// @param request - DescribeUserPermissionsRequest
//
// @return DescribeUserPermissionsResponse
func (client *Client) DescribeUserPermissions(request *DescribeUserPermissionsRequest) (_result *DescribeUserPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserPermissionsResponse{}
	_body, _err := client.DescribeUserPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Alibaba Cloud CLI allows you to search for API operations, call and debug API operations online, and dynamically generate executable sample code for SDKs.
//
// @param request - DetachClusterFromHubRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachClusterFromHubResponse
func (client *Client) DetachClusterFromHubWithOptions(request *DetachClusterFromHubRequest, runtime *util.RuntimeOptions) (_result *DetachClusterFromHubResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DetachFromMesh)) {
		query["DetachFromMesh"] = request.DetachFromMesh
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterIds)) {
		body["ClusterIds"] = request.ClusterIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachClusterFromHub"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachClusterFromHubResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Alibaba Cloud CLI allows you to search for API operations, call and debug API operations online, and dynamically generate executable sample code for SDKs.
//
// @param request - DetachClusterFromHubRequest
//
// @return DetachClusterFromHubResponse
func (client *Client) DetachClusterFromHub(request *DetachClusterFromHubRequest) (_result *DetachClusterFromHubResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachClusterFromHubResponse{}
	_body, _err := client.DetachClusterFromHubWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Schema of Response
//
// @param request - GrantUserPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantUserPermissionResponse
func (client *Client) GrantUserPermissionWithOptions(request *GrantUserPermissionRequest, runtime *util.RuntimeOptions) (_result *GrantUserPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IsRamRole)) {
		query["IsRamRole"] = request.IsRamRole
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		query["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantUserPermission"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantUserPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Schema of Response
//
// @param request - GrantUserPermissionRequest
//
// @return GrantUserPermissionResponse
func (client *Client) GrantUserPermission(request *GrantUserPermissionRequest) (_result *GrantUserPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantUserPermissionResponse{}
	_body, _err := client.GrantUserPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GrantUserPermissions is deprecated, please use adcp::2022-01-01::GrantUserPermission instead.
//
// Summary:
//
// Grant permissions to a Resource Access Management (RAM) user.
//
// @param tmpReq - GrantUserPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantUserPermissionsResponse
// Deprecated
func (client *Client) GrantUserPermissionsWithOptions(tmpReq *GrantUserPermissionsRequest, runtime *util.RuntimeOptions) (_result *GrantUserPermissionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GrantUserPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Permissions)) {
		request.PermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permissions, tea.String("Permissions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PermissionsShrink)) {
		query["Permissions"] = request.PermissionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantUserPermissions"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantUserPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GrantUserPermissions is deprecated, please use adcp::2022-01-01::GrantUserPermission instead.
//
// Summary:
//
// Grant permissions to a Resource Access Management (RAM) user.
//
// @param request - GrantUserPermissionsRequest
//
// @return GrantUserPermissionsResponse
// Deprecated
func (client *Client) GrantUserPermissions(request *GrantUserPermissionsRequest) (_result *GrantUserPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantUserPermissionsResponse{}
	_body, _err := client.GrantUserPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of a Container Service for Kubernetes (ACK) cluster that serves as a master instance.
//
// @param tmpReq - UpdateHubClusterFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHubClusterFeatureResponse
func (client *Client) UpdateHubClusterFeatureWithOptions(tmpReq *UpdateHubClusterFeatureRequest, runtime *util.RuntimeOptions) (_result *UpdateHubClusterFeatureResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateHubClusterFeatureShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccessControlList)) {
		request.AccessControlListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccessControlList, tea.String("AccessControlList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VSwitches)) {
		request.VSwitchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitches, tea.String("VSwitches"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessControlListShrink)) {
		query["AccessControlList"] = request.AccessControlListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ApiServerEipId)) {
		query["ApiServerEipId"] = request.ApiServerEipId
	}

	if !tea.BoolValue(util.IsUnset(request.ArgoCDEnabled)) {
		query["ArgoCDEnabled"] = request.ArgoCDEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ArgoCDHAEnabled)) {
		query["ArgoCDHAEnabled"] = request.ArgoCDHAEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ArgoEventsEnabled)) {
		query["ArgoEventsEnabled"] = request.ArgoEventsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ArgoServerEnabled)) {
		query["ArgoServerEnabled"] = request.ArgoServerEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AuditLogEnabled)) {
		query["AuditLogEnabled"] = request.AuditLogEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DeletionProtection)) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !tea.BoolValue(util.IsUnset(request.EnableMesh)) {
		query["EnableMesh"] = request.EnableMesh
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayEnabled)) {
		query["GatewayEnabled"] = request.GatewayEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorEnabled)) {
		query["MonitorEnabled"] = request.MonitorEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PriceLimit)) {
		query["PriceLimit"] = request.PriceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.PublicAccessEnabled)) {
		query["PublicAccessEnabled"] = request.PublicAccessEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.PublicApiServerEnabled)) {
		query["PublicApiServerEnabled"] = request.PublicApiServerEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchesShrink)) {
		query["VSwitches"] = request.VSwitchesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowScheduleMode)) {
		query["WorkflowScheduleMode"] = request.WorkflowScheduleMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateHubClusterFeature"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateHubClusterFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a Container Service for Kubernetes (ACK) cluster that serves as a master instance.
//
// @param request - UpdateHubClusterFeatureRequest
//
// @return UpdateHubClusterFeatureResponse
func (client *Client) UpdateHubClusterFeature(request *UpdateHubClusterFeatureRequest) (_result *UpdateHubClusterFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHubClusterFeatureResponse{}
	_body, _err := client.UpdateHubClusterFeatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the role-based access control (RBAC) permissions of a RAM user.
//
// @param request - UpdateUserPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPermissionResponse
func (client *Client) UpdateUserPermissionWithOptions(request *UpdateUserPermissionRequest, runtime *util.RuntimeOptions) (_result *UpdateUserPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		query["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserPermission"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the role-based access control (RBAC) permissions of a RAM user.
//
// @param request - UpdateUserPermissionRequest
//
// @return UpdateUserPermissionResponse
func (client *Client) UpdateUserPermission(request *UpdateUserPermissionRequest) (_result *UpdateUserPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserPermissionResponse{}
	_body, _err := client.UpdateUserPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
