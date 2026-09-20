// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetProjectDetailResponseBodyData) *GetProjectDetailResponseBody
	GetData() *GetProjectDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetProjectDetailResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetProjectDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProjectDetailResponseBody
	GetSuccess() *bool
}

type GetProjectDetailResponseBody struct {
	// The description of the workspace.
	Data *GetProjectDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 0bc1411515937635973****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the workspace information was queried successfully.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProjectDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectDetailResponseBody) GetData() *GetProjectDetailResponseBodyData {
	return s.Data
}

func (s *GetProjectDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetProjectDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectDetailResponseBody) SetData(v *GetProjectDetailResponseBodyData) *GetProjectDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetProjectDetailResponseBody) SetHttpStatusCode(v int32) *GetProjectDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetProjectDetailResponseBody) SetRequestId(v string) *GetProjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectDetailResponseBody) SetSuccess(v bool) *GetProjectDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectDetailResponseBodyData struct {
	// The identifier of the default data integration resource group for nodes.
	//
	// example:
	//
	// group_280749521****
	DefaultDiResourceGroupIdentifier *string `json:"DefaultDiResourceGroupIdentifier,omitempty" xml:"DefaultDiResourceGroupIdentifier,omitempty"`
	// The project development type. The value is 4.
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
	// The environments contained in the workspace. A workspace in basic mode contains only the production environment (PROD). A workspace in standard mode contains the development environment (DEV) and the production environment (PROD).
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// The time when the workspace was created.
	//
	// example:
	//
	// Oct 10, 2019 3:42:53 PM
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the workspace was last modified.
	//
	// example:
	//
	// Dec 3, 2019 9:12:20 PM
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether downloads are allowed.
	//
	// example:
	//
	// 1
	IsAllowDownload *int32 `json:"IsAllowDownload,omitempty" xml:"IsAllowDownload,omitempty"`
	// Indicates whether the workspace is the default workspace. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// example:
	//
	// 1
	IsDefault *int32 `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// abc
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 27
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The identifier of the workspace.
	//
	// example:
	//
	// abc
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The workspace mode. A value of 2 indicates basic mode. A value of 3 indicates standard mode.
	//
	// example:
	//
	// 1
	ProjectMode *int32 `json:"ProjectMode,omitempty" xml:"ProjectMode,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// abc
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The Alibaba Cloud ID of the workspace owner.
	//
	// example:
	//
	// 18229311****
	ProjectOwnerBaseId *string `json:"ProjectOwnerBaseId,omitempty" xml:"ProjectOwnerBaseId,omitempty"`
	// Indicates whether protected mode is enabled for the workspace.
	//
	// example:
	//
	// 1
	ProtectedMode *int32 `json:"ProtectedMode,omitempty" xml:"ProtectedMode,omitempty"`
	// The type of the workspace. A value of private indicates a private zone. A value of swap indicates a swap zone.
	//
	// example:
	//
	// private
	ResidentArea *string `json:"ResidentArea,omitempty" xml:"ResidentArea,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzbn7pti3zfa
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The default maximum number of automatic reruns upon an error.
	//
	// example:
	//
	// 3
	SchedulerMaxRetryTimes *int32 `json:"SchedulerMaxRetryTimes,omitempty" xml:"SchedulerMaxRetryTimes,omitempty"`
	// The default interval between automatic reruns upon an error.
	//
	// example:
	//
	// 120000
	SchedulerRetryInterval *int32 `json:"SchedulerRetryInterval,omitempty" xml:"SchedulerRetryInterval,omitempty"`
	// The status of the workspace.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The visibility of MaxCompute tables. Valid values:
	//
	// - 0: MaxCompute tables are not visible to tenant users.
	//
	// - 1: MaxCompute tables are visible to tenant users.
	//
	// example:
	//
	// 1
	TablePrivacyMode *int32 `json:"TablePrivacyMode,omitempty" xml:"TablePrivacyMode,omitempty"`
	// The tag information.
	Tags []*GetProjectDetailResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 280749521950784
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Indicates whether a proxy account is used to access the MaxCompute engine.
	//
	// example:
	//
	// true
	UseProxyOdpsAccount *bool `json:"UseProxyOdpsAccount,omitempty" xml:"UseProxyOdpsAccount,omitempty"`
}

func (s GetProjectDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProjectDetailResponseBodyData) GetDefaultDiResourceGroupIdentifier() *string {
	return s.DefaultDiResourceGroupIdentifier
}

func (s *GetProjectDetailResponseBodyData) GetDevelopmentType() *int32 {
	return s.DevelopmentType
}

func (s *GetProjectDetailResponseBodyData) GetDisableDevelopment() *bool {
	return s.DisableDevelopment
}

func (s *GetProjectDetailResponseBodyData) GetEnvTypes() []*string {
	return s.EnvTypes
}

func (s *GetProjectDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetProjectDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetProjectDetailResponseBodyData) GetIsAllowDownload() *int32 {
	return s.IsAllowDownload
}

func (s *GetProjectDetailResponseBodyData) GetIsDefault() *int32 {
	return s.IsDefault
}

func (s *GetProjectDetailResponseBodyData) GetProjectDescription() *string {
	return s.ProjectDescription
}

func (s *GetProjectDetailResponseBodyData) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetProjectDetailResponseBodyData) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetProjectDetailResponseBodyData) GetProjectMode() *int32 {
	return s.ProjectMode
}

func (s *GetProjectDetailResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetProjectDetailResponseBodyData) GetProjectOwnerBaseId() *string {
	return s.ProjectOwnerBaseId
}

func (s *GetProjectDetailResponseBodyData) GetProtectedMode() *int32 {
	return s.ProtectedMode
}

func (s *GetProjectDetailResponseBodyData) GetResidentArea() *string {
	return s.ResidentArea
}

func (s *GetProjectDetailResponseBodyData) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *GetProjectDetailResponseBodyData) GetSchedulerMaxRetryTimes() *int32 {
	return s.SchedulerMaxRetryTimes
}

func (s *GetProjectDetailResponseBodyData) GetSchedulerRetryInterval() *int32 {
	return s.SchedulerRetryInterval
}

func (s *GetProjectDetailResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetProjectDetailResponseBodyData) GetTablePrivacyMode() *int32 {
	return s.TablePrivacyMode
}

func (s *GetProjectDetailResponseBodyData) GetTags() []*GetProjectDetailResponseBodyDataTags {
	return s.Tags
}

func (s *GetProjectDetailResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetProjectDetailResponseBodyData) GetUseProxyOdpsAccount() *bool {
	return s.UseProxyOdpsAccount
}

func (s *GetProjectDetailResponseBodyData) SetDefaultDiResourceGroupIdentifier(v string) *GetProjectDetailResponseBodyData {
	s.DefaultDiResourceGroupIdentifier = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetDevelopmentType(v int32) *GetProjectDetailResponseBodyData {
	s.DevelopmentType = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetDisableDevelopment(v bool) *GetProjectDetailResponseBodyData {
	s.DisableDevelopment = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetEnvTypes(v []*string) *GetProjectDetailResponseBodyData {
	s.EnvTypes = v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetGmtCreate(v string) *GetProjectDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetGmtModified(v string) *GetProjectDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetIsAllowDownload(v int32) *GetProjectDetailResponseBodyData {
	s.IsAllowDownload = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetIsDefault(v int32) *GetProjectDetailResponseBodyData {
	s.IsDefault = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetProjectDescription(v string) *GetProjectDetailResponseBodyData {
	s.ProjectDescription = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetProjectId(v int32) *GetProjectDetailResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetProjectIdentifier(v string) *GetProjectDetailResponseBodyData {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetProjectMode(v int32) *GetProjectDetailResponseBodyData {
	s.ProjectMode = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetProjectName(v string) *GetProjectDetailResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetProjectOwnerBaseId(v string) *GetProjectDetailResponseBodyData {
	s.ProjectOwnerBaseId = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetProtectedMode(v int32) *GetProjectDetailResponseBodyData {
	s.ProtectedMode = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetResidentArea(v string) *GetProjectDetailResponseBodyData {
	s.ResidentArea = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetResourceManagerResourceGroupId(v string) *GetProjectDetailResponseBodyData {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetSchedulerMaxRetryTimes(v int32) *GetProjectDetailResponseBodyData {
	s.SchedulerMaxRetryTimes = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetSchedulerRetryInterval(v int32) *GetProjectDetailResponseBodyData {
	s.SchedulerRetryInterval = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetStatus(v int32) *GetProjectDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetTablePrivacyMode(v int32) *GetProjectDetailResponseBodyData {
	s.TablePrivacyMode = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetTags(v []*GetProjectDetailResponseBodyDataTags) *GetProjectDetailResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetTenantId(v int64) *GetProjectDetailResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) SetUseProxyOdpsAccount(v bool) *GetProjectDetailResponseBodyData {
	s.UseProxyOdpsAccount = &v
	return s
}

func (s *GetProjectDetailResponseBodyData) Validate() error {
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

type GetProjectDetailResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// Env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProjectDetailResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDetailResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetProjectDetailResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetProjectDetailResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetProjectDetailResponseBodyDataTags) SetKey(v string) *GetProjectDetailResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetProjectDetailResponseBodyDataTags) SetValue(v string) *GetProjectDetailResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetProjectDetailResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
