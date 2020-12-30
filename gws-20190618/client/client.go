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

type CreateClusterRequest struct {
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
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

type CreateClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateImageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type CreateImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponseBody) SetImageId(v string) *CreateImageResponseBody {
	s.ImageId = &v
	return s
}

type CreateImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetHeaders(v map[string]*string) *CreateImageResponse {
	s.Headers = v
	return s
}

func (s *CreateImageResponse) SetBody(v *CreateImageResponseBody) *CreateImageResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	ClusterId               *string                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	VSwitchId               *string                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Name                    *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	ImageId                 *string                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	SystemDiskSize          *int32                          `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	SystemDiskCategory      *string                         `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	AllocatePublicAddress   *string                         `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	InternetChargeType      *string                         `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn  *int32                          `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut *int32                          `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	InstanceType            *string                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceChargeType      *string                         `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	AutoRenew               *string                         `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	Period                  *int32                          `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string                         `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	WorkMode                *string                         `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
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

func (s *CreateInstanceRequest) SetSystemDiskSize(v int32) *CreateInstanceRequest {
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

func (s *CreateInstanceRequest) SetInternetMaxBandwidthIn(v int32) *CreateInstanceRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateInstanceRequest {
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

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
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

type CreateInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
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

type CreateServiceLinkedRoleResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AlreadyExists *bool   `json:"AlreadyExists,omitempty" xml:"AlreadyExists,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetAlreadyExists(v bool) *CreateServiceLinkedRoleResponseBody {
	s.AlreadyExists = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

type DeleteClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
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

type DeleteImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
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

type DescribeAvailableResourceResponseBody struct {
	RequestId          *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AvailableResources []*DescribeAvailableResourceResponseBodyAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableResources(v []*DescribeAvailableResourceResponseBodyAvailableResources) *DescribeAvailableResourceResponseBody {
	s.AvailableResources = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableResources struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Zone        *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableResources) SetClusterType(v string) *DescribeAvailableResourceResponseBodyAvailableResources {
	s.ClusterType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableResources) SetZone(v string) *DescribeAvailableResourceResponseBodyAvailableResources {
	s.Zone = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeClusterADDomainRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

type DescribeClusterADDomainResponseBody struct {
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	TaskFinished *bool   `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty"`
	IsSupported  *bool   `json:"IsSupported,omitempty" xml:"IsSupported,omitempty"`
	DomainDnsIp  *string `json:"DomainDnsIp,omitempty" xml:"DomainDnsIp,omitempty"`
}

func (s DescribeClusterADDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterADDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterADDomainResponseBody) SetTaskId(v string) *DescribeClusterADDomainResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterADDomainResponseBody) SetRequestId(v string) *DescribeClusterADDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterADDomainResponseBody) SetDomainName(v string) *DescribeClusterADDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeClusterADDomainResponseBody) SetTaskFinished(v bool) *DescribeClusterADDomainResponseBody {
	s.TaskFinished = &v
	return s
}

func (s *DescribeClusterADDomainResponseBody) SetIsSupported(v bool) *DescribeClusterADDomainResponseBody {
	s.IsSupported = &v
	return s
}

func (s *DescribeClusterADDomainResponseBody) SetDomainDnsIp(v string) *DescribeClusterADDomainResponseBody {
	s.DomainDnsIp = &v
	return s
}

type DescribeClusterADDomainResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterADDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterADDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterADDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterADDomainResponse) SetHeaders(v map[string]*string) *DescribeClusterADDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterADDomainResponse) SetBody(v *DescribeClusterADDomainResponseBody) *DescribeClusterADDomainResponse {
	s.Body = v
	return s
}

type DescribeClusterConnectionsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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

type DescribeClusterConnectionsResponseBody struct {
	Connections  []*DescribeClusterConnectionsResponseBodyConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	TotalCount   *int64                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TaskId       *string                                              `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskFinished *bool                                                `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty"`
}

func (s DescribeClusterConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionsResponseBody) SetConnections(v []*DescribeClusterConnectionsResponseBodyConnections) *DescribeClusterConnectionsResponseBody {
	s.Connections = v
	return s
}

