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

type CloneFlowJobRequest struct {
	// 克隆的目标作业ID。您可以调用ListFlowJob查看。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 克隆的目标作业名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 克隆的目标作业所属项目。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// 新产生的作业ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 请求ID。
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

type CreateFlowJobRequest struct {
	// 是否临时查询。
	Adhoc *bool `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	// 保留参数。
	AlertConf *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	// 保留参数。
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 集群ID。您可以调用ListClusters查看集群的ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 自定义变量。
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// 作业的描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 环境变量设置。
	EnvConf *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	// 失败策略，可能的取值：CONTINUE（提过本次作业），STOP（停止作业）
	FailAct *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	// 模型模式，取值如下：  YARN：将作业包装成一个Launcher提交至YARN中执行，LOCAL：作业直接在机器上以进程方式运行。
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 监控配置，仅SPARK_STREAMING类型作业支持监控配置。
	MonitorConf *string `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	// 作业的名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数设置。
	ParamConf *string `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	// 作业内容。
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 父目录ID。您可以调用DescribeFlowCategory查看。
	ParentCategory *string `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 保留参数。
	ResourceList []*CreateFlowJobRequestResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// 重试策略，保留参数。
	RetryPolicy *string `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	// 运行配置，取值如下：priority（优先级），userName（任务的Linux提交用户），memory（内存，单位为MB），cores（核数）
	RunConf *string `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	// 作业的类型，可能的取值有：SPARK，SPARK_STREAMING，ZEPPELIN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// 保留参数。
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// 保留参数。
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
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
	// 作业ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 请求ID。
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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 项目描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 项目名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 产品类型，固定值DATABRICKS_DATAINSIGHT
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// 地域ID
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组ID
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
	// 项目ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 请求ID
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

type DescribeFlowJobRequest struct {
	// 作业ID。您可以调用ListFlowJob查看作业ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// 是否临时查询。
	Adhoc *string `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	// 报警配置。
	AlertConf *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	// 作业所在目录ID。
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// 自定义变量。
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// 作业的描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 保留参数。
	EditLockDetail *string `json:"EditLockDetail,omitempty" xml:"EditLockDetail,omitempty"`
	// 环境变量设置。
	EnvConf *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	// 失败策略，可能的取值：CONTINUE（提过本次作业），STOP（停止作业）
	FailAct *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	// 创建时间。
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间。
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 作业ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Knox的用户密码，执行Zeppelin Notebook时必须提供。
	KnoxPassword *string `json:"KnoxPassword,omitempty" xml:"KnoxPassword,omitempty"`
	// Knox的用户名，执行Zeppelin Notebook时必须提供。
	KnoxUser *string `json:"KnoxUser,omitempty" xml:"KnoxUser,omitempty"`
	// 最后一次执行的实例ID。
	LastInstanceId *string `json:"LastInstanceId,omitempty" xml:"LastInstanceId,omitempty"`
	// 最大重试次数。
	MaxRetry *int32 `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	// 保留参数。
	MaxRunningTimeSec *int64 `json:"MaxRunningTimeSec,omitempty" xml:"MaxRunningTimeSec,omitempty"`
	// 模型模式，取值如下：  YARN：将作业包装成一个Launcher提交至YARN中执行，LOCAL：作业直接在机器上以进程方式运行。
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 监控配置，仅SPARK_STREAMING类型作业支持监控配置。
	MonitorConf *string `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	// 作业名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数设置。
	ParamConf *string `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	// 作业内容。
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 请求ID。
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceList *DescribeFlowJobResponseBodyResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Struct"`
	// 重试间隔 0~300（秒）。
	RetryInterval *int64 `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	// 重试策略，保留参数。
	RetryPolicy *string `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	// 运行配置，取值如下：priority（优先级），userName（任务的Linux提交用户），memory（内存，单位为MB），cores（核数）
	RunConf *string `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	// 作业的类型，可能的取值有：SPARK，SPARK_STREAMING，ZEPPELIN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// 保留参数。
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// 保留参数。
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
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

type KillFlowJobRequest struct {
	// 作业实例ID。您可以调用DescribeFlowJob查看作业实例ID。
	JobInstanceId *string `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// 返回执行结果，包含如下：true（执行成功），false（执行失败）
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// 请求ID。
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
	// 集群ID。您可以调用ListClusters查看集群的ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 工作流ID。您可以调用ListFlowInstance查看工作流ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 作业ID。您可以调用ListFlowJob查看。
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 工作流名称。您可以调用ListFlowInstance查看。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 页码。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页查询数量。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 是否调度。
	Periodic *bool `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 状态：  STOP_SCHEDULE（停止调度） UNDER_SCHEDULE（调度中）
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// 工作流列表
	Flow *ListFlowResponseBodyFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Struct"`
	// 页码。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页数量。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总数。
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
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

type ListFlowJobHistoryRequest struct {
	// 作业ID。您可以调用ListFlowJob查看作业ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 作业实例ID。您可以调用DescribeFlowJob查看作业实例ID。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 作业的类型，可能的取值有：SPARK，SPARK_STREAMING，ZEPPELIN
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 当前页码。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时每页行数。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 状态列表。取值如下：SUBMITTED, RUNNING, SUCCESS, FAILED, KILL_FAILED, KILL_SUCCESS
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// 查询的时间范围参数，参数列表：type: range，from: 开始时间（long型时间戳），to: 结束时间（long型时间戳）
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
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
	// 作业实例列表。
	NodeInstances *ListFlowJobHistoryResponseBodyNodeInstances `json:"NodeInstances,omitempty" xml:"NodeInstances,omitempty" type:"Struct"`
	// 当前页码。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时设置的每页行数。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 记录总数。
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 运行结束时间。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 环境变量设置。
	EnvConf *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	// 启动器的application的ID。
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// 外部信息。例如，运行作业的错误诊断信息。
	ExternalInfo *string `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty"`
	// 实例对应的Container的状态：SUBMITTED, RUNNING, SUCCESS, FAIL, KILL_FAIL, KILL_SUCCESS
	ExternalStatus *string `json:"ExternalStatus,omitempty" xml:"ExternalStatus,omitempty"`
	// 失败策略，可能的取值：CONTINUE（提过本次作业），STOP（停止作业）
	FailAct *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	// 创建时间。
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 创建时间。
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 保留参数。
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 作业实例ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 作业ID。
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 作业名称。
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// 作业内容。
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// 作业类型。
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 最大重试次数。
	MaxRetry *int32 `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	// 保留参数。
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// 参数设置。
	ParamConf *string `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 重试次数。
	Retries *int32 `json:"Retries,omitempty" xml:"Retries,omitempty"`
	// 重试间隔 0-300（秒）。
	RetryInterval *int64 `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	// 运行配置，取值如下：priority（优先级），userName（任务的Linux提交用户），memory（内存，单位为MB），cores（核数）
	RunConf *string `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	// 运行开始时间。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 实例的执行状态：PREP：准备启动，SUBMITTING：提交中，RUNNING：运行中DONE：已完成，OK：执行成功，FAILED：执行失败，KILLED：已终止，KILL_FAILED：终止失败，START_RETRY：开始重试
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 节点类型：JOB：作业，CLUSTER：集群，START：开始，END：结束
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 是否结束。
	Pending *bool `json:"pending,omitempty" xml:"pending,omitempty"`
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
	// 是否为临时查询。用于过滤作业。
	Adhoc     *bool   `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	ExactName *string `json:"ExactName,omitempty" xml:"ExactName,omitempty"`
	// 作业ID。您可以调用ListFlowJob查看作业ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 作业名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 当前页数。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页的作业数量。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 作业类型。用于过滤作业，支持的类型有：SPARK，SPARK_STREAMING，ZEPPELIN。
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *ListFlowJobsRequest) SetExactName(v string) *ListFlowJobsRequest {
	s.ExactName = &v
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
	JobList *ListFlowJobsResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Struct"`
	// 当前页数。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页的作业数量。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 作业数量。
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// 是否临时查询。
	Adhoc *string `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	// 报警配置。
	AlertConf *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	// 作业所在目录ID。
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// 自定义变量。
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// 作业的描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 环境变量设置。
	EnvConf *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	// 失败策略，可能的取值：CONTINUE（提过本次作业），STOP（停止作业）
	FailAct *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	// 创建时间。
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间。
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 作业ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 最后一次执行的实例ID。
	LastInstanceDetail *string `json:"LastInstanceDetail,omitempty" xml:"LastInstanceDetail,omitempty"`
	// 最大重试次数。
	MaxRetry *int32 `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	// 模型模式，取值如下：  YARN：将作业包装成一个Launcher提交至YARN中执行，LOCAL：作业直接在机器上以进程方式运行。
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 监控配置，仅SPARK_STREAMING类型作业支持监控配置。
	MonitorConf *string `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	// 作业名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数设置。
	ParamConf *string `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	// 作业内容。
	Params       *string                                         `json:"Params,omitempty" xml:"Params,omitempty"`
	ResourceList *ListFlowJobsResponseBodyJobListJobResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Struct"`
	// 重试间隔 0~300（秒）。
	RetryInterval *int64 `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	// 运行配置，取值如下：priority（优先级），userName（任务的Linux提交用户），memory（内存，单位为MB），cores（核数）
	RunConf *string `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	// 作业的类型，可能的取值有：SPARK，SPARK_STREAMING，ZEPPELIN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// 保留参数。
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// 保留参数。
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
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
	// 项目名称，用于过滤项目
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 页码，用于分页
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页数量
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 产品类型。固定值DATABIRCKS_DATAINSIGHT
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// 项目ID。您可以调用ListFlowProjects查看项目的ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组ID
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
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页数量
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 项目列表
	Projects *ListFlowProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Struct"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总数
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// 项目描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 创建时间戳
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 修改时间戳
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 项目ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 项目名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 主账号ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

type ModifyFlowJobRequest struct {
	// 保留参数。
	AlertConf *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	// 集群ID。您可以调用ListClusters查看集群的ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 自定义变量。
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// 修改后的作业描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 环境变量设置。
	EnvConf *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	// 失败策略，可能的取值：CONTINUE（提过本次作业），STOP（停止作业）
	FailAct *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	// 需要修改的作业的ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Knox的用户密码，执行Zeppelin Notebook时必须提供。
	KnoxPassword *string `json:"KnoxPassword,omitempty" xml:"KnoxPassword,omitempty"`
	// Knox的用户名，执行Zeppelin Notebook时必须提供。
	KnoxUser *string `json:"KnoxUser,omitempty" xml:"KnoxUser,omitempty"`
	// 模型模式，取值如下：  YARN：将作业包装成一个Launcher提交至YARN中执行，LOCAL：作业直接在机器上以进程方式运行。
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 监控配置，仅SPARK_STREAMING类型作业支持监控配置。
	MonitorConf *string `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	// 修改后的作业名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参数设置。
	ParamConf *string `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	// 作业内容。如果是spark作业，该参数的内容会作为spark-submit的参数。
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 保留参数。
	ResourceList []*ModifyFlowJobRequestResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// 重试策略，保留参数。
	RetryPolicy *string `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	// 运行配置，取值如下：priority（优先级），userName（任务的Linux提交用户），memory（内存，单位为MB），cores（核数）
	RunConf *string `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
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
	// 保留参数。
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// 保留参数。
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
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
	// API调用结果：true（修改成功），false（修改失败）
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// 请求ID。
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
	// 工作流实例ID。您可以调用ListFlowInstance查看工作流实例ID。
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 是否只重试失败节点。
	ReRunFail *bool `json:"ReRunFail,omitempty" xml:"ReRunFail,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// 返回执行结果，包含如下：true: 重试工作流成功，false: 重试工作流失败。
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// 请求ID。
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

type ResumeFlowRequest struct {
	// 工作流实例ID。您可以调用ListFlowInstance查看工作流ID。
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 区域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// 返回执行结果。
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// 请求ID。
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

type SubmitFlowRequest struct {
	// 配置信息{"key":"value"}格式。  本示例中cyctime表示实际调度运行的时间（长整型时间戳）。
	Conf *string `json:"Conf,omitempty" xml:"Conf,omitempty"`
	// 工作流ID。您可以调用ListFlowInstance查看工作流ID。
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// 过期参数。
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 工作流实例ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 过期参数。
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// 集群ID。您可以调用ListClusters查看集群的ID。
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 配置参数信息：{"key1":"value1"}。key为params的参数值会覆盖实际作业中运行的内容。
	Conf *string `json:"Conf,omitempty" xml:"Conf,omitempty"`
	// 保留参数。
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 作业ID。您可以调用ListFlowJob查看作业ID。
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 项目ID。您可以调用ListFlowProject查看项目的ID。
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 地域ID。您可以调用DescribeRegions查看最新的阿里云地域列表。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// 运行的作业实例ID。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 请求ID。
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

func (client *Client) CloneFlowJobWithOptions(request *CloneFlowJobRequest, runtime *util.RuntimeOptions) (_result *CloneFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

func (client *Client) CreateClusterV2WithOptions(request *CreateClusterV2Request, runtime *util.RuntimeOptions) (_result *CreateClusterV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizeContent)) {
		query["AuthorizeContent"] = request.AuthorizeContent
	}

	if !tea.BoolValue(util.IsUnset(request.Auto)) {
		query["Auto"] = request.Auto
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPayOrder)) {
		query["AutoPayOrder"] = request.AutoPayOrder
	}

	if !tea.BoolValue(util.IsUnset(request.BootstrapAction)) {
		query["BootstrapAction"] = request.BootstrapAction
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClickHouseConf)) {
		query["ClickHouseConf"] = request.ClickHouseConf
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Configurations)) {
		query["Configurations"] = request.Configurations
	}

	if !tea.BoolValue(util.IsUnset(request.DepositType)) {
		query["DepositType"] = request.DepositType
	}

	if !tea.BoolValue(util.IsUnset(request.EmrVer)) {
		query["EmrVer"] = request.EmrVer
	}

	if !tea.BoolValue(util.IsUnset(request.EnableEas)) {
		query["EnableEas"] = request.EnableEas
	}

	if !tea.BoolValue(util.IsUnset(request.EnableHighAvailability)) {
		query["EnableHighAvailability"] = request.EnableHighAvailability
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSsh)) {
		query["EnableSsh"] = request.EnableSsh
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraAttributes)) {
		query["ExtraAttributes"] = request.ExtraAttributes
	}

	if !tea.BoolValue(util.IsUnset(request.HostComponentInfo)) {
		query["HostComponentInfo"] = request.HostComponentInfo
	}

	if !tea.BoolValue(util.IsUnset(request.HostGroup)) {
		query["HostGroup"] = request.HostGroup
	}

	if !tea.BoolValue(util.IsUnset(request.InitCustomHiveMetaDB)) {
		query["InitCustomHiveMetaDB"] = request.InitCustomHiveMetaDB
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceGeneration)) {
		query["InstanceGeneration"] = request.InstanceGeneration
	}

	if !tea.BoolValue(util.IsUnset(request.IsOpenPublicIp)) {
		query["IsOpenPublicIp"] = request.IsOpenPublicIp
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.LogPath)) {
		query["LogPath"] = request.LogPath
	}

	if !tea.BoolValue(util.IsUnset(request.MachineType)) {
		query["MachineType"] = request.MachineType
	}

	if !tea.BoolValue(util.IsUnset(request.MasterPwd)) {
		query["MasterPwd"] = request.MasterPwd
	}

	if !tea.BoolValue(util.IsUnset(request.MetaStoreConf)) {
		query["MetaStoreConf"] = request.MetaStoreConf
	}

	if !tea.BoolValue(util.IsUnset(request.MetaStoreType)) {
		query["MetaStoreType"] = request.MetaStoreType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NetType)) {
		query["NetType"] = request.NetType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionInfo)) {
		query["PromotionInfo"] = request.PromotionInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedClusterId)) {
		query["RelatedClusterId"] = request.RelatedClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupName)) {
		query["SecurityGroupName"] = request.SecurityGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInfo)) {
		query["ServiceInfo"] = request.ServiceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UseCustomHiveMetaDB)) {
		query["UseCustomHiveMetaDB"] = request.UseCustomHiveMetaDB
	}

	if !tea.BoolValue(util.IsUnset(request.UseLocalMetaDb)) {
		query["UseLocalMetaDb"] = request.UseLocalMetaDb
	}

	if !tea.BoolValue(util.IsUnset(request.UserDefinedEmrEcsRole)) {
		query["UserDefinedEmrEcsRole"] = request.UserDefinedEmrEcsRole
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfo)) {
		query["UserInfo"] = request.UserInfo
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteListType)) {
		query["WhiteListType"] = request.WhiteListType
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

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
	if !tea.BoolValue(util.IsUnset(request.AlertConf)) {
		query["AlertConf"] = request.AlertConf
	}

	if !tea.BoolValue(util.IsUnset(request.AlertDingDingGroupBizId)) {
		query["AlertDingDingGroupBizId"] = request.AlertDingDingGroupBizId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertUserGroupBizId)) {
		query["AlertUserGroupBizId"] = request.AlertUserGroupBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Application)) {
		query["Application"] = request.Application
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateCluster)) {
		query["CreateCluster"] = request.CreateCluster
	}

	if !tea.BoolValue(util.IsUnset(request.CronExpression)) {
		query["CronExpression"] = request.CronExpression
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndSchedule)) {
		query["EndSchedule"] = request.EndSchedule
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCategory)) {
		query["ParentCategory"] = request.ParentCategory
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFlowList)) {
		query["ParentFlowList"] = request.ParentFlowList
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartSchedule)) {
		query["StartSchedule"] = request.StartSchedule
	}

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
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

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

func (client *Client) CreateFlowJobWithOptions(request *CreateFlowJobRequest, runtime *util.RuntimeOptions) (_result *CreateFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Adhoc)) {
		query["Adhoc"] = request.Adhoc
	}

	if !tea.BoolValue(util.IsUnset(request.AlertConf)) {
		query["AlertConf"] = request.AlertConf
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomVariables)) {
		query["CustomVariables"] = request.CustomVariables
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvConf)) {
		query["EnvConf"] = request.EnvConf
	}

	if !tea.BoolValue(util.IsUnset(request.FailAct)) {
		query["FailAct"] = request.FailAct
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorConf)) {
		query["MonitorConf"] = request.MonitorConf
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParamConf)) {
		query["ParamConf"] = request.ParamConf
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCategory)) {
		query["ParentCategory"] = request.ParentCategory
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceList)) {
		query["ResourceList"] = request.ResourceList
	}

	if !tea.BoolValue(util.IsUnset(request.RetryPolicy)) {
		query["RetryPolicy"] = request.RetryPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.RunConf)) {
		query["RunConf"] = request.RunConf
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

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
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

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

func (client *Client) CreateFlowProjectUserWithOptions(request *CreateFlowProjectUserRequest, runtime *util.RuntimeOptions) (_result *CreateFlowProjectUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

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
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

func (client *Client) DeleteFlowProjectWithOptions(request *DeleteFlowProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

func (client *Client) DeleteFlowProjectUserWithOptions(request *DeleteFlowProjectUserRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowProjectUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

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
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

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
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

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

func (client *Client) DescribeFlowJobWithOptions(request *DescribeFlowJobRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

func (client *Client) DescribeFlowProjectWithOptions(request *DescribeFlowProjectRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

func (client *Client) KillFlowJobWithOptions(request *KillFlowJobRequest, runtime *util.RuntimeOptions) (_result *KillFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobInstanceId)) {
		query["JobInstanceId"] = request.JobInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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
	if !tea.BoolValue(util.IsUnset(request.ClusterTypeList)) {
		query["ClusterTypeList"] = request.ClusterTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.CreateType)) {
		query["CreateType"] = request.CreateType
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultStatus)) {
		query["DefaultStatus"] = request.DefaultStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DepositType)) {
		query["DepositType"] = request.DepositType
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTagList)) {
		query["ExpiredTagList"] = request.ExpiredTagList
	}

	if !tea.BoolValue(util.IsUnset(request.IsDesc)) {
		query["IsDesc"] = request.IsDesc
	}

	if !tea.BoolValue(util.IsUnset(request.MachineType)) {
		query["MachineType"] = request.MachineType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		query["StatusList"] = request.StatusList
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

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
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Periodic)) {
		query["Periodic"] = request.Periodic
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

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

func (client *Client) ListFlowJobHistoryWithOptions(request *ListFlowJobHistoryRequest, runtime *util.RuntimeOptions) (_result *ListFlowJobHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		query["StatusList"] = request.StatusList
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRange)) {
		query["TimeRange"] = request.TimeRange
	}

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
	if !tea.BoolValue(util.IsUnset(request.Adhoc)) {
		query["Adhoc"] = request.Adhoc
	}

	if !tea.BoolValue(util.IsUnset(request.ExactName)) {
		query["ExactName"] = request.ExactName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

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

func (client *Client) ListFlowProjectUserWithOptions(request *ListFlowProjectUserRequest, runtime *util.RuntimeOptions) (_result *ListFlowProjectUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

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

func (client *Client) ModifyFlowJobWithOptions(request *ModifyFlowJobRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertConf)) {
		query["AlertConf"] = request.AlertConf
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomVariables)) {
		query["CustomVariables"] = request.CustomVariables
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvConf)) {
		query["EnvConf"] = request.EnvConf
	}

	if !tea.BoolValue(util.IsUnset(request.FailAct)) {
		query["FailAct"] = request.FailAct
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.KnoxPassword)) {
		query["KnoxPassword"] = request.KnoxPassword
	}

	if !tea.BoolValue(util.IsUnset(request.KnoxUser)) {
		query["KnoxUser"] = request.KnoxUser
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorConf)) {
		query["MonitorConf"] = request.MonitorConf
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParamConf)) {
		query["ParamConf"] = request.ParamConf
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceList)) {
		query["ResourceList"] = request.ResourceList
	}

	if !tea.BoolValue(util.IsUnset(request.RetryPolicy)) {
		query["RetryPolicy"] = request.RetryPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.RunConf)) {
		query["RunConf"] = request.RunConf
	}

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
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

func (client *Client) ReleaseClusterWithOptions(request *ReleaseClusterRequest, runtime *util.RuntimeOptions) (_result *ReleaseClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceRelease)) {
		query["ForceRelease"] = request.ForceRelease
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

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
	if !tea.BoolValue(util.IsUnset(request.FlowInstanceId)) {
		query["FlowInstanceId"] = request.FlowInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ReRunFail)) {
		query["ReRunFail"] = request.ReRunFail
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

func (client *Client) ResumeFlowWithOptions(request *ResumeFlowRequest, runtime *util.RuntimeOptions) (_result *ResumeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlowInstanceId)) {
		query["FlowInstanceId"] = request.FlowInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

func (client *Client) SubmitFlowWithOptions(request *SubmitFlowRequest, runtime *util.RuntimeOptions) (_result *SubmitFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Conf)) {
		query["Conf"] = request.Conf
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		query["FlowId"] = request.FlowId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Conf)) {
		query["Conf"] = request.Conf
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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
