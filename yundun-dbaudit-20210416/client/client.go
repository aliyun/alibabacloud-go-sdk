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

type ClearInstanceStorageRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StorageSpace    *string `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
}

func (s ClearInstanceStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearInstanceStorageRequest) GoString() string {
	return s.String()
}

func (s *ClearInstanceStorageRequest) SetInstanceId(v string) *ClearInstanceStorageRequest {
	s.InstanceId = &v
	return s
}

func (s *ClearInstanceStorageRequest) SetRegionId(v string) *ClearInstanceStorageRequest {
	s.RegionId = &v
	return s
}

func (s *ClearInstanceStorageRequest) SetStorageSpace(v string) *ClearInstanceStorageRequest {
	s.StorageSpace = &v
	return s
}

func (s *ClearInstanceStorageRequest) SetStorageCategory(v string) *ClearInstanceStorageRequest {
	s.StorageCategory = &v
	return s
}

type ClearInstanceStorageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearInstanceStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearInstanceStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ClearInstanceStorageResponseBody) SetRequestId(v string) *ClearInstanceStorageResponseBody {
	s.RequestId = &v
	return s
}

type ClearInstanceStorageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearInstanceStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClearInstanceStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearInstanceStorageResponse) GoString() string {
	return s.String()
}

func (s *ClearInstanceStorageResponse) SetHeaders(v map[string]*string) *ClearInstanceStorageResponse {
	s.Headers = v
	return s
}

func (s *ClearInstanceStorageResponse) SetBody(v *ClearInstanceStorageResponseBody) *ClearInstanceStorageResponse {
	s.Body = v
	return s
}

type CreateInstanceConnectionRequest struct {
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId            *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Bandwidth        *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	VswitchIds       []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s CreateInstanceConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceConnectionRequest) SetInstanceId(v string) *CreateInstanceConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceConnectionRequest) SetRegionId(v string) *CreateInstanceConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceConnectionRequest) SetVpcId(v string) *CreateInstanceConnectionRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceConnectionRequest) SetBandwidth(v int32) *CreateInstanceConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateInstanceConnectionRequest) SetVswitchIds(v []*string) *CreateInstanceConnectionRequest {
	s.VswitchIds = v
	return s
}

func (s *CreateInstanceConnectionRequest) SetSecurityGroupIds(v []*string) *CreateInstanceConnectionRequest {
	s.SecurityGroupIds = v
	return s
}

type CreateInstanceConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceConnectionResponseBody) SetRequestId(v string) *CreateInstanceConnectionResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceConnectionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceConnectionResponse) SetHeaders(v map[string]*string) *CreateInstanceConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceConnectionResponse) SetBody(v *CreateInstanceConnectionResponseBody) *CreateInstanceConnectionResponse {
	s.Body = v
	return s
}

type DeleteInstanceConnectionRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteInstanceConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceConnectionRequest) SetInstanceId(v string) *DeleteInstanceConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceConnectionRequest) SetRegionId(v string) *DeleteInstanceConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteInstanceConnectionRequest) SetVpcId(v string) *DeleteInstanceConnectionRequest {
	s.VpcId = &v
	return s
}

type DeleteInstanceConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceConnectionResponseBody) SetRequestId(v string) *DeleteInstanceConnectionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceConnectionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceConnectionResponse) SetHeaders(v map[string]*string) *DeleteInstanceConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceConnectionResponse) SetBody(v *DeleteInstanceConnectionResponseBody) *DeleteInstanceConnectionResponse {
	s.Body = v
	return s
}

type DescribeInstanceAttributeRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeRequest) SetInstanceId(v string) *DescribeInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetRegionId(v string) *DescribeInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceAttributeResponseBody struct {
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceAttribute *DescribeInstanceAttributeResponseBodyInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Struct"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody {
	s.InstanceAttribute = v
	return s
}

type DescribeInstanceAttributeResponseBodyInstanceAttribute struct {
	InstancePurchaseStorage        *int64                                                                       `json:"InstancePurchaseStorage,omitempty" xml:"InstancePurchaseStorage,omitempty"`
	ExpireTime                     *int64                                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceTotalStorage           *int64                                                                       `json:"InstanceTotalStorage,omitempty" xml:"InstanceTotalStorage,omitempty"`
	ImageVersion                   *string                                                                      `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	InstanceId                     *string                                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceConnectionQuota        *int32                                                                       `json:"InstanceConnectionQuota,omitempty" xml:"InstanceConnectionQuota,omitempty"`
	RegionId                       *string                                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceDatabaseQuota          *int64                                                                       `json:"InstanceDatabaseQuota,omitempty" xml:"InstanceDatabaseQuota,omitempty"`
	StartTime                      *int64                                                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SeriesCode                     *string                                                                      `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	Description                    *string                                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceStatus                 *string                                                                      `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LicenseCode                    *string                                                                      `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	InstanceConnectionMaxBandwidth *int32                                                                       `json:"InstanceConnectionMaxBandwidth,omitempty" xml:"InstanceConnectionMaxBandwidth,omitempty"`
	InstanceConnections            []*DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections `json:"InstanceConnections,omitempty" xml:"InstanceConnections,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstancePurchaseStorage(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstancePurchaseStorage = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetExpireTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceTotalStorage(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceTotalStorage = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetImageVersion(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ImageVersion = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceConnectionQuota(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceConnectionQuota = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceDatabaseQuota(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceDatabaseQuota = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStartTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSeriesCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDescription(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetLicenseCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceConnectionMaxBandwidth(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceConnectionMaxBandwidth = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceConnections(v []*DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceConnections = v
	return s
}

type DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections struct {
	VpcId                       *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	InstanceConnectionBandwidth *int32    `json:"InstanceConnectionBandwidth,omitempty" xml:"InstanceConnectionBandwidth,omitempty"`
	InstanceConnectionDomain    *string   `json:"InstanceConnectionDomain,omitempty" xml:"InstanceConnectionDomain,omitempty"`
	SecurityGroupIds            []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	VswitchIds                  []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections) SetInstanceConnectionBandwidth(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections {
	s.InstanceConnectionBandwidth = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections) SetInstanceConnectionDomain(v string) *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections {
	s.InstanceConnectionDomain = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections) SetSecurityGroupIds(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections) SetVswitchIds(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttributeInstanceConnections {
	s.VswitchIds = v
	return s
}

type DescribeInstanceAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAttributeResponse) SetBody(v *DescribeInstanceAttributeResponseBody) *DescribeInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	PageSize        *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RegionId        *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceStatus  *string                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceId      []*string                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetCurrentPage(v int32) *DescribeInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceStatus(v string) *DescribeInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v []*string) *DescribeInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponseBody struct {
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Instances  []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	InstancePurchaseStorage        *int64  `json:"InstancePurchaseStorage,omitempty" xml:"InstancePurchaseStorage,omitempty"`
	ExpireTime                     *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceTotalStorage           *int64  `json:"InstanceTotalStorage,omitempty" xml:"InstanceTotalStorage,omitempty"`
	ImageVersion                   *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	InstanceId                     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceConnectionQuota        *int32  `json:"InstanceConnectionQuota,omitempty" xml:"InstanceConnectionQuota,omitempty"`
	RegionId                       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceDatabaseQuota          *int64  `json:"InstanceDatabaseQuota,omitempty" xml:"InstanceDatabaseQuota,omitempty"`
	StartTime                      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SeriesCode                     *string `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	Description                    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceStatus                 *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LicenseCode                    *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	InstanceConnectionMaxBandwidth *int32  `json:"InstanceConnectionMaxBandwidth,omitempty" xml:"InstanceConnectionMaxBandwidth,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetInstancePurchaseStorage(v int64) *DescribeInstancesResponseBodyInstances {
	s.InstancePurchaseStorage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceTotalStorage(v int64) *DescribeInstancesResponseBodyInstances {
	s.InstanceTotalStorage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetImageVersion(v string) *DescribeInstancesResponseBodyInstances {
	s.ImageVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceConnectionQuota(v int32) *DescribeInstancesResponseBodyInstances {
	s.InstanceConnectionQuota = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegionId(v string) *DescribeInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceDatabaseQuota(v int64) *DescribeInstancesResponseBodyInstances {
	s.InstanceDatabaseQuota = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStartTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSeriesCode(v string) *DescribeInstancesResponseBodyInstances {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDescription(v string) *DescribeInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLicenseCode(v string) *DescribeInstancesResponseBodyInstances {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceConnectionMaxBandwidth(v int32) *DescribeInstancesResponseBodyInstances {
	s.InstanceConnectionMaxBandwidth = &v
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

type DescribeInstanceStorageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageRequest) SetInstanceId(v string) *DescribeInstanceStorageRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStorageRequest) SetRegionId(v string) *DescribeInstanceStorageRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceStorageResponseBody struct {
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceStorages []*DescribeInstanceStorageResponseBodyInstanceStorages `json:"InstanceStorages,omitempty" xml:"InstanceStorages,omitempty" type:"Repeated"`
}

func (s DescribeInstanceStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageResponseBody) SetRequestId(v string) *DescribeInstanceStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStorageResponseBody) SetInstanceStorages(v []*DescribeInstanceStorageResponseBodyInstanceStorages) *DescribeInstanceStorageResponseBody {
	s.InstanceStorages = v
	return s
}

type DescribeInstanceStorageResponseBodyInstanceStorages struct {
	StorageTime     *int64  `json:"StorageTime,omitempty" xml:"StorageTime,omitempty"`
	StorageCapacity *int64  `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
	StorageSpace    *string `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	StorageUsed     *int64  `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
}

func (s DescribeInstanceStorageResponseBodyInstanceStorages) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageResponseBodyInstanceStorages) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageTime(v int64) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageTime = &v
	return s
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageCapacity(v int64) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageCapacity = &v
	return s
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageCategory(v string) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageCategory = &v
	return s
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageSpace(v string) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageSpace = &v
	return s
}

func (s *DescribeInstanceStorageResponseBodyInstanceStorages) SetStorageUsed(v int64) *DescribeInstanceStorageResponseBodyInstanceStorages {
	s.StorageUsed = &v
	return s
}

type DescribeInstanceStorageResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageResponse) SetHeaders(v map[string]*string) *DescribeInstanceStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStorageResponse) SetBody(v *DescribeInstanceStorageResponseBody) *DescribeInstanceStorageResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
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

type ListTagKeysRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

type ListTagKeysResponseBody struct {
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TagKeys     []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagCount(v int32) *ListTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyInstanceAttributeRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetDescription(v string) *ModifyInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRegionId(v string) *ModifyInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetBody(v *ModifyInstanceAttributeResponseBody) *ModifyInstanceAttributeResponse {
	s.Body = v
	return s
}

type ModifyInstanceStorageRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StorageSpace    *string `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
	StorageTime     *int32  `json:"StorageTime,omitempty" xml:"StorageTime,omitempty"`
}

func (s ModifyInstanceStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceStorageRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageRequest) SetInstanceId(v string) *ModifyInstanceStorageRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceStorageRequest) SetRegionId(v string) *ModifyInstanceStorageRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceStorageRequest) SetStorageSpace(v string) *ModifyInstanceStorageRequest {
	s.StorageSpace = &v
	return s
}

func (s *ModifyInstanceStorageRequest) SetStorageCategory(v string) *ModifyInstanceStorageRequest {
	s.StorageCategory = &v
	return s
}

func (s *ModifyInstanceStorageRequest) SetStorageTime(v int32) *ModifyInstanceStorageRequest {
	s.StorageTime = &v
	return s
}

type ModifyInstanceStorageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageResponseBody) SetRequestId(v string) *ModifyInstanceStorageResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceStorageResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceStorageResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageResponse) SetHeaders(v map[string]*string) *ModifyInstanceStorageResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceStorageResponse) SetBody(v *ModifyInstanceStorageResponseBody) *ModifyInstanceStorageResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RefundInstanceRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s RefundInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundInstanceRequest) GoString() string {
	return s.String()
}

