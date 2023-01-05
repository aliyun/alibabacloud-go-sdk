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

type AttachClusterToHubRequest struct {
	// Specifies whether to associate the clusters with Service Mesh (ASM) instances. Valid values:
	AttachToMesh *bool `json:"AttachToMesh,omitempty" xml:"AttachToMesh,omitempty"`
	// The ID of the master instance.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// A JSON string that can be parsed into a string array. The string specifies the clusters that you want to associate with the master instance.
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
	// The ID of the master instance.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// A list of the IDs of the clusters that you want to associate with the master instance.
	ManagedClusterIds []*string `json:"ManagedClusterIds,omitempty" xml:"ManagedClusterIds,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachClusterToHubResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateHubClusterRequest struct {
	// Specifies whether to use a public IP address to expose the API server. Valid values: - true: uses a public IP address to expose the API server. - true: uses an internal IP address to expose the API server.
	ApiServerPublicEip *bool `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	// Specifies whether to enable audit logs. Valid values: - true: enables audit logs. - false: disables audit logs.
	AuditLogEnabled *bool `json:"AuditLogEnabled,omitempty" xml:"AuditLogEnabled,omitempty"`
	// Specifies whether the security group is an advanced security group.
	IsEnterpriseSecurityGroup *bool `json:"IsEnterpriseSecurityGroup,omitempty" xml:"IsEnterpriseSecurityGroup,omitempty"`
	// The name of the master instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Scenario-oriented master control type. The value can be:
	//
	// - `Default`: Standard scenario Master instance.
	// - `XFlow`: Workflow scenario master instance.
	//
	// Default Value: `Default`.
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the region. You can call the DescribeRegions operation to query available regions.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the vSwitch.
	VSwitches *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the master instance belongs. You can call the DescribeVpcs operation to query available VPCs.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *CreateHubClusterRequest) SetProfile(v string) *CreateHubClusterRequest {
	s.Profile = &v
	return s
}

