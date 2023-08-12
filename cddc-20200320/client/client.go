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

type CreateDedicatedHostRequest struct {
	AutoRenew            *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	HostClass            *string `json:"HostClass,omitempty" xml:"HostClass,omitempty"`
	HostStorage          *string `json:"HostStorage,omitempty" xml:"HostStorage,omitempty"`
	HostStorageType      *string `json:"HostStorageType,omitempty" xml:"HostStorageType,omitempty"`
	ImageCategory        *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	OsPassword           *string `json:"OsPassword,omitempty" xml:"OsPassword,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UsedTime             *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDedicatedHostRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostRequest) SetAutoRenew(v string) *CreateDedicatedHostRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetClientToken(v string) *CreateDedicatedHostRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetDedicatedHostGroupId(v string) *CreateDedicatedHostRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetHostClass(v string) *CreateDedicatedHostRequest {
	s.HostClass = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetHostStorage(v string) *CreateDedicatedHostRequest {
	s.HostStorage = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetHostStorageType(v string) *CreateDedicatedHostRequest {
	s.HostStorageType = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetImageCategory(v string) *CreateDedicatedHostRequest {
	s.ImageCategory = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetOsPassword(v string) *CreateDedicatedHostRequest {
	s.OsPassword = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetOwnerId(v int64) *CreateDedicatedHostRequest {
	s.OwnerId = &v
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

func (s *CreateDedicatedHostRequest) SetRegionId(v string) *CreateDedicatedHostRequest {
	s.RegionId = &v
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

func (s *CreateDedicatedHostRequest) SetUsedTime(v string) *CreateDedicatedHostRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetVSwitchId(v string) *CreateDedicatedHostRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDedicatedHostRequest) SetZoneId(v string) *CreateDedicatedHostRequest {
	s.ZoneId = &v
	return s
}

type CreateDedicatedHostResponseBody struct {
	DedicateHostList *CreateDedicatedHostResponseBodyDedicateHostList `json:"DedicateHostList,omitempty" xml:"DedicateHostList,omitempty" type:"Struct"`
	OrderId          *int64                                           `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDedicatedHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostResponseBody) SetDedicateHostList(v *CreateDedicatedHostResponseBodyDedicateHostList) *CreateDedicatedHostResponseBody {
	s.DedicateHostList = v
	return s
}

func (s *CreateDedicatedHostResponseBody) SetOrderId(v int64) *CreateDedicatedHostResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDedicatedHostResponseBody) SetRequestId(v string) *CreateDedicatedHostResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDedicatedHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDedicatedHostResponse) SetStatusCode(v int32) *CreateDedicatedHostResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDedicatedHostResponse) SetBody(v *CreateDedicatedHostResponseBody) *CreateDedicatedHostResponse {
	s.Body = v
	return s
}

type CreateDedicatedHostAccountRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	BastionInstanceId    *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDedicatedHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostAccountRequest) SetAccountName(v string) *CreateDedicatedHostAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetAccountPassword(v string) *CreateDedicatedHostAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetAccountType(v string) *CreateDedicatedHostAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetBastionInstanceId(v string) *CreateDedicatedHostAccountRequest {
	s.BastionInstanceId = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetClientToken(v string) *CreateDedicatedHostAccountRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetDedicatedHostId(v string) *CreateDedicatedHostAccountRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetOwnerId(v int64) *CreateDedicatedHostAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDedicatedHostAccountRequest) SetRegionId(v string) *CreateDedicatedHostAccountRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDedicatedHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDedicatedHostAccountResponse) SetStatusCode(v int32) *CreateDedicatedHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDedicatedHostAccountResponse) SetBody(v *CreateDedicatedHostAccountResponseBody) *CreateDedicatedHostAccountResponse {
	s.Body = v
	return s
}

type CreateDedicatedHostGroupRequest struct {
	AllocationPolicy       *string `json:"AllocationPolicy,omitempty" xml:"AllocationPolicy,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CpuAllocationRatio     *int32  `json:"CpuAllocationRatio,omitempty" xml:"CpuAllocationRatio,omitempty"`
	DedicatedHostGroupDesc *string `json:"DedicatedHostGroupDesc,omitempty" xml:"DedicatedHostGroupDesc,omitempty"`
	DiskAllocationRatio    *int32  `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	Engine                 *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	HostReplacePolicy      *string `json:"HostReplacePolicy,omitempty" xml:"HostReplacePolicy,omitempty"`
	MemAllocationRatio     *int32  `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	OpenPermission         *int32  `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VPCId                  *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
}

func (s CreateDedicatedHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostGroupRequest) SetAllocationPolicy(v string) *CreateDedicatedHostGroupRequest {
	s.AllocationPolicy = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetClientToken(v string) *CreateDedicatedHostGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetCpuAllocationRatio(v int32) *CreateDedicatedHostGroupRequest {
	s.CpuAllocationRatio = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetDedicatedHostGroupDesc(v string) *CreateDedicatedHostGroupRequest {
	s.DedicatedHostGroupDesc = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetDiskAllocationRatio(v int32) *CreateDedicatedHostGroupRequest {
	s.DiskAllocationRatio = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetEngine(v string) *CreateDedicatedHostGroupRequest {
	s.Engine = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetHostReplacePolicy(v string) *CreateDedicatedHostGroupRequest {
	s.HostReplacePolicy = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetMemAllocationRatio(v int32) *CreateDedicatedHostGroupRequest {
	s.MemAllocationRatio = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetOpenPermission(v int32) *CreateDedicatedHostGroupRequest {
	s.OpenPermission = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetOwnerId(v int64) *CreateDedicatedHostGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDedicatedHostGroupRequest) SetRegionId(v string) *CreateDedicatedHostGroupRequest {
	s.RegionId = &v
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

func (s *CreateDedicatedHostGroupRequest) SetVPCId(v string) *CreateDedicatedHostGroupRequest {
	s.VPCId = &v
	return s
}

type CreateDedicatedHostGroupResponseBody struct {
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDedicatedHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostGroupResponseBody) SetDedicatedHostGroupId(v string) *CreateDedicatedHostGroupResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateDedicatedHostGroupResponseBody) SetRequestId(v string) *CreateDedicatedHostGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateDedicatedHostGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDedicatedHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDedicatedHostGroupResponse) SetStatusCode(v int32) *CreateDedicatedHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDedicatedHostGroupResponse) SetBody(v *CreateDedicatedHostGroupResponseBody) *CreateDedicatedHostGroupResponse {
	s.Body = v
	return s
}

type CreateMyBaseRequest struct {
	AutoRenew                     *string                            `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientToken                   *string                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DedicatedHostGroupDescription *string                            `json:"DedicatedHostGroupDescription,omitempty" xml:"DedicatedHostGroupDescription,omitempty"`
	DedicatedHostGroupId          *string                            `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ECSClassList                  []*CreateMyBaseRequestECSClassList `json:"ECSClassList,omitempty" xml:"ECSClassList,omitempty" type:"Repeated"`
	EcsDeploymentSetId            *string                            `json:"EcsDeploymentSetId,omitempty" xml:"EcsDeploymentSetId,omitempty"`
	EcsHostName                   *string                            `json:"EcsHostName,omitempty" xml:"EcsHostName,omitempty"`
	EcsInstanceName               *string                            `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	EcsUniqueSuffix               *string                            `json:"EcsUniqueSuffix,omitempty" xml:"EcsUniqueSuffix,omitempty"`
	Engine                        *string                            `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ImageId                       *string                            `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	KeyPairName                   *string                            `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OsPassword                    *string                            `json:"OsPassword,omitempty" xml:"OsPassword,omitempty"`
	OwnerId                       *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PasswordInherit               *string                            `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	PayType                       *string                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                        *string                            `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodType                    *string                            `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	RegionId                      *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount          *string                            `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId               *int64                             `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityGroupId               *string                            `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId                     *string                            `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                         *string                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                        *string                            `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateMyBaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMyBaseRequest) GoString() string {
	return s.String()
}

func (s *CreateMyBaseRequest) SetAutoRenew(v string) *CreateMyBaseRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateMyBaseRequest) SetClientToken(v string) *CreateMyBaseRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMyBaseRequest) SetDedicatedHostGroupDescription(v string) *CreateMyBaseRequest {
	s.DedicatedHostGroupDescription = &v
	return s
}

func (s *CreateMyBaseRequest) SetDedicatedHostGroupId(v string) *CreateMyBaseRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateMyBaseRequest) SetECSClassList(v []*CreateMyBaseRequestECSClassList) *CreateMyBaseRequest {
	s.ECSClassList = v
	return s
}

func (s *CreateMyBaseRequest) SetEcsDeploymentSetId(v string) *CreateMyBaseRequest {
	s.EcsDeploymentSetId = &v
	return s
}

func (s *CreateMyBaseRequest) SetEcsHostName(v string) *CreateMyBaseRequest {
	s.EcsHostName = &v
	return s
}

func (s *CreateMyBaseRequest) SetEcsInstanceName(v string) *CreateMyBaseRequest {
	s.EcsInstanceName = &v
	return s
}

func (s *CreateMyBaseRequest) SetEcsUniqueSuffix(v string) *CreateMyBaseRequest {
	s.EcsUniqueSuffix = &v
	return s
}

func (s *CreateMyBaseRequest) SetEngine(v string) *CreateMyBaseRequest {
	s.Engine = &v
	return s
}

func (s *CreateMyBaseRequest) SetImageId(v string) *CreateMyBaseRequest {
	s.ImageId = &v
	return s
}

func (s *CreateMyBaseRequest) SetKeyPairName(v string) *CreateMyBaseRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateMyBaseRequest) SetOsPassword(v string) *CreateMyBaseRequest {
	s.OsPassword = &v
	return s
}

func (s *CreateMyBaseRequest) SetOwnerId(v int64) *CreateMyBaseRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMyBaseRequest) SetPasswordInherit(v string) *CreateMyBaseRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateMyBaseRequest) SetPayType(v string) *CreateMyBaseRequest {
	s.PayType = &v
	return s
}

