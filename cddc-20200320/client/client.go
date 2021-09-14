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

type ModifyDBInstanceSwitchWeightRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceName       *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	SwitchWeight         *string `json:"SwitchWeight,omitempty" xml:"SwitchWeight,omitempty"`
}

func (s ModifyDBInstanceSwitchWeightRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSwitchWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSwitchWeightRequest) SetOwnerId(v int64) *ModifyDBInstanceSwitchWeightRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSwitchWeightRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSwitchWeightRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSwitchWeightRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSwitchWeightRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSwitchWeightRequest) SetRegionId(v string) *ModifyDBInstanceSwitchWeightRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceSwitchWeightRequest) SetDBInstanceName(v string) *ModifyDBInstanceSwitchWeightRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceSwitchWeightRequest) SetSwitchWeight(v string) *ModifyDBInstanceSwitchWeightRequest {
	s.SwitchWeight = &v
	return s
}

type ModifyDBInstanceSwitchWeightResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceSwitchWeightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSwitchWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSwitchWeightResponseBody) SetRequestId(v string) *ModifyDBInstanceSwitchWeightResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceSwitchWeightResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBInstanceSwitchWeightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBInstanceSwitchWeightResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSwitchWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSwitchWeightResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceSwitchWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceSwitchWeightResponse) SetBody(v *ModifyDBInstanceSwitchWeightResponseBody) *ModifyDBInstanceSwitchWeightResponse {
	s.Body = v
	return s
}

type DescribeAvailableDedicatedHostZonesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DbType               *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
}

func (s DescribeAvailableDedicatedHostZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostZonesRequest) SetOwnerId(v int64) *DescribeAvailableDedicatedHostZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostZonesRequest) SetResourceOwnerAccount(v string) *DescribeAvailableDedicatedHostZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableDedicatedHostZonesRequest) SetResourceOwnerId(v int64) *DescribeAvailableDedicatedHostZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostZonesRequest) SetRegionId(v string) *DescribeAvailableDedicatedHostZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostZonesRequest) SetDbType(v string) *DescribeAvailableDedicatedHostZonesRequest {
	s.DbType = &v
	return s
}

type DescribeAvailableDedicatedHostZonesResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *DescribeAvailableDedicatedHostZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeAvailableDedicatedHostZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostZonesResponseBody) SetRequestId(v string) *DescribeAvailableDedicatedHostZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostZonesResponseBody) SetZones(v *DescribeAvailableDedicatedHostZonesResponseBodyZones) *DescribeAvailableDedicatedHostZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeAvailableDedicatedHostZonesResponseBodyZones struct {
	DedicatedHostZones []*DescribeAvailableDedicatedHostZonesResponseBodyZonesDedicatedHostZones `json:"DedicatedHostZones,omitempty" xml:"DedicatedHostZones,omitempty" type:"Repeated"`
}

func (s DescribeAvailableDedicatedHostZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostZonesResponseBodyZones) SetDedicatedHostZones(v []*DescribeAvailableDedicatedHostZonesResponseBodyZonesDedicatedHostZones) *DescribeAvailableDedicatedHostZonesResponseBodyZones {
	s.DedicatedHostZones = v
	return s
}

type DescribeAvailableDedicatedHostZonesResponseBodyZonesDedicatedHostZones struct {
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeAvailableDedicatedHostZonesResponseBodyZonesDedicatedHostZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostZonesResponseBodyZonesDedicatedHostZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostZonesResponseBodyZonesDedicatedHostZones) SetZoneId(v string) *DescribeAvailableDedicatedHostZonesResponseBodyZonesDedicatedHostZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostZonesResponseBodyZonesDedicatedHostZones) SetDescription(v string) *DescribeAvailableDedicatedHostZonesResponseBodyZonesDedicatedHostZones {
	s.Description = &v
	return s
}

type DescribeAvailableDedicatedHostZonesResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableDedicatedHostZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableDedicatedHostZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostZonesResponse) SetHeaders(v map[string]*string) *DescribeAvailableDedicatedHostZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableDedicatedHostZonesResponse) SetBody(v *DescribeAvailableDedicatedHostZonesResponseBody) *DescribeAvailableDedicatedHostZonesResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostGroupsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ImageCategory        *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDedicatedHostGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsRequest) SetOwnerId(v int64) *DescribeDedicatedHostGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetRegionId(v string) *DescribeDedicatedHostGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostGroupsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetImageCategory(v string) *DescribeDedicatedHostGroupsRequest {
	s.ImageCategory = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetEngine(v string) *DescribeDedicatedHostGroupsRequest {
	s.Engine = &v
	return s
}

type DescribeDedicatedHostGroupsResponseBody struct {
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DedicatedHostGroups *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups `json:"DedicatedHostGroups,omitempty" xml:"DedicatedHostGroups,omitempty" type:"Struct"`
}

func (s DescribeDedicatedHostGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBody) SetRequestId(v string) *DescribeDedicatedHostGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBody) SetDedicatedHostGroups(v *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) *DescribeDedicatedHostGroupsResponseBody {
	s.DedicatedHostGroups = v
	return s
}

type DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups struct {
	DedicatedHostGroups []*DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups `json:"DedicatedHostGroups,omitempty" xml:"DedicatedHostGroups,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) SetDedicatedHostGroups(v []*DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups {
	s.DedicatedHostGroups = v
	return s
}

type DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups struct {
	DiskAllocateRation                *float32                                                                                 `json:"DiskAllocateRation,omitempty" xml:"DiskAllocateRation,omitempty"`
	DeployType                        *string                                                                                  `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	CreateTime                        *string                                                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DedicatedHostCountGroupByHostType map[string]interface{}                                                                   `json:"DedicatedHostCountGroupByHostType,omitempty" xml:"DedicatedHostCountGroupByHostType,omitempty"`
	Text                              *string                                                                                  `json:"Text,omitempty" xml:"Text,omitempty"`
	DedicatedHostGroupId              *string                                                                                  `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DiskUtility                       *float32                                                                                 `json:"DiskUtility,omitempty" xml:"DiskUtility,omitempty"`
	MemUsedAmount                     *float32                                                                                 `json:"MemUsedAmount,omitempty" xml:"MemUsedAmount,omitempty"`
	MemAllocatedAmount                *float32                                                                                 `json:"MemAllocatedAmount,omitempty" xml:"MemAllocatedAmount,omitempty"`
	CpuAllocationRatio                *int32                                                                                   `json:"CpuAllocationRatio,omitempty" xml:"CpuAllocationRatio,omitempty"`
	MemAllocationRatio                *int32                                                                                   `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	MemAllocateRation                 *float32                                                                                 `json:"MemAllocateRation,omitempty" xml:"MemAllocateRation,omitempty"`
	MemUtility                        *float32                                                                                 `json:"MemUtility,omitempty" xml:"MemUtility,omitempty"`
	CpuAllocatedAmount                *float32                                                                                 `json:"CpuAllocatedAmount,omitempty" xml:"CpuAllocatedAmount,omitempty"`
	DedicatedHostGroupDesc            *string                                                                                  `json:"DedicatedHostGroupDesc,omitempty" xml:"DedicatedHostGroupDesc,omitempty"`
	CpuAllocateRation                 *float32                                                                                 `json:"CpuAllocateRation,omitempty" xml:"CpuAllocateRation,omitempty"`
	InstanceNumber                    *int32                                                                                   `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	OpenPermission                    *string                                                                                  `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	VPCId                             *string                                                                                  `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	DiskAllocatedAmount               *float32                                                                                 `json:"DiskAllocatedAmount,omitempty" xml:"DiskAllocatedAmount,omitempty"`
	HostNumber                        *int32                                                                                   `json:"HostNumber,omitempty" xml:"HostNumber,omitempty"`
	DiskUsedAmount                    *float32                                                                                 `json:"DiskUsedAmount,omitempty" xml:"DiskUsedAmount,omitempty"`
	AllocationPolicy                  *string                                                                                  `json:"AllocationPolicy,omitempty" xml:"AllocationPolicy,omitempty"`
	Engine                            *string                                                                                  `json:"Engine,omitempty" xml:"Engine,omitempty"`
	DiskAllocationRatio               *int32                                                                                   `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	BastionInstanceId                 *string                                                                                  `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	HostReplacePolicy                 *string                                                                                  `json:"HostReplacePolicy,omitempty" xml:"HostReplacePolicy,omitempty"`
	ZoneIDList                        *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList `json:"ZoneIDList,omitempty" xml:"ZoneIDList,omitempty" type:"Struct"`
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskAllocateRation(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskAllocateRation = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDeployType(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DeployType = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCreateTime(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDedicatedHostCountGroupByHostType(v map[string]interface{}) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DedicatedHostCountGroupByHostType = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetText(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.Text = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskUtility(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskUtility = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemUsedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemUsedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemAllocatedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemAllocatedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCpuAllocationRatio(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CpuAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemAllocationRatio(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemAllocateRation(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemAllocateRation = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemUtility(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemUtility = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCpuAllocatedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CpuAllocatedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDedicatedHostGroupDesc(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DedicatedHostGroupDesc = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCpuAllocateRation(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CpuAllocateRation = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetInstanceNumber(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetOpenPermission(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.OpenPermission = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetVPCId(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.VPCId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskAllocatedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskAllocatedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetHostNumber(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.HostNumber = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskUsedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskUsedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetAllocationPolicy(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.AllocationPolicy = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetEngine(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskAllocationRatio(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetBastionInstanceId(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.BastionInstanceId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetHostReplacePolicy(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.HostReplacePolicy = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetZoneIDList(v *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.ZoneIDList = v
	return s
}

type DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList struct {
	ZoneIDList []*string `json:"ZoneIDList,omitempty" xml:"ZoneIDList,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList) SetZoneIDList(v []*string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList {
	s.ZoneIDList = v
	return s
}

type DescribeDedicatedHostGroupsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedHostGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedHostGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponse) SetBody(v *DescribeDedicatedHostGroupsResponseBody) *DescribeDedicatedHostGroupsResponse {
	s.Body = v
	return s
}

type DescribeMyBaseHostOverViewRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeMyBaseHostOverViewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseHostOverViewRequest) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseHostOverViewRequest) SetOwnerId(v int64) *DescribeMyBaseHostOverViewRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMyBaseHostOverViewRequest) SetResourceOwnerAccount(v string) *DescribeMyBaseHostOverViewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMyBaseHostOverViewRequest) SetResourceOwnerId(v int64) *DescribeMyBaseHostOverViewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMyBaseHostOverViewRequest) SetRegionId(v string) *DescribeMyBaseHostOverViewRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMyBaseHostOverViewRequest) SetRegion(v string) *DescribeMyBaseHostOverViewRequest {
	s.Region = &v
	return s
}

type DescribeMyBaseHostOverViewResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeMyBaseHostOverViewResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeMyBaseHostOverViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseHostOverViewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseHostOverViewResponseBody) SetRequestId(v string) *DescribeMyBaseHostOverViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMyBaseHostOverViewResponseBody) SetRegions(v *DescribeMyBaseHostOverViewResponseBodyRegions) *DescribeMyBaseHostOverViewResponseBody {
	s.Regions = v
	return s
}

type DescribeMyBaseHostOverViewResponseBodyRegions struct {
	RegionModel []*DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel `json:"RegionModel,omitempty" xml:"RegionModel,omitempty" type:"Repeated"`
}

func (s DescribeMyBaseHostOverViewResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseHostOverViewResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegions) SetRegionModel(v []*DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel) *DescribeMyBaseHostOverViewResponseBodyRegions {
	s.RegionModel = v
	return s
}

type DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel struct {
	EngineCount    *string                                                             `json:"EngineCount,omitempty" xml:"EngineCount,omitempty"`
	TotalCount     *int32                                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	HostGroupCount *int32                                                              `json:"HostGroupCount,omitempty" xml:"HostGroupCount,omitempty"`
	Region         *string                                                             `json:"Region,omitempty" xml:"Region,omitempty"`
	TypeModels     *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModels `json:"TypeModels,omitempty" xml:"TypeModels,omitempty" type:"Struct"`
}

func (s DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel) SetEngineCount(v string) *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel {
	s.EngineCount = &v
	return s
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel) SetTotalCount(v int32) *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel {
	s.TotalCount = &v
	return s
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel) SetHostGroupCount(v int32) *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel {
	s.HostGroupCount = &v
	return s
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel) SetRegion(v string) *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel {
	s.Region = &v
	return s
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel) SetTypeModels(v *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModels) *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModel {
	s.TypeModels = v
	return s
}

type DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModels struct {
	TypeModel []*DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel `json:"TypeModel,omitempty" xml:"TypeModel,omitempty" type:"Repeated"`
}

func (s DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModels) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModels) SetTypeModel(v []*DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModels {
	s.TypeModel = v
	return s
}

type DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel struct {
	HostEngineCount *string `json:"HostEngineCount,omitempty" xml:"HostEngineCount,omitempty"`
	HostDateType    *string `json:"HostDateType,omitempty" xml:"HostDateType,omitempty"`
	Count           *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) SetHostEngineCount(v string) *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel {
	s.HostEngineCount = &v
	return s
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) SetHostDateType(v string) *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel {
	s.HostDateType = &v
	return s
}

func (s *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) SetCount(v int32) *DescribeMyBaseHostOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel {
	s.Count = &v
	return s
}

type DescribeMyBaseHostOverViewResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMyBaseHostOverViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMyBaseHostOverViewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseHostOverViewResponse) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseHostOverViewResponse) SetHeaders(v map[string]*string) *DescribeMyBaseHostOverViewResponse {
	s.Headers = v
	return s
}

func (s *DescribeMyBaseHostOverViewResponse) SetBody(v *DescribeMyBaseHostOverViewResponseBody) *DescribeMyBaseHostOverViewResponse {
	s.Body = v
	return s
}

type DescribeBriefDedicatedHostsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	PageNumbers          *int32  `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBriefDedicatedHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBriefDedicatedHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBriefDedicatedHostsRequest) SetOwnerId(v int64) *DescribeBriefDedicatedHostsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBriefDedicatedHostsRequest) SetResourceOwnerAccount(v string) *DescribeBriefDedicatedHostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBriefDedicatedHostsRequest) SetResourceOwnerId(v int64) *DescribeBriefDedicatedHostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBriefDedicatedHostsRequest) SetRegionId(v string) *DescribeBriefDedicatedHostsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBriefDedicatedHostsRequest) SetDedicatedHostGroupId(v string) *DescribeBriefDedicatedHostsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeBriefDedicatedHostsRequest) SetZoneId(v string) *DescribeBriefDedicatedHostsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeBriefDedicatedHostsRequest) SetPageNumbers(v int32) *DescribeBriefDedicatedHostsRequest {
	s.PageNumbers = &v
	return s
}

func (s *DescribeBriefDedicatedHostsRequest) SetPageSize(v int32) *DescribeBriefDedicatedHostsRequest {
	s.PageSize = &v
	return s
}

type DescribeBriefDedicatedHostsResponseBody struct {
	TotalRecords         *int32                                                   `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
	PageSize             *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumbers          *int32                                                   `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	DedicatedHostGroupId *string                                                  `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHosts       []*DescribeBriefDedicatedHostsResponseBodyDedicatedHosts `json:"DedicatedHosts,omitempty" xml:"DedicatedHosts,omitempty" type:"Repeated"`
}

func (s DescribeBriefDedicatedHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBriefDedicatedHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBriefDedicatedHostsResponseBody) SetTotalRecords(v int32) *DescribeBriefDedicatedHostsResponseBody {
	s.TotalRecords = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBody) SetPageSize(v int32) *DescribeBriefDedicatedHostsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBody) SetRequestId(v string) *DescribeBriefDedicatedHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBody) SetPageNumbers(v int32) *DescribeBriefDedicatedHostsResponseBody {
	s.PageNumbers = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBody) SetDedicatedHostGroupId(v string) *DescribeBriefDedicatedHostsResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBody) SetDedicatedHosts(v []*DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) *DescribeBriefDedicatedHostsResponseBody {
	s.DedicatedHosts = v
	return s
}

type DescribeBriefDedicatedHostsResponseBodyDedicatedHosts struct {
	AllocationStatus *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	HostStatus       *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	HostCPU          *int32  `json:"HostCPU,omitempty" xml:"HostCPU,omitempty"`
	DedicatedHostId  *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Engine           *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	HostMem          *int32  `json:"HostMem,omitempty" xml:"HostMem,omitempty"`
	CreatedTime      *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
}

func (s DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) GoString() string {
	return s.String()
}

func (s *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) SetAllocationStatus(v string) *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) SetRegion(v string) *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts {
	s.Region = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) SetHostStatus(v string) *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts {
	s.HostStatus = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) SetZoneId(v string) *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts {
	s.ZoneId = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) SetHostCPU(v int32) *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts {
	s.HostCPU = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) SetDedicatedHostId(v string) *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) SetEngine(v string) *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts {
	s.Engine = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) SetHostMem(v int32) *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts {
	s.HostMem = &v
	return s
}

func (s *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts) SetCreatedTime(v string) *DescribeBriefDedicatedHostsResponseBodyDedicatedHosts {
	s.CreatedTime = &v
	return s
}

type DescribeBriefDedicatedHostsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBriefDedicatedHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBriefDedicatedHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBriefDedicatedHostsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBriefDedicatedHostsResponse) SetHeaders(v map[string]*string) *DescribeBriefDedicatedHostsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBriefDedicatedHostsResponse) SetBody(v *DescribeBriefDedicatedHostsResponseBody) *DescribeBriefDedicatedHostsResponse {
	s.Body = v
	return s
}

type DescribeDedicatedResourceAdvisorRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDedicatedResourceAdvisorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedResourceAdvisorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedResourceAdvisorRequest) SetOwnerId(v int64) *DescribeDedicatedResourceAdvisorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedResourceAdvisorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorRequest) SetResourceOwnerId(v int64) *DescribeDedicatedResourceAdvisorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedResourceAdvisorRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorRequest) SetRegionId(v string) *DescribeDedicatedResourceAdvisorRequest {
	s.RegionId = &v
	return s
}

type DescribeDedicatedResourceAdvisorResponseBody struct {
	DedicatedHostGroupId *string                                                       `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceAdvisors     *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisors `json:"ResourceAdvisors,omitempty" xml:"ResourceAdvisors,omitempty" type:"Struct"`
}

func (s DescribeDedicatedResourceAdvisorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedResourceAdvisorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedResourceAdvisorResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedResourceAdvisorResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponseBody) SetRequestId(v string) *DescribeDedicatedResourceAdvisorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponseBody) SetResourceAdvisors(v *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisors) *DescribeDedicatedResourceAdvisorResponseBody {
	s.ResourceAdvisors = v
	return s
}

type DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisors struct {
	ResourceAdvisors []*DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors `json:"ResourceAdvisors,omitempty" xml:"ResourceAdvisors,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisors) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisors) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisors) SetResourceAdvisors(v []*DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisors {
	s.ResourceAdvisors = v
	return s
}

type DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors struct {
	CpuUsageAfterAction         *float32 `json:"CpuUsageAfterAction,omitempty" xml:"CpuUsageAfterAction,omitempty"`
	Action                      *string  `json:"Action,omitempty" xml:"Action,omitempty"`
	DestinationDedicatedHostId  *string  `json:"DestinationDedicatedHostId,omitempty" xml:"DestinationDedicatedHostId,omitempty"`
	SourceDedicatedInstanceId   *string  `json:"SourceDedicatedInstanceId,omitempty" xml:"SourceDedicatedInstanceId,omitempty"`
	ActionDescription           *string  `json:"ActionDescription,omitempty" xml:"ActionDescription,omitempty"`
	SourceDedicatedInstanceRole *string  `json:"SourceDedicatedInstanceRole,omitempty" xml:"SourceDedicatedInstanceRole,omitempty"`
	SourceDedicatedHostId       *string  `json:"SourceDedicatedHostId,omitempty" xml:"SourceDedicatedHostId,omitempty"`
	CpuUsageBeforeAction        *float32 `json:"CpuUsageBeforeAction,omitempty" xml:"CpuUsageBeforeAction,omitempty"`
}

func (s DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) SetCpuUsageAfterAction(v float32) *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors {
	s.CpuUsageAfterAction = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) SetAction(v string) *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors {
	s.Action = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) SetDestinationDedicatedHostId(v string) *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors {
	s.DestinationDedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) SetSourceDedicatedInstanceId(v string) *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors {
	s.SourceDedicatedInstanceId = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) SetActionDescription(v string) *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors {
	s.ActionDescription = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) SetSourceDedicatedInstanceRole(v string) *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors {
	s.SourceDedicatedInstanceRole = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) SetSourceDedicatedHostId(v string) *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors {
	s.SourceDedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors) SetCpuUsageBeforeAction(v float32) *DescribeDedicatedResourceAdvisorResponseBodyResourceAdvisorsResourceAdvisors {
	s.CpuUsageBeforeAction = &v
	return s
}

type DescribeDedicatedResourceAdvisorResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedResourceAdvisorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedResourceAdvisorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedResourceAdvisorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedResourceAdvisorResponse) SetHeaders(v map[string]*string) *DescribeDedicatedResourceAdvisorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedResourceAdvisorResponse) SetBody(v *DescribeDedicatedResourceAdvisorResponseBody) *DescribeDedicatedResourceAdvisorResponse {
	s.Body = v
	return s
}

type DescribeListUserBackupFileRecordRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OpsServiceVersion    *string `json:"OpsServiceVersion,omitempty" xml:"OpsServiceVersion,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeListUserBackupFileRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeListUserBackupFileRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeListUserBackupFileRecordRequest) SetOwnerId(v int64) *DescribeListUserBackupFileRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeListUserBackupFileRecordRequest) SetResourceOwnerAccount(v string) *DescribeListUserBackupFileRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeListUserBackupFileRecordRequest) SetResourceOwnerId(v int64) *DescribeListUserBackupFileRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeListUserBackupFileRecordRequest) SetStatus(v string) *DescribeListUserBackupFileRecordRequest {
	s.Status = &v
	return s
}

func (s *DescribeListUserBackupFileRecordRequest) SetEngine(v string) *DescribeListUserBackupFileRecordRequest {
	s.Engine = &v
	return s
}

func (s *DescribeListUserBackupFileRecordRequest) SetOpsServiceVersion(v string) *DescribeListUserBackupFileRecordRequest {
	s.OpsServiceVersion = &v
	return s
}

func (s *DescribeListUserBackupFileRecordRequest) SetRegionId(v string) *DescribeListUserBackupFileRecordRequest {
	s.RegionId = &v
	return s
}

type DescribeListUserBackupFileRecordResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Records   []*DescribeListUserBackupFileRecordResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeListUserBackupFileRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeListUserBackupFileRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListUserBackupFileRecordResponseBody) SetRequestId(v string) *DescribeListUserBackupFileRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBody) SetRecords(v []*DescribeListUserBackupFileRecordResponseBodyRecords) *DescribeListUserBackupFileRecordResponseBody {
	s.Records = v
	return s
}

type DescribeListUserBackupFileRecordResponseBodyRecords struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	OssFilePath     *string `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
	OssBucket       *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Bid             *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	EngineVersion   *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OssFileName     *string `json:"OssFileName,omitempty" xml:"OssFileName,omitempty"`
	OssFileSize     *int64  `json:"OssFileSize,omitempty" xml:"OssFileSize,omitempty"`
	GmtCreated      *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	BackupSetId     *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	DbInstanceName  *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	Engine          *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	BinlogInfo      *string `json:"BinlogInfo,omitempty" xml:"BinlogInfo,omitempty"`
	CustinsId       *string `json:"CustinsId,omitempty" xml:"CustinsId,omitempty"`
	OssFileMetaData *string `json:"OssFileMetaData,omitempty" xml:"OssFileMetaData,omitempty"`
	OssUrl          *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Reason          *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Uid             *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeListUserBackupFileRecordResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeListUserBackupFileRecordResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetStatus(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.Status = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetOssFilePath(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.OssFilePath = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetOssBucket(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.OssBucket = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetGmtModified(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.GmtModified = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetBid(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.Bid = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetEngineVersion(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.EngineVersion = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetOssFileName(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.OssFileName = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetOssFileSize(v int64) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.OssFileSize = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetGmtCreated(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.GmtCreated = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetBackupSetId(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.BackupSetId = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetInstanceName(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.InstanceName = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetDbInstanceName(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetEngine(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.Engine = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetBinlogInfo(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.BinlogInfo = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetCustinsId(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.CustinsId = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetOssFileMetaData(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.OssFileMetaData = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetOssUrl(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.OssUrl = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetTaskId(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.TaskId = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetId(v int64) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.Id = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetReason(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.Reason = &v
	return s
}

func (s *DescribeListUserBackupFileRecordResponseBodyRecords) SetUid(v string) *DescribeListUserBackupFileRecordResponseBodyRecords {
	s.Uid = &v
	return s
}

type DescribeListUserBackupFileRecordResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeListUserBackupFileRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeListUserBackupFileRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeListUserBackupFileRecordResponse) GoString() string {
	return s.String()
}

func (s *DescribeListUserBackupFileRecordResponse) SetHeaders(v map[string]*string) *DescribeListUserBackupFileRecordResponse {
	s.Headers = v
	return s
}

func (s *DescribeListUserBackupFileRecordResponse) SetBody(v *DescribeListUserBackupFileRecordResponseBody) *DescribeListUserBackupFileRecordResponse {
	s.Body = v
	return s
}

type CreateDedicatedHostAccountRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BastionInstanceId    *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateDedicatedHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostAccountRequest) SetOwnerId(v int64) *CreateDedicatedHostAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetResourceOwnerAccount(v string) *CreateDedicatedHostAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetResourceOwnerId(v int64) *CreateDedicatedHostAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetDedicatedHostId(v string) *CreateDedicatedHostAccountRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetAccountName(v string) *CreateDedicatedHostAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetAccountPassword(v string) *CreateDedicatedHostAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetRegionId(v string) *CreateDedicatedHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetBastionInstanceId(v string) *CreateDedicatedHostAccountRequest {
	s.BastionInstanceId = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetAccountType(v string) *CreateDedicatedHostAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetClientToken(v string) *CreateDedicatedHostAccountRequest {
	s.ClientToken = &v
	return s
}

type CreateDedicatedHostAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDedicatedHostAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostAccountResponseBody) SetRequestId(v string) *CreateDedicatedHostAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateDedicatedHostAccountResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDedicatedHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDedicatedHostAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostAccountResponse) SetHeaders(v map[string]*string) *CreateDedicatedHostAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateDedicatedHostAccountResponse) SetBody(v *CreateDedicatedHostAccountResponseBody) *CreateDedicatedHostAccountResponse {
	s.Body = v
	return s
}

type DeleteDedicatedHostAccountRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDedicatedHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedHostAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostAccountRequest) SetOwnerId(v int64) *DeleteDedicatedHostAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDedicatedHostAccountRequest) SetResourceOwnerAccount(v string) *DeleteDedicatedHostAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDedicatedHostAccountRequest) SetResourceOwnerId(v int64) *DeleteDedicatedHostAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDedicatedHostAccountRequest) SetDedicatedHostId(v string) *DeleteDedicatedHostAccountRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DeleteDedicatedHostAccountRequest) SetAccountName(v string) *DeleteDedicatedHostAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteDedicatedHostAccountRequest) SetRegionId(v string) *DeleteDedicatedHostAccountRequest {
	s.RegionId = &v
	return s
}

type DeleteDedicatedHostAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDedicatedHostAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostAccountResponseBody) SetRequestId(v string) *DeleteDedicatedHostAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDedicatedHostAccountResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDedicatedHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDedicatedHostAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedHostAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostAccountResponse) SetHeaders(v map[string]*string) *DeleteDedicatedHostAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteDedicatedHostAccountResponse) SetBody(v *DeleteDedicatedHostAccountResponseBody) *DeleteDedicatedHostAccountResponse {
	s.Body = v
	return s
}

type DeleteDedicatedHostGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDedicatedHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedHostGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostGroupRequest) SetOwnerId(v int64) *DeleteDedicatedHostGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDedicatedHostGroupRequest) SetResourceOwnerAccount(v string) *DeleteDedicatedHostGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDedicatedHostGroupRequest) SetResourceOwnerId(v int64) *DeleteDedicatedHostGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDedicatedHostGroupRequest) SetDedicatedHostGroupId(v string) *DeleteDedicatedHostGroupRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DeleteDedicatedHostGroupRequest) SetRegionId(v string) *DeleteDedicatedHostGroupRequest {
	s.RegionId = &v
	return s
}

type DeleteDedicatedHostGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDedicatedHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostGroupResponseBody) SetRequestId(v string) *DeleteDedicatedHostGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDedicatedHostGroupResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDedicatedHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDedicatedHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedHostGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostGroupResponse) SetHeaders(v map[string]*string) *DeleteDedicatedHostGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDedicatedHostGroupResponse) SetBody(v *DeleteDedicatedHostGroupResponseBody) *DeleteDedicatedHostGroupResponse {
	s.Body = v
	return s
}

type CheckUserIfAuthoriseMyBaseSystemRoleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	RamRoleName          *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
}

func (s CheckUserIfAuthoriseMyBaseSystemRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUserIfAuthoriseMyBaseSystemRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleRequest) SetOwnerId(v int64) *CheckUserIfAuthoriseMyBaseSystemRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleRequest) SetResourceOwnerAccount(v string) *CheckUserIfAuthoriseMyBaseSystemRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleRequest) SetResourceOwnerId(v int64) *CheckUserIfAuthoriseMyBaseSystemRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleRequest) SetRegionId(v string) *CheckUserIfAuthoriseMyBaseSystemRoleRequest {
	s.RegionId = &v
	return s
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleRequest) SetRegion(v string) *CheckUserIfAuthoriseMyBaseSystemRoleRequest {
	s.Region = &v
	return s
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleRequest) SetZoneId(v string) *CheckUserIfAuthoriseMyBaseSystemRoleRequest {
	s.ZoneId = &v
	return s
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleRequest) SetRamRoleName(v string) *CheckUserIfAuthoriseMyBaseSystemRoleRequest {
	s.RamRoleName = &v
	return s
}

type CheckUserIfAuthoriseMyBaseSystemRoleResponseBody struct {
	AliUid    *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleName  *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CheckUserIfAuthoriseMyBaseSystemRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserIfAuthoriseMyBaseSystemRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleResponseBody) SetAliUid(v string) *CheckUserIfAuthoriseMyBaseSystemRoleResponseBody {
	s.AliUid = &v
	return s
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleResponseBody) SetRequestId(v string) *CheckUserIfAuthoriseMyBaseSystemRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleResponseBody) SetRoleName(v string) *CheckUserIfAuthoriseMyBaseSystemRoleResponseBody {
	s.RoleName = &v
	return s
}

type CheckUserIfAuthoriseMyBaseSystemRoleResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckUserIfAuthoriseMyBaseSystemRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckUserIfAuthoriseMyBaseSystemRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserIfAuthoriseMyBaseSystemRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleResponse) SetHeaders(v map[string]*string) *CheckUserIfAuthoriseMyBaseSystemRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckUserIfAuthoriseMyBaseSystemRoleResponse) SetBody(v *CheckUserIfAuthoriseMyBaseSystemRoleResponseBody) *CheckUserIfAuthoriseMyBaseSystemRoleResponse {
	s.Body = v
	return s
}

type DescribeScheduleInstanceRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ScheduleId           *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
}

func (s DescribeScheduleInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduleInstanceRequest) SetOwnerId(v int64) *DescribeScheduleInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScheduleInstanceRequest) SetResourceOwnerAccount(v string) *DescribeScheduleInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScheduleInstanceRequest) SetResourceOwnerId(v int64) *DescribeScheduleInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScheduleInstanceRequest) SetRegionId(v string) *DescribeScheduleInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScheduleInstanceRequest) SetDedicatedHostGroupId(v string) *DescribeScheduleInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeScheduleInstanceRequest) SetScheduleId(v string) *DescribeScheduleInstanceRequest {
	s.ScheduleId = &v
	return s
}

type DescribeScheduleInstanceResponseBody struct {
	ScheduleId                 *string                                                           `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	RequestId                  *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceScheduleStatusList []*DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList `json:"InstanceScheduleStatusList,omitempty" xml:"InstanceScheduleStatusList,omitempty" type:"Repeated"`
}

func (s DescribeScheduleInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduleInstanceResponseBody) SetScheduleId(v string) *DescribeScheduleInstanceResponseBody {
	s.ScheduleId = &v
	return s
}

func (s *DescribeScheduleInstanceResponseBody) SetRequestId(v string) *DescribeScheduleInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduleInstanceResponseBody) SetInstanceScheduleStatusList(v []*DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList) *DescribeScheduleInstanceResponseBody {
	s.InstanceScheduleStatusList = v
	return s
}

type DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList struct {
	ScheduleStatus           *string `json:"ScheduleStatus,omitempty" xml:"ScheduleStatus,omitempty"`
	InstanceNodeTargetHostId *string `json:"InstanceNodeTargetHostId,omitempty" xml:"InstanceNodeTargetHostId,omitempty"`
	InstanceNodeRole         *string `json:"InstanceNodeRole,omitempty" xml:"InstanceNodeRole,omitempty"`
	Engine                   *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceId               *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceNodeSourceHostId *string `json:"InstanceNodeSourceHostId,omitempty" xml:"InstanceNodeSourceHostId,omitempty"`
}

func (s DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList) GoString() string {
	return s.String()
}

func (s *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList) SetScheduleStatus(v string) *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList {
	s.ScheduleStatus = &v
	return s
}

func (s *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList) SetInstanceNodeTargetHostId(v string) *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList {
	s.InstanceNodeTargetHostId = &v
	return s
}

func (s *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList) SetInstanceNodeRole(v string) *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList {
	s.InstanceNodeRole = &v
	return s
}

func (s *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList) SetEngine(v string) *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList {
	s.Engine = &v
	return s
}

func (s *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList) SetInstanceId(v string) *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList {
	s.InstanceId = &v
	return s
}

func (s *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList) SetInstanceNodeSourceHostId(v string) *DescribeScheduleInstanceResponseBodyInstanceScheduleStatusList {
	s.InstanceNodeSourceHostId = &v
	return s
}

type DescribeScheduleInstanceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScheduleInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScheduleInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduleInstanceResponse) SetHeaders(v map[string]*string) *DescribeScheduleInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduleInstanceResponse) SetBody(v *DescribeScheduleInstanceResponseBody) *DescribeScheduleInstanceResponse {
	s.Body = v
	return s
}

type ExcuteScheduleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ActionType           *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
}

func (s ExcuteScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s ExcuteScheduleRequest) GoString() string {
	return s.String()
}

func (s *ExcuteScheduleRequest) SetOwnerId(v int64) *ExcuteScheduleRequest {
	s.OwnerId = &v
	return s
}

func (s *ExcuteScheduleRequest) SetResourceOwnerAccount(v string) *ExcuteScheduleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ExcuteScheduleRequest) SetResourceOwnerId(v int64) *ExcuteScheduleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ExcuteScheduleRequest) SetRegionId(v string) *ExcuteScheduleRequest {
	s.RegionId = &v
	return s
}

func (s *ExcuteScheduleRequest) SetDedicatedHostGroupId(v string) *ExcuteScheduleRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *ExcuteScheduleRequest) SetActionType(v string) *ExcuteScheduleRequest {
	s.ActionType = &v
	return s
}

type ExcuteScheduleResponseBody struct {
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExcuteScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExcuteScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *ExcuteScheduleResponseBody) SetScheduleId(v string) *ExcuteScheduleResponseBody {
	s.ScheduleId = &v
	return s
}

func (s *ExcuteScheduleResponseBody) SetRequestId(v string) *ExcuteScheduleResponseBody {
	s.RequestId = &v
	return s
}

type ExcuteScheduleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExcuteScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExcuteScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s ExcuteScheduleResponse) GoString() string {
	return s.String()
}

func (s *ExcuteScheduleResponse) SetHeaders(v map[string]*string) *ExcuteScheduleResponse {
	s.Headers = v
	return s
}

func (s *ExcuteScheduleResponse) SetBody(v *ExcuteScheduleResponseBody) *ExcuteScheduleResponse {
	s.Body = v
	return s
}

type ReplaceDedicatedHostRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	FailoverMode         *string `json:"FailoverMode,omitempty" xml:"FailoverMode,omitempty"`
}

func (s ReplaceDedicatedHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceDedicatedHostRequest) GoString() string {
	return s.String()
}

func (s *ReplaceDedicatedHostRequest) SetOwnerId(v int64) *ReplaceDedicatedHostRequest {
	s.OwnerId = &v
	return s
}

func (s *ReplaceDedicatedHostRequest) SetResourceOwnerAccount(v string) *ReplaceDedicatedHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReplaceDedicatedHostRequest) SetResourceOwnerId(v int64) *ReplaceDedicatedHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReplaceDedicatedHostRequest) SetRegionId(v string) *ReplaceDedicatedHostRequest {
	s.RegionId = &v
	return s
}

func (s *ReplaceDedicatedHostRequest) SetDedicatedHostId(v string) *ReplaceDedicatedHostRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ReplaceDedicatedHostRequest) SetFailoverMode(v string) *ReplaceDedicatedHostRequest {
	s.FailoverMode = &v
	return s
}

type ReplaceDedicatedHostResponseBody struct {
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId          *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ReplaceDedicatedHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplaceDedicatedHostResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceDedicatedHostResponseBody) SetDedicatedHostId(v string) *ReplaceDedicatedHostResponseBody {
	s.DedicatedHostId = &v
	return s
}

func (s *ReplaceDedicatedHostResponseBody) SetRequestId(v string) *ReplaceDedicatedHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceDedicatedHostResponseBody) SetTaskId(v int32) *ReplaceDedicatedHostResponseBody {
	s.TaskId = &v
	return s
}

type ReplaceDedicatedHostResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReplaceDedicatedHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceDedicatedHostResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceDedicatedHostResponse) GoString() string {
	return s.String()
}

func (s *ReplaceDedicatedHostResponse) SetHeaders(v map[string]*string) *ReplaceDedicatedHostResponse {
	s.Headers = v
	return s
}

func (s *ReplaceDedicatedHostResponse) SetBody(v *ReplaceDedicatedHostResponseBody) *ReplaceDedicatedHostResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostAccountRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDedicatedHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAccountRequest) SetOwnerId(v int64) *ModifyDedicatedHostAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetDedicatedHostId(v string) *ModifyDedicatedHostAccountRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetAccountName(v string) *ModifyDedicatedHostAccountRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetAccountPassword(v string) *ModifyDedicatedHostAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetRegionId(v string) *ModifyDedicatedHostAccountRequest {
	s.RegionId = &v
	return s
}

type ModifyDedicatedHostAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedHostAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAccountResponseBody) SetRequestId(v string) *ModifyDedicatedHostAccountResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDedicatedHostAccountResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDedicatedHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDedicatedHostAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostAccountResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAccountResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostAccountResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostAccountResponse) SetBody(v *ModifyDedicatedHostAccountResponseBody) *ModifyDedicatedHostAccountResponse {
	s.Body = v
	return s
}

type DescribeHostEcsLevelInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DbType               *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	StorageType          *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeHostEcsLevelInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostEcsLevelInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostEcsLevelInfoRequest) SetOwnerId(v int64) *DescribeHostEcsLevelInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetResourceOwnerAccount(v string) *DescribeHostEcsLevelInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetResourceOwnerId(v int64) *DescribeHostEcsLevelInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetDbType(v string) *DescribeHostEcsLevelInfoRequest {
	s.DbType = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetRegionId(v string) *DescribeHostEcsLevelInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetZoneId(v string) *DescribeHostEcsLevelInfoRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetStorageType(v string) *DescribeHostEcsLevelInfoRequest {
	s.StorageType = &v
	return s
}

type DescribeHostEcsLevelInfoResponseBody struct {
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostEcsLevelInfos []*DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos `json:"HostEcsLevelInfos,omitempty" xml:"HostEcsLevelInfos,omitempty" type:"Repeated"`
}

func (s DescribeHostEcsLevelInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostEcsLevelInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostEcsLevelInfoResponseBody) SetRequestId(v string) *DescribeHostEcsLevelInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBody) SetHostEcsLevelInfos(v []*DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos) *DescribeHostEcsLevelInfoResponseBody {
	s.HostEcsLevelInfos = v
	return s
}

type DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos struct {
	CddcHostType *string                                                       `json:"CddcHostType,omitempty" xml:"CddcHostType,omitempty"`
	Items        []*DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos) GoString() string {
	return s.String()
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos) SetCddcHostType(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos {
	s.CddcHostType = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos) SetItems(v []*DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos {
	s.Items = v
	return s
}

type DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems struct {
	NetBandWidth          *float32 `json:"NetBandWidth,omitempty" xml:"NetBandWidth,omitempty"`
	EcsClass              *string  `json:"EcsClass,omitempty" xml:"EcsClass,omitempty"`
	RdsClassCode          *string  `json:"RdsClassCode,omitempty" xml:"RdsClassCode,omitempty"`
	Cpu                   *int32   `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CpuFrequency          *string  `json:"CpuFrequency,omitempty" xml:"CpuFrequency,omitempty"`
	StorageIops           *int32   `json:"StorageIops,omitempty" xml:"StorageIops,omitempty"`
	CloudStorageBandwidth *float32 `json:"CloudStorageBandwidth,omitempty" xml:"CloudStorageBandwidth,omitempty"`
	EcsClassCode          *string  `json:"EcsClassCode,omitempty" xml:"EcsClassCode,omitempty"`
	IsCloudDisk           *int32   `json:"IsCloudDisk,omitempty" xml:"IsCloudDisk,omitempty"`
	Memory                *int32   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NetPackage            *int32   `json:"NetPackage,omitempty" xml:"NetPackage,omitempty"`
	CpuVersion            *string  `json:"CpuVersion,omitempty" xml:"CpuVersion,omitempty"`
	LocalStorage          *string  `json:"LocalStorage,omitempty" xml:"LocalStorage,omitempty"`
	Description           *string  `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) GoString() string {
	return s.String()
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetNetBandWidth(v float32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.NetBandWidth = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetEcsClass(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.EcsClass = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetRdsClassCode(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.RdsClassCode = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetCpu(v int32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.Cpu = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetCpuFrequency(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.CpuFrequency = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetStorageIops(v int32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.StorageIops = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetCloudStorageBandwidth(v float32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.CloudStorageBandwidth = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetEcsClassCode(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.EcsClassCode = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetIsCloudDisk(v int32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.IsCloudDisk = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetMemory(v int32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.Memory = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetNetPackage(v int32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.NetPackage = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetCpuVersion(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.CpuVersion = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetLocalStorage(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.LocalStorage = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetDescription(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.Description = &v
	return s
}

type DescribeHostEcsLevelInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHostEcsLevelInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHostEcsLevelInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostEcsLevelInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeHostEcsLevelInfoResponse) SetHeaders(v map[string]*string) *DescribeHostEcsLevelInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeHostEcsLevelInfoResponse) SetBody(v *DescribeHostEcsLevelInfoResponseBody) *DescribeHostEcsLevelInfoResponse {
	s.Body = v
	return s
}

type AllocateHostInstanceCrossVpcVipRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ConnectionString     *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	BindVpcId            *string `json:"BindVpcId,omitempty" xml:"BindVpcId,omitempty"`
	BindVSwitchId        *string `json:"BindVSwitchId,omitempty" xml:"BindVSwitchId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocateHostInstanceCrossVpcVipRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateHostInstanceCrossVpcVipRequest) GoString() string {
	return s.String()
}

