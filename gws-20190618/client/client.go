// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SetClusterDnatRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	NatId     *string `json:"NatId,omitempty" xml:"NatId,omitempty" require:"true"`
	NatEip    *string `json:"NatEip,omitempty" xml:"NatEip,omitempty"`
}

func (s SetClusterDnatRequest) String() string {
	return tea.Prettify(s)
}

func (s SetClusterDnatRequest) GoString() string {
	return s.String()
}

func (s *SetClusterDnatRequest) SetClusterId(v string) *SetClusterDnatRequest {
	s.ClusterId = &v
	return s
}

func (s *SetClusterDnatRequest) SetNatId(v string) *SetClusterDnatRequest {
	s.NatId = &v
	return s
}

func (s *SetClusterDnatRequest) SetNatEip(v string) *SetClusterDnatRequest {
	s.NatEip = &v
	return s
}

type SetClusterDnatResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetClusterDnatResponse) String() string {
	return tea.Prettify(s)
}

func (s SetClusterDnatResponse) GoString() string {
	return s.String()
}

func (s *SetClusterDnatResponse) SetRequestId(v string) *SetClusterDnatResponse {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleRequest struct {
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

type CreateServiceLinkedRoleResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AlreadyExists *bool   `json:"AlreadyExists,omitempty" xml:"AlreadyExists,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetRequestId(v string) *CreateServiceLinkedRoleResponse {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetAlreadyExists(v bool) *CreateServiceLinkedRoleResponse {
	s.AlreadyExists = &v
	return s
}

type DescribeClusterConnectionsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeClusterConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionsRequest) SetClusterId(v string) *DescribeClusterConnectionsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterConnectionsRequest) SetStartTime(v string) *DescribeClusterConnectionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClusterConnectionsRequest) SetEndTime(v string) *DescribeClusterConnectionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClusterConnectionsRequest) SetTaskId(v string) *DescribeClusterConnectionsRequest {
	s.TaskId = &v
	return s
}

type DescribeClusterConnectionsResponse struct {
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskFinished *bool                                            `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty" require:"true"`
	TaskId       *string                                          `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	TotalCount   *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Connections  []*DescribeClusterConnectionsResponseConnections `json:"Connections,omitempty" xml:"Connections,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeClusterConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionsResponse) SetRequestId(v string) *DescribeClusterConnectionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterConnectionsResponse) SetTaskFinished(v bool) *DescribeClusterConnectionsResponse {
	s.TaskFinished = &v
	return s
}

func (s *DescribeClusterConnectionsResponse) SetTaskId(v string) *DescribeClusterConnectionsResponse {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterConnectionsResponse) SetTotalCount(v int64) *DescribeClusterConnectionsResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeClusterConnectionsResponse) SetConnections(v []*DescribeClusterConnectionsResponseConnections) *DescribeClusterConnectionsResponse {
	s.Connections = v
	return s
}

type DescribeClusterConnectionsResponseConnections struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	HostName     *string `json:"HostName,omitempty" xml:"HostName,omitempty" require:"true"`
	ClientName   *string `json:"ClientName,omitempty" xml:"ClientName,omitempty" require:"true"`
	LogonTime    *string `json:"LogonTime,omitempty" xml:"LogonTime,omitempty" require:"true"`
	LogoffTime   *string `json:"LogoffTime,omitempty" xml:"LogoffTime,omitempty" require:"true"`
	LogoffStatus *string `json:"LogoffStatus,omitempty" xml:"LogoffStatus,omitempty" require:"true"`
}

func (s DescribeClusterConnectionsResponseConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionsResponseConnections) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionsResponseConnections) SetInstanceId(v string) *DescribeClusterConnectionsResponseConnections {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterConnectionsResponseConnections) SetInstanceName(v string) *DescribeClusterConnectionsResponseConnections {
	s.InstanceName = &v
	return s
}

func (s *DescribeClusterConnectionsResponseConnections) SetHostName(v string) *DescribeClusterConnectionsResponseConnections {
	s.HostName = &v
	return s
}

func (s *DescribeClusterConnectionsResponseConnections) SetClientName(v string) *DescribeClusterConnectionsResponseConnections {
	s.ClientName = &v
	return s
}

func (s *DescribeClusterConnectionsResponseConnections) SetLogonTime(v string) *DescribeClusterConnectionsResponseConnections {
	s.LogonTime = &v
	return s
}

func (s *DescribeClusterConnectionsResponseConnections) SetLogoffTime(v string) *DescribeClusterConnectionsResponseConnections {
	s.LogoffTime = &v
	return s
}

func (s *DescribeClusterConnectionsResponseConnections) SetLogoffStatus(v string) *DescribeClusterConnectionsResponseConnections {
	s.LogoffStatus = &v
	return s
}

type DescribeClusterADDomainRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeClusterADDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterADDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterADDomainRequest) SetClusterId(v string) *DescribeClusterADDomainRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterADDomainRequest) SetTaskId(v string) *DescribeClusterADDomainRequest {
	s.TaskId = &v
	return s
}

type DescribeClusterADDomainResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsSupported  *bool   `json:"IsSupported,omitempty" xml:"IsSupported,omitempty" require:"true"`
	TaskFinished *bool   `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty" require:"true"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainDnsIp  *string `json:"DomainDnsIp,omitempty" xml:"DomainDnsIp,omitempty" require:"true"`
}

func (s DescribeClusterADDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterADDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterADDomainResponse) SetRequestId(v string) *DescribeClusterADDomainResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterADDomainResponse) SetIsSupported(v bool) *DescribeClusterADDomainResponse {
	s.IsSupported = &v
	return s
}

func (s *DescribeClusterADDomainResponse) SetTaskFinished(v bool) *DescribeClusterADDomainResponse {
	s.TaskFinished = &v
	return s
}

func (s *DescribeClusterADDomainResponse) SetTaskId(v string) *DescribeClusterADDomainResponse {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterADDomainResponse) SetDomainName(v string) *DescribeClusterADDomainResponse {
	s.DomainName = &v
	return s
}

func (s *DescribeClusterADDomainResponse) SetDomainDnsIp(v string) *DescribeClusterADDomainResponse {
	s.DomainDnsIp = &v
	return s
}

type SetClusterADDomainRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	DomainDnsIp    *string `json:"DomainDnsIp,omitempty" xml:"DomainDnsIp,omitempty" require:"true"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainPassword *string `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	DomainAdmin    *string `json:"DomainAdmin,omitempty" xml:"DomainAdmin,omitempty"`
	DomainDelete   *bool   `json:"DomainDelete,omitempty" xml:"DomainDelete,omitempty"`
}

func (s SetClusterADDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s SetClusterADDomainRequest) GoString() string {
	return s.String()
}

func (s *SetClusterADDomainRequest) SetClusterId(v string) *SetClusterADDomainRequest {
	s.ClusterId = &v
	return s
}

func (s *SetClusterADDomainRequest) SetDomainDnsIp(v string) *SetClusterADDomainRequest {
	s.DomainDnsIp = &v
	return s
}

func (s *SetClusterADDomainRequest) SetDomainName(v string) *SetClusterADDomainRequest {
	s.DomainName = &v
	return s
}

func (s *SetClusterADDomainRequest) SetDomainPassword(v string) *SetClusterADDomainRequest {
	s.DomainPassword = &v
	return s
}

func (s *SetClusterADDomainRequest) SetDomainAdmin(v string) *SetClusterADDomainRequest {
	s.DomainAdmin = &v
	return s
}

func (s *SetClusterADDomainRequest) SetDomainDelete(v bool) *SetClusterADDomainRequest {
	s.DomainDelete = &v
	return s
}

type SetClusterADDomainResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s SetClusterADDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s SetClusterADDomainResponse) GoString() string {
	return s.String()
}

func (s *SetClusterADDomainResponse) SetRequestId(v string) *SetClusterADDomainResponse {
	s.RequestId = &v
	return s
}

func (s *SetClusterADDomainResponse) SetTaskId(v string) *SetClusterADDomainResponse {
	s.TaskId = &v
	return s
}

type DescribeInstancePolicyRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	AsyncMode  *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
}

func (s DescribeInstancePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancePolicyRequest) SetInstanceId(v string) *DescribeInstancePolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancePolicyRequest) SetTaskId(v string) *DescribeInstancePolicyRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeInstancePolicyRequest) SetAsyncMode(v bool) *DescribeInstancePolicyRequest {
	s.AsyncMode = &v
	return s
}

type DescribeInstancePolicyResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VisualLossless *string `json:"VisualLossless,omitempty" xml:"VisualLossless,omitempty" require:"true"`
	OptimizeFor3d  *string `json:"OptimizeFor3d,omitempty" xml:"OptimizeFor3d,omitempty" require:"true"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	TaskFinished   *bool   `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty" require:"true"`
}

func (s DescribeInstancePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancePolicyResponse) SetRequestId(v string) *DescribeInstancePolicyResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancePolicyResponse) SetVisualLossless(v string) *DescribeInstancePolicyResponse {
	s.VisualLossless = &v
	return s
}

func (s *DescribeInstancePolicyResponse) SetOptimizeFor3d(v string) *DescribeInstancePolicyResponse {
	s.OptimizeFor3d = &v
	return s
}

func (s *DescribeInstancePolicyResponse) SetTaskId(v string) *DescribeInstancePolicyResponse {
	s.TaskId = &v
	return s
}

func (s *DescribeInstancePolicyResponse) SetTaskFinished(v bool) *DescribeInstancePolicyResponse {
	s.TaskFinished = &v
	return s
}

type SetInstancePolicyRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VisualLossless *string `json:"VisualLossless,omitempty" xml:"VisualLossless,omitempty"`
	OptimizeFor3d  *string `json:"OptimizeFor3d,omitempty" xml:"OptimizeFor3d,omitempty"`
	AsyncMode      *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
}

func (s SetInstancePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstancePolicyRequest) GoString() string {
	return s.String()
}

func (s *SetInstancePolicyRequest) SetInstanceId(v string) *SetInstancePolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstancePolicyRequest) SetVisualLossless(v string) *SetInstancePolicyRequest {
	s.VisualLossless = &v
	return s
}

func (s *SetInstancePolicyRequest) SetOptimizeFor3d(v string) *SetInstancePolicyRequest {
	s.OptimizeFor3d = &v
	return s
}

func (s *SetInstancePolicyRequest) SetAsyncMode(v bool) *SetInstancePolicyRequest {
	s.AsyncMode = &v
	return s
}

type SetInstancePolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s SetInstancePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *SetInstancePolicyResponse) SetRequestId(v string) *SetInstancePolicyResponse {
	s.RequestId = &v
	return s
}

func (s *SetInstancePolicyResponse) SetTaskId(v string) *SetInstancePolicyResponse {
	s.TaskId = &v
	return s
}

type DescribeAvailableResourceRequest struct {
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

type DescribeAvailableResourceResponse struct {
	RequestId          *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AvailableResources []*DescribeAvailableResourceResponseAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetRequestId(v string) *DescribeAvailableResourceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponse) SetAvailableResources(v []*DescribeAvailableResourceResponseAvailableResources) *DescribeAvailableResourceResponse {
	s.AvailableResources = v
	return s
}

type DescribeAvailableResourceResponseAvailableResources struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	Zone        *string `json:"Zone,omitempty" xml:"Zone,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponseAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseAvailableResources) SetClusterType(v string) *DescribeAvailableResourceResponseAvailableResources {
	s.ClusterType = &v
	return s
}

func (s *DescribeAvailableResourceResponseAvailableResources) SetZone(v string) *DescribeAvailableResourceResponseAvailableResources {
	s.Zone = &v
	return s
}

type SetClusterPolicyRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty" require:"true"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty" require:"true"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty" require:"true"`
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty" require:"true"`
	UdpPort     *string `json:"UdpPort,omitempty" xml:"UdpPort,omitempty"`
	DomainList  *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	AsyncMode   *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
}

func (s SetClusterPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetClusterPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetClusterPolicyRequest) SetClusterId(v string) *SetClusterPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *SetClusterPolicyRequest) SetUsbRedirect(v string) *SetClusterPolicyRequest {
	s.UsbRedirect = &v
	return s
}

func (s *SetClusterPolicyRequest) SetWatermark(v string) *SetClusterPolicyRequest {
	s.Watermark = &v
	return s
}

func (s *SetClusterPolicyRequest) SetLocalDrive(v string) *SetClusterPolicyRequest {
	s.LocalDrive = &v
	return s
}

func (s *SetClusterPolicyRequest) SetClipboard(v string) *SetClusterPolicyRequest {
	s.Clipboard = &v
	return s
}

func (s *SetClusterPolicyRequest) SetUdpPort(v string) *SetClusterPolicyRequest {
	s.UdpPort = &v
	return s
}

func (s *SetClusterPolicyRequest) SetDomainList(v string) *SetClusterPolicyRequest {
	s.DomainList = &v
	return s
}

func (s *SetClusterPolicyRequest) SetAsyncMode(v bool) *SetClusterPolicyRequest {
	s.AsyncMode = &v
	return s
}

type SetClusterPolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s SetClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetClusterPolicyResponse) SetRequestId(v string) *SetClusterPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *SetClusterPolicyResponse) SetTaskId(v string) *SetClusterPolicyResponse {
	s.TaskId = &v
	return s
}

type DescribeClusterPolicyRequest struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	AsyncMode *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
}

func (s DescribeClusterPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterPolicyRequest) SetTaskId(v string) *DescribeClusterPolicyRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterPolicyRequest) SetAsyncMode(v bool) *DescribeClusterPolicyRequest {
	s.AsyncMode = &v
	return s
}

func (s *DescribeClusterPolicyRequest) SetClusterId(v string) *DescribeClusterPolicyRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterPolicyResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UsbRedirect  *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty" require:"true"`
	Watermark    *string `json:"Watermark,omitempty" xml:"Watermark,omitempty" require:"true"`
	LocalDrive   *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty" require:"true"`
	Clipboard    *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty" require:"true"`
	UdpPort      *string `json:"UdpPort,omitempty" xml:"UdpPort,omitempty" require:"true"`
	DomainList   *string `json:"DomainList,omitempty" xml:"DomainList,omitempty" require:"true"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	TaskFinished *bool   `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty" require:"true"`
}

