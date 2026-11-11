// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveRightsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIp(v string) *ActiveRightsRequest
	GetClientIp() *string
	SetCustomImageId(v string) *ActiveRightsRequest
	GetCustomImageId() *string
	SetDataDiskSize(v int32) *ActiveRightsRequest
	GetDataDiskSize() *int32
	SetDesktopType(v string) *ActiveRightsRequest
	GetDesktopType() *string
	SetEndTime(v string) *ActiveRightsRequest
	GetEndTime() *string
	SetHibernate(v int32) *ActiveRightsRequest
	GetHibernate() *int32
	SetLanguage(v string) *ActiveRightsRequest
	GetLanguage() *string
	SetOfficeSiteId(v string) *ActiveRightsRequest
	GetOfficeSiteId() *string
	SetPolicyGroupId(v string) *ActiveRightsRequest
	GetPolicyGroupId() *string
	SetResourceGroupId(v string) *ActiveRightsRequest
	GetResourceGroupId() *string
	SetRightsAssignedUser(v []*string) *ActiveRightsRequest
	GetRightsAssignedUser() []*string
	SetRightsGpuSpec(v string) *ActiveRightsRequest
	GetRightsGpuSpec() *string
	SetRightsIds(v []*string) *ActiveRightsRequest
	GetRightsIds() []*string
	SetRightsResourceArea(v string) *ActiveRightsRequest
	GetRightsResourceArea() *string
	SetRightsResourceImageOs(v string) *ActiveRightsRequest
	GetRightsResourceImageOs() *string
	SetRightsResourceRegion(v string) *ActiveRightsRequest
	GetRightsResourceRegion() *string
	SetRightsResourceType(v string) *ActiveRightsRequest
	GetRightsResourceType() *string
	SetRightsSpec(v string) *ActiveRightsRequest
	GetRightsSpec() *string
	SetRootDiskSize(v int32) *ActiveRightsRequest
	GetRootDiskSize() *int32
	SetStartTime(v string) *ActiveRightsRequest
	GetStartTime() *string
	SetVSwitchId(v string) *ActiveRightsRequest
	GetVSwitchId() *string
	SetWorkspaceIds(v []*string) *ActiveRightsRequest
	GetWorkspaceIds() []*string
	SetWyidList(v []*string) *ActiveRightsRequest
	GetWyidList() []*string
}

