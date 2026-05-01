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
	Data []*ListDesktopAgentRuntimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AgentImInfo           *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo             `json:"AgentImInfo,omitempty" xml:"AgentImInfo,omitempty" type:"Struct"`
	AgentInstanceInfoList []*ListDesktopAgentRuntimeResponseBodyDataAgentInstanceInfoList `json:"AgentInstanceInfoList,omitempty" xml:"AgentInstanceInfoList,omitempty" type:"Repeated"`
	AuthUsers             []*string                                                       `json:"AuthUsers,omitempty" xml:"AuthUsers,omitempty" type:"Repeated"`
	// example:
	//
	// true
	ChannelConfigure      *bool     `json:"ChannelConfigure,omitempty" xml:"ChannelConfigure,omitempty"`
	ChannelConfiguredList []*string `json:"ChannelConfiguredList,omitempty" xml:"ChannelConfiguredList,omitempty" type:"Repeated"`
	// example:
	//
	// ecd-xxxx
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// Agent-001
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// example:
	//
	// true
	HasAuthUser *bool `json:"HasAuthUser,omitempty" xml:"HasAuthUser,omitempty"`
	// example:
	//
	// true
	ModelConfigure *bool `json:"ModelConfigure,omitempty" xml:"ModelConfigure,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId      *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroup *ListDesktopAgentRuntimeResponseBodyDataResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
	// example:
	//
	// ecd-xxxx
	ResourceId *string                                          `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	RiskInfo   *ListDesktopAgentRuntimeResponseBodyDataRiskInfo `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty" type:"Struct"`
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

func (s *ListDesktopAgentRuntimeResponseBodyData) GetModelConfigure() *bool {
	return s.ModelConfigure
}

func (s *ListDesktopAgentRuntimeResponseBodyData) GetRegionId() *string {
	return s.RegionId
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

func (s *ListDesktopAgentRuntimeResponseBodyData) SetModelConfigure(v bool) *ListDesktopAgentRuntimeResponseBodyData {
	s.ModelConfigure = &v
	return s
}

func (s *ListDesktopAgentRuntimeResponseBodyData) SetRegionId(v string) *ListDesktopAgentRuntimeResponseBodyData {
	s.RegionId = &v
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
	// example:
	//
	// Enabled
	AgentImStatus *string `json:"AgentImStatus,omitempty" xml:"AgentImStatus,omitempty"`
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

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) GetAgentImStatus() *string {
	return s.AgentImStatus
}

func (s *ListDesktopAgentRuntimeResponseBodyDataAgentImInfo) GetCloudSpaceStatus() *string {
	return s.CloudSpaceStatus
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
	// example:
	//
	// Running
	AgentInstanceStatus *string `json:"AgentInstanceStatus,omitempty" xml:"AgentInstanceStatus,omitempty"`
	// example:
	//
	// 2026.3.13
	AgentInstanceVersion *string `json:"AgentInstanceVersion,omitempty" xml:"AgentInstanceVersion,omitempty"`
	// example:
	//
	// true
	ChannelConfigure      *bool     `json:"ChannelConfigure,omitempty" xml:"ChannelConfigure,omitempty"`
	ChannelConfiguredList []*string `json:"ChannelConfiguredList,omitempty" xml:"ChannelConfiguredList,omitempty" type:"Repeated"`
	// example:
	//
	// Admin
	DeploymentSource *string `json:"DeploymentSource,omitempty" xml:"DeploymentSource,omitempty"`
	// example:
	//
	// true
	ModelConfigure *bool `json:"ModelConfigure,omitempty" xml:"ModelConfigure,omitempty"`
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
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
	// example:
	//
	// true
	AgentUninstalled *bool `json:"AgentUninstalled,omitempty" xml:"AgentUninstalled,omitempty"`
	// example:
	//
	// true
	ChannelModified *bool `json:"ChannelModified,omitempty" xml:"ChannelModified,omitempty"`
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