func (s DescribeClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterPolicyResponse) SetRequestId(v string) *DescribeClusterPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterPolicyResponse) SetUsbRedirect(v string) *DescribeClusterPolicyResponse {
	s.UsbRedirect = &v
	return s
}

func (s *DescribeClusterPolicyResponse) SetWatermark(v string) *DescribeClusterPolicyResponse {
	s.Watermark = &v
	return s
}

func (s *DescribeClusterPolicyResponse) SetLocalDrive(v string) *DescribeClusterPolicyResponse {
	s.LocalDrive = &v
	return s
}

func (s *DescribeClusterPolicyResponse) SetClipboard(v string) *DescribeClusterPolicyResponse {
	s.Clipboard = &v
	return s
}

func (s *DescribeClusterPolicyResponse) SetUdpPort(v string) *DescribeClusterPolicyResponse {
	s.UdpPort = &v
	return s
}

func (s *DescribeClusterPolicyResponse) SetDomainList(v string) *DescribeClusterPolicyResponse {
	s.DomainList = &v
	return s
}

func (s *DescribeClusterPolicyResponse) SetTaskId(v string) *DescribeClusterPolicyResponse {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterPolicyResponse) SetTaskFinished(v bool) *DescribeClusterPolicyResponse {
	s.TaskFinished = &v
	return s
}

type SetInstanceNameRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s SetInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceNameRequest) SetInstanceId(v string) *SetInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstanceNameRequest) SetName(v string) *SetInstanceNameRequest {
	s.Name = &v
	return s
}

type SetInstanceNameResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceNameResponse) SetRequestId(v string) *SetInstanceNameResponse {
	s.RequestId = &v
	return s
}

type SetImageNameRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s SetImageNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetImageNameRequest) GoString() string {
	return s.String()
}

func (s *SetImageNameRequest) SetImageId(v string) *SetImageNameRequest {
	s.ImageId = &v
	return s
}

func (s *SetImageNameRequest) SetName(v string) *SetImageNameRequest {
	s.Name = &v
	return s
}

type SetImageNameResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetImageNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetImageNameResponse) GoString() string {
	return s.String()
}

func (s *SetImageNameResponse) SetRequestId(v string) *SetImageNameResponse {
	s.RequestId = &v
	return s
}

type SetClusterNameRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s SetClusterNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetClusterNameRequest) GoString() string {
	return s.String()
}

func (s *SetClusterNameRequest) SetClusterId(v string) *SetClusterNameRequest {
	s.ClusterId = &v
	return s
}

func (s *SetClusterNameRequest) SetName(v string) *SetClusterNameRequest {
	s.Name = &v
	return s
}

type SetClusterNameResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetClusterNameResponse) GoString() string {
	return s.String()
}

func (s *SetClusterNameResponse) SetRequestId(v string) *SetClusterNameResponse {
	s.RequestId = &v
	return s
}

type StopInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetInstanceId(v string) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

type StopInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetRequestId(v string) *StopInstanceResponse {
	s.RequestId = &v
	return s
}

type StartInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

type StartInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetRequestId(v string) *StartInstanceResponse {
	s.RequestId = &v
	return s
}

type SetInstanceUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	UserUid    *int64  `json:"UserUid,omitempty" xml:"UserUid,omitempty" require:"true"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
}

func (s SetInstanceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceUserRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceUserRequest) SetInstanceId(v string) *SetInstanceUserRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstanceUserRequest) SetUserUid(v int64) *SetInstanceUserRequest {
	s.UserUid = &v
	return s
}

func (s *SetInstanceUserRequest) SetUserName(v string) *SetInstanceUserRequest {
	s.UserName = &v
	return s
}

type SetInstanceUserResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetInstanceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceUserResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceUserResponse) SetRequestId(v string) *SetInstanceUserResponse {
	s.RequestId = &v
	return s
}

type RestartInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s RestartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) SetInstanceId(v string) *RestartInstanceRequest {
	s.InstanceId = &v
	return s
}

type RestartInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RestartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) SetRequestId(v string) *RestartInstanceResponse {
	s.RequestId = &v
	return s
}

type IsUserAdminRequest struct {
}

func (s IsUserAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s IsUserAdminRequest) GoString() string {
	return s.String()
}

type IsUserAdminResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsAdmin   *bool   `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty" require:"true"`
	IsAllow   *bool   `json:"IsAllow,omitempty" xml:"IsAllow,omitempty" require:"true"`
}

func (s IsUserAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s IsUserAdminResponse) GoString() string {
	return s.String()
}

func (s *IsUserAdminResponse) SetRequestId(v string) *IsUserAdminResponse {
	s.RequestId = &v
	return s
}

func (s *IsUserAdminResponse) SetIsAdmin(v bool) *IsUserAdminResponse {
	s.IsAdmin = &v
	return s
}

func (s *IsUserAdminResponse) SetIsAllow(v bool) *IsUserAdminResponse {
	s.IsAllow = &v
	return s
}

type GetConnectTicketRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password   *string `json:"Password,omitempty" xml:"Password,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	AsyncMode  *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
}

func (s GetConnectTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectTicketRequest) SetInstanceId(v string) *GetConnectTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConnectTicketRequest) SetAppName(v string) *GetConnectTicketRequest {
	s.AppName = &v
	return s
}

func (s *GetConnectTicketRequest) SetUserName(v string) *GetConnectTicketRequest {
	s.UserName = &v
	return s
}

func (s *GetConnectTicketRequest) SetPassword(v string) *GetConnectTicketRequest {
	s.Password = &v
	return s
}

func (s *GetConnectTicketRequest) SetTaskId(v string) *GetConnectTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectTicketRequest) SetAsyncMode(v bool) *GetConnectTicketRequest {
	s.AsyncMode = &v
	return s
}

type GetConnectTicketResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Ticket       *string `json:"Ticket,omitempty" xml:"Ticket,omitempty" require:"true"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	TaskFinished *bool   `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty" require:"true"`
}

func (s GetConnectTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectTicketResponse) SetRequestId(v string) *GetConnectTicketResponse {
	s.RequestId = &v
	return s
}

func (s *GetConnectTicketResponse) SetTicket(v string) *GetConnectTicketResponse {
	s.Ticket = &v
	return s
}

func (s *GetConnectTicketResponse) SetTaskId(v string) *GetConnectTicketResponse {
	s.TaskId = &v
	return s
}

func (s *GetConnectTicketResponse) SetTaskFinished(v bool) *GetConnectTicketResponse {
	s.TaskFinished = &v
	return s
}

type DescribeInstancesRequest struct {
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserUid    *int64  `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetPageNumber(v int64) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int64) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetClusterId(v string) *DescribeInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetUserUid(v int64) *DescribeInstancesRequest {
	s.UserUid = &v
	return s
}

func (s *DescribeInstancesRequest) SetUserName(v string) *DescribeInstancesRequest {
	s.UserName = &v
	return s
}

type DescribeInstancesResponse struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int64                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Instances  []*DescribeInstancesResponseInstances `json:"Instances,omitempty" xml:"Instances,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetRequestId(v string) *DescribeInstancesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponse) SetTotalCount(v int64) *DescribeInstancesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponse) SetPageNumber(v int64) *DescribeInstancesResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponse) SetPageSize(v int64) *DescribeInstancesResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponse) SetInstances(v []*DescribeInstancesResponseInstances) *DescribeInstancesResponse {
	s.Instances = v
	return s
}