func (s *AllocateHostInstanceCrossVpcVipRequest) SetOwnerId(v int64) *AllocateHostInstanceCrossVpcVipRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipRequest) SetResourceOwnerAccount(v string) *AllocateHostInstanceCrossVpcVipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipRequest) SetResourceOwnerId(v int64) *AllocateHostInstanceCrossVpcVipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipRequest) SetConnectionString(v string) *AllocateHostInstanceCrossVpcVipRequest {
	s.ConnectionString = &v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipRequest) SetPort(v string) *AllocateHostInstanceCrossVpcVipRequest {
	s.Port = &v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipRequest) SetDBInstanceId(v string) *AllocateHostInstanceCrossVpcVipRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipRequest) SetBindVpcId(v string) *AllocateHostInstanceCrossVpcVipRequest {
	s.BindVpcId = &v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipRequest) SetBindVSwitchId(v string) *AllocateHostInstanceCrossVpcVipRequest {
	s.BindVSwitchId = &v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipRequest) SetRegionId(v string) *AllocateHostInstanceCrossVpcVipRequest {
	s.RegionId = &v
	return s
}

type AllocateHostInstanceCrossVpcVipResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateHostInstanceCrossVpcVipResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateHostInstanceCrossVpcVipResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateHostInstanceCrossVpcVipResponseBody) SetTaskId(v string) *AllocateHostInstanceCrossVpcVipResponseBody {
	s.TaskId = &v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipResponseBody) SetRequestId(v string) *AllocateHostInstanceCrossVpcVipResponseBody {
	s.RequestId = &v
	return s
}

type AllocateHostInstanceCrossVpcVipResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocateHostInstanceCrossVpcVipResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateHostInstanceCrossVpcVipResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateHostInstanceCrossVpcVipResponse) GoString() string {
	return s.String()
}

func (s *AllocateHostInstanceCrossVpcVipResponse) SetHeaders(v map[string]*string) *AllocateHostInstanceCrossVpcVipResponse {
	s.Headers = v
	return s
}

func (s *AllocateHostInstanceCrossVpcVipResponse) SetBody(v *AllocateHostInstanceCrossVpcVipResponseBody) *AllocateHostInstanceCrossVpcVipResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	OrderId              *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	HostType             *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	HostStatus           *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	AllocationStatus     *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	PageNumbers          *int32  `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDedicatedHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsRequest) SetOwnerId(v int64) *DescribeDedicatedHostsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetRegionId(v string) *DescribeDedicatedHostsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetOrderId(v int64) *DescribeDedicatedHostsRequest {
	s.OrderId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetHostType(v string) *DescribeDedicatedHostsRequest {
	s.HostType = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetHostStatus(v string) *DescribeDedicatedHostsRequest {
	s.HostStatus = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetAllocationStatus(v string) *DescribeDedicatedHostsRequest {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetZoneId(v string) *DescribeDedicatedHostsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostId(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetPageNumbers(v int32) *DescribeDedicatedHostsRequest {
	s.PageNumbers = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetPageSize(v int32) *DescribeDedicatedHostsRequest {
	s.PageSize = &v
	return s
}

type DescribeDedicatedHostsResponseBody struct {
	TotalRecords         *int32                                            `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
	PageSize             *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumbers          *int32                                            `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	DedicatedHostGroupId *string                                           `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHosts       *DescribeDedicatedHostsResponseBodyDedicatedHosts `json:"DedicatedHosts,omitempty" xml:"DedicatedHosts,omitempty" type:"Struct"`
}

func (s DescribeDedicatedHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBody) SetTotalRecords(v int32) *DescribeDedicatedHostsResponseBody {
	s.TotalRecords = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetPageSize(v int32) *DescribeDedicatedHostsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetRequestId(v string) *DescribeDedicatedHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetPageNumbers(v int32) *DescribeDedicatedHostsResponseBody {
	s.PageNumbers = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetDedicatedHosts(v *DescribeDedicatedHostsResponseBodyDedicatedHosts) *DescribeDedicatedHostsResponseBody {
	s.DedicatedHosts = v
	return s
}

type DescribeDedicatedHostsResponseBodyDedicatedHosts struct {
	DedicatedHosts []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts `json:"DedicatedHosts,omitempty" xml:"DedicatedHosts,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHosts) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHosts) SetDedicatedHosts(v []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) *DescribeDedicatedHostsResponseBodyDedicatedHosts {
	s.DedicatedHosts = v
	return s
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts struct {
	DeployType           *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	HostType             *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	HostStorage          *string `json:"HostStorage,omitempty" xml:"HostStorage,omitempty"`
	MemoryUsed           *string `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	AllocationStatus     *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	StorageUsed          *string `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	EcsClassCode         *string `json:"EcsClassCode,omitempty" xml:"EcsClassCode,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	MemAllocationRatio   *string `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	CreatedTime          *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	IPAddress            *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	HostStatus           *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	HostName             *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostCPU              *string `json:"HostCPU,omitempty" xml:"HostCPU,omitempty"`
	CpuUsed              *string `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty"`
	InstanceNumber       *string `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	OpenPermission       *string `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	DistributionSymbol   *string `json:"DistributionSymbol,omitempty" xml:"DistributionSymbol,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	HostClass            *string `json:"HostClass,omitempty" xml:"HostClass,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	CPUAllocationRatio   *string `json:"CPUAllocationRatio,omitempty" xml:"CPUAllocationRatio,omitempty"`
	ImageCategory        *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	DiskAllocationRatio  *string `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	HostMem              *string `json:"HostMem,omitempty" xml:"HostMem,omitempty"`
	BastionInstanceId    *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDeployType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DeployType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostStorage(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostStorage = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetMemoryUsed(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.MemoryUsed = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetAllocationStatus(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetStorageUsed(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetEcsClassCode(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.EcsClassCode = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDedicatedHostId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetMemAllocationRatio(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.MemAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCreatedTime(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetIPAddress(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.IPAddress = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostStatus(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostStatus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostName(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostName = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostCPU(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostCPU = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCpuUsed(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.CpuUsed = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetInstanceNumber(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetOpenPermission(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.OpenPermission = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDistributionSymbol(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DistributionSymbol = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetVPCId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.VPCId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostClass(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostClass = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetEndTime(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.EndTime = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetVSwitchId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetZoneId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCPUAllocationRatio(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.CPUAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetImageCategory(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.ImageCategory = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetEngine(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDiskAllocationRatio(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DiskAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostMem(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostMem = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetBastionInstanceId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.BastionInstanceId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetAccountName(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.AccountName = &v
	return s
}

type DescribeDedicatedHostsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostsResponse) SetBody(v *DescribeDedicatedHostsResponseBody) *DescribeDedicatedHostsResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostDisksRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDedicatedHostDisksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostDisksRequest) SetOwnerId(v int64) *DescribeDedicatedHostDisksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostDisksRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostDisksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostDisksRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostDisksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostDisksRequest) SetDedicatedHostId(v string) *DescribeDedicatedHostDisksRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostDisksRequest) SetRegionId(v string) *DescribeDedicatedHostDisksRequest {
	s.RegionId = &v
	return s
}

type DescribeDedicatedHostDisksResponseBody struct {
	DedicatedHostId *string                                        `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Disks           []*DescribeDedicatedHostDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostDisksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostDisksResponseBody) SetDedicatedHostId(v string) *DescribeDedicatedHostDisksResponseBody {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBody) SetRequestId(v string) *DescribeDedicatedHostDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBody) SetDisks(v []*DescribeDedicatedHostDisksResponseBodyDisks) *DescribeDedicatedHostDisksResponseBody {
	s.Disks = v
	return s
}

type DescribeDedicatedHostDisksResponseBodyDisks struct {
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	DiskId           *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	MaxThroughput    *int32  `json:"MaxThroughput,omitempty" xml:"MaxThroughput,omitempty"`
	MaxIOPS          *int32  `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	HasDBInstance    *bool   `json:"HasDBInstance,omitempty" xml:"HasDBInstance,omitempty"`
	Device           *string `json:"Device,omitempty" xml:"Device,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDedicatedHostDisksResponseBodyDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetType(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Type = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetStatus(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetPerformanceLevel(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetDiskId(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetMaxThroughput(v int32) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.MaxThroughput = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetMaxIOPS(v int32) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetHasDBInstance(v bool) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.HasDBInstance = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetDevice(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Device = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetSize(v int32) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetZoneId(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetCategory(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetDBInstanceId(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.DBInstanceId = &v
	return s
}

type DescribeDedicatedHostDisksResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedHostDisksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedHostDisksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostDisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostDisksResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostDisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostDisksResponse) SetBody(v *DescribeDedicatedHostDisksResponseBody) *DescribeDedicatedHostDisksResponse {
	s.Body = v
	return s
}

type DescribeMyBaseInstanceOverViewRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeMyBaseInstanceOverViewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseInstanceOverViewRequest) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseInstanceOverViewRequest) SetOwnerId(v int64) *DescribeMyBaseInstanceOverViewRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewRequest) SetResourceOwnerAccount(v string) *DescribeMyBaseInstanceOverViewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewRequest) SetResourceOwnerId(v int64) *DescribeMyBaseInstanceOverViewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewRequest) SetRegionId(v string) *DescribeMyBaseInstanceOverViewRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewRequest) SetRegion(v string) *DescribeMyBaseInstanceOverViewRequest {
	s.Region = &v
	return s
}

type DescribeMyBaseInstanceOverViewResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeMyBaseInstanceOverViewResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeMyBaseInstanceOverViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseInstanceOverViewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseInstanceOverViewResponseBody) SetRequestId(v string) *DescribeMyBaseInstanceOverViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBody) SetRegions(v *DescribeMyBaseInstanceOverViewResponseBodyRegions) *DescribeMyBaseInstanceOverViewResponseBody {
	s.Regions = v
	return s
}

type DescribeMyBaseInstanceOverViewResponseBodyRegions struct {
	RegionModel []*DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel `json:"RegionModel,omitempty" xml:"RegionModel,omitempty" type:"Repeated"`
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegions) SetRegionModel(v []*DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel) *DescribeMyBaseInstanceOverViewResponseBodyRegions {
	s.RegionModel = v
	return s
}

type DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel struct {
	EngineCount *string                                                                 `json:"EngineCount,omitempty" xml:"EngineCount,omitempty"`
	TotalCount  *int32                                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Region      *string                                                                 `json:"Region,omitempty" xml:"Region,omitempty"`
	TypeModels  *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModels `json:"TypeModels,omitempty" xml:"TypeModels,omitempty" type:"Struct"`
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel) SetEngineCount(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel {
	s.EngineCount = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel) SetTotalCount(v int32) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel {
	s.TotalCount = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel) SetRegion(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel {
	s.Region = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel) SetTypeModels(v *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModels) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModel {
	s.TypeModels = v
	return s
}

type DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModels struct {
	TypeModel []*DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel `json:"TypeModel,omitempty" xml:"TypeModel,omitempty" type:"Repeated"`
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModels) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModels) SetTypeModel(v []*DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModels {
	s.TypeModel = v
	return s
}

type DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel struct {
	Count               *int32                                                                                         `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceDateType    *string                                                                                        `json:"InstanceDateType,omitempty" xml:"InstanceDateType,omitempty"`
	InstanceEngineCount *string                                                                                        `json:"InstanceEngineCount,omitempty" xml:"InstanceEngineCount,omitempty"`
	InstanceModels      *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModels `json:"InstanceModels,omitempty" xml:"InstanceModels,omitempty" type:"Struct"`
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) SetCount(v int32) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel {
	s.Count = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) SetInstanceDateType(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel {
	s.InstanceDateType = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) SetInstanceEngineCount(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel {
	s.InstanceEngineCount = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel) SetInstanceModels(v *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModels) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModel {
	s.InstanceModels = v
	return s
}

type DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModels struct {
	InstanceModel []*DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel `json:"InstanceModel,omitempty" xml:"InstanceModel,omitempty" type:"Repeated"`
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModels) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModels) SetInstanceModel(v []*DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModels {
	s.InstanceModel = v
	return s
}

type DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel struct {
	ExpireTime       *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	CreatedTime      *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel) SetExpireTime(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel {
	s.ExpireTime = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel) SetPayType(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel {
	s.PayType = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel) SetZoneId(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel {
	s.ZoneId = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel) SetDBInstanceId(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel) SetDBInstanceStatus(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel) SetCreatedTime(v string) *DescribeMyBaseInstanceOverViewResponseBodyRegionsRegionModelTypeModelsTypeModelInstanceModelsInstanceModel {
	s.CreatedTime = &v
	return s
}

type DescribeMyBaseInstanceOverViewResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMyBaseInstanceOverViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMyBaseInstanceOverViewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMyBaseInstanceOverViewResponse) GoString() string {
	return s.String()
}

func (s *DescribeMyBaseInstanceOverViewResponse) SetHeaders(v map[string]*string) *DescribeMyBaseInstanceOverViewResponse {
	s.Headers = v
	return s
}

func (s *DescribeMyBaseInstanceOverViewResponse) SetBody(v *DescribeMyBaseInstanceOverViewResponseBody) *DescribeMyBaseInstanceOverViewResponse {
	s.Body = v
	return s
}

type ModifyScheduleMethodRequest struct {
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostGroupId      *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	Engine                    *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceDistribution      *string `json:"InstanceDistribution,omitempty" xml:"InstanceDistribution,omitempty"`
	ZoneDistribution          *string `json:"ZoneDistribution,omitempty" xml:"ZoneDistribution,omitempty"`
	CpuUtilityThreshold       *string `json:"CpuUtilityThreshold,omitempty" xml:"CpuUtilityThreshold,omitempty"`
	MemoryUtilityThreshold    *string `json:"MemoryUtilityThreshold,omitempty" xml:"MemoryUtilityThreshold,omitempty"`
	LocalDiskUtilityThreshold *string `json:"LocalDiskUtilityThreshold,omitempty" xml:"LocalDiskUtilityThreshold,omitempty"`
	ZonesOrder                *string `json:"ZonesOrder,omitempty" xml:"ZonesOrder,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyScheduleMethodRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduleMethodRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduleMethodRequest) SetOwnerId(v int64) *ModifyScheduleMethodRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetResourceOwnerAccount(v string) *ModifyScheduleMethodRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetResourceOwnerId(v int64) *ModifyScheduleMethodRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetDedicatedHostGroupId(v string) *ModifyScheduleMethodRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetEngine(v string) *ModifyScheduleMethodRequest {
	s.Engine = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetInstanceDistribution(v string) *ModifyScheduleMethodRequest {
	s.InstanceDistribution = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetZoneDistribution(v string) *ModifyScheduleMethodRequest {
	s.ZoneDistribution = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetCpuUtilityThreshold(v string) *ModifyScheduleMethodRequest {
	s.CpuUtilityThreshold = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetMemoryUtilityThreshold(v string) *ModifyScheduleMethodRequest {
	s.MemoryUtilityThreshold = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetLocalDiskUtilityThreshold(v string) *ModifyScheduleMethodRequest {
	s.LocalDiskUtilityThreshold = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetZonesOrder(v string) *ModifyScheduleMethodRequest {
	s.ZonesOrder = &v
	return s
}

func (s *ModifyScheduleMethodRequest) SetRegionId(v string) *ModifyScheduleMethodRequest {
	s.RegionId = &v
	return s
}

type ModifyScheduleMethodResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScheduleMethodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduleMethodResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduleMethodResponseBody) SetRequestId(v string) *ModifyScheduleMethodResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScheduleMethodResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScheduleMethodResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScheduleMethodResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduleMethodResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduleMethodResponse) SetHeaders(v map[string]*string) *ModifyScheduleMethodResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduleMethodResponse) SetBody(v *ModifyScheduleMethodResponseBody) *ModifyScheduleMethodResponse {
	s.Body = v
	return s
}

type DescribeAvailableDedicatedHostClassesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	StorageType          *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	DbType               *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
}

func (s DescribeAvailableDedicatedHostClassesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostClassesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostClassesRequest) SetOwnerId(v int64) *DescribeAvailableDedicatedHostClassesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostClassesRequest) SetResourceOwnerAccount(v string) *DescribeAvailableDedicatedHostClassesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableDedicatedHostClassesRequest) SetResourceOwnerId(v int64) *DescribeAvailableDedicatedHostClassesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostClassesRequest) SetRegionId(v string) *DescribeAvailableDedicatedHostClassesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostClassesRequest) SetZoneId(v string) *DescribeAvailableDedicatedHostClassesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostClassesRequest) SetStorageType(v string) *DescribeAvailableDedicatedHostClassesRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailableDedicatedHostClassesRequest) SetDbType(v string) *DescribeAvailableDedicatedHostClassesRequest {
	s.DbType = &v
	return s
}

type DescribeAvailableDedicatedHostClassesResponseBody struct {
	RequestId   *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostClasses *DescribeAvailableDedicatedHostClassesResponseBodyHostClasses `json:"HostClasses,omitempty" xml:"HostClasses,omitempty" type:"Struct"`
}

func (s DescribeAvailableDedicatedHostClassesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostClassesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostClassesResponseBody) SetRequestId(v string) *DescribeAvailableDedicatedHostClassesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableDedicatedHostClassesResponseBody) SetHostClasses(v *DescribeAvailableDedicatedHostClassesResponseBodyHostClasses) *DescribeAvailableDedicatedHostClassesResponseBody {
	s.HostClasses = v
	return s
}

type DescribeAvailableDedicatedHostClassesResponseBodyHostClasses struct {
	HostClasses []*DescribeAvailableDedicatedHostClassesResponseBodyHostClassesHostClasses `json:"HostClasses,omitempty" xml:"HostClasses,omitempty" type:"Repeated"`
}

func (s DescribeAvailableDedicatedHostClassesResponseBodyHostClasses) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostClassesResponseBodyHostClasses) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostClassesResponseBodyHostClasses) SetHostClasses(v []*DescribeAvailableDedicatedHostClassesResponseBodyHostClassesHostClasses) *DescribeAvailableDedicatedHostClassesResponseBodyHostClasses {
	s.HostClasses = v
	return s
}

type DescribeAvailableDedicatedHostClassesResponseBodyHostClassesHostClasses struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HostClassName *string `json:"HostClassName,omitempty" xml:"HostClassName,omitempty"`
}

func (s DescribeAvailableDedicatedHostClassesResponseBodyHostClassesHostClasses) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostClassesResponseBodyHostClassesHostClasses) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostClassesResponseBodyHostClassesHostClasses) SetDescription(v string) *DescribeAvailableDedicatedHostClassesResponseBodyHostClassesHostClasses {
	s.Description = &v
	return s
}

func (s *DescribeAvailableDedicatedHostClassesResponseBodyHostClassesHostClasses) SetHostClassName(v string) *DescribeAvailableDedicatedHostClassesResponseBodyHostClassesHostClasses {
	s.HostClassName = &v
	return s
}

type DescribeAvailableDedicatedHostClassesResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableDedicatedHostClassesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableDedicatedHostClassesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableDedicatedHostClassesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableDedicatedHostClassesResponse) SetHeaders(v map[string]*string) *DescribeAvailableDedicatedHostClassesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableDedicatedHostClassesResponse) SetBody(v *DescribeAvailableDedicatedHostClassesResponseBody) *DescribeAvailableDedicatedHostClassesResponse {
	s.Body = v
	return s
}

type DescribeCheckUserBackupFileRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	BucketRegion         *string `json:"BucketRegion,omitempty" xml:"BucketRegion,omitempty"`
	BackupFile           *string `json:"BackupFile,omitempty" xml:"BackupFile,omitempty"`
}

func (s DescribeCheckUserBackupFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckUserBackupFileRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckUserBackupFileRequest) SetOwnerId(v int64) *DescribeCheckUserBackupFileRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCheckUserBackupFileRequest) SetResourceOwnerAccount(v string) *DescribeCheckUserBackupFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCheckUserBackupFileRequest) SetResourceOwnerId(v int64) *DescribeCheckUserBackupFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCheckUserBackupFileRequest) SetRegionId(v string) *DescribeCheckUserBackupFileRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCheckUserBackupFileRequest) SetEngineVersion(v string) *DescribeCheckUserBackupFileRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeCheckUserBackupFileRequest) SetBucketRegion(v string) *DescribeCheckUserBackupFileRequest {
	s.BucketRegion = &v
	return s
}

func (s *DescribeCheckUserBackupFileRequest) SetBackupFile(v string) *DescribeCheckUserBackupFileRequest {
	s.BackupFile = &v
	return s
}