func (s *DescribeClusterConnectionsResponseBody) SetTotalCount(v int64) *DescribeClusterConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeClusterConnectionsResponseBody) SetTaskId(v string) *DescribeClusterConnectionsResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterConnectionsResponseBody) SetRequestId(v string) *DescribeClusterConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterConnectionsResponseBody) SetTaskFinished(v bool) *DescribeClusterConnectionsResponseBody {
	s.TaskFinished = &v
	return s
}

type DescribeClusterConnectionsResponseBodyConnections struct {
	LogoffStatus *string `json:"LogoffStatus,omitempty" xml:"LogoffStatus,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	LogonTime    *string `json:"LogonTime,omitempty" xml:"LogonTime,omitempty"`
	HostName     *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	LogoffTime   *string `json:"LogoffTime,omitempty" xml:"LogoffTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ClientName   *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
}

func (s DescribeClusterConnectionsResponseBodyConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionsResponseBodyConnections) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionsResponseBodyConnections) SetLogoffStatus(v string) *DescribeClusterConnectionsResponseBodyConnections {
	s.LogoffStatus = &v
	return s
}

func (s *DescribeClusterConnectionsResponseBodyConnections) SetInstanceName(v string) *DescribeClusterConnectionsResponseBodyConnections {
	s.InstanceName = &v
	return s
}

func (s *DescribeClusterConnectionsResponseBodyConnections) SetLogonTime(v string) *DescribeClusterConnectionsResponseBodyConnections {
	s.LogonTime = &v
	return s
}

func (s *DescribeClusterConnectionsResponseBodyConnections) SetHostName(v string) *DescribeClusterConnectionsResponseBodyConnections {
	s.HostName = &v
	return s
}

func (s *DescribeClusterConnectionsResponseBodyConnections) SetLogoffTime(v string) *DescribeClusterConnectionsResponseBodyConnections {
	s.LogoffTime = &v
	return s
}

func (s *DescribeClusterConnectionsResponseBodyConnections) SetInstanceId(v string) *DescribeClusterConnectionsResponseBodyConnections {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterConnectionsResponseBodyConnections) SetClientName(v string) *DescribeClusterConnectionsResponseBodyConnections {
	s.ClientName = &v
	return s
}

type DescribeClusterConnectionsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionsResponse) SetHeaders(v map[string]*string) *DescribeClusterConnectionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterConnectionsResponse) SetBody(v *DescribeClusterConnectionsResponseBody) *DescribeClusterConnectionsResponse {
	s.Body = v
	return s
}

type DescribeClusterPolicyRequest struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	AsyncMode *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

type DescribeClusterPolicyResponseBody struct {
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UdpPort      *string `json:"UdpPort,omitempty" xml:"UdpPort,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LocalDrive   *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect  *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	TaskFinished *bool   `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty"`
	Clipboard    *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	DomainList   *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	Watermark    *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s DescribeClusterPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterPolicyResponseBody) SetTaskId(v string) *DescribeClusterPolicyResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterPolicyResponseBody) SetUdpPort(v string) *DescribeClusterPolicyResponseBody {
	s.UdpPort = &v
	return s
}

func (s *DescribeClusterPolicyResponseBody) SetRequestId(v string) *DescribeClusterPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterPolicyResponseBody) SetLocalDrive(v string) *DescribeClusterPolicyResponseBody {
	s.LocalDrive = &v
	return s
}

func (s *DescribeClusterPolicyResponseBody) SetUsbRedirect(v string) *DescribeClusterPolicyResponseBody {
	s.UsbRedirect = &v
	return s
}

func (s *DescribeClusterPolicyResponseBody) SetTaskFinished(v bool) *DescribeClusterPolicyResponseBody {
	s.TaskFinished = &v
	return s
}

func (s *DescribeClusterPolicyResponseBody) SetClipboard(v string) *DescribeClusterPolicyResponseBody {
	s.Clipboard = &v
	return s
}

func (s *DescribeClusterPolicyResponseBody) SetDomainList(v string) *DescribeClusterPolicyResponseBody {
	s.DomainList = &v
	return s
}

func (s *DescribeClusterPolicyResponseBody) SetWatermark(v string) *DescribeClusterPolicyResponseBody {
	s.Watermark = &v
	return s
}

type DescribeClusterPolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterPolicyResponse) SetHeaders(v map[string]*string) *DescribeClusterPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterPolicyResponse) SetBody(v *DescribeClusterPolicyResponseBody) *DescribeClusterPolicyResponse {
	s.Body = v
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

type DescribeClustersResponseBody struct {
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Clusters   []*DescribeClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
}

func (s DescribeClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBody) SetTotalCount(v int64) *DescribeClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeClustersResponseBody) SetPageSize(v int64) *DescribeClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersResponseBody) SetRequestId(v string) *DescribeClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetPageNumber(v int64) *DescribeClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersResponseBody) SetClusters(v []*DescribeClustersResponseBodyClusters) *DescribeClustersResponseBody {
	s.Clusters = v
	return s
}

type DescribeClustersResponseBodyClusters struct {
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	NatId         *string `json:"NatId,omitempty" xml:"NatId,omitempty"`
	InstanceCount *int64  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	NatEip        *string `json:"NatEip,omitempty" xml:"NatEip,omitempty"`
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClusters) SetVpcId(v string) *DescribeClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetStatus(v string) *DescribeClustersResponseBodyClusters {
	s.Status = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetNatId(v string) *DescribeClustersResponseBodyClusters {
	s.NatId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetInstanceCount(v int64) *DescribeClustersResponseBodyClusters {
	s.InstanceCount = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetCreateTime(v string) *DescribeClustersResponseBodyClusters {
	s.CreateTime = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetNatEip(v string) *DescribeClustersResponseBodyClusters {
	s.NatEip = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetSecurityGroup(v string) *DescribeClustersResponseBodyClusters {
	s.SecurityGroup = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetName(v string) *DescribeClustersResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetDomainName(v string) *DescribeClustersResponseBodyClusters {
	s.DomainName = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetClusterId(v string) *DescribeClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

type DescribeClustersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) SetHeaders(v map[string]*string) *DescribeClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersResponse) SetBody(v *DescribeClustersResponseBody) *DescribeClustersResponse {
	s.Body = v
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

type DescribeImagesResponseBody struct {
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int64                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Images     []*DescribeImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBody) SetTotalCount(v int64) *DescribeImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImagesResponseBody) SetPageSize(v int64) *DescribeImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImagesResponseBody) SetRequestId(v string) *DescribeImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponseBody) SetPageNumber(v int64) *DescribeImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesResponseBody) SetImages(v []*DescribeImagesResponseBodyImages) *DescribeImagesResponseBody {
	s.Images = v
	return s
}

type DescribeImagesResponseBodyImages struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ImageType   *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Progress    *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImages) SetStatus(v string) *DescribeImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageType(v string) *DescribeImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetProgress(v string) *DescribeImagesResponseBodyImages {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetSize(v int64) *DescribeImagesResponseBodyImages {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetCreateTime(v string) *DescribeImagesResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetName(v string) *DescribeImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageId(v string) *DescribeImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetProductCode(v string) *DescribeImagesResponseBodyImages {
	s.ProductCode = &v
	return s
}

type DescribeImagesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) SetHeaders(v map[string]*string) *DescribeImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagesResponse) SetBody(v *DescribeImagesResponseBody) *DescribeImagesResponse {
	s.Body = v
	return s
}

type DescribeInstancePolicyRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type DescribeInstancePolicyResponseBody struct {
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskFinished   *bool   `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty"`
	VisualLossless *string `json:"VisualLossless,omitempty" xml:"VisualLossless,omitempty"`
	OptimizeFor3d  *string `json:"OptimizeFor3d,omitempty" xml:"OptimizeFor3d,omitempty"`
}

func (s DescribeInstancePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancePolicyResponseBody) SetTaskId(v string) *DescribeInstancePolicyResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeInstancePolicyResponseBody) SetRequestId(v string) *DescribeInstancePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancePolicyResponseBody) SetTaskFinished(v bool) *DescribeInstancePolicyResponseBody {
	s.TaskFinished = &v
	return s
}

func (s *DescribeInstancePolicyResponseBody) SetVisualLossless(v string) *DescribeInstancePolicyResponseBody {
	s.VisualLossless = &v
	return s
}

func (s *DescribeInstancePolicyResponseBody) SetOptimizeFor3d(v string) *DescribeInstancePolicyResponseBody {
	s.OptimizeFor3d = &v
	return s
}

type DescribeInstancePolicyResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancePolicyResponse) SetHeaders(v map[string]*string) *DescribeInstancePolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancePolicyResponse) SetBody(v *DescribeInstancePolicyResponseBody) *DescribeInstancePolicyResponse {
	s.Body = v
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

type DescribeInstancesResponseBody struct {
	Instances  []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int64                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int64) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int64) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	Status             *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkMode           *string                                          `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	ExpireTime         *string                                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime         *string                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	StoppedMode        *string                                          `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
	UserUid            *int64                                           `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
	InstanceId         *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType       *string                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	DomainName         *string                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	InstanceChargeType *string                                          `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	AppList            []*DescribeInstancesResponseBodyInstancesAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	MaxBandwidthIn     *int64                                           `json:"MaxBandwidthIn,omitempty" xml:"MaxBandwidthIn,omitempty"`
	IsBoundUser        *bool                                            `json:"IsBoundUser,omitempty" xml:"IsBoundUser,omitempty"`
	MaxBandwidthOut    *int64                                           `json:"MaxBandwidthOut,omitempty" xml:"MaxBandwidthOut,omitempty"`
	Name               *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	UserName           *string                                          `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ImageId            *string                                          `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ClusterId          *string                                          `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetWorkMode(v string) *DescribeInstancesResponseBodyInstances {
	s.WorkMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v string) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetCreateTime(v string) *DescribeInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStoppedMode(v string) *DescribeInstancesResponseBodyInstances {
	s.StoppedMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUserUid(v int64) *DescribeInstancesResponseBodyInstances {
	s.UserUid = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceType(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDomainName(v string) *DescribeInstancesResponseBodyInstances {
	s.DomainName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceChargeType(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetAppList(v []*DescribeInstancesResponseBodyInstancesAppList) *DescribeInstancesResponseBodyInstances {
	s.AppList = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetMaxBandwidthIn(v int64) *DescribeInstancesResponseBodyInstances {
	s.MaxBandwidthIn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIsBoundUser(v bool) *DescribeInstancesResponseBodyInstances {
	s.IsBoundUser = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetMaxBandwidthOut(v int64) *DescribeInstancesResponseBodyInstances {
	s.MaxBandwidthOut = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetName(v string) *DescribeInstancesResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUserName(v string) *DescribeInstancesResponseBodyInstances {
	s.UserName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetImageId(v string) *DescribeInstancesResponseBodyInstances {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterId(v string) *DescribeInstancesResponseBodyInstances {
	s.ClusterId = &v
	return s
}

type DescribeInstancesResponseBodyInstancesAppList struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPath *string `json:"AppPath,omitempty" xml:"AppPath,omitempty"`
	AppArgs *string `json:"AppArgs,omitempty" xml:"AppArgs,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesAppList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesAppList) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesAppList) SetAppName(v string) *DescribeInstancesResponseBodyInstancesAppList {
	s.AppName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesAppList) SetAppPath(v string) *DescribeInstancesResponseBodyInstancesAppList {
	s.AppPath = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesAppList) SetAppArgs(v string) *DescribeInstancesResponseBodyInstancesAppList {
	s.AppArgs = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type GetConnectTicketRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type GetConnectTicketResponseBody struct {
	Ticket       *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskFinished *bool   `json:"TaskFinished,omitempty" xml:"TaskFinished,omitempty"`
}

func (s GetConnectTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectTicketResponseBody) SetTicket(v string) *GetConnectTicketResponseBody {
	s.Ticket = &v
	return s
}

func (s *GetConnectTicketResponseBody) SetTaskId(v string) *GetConnectTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectTicketResponseBody) SetRequestId(v string) *GetConnectTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectTicketResponseBody) SetTaskFinished(v bool) *GetConnectTicketResponseBody {
	s.TaskFinished = &v
	return s
}

type GetConnectTicketResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConnectTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConnectTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectTicketResponse) SetHeaders(v map[string]*string) *GetConnectTicketResponse {
	s.Headers = v
	return s
}

func (s *GetConnectTicketResponse) SetBody(v *GetConnectTicketResponseBody) *GetConnectTicketResponse {
	s.Body = v
	return s
}