type DescribeInstancesResponseInstances struct {
	InstanceId         *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Name               *string                                      `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ClusterId          *string                                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Status             *string                                      `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	WorkMode           *string                                      `json:"WorkMode,omitempty" xml:"WorkMode,omitempty" require:"true"`
	StoppedMode        *string                                      `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty" require:"true"`
	ImageId            *string                                      `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	InstanceChargeType *string                                      `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty" require:"true"`
	InstanceType       *string                                      `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" require:"true"`
	CreateTime         *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ExpireTime         *string                                      `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	UserUid            *int64                                       `json:"UserUid,omitempty" xml:"UserUid,omitempty" require:"true"`
	UserName           *string                                      `json:"UserName,omitempty" xml:"UserName,omitempty" require:"true"`
	DomainName         *string                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	MaxBandwidthIn     *int64                                       `json:"MaxBandwidthIn,omitempty" xml:"MaxBandwidthIn,omitempty" require:"true"`
	MaxBandwidthOut    *int64                                       `json:"MaxBandwidthOut,omitempty" xml:"MaxBandwidthOut,omitempty" require:"true"`
	IsBoundUser        *bool                                        `json:"IsBoundUser,omitempty" xml:"IsBoundUser,omitempty" require:"true"`
	AppList            []*DescribeInstancesResponseInstancesAppList `json:"AppList,omitempty" xml:"AppList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponseInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstances) SetInstanceId(v string) *DescribeInstancesResponseInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetName(v string) *DescribeInstancesResponseInstances {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetClusterId(v string) *DescribeInstancesResponseInstances {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetStatus(v string) *DescribeInstancesResponseInstances {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetWorkMode(v string) *DescribeInstancesResponseInstances {
	s.WorkMode = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetStoppedMode(v string) *DescribeInstancesResponseInstances {
	s.StoppedMode = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetImageId(v string) *DescribeInstancesResponseInstances {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetInstanceChargeType(v string) *DescribeInstancesResponseInstances {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetInstanceType(v string) *DescribeInstancesResponseInstances {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetCreateTime(v string) *DescribeInstancesResponseInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetExpireTime(v string) *DescribeInstancesResponseInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetUserUid(v int64) *DescribeInstancesResponseInstances {
	s.UserUid = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetUserName(v string) *DescribeInstancesResponseInstances {
	s.UserName = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetDomainName(v string) *DescribeInstancesResponseInstances {
	s.DomainName = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetMaxBandwidthIn(v int64) *DescribeInstancesResponseInstances {
	s.MaxBandwidthIn = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetMaxBandwidthOut(v int64) *DescribeInstancesResponseInstances {
	s.MaxBandwidthOut = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetIsBoundUser(v bool) *DescribeInstancesResponseInstances {
	s.IsBoundUser = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetAppList(v []*DescribeInstancesResponseInstancesAppList) *DescribeInstancesResponseInstances {
	s.AppList = v
	return s
}

type DescribeInstancesResponseInstancesAppList struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	AppPath *string `json:"AppPath,omitempty" xml:"AppPath,omitempty" require:"true"`
	AppArgs *string `json:"AppArgs,omitempty" xml:"AppArgs,omitempty" require:"true"`
}

func (s DescribeInstancesResponseInstancesAppList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstancesAppList) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstancesAppList) SetAppName(v string) *DescribeInstancesResponseInstancesAppList {
	s.AppName = &v
	return s
}

func (s *DescribeInstancesResponseInstancesAppList) SetAppPath(v string) *DescribeInstancesResponseInstancesAppList {
	s.AppPath = &v
	return s
}

func (s *DescribeInstancesResponseInstancesAppList) SetAppArgs(v string) *DescribeInstancesResponseInstancesAppList {
	s.AppArgs = &v
	return s
}

type DescribeImagesRequest struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) SetInstanceType(v string) *DescribeImagesRequest {
	s.InstanceType = &v
	return s
}

type DescribeImagesResponse struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int64                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Images     []*DescribeImagesResponseImages `json:"Images,omitempty" xml:"Images,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) SetRequestId(v string) *DescribeImagesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponse) SetTotalCount(v int64) *DescribeImagesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeImagesResponse) SetPageNumber(v int64) *DescribeImagesResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesResponse) SetPageSize(v int64) *DescribeImagesResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeImagesResponse) SetImages(v []*DescribeImagesResponseImages) *DescribeImagesResponse {
	s.Images = v
	return s
}

type DescribeImagesResponseImages struct {
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	Progress    *string `json:"Progress,omitempty" xml:"Progress,omitempty" require:"true"`
	ImageType   *string `json:"ImageType,omitempty" xml:"ImageType,omitempty" require:"true"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty" require:"true"`
}

func (s DescribeImagesResponseImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseImages) SetImageId(v string) *DescribeImagesResponseImages {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseImages) SetName(v string) *DescribeImagesResponseImages {
	s.Name = &v
	return s
}

func (s *DescribeImagesResponseImages) SetSize(v int64) *DescribeImagesResponseImages {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseImages) SetStatus(v string) *DescribeImagesResponseImages {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseImages) SetCreateTime(v string) *DescribeImagesResponseImages {
	s.CreateTime = &v
	return s
}

func (s *DescribeImagesResponseImages) SetProgress(v string) *DescribeImagesResponseImages {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseImages) SetImageType(v string) *DescribeImagesResponseImages {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesResponseImages) SetProductCode(v string) *DescribeImagesResponseImages {
	s.ProductCode = &v
	return s
}

type DescribeClustersRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequest) SetClusterId(v string) *DescribeClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersRequest) SetPageNumber(v int64) *DescribeClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersRequest) SetPageSize(v int64) *DescribeClustersRequest {
	s.PageSize = &v
	return s
}

type DescribeClustersResponse struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int64                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int64                              `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Clusters   []*DescribeClustersResponseClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) SetRequestId(v string) *DescribeClustersResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeClustersResponse) SetTotalCount(v int64) *DescribeClustersResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeClustersResponse) SetPageNumber(v int64) *DescribeClustersResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersResponse) SetPageSize(v int64) *DescribeClustersResponse {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersResponse) SetClusters(v []*DescribeClustersResponseClusters) *DescribeClustersResponse {
	s.Clusters = v
	return s
}

type DescribeClustersResponseClusters struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" require:"true"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	NatId         *string `json:"NatId,omitempty" xml:"NatId,omitempty" require:"true"`
	NatEip        *string `json:"NatEip,omitempty" xml:"NatEip,omitempty" require:"true"`
	InstanceCount *int64  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty" require:"true"`
}

func (s DescribeClustersResponseClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseClusters) SetClusterId(v string) *DescribeClustersResponseClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersResponseClusters) SetName(v string) *DescribeClustersResponseClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersResponseClusters) SetStatus(v string) *DescribeClustersResponseClusters {
	s.Status = &v
	return s
}

func (s *DescribeClustersResponseClusters) SetVpcId(v string) *DescribeClustersResponseClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersResponseClusters) SetCreateTime(v string) *DescribeClustersResponseClusters {
	s.CreateTime = &v
	return s
}