type DescribeCheckUserBackupFileResponseBody struct {
	Status         *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId         *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeCheckUserBackupFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckUserBackupFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckUserBackupFileResponseBody) SetStatus(v bool) *DescribeCheckUserBackupFileResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCheckUserBackupFileResponseBody) SetDBInstanceName(v string) *DescribeCheckUserBackupFileResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCheckUserBackupFileResponseBody) SetRequestId(v string) *DescribeCheckUserBackupFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckUserBackupFileResponseBody) SetTaskId(v int64) *DescribeCheckUserBackupFileResponseBody {
	s.TaskId = &v
	return s
}

type DescribeCheckUserBackupFileResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCheckUserBackupFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCheckUserBackupFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckUserBackupFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckUserBackupFileResponse) SetHeaders(v map[string]*string) *DescribeCheckUserBackupFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckUserBackupFileResponse) SetBody(v *DescribeCheckUserBackupFileResponseBody) *DescribeCheckUserBackupFileResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostPasswordRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NewPassword          *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
}

func (s ModifyDedicatedHostPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostPasswordRequest) SetOwnerId(v int64) *ModifyDedicatedHostPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostPasswordRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostPasswordRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostPasswordRequest) SetRegionId(v string) *ModifyDedicatedHostPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostPasswordRequest) SetNewPassword(v string) *ModifyDedicatedHostPasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ModifyDedicatedHostPasswordRequest) SetDedicatedHostId(v string) *ModifyDedicatedHostPasswordRequest {
	s.DedicatedHostId = &v
	return s
}

type ModifyDedicatedHostPasswordResponseBody struct {
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedHostPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostPasswordResponseBody) SetDedicatedHostName(v string) *ModifyDedicatedHostPasswordResponseBody {
	s.DedicatedHostName = &v
	return s
}

func (s *ModifyDedicatedHostPasswordResponseBody) SetRequestId(v string) *ModifyDedicatedHostPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDedicatedHostPasswordResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDedicatedHostPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDedicatedHostPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostPasswordResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostPasswordResponse) SetBody(v *ModifyDedicatedHostPasswordResponseBody) *ModifyDedicatedHostPasswordResponse {
	s.Body = v
	return s
}

type DescribeScheduleMethodsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
}

func (s DescribeScheduleMethodsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleMethodsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduleMethodsRequest) SetOwnerId(v int64) *DescribeScheduleMethodsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScheduleMethodsRequest) SetResourceOwnerAccount(v string) *DescribeScheduleMethodsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScheduleMethodsRequest) SetResourceOwnerId(v int64) *DescribeScheduleMethodsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScheduleMethodsRequest) SetRegionId(v string) *DescribeScheduleMethodsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScheduleMethodsRequest) SetDedicatedHostGroupId(v string) *DescribeScheduleMethodsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

type DescribeScheduleMethodsResponseBody struct {
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScheduleMethodList []*DescribeScheduleMethodsResponseBodyScheduleMethodList `json:"ScheduleMethodList,omitempty" xml:"ScheduleMethodList,omitempty" type:"Repeated"`
}

func (s DescribeScheduleMethodsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleMethodsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduleMethodsResponseBody) SetRequestId(v string) *DescribeScheduleMethodsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduleMethodsResponseBody) SetScheduleMethodList(v []*DescribeScheduleMethodsResponseBodyScheduleMethodList) *DescribeScheduleMethodsResponseBody {
	s.ScheduleMethodList = v
	return s
}

type DescribeScheduleMethodsResponseBodyScheduleMethodList struct {
	ZonesOrder                *string `json:"ZonesOrder,omitempty" xml:"ZonesOrder,omitempty"`
	ZoneDistribution          *string `json:"ZoneDistribution,omitempty" xml:"ZoneDistribution,omitempty"`
	CPUUtilityThreshold       *int32  `json:"CPUUtilityThreshold,omitempty" xml:"CPUUtilityThreshold,omitempty"`
	DedicatedHostGroupDesc    *string `json:"DedicatedHostGroupDesc,omitempty" xml:"DedicatedHostGroupDesc,omitempty"`
	MemoryUtilityThreshold    *int32  `json:"MemoryUtilityThreshold,omitempty" xml:"MemoryUtilityThreshold,omitempty"`
	Engine                    *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceDistribution      *string `json:"InstanceDistribution,omitempty" xml:"InstanceDistribution,omitempty"`
	LocalDiskUtilityThreshold *int32  `json:"LocalDiskUtilityThreshold,omitempty" xml:"LocalDiskUtilityThreshold,omitempty"`
	DedicatedHostGroupId      *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
}

func (s DescribeScheduleMethodsResponseBodyScheduleMethodList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleMethodsResponseBodyScheduleMethodList) GoString() string {
	return s.String()
}

func (s *DescribeScheduleMethodsResponseBodyScheduleMethodList) SetZonesOrder(v string) *DescribeScheduleMethodsResponseBodyScheduleMethodList {
	s.ZonesOrder = &v
	return s
}

func (s *DescribeScheduleMethodsResponseBodyScheduleMethodList) SetZoneDistribution(v string) *DescribeScheduleMethodsResponseBodyScheduleMethodList {
	s.ZoneDistribution = &v
	return s
}

func (s *DescribeScheduleMethodsResponseBodyScheduleMethodList) SetCPUUtilityThreshold(v int32) *DescribeScheduleMethodsResponseBodyScheduleMethodList {
	s.CPUUtilityThreshold = &v
	return s
}

func (s *DescribeScheduleMethodsResponseBodyScheduleMethodList) SetDedicatedHostGroupDesc(v string) *DescribeScheduleMethodsResponseBodyScheduleMethodList {
	s.DedicatedHostGroupDesc = &v
	return s
}

func (s *DescribeScheduleMethodsResponseBodyScheduleMethodList) SetMemoryUtilityThreshold(v int32) *DescribeScheduleMethodsResponseBodyScheduleMethodList {
	s.MemoryUtilityThreshold = &v
	return s
}

func (s *DescribeScheduleMethodsResponseBodyScheduleMethodList) SetEngine(v string) *DescribeScheduleMethodsResponseBodyScheduleMethodList {
	s.Engine = &v
	return s
}

func (s *DescribeScheduleMethodsResponseBodyScheduleMethodList) SetInstanceDistribution(v string) *DescribeScheduleMethodsResponseBodyScheduleMethodList {
	s.InstanceDistribution = &v
	return s
}

func (s *DescribeScheduleMethodsResponseBodyScheduleMethodList) SetLocalDiskUtilityThreshold(v int32) *DescribeScheduleMethodsResponseBodyScheduleMethodList {
	s.LocalDiskUtilityThreshold = &v
	return s
}

func (s *DescribeScheduleMethodsResponseBodyScheduleMethodList) SetDedicatedHostGroupId(v string) *DescribeScheduleMethodsResponseBodyScheduleMethodList {
	s.DedicatedHostGroupId = &v
	return s
}

type DescribeScheduleMethodsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScheduleMethodsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScheduleMethodsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleMethodsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduleMethodsResponse) SetHeaders(v map[string]*string) *DescribeScheduleMethodsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduleMethodsResponse) SetBody(v *DescribeScheduleMethodsResponseBody) *DescribeScheduleMethodsResponse {
	s.Body = v
	return s
}

type ClearDedicatedHostRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	FailoverMode         *string `json:"FailoverMode,omitempty" xml:"FailoverMode,omitempty"`
}

func (s ClearDedicatedHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearDedicatedHostRequest) GoString() string {
	return s.String()
}

func (s *ClearDedicatedHostRequest) SetOwnerId(v int64) *ClearDedicatedHostRequest {
	s.OwnerId = &v
	return s
}

func (s *ClearDedicatedHostRequest) SetResourceOwnerAccount(v string) *ClearDedicatedHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClearDedicatedHostRequest) SetResourceOwnerId(v int64) *ClearDedicatedHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClearDedicatedHostRequest) SetRegionId(v string) *ClearDedicatedHostRequest {
	s.RegionId = &v
	return s
}

func (s *ClearDedicatedHostRequest) SetDedicatedHostId(v string) *ClearDedicatedHostRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ClearDedicatedHostRequest) SetFailoverMode(v string) *ClearDedicatedHostRequest {
	s.FailoverMode = &v
	return s
}

type ClearDedicatedHostResponseBody struct {
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ClearDedicatedHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearDedicatedHostResponseBody) GoString() string {
	return s.String()
}

func (s *ClearDedicatedHostResponseBody) SetDedicatedHostId(v string) *ClearDedicatedHostResponseBody {
	s.DedicatedHostId = &v
	return s
}

func (s *ClearDedicatedHostResponseBody) SetRequestId(v string) *ClearDedicatedHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearDedicatedHostResponseBody) SetTaskId(v string) *ClearDedicatedHostResponseBody {
	s.TaskId = &v
	return s
}

type ClearDedicatedHostResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearDedicatedHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClearDedicatedHostResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearDedicatedHostResponse) GoString() string {
	return s.String()
}

func (s *ClearDedicatedHostResponse) SetHeaders(v map[string]*string) *ClearDedicatedHostResponse {
	s.Headers = v
	return s
}

func (s *ClearDedicatedHostResponse) SetBody(v *ClearDedicatedHostResponseBody) *ClearDedicatedHostResponse {
	s.Body = v
	return s
}

type DeleteUserBackupFileRecordRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OpsServiceVersion    *string `json:"OpsServiceVersion,omitempty" xml:"OpsServiceVersion,omitempty"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteUserBackupFileRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserBackupFileRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserBackupFileRecordRequest) SetOwnerId(v int64) *DeleteUserBackupFileRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUserBackupFileRecordRequest) SetResourceOwnerAccount(v string) *DeleteUserBackupFileRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteUserBackupFileRecordRequest) SetResourceOwnerId(v int64) *DeleteUserBackupFileRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUserBackupFileRecordRequest) SetEngine(v string) *DeleteUserBackupFileRecordRequest {
	s.Engine = &v
	return s
}

func (s *DeleteUserBackupFileRecordRequest) SetOpsServiceVersion(v string) *DeleteUserBackupFileRecordRequest {
	s.OpsServiceVersion = &v
	return s
}

func (s *DeleteUserBackupFileRecordRequest) SetInstanceName(v string) *DeleteUserBackupFileRecordRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteUserBackupFileRecordRequest) SetRegionId(v string) *DeleteUserBackupFileRecordRequest {
	s.RegionId = &v
	return s
}

type DeleteUserBackupFileRecordResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DeleteUserBackupFileRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserBackupFileRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserBackupFileRecordResponseBody) SetRequestId(v string) *DeleteUserBackupFileRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserBackupFileRecordResponseBody) SetSuccess(v string) *DeleteUserBackupFileRecordResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserBackupFileRecordResponseBody) SetId(v string) *DeleteUserBackupFileRecordResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteUserBackupFileRecordResponseBody) SetInstanceName(v string) *DeleteUserBackupFileRecordResponseBody {
	s.InstanceName = &v
	return s
}

type DeleteUserBackupFileRecordResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserBackupFileRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserBackupFileRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserBackupFileRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserBackupFileRecordResponse) SetHeaders(v map[string]*string) *DeleteUserBackupFileRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserBackupFileRecordResponse) SetBody(v *DeleteUserBackupFileRecordResponseBody) *DeleteUserBackupFileRecordResponse {
	s.Body = v
	return s
}

type DescribeEvaluateDedicatedHostsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ClassCode            *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	StorageType          *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StorageSize          *int32  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	NodeType             *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeEvaluateDedicatedHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluateDedicatedHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetOwnerId(v int64) *DescribeEvaluateDedicatedHostsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetResourceOwnerAccount(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetResourceOwnerId(v int64) *DescribeEvaluateDedicatedHostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetRegionId(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetDedicatedHostGroupId(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetClassCode(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.ClassCode = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetEngine(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.Engine = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetEngineVersion(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetZoneId(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetDedicatedHostId(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetStorageType(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetStorageSize(v int32) *DescribeEvaluateDedicatedHostsRequest {
	s.StorageSize = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsRequest) SetNodeType(v string) *DescribeEvaluateDedicatedHostsRequest {
	s.NodeType = &v
	return s
}

type DescribeEvaluateDedicatedHostsResponseBody struct {
	RequestId                 *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceRequiredCPU       *int32                                                      `json:"InstanceRequiredCPU,omitempty" xml:"InstanceRequiredCPU,omitempty"`
	CpuOverAllocationRatio    *int32                                                      `json:"CpuOverAllocationRatio,omitempty" xml:"CpuOverAllocationRatio,omitempty"`
	MemoryOverAllocationRatio *int32                                                      `json:"MemoryOverAllocationRatio,omitempty" xml:"MemoryOverAllocationRatio,omitempty"`
	InstanceRequiredMemory    *int32                                                      `json:"InstanceRequiredMemory,omitempty" xml:"InstanceRequiredMemory,omitempty"`
	InstanceClassCode         *string                                                     `json:"InstanceClassCode,omitempty" xml:"InstanceClassCode,omitempty"`
	InstanceRequiredStorage   *int32                                                      `json:"InstanceRequiredStorage,omitempty" xml:"InstanceRequiredStorage,omitempty"`
	DiskOverAllocationRatio   *int32                                                      `json:"DiskOverAllocationRatio,omitempty" xml:"DiskOverAllocationRatio,omitempty"`
	DedicatedHostGroupId      *string                                                     `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHosts            []*DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts `json:"DedicatedHosts,omitempty" xml:"DedicatedHosts,omitempty" type:"Repeated"`
}

func (s DescribeEvaluateDedicatedHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluateDedicatedHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetRequestId(v string) *DescribeEvaluateDedicatedHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetInstanceRequiredCPU(v int32) *DescribeEvaluateDedicatedHostsResponseBody {
	s.InstanceRequiredCPU = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetCpuOverAllocationRatio(v int32) *DescribeEvaluateDedicatedHostsResponseBody {
	s.CpuOverAllocationRatio = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetMemoryOverAllocationRatio(v int32) *DescribeEvaluateDedicatedHostsResponseBody {
	s.MemoryOverAllocationRatio = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetInstanceRequiredMemory(v int32) *DescribeEvaluateDedicatedHostsResponseBody {
	s.InstanceRequiredMemory = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetInstanceClassCode(v string) *DescribeEvaluateDedicatedHostsResponseBody {
	s.InstanceClassCode = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetInstanceRequiredStorage(v int32) *DescribeEvaluateDedicatedHostsResponseBody {
	s.InstanceRequiredStorage = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetDiskOverAllocationRatio(v int32) *DescribeEvaluateDedicatedHostsResponseBody {
	s.DiskOverAllocationRatio = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetDedicatedHostGroupId(v string) *DescribeEvaluateDedicatedHostsResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBody) SetDedicatedHosts(v []*DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) *DescribeEvaluateDedicatedHostsResponseBody {
	s.DedicatedHosts = v
	return s
}

type DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts struct {
	HostAllocationStatus   *bool   `json:"HostAllocationStatus,omitempty" xml:"HostAllocationStatus,omitempty"`
	Comment                *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostFreeMem            *int32  `json:"HostFreeMem,omitempty" xml:"HostFreeMem,omitempty"`
	HostStatus             *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	HostFreeCPU            *int32  `json:"HostFreeCPU,omitempty" xml:"HostFreeCPU,omitempty"`
	HostStorage            *int32  `json:"HostStorage,omitempty" xml:"HostStorage,omitempty"`
	HostCPU                *int32  `json:"HostCPU,omitempty" xml:"HostCPU,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostFreeStorage        *int32  `json:"HostFreeStorage,omitempty" xml:"HostFreeStorage,omitempty"`
	VSwitchId              *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	DedicatedHostName      *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	ZoneId                 *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ImageCategory          *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	DedicatedHostId        *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Engine                 *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsCouldSSD             *bool   `json:"IsCouldSSD,omitempty" xml:"IsCouldSSD,omitempty"`
	IsAvailableForInstance *bool   `json:"IsAvailableForInstance,omitempty" xml:"IsAvailableForInstance,omitempty"`
	HostMem                *int32  `json:"HostMem,omitempty" xml:"HostMem,omitempty"`
	CreatedTime            *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	IPAddress              *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetHostAllocationStatus(v bool) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.HostAllocationStatus = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetComment(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.Comment = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetHostFreeMem(v int32) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.HostFreeMem = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetHostStatus(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.HostStatus = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetHostFreeCPU(v int32) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.HostFreeCPU = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetHostStorage(v int32) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.HostStorage = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetHostCPU(v int32) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.HostCPU = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetRegionId(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.RegionId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetHostFreeStorage(v int32) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.HostFreeStorage = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetVSwitchId(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetDedicatedHostName(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetZoneId(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.ZoneId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetImageCategory(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.ImageCategory = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetDedicatedHostId(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetEngine(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.Engine = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetIsCouldSSD(v bool) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.IsCouldSSD = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetIsAvailableForInstance(v bool) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.IsAvailableForInstance = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetHostMem(v int32) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.HostMem = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetCreatedTime(v int64) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.CreatedTime = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts) SetIPAddress(v string) *DescribeEvaluateDedicatedHostsResponseBodyDedicatedHosts {
	s.IPAddress = &v
	return s
}

type DescribeEvaluateDedicatedHostsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEvaluateDedicatedHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEvaluateDedicatedHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluateDedicatedHostsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateDedicatedHostsResponse) SetHeaders(v map[string]*string) *DescribeEvaluateDedicatedHostsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEvaluateDedicatedHostsResponse) SetBody(v *DescribeEvaluateDedicatedHostsResponseBody) *DescribeEvaluateDedicatedHostsResponse {
	s.Body = v
	return s
}

type DescribeHostInstanceMonitorInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeHostInstanceMonitorInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostInstanceMonitorInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostInstanceMonitorInfoRequest) SetOwnerId(v int64) *DescribeHostInstanceMonitorInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoRequest) SetResourceOwnerAccount(v string) *DescribeHostInstanceMonitorInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoRequest) SetResourceOwnerId(v int64) *DescribeHostInstanceMonitorInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoRequest) SetDedicatedHostId(v string) *DescribeHostInstanceMonitorInfoRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoRequest) SetRegionId(v string) *DescribeHostInstanceMonitorInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoRequest) SetStartTime(v string) *DescribeHostInstanceMonitorInfoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoRequest) SetEndTime(v string) *DescribeHostInstanceMonitorInfoRequest {
	s.EndTime = &v
	return s
}

type DescribeHostInstanceMonitorInfoResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     []*DescribeHostInstanceMonitorInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeHostInstanceMonitorInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostInstanceMonitorInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostInstanceMonitorInfoResponseBody) SetRequestId(v string) *DescribeHostInstanceMonitorInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBody) SetItems(v []*DescribeHostInstanceMonitorInfoResponseBodyItems) *DescribeHostInstanceMonitorInfoResponseBody {
	s.Items = v
	return s
}

type DescribeHostInstanceMonitorInfoResponseBodyItems struct {
	Status                  *string                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	MemSize                 *int32                                                                   `json:"MemSize,omitempty" xml:"MemSize,omitempty"`
	DiskSize                *int32                                                                   `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	Ip                      *string                                                                  `json:"Ip,omitempty" xml:"Ip,omitempty"`
	DBInstanceId            *string                                                                  `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	LevelName               *string                                                                  `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	Engine                  *string                                                                  `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Role                    *string                                                                  `json:"Role,omitempty" xml:"Role,omitempty"`
	Port                    *string                                                                  `json:"Port,omitempty" xml:"Port,omitempty"`
	CpuCores                *int32                                                                   `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	EngineVersion           *string                                                                  `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	AuroraSwitchInstanceLog *DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog `json:"AuroraSwitchInstanceLog,omitempty" xml:"AuroraSwitchInstanceLog,omitempty" type:"Struct"`
}

func (s DescribeHostInstanceMonitorInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostInstanceMonitorInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetStatus(v string) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetMemSize(v int32) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.MemSize = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetDiskSize(v int32) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.DiskSize = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetIp(v string) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.Ip = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetDBInstanceId(v string) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetLevelName(v string) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.LevelName = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetEngine(v string) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.Engine = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetRole(v string) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.Role = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetPort(v string) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.Port = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetCpuCores(v int32) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.CpuCores = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetEngineVersion(v string) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.EngineVersion = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItems) SetAuroraSwitchInstanceLog(v *DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog) *DescribeHostInstanceMonitorInfoResponseBodyItems {
	s.AuroraSwitchInstanceLog = v
	return s
}

type DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog struct {
	SwitchFrom *float32 `json:"SwitchFrom,omitempty" xml:"SwitchFrom,omitempty"`
	SwitchType *float32 `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
	SwitchTime *int64   `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog) GoString() string {
	return s.String()
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog) SetSwitchFrom(v float32) *DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog {
	s.SwitchFrom = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog) SetSwitchType(v float32) *DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog {
	s.SwitchType = &v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog) SetSwitchTime(v int64) *DescribeHostInstanceMonitorInfoResponseBodyItemsAuroraSwitchInstanceLog {
	s.SwitchTime = &v
	return s
}

type DescribeHostInstanceMonitorInfoResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHostInstanceMonitorInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHostInstanceMonitorInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostInstanceMonitorInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeHostInstanceMonitorInfoResponse) SetHeaders(v map[string]*string) *DescribeHostInstanceMonitorInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeHostInstanceMonitorInfoResponse) SetBody(v *DescribeHostInstanceMonitorInfoResponseBody) *DescribeHostInstanceMonitorInfoResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostMetricRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDedicatedHostMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostMetricRequest) SetOwnerId(v int64) *DescribeDedicatedHostMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostMetricRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostMetricRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostMetricRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostMetricRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostMetricRequest) SetRegionId(v string) *DescribeDedicatedHostMetricRequest {
	s.RegionId = &v
	return s
}

type DescribeDedicatedHostMetricResponseBody struct {
	DedicatedHostGroupId *string                                         `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	RequestId            *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Metrics              *DescribeDedicatedHostMetricResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s DescribeDedicatedHostMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostMetricResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostMetricResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBody) SetRequestId(v string) *DescribeDedicatedHostMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBody) SetMetrics(v *DescribeDedicatedHostMetricResponseBodyMetrics) *DescribeDedicatedHostMetricResponseBody {
	s.Metrics = v
	return s
}

type DescribeDedicatedHostMetricResponseBodyMetrics struct {
	Metrics []*DescribeDedicatedHostMetricResponseBodyMetricsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostMetricResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostMetricResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostMetricResponseBodyMetrics) SetMetrics(v []*DescribeDedicatedHostMetricResponseBodyMetricsMetrics) *DescribeDedicatedHostMetricResponseBodyMetrics {
	s.Metrics = v
	return s
}

type DescribeDedicatedHostMetricResponseBodyMetricsMetrics struct {
	AvgMemUsage     *float32                                                    `json:"AvgMemUsage,omitempty" xml:"AvgMemUsage,omitempty"`
	EndDate         *string                                                     `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	AvgIops         *int32                                                      `json:"AvgIops,omitempty" xml:"AvgIops,omitempty"`
	MaxMemUsage     *float32                                                    `json:"MaxMemUsage,omitempty" xml:"MaxMemUsage,omitempty"`
	StartDate       *string                                                     `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	DedicatedHostId *string                                                     `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	MaxIops         *int32                                                      `json:"MaxIops,omitempty" xml:"MaxIops,omitempty"`
	AvgCpuUsage     *float32                                                    `json:"AvgCpuUsage,omitempty" xml:"AvgCpuUsage,omitempty"`
	DiskUsage       *float32                                                    `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	MaxCpuUsage     *float32                                                    `json:"MaxCpuUsage,omitempty" xml:"MaxCpuUsage,omitempty"`
	Risks           *DescribeDedicatedHostMetricResponseBodyMetricsMetricsRisks `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Struct"`
}

func (s DescribeDedicatedHostMetricResponseBodyMetricsMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostMetricResponseBodyMetricsMetrics) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetAvgMemUsage(v float32) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.AvgMemUsage = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetEndDate(v string) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.EndDate = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetAvgIops(v int32) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.AvgIops = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetMaxMemUsage(v float32) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.MaxMemUsage = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetStartDate(v string) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.StartDate = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetDedicatedHostId(v string) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetMaxIops(v int32) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.MaxIops = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetAvgCpuUsage(v float32) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.AvgCpuUsage = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetDiskUsage(v float32) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.DiskUsage = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetMaxCpuUsage(v float32) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.MaxCpuUsage = &v
	return s
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetrics) SetRisks(v *DescribeDedicatedHostMetricResponseBodyMetricsMetricsRisks) *DescribeDedicatedHostMetricResponseBodyMetricsMetrics {
	s.Risks = v
	return s
}

type DescribeDedicatedHostMetricResponseBodyMetricsMetricsRisks struct {
	Risks []*string `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostMetricResponseBodyMetricsMetricsRisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostMetricResponseBodyMetricsMetricsRisks) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostMetricResponseBodyMetricsMetricsRisks) SetRisks(v []*string) *DescribeDedicatedHostMetricResponseBodyMetricsMetricsRisks {
	s.Risks = v
	return s
}

type DescribeDedicatedHostMetricResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedHostMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedHostMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostMetricResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostMetricResponse) SetBody(v *DescribeDedicatedHostMetricResponseBody) *DescribeDedicatedHostMetricResponse {
	s.Body = v
	return s
}

type CreateDedicatedHostRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	HostName             *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	HostClass            *string `json:"HostClass,omitempty" xml:"HostClass,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime             *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AutoRenew            *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ImageCategory        *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
}

func (s CreateDedicatedHostRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostRequest) SetOwnerId(v int64) *CreateDedicatedHostRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetResourceOwnerAccount(v string) *CreateDedicatedHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetResourceOwnerId(v int64) *CreateDedicatedHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetRegionId(v string) *CreateDedicatedHostRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetDedicatedHostGroupId(v string) *CreateDedicatedHostRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetHostName(v string) *CreateDedicatedHostRequest {
	s.HostName = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetZoneId(v string) *CreateDedicatedHostRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetVSwitchId(v string) *CreateDedicatedHostRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetHostClass(v string) *CreateDedicatedHostRequest {
	s.HostClass = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetPayType(v string) *CreateDedicatedHostRequest {
	s.PayType = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetPeriod(v string) *CreateDedicatedHostRequest {
	s.Period = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetUsedTime(v string) *CreateDedicatedHostRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetClientToken(v string) *CreateDedicatedHostRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetAutoRenew(v string) *CreateDedicatedHostRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetImageCategory(v string) *CreateDedicatedHostRequest {
	s.ImageCategory = &v
	return s
}

type CreateDedicatedHostResponseBody struct {
	OrderId          *int64                                           `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DedicateHostList *CreateDedicatedHostResponseBodyDedicateHostList `json:"DedicateHostList,omitempty" xml:"DedicateHostList,omitempty" type:"Struct"`
}

func (s CreateDedicatedHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostResponseBody) SetOrderId(v int64) *CreateDedicatedHostResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDedicatedHostResponseBody) SetRequestId(v string) *CreateDedicatedHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDedicatedHostResponseBody) SetDedicateHostList(v *CreateDedicatedHostResponseBodyDedicateHostList) *CreateDedicatedHostResponseBody {
	s.DedicateHostList = v
	return s
}

type CreateDedicatedHostResponseBodyDedicateHostList struct {
	DedicateHostList []*CreateDedicatedHostResponseBodyDedicateHostListDedicateHostList `json:"DedicateHostList,omitempty" xml:"DedicateHostList,omitempty" type:"Repeated"`
}

func (s CreateDedicatedHostResponseBodyDedicateHostList) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostResponseBodyDedicateHostList) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostResponseBodyDedicateHostList) SetDedicateHostList(v []*CreateDedicatedHostResponseBodyDedicateHostListDedicateHostList) *CreateDedicatedHostResponseBodyDedicateHostList {
	s.DedicateHostList = v
	return s
}

type CreateDedicatedHostResponseBodyDedicateHostListDedicateHostList struct {
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
}

func (s CreateDedicatedHostResponseBodyDedicateHostListDedicateHostList) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostResponseBodyDedicateHostListDedicateHostList) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostResponseBodyDedicateHostListDedicateHostList) SetDedicatedHostId(v string) *CreateDedicatedHostResponseBodyDedicateHostListDedicateHostList {
	s.DedicatedHostId = &v
	return s
}

type CreateDedicatedHostResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDedicatedHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDedicatedHostResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostResponse) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostResponse) SetHeaders(v map[string]*string) *CreateDedicatedHostResponse {
	s.Headers = v
	return s
}

func (s *CreateDedicatedHostResponse) SetBody(v *CreateDedicatedHostResponseBody) *CreateDedicatedHostResponse {
	s.Body = v
	return s
}

type DescribeDedicatedInstanceMetricRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDedicatedInstanceMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedInstanceMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedInstanceMetricRequest) SetOwnerId(v int64) *DescribeDedicatedInstanceMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedInstanceMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricRequest) SetResourceOwnerId(v int64) *DescribeDedicatedInstanceMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricRequest) SetDedicatedHostId(v string) *DescribeDedicatedInstanceMetricRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricRequest) SetRegionId(v string) *DescribeDedicatedInstanceMetricRequest {
	s.RegionId = &v
	return s
}

type DescribeDedicatedInstanceMetricResponseBody struct {
	DedicatedHostId *string                                             `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Metrics         *DescribeDedicatedInstanceMetricResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
}

func (s DescribeDedicatedInstanceMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedInstanceMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedInstanceMetricResponseBody) SetDedicatedHostId(v string) *DescribeDedicatedInstanceMetricResponseBody {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBody) SetRequestId(v string) *DescribeDedicatedInstanceMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBody) SetMetrics(v *DescribeDedicatedInstanceMetricResponseBodyMetrics) *DescribeDedicatedInstanceMetricResponseBody {
	s.Metrics = v
	return s
}

type DescribeDedicatedInstanceMetricResponseBodyMetrics struct {
	Metrics []*DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedInstanceMetricResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedInstanceMetricResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetrics) SetMetrics(v []*DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) *DescribeDedicatedInstanceMetricResponseBodyMetrics {
	s.Metrics = v
	return s
}

type DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics struct {
	EndDate             *string                                                         `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	MaxMemUsage         *float32                                                        `json:"MaxMemUsage,omitempty" xml:"MaxMemUsage,omitempty"`
	DedicatedInstanceId *string                                                         `json:"DedicatedInstanceId,omitempty" xml:"DedicatedInstanceId,omitempty"`
	StartDate           *string                                                         `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	DiskUsage           *float32                                                        `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	AvgMemUsage         *float32                                                        `json:"AvgMemUsage,omitempty" xml:"AvgMemUsage,omitempty"`
	MaxIOPS             *int32                                                          `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	Memory              *float32                                                        `json:"Memory,omitempty" xml:"Memory,omitempty"`
	AvgCpuUsage         *float32                                                        `json:"AvgCpuUsage,omitempty" xml:"AvgCpuUsage,omitempty"`
	AvgIOPS             *int32                                                          `json:"AvgIOPS,omitempty" xml:"AvgIOPS,omitempty"`
	Role                *string                                                         `json:"Role,omitempty" xml:"Role,omitempty"`
	MaxCpuUsage         *float32                                                        `json:"MaxCpuUsage,omitempty" xml:"MaxCpuUsage,omitempty"`
	Risks               *DescribeDedicatedInstanceMetricResponseBodyMetricsMetricsRisks `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Struct"`
}

func (s DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetEndDate(v string) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.EndDate = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetMaxMemUsage(v float32) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.MaxMemUsage = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetDedicatedInstanceId(v string) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.DedicatedInstanceId = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetStartDate(v string) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.StartDate = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetDiskUsage(v float32) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.DiskUsage = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetAvgMemUsage(v float32) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.AvgMemUsage = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetMaxIOPS(v int32) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetMemory(v float32) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.Memory = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetAvgCpuUsage(v float32) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.AvgCpuUsage = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetAvgIOPS(v int32) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.AvgIOPS = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetRole(v string) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.Role = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetMaxCpuUsage(v float32) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.MaxCpuUsage = &v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics) SetRisks(v *DescribeDedicatedInstanceMetricResponseBodyMetricsMetricsRisks) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetrics {
	s.Risks = v
	return s
}

type DescribeDedicatedInstanceMetricResponseBodyMetricsMetricsRisks struct {
	Risks []*string `json:"Risks,omitempty" xml:"Risks,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedInstanceMetricResponseBodyMetricsMetricsRisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedInstanceMetricResponseBodyMetricsMetricsRisks) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedInstanceMetricResponseBodyMetricsMetricsRisks) SetRisks(v []*string) *DescribeDedicatedInstanceMetricResponseBodyMetricsMetricsRisks {
	s.Risks = v
	return s
}

type DescribeDedicatedInstanceMetricResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedInstanceMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedInstanceMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedInstanceMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedInstanceMetricResponse) SetHeaders(v map[string]*string) *DescribeDedicatedInstanceMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedInstanceMetricResponse) SetBody(v *DescribeDedicatedInstanceMetricResponseBody) *DescribeDedicatedInstanceMetricResponse {
	s.Body = v
	return s
}

type DescribeCrossVpcInstanceRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeCrossVpcInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrossVpcInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossVpcInstanceRequest) SetOwnerId(v int64) *DescribeCrossVpcInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCrossVpcInstanceRequest) SetResourceOwnerAccount(v string) *DescribeCrossVpcInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCrossVpcInstanceRequest) SetResourceOwnerId(v int64) *DescribeCrossVpcInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCrossVpcInstanceRequest) SetRegionId(v string) *DescribeCrossVpcInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCrossVpcInstanceRequest) SetDBInstanceId(v string) *DescribeCrossVpcInstanceRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeCrossVpcInstanceResponseBody struct {
	HasDelInProcessTask    *bool                                        `json:"HasDelInProcessTask,omitempty" xml:"HasDelInProcessTask,omitempty"`
	RequestId              *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HasCreateInProcessTask *bool                                        `json:"HasCreateInProcessTask,omitempty" xml:"HasCreateInProcessTask,omitempty"`
	Items                  []*DescribeCrossVpcInstanceResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeCrossVpcInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrossVpcInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrossVpcInstanceResponseBody) SetHasDelInProcessTask(v bool) *DescribeCrossVpcInstanceResponseBody {
	s.HasDelInProcessTask = &v
	return s
}

func (s *DescribeCrossVpcInstanceResponseBody) SetRequestId(v string) *DescribeCrossVpcInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrossVpcInstanceResponseBody) SetHasCreateInProcessTask(v bool) *DescribeCrossVpcInstanceResponseBody {
	s.HasCreateInProcessTask = &v
	return s
}

func (s *DescribeCrossVpcInstanceResponseBody) SetItems(v []*DescribeCrossVpcInstanceResponseBodyItems) *DescribeCrossVpcInstanceResponseBody {
	s.Items = v
	return s
}

type DescribeCrossVpcInstanceResponseBodyItems struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
}

func (s DescribeCrossVpcInstanceResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrossVpcInstanceResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCrossVpcInstanceResponseBodyItems) SetConnectionString(v string) *DescribeCrossVpcInstanceResponseBodyItems {
	s.ConnectionString = &v
	return s
}

type DescribeCrossVpcInstanceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCrossVpcInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCrossVpcInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrossVpcInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrossVpcInstanceResponse) SetHeaders(v map[string]*string) *DescribeCrossVpcInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrossVpcInstanceResponse) SetBody(v *DescribeCrossVpcInstanceResponseBody) *DescribeCrossVpcInstanceResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostClassRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	TargetClassCode      *string `json:"TargetClassCode,omitempty" xml:"TargetClassCode,omitempty"`
	SwitchTime           *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode       *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s ModifyDedicatedHostClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostClassRequest) SetOwnerId(v int64) *ModifyDedicatedHostClassRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostClassRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostClassRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetRegionId(v string) *ModifyDedicatedHostClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetDedicatedHostId(v string) *ModifyDedicatedHostClassRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetTargetClassCode(v string) *ModifyDedicatedHostClassRequest {
	s.TargetClassCode = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetSwitchTime(v string) *ModifyDedicatedHostClassRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetSwitchTimeMode(v string) *ModifyDedicatedHostClassRequest {
	s.SwitchTimeMode = &v
	return s
}

type ModifyDedicatedHostClassResponseBody struct {
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDedicatedHostClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostClassResponseBody) SetDedicatedHostId(v string) *ModifyDedicatedHostClassResponseBody {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyDedicatedHostClassResponseBody) SetRequestId(v string) *ModifyDedicatedHostClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDedicatedHostClassResponseBody) SetTaskId(v string) *ModifyDedicatedHostClassResponseBody {
	s.TaskId = &v
	return s
}

type ModifyDedicatedHostClassResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDedicatedHostClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDedicatedHostClassResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostClassResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostClassResponse) SetBody(v *ModifyDedicatedHostClassResponseBody) *ModifyDedicatedHostClassResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostsCheckInBastionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	BastionInstanceId    *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
}

func (s DescribeDedicatedHostsCheckInBastionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsCheckInBastionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsCheckInBastionRequest) SetOwnerId(v int64) *DescribeDedicatedHostsCheckInBastionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostsCheckInBastionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostsCheckInBastionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionRequest) SetRegionId(v string) *DescribeDedicatedHostsCheckInBastionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsCheckInBastionRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionRequest) SetBastionInstanceId(v string) *DescribeDedicatedHostsCheckInBastionRequest {
	s.BastionInstanceId = &v
	return s
}

type DescribeDedicatedHostsCheckInBastionResponseBody struct {
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BastionInstanceId    *string                                                  `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	DedicatedHostGroupId *string                                                  `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	Hosts                []*DescribeDedicatedHostsCheckInBastionResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsCheckInBastionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsCheckInBastionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBody) SetRequestId(v string) *DescribeDedicatedHostsCheckInBastionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBody) SetBastionInstanceId(v string) *DescribeDedicatedHostsCheckInBastionResponseBody {
	s.BastionInstanceId = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsCheckInBastionResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBody) SetHosts(v []*DescribeDedicatedHostsCheckInBastionResponseBodyHosts) *DescribeDedicatedHostsCheckInBastionResponseBody {
	s.Hosts = v
	return s
}

type DescribeDedicatedHostsCheckInBastionResponseBodyHosts struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	AllocationStatus  *bool   `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	InBastion         *bool   `json:"InBastion,omitempty" xml:"InBastion,omitempty"`
	HostName          *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Ip                *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDedicatedHostsCheckInBastionResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsCheckInBastionResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBodyHosts) SetStatus(v string) *DescribeDedicatedHostsCheckInBastionResponseBodyHosts {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBodyHosts) SetDedicatedHostName(v string) *DescribeDedicatedHostsCheckInBastionResponseBodyHosts {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBodyHosts) SetAllocationStatus(v bool) *DescribeDedicatedHostsCheckInBastionResponseBodyHosts {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBodyHosts) SetInBastion(v bool) *DescribeDedicatedHostsCheckInBastionResponseBodyHosts {
	s.InBastion = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBodyHosts) SetHostName(v string) *DescribeDedicatedHostsCheckInBastionResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionResponseBodyHosts) SetIp(v string) *DescribeDedicatedHostsCheckInBastionResponseBodyHosts {
	s.Ip = &v
	return s
}

type DescribeDedicatedHostsCheckInBastionResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedHostsCheckInBastionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedHostsCheckInBastionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsCheckInBastionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsCheckInBastionResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostsCheckInBastionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostsCheckInBastionResponse) SetBody(v *DescribeDedicatedHostsCheckInBastionResponseBody) *DescribeDedicatedHostsCheckInBastionResponse {
	s.Body = v
	return s
}

type DropDedicatedHostUserRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostName    *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	UserName             *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DropDedicatedHostUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DropDedicatedHostUserRequest) GoString() string {
	return s.String()
}

func (s *DropDedicatedHostUserRequest) SetOwnerId(v int64) *DropDedicatedHostUserRequest {
	s.OwnerId = &v
	return s
}

func (s *DropDedicatedHostUserRequest) SetResourceOwnerAccount(v string) *DropDedicatedHostUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DropDedicatedHostUserRequest) SetResourceOwnerId(v int64) *DropDedicatedHostUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DropDedicatedHostUserRequest) SetDedicatedHostName(v string) *DropDedicatedHostUserRequest {
	s.DedicatedHostName = &v
	return s
}

func (s *DropDedicatedHostUserRequest) SetUserName(v string) *DropDedicatedHostUserRequest {
	s.UserName = &v
	return s
}

func (s *DropDedicatedHostUserRequest) SetRegionId(v string) *DropDedicatedHostUserRequest {
	s.RegionId = &v
	return s
}

type DropDedicatedHostUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DropDedicatedHostUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DropDedicatedHostUserResponseBody) GoString() string {
	return s.String()
}

func (s *DropDedicatedHostUserResponseBody) SetRequestId(v string) *DropDedicatedHostUserResponseBody {
	s.RequestId = &v
	return s
}

type DropDedicatedHostUserResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DropDedicatedHostUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DropDedicatedHostUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DropDedicatedHostUserResponse) GoString() string {
	return s.String()
}

func (s *DropDedicatedHostUserResponse) SetHeaders(v map[string]*string) *DropDedicatedHostUserResponse {
	s.Headers = v
	return s
}

func (s *DropDedicatedHostUserResponse) SetBody(v *DropDedicatedHostUserResponseBody) *DropDedicatedHostUserResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostsInBastionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	BastionInstanceId    *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
}

func (s DescribeDedicatedHostsInBastionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsInBastionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsInBastionRequest) SetOwnerId(v int64) *DescribeDedicatedHostsInBastionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostsInBastionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostsInBastionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionRequest) SetRegionId(v string) *DescribeDedicatedHostsInBastionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsInBastionRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionRequest) SetBastionInstanceId(v string) *DescribeDedicatedHostsInBastionRequest {
	s.BastionInstanceId = &v
	return s
}

type DescribeDedicatedHostsInBastionResponseBody struct {
	RequestId            *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BastionInstanceId    *string                                             `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	DedicatedHostGroupId *string                                             `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	Hosts                []*DescribeDedicatedHostsInBastionResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsInBastionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsInBastionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsInBastionResponseBody) SetRequestId(v string) *DescribeDedicatedHostsInBastionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBody) SetBastionInstanceId(v string) *DescribeDedicatedHostsInBastionResponseBody {
	s.BastionInstanceId = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsInBastionResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBody) SetHosts(v []*DescribeDedicatedHostsInBastionResponseBodyHosts) *DescribeDedicatedHostsInBastionResponseBody {
	s.Hosts = v
	return s
}

