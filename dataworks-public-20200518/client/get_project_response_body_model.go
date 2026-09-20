// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetProjectResponseBodyData) *GetProjectResponseBody
	GetData() *GetProjectResponseBodyData
	SetHttpStatusCode(v int32) *GetProjectResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProjectResponseBody
	GetSuccess() *bool
}

type GetProjectResponseBody struct {
	// The details of the workspace.
	Data *GetProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) GetData() *GetProjectResponseBodyData {
	return s.Data
}

func (s *GetProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectResponseBody) SetData(v *GetProjectResponseBodyData) *GetProjectResponseBody {
	s.Data = v
	return s
}

func (s *GetProjectResponseBody) SetHttpStatusCode(v int32) *GetProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) SetSuccess(v bool) *GetProjectResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectResponseBodyData struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 0
	Appkey *string `json:"Appkey,omitempty" xml:"Appkey,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// false
	BaseProject *bool `json:"BaseProject,omitempty" xml:"BaseProject,omitempty"`
	// The identifier of the default resource group that is automatically assigned when you purchase a MaxCompute exclusive resource group.
	//
	// example:
	//
	// group_280749521****
	DefaultDiResourceGroupIdentifier *string `json:"DefaultDiResourceGroupIdentifier,omitempty" xml:"DefaultDiResourceGroupIdentifier,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 0
	Destination *int32 `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 0
	DevStorageQuota *string `json:"DevStorageQuota,omitempty" xml:"DevStorageQuota,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 4
	DevelopmentType *int32 `json:"DevelopmentType,omitempty" xml:"DevelopmentType,omitempty"`
	// Indicates whether the development role is disabled. Valid values:
	//
	//  - **false*	- (default): The development role is enabled.
	//
	//  - **true**: The development role is disabled.
	//
	// example:
	//
	// true
	DisableDevelopment *bool `json:"DisableDevelopment,omitempty" xml:"DisableDevelopment,omitempty"`
	// The environment context of the workspace.
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// The time when the workspace was created. Example: `Dec 3, 2019 9:12:20 PM`.
	//
	// example:
	//
	// Oct 10, 2019 3:42:53 PM
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the workspace was last modified. Example: `Dec 3, 2019 9:12:20 PM`.
	//
	// example:
	//
	// Dec 3, 2019 9:12:20 PM
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether downloading query results from the IDE is allowed. Valid values:
	//
	// - **1**: Downloading is allowed.
	//
	// - **0**: Downloading is not allowed.
	//
	// example:
	//
	// 1
	IsAllowDownload *int32 `json:"IsAllowDownload,omitempty" xml:"IsAllowDownload,omitempty"`
	// Indicates whether the workspace is the default workspace. Valid values:
	//
	// - **1**: Yes.
	//
	// - **0**: No.
	//
	// example:
	//
	// 1
	IsDefault *int32 `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MaxFlowNode *int32 `json:"MaxFlowNode,omitempty" xml:"MaxFlowNode,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	ProdStorageQuota *string `json:"ProdStorageQuota,omitempty" xml:"ProdStorageQuota,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// test_describe
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 466230
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// test_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The mode of the workspace. Valid values:
	//
	// - **2**: basic mode.
	//
	// - **3**: standard mode.
	//
	// example:
	//
	// 2
	ProjectMode *int32 `json:"ProjectMode,omitempty" xml:"ProjectMode,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The Alibaba Cloud ID of the workspace owner.
	//
	// example:
	//
	// 18229311****
	ProjectOwnerBaseId *string `json:"ProjectOwnerBaseId,omitempty" xml:"ProjectOwnerBaseId,omitempty"`
	// Indicates whether protected mode is enabled for the workspace. Valid values:
	//
	// - **1**: Protected mode is enabled.
	//
	// - **0**: Protected mode is not enabled.
	//
	// example:
	//
	// 1
	ProtectedMode *int32 `json:"ProtectedMode,omitempty" xml:"ProtectedMode,omitempty"`
	// The type of the workspace. Valid values:
	//
	// - **private**: private zone.
	//
	// - **swap**: swap zone.
	//
	// example:
	//
	// private
	ResidentArea *string `json:"ResidentArea,omitempty" xml:"ResidentArea,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzbn7****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The default maximum number of automatic reruns upon an error.
	//
	// example:
	//
	// 3
	SchedulerMaxRetryTimes *int32 `json:"SchedulerMaxRetryTimes,omitempty" xml:"SchedulerMaxRetryTimes,omitempty"`
	// The default interval between automatic reruns upon an error. Unit: ms. The maximum value is 30 minutes. Note the unit conversion.
	//
	// example:
	//
	// 120000
	SchedulerRetryInterval *int32 `json:"SchedulerRetryInterval,omitempty" xml:"SchedulerRetryInterval,omitempty"`
	// The status of the workspace. Valid values:
	//
	// - **0**: AVAILABLE. The workspace is running normally.
	//
	// - **1**: DELETED. The workspace has been deleted.
	//
	// - **2**: INITIALIZING. The workspace is being initialized.
	//
	// - **3**: INIT_FAILED. The workspace failed to be initialized.
	//
	// - **4**: FORBIDDEN. The workspace is manually disabled.
	//
	// - **5**: DELETING. The workspace is being deleted.
	//
	// - **6**: DEL_FAILED. The workspace failed to be deleted.
	//
	// - **7**: FROZEN. The workspace is frozen due to overdue payment.
	//
	// - **8**: UPDATING. The workspace is being updated (a compute engine is being added and initialized).
	//
	// - **9**: UPDATE_FAILED. The workspace failed to be updated (a compute engine failed to be added and initialized).
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The visibility of MaxCompute tables. Valid values:
	//
	// - **0**: MaxCompute tables are not visible to users within the tenant.
	//
	// - **1**: MaxCompute tables are visible to users within the tenant.
	//
	// example:
	//
	// 1
	TablePrivacyMode *int32 `json:"TablePrivacyMode,omitempty" xml:"TablePrivacyMode,omitempty"`
	// The list of tags bound to the workspace.
	Tags []*GetProjectResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 280749521
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Indicates whether a proxy account is used to access the MaxCompute engine.
	//
	// example:
	//
	// true
	UseProxyOdpsAccount *bool `json:"UseProxyOdpsAccount,omitempty" xml:"UseProxyOdpsAccount,omitempty"`
}