type ActiveRightsRequest struct {
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.1
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The ID of the image that is available for the specification in the target region.
	//
	// example:
	//
	// m-example
	CustomImageId *string `json:"CustomImageId,omitempty" xml:"CustomImageId,omitempty"`
	// The size of the data cloud disk, in GiB. Pass this parameter only when the commodity supports this activation parameter.
	//
	// example:
	//
	// 80
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The desktop type. Specify the value based on the types supported by the selected commodity. Omit this parameter if regular cloud desktops do not have this configuration.
	//
	// example:
	//
	// eds.business_edition_office.4c8g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The end time of the non-hibernation period, in the HH:mm:ss format. This parameter is used together with StartTime.
	//
	// example:
	//
	// 18:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The hibernation-related options. The compute plan and existing activation rules corresponding to the right are used.
	//
	// example:
	//
	// 0
	Hibernate *int32 `json:"Hibernate,omitempty" xml:"Hibernate,omitempty"`
	// The display language of the cloud desktop. Valid values: en-US, ja-JP, zh-CN, and zh-HK.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the office network in the target region. This parameter is required when you activate a cloud desktop.
	//
	// example:
	//
	// cn-hangzhou+dir-example
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The ID of the policy group. This parameter is optional. If you do not specify this parameter, the default policy of the workspace is used.
	//
	// example:
	//
	// pg-example
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The ID of the resource group. This parameter is optional. The ID is specified by the caller and passed through to the activation request.
	//
	// example:
	//
	// rg-example
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of IDs of the assigned end users. The IDs correspond to the rights in this batch by position. For convenience accounts, pass the EndUserId. The number of IDs must be the same as the number of rights.
	//
	// example:
	//
	// ["user-example"]
	RightsAssignedUser []*string `json:"RightsAssignedUser,omitempty" xml:"RightsAssignedUser,omitempty" type:"Repeated"`
	// The GPU specification. The value is subject to the return value of the rights details.
	//
	// example:
	//
	// eds.graphics_flagship_pro.8c16g.2g
	RightsGpuSpec *string `json:"RightsGpuSpec,omitempty" xml:"RightsGpuSpec,omitempty"`
	// The list of IDs of the rights to be activated. Select the rights accessible to the current identity from ListRights. The parameters in the same batch apply to these rights.
	//
	// example:
	//
	// ["New_beds-example_0001"]
	RightsIds []*string `json:"RightsIds,omitempty" xml:"RightsIds,omitempty" type:"Repeated"`
	// The area to which the right belongs. This parameter cannot replace the resource region for a new activation request.
	//
	// example:
	//
	// cn-east
	RightsResourceArea *string `json:"RightsResourceArea,omitempty" xml:"RightsResourceArea,omitempty"`
	// The operating system type of the image used for rights activation.
	//
	// example:
	//
	// Windows
	RightsResourceImageOs *string `json:"RightsResourceImageOs,omitempty" xml:"RightsResourceImageOs,omitempty"`
	// The ID of the region where the resource resides. This parameter is configured separately from the routing region of the POP request.
	//
	// example:
	//
	// cn-hangzhou
	RightsResourceRegion *string `json:"RightsResourceRegion,omitempty" xml:"RightsResourceRegion,omitempty"`
	// The resource type of the right. For regular cloud desktops, use Desktop. For other commodities, pass the actual type of the right.
	//
	// example:
	//
	// Desktop
	RightsResourceType *string `json:"RightsResourceType,omitempty" xml:"RightsResourceType,omitempty"`
	// The specification code of the right. The value is subject to the rights query result.
	//
	// example:
	//
	// eds.business_edition_office.4c8g
	RightsSpec *string `json:"RightsSpec,omitempty" xml:"RightsSpec,omitempty"`
	// The size of the system cloud disk, in GiB. Pass this parameter only when the commodity supports this activation parameter.
	//
	// example:
	//
	// 80
	RootDiskSize *int32 `json:"RootDiskSize,omitempty" xml:"RootDiskSize,omitempty"`
	// The start time of the non-hibernation period, in the HH:mm:ss format. This parameter applies only to the corresponding compute plan.
	//
	// example:
	//
	// 09:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the vSwitch in the target office network. Pass the value based on the network type.
	//
	// example:
	//
	// vsw-example
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The list of AD workspace IDs used for BusinessJvsCopilotAgent activation.
	//
	// example:
	//
	// ["ws-example"]
	WorkspaceIds []*string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty" type:"Repeated"`
	// The list of WAIDs used for BusinessJvsCopilotAgent activation. Do not pass this parameter for regular cloud desktop activation.
	//
	// example:
	//
	// ["waid-example"]
	WyidList []*string `json:"WyidList,omitempty" xml:"WyidList,omitempty" type:"Repeated"`
}

func (s ActiveRightsRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveRightsRequest) GoString() string {
	return s.String()
}

func (s *ActiveRightsRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *ActiveRightsRequest) GetCustomImageId() *string {
	return s.CustomImageId
}

func (s *ActiveRightsRequest) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *ActiveRightsRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *ActiveRightsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ActiveRightsRequest) GetHibernate() *int32 {
	return s.Hibernate
}

func (s *ActiveRightsRequest) GetLanguage() *string {
	return s.Language
}

func (s *ActiveRightsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ActiveRightsRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ActiveRightsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ActiveRightsRequest) GetRightsAssignedUser() []*string {
	return s.RightsAssignedUser
}

func (s *ActiveRightsRequest) GetRightsGpuSpec() *string {
	return s.RightsGpuSpec
}

func (s *ActiveRightsRequest) GetRightsIds() []*string {
	return s.RightsIds
}

func (s *ActiveRightsRequest) GetRightsResourceArea() *string {
	return s.RightsResourceArea
}

func (s *ActiveRightsRequest) GetRightsResourceImageOs() *string {
	return s.RightsResourceImageOs
}

