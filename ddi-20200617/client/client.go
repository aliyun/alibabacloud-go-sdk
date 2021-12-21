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

type CloneFlowRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CloneFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowRequest) GoString() string {
	return s.String()
}

func (s *CloneFlowRequest) SetId(v string) *CloneFlowRequest {
	s.Id = &v
	return s
}

func (s *CloneFlowRequest) SetProjectId(v string) *CloneFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *CloneFlowRequest) SetRegionId(v string) *CloneFlowRequest {
	s.RegionId = &v
	return s
}

type CloneFlowResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CloneFlowResponseBody) SetId(v string) *CloneFlowResponseBody {
	s.Id = &v
	return s
}

func (s *CloneFlowResponseBody) SetRequestId(v string) *CloneFlowResponseBody {
	s.RequestId = &v
	return s
}

type CloneFlowResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloneFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloneFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowResponse) GoString() string {
	return s.String()
}

func (s *CloneFlowResponse) SetHeaders(v map[string]*string) *CloneFlowResponse {
	s.Headers = v
	return s
}

func (s *CloneFlowResponse) SetBody(v *CloneFlowResponseBody) *CloneFlowResponse {
	s.Body = v
	return s
}

type CloneFlowJobRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CloneFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowJobRequest) GoString() string {
	return s.String()
}

func (s *CloneFlowJobRequest) SetId(v string) *CloneFlowJobRequest {
	s.Id = &v
	return s
}

func (s *CloneFlowJobRequest) SetName(v string) *CloneFlowJobRequest {
	s.Name = &v
	return s
}

func (s *CloneFlowJobRequest) SetProjectId(v string) *CloneFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *CloneFlowJobRequest) SetRegionId(v string) *CloneFlowJobRequest {
	s.RegionId = &v
	return s
}

type CloneFlowJobResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *CloneFlowJobResponseBody) SetId(v string) *CloneFlowJobResponseBody {
	s.Id = &v
	return s
}

func (s *CloneFlowJobResponseBody) SetRequestId(v string) *CloneFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type CloneFlowJobResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloneFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloneFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowJobResponse) GoString() string {
	return s.String()
}

func (s *CloneFlowJobResponse) SetHeaders(v map[string]*string) *CloneFlowJobResponse {
	s.Headers = v
	return s
}

func (s *CloneFlowJobResponse) SetBody(v *CloneFlowJobResponseBody) *CloneFlowJobResponse {
	s.Body = v
	return s
}

type CommitFlowEntitySnapshotRequest struct {
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType      *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CommitFlowEntitySnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CommitFlowEntitySnapshotRequest) GoString() string {
	return s.String()
}

func (s *CommitFlowEntitySnapshotRequest) SetEntityId(v string) *CommitFlowEntitySnapshotRequest {
	s.EntityId = &v
	return s
}

func (s *CommitFlowEntitySnapshotRequest) SetEntityType(v string) *CommitFlowEntitySnapshotRequest {
	s.EntityType = &v
	return s
}

func (s *CommitFlowEntitySnapshotRequest) SetMessage(v string) *CommitFlowEntitySnapshotRequest {
	s.Message = &v
	return s
}

func (s *CommitFlowEntitySnapshotRequest) SetRegionId(v string) *CommitFlowEntitySnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CommitFlowEntitySnapshotRequest) SetResourceOwnerId(v int64) *CommitFlowEntitySnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

type CommitFlowEntitySnapshotResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CommitFlowEntitySnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommitFlowEntitySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CommitFlowEntitySnapshotResponseBody) SetData(v string) *CommitFlowEntitySnapshotResponseBody {
	s.Data = &v
	return s
}

func (s *CommitFlowEntitySnapshotResponseBody) SetRequestId(v string) *CommitFlowEntitySnapshotResponseBody {
	s.RequestId = &v
	return s
}

type CommitFlowEntitySnapshotResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CommitFlowEntitySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommitFlowEntitySnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CommitFlowEntitySnapshotResponse) GoString() string {
	return s.String()
}

func (s *CommitFlowEntitySnapshotResponse) SetHeaders(v map[string]*string) *CommitFlowEntitySnapshotResponse {
	s.Headers = v
	return s
}

func (s *CommitFlowEntitySnapshotResponse) SetBody(v *CommitFlowEntitySnapshotResponseBody) *CommitFlowEntitySnapshotResponse {
	s.Body = v
	return s
}

type CreateClusterV2Request struct {
	AuthorizeContent       *string                                    `json:"AuthorizeContent,omitempty" xml:"AuthorizeContent,omitempty"`
	Auto                   *bool                                      `json:"Auto,omitempty" xml:"Auto,omitempty"`
	AutoPayOrder           *bool                                      `json:"AutoPayOrder,omitempty" xml:"AutoPayOrder,omitempty"`
	BootstrapAction        []*CreateClusterV2RequestBootstrapAction   `json:"BootstrapAction,omitempty" xml:"BootstrapAction,omitempty" type:"Repeated"`
	ChargeType             *string                                    `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClickHouseConf         *string                                    `json:"ClickHouseConf,omitempty" xml:"ClickHouseConf,omitempty"`
	ClientToken            *string                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterType            *string                                    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Config                 []*CreateClusterV2RequestConfig            `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
	Configurations         *string                                    `json:"Configurations,omitempty" xml:"Configurations,omitempty"`
	DepositType            *string                                    `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	EmrVer                 *string                                    `json:"EmrVer,omitempty" xml:"EmrVer,omitempty"`
	EnableEas              *bool                                      `json:"EnableEas,omitempty" xml:"EnableEas,omitempty"`
	EnableHighAvailability *bool                                      `json:"EnableHighAvailability,omitempty" xml:"EnableHighAvailability,omitempty"`
	EnableSsh              *bool                                      `json:"EnableSsh,omitempty" xml:"EnableSsh,omitempty"`
	ExtraAttributes        *string                                    `json:"ExtraAttributes,omitempty" xml:"ExtraAttributes,omitempty"`
	HostComponentInfo      []*CreateClusterV2RequestHostComponentInfo `json:"HostComponentInfo,omitempty" xml:"HostComponentInfo,omitempty" type:"Repeated"`
	HostGroup              []*CreateClusterV2RequestHostGroup         `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Repeated"`
	InitCustomHiveMetaDB   *bool                                      `json:"InitCustomHiveMetaDB,omitempty" xml:"InitCustomHiveMetaDB,omitempty"`
	InstanceGeneration     *string                                    `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	IsOpenPublicIp         *bool                                      `json:"IsOpenPublicIp,omitempty" xml:"IsOpenPublicIp,omitempty"`
	KeyPairName            *string                                    `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LogPath                *string                                    `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	MachineType            *string                                    `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	MasterPwd              *string                                    `json:"MasterPwd,omitempty" xml:"MasterPwd,omitempty"`
	MetaStoreConf          *string                                    `json:"MetaStoreConf,omitempty" xml:"MetaStoreConf,omitempty"`
	MetaStoreType          *string                                    `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	Name                   *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	NetType                *string                                    `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Period                 *int32                                     `json:"Period,omitempty" xml:"Period,omitempty"`
	PromotionInfo          []*CreateClusterV2RequestPromotionInfo     `json:"PromotionInfo,omitempty" xml:"PromotionInfo,omitempty" type:"Repeated"`
	RegionId               *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelatedClusterId       *string                                    `json:"RelatedClusterId,omitempty" xml:"RelatedClusterId,omitempty"`
	ResourceGroupId        *string                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId        *int64                                     `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityGroupId        *string                                    `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName      *string                                    `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	ServiceInfo            []*CreateClusterV2RequestServiceInfo       `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	Tag                    []*CreateClusterV2RequestTag               `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UseCustomHiveMetaDB    *bool                                      `json:"UseCustomHiveMetaDB,omitempty" xml:"UseCustomHiveMetaDB,omitempty"`
	UseLocalMetaDb         *bool                                      `json:"UseLocalMetaDb,omitempty" xml:"UseLocalMetaDb,omitempty"`
	UserDefinedEmrEcsRole  *string                                    `json:"UserDefinedEmrEcsRole,omitempty" xml:"UserDefinedEmrEcsRole,omitempty"`
	UserInfo               []*CreateClusterV2RequestUserInfo          `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Repeated"`
	VSwitchId              *string                                    `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                  *string                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WhiteListType          *string                                    `json:"WhiteListType,omitempty" xml:"WhiteListType,omitempty"`
	ZoneId                 *string                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterV2Request) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2Request) GoString() string {
	return s.String()
}

func (s *CreateClusterV2Request) SetAuthorizeContent(v string) *CreateClusterV2Request {
	s.AuthorizeContent = &v
	return s
}

func (s *CreateClusterV2Request) SetAuto(v bool) *CreateClusterV2Request {
	s.Auto = &v
	return s
}

func (s *CreateClusterV2Request) SetAutoPayOrder(v bool) *CreateClusterV2Request {
	s.AutoPayOrder = &v
	return s
}

func (s *CreateClusterV2Request) SetBootstrapAction(v []*CreateClusterV2RequestBootstrapAction) *CreateClusterV2Request {
	s.BootstrapAction = v
	return s
}

func (s *CreateClusterV2Request) SetChargeType(v string) *CreateClusterV2Request {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterV2Request) SetClickHouseConf(v string) *CreateClusterV2Request {
	s.ClickHouseConf = &v
	return s
}

func (s *CreateClusterV2Request) SetClientToken(v string) *CreateClusterV2Request {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterV2Request) SetClusterType(v string) *CreateClusterV2Request {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterV2Request) SetConfig(v []*CreateClusterV2RequestConfig) *CreateClusterV2Request {
	s.Config = v
	return s
}

func (s *CreateClusterV2Request) SetConfigurations(v string) *CreateClusterV2Request {
	s.Configurations = &v
	return s
}

func (s *CreateClusterV2Request) SetDepositType(v string) *CreateClusterV2Request {
	s.DepositType = &v
	return s
}

func (s *CreateClusterV2Request) SetEmrVer(v string) *CreateClusterV2Request {
	s.EmrVer = &v
	return s
}

func (s *CreateClusterV2Request) SetEnableEas(v bool) *CreateClusterV2Request {
	s.EnableEas = &v
	return s
}

func (s *CreateClusterV2Request) SetEnableHighAvailability(v bool) *CreateClusterV2Request {
	s.EnableHighAvailability = &v
	return s
}

func (s *CreateClusterV2Request) SetEnableSsh(v bool) *CreateClusterV2Request {
	s.EnableSsh = &v
	return s
}

func (s *CreateClusterV2Request) SetExtraAttributes(v string) *CreateClusterV2Request {
	s.ExtraAttributes = &v
	return s
}

func (s *CreateClusterV2Request) SetHostComponentInfo(v []*CreateClusterV2RequestHostComponentInfo) *CreateClusterV2Request {
	s.HostComponentInfo = v
	return s
}

func (s *CreateClusterV2Request) SetHostGroup(v []*CreateClusterV2RequestHostGroup) *CreateClusterV2Request {
	s.HostGroup = v
	return s
}

func (s *CreateClusterV2Request) SetInitCustomHiveMetaDB(v bool) *CreateClusterV2Request {
	s.InitCustomHiveMetaDB = &v
	return s
}

func (s *CreateClusterV2Request) SetInstanceGeneration(v string) *CreateClusterV2Request {
	s.InstanceGeneration = &v
	return s
}

func (s *CreateClusterV2Request) SetIsOpenPublicIp(v bool) *CreateClusterV2Request {
	s.IsOpenPublicIp = &v
	return s
}

func (s *CreateClusterV2Request) SetKeyPairName(v string) *CreateClusterV2Request {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterV2Request) SetLogPath(v string) *CreateClusterV2Request {
	s.LogPath = &v
	return s
}

func (s *CreateClusterV2Request) SetMachineType(v string) *CreateClusterV2Request {
	s.MachineType = &v
	return s
}

func (s *CreateClusterV2Request) SetMasterPwd(v string) *CreateClusterV2Request {
	s.MasterPwd = &v
	return s
}

func (s *CreateClusterV2Request) SetMetaStoreConf(v string) *CreateClusterV2Request {
	s.MetaStoreConf = &v
	return s
}

func (s *CreateClusterV2Request) SetMetaStoreType(v string) *CreateClusterV2Request {
	s.MetaStoreType = &v
	return s
}

func (s *CreateClusterV2Request) SetName(v string) *CreateClusterV2Request {
	s.Name = &v
	return s
}

func (s *CreateClusterV2Request) SetNetType(v string) *CreateClusterV2Request {
	s.NetType = &v
	return s
}

func (s *CreateClusterV2Request) SetPeriod(v int32) *CreateClusterV2Request {
	s.Period = &v
	return s
}

func (s *CreateClusterV2Request) SetPromotionInfo(v []*CreateClusterV2RequestPromotionInfo) *CreateClusterV2Request {
	s.PromotionInfo = v
	return s
}

func (s *CreateClusterV2Request) SetRegionId(v string) *CreateClusterV2Request {
	s.RegionId = &v
	return s
}

func (s *CreateClusterV2Request) SetRelatedClusterId(v string) *CreateClusterV2Request {
	s.RelatedClusterId = &v
	return s
}

func (s *CreateClusterV2Request) SetResourceGroupId(v string) *CreateClusterV2Request {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterV2Request) SetResourceOwnerId(v int64) *CreateClusterV2Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateClusterV2Request) SetSecurityGroupId(v string) *CreateClusterV2Request {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterV2Request) SetSecurityGroupName(v string) *CreateClusterV2Request {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateClusterV2Request) SetServiceInfo(v []*CreateClusterV2RequestServiceInfo) *CreateClusterV2Request {
	s.ServiceInfo = v
	return s
}

func (s *CreateClusterV2Request) SetTag(v []*CreateClusterV2RequestTag) *CreateClusterV2Request {
	s.Tag = v
	return s
}

func (s *CreateClusterV2Request) SetUseCustomHiveMetaDB(v bool) *CreateClusterV2Request {
	s.UseCustomHiveMetaDB = &v
	return s
}

func (s *CreateClusterV2Request) SetUseLocalMetaDb(v bool) *CreateClusterV2Request {
	s.UseLocalMetaDb = &v
	return s
}

func (s *CreateClusterV2Request) SetUserDefinedEmrEcsRole(v string) *CreateClusterV2Request {
	s.UserDefinedEmrEcsRole = &v
	return s
}

func (s *CreateClusterV2Request) SetUserInfo(v []*CreateClusterV2RequestUserInfo) *CreateClusterV2Request {
	s.UserInfo = v
	return s
}

func (s *CreateClusterV2Request) SetVSwitchId(v string) *CreateClusterV2Request {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterV2Request) SetVpcId(v string) *CreateClusterV2Request {
	s.VpcId = &v
	return s
}

func (s *CreateClusterV2Request) SetWhiteListType(v string) *CreateClusterV2Request {
	s.WhiteListType = &v
	return s
}

func (s *CreateClusterV2Request) SetZoneId(v string) *CreateClusterV2Request {
	s.ZoneId = &v
	return s
}

type CreateClusterV2RequestBootstrapAction struct {
	Arg  *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateClusterV2RequestBootstrapAction) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestBootstrapAction) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestBootstrapAction) SetArg(v string) *CreateClusterV2RequestBootstrapAction {
	s.Arg = &v
	return s
}

func (s *CreateClusterV2RequestBootstrapAction) SetName(v string) *CreateClusterV2RequestBootstrapAction {
	s.Name = &v
	return s
}

func (s *CreateClusterV2RequestBootstrapAction) SetPath(v string) *CreateClusterV2RequestBootstrapAction {
	s.Path = &v
	return s
}

type CreateClusterV2RequestConfig struct {
	ConfigKey   *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Encrypt     *string `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Replace     *string `json:"Replace,omitempty" xml:"Replace,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateClusterV2RequestConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestConfig) SetConfigKey(v string) *CreateClusterV2RequestConfig {
	s.ConfigKey = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetConfigValue(v string) *CreateClusterV2RequestConfig {
	s.ConfigValue = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetEncrypt(v string) *CreateClusterV2RequestConfig {
	s.Encrypt = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetFileName(v string) *CreateClusterV2RequestConfig {
	s.FileName = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetReplace(v string) *CreateClusterV2RequestConfig {
	s.Replace = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetServiceName(v string) *CreateClusterV2RequestConfig {
	s.ServiceName = &v
	return s
}

type CreateClusterV2RequestHostComponentInfo struct {
	ComponentNameList []*string `json:"ComponentNameList,omitempty" xml:"ComponentNameList,omitempty" type:"Repeated"`
	HostName          *string   `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ServiceName       *string   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateClusterV2RequestHostComponentInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestHostComponentInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestHostComponentInfo) SetComponentNameList(v []*string) *CreateClusterV2RequestHostComponentInfo {
	s.ComponentNameList = v
	return s
}

func (s *CreateClusterV2RequestHostComponentInfo) SetHostName(v string) *CreateClusterV2RequestHostComponentInfo {
	s.HostName = &v
	return s
}

func (s *CreateClusterV2RequestHostComponentInfo) SetServiceName(v string) *CreateClusterV2RequestHostComponentInfo {
	s.ServiceName = &v
	return s
}

type CreateClusterV2RequestHostGroup struct {
	AutoRenew       *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType      *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Comment         *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	DiskCapacity    *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	DiskCount       *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	DiskType        *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	GpuDriver       *string `json:"GpuDriver,omitempty" xml:"GpuDriver,omitempty"`
	HostGroupId     *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName   *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	HostGroupType   *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NodeCount       *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	SysDiskCapacity *int32  `json:"SysDiskCapacity,omitempty" xml:"SysDiskCapacity,omitempty"`
	SysDiskType     *string `json:"SysDiskType,omitempty" xml:"SysDiskType,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateClusterV2RequestHostGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestHostGroup) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestHostGroup) SetAutoRenew(v bool) *CreateClusterV2RequestHostGroup {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetChargeType(v string) *CreateClusterV2RequestHostGroup {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetClusterId(v string) *CreateClusterV2RequestHostGroup {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetComment(v string) *CreateClusterV2RequestHostGroup {
	s.Comment = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetCreateType(v string) *CreateClusterV2RequestHostGroup {
	s.CreateType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetDiskCapacity(v int32) *CreateClusterV2RequestHostGroup {
	s.DiskCapacity = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetDiskCount(v int32) *CreateClusterV2RequestHostGroup {
	s.DiskCount = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetDiskType(v string) *CreateClusterV2RequestHostGroup {
	s.DiskType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetGpuDriver(v string) *CreateClusterV2RequestHostGroup {
	s.GpuDriver = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetHostGroupId(v string) *CreateClusterV2RequestHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetHostGroupName(v string) *CreateClusterV2RequestHostGroup {
	s.HostGroupName = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetHostGroupType(v string) *CreateClusterV2RequestHostGroup {
	s.HostGroupType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetInstanceType(v string) *CreateClusterV2RequestHostGroup {
	s.InstanceType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetNodeCount(v int32) *CreateClusterV2RequestHostGroup {
	s.NodeCount = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetPeriod(v int32) *CreateClusterV2RequestHostGroup {
	s.Period = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetSysDiskCapacity(v int32) *CreateClusterV2RequestHostGroup {
	s.SysDiskCapacity = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetSysDiskType(v string) *CreateClusterV2RequestHostGroup {
	s.SysDiskType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetVSwitchId(v string) *CreateClusterV2RequestHostGroup {
	s.VSwitchId = &v
	return s
}

type CreateClusterV2RequestPromotionInfo struct {
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	PromotionOptionNo   *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s CreateClusterV2RequestPromotionInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestPromotionInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestPromotionInfo) SetProductCode(v string) *CreateClusterV2RequestPromotionInfo {
	s.ProductCode = &v
	return s
}

func (s *CreateClusterV2RequestPromotionInfo) SetPromotionOptionCode(v string) *CreateClusterV2RequestPromotionInfo {
	s.PromotionOptionCode = &v
	return s
}

func (s *CreateClusterV2RequestPromotionInfo) SetPromotionOptionNo(v string) *CreateClusterV2RequestPromotionInfo {
	s.PromotionOptionNo = &v
	return s
}

type CreateClusterV2RequestServiceInfo struct {
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s CreateClusterV2RequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestServiceInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestServiceInfo) SetServiceName(v string) *CreateClusterV2RequestServiceInfo {
	s.ServiceName = &v
	return s
}

func (s *CreateClusterV2RequestServiceInfo) SetServiceVersion(v string) *CreateClusterV2RequestServiceInfo {
	s.ServiceVersion = &v
	return s
}

type CreateClusterV2RequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterV2RequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestTag) SetKey(v string) *CreateClusterV2RequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterV2RequestTag) SetValue(v string) *CreateClusterV2RequestTag {
	s.Value = &v
	return s
}

type CreateClusterV2RequestUserInfo struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateClusterV2RequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestUserInfo) SetPassword(v string) *CreateClusterV2RequestUserInfo {
	s.Password = &v
	return s
}

func (s *CreateClusterV2RequestUserInfo) SetUserId(v string) *CreateClusterV2RequestUserInfo {
	s.UserId = &v
	return s
}

func (s *CreateClusterV2RequestUserInfo) SetUserName(v string) *CreateClusterV2RequestUserInfo {
	s.UserName = &v
	return s
}

type CreateClusterV2ResponseBody struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CoreOrderId   *string `json:"CoreOrderId,omitempty" xml:"CoreOrderId,omitempty"`
	EmrOrderId    *string `json:"EmrOrderId,omitempty" xml:"EmrOrderId,omitempty"`
	MasterOrderId *string `json:"MasterOrderId,omitempty" xml:"MasterOrderId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterV2ResponseBody) SetClusterId(v string) *CreateClusterV2ResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterV2ResponseBody) SetCoreOrderId(v string) *CreateClusterV2ResponseBody {
	s.CoreOrderId = &v
	return s
}

func (s *CreateClusterV2ResponseBody) SetEmrOrderId(v string) *CreateClusterV2ResponseBody {
	s.EmrOrderId = &v
	return s
}

func (s *CreateClusterV2ResponseBody) SetMasterOrderId(v string) *CreateClusterV2ResponseBody {
	s.MasterOrderId = &v
	return s
}

func (s *CreateClusterV2ResponseBody) SetRequestId(v string) *CreateClusterV2ResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterV2Response struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterV2Response) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2Response) GoString() string {
	return s.String()
}

func (s *CreateClusterV2Response) SetHeaders(v map[string]*string) *CreateClusterV2Response {
	s.Headers = v
	return s
}

func (s *CreateClusterV2Response) SetBody(v *CreateClusterV2ResponseBody) *CreateClusterV2Response {
	s.Body = v
	return s
}

type CreateFlowRequest struct {
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	Application             *string `json:"Application,omitempty" xml:"Application,omitempty"`
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	CronExpression          *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace               *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParentCategory          *string `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	ParentFlowList          *string `json:"ParentFlowList,omitempty" xml:"ParentFlowList,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) SetAlertConf(v string) *CreateFlowRequest {
	s.AlertConf = &v
	return s
}

func (s *CreateFlowRequest) SetAlertDingDingGroupBizId(v string) *CreateFlowRequest {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *CreateFlowRequest) SetAlertUserGroupBizId(v string) *CreateFlowRequest {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *CreateFlowRequest) SetApplication(v string) *CreateFlowRequest {
	s.Application = &v
	return s
}

func (s *CreateFlowRequest) SetClientToken(v string) *CreateFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowRequest) SetClusterId(v string) *CreateFlowRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateFlowRequest) SetCreateCluster(v bool) *CreateFlowRequest {
	s.CreateCluster = &v
	return s
}

func (s *CreateFlowRequest) SetCronExpression(v string) *CreateFlowRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateFlowRequest) SetDescription(v string) *CreateFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowRequest) SetEndSchedule(v int64) *CreateFlowRequest {
	s.EndSchedule = &v
	return s
}

func (s *CreateFlowRequest) SetHostName(v string) *CreateFlowRequest {
	s.HostName = &v
	return s
}

func (s *CreateFlowRequest) SetName(v string) *CreateFlowRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowRequest) SetNamespace(v string) *CreateFlowRequest {
	s.Namespace = &v
	return s
}

func (s *CreateFlowRequest) SetParentCategory(v string) *CreateFlowRequest {
	s.ParentCategory = &v
	return s
}

func (s *CreateFlowRequest) SetParentFlowList(v string) *CreateFlowRequest {
	s.ParentFlowList = &v
	return s
}

func (s *CreateFlowRequest) SetProjectId(v string) *CreateFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowRequest) SetRegionId(v string) *CreateFlowRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowRequest) SetStartSchedule(v int64) *CreateFlowRequest {
	s.StartSchedule = &v
	return s
}

type CreateFlowResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBody) SetId(v string) *CreateFlowResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowResponseBody) SetRequestId(v string) *CreateFlowResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowResponse) SetHeaders(v map[string]*string) *CreateFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowResponse) SetBody(v *CreateFlowResponseBody) *CreateFlowResponse {
	s.Body = v
	return s
}

type CreateFlowCategoryRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFlowCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowCategoryRequest) SetClientToken(v string) *CreateFlowCategoryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetName(v string) *CreateFlowCategoryRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetParentId(v string) *CreateFlowCategoryRequest {
	s.ParentId = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetProjectId(v string) *CreateFlowCategoryRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetRegionId(v string) *CreateFlowCategoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetType(v string) *CreateFlowCategoryRequest {
	s.Type = &v
	return s
}

type CreateFlowCategoryResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowCategoryResponseBody) SetId(v string) *CreateFlowCategoryResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowCategoryResponseBody) SetRequestId(v string) *CreateFlowCategoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowCategoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowCategoryResponse) SetHeaders(v map[string]*string) *CreateFlowCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowCategoryResponse) SetBody(v *CreateFlowCategoryResponseBody) *CreateFlowCategoryResponse {
	s.Body = v
	return s
}

type CreateFlowEditLockRequest struct {
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Force           *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateFlowEditLockRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowEditLockRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowEditLockRequest) SetEntityId(v string) *CreateFlowEditLockRequest {
	s.EntityId = &v
	return s
}

func (s *CreateFlowEditLockRequest) SetForce(v bool) *CreateFlowEditLockRequest {
	s.Force = &v
	return s
}

