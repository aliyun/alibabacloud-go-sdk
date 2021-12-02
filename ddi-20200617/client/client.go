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
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClusterV2"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterV2"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseCluster"),
		Version:     tea.String("2020-06-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