func (s *DescribeClustersResponseClusters) SetSecurityGroup(v string) *DescribeClustersResponseClusters {
	s.SecurityGroup = &v
	return s
}

func (s *DescribeClustersResponseClusters) SetDomainName(v string) *DescribeClustersResponseClusters {
	s.DomainName = &v
	return s
}

func (s *DescribeClustersResponseClusters) SetNatId(v string) *DescribeClustersResponseClusters {
	s.NatId = &v
	return s
}

func (s *DescribeClustersResponseClusters) SetNatEip(v string) *DescribeClustersResponseClusters {
	s.NatEip = &v
	return s
}

func (s *DescribeClustersResponseClusters) SetInstanceCount(v int64) *DescribeClustersResponseClusters {
	s.InstanceCount = &v
	return s
}

type DeleteInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

type DeleteInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetRequestId(v string) *DeleteInstanceResponse {
	s.RequestId = &v
	return s
}

type DeleteImageRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetImageId(v string) *DeleteImageRequest {
	s.ImageId = &v
	return s
}

type DeleteImageResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetRequestId(v string) *DeleteImageResponse {
	s.RequestId = &v
	return s
}

type DeleteClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

type DeleteClusterResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetRequestId(v string) *DeleteClusterResponse {
	s.RequestId = &v
	return s
}

type CreateInstanceRequest struct {
	ClusterId               *string                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	VSwitchId               *string                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Name                    *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	ImageId                 *string                         `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	SystemDiskSize          *int                            `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	SystemDiskCategory      *string                         `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" require:"true"`
	AllocatePublicAddress   *string                         `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	InternetChargeType      *string                         `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn  *int                            `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut *int                            `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	InstanceType            *string                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" require:"true"`
	InstanceChargeType      *string                         `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	AutoRenew               *string                         `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	Period                  *int                            `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string                         `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	WorkMode                *string                         `json:"WorkMode,omitempty" xml:"WorkMode,omitempty" require:"true"`
	AppList                 []*CreateInstanceRequestAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetClusterId(v string) *CreateInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetSystemDiskSize(v int) *CreateInstanceRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateInstanceRequest) SetSystemDiskCategory(v string) *CreateInstanceRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateInstanceRequest) SetAllocatePublicAddress(v string) *CreateInstanceRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetChargeType(v string) *CreateInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetMaxBandwidthIn(v int) *CreateInstanceRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetMaxBandwidthOut(v int) *CreateInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceChargeType(v string) *CreateInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v string) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodUnit(v string) *CreateInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetWorkMode(v string) *CreateInstanceRequest {
	s.WorkMode = &v
	return s
}

func (s *CreateInstanceRequest) SetAppList(v []*CreateInstanceRequestAppList) *CreateInstanceRequest {
	s.AppList = v
	return s
}

type CreateInstanceRequestAppList struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPath *string `json:"AppPath,omitempty" xml:"AppPath,omitempty"`
	AppArgs *string `json:"AppArgs,omitempty" xml:"AppArgs,omitempty"`
}

func (s CreateInstanceRequestAppList) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestAppList) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestAppList) SetAppName(v string) *CreateInstanceRequestAppList {
	s.AppName = &v
	return s
}

func (s *CreateInstanceRequestAppList) SetAppPath(v string) *CreateInstanceRequestAppList {
	s.AppPath = &v
	return s
}

func (s *CreateInstanceRequestAppList) SetAppArgs(v string) *CreateInstanceRequestAppList {
	s.AppArgs = &v
	return s
}

type CreateInstanceResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetRequestId(v string) *CreateInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponse) SetInstanceId(v string) *CreateInstanceResponse {
	s.InstanceId = &v
	return s
}

type CreateImageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetInstanceId(v string) *CreateImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateImageRequest) SetName(v string) *CreateImageRequest {
	s.Name = &v
	return s
}

type CreateImageResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetRequestId(v string) *CreateImageResponse {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponse) SetImageId(v string) *CreateImageResponse {
	s.ImageId = &v
	return s
}

