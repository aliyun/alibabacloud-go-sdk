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

type AttachClusterToHubRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
}

func (s AttachClusterToHubRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachClusterToHubRequest) GoString() string {
	return s.String()
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
	ClusterId         *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ManagedClusterIds []*string `json:"ManagedClusterIds,omitempty" xml:"ManagedClusterIds,omitempty" type:"Repeated"`
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type AttachClusterToHubResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachClusterToHubResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachClusterToHubResponse) SetBody(v *AttachClusterToHubResponseBody) *AttachClusterToHubResponse {
	s.Body = v
	return s
}

type CreateHubClusterRequest struct {
	ApiServerPublicEip     *bool   `json:"ApiServerPublicEip,omitempty" xml:"ApiServerPublicEip,omitempty"`
	AuditLogEnabled        *bool   `json:"AuditLogEnabled,omitempty" xml:"AuditLogEnabled,omitempty"`
	AuditLogProject        *string `json:"AuditLogProject,omitempty" xml:"AuditLogProject,omitempty"`
	AuditLogStoreTTL       *string `json:"AuditLogStoreTTL,omitempty" xml:"AuditLogStoreTTL,omitempty"`
	ControlPlaneLogEnabled *bool   `json:"ControlPlaneLogEnabled,omitempty" xml:"ControlPlaneLogEnabled,omitempty"`
	ControlPlaneLogProject *string `json:"ControlPlaneLogProject,omitempty" xml:"ControlPlaneLogProject,omitempty"`
	ControlPlaneLogTTL     *string `json:"ControlPlaneLogTTL,omitempty" xml:"ControlPlaneLogTTL,omitempty"`
	// 是否企业安全组
	IsEnterpriseSecurityGroup *bool `json:"IsEnterpriseSecurityGroup,omitempty" xml:"IsEnterpriseSecurityGroup,omitempty"`
	// 集群名称
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitches *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *CreateHubClusterRequest) SetAuditLogProject(v string) *CreateHubClusterRequest {
	s.AuditLogProject = &v
	return s
}

func (s *CreateHubClusterRequest) SetAuditLogStoreTTL(v string) *CreateHubClusterRequest {
	s.AuditLogStoreTTL = &v
	return s
}

func (s *CreateHubClusterRequest) SetControlPlaneLogEnabled(v bool) *CreateHubClusterRequest {
	s.ControlPlaneLogEnabled = &v
	return s
}

func (s *CreateHubClusterRequest) SetControlPlaneLogProject(v string) *CreateHubClusterRequest {
	s.ControlPlaneLogProject = &v
	return s
}

func (s *CreateHubClusterRequest) SetControlPlaneLogTTL(v string) *CreateHubClusterRequest {
	s.ControlPlaneLogTTL = &v
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
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHubClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateHubClusterResponse) SetBody(v *CreateHubClusterResponseBody) *CreateHubClusterResponse {
	s.Body = v
	return s
}

type DeleteHubClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Force     *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
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

type DeleteHubClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHubClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteHubClusterResponse) SetBody(v *DeleteHubClusterResponseBody) *DeleteHubClusterResponse {
	s.Body = v
	return s
}