func (s *ActiveRightsRequest) GetRightsResourceRegion() *string {
	return s.RightsResourceRegion
}

func (s *ActiveRightsRequest) GetRightsResourceType() *string {
	return s.RightsResourceType
}

func (s *ActiveRightsRequest) GetRightsSpec() *string {
	return s.RightsSpec
}

func (s *ActiveRightsRequest) GetRootDiskSize() *int32 {
	return s.RootDiskSize
}

func (s *ActiveRightsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ActiveRightsRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ActiveRightsRequest) GetWorkspaceIds() []*string {
	return s.WorkspaceIds
}

func (s *ActiveRightsRequest) GetWyidList() []*string {
	return s.WyidList
}

func (s *ActiveRightsRequest) SetClientIp(v string) *ActiveRightsRequest {
	s.ClientIp = &v
	return s
}

func (s *ActiveRightsRequest) SetCustomImageId(v string) *ActiveRightsRequest {
	s.CustomImageId = &v
	return s
}

func (s *ActiveRightsRequest) SetDataDiskSize(v int32) *ActiveRightsRequest {
	s.DataDiskSize = &v
	return s
}

func (s *ActiveRightsRequest) SetDesktopType(v string) *ActiveRightsRequest {
	s.DesktopType = &v
	return s
}

func (s *ActiveRightsRequest) SetEndTime(v string) *ActiveRightsRequest {
	s.EndTime = &v
	return s
}

func (s *ActiveRightsRequest) SetHibernate(v int32) *ActiveRightsRequest {
	s.Hibernate = &v
	return s
}

func (s *ActiveRightsRequest) SetLanguage(v string) *ActiveRightsRequest {
	s.Language = &v
	return s
}

func (s *ActiveRightsRequest) SetOfficeSiteId(v string) *ActiveRightsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ActiveRightsRequest) SetPolicyGroupId(v string) *ActiveRightsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ActiveRightsRequest) SetResourceGroupId(v string) *ActiveRightsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ActiveRightsRequest) SetRightsAssignedUser(v []*string) *ActiveRightsRequest {
	s.RightsAssignedUser = v
	return s
}

func (s *ActiveRightsRequest) SetRightsGpuSpec(v string) *ActiveRightsRequest {
	s.RightsGpuSpec = &v
	return s
}

func (s *ActiveRightsRequest) SetRightsIds(v []*string) *ActiveRightsRequest {
	s.RightsIds = v
	return s
}

func (s *ActiveRightsRequest) SetRightsResourceArea(v string) *ActiveRightsRequest {
	s.RightsResourceArea = &v
	return s
}

func (s *ActiveRightsRequest) SetRightsResourceImageOs(v string) *ActiveRightsRequest {
	s.RightsResourceImageOs = &v
	return s
}

func (s *ActiveRightsRequest) SetRightsResourceRegion(v string) *ActiveRightsRequest {
	s.RightsResourceRegion = &v
	return s
}

func (s *ActiveRightsRequest) SetRightsResourceType(v string) *ActiveRightsRequest {
	s.RightsResourceType = &v
	return s
}

func (s *ActiveRightsRequest) SetRightsSpec(v string) *ActiveRightsRequest {
	s.RightsSpec = &v
	return s
}

func (s *ActiveRightsRequest) SetRootDiskSize(v int32) *ActiveRightsRequest {
	s.RootDiskSize = &v
	return s
}

func (s *ActiveRightsRequest) SetStartTime(v string) *ActiveRightsRequest {
	s.StartTime = &v
	return s
}

func (s *ActiveRightsRequest) SetVSwitchId(v string) *ActiveRightsRequest {
	s.VSwitchId = &v
	return s
}

func (s *ActiveRightsRequest) SetWorkspaceIds(v []*string) *ActiveRightsRequest {
	s.WorkspaceIds = v
	return s
}

func (s *ActiveRightsRequest) SetWyidList(v []*string) *ActiveRightsRequest {
	s.WyidList = v
	return s
}

func (s *ActiveRightsRequest) Validate() error {
	return dara.Validate(s)
}