func (s *CreateFlowEditLockRequest) SetRegionId(v string) *CreateFlowEditLockRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowEditLockRequest) SetResourceOwnerId(v int64) *CreateFlowEditLockRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateFlowEditLockResponseBody struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateFlowEditLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowEditLockResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowEditLockResponseBody) SetBizId(v string) *CreateFlowEditLockResponseBody {
	s.BizId = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetEntityId(v string) *CreateFlowEditLockResponseBody {
	s.EntityId = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetOwnerId(v string) *CreateFlowEditLockResponseBody {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetRequestId(v string) *CreateFlowEditLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetStatus(v string) *CreateFlowEditLockResponseBody {
	s.Status = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetUserId(v string) *CreateFlowEditLockResponseBody {
	s.UserId = &v
	return s
}

type CreateFlowEditLockResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowEditLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowEditLockResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowEditLockResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowEditLockResponse) SetHeaders(v map[string]*string) *CreateFlowEditLockResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowEditLockResponse) SetBody(v *CreateFlowEditLockResponseBody) *CreateFlowEditLockResponse {
	s.Body = v
	return s
}

type CreateFlowJobRequest struct {
	Adhoc           *bool                               `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	AlertConf       *string                             `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	ClientToken     *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId       *string                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CustomVariables *string                             `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	Description     *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvConf         *string                             `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	FailAct         *string                             `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	Mode            *string                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
	MonitorConf     *string                             `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	Name            *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamConf       *string                             `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	Params          *string                             `json:"Params,omitempty" xml:"Params,omitempty"`
	ParentCategory  *string                             `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	ProjectId       *string                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceList    []*CreateFlowJobRequestResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	RetryPolicy     *string                             `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	RunConf         *string                             `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	Type            *string                             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowJobRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowJobRequest) SetAdhoc(v bool) *CreateFlowJobRequest {
	s.Adhoc = &v
	return s
}

func (s *CreateFlowJobRequest) SetAlertConf(v string) *CreateFlowJobRequest {
	s.AlertConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetClientToken(v string) *CreateFlowJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowJobRequest) SetClusterId(v string) *CreateFlowJobRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateFlowJobRequest) SetCustomVariables(v string) *CreateFlowJobRequest {
	s.CustomVariables = &v
	return s
}

func (s *CreateFlowJobRequest) SetDescription(v string) *CreateFlowJobRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowJobRequest) SetEnvConf(v string) *CreateFlowJobRequest {
	s.EnvConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetFailAct(v string) *CreateFlowJobRequest {
	s.FailAct = &v
	return s
}

func (s *CreateFlowJobRequest) SetMode(v string) *CreateFlowJobRequest {
	s.Mode = &v
	return s
}

func (s *CreateFlowJobRequest) SetMonitorConf(v string) *CreateFlowJobRequest {
	s.MonitorConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetName(v string) *CreateFlowJobRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowJobRequest) SetParamConf(v string) *CreateFlowJobRequest {
	s.ParamConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetParams(v string) *CreateFlowJobRequest {
	s.Params = &v
	return s
}

func (s *CreateFlowJobRequest) SetParentCategory(v string) *CreateFlowJobRequest {
	s.ParentCategory = &v
	return s
}

func (s *CreateFlowJobRequest) SetProjectId(v string) *CreateFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowJobRequest) SetRegionId(v string) *CreateFlowJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowJobRequest) SetResourceList(v []*CreateFlowJobRequestResourceList) *CreateFlowJobRequest {
	s.ResourceList = v
	return s
}

func (s *CreateFlowJobRequest) SetRetryPolicy(v string) *CreateFlowJobRequest {
	s.RetryPolicy = &v
	return s
}

func (s *CreateFlowJobRequest) SetRunConf(v string) *CreateFlowJobRequest {
	s.RunConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetType(v string) *CreateFlowJobRequest {
	s.Type = &v
	return s
}

type CreateFlowJobRequestResourceList struct {
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateFlowJobRequestResourceList) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowJobRequestResourceList) GoString() string {
	return s.String()
}

func (s *CreateFlowJobRequestResourceList) SetAlias(v string) *CreateFlowJobRequestResourceList {
	s.Alias = &v
	return s
}

func (s *CreateFlowJobRequestResourceList) SetPath(v string) *CreateFlowJobRequestResourceList {
	s.Path = &v
	return s
}

type CreateFlowJobResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowJobResponseBody) SetId(v string) *CreateFlowJobResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowJobResponseBody) SetRequestId(v string) *CreateFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowJobResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowJobResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowJobResponse) SetHeaders(v map[string]*string) *CreateFlowJobResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowJobResponse) SetBody(v *CreateFlowJobResponseBody) *CreateFlowJobResponse {
	s.Body = v
	return s
}

type CreateFlowProjectRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductType     *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateFlowProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectRequest) SetClientToken(v string) *CreateFlowProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowProjectRequest) SetDescription(v string) *CreateFlowProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowProjectRequest) SetName(v string) *CreateFlowProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowProjectRequest) SetProductType(v string) *CreateFlowProjectRequest {
	s.ProductType = &v
	return s
}

func (s *CreateFlowProjectRequest) SetRegionId(v string) *CreateFlowProjectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowProjectRequest) SetResourceGroupId(v string) *CreateFlowProjectRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateFlowProjectResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectResponseBody) SetId(v string) *CreateFlowProjectResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowProjectResponseBody) SetRequestId(v string) *CreateFlowProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectResponse) SetHeaders(v map[string]*string) *CreateFlowProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowProjectResponse) SetBody(v *CreateFlowProjectResponseBody) *CreateFlowProjectResponse {
	s.Body = v
	return s
}

type CreateFlowProjectClusterSettingRequest struct {
	ClientToken  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId    *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DefaultQueue *string   `json:"DefaultQueue,omitempty" xml:"DefaultQueue,omitempty"`
	DefaultUser  *string   `json:"DefaultUser,omitempty" xml:"DefaultUser,omitempty"`
	HostList     []*string `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Repeated"`
	ProjectId    *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QueueList    []*string `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserList     []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s CreateFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectClusterSettingRequest) SetClientToken(v string) *CreateFlowProjectClusterSettingRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetClusterId(v string) *CreateFlowProjectClusterSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetDefaultQueue(v string) *CreateFlowProjectClusterSettingRequest {
	s.DefaultQueue = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetDefaultUser(v string) *CreateFlowProjectClusterSettingRequest {
	s.DefaultUser = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetHostList(v []*string) *CreateFlowProjectClusterSettingRequest {
	s.HostList = v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetProjectId(v string) *CreateFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetQueueList(v []*string) *CreateFlowProjectClusterSettingRequest {
	s.QueueList = v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetRegionId(v string) *CreateFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetUserList(v []*string) *CreateFlowProjectClusterSettingRequest {
	s.UserList = v
	return s
}

type CreateFlowProjectClusterSettingResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectClusterSettingResponseBody) SetData(v bool) *CreateFlowProjectClusterSettingResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFlowProjectClusterSettingResponseBody) SetRequestId(v string) *CreateFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *CreateFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowProjectClusterSettingResponse) SetBody(v *CreateFlowProjectClusterSettingResponseBody) *CreateFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type CreateFlowProjectUserRequest struct {
	ClientToken *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ProjectId   *string                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId    *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	User        []*CreateFlowProjectUserRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s CreateFlowProjectUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectUserRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectUserRequest) SetClientToken(v string) *CreateFlowProjectUserRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowProjectUserRequest) SetProjectId(v string) *CreateFlowProjectUserRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowProjectUserRequest) SetRegionId(v string) *CreateFlowProjectUserRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowProjectUserRequest) SetUser(v []*CreateFlowProjectUserRequestUser) *CreateFlowProjectUserRequest {
	s.User = v
	return s
}

type CreateFlowProjectUserRequestUser struct {
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateFlowProjectUserRequestUser) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectUserRequestUser) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectUserRequestUser) SetUserId(v string) *CreateFlowProjectUserRequestUser {
	s.UserId = &v
	return s
}

func (s *CreateFlowProjectUserRequestUser) SetUserName(v string) *CreateFlowProjectUserRequestUser {
	s.UserName = &v
	return s
}

type CreateFlowProjectUserResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowProjectUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectUserResponseBody) SetData(v bool) *CreateFlowProjectUserResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFlowProjectUserResponseBody) SetRequestId(v string) *CreateFlowProjectUserResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowProjectUserResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowProjectUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowProjectUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectUserResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectUserResponse) SetHeaders(v map[string]*string) *CreateFlowProjectUserResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowProjectUserResponse) SetBody(v *CreateFlowProjectUserResponseBody) *CreateFlowProjectUserResponse {
	s.Body = v
	return s
}

type DeleteFlowRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) SetId(v string) *DeleteFlowRequest {
	s.Id = &v
	return s
}

func (s *DeleteFlowRequest) SetProjectId(v string) *DeleteFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFlowRequest) SetRegionId(v string) *DeleteFlowRequest {
	s.RegionId = &v
	return s
}

type DeleteFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponseBody) SetData(v bool) *DeleteFlowResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowResponseBody) SetRequestId(v string) *DeleteFlowResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponse) SetHeaders(v map[string]*string) *DeleteFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowResponse) SetBody(v *DeleteFlowResponseBody) *DeleteFlowResponse {
	s.Body = v
	return s
}

type DeleteFlowCategoryRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFlowCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowCategoryRequest) SetId(v string) *DeleteFlowCategoryRequest {
	s.Id = &v
	return s
}

func (s *DeleteFlowCategoryRequest) SetProjectId(v string) *DeleteFlowCategoryRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFlowCategoryRequest) SetRegionId(v string) *DeleteFlowCategoryRequest {
	s.RegionId = &v
	return s
}

type DeleteFlowCategoryResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowCategoryResponseBody) SetData(v bool) *DeleteFlowCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowCategoryResponseBody) SetRequestId(v string) *DeleteFlowCategoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowCategoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowCategoryResponse) SetHeaders(v map[string]*string) *DeleteFlowCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowCategoryResponse) SetBody(v *DeleteFlowCategoryResponseBody) *DeleteFlowCategoryResponse {
	s.Body = v
	return s
}

type DeleteFlowEditLockRequest struct {
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteFlowEditLockRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowEditLockRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowEditLockRequest) SetEntityId(v string) *DeleteFlowEditLockRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteFlowEditLockRequest) SetRegionId(v string) *DeleteFlowEditLockRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowEditLockRequest) SetResourceOwnerId(v int64) *DeleteFlowEditLockRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteFlowEditLockResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowEditLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowEditLockResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowEditLockResponseBody) SetData(v bool) *DeleteFlowEditLockResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowEditLockResponseBody) SetRequestId(v string) *DeleteFlowEditLockResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowEditLockResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowEditLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowEditLockResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowEditLockResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowEditLockResponse) SetHeaders(v map[string]*string) *DeleteFlowEditLockResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowEditLockResponse) SetBody(v *DeleteFlowEditLockResponseBody) *DeleteFlowEditLockResponse {
	s.Body = v
	return s
}

type DeleteFlowProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFlowProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectRequest) SetProjectId(v string) *DeleteFlowProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFlowProjectRequest) SetRegionId(v string) *DeleteFlowProjectRequest {
	s.RegionId = &v
	return s
}

type DeleteFlowProjectResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectResponseBody) SetData(v bool) *DeleteFlowProjectResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowProjectResponseBody) SetRequestId(v string) *DeleteFlowProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectResponse) SetHeaders(v map[string]*string) *DeleteFlowProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowProjectResponse) SetBody(v *DeleteFlowProjectResponseBody) *DeleteFlowProjectResponse {
	s.Body = v
	return s
}

type DeleteFlowProjectClusterSettingRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectClusterSettingRequest) SetClusterId(v string) *DeleteFlowProjectClusterSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteFlowProjectClusterSettingRequest) SetProjectId(v string) *DeleteFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFlowProjectClusterSettingRequest) SetRegionId(v string) *DeleteFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

type DeleteFlowProjectClusterSettingResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectClusterSettingResponseBody) SetData(v bool) *DeleteFlowProjectClusterSettingResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowProjectClusterSettingResponseBody) SetRequestId(v string) *DeleteFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *DeleteFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowProjectClusterSettingResponse) SetBody(v *DeleteFlowProjectClusterSettingResponseBody) *DeleteFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type DeleteFlowProjectUserRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteFlowProjectUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectUserRequest) SetProjectId(v string) *DeleteFlowProjectUserRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFlowProjectUserRequest) SetRegionId(v string) *DeleteFlowProjectUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowProjectUserRequest) SetUserName(v string) *DeleteFlowProjectUserRequest {
	s.UserName = &v
	return s
}

type DeleteFlowProjectUserResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowProjectUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectUserResponseBody) SetData(v bool) *DeleteFlowProjectUserResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowProjectUserResponseBody) SetRequestId(v string) *DeleteFlowProjectUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowProjectUserResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowProjectUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowProjectUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectUserResponse) SetHeaders(v map[string]*string) *DeleteFlowProjectUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowProjectUserResponse) SetBody(v *DeleteFlowProjectUserResponseBody) *DeleteFlowProjectUserResponse {
	s.Body = v
	return s
}

type DescribeClusterV2Request struct {
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeClusterV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2Request) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2Request) SetId(v string) *DescribeClusterV2Request {
	s.Id = &v
	return s
}

func (s *DescribeClusterV2Request) SetRegionId(v string) *DescribeClusterV2Request {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterV2Request) SetResourceOwnerId(v int64) *DescribeClusterV2Request {
	s.ResourceOwnerId = &v
	return s
}

type DescribeClusterV2ResponseBody struct {
	ClusterInfo *DescribeClusterV2ResponseBodyClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBody) SetClusterInfo(v *DescribeClusterV2ResponseBodyClusterInfo) *DescribeClusterV2ResponseBody {
	s.ClusterInfo = v
	return s
}

func (s *DescribeClusterV2ResponseBody) SetRequestId(v string) *DescribeClusterV2ResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfo struct {
	AccessInfo                      *DescribeClusterV2ResponseBodyClusterInfoAccessInfo             `json:"AccessInfo,omitempty" xml:"AccessInfo,omitempty" type:"Struct"`
	AutoScalingAllowed              *bool                                                           `json:"AutoScalingAllowed,omitempty" xml:"AutoScalingAllowed,omitempty"`
	AutoScalingByLoadAllowed        *bool                                                           `json:"AutoScalingByLoadAllowed,omitempty" xml:"AutoScalingByLoadAllowed,omitempty"`
	AutoScalingEnable               *bool                                                           `json:"AutoScalingEnable,omitempty" xml:"AutoScalingEnable,omitempty"`
	AutoScalingSpotWithLimitAllowed *bool                                                           `json:"AutoScalingSpotWithLimitAllowed,omitempty" xml:"AutoScalingSpotWithLimitAllowed,omitempty"`
	AutoScalingVersion              *string                                                         `json:"AutoScalingVersion,omitempty" xml:"AutoScalingVersion,omitempty"`
	AutoScalingWithGraceAllowed     *bool                                                           `json:"AutoScalingWithGraceAllowed,omitempty" xml:"AutoScalingWithGraceAllowed,omitempty"`
	BootstrapActionList             *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList    `json:"BootstrapActionList,omitempty" xml:"BootstrapActionList,omitempty" type:"Struct"`
	BootstrapFailed                 *bool                                                           `json:"BootstrapFailed,omitempty" xml:"BootstrapFailed,omitempty"`
	ChargeType                      *string                                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Configurations                  *string                                                         `json:"Configurations,omitempty" xml:"Configurations,omitempty"`
	CoreNodeInService               *int32                                                          `json:"CoreNodeInService,omitempty" xml:"CoreNodeInService,omitempty"`
	CoreNodeTotal                   *int32                                                          `json:"CoreNodeTotal,omitempty" xml:"CoreNodeTotal,omitempty"`
	CreateResource                  *string                                                         `json:"CreateResource,omitempty" xml:"CreateResource,omitempty"`
	CreateType                      *string                                                         `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	DepositType                     *string                                                         `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	EasEnable                       *bool                                                           `json:"EasEnable,omitempty" xml:"EasEnable,omitempty"`
	ExpiredTime                     *int64                                                          `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ExtraInfo                       *string                                                         `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	FailReason                      *DescribeClusterV2ResponseBodyClusterInfoFailReason             `json:"FailReason,omitempty" xml:"FailReason,omitempty" type:"Struct"`
	GatewayClusterIds               *string                                                         `json:"GatewayClusterIds,omitempty" xml:"GatewayClusterIds,omitempty"`
	GatewayClusterInfoList          *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList `json:"GatewayClusterInfoList,omitempty" xml:"GatewayClusterInfoList,omitempty" type:"Struct"`
	HighAvailabilityEnable          *bool                                                           `json:"HighAvailabilityEnable,omitempty" xml:"HighAvailabilityEnable,omitempty"`
	HostGroupList                   *DescribeClusterV2ResponseBodyClusterInfoHostGroupList          `json:"HostGroupList,omitempty" xml:"HostGroupList,omitempty" type:"Struct"`
	HostPoolInfo                    *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo           `json:"HostPoolInfo,omitempty" xml:"HostPoolInfo,omitempty" type:"Struct"`
	Id                              *string                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId                         *string                                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceGeneration              *string                                                         `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	IoOptimized                     *bool                                                           `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	K8sClusterId                    *string                                                         `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	LocalMetaDb                     *bool                                                           `json:"LocalMetaDb,omitempty" xml:"LocalMetaDb,omitempty"`
	LogEnable                       *bool                                                           `json:"LogEnable,omitempty" xml:"LogEnable,omitempty"`
	LogPath                         *string                                                         `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	MachineType                     *string                                                         `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	MasterNodeInService             *int32                                                          `json:"MasterNodeInService,omitempty" xml:"MasterNodeInService,omitempty"`
	MasterNodeTotal                 *int32                                                          `json:"MasterNodeTotal,omitempty" xml:"MasterNodeTotal,omitempty"`
	MetaStoreType                   *string                                                         `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	Name                            *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	NetType                         *string                                                         `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Period                          *int32                                                          `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId                        *string                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelateClusterId                 *string                                                         `json:"RelateClusterId,omitempty" xml:"RelateClusterId,omitempty"`
	RelateClusterInfo               *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo      `json:"RelateClusterInfo,omitempty" xml:"RelateClusterInfo,omitempty" type:"Struct"`
	ResizeDiskEnable                *bool                                                           `json:"ResizeDiskEnable,omitempty" xml:"ResizeDiskEnable,omitempty"`
	RunningTime                     *int32                                                          `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SecurityGroupId                 *string                                                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName               *string                                                         `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	ShowSoftwareInterface           *bool                                                           `json:"ShowSoftwareInterface,omitempty" xml:"ShowSoftwareInterface,omitempty"`
	SoftwareInfo                    *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo           `json:"SoftwareInfo,omitempty" xml:"SoftwareInfo,omitempty" type:"Struct"`
	StartTime                       *int64                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                          *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	StopTime                        *int64                                                          `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	TaskNodeInService               *int32                                                          `json:"TaskNodeInService,omitempty" xml:"TaskNodeInService,omitempty"`
	TaskNodeTotal                   *int32                                                          `json:"TaskNodeTotal,omitempty" xml:"TaskNodeTotal,omitempty"`
	UserDefinedEmrEcsRole           *string                                                         `json:"UserDefinedEmrEcsRole,omitempty" xml:"UserDefinedEmrEcsRole,omitempty"`
	UserId                          *string                                                         `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VSwitchId                       *string                                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                           *string                                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                          *string                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAccessInfo(v *DescribeClusterV2ResponseBodyClusterInfoAccessInfo) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AccessInfo = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingAllowed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingAllowed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingByLoadAllowed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingByLoadAllowed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingSpotWithLimitAllowed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingSpotWithLimitAllowed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingVersion(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingVersion = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingWithGraceAllowed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingWithGraceAllowed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetBootstrapActionList(v *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList) *DescribeClusterV2ResponseBodyClusterInfo {
	s.BootstrapActionList = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetBootstrapFailed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.BootstrapFailed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetChargeType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ChargeType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetConfigurations(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Configurations = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetCoreNodeInService(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.CoreNodeInService = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetCoreNodeTotal(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.CoreNodeTotal = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetCreateResource(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.CreateResource = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetCreateType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.CreateType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetDepositType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.DepositType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetEasEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.EasEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetExpiredTime(v int64) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetExtraInfo(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ExtraInfo = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetFailReason(v *DescribeClusterV2ResponseBodyClusterInfoFailReason) *DescribeClusterV2ResponseBodyClusterInfo {
	s.FailReason = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetGatewayClusterIds(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.GatewayClusterIds = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetGatewayClusterInfoList(v *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList) *DescribeClusterV2ResponseBodyClusterInfo {
	s.GatewayClusterInfoList = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetHighAvailabilityEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.HighAvailabilityEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetHostGroupList(v *DescribeClusterV2ResponseBodyClusterInfoHostGroupList) *DescribeClusterV2ResponseBodyClusterInfo {
	s.HostGroupList = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetHostPoolInfo(v *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) *DescribeClusterV2ResponseBodyClusterInfo {
	s.HostPoolInfo = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Id = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetImageId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetInstanceGeneration(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.InstanceGeneration = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetIoOptimized(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.IoOptimized = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetK8sClusterId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetLocalMetaDb(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.LocalMetaDb = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetLogEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.LogEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetLogPath(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.LogPath = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetMachineType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.MachineType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetMasterNodeInService(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.MasterNodeInService = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetMasterNodeTotal(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.MasterNodeTotal = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetMetaStoreType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.MetaStoreType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetName(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetNetType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetPeriod(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Period = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetRegionId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetRelateClusterId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.RelateClusterId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetRelateClusterInfo(v *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) *DescribeClusterV2ResponseBodyClusterInfo {
	s.RelateClusterInfo = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetResizeDiskEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ResizeDiskEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetRunningTime(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.RunningTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetSecurityGroupId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetSecurityGroupName(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetShowSoftwareInterface(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ShowSoftwareInterface = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetSoftwareInfo(v *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) *DescribeClusterV2ResponseBodyClusterInfo {
	s.SoftwareInfo = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetStartTime(v int64) *DescribeClusterV2ResponseBodyClusterInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetStatus(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetStopTime(v int64) *DescribeClusterV2ResponseBodyClusterInfo {
	s.StopTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetTaskNodeInService(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.TaskNodeInService = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetTaskNodeTotal(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.TaskNodeTotal = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetUserDefinedEmrEcsRole(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.UserDefinedEmrEcsRole = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetUserId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.UserId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetVSwitchId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetVpcId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetZoneId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ZoneId = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoAccessInfo struct {
	ZKLinks *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks `json:"ZKLinks,omitempty" xml:"ZKLinks,omitempty" type:"Struct"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoAccessInfo) SetZKLinks(v *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks) *DescribeClusterV2ResponseBodyClusterInfoAccessInfo {
	s.ZKLinks = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks struct {
	ZKLink []*DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink `json:"ZKLink,omitempty" xml:"ZKLink,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks) SetZKLink(v []*DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks {
	s.ZKLink = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink struct {
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) SetLink(v string) *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink {
	s.Link = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) SetPort(v string) *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink {
	s.Port = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList struct {
	BootstrapAction []*DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction `json:"BootstrapAction,omitempty" xml:"BootstrapAction,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList) SetBootstrapAction(v []*DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList {
	s.BootstrapAction = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction struct {
	Arg  *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) SetArg(v string) *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction {
	s.Arg = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) SetName(v string) *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction {
	s.Name = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) SetPath(v string) *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction {
	s.Path = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoFailReason struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoFailReason) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoFailReason) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoFailReason) SetErrorCode(v string) *DescribeClusterV2ResponseBodyClusterInfoFailReason {
	s.ErrorCode = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoFailReason) SetErrorMsg(v string) *DescribeClusterV2ResponseBodyClusterInfoFailReason {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoFailReason) SetRequestId(v string) *DescribeClusterV2ResponseBodyClusterInfoFailReason {
	s.RequestId = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList struct {
	GatewayClusterInfo []*DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo `json:"GatewayClusterInfo,omitempty" xml:"GatewayClusterInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList) SetGatewayClusterInfo(v []*DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList {
	s.GatewayClusterInfo = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) SetClusterId(v string) *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) SetClusterName(v string) *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) SetStatus(v string) *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo {
	s.Status = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupList struct {
	HostGroup []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupList) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupList) SetHostGroup(v []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) *DescribeClusterV2ResponseBodyClusterInfoHostGroupList {
	s.HostGroup = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup struct {
	BandWidth             *string                                                              `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	ChargeType            *string                                                              `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CpuCore               *int32                                                               `json:"CpuCore,omitempty" xml:"CpuCore,omitempty"`
	DiskCapacity          *int32                                                               `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	DiskCount             *int32                                                               `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	DiskType              *string                                                              `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HostGroupChangeStatus *string                                                              `json:"HostGroupChangeStatus,omitempty" xml:"HostGroupChangeStatus,omitempty"`
	HostGroupChangeType   *string                                                              `json:"HostGroupChangeType,omitempty" xml:"HostGroupChangeType,omitempty"`
	HostGroupId           *string                                                              `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName         *string                                                              `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	HostGroupSubType      *string                                                              `json:"HostGroupSubType,omitempty" xml:"HostGroupSubType,omitempty"`
	HostGroupType         *string                                                              `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	InstanceType          *string                                                              `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LockReason            *string                                                              `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	LockType              *string                                                              `json:"LockType,omitempty" xml:"LockType,omitempty"`
	MemoryCapacity        *int32                                                               `json:"MemoryCapacity,omitempty" xml:"MemoryCapacity,omitempty"`
	NodeCount             *int32                                                               `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	Nodes                 *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	Period                *string                                                              `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetBandWidth(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.BandWidth = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetChargeType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.ChargeType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetCpuCore(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.CpuCore = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetDiskCapacity(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.DiskCapacity = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetDiskCount(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.DiskCount = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetDiskType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.DiskType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupChangeStatus(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupChangeStatus = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupChangeType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupChangeType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupName(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupSubType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupSubType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetInstanceType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetLockReason(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.LockReason = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetLockType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.LockType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetMemoryCapacity(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.MemoryCapacity = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetNodeCount(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.NodeCount = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetNodes(v *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.Nodes = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetPeriod(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.Period = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes struct {
	Node []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes) SetNode(v []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes {
	s.Node = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode struct {
	CreateTime     *string                                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DaemonInfos    *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos `json:"DaemonInfos,omitempty" xml:"DaemonInfos,omitempty" type:"Struct"`
	DiskInfos      *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos   `json:"DiskInfos,omitempty" xml:"DiskInfos,omitempty" type:"Struct"`
	EmrExpiredTime *string                                                                             `json:"EmrExpiredTime,omitempty" xml:"EmrExpiredTime,omitempty"`
	ExpiredTime    *string                                                                             `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InnerIp        *string                                                                             `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	InstanceId     *string                                                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PubIp          *string                                                                             `json:"PubIp,omitempty" xml:"PubIp,omitempty"`
	Status         *string                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportIpV6    *bool                                                                               `json:"SupportIpV6,omitempty" xml:"SupportIpV6,omitempty"`
	ZoneId         *string                                                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetCreateTime(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetDaemonInfos(v *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.DaemonInfos = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetDiskInfos(v *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.DiskInfos = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetEmrExpiredTime(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.EmrExpiredTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetExpiredTime(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetInnerIp(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.InnerIp = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetInstanceId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetPubIp(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.PubIp = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetStatus(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.Status = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetSupportIpV6(v bool) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.SupportIpV6 = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetZoneId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.ZoneId = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos struct {
	DaemonInfo []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo `json:"DaemonInfo,omitempty" xml:"DaemonInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos) SetDaemonInfo(v []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos {
	s.DaemonInfo = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo) SetName(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo {
	s.Name = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos struct {
	DiskInfo []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos) SetDiskInfo(v []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos {
	s.DiskInfo = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo struct {
	Device   *string `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetDevice(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.Device = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetDiskId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.DiskId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetDiskName(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.DiskName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetSize(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.Size = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.Type = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo struct {
	HpBizId *string `json:"HpBizId,omitempty" xml:"HpBizId,omitempty"`
	HpName  *string `json:"HpName,omitempty" xml:"HpName,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) SetHpBizId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo {
	s.HpBizId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) SetHpName(v string) *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo {
	s.HpName = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) SetClusterId(v string) *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) SetClusterName(v string) *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) SetStatus(v string) *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo {
	s.Status = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo struct {
	ClusterType *string                                                        `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	EmrVer      *string                                                        `json:"EmrVer,omitempty" xml:"EmrVer,omitempty"`
	Softwares   *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Struct"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) SetClusterType(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) SetEmrVer(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo {
	s.EmrVer = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) SetSoftwares(v *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo {
	s.Softwares = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares struct {
	Software []*DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware `json:"Software,omitempty" xml:"Software,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares) SetSoftware(v []*DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares {
	s.Software = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OnlyDisplay *bool   `json:"OnlyDisplay,omitempty" xml:"OnlyDisplay,omitempty"`
	StartTpe    *int32  `json:"StartTpe,omitempty" xml:"StartTpe,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetDisplayName(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.DisplayName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetName(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.Name = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetOnlyDisplay(v bool) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.OnlyDisplay = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetStartTpe(v int32) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.StartTpe = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetVersion(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.Version = &v
	return s
}

type DescribeClusterV2Response struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2Response) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2Response) SetHeaders(v map[string]*string) *DescribeClusterV2Response {
	s.Headers = v
	return s
}

func (s *DescribeClusterV2Response) SetBody(v *DescribeClusterV2ResponseBody) *DescribeClusterV2Response {
	s.Body = v
	return s
}

type DescribeFlowRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowRequest) SetId(v string) *DescribeFlowRequest {
	s.Id = &v
	return s
}

func (s *DescribeFlowRequest) SetProjectId(v string) *DescribeFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowRequest) SetRegionId(v string) *DescribeFlowRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowResponseBody struct {
	AlertConf               *string                                 `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	AlertDingDingGroupBizId *string                                 `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	AlertUserGroupBizId     *string                                 `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	Application             *string                                 `json:"Application,omitempty" xml:"Application,omitempty"`
	CategoryId              *string                                 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ClusterId               *string                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateCluster           *bool                                   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	CronExpr                *string                                 `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	Description             *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	EditLockDetail          *string                                 `json:"EditLockDetail,omitempty" xml:"EditLockDetail,omitempty"`
	EndSchedule             *int64                                  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	GmtCreate               *int64                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified             *int64                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Graph                   *string                                 `json:"Graph,omitempty" xml:"Graph,omitempty"`
	HostName                *string                                 `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id                      *string                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                    *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace               *string                                 `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParentFlowList          *DescribeFlowResponseBodyParentFlowList `json:"ParentFlowList,omitempty" xml:"ParentFlowList,omitempty" type:"Struct"`
	Periodic                *bool                                   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	RequestId               *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartSchedule           *int64                                  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	Status                  *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                    *string                                 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBody) SetAlertConf(v string) *DescribeFlowResponseBody {
	s.AlertConf = &v
	return s
}

func (s *DescribeFlowResponseBody) SetAlertDingDingGroupBizId(v string) *DescribeFlowResponseBody {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetAlertUserGroupBizId(v string) *DescribeFlowResponseBody {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetApplication(v string) *DescribeFlowResponseBody {
	s.Application = &v
	return s
}

func (s *DescribeFlowResponseBody) SetCategoryId(v string) *DescribeFlowResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetClusterId(v string) *DescribeFlowResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetCreateCluster(v bool) *DescribeFlowResponseBody {
	s.CreateCluster = &v
	return s
}

func (s *DescribeFlowResponseBody) SetCronExpr(v string) *DescribeFlowResponseBody {
	s.CronExpr = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDescription(v string) *DescribeFlowResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowResponseBody) SetEditLockDetail(v string) *DescribeFlowResponseBody {
	s.EditLockDetail = &v
	return s
}

func (s *DescribeFlowResponseBody) SetEndSchedule(v int64) *DescribeFlowResponseBody {
	s.EndSchedule = &v
	return s
}

func (s *DescribeFlowResponseBody) SetGmtCreate(v int64) *DescribeFlowResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowResponseBody) SetGmtModified(v int64) *DescribeFlowResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowResponseBody) SetGraph(v string) *DescribeFlowResponseBody {
	s.Graph = &v
	return s
}

func (s *DescribeFlowResponseBody) SetHostName(v string) *DescribeFlowResponseBody {
	s.HostName = &v
	return s
}

func (s *DescribeFlowResponseBody) SetId(v string) *DescribeFlowResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowResponseBody) SetName(v string) *DescribeFlowResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFlowResponseBody) SetNamespace(v string) *DescribeFlowResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeFlowResponseBody) SetParentFlowList(v *DescribeFlowResponseBodyParentFlowList) *DescribeFlowResponseBody {
	s.ParentFlowList = v
	return s
}

func (s *DescribeFlowResponseBody) SetPeriodic(v bool) *DescribeFlowResponseBody {
	s.Periodic = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRequestId(v string) *DescribeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetStartSchedule(v int64) *DescribeFlowResponseBody {
	s.StartSchedule = &v
	return s
}

func (s *DescribeFlowResponseBody) SetStatus(v string) *DescribeFlowResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFlowResponseBody) SetType(v string) *DescribeFlowResponseBody {
	s.Type = &v
	return s
}

type DescribeFlowResponseBodyParentFlowList struct {
	ParentFlow []*DescribeFlowResponseBodyParentFlowListParentFlow `json:"ParentFlow,omitempty" xml:"ParentFlow,omitempty" type:"Repeated"`
}

func (s DescribeFlowResponseBodyParentFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponseBodyParentFlowList) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBodyParentFlowList) SetParentFlow(v []*DescribeFlowResponseBodyParentFlowListParentFlow) *DescribeFlowResponseBodyParentFlowList {
	s.ParentFlow = v
	return s
}

type DescribeFlowResponseBodyParentFlowListParentFlow struct {
	ParentFlowId   *string `json:"ParentFlowId,omitempty" xml:"ParentFlowId,omitempty"`
	ParentFlowName *string `json:"ParentFlowName,omitempty" xml:"ParentFlowName,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName    *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DescribeFlowResponseBodyParentFlowListParentFlow) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponseBodyParentFlowListParentFlow) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBodyParentFlowListParentFlow) SetParentFlowId(v string) *DescribeFlowResponseBodyParentFlowListParentFlow {
	s.ParentFlowId = &v
	return s
}

func (s *DescribeFlowResponseBodyParentFlowListParentFlow) SetParentFlowName(v string) *DescribeFlowResponseBodyParentFlowListParentFlow {
	s.ParentFlowName = &v
	return s
}

func (s *DescribeFlowResponseBodyParentFlowListParentFlow) SetProjectId(v string) *DescribeFlowResponseBodyParentFlowListParentFlow {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowResponseBodyParentFlowListParentFlow) SetProjectName(v string) *DescribeFlowResponseBodyParentFlowListParentFlow {
	s.ProjectName = &v
	return s
}

type DescribeFlowResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponse) SetHeaders(v map[string]*string) *DescribeFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowResponse) SetBody(v *DescribeFlowResponseBody) *DescribeFlowResponse {
	s.Body = v
	return s
}

type DescribeFlowCategoryTreeRequest struct {
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFlowCategoryTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowCategoryTreeRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowCategoryTreeRequest) SetCategoryId(v string) *DescribeFlowCategoryTreeRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetKeyword(v string) *DescribeFlowCategoryTreeRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetMode(v string) *DescribeFlowCategoryTreeRequest {
	s.Mode = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetProjectId(v string) *DescribeFlowCategoryTreeRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetRegionId(v string) *DescribeFlowCategoryTreeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetType(v string) *DescribeFlowCategoryTreeRequest {
	s.Type = &v
	return s
}

type DescribeFlowCategoryTreeResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowCategoryTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowCategoryTreeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowCategoryTreeResponseBody) SetData(v string) *DescribeFlowCategoryTreeResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeFlowCategoryTreeResponseBody) SetRequestId(v string) *DescribeFlowCategoryTreeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFlowCategoryTreeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowCategoryTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowCategoryTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowCategoryTreeResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowCategoryTreeResponse) SetHeaders(v map[string]*string) *DescribeFlowCategoryTreeResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowCategoryTreeResponse) SetBody(v *DescribeFlowCategoryTreeResponseBody) *DescribeFlowCategoryTreeResponse {
	s.Body = v
	return s
}

type DescribeFlowInstanceRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceRequest) SetId(v string) *DescribeFlowInstanceRequest {
	s.Id = &v
	return s
}

func (s *DescribeFlowInstanceRequest) SetProjectId(v string) *DescribeFlowInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowInstanceRequest) SetRegionId(v string) *DescribeFlowInstanceRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowInstanceResponseBody struct {
	ClusterId          *string                                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CronExpression     *string                                             `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	DependencyFlowList *DescribeFlowInstanceResponseBodyDependencyFlowList `json:"DependencyFlowList,omitempty" xml:"DependencyFlowList,omitempty" type:"Struct"`
	Duration           *int64                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTime            *int64                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FlowId             *string                                             `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName           *string                                             `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	GmtCreate          *int64                                              `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *int64                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Graph              *string                                             `json:"Graph,omitempty" xml:"Graph,omitempty"`
	HasNodeFailed      *bool                                               `json:"HasNodeFailed,omitempty" xml:"HasNodeFailed,omitempty"`
	Id                 *string                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Namespace          *string                                             `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NodeInstance       *DescribeFlowInstanceResponseBodyNodeInstance       `json:"NodeInstance,omitempty" xml:"NodeInstance,omitempty" type:"Struct"`
	ProjectId          *string                                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RequestId          *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScheduleTime       *int64                                              `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	StartTime          *int64                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status             *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFlowInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBody) SetClusterId(v string) *DescribeFlowInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetCronExpression(v string) *DescribeFlowInstanceResponseBody {
	s.CronExpression = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetDependencyFlowList(v *DescribeFlowInstanceResponseBodyDependencyFlowList) *DescribeFlowInstanceResponseBody {
	s.DependencyFlowList = v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetDuration(v int64) *DescribeFlowInstanceResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetEndTime(v int64) *DescribeFlowInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetFlowId(v string) *DescribeFlowInstanceResponseBody {
	s.FlowId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetFlowName(v string) *DescribeFlowInstanceResponseBody {
	s.FlowName = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetGmtCreate(v int64) *DescribeFlowInstanceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetGmtModified(v int64) *DescribeFlowInstanceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetGraph(v string) *DescribeFlowInstanceResponseBody {
	s.Graph = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetHasNodeFailed(v bool) *DescribeFlowInstanceResponseBody {
	s.HasNodeFailed = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetId(v string) *DescribeFlowInstanceResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetNamespace(v string) *DescribeFlowInstanceResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetNodeInstance(v *DescribeFlowInstanceResponseBodyNodeInstance) *DescribeFlowInstanceResponseBody {
	s.NodeInstance = v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetProjectId(v string) *DescribeFlowInstanceResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetRequestId(v string) *DescribeFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetScheduleTime(v int64) *DescribeFlowInstanceResponseBody {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetStartTime(v int64) *DescribeFlowInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetStatus(v string) *DescribeFlowInstanceResponseBody {
	s.Status = &v
	return s
}

type DescribeFlowInstanceResponseBodyDependencyFlowList struct {
	ParentFlow []*DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow `json:"ParentFlow,omitempty" xml:"ParentFlow,omitempty" type:"Repeated"`
}

func (s DescribeFlowInstanceResponseBodyDependencyFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBodyDependencyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowList) SetParentFlow(v []*DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) *DescribeFlowInstanceResponseBodyDependencyFlowList {
	s.ParentFlow = v
	return s
}

type DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow struct {
	BizDate              *int64  `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	DependencyFlowId     *string `json:"DependencyFlowId,omitempty" xml:"DependencyFlowId,omitempty"`
	DependencyInstanceId *string `json:"DependencyInstanceId,omitempty" xml:"DependencyInstanceId,omitempty"`
	FlowId               *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowInstanceId       *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	Meet                 *bool   `json:"Meet,omitempty" xml:"Meet,omitempty"`
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ScheduleKey          *string `json:"ScheduleKey,omitempty" xml:"ScheduleKey,omitempty"`
}

func (s DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetBizDate(v int64) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.BizDate = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetDependencyFlowId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.DependencyFlowId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetDependencyInstanceId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.DependencyInstanceId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetFlowId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.FlowId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetFlowInstanceId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.FlowInstanceId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetMeet(v bool) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.Meet = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetProjectId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetScheduleKey(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.ScheduleKey = &v
	return s
}

type DescribeFlowInstanceResponseBodyNodeInstance struct {
	NodeInstance []*DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance `json:"NodeInstance,omitempty" xml:"NodeInstance,omitempty" type:"Repeated"`
}

func (s DescribeFlowInstanceResponseBodyNodeInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBodyNodeInstance) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBodyNodeInstance) SetNodeInstance(v []*DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) *DescribeFlowInstanceResponseBodyNodeInstance {
	s.NodeInstance = v
	return s
}

type DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Duration       *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExternalId     *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	ExternalInfo   *string `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty"`
	ExternalStatus *string `json:"ExternalStatus,omitempty" xml:"ExternalStatus,omitempty"`
	FailAct        *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HostName       *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName        *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobType        *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MaxRetry       *string `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	NodeName       *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Pending        *bool   `json:"Pending,omitempty" xml:"Pending,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Retries        *int32  `json:"Retries,omitempty" xml:"Retries,omitempty"`
	RetryInterval  *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetClusterId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetDuration(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Duration = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetEndTime(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetExternalId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ExternalId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetExternalInfo(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ExternalInfo = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetExternalStatus(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ExternalStatus = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetFailAct(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.FailAct = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetGmtCreate(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetGmtModified(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetHostName(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.HostName = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Id = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetJobId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.JobId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetJobName(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.JobName = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetJobType(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.JobType = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetMaxRetry(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.MaxRetry = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetNodeName(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.NodeName = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetPending(v bool) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Pending = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetProjectId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetRetries(v int32) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Retries = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetRetryInterval(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.RetryInterval = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetStartTime(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetStatus(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Status = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetType(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Type = &v
	return s
}

type DescribeFlowInstanceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponse) SetHeaders(v map[string]*string) *DescribeFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowInstanceResponse) SetBody(v *DescribeFlowInstanceResponseBody) *DescribeFlowInstanceResponse {
	s.Body = v
	return s
}

type DescribeFlowJobRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobRequest) SetId(v string) *DescribeFlowJobRequest {
	s.Id = &v
	return s
}

func (s *DescribeFlowJobRequest) SetProjectId(v string) *DescribeFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowJobRequest) SetRegionId(v string) *DescribeFlowJobRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowJobResponseBody struct {
	Adhoc             *string                                  `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	AlertConf         *string                                  `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	CategoryId        *string                                  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CustomVariables   *string                                  `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	Description       *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EditLockDetail    *string                                  `json:"EditLockDetail,omitempty" xml:"EditLockDetail,omitempty"`
	EnvConf           *string                                  `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	FailAct           *string                                  `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	GmtCreate         *int64                                   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified       *int64                                   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                *string                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	KnoxPassword      *string                                  `json:"KnoxPassword,omitempty" xml:"KnoxPassword,omitempty"`
	KnoxUser          *string                                  `json:"KnoxUser,omitempty" xml:"KnoxUser,omitempty"`
	LastInstanceId    *string                                  `json:"LastInstanceId,omitempty" xml:"LastInstanceId,omitempty"`
	MaxRetry          *int32                                   `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	MaxRunningTimeSec *int64                                   `json:"MaxRunningTimeSec,omitempty" xml:"MaxRunningTimeSec,omitempty"`
	Mode              *string                                  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	MonitorConf       *string                                  `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	Name              *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamConf         *string                                  `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	Params            *string                                  `json:"Params,omitempty" xml:"Params,omitempty"`
	RequestId         *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceList      *DescribeFlowJobResponseBodyResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Struct"`
	RetryInterval     *int64                                   `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RetryPolicy       *string                                  `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	RunConf           *string                                  `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	Type              *string                                  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobResponseBody) SetAdhoc(v string) *DescribeFlowJobResponseBody {
	s.Adhoc = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetAlertConf(v string) *DescribeFlowJobResponseBody {
	s.AlertConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetCategoryId(v string) *DescribeFlowJobResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetCustomVariables(v string) *DescribeFlowJobResponseBody {
	s.CustomVariables = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetDescription(v string) *DescribeFlowJobResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetEditLockDetail(v string) *DescribeFlowJobResponseBody {
	s.EditLockDetail = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetEnvConf(v string) *DescribeFlowJobResponseBody {
	s.EnvConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetFailAct(v string) *DescribeFlowJobResponseBody {
	s.FailAct = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetGmtCreate(v int64) *DescribeFlowJobResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetGmtModified(v int64) *DescribeFlowJobResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetId(v string) *DescribeFlowJobResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetKnoxPassword(v string) *DescribeFlowJobResponseBody {
	s.KnoxPassword = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetKnoxUser(v string) *DescribeFlowJobResponseBody {
	s.KnoxUser = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetLastInstanceId(v string) *DescribeFlowJobResponseBody {
	s.LastInstanceId = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetMaxRetry(v int32) *DescribeFlowJobResponseBody {
	s.MaxRetry = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetMaxRunningTimeSec(v int64) *DescribeFlowJobResponseBody {
	s.MaxRunningTimeSec = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetMode(v string) *DescribeFlowJobResponseBody {
	s.Mode = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetMonitorConf(v string) *DescribeFlowJobResponseBody {
	s.MonitorConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetName(v string) *DescribeFlowJobResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetParamConf(v string) *DescribeFlowJobResponseBody {
	s.ParamConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetParams(v string) *DescribeFlowJobResponseBody {
	s.Params = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetRequestId(v string) *DescribeFlowJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetResourceList(v *DescribeFlowJobResponseBodyResourceList) *DescribeFlowJobResponseBody {
	s.ResourceList = v
	return s
}

func (s *DescribeFlowJobResponseBody) SetRetryInterval(v int64) *DescribeFlowJobResponseBody {
	s.RetryInterval = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetRetryPolicy(v string) *DescribeFlowJobResponseBody {
	s.RetryPolicy = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetRunConf(v string) *DescribeFlowJobResponseBody {
	s.RunConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetType(v string) *DescribeFlowJobResponseBody {
	s.Type = &v
	return s
}

type DescribeFlowJobResponseBodyResourceList struct {
	Resource []*DescribeFlowJobResponseBodyResourceListResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeFlowJobResponseBodyResourceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobResponseBodyResourceList) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobResponseBodyResourceList) SetResource(v []*DescribeFlowJobResponseBodyResourceListResource) *DescribeFlowJobResponseBodyResourceList {
	s.Resource = v
	return s
}

type DescribeFlowJobResponseBodyResourceListResource struct {
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeFlowJobResponseBodyResourceListResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobResponseBodyResourceListResource) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobResponseBodyResourceListResource) SetAlias(v string) *DescribeFlowJobResponseBodyResourceListResource {
	s.Alias = &v
	return s
}

func (s *DescribeFlowJobResponseBodyResourceListResource) SetPath(v string) *DescribeFlowJobResponseBodyResourceListResource {
	s.Path = &v
	return s
}

type DescribeFlowJobResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobResponse) SetHeaders(v map[string]*string) *DescribeFlowJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowJobResponse) SetBody(v *DescribeFlowJobResponseBody) *DescribeFlowJobResponse {
	s.Body = v
	return s
}

type DescribeFlowNodeInstanceRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowNodeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceRequest) SetId(v string) *DescribeFlowNodeInstanceRequest {
	s.Id = &v
	return s
}

func (s *DescribeFlowNodeInstanceRequest) SetProjectId(v string) *DescribeFlowNodeInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowNodeInstanceRequest) SetRegionId(v string) *DescribeFlowNodeInstanceRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowNodeInstanceResponseBody struct {
	Adhoc            *bool   `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName      *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	Duration         *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EnvConf          *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	ExternalChildIds *string `json:"ExternalChildIds,omitempty" xml:"ExternalChildIds,omitempty"`
	ExternalId       *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	ExternalInfo     *string `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty"`
	ExternalStatus   *string `json:"ExternalStatus,omitempty" xml:"ExternalStatus,omitempty"`
	ExternalSubId    *string `json:"ExternalSubId,omitempty" xml:"ExternalSubId,omitempty"`
	FailAct          *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	FlowId           *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowInstanceId   *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HostName         *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName          *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobParams        *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	JobType          *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MaxRetry         *string `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	Mode             *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	MonitorConf      *string `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	NodeName         *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	ParamConf        *string `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	Pending          *bool   `json:"Pending,omitempty" xml:"Pending,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Retries          *int32  `json:"Retries,omitempty" xml:"Retries,omitempty"`
	RetryInterval    *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RetryPolicy      *string `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	RunConf          *string `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFlowNodeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceResponseBody) SetAdhoc(v bool) *DescribeFlowNodeInstanceResponseBody {
	s.Adhoc = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetClusterId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetClusterName(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetDuration(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetEndTime(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetEnvConf(v string) *DescribeFlowNodeInstanceResponseBody {
	s.EnvConf = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalChildIds(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalChildIds = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalInfo(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalInfo = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalStatus(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalStatus = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalSubId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalSubId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetFailAct(v string) *DescribeFlowNodeInstanceResponseBody {
	s.FailAct = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetFlowId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.FlowId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetFlowInstanceId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.FlowInstanceId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetGmtCreate(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetGmtModified(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetHostName(v string) *DescribeFlowNodeInstanceResponseBody {
	s.HostName = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetJobId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.JobId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetJobName(v string) *DescribeFlowNodeInstanceResponseBody {
	s.JobName = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetJobParams(v string) *DescribeFlowNodeInstanceResponseBody {
	s.JobParams = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetJobType(v string) *DescribeFlowNodeInstanceResponseBody {
	s.JobType = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetMaxRetry(v string) *DescribeFlowNodeInstanceResponseBody {
	s.MaxRetry = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetMode(v string) *DescribeFlowNodeInstanceResponseBody {
	s.Mode = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetMonitorConf(v string) *DescribeFlowNodeInstanceResponseBody {
	s.MonitorConf = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetNodeName(v string) *DescribeFlowNodeInstanceResponseBody {
	s.NodeName = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetParamConf(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ParamConf = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetPending(v bool) *DescribeFlowNodeInstanceResponseBody {
	s.Pending = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetProjectId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRequestId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRetries(v int32) *DescribeFlowNodeInstanceResponseBody {
	s.Retries = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRetryInterval(v string) *DescribeFlowNodeInstanceResponseBody {
	s.RetryInterval = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRetryPolicy(v string) *DescribeFlowNodeInstanceResponseBody {
	s.RetryPolicy = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRunConf(v string) *DescribeFlowNodeInstanceResponseBody {
	s.RunConf = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetStartTime(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetStatus(v string) *DescribeFlowNodeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetType(v string) *DescribeFlowNodeInstanceResponseBody {
	s.Type = &v
	return s
}

type DescribeFlowNodeInstanceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowNodeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowNodeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceResponse) SetHeaders(v map[string]*string) *DescribeFlowNodeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowNodeInstanceResponse) SetBody(v *DescribeFlowNodeInstanceResponseBody) *DescribeFlowNodeInstanceResponse {
	s.Body = v
	return s
}

type DescribeFlowNodeInstanceContainerLogRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ContainerId    *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	Length         *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	LogName        *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	Offset         *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowNodeInstanceContainerLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetAppId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetContainerId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.ContainerId = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetLength(v int32) *DescribeFlowNodeInstanceContainerLogRequest {
	s.Length = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetLogName(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.LogName = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetNodeInstanceId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.NodeInstanceId = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetOffset(v int32) *DescribeFlowNodeInstanceContainerLogRequest {
	s.Offset = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetProjectId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetRegionId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowNodeInstanceContainerLogResponseBody struct {
	LogEnd    *bool                                                      `json:"LogEnd,omitempty" xml:"LogEnd,omitempty"`
	LogEntrys *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys `json:"LogEntrys,omitempty" xml:"LogEntrys,omitempty" type:"Struct"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowNodeInstanceContainerLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBody) SetLogEnd(v bool) *DescribeFlowNodeInstanceContainerLogResponseBody {
	s.LogEnd = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBody) SetLogEntrys(v *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys) *DescribeFlowNodeInstanceContainerLogResponseBody {
	s.LogEntrys = v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBody) SetRequestId(v string) *DescribeFlowNodeInstanceContainerLogResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys struct {
	LogEntry []*DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry `json:"LogEntry,omitempty" xml:"LogEntry,omitempty" type:"Repeated"`
}

func (s DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys) SetLogEntry(v []*DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry) *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys {
	s.LogEntry = v
	return s
}

type DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry) SetContent(v string) *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry {
	s.Content = &v
	return s
}

type DescribeFlowNodeInstanceContainerLogResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowNodeInstanceContainerLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowNodeInstanceContainerLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogResponse) SetHeaders(v map[string]*string) *DescribeFlowNodeInstanceContainerLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogResponse) SetBody(v *DescribeFlowNodeInstanceContainerLogResponseBody) *DescribeFlowNodeInstanceContainerLogResponse {
	s.Body = v
	return s
}

type DescribeFlowNodeInstanceLauncherLogRequest struct {
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Length         *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Lines          *int32  `json:"Lines,omitempty" xml:"Lines,omitempty"`
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	Offset         *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse        *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	Start          *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFlowNodeInstanceLauncherLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetEndTime(v int64) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetLength(v int32) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Length = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetLines(v int32) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Lines = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetNodeInstanceId(v string) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.NodeInstanceId = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetOffset(v int32) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Offset = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetProjectId(v string) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetRegionId(v string) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetReverse(v bool) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Reverse = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetStart(v int32) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Start = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetStartTime(v int64) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.StartTime = &v
	return s
}

type DescribeFlowNodeInstanceLauncherLogResponseBody struct {
	LogEnd    *bool                                                     `json:"LogEnd,omitempty" xml:"LogEnd,omitempty"`
	LogEntrys *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys `json:"LogEntrys,omitempty" xml:"LogEntrys,omitempty" type:"Struct"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBody) SetLogEnd(v bool) *DescribeFlowNodeInstanceLauncherLogResponseBody {
	s.LogEnd = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBody) SetLogEntrys(v *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys) *DescribeFlowNodeInstanceLauncherLogResponseBody {
	s.LogEntrys = v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBody) SetRequestId(v string) *DescribeFlowNodeInstanceLauncherLogResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys struct {
	LogEntry []*DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry `json:"LogEntry,omitempty" xml:"LogEntry,omitempty" type:"Repeated"`
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys) SetLogEntry(v []*DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry) *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys {
	s.LogEntry = v
	return s
}

type DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry) SetContent(v string) *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry {
	s.Content = &v
	return s
}

type DescribeFlowNodeInstanceLauncherLogResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowNodeInstanceLauncherLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowNodeInstanceLauncherLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogResponse) SetHeaders(v map[string]*string) *DescribeFlowNodeInstanceLauncherLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogResponse) SetBody(v *DescribeFlowNodeInstanceLauncherLogResponseBody) *DescribeFlowNodeInstanceLauncherLogResponse {
	s.Body = v
	return s
}

type DescribeFlowProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectRequest) SetProjectId(v string) *DescribeFlowProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowProjectRequest) SetRegionId(v string) *DescribeFlowProjectRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowProjectResponseBody struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFlowProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectResponseBody) SetDescription(v string) *DescribeFlowProjectResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetGmtCreate(v int64) *DescribeFlowProjectResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetGmtModified(v int64) *DescribeFlowProjectResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetId(v string) *DescribeFlowProjectResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetName(v string) *DescribeFlowProjectResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetRequestId(v string) *DescribeFlowProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetUserId(v string) *DescribeFlowProjectResponseBody {
	s.UserId = &v
	return s
}

type DescribeFlowProjectResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectResponse) SetHeaders(v map[string]*string) *DescribeFlowProjectResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowProjectResponse) SetBody(v *DescribeFlowProjectResponseBody) *DescribeFlowProjectResponse {
	s.Body = v
	return s
}

type DescribeFlowProjectClusterSettingRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingRequest) SetClusterId(v string) *DescribeFlowProjectClusterSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingRequest) SetProjectId(v string) *DescribeFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingRequest) SetRegionId(v string) *DescribeFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowProjectClusterSettingResponseBody struct {
	ClusterId    *string                                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DefaultQueue *string                                                 `json:"DefaultQueue,omitempty" xml:"DefaultQueue,omitempty"`
	DefaultUser  *string                                                 `json:"DefaultUser,omitempty" xml:"DefaultUser,omitempty"`
	GmtCreate    *int64                                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *int64                                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HostList     *DescribeFlowProjectClusterSettingResponseBodyHostList  `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Struct"`
	K8sClusterId *string                                                 `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	ProjectId    *string                                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QueueList    *DescribeFlowProjectClusterSettingResponseBodyQueueList `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Struct"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserList     *DescribeFlowProjectClusterSettingResponseBodyUserList  `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s DescribeFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetClusterId(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetDefaultQueue(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.DefaultQueue = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetDefaultUser(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.DefaultUser = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetGmtCreate(v int64) *DescribeFlowProjectClusterSettingResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetGmtModified(v int64) *DescribeFlowProjectClusterSettingResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetHostList(v *DescribeFlowProjectClusterSettingResponseBodyHostList) *DescribeFlowProjectClusterSettingResponseBody {
	s.HostList = v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetK8sClusterId(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetProjectId(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetQueueList(v *DescribeFlowProjectClusterSettingResponseBodyQueueList) *DescribeFlowProjectClusterSettingResponseBody {
	s.QueueList = v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetRequestId(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetUserList(v *DescribeFlowProjectClusterSettingResponseBodyUserList) *DescribeFlowProjectClusterSettingResponseBody {
	s.UserList = v
	return s
}

type DescribeFlowProjectClusterSettingResponseBodyHostList struct {
	Host []*string `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s DescribeFlowProjectClusterSettingResponseBodyHostList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponseBodyHostList) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponseBodyHostList) SetHost(v []*string) *DescribeFlowProjectClusterSettingResponseBodyHostList {
	s.Host = v
	return s
}

type DescribeFlowProjectClusterSettingResponseBodyQueueList struct {
	Queue []*string `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Repeated"`
}

func (s DescribeFlowProjectClusterSettingResponseBodyQueueList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponseBodyQueueList) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponseBodyQueueList) SetQueue(v []*string) *DescribeFlowProjectClusterSettingResponseBodyQueueList {
	s.Queue = v
	return s
}

type DescribeFlowProjectClusterSettingResponseBodyUserList struct {
	User []*string `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DescribeFlowProjectClusterSettingResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponseBodyUserList) SetUser(v []*string) *DescribeFlowProjectClusterSettingResponseBodyUserList {
	s.User = v
	return s
}

type DescribeFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *DescribeFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponse) SetBody(v *DescribeFlowProjectClusterSettingResponseBody) *DescribeFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type DescribeFlowSLARequest struct {
	From            *int64  `json:"From,omitempty" xml:"From,omitempty"`
	Metrics         *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	PeriodType      *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	To              *int64  `json:"To,omitempty" xml:"To,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFlowSLARequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowSLARequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowSLARequest) SetFrom(v int64) *DescribeFlowSLARequest {
	s.From = &v
	return s
}

func (s *DescribeFlowSLARequest) SetMetrics(v string) *DescribeFlowSLARequest {
	s.Metrics = &v
	return s
}

func (s *DescribeFlowSLARequest) SetPeriodType(v string) *DescribeFlowSLARequest {
	s.PeriodType = &v
	return s
}

func (s *DescribeFlowSLARequest) SetProjectId(v string) *DescribeFlowSLARequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowSLARequest) SetRegionId(v string) *DescribeFlowSLARequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowSLARequest) SetResourceGroupId(v string) *DescribeFlowSLARequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFlowSLARequest) SetResourceOwnerId(v int64) *DescribeFlowSLARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeFlowSLARequest) SetTo(v int64) *DescribeFlowSLARequest {
	s.To = &v
	return s
}

func (s *DescribeFlowSLARequest) SetType(v string) *DescribeFlowSLARequest {
	s.Type = &v
	return s
}

type DescribeFlowSLAResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeFlowSLAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowSLAResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowSLAResponseBody) SetRequestId(v string) *DescribeFlowSLAResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowSLAResponseBody) SetResult(v string) *DescribeFlowSLAResponseBody {
	s.Result = &v
	return s
}

type DescribeFlowSLAResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowSLAResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowSLAResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowSLAResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowSLAResponse) SetHeaders(v map[string]*string) *DescribeFlowSLAResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowSLAResponse) SetBody(v *DescribeFlowSLAResponseBody) *DescribeFlowSLAResponse {
	s.Body = v
	return s
}

type DescribeFlowVariableCollectionRequest struct {
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeFlowVariableCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowVariableCollectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowVariableCollectionRequest) SetEntityId(v string) *DescribeFlowVariableCollectionRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeFlowVariableCollectionRequest) SetRegionId(v string) *DescribeFlowVariableCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowVariableCollectionRequest) SetResourceGroupId(v string) *DescribeFlowVariableCollectionRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeFlowVariableCollectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeFlowVariableCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowVariableCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowVariableCollectionResponseBody) SetRequestId(v string) *DescribeFlowVariableCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowVariableCollectionResponseBody) SetResult(v string) *DescribeFlowVariableCollectionResponseBody {
	s.Result = &v
	return s
}

type DescribeFlowVariableCollectionResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowVariableCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowVariableCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowVariableCollectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowVariableCollectionResponse) SetHeaders(v map[string]*string) *DescribeFlowVariableCollectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowVariableCollectionResponse) SetBody(v *DescribeFlowVariableCollectionResponseBody) *DescribeFlowVariableCollectionResponse {
	s.Body = v
	return s
}

type GetFlowAuditLogsRequest struct {
	CurrentSize     *int32  `json:"CurrentSize,omitempty" xml:"CurrentSize,omitempty"`
	EntityGroupId   *string `json:"EntityGroupId,omitempty" xml:"EntityGroupId,omitempty"`
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType      *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Operation       *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	OperatorId      *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	OrderField      *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	OrderMode       *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	PageCount       *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFlowAuditLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowAuditLogsRequest) GoString() string {
	return s.String()
}

func (s *GetFlowAuditLogsRequest) SetCurrentSize(v int32) *GetFlowAuditLogsRequest {
	s.CurrentSize = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetEntityGroupId(v string) *GetFlowAuditLogsRequest {
	s.EntityGroupId = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetEntityId(v string) *GetFlowAuditLogsRequest {
	s.EntityId = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetEntityType(v string) *GetFlowAuditLogsRequest {
	s.EntityType = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetLimit(v int32) *GetFlowAuditLogsRequest {
	s.Limit = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetOperation(v string) *GetFlowAuditLogsRequest {
	s.Operation = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetOperatorId(v string) *GetFlowAuditLogsRequest {
	s.OperatorId = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetOrderField(v string) *GetFlowAuditLogsRequest {
	s.OrderField = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetOrderMode(v string) *GetFlowAuditLogsRequest {
	s.OrderMode = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetPageCount(v int32) *GetFlowAuditLogsRequest {
	s.PageCount = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetPageNumber(v int32) *GetFlowAuditLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetPageSize(v int32) *GetFlowAuditLogsRequest {
	s.PageSize = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetRegionId(v string) *GetFlowAuditLogsRequest {
	s.RegionId = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetResourceOwnerId(v int64) *GetFlowAuditLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFlowAuditLogsRequest) SetStatus(v string) *GetFlowAuditLogsRequest {
	s.Status = &v
	return s
}

type GetFlowAuditLogsResponseBody struct {
	Items      *GetFlowAuditLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetFlowAuditLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowAuditLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowAuditLogsResponseBody) SetItems(v *GetFlowAuditLogsResponseBodyItems) *GetFlowAuditLogsResponseBody {
	s.Items = v
	return s
}

func (s *GetFlowAuditLogsResponseBody) SetPageNumber(v int32) *GetFlowAuditLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetFlowAuditLogsResponseBody) SetPageSize(v int32) *GetFlowAuditLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetFlowAuditLogsResponseBody) SetRequestId(v string) *GetFlowAuditLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowAuditLogsResponseBody) SetTotalCount(v int32) *GetFlowAuditLogsResponseBody {
	s.TotalCount = &v
	return s
}

type GetFlowAuditLogsResponseBodyItems struct {
	Item []*GetFlowAuditLogsResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s GetFlowAuditLogsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s GetFlowAuditLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetFlowAuditLogsResponseBodyItems) SetItem(v []*GetFlowAuditLogsResponseBodyItemsItem) *GetFlowAuditLogsResponseBodyItems {
	s.Item = v
	return s
}

type GetFlowAuditLogsResponseBodyItemsItem struct {
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EntityGroupId *string `json:"EntityGroupId,omitempty" xml:"EntityGroupId,omitempty"`
	EntityId      *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType    *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Operation     *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	OperatorId    *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Ts            *int64  `json:"Ts,omitempty" xml:"Ts,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetFlowAuditLogsResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s GetFlowAuditLogsResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *GetFlowAuditLogsResponseBodyItemsItem) SetContent(v string) *GetFlowAuditLogsResponseBodyItemsItem {
	s.Content = &v
	return s
}

func (s *GetFlowAuditLogsResponseBodyItemsItem) SetEntityGroupId(v string) *GetFlowAuditLogsResponseBodyItemsItem {
	s.EntityGroupId = &v
	return s
}

func (s *GetFlowAuditLogsResponseBodyItemsItem) SetEntityId(v string) *GetFlowAuditLogsResponseBodyItemsItem {
	s.EntityId = &v
	return s
}

func (s *GetFlowAuditLogsResponseBodyItemsItem) SetEntityType(v string) *GetFlowAuditLogsResponseBodyItemsItem {
	s.EntityType = &v
	return s
}

func (s *GetFlowAuditLogsResponseBodyItemsItem) SetOperation(v string) *GetFlowAuditLogsResponseBodyItemsItem {
	s.Operation = &v
	return s
}

func (s *GetFlowAuditLogsResponseBodyItemsItem) SetOperatorId(v string) *GetFlowAuditLogsResponseBodyItemsItem {
	s.OperatorId = &v
	return s
}

func (s *GetFlowAuditLogsResponseBodyItemsItem) SetStatus(v string) *GetFlowAuditLogsResponseBodyItemsItem {
	s.Status = &v
	return s
}

func (s *GetFlowAuditLogsResponseBodyItemsItem) SetTs(v int64) *GetFlowAuditLogsResponseBodyItemsItem {
	s.Ts = &v
	return s
}

func (s *GetFlowAuditLogsResponseBodyItemsItem) SetUserId(v string) *GetFlowAuditLogsResponseBodyItemsItem {
	s.UserId = &v
	return s
}

type GetFlowAuditLogsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFlowAuditLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFlowAuditLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowAuditLogsResponse) GoString() string {
	return s.String()
}

func (s *GetFlowAuditLogsResponse) SetHeaders(v map[string]*string) *GetFlowAuditLogsResponse {
	s.Headers = v
	return s
}

func (s *GetFlowAuditLogsResponse) SetBody(v *GetFlowAuditLogsResponseBody) *GetFlowAuditLogsResponse {
	s.Body = v
	return s
}

type KillFlowRequest struct {
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s KillFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s KillFlowRequest) GoString() string {
	return s.String()
}

func (s *KillFlowRequest) SetFlowInstanceId(v string) *KillFlowRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *KillFlowRequest) SetProjectId(v string) *KillFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *KillFlowRequest) SetRegionId(v string) *KillFlowRequest {
	s.RegionId = &v
	return s
}

type KillFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillFlowResponseBody) GoString() string {
	return s.String()
}

func (s *KillFlowResponseBody) SetData(v bool) *KillFlowResponseBody {
	s.Data = &v
	return s
}

func (s *KillFlowResponseBody) SetRequestId(v string) *KillFlowResponseBody {
	s.RequestId = &v
	return s
}

type KillFlowResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *KillFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KillFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s KillFlowResponse) GoString() string {
	return s.String()
}

func (s *KillFlowResponse) SetHeaders(v map[string]*string) *KillFlowResponse {
	s.Headers = v
	return s
}

func (s *KillFlowResponse) SetBody(v *KillFlowResponseBody) *KillFlowResponse {
	s.Body = v
	return s
}

type KillFlowJobRequest struct {
	JobInstanceId *string `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s KillFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s KillFlowJobRequest) GoString() string {
	return s.String()
}

func (s *KillFlowJobRequest) SetJobInstanceId(v string) *KillFlowJobRequest {
	s.JobInstanceId = &v
	return s
}

func (s *KillFlowJobRequest) SetProjectId(v string) *KillFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *KillFlowJobRequest) SetRegionId(v string) *KillFlowJobRequest {
	s.RegionId = &v
	return s
}

type KillFlowJobResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *KillFlowJobResponseBody) SetData(v bool) *KillFlowJobResponseBody {
	s.Data = &v
	return s
}

func (s *KillFlowJobResponseBody) SetRequestId(v string) *KillFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type KillFlowJobResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *KillFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KillFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s KillFlowJobResponse) GoString() string {
	return s.String()
}

func (s *KillFlowJobResponse) SetHeaders(v map[string]*string) *KillFlowJobResponse {
	s.Headers = v
	return s
}

func (s *KillFlowJobResponse) SetBody(v *KillFlowJobResponseBody) *KillFlowJobResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	ClusterTypeList []*string                 `json:"ClusterTypeList,omitempty" xml:"ClusterTypeList,omitempty" type:"Repeated"`
	CreateType      *string                   `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	DefaultStatus   *bool                     `json:"DefaultStatus,omitempty" xml:"DefaultStatus,omitempty"`
	DepositType     *string                   `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	ExpiredTagList  []*string                 `json:"ExpiredTagList,omitempty" xml:"ExpiredTagList,omitempty" type:"Repeated"`
	IsDesc          *bool                     `json:"IsDesc,omitempty" xml:"IsDesc,omitempty"`
	MachineType     *string                   `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	Name            *string                   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber      *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StatusList      []*string                 `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	Tag             []*ListClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetClusterTypeList(v []*string) *ListClustersRequest {
	s.ClusterTypeList = v
	return s
}

func (s *ListClustersRequest) SetCreateType(v string) *ListClustersRequest {
	s.CreateType = &v
	return s
}

func (s *ListClustersRequest) SetDefaultStatus(v bool) *ListClustersRequest {
	s.DefaultStatus = &v
	return s
}

func (s *ListClustersRequest) SetDepositType(v string) *ListClustersRequest {
	s.DepositType = &v
	return s
}

func (s *ListClustersRequest) SetExpiredTagList(v []*string) *ListClustersRequest {
	s.ExpiredTagList = v
	return s
}

func (s *ListClustersRequest) SetIsDesc(v bool) *ListClustersRequest {
	s.IsDesc = &v
	return s
}

func (s *ListClustersRequest) SetMachineType(v string) *ListClustersRequest {
	s.MachineType = &v
	return s
}

func (s *ListClustersRequest) SetName(v string) *ListClustersRequest {
	s.Name = &v
	return s
}

func (s *ListClustersRequest) SetPageNumber(v int32) *ListClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersRequest) SetRegionId(v string) *ListClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListClustersRequest) SetResourceGroupId(v string) *ListClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersRequest) SetResourceOwnerId(v int64) *ListClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClustersRequest) SetStatusList(v []*string) *ListClustersRequest {
	s.StatusList = v
	return s
}

func (s *ListClustersRequest) SetTag(v []*ListClustersRequestTag) *ListClustersRequest {
	s.Tag = v
	return s
}

type ListClustersRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequestTag) GoString() string {
	return s.String()
}

func (s *ListClustersRequestTag) SetKey(v string) *ListClustersRequestTag {
	s.Key = &v
	return s
}

func (s *ListClustersRequestTag) SetValue(v string) *ListClustersRequestTag {
	s.Value = &v
	return s
}

type ListClustersResponseBody struct {
	Clusters   *ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetClusters(v *ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetPageNumber(v int32) *ListClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

type ListClustersResponseBodyClusters struct {
	ClusterInfo []*ListClustersResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetClusterInfo(v []*ListClustersResponseBodyClustersClusterInfo) *ListClustersResponseBodyClusters {
	s.ClusterInfo = v
	return s
}

type ListClustersResponseBodyClustersClusterInfo struct {
	ChargeType          *string                                                   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateResource      *string                                                   `json:"CreateResource,omitempty" xml:"CreateResource,omitempty"`
	CreateTime          *int64                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DepositType         *string                                                   `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	ExpiredTime         *int64                                                    `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FailReason          *ListClustersResponseBodyClustersClusterInfoFailReason    `json:"FailReason,omitempty" xml:"FailReason,omitempty" type:"Struct"`
	HasUncompletedOrder *bool                                                     `json:"HasUncompletedOrder,omitempty" xml:"HasUncompletedOrder,omitempty"`
	Id                  *string                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	K8sClusterId        *string                                                   `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	MachineType         *string                                                   `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	MetaStoreType       *string                                                   `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	Name                *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderList           *string                                                   `json:"OrderList,omitempty" xml:"OrderList,omitempty"`
	OrderTaskInfo       *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo `json:"OrderTaskInfo,omitempty" xml:"OrderTaskInfo,omitempty" type:"Struct"`
	Period              *int32                                                    `json:"Period,omitempty" xml:"Period,omitempty"`
	RunningTime         *int32                                                    `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status              *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                *ListClustersResponseBodyClustersClusterInfoTags          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Type                *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfo) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetChargeType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.ChargeType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetCreateResource(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.CreateResource = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetCreateTime(v int64) *ListClustersResponseBodyClustersClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetDepositType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.DepositType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetExpiredTime(v int64) *ListClustersResponseBodyClustersClusterInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetFailReason(v *ListClustersResponseBodyClustersClusterInfoFailReason) *ListClustersResponseBodyClustersClusterInfo {
	s.FailReason = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetHasUncompletedOrder(v bool) *ListClustersResponseBodyClustersClusterInfo {
	s.HasUncompletedOrder = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetId(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.Id = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetK8sClusterId(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.K8sClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetMachineType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.MachineType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetMetaStoreType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.MetaStoreType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetName(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetOrderList(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.OrderList = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetOrderTaskInfo(v *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) *ListClustersResponseBodyClustersClusterInfo {
	s.OrderTaskInfo = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetPeriod(v int32) *ListClustersResponseBodyClustersClusterInfo {
	s.Period = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetRunningTime(v int32) *ListClustersResponseBodyClustersClusterInfo {
	s.RunningTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetStatus(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetTags(v *ListClustersResponseBodyClustersClusterInfoTags) *ListClustersResponseBodyClustersClusterInfo {
	s.Tags = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.Type = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoFailReason struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoFailReason) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoFailReason) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoFailReason) SetErrorCode(v string) *ListClustersResponseBodyClustersClusterInfoFailReason {
	s.ErrorCode = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoFailReason) SetErrorMsg(v string) *ListClustersResponseBodyClustersClusterInfoFailReason {
	s.ErrorMsg = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoFailReason) SetRequestId(v string) *ListClustersResponseBodyClustersClusterInfoFailReason {
	s.RequestId = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoOrderTaskInfo struct {
	CurrentCount *int32  `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	OrderIdList  *string `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty"`
	TargetCount  *int32  `json:"TargetCount,omitempty" xml:"TargetCount,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) SetCurrentCount(v int32) *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo {
	s.CurrentCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) SetOrderIdList(v string) *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo {
	s.OrderIdList = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) SetTargetCount(v int32) *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo {
	s.TargetCount = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoTags struct {
	Tag []*ListClustersResponseBodyClustersClusterInfoTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBodyClustersClusterInfoTags) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoTags) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoTags) SetTag(v []*ListClustersResponseBodyClustersClusterInfoTagsTag) *ListClustersResponseBodyClustersClusterInfoTags {
	s.Tag = v
	return s
}

type ListClustersResponseBodyClustersClusterInfoTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoTagsTag) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoTagsTag) SetTagKey(v string) *ListClustersResponseBodyClustersClusterInfoTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoTagsTag) SetTagValue(v string) *ListClustersResponseBodyClustersClusterInfoTagsTag {
	s.TagValue = &v
	return s
}

type ListClustersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListFlowRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Periodic   *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRequest) GoString() string {
	return s.String()
}

func (s *ListFlowRequest) SetClusterId(v string) *ListFlowRequest {
	s.ClusterId = &v
	return s
}

func (s *ListFlowRequest) SetId(v string) *ListFlowRequest {
	s.Id = &v
	return s
}

func (s *ListFlowRequest) SetJobId(v string) *ListFlowRequest {
	s.JobId = &v
	return s
}

func (s *ListFlowRequest) SetName(v string) *ListFlowRequest {
	s.Name = &v
	return s
}

func (s *ListFlowRequest) SetPageNumber(v int32) *ListFlowRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowRequest) SetPageSize(v int32) *ListFlowRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowRequest) SetPeriodic(v bool) *ListFlowRequest {
	s.Periodic = &v
	return s
}

func (s *ListFlowRequest) SetProjectId(v string) *ListFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowRequest) SetRegionId(v string) *ListFlowRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowRequest) SetStatus(v string) *ListFlowRequest {
	s.Status = &v
	return s
}

type ListFlowResponseBody struct {
	Flow       *ListFlowResponseBodyFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Struct"`
	PageNumber *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBody) SetFlow(v *ListFlowResponseBodyFlow) *ListFlowResponseBody {
	s.Flow = v
	return s
}

func (s *ListFlowResponseBody) SetPageNumber(v int32) *ListFlowResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowResponseBody) SetPageSize(v int32) *ListFlowResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowResponseBody) SetRequestId(v string) *ListFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowResponseBody) SetTotal(v int32) *ListFlowResponseBody {
	s.Total = &v
	return s
}

type ListFlowResponseBodyFlow struct {
	Flow []*ListFlowResponseBodyFlowFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Repeated"`
}

func (s ListFlowResponseBodyFlow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBodyFlow) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBodyFlow) SetFlow(v []*ListFlowResponseBodyFlowFlow) *ListFlowResponseBodyFlow {
	s.Flow = v
	return s
}

type ListFlowResponseBodyFlowFlow struct {
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	CategoryId              *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	CronExpr                *string `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	GmtCreate               *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified             *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Graph                   *string `json:"Graph,omitempty" xml:"Graph,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Periodic                *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowResponseBodyFlowFlow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBodyFlowFlow) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBodyFlowFlow) SetAlertConf(v string) *ListFlowResponseBodyFlowFlow {
	s.AlertConf = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetAlertDingDingGroupBizId(v string) *ListFlowResponseBodyFlowFlow {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetAlertUserGroupBizId(v string) *ListFlowResponseBodyFlowFlow {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetCategoryId(v string) *ListFlowResponseBodyFlowFlow {
	s.CategoryId = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetClusterId(v string) *ListFlowResponseBodyFlowFlow {
	s.ClusterId = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetCreateCluster(v bool) *ListFlowResponseBodyFlowFlow {
	s.CreateCluster = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetCronExpr(v string) *ListFlowResponseBodyFlowFlow {
	s.CronExpr = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetDescription(v string) *ListFlowResponseBodyFlowFlow {
	s.Description = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetEndSchedule(v int64) *ListFlowResponseBodyFlowFlow {
	s.EndSchedule = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetGmtCreate(v int64) *ListFlowResponseBodyFlowFlow {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetGmtModified(v int64) *ListFlowResponseBodyFlowFlow {
	s.GmtModified = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetGraph(v string) *ListFlowResponseBodyFlowFlow {
	s.Graph = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetHostName(v string) *ListFlowResponseBodyFlowFlow {
	s.HostName = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetId(v string) *ListFlowResponseBodyFlowFlow {
	s.Id = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetName(v string) *ListFlowResponseBodyFlowFlow {
	s.Name = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetPeriodic(v bool) *ListFlowResponseBodyFlowFlow {
	s.Periodic = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetProjectId(v string) *ListFlowResponseBodyFlowFlow {
	s.ProjectId = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetStartSchedule(v int64) *ListFlowResponseBodyFlowFlow {
	s.StartSchedule = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetStatus(v string) *ListFlowResponseBodyFlowFlow {
	s.Status = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetType(v string) *ListFlowResponseBodyFlowFlow {
	s.Type = &v
	return s
}

type ListFlowResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponse) GoString() string {
	return s.String()
}

func (s *ListFlowResponse) SetHeaders(v map[string]*string) *ListFlowResponse {
	s.Headers = v
	return s
}

func (s *ListFlowResponse) SetBody(v *ListFlowResponseBody) *ListFlowResponse {
	s.Body = v
	return s
}

type ListFlowClusterRequest struct {
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterRequest) GoString() string {
	return s.String()
}

func (s *ListFlowClusterRequest) SetPageNumber(v int32) *ListFlowClusterRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowClusterRequest) SetPageSize(v int32) *ListFlowClusterRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowClusterRequest) SetProjectId(v string) *ListFlowClusterRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowClusterRequest) SetRegionId(v string) *ListFlowClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowClusterRequest) SetResourceGroupId(v string) *ListFlowClusterRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowClusterResponseBody struct {
	Clusters   *ListFlowClusterResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFlowClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBody) SetClusters(v *ListFlowClusterResponseBodyClusters) *ListFlowClusterResponseBody {
	s.Clusters = v
	return s
}

func (s *ListFlowClusterResponseBody) SetPageNumber(v int32) *ListFlowClusterResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowClusterResponseBody) SetPageSize(v int32) *ListFlowClusterResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowClusterResponseBody) SetRequestId(v string) *ListFlowClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowClusterResponseBody) SetTotalCount(v int32) *ListFlowClusterResponseBody {
	s.TotalCount = &v
	return s
}

type ListFlowClusterResponseBodyClusters struct {
	ClusterInfo []*ListFlowClusterResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Repeated"`
}

func (s ListFlowClusterResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBodyClusters) SetClusterInfo(v []*ListFlowClusterResponseBodyClustersClusterInfo) *ListFlowClusterResponseBodyClusters {
	s.ClusterInfo = v
	return s
}

type ListFlowClusterResponseBodyClustersClusterInfo struct {
	ChargeType          *string                                                      `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateResource      *string                                                      `json:"CreateResource,omitempty" xml:"CreateResource,omitempty"`
	CreateTime          *int64                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpiredTime         *int64                                                       `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FailReason          *ListFlowClusterResponseBodyClustersClusterInfoFailReason    `json:"FailReason,omitempty" xml:"FailReason,omitempty" type:"Struct"`
	HasUncompletedOrder *bool                                                        `json:"HasUncompletedOrder,omitempty" xml:"HasUncompletedOrder,omitempty"`
	Id                  *string                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	K8sClusterId        *string                                                      `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	Name                *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderList           *string                                                      `json:"OrderList,omitempty" xml:"OrderList,omitempty"`
	OrderTaskInfo       *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo `json:"OrderTaskInfo,omitempty" xml:"OrderTaskInfo,omitempty" type:"Struct"`
	Period              *int32                                                       `json:"Period,omitempty" xml:"Period,omitempty"`
	RunningTime         *int32                                                       `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status              *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowClusterResponseBodyClustersClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBodyClustersClusterInfo) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetChargeType(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.ChargeType = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetCreateResource(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.CreateResource = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetCreateTime(v int64) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetExpiredTime(v int64) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetFailReason(v *ListFlowClusterResponseBodyClustersClusterInfoFailReason) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.FailReason = v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetHasUncompletedOrder(v bool) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.HasUncompletedOrder = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetId(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Id = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetK8sClusterId(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.K8sClusterId = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetName(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Name = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetOrderList(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.OrderList = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetOrderTaskInfo(v *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.OrderTaskInfo = v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetPeriod(v int32) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Period = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetRunningTime(v int32) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.RunningTime = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetStatus(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Status = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetType(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Type = &v
	return s
}

type ListFlowClusterResponseBodyClustersClusterInfoFailReason struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowClusterResponseBodyClustersClusterInfoFailReason) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBodyClustersClusterInfoFailReason) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoFailReason) SetErrorCode(v string) *ListFlowClusterResponseBodyClustersClusterInfoFailReason {
	s.ErrorCode = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoFailReason) SetErrorMsg(v string) *ListFlowClusterResponseBodyClustersClusterInfoFailReason {
	s.ErrorMsg = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoFailReason) SetRequestId(v string) *ListFlowClusterResponseBodyClustersClusterInfoFailReason {
	s.RequestId = &v
	return s
}

type ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo struct {
	CurrentCount *int32  `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	OrderIdList  *string `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty"`
	TargetCount  *int32  `json:"TargetCount,omitempty" xml:"TargetCount,omitempty"`
}

func (s ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) SetCurrentCount(v int32) *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo {
	s.CurrentCount = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) SetOrderIdList(v string) *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo {
	s.OrderIdList = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) SetTargetCount(v int32) *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo {
	s.TargetCount = &v
	return s
}

type ListFlowClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponse) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponse) SetHeaders(v map[string]*string) *ListFlowClusterResponse {
	s.Headers = v
	return s
}

func (s *ListFlowClusterResponse) SetBody(v *ListFlowClusterResponseBody) *ListFlowClusterResponse {
	s.Body = v
	return s
}

type ListFlowClusterAllRequest struct {
	ProductType     *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowClusterAllRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllRequest) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllRequest) SetProductType(v string) *ListFlowClusterAllRequest {
	s.ProductType = &v
	return s
}

func (s *ListFlowClusterAllRequest) SetRegionId(v string) *ListFlowClusterAllRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowClusterAllRequest) SetResourceGroupId(v string) *ListFlowClusterAllRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowClusterAllResponseBody struct {
	Clusters   *ListFlowClusterAllResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFlowClusterAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBody) SetClusters(v *ListFlowClusterAllResponseBodyClusters) *ListFlowClusterAllResponseBody {
	s.Clusters = v
	return s
}

func (s *ListFlowClusterAllResponseBody) SetPageNumber(v int32) *ListFlowClusterAllResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowClusterAllResponseBody) SetPageSize(v int32) *ListFlowClusterAllResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowClusterAllResponseBody) SetRequestId(v string) *ListFlowClusterAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowClusterAllResponseBody) SetTotalCount(v int32) *ListFlowClusterAllResponseBody {
	s.TotalCount = &v
	return s
}

type ListFlowClusterAllResponseBodyClusters struct {
	ClusterInfo []*ListFlowClusterAllResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Repeated"`
}

func (s ListFlowClusterAllResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBodyClusters) SetClusterInfo(v []*ListFlowClusterAllResponseBodyClustersClusterInfo) *ListFlowClusterAllResponseBodyClusters {
	s.ClusterInfo = v
	return s
}

type ListFlowClusterAllResponseBodyClustersClusterInfo struct {
	ChargeType          *string                                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateResource      *string                                                         `json:"CreateResource,omitempty" xml:"CreateResource,omitempty"`
	CreateTime          *int64                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpiredTime         *int64                                                          `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FailReason          *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason    `json:"FailReason,omitempty" xml:"FailReason,omitempty" type:"Struct"`
	HasUncompletedOrder *bool                                                           `json:"HasUncompletedOrder,omitempty" xml:"HasUncompletedOrder,omitempty"`
	Id                  *string                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	K8sClusterId        *string                                                         `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	Name                *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderList           *string                                                         `json:"OrderList,omitempty" xml:"OrderList,omitempty"`
	OrderTaskInfo       *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo `json:"OrderTaskInfo,omitempty" xml:"OrderTaskInfo,omitempty" type:"Struct"`
	Period              *int32                                                          `json:"Period,omitempty" xml:"Period,omitempty"`
	RunningTime         *int32                                                          `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Status              *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfo) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetChargeType(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.ChargeType = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetCreateResource(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.CreateResource = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetCreateTime(v int64) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetExpiredTime(v int64) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetFailReason(v *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.FailReason = v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetHasUncompletedOrder(v bool) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.HasUncompletedOrder = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetId(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Id = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetK8sClusterId(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.K8sClusterId = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetName(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Name = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetOrderList(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.OrderList = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetOrderTaskInfo(v *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.OrderTaskInfo = v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetPeriod(v int32) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Period = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetRunningTime(v int32) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.RunningTime = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetStatus(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Status = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetType(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Type = &v
	return s
}

type ListFlowClusterAllResponseBodyClustersClusterInfoFailReason struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) SetErrorCode(v string) *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason {
	s.ErrorCode = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) SetErrorMsg(v string) *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason {
	s.ErrorMsg = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) SetRequestId(v string) *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason {
	s.RequestId = &v
	return s
}

type ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo struct {
	CurrentCount *int32  `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	OrderIdList  *string `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty"`
	TargetCount  *int32  `json:"TargetCount,omitempty" xml:"TargetCount,omitempty"`
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) SetCurrentCount(v int32) *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo {
	s.CurrentCount = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) SetOrderIdList(v string) *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo {
	s.OrderIdList = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) SetTargetCount(v int32) *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo {
	s.TargetCount = &v
	return s
}

type ListFlowClusterAllResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowClusterAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowClusterAllResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponse) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponse) SetHeaders(v map[string]*string) *ListFlowClusterAllResponse {
	s.Headers = v
	return s
}

func (s *ListFlowClusterAllResponse) SetBody(v *ListFlowClusterAllResponseBody) *ListFlowClusterAllResponse {
	s.Body = v
	return s
}

type ListFlowClusterAllHostsRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowClusterAllHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsRequest) SetClusterId(v string) *ListFlowClusterAllHostsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListFlowClusterAllHostsRequest) SetProjectId(v string) *ListFlowClusterAllHostsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowClusterAllHostsRequest) SetRegionId(v string) *ListFlowClusterAllHostsRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowClusterAllHostsRequest) SetResourceGroupId(v string) *ListFlowClusterAllHostsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowClusterAllHostsResponseBody struct {
	HostList  *ListFlowClusterAllHostsResponseBodyHostList `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowClusterAllHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsResponseBody) SetHostList(v *ListFlowClusterAllHostsResponseBodyHostList) *ListFlowClusterAllHostsResponseBody {
	s.HostList = v
	return s
}

func (s *ListFlowClusterAllHostsResponseBody) SetRequestId(v string) *ListFlowClusterAllHostsResponseBody {
	s.RequestId = &v
	return s
}

type ListFlowClusterAllHostsResponseBodyHostList struct {
	Host []*ListFlowClusterAllHostsResponseBodyHostListHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s ListFlowClusterAllHostsResponseBodyHostList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsResponseBodyHostList) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsResponseBodyHostList) SetHost(v []*ListFlowClusterAllHostsResponseBodyHostListHost) *ListFlowClusterAllHostsResponseBodyHostList {
	s.Host = v
	return s
}

type ListFlowClusterAllHostsResponseBodyHostListHost struct {
	Cpu            *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	HostId         *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostInstanceId *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	HostName       *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Memory         *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PrivateIp      *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	PublicIp       *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	Role           *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SerialNumber   *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowClusterAllHostsResponseBodyHostListHost) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsResponseBodyHostListHost) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetCpu(v int32) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Cpu = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetHostId(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.HostId = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetHostInstanceId(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.HostInstanceId = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetHostName(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.HostName = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetInstanceType(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.InstanceType = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetMemory(v int32) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Memory = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetPrivateIp(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.PrivateIp = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetPublicIp(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.PublicIp = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetRole(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Role = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetSerialNumber(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.SerialNumber = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetStatus(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Status = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetType(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Type = &v
	return s
}

type ListFlowClusterAllHostsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowClusterAllHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowClusterAllHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsResponse) SetHeaders(v map[string]*string) *ListFlowClusterAllHostsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowClusterAllHostsResponse) SetBody(v *ListFlowClusterAllHostsResponseBody) *ListFlowClusterAllHostsResponse {
	s.Body = v
	return s
}

type ListFlowClusterHostRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowClusterHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostRequest) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostRequest) SetClusterId(v string) *ListFlowClusterHostRequest {
	s.ClusterId = &v
	return s
}

func (s *ListFlowClusterHostRequest) SetProjectId(v string) *ListFlowClusterHostRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowClusterHostRequest) SetRegionId(v string) *ListFlowClusterHostRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowClusterHostRequest) SetResourceGroupId(v string) *ListFlowClusterHostRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowClusterHostResponseBody struct {
	HostList  *ListFlowClusterHostResponseBodyHostList `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowClusterHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostResponseBody) SetHostList(v *ListFlowClusterHostResponseBodyHostList) *ListFlowClusterHostResponseBody {
	s.HostList = v
	return s
}

func (s *ListFlowClusterHostResponseBody) SetRequestId(v string) *ListFlowClusterHostResponseBody {
	s.RequestId = &v
	return s
}

type ListFlowClusterHostResponseBodyHostList struct {
	Host []*ListFlowClusterHostResponseBodyHostListHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s ListFlowClusterHostResponseBodyHostList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostResponseBodyHostList) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostResponseBodyHostList) SetHost(v []*ListFlowClusterHostResponseBodyHostListHost) *ListFlowClusterHostResponseBodyHostList {
	s.Host = v
	return s
}

type ListFlowClusterHostResponseBodyHostListHost struct {
	Cpu            *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	HostId         *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostInstanceId *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	HostName       *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Memory         *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PrivateIp      *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	PublicIp       *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	Role           *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SerialNumber   *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowClusterHostResponseBodyHostListHost) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostResponseBodyHostListHost) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetCpu(v int32) *ListFlowClusterHostResponseBodyHostListHost {
	s.Cpu = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetHostId(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.HostId = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetHostInstanceId(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.HostInstanceId = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetHostName(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.HostName = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetInstanceType(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.InstanceType = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetMemory(v int32) *ListFlowClusterHostResponseBodyHostListHost {
	s.Memory = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetPrivateIp(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.PrivateIp = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetPublicIp(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.PublicIp = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetRole(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.Role = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetSerialNumber(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.SerialNumber = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetStatus(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.Status = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetType(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.Type = &v
	return s
}

type ListFlowClusterHostResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowClusterHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowClusterHostResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostResponse) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostResponse) SetHeaders(v map[string]*string) *ListFlowClusterHostResponse {
	s.Headers = v
	return s
}

func (s *ListFlowClusterHostResponse) SetBody(v *ListFlowClusterHostResponseBody) *ListFlowClusterHostResponse {
	s.Body = v
	return s
}

type ListFlowEntitySnapshotRequest struct {
	CommitterId     *string `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	CurrentSize     *int32  `json:"CurrentSize,omitempty" xml:"CurrentSize,omitempty"`
	EntityGroupId   *string `json:"EntityGroupId,omitempty" xml:"EntityGroupId,omitempty"`
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType      *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OrderField      *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	OrderMode       *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	PageCount       *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Revision        *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
}

func (s ListFlowEntitySnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotRequest) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotRequest) SetCommitterId(v string) *ListFlowEntitySnapshotRequest {
	s.CommitterId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetCurrentSize(v int32) *ListFlowEntitySnapshotRequest {
	s.CurrentSize = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetEntityGroupId(v string) *ListFlowEntitySnapshotRequest {
	s.EntityGroupId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetEntityId(v string) *ListFlowEntitySnapshotRequest {
	s.EntityId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetEntityType(v string) *ListFlowEntitySnapshotRequest {
	s.EntityType = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetLimit(v int32) *ListFlowEntitySnapshotRequest {
	s.Limit = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetOrderField(v string) *ListFlowEntitySnapshotRequest {
	s.OrderField = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetOrderMode(v string) *ListFlowEntitySnapshotRequest {
	s.OrderMode = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetPageCount(v int32) *ListFlowEntitySnapshotRequest {
	s.PageCount = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetPageNumber(v int32) *ListFlowEntitySnapshotRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetPageSize(v int32) *ListFlowEntitySnapshotRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetRegionId(v string) *ListFlowEntitySnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetResourceOwnerId(v int64) *ListFlowEntitySnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetRevision(v string) *ListFlowEntitySnapshotRequest {
	s.Revision = &v
	return s
}

type ListFlowEntitySnapshotResponseBody struct {
	Items      *ListFlowEntitySnapshotResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFlowEntitySnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotResponseBody) SetItems(v *ListFlowEntitySnapshotResponseBodyItems) *ListFlowEntitySnapshotResponseBody {
	s.Items = v
	return s
}

func (s *ListFlowEntitySnapshotResponseBody) SetPageNumber(v int32) *ListFlowEntitySnapshotResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBody) SetPageSize(v int32) *ListFlowEntitySnapshotResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBody) SetRequestId(v string) *ListFlowEntitySnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBody) SetTotalCount(v int32) *ListFlowEntitySnapshotResponseBody {
	s.TotalCount = &v
	return s
}

type ListFlowEntitySnapshotResponseBodyItems struct {
	Item []*ListFlowEntitySnapshotResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s ListFlowEntitySnapshotResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotResponseBodyItems) SetItem(v []*ListFlowEntitySnapshotResponseBodyItemsItem) *ListFlowEntitySnapshotResponseBodyItems {
	s.Item = v
	return s
}

type ListFlowEntitySnapshotResponseBodyItemsItem struct {
	Active        *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	CommitterId   *string `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	Data          *string `json:"Data,omitempty" xml:"Data,omitempty"`
	EntityGroupId *string `json:"EntityGroupId,omitempty" xml:"EntityGroupId,omitempty"`
	EntityId      *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType    *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Revision      *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListFlowEntitySnapshotResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetActive(v bool) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.Active = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetCommitterId(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.CommitterId = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetData(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.Data = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetEntityGroupId(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.EntityGroupId = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetEntityId(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.EntityId = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetEntityType(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.EntityType = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetGmtCreate(v int64) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetMessage(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.Message = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetRevision(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.Revision = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetUserId(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.UserId = &v
	return s
}

type ListFlowEntitySnapshotResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowEntitySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowEntitySnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotResponse) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotResponse) SetHeaders(v map[string]*string) *ListFlowEntitySnapshotResponse {
	s.Headers = v
	return s
}

func (s *ListFlowEntitySnapshotResponse) SetBody(v *ListFlowEntitySnapshotResponseBody) *ListFlowEntitySnapshotResponse {
	s.Body = v
	return s
}

type ListFlowInstanceRequest struct {
	FlowId     *string   `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName   *string   `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	Id         *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderBy    *string   `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderType  *string   `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Owner      *string   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	PageNumber *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId  *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	TimeRange  *string   `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s ListFlowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceRequest) SetFlowId(v string) *ListFlowInstanceRequest {
	s.FlowId = &v
	return s
}

func (s *ListFlowInstanceRequest) SetFlowName(v string) *ListFlowInstanceRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowInstanceRequest) SetId(v string) *ListFlowInstanceRequest {
	s.Id = &v
	return s
}

func (s *ListFlowInstanceRequest) SetInstanceId(v string) *ListFlowInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlowInstanceRequest) SetOrderBy(v string) *ListFlowInstanceRequest {
	s.OrderBy = &v
	return s
}

func (s *ListFlowInstanceRequest) SetOrderType(v string) *ListFlowInstanceRequest {
	s.OrderType = &v
	return s
}

func (s *ListFlowInstanceRequest) SetOwner(v string) *ListFlowInstanceRequest {
	s.Owner = &v
	return s
}

func (s *ListFlowInstanceRequest) SetPageNumber(v int32) *ListFlowInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowInstanceRequest) SetPageSize(v int32) *ListFlowInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowInstanceRequest) SetProjectId(v string) *ListFlowInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowInstanceRequest) SetRegionId(v string) *ListFlowInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowInstanceRequest) SetStatusList(v []*string) *ListFlowInstanceRequest {
	s.StatusList = v
	return s
}

func (s *ListFlowInstanceRequest) SetTimeRange(v string) *ListFlowInstanceRequest {
	s.TimeRange = &v
	return s
}

type ListFlowInstanceResponseBody struct {
	FlowInstances *ListFlowInstanceResponseBodyFlowInstances `json:"FlowInstances,omitempty" xml:"FlowInstances,omitempty" type:"Struct"`
	PageNumber    *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total         *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFlowInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceResponseBody) SetFlowInstances(v *ListFlowInstanceResponseBodyFlowInstances) *ListFlowInstanceResponseBody {
	s.FlowInstances = v
	return s
}

func (s *ListFlowInstanceResponseBody) SetPageNumber(v int32) *ListFlowInstanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowInstanceResponseBody) SetPageSize(v int32) *ListFlowInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowInstanceResponseBody) SetRequestId(v string) *ListFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowInstanceResponseBody) SetTotal(v int32) *ListFlowInstanceResponseBody {
	s.Total = &v
	return s
}

type ListFlowInstanceResponseBodyFlowInstances struct {
	FlowInstance []*ListFlowInstanceResponseBodyFlowInstancesFlowInstance `json:"FlowInstance,omitempty" xml:"FlowInstance,omitempty" type:"Repeated"`
}

func (s ListFlowInstanceResponseBodyFlowInstances) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceResponseBodyFlowInstances) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceResponseBodyFlowInstances) SetFlowInstance(v []*ListFlowInstanceResponseBodyFlowInstancesFlowInstance) *ListFlowInstanceResponseBodyFlowInstances {
	s.FlowInstance = v
	return s
}

type ListFlowInstanceResponseBodyFlowInstancesFlowInstance struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FlowId        *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName      *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HasNodeFailed *bool   `json:"HasNodeFailed,omitempty" xml:"HasNodeFailed,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Owner         *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ScheduleTime  *int64  `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFlowInstanceResponseBodyFlowInstancesFlowInstance) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceResponseBodyFlowInstancesFlowInstance) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetClusterId(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.ClusterId = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetDuration(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.Duration = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetEndTime(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.EndTime = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetFlowId(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.FlowId = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetFlowName(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.FlowName = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetGmtCreate(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetGmtModified(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.GmtModified = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetHasNodeFailed(v bool) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.HasNodeFailed = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetId(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.Id = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetOwner(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.Owner = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetProjectId(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.ProjectId = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetScheduleTime(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.ScheduleTime = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetStartTime(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.StartTime = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetStatus(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.Status = &v
	return s
}

type ListFlowInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceResponse) SetHeaders(v map[string]*string) *ListFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListFlowInstanceResponse) SetBody(v *ListFlowInstanceResponseBody) *ListFlowInstanceResponse {
	s.Body = v
	return s
}

type ListFlowJobHistoryRequest struct {
	Id         *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobType    *string   `json:"JobType,omitempty" xml:"JobType,omitempty"`
	PageNumber *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId  *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	TimeRange  *string   `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s ListFlowJobHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryRequest) SetId(v string) *ListFlowJobHistoryRequest {
	s.Id = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetInstanceId(v string) *ListFlowJobHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetJobType(v string) *ListFlowJobHistoryRequest {
	s.JobType = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetPageNumber(v int32) *ListFlowJobHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetPageSize(v int32) *ListFlowJobHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetProjectId(v string) *ListFlowJobHistoryRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetRegionId(v string) *ListFlowJobHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetStatusList(v []*string) *ListFlowJobHistoryRequest {
	s.StatusList = v
	return s
}

func (s *ListFlowJobHistoryRequest) SetTimeRange(v string) *ListFlowJobHistoryRequest {
	s.TimeRange = &v
	return s
}

type ListFlowJobHistoryResponseBody struct {
	NodeInstances *ListFlowJobHistoryResponseBodyNodeInstances `json:"NodeInstances,omitempty" xml:"NodeInstances,omitempty" type:"Struct"`
	PageNumber    *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total         *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFlowJobHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryResponseBody) SetNodeInstances(v *ListFlowJobHistoryResponseBodyNodeInstances) *ListFlowJobHistoryResponseBody {
	s.NodeInstances = v
	return s
}

func (s *ListFlowJobHistoryResponseBody) SetPageNumber(v int32) *ListFlowJobHistoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowJobHistoryResponseBody) SetPageSize(v int32) *ListFlowJobHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowJobHistoryResponseBody) SetRequestId(v string) *ListFlowJobHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBody) SetTotal(v int32) *ListFlowJobHistoryResponseBody {
	s.Total = &v
	return s
}

type ListFlowJobHistoryResponseBodyNodeInstances struct {
	NodeInstance []*ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance `json:"NodeInstance,omitempty" xml:"NodeInstance,omitempty" type:"Repeated"`
}

func (s ListFlowJobHistoryResponseBodyNodeInstances) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryResponseBodyNodeInstances) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryResponseBodyNodeInstances) SetNodeInstance(v []*ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) *ListFlowJobHistoryResponseBodyNodeInstances {
	s.NodeInstance = v
	return s
}

type ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EnvConf        *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	ExternalId     *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	ExternalInfo   *string `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty"`
	ExternalStatus *string `json:"ExternalStatus,omitempty" xml:"ExternalStatus,omitempty"`
	FailAct        *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HostName       *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName        *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobParams      *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	JobType        *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MaxRetry       *int32  `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	NodeName       *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	ParamConf      *string `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Retries        *int32  `json:"Retries,omitempty" xml:"Retries,omitempty"`
	RetryInterval  *int64  `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RunConf        *string `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Pending        *bool   `json:"pending,omitempty" xml:"pending,omitempty"`
}

func (s ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetClusterId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ClusterId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetEndTime(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.EndTime = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetEnvConf(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.EnvConf = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetExternalId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ExternalId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetExternalInfo(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ExternalInfo = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetExternalStatus(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ExternalStatus = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetFailAct(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.FailAct = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetGmtCreate(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetGmtModified(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.GmtModified = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetHostName(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.HostName = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Id = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetJobId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.JobId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetJobName(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.JobName = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetJobParams(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.JobParams = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetJobType(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.JobType = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetMaxRetry(v int32) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.MaxRetry = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetNodeName(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.NodeName = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetParamConf(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ParamConf = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetProjectId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ProjectId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetRetries(v int32) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Retries = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetRetryInterval(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.RetryInterval = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetRunConf(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.RunConf = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetStartTime(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.StartTime = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetStatus(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Status = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetType(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Type = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetPending(v bool) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Pending = &v
	return s
}

type ListFlowJobHistoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowJobHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowJobHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryResponse) SetHeaders(v map[string]*string) *ListFlowJobHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListFlowJobHistoryResponse) SetBody(v *ListFlowJobHistoryResponseBody) *ListFlowJobHistoryResponse {
	s.Body = v
	return s
}

type ListFlowJobsRequest struct {
	Adhoc      *bool   `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowJobsRequest) SetAdhoc(v bool) *ListFlowJobsRequest {
	s.Adhoc = &v
	return s
}

func (s *ListFlowJobsRequest) SetId(v string) *ListFlowJobsRequest {
	s.Id = &v
	return s
}

func (s *ListFlowJobsRequest) SetName(v string) *ListFlowJobsRequest {
	s.Name = &v
	return s
}

func (s *ListFlowJobsRequest) SetPageNumber(v int32) *ListFlowJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowJobsRequest) SetPageSize(v int32) *ListFlowJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowJobsRequest) SetProjectId(v string) *ListFlowJobsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowJobsRequest) SetRegionId(v string) *ListFlowJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowJobsRequest) SetType(v string) *ListFlowJobsRequest {
	s.Type = &v
	return s
}

type ListFlowJobsResponseBody struct {
	JobList    *ListFlowJobsResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Struct"`
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFlowJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBody) SetJobList(v *ListFlowJobsResponseBodyJobList) *ListFlowJobsResponseBody {
	s.JobList = v
	return s
}

func (s *ListFlowJobsResponseBody) SetPageNumber(v int32) *ListFlowJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowJobsResponseBody) SetPageSize(v int32) *ListFlowJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowJobsResponseBody) SetRequestId(v string) *ListFlowJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowJobsResponseBody) SetTotal(v int32) *ListFlowJobsResponseBody {
	s.Total = &v
	return s
}

type ListFlowJobsResponseBodyJobList struct {
	Job []*ListFlowJobsResponseBodyJobListJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Repeated"`
}

func (s ListFlowJobsResponseBodyJobList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBodyJobList) SetJob(v []*ListFlowJobsResponseBodyJobListJob) *ListFlowJobsResponseBodyJobList {
	s.Job = v
	return s
}

type ListFlowJobsResponseBodyJobListJob struct {
	Adhoc              *string                                         `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	AlertConf          *string                                         `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	CategoryId         *string                                         `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CustomVariables    *string                                         `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	Description        *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvConf            *string                                         `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	FailAct            *string                                         `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	GmtCreate          *int64                                          `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *int64                                          `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                 *string                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	LastInstanceDetail *string                                         `json:"LastInstanceDetail,omitempty" xml:"LastInstanceDetail,omitempty"`
	MaxRetry           *int32                                          `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	Mode               *string                                         `json:"Mode,omitempty" xml:"Mode,omitempty"`
	MonitorConf        *string                                         `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	Name               *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamConf          *string                                         `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	Params             *string                                         `json:"Params,omitempty" xml:"Params,omitempty"`
	ResourceList       *ListFlowJobsResponseBodyJobListJobResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Struct"`
	RetryInterval      *int64                                          `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RunConf            *string                                         `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	Type               *string                                         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowJobsResponseBodyJobListJob) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBodyJobListJob) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBodyJobListJob) SetAdhoc(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Adhoc = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetAlertConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.AlertConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetCategoryId(v string) *ListFlowJobsResponseBodyJobListJob {
	s.CategoryId = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetCustomVariables(v string) *ListFlowJobsResponseBodyJobListJob {
	s.CustomVariables = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetDescription(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Description = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetEnvConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.EnvConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetFailAct(v string) *ListFlowJobsResponseBodyJobListJob {
	s.FailAct = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetGmtCreate(v int64) *ListFlowJobsResponseBodyJobListJob {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetGmtModified(v int64) *ListFlowJobsResponseBodyJobListJob {
	s.GmtModified = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetId(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Id = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetLastInstanceDetail(v string) *ListFlowJobsResponseBodyJobListJob {
	s.LastInstanceDetail = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetMaxRetry(v int32) *ListFlowJobsResponseBodyJobListJob {
	s.MaxRetry = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetMode(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Mode = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetMonitorConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.MonitorConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetName(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Name = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetParamConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.ParamConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetParams(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Params = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetResourceList(v *ListFlowJobsResponseBodyJobListJobResourceList) *ListFlowJobsResponseBodyJobListJob {
	s.ResourceList = v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetRetryInterval(v int64) *ListFlowJobsResponseBodyJobListJob {
	s.RetryInterval = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetRunConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.RunConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetType(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Type = &v
	return s
}

type ListFlowJobsResponseBodyJobListJobResourceList struct {
	Resource []*ListFlowJobsResponseBodyJobListJobResourceListResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s ListFlowJobsResponseBodyJobListJobResourceList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBodyJobListJobResourceList) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBodyJobListJobResourceList) SetResource(v []*ListFlowJobsResponseBodyJobListJobResourceListResource) *ListFlowJobsResponseBodyJobListJobResourceList {
	s.Resource = v
	return s
}

type ListFlowJobsResponseBodyJobListJobResourceListResource struct {
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListFlowJobsResponseBodyJobListJobResourceListResource) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBodyJobListJobResourceListResource) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBodyJobListJobResourceListResource) SetAlias(v string) *ListFlowJobsResponseBodyJobListJobResourceListResource {
	s.Alias = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJobResourceListResource) SetPath(v string) *ListFlowJobsResponseBodyJobListJobResourceListResource {
	s.Path = &v
	return s
}

type ListFlowJobsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponse) SetHeaders(v map[string]*string) *ListFlowJobsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowJobsResponse) SetBody(v *ListFlowJobsResponseBody) *ListFlowJobsResponse {
	s.Body = v
	return s
}

type ListFlowNodeInstanceContainerStatusRequest struct {
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListFlowNodeInstanceContainerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusRequest) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetNodeInstanceId(v string) *ListFlowNodeInstanceContainerStatusRequest {
	s.NodeInstanceId = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetPageNumber(v int32) *ListFlowNodeInstanceContainerStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetPageSize(v int32) *ListFlowNodeInstanceContainerStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetProjectId(v string) *ListFlowNodeInstanceContainerStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetRegionId(v string) *ListFlowNodeInstanceContainerStatusRequest {
	s.RegionId = &v
	return s
}

type ListFlowNodeInstanceContainerStatusResponseBody struct {
	ContainerStatusList *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList `json:"ContainerStatusList,omitempty" xml:"ContainerStatusList,omitempty" type:"Struct"`
	PageNumber          *int32                                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total               *int32                                                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFlowNodeInstanceContainerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetContainerStatusList(v *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.ContainerStatusList = v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetPageNumber(v int32) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetPageSize(v int32) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetRequestId(v string) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetTotal(v int32) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.Total = &v
	return s
}

type ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList struct {
	ContainerStatus []*ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus `json:"ContainerStatus,omitempty" xml:"ContainerStatus,omitempty" type:"Repeated"`
}

func (s ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList) SetContainerStatus(v []*ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList {
	s.ContainerStatus = v
	return s
}

type ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ContainerId   *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	HostName      *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) SetApplicationId(v string) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus {
	s.ApplicationId = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) SetContainerId(v string) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus {
	s.ContainerId = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) SetHostName(v string) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus {
	s.HostName = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) SetStatus(v string) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus {
	s.Status = &v
	return s
}

type ListFlowNodeInstanceContainerStatusResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowNodeInstanceContainerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowNodeInstanceContainerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusResponse) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusResponse) SetHeaders(v map[string]*string) *ListFlowNodeInstanceContainerStatusResponse {
	s.Headers = v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponse) SetBody(v *ListFlowNodeInstanceContainerStatusResponseBody) *ListFlowNodeInstanceContainerStatusResponse {
	s.Body = v
	return s
}

type ListFlowNodeSqlResultRequest struct {
	Length         *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	Offset         *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SqlIndex       *int32  `json:"SqlIndex,omitempty" xml:"SqlIndex,omitempty"`
}

func (s ListFlowNodeSqlResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultRequest) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultRequest) SetLength(v int32) *ListFlowNodeSqlResultRequest {
	s.Length = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetNodeInstanceId(v string) *ListFlowNodeSqlResultRequest {
	s.NodeInstanceId = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetOffset(v int32) *ListFlowNodeSqlResultRequest {
	s.Offset = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetProjectId(v string) *ListFlowNodeSqlResultRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetRegionId(v string) *ListFlowNodeSqlResultRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetSqlIndex(v int32) *ListFlowNodeSqlResultRequest {
	s.SqlIndex = &v
	return s
}

type ListFlowNodeSqlResultResponseBody struct {
	End        *bool                                        `json:"End,omitempty" xml:"End,omitempty"`
	HeaderList *ListFlowNodeSqlResultResponseBodyHeaderList `json:"HeaderList,omitempty" xml:"HeaderList,omitempty" type:"Struct"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RowList    *ListFlowNodeSqlResultResponseBodyRowList    `json:"RowList,omitempty" xml:"RowList,omitempty" type:"Struct"`
}

func (s ListFlowNodeSqlResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBody) SetEnd(v bool) *ListFlowNodeSqlResultResponseBody {
	s.End = &v
	return s
}

func (s *ListFlowNodeSqlResultResponseBody) SetHeaderList(v *ListFlowNodeSqlResultResponseBodyHeaderList) *ListFlowNodeSqlResultResponseBody {
	s.HeaderList = v
	return s
}

func (s *ListFlowNodeSqlResultResponseBody) SetRequestId(v string) *ListFlowNodeSqlResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowNodeSqlResultResponseBody) SetRowList(v *ListFlowNodeSqlResultResponseBodyRowList) *ListFlowNodeSqlResultResponseBody {
	s.RowList = v
	return s
}

type ListFlowNodeSqlResultResponseBodyHeaderList struct {
	Header []*string `json:"Header,omitempty" xml:"Header,omitempty" type:"Repeated"`
}

func (s ListFlowNodeSqlResultResponseBodyHeaderList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBodyHeaderList) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBodyHeaderList) SetHeader(v []*string) *ListFlowNodeSqlResultResponseBodyHeaderList {
	s.Header = v
	return s
}

type ListFlowNodeSqlResultResponseBodyRowList struct {
	Row []*ListFlowNodeSqlResultResponseBodyRowListRow `json:"Row,omitempty" xml:"Row,omitempty" type:"Repeated"`
}

func (s ListFlowNodeSqlResultResponseBodyRowList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBodyRowList) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBodyRowList) SetRow(v []*ListFlowNodeSqlResultResponseBodyRowListRow) *ListFlowNodeSqlResultResponseBodyRowList {
	s.Row = v
	return s
}

type ListFlowNodeSqlResultResponseBodyRowListRow struct {
	RowIndex    *int32                                                  `json:"RowIndex,omitempty" xml:"RowIndex,omitempty"`
	RowItemList *ListFlowNodeSqlResultResponseBodyRowListRowRowItemList `json:"RowItemList,omitempty" xml:"RowItemList,omitempty" type:"Struct"`
}

func (s ListFlowNodeSqlResultResponseBodyRowListRow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBodyRowListRow) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBodyRowListRow) SetRowIndex(v int32) *ListFlowNodeSqlResultResponseBodyRowListRow {
	s.RowIndex = &v
	return s
}

func (s *ListFlowNodeSqlResultResponseBodyRowListRow) SetRowItemList(v *ListFlowNodeSqlResultResponseBodyRowListRowRowItemList) *ListFlowNodeSqlResultResponseBodyRowListRow {
	s.RowItemList = v
	return s
}

type ListFlowNodeSqlResultResponseBodyRowListRowRowItemList struct {
	RowItem []*string `json:"RowItem,omitempty" xml:"RowItem,omitempty" type:"Repeated"`
}

func (s ListFlowNodeSqlResultResponseBodyRowListRowRowItemList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBodyRowListRowRowItemList) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBodyRowListRowRowItemList) SetRowItem(v []*string) *ListFlowNodeSqlResultResponseBodyRowListRowRowItemList {
	s.RowItem = v
	return s
}

type ListFlowNodeSqlResultResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowNodeSqlResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowNodeSqlResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponse) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponse) SetHeaders(v map[string]*string) *ListFlowNodeSqlResultResponse {
	s.Headers = v
	return s
}

func (s *ListFlowNodeSqlResultResponse) SetBody(v *ListFlowNodeSqlResultResponseBody) *ListFlowNodeSqlResultResponse {
	s.Body = v
	return s
}

type ListFlowProjectClusterSettingRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingRequest) SetPageNumber(v int32) *ListFlowProjectClusterSettingRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectClusterSettingRequest) SetPageSize(v int32) *ListFlowProjectClusterSettingRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectClusterSettingRequest) SetProjectId(v string) *ListFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectClusterSettingRequest) SetRegionId(v string) *ListFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

type ListFlowProjectClusterSettingResponseBody struct {
	ClusterSettings *ListFlowProjectClusterSettingResponseBodyClusterSettings `json:"ClusterSettings,omitempty" xml:"ClusterSettings,omitempty" type:"Struct"`
	PageNumber      *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total           *int32                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBody) SetClusterSettings(v *ListFlowProjectClusterSettingResponseBodyClusterSettings) *ListFlowProjectClusterSettingResponseBody {
	s.ClusterSettings = v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBody) SetPageNumber(v int32) *ListFlowProjectClusterSettingResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBody) SetPageSize(v int32) *ListFlowProjectClusterSettingResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBody) SetRequestId(v string) *ListFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBody) SetTotal(v int32) *ListFlowProjectClusterSettingResponseBody {
	s.Total = &v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettings struct {
	ClusterSetting []*ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting `json:"ClusterSetting,omitempty" xml:"ClusterSetting,omitempty" type:"Repeated"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettings) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettings) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettings) SetClusterSetting(v []*ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) *ListFlowProjectClusterSettingResponseBodyClusterSettings {
	s.ClusterSetting = v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting struct {
	ClusterId    *string                                                                          `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName  *string                                                                          `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	DefaultQueue *string                                                                          `json:"DefaultQueue,omitempty" xml:"DefaultQueue,omitempty"`
	DefaultUser  *string                                                                          `json:"DefaultUser,omitempty" xml:"DefaultUser,omitempty"`
	GmtCreate    *int64                                                                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *int64                                                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HostList     *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList  `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Struct"`
	K8sClusterId *string                                                                          `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	ProjectId    *string                                                                          `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QueueList    *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Struct"`
	UserList     *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList  `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetClusterId(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.ClusterId = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetClusterName(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.ClusterName = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetDefaultQueue(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.DefaultQueue = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetDefaultUser(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.DefaultUser = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetGmtCreate(v int64) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetGmtModified(v int64) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.GmtModified = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetHostList(v *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.HostList = v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetK8sClusterId(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.K8sClusterId = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetProjectId(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetQueueList(v *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.QueueList = v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetUserList(v *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.UserList = v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList struct {
	Host []*string `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList) SetHost(v []*string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList {
	s.Host = v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList struct {
	Queue []*string `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Repeated"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList) SetQueue(v []*string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList {
	s.Queue = v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList struct {
	User []*string `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList) SetUser(v []*string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList {
	s.User = v
	return s
}

type ListFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *ListFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *ListFlowProjectClusterSettingResponse) SetBody(v *ListFlowProjectClusterSettingResponseBody) *ListFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type ListFlowProjectUserRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListFlowProjectUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserRequest) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserRequest) SetPageNumber(v int32) *ListFlowProjectUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectUserRequest) SetPageSize(v int32) *ListFlowProjectUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectUserRequest) SetProjectId(v string) *ListFlowProjectUserRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectUserRequest) SetRegionId(v string) *ListFlowProjectUserRequest {
	s.RegionId = &v
	return s
}

type ListFlowProjectUserResponseBody struct {
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Users      *ListFlowProjectUserResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListFlowProjectUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserResponseBody) SetPageNumber(v int32) *ListFlowProjectUserResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectUserResponseBody) SetPageSize(v int32) *ListFlowProjectUserResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectUserResponseBody) SetRequestId(v string) *ListFlowProjectUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowProjectUserResponseBody) SetTotal(v int32) *ListFlowProjectUserResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowProjectUserResponseBody) SetUsers(v *ListFlowProjectUserResponseBodyUsers) *ListFlowProjectUserResponseBody {
	s.Users = v
	return s
}

type ListFlowProjectUserResponseBodyUsers struct {
	User []*ListFlowProjectUserResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListFlowProjectUserResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserResponseBodyUsers) SetUser(v []*ListFlowProjectUserResponseBodyUsersUser) *ListFlowProjectUserResponseBodyUsers {
	s.User = v
	return s
}

type ListFlowProjectUserResponseBodyUsersUser struct {
	AccountUserId *string `json:"AccountUserId,omitempty" xml:"AccountUserId,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListFlowProjectUserResponseBodyUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetAccountUserId(v string) *ListFlowProjectUserResponseBodyUsersUser {
	s.AccountUserId = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetGmtCreate(v int64) *ListFlowProjectUserResponseBodyUsersUser {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetGmtModified(v int64) *ListFlowProjectUserResponseBodyUsersUser {
	s.GmtModified = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetOwnerId(v string) *ListFlowProjectUserResponseBodyUsersUser {
	s.OwnerId = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetProjectId(v string) *ListFlowProjectUserResponseBodyUsersUser {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetUserName(v string) *ListFlowProjectUserResponseBodyUsersUser {
	s.UserName = &v
	return s
}

type ListFlowProjectUserResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowProjectUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowProjectUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserResponse) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserResponse) SetHeaders(v map[string]*string) *ListFlowProjectUserResponse {
	s.Headers = v
	return s
}

func (s *ListFlowProjectUserResponse) SetBody(v *ListFlowProjectUserResponseBody) *ListFlowProjectUserResponse {
	s.Body = v
	return s
}

type ListFlowProjectsRequest struct {
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType     *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsRequest) SetName(v string) *ListFlowProjectsRequest {
	s.Name = &v
	return s
}

func (s *ListFlowProjectsRequest) SetPageNumber(v int32) *ListFlowProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectsRequest) SetPageSize(v int32) *ListFlowProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectsRequest) SetProductType(v string) *ListFlowProjectsRequest {
	s.ProductType = &v
	return s
}

func (s *ListFlowProjectsRequest) SetProjectId(v string) *ListFlowProjectsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectsRequest) SetRegionId(v string) *ListFlowProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowProjectsRequest) SetResourceGroupId(v string) *ListFlowProjectsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowProjectsResponseBody struct {
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Projects   *ListFlowProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Struct"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFlowProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsResponseBody) SetPageNumber(v int32) *ListFlowProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectsResponseBody) SetPageSize(v int32) *ListFlowProjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectsResponseBody) SetProjects(v *ListFlowProjectsResponseBodyProjects) *ListFlowProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListFlowProjectsResponseBody) SetRequestId(v string) *ListFlowProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowProjectsResponseBody) SetTotal(v int32) *ListFlowProjectsResponseBody {
	s.Total = &v
	return s
}

type ListFlowProjectsResponseBodyProjects struct {
	Project []*ListFlowProjectsResponseBodyProjectsProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Repeated"`
}

func (s ListFlowProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsResponseBodyProjects) SetProject(v []*ListFlowProjectsResponseBodyProjectsProject) *ListFlowProjectsResponseBodyProjects {
	s.Project = v
	return s
}

type ListFlowProjectsResponseBodyProjectsProject struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListFlowProjectsResponseBodyProjectsProject) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsResponseBodyProjectsProject) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetDescription(v string) *ListFlowProjectsResponseBodyProjectsProject {
	s.Description = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetGmtCreate(v int64) *ListFlowProjectsResponseBodyProjectsProject {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetGmtModified(v int64) *ListFlowProjectsResponseBodyProjectsProject {
	s.GmtModified = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetId(v string) *ListFlowProjectsResponseBodyProjectsProject {
	s.Id = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetName(v string) *ListFlowProjectsResponseBodyProjectsProject {
	s.Name = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetUserId(v string) *ListFlowProjectsResponseBodyProjectsProject {
	s.UserId = &v
	return s
}

type ListFlowProjectsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsResponse) SetHeaders(v map[string]*string) *ListFlowProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowProjectsResponse) SetBody(v *ListFlowProjectsResponseBody) *ListFlowProjectsResponse {
	s.Body = v
	return s
}

type ListFlowsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Periodic   *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowsRequest) SetClusterId(v string) *ListFlowsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListFlowsRequest) SetId(v string) *ListFlowsRequest {
	s.Id = &v
	return s
}

func (s *ListFlowsRequest) SetJobId(v string) *ListFlowsRequest {
	s.JobId = &v
	return s
}

func (s *ListFlowsRequest) SetName(v string) *ListFlowsRequest {
	s.Name = &v
	return s
}

func (s *ListFlowsRequest) SetPageNumber(v int32) *ListFlowsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowsRequest) SetPageSize(v int32) *ListFlowsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowsRequest) SetPeriodic(v bool) *ListFlowsRequest {
	s.Periodic = &v
	return s
}

func (s *ListFlowsRequest) SetProjectId(v string) *ListFlowsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowsRequest) SetRegionId(v string) *ListFlowsRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowsRequest) SetStatus(v string) *ListFlowsRequest {
	s.Status = &v
	return s
}

type ListFlowsResponseBody struct {
	Flow       *ListFlowsResponseBodyFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Struct"`
	PageNumber *int32                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBody) SetFlow(v *ListFlowsResponseBodyFlow) *ListFlowsResponseBody {
	s.Flow = v
	return s
}

func (s *ListFlowsResponseBody) SetPageNumber(v int32) *ListFlowsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowsResponseBody) SetPageSize(v int32) *ListFlowsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowsResponseBody) SetRequestId(v string) *ListFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowsResponseBody) SetTotal(v int32) *ListFlowsResponseBody {
	s.Total = &v
	return s
}

type ListFlowsResponseBodyFlow struct {
	Flow []*ListFlowsResponseBodyFlowFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Repeated"`
}

func (s ListFlowsResponseBodyFlow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBodyFlow) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlow) SetFlow(v []*ListFlowsResponseBodyFlowFlow) *ListFlowsResponseBodyFlow {
	s.Flow = v
	return s
}

type ListFlowsResponseBodyFlowFlow struct {
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	CategoryId              *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	CronExpr                *string `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	GmtCreate               *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified             *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Graph                   *string `json:"Graph,omitempty" xml:"Graph,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Periodic                *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowsResponseBodyFlowFlow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBodyFlowFlow) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlowFlow) SetAlertConf(v string) *ListFlowsResponseBodyFlowFlow {
	s.AlertConf = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetAlertDingDingGroupBizId(v string) *ListFlowsResponseBodyFlowFlow {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetAlertUserGroupBizId(v string) *ListFlowsResponseBodyFlowFlow {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetCategoryId(v string) *ListFlowsResponseBodyFlowFlow {
	s.CategoryId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetClusterId(v string) *ListFlowsResponseBodyFlowFlow {
	s.ClusterId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetCreateCluster(v bool) *ListFlowsResponseBodyFlowFlow {
	s.CreateCluster = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetCronExpr(v string) *ListFlowsResponseBodyFlowFlow {
	s.CronExpr = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetDescription(v string) *ListFlowsResponseBodyFlowFlow {
	s.Description = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetEndSchedule(v int64) *ListFlowsResponseBodyFlowFlow {
	s.EndSchedule = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetGmtCreate(v int64) *ListFlowsResponseBodyFlowFlow {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetGmtModified(v int64) *ListFlowsResponseBodyFlowFlow {
	s.GmtModified = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetGraph(v string) *ListFlowsResponseBodyFlowFlow {
	s.Graph = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetHostName(v string) *ListFlowsResponseBodyFlowFlow {
	s.HostName = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetId(v string) *ListFlowsResponseBodyFlowFlow {
	s.Id = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetName(v string) *ListFlowsResponseBodyFlowFlow {
	s.Name = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetPeriodic(v bool) *ListFlowsResponseBodyFlowFlow {
	s.Periodic = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetProjectId(v string) *ListFlowsResponseBodyFlowFlow {
	s.ProjectId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetStartSchedule(v int64) *ListFlowsResponseBodyFlowFlow {
	s.StartSchedule = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetStatus(v string) *ListFlowsResponseBodyFlowFlow {
	s.Status = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetType(v string) *ListFlowsResponseBodyFlowFlow {
	s.Type = &v
	return s
}

type ListFlowsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowsResponse) SetHeaders(v map[string]*string) *ListFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowsResponse) SetBody(v *ListFlowsResponseBody) *ListFlowsResponse {
	s.Body = v
	return s
}

type ListMainVersionsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListMainVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMainVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListMainVersionsRequest) SetRegionId(v string) *ListMainVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListMainVersionsRequest) SetResourceGroupId(v string) *ListMainVersionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListMainVersionsRequest) SetResourceOwnerId(v int64) *ListMainVersionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListMainVersionsResponseBody struct {
	MainVersionList []*ListMainVersionsResponseBodyMainVersionList `json:"MainVersionList,omitempty" xml:"MainVersionList,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMainVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMainVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMainVersionsResponseBody) SetMainVersionList(v []*ListMainVersionsResponseBodyMainVersionList) *ListMainVersionsResponseBody {
	s.MainVersionList = v
	return s
}

func (s *ListMainVersionsResponseBody) SetRequestId(v string) *ListMainVersionsResponseBody {
	s.RequestId = &v
	return s
}

type ListMainVersionsResponseBodyMainVersionList struct {
	ClusterTypeInfoList []*ListMainVersionsResponseBodyMainVersionListClusterTypeInfoList `json:"ClusterTypeInfoList,omitempty" xml:"ClusterTypeInfoList,omitempty" type:"Repeated"`
	ExtraInfo           *string                                                           `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	MainVersionName     *string                                                           `json:"MainVersionName,omitempty" xml:"MainVersionName,omitempty"`
	RegionId            *string                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListMainVersionsResponseBodyMainVersionList) String() string {
	return tea.Prettify(s)
}

func (s ListMainVersionsResponseBodyMainVersionList) GoString() string {
	return s.String()
}

func (s *ListMainVersionsResponseBodyMainVersionList) SetClusterTypeInfoList(v []*ListMainVersionsResponseBodyMainVersionListClusterTypeInfoList) *ListMainVersionsResponseBodyMainVersionList {
	s.ClusterTypeInfoList = v
	return s
}

func (s *ListMainVersionsResponseBodyMainVersionList) SetExtraInfo(v string) *ListMainVersionsResponseBodyMainVersionList {
	s.ExtraInfo = &v
	return s
}

func (s *ListMainVersionsResponseBodyMainVersionList) SetMainVersionName(v string) *ListMainVersionsResponseBodyMainVersionList {
	s.MainVersionName = &v
	return s
}

func (s *ListMainVersionsResponseBodyMainVersionList) SetRegionId(v string) *ListMainVersionsResponseBodyMainVersionList {
	s.RegionId = &v
	return s
}

type ListMainVersionsResponseBodyMainVersionListClusterTypeInfoList struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
}

func (s ListMainVersionsResponseBodyMainVersionListClusterTypeInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListMainVersionsResponseBodyMainVersionListClusterTypeInfoList) GoString() string {
	return s.String()
}

func (s *ListMainVersionsResponseBodyMainVersionListClusterTypeInfoList) SetClusterType(v string) *ListMainVersionsResponseBodyMainVersionListClusterTypeInfoList {
	s.ClusterType = &v
	return s
}

type ListMainVersionsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMainVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMainVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMainVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListMainVersionsResponse) SetHeaders(v map[string]*string) *ListMainVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListMainVersionsResponse) SetBody(v *ListMainVersionsResponseBody) *ListMainVersionsResponse {
	s.Body = v
	return s
}

type ModifyFlowRequest struct {
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	Application             *string `json:"Application,omitempty" xml:"Application,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	CronExpr                *string `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentCategory          *string `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	ParentFlowList          *string `json:"ParentFlowList,omitempty" xml:"ParentFlowList,omitempty"`
	Periodic                *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowRequest) SetAlertConf(v string) *ModifyFlowRequest {
	s.AlertConf = &v
	return s
}

func (s *ModifyFlowRequest) SetAlertDingDingGroupBizId(v string) *ModifyFlowRequest {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *ModifyFlowRequest) SetAlertUserGroupBizId(v string) *ModifyFlowRequest {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *ModifyFlowRequest) SetApplication(v string) *ModifyFlowRequest {
	s.Application = &v
	return s
}

func (s *ModifyFlowRequest) SetClusterId(v string) *ModifyFlowRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyFlowRequest) SetCreateCluster(v bool) *ModifyFlowRequest {
	s.CreateCluster = &v
	return s
}

func (s *ModifyFlowRequest) SetCronExpr(v string) *ModifyFlowRequest {
	s.CronExpr = &v
	return s
}

func (s *ModifyFlowRequest) SetDescription(v string) *ModifyFlowRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowRequest) SetEndSchedule(v int64) *ModifyFlowRequest {
	s.EndSchedule = &v
	return s
}

func (s *ModifyFlowRequest) SetHostName(v string) *ModifyFlowRequest {
	s.HostName = &v
	return s
}

func (s *ModifyFlowRequest) SetId(v string) *ModifyFlowRequest {
	s.Id = &v
	return s
}

func (s *ModifyFlowRequest) SetName(v string) *ModifyFlowRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowRequest) SetParentCategory(v string) *ModifyFlowRequest {
	s.ParentCategory = &v
	return s
}

func (s *ModifyFlowRequest) SetParentFlowList(v string) *ModifyFlowRequest {
	s.ParentFlowList = &v
	return s
}

func (s *ModifyFlowRequest) SetPeriodic(v bool) *ModifyFlowRequest {
	s.Periodic = &v
	return s
}

func (s *ModifyFlowRequest) SetProjectId(v string) *ModifyFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowRequest) SetRegionId(v string) *ModifyFlowRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowRequest) SetStartSchedule(v int64) *ModifyFlowRequest {
	s.StartSchedule = &v
	return s
}

func (s *ModifyFlowRequest) SetStatus(v string) *ModifyFlowRequest {
	s.Status = &v
	return s
}

type ModifyFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponseBody) SetData(v bool) *ModifyFlowResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowResponseBody) SetRequestId(v string) *ModifyFlowResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponse) SetHeaders(v map[string]*string) *ModifyFlowResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowResponse) SetBody(v *ModifyFlowResponseBody) *ModifyFlowResponse {
	s.Body = v
	return s
}

type ModifyFlowCategoryRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId  *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyFlowCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowCategoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowCategoryRequest) SetId(v string) *ModifyFlowCategoryRequest {
	s.Id = &v
	return s
}

func (s *ModifyFlowCategoryRequest) SetName(v string) *ModifyFlowCategoryRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowCategoryRequest) SetParentId(v string) *ModifyFlowCategoryRequest {
	s.ParentId = &v
	return s
}

func (s *ModifyFlowCategoryRequest) SetProjectId(v string) *ModifyFlowCategoryRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowCategoryRequest) SetRegionId(v string) *ModifyFlowCategoryRequest {
	s.RegionId = &v
	return s
}

type ModifyFlowCategoryResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowCategoryResponseBody) SetData(v bool) *ModifyFlowCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowCategoryResponseBody) SetRequestId(v string) *ModifyFlowCategoryResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowCategoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowCategoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowCategoryResponse) SetHeaders(v map[string]*string) *ModifyFlowCategoryResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowCategoryResponse) SetBody(v *ModifyFlowCategoryResponseBody) *ModifyFlowCategoryResponse {
	s.Body = v
	return s
}

type ModifyFlowForWebRequest struct {
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	CronExpr                *string `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	Graph                   *string `json:"Graph,omitempty" xml:"Graph,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace               *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParentCategory          *string `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	ParentFlowList          *string `json:"ParentFlowList,omitempty" xml:"ParentFlowList,omitempty"`
	Periodic                *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyFlowForWebRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowForWebRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowForWebRequest) SetAlertConf(v string) *ModifyFlowForWebRequest {
	s.AlertConf = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetAlertDingDingGroupBizId(v string) *ModifyFlowForWebRequest {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetAlertUserGroupBizId(v string) *ModifyFlowForWebRequest {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetClusterId(v string) *ModifyFlowForWebRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetCreateCluster(v bool) *ModifyFlowForWebRequest {
	s.CreateCluster = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetCronExpr(v string) *ModifyFlowForWebRequest {
	s.CronExpr = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetDescription(v string) *ModifyFlowForWebRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetEndSchedule(v int64) *ModifyFlowForWebRequest {
	s.EndSchedule = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetGraph(v string) *ModifyFlowForWebRequest {
	s.Graph = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetHostName(v string) *ModifyFlowForWebRequest {
	s.HostName = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetId(v string) *ModifyFlowForWebRequest {
	s.Id = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetName(v string) *ModifyFlowForWebRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetNamespace(v string) *ModifyFlowForWebRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetParentCategory(v string) *ModifyFlowForWebRequest {
	s.ParentCategory = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetParentFlowList(v string) *ModifyFlowForWebRequest {
	s.ParentFlowList = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetPeriodic(v bool) *ModifyFlowForWebRequest {
	s.Periodic = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetProjectId(v string) *ModifyFlowForWebRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetRegionId(v string) *ModifyFlowForWebRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetStartSchedule(v int64) *ModifyFlowForWebRequest {
	s.StartSchedule = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetStatus(v string) *ModifyFlowForWebRequest {
	s.Status = &v
	return s
}

type ModifyFlowForWebResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowForWebResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowForWebResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowForWebResponseBody) SetData(v bool) *ModifyFlowForWebResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowForWebResponseBody) SetRequestId(v string) *ModifyFlowForWebResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowForWebResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowForWebResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowForWebResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowForWebResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowForWebResponse) SetHeaders(v map[string]*string) *ModifyFlowForWebResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowForWebResponse) SetBody(v *ModifyFlowForWebResponseBody) *ModifyFlowForWebResponse {
	s.Body = v
	return s
}

type ModifyFlowJobRequest struct {
	AlertConf       *string                             `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	ClusterId       *string                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CustomVariables *string                             `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	Description     *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvConf         *string                             `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	FailAct         *string                             `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	Id              *string                             `json:"Id,omitempty" xml:"Id,omitempty"`
	KnoxPassword    *string                             `json:"KnoxPassword,omitempty" xml:"KnoxPassword,omitempty"`
	KnoxUser        *string                             `json:"KnoxUser,omitempty" xml:"KnoxUser,omitempty"`
	Mode            *string                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
	MonitorConf     *string                             `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	Name            *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamConf       *string                             `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	Params          *string                             `json:"Params,omitempty" xml:"Params,omitempty"`
	ProjectId       *string                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceList    []*ModifyFlowJobRequestResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	RetryPolicy     *string                             `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	RunConf         *string                             `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
}

func (s ModifyFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowJobRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowJobRequest) SetAlertConf(v string) *ModifyFlowJobRequest {
	s.AlertConf = &v
	return s
}

func (s *ModifyFlowJobRequest) SetClusterId(v string) *ModifyFlowJobRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyFlowJobRequest) SetCustomVariables(v string) *ModifyFlowJobRequest {
	s.CustomVariables = &v
	return s
}

func (s *ModifyFlowJobRequest) SetDescription(v string) *ModifyFlowJobRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowJobRequest) SetEnvConf(v string) *ModifyFlowJobRequest {
	s.EnvConf = &v
	return s
}

func (s *ModifyFlowJobRequest) SetFailAct(v string) *ModifyFlowJobRequest {
	s.FailAct = &v
	return s
}

func (s *ModifyFlowJobRequest) SetId(v string) *ModifyFlowJobRequest {
	s.Id = &v
	return s
}

func (s *ModifyFlowJobRequest) SetKnoxPassword(v string) *ModifyFlowJobRequest {
	s.KnoxPassword = &v
	return s
}

func (s *ModifyFlowJobRequest) SetKnoxUser(v string) *ModifyFlowJobRequest {
	s.KnoxUser = &v
	return s
}

func (s *ModifyFlowJobRequest) SetMode(v string) *ModifyFlowJobRequest {
	s.Mode = &v
	return s
}

func (s *ModifyFlowJobRequest) SetMonitorConf(v string) *ModifyFlowJobRequest {
	s.MonitorConf = &v
	return s
}

func (s *ModifyFlowJobRequest) SetName(v string) *ModifyFlowJobRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowJobRequest) SetParamConf(v string) *ModifyFlowJobRequest {
	s.ParamConf = &v
	return s
}

func (s *ModifyFlowJobRequest) SetParams(v string) *ModifyFlowJobRequest {
	s.Params = &v
	return s
}

func (s *ModifyFlowJobRequest) SetProjectId(v string) *ModifyFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowJobRequest) SetRegionId(v string) *ModifyFlowJobRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowJobRequest) SetResourceList(v []*ModifyFlowJobRequestResourceList) *ModifyFlowJobRequest {
	s.ResourceList = v
	return s
}

func (s *ModifyFlowJobRequest) SetRetryPolicy(v string) *ModifyFlowJobRequest {
	s.RetryPolicy = &v
	return s
}

func (s *ModifyFlowJobRequest) SetRunConf(v string) *ModifyFlowJobRequest {
	s.RunConf = &v
	return s
}

type ModifyFlowJobRequestResourceList struct {
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ModifyFlowJobRequestResourceList) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowJobRequestResourceList) GoString() string {
	return s.String()
}

func (s *ModifyFlowJobRequestResourceList) SetAlias(v string) *ModifyFlowJobRequestResourceList {
	s.Alias = &v
	return s
}

func (s *ModifyFlowJobRequestResourceList) SetPath(v string) *ModifyFlowJobRequestResourceList {
	s.Path = &v
	return s
}

type ModifyFlowJobResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowJobResponseBody) SetData(v bool) *ModifyFlowJobResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowJobResponseBody) SetRequestId(v string) *ModifyFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowJobResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowJobResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowJobResponse) SetHeaders(v map[string]*string) *ModifyFlowJobResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowJobResponse) SetBody(v *ModifyFlowJobResponseBody) *ModifyFlowJobResponse {
	s.Body = v
	return s
}

type ModifyFlowProjectRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyFlowProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectRequest) SetDescription(v string) *ModifyFlowProjectRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowProjectRequest) SetName(v string) *ModifyFlowProjectRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowProjectRequest) SetProjectId(v string) *ModifyFlowProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowProjectRequest) SetRegionId(v string) *ModifyFlowProjectRequest {
	s.RegionId = &v
	return s
}

type ModifyFlowProjectResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectResponseBody) SetData(v bool) *ModifyFlowProjectResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowProjectResponseBody) SetRequestId(v string) *ModifyFlowProjectResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectResponse) SetHeaders(v map[string]*string) *ModifyFlowProjectResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowProjectResponse) SetBody(v *ModifyFlowProjectResponseBody) *ModifyFlowProjectResponse {
	s.Body = v
	return s
}

type ModifyFlowProjectClusterSettingRequest struct {
	ClusterId    *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DefaultQueue *string   `json:"DefaultQueue,omitempty" xml:"DefaultQueue,omitempty"`
	DefaultUser  *string   `json:"DefaultUser,omitempty" xml:"DefaultUser,omitempty"`
	HostList     []*string `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Repeated"`
	ProjectId    *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QueueList    []*string `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserList     []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ModifyFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectClusterSettingRequest) SetClusterId(v string) *ModifyFlowProjectClusterSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetDefaultQueue(v string) *ModifyFlowProjectClusterSettingRequest {
	s.DefaultQueue = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetDefaultUser(v string) *ModifyFlowProjectClusterSettingRequest {
	s.DefaultUser = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetHostList(v []*string) *ModifyFlowProjectClusterSettingRequest {
	s.HostList = v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetProjectId(v string) *ModifyFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetQueueList(v []*string) *ModifyFlowProjectClusterSettingRequest {
	s.QueueList = v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetRegionId(v string) *ModifyFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetUserList(v []*string) *ModifyFlowProjectClusterSettingRequest {
	s.UserList = v
	return s
}

type ModifyFlowProjectClusterSettingResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectClusterSettingResponseBody) SetData(v bool) *ModifyFlowProjectClusterSettingResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingResponseBody) SetRequestId(v string) *ModifyFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *ModifyFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowProjectClusterSettingResponse) SetBody(v *ModifyFlowProjectClusterSettingResponseBody) *ModifyFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type ModifyFlowVariableCollectionRequest struct {
	Data            *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyFlowVariableCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowVariableCollectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowVariableCollectionRequest) SetData(v string) *ModifyFlowVariableCollectionRequest {
	s.Data = &v
	return s
}

func (s *ModifyFlowVariableCollectionRequest) SetRegionId(v string) *ModifyFlowVariableCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowVariableCollectionRequest) SetResourceGroupId(v string) *ModifyFlowVariableCollectionRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyFlowVariableCollectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowVariableCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowVariableCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowVariableCollectionResponseBody) SetRequestId(v string) *ModifyFlowVariableCollectionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowVariableCollectionResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowVariableCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowVariableCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowVariableCollectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowVariableCollectionResponse) SetHeaders(v map[string]*string) *ModifyFlowVariableCollectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowVariableCollectionResponse) SetBody(v *ModifyFlowVariableCollectionResponseBody) *ModifyFlowVariableCollectionResponse {
	s.Body = v
	return s
}

type ReleaseClusterRequest struct {
	ForceRelease    *bool   `json:"ForceRelease,omitempty" xml:"ForceRelease,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterRequest) GoString() string {
	return s.String()
}

func (s *ReleaseClusterRequest) SetForceRelease(v bool) *ReleaseClusterRequest {
	s.ForceRelease = &v
	return s
}

func (s *ReleaseClusterRequest) SetId(v string) *ReleaseClusterRequest {
	s.Id = &v
	return s
}

func (s *ReleaseClusterRequest) SetRegionId(v string) *ReleaseClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseClusterRequest) SetResourceOwnerId(v int64) *ReleaseClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type ReleaseClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseClusterResponseBody) SetRequestId(v string) *ReleaseClusterResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseClusterResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterResponse) GoString() string {
	return s.String()
}

func (s *ReleaseClusterResponse) SetHeaders(v map[string]*string) *ReleaseClusterResponse {
	s.Headers = v
	return s
}

func (s *ReleaseClusterResponse) SetBody(v *ReleaseClusterResponseBody) *ReleaseClusterResponse {
	s.Body = v
	return s
}

type RerunFlowRequest struct {
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ReRunFail      *bool   `json:"ReRunFail,omitempty" xml:"ReRunFail,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RerunFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s RerunFlowRequest) GoString() string {
	return s.String()
}

func (s *RerunFlowRequest) SetFlowInstanceId(v string) *RerunFlowRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *RerunFlowRequest) SetProjectId(v string) *RerunFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *RerunFlowRequest) SetReRunFail(v bool) *RerunFlowRequest {
	s.ReRunFail = &v
	return s
}

func (s *RerunFlowRequest) SetRegionId(v string) *RerunFlowRequest {
	s.RegionId = &v
	return s
}

type RerunFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RerunFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RerunFlowResponseBody) GoString() string {
	return s.String()
}

func (s *RerunFlowResponseBody) SetData(v bool) *RerunFlowResponseBody {
	s.Data = &v
	return s
}

func (s *RerunFlowResponseBody) SetRequestId(v string) *RerunFlowResponseBody {
	s.RequestId = &v
	return s
}

type RerunFlowResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RerunFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RerunFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s RerunFlowResponse) GoString() string {
	return s.String()
}

func (s *RerunFlowResponse) SetHeaders(v map[string]*string) *RerunFlowResponse {
	s.Headers = v
	return s
}

func (s *RerunFlowResponse) SetBody(v *RerunFlowResponseBody) *RerunFlowResponse {
	s.Body = v
	return s
}

type RestoreFlowEntitySnapshotRequest struct {
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType      *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	OperatorId      *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Revision        *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
}

func (s RestoreFlowEntitySnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreFlowEntitySnapshotRequest) GoString() string {
	return s.String()
}

func (s *RestoreFlowEntitySnapshotRequest) SetEntityId(v string) *RestoreFlowEntitySnapshotRequest {
	s.EntityId = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetEntityType(v string) *RestoreFlowEntitySnapshotRequest {
	s.EntityType = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetOperatorId(v string) *RestoreFlowEntitySnapshotRequest {
	s.OperatorId = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetRegionId(v string) *RestoreFlowEntitySnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetResourceOwnerId(v int64) *RestoreFlowEntitySnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetRevision(v string) *RestoreFlowEntitySnapshotRequest {
	s.Revision = &v
	return s
}

type RestoreFlowEntitySnapshotResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreFlowEntitySnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreFlowEntitySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreFlowEntitySnapshotResponseBody) SetData(v bool) *RestoreFlowEntitySnapshotResponseBody {
	s.Data = &v
	return s
}

func (s *RestoreFlowEntitySnapshotResponseBody) SetRequestId(v string) *RestoreFlowEntitySnapshotResponseBody {
	s.RequestId = &v
	return s
}

type RestoreFlowEntitySnapshotResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestoreFlowEntitySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestoreFlowEntitySnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreFlowEntitySnapshotResponse) GoString() string {
	return s.String()
}

func (s *RestoreFlowEntitySnapshotResponse) SetHeaders(v map[string]*string) *RestoreFlowEntitySnapshotResponse {
	s.Headers = v
	return s
}

func (s *RestoreFlowEntitySnapshotResponse) SetBody(v *RestoreFlowEntitySnapshotResponseBody) *RestoreFlowEntitySnapshotResponse {
	s.Body = v
	return s
}

type ResumeFlowRequest struct {
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResumeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeFlowRequest) GoString() string {
	return s.String()
}

func (s *ResumeFlowRequest) SetFlowInstanceId(v string) *ResumeFlowRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *ResumeFlowRequest) SetProjectId(v string) *ResumeFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *ResumeFlowRequest) SetRegionId(v string) *ResumeFlowRequest {
	s.RegionId = &v
	return s
}

type ResumeFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeFlowResponseBody) SetData(v bool) *ResumeFlowResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeFlowResponseBody) SetRequestId(v string) *ResumeFlowResponseBody {
	s.RequestId = &v
	return s
}

type ResumeFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeFlowResponse) GoString() string {
	return s.String()
}

func (s *ResumeFlowResponse) SetHeaders(v map[string]*string) *ResumeFlowResponse {
	s.Headers = v
	return s
}

func (s *ResumeFlowResponse) SetBody(v *ResumeFlowResponseBody) *ResumeFlowResponse {
	s.Body = v
	return s
}

type StartFlowRequest struct {
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s StartFlowRequest) GoString() string {
	return s.String()
}

func (s *StartFlowRequest) SetFlowInstanceId(v string) *StartFlowRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *StartFlowRequest) SetProjectId(v string) *StartFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *StartFlowRequest) SetRegionId(v string) *StartFlowRequest {
	s.RegionId = &v
	return s
}

type StartFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartFlowResponseBody) GoString() string {
	return s.String()
}

func (s *StartFlowResponseBody) SetData(v bool) *StartFlowResponseBody {
	s.Data = &v
	return s
}

func (s *StartFlowResponseBody) SetRequestId(v string) *StartFlowResponseBody {
	s.RequestId = &v
	return s
}

type StartFlowResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s StartFlowResponse) GoString() string {
	return s.String()
}

func (s *StartFlowResponse) SetHeaders(v map[string]*string) *StartFlowResponse {
	s.Headers = v
	return s
}

func (s *StartFlowResponse) SetBody(v *StartFlowResponseBody) *StartFlowResponse {
	s.Body = v
	return s
}

type SubmitFlowRequest struct {
	Conf      *string `json:"Conf,omitempty" xml:"Conf,omitempty"`
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SubmitFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowRequest) GoString() string {
	return s.String()
}

func (s *SubmitFlowRequest) SetConf(v string) *SubmitFlowRequest {
	s.Conf = &v
	return s
}

func (s *SubmitFlowRequest) SetFlowId(v string) *SubmitFlowRequest {
	s.FlowId = &v
	return s
}

func (s *SubmitFlowRequest) SetProjectId(v string) *SubmitFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitFlowRequest) SetRegionId(v string) *SubmitFlowRequest {
	s.RegionId = &v
	return s
}

type SubmitFlowResponseBody struct {
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFlowResponseBody) SetData(v string) *SubmitFlowResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitFlowResponseBody) SetId(v string) *SubmitFlowResponseBody {
	s.Id = &v
	return s
}

func (s *SubmitFlowResponseBody) SetInstanceId(v string) *SubmitFlowResponseBody {
	s.InstanceId = &v
	return s
}

func (s *SubmitFlowResponseBody) SetRequestId(v string) *SubmitFlowResponseBody {
	s.RequestId = &v
	return s
}

type SubmitFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowResponse) GoString() string {
	return s.String()
}

func (s *SubmitFlowResponse) SetHeaders(v map[string]*string) *SubmitFlowResponse {
	s.Headers = v
	return s
}

func (s *SubmitFlowResponse) SetBody(v *SubmitFlowResponseBody) *SubmitFlowResponse {
	s.Body = v
	return s
}

type SubmitFlowJobRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Conf      *string `json:"Conf,omitempty" xml:"Conf,omitempty"`
	HostName  *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SubmitFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitFlowJobRequest) SetClusterId(v string) *SubmitFlowJobRequest {
	s.ClusterId = &v
	return s
}

func (s *SubmitFlowJobRequest) SetConf(v string) *SubmitFlowJobRequest {
	s.Conf = &v
	return s
}

func (s *SubmitFlowJobRequest) SetHostName(v string) *SubmitFlowJobRequest {
	s.HostName = &v
	return s
}

func (s *SubmitFlowJobRequest) SetJobId(v string) *SubmitFlowJobRequest {
	s.JobId = &v
	return s
}

func (s *SubmitFlowJobRequest) SetProjectId(v string) *SubmitFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitFlowJobRequest) SetRegionId(v string) *SubmitFlowJobRequest {
	s.RegionId = &v
	return s
}

type SubmitFlowJobResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFlowJobResponseBody) SetId(v string) *SubmitFlowJobResponseBody {
	s.Id = &v
	return s
}

func (s *SubmitFlowJobResponseBody) SetRequestId(v string) *SubmitFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitFlowJobResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitFlowJobResponse) SetHeaders(v map[string]*string) *SubmitFlowJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitFlowJobResponse) SetBody(v *SubmitFlowJobResponseBody) *SubmitFlowJobResponse {
	s.Body = v
	return s
}

type SuspendFlowRequest struct {
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SuspendFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendFlowRequest) GoString() string {
	return s.String()
}

func (s *SuspendFlowRequest) SetFlowInstanceId(v string) *SuspendFlowRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *SuspendFlowRequest) SetProjectId(v string) *SuspendFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *SuspendFlowRequest) SetRegionId(v string) *SuspendFlowRequest {
	s.RegionId = &v
	return s
}

type SuspendFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendFlowResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendFlowResponseBody) SetData(v bool) *SuspendFlowResponseBody {
	s.Data = &v
	return s
}

func (s *SuspendFlowResponseBody) SetRequestId(v string) *SuspendFlowResponseBody {
	s.RequestId = &v
	return s
}

type SuspendFlowResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuspendFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendFlowResponse) GoString() string {
	return s.String()
}

func (s *SuspendFlowResponse) SetHeaders(v map[string]*string) *SuspendFlowResponse {
	s.Headers = v
	return s
}

func (s *SuspendFlowResponse) SetBody(v *SuspendFlowResponseBody) *SuspendFlowResponse {
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
		"cn-qingdao":     tea.String("ddi.cn-qingdao.aliyuncs.com"),
		"cn-chengdu":     tea.String("ddi.cn-chengdu.aliyuncs.com"),
		"cn-zhangjiakou": tea.String("ddi.cn-zhangjiakou.aliyuncs.com"),
		"cn-huhehaote":   tea.String("ddi.cn-huhehaote.aliyuncs.com"),
		"cn-hongkong":    tea.String("ddi.cn-hongkong.aliyuncs.com"),
		"ap-southeast-2": tea.String("ddi.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3": tea.String("ddi.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5": tea.String("ddi.ap-southeast-5.aliyuncs.com"),
		"ap-northeast-1": tea.String("ddi.ap-northeast-1.aliyuncs.com"),
		"eu-west-1":      tea.String("ddi.eu-west-1.aliyuncs.com"),
		"us-east-1":      tea.String("ddi.us-east-1.aliyuncs.com"),
		"eu-central-1":   tea.String("ddi.eu-central-1.aliyuncs.com"),
		"me-east-1":      tea.String("ddi.me-east-1.aliyuncs.com"),
		"ap-south-1":     tea.String("ddi.ap-south-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CloneFlowWithOptions(request *CloneFlowRequest, runtime *util.RuntimeOptions) (_result *CloneFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneFlow(request *CloneFlowRequest) (_result *CloneFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneFlowResponse{}
	_body, _err := client.CloneFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloneFlowJobWithOptions(request *CloneFlowJobRequest, runtime *util.RuntimeOptions) (_result *CloneFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["Name"] = request.Name
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneFlowJob"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneFlowJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneFlowJob(request *CloneFlowJobRequest) (_result *CloneFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneFlowJobResponse{}
	_body, _err := client.CloneFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CommitFlowEntitySnapshotWithOptions(request *CommitFlowEntitySnapshotRequest, runtime *util.RuntimeOptions) (_result *CommitFlowEntitySnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EntityId"] = request.EntityId
	query["EntityType"] = request.EntityType
	query["Message"] = request.Message
	query["RegionId"] = request.RegionId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CommitFlowEntitySnapshot"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CommitFlowEntitySnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CommitFlowEntitySnapshot(request *CommitFlowEntitySnapshotRequest) (_result *CommitFlowEntitySnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CommitFlowEntitySnapshotResponse{}
	_body, _err := client.CommitFlowEntitySnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterV2WithOptions(request *CreateClusterV2Request, runtime *util.RuntimeOptions) (_result *CreateClusterV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AuthorizeContent"] = request.AuthorizeContent
	query["Auto"] = request.Auto
	query["AutoPayOrder"] = request.AutoPayOrder
	query["BootstrapAction"] = request.BootstrapAction
	query["ChargeType"] = request.ChargeType
	query["ClickHouseConf"] = request.ClickHouseConf
	query["ClientToken"] = request.ClientToken
	query["ClusterType"] = request.ClusterType
	query["Config"] = request.Config
	query["Configurations"] = request.Configurations
	query["DepositType"] = request.DepositType
	query["EmrVer"] = request.EmrVer
	query["EnableEas"] = request.EnableEas
	query["EnableHighAvailability"] = request.EnableHighAvailability
	query["EnableSsh"] = request.EnableSsh
	query["ExtraAttributes"] = request.ExtraAttributes
	query["HostComponentInfo"] = request.HostComponentInfo
	query["HostGroup"] = request.HostGroup
	query["InitCustomHiveMetaDB"] = request.InitCustomHiveMetaDB
	query["InstanceGeneration"] = request.InstanceGeneration
	query["IsOpenPublicIp"] = request.IsOpenPublicIp
	query["KeyPairName"] = request.KeyPairName
	query["LogPath"] = request.LogPath
	query["MachineType"] = request.MachineType
	query["MasterPwd"] = request.MasterPwd
	query["MetaStoreConf"] = request.MetaStoreConf
	query["MetaStoreType"] = request.MetaStoreType
	query["Name"] = request.Name
	query["NetType"] = request.NetType
	query["Period"] = request.Period
	query["PromotionInfo"] = request.PromotionInfo
	query["RegionId"] = request.RegionId
	query["RelatedClusterId"] = request.RelatedClusterId
	query["ResourceGroupId"] = request.ResourceGroupId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SecurityGroupId"] = request.SecurityGroupId
	query["SecurityGroupName"] = request.SecurityGroupName
	query["ServiceInfo"] = request.ServiceInfo
	query["Tag"] = request.Tag
	query["UseCustomHiveMetaDB"] = request.UseCustomHiveMetaDB
	query["UseLocalMetaDb"] = request.UseLocalMetaDb
	query["UserDefinedEmrEcsRole"] = request.UserDefinedEmrEcsRole
	query["UserInfo"] = request.UserInfo
	query["VSwitchId"] = request.VSwitchId
	query["VpcId"] = request.VpcId
	query["WhiteListType"] = request.WhiteListType
	query["ZoneId"] = request.ZoneId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClusterV2"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClusterV2(request *CreateClusterV2Request) (_result *CreateClusterV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterV2Response{}
	_body, _err := client.CreateClusterV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowWithOptions(request *CreateFlowRequest, runtime *util.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AlertConf"] = request.AlertConf
	query["AlertDingDingGroupBizId"] = request.AlertDingDingGroupBizId
	query["AlertUserGroupBizId"] = request.AlertUserGroupBizId
	query["Application"] = request.Application
	query["ClientToken"] = request.ClientToken
	query["ClusterId"] = request.ClusterId
	query["CreateCluster"] = request.CreateCluster
	query["CronExpression"] = request.CronExpression
	query["Description"] = request.Description
	query["EndSchedule"] = request.EndSchedule
	query["HostName"] = request.HostName
	query["Name"] = request.Name
	query["Namespace"] = request.Namespace
	query["ParentCategory"] = request.ParentCategory
	query["ParentFlowList"] = request.ParentFlowList
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["StartSchedule"] = request.StartSchedule
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlow(request *CreateFlowRequest) (_result *CreateFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowResponse{}
	_body, _err := client.CreateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowCategoryWithOptions(request *CreateFlowCategoryRequest, runtime *util.RuntimeOptions) (_result *CreateFlowCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["Name"] = request.Name
	query["ParentId"] = request.ParentId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowCategory"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowCategory(request *CreateFlowCategoryRequest) (_result *CreateFlowCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowCategoryResponse{}
	_body, _err := client.CreateFlowCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowEditLockWithOptions(request *CreateFlowEditLockRequest, runtime *util.RuntimeOptions) (_result *CreateFlowEditLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EntityId"] = request.EntityId
	query["Force"] = request.Force
	query["RegionId"] = request.RegionId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowEditLock"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowEditLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowEditLock(request *CreateFlowEditLockRequest) (_result *CreateFlowEditLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowEditLockResponse{}
	_body, _err := client.CreateFlowEditLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowJobWithOptions(request *CreateFlowJobRequest, runtime *util.RuntimeOptions) (_result *CreateFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Adhoc"] = request.Adhoc
	query["AlertConf"] = request.AlertConf
	query["ClientToken"] = request.ClientToken
	query["ClusterId"] = request.ClusterId
	query["CustomVariables"] = request.CustomVariables
	query["Description"] = request.Description
	query["EnvConf"] = request.EnvConf
	query["FailAct"] = request.FailAct
	query["Mode"] = request.Mode
	query["MonitorConf"] = request.MonitorConf
	query["Name"] = request.Name
	query["ParamConf"] = request.ParamConf
	query["Params"] = request.Params
	query["ParentCategory"] = request.ParentCategory
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["ResourceList"] = request.ResourceList
	query["RetryPolicy"] = request.RetryPolicy
	query["RunConf"] = request.RunConf
	query["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowJob"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowJob(request *CreateFlowJobRequest) (_result *CreateFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowJobResponse{}
	_body, _err := client.CreateFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowProjectWithOptions(request *CreateFlowProjectRequest, runtime *util.RuntimeOptions) (_result *CreateFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["Name"] = request.Name
	query["ProductType"] = request.ProductType
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowProject"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowProject(request *CreateFlowProjectRequest) (_result *CreateFlowProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowProjectResponse{}
	_body, _err := client.CreateFlowProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowProjectClusterSettingWithOptions(request *CreateFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *CreateFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["ClusterId"] = request.ClusterId
	query["DefaultQueue"] = request.DefaultQueue
	query["DefaultUser"] = request.DefaultUser
	query["HostList"] = request.HostList
	query["ProjectId"] = request.ProjectId
	query["QueueList"] = request.QueueList
	query["RegionId"] = request.RegionId
	query["UserList"] = request.UserList
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowProjectClusterSetting"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowProjectClusterSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowProjectClusterSetting(request *CreateFlowProjectClusterSettingRequest) (_result *CreateFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowProjectClusterSettingResponse{}
	_body, _err := client.CreateFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowProjectUserWithOptions(request *CreateFlowProjectUserRequest, runtime *util.RuntimeOptions) (_result *CreateFlowProjectUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["User"] = request.User
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlowProjectUser"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowProjectUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowProjectUser(request *CreateFlowProjectUserRequest) (_result *CreateFlowProjectUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowProjectUserResponse{}
	_body, _err := client.CreateFlowProjectUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowWithOptions(request *DeleteFlowRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlow(request *DeleteFlowRequest) (_result *DeleteFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DeleteFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowCategoryWithOptions(request *DeleteFlowCategoryRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowCategory"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowCategory(request *DeleteFlowCategoryRequest) (_result *DeleteFlowCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowCategoryResponse{}
	_body, _err := client.DeleteFlowCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowEditLockWithOptions(request *DeleteFlowEditLockRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowEditLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EntityId"] = request.EntityId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowEditLock"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowEditLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowEditLock(request *DeleteFlowEditLockRequest) (_result *DeleteFlowEditLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowEditLockResponse{}
	_body, _err := client.DeleteFlowEditLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowProjectWithOptions(request *DeleteFlowProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowProject"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowProject(request *DeleteFlowProjectRequest) (_result *DeleteFlowProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowProjectResponse{}
	_body, _err := client.DeleteFlowProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowProjectClusterSettingWithOptions(request *DeleteFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowProjectClusterSetting"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowProjectClusterSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowProjectClusterSetting(request *DeleteFlowProjectClusterSettingRequest) (_result *DeleteFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowProjectClusterSettingResponse{}
	_body, _err := client.DeleteFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowProjectUserWithOptions(request *DeleteFlowProjectUserRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowProjectUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["UserName"] = request.UserName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlowProjectUser"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowProjectUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowProjectUser(request *DeleteFlowProjectUserRequest) (_result *DeleteFlowProjectUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowProjectUserResponse{}
	_body, _err := client.DeleteFlowProjectUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterV2WithOptions(request *DescribeClusterV2Request, runtime *util.RuntimeOptions) (_result *DescribeClusterV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["RegionId"] = request.RegionId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterV2"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterV2(request *DescribeClusterV2Request) (_result *DescribeClusterV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterV2Response{}
	_body, _err := client.DescribeClusterV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowWithOptions(request *DescribeFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlow(request *DescribeFlowRequest) (_result *DescribeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowResponse{}
	_body, _err := client.DescribeFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowCategoryTreeWithOptions(request *DescribeFlowCategoryTreeRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowCategoryTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CategoryId"] = request.CategoryId
	query["Keyword"] = request.Keyword
	query["Mode"] = request.Mode
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowCategoryTree"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowCategoryTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowCategoryTree(request *DescribeFlowCategoryTreeRequest) (_result *DescribeFlowCategoryTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowCategoryTreeResponse{}
	_body, _err := client.DescribeFlowCategoryTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowInstanceWithOptions(request *DescribeFlowInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowInstance"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowInstance(request *DescribeFlowInstanceRequest) (_result *DescribeFlowInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowInstanceResponse{}
	_body, _err := client.DescribeFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowJobWithOptions(request *DescribeFlowJobRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowJob"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowJob(request *DescribeFlowJobRequest) (_result *DescribeFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowJobResponse{}
	_body, _err := client.DescribeFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceWithOptions(request *DescribeFlowNodeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowNodeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowNodeInstance"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowNodeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstance(request *DescribeFlowNodeInstanceRequest) (_result *DescribeFlowNodeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowNodeInstanceResponse{}
	_body, _err := client.DescribeFlowNodeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceContainerLogWithOptions(request *DescribeFlowNodeInstanceContainerLogRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowNodeInstanceContainerLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppId"] = request.AppId
	query["ContainerId"] = request.ContainerId
	query["Length"] = request.Length
	query["LogName"] = request.LogName
	query["NodeInstanceId"] = request.NodeInstanceId
	query["Offset"] = request.Offset
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowNodeInstanceContainerLog"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowNodeInstanceContainerLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceContainerLog(request *DescribeFlowNodeInstanceContainerLogRequest) (_result *DescribeFlowNodeInstanceContainerLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowNodeInstanceContainerLogResponse{}
	_body, _err := client.DescribeFlowNodeInstanceContainerLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceLauncherLogWithOptions(request *DescribeFlowNodeInstanceLauncherLogRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowNodeInstanceLauncherLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EndTime"] = request.EndTime
	query["Length"] = request.Length
	query["Lines"] = request.Lines
	query["NodeInstanceId"] = request.NodeInstanceId
	query["Offset"] = request.Offset
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["Reverse"] = request.Reverse
	query["Start"] = request.Start
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowNodeInstanceLauncherLog"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowNodeInstanceLauncherLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceLauncherLog(request *DescribeFlowNodeInstanceLauncherLogRequest) (_result *DescribeFlowNodeInstanceLauncherLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowNodeInstanceLauncherLogResponse{}
	_body, _err := client.DescribeFlowNodeInstanceLauncherLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowProjectWithOptions(request *DescribeFlowProjectRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowProject"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowProject(request *DescribeFlowProjectRequest) (_result *DescribeFlowProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowProjectResponse{}
	_body, _err := client.DescribeFlowProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowProjectClusterSettingWithOptions(request *DescribeFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowProjectClusterSetting"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowProjectClusterSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowProjectClusterSetting(request *DescribeFlowProjectClusterSettingRequest) (_result *DescribeFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowProjectClusterSettingResponse{}
	_body, _err := client.DescribeFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowSLAWithOptions(request *DescribeFlowSLARequest, runtime *util.RuntimeOptions) (_result *DescribeFlowSLAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["From"] = request.From
	query["Metrics"] = request.Metrics
	query["PeriodType"] = request.PeriodType
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["To"] = request.To
	query["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowSLA"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowSLAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowSLA(request *DescribeFlowSLARequest) (_result *DescribeFlowSLAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowSLAResponse{}
	_body, _err := client.DescribeFlowSLAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowVariableCollectionWithOptions(request *DescribeFlowVariableCollectionRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowVariableCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EntityId"] = request.EntityId
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowVariableCollection"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowVariableCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowVariableCollection(request *DescribeFlowVariableCollectionRequest) (_result *DescribeFlowVariableCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowVariableCollectionResponse{}
	_body, _err := client.DescribeFlowVariableCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlowAuditLogsWithOptions(request *GetFlowAuditLogsRequest, runtime *util.RuntimeOptions) (_result *GetFlowAuditLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CurrentSize"] = request.CurrentSize
	query["EntityGroupId"] = request.EntityGroupId
	query["EntityId"] = request.EntityId
	query["EntityType"] = request.EntityType
	query["Limit"] = request.Limit
	query["Operation"] = request.Operation
	query["OperatorId"] = request.OperatorId
	query["OrderField"] = request.OrderField
	query["OrderMode"] = request.OrderMode
	query["PageCount"] = request.PageCount
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFlowAuditLogs"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFlowAuditLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlowAuditLogs(request *GetFlowAuditLogsRequest) (_result *GetFlowAuditLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFlowAuditLogsResponse{}
	_body, _err := client.GetFlowAuditLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillFlowWithOptions(request *KillFlowRequest, runtime *util.RuntimeOptions) (_result *KillFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["FlowInstanceId"] = request.FlowInstanceId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("KillFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KillFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillFlow(request *KillFlowRequest) (_result *KillFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillFlowResponse{}
	_body, _err := client.KillFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillFlowJobWithOptions(request *KillFlowJobRequest, runtime *util.RuntimeOptions) (_result *KillFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["JobInstanceId"] = request.JobInstanceId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("KillFlowJob"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KillFlowJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillFlowJob(request *KillFlowJobRequest) (_result *KillFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillFlowJobResponse{}
	_body, _err := client.KillFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterTypeList"] = request.ClusterTypeList
	query["CreateType"] = request.CreateType
	query["DefaultStatus"] = request.DefaultStatus
	query["DepositType"] = request.DepositType
	query["ExpiredTagList"] = request.ExpiredTagList
	query["IsDesc"] = request.IsDesc
	query["MachineType"] = request.MachineType
	query["Name"] = request.Name
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["StatusList"] = request.StatusList
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowWithOptions(request *ListFlowRequest, runtime *util.RuntimeOptions) (_result *ListFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Id"] = request.Id
	query["JobId"] = request.JobId
	query["Name"] = request.Name
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["Periodic"] = request.Periodic
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlow(request *ListFlowRequest) (_result *ListFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowResponse{}
	_body, _err := client.ListFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowClusterWithOptions(request *ListFlowClusterRequest, runtime *util.RuntimeOptions) (_result *ListFlowClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowCluster"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowCluster(request *ListFlowClusterRequest) (_result *ListFlowClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowClusterResponse{}
	_body, _err := client.ListFlowClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowClusterAllWithOptions(request *ListFlowClusterAllRequest, runtime *util.RuntimeOptions) (_result *ListFlowClusterAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ProductType"] = request.ProductType
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowClusterAll"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowClusterAllResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowClusterAll(request *ListFlowClusterAllRequest) (_result *ListFlowClusterAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowClusterAllResponse{}
	_body, _err := client.ListFlowClusterAllWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowClusterAllHostsWithOptions(request *ListFlowClusterAllHostsRequest, runtime *util.RuntimeOptions) (_result *ListFlowClusterAllHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowClusterAllHosts"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowClusterAllHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowClusterAllHosts(request *ListFlowClusterAllHostsRequest) (_result *ListFlowClusterAllHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowClusterAllHostsResponse{}
	_body, _err := client.ListFlowClusterAllHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowClusterHostWithOptions(request *ListFlowClusterHostRequest, runtime *util.RuntimeOptions) (_result *ListFlowClusterHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowClusterHost"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowClusterHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowClusterHost(request *ListFlowClusterHostRequest) (_result *ListFlowClusterHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowClusterHostResponse{}
	_body, _err := client.ListFlowClusterHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowEntitySnapshotWithOptions(request *ListFlowEntitySnapshotRequest, runtime *util.RuntimeOptions) (_result *ListFlowEntitySnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CommitterId"] = request.CommitterId
	query["CurrentSize"] = request.CurrentSize
	query["EntityGroupId"] = request.EntityGroupId
	query["EntityId"] = request.EntityId
	query["EntityType"] = request.EntityType
	query["Limit"] = request.Limit
	query["OrderField"] = request.OrderField
	query["OrderMode"] = request.OrderMode
	query["PageCount"] = request.PageCount
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Revision"] = request.Revision
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowEntitySnapshot"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowEntitySnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowEntitySnapshot(request *ListFlowEntitySnapshotRequest) (_result *ListFlowEntitySnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowEntitySnapshotResponse{}
	_body, _err := client.ListFlowEntitySnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowInstanceWithOptions(request *ListFlowInstanceRequest, runtime *util.RuntimeOptions) (_result *ListFlowInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["FlowId"] = request.FlowId
	query["FlowName"] = request.FlowName
	query["Id"] = request.Id
	query["InstanceId"] = request.InstanceId
	query["OrderBy"] = request.OrderBy
	query["OrderType"] = request.OrderType
	query["Owner"] = request.Owner
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["StatusList"] = request.StatusList
	query["TimeRange"] = request.TimeRange
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowInstance"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowInstance(request *ListFlowInstanceRequest) (_result *ListFlowInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowInstanceResponse{}
	_body, _err := client.ListFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowJobHistoryWithOptions(request *ListFlowJobHistoryRequest, runtime *util.RuntimeOptions) (_result *ListFlowJobHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["InstanceId"] = request.InstanceId
	query["JobType"] = request.JobType
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["StatusList"] = request.StatusList
	query["TimeRange"] = request.TimeRange
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowJobHistory"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowJobHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowJobHistory(request *ListFlowJobHistoryRequest) (_result *ListFlowJobHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowJobHistoryResponse{}
	_body, _err := client.ListFlowJobHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowJobsWithOptions(request *ListFlowJobsRequest, runtime *util.RuntimeOptions) (_result *ListFlowJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Adhoc"] = request.Adhoc
	query["Id"] = request.Id
	query["Name"] = request.Name
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowJobs"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowJobs(request *ListFlowJobsRequest) (_result *ListFlowJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowJobsResponse{}
	_body, _err := client.ListFlowJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowNodeInstanceContainerStatusWithOptions(request *ListFlowNodeInstanceContainerStatusRequest, runtime *util.RuntimeOptions) (_result *ListFlowNodeInstanceContainerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["NodeInstanceId"] = request.NodeInstanceId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowNodeInstanceContainerStatus"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowNodeInstanceContainerStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowNodeInstanceContainerStatus(request *ListFlowNodeInstanceContainerStatusRequest) (_result *ListFlowNodeInstanceContainerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowNodeInstanceContainerStatusResponse{}
	_body, _err := client.ListFlowNodeInstanceContainerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowNodeSqlResultWithOptions(request *ListFlowNodeSqlResultRequest, runtime *util.RuntimeOptions) (_result *ListFlowNodeSqlResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Length"] = request.Length
	query["NodeInstanceId"] = request.NodeInstanceId
	query["Offset"] = request.Offset
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["SqlIndex"] = request.SqlIndex
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowNodeSqlResult"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowNodeSqlResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowNodeSqlResult(request *ListFlowNodeSqlResultRequest) (_result *ListFlowNodeSqlResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowNodeSqlResultResponse{}
	_body, _err := client.ListFlowNodeSqlResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowProjectClusterSettingWithOptions(request *ListFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *ListFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowProjectClusterSetting"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowProjectClusterSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowProjectClusterSetting(request *ListFlowProjectClusterSettingRequest) (_result *ListFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowProjectClusterSettingResponse{}
	_body, _err := client.ListFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowProjectUserWithOptions(request *ListFlowProjectUserRequest, runtime *util.RuntimeOptions) (_result *ListFlowProjectUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowProjectUser"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowProjectUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowProjectUser(request *ListFlowProjectUserRequest) (_result *ListFlowProjectUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowProjectUserResponse{}
	_body, _err := client.ListFlowProjectUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowProjectsWithOptions(request *ListFlowProjectsRequest, runtime *util.RuntimeOptions) (_result *ListFlowProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Name"] = request.Name
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProductType"] = request.ProductType
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlowProjects"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowProjects(request *ListFlowProjectsRequest) (_result *ListFlowProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowProjectsResponse{}
	_body, _err := client.ListFlowProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowsWithOptions(request *ListFlowsRequest, runtime *util.RuntimeOptions) (_result *ListFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Id"] = request.Id
	query["JobId"] = request.JobId
	query["Name"] = request.Name
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["Periodic"] = request.Periodic
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlows"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlows(request *ListFlowsRequest) (_result *ListFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowsResponse{}
	_body, _err := client.ListFlowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMainVersionsWithOptions(request *ListMainVersionsRequest, runtime *util.RuntimeOptions) (_result *ListMainVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMainVersions"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMainVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMainVersions(request *ListMainVersionsRequest) (_result *ListMainVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMainVersionsResponse{}
	_body, _err := client.ListMainVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowWithOptions(request *ModifyFlowRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AlertConf"] = request.AlertConf
	query["AlertDingDingGroupBizId"] = request.AlertDingDingGroupBizId
	query["AlertUserGroupBizId"] = request.AlertUserGroupBizId
	query["Application"] = request.Application
	query["ClusterId"] = request.ClusterId
	query["CreateCluster"] = request.CreateCluster
	query["CronExpr"] = request.CronExpr
	query["Description"] = request.Description
	query["EndSchedule"] = request.EndSchedule
	query["HostName"] = request.HostName
	query["Id"] = request.Id
	query["Name"] = request.Name
	query["ParentCategory"] = request.ParentCategory
	query["ParentFlowList"] = request.ParentFlowList
	query["Periodic"] = request.Periodic
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["StartSchedule"] = request.StartSchedule
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlow(request *ModifyFlowRequest) (_result *ModifyFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowResponse{}
	_body, _err := client.ModifyFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowCategoryWithOptions(request *ModifyFlowCategoryRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["Name"] = request.Name
	query["ParentId"] = request.ParentId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlowCategory"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowCategory(request *ModifyFlowCategoryRequest) (_result *ModifyFlowCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowCategoryResponse{}
	_body, _err := client.ModifyFlowCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowForWebWithOptions(request *ModifyFlowForWebRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowForWebResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AlertConf"] = request.AlertConf
	query["AlertDingDingGroupBizId"] = request.AlertDingDingGroupBizId
	query["AlertUserGroupBizId"] = request.AlertUserGroupBizId
	query["ClusterId"] = request.ClusterId
	query["CreateCluster"] = request.CreateCluster
	query["CronExpr"] = request.CronExpr
	query["Description"] = request.Description
	query["EndSchedule"] = request.EndSchedule
	query["Graph"] = request.Graph
	query["HostName"] = request.HostName
	query["Id"] = request.Id
	query["Name"] = request.Name
	query["Namespace"] = request.Namespace
	query["ParentCategory"] = request.ParentCategory
	query["ParentFlowList"] = request.ParentFlowList
	query["Periodic"] = request.Periodic
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["StartSchedule"] = request.StartSchedule
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlowForWeb"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowForWebResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowForWeb(request *ModifyFlowForWebRequest) (_result *ModifyFlowForWebResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowForWebResponse{}
	_body, _err := client.ModifyFlowForWebWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowJobWithOptions(request *ModifyFlowJobRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AlertConf"] = request.AlertConf
	query["ClusterId"] = request.ClusterId
	query["CustomVariables"] = request.CustomVariables
	query["Description"] = request.Description
	query["EnvConf"] = request.EnvConf
	query["FailAct"] = request.FailAct
	query["Id"] = request.Id
	query["KnoxPassword"] = request.KnoxPassword
	query["KnoxUser"] = request.KnoxUser
	query["Mode"] = request.Mode
	query["MonitorConf"] = request.MonitorConf
	query["Name"] = request.Name
	query["ParamConf"] = request.ParamConf
	query["Params"] = request.Params
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	query["ResourceList"] = request.ResourceList
	query["RetryPolicy"] = request.RetryPolicy
	query["RunConf"] = request.RunConf
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlowJob"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowJob(request *ModifyFlowJobRequest) (_result *ModifyFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowJobResponse{}
	_body, _err := client.ModifyFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowProjectWithOptions(request *ModifyFlowProjectRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Description"] = request.Description
	query["Name"] = request.Name
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlowProject"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowProject(request *ModifyFlowProjectRequest) (_result *ModifyFlowProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowProjectResponse{}
	_body, _err := client.ModifyFlowProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowProjectClusterSettingWithOptions(request *ModifyFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["DefaultQueue"] = request.DefaultQueue
	query["DefaultUser"] = request.DefaultUser
	query["HostList"] = request.HostList
	query["ProjectId"] = request.ProjectId
	query["QueueList"] = request.QueueList
	query["RegionId"] = request.RegionId
	query["UserList"] = request.UserList
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlowProjectClusterSetting"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowProjectClusterSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowProjectClusterSetting(request *ModifyFlowProjectClusterSettingRequest) (_result *ModifyFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowProjectClusterSettingResponse{}
	_body, _err := client.ModifyFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowVariableCollectionWithOptions(request *ModifyFlowVariableCollectionRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowVariableCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Data"] = request.Data
	query["RegionId"] = request.RegionId
	query["ResourceGroupId"] = request.ResourceGroupId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFlowVariableCollection"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFlowVariableCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowVariableCollection(request *ModifyFlowVariableCollectionRequest) (_result *ModifyFlowVariableCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowVariableCollectionResponse{}
	_body, _err := client.ModifyFlowVariableCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseClusterWithOptions(request *ReleaseClusterRequest, runtime *util.RuntimeOptions) (_result *ReleaseClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ForceRelease"] = request.ForceRelease
	query["Id"] = request.Id
	query["RegionId"] = request.RegionId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseCluster"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseCluster(request *ReleaseClusterRequest) (_result *ReleaseClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseClusterResponse{}
	_body, _err := client.ReleaseClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RerunFlowWithOptions(request *RerunFlowRequest, runtime *util.RuntimeOptions) (_result *RerunFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["FlowInstanceId"] = request.FlowInstanceId
	query["ProjectId"] = request.ProjectId
	query["ReRunFail"] = request.ReRunFail
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RerunFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RerunFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RerunFlow(request *RerunFlowRequest) (_result *RerunFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RerunFlowResponse{}
	_body, _err := client.RerunFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestoreFlowEntitySnapshotWithOptions(request *RestoreFlowEntitySnapshotRequest, runtime *util.RuntimeOptions) (_result *RestoreFlowEntitySnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EntityId"] = request.EntityId
	query["EntityType"] = request.EntityType
	query["OperatorId"] = request.OperatorId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Revision"] = request.Revision
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreFlowEntitySnapshot"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreFlowEntitySnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestoreFlowEntitySnapshot(request *RestoreFlowEntitySnapshotRequest) (_result *RestoreFlowEntitySnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreFlowEntitySnapshotResponse{}
	_body, _err := client.RestoreFlowEntitySnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeFlowWithOptions(request *ResumeFlowRequest, runtime *util.RuntimeOptions) (_result *ResumeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["FlowInstanceId"] = request.FlowInstanceId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeFlow(request *ResumeFlowRequest) (_result *ResumeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeFlowResponse{}
	_body, _err := client.ResumeFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartFlowWithOptions(request *StartFlowRequest, runtime *util.RuntimeOptions) (_result *StartFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["FlowInstanceId"] = request.FlowInstanceId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartFlow(request *StartFlowRequest) (_result *StartFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartFlowResponse{}
	_body, _err := client.StartFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitFlowWithOptions(request *SubmitFlowRequest, runtime *util.RuntimeOptions) (_result *SubmitFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Conf"] = request.Conf
	query["FlowId"] = request.FlowId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitFlow(request *SubmitFlowRequest) (_result *SubmitFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitFlowResponse{}
	_body, _err := client.SubmitFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitFlowJobWithOptions(request *SubmitFlowJobRequest, runtime *util.RuntimeOptions) (_result *SubmitFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Conf"] = request.Conf
	query["HostName"] = request.HostName
	query["JobId"] = request.JobId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitFlowJob"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitFlowJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitFlowJob(request *SubmitFlowJobRequest) (_result *SubmitFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitFlowJobResponse{}
	_body, _err := client.SubmitFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendFlowWithOptions(request *SuspendFlowRequest, runtime *util.RuntimeOptions) (_result *SuspendFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["FlowInstanceId"] = request.FlowInstanceId
	query["ProjectId"] = request.ProjectId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendFlow"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SuspendFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendFlow(request *SuspendFlowRequest) (_result *SuspendFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendFlowResponse{}
	_body, _err := client.SuspendFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