func (s GetProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyData) GetAppkey() *string {
	return s.Appkey
}

func (s *GetProjectResponseBodyData) GetBaseProject() *bool {
	return s.BaseProject
}

func (s *GetProjectResponseBodyData) GetDefaultDiResourceGroupIdentifier() *string {
	return s.DefaultDiResourceGroupIdentifier
}

func (s *GetProjectResponseBodyData) GetDestination() *int32 {
	return s.Destination
}

func (s *GetProjectResponseBodyData) GetDevStorageQuota() *string {
	return s.DevStorageQuota
}

func (s *GetProjectResponseBodyData) GetDevelopmentType() *int32 {
	return s.DevelopmentType
}

func (s *GetProjectResponseBodyData) GetDisableDevelopment() *bool {
	return s.DisableDevelopment
}

func (s *GetProjectResponseBodyData) GetEnvTypes() []*string {
	return s.EnvTypes
}

func (s *GetProjectResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetProjectResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetProjectResponseBodyData) GetIsAllowDownload() *int32 {
	return s.IsAllowDownload
}

func (s *GetProjectResponseBodyData) GetIsDefault() *int32 {
	return s.IsDefault
}

func (s *GetProjectResponseBodyData) GetMaxFlowNode() *int32 {
	return s.MaxFlowNode
}

func (s *GetProjectResponseBodyData) GetProdStorageQuota() *string {
	return s.ProdStorageQuota
}

func (s *GetProjectResponseBodyData) GetProjectDescription() *string {
	return s.ProjectDescription
}

func (s *GetProjectResponseBodyData) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetProjectResponseBodyData) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetProjectResponseBodyData) GetProjectMode() *int32 {
	return s.ProjectMode
}

func (s *GetProjectResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetProjectResponseBodyData) GetProjectOwnerBaseId() *string {
	return s.ProjectOwnerBaseId
}

func (s *GetProjectResponseBodyData) GetProtectedMode() *int32 {
	return s.ProtectedMode
}

func (s *GetProjectResponseBodyData) GetResidentArea() *string {
	return s.ResidentArea
}

func (s *GetProjectResponseBodyData) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *GetProjectResponseBodyData) GetSchedulerMaxRetryTimes() *int32 {
	return s.SchedulerMaxRetryTimes
}

func (s *GetProjectResponseBodyData) GetSchedulerRetryInterval() *int32 {
	return s.SchedulerRetryInterval
}

func (s *GetProjectResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetProjectResponseBodyData) GetTablePrivacyMode() *int32 {
	return s.TablePrivacyMode
}