type DescribeDedicatedHostsInBastionResponseBodyHosts struct {
	HostDescription  *string `json:"HostDescription,omitempty" xml:"HostDescription,omitempty"`
	Comment          *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	BastionHostId    *string `json:"BastionHostId,omitempty" xml:"BastionHostId,omitempty"`
	OsName           *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	HostAccountCount *string `json:"HostAccountCount,omitempty" xml:"HostAccountCount,omitempty"`
	DedicatedHostId  *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	HostPrivateIp    *string `json:"HostPrivateIp,omitempty" xml:"HostPrivateIp,omitempty"`
	AccountCreating  *bool   `json:"AccountCreating,omitempty" xml:"AccountCreating,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeDedicatedHostsInBastionResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsInBastionResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsInBastionResponseBodyHosts) SetHostDescription(v string) *DescribeDedicatedHostsInBastionResponseBodyHosts {
	s.HostDescription = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBodyHosts) SetComment(v string) *DescribeDedicatedHostsInBastionResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBodyHosts) SetBastionHostId(v string) *DescribeDedicatedHostsInBastionResponseBodyHosts {
	s.BastionHostId = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBodyHosts) SetOsName(v string) *DescribeDedicatedHostsInBastionResponseBodyHosts {
	s.OsName = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBodyHosts) SetHostAccountCount(v string) *DescribeDedicatedHostsInBastionResponseBodyHosts {
	s.HostAccountCount = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBodyHosts) SetDedicatedHostId(v string) *DescribeDedicatedHostsInBastionResponseBodyHosts {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBodyHosts) SetHostPrivateIp(v string) *DescribeDedicatedHostsInBastionResponseBodyHosts {
	s.HostPrivateIp = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBodyHosts) SetAccountCreating(v bool) *DescribeDedicatedHostsInBastionResponseBodyHosts {
	s.AccountCreating = &v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponseBodyHosts) SetAccountName(v string) *DescribeDedicatedHostsInBastionResponseBodyHosts {
	s.AccountName = &v
	return s
}

type DescribeDedicatedHostsInBastionResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedHostsInBastionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedHostsInBastionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsInBastionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsInBastionResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostsInBastionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostsInBastionResponse) SetBody(v *DescribeDedicatedHostsInBastionResponseBody) *DescribeDedicatedHostsInBastionResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostGroupAttributeRequest struct {
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId   *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHostGroupDesc *string `json:"DedicatedHostGroupDesc,omitempty" xml:"DedicatedHostGroupDesc,omitempty"`
	CpuAllocationRatio     *int32  `json:"CpuAllocationRatio,omitempty" xml:"CpuAllocationRatio,omitempty"`
	MemAllocationRatio     *int32  `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	DiskAllocationRatio    *int32  `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	AllocationPolicy       *string `json:"AllocationPolicy,omitempty" xml:"AllocationPolicy,omitempty"`
	HostReplacePolicy      *string `json:"HostReplacePolicy,omitempty" xml:"HostReplacePolicy,omitempty"`
	OpenPermission         *string `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
}

func (s ModifyDedicatedHostGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetOwnerId(v int64) *ModifyDedicatedHostGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetRegionId(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetDedicatedHostGroupId(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetDedicatedHostGroupDesc(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.DedicatedHostGroupDesc = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetCpuAllocationRatio(v int32) *ModifyDedicatedHostGroupAttributeRequest {
	s.CpuAllocationRatio = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetMemAllocationRatio(v int32) *ModifyDedicatedHostGroupAttributeRequest {
	s.MemAllocationRatio = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetDiskAllocationRatio(v int32) *ModifyDedicatedHostGroupAttributeRequest {
	s.DiskAllocationRatio = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetAllocationPolicy(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.AllocationPolicy = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetHostReplacePolicy(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.HostReplacePolicy = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetOpenPermission(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.OpenPermission = &v
	return s
}

type ModifyDedicatedHostGroupAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedHostGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostGroupAttributeResponseBody) SetRequestId(v string) *ModifyDedicatedHostGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDedicatedHostGroupAttributeResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDedicatedHostGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDedicatedHostGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostGroupAttributeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeResponse) SetBody(v *ModifyDedicatedHostGroupAttributeResponseBody) *ModifyDedicatedHostGroupAttributeResponse {
	s.Body = v
	return s
}

type QueryHostInstanceConsoleInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
}

func (s QueryHostInstanceConsoleInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHostInstanceConsoleInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHostInstanceConsoleInfoRequest) SetOwnerId(v int64) *QueryHostInstanceConsoleInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoRequest) SetResourceOwnerAccount(v string) *QueryHostInstanceConsoleInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoRequest) SetResourceOwnerId(v int64) *QueryHostInstanceConsoleInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoRequest) SetRegionId(v string) *QueryHostInstanceConsoleInfoRequest {
	s.RegionId = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoRequest) SetDedicatedHostId(v string) *QueryHostInstanceConsoleInfoRequest {
	s.DedicatedHostId = &v
	return s
}

type QueryHostInstanceConsoleInfoResponseBody struct {
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostInstanceConsoleInfos []*QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos `json:"HostInstanceConsoleInfos,omitempty" xml:"HostInstanceConsoleInfos,omitempty" type:"Repeated"`
}

func (s QueryHostInstanceConsoleInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHostInstanceConsoleInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHostInstanceConsoleInfoResponseBody) SetRequestId(v string) *QueryHostInstanceConsoleInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBody) SetHostInstanceConsoleInfos(v []*QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) *QueryHostInstanceConsoleInfoResponseBody {
	s.HostInstanceConsoleInfos = v
	return s
}

type QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos struct {
	Status                    *string                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	MaxConnIncreaseRatioValue *int32                                                                    `json:"MaxConnIncreaseRatioValue,omitempty" xml:"MaxConnIncreaseRatioValue,omitempty"`
	MemSize                   *int32                                                                    `json:"MemSize,omitempty" xml:"MemSize,omitempty"`
	DiskSize                  *int32                                                                    `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	Ip                        *string                                                                   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port                      *string                                                                   `json:"Port,omitempty" xml:"Port,omitempty"`
	EngineVersion             *string                                                                   `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	MemoryIncreaseRatioValue  *int32                                                                    `json:"MemoryIncreaseRatioValue,omitempty" xml:"MemoryIncreaseRatioValue,omitempty"`
	CpuIncreaseRatioValue     *int32                                                                    `json:"CpuIncreaseRatioValue,omitempty" xml:"CpuIncreaseRatioValue,omitempty"`
	DBInstanceId              *string                                                                   `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Engine                    *string                                                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
	LevelName                 *string                                                                   `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	Role                      *string                                                                   `json:"Role,omitempty" xml:"Role,omitempty"`
	DBInstanceDescription     *string                                                                   `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	CpuCores                  *int32                                                                    `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	PerfInfo                  *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo `json:"PerfInfo,omitempty" xml:"PerfInfo,omitempty" type:"Struct"`
}

func (s QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) GoString() string {
	return s.String()
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetStatus(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Status = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetMaxConnIncreaseRatioValue(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.MaxConnIncreaseRatioValue = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetMemSize(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.MemSize = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetDiskSize(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.DiskSize = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetIp(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Ip = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetPort(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Port = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetEngineVersion(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.EngineVersion = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetMemoryIncreaseRatioValue(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.MemoryIncreaseRatioValue = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetCpuIncreaseRatioValue(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.CpuIncreaseRatioValue = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetDBInstanceId(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.DBInstanceId = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetEngine(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Engine = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetLevelName(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.LevelName = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetRole(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Role = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetDBInstanceDescription(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.DBInstanceDescription = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetCpuCores(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.CpuCores = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetPerfInfo(v *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.PerfInfo = v
	return s
}

type QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo struct {
	PerfIdbPio *float32 `json:"PerfIdbPio,omitempty" xml:"PerfIdbPio,omitempty"`
	DiskCurr   *float32 `json:"DiskCurr,omitempty" xml:"DiskCurr,omitempty"`
	MemRatio   *float32 `json:"MemRatio,omitempty" xml:"MemRatio,omitempty"`
	CpuRatio   *float32 `json:"CpuRatio,omitempty" xml:"CpuRatio,omitempty"`
}

func (s QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) GoString() string {
	return s.String()
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) SetPerfIdbPio(v float32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo {
	s.PerfIdbPio = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) SetDiskCurr(v float32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo {
	s.DiskCurr = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) SetMemRatio(v float32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo {
	s.MemRatio = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) SetCpuRatio(v float32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo {
	s.CpuRatio = &v
	return s
}

type QueryHostInstanceConsoleInfoResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHostInstanceConsoleInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHostInstanceConsoleInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHostInstanceConsoleInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryHostInstanceConsoleInfoResponse) SetHeaders(v map[string]*string) *QueryHostInstanceConsoleInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponse) SetBody(v *QueryHostInstanceConsoleInfoResponseBody) *QueryHostInstanceConsoleInfoResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
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

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RDSRegion []*DescribeRegionsResponseBodyRegionsRDSRegion `json:"RDSRegion,omitempty" xml:"RDSRegion,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRDSRegion(v []*DescribeRegionsResponseBodyRegionsRDSRegion) *DescribeRegionsResponseBodyRegions {
	s.RDSRegion = v
	return s
}

type DescribeRegionsResponseBodyRegionsRDSRegion struct {
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRDSRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRDSRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRDSRegion {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRDSRegion {
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

type QueryHostBaseInfoByInstanceRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s QueryHostBaseInfoByInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHostBaseInfoByInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryHostBaseInfoByInstanceRequest) SetOwnerId(v int64) *QueryHostBaseInfoByInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceRequest) SetResourceOwnerAccount(v string) *QueryHostBaseInfoByInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceRequest) SetResourceOwnerId(v int64) *QueryHostBaseInfoByInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceRequest) SetRegionId(v string) *QueryHostBaseInfoByInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceRequest) SetDBInstanceId(v string) *QueryHostBaseInfoByInstanceRequest {
	s.DBInstanceId = &v
	return s
}

type QueryHostBaseInfoByInstanceResponseBody struct {
	RequestId                *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostInstanceConsoleInfos []*QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos `json:"HostInstanceConsoleInfos,omitempty" xml:"HostInstanceConsoleInfos,omitempty" type:"Repeated"`
}

func (s QueryHostBaseInfoByInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHostBaseInfoByInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHostBaseInfoByInstanceResponseBody) SetRequestId(v string) *QueryHostBaseInfoByInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBody) SetHostInstanceConsoleInfos(v []*QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) *QueryHostBaseInfoByInstanceResponseBody {
	s.HostInstanceConsoleInfos = v
	return s
}

type QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos struct {
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ExpiredTime   *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ClusterName   *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	HostName      *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Role          *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Port          *string `json:"Port,omitempty" xml:"Port,omitempty"`
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
}

func (s QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) GoString() string {
	return s.String()
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetVpcId(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.VpcId = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetStatus(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Status = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetExpiredTime(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.ExpiredTime = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetClusterName(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.ClusterName = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetIp(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Ip = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetHostName(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.HostName = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetEngine(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Engine = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetRole(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Role = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetPort(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Port = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetEngineVersion(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.EngineVersion = &v
	return s
}

type QueryHostBaseInfoByInstanceResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHostBaseInfoByInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHostBaseInfoByInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHostBaseInfoByInstanceResponse) GoString() string {
	return s.String()
}

func (s *QueryHostBaseInfoByInstanceResponse) SetHeaders(v map[string]*string) *QueryHostBaseInfoByInstanceResponse {
	s.Headers = v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponse) SetBody(v *QueryHostBaseInfoByInstanceResponseBody) *QueryHostBaseInfoByInstanceResponse {
	s.Body = v
	return s
}

type DescribeDedicatedInstanceDistributionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDedicatedInstanceDistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedInstanceDistributionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedInstanceDistributionRequest) SetOwnerId(v int64) *DescribeDedicatedInstanceDistributionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedInstanceDistributionRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedInstanceDistributionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedInstanceDistributionRequest) SetResourceOwnerId(v int64) *DescribeDedicatedInstanceDistributionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedInstanceDistributionRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedInstanceDistributionRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedInstanceDistributionRequest) SetRegionId(v string) *DescribeDedicatedInstanceDistributionRequest {
	s.RegionId = &v
	return s
}

type DescribeDedicatedInstanceDistributionResponseBody struct {
	DBClass              map[string]interface{} `json:"DBClass,omitempty" xml:"DBClass,omitempty"`
	DBVersion            map[string]interface{} `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	RequestId            *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceCount        *int32                 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	DBType               map[string]interface{} `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DedicatedHostGroupId *string                `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
}

func (s DescribeDedicatedInstanceDistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedInstanceDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedInstanceDistributionResponseBody) SetDBClass(v map[string]interface{}) *DescribeDedicatedInstanceDistributionResponseBody {
	s.DBClass = v
	return s
}

func (s *DescribeDedicatedInstanceDistributionResponseBody) SetDBVersion(v map[string]interface{}) *DescribeDedicatedInstanceDistributionResponseBody {
	s.DBVersion = v
	return s
}

func (s *DescribeDedicatedInstanceDistributionResponseBody) SetRequestId(v string) *DescribeDedicatedInstanceDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedInstanceDistributionResponseBody) SetInstanceCount(v int32) *DescribeDedicatedInstanceDistributionResponseBody {
	s.InstanceCount = &v
	return s
}

func (s *DescribeDedicatedInstanceDistributionResponseBody) SetDBType(v map[string]interface{}) *DescribeDedicatedInstanceDistributionResponseBody {
	s.DBType = v
	return s
}

func (s *DescribeDedicatedInstanceDistributionResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedInstanceDistributionResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

type DescribeDedicatedInstanceDistributionResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedInstanceDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedInstanceDistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedInstanceDistributionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedInstanceDistributionResponse) SetHeaders(v map[string]*string) *DescribeDedicatedInstanceDistributionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedInstanceDistributionResponse) SetBody(v *DescribeDedicatedInstanceDistributionResponseBody) *DescribeDedicatedInstanceDistributionResponse {
	s.Body = v
	return s
}

type DescribeScheduleHostRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ScheduleId           *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
}

func (s DescribeScheduleHostRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleHostRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduleHostRequest) SetOwnerId(v int64) *DescribeScheduleHostRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScheduleHostRequest) SetResourceOwnerAccount(v string) *DescribeScheduleHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScheduleHostRequest) SetResourceOwnerId(v int64) *DescribeScheduleHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScheduleHostRequest) SetRegionId(v string) *DescribeScheduleHostRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScheduleHostRequest) SetDedicatedHostGroupId(v string) *DescribeScheduleHostRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeScheduleHostRequest) SetScheduleId(v string) *DescribeScheduleHostRequest {
	s.ScheduleId = &v
	return s
}

type DescribeScheduleHostResponseBody struct {
	ScheduleId             *string                                                   `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	RequestId              *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostScheduleStatusList []*DescribeScheduleHostResponseBodyHostScheduleStatusList `json:"HostScheduleStatusList,omitempty" xml:"HostScheduleStatusList,omitempty" type:"Repeated"`
}

func (s DescribeScheduleHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleHostResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduleHostResponseBody) SetScheduleId(v string) *DescribeScheduleHostResponseBody {
	s.ScheduleId = &v
	return s
}

func (s *DescribeScheduleHostResponseBody) SetRequestId(v string) *DescribeScheduleHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduleHostResponseBody) SetHostScheduleStatusList(v []*DescribeScheduleHostResponseBodyHostScheduleStatusList) *DescribeScheduleHostResponseBody {
	s.HostScheduleStatusList = v
	return s
}

type DescribeScheduleHostResponseBodyHostScheduleStatusList struct {
	ScheduleStatus  *string `json:"ScheduleStatus,omitempty" xml:"ScheduleStatus,omitempty"`
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	InstanceOutId   *string `json:"InstanceOutId,omitempty" xml:"InstanceOutId,omitempty"`
	InstanceInId    *string `json:"InstanceInId,omitempty" xml:"InstanceInId,omitempty"`
}

func (s DescribeScheduleHostResponseBodyHostScheduleStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleHostResponseBodyHostScheduleStatusList) GoString() string {
	return s.String()
}

func (s *DescribeScheduleHostResponseBodyHostScheduleStatusList) SetScheduleStatus(v string) *DescribeScheduleHostResponseBodyHostScheduleStatusList {
	s.ScheduleStatus = &v
	return s
}

func (s *DescribeScheduleHostResponseBodyHostScheduleStatusList) SetDedicatedHostId(v string) *DescribeScheduleHostResponseBodyHostScheduleStatusList {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeScheduleHostResponseBodyHostScheduleStatusList) SetInstanceOutId(v string) *DescribeScheduleHostResponseBodyHostScheduleStatusList {
	s.InstanceOutId = &v
	return s
}

func (s *DescribeScheduleHostResponseBodyHostScheduleStatusList) SetInstanceInId(v string) *DescribeScheduleHostResponseBodyHostScheduleStatusList {
	s.InstanceInId = &v
	return s
}

type DescribeScheduleHostResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScheduleHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScheduleHostResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleHostResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduleHostResponse) SetHeaders(v map[string]*string) *DescribeScheduleHostResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduleHostResponse) SetBody(v *DescribeScheduleHostResponseBody) *DescribeScheduleHostResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	HostName             *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	AllocationStatus     *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
}

func (s ModifyDedicatedHostAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAttributeRequest) SetOwnerId(v int64) *ModifyDedicatedHostAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetRegionId(v string) *ModifyDedicatedHostAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetDedicatedHostId(v string) *ModifyDedicatedHostAttributeRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetHostName(v string) *ModifyDedicatedHostAttributeRequest {
	s.HostName = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetAllocationStatus(v string) *ModifyDedicatedHostAttributeRequest {
	s.AllocationStatus = &v
	return s
}

type ModifyDedicatedHostAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedHostAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAttributeResponseBody) SetRequestId(v string) *ModifyDedicatedHostAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDedicatedHostAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDedicatedHostAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDedicatedHostAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAttributeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostAttributeResponse) SetBody(v *ModifyDedicatedHostAttributeResponseBody) *ModifyDedicatedHostAttributeResponse {
	s.Body = v
	return s
}

type CreateDedicatedHostGroupRequest struct {
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Engine                 *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	CpuAllocationRatio     *int32  `json:"CpuAllocationRatio,omitempty" xml:"CpuAllocationRatio,omitempty"`
	MemAllocationRatio     *int32  `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	DiskAllocationRatio    *int32  `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	AllocationPolicy       *string `json:"AllocationPolicy,omitempty" xml:"AllocationPolicy,omitempty"`
	VPCId                  *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	HostReplacePolicy      *string `json:"HostReplacePolicy,omitempty" xml:"HostReplacePolicy,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OpenPermission         *int32  `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	DedicatedHostGroupDesc *string `json:"DedicatedHostGroupDesc,omitempty" xml:"DedicatedHostGroupDesc,omitempty"`
}

func (s CreateDedicatedHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostGroupRequest) SetOwnerId(v int64) *CreateDedicatedHostGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetResourceOwnerAccount(v string) *CreateDedicatedHostGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetResourceOwnerId(v int64) *CreateDedicatedHostGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetRegionId(v string) *CreateDedicatedHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetEngine(v string) *CreateDedicatedHostGroupRequest {
	s.Engine = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetCpuAllocationRatio(v int32) *CreateDedicatedHostGroupRequest {
	s.CpuAllocationRatio = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetMemAllocationRatio(v int32) *CreateDedicatedHostGroupRequest {
	s.MemAllocationRatio = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetDiskAllocationRatio(v int32) *CreateDedicatedHostGroupRequest {
	s.DiskAllocationRatio = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetAllocationPolicy(v string) *CreateDedicatedHostGroupRequest {
	s.AllocationPolicy = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetVPCId(v string) *CreateDedicatedHostGroupRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetHostReplacePolicy(v string) *CreateDedicatedHostGroupRequest {
	s.HostReplacePolicy = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetClientToken(v string) *CreateDedicatedHostGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetOpenPermission(v int32) *CreateDedicatedHostGroupRequest {
	s.OpenPermission = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetDedicatedHostGroupDesc(v string) *CreateDedicatedHostGroupRequest {
	s.DedicatedHostGroupDesc = &v
	return s
}

type CreateDedicatedHostGroupResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
}

func (s CreateDedicatedHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostGroupResponseBody) SetRequestId(v string) *CreateDedicatedHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDedicatedHostGroupResponseBody) SetDedicatedHostGroupId(v string) *CreateDedicatedHostGroupResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

type CreateDedicatedHostGroupResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDedicatedHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDedicatedHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostGroupResponse) SetHeaders(v map[string]*string) *CreateDedicatedHostGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDedicatedHostGroupResponse) SetBody(v *CreateDedicatedHostGroupResponseBody) *CreateDedicatedHostGroupResponse {
	s.Body = v
	return s
}

type AddHostsToBastionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	BastionInstanceId    *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	Hosts                *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
}

func (s AddHostsToBastionRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHostsToBastionRequest) GoString() string {
	return s.String()
}

func (s *AddHostsToBastionRequest) SetOwnerId(v int64) *AddHostsToBastionRequest {
	s.OwnerId = &v
	return s
}

func (s *AddHostsToBastionRequest) SetResourceOwnerAccount(v string) *AddHostsToBastionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddHostsToBastionRequest) SetResourceOwnerId(v int64) *AddHostsToBastionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddHostsToBastionRequest) SetRegionId(v string) *AddHostsToBastionRequest {
	s.RegionId = &v
	return s
}

func (s *AddHostsToBastionRequest) SetDedicatedHostGroupId(v string) *AddHostsToBastionRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *AddHostsToBastionRequest) SetBastionInstanceId(v string) *AddHostsToBastionRequest {
	s.BastionInstanceId = &v
	return s
}

func (s *AddHostsToBastionRequest) SetHosts(v string) *AddHostsToBastionRequest {
	s.Hosts = &v
	return s
}

type AddHostsToBastionResponseBody struct {
	RequestId            *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BastionInstanceId    *string                               `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	DedicatedHostGroupId *string                               `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	Hosts                []*AddHostsToBastionResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s AddHostsToBastionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHostsToBastionResponseBody) GoString() string {
	return s.String()
}

func (s *AddHostsToBastionResponseBody) SetRequestId(v string) *AddHostsToBastionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHostsToBastionResponseBody) SetBastionInstanceId(v string) *AddHostsToBastionResponseBody {
	s.BastionInstanceId = &v
	return s
}

func (s *AddHostsToBastionResponseBody) SetDedicatedHostGroupId(v string) *AddHostsToBastionResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *AddHostsToBastionResponseBody) SetHosts(v []*AddHostsToBastionResponseBodyHosts) *AddHostsToBastionResponseBody {
	s.Hosts = v
	return s
}

type AddHostsToBastionResponseBodyHosts struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostName      *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostPrivateIp *string `json:"HostPrivateIp,omitempty" xml:"HostPrivateIp,omitempty"`
}

func (s AddHostsToBastionResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s AddHostsToBastionResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *AddHostsToBastionResponseBodyHosts) SetCode(v string) *AddHostsToBastionResponseBodyHosts {
	s.Code = &v
	return s
}

func (s *AddHostsToBastionResponseBodyHosts) SetHostName(v string) *AddHostsToBastionResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *AddHostsToBastionResponseBodyHosts) SetMessage(v string) *AddHostsToBastionResponseBodyHosts {
	s.Message = &v
	return s
}

func (s *AddHostsToBastionResponseBodyHosts) SetHostPrivateIp(v string) *AddHostsToBastionResponseBodyHosts {
	s.HostPrivateIp = &v
	return s
}

type AddHostsToBastionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddHostsToBastionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddHostsToBastionResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHostsToBastionResponse) GoString() string {
	return s.String()
}

func (s *AddHostsToBastionResponse) SetHeaders(v map[string]*string) *AddHostsToBastionResponse {
	s.Headers = v
	return s
}

func (s *AddHostsToBastionResponse) SetBody(v *AddHostsToBastionResponseBody) *AddHostsToBastionResponse {
	s.Body = v
	return s
}

type AttachHostsToBastionUserRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BastionUser          *string `json:"BastionUser,omitempty" xml:"BastionUser,omitempty"`
	BastionInstanceId    *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	Hosts                *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
}

func (s AttachHostsToBastionUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostsToBastionUserRequest) GoString() string {
	return s.String()
}

func (s *AttachHostsToBastionUserRequest) SetOwnerId(v int64) *AttachHostsToBastionUserRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachHostsToBastionUserRequest) SetResourceOwnerAccount(v string) *AttachHostsToBastionUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachHostsToBastionUserRequest) SetResourceOwnerId(v int64) *AttachHostsToBastionUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachHostsToBastionUserRequest) SetRegionId(v string) *AttachHostsToBastionUserRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostsToBastionUserRequest) SetBastionUser(v string) *AttachHostsToBastionUserRequest {
	s.BastionUser = &v
	return s
}