func (s *CreateHubClusterRequest) SetRegionId(v string) *CreateHubClusterRequest {
	s.RegionId = &v
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

type CreateHubClusterResponseBody struct {
	// The ID of the master instance.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHubClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to forcefully delete the master instance. Valid values: - true: forcefully delete the master instance. - false: does not forcefully delete the master instance. Default value: false.
	Force           *bool     `json:"Force,omitempty" xml:"Force,omitempty"`
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
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to forcefully delete the master instance. Valid values: - true: forcefully delete the master instance. - false: does not forcefully delete the master instance. Default value: false.
	Force                 *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
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
	// The ID of the master instance.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the master instance.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHubClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeHubClusterDetailsRequest struct {
	// The ID of the master instance.
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
	// The details about the master instance.
	Cluster *DescribeHubClusterDetailsResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	// The ID of the request.
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
	// Information about the API server of the master instance.
	ApiServer *DescribeHubClusterDetailsResponseBodyClusterApiServer `json:"ApiServer,omitempty" xml:"ApiServer,omitempty" type:"Struct"`
	// The details about the master instance.
	ClusterInfo *DescribeHubClusterDetailsResponseBodyClusterClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// 集群删除条件信息列表
	Conditions []*DescribeHubClusterDetailsResponseBodyClusterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The endpoint of the master instance.
	Endpoints *DescribeHubClusterDetailsResponseBodyClusterEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The logging configuration.
	LogConfig *DescribeHubClusterDetailsResponseBodyClusterLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// The Service Mesh (ASM) configurations.
	MeshConfig *DescribeHubClusterDetailsResponseBodyClusterMeshConfig `json:"MeshConfig,omitempty" xml:"MeshConfig,omitempty" type:"Struct"`
	// The network configurations of the master instance.
	Network *DescribeHubClusterDetailsResponseBodyClusterNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
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

type DescribeHubClusterDetailsResponseBodyClusterApiServer struct {
	// The ID of the elastic IP address (EIP).
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// Indicates whether a public endpoint is used to expose the API server. Valid values: - true: a public endpoint is used to expose the API server. - false: no public endpoint is used to expose the API server.
	EnabledPublic *bool `json:"EnabledPublic,omitempty" xml:"EnabledPublic,omitempty"`
	// The ID of the Server Load Balancer (SLB) instance.
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
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The specification of the master instance. Valid values: - ack.pro.small: ACK Pro
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The time when the master instance was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message that is returned when the system fails to create the master instance.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the master instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of the master instance.
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the region in which the master instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the master instance. Valid values: - initial: The master instance is being initialized. - failed: The master instance failed to be created. - running: The master instance is running. - inactive: The master instance is inactive. - deleting: The master instance is being deleted. - delete_failed: The master instance failed to be deleted. - deleted: The master instance is deleted.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the master instance was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The Kubernetes version of the master instance.
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

func (s *DescribeHubClusterDetailsResponseBodyClusterClusterInfo) SetState(v string) *DescribeHubClusterDetailsResponseBodyClusterClusterInfo {
	s.State = &v
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

type DescribeHubClusterDetailsResponseBodyClusterConditions struct {
	// 删除条件错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 删除条件原因
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 删除条件状态，取值
	// - True 不能删除
	// - False 允许删除
	// - Unknow 未知
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 删除条件类型
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
	// The internal endpoint of the API server of the master instance.
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	// The public endpoint of the API server of the master instance.
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
	// Indicates whether audit logs are enabled. Valid values: - true: audit logs are enabled. - false: audit logs are disabled.
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// The name of the Log Service project.
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The retention period of the logs.
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
	// Indicates whether ASM is enabled. Valid values: - true: ASM is enabled. - false: ASM is disabled.
	EnableMesh *bool `json:"EnableMesh,omitempty" xml:"EnableMesh,omitempty"`
	// The ID of the ASM instance.
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
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The IP version that is supported by the master instance. Valid values: - ipv4: IPv4. - ipv6: IPv6. - dual: IPv4 and IPv6.
	IPStack *string `json:"IPStack,omitempty" xml:"IPStack,omitempty"`
	// The ID of the associated security group.
	SecurityGroupIDs []*string `json:"SecurityGroupIDs,omitempty" xml:"SecurityGroupIDs,omitempty" type:"Repeated"`
	// A list of the vSwitches that are used by the master instance.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) in which the master instance resides.
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

type DescribeHubClusterDetailsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHubClusterDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to obtain the credential that is used to connect to the master instance over the internal network. Valid values: - `true`: obtains only the credential that is used to access the master instance over the internal network. - `false`: obtains only the credential that is used to access the master instance over the Internet. Default value: `false`.
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
	// The content of the kubeconfig file of the master instance.
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHubClusterKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// A log of the master instance.
	ClusterLog *string `json:"ClusterLog,omitempty" xml:"ClusterLog,omitempty"`
	// The time when the log was created. Format: <i>yyyy-mm-dd</i>t<i>hh:mm:ss</i>z (UTC time).
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The severity level of the log. Valid values: - error: errors. - warn: warnings. - info: information.
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHubClusterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The scenario where master instances are used. Valid values:
	//
	// *   `Default`: standard scenarios.
	// *   `XFlow`: workflow scenarios.
	//
	// Default value: `Default`.
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
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

type DescribeHubClustersResponseBody struct {
	// The list of the master instances returned.
	Clusters []*DescribeHubClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The ID of the request.
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
	// The details of the Kubernetes API server.
	ApiServer *DescribeHubClustersResponseBodyClustersApiServer `json:"ApiServer,omitempty" xml:"ApiServer,omitempty" type:"Struct"`
	// The details of the master instance.
	ClusterInfo *DescribeHubClustersResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// The list of the deletion conditions of the master instance.
	Conditions []*DescribeHubClustersResponseBodyClustersConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The endpoint of the master instance.
	Endpoints *DescribeHubClustersResponseBodyClustersEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The logging configurations.
	LogConfig *DescribeHubClustersResponseBodyClustersLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// The configurations of Alibaba Cloud Service Mesh (ASM).
	MeshConfig *DescribeHubClustersResponseBodyClustersMeshConfig `json:"MeshConfig,omitempty" xml:"MeshConfig,omitempty" type:"Struct"`
	// The network configurations of the master instance.
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
	// The ID of the elastic IP address (EIP).
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// Indicates whether the API server is accessible over the Internet. Valid values:
	//
	// *   true: The API server is accessible over the Internet.
	// *   false: The API server is inaccessible over the Internet.
	EnabledPublic *bool `json:"EnabledPublic,omitempty" xml:"EnabledPublic,omitempty"`
	// The ID of the Server Load Balancer (SLB) instance that is associated with the Kubernetes API server.
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
	// The ID of the master instance.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The specification of the master instance.
	//
	// *   ack.pro.small: ACK Pro cluster
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The time when the master instance was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message returned when the master instance failed to be created.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the master instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of the master instance.
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The ID of the region in which the master instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the master instance. Valid values:
	//
	// *   initial: The master instance is being initialized.
	// *   failed: The master instance failed to be created.
	// *   running: The master instance is running
	// *   inactive: The master instance is pending.
	// *   deleting: The master instance is being deleted.
	// *   delete_failed: The master instance failed to be deleted.
	// *   deleted: The master instance is deleted.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The last time when the master instance was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The Kubernetes version of the master instance.
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

func (s *DescribeHubClustersResponseBodyClustersClusterInfo) SetState(v string) *DescribeHubClustersResponseBodyClustersClusterInfo {
	s.State = &v
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

type DescribeHubClustersResponseBodyClustersConditions struct {
	// The error message of the deletion condition.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the deletion condition.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the deletion condition. Valid values:
	//
	// *   True: The master instance cannot be deleted.
	// *   False: The master instance can be deleted.
	// *   Unknow: Whether the master instance can be deleted is unknown.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of deletion condition.
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
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	// The public endpoint of the API server.
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
	// Indicates whether audit logging is enabled. Valid values:
	//
	// *   true: Audit logging is enabled.
	// *   false: Audit logging is disabled.
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// The name of the project of Log Service.
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The number of days that logs are retained by Log Service.
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
	// *   true: ASM is enabled.
	// *   false: ASM is disabled.
	EnableMesh *bool `json:"EnableMesh,omitempty" xml:"EnableMesh,omitempty"`
	// The ID of the ASM instance.
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
	// The domain name of the master instance.
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The security group IDs of the master instance.
	SecurityGroupIDs []*string `json:"SecurityGroupIDs,omitempty" xml:"SecurityGroupIDs,omitempty" type:"Repeated"`
	// The IDs of the vSwitches to which the master instance is connected.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the master instance belongs.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHubClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The ID of the master instance.
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
	// Information about the master instance.
	Clusters []*DescribeManagedClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The ID of the request.
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
	// The status of the association between the clusters and Service Mesh (ASM).
	MeshStatus *DescribeManagedClustersResponseBodyClustersMeshStatus `json:"MeshStatus,omitempty" xml:"MeshStatus,omitempty" type:"Struct"`
	// The status of the association between the clusters and the master instance.
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
	// The ID of the master instance.
	ClusterID *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	// The specification of the master instance. Valid values: - ack.pro.small: ACK Pro.
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// The type of the master instance.
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the master instance was created.
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The current Kubernetes version of the master instance.
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The original Kubernetes version of the master instance.
	InitVersion *string `json:"InitVersion,omitempty" xml:"InitVersion,omitempty"`
	// The name of the master instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the master instance.
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The region in which the master instance resides.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the associated clusters. Valid values: - initial: The associated clusters are being initialized. - failed: The associated clustersfailed to be created. - running: The associated clusters are running. - inactive: The associated clusters are inactive. - deleting: The associated clusters are being deleted. - deleted: The associated clusters are deleted.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the master instance was updated.
	Updated *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	// The ID of the vSwitch.
	VSwitchID *string `json:"VSwitchID,omitempty" xml:"VSwitchID,omitempty"`
	// VPC ID.
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
	// Indicates whether the clusters are associated with ASM instances. Valid values: - true: The clusters are associated with ASM instances. - false: The clusters are not associated with ASM instances.
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
	// The status information.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The status of the association between the clusters and the master instance. Valid values: - Installing: The clusters are being associated with the master instance. - Successed: The clusters are associated with the master instance. - Failed: The clusters failed to be associated with the master instance. - Deleting: The clusters are being disassociated from the master instance. - Deleted: The clusters are disassociated from the master instance.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeManagedClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeRegionsRequest struct {
	// The language. Valid values: zh, en, and jp.
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
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DetachClusterFromHubRequest struct {
	// The ID of the master instance.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// A JSON string that can be parsed into a string array. The string specifies the clusters that you want to disassociate from the master instance.
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// Specifies whether to only disassociate the clusters from Service Mesh (ASM) instances. Valid values: - true: only disassociates the clusters from ASM instances. - false: disassociates the clusters from the master instance and ASM instances.
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
	// The ID of the master instance.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The IDs of the clusters that are disassociated from the master instance.
	ManagedClusterIds []*string `json:"ManagedClusterIds,omitempty" xml:"ManagedClusterIds,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachClusterFromHubResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateHubClusterFeatureRequest struct {
	// The ID of the EIP.
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// Specifies whether to enable audit logs. Valid values: - true: enable audit logs. - false: disables audit logs.
	AuditLogEnabled *bool `json:"AuditLogEnabled,omitempty" xml:"AuditLogEnabled,omitempty"`
	// The ID of the master instance.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable deletion protection for the master instance. After you enable deletion protection, you cannot delete the master instance in the console or by calling API operations. Valid values:
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Whether to enable ArgoCD.
	//
	// - true Enabled
	// - false Disabled
	EnableArgoCD *bool `json:"EnableArgoCD,omitempty" xml:"EnableArgoCD,omitempty"`
	// Specifies whether to enable Service Mesh (ASM). Valid values: true: enables ASM. false: disables ASM.
	EnableMesh *bool `json:"EnableMesh,omitempty" xml:"EnableMesh,omitempty"`
	Enabled    *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the master instance. The name must be 1 to 63 characters in length, and can contain letters and digits. The name must start with a letter. The name can contain letters, digits, underscores (_), and hyphens (-).
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PriceLimit *string `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	// Specifies whether to associate an elastic IP address (EIP) with the API server. Default value: false. To associate an EIP with the API server, set the value to true. You can use a custom EIP by setting the ApiServerEipId parameter. If you do not set the ApiServerEipId parameter, the system automatically creates an EIP.
	PublicApiServerEnabled *bool                                  `json:"PublicApiServerEnabled,omitempty" xml:"PublicApiServerEnabled,omitempty"`
	ScheduleMode           *string                                `json:"ScheduleMode,omitempty" xml:"ScheduleMode,omitempty"`
	ServerEnabled          *bool                                  `json:"ServerEnabled,omitempty" xml:"ServerEnabled,omitempty"`
	Units                  []*UpdateHubClusterFeatureRequestUnits `json:"Units,omitempty" xml:"Units,omitempty" type:"Repeated"`
}

func (s UpdateHubClusterFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHubClusterFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateHubClusterFeatureRequest) SetApiServerEipId(v string) *UpdateHubClusterFeatureRequest {
	s.ApiServerEipId = &v
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

func (s *UpdateHubClusterFeatureRequest) SetEnableArgoCD(v bool) *UpdateHubClusterFeatureRequest {
	s.EnableArgoCD = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetEnableMesh(v bool) *UpdateHubClusterFeatureRequest {
	s.EnableMesh = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.Enabled = &v
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

func (s *UpdateHubClusterFeatureRequest) SetPublicApiServerEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.PublicApiServerEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetScheduleMode(v string) *UpdateHubClusterFeatureRequest {
	s.ScheduleMode = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetServerEnabled(v bool) *UpdateHubClusterFeatureRequest {
	s.ServerEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureRequest) SetUnits(v []*UpdateHubClusterFeatureRequestUnits) *UpdateHubClusterFeatureRequest {
	s.Units = v
	return s
}

type UpdateHubClusterFeatureRequestUnits struct {
	RegionId  *string                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitches []*UpdateHubClusterFeatureRequestUnitsVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VpcId     *string                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateHubClusterFeatureRequestUnits) String() string {
	return tea.Prettify(s)
}

func (s UpdateHubClusterFeatureRequestUnits) GoString() string {
	return s.String()
}

func (s *UpdateHubClusterFeatureRequestUnits) SetRegionId(v string) *UpdateHubClusterFeatureRequestUnits {
	s.RegionId = &v
	return s
}

func (s *UpdateHubClusterFeatureRequestUnits) SetVSwitches(v []*UpdateHubClusterFeatureRequestUnitsVSwitches) *UpdateHubClusterFeatureRequestUnits {
	s.VSwitches = v
	return s
}

func (s *UpdateHubClusterFeatureRequestUnits) SetVpcId(v string) *UpdateHubClusterFeatureRequestUnits {
	s.VpcId = &v
	return s
}

type UpdateHubClusterFeatureRequestUnitsVSwitches struct {
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateHubClusterFeatureRequestUnitsVSwitches) String() string {
	return tea.Prettify(s)
}

func (s UpdateHubClusterFeatureRequestUnitsVSwitches) GoString() string {
	return s.String()
}

func (s *UpdateHubClusterFeatureRequestUnitsVSwitches) SetVswitchId(v string) *UpdateHubClusterFeatureRequestUnitsVSwitches {
	s.VswitchId = &v
	return s
}

func (s *UpdateHubClusterFeatureRequestUnitsVSwitches) SetZoneId(v string) *UpdateHubClusterFeatureRequestUnitsVSwitches {
	s.ZoneId = &v
	return s
}

type UpdateHubClusterFeatureShrinkRequest struct {
	// The ID of the EIP.
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// Specifies whether to enable audit logs. Valid values: - true: enable audit logs. - false: disables audit logs.
	AuditLogEnabled *bool `json:"AuditLogEnabled,omitempty" xml:"AuditLogEnabled,omitempty"`
	// The ID of the master instance.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable deletion protection for the master instance. After you enable deletion protection, you cannot delete the master instance in the console or by calling API operations. Valid values:
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Whether to enable ArgoCD.
	//
	// - true Enabled
	// - false Disabled
	EnableArgoCD *bool `json:"EnableArgoCD,omitempty" xml:"EnableArgoCD,omitempty"`
	// Specifies whether to enable Service Mesh (ASM). Valid values: true: enables ASM. false: disables ASM.
	EnableMesh *bool `json:"EnableMesh,omitempty" xml:"EnableMesh,omitempty"`
	Enabled    *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the master instance. The name must be 1 to 63 characters in length, and can contain letters and digits. The name must start with a letter. The name can contain letters, digits, underscores (_), and hyphens (-).
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PriceLimit *string `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	// Specifies whether to associate an elastic IP address (EIP) with the API server. Default value: false. To associate an EIP with the API server, set the value to true. You can use a custom EIP by setting the ApiServerEipId parameter. If you do not set the ApiServerEipId parameter, the system automatically creates an EIP.
	PublicApiServerEnabled *bool   `json:"PublicApiServerEnabled,omitempty" xml:"PublicApiServerEnabled,omitempty"`
	ScheduleMode           *string `json:"ScheduleMode,omitempty" xml:"ScheduleMode,omitempty"`
	ServerEnabled          *bool   `json:"ServerEnabled,omitempty" xml:"ServerEnabled,omitempty"`
	UnitsShrink            *string `json:"Units,omitempty" xml:"Units,omitempty"`
}

func (s UpdateHubClusterFeatureShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHubClusterFeatureShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetApiServerEipId(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.ApiServerEipId = &v
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

func (s *UpdateHubClusterFeatureShrinkRequest) SetEnableArgoCD(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.EnableArgoCD = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetEnableMesh(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.EnableMesh = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.Enabled = &v
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

func (s *UpdateHubClusterFeatureShrinkRequest) SetPublicApiServerEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.PublicApiServerEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetScheduleMode(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.ScheduleMode = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetServerEnabled(v bool) *UpdateHubClusterFeatureShrinkRequest {
	s.ServerEnabled = &v
	return s
}

func (s *UpdateHubClusterFeatureShrinkRequest) SetUnitsShrink(v string) *UpdateHubClusterFeatureShrinkRequest {
	s.UnitsShrink = &v
	return s
}

type UpdateHubClusterFeatureResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateHubClusterFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.SignatureAlgorithm = tea.String("v2")
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

func (client *Client) CreateHubClusterWithOptions(request *CreateHubClusterRequest, runtime *util.RuntimeOptions) (_result *CreateHubClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiServerPublicEip)) {
		body["ApiServerPublicEip"] = request.ApiServerPublicEip
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

	if !tea.BoolValue(util.IsUnset(request.Profile)) {
		body["Profile"] = request.Profile
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitches)) {
		body["VSwitches"] = request.VSwitches
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
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

func (client *Client) DescribeHubClustersWithOptions(request *DescribeHubClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeHubClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Profile)) {
		query["Profile"] = request.Profile
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

func (client *Client) UpdateHubClusterFeatureWithOptions(tmpReq *UpdateHubClusterFeatureRequest, runtime *util.RuntimeOptions) (_result *UpdateHubClusterFeatureResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateHubClusterFeatureShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Units)) {
		request.UnitsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Units, tea.String("Units"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiServerEipId)) {
		query["ApiServerEipId"] = request.ApiServerEipId
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

	if !tea.BoolValue(util.IsUnset(request.EnableArgoCD)) {
		query["EnableArgoCD"] = request.EnableArgoCD
	}

	if !tea.BoolValue(util.IsUnset(request.EnableMesh)) {
		query["EnableMesh"] = request.EnableMesh
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PriceLimit)) {
		query["PriceLimit"] = request.PriceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.PublicApiServerEnabled)) {
		query["PublicApiServerEnabled"] = request.PublicApiServerEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleMode)) {
		query["ScheduleMode"] = request.ScheduleMode
	}

	if !tea.BoolValue(util.IsUnset(request.ServerEnabled)) {
		query["ServerEnabled"] = request.ServerEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.UnitsShrink)) {
		query["Units"] = request.UnitsShrink
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