func (s *CreateMyBaseRequest) SetPeriod(v string) *CreateMyBaseRequest {
	s.Period = &v
	return s
}

func (s *CreateMyBaseRequest) SetPeriodType(v string) *CreateMyBaseRequest {
	s.PeriodType = &v
	return s
}

func (s *CreateMyBaseRequest) SetRegionId(v string) *CreateMyBaseRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMyBaseRequest) SetResourceOwnerAccount(v string) *CreateMyBaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMyBaseRequest) SetResourceOwnerId(v int64) *CreateMyBaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMyBaseRequest) SetSecurityGroupId(v string) *CreateMyBaseRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMyBaseRequest) SetVSwitchId(v string) *CreateMyBaseRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateMyBaseRequest) SetVpcId(v string) *CreateMyBaseRequest {
	s.VpcId = &v
	return s
}

func (s *CreateMyBaseRequest) SetZoneId(v string) *CreateMyBaseRequest {
	s.ZoneId = &v
	return s
}

type CreateMyBaseRequestECSClassList struct {
	DataDiskPerformanceLevel   *string `json:"dataDiskPerformanceLevel,omitempty" xml:"dataDiskPerformanceLevel,omitempty"`
	DiskCapacity               *int32  `json:"diskCapacity,omitempty" xml:"diskCapacity,omitempty"`
	DiskCount                  *int32  `json:"diskCount,omitempty" xml:"diskCount,omitempty"`
	DiskType                   *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	InstanceType               *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	NodeCount                  *int32  `json:"nodeCount,omitempty" xml:"nodeCount,omitempty"`
	SysDiskCapacity            *int32  `json:"sysDiskCapacity,omitempty" xml:"sysDiskCapacity,omitempty"`
	SysDiskType                *string `json:"sysDiskType,omitempty" xml:"sysDiskType,omitempty"`
	SystemDiskPerformanceLevel *string `json:"systemDiskPerformanceLevel,omitempty" xml:"systemDiskPerformanceLevel,omitempty"`
}

func (s CreateMyBaseRequestECSClassList) String() string {
	return tea.Prettify(s)
}

func (s CreateMyBaseRequestECSClassList) GoString() string {
	return s.String()
}

func (s *CreateMyBaseRequestECSClassList) SetDataDiskPerformanceLevel(v string) *CreateMyBaseRequestECSClassList {
	s.DataDiskPerformanceLevel = &v
	return s
}

func (s *CreateMyBaseRequestECSClassList) SetDiskCapacity(v int32) *CreateMyBaseRequestECSClassList {
	s.DiskCapacity = &v
	return s
}

func (s *CreateMyBaseRequestECSClassList) SetDiskCount(v int32) *CreateMyBaseRequestECSClassList {
	s.DiskCount = &v
	return s
}

func (s *CreateMyBaseRequestECSClassList) SetDiskType(v string) *CreateMyBaseRequestECSClassList {
	s.DiskType = &v
	return s
}

func (s *CreateMyBaseRequestECSClassList) SetInstanceType(v string) *CreateMyBaseRequestECSClassList {
	s.InstanceType = &v
	return s
}

func (s *CreateMyBaseRequestECSClassList) SetNodeCount(v int32) *CreateMyBaseRequestECSClassList {
	s.NodeCount = &v
	return s
}

func (s *CreateMyBaseRequestECSClassList) SetSysDiskCapacity(v int32) *CreateMyBaseRequestECSClassList {
	s.SysDiskCapacity = &v
	return s
}

func (s *CreateMyBaseRequestECSClassList) SetSysDiskType(v string) *CreateMyBaseRequestECSClassList {
	s.SysDiskType = &v
	return s
}

func (s *CreateMyBaseRequestECSClassList) SetSystemDiskPerformanceLevel(v string) *CreateMyBaseRequestECSClassList {
	s.SystemDiskPerformanceLevel = &v
	return s
}

type CreateMyBaseShrinkRequest struct {
	AutoRenew                     *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClientToken                   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DedicatedHostGroupDescription *string `json:"DedicatedHostGroupDescription,omitempty" xml:"DedicatedHostGroupDescription,omitempty"`
	DedicatedHostGroupId          *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ECSClassListShrink            *string `json:"ECSClassList,omitempty" xml:"ECSClassList,omitempty"`
	EcsDeploymentSetId            *string `json:"EcsDeploymentSetId,omitempty" xml:"EcsDeploymentSetId,omitempty"`
	EcsHostName                   *string `json:"EcsHostName,omitempty" xml:"EcsHostName,omitempty"`
	EcsInstanceName               *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	EcsUniqueSuffix               *string `json:"EcsUniqueSuffix,omitempty" xml:"EcsUniqueSuffix,omitempty"`
	Engine                        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ImageId                       *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	KeyPairName                   *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OsPassword                    *string `json:"OsPassword,omitempty" xml:"OsPassword,omitempty"`
	OwnerId                       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PasswordInherit               *string `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	PayType                       *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                        *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodType                    *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	RegionId                      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount          *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId               *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityGroupId               *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId                     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateMyBaseShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMyBaseShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMyBaseShrinkRequest) SetAutoRenew(v string) *CreateMyBaseShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetClientToken(v string) *CreateMyBaseShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetDedicatedHostGroupDescription(v string) *CreateMyBaseShrinkRequest {
	s.DedicatedHostGroupDescription = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetDedicatedHostGroupId(v string) *CreateMyBaseShrinkRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetECSClassListShrink(v string) *CreateMyBaseShrinkRequest {
	s.ECSClassListShrink = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetEcsDeploymentSetId(v string) *CreateMyBaseShrinkRequest {
	s.EcsDeploymentSetId = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetEcsHostName(v string) *CreateMyBaseShrinkRequest {
	s.EcsHostName = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetEcsInstanceName(v string) *CreateMyBaseShrinkRequest {
	s.EcsInstanceName = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetEcsUniqueSuffix(v string) *CreateMyBaseShrinkRequest {
	s.EcsUniqueSuffix = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetEngine(v string) *CreateMyBaseShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetImageId(v string) *CreateMyBaseShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetKeyPairName(v string) *CreateMyBaseShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetOsPassword(v string) *CreateMyBaseShrinkRequest {
	s.OsPassword = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetOwnerId(v int64) *CreateMyBaseShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetPasswordInherit(v string) *CreateMyBaseShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetPayType(v string) *CreateMyBaseShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetPeriod(v string) *CreateMyBaseShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetPeriodType(v string) *CreateMyBaseShrinkRequest {
	s.PeriodType = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetRegionId(v string) *CreateMyBaseShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetResourceOwnerAccount(v string) *CreateMyBaseShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetResourceOwnerId(v int64) *CreateMyBaseShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetSecurityGroupId(v string) *CreateMyBaseShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetVSwitchId(v string) *CreateMyBaseShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetVpcId(v string) *CreateMyBaseShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateMyBaseShrinkRequest) SetZoneId(v string) *CreateMyBaseShrinkRequest {
	s.ZoneId = &v
	return s
}

type CreateMyBaseResponseBody struct {
	OrderList *CreateMyBaseResponseBodyOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMyBaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMyBaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMyBaseResponseBody) SetOrderList(v *CreateMyBaseResponseBodyOrderList) *CreateMyBaseResponseBody {
	s.OrderList = v
	return s
}

func (s *CreateMyBaseResponseBody) SetRequestId(v string) *CreateMyBaseResponseBody {
	s.RequestId = &v
	return s
}

type CreateMyBaseResponseBodyOrderList struct {
	OrderList []*CreateMyBaseResponseBodyOrderListOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Repeated"`
}

