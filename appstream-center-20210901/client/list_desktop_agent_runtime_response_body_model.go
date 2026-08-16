// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDesktopAgentRuntimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDesktopAgentRuntimeResponseBodyData) *ListDesktopAgentRuntimeResponseBody
	GetData() []*ListDesktopAgentRuntimeResponseBodyData
	SetPageNumber(v int32) *ListDesktopAgentRuntimeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDesktopAgentRuntimeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDesktopAgentRuntimeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDesktopAgentRuntimeResponseBody
	GetTotalCount() *int32
}

type ListDesktopAgentRuntimeResponseBody struct {
	// The list of returned result objects.
	Data []*ListDesktopAgentRuntimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDesktopAgentRuntimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDesktopAgentRuntimeResponseBody) GoString() string {
	return s.String()
}

func (s *ListDesktopAgentRuntimeResponseBody) GetData() []*ListDesktopAgentRuntimeResponseBodyData {
	return s.Data
}

func (s *ListDesktopAgentRuntimeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDesktopAgentRuntimeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDesktopAgentRuntimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDesktopAgentRuntimeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDesktopAgentRuntimeResponseBody) SetData(v []*ListDesktopAgentRuntimeResponseBodyData) *ListDesktopAgentRuntimeResponseBody {
	s.Data = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBody) SetPageNumber(v int32) *ListDesktopAgentRuntimeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBody) SetPageSize(v int32) *ListDesktopAgentRuntimeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBody) SetRequestId(v string) *ListDesktopAgentRuntimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBody) SetTotalCount(v int32) *ListDesktopAgentRuntimeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDesktopAgentRuntimeResponseBodyData struct {
	// The agent IM information.
	AgentImInfo *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo `json:"AgentImInfo,omitempty" xml:"AgentImInfo,omitempty" type:"Struct"`
	// The list of agent instance information.
	AgentInstanceInfoList []*ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList `json:"AgentInstanceInfoList,omitempty" xml:"AgentInstanceInfoList,omitempty" type:"Repeated"`
	// The list of authorized users.
	AuthUsers []*string `json:"AuthUsers,omitempty" xml:"AuthUsers,omitempty" type:"Repeated"`
	// Indicates whether a third-party channel has been configured for the agent runtime.
	//
	// example:
	//
	// true
	ChannelConfigure *bool `json:"ChannelConfigure,omitempty" xml:"ChannelConfigure,omitempty"`
	// The list of third-party channels configured for the agent runtime.
	ChannelConfiguredList []*string `json:"ChannelConfiguredList,omitempty" xml:"ChannelConfiguredList,omitempty" type:"Repeated"`
	// The agent runtime ID.
	//
	// example:
	//
	// ecd-xxxx
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The agent runtime name.
	//
	// example:
	//
	// Agent-001
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The Cloud Desktop status.
	//
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// Indicates whether there is an authorized user with authorization.
	//
	// example:
	//
	// true
	HasAuthUser *bool `json:"HasAuthUser,omitempty" xml:"HasAuthUser,omitempty"`
	// The management status list, parsed from the managementStatus composite value.
	ManagementStatuses []*string `json:"ManagementStatuses,omitempty" xml:"ManagementStatuses,omitempty" type:"Repeated"`
	// Indicates whether a model has been configured for the agent runtime.
	//
	// example:
	//
	// true
	ModelConfigure *bool `json:"ModelConfigure,omitempty" xml:"ModelConfigure,omitempty"`
	// The active model template ID. This parameter is returned only when modelConfigure is set to true.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// The active model template name. This parameter is returned only when modelConfigure is set to true.
	//
	// example:
	//
	// model-template-001
	ModelTemplateName *string `json:"ModelTemplateName,omitempty" xml:"ModelTemplateName,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The list of channel codes in QR code configuration.
	QrCodeConfiguringList []*string `json:"QrCodeConfiguringList,omitempty" xml:"QrCodeConfiguringList,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region location (the Chinese mainland or outside China).
	//
	// example:
	//
	// Mainland
	RegionLocation *string `json:"RegionLocation,omitempty" xml:"RegionLocation,omitempty"`
	// The resource group information.
	ResourceGroup *ListDesktopAgentRuntimeResponseBodyDataResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
	// The resource ID, which is the Cloud Desktop ID.
	//
	// example:
	//
	// ecd-xxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The risk information. This parameter is returned only when the request parameter IncludeRiskInfo is set to true. Otherwise, null is returned.
	RiskInfo *ListDesktopAgentRuntimeResponseBodyDataRiskInfo `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty" type:"Struct"`
}

func (s ListDesktopAgentRuntimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDesktopAgentRuntimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetAgentImInfo() *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo {
	return s.AgentImInfo
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetAgentInstanceInfoList() []*ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	return s.AgentInstanceInfoList
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetAuthUsers() []*string {
	return s.AuthUsers
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetChannelConfigure() *bool {
	return s.ChannelConfigure
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetChannelConfiguredList() []*string {
	return s.ChannelConfiguredList
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetDesktopName() *string {
	return s.DesktopName
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetHasAuthUser() *bool {
	return s.HasAuthUser
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetManagementStatuses() []*string {
	return s.ManagementStatuses
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetModelConfigure() *bool {
	return s.ModelConfigure
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetModelTemplateName() *string {
	return s.ModelTemplateName
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetQrCodeConfiguringList() []*string {
	return s.QrCodeConfiguringList
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetRegionLocation() *string {
	return s.RegionLocation
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetResourceGroup() *ListDesktopAgentRuntimeResponseBodyDataResourceGroup {
	return s.ResourceGroup
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetRiskInfo() *ListDesktopAgentRuntimeResponseBodyDataRiskInfo {
	return s.RiskInfo
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetAgentImInfo(v *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) *ListDesktopAgentRuntimeResponseBodyData {
	s.AgentImInfo = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetAgentInstanceInfoList(v []*ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) *ListDesktopAgentRuntimeResponseBodyData {
	s.AgentInstanceInfoList = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetAuthUsers(v []*string) *ListDesktopAgentRuntimeResponseBodyData {
	s.AuthUsers = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetChannelConfigure(v bool) *ListDesktopAgentRuntimeResponseBodyData {
	s.ChannelConfigure = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetChannelConfiguredList(v []*string) *ListDesktopAgentRuntimeResponseBodyData {
	s.ChannelConfiguredList = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetDesktopId(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.DesktopId = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetDesktopName(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.DesktopName = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetDesktopStatus(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.DesktopStatus = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetHasAuthUser(v bool) *ListDesktopAgentRuntimeResponseBodyData {
	s.HasAuthUser = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetManagementStatuses(v []*string) *ListDesktopAgentRuntimeResponseBodyData {
	s.ManagementStatuses = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetModelConfigure(v bool) *ListDesktopAgentRuntimeResponseBodyData {
	s.ModelConfigure = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetModelTemplateId(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.ModelTemplateId = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetModelTemplateName(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.ModelTemplateName = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetOsType(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.OsType = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetQrCodeConfiguringList(v []*string) *ListDesktopAgentRuntimeResponseBodyData {
	s.QrCodeConfiguringList = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetRegionId(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetRegionLocation(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.RegionLocation = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetResourceGroup(v *ListDesktopAgentRuntimeResponseBodyDataResourceGroup) *ListDesktopAgentRuntimeResponseBodyData {
	s.ResourceGroup = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetResourceId(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetRiskInfo(v *ListDesktopAgentRuntimeResponseBodyDataRiskInfo) *ListDesktopAgentRuntimeResponseBodyData {
	s.RiskInfo = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) Validate() error {
	if s.AgentImInfo != nil {
		if err := s.AgentImInfo.Validate(); err != nil {
			return err
		}
	}
	if s.AgentInstanceInfoList != nil {
		for _, item := range s.AgentInstanceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceGroup != nil {
		if err := s.ResourceGroup.Validate(); err != nil {
			return err
		}
	}
	if s.RiskInfo != nil {
		if err := s.RiskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDesktopAgentRuntimeResponseBodyDataAgentImInfo struct {
	// The agent IM online status (Online/Offline). Default value: Offline.
	AgentImOnlineStatus *string `json:"AgentImOnlineStatus,omitempty" xml:"AgentImOnlineStatus,omitempty"`
	// The agent IM status.
	//
	// example:
	//
	// Enabled
	AgentImStatus *string `json:"AgentImStatus,omitempty" xml:"AgentImStatus,omitempty"`
	// The CloudSpace status.
	//
	// example:
	//
	// Enabled
	CloudSpaceStatus *string `json:"CloudSpaceStatus,omitempty" xml:"CloudSpaceStatus,omitempty"`
}

func (s ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) GoString() string {
	return s.String()
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) GetAgentImOnlineStatus() *string {
	return s.AgentImOnlineStatus
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) GetAgentImStatus() *string {
	return s.AgentImStatus
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) GetCloudSpaceStatus() *string {
	return s.CloudSpaceStatus
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) SetAgentImOnlineStatus(v string) *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo {
	s.AgentImOnlineStatus = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) SetAgentImStatus(v string) *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo {
	s.AgentImStatus = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) SetCloudSpaceStatus(v string) *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo {
	s.CloudSpaceStatus = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) Validate() error {
	return dara.Validate(s)
}

type ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList struct {
	// The agent instance status.
	//
	// example:
	//
	// Running
	AgentInstanceStatus *string `json:"AgentInstanceStatus,omitempty" xml:"AgentInstanceStatus,omitempty"`
	// The agent instance version.
	//
	// example:
	//
	// 2026.3.13
	AgentInstanceVersion *string `json:"AgentInstanceVersion,omitempty" xml:"AgentInstanceVersion,omitempty"`
	// The agent platform (enum name, such as ENTERPRISE, JVS, or ENTERPRISE_JVS).
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider (enum name, such as OPEN_CLAW or HERMES_AGENT).
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// Indicates whether a third-party channel has been configured for the agent instance.
	//
	// example:
	//
	// true
	ChannelConfigure *bool `json:"ChannelConfigure,omitempty" xml:"ChannelConfigure,omitempty"`
	// The list of third-party channels configured for the agent instance.
	ChannelConfiguredList []*string `json:"ChannelConfiguredList,omitempty" xml:"ChannelConfiguredList,omitempty" type:"Repeated"`
	// The deployment source.
	//
	// example:
	//
	// Admin
	DeploymentSource *string `json:"DeploymentSource,omitempty" xml:"DeploymentSource,omitempty"`
	// Indicates whether a model has been configured for the agent instance.
	//
	// example:
	//
	// true
	ModelConfigure *bool `json:"ModelConfigure,omitempty" xml:"ModelConfigure,omitempty"`
	// The configured model group ID.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
}

func (s ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GoString() string {
	return s.String()
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GetAgentInstanceStatus() *string {
	return s.AgentInstanceStatus
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GetAgentInstanceVersion() *string {
	return s.AgentInstanceVersion
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GetChannelConfigure() *bool {
	return s.ChannelConfigure
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GetChannelConfiguredList() []*string {
	return s.ChannelConfiguredList
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GetDeploymentSource() *string {
	return s.DeploymentSource
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GetModelConfigure() *bool {
	return s.ModelConfigure
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) SetAgentInstanceStatus(v string) *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	s.AgentInstanceStatus = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) SetAgentInstanceVersion(v string) *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	s.AgentInstanceVersion = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) SetAgentPlatform(v string) *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	s.AgentPlatform = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) SetAgentProvider(v string) *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	s.AgentProvider = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) SetChannelConfigure(v bool) *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	s.ChannelConfigure = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) SetChannelConfiguredList(v []*string) *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	s.ChannelConfiguredList = v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) SetDeploymentSource(v string) *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	s.DeploymentSource = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) SetModelConfigure(v bool) *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	s.ModelConfigure = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) SetModelTemplateId(v string) *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList {
	s.ModelTemplateId = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList) Validate() error {
	return dara.Validate(s)
}

type ListDesktopAgentRuntimeResponseBodyDataResourceGroup struct {
	// The resource group ID.
	//
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource group name.
	//
	// example:
	//
	// resource-group-001
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ListDesktopAgentRuntimeResponseBodyDataResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s ListDesktopAgentRuntimeResponseBodyDataResourceGroup) GoString() string {
	return s.String()
}

func (s *ListDesktopAgentRuntimeResponseBodyDataResourceGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDesktopAgentRuntimeResponseBodyDataResourceGroup) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListDesktopAgentRuntimeResponseBodyDataResourceGroup) SetResourceGroupId(v string) *ListDesktopAgentRuntimeResponseBodyDataResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataResourceGroup) SetResourceGroupName(v string) *ListDesktopAgentRuntimeResponseBodyDataResourceGroup {
	s.ResourceGroupName = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataResourceGroup) Validate() error {
	return dara.Validate(s)
}

type ListDesktopAgentRuntimeResponseBodyDataRiskInfo struct {
	// Indicates whether the agent has been uninstalled.
	//
	// example:
	//
	// true
	AgentUninstalled *bool `json:"AgentUninstalled,omitempty" xml:"AgentUninstalled,omitempty"`
	// Indicates whether the third-party channel configuration has been modified (inconsistent with the administrator-distributed configuration).
	//
	// example:
	//
	// true
	ChannelModified *bool `json:"ChannelModified,omitempty" xml:"ChannelModified,omitempty"`
	// Indicates whether the model configuration has been modified (inconsistent with the administrator-distributed configuration).
	//
	// example:
	//
	// true
	ModelModified *bool `json:"ModelModified,omitempty" xml:"ModelModified,omitempty"`
}

func (s ListDesktopAgentRuntimeResponseBodyDataRiskInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDesktopAgentRuntimeResponseBodyDataRiskInfo) GoString() string {
	return s.String()
}

func (s *ListDesktopAgentRuntimeResponseBodyDataRiskInfo) GetAgentUninstalled() *bool {
	return s.AgentUninstalled
}

func (s *ListDesktopAgentRuntimeResponseBodyDataRiskInfo) GetChannelModified() *bool {
	return s.ChannelModified
}

func (s *ListDesktopAgentRuntimeResponseBodyDataRiskInfo) GetModelModified() *bool {
	return s.ModelModified
}

func (s *ListDesktopAgentRuntimeResponseBodyDataRiskInfo) SetAgentUninstalled(v bool) *ListDesktopAgentRuntimeResponseBodyDataRiskInfo {
	s.AgentUninstalled = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataRiskInfo) SetChannelModified(v bool) *ListDesktopAgentRuntimeResponseBodyDataRiskInfo {
	s.ChannelModified = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataRiskInfo) SetModelModified(v bool) *ListDesktopAgentRuntimeResponseBodyDataRiskInfo {
	s.ModelModified = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyDataRiskInfo) Validate() error {
	return dara.Validate(s)
}