type CreateClusterRequest struct {
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

type CreateClusterResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetRequestId(v string) *CreateClusterResponse {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponse) SetClusterId(v string) *CreateClusterResponse {
	s.ClusterId = &v
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
		"ap-southeast-3":      tea.String("gws.ap-northeast-3.aliyuncs.com"),
		"cn-hangzhou-finance": tea.String("ecd.cn-hangzhou-finance.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("gws"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) SetClusterDnatWithOptions(request *SetClusterDnatRequest, runtime *util.RuntimeOptions) (_result *SetClusterDnatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetClusterDnatResponse{}
	_body, _err := client.DoRequest(tea.String("SetClusterDnat"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetClusterDnat(request *SetClusterDnatRequest) (_result *SetClusterDnatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetClusterDnatResponse{}
	_body, _err := client.SetClusterDnatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.DoRequest(tea.String("CreateServiceLinkedRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterConnectionsWithOptions(request *DescribeClusterConnectionsRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterConnectionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeClusterConnections"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterConnections(request *DescribeClusterConnectionsRequest) (_result *DescribeClusterConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterConnectionsResponse{}
	_body, _err := client.DescribeClusterConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterADDomainWithOptions(request *DescribeClusterADDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterADDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterADDomainResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeClusterADDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterADDomain(request *DescribeClusterADDomainRequest) (_result *DescribeClusterADDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterADDomainResponse{}
	_body, _err := client.DescribeClusterADDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetClusterADDomainWithOptions(request *SetClusterADDomainRequest, runtime *util.RuntimeOptions) (_result *SetClusterADDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetClusterADDomainResponse{}
	_body, _err := client.DoRequest(tea.String("SetClusterADDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetClusterADDomain(request *SetClusterADDomainRequest) (_result *SetClusterADDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetClusterADDomainResponse{}
	_body, _err := client.SetClusterADDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancePolicyWithOptions(request *DescribeInstancePolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstancePolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstancePolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstancePolicy(request *DescribeInstancePolicyRequest) (_result *DescribeInstancePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancePolicyResponse{}
	_body, _err := client.DescribeInstancePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetInstancePolicyWithOptions(request *SetInstancePolicyRequest, runtime *util.RuntimeOptions) (_result *SetInstancePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetInstancePolicyResponse{}
	_body, _err := client.DoRequest(tea.String("SetInstancePolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetInstancePolicy(request *SetInstancePolicyRequest) (_result *SetInstancePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetInstancePolicyResponse{}
	_body, _err := client.SetInstancePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAvailableResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetClusterPolicyWithOptions(request *SetClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *SetClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetClusterPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("SetClusterPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetClusterPolicy(request *SetClusterPolicyRequest) (_result *SetClusterPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetClusterPolicyResponse{}
	_body, _err := client.SetClusterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterPolicyWithOptions(request *DescribeClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeClusterPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterPolicy(request *DescribeClusterPolicyRequest) (_result *DescribeClusterPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterPolicyResponse{}
	_body, _err := client.DescribeClusterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetInstanceNameWithOptions(request *SetInstanceNameRequest, runtime *util.RuntimeOptions) (_result *SetInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetInstanceNameResponse{}
	_body, _err := client.DoRequest(tea.String("SetInstanceName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetInstanceName(request *SetInstanceNameRequest) (_result *SetInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetInstanceNameResponse{}
	_body, _err := client.SetInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetImageNameWithOptions(request *SetImageNameRequest, runtime *util.RuntimeOptions) (_result *SetImageNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetImageNameResponse{}
	_body, _err := client.DoRequest(tea.String("SetImageName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetImageName(request *SetImageNameRequest) (_result *SetImageNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetImageNameResponse{}
	_body, _err := client.SetImageNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetClusterNameWithOptions(request *SetClusterNameRequest, runtime *util.RuntimeOptions) (_result *SetClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetClusterNameResponse{}
	_body, _err := client.DoRequest(tea.String("SetClusterName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetClusterName(request *SetClusterNameRequest) (_result *SetClusterNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetClusterNameResponse{}
	_body, _err := client.SetClusterNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("StopInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("StartInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetInstanceUserWithOptions(request *SetInstanceUserRequest, runtime *util.RuntimeOptions) (_result *SetInstanceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetInstanceUserResponse{}
	_body, _err := client.DoRequest(tea.String("SetInstanceUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetInstanceUser(request *SetInstanceUserRequest) (_result *SetInstanceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetInstanceUserResponse{}
	_body, _err := client.SetInstanceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("RestartInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartInstance(request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IsUserAdminWithOptions(request *IsUserAdminRequest, runtime *util.RuntimeOptions) (_result *IsUserAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &IsUserAdminResponse{}
	_body, _err := client.DoRequest(tea.String("IsUserAdmin"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsUserAdmin(request *IsUserAdminRequest) (_result *IsUserAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IsUserAdminResponse{}
	_body, _err := client.IsUserAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConnectTicketWithOptions(request *GetConnectTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetConnectTicketResponse{}
	_body, _err := client.DoRequest(tea.String("GetConnectTicket"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConnectTicket(request *GetConnectTicketRequest) (_result *GetConnectTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectTicketResponse{}
	_body, _err := client.GetConnectTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeImages"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImages(request *DescribeImagesRequest) (_result *DescribeImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DescribeImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClustersWithOptions(request *DescribeClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeClusters"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusters(request *DescribeClustersRequest) (_result *DescribeClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DescribeClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteCluster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("CreateInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateImageResponse{}
	_body, _err := client.DoRequest(tea.String("CreateImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoRequest(tea.String("CreateCluster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-06-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
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