func (s *AttachHostsToBastionUserRequest) SetBastionInstanceId(v string) *AttachHostsToBastionUserRequest {
	s.BastionInstanceId = &v
	return s
}

func (s *AttachHostsToBastionUserRequest) SetHosts(v string) *AttachHostsToBastionUserRequest {
	s.Hosts = &v
	return s
}

func (s *AttachHostsToBastionUserRequest) SetDedicatedHostGroupId(v string) *AttachHostsToBastionUserRequest {
	s.DedicatedHostGroupId = &v
	return s
}

type AttachHostsToBastionUserResponseBody struct {
	RequestId         *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BastionInstanceId *string                                      `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	BastionUser       *string                                      `json:"BastionUser,omitempty" xml:"BastionUser,omitempty"`
	Hosts             []*AttachHostsToBastionUserResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s AttachHostsToBastionUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachHostsToBastionUserResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostsToBastionUserResponseBody) SetRequestId(v string) *AttachHostsToBastionUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostsToBastionUserResponseBody) SetBastionInstanceId(v string) *AttachHostsToBastionUserResponseBody {
	s.BastionInstanceId = &v
	return s
}

func (s *AttachHostsToBastionUserResponseBody) SetBastionUser(v string) *AttachHostsToBastionUserResponseBody {
	s.BastionUser = &v
	return s
}

func (s *AttachHostsToBastionUserResponseBody) SetHosts(v []*AttachHostsToBastionUserResponseBodyHosts) *AttachHostsToBastionUserResponseBody {
	s.Hosts = v
	return s
}

type AttachHostsToBastionUserResponseBodyHosts struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AttachHostsToBastionUserResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s AttachHostsToBastionUserResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *AttachHostsToBastionUserResponseBodyHosts) SetCode(v string) *AttachHostsToBastionUserResponseBodyHosts {
	s.Code = &v
	return s
}

func (s *AttachHostsToBastionUserResponseBodyHosts) SetHostName(v string) *AttachHostsToBastionUserResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *AttachHostsToBastionUserResponseBodyHosts) SetMessage(v string) *AttachHostsToBastionUserResponseBodyHosts {
	s.Message = &v
	return s
}

type AttachHostsToBastionUserResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachHostsToBastionUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachHostsToBastionUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachHostsToBastionUserResponse) GoString() string {
	return s.String()
}

func (s *AttachHostsToBastionUserResponse) SetHeaders(v map[string]*string) *AttachHostsToBastionUserResponse {
	s.Headers = v
	return s
}

func (s *AttachHostsToBastionUserResponse) SetBody(v *AttachHostsToBastionUserResponseBody) *AttachHostsToBastionUserResponse {
	s.Body = v
	return s
}

type DescribeEvaluateDedicatedHostsForInstanceRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ExcludeType          *string `json:"ExcludeType,omitempty" xml:"ExcludeType,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
}

func (s DescribeEvaluateDedicatedHostsForInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluateDedicatedHostsForInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateDedicatedHostsForInstanceRequest) SetOwnerId(v int64) *DescribeEvaluateDedicatedHostsForInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceRequest) SetResourceOwnerAccount(v string) *DescribeEvaluateDedicatedHostsForInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceRequest) SetResourceOwnerId(v int64) *DescribeEvaluateDedicatedHostsForInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceRequest) SetRegionId(v string) *DescribeEvaluateDedicatedHostsForInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceRequest) SetDedicatedHostGroupId(v string) *DescribeEvaluateDedicatedHostsForInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceRequest) SetDBInstanceId(v string) *DescribeEvaluateDedicatedHostsForInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceRequest) SetExcludeType(v string) *DescribeEvaluateDedicatedHostsForInstanceRequest {
	s.ExcludeType = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceRequest) SetZoneId(v string) *DescribeEvaluateDedicatedHostsForInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceRequest) SetDedicatedHostId(v string) *DescribeEvaluateDedicatedHostsForInstanceRequest {
	s.DedicatedHostId = &v
	return s
}