func (s CreateMyBaseResponseBodyOrderList) String() string {
	return tea.Prettify(s)
}

func (s CreateMyBaseResponseBodyOrderList) GoString() string {
	return s.String()
}

func (s *CreateMyBaseResponseBodyOrderList) SetOrderList(v []*CreateMyBaseResponseBodyOrderListOrderList) *CreateMyBaseResponseBodyOrderList {
	s.OrderList = v
	return s
}

type CreateMyBaseResponseBodyOrderListOrderList struct {
	CreateTimestamp        *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	DedicatedHostGroupName *string `json:"DedicatedHostGroupName,omitempty" xml:"DedicatedHostGroupName,omitempty"`
	ECSInstanceIds         *string `json:"ECSInstanceIds,omitempty" xml:"ECSInstanceIds,omitempty"`
	OrderId                *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateMyBaseResponseBodyOrderListOrderList) String() string {
	return tea.Prettify(s)
}

func (s CreateMyBaseResponseBodyOrderListOrderList) GoString() string {
	return s.String()
}

func (s *CreateMyBaseResponseBodyOrderListOrderList) SetCreateTimestamp(v int64) *CreateMyBaseResponseBodyOrderListOrderList {
	s.CreateTimestamp = &v
	return s
}

func (s *CreateMyBaseResponseBodyOrderListOrderList) SetDedicatedHostGroupName(v string) *CreateMyBaseResponseBodyOrderListOrderList {
	s.DedicatedHostGroupName = &v
	return s
}

func (s *CreateMyBaseResponseBodyOrderListOrderList) SetECSInstanceIds(v string) *CreateMyBaseResponseBodyOrderListOrderList {
	s.ECSInstanceIds = &v
	return s
}

func (s *CreateMyBaseResponseBodyOrderListOrderList) SetOrderId(v string) *CreateMyBaseResponseBodyOrderListOrderList {
	s.OrderId = &v
	return s
}

type CreateMyBaseResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMyBaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMyBaseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMyBaseResponse) GoString() string {
	return s.String()
}

func (s *CreateMyBaseResponse) SetHeaders(v map[string]*string) *CreateMyBaseResponse {
	s.Headers = v
	return s
}

func (s *CreateMyBaseResponse) SetStatusCode(v int32) *CreateMyBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMyBaseResponse) SetBody(v *CreateMyBaseResponseBody) *CreateMyBaseResponse {
	s.Body = v
	return s
}

type DeleteDedicatedHostAccountRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDedicatedHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedHostAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostAccountRequest) SetAccountName(v string) *DeleteDedicatedHostAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteDedicatedHostAccountRequest) SetDedicatedHostId(v string) *DeleteDedicatedHostAccountRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DeleteDedicatedHostAccountRequest) SetOwnerId(v int64) *DeleteDedicatedHostAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDedicatedHostAccountRequest) SetRegionId(v string) *DeleteDedicatedHostAccountRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDedicatedHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDedicatedHostAccountResponse) SetStatusCode(v int32) *DeleteDedicatedHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDedicatedHostAccountResponse) SetBody(v *DeleteDedicatedHostAccountResponseBody) *DeleteDedicatedHostAccountResponse {
	s.Body = v
	return s
}