func (s *RefundInstanceRequest) SetInstanceId(v string) *RefundInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundInstanceRequest) SetRegionId(v string) *RefundInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RefundInstanceRequest) SetServiceCode(v string) *RefundInstanceRequest {
	s.ServiceCode = &v
	return s
}

type RefundInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponseBody) SetRequestId(v string) *RefundInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RefundInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefundInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefundInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundInstanceResponse) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponse) SetHeaders(v map[string]*string) *RefundInstanceResponse {
	s.Headers = v
	return s
}

func (s *RefundInstanceResponse) SetBody(v *RefundInstanceResponseBody) *RefundInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateInstanceConnectionRequest struct {
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId            *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Bandwidth        *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	VswitchIds       []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s UpdateInstanceConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceConnectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceConnectionRequest) SetInstanceId(v string) *UpdateInstanceConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceConnectionRequest) SetRegionId(v string) *UpdateInstanceConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceConnectionRequest) SetVpcId(v string) *UpdateInstanceConnectionRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateInstanceConnectionRequest) SetBandwidth(v int32) *UpdateInstanceConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateInstanceConnectionRequest) SetVswitchIds(v []*string) *UpdateInstanceConnectionRequest {
	s.VswitchIds = v
	return s
}