type DescribeEvaluateDedicatedHostsForInstanceResponseBody struct {
	RequestId                 *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceRequiredCPU       *int32                                                                 `json:"InstanceRequiredCPU,omitempty" xml:"InstanceRequiredCPU,omitempty"`
	CpuOverAllocationRatio    *int32                                                                 `json:"CpuOverAllocationRatio,omitempty" xml:"CpuOverAllocationRatio,omitempty"`
	MemoryOverAllocationRatio *int32                                                                 `json:"MemoryOverAllocationRatio,omitempty" xml:"MemoryOverAllocationRatio,omitempty"`
	InstanceRequiredMemory    *int32                                                                 `json:"InstanceRequiredMemory,omitempty" xml:"InstanceRequiredMemory,omitempty"`
	InstanceClassCode         *string                                                                `json:"InstanceClassCode,omitempty" xml:"InstanceClassCode,omitempty"`
	AvailableHostNum          *int32                                                                 `json:"AvailableHostNum,omitempty" xml:"AvailableHostNum,omitempty"`
	InstanceRequiredStorage   *int32                                                                 `json:"InstanceRequiredStorage,omitempty" xml:"InstanceRequiredStorage,omitempty"`
	DiskOverAllocationRatio   *int32                                                                 `json:"DiskOverAllocationRatio,omitempty" xml:"DiskOverAllocationRatio,omitempty"`
	DedicatedHostGroupId      *string                                                                `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHosts            []*DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts `json:"DedicatedHosts,omitempty" xml:"DedicatedHosts,omitempty" type:"Repeated"`
}

func (s DescribeEvaluateDedicatedHostsForInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluateDedicatedHostsForInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetRequestId(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetInstanceRequiredCPU(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.InstanceRequiredCPU = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetCpuOverAllocationRatio(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.CpuOverAllocationRatio = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetMemoryOverAllocationRatio(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.MemoryOverAllocationRatio = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetInstanceRequiredMemory(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.InstanceRequiredMemory = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetInstanceClassCode(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.InstanceClassCode = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetAvailableHostNum(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.AvailableHostNum = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetInstanceRequiredStorage(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.InstanceRequiredStorage = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetDiskOverAllocationRatio(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.DiskOverAllocationRatio = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetDedicatedHostGroupId(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBody) SetDedicatedHosts(v []*DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) *DescribeEvaluateDedicatedHostsForInstanceResponseBody {
	s.DedicatedHosts = v
	return s
}

type DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts struct {
	IsContainReadOnlyInstance *bool   `json:"IsContainReadOnlyInstance,omitempty" xml:"IsContainReadOnlyInstance,omitempty"`
	HostAllocationStatus      *bool   `json:"HostAllocationStatus,omitempty" xml:"HostAllocationStatus,omitempty"`
	Comment                   *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostFreeMem               *int32  `json:"HostFreeMem,omitempty" xml:"HostFreeMem,omitempty"`
	HostStatus                *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	HostFreeCPU               *int32  `json:"HostFreeCPU,omitempty" xml:"HostFreeCPU,omitempty"`
	HostStorage               *int32  `json:"HostStorage,omitempty" xml:"HostStorage,omitempty"`
	HostCPU                   *int32  `json:"HostCPU,omitempty" xml:"HostCPU,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostFreeStorage           *int32  `json:"HostFreeStorage,omitempty" xml:"HostFreeStorage,omitempty"`
	VSwitchId                 *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	DedicatedHostName         *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	ZoneId                    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DedicatedHostId           *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Engine                    *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsCouldSSD                *bool   `json:"IsCouldSSD,omitempty" xml:"IsCouldSSD,omitempty"`
	IsAvailableForInstance    *bool   `json:"IsAvailableForInstance,omitempty" xml:"IsAvailableForInstance,omitempty"`
	HostMem                   *int32  `json:"HostMem,omitempty" xml:"HostMem,omitempty"`
	CreatedTime               *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	IPAddress                 *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetIsContainReadOnlyInstance(v bool) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.IsContainReadOnlyInstance = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetHostAllocationStatus(v bool) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.HostAllocationStatus = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetComment(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.Comment = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetHostFreeMem(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.HostFreeMem = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetHostStatus(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.HostStatus = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetHostFreeCPU(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.HostFreeCPU = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetHostStorage(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.HostStorage = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetHostCPU(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.HostCPU = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetRegionId(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.RegionId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetHostFreeStorage(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.HostFreeStorage = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetVSwitchId(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetDedicatedHostName(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetZoneId(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.ZoneId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetDedicatedHostId(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetEngine(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.Engine = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetIsCouldSSD(v bool) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.IsCouldSSD = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetIsAvailableForInstance(v bool) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.IsAvailableForInstance = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetHostMem(v int32) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.HostMem = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetCreatedTime(v int64) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.CreatedTime = &v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts) SetIPAddress(v string) *DescribeEvaluateDedicatedHostsForInstanceResponseBodyDedicatedHosts {
	s.IPAddress = &v
	return s
}

type DescribeEvaluateDedicatedHostsForInstanceResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEvaluateDedicatedHostsForInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEvaluateDedicatedHostsForInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluateDedicatedHostsForInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponse) SetHeaders(v map[string]*string) *DescribeEvaluateDedicatedHostsForInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeEvaluateDedicatedHostsForInstanceResponse) SetBody(v *DescribeEvaluateDedicatedHostsForInstanceResponseBody) *DescribeEvaluateDedicatedHostsForInstanceResponse {
	s.Body = v
	return s
}

type RestartDedicatedHostRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	FailoverMode         *string `json:"FailoverMode,omitempty" xml:"FailoverMode,omitempty"`
	ForceStop            *bool   `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
}

func (s RestartDedicatedHostRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDedicatedHostRequest) GoString() string {
	return s.String()
}

func (s *RestartDedicatedHostRequest) SetOwnerId(v int64) *RestartDedicatedHostRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDedicatedHostRequest) SetResourceOwnerAccount(v string) *RestartDedicatedHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDedicatedHostRequest) SetResourceOwnerId(v int64) *RestartDedicatedHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDedicatedHostRequest) SetRegionId(v string) *RestartDedicatedHostRequest {
	s.RegionId = &v
	return s
}

func (s *RestartDedicatedHostRequest) SetDedicatedHostId(v string) *RestartDedicatedHostRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *RestartDedicatedHostRequest) SetFailoverMode(v string) *RestartDedicatedHostRequest {
	s.FailoverMode = &v
	return s
}

func (s *RestartDedicatedHostRequest) SetForceStop(v bool) *RestartDedicatedHostRequest {
	s.ForceStop = &v
	return s
}

type RestartDedicatedHostResponseBody struct {
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId          *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartDedicatedHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDedicatedHostResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDedicatedHostResponseBody) SetDedicatedHostId(v string) *RestartDedicatedHostResponseBody {
	s.DedicatedHostId = &v
	return s
}

func (s *RestartDedicatedHostResponseBody) SetRequestId(v string) *RestartDedicatedHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDedicatedHostResponseBody) SetTaskId(v int32) *RestartDedicatedHostResponseBody {
	s.TaskId = &v
	return s
}

type RestartDedicatedHostResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartDedicatedHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartDedicatedHostResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDedicatedHostResponse) GoString() string {
	return s.String()
}

func (s *RestartDedicatedHostResponse) SetHeaders(v map[string]*string) *RestartDedicatedHostResponse {
	s.Headers = v
	return s
}

func (s *RestartDedicatedHostResponse) SetBody(v *RestartDedicatedHostResponseBody) *RestartDedicatedHostResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostHealthRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDedicatedHostHealthRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostHealthRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostHealthRequest) SetOwnerId(v int64) *DescribeDedicatedHostHealthRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostHealthRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostHealthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostHealthRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostHealthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostHealthRequest) SetDedicatedHostId(v string) *DescribeDedicatedHostHealthRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostHealthRequest) SetRegionId(v string) *DescribeDedicatedHostHealthRequest {
	s.RegionId = &v
	return s
}

type DescribeDedicatedHostHealthResponseBody struct {
	DedicatedHostId *string                                            `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	ResourceEvent   map[string]interface{}                             `json:"ResourceEvent,omitempty" xml:"ResourceEvent,omitempty"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HealthStatus    map[string]interface{}                             `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	HostEvents      *DescribeDedicatedHostHealthResponseBodyHostEvents `json:"HostEvents,omitempty" xml:"HostEvents,omitempty" type:"Struct"`
}

func (s DescribeDedicatedHostHealthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostHealthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostHealthResponseBody) SetDedicatedHostId(v string) *DescribeDedicatedHostHealthResponseBody {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostHealthResponseBody) SetResourceEvent(v map[string]interface{}) *DescribeDedicatedHostHealthResponseBody {
	s.ResourceEvent = v
	return s
}

func (s *DescribeDedicatedHostHealthResponseBody) SetRequestId(v string) *DescribeDedicatedHostHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostHealthResponseBody) SetHealthStatus(v map[string]interface{}) *DescribeDedicatedHostHealthResponseBody {
	s.HealthStatus = v
	return s
}

func (s *DescribeDedicatedHostHealthResponseBody) SetHostEvents(v *DescribeDedicatedHostHealthResponseBodyHostEvents) *DescribeDedicatedHostHealthResponseBody {
	s.HostEvents = v
	return s
}

type DescribeDedicatedHostHealthResponseBodyHostEvents struct {
	HostEvents []*DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents `json:"HostEvents,omitempty" xml:"HostEvents,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostHealthResponseBodyHostEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostHealthResponseBodyHostEvents) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostHealthResponseBodyHostEvents) SetHostEvents(v []*DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents) *DescribeDedicatedHostHealthResponseBodyHostEvents {
	s.HostEvents = v
	return s
}

type DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents struct {
	EventName        *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventTime        *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	EventType        *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	EventDescription *string `json:"EventDescription,omitempty" xml:"EventDescription,omitempty"`
}

func (s DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents) SetEventName(v string) *DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents {
	s.EventName = &v
	return s
}

func (s *DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents) SetEventTime(v string) *DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents {
	s.EventTime = &v
	return s
}

func (s *DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents) SetEventType(v string) *DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents {
	s.EventType = &v
	return s
}

func (s *DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents) SetEventDescription(v string) *DescribeDedicatedHostHealthResponseBodyHostEventsHostEvents {
	s.EventDescription = &v
	return s
}

type DescribeDedicatedHostHealthResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedHostHealthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedHostHealthResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostHealthResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostHealthResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostHealthResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostHealthResponse) SetBody(v *DescribeDedicatedHostHealthResponseBody) *DescribeDedicatedHostHealthResponse {
	s.Body = v
	return s
}

type DescribeHostWebShellRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeHostWebShellRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostWebShellRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostWebShellRequest) SetOwnerId(v int64) *DescribeHostWebShellRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetResourceOwnerAccount(v string) *DescribeHostWebShellRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetResourceOwnerId(v int64) *DescribeHostWebShellRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetDedicatedHostId(v string) *DescribeHostWebShellRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetRegionId(v string) *DescribeHostWebShellRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetZoneId(v string) *DescribeHostWebShellRequest {
	s.ZoneId = &v
	return s
}

type DescribeHostWebShellResponseBody struct {
	LoginUrl  *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHostWebShellResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostWebShellResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostWebShellResponseBody) SetLoginUrl(v string) *DescribeHostWebShellResponseBody {
	s.LoginUrl = &v
	return s
}

func (s *DescribeHostWebShellResponseBody) SetRequestId(v string) *DescribeHostWebShellResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHostWebShellResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHostWebShellResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHostWebShellResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostWebShellResponse) GoString() string {
	return s.String()
}

func (s *DescribeHostWebShellResponse) SetHeaders(v map[string]*string) *DescribeHostWebShellResponse {
	s.Headers = v
	return s
}

func (s *DescribeHostWebShellResponse) SetBody(v *DescribeHostWebShellResponseBody) *DescribeHostWebShellResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
}

func (s DescribeDedicatedHostAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAttributeRequest) SetOwnerId(v int64) *DescribeDedicatedHostAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostAttributeRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeRequest) SetRegionId(v string) *DescribeDedicatedHostAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeRequest) SetDedicatedHostId(v string) *DescribeDedicatedHostAttributeRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostAttributeRequest {
	s.DedicatedHostGroupId = &v
	return s
}

type DescribeDedicatedHostAttributeResponseBody struct {
	HostType               *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	HostStorage            *int32  `json:"HostStorage,omitempty" xml:"HostStorage,omitempty"`
	InstanceNumberROSlave  *int32  `json:"InstanceNumberROSlave,omitempty" xml:"InstanceNumberROSlave,omitempty"`
	AccountType            *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	MemoryUsed             *string `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	DedicatedHostGroupId   *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceNumberROMaster *int32  `json:"InstanceNumberROMaster,omitempty" xml:"InstanceNumberROMaster,omitempty"`
	AllocationStatus       *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	StorageUsed            *string `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	EcsClassCode           *string `json:"EcsClassCode,omitempty" xml:"EcsClassCode,omitempty"`
	DedicatedHostId        *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	MemAllocationRatio     *string `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	CreatedTime            *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	IPAddress              *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	AutoRenew              *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	HostStatus             *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	HostName               *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostCPU                *int32  `json:"HostCPU,omitempty" xml:"HostCPU,omitempty"`
	OpenPermission         *string `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	InstanceNumber         *int32  `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	CpuUsed                *string `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty"`
	VPCId                  *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	HostClass              *string `json:"HostClass,omitempty" xml:"HostClass,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceNumberMaster   *int32  `json:"InstanceNumberMaster,omitempty" xml:"InstanceNumberMaster,omitempty"`
	VSwitchId              *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	InstanceNumberSlave    *int32  `json:"InstanceNumberSlave,omitempty" xml:"InstanceNumberSlave,omitempty"`
	ExpiredTime            *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ZoneId                 *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	CPUAllocationRatio     *string `json:"CPUAllocationRatio,omitempty" xml:"CPUAllocationRatio,omitempty"`
	ImageCategory          *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	DiskAllocationRatio    *string `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	HostMem                *int32  `json:"HostMem,omitempty" xml:"HostMem,omitempty"`
	AccountName            *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeDedicatedHostAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostType(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.HostType = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostStorage(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.HostStorage = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumberROSlave(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumberROSlave = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetAccountType(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.AccountType = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetMemoryUsed(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.MemoryUsed = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetRequestId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumberROMaster(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumberROMaster = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetAllocationStatus(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetStorageUsed(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetEcsClassCode(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.EcsClassCode = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetDedicatedHostId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetMemAllocationRatio(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.MemAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetCreatedTime(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetIPAddress(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.IPAddress = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetAutoRenew(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostStatus(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.HostStatus = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostName(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.HostName = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostCPU(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.HostCPU = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetOpenPermission(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.OpenPermission = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumber(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetCpuUsed(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.CpuUsed = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetVPCId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostClass(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.HostClass = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetRegionId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumberMaster(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumberMaster = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetVSwitchId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumberSlave(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumberSlave = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetExpiredTime(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetZoneId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetCPUAllocationRatio(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.CPUAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetImageCategory(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.ImageCategory = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetDiskAllocationRatio(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.DiskAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostMem(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.HostMem = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetAccountName(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.AccountName = &v
	return s
}

type DescribeDedicatedHostAttributeResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDedicatedHostAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedHostAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAttributeResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostAttributeResponse) SetBody(v *DescribeDedicatedHostAttributeResponseBody) *DescribeDedicatedHostAttributeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cddc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ModifyDBInstanceSwitchWeightWithOptions(request *ModifyDBInstanceSwitchWeightRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceSwitchWeightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBInstanceSwitchWeightResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBInstanceSwitchWeight"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceSwitchWeight(request *ModifyDBInstanceSwitchWeightRequest) (_result *ModifyDBInstanceSwitchWeightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceSwitchWeightResponse{}
	_body, _err := client.ModifyDBInstanceSwitchWeightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableDedicatedHostZonesWithOptions(request *DescribeAvailableDedicatedHostZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableDedicatedHostZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAvailableDedicatedHostZonesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableDedicatedHostZones"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableDedicatedHostZones(request *DescribeAvailableDedicatedHostZonesRequest) (_result *DescribeAvailableDedicatedHostZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableDedicatedHostZonesResponse{}
	_body, _err := client.DescribeAvailableDedicatedHostZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedHostGroupsWithOptions(request *DescribeDedicatedHostGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedHostGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedHostGroups"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedHostGroups(request *DescribeDedicatedHostGroupsRequest) (_result *DescribeDedicatedHostGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedHostGroupsResponse{}
	_body, _err := client.DescribeDedicatedHostGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMyBaseHostOverViewWithOptions(request *DescribeMyBaseHostOverViewRequest, runtime *util.RuntimeOptions) (_result *DescribeMyBaseHostOverViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMyBaseHostOverViewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMyBaseHostOverView"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMyBaseHostOverView(request *DescribeMyBaseHostOverViewRequest) (_result *DescribeMyBaseHostOverViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMyBaseHostOverViewResponse{}
	_body, _err := client.DescribeMyBaseHostOverViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBriefDedicatedHostsWithOptions(request *DescribeBriefDedicatedHostsRequest, runtime *util.RuntimeOptions) (_result *DescribeBriefDedicatedHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBriefDedicatedHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBriefDedicatedHosts"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBriefDedicatedHosts(request *DescribeBriefDedicatedHostsRequest) (_result *DescribeBriefDedicatedHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBriefDedicatedHostsResponse{}
	_body, _err := client.DescribeBriefDedicatedHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedResourceAdvisorWithOptions(request *DescribeDedicatedResourceAdvisorRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedResourceAdvisorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedResourceAdvisorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedResourceAdvisor"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedResourceAdvisor(request *DescribeDedicatedResourceAdvisorRequest) (_result *DescribeDedicatedResourceAdvisorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedResourceAdvisorResponse{}
	_body, _err := client.DescribeDedicatedResourceAdvisorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeListUserBackupFileRecordWithOptions(request *DescribeListUserBackupFileRecordRequest, runtime *util.RuntimeOptions) (_result *DescribeListUserBackupFileRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeListUserBackupFileRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeListUserBackupFileRecord"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeListUserBackupFileRecord(request *DescribeListUserBackupFileRecordRequest) (_result *DescribeListUserBackupFileRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeListUserBackupFileRecordResponse{}
	_body, _err := client.DescribeListUserBackupFileRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDedicatedHostAccountWithOptions(request *CreateDedicatedHostAccountRequest, runtime *util.RuntimeOptions) (_result *CreateDedicatedHostAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDedicatedHostAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDedicatedHostAccount"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDedicatedHostAccount(request *CreateDedicatedHostAccountRequest) (_result *CreateDedicatedHostAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDedicatedHostAccountResponse{}
	_body, _err := client.CreateDedicatedHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDedicatedHostAccountWithOptions(request *DeleteDedicatedHostAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteDedicatedHostAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDedicatedHostAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDedicatedHostAccount"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDedicatedHostAccount(request *DeleteDedicatedHostAccountRequest) (_result *DeleteDedicatedHostAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDedicatedHostAccountResponse{}
	_body, _err := client.DeleteDedicatedHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDedicatedHostGroupWithOptions(request *DeleteDedicatedHostGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDedicatedHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDedicatedHostGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDedicatedHostGroup"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDedicatedHostGroup(request *DeleteDedicatedHostGroupRequest) (_result *DeleteDedicatedHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDedicatedHostGroupResponse{}
	_body, _err := client.DeleteDedicatedHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckUserIfAuthoriseMyBaseSystemRoleWithOptions(request *CheckUserIfAuthoriseMyBaseSystemRoleRequest, runtime *util.RuntimeOptions) (_result *CheckUserIfAuthoriseMyBaseSystemRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckUserIfAuthoriseMyBaseSystemRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckUserIfAuthoriseMyBaseSystemRole"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUserIfAuthoriseMyBaseSystemRole(request *CheckUserIfAuthoriseMyBaseSystemRoleRequest) (_result *CheckUserIfAuthoriseMyBaseSystemRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUserIfAuthoriseMyBaseSystemRoleResponse{}
	_body, _err := client.CheckUserIfAuthoriseMyBaseSystemRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScheduleInstanceWithOptions(request *DescribeScheduleInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeScheduleInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScheduleInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScheduleInstance"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScheduleInstance(request *DescribeScheduleInstanceRequest) (_result *DescribeScheduleInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScheduleInstanceResponse{}
	_body, _err := client.DescribeScheduleInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExcuteScheduleWithOptions(request *ExcuteScheduleRequest, runtime *util.RuntimeOptions) (_result *ExcuteScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExcuteScheduleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExcuteSchedule"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExcuteSchedule(request *ExcuteScheduleRequest) (_result *ExcuteScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExcuteScheduleResponse{}
	_body, _err := client.ExcuteScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplaceDedicatedHostWithOptions(request *ReplaceDedicatedHostRequest, runtime *util.RuntimeOptions) (_result *ReplaceDedicatedHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReplaceDedicatedHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReplaceDedicatedHost"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplaceDedicatedHost(request *ReplaceDedicatedHostRequest) (_result *ReplaceDedicatedHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplaceDedicatedHostResponse{}
	_body, _err := client.ReplaceDedicatedHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDedicatedHostAccountWithOptions(request *ModifyDedicatedHostAccountRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDedicatedHostAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDedicatedHostAccount"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDedicatedHostAccount(request *ModifyDedicatedHostAccountRequest) (_result *ModifyDedicatedHostAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDedicatedHostAccountResponse{}
	_body, _err := client.ModifyDedicatedHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHostEcsLevelInfoWithOptions(request *DescribeHostEcsLevelInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeHostEcsLevelInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHostEcsLevelInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHostEcsLevelInfo"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHostEcsLevelInfo(request *DescribeHostEcsLevelInfoRequest) (_result *DescribeHostEcsLevelInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHostEcsLevelInfoResponse{}
	_body, _err := client.DescribeHostEcsLevelInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocateHostInstanceCrossVpcVipWithOptions(request *AllocateHostInstanceCrossVpcVipRequest, runtime *util.RuntimeOptions) (_result *AllocateHostInstanceCrossVpcVipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocateHostInstanceCrossVpcVipResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocateHostInstanceCrossVpcVip"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateHostInstanceCrossVpcVip(request *AllocateHostInstanceCrossVpcVipRequest) (_result *AllocateHostInstanceCrossVpcVipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateHostInstanceCrossVpcVipResponse{}
	_body, _err := client.AllocateHostInstanceCrossVpcVipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedHostsWithOptions(request *DescribeDedicatedHostsRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedHosts"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedHosts(request *DescribeDedicatedHostsRequest) (_result *DescribeDedicatedHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedHostsResponse{}
	_body, _err := client.DescribeDedicatedHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedHostDisksWithOptions(request *DescribeDedicatedHostDisksRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostDisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedHostDisksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedHostDisks"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedHostDisks(request *DescribeDedicatedHostDisksRequest) (_result *DescribeDedicatedHostDisksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedHostDisksResponse{}
	_body, _err := client.DescribeDedicatedHostDisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMyBaseInstanceOverViewWithOptions(request *DescribeMyBaseInstanceOverViewRequest, runtime *util.RuntimeOptions) (_result *DescribeMyBaseInstanceOverViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMyBaseInstanceOverViewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMyBaseInstanceOverView"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMyBaseInstanceOverView(request *DescribeMyBaseInstanceOverViewRequest) (_result *DescribeMyBaseInstanceOverViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMyBaseInstanceOverViewResponse{}
	_body, _err := client.DescribeMyBaseInstanceOverViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScheduleMethodWithOptions(request *ModifyScheduleMethodRequest, runtime *util.RuntimeOptions) (_result *ModifyScheduleMethodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyScheduleMethodResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyScheduleMethod"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScheduleMethod(request *ModifyScheduleMethodRequest) (_result *ModifyScheduleMethodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScheduleMethodResponse{}
	_body, _err := client.ModifyScheduleMethodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableDedicatedHostClassesWithOptions(request *DescribeAvailableDedicatedHostClassesRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableDedicatedHostClassesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAvailableDedicatedHostClassesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAvailableDedicatedHostClasses"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableDedicatedHostClasses(request *DescribeAvailableDedicatedHostClassesRequest) (_result *DescribeAvailableDedicatedHostClassesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableDedicatedHostClassesResponse{}
	_body, _err := client.DescribeAvailableDedicatedHostClassesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCheckUserBackupFileWithOptions(request *DescribeCheckUserBackupFileRequest, runtime *util.RuntimeOptions) (_result *DescribeCheckUserBackupFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCheckUserBackupFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCheckUserBackupFile"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCheckUserBackupFile(request *DescribeCheckUserBackupFileRequest) (_result *DescribeCheckUserBackupFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCheckUserBackupFileResponse{}
	_body, _err := client.DescribeCheckUserBackupFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDedicatedHostPasswordWithOptions(request *ModifyDedicatedHostPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDedicatedHostPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDedicatedHostPassword"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDedicatedHostPassword(request *ModifyDedicatedHostPasswordRequest) (_result *ModifyDedicatedHostPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDedicatedHostPasswordResponse{}
	_body, _err := client.ModifyDedicatedHostPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScheduleMethodsWithOptions(request *DescribeScheduleMethodsRequest, runtime *util.RuntimeOptions) (_result *DescribeScheduleMethodsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScheduleMethodsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScheduleMethods"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScheduleMethods(request *DescribeScheduleMethodsRequest) (_result *DescribeScheduleMethodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScheduleMethodsResponse{}
	_body, _err := client.DescribeScheduleMethodsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearDedicatedHostWithOptions(request *ClearDedicatedHostRequest, runtime *util.RuntimeOptions) (_result *ClearDedicatedHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClearDedicatedHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClearDedicatedHost"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearDedicatedHost(request *ClearDedicatedHostRequest) (_result *ClearDedicatedHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearDedicatedHostResponse{}
	_body, _err := client.ClearDedicatedHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserBackupFileRecordWithOptions(request *DeleteUserBackupFileRecordRequest, runtime *util.RuntimeOptions) (_result *DeleteUserBackupFileRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserBackupFileRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserBackupFileRecord"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserBackupFileRecord(request *DeleteUserBackupFileRecordRequest) (_result *DeleteUserBackupFileRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserBackupFileRecordResponse{}
	_body, _err := client.DeleteUserBackupFileRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEvaluateDedicatedHostsWithOptions(request *DescribeEvaluateDedicatedHostsRequest, runtime *util.RuntimeOptions) (_result *DescribeEvaluateDedicatedHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEvaluateDedicatedHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEvaluateDedicatedHosts"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEvaluateDedicatedHosts(request *DescribeEvaluateDedicatedHostsRequest) (_result *DescribeEvaluateDedicatedHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEvaluateDedicatedHostsResponse{}
	_body, _err := client.DescribeEvaluateDedicatedHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHostInstanceMonitorInfoWithOptions(request *DescribeHostInstanceMonitorInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeHostInstanceMonitorInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHostInstanceMonitorInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHostInstanceMonitorInfo"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHostInstanceMonitorInfo(request *DescribeHostInstanceMonitorInfoRequest) (_result *DescribeHostInstanceMonitorInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHostInstanceMonitorInfoResponse{}
	_body, _err := client.DescribeHostInstanceMonitorInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedHostMetricWithOptions(request *DescribeDedicatedHostMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedHostMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedHostMetric"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedHostMetric(request *DescribeDedicatedHostMetricRequest) (_result *DescribeDedicatedHostMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedHostMetricResponse{}
	_body, _err := client.DescribeDedicatedHostMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDedicatedHostWithOptions(request *CreateDedicatedHostRequest, runtime *util.RuntimeOptions) (_result *CreateDedicatedHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDedicatedHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDedicatedHost"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDedicatedHost(request *CreateDedicatedHostRequest) (_result *CreateDedicatedHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDedicatedHostResponse{}
	_body, _err := client.CreateDedicatedHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedInstanceMetricWithOptions(request *DescribeDedicatedInstanceMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedInstanceMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedInstanceMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedInstanceMetric"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedInstanceMetric(request *DescribeDedicatedInstanceMetricRequest) (_result *DescribeDedicatedInstanceMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedInstanceMetricResponse{}
	_body, _err := client.DescribeDedicatedInstanceMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCrossVpcInstanceWithOptions(request *DescribeCrossVpcInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeCrossVpcInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCrossVpcInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCrossVpcInstance"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCrossVpcInstance(request *DescribeCrossVpcInstanceRequest) (_result *DescribeCrossVpcInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCrossVpcInstanceResponse{}
	_body, _err := client.DescribeCrossVpcInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDedicatedHostClassWithOptions(request *ModifyDedicatedHostClassRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDedicatedHostClassResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDedicatedHostClass"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDedicatedHostClass(request *ModifyDedicatedHostClassRequest) (_result *ModifyDedicatedHostClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDedicatedHostClassResponse{}
	_body, _err := client.ModifyDedicatedHostClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedHostsCheckInBastionWithOptions(request *DescribeDedicatedHostsCheckInBastionRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostsCheckInBastionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedHostsCheckInBastionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedHostsCheckInBastion"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedHostsCheckInBastion(request *DescribeDedicatedHostsCheckInBastionRequest) (_result *DescribeDedicatedHostsCheckInBastionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedHostsCheckInBastionResponse{}
	_body, _err := client.DescribeDedicatedHostsCheckInBastionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DropDedicatedHostUserWithOptions(request *DropDedicatedHostUserRequest, runtime *util.RuntimeOptions) (_result *DropDedicatedHostUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DropDedicatedHostUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DropDedicatedHostUser"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DropDedicatedHostUser(request *DropDedicatedHostUserRequest) (_result *DropDedicatedHostUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DropDedicatedHostUserResponse{}
	_body, _err := client.DropDedicatedHostUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedHostsInBastionWithOptions(request *DescribeDedicatedHostsInBastionRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostsInBastionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedHostsInBastionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedHostsInBastion"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedHostsInBastion(request *DescribeDedicatedHostsInBastionRequest) (_result *DescribeDedicatedHostsInBastionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedHostsInBastionResponse{}
	_body, _err := client.DescribeDedicatedHostsInBastionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDedicatedHostGroupAttributeWithOptions(request *ModifyDedicatedHostGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDedicatedHostGroupAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDedicatedHostGroupAttribute"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDedicatedHostGroupAttribute(request *ModifyDedicatedHostGroupAttributeRequest) (_result *ModifyDedicatedHostGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDedicatedHostGroupAttributeResponse{}
	_body, _err := client.ModifyDedicatedHostGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHostInstanceConsoleInfoWithOptions(request *QueryHostInstanceConsoleInfoRequest, runtime *util.RuntimeOptions) (_result *QueryHostInstanceConsoleInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryHostInstanceConsoleInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryHostInstanceConsoleInfo"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHostInstanceConsoleInfo(request *QueryHostInstanceConsoleInfoRequest) (_result *QueryHostInstanceConsoleInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHostInstanceConsoleInfoResponse{}
	_body, _err := client.QueryHostInstanceConsoleInfoWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryHostBaseInfoByInstanceWithOptions(request *QueryHostBaseInfoByInstanceRequest, runtime *util.RuntimeOptions) (_result *QueryHostBaseInfoByInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryHostBaseInfoByInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryHostBaseInfoByInstance"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHostBaseInfoByInstance(request *QueryHostBaseInfoByInstanceRequest) (_result *QueryHostBaseInfoByInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHostBaseInfoByInstanceResponse{}
	_body, _err := client.QueryHostBaseInfoByInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedInstanceDistributionWithOptions(request *DescribeDedicatedInstanceDistributionRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedInstanceDistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedInstanceDistributionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedInstanceDistribution"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedInstanceDistribution(request *DescribeDedicatedInstanceDistributionRequest) (_result *DescribeDedicatedInstanceDistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedInstanceDistributionResponse{}
	_body, _err := client.DescribeDedicatedInstanceDistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScheduleHostWithOptions(request *DescribeScheduleHostRequest, runtime *util.RuntimeOptions) (_result *DescribeScheduleHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScheduleHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScheduleHost"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScheduleHost(request *DescribeScheduleHostRequest) (_result *DescribeScheduleHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScheduleHostResponse{}
	_body, _err := client.DescribeScheduleHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDedicatedHostAttributeWithOptions(request *ModifyDedicatedHostAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDedicatedHostAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDedicatedHostAttribute"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDedicatedHostAttribute(request *ModifyDedicatedHostAttributeRequest) (_result *ModifyDedicatedHostAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDedicatedHostAttributeResponse{}
	_body, _err := client.ModifyDedicatedHostAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDedicatedHostGroupWithOptions(request *CreateDedicatedHostGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDedicatedHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDedicatedHostGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDedicatedHostGroup"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDedicatedHostGroup(request *CreateDedicatedHostGroupRequest) (_result *CreateDedicatedHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDedicatedHostGroupResponse{}
	_body, _err := client.CreateDedicatedHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddHostsToBastionWithOptions(request *AddHostsToBastionRequest, runtime *util.RuntimeOptions) (_result *AddHostsToBastionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddHostsToBastionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddHostsToBastion"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddHostsToBastion(request *AddHostsToBastionRequest) (_result *AddHostsToBastionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddHostsToBastionResponse{}
	_body, _err := client.AddHostsToBastionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachHostsToBastionUserWithOptions(request *AttachHostsToBastionUserRequest, runtime *util.RuntimeOptions) (_result *AttachHostsToBastionUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachHostsToBastionUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachHostsToBastionUser"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachHostsToBastionUser(request *AttachHostsToBastionUserRequest) (_result *AttachHostsToBastionUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachHostsToBastionUserResponse{}
	_body, _err := client.AttachHostsToBastionUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEvaluateDedicatedHostsForInstanceWithOptions(request *DescribeEvaluateDedicatedHostsForInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeEvaluateDedicatedHostsForInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEvaluateDedicatedHostsForInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEvaluateDedicatedHostsForInstance"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEvaluateDedicatedHostsForInstance(request *DescribeEvaluateDedicatedHostsForInstanceRequest) (_result *DescribeEvaluateDedicatedHostsForInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEvaluateDedicatedHostsForInstanceResponse{}
	_body, _err := client.DescribeEvaluateDedicatedHostsForInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartDedicatedHostWithOptions(request *RestartDedicatedHostRequest, runtime *util.RuntimeOptions) (_result *RestartDedicatedHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartDedicatedHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartDedicatedHost"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDedicatedHost(request *RestartDedicatedHostRequest) (_result *RestartDedicatedHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDedicatedHostResponse{}
	_body, _err := client.RestartDedicatedHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedHostHealthWithOptions(request *DescribeDedicatedHostHealthRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostHealthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedHostHealthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedHostHealth"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedHostHealth(request *DescribeDedicatedHostHealthRequest) (_result *DescribeDedicatedHostHealthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedHostHealthResponse{}
	_body, _err := client.DescribeDedicatedHostHealthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHostWebShellWithOptions(request *DescribeHostWebShellRequest, runtime *util.RuntimeOptions) (_result *DescribeHostWebShellResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHostWebShellResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHostWebShell"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHostWebShell(request *DescribeHostWebShellRequest) (_result *DescribeHostWebShellResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHostWebShellResponse{}
	_body, _err := client.DescribeHostWebShellWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedHostAttributeWithOptions(request *DescribeDedicatedHostAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDedicatedHostAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDedicatedHostAttribute"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedHostAttribute(request *DescribeDedicatedHostAttributeRequest) (_result *DescribeDedicatedHostAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedHostAttributeResponse{}
	_body, _err := client.DescribeDedicatedHostAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