type DeleteDedicatedHostGroupRequest struct {
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDedicatedHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDedicatedHostGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostGroupRequest) SetDedicatedHostGroupId(v string) *DeleteDedicatedHostGroupRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DeleteDedicatedHostGroupRequest) SetOwnerId(v int64) *DeleteDedicatedHostGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDedicatedHostGroupRequest) SetRegionId(v string) *DeleteDedicatedHostGroupRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDedicatedHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDedicatedHostGroupResponse) SetStatusCode(v int32) *DeleteDedicatedHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDedicatedHostGroupResponse) SetBody(v *DeleteDedicatedHostGroupResponseBody) *DeleteDedicatedHostGroupResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostAttributeRequest struct {
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDedicatedHostAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAttributeRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostAttributeRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeRequest) SetDedicatedHostId(v string) *DescribeDedicatedHostAttributeRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeRequest) SetOwnerId(v int64) *DescribeDedicatedHostAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeRequest) SetRegionId(v string) *DescribeDedicatedHostAttributeRequest {
	s.RegionId = &v
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

type DescribeDedicatedHostAttributeResponseBody struct {
	AccountName            *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType            *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AllocationStatus       *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	AutoRenew              *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CPUAllocationRatio     *string `json:"CPUAllocationRatio,omitempty" xml:"CPUAllocationRatio,omitempty"`
	CpuUsed                *string `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty"`
	CreatedTime            *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DedicatedHostGroupId   *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHostId        *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DiskAllocationRatio    *string `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	DistributionTag        *string `json:"DistributionTag,omitempty" xml:"DistributionTag,omitempty"`
	EcsClassCode           *string `json:"EcsClassCode,omitempty" xml:"EcsClassCode,omitempty"`
	ExpiredTime            *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HostCPU                *int32  `json:"HostCPU,omitempty" xml:"HostCPU,omitempty"`
	HostClass              *string `json:"HostClass,omitempty" xml:"HostClass,omitempty"`
	HostMem                *int32  `json:"HostMem,omitempty" xml:"HostMem,omitempty"`
	HostName               *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostStatus             *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	HostStorage            *int32  `json:"HostStorage,omitempty" xml:"HostStorage,omitempty"`
	HostType               *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	IPAddress              *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	ImageCategory          *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	InstanceNumber         *int32  `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	InstanceNumberMaster   *int32  `json:"InstanceNumberMaster,omitempty" xml:"InstanceNumberMaster,omitempty"`
	InstanceNumberROMaster *int32  `json:"InstanceNumberROMaster,omitempty" xml:"InstanceNumberROMaster,omitempty"`
	InstanceNumberROSlave  *int32  `json:"InstanceNumberROSlave,omitempty" xml:"InstanceNumberROSlave,omitempty"`
	InstanceNumberSlave    *int32  `json:"InstanceNumberSlave,omitempty" xml:"InstanceNumberSlave,omitempty"`
	MemAllocationRatio     *string `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	MemoryUsed             *string `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	OpenPermission         *string `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageUsed            *string `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	VPCId                  *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// VSwitch ID。
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetAccountName(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.AccountName = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetAccountType(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.AccountType = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetAllocationStatus(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetAutoRenew(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetCPUAllocationRatio(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.CPUAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetCpuUsed(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.CpuUsed = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetCreatedTime(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetDedicatedHostId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetDiskAllocationRatio(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.DiskAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetDistributionTag(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.DistributionTag = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetEcsClassCode(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.EcsClassCode = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetExpiredTime(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostCPU(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.HostCPU = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostClass(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.HostClass = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostMem(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.HostMem = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostName(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.HostName = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostStatus(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.HostStatus = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostStorage(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.HostStorage = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetHostType(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.HostType = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetIPAddress(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.IPAddress = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetImageCategory(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.ImageCategory = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumber(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumberMaster(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumberMaster = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumberROMaster(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumberROMaster = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumberROSlave(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumberROSlave = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetInstanceNumberSlave(v int32) *DescribeDedicatedHostAttributeResponseBody {
	s.InstanceNumberSlave = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetMemAllocationRatio(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.MemAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetMemoryUsed(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.MemoryUsed = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetOpenPermission(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.OpenPermission = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetRegionId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetRequestId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetStorageUsed(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetVPCId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetVSwitchId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponseBody) SetZoneId(v string) *DescribeDedicatedHostAttributeResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeDedicatedHostAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDedicatedHostAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDedicatedHostAttributeResponse) SetStatusCode(v int32) *DescribeDedicatedHostAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedHostAttributeResponse) SetBody(v *DescribeDedicatedHostAttributeResponseBody) *DescribeDedicatedHostAttributeResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostDisksRequest struct {
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDedicatedHostDisksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostDisksRequest) SetDedicatedHostId(v string) *DescribeDedicatedHostDisksRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostDisksRequest) SetOwnerId(v int64) *DescribeDedicatedHostDisksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostDisksRequest) SetRegionId(v string) *DescribeDedicatedHostDisksRequest {
	s.RegionId = &v
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

type DescribeDedicatedHostDisksResponseBody struct {
	DedicatedHostId *string                                        `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Disks           []*DescribeDedicatedHostDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeDedicatedHostDisksResponseBody) SetDisks(v []*DescribeDedicatedHostDisksResponseBodyDisks) *DescribeDedicatedHostDisksResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBody) SetRequestId(v string) *DescribeDedicatedHostDisksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDedicatedHostDisksResponseBodyDisks struct {
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Device           *string `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskId           *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	HasDBInstance    *bool   `json:"HasDBInstance,omitempty" xml:"HasDBInstance,omitempty"`
	MaxIOPS          *int32  `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	MaxThroughput    *int32  `json:"MaxThroughput,omitempty" xml:"MaxThroughput,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostDisksResponseBodyDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetCategory(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetDBInstanceId(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetDevice(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Device = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetDiskId(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetHasDBInstance(v bool) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.HasDBInstance = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetMaxIOPS(v int32) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetMaxThroughput(v int32) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.MaxThroughput = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetPerformanceLevel(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetSize(v int32) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetStatus(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetType(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.Type = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponseBodyDisks) SetZoneId(v string) *DescribeDedicatedHostDisksResponseBodyDisks {
	s.ZoneId = &v
	return s
}

type DescribeDedicatedHostDisksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDedicatedHostDisksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDedicatedHostDisksResponse) SetStatusCode(v int32) *DescribeDedicatedHostDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedHostDisksResponse) SetBody(v *DescribeDedicatedHostDisksResponseBody) *DescribeDedicatedHostDisksResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostGroupsRequest struct {
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ImageCategory        *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDedicatedHostGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostGroupsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetEngine(v string) *DescribeDedicatedHostGroupsRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetImageCategory(v string) *DescribeDedicatedHostGroupsRequest {
	s.ImageCategory = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetOwnerId(v int64) *DescribeDedicatedHostGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetRegionId(v string) *DescribeDedicatedHostGroupsRequest {
	s.RegionId = &v
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

type DescribeDedicatedHostGroupsResponseBody struct {
	DedicatedHostGroups *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups `json:"DedicatedHostGroups,omitempty" xml:"DedicatedHostGroups,omitempty" type:"Struct"`
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDedicatedHostGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBody) SetDedicatedHostGroups(v *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroups) *DescribeDedicatedHostGroupsResponseBody {
	s.DedicatedHostGroups = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBody) SetRequestId(v string) *DescribeDedicatedHostGroupsResponseBody {
	s.RequestId = &v
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
	AllocationPolicy                  *string                                                                                  `json:"AllocationPolicy,omitempty" xml:"AllocationPolicy,omitempty"`
	BastionInstanceId                 *string                                                                                  `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	Category                          *string                                                                                  `json:"Category,omitempty" xml:"Category,omitempty"`
	CpuAllocateRation                 *float32                                                                                 `json:"CpuAllocateRation,omitempty" xml:"CpuAllocateRation,omitempty"`
	CpuAllocatedAmount                *float32                                                                                 `json:"CpuAllocatedAmount,omitempty" xml:"CpuAllocatedAmount,omitempty"`
	CpuAllocationRatio                *int32                                                                                   `json:"CpuAllocationRatio,omitempty" xml:"CpuAllocationRatio,omitempty"`
	CreateTime                        *string                                                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DedicatedHostCountGroupByHostType map[string]interface{}                                                                   `json:"DedicatedHostCountGroupByHostType,omitempty" xml:"DedicatedHostCountGroupByHostType,omitempty"`
	DedicatedHostGroupDesc            *string                                                                                  `json:"DedicatedHostGroupDesc,omitempty" xml:"DedicatedHostGroupDesc,omitempty"`
	DedicatedHostGroupId              *string                                                                                  `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DeployType                        *string                                                                                  `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DiskAllocateRation                *float32                                                                                 `json:"DiskAllocateRation,omitempty" xml:"DiskAllocateRation,omitempty"`
	DiskAllocatedAmount               *float32                                                                                 `json:"DiskAllocatedAmount,omitempty" xml:"DiskAllocatedAmount,omitempty"`
	DiskAllocationRatio               *int32                                                                                   `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	DiskUsedAmount                    *float32                                                                                 `json:"DiskUsedAmount,omitempty" xml:"DiskUsedAmount,omitempty"`
	DiskUtility                       *float32                                                                                 `json:"DiskUtility,omitempty" xml:"DiskUtility,omitempty"`
	Engine                            *string                                                                                  `json:"Engine,omitempty" xml:"Engine,omitempty"`
	HostNumber                        *int32                                                                                   `json:"HostNumber,omitempty" xml:"HostNumber,omitempty"`
	HostReplacePolicy                 *string                                                                                  `json:"HostReplacePolicy,omitempty" xml:"HostReplacePolicy,omitempty"`
	InstanceNumber                    *int32                                                                                   `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	MemAllocateRation                 *float32                                                                                 `json:"MemAllocateRation,omitempty" xml:"MemAllocateRation,omitempty"`
	MemAllocatedAmount                *float32                                                                                 `json:"MemAllocatedAmount,omitempty" xml:"MemAllocatedAmount,omitempty"`
	MemAllocationRatio                *int32                                                                                   `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	MemUsedAmount                     *float32                                                                                 `json:"MemUsedAmount,omitempty" xml:"MemUsedAmount,omitempty"`
	MemUtility                        *float32                                                                                 `json:"MemUtility,omitempty" xml:"MemUtility,omitempty"`
	OpenPermission                    *string                                                                                  `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	Text                              *string                                                                                  `json:"Text,omitempty" xml:"Text,omitempty"`
	VPCId                             *string                                                                                  `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	ZoneIDList                        *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroupsZoneIDList `json:"ZoneIDList,omitempty" xml:"ZoneIDList,omitempty" type:"Struct"`
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetAllocationPolicy(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.AllocationPolicy = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetBastionInstanceId(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.BastionInstanceId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCategory(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCpuAllocateRation(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CpuAllocateRation = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCpuAllocatedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CpuAllocatedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetCpuAllocationRatio(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.CpuAllocationRatio = &v
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

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDedicatedHostGroupDesc(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DedicatedHostGroupDesc = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDeployType(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DeployType = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskAllocateRation(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskAllocateRation = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskAllocatedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskAllocatedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskAllocationRatio(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskUsedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskUsedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetDiskUtility(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.DiskUtility = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetEngine(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetHostNumber(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.HostNumber = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetHostReplacePolicy(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.HostReplacePolicy = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetInstanceNumber(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemAllocateRation(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemAllocateRation = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemAllocatedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemAllocatedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemAllocationRatio(v int32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemUsedAmount(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemUsedAmount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetMemUtility(v float32) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.MemUtility = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetOpenPermission(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.OpenPermission = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetText(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.Text = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups) SetVPCId(v string) *DescribeDedicatedHostGroupsResponseBodyDedicatedHostGroupsDedicatedHostGroups {
	s.VPCId = &v
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDedicatedHostGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDedicatedHostGroupsResponse) SetStatusCode(v int32) *DescribeDedicatedHostGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponse) SetBody(v *DescribeDedicatedHostGroupsResponseBody) *DescribeDedicatedHostGroupsResponse {
	s.Body = v
	return s
}

type DescribeDedicatedHostsRequest struct {
	AllocationStatus     *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	HostStatus           *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	HostType             *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	OrderId              *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumbers          *int32  `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsRequest) SetAllocationStatus(v string) *DescribeDedicatedHostsRequest {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostId(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetHostStatus(v string) *DescribeDedicatedHostsRequest {
	s.HostStatus = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetHostType(v string) *DescribeDedicatedHostsRequest {
	s.HostType = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetOrderId(v int64) *DescribeDedicatedHostsRequest {
	s.OrderId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetOwnerId(v int64) *DescribeDedicatedHostsRequest {
	s.OwnerId = &v
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

func (s *DescribeDedicatedHostsRequest) SetRegionId(v string) *DescribeDedicatedHostsRequest {
	s.RegionId = &v
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

func (s *DescribeDedicatedHostsRequest) SetTags(v string) *DescribeDedicatedHostsRequest {
	s.Tags = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetZoneId(v string) *DescribeDedicatedHostsRequest {
	s.ZoneId = &v
	return s
}

type DescribeDedicatedHostsResponseBody struct {
	DedicatedHostGroupId    *string                                           `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHosts          *DescribeDedicatedHostsResponseBodyDedicatedHosts `json:"DedicatedHosts,omitempty" xml:"DedicatedHosts,omitempty" type:"Struct"`
	MaxAutoScaleHostStorage *int64                                            `json:"MaxAutoScaleHostStorage,omitempty" xml:"MaxAutoScaleHostStorage,omitempty"`
	PageNumbers             *int32                                            `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	PageSize                *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId               *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecords            *int32                                            `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeDedicatedHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetDedicatedHosts(v *DescribeDedicatedHostsResponseBodyDedicatedHosts) *DescribeDedicatedHostsResponseBody {
	s.DedicatedHosts = v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetMaxAutoScaleHostStorage(v int64) *DescribeDedicatedHostsResponseBody {
	s.MaxAutoScaleHostStorage = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetPageNumbers(v int32) *DescribeDedicatedHostsResponseBody {
	s.PageNumbers = &v
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

func (s *DescribeDedicatedHostsResponseBody) SetTotalRecords(v int32) *DescribeDedicatedHostsResponseBody {
	s.TotalRecords = &v
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
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AllocationStatus     *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	BastionInstanceId    *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	CPUAllocationRatio   *string `json:"CPUAllocationRatio,omitempty" xml:"CPUAllocationRatio,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CpuUsed              *string `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty"`
	CreatedTime          *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DiskAllocationRatio  *string `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	DiskInfo             *string `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty"`
	DistributionSymbol   *string `json:"DistributionSymbol,omitempty" xml:"DistributionSymbol,omitempty"`
	DistributionTag      *string `json:"DistributionTag,omitempty" xml:"DistributionTag,omitempty"`
	EcsClassCode         *string `json:"EcsClassCode,omitempty" xml:"EcsClassCode,omitempty"`
	EcsId                *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	HostCPU              *string `json:"HostCPU,omitempty" xml:"HostCPU,omitempty"`
	HostClass            *string `json:"HostClass,omitempty" xml:"HostClass,omitempty"`
	HostMem              *string `json:"HostMem,omitempty" xml:"HostMem,omitempty"`
	HostName             *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostStatus           *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	HostStorage          *string `json:"HostStorage,omitempty" xml:"HostStorage,omitempty"`
	HostType             *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	IPAddress            *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	ImageCategory        *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	InstanceNumber       *string `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	MemAllocationRatio   *string `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	MemoryUsed           *string `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	MssqlSupportVersion  *string `json:"MssqlSupportVersion,omitempty" xml:"MssqlSupportVersion,omitempty"`
	OpenPermission       *string `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	StorageUsed          *string `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetAccountName(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.AccountName = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetAccountType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.AccountType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetAllocationStatus(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetBastionInstanceId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.BastionInstanceId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCPUAllocationRatio(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.CPUAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCategory(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetChargeType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.ChargeType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCpuUsed(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.CpuUsed = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCreatedTime(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDedicatedHostId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDiskAllocationRatio(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DiskAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDiskInfo(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DiskInfo = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDistributionSymbol(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DistributionSymbol = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDistributionTag(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DistributionTag = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetEcsClassCode(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.EcsClassCode = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetEcsId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.EcsId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetEndTime(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.EndTime = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetEngine(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostCPU(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostCPU = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostClass(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostClass = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostMem(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostMem = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostName(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostName = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostStatus(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostStatus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostStorage(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostStorage = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetIPAddress(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.IPAddress = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetImageCategory(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.ImageCategory = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetInstanceNumber(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetMemAllocationRatio(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.MemAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetMemoryUsed(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.MemoryUsed = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetMssqlSupportVersion(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.MssqlSupportVersion = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetOpenPermission(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.OpenPermission = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetStorageUsed(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetVPCId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.VPCId = &v
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

type DescribeDedicatedHostsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDedicatedHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDedicatedHostsResponse) SetStatusCode(v int32) *DescribeDedicatedHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedHostsResponse) SetBody(v *DescribeDedicatedHostsResponseBody) *DescribeDedicatedHostsResponse {
	s.Body = v
	return s
}

type DescribeHostEcsLevelInfoRequest struct {
	DbType               *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ImageCategory        *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StorageType          *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeHostEcsLevelInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostEcsLevelInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostEcsLevelInfoRequest) SetDbType(v string) *DescribeHostEcsLevelInfoRequest {
	s.DbType = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetImageCategory(v string) *DescribeHostEcsLevelInfoRequest {
	s.ImageCategory = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetOwnerId(v int64) *DescribeHostEcsLevelInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetRegionId(v string) *DescribeHostEcsLevelInfoRequest {
	s.RegionId = &v
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

func (s *DescribeHostEcsLevelInfoRequest) SetStorageType(v string) *DescribeHostEcsLevelInfoRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeHostEcsLevelInfoRequest) SetZoneId(v string) *DescribeHostEcsLevelInfoRequest {
	s.ZoneId = &v
	return s
}

type DescribeHostEcsLevelInfoResponseBody struct {
	HostEcsLevelInfos []*DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos `json:"HostEcsLevelInfos,omitempty" xml:"HostEcsLevelInfos,omitempty" type:"Repeated"`
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHostEcsLevelInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostEcsLevelInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostEcsLevelInfoResponseBody) SetHostEcsLevelInfos(v []*DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfos) *DescribeHostEcsLevelInfoResponseBody {
	s.HostEcsLevelInfos = v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBody) SetRequestId(v string) *DescribeHostEcsLevelInfoResponseBody {
	s.RequestId = &v
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
	CloudStorageBandwidth *float32 `json:"CloudStorageBandwidth,omitempty" xml:"CloudStorageBandwidth,omitempty"`
	Cpu                   *int32   `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CpuFrequency          *string  `json:"CpuFrequency,omitempty" xml:"CpuFrequency,omitempty"`
	CpuVersion            *string  `json:"CpuVersion,omitempty" xml:"CpuVersion,omitempty"`
	Description           *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsClass              *string  `json:"EcsClass,omitempty" xml:"EcsClass,omitempty"`
	EcsClassCode          *string  `json:"EcsClassCode,omitempty" xml:"EcsClassCode,omitempty"`
	IsCloudDisk           *int32   `json:"IsCloudDisk,omitempty" xml:"IsCloudDisk,omitempty"`
	LocalStorage          *string  `json:"LocalStorage,omitempty" xml:"LocalStorage,omitempty"`
	Memory                *int32   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NetBandWidth          *float32 `json:"NetBandWidth,omitempty" xml:"NetBandWidth,omitempty"`
	NetPackage            *int32   `json:"NetPackage,omitempty" xml:"NetPackage,omitempty"`
	RdsClassCode          *string  `json:"RdsClassCode,omitempty" xml:"RdsClassCode,omitempty"`
	StorageIops           *int32   `json:"StorageIops,omitempty" xml:"StorageIops,omitempty"`
}

func (s DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) GoString() string {
	return s.String()
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetCloudStorageBandwidth(v float32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.CloudStorageBandwidth = &v
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

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetCpuVersion(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.CpuVersion = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetDescription(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.Description = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetEcsClass(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.EcsClass = &v
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

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetLocalStorage(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.LocalStorage = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetMemory(v int32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.Memory = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetNetBandWidth(v float32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.NetBandWidth = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetNetPackage(v int32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.NetPackage = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetRdsClassCode(v string) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.RdsClassCode = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems) SetStorageIops(v int32) *DescribeHostEcsLevelInfoResponseBodyHostEcsLevelInfosItems {
	s.StorageIops = &v
	return s
}

type DescribeHostEcsLevelInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHostEcsLevelInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeHostEcsLevelInfoResponse) SetStatusCode(v int32) *DescribeHostEcsLevelInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHostEcsLevelInfoResponse) SetBody(v *DescribeHostEcsLevelInfoResponseBody) *DescribeHostEcsLevelInfoResponse {
	s.Body = v
	return s
}

type DescribeHostWebShellRequest struct {
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeHostWebShellRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostWebShellRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostWebShellRequest) SetDedicatedHostId(v string) *DescribeHostWebShellRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetOwnerId(v int64) *DescribeHostWebShellRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHostWebShellRequest) SetRegionId(v string) *DescribeHostWebShellRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHostWebShellResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeHostWebShellResponse) SetStatusCode(v int32) *DescribeHostWebShellResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHostWebShellResponse) SetBody(v *DescribeHostWebShellResponseBody) *DescribeHostWebShellResponse {
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
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRDSRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRDSRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRDSRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRDSRegion) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRDSRegion {
	s.ZoneId = &v
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

type ListTagResourcesRequest struct {
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
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
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostAccountRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDedicatedHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAccountRequest) SetAccountName(v string) *ModifyDedicatedHostAccountRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetAccountPassword(v string) *ModifyDedicatedHostAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetDedicatedHostId(v string) *ModifyDedicatedHostAccountRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetOwnerId(v int64) *ModifyDedicatedHostAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAccountRequest) SetRegionId(v string) *ModifyDedicatedHostAccountRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDedicatedHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDedicatedHostAccountResponse) SetStatusCode(v int32) *ModifyDedicatedHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostAccountResponse) SetBody(v *ModifyDedicatedHostAccountResponseBody) *ModifyDedicatedHostAccountResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostAttributeRequest struct {
	AllocationStatus     *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	HostName             *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDedicatedHostAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAttributeRequest) SetAllocationStatus(v string) *ModifyDedicatedHostAttributeRequest {
	s.AllocationStatus = &v
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

func (s *ModifyDedicatedHostAttributeRequest) SetOwnerId(v int64) *ModifyDedicatedHostAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeRequest) SetRegionId(v string) *ModifyDedicatedHostAttributeRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDedicatedHostAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDedicatedHostAttributeResponse) SetStatusCode(v int32) *ModifyDedicatedHostAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostAttributeResponse) SetBody(v *ModifyDedicatedHostAttributeResponseBody) *ModifyDedicatedHostAttributeResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostClassRequest struct {
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchTime           *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode       *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	TargetClassCode      *string `json:"TargetClassCode,omitempty" xml:"TargetClassCode,omitempty"`
}

func (s ModifyDedicatedHostClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostClassRequest) SetDedicatedHostId(v string) *ModifyDedicatedHostClassRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetOwnerId(v int64) *ModifyDedicatedHostClassRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetRegionId(v string) *ModifyDedicatedHostClassRequest {
	s.RegionId = &v
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

func (s *ModifyDedicatedHostClassRequest) SetSwitchTime(v string) *ModifyDedicatedHostClassRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetSwitchTimeMode(v string) *ModifyDedicatedHostClassRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ModifyDedicatedHostClassRequest) SetTargetClassCode(v string) *ModifyDedicatedHostClassRequest {
	s.TargetClassCode = &v
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDedicatedHostClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDedicatedHostClassResponse) SetStatusCode(v int32) *ModifyDedicatedHostClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostClassResponse) SetBody(v *ModifyDedicatedHostClassResponseBody) *ModifyDedicatedHostClassResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostGroupAttributeRequest struct {
	AllocationPolicy       *string `json:"AllocationPolicy,omitempty" xml:"AllocationPolicy,omitempty"`
	CpuAllocationRatio     *int32  `json:"CpuAllocationRatio,omitempty" xml:"CpuAllocationRatio,omitempty"`
	DedicatedHostGroupDesc *string `json:"DedicatedHostGroupDesc,omitempty" xml:"DedicatedHostGroupDesc,omitempty"`
	DedicatedHostGroupId   *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	DiskAllocationRatio    *int32  `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	HostReplacePolicy      *string `json:"HostReplacePolicy,omitempty" xml:"HostReplacePolicy,omitempty"`
	MemAllocationRatio     *int32  `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	OpenPermission         *string `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDedicatedHostGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetAllocationPolicy(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.AllocationPolicy = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetCpuAllocationRatio(v int32) *ModifyDedicatedHostGroupAttributeRequest {
	s.CpuAllocationRatio = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetDedicatedHostGroupDesc(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.DedicatedHostGroupDesc = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetDedicatedHostGroupId(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetDiskAllocationRatio(v int32) *ModifyDedicatedHostGroupAttributeRequest {
	s.DiskAllocationRatio = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetHostReplacePolicy(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.HostReplacePolicy = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetMemAllocationRatio(v int32) *ModifyDedicatedHostGroupAttributeRequest {
	s.MemAllocationRatio = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetOpenPermission(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.OpenPermission = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetOwnerId(v int64) *ModifyDedicatedHostGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeRequest) SetRegionId(v string) *ModifyDedicatedHostGroupAttributeRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDedicatedHostGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDedicatedHostGroupAttributeResponse) SetStatusCode(v int32) *ModifyDedicatedHostGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostGroupAttributeResponse) SetBody(v *ModifyDedicatedHostGroupAttributeResponseBody) *ModifyDedicatedHostGroupAttributeResponse {
	s.Body = v
	return s
}

type ModifyDedicatedHostPasswordRequest struct {
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	NewPassword          *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDedicatedHostPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedHostPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostPasswordRequest) SetDedicatedHostId(v string) *ModifyDedicatedHostPasswordRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyDedicatedHostPasswordRequest) SetNewPassword(v string) *ModifyDedicatedHostPasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ModifyDedicatedHostPasswordRequest) SetOwnerId(v int64) *ModifyDedicatedHostPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostPasswordRequest) SetRegionId(v string) *ModifyDedicatedHostPasswordRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDedicatedHostPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDedicatedHostPasswordResponse) SetStatusCode(v int32) *ModifyDedicatedHostPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostPasswordResponse) SetBody(v *ModifyDedicatedHostPasswordResponseBody) *ModifyDedicatedHostPasswordResponse {
	s.Body = v
	return s
}

type QueryHostBaseInfoByInstanceRequest struct {
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryHostBaseInfoByInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHostBaseInfoByInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryHostBaseInfoByInstanceRequest) SetDBInstanceId(v string) *QueryHostBaseInfoByInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceRequest) SetOwnerId(v int64) *QueryHostBaseInfoByInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceRequest) SetRegionId(v string) *QueryHostBaseInfoByInstanceRequest {
	s.RegionId = &v
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

type QueryHostBaseInfoByInstanceResponseBody struct {
	HostInstanceConsoleInfos []*QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos `json:"HostInstanceConsoleInfos,omitempty" xml:"HostInstanceConsoleInfos,omitempty" type:"Repeated"`
	RequestId                *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryHostBaseInfoByInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHostBaseInfoByInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHostBaseInfoByInstanceResponseBody) SetHostInstanceConsoleInfos(v []*QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) *QueryHostBaseInfoByInstanceResponseBody {
	s.HostInstanceConsoleInfos = v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBody) SetRequestId(v string) *QueryHostBaseInfoByInstanceResponseBody {
	s.RequestId = &v
	return s
}

type QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos struct {
	ClusterName   *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpiredTime   *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HostName      *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port          *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Role          *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) GoString() string {
	return s.String()
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetClusterName(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.ClusterName = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetEngine(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Engine = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetEngineVersion(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.EngineVersion = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetExpiredTime(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.ExpiredTime = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetHostName(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.HostName = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetIp(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Ip = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetPort(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Port = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetRole(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Role = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetStatus(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.Status = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos) SetVpcId(v string) *QueryHostBaseInfoByInstanceResponseBodyHostInstanceConsoleInfos {
	s.VpcId = &v
	return s
}

type QueryHostBaseInfoByInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHostBaseInfoByInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryHostBaseInfoByInstanceResponse) SetStatusCode(v int32) *QueryHostBaseInfoByInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHostBaseInfoByInstanceResponse) SetBody(v *QueryHostBaseInfoByInstanceResponseBody) *QueryHostBaseInfoByInstanceResponse {
	s.Body = v
	return s
}

type QueryHostInstanceConsoleInfoRequest struct {
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryHostInstanceConsoleInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHostInstanceConsoleInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHostInstanceConsoleInfoRequest) SetDedicatedHostId(v string) *QueryHostInstanceConsoleInfoRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoRequest) SetOwnerId(v int64) *QueryHostInstanceConsoleInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoRequest) SetRegionId(v string) *QueryHostInstanceConsoleInfoRequest {
	s.RegionId = &v
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

type QueryHostInstanceConsoleInfoResponseBody struct {
	HostInstanceConsoleInfos []*QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos `json:"HostInstanceConsoleInfos,omitempty" xml:"HostInstanceConsoleInfos,omitempty" type:"Repeated"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryHostInstanceConsoleInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHostInstanceConsoleInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHostInstanceConsoleInfoResponseBody) SetHostInstanceConsoleInfos(v []*QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) *QueryHostInstanceConsoleInfoResponseBody {
	s.HostInstanceConsoleInfos = v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBody) SetRequestId(v string) *QueryHostInstanceConsoleInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos struct {
	CpuCores                  *int32                                                                    `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	CpuIncreaseRatioValue     *int32                                                                    `json:"CpuIncreaseRatioValue,omitempty" xml:"CpuIncreaseRatioValue,omitempty"`
	DBInstanceDescription     *string                                                                   `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId              *string                                                                   `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DiskSize                  *int32                                                                    `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	Engine                    *string                                                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion             *string                                                                   `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Ip                        *string                                                                   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	LevelName                 *string                                                                   `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	MaxConnIncreaseRatioValue *int32                                                                    `json:"MaxConnIncreaseRatioValue,omitempty" xml:"MaxConnIncreaseRatioValue,omitempty"`
	MemSize                   *int32                                                                    `json:"MemSize,omitempty" xml:"MemSize,omitempty"`
	MemoryIncreaseRatioValue  *int32                                                                    `json:"MemoryIncreaseRatioValue,omitempty" xml:"MemoryIncreaseRatioValue,omitempty"`
	PerfInfo                  *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo `json:"PerfInfo,omitempty" xml:"PerfInfo,omitempty" type:"Struct"`
	Port                      *string                                                                   `json:"Port,omitempty" xml:"Port,omitempty"`
	Role                      *string                                                                   `json:"Role,omitempty" xml:"Role,omitempty"`
	Status                    *string                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) GoString() string {
	return s.String()
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetCpuCores(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.CpuCores = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetCpuIncreaseRatioValue(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.CpuIncreaseRatioValue = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetDBInstanceDescription(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.DBInstanceDescription = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetDBInstanceId(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.DBInstanceId = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetDiskSize(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.DiskSize = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetEngine(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Engine = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetEngineVersion(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.EngineVersion = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetIp(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Ip = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetLevelName(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.LevelName = &v
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

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetMemoryIncreaseRatioValue(v int32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.MemoryIncreaseRatioValue = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetPerfInfo(v *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.PerfInfo = v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetPort(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Port = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetRole(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Role = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos) SetStatus(v string) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfos {
	s.Status = &v
	return s
}

type QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo struct {
	CpuRatio   *float32 `json:"CpuRatio,omitempty" xml:"CpuRatio,omitempty"`
	DiskCurr   *float32 `json:"DiskCurr,omitempty" xml:"DiskCurr,omitempty"`
	MemRatio   *float32 `json:"MemRatio,omitempty" xml:"MemRatio,omitempty"`
	PerfIdbPio *float32 `json:"PerfIdbPio,omitempty" xml:"PerfIdbPio,omitempty"`
}

func (s QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) GoString() string {
	return s.String()
}

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) SetCpuRatio(v float32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo {
	s.CpuRatio = &v
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

func (s *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo) SetPerfIdbPio(v float32) *QueryHostInstanceConsoleInfoResponseBodyHostInstanceConsoleInfosPerfInfo {
	s.PerfIdbPio = &v
	return s
}

type QueryHostInstanceConsoleInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHostInstanceConsoleInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryHostInstanceConsoleInfoResponse) SetStatusCode(v int32) *QueryHostInstanceConsoleInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHostInstanceConsoleInfoResponse) SetBody(v *QueryHostInstanceConsoleInfoResponseBody) *QueryHostInstanceConsoleInfoResponse {
	s.Body = v
	return s
}

type ReplaceDedicatedHostRequest struct {
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	FailoverMode         *string `json:"FailoverMode,omitempty" xml:"FailoverMode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReplaceDedicatedHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceDedicatedHostRequest) GoString() string {
	return s.String()
}

func (s *ReplaceDedicatedHostRequest) SetDedicatedHostId(v string) *ReplaceDedicatedHostRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ReplaceDedicatedHostRequest) SetFailoverMode(v string) *ReplaceDedicatedHostRequest {
	s.FailoverMode = &v
	return s
}

func (s *ReplaceDedicatedHostRequest) SetOwnerId(v int64) *ReplaceDedicatedHostRequest {
	s.OwnerId = &v
	return s
}

func (s *ReplaceDedicatedHostRequest) SetRegionId(v string) *ReplaceDedicatedHostRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReplaceDedicatedHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReplaceDedicatedHostResponse) SetStatusCode(v int32) *ReplaceDedicatedHostResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceDedicatedHostResponse) SetBody(v *ReplaceDedicatedHostResponseBody) *ReplaceDedicatedHostResponse {
	s.Body = v
	return s
}

type RestartDedicatedHostRequest struct {
	DedicatedHostId      *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	FailoverMode         *string `json:"FailoverMode,omitempty" xml:"FailoverMode,omitempty"`
	ForceStop            *bool   `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartDedicatedHostRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDedicatedHostRequest) GoString() string {
	return s.String()
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

func (s *RestartDedicatedHostRequest) SetOwnerId(v int64) *RestartDedicatedHostRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDedicatedHostRequest) SetRegionId(v string) *RestartDedicatedHostRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartDedicatedHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RestartDedicatedHostResponse) SetStatusCode(v int32) *RestartDedicatedHostResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDedicatedHostResponse) SetBody(v *RestartDedicatedHostResponseBody) *RestartDedicatedHostResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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

func (client *Client) CreateDedicatedHostWithOptions(request *CreateDedicatedHostRequest, runtime *util.RuntimeOptions) (_result *CreateDedicatedHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupId)) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HostClass)) {
		query["HostClass"] = request.HostClass
	}

	if !tea.BoolValue(util.IsUnset(request.HostStorage)) {
		query["HostStorage"] = request.HostStorage
	}

	if !tea.BoolValue(util.IsUnset(request.HostStorageType)) {
		query["HostStorageType"] = request.HostStorageType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCategory)) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !tea.BoolValue(util.IsUnset(request.OsPassword)) {
		query["OsPassword"] = request.OsPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDedicatedHost"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDedicatedHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateDedicatedHostAccountWithOptions(request *CreateDedicatedHostAccountRequest, runtime *util.RuntimeOptions) (_result *CreateDedicatedHostAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.BastionInstanceId)) {
		query["BastionInstanceId"] = request.BastionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDedicatedHostAccount"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDedicatedHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateDedicatedHostGroupWithOptions(request *CreateDedicatedHostGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDedicatedHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllocationPolicy)) {
		query["AllocationPolicy"] = request.AllocationPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CpuAllocationRatio)) {
		query["CpuAllocationRatio"] = request.CpuAllocationRatio
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupDesc)) {
		query["DedicatedHostGroupDesc"] = request.DedicatedHostGroupDesc
	}

	if !tea.BoolValue(util.IsUnset(request.DiskAllocationRatio)) {
		query["DiskAllocationRatio"] = request.DiskAllocationRatio
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.HostReplacePolicy)) {
		query["HostReplacePolicy"] = request.HostReplacePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.MemAllocationRatio)) {
		query["MemAllocationRatio"] = request.MemAllocationRatio
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPermission)) {
		query["OpenPermission"] = request.OpenPermission
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDedicatedHostGroup"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDedicatedHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateMyBaseWithOptions(tmpReq *CreateMyBaseRequest, runtime *util.RuntimeOptions) (_result *CreateMyBaseResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateMyBaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ECSClassList)) {
		request.ECSClassListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ECSClassList, tea.String("ECSClassList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupDescription)) {
		query["DedicatedHostGroupDescription"] = request.DedicatedHostGroupDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupId)) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ECSClassListShrink)) {
		query["ECSClassList"] = request.ECSClassListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EcsDeploymentSetId)) {
		query["EcsDeploymentSetId"] = request.EcsDeploymentSetId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsHostName)) {
		query["EcsHostName"] = request.EcsHostName
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceName)) {
		query["EcsInstanceName"] = request.EcsInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.EcsUniqueSuffix)) {
		query["EcsUniqueSuffix"] = request.EcsUniqueSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.OsPassword)) {
		query["OsPassword"] = request.OsPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordInherit)) {
		query["PasswordInherit"] = request.PasswordInherit
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodType)) {
		query["PeriodType"] = request.PeriodType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMyBase"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMyBaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMyBase(request *CreateMyBaseRequest) (_result *CreateMyBaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMyBaseResponse{}
	_body, _err := client.CreateMyBaseWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDedicatedHostAccount"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDedicatedHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupId)) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDedicatedHostGroup"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDedicatedHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDedicatedHostAttributeWithOptions(request *DescribeDedicatedHostAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupId)) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDedicatedHostAttribute"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDedicatedHostAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDedicatedHostDisksWithOptions(request *DescribeDedicatedHostDisksRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostDisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDedicatedHostDisks"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDedicatedHostDisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDedicatedHostGroupsWithOptions(request *DescribeDedicatedHostGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupId)) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCategory)) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDedicatedHostGroups"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDedicatedHostGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeDedicatedHostsWithOptions(request *DescribeDedicatedHostsRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllocationStatus)) {
		query["AllocationStatus"] = request.AllocationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupId)) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.HostStatus)) {
		query["HostStatus"] = request.HostStatus
	}

	if !tea.BoolValue(util.IsUnset(request.HostType)) {
		query["HostType"] = request.HostType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumbers)) {
		query["PageNumbers"] = request.PageNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDedicatedHosts"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDedicatedHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeHostEcsLevelInfoWithOptions(request *DescribeHostEcsLevelInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeHostEcsLevelInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		query["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCategory)) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHostEcsLevelInfo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHostEcsLevelInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeHostWebShellWithOptions(request *DescribeHostWebShellRequest, runtime *util.RuntimeOptions) (_result *DescribeHostWebShellResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHostWebShell"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHostWebShellResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyDedicatedHostAccountWithOptions(request *ModifyDedicatedHostAccountRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDedicatedHostAccount"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDedicatedHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyDedicatedHostAttributeWithOptions(request *ModifyDedicatedHostAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllocationStatus)) {
		query["AllocationStatus"] = request.AllocationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDedicatedHostAttribute"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDedicatedHostAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyDedicatedHostClassWithOptions(request *ModifyDedicatedHostClassRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTimeMode)) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetClassCode)) {
		query["TargetClassCode"] = request.TargetClassCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDedicatedHostClass"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDedicatedHostClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyDedicatedHostGroupAttributeWithOptions(request *ModifyDedicatedHostGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllocationPolicy)) {
		query["AllocationPolicy"] = request.AllocationPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.CpuAllocationRatio)) {
		query["CpuAllocationRatio"] = request.CpuAllocationRatio
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupDesc)) {
		query["DedicatedHostGroupDesc"] = request.DedicatedHostGroupDesc
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedHostGroupId)) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DiskAllocationRatio)) {
		query["DiskAllocationRatio"] = request.DiskAllocationRatio
	}

	if !tea.BoolValue(util.IsUnset(request.HostReplacePolicy)) {
		query["HostReplacePolicy"] = request.HostReplacePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.MemAllocationRatio)) {
		query["MemAllocationRatio"] = request.MemAllocationRatio
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPermission)) {
		query["OpenPermission"] = request.OpenPermission
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDedicatedHostGroupAttribute"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDedicatedHostGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyDedicatedHostPasswordWithOptions(request *ModifyDedicatedHostPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedHostPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.NewPassword)) {
		query["NewPassword"] = request.NewPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDedicatedHostPassword"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDedicatedHostPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QueryHostBaseInfoByInstanceWithOptions(request *QueryHostBaseInfoByInstanceRequest, runtime *util.RuntimeOptions) (_result *QueryHostBaseInfoByInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceId)) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHostBaseInfoByInstance"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHostBaseInfoByInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QueryHostInstanceConsoleInfoWithOptions(request *QueryHostInstanceConsoleInfoRequest, runtime *util.RuntimeOptions) (_result *QueryHostInstanceConsoleInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHostInstanceConsoleInfo"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHostInstanceConsoleInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ReplaceDedicatedHostWithOptions(request *ReplaceDedicatedHostRequest, runtime *util.RuntimeOptions) (_result *ReplaceDedicatedHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.FailoverMode)) {
		query["FailoverMode"] = request.FailoverMode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplaceDedicatedHost"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplaceDedicatedHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RestartDedicatedHostWithOptions(request *RestartDedicatedHostRequest, runtime *util.RuntimeOptions) (_result *RestartDedicatedHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DedicatedHostId)) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !tea.BoolValue(util.IsUnset(request.FailoverMode)) {
		query["FailoverMode"] = request.FailoverMode
	}

	if !tea.BoolValue(util.IsUnset(request.ForceStop)) {
		query["ForceStop"] = request.ForceStop
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDedicatedHost"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartDedicatedHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