func (s *UpdateInstanceConnectionRequest) SetSecurityGroupIds(v []*string) *UpdateInstanceConnectionRequest {
	s.SecurityGroupIds = v
	return s
}

type UpdateInstanceConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceConnectionResponseBody) SetRequestId(v string) *UpdateInstanceConnectionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceConnectionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceConnectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceConnectionResponse) SetHeaders(v map[string]*string) *UpdateInstanceConnectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceConnectionResponse) SetBody(v *UpdateInstanceConnectionResponseBody) *UpdateInstanceConnectionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("yundun-dbaudit"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ClearInstanceStorageWithOptions(request *ClearInstanceStorageRequest, runtime *util.RuntimeOptions) (_result *ClearInstanceStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClearInstanceStorageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClearInstanceStorage"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearInstanceStorage(request *ClearInstanceStorageRequest) (_result *ClearInstanceStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearInstanceStorageResponse{}
	_body, _err := client.ClearInstanceStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceConnectionWithOptions(request *CreateInstanceConnectionRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstanceConnection"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceConnection(request *CreateInstanceConnectionRequest) (_result *CreateInstanceConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceConnectionResponse{}
	_body, _err := client.CreateInstanceConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceConnectionWithOptions(request *DeleteInstanceConnectionRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInstanceConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInstanceConnection"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstanceConnection(request *DeleteInstanceConnectionRequest) (_result *DeleteInstanceConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceConnectionResponse{}
	_body, _err := client.DeleteInstanceConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAttributeWithOptions(request *DescribeInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAttribute"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (_result *DescribeInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DescribeInstanceAttributeWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeInstanceStorageWithOptions(request *DescribeInstanceStorageRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceStorageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceStorage"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceStorage(request *DescribeInstanceStorageRequest) (_result *DescribeInstanceStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceStorageResponse{}
	_body, _err := client.DescribeInstanceStorageWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceAttribute"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceStorageWithOptions(request *ModifyInstanceStorageRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceStorageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceStorage"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceStorage(request *ModifyInstanceStorageRequest) (_result *ModifyInstanceStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceStorageResponse{}
	_body, _err := client.ModifyInstanceStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveResourceGroup"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundInstanceWithOptions(request *RefundInstanceRequest, runtime *util.RuntimeOptions) (_result *RefundInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefundInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefundInstance"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundInstance(request *RefundInstanceRequest) (_result *RefundInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundInstanceResponse{}
	_body, _err := client.RefundInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceConnectionWithOptions(request *UpdateInstanceConnectionRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateInstanceConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateInstanceConnection"), tea.String("2021-04-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceConnection(request *UpdateInstanceConnectionRequest) (_result *UpdateInstanceConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceConnectionResponse{}
	_body, _err := client.UpdateInstanceConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