type IsUserAdminResponseBody struct {
	IsAdmin   *bool   `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IsAllow   *bool   `json:"IsAllow,omitempty" xml:"IsAllow,omitempty"`
}

func (s IsUserAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsUserAdminResponseBody) GoString() string {
	return s.String()
}

func (s *IsUserAdminResponseBody) SetIsAdmin(v bool) *IsUserAdminResponseBody {
	s.IsAdmin = &v
	return s
}

func (s *IsUserAdminResponseBody) SetRequestId(v string) *IsUserAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsUserAdminResponseBody) SetIsAllow(v bool) *IsUserAdminResponseBody {
	s.IsAllow = &v
	return s
}

type IsUserAdminResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IsUserAdminResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IsUserAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s IsUserAdminResponse) GoString() string {
	return s.String()
}

func (s *IsUserAdminResponse) SetHeaders(v map[string]*string) *IsUserAdminResponse {
	s.Headers = v
	return s
}

func (s *IsUserAdminResponse) SetBody(v *IsUserAdminResponseBody) *IsUserAdminResponse {
	s.Body = v
	return s
}

type RestartInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type RestartInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) SetHeaders(v map[string]*string) *RestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

type SetClusterADDomainRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DomainDnsIp    *string `json:"DomainDnsIp,omitempty" xml:"DomainDnsIp,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainPassword *string `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
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

type SetClusterADDomainResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetClusterADDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetClusterADDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetClusterADDomainResponseBody) SetTaskId(v string) *SetClusterADDomainResponseBody {
	s.TaskId = &v
	return s
}

func (s *SetClusterADDomainResponseBody) SetRequestId(v string) *SetClusterADDomainResponseBody {
	s.RequestId = &v
	return s
}

type SetClusterADDomainResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetClusterADDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetClusterADDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s SetClusterADDomainResponse) GoString() string {
	return s.String()
}

func (s *SetClusterADDomainResponse) SetHeaders(v map[string]*string) *SetClusterADDomainResponse {
	s.Headers = v
	return s
}

func (s *SetClusterADDomainResponse) SetBody(v *SetClusterADDomainResponseBody) *SetClusterADDomainResponse {
	s.Body = v
	return s
}

type SetClusterDnatRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NatId     *string `json:"NatId,omitempty" xml:"NatId,omitempty"`
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

type SetClusterDnatResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetClusterDnatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetClusterDnatResponseBody) GoString() string {
	return s.String()
}

func (s *SetClusterDnatResponseBody) SetRequestId(v string) *SetClusterDnatResponseBody {
	s.RequestId = &v
	return s
}

type SetClusterDnatResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetClusterDnatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetClusterDnatResponse) String() string {
	return tea.Prettify(s)
}

func (s SetClusterDnatResponse) GoString() string {
	return s.String()
}

func (s *SetClusterDnatResponse) SetHeaders(v map[string]*string) *SetClusterDnatResponse {
	s.Headers = v
	return s
}

func (s *SetClusterDnatResponse) SetBody(v *SetClusterDnatResponseBody) *SetClusterDnatResponse {
	s.Body = v
	return s
}

type SetClusterNameRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type SetClusterNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetClusterNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetClusterNameResponseBody) SetRequestId(v string) *SetClusterNameResponseBody {
	s.RequestId = &v
	return s
}

type SetClusterNameResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetClusterNameResponse) GoString() string {
	return s.String()
}

func (s *SetClusterNameResponse) SetHeaders(v map[string]*string) *SetClusterNameResponse {
	s.Headers = v
	return s
}

func (s *SetClusterNameResponse) SetBody(v *SetClusterNameResponseBody) *SetClusterNameResponse {
	s.Body = v
	return s
}

type SetClusterPolicyRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
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

type SetClusterPolicyResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetClusterPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetClusterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetClusterPolicyResponseBody) SetTaskId(v string) *SetClusterPolicyResponseBody {
	s.TaskId = &v
	return s
}

func (s *SetClusterPolicyResponseBody) SetRequestId(v string) *SetClusterPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetClusterPolicyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetClusterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetClusterPolicyResponse) SetHeaders(v map[string]*string) *SetClusterPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetClusterPolicyResponse) SetBody(v *SetClusterPolicyResponseBody) *SetClusterPolicyResponse {
	s.Body = v
	return s
}

type SetImageNameRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type SetImageNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetImageNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetImageNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetImageNameResponseBody) SetRequestId(v string) *SetImageNameResponseBody {
	s.RequestId = &v
	return s
}

type SetImageNameResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetImageNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetImageNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetImageNameResponse) GoString() string {
	return s.String()
}

func (s *SetImageNameResponse) SetHeaders(v map[string]*string) *SetImageNameResponse {
	s.Headers = v
	return s
}

func (s *SetImageNameResponse) SetBody(v *SetImageNameResponseBody) *SetImageNameResponse {
	s.Body = v
	return s
}

type SetInstanceNameRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type SetInstanceNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceNameResponseBody) SetRequestId(v string) *SetInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

type SetInstanceNameResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceNameResponse) SetHeaders(v map[string]*string) *SetInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceNameResponse) SetBody(v *SetInstanceNameResponseBody) *SetInstanceNameResponse {
	s.Body = v
	return s
}

type SetInstancePolicyRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type SetInstancePolicyResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstancePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInstancePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstancePolicyResponseBody) SetTaskId(v string) *SetInstancePolicyResponseBody {
	s.TaskId = &v
	return s
}

func (s *SetInstancePolicyResponseBody) SetRequestId(v string) *SetInstancePolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetInstancePolicyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetInstancePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetInstancePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *SetInstancePolicyResponse) SetHeaders(v map[string]*string) *SetInstancePolicyResponse {
	s.Headers = v
	return s
}

func (s *SetInstancePolicyResponse) SetBody(v *SetInstancePolicyResponseBody) *SetInstancePolicyResponse {
	s.Body = v
	return s
}

type SetInstanceUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserUid    *int64  `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

type SetInstanceUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstanceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceUserResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceUserResponseBody) SetRequestId(v string) *SetInstanceUserResponseBody {
	s.RequestId = &v
	return s
}

type SetInstanceUserResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetInstanceUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetInstanceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceUserResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceUserResponse) SetHeaders(v map[string]*string) *SetInstanceUserResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceUserResponse) SetBody(v *SetInstanceUserResponseBody) *SetInstanceUserResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type StartInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type StopInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
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

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCluster"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateImage"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateServiceLinkedRoleWithOptions(runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServiceLinkedRole"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole() (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCluster"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImage"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInstance"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAvailableResourceWithOptions(runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableResource"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResource() (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterADDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterADDomain"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeClusterConnectionsWithOptions(request *DescribeClusterConnectionsRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterConnectionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterConnections"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeClusterPolicyWithOptions(request *DescribeClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterPolicy"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeClustersWithOptions(request *DescribeClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusters"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImages"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeInstancePolicyWithOptions(request *DescribeInstancePolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstancePolicy"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetConnectTicketWithOptions(request *GetConnectTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetConnectTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConnectTicket"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) IsUserAdminWithOptions(runtime *util.RuntimeOptions) (_result *IsUserAdminResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &IsUserAdminResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IsUserAdmin"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsUserAdmin() (_result *IsUserAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IsUserAdminResponse{}
	_body, _err := client.IsUserAdminWithOptions(runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartInstance"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SetClusterADDomainWithOptions(request *SetClusterADDomainRequest, runtime *util.RuntimeOptions) (_result *SetClusterADDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetClusterADDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetClusterADDomain"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SetClusterDnatWithOptions(request *SetClusterDnatRequest, runtime *util.RuntimeOptions) (_result *SetClusterDnatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetClusterDnatResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetClusterDnat"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SetClusterNameWithOptions(request *SetClusterNameRequest, runtime *util.RuntimeOptions) (_result *SetClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetClusterNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetClusterName"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SetClusterPolicyWithOptions(request *SetClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *SetClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetClusterPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetClusterPolicy"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SetImageNameWithOptions(request *SetImageNameRequest, runtime *util.RuntimeOptions) (_result *SetImageNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetImageNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetImageName"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SetInstanceNameWithOptions(request *SetInstanceNameRequest, runtime *util.RuntimeOptions) (_result *SetInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetInstanceNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetInstanceName"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SetInstancePolicyWithOptions(request *SetInstancePolicyRequest, runtime *util.RuntimeOptions) (_result *SetInstancePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetInstancePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetInstancePolicy"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SetInstanceUserWithOptions(request *SetInstanceUserRequest, runtime *util.RuntimeOptions) (_result *SetInstanceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetInstanceUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetInstanceUser"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartInstance"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopInstance"), tea.String("2019-06-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