func (s *GetProjectResponseBodyData) GetTags() []*GetProjectResponseBodyDataTags {
	return s.Tags
}

func (s *GetProjectResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetProjectResponseBodyData) GetUseProxyOdpsAccount() *bool {
	return s.UseProxyOdpsAccount
}

func (s *GetProjectResponseBodyData) SetAppkey(v string) *GetProjectResponseBodyData {
	s.Appkey = &v
	return s
}

func (s *GetProjectResponseBodyData) SetBaseProject(v bool) *GetProjectResponseBodyData {
	s.BaseProject = &v
	return s
}

func (s *GetProjectResponseBodyData) SetDefaultDiResourceGroupIdentifier(v string) *GetProjectResponseBodyData {
	s.DefaultDiResourceGroupIdentifier = &v
	return s
}

func (s *GetProjectResponseBodyData) SetDestination(v int32) *GetProjectResponseBodyData {
	s.Destination = &v
	return s
}

func (s *GetProjectResponseBodyData) SetDevStorageQuota(v string) *GetProjectResponseBodyData {
	s.DevStorageQuota = &v
	return s
}

func (s *GetProjectResponseBodyData) SetDevelopmentType(v int32) *GetProjectResponseBodyData {
	s.DevelopmentType = &v
	return s
}

func (s *GetProjectResponseBodyData) SetDisableDevelopment(v bool) *GetProjectResponseBodyData {
	s.DisableDevelopment = &v
	return s
}

func (s *GetProjectResponseBodyData) SetEnvTypes(v []*string) *GetProjectResponseBodyData {
	s.EnvTypes = v
	return s
}

func (s *GetProjectResponseBodyData) SetGmtCreate(v string) *GetProjectResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetProjectResponseBodyData) SetGmtModified(v string) *GetProjectResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetProjectResponseBodyData) SetIsAllowDownload(v int32) *GetProjectResponseBodyData {
	s.IsAllowDownload = &v
	return s
}

func (s *GetProjectResponseBodyData) SetIsDefault(v int32) *GetProjectResponseBodyData {
	s.IsDefault = &v
	return s
}

func (s *GetProjectResponseBodyData) SetMaxFlowNode(v int32) *GetProjectResponseBodyData {
	s.MaxFlowNode = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProdStorageQuota(v string) *GetProjectResponseBodyData {
	s.ProdStorageQuota = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectDescription(v string) *GetProjectResponseBodyData {
	s.ProjectDescription = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectId(v int32) *GetProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectIdentifier(v string) *GetProjectResponseBodyData {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectMode(v int32) *GetProjectResponseBodyData {
	s.ProjectMode = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectName(v string) *GetProjectResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectOwnerBaseId(v string) *GetProjectResponseBodyData {
	s.ProjectOwnerBaseId = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProtectedMode(v int32) *GetProjectResponseBodyData {
	s.ProtectedMode = &v
	return s
}

func (s *GetProjectResponseBodyData) SetResidentArea(v string) *GetProjectResponseBodyData {
	s.ResidentArea = &v
	return s
}

func (s *GetProjectResponseBodyData) SetResourceManagerResourceGroupId(v string) *GetProjectResponseBodyData {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *GetProjectResponseBodyData) SetSchedulerMaxRetryTimes(v int32) *GetProjectResponseBodyData {
	s.SchedulerMaxRetryTimes = &v
	return s
}

func (s *GetProjectResponseBodyData) SetSchedulerRetryInterval(v int32) *GetProjectResponseBodyData {
	s.SchedulerRetryInterval = &v
	return s
}

func (s *GetProjectResponseBodyData) SetStatus(v int32) *GetProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetProjectResponseBodyData) SetTablePrivacyMode(v int32) *GetProjectResponseBodyData {
	s.TablePrivacyMode = &v
	return s
}

func (s *GetProjectResponseBodyData) SetTags(v []*GetProjectResponseBodyDataTags) *GetProjectResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetProjectResponseBodyData) SetTenantId(v int64) *GetProjectResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetProjectResponseBodyData) SetUseProxyOdpsAccount(v bool) *GetProjectResponseBodyData {
	s.UseProxyOdpsAccount = &v
	return s
}

func (s *GetProjectResponseBodyData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProjectResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProjectResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetProjectResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetProjectResponseBodyDataTags) SetKey(v string) *GetProjectResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetProjectResponseBodyDataTags) SetValue(v string) *GetProjectResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetProjectResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