type DescribeHubClusterDetailsRequest struct {
	// 集群ID
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
	Cluster   *DescribeHubClusterDetailsResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ApiServer   *DescribeHubClusterDetailsResponseBodyClusterApiServer   `json:"ApiServer,omitempty" xml:"ApiServer,omitempty" type:"Struct"`
	ClusterInfo *DescribeHubClusterDetailsResponseBodyClusterClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	Endpoints   *DescribeHubClusterDetailsResponseBodyClusterEndpoints   `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	Network     *DescribeHubClusterDetailsResponseBodyClusterNetwork     `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
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

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetEndpoints(v *DescribeHubClusterDetailsResponseBodyClusterEndpoints) *DescribeHubClusterDetailsResponseBodyCluster {
	s.Endpoints = v
	return s
}

func (s *DescribeHubClusterDetailsResponseBodyCluster) SetNetwork(v *DescribeHubClusterDetailsResponseBodyClusterNetwork) *DescribeHubClusterDetailsResponseBodyCluster {
	s.Network = v
	return s
}

type DescribeHubClusterDetailsResponseBodyClusterApiServer struct {
	EnabledPublic  *bool   `json:"EnabledPublic,omitempty" xml:"EnabledPublic,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DescribeHubClusterDetailsResponseBodyClusterApiServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClusterDetailsResponseBodyClusterApiServer) GoString() string {
	return s.String()
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
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterSpec  *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Profile      *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdateTime   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type DescribeHubClusterDetailsResponseBodyClusterEndpoints struct {
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	PublicApiServerEndpoint   *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
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

type DescribeHubClusterDetailsResponseBodyClusterNetwork struct {
	ClusterDomain    *string   `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	IPStack          *string   `json:"IPStack,omitempty" xml:"IPStack,omitempty"`
	SecurityGroupIDs []*string `json:"SecurityGroupIDs,omitempty" xml:"SecurityGroupIDs,omitempty" type:"Repeated"`
	VSwitches        []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VpcId            *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHubClusterDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeHubClusterDetailsResponse) SetBody(v *DescribeHubClusterDetailsResponseBody) *DescribeHubClusterDetailsResponse {
	s.Body = v
	return s
}

type DescribeHubClusterKubeconfigRequest struct {
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PrivateIpAddress *bool   `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
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
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHubClusterKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeHubClusterKubeconfigResponse) SetBody(v *DescribeHubClusterKubeconfigResponseBody) *DescribeHubClusterKubeconfigResponse {
	s.Body = v
	return s
}

type DescribeHubClusterLogsRequest struct {
	// 集群ID
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
	Logs []*DescribeHubClusterLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// Id of the request
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
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterLog   *string `json:"ClusterLog,omitempty" xml:"ClusterLog,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	LogLevel     *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHubClusterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeHubClusterLogsResponse) SetBody(v *DescribeHubClusterLogsResponseBody) *DescribeHubClusterLogsResponse {
	s.Body = v
	return s
}

type DescribeHubClustersResponseBody struct {
	Clusters  []*DescribeHubClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ApiServer   *DescribeHubClustersResponseBodyClustersApiServer   `json:"ApiServer,omitempty" xml:"ApiServer,omitempty" type:"Struct"`
	ClusterInfo *DescribeHubClustersResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	Endpoints   *DescribeHubClustersResponseBodyClustersEndpoints   `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	Network     *DescribeHubClustersResponseBodyClustersNetwork     `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
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

func (s *DescribeHubClustersResponseBodyClusters) SetEndpoints(v *DescribeHubClustersResponseBodyClustersEndpoints) *DescribeHubClustersResponseBodyClusters {
	s.Endpoints = v
	return s
}

func (s *DescribeHubClustersResponseBodyClusters) SetNetwork(v *DescribeHubClustersResponseBodyClustersNetwork) *DescribeHubClustersResponseBodyClusters {
	s.Network = v
	return s
}

type DescribeHubClustersResponseBodyClustersApiServer struct {
	EnabledPublic  *bool   `json:"EnabledPublic,omitempty" xml:"EnabledPublic,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DescribeHubClustersResponseBodyClustersApiServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeHubClustersResponseBodyClustersApiServer) GoString() string {
	return s.String()
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
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterSpec  *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Profile      *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdateTime   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

type DescribeHubClustersResponseBodyClustersEndpoints struct {
	IntranetApiServerEndpoint *string `json:"IntranetApiServerEndpoint,omitempty" xml:"IntranetApiServerEndpoint,omitempty"`
	PublicApiServerEndpoint   *string `json:"PublicApiServerEndpoint,omitempty" xml:"PublicApiServerEndpoint,omitempty"`
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

type DescribeHubClustersResponseBodyClustersNetwork struct {
	ClusterDomain    *string   `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	SecurityGroupIDs []*string `json:"SecurityGroupIDs,omitempty" xml:"SecurityGroupIDs,omitempty" type:"Repeated"`
	VSwitches        []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VpcId            *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHubClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeHubClustersResponse) SetBody(v *DescribeHubClustersResponseBody) *DescribeHubClustersResponse {
	s.Body = v
	return s
}

type DescribeManagedClustersRequest struct {
	// 集群ID
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
	Clusters  []*DescribeManagedClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Cluster *DescribeManagedClustersResponseBodyClustersCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	Status  *DescribeManagedClustersResponseBodyClustersStatus  `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
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

func (s *DescribeManagedClustersResponseBodyClusters) SetStatus(v *DescribeManagedClustersResponseBodyClustersStatus) *DescribeManagedClustersResponseBodyClusters {
	s.Status = v
	return s
}

type DescribeManagedClustersResponseBodyClustersCluster struct {
	ClusterID       *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	ClusterSpec     *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	ClusterType     *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Created         *string `json:"Created,omitempty" xml:"Created,omitempty"`
	CurrentVersion  *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	InitVersion     *string `json:"InitVersion,omitempty" xml:"InitVersion,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Profile         *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Updated         *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	VSwitchID       *string `json:"VSwitchID,omitempty" xml:"VSwitchID,omitempty"`
	VpcID           *string `json:"VpcID,omitempty" xml:"VpcID,omitempty"`
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

type DescribeManagedClustersResponseBodyClustersStatus struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	State   *string `json:"State,omitempty" xml:"State,omitempty"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeManagedClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeManagedClustersResponse) SetBody(v *DescribeManagedClustersResponseBody) *DescribeManagedClustersResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Id of the request
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
	LocalName         *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint    *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionVpcEndpoint *string `json:"RegionVpcEndpoint,omitempty" xml:"RegionVpcEndpoint,omitempty"`
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

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionVpcEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionVpcEndpoint = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DetachClusterFromHubRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
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

type DetachClusterFromHubResponseBody struct {
	ClusterId         *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ManagedClusterIds []*string `json:"ManagedClusterIds,omitempty" xml:"ManagedClusterIds,omitempty" type:"Repeated"`
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type DetachClusterFromHubResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachClusterFromHubResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachClusterFromHubResponse) SetBody(v *DetachClusterFromHubResponseBody) *DetachClusterFromHubResponse {
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

	if !tea.BoolValue(util.IsUnset(request.AuditLogProject)) {
		body["AuditLogProject"] = request.AuditLogProject
	}

	if !tea.BoolValue(util.IsUnset(request.AuditLogStoreTTL)) {
		body["AuditLogStoreTTL"] = request.AuditLogStoreTTL
	}

	if !tea.BoolValue(util.IsUnset(request.ControlPlaneLogEnabled)) {
		body["ControlPlaneLogEnabled"] = request.ControlPlaneLogEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ControlPlaneLogProject)) {
		body["ControlPlaneLogProject"] = request.ControlPlaneLogProject
	}

	if !tea.BoolValue(util.IsUnset(request.ControlPlaneLogTTL)) {
		body["ControlPlaneLogTTL"] = request.ControlPlaneLogTTL
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnterpriseSecurityGroup)) {
		body["IsEnterpriseSecurityGroup"] = request.IsEnterpriseSecurityGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
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

func (client *Client) DeleteHubClusterWithOptions(request *DeleteHubClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteHubClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
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

func (client *Client) DescribeHubClustersWithOptions(runtime *util.RuntimeOptions) (_result *DescribeHubClustersResponse, _err error) {
	req := &openapi.OpenApiRequest{}
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

func (client *Client) DescribeHubClusters() (_result *DescribeHubClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHubClustersResponse{}
	_body, _err := client.DescribeHubClustersWithOptions(runtime)
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

func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
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

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
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
