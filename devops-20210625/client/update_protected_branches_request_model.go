// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProtectedBranchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateProtectedBranchesRequest
	GetAccessToken() *string
	SetAllowMergeRoles(v []*int32) *UpdateProtectedBranchesRequest
	GetAllowMergeRoles() []*int32
	SetAllowMergeUserIds(v []*string) *UpdateProtectedBranchesRequest
	GetAllowMergeUserIds() []*string
	SetAllowPushRoles(v []*int32) *UpdateProtectedBranchesRequest
	GetAllowPushRoles() []*int32
	SetAllowPushUserIds(v []*string) *UpdateProtectedBranchesRequest
	GetAllowPushUserIds() []*string
	SetBranch(v string) *UpdateProtectedBranchesRequest
	GetBranch() *string
	SetId(v int64) *UpdateProtectedBranchesRequest
	GetId() *int64
	SetMergeRequestSetting(v *UpdateProtectedBranchesRequestMergeRequestSetting) *UpdateProtectedBranchesRequest
	GetMergeRequestSetting() *UpdateProtectedBranchesRequestMergeRequestSetting
	SetTestSettingDTO(v *UpdateProtectedBranchesRequestTestSettingDTO) *UpdateProtectedBranchesRequest
	GetTestSettingDTO() *UpdateProtectedBranchesRequestTestSettingDTO
	SetOrganizationId(v string) *UpdateProtectedBranchesRequest
	GetOrganizationId() *string
}

type UpdateProtectedBranchesRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken       *string   `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	AllowMergeRoles   []*int32  `json:"allowMergeRoles,omitempty" xml:"allowMergeRoles,omitempty" type:"Repeated"`
	AllowMergeUserIds []*string `json:"allowMergeUserIds,omitempty" xml:"allowMergeUserIds,omitempty" type:"Repeated"`
	AllowPushRoles    []*int32  `json:"allowPushRoles,omitempty" xml:"allowPushRoles,omitempty" type:"Repeated"`
	AllowPushUserIds  []*string `json:"allowPushUserIds,omitempty" xml:"allowPushUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 19224
	Id                  *int64                                             `json:"id,omitempty" xml:"id,omitempty"`
	MergeRequestSetting *UpdateProtectedBranchesRequestMergeRequestSetting `json:"mergeRequestSetting,omitempty" xml:"mergeRequestSetting,omitempty" type:"Struct"`
	TestSettingDTO      *UpdateProtectedBranchesRequestTestSettingDTO      `json:"testSettingDTO,omitempty" xml:"testSettingDTO,omitempty" type:"Struct"`
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateProtectedBranchesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesRequest) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateProtectedBranchesRequest) GetAllowMergeRoles() []*int32 {
	return s.AllowMergeRoles
}

func (s *UpdateProtectedBranchesRequest) GetAllowMergeUserIds() []*string {
	return s.AllowMergeUserIds
}

func (s *UpdateProtectedBranchesRequest) GetAllowPushRoles() []*int32 {
	return s.AllowPushRoles
}

func (s *UpdateProtectedBranchesRequest) GetAllowPushUserIds() []*string {
	return s.AllowPushUserIds
}

func (s *UpdateProtectedBranchesRequest) GetBranch() *string {
	return s.Branch
}

func (s *UpdateProtectedBranchesRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateProtectedBranchesRequest) GetMergeRequestSetting() *UpdateProtectedBranchesRequestMergeRequestSetting {
	return s.MergeRequestSetting
}

func (s *UpdateProtectedBranchesRequest) GetTestSettingDTO() *UpdateProtectedBranchesRequestTestSettingDTO {
	return s.TestSettingDTO
}

func (s *UpdateProtectedBranchesRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateProtectedBranchesRequest) SetAccessToken(v string) *UpdateProtectedBranchesRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetAllowMergeRoles(v []*int32) *UpdateProtectedBranchesRequest {
	s.AllowMergeRoles = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetAllowMergeUserIds(v []*string) *UpdateProtectedBranchesRequest {
	s.AllowMergeUserIds = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetAllowPushRoles(v []*int32) *UpdateProtectedBranchesRequest {
	s.AllowPushRoles = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetAllowPushUserIds(v []*string) *UpdateProtectedBranchesRequest {
	s.AllowPushUserIds = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetBranch(v string) *UpdateProtectedBranchesRequest {
	s.Branch = &v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetId(v int64) *UpdateProtectedBranchesRequest {
	s.Id = &v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetMergeRequestSetting(v *UpdateProtectedBranchesRequestMergeRequestSetting) *UpdateProtectedBranchesRequest {
	s.MergeRequestSetting = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetTestSettingDTO(v *UpdateProtectedBranchesRequestTestSettingDTO) *UpdateProtectedBranchesRequest {
	s.TestSettingDTO = v
	return s
}

func (s *UpdateProtectedBranchesRequest) SetOrganizationId(v string) *UpdateProtectedBranchesRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateProtectedBranchesRequest) Validate() error {
	if s.MergeRequestSetting != nil {
		if err := s.MergeRequestSetting.Validate(); err != nil {
			return err
		}
	}
	if s.TestSettingDTO != nil {
		if err := s.TestSettingDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProtectedBranchesRequestMergeRequestSetting struct {
	AllowMergeRequestRoles []*int32  `json:"allowMergeRequestRoles,omitempty" xml:"allowMergeRequestRoles,omitempty" type:"Repeated"`
	DefaultAssignees       []*string `json:"defaultAssignees,omitempty" xml:"defaultAssignees,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsAllowSelfApproval *bool `json:"isAllowSelfApproval,omitempty" xml:"isAllowSelfApproval,omitempty"`
	// example:
	//
	// true
	IsRequireDiscussionProcessed *bool `json:"isRequireDiscussionProcessed,omitempty" xml:"isRequireDiscussionProcessed,omitempty"`
	// example:
	//
	// true
	IsRequired *bool `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	// example:
	//
	// false
	IsResetApprovalWhenNewPush *bool `json:"isResetApprovalWhenNewPush,omitempty" xml:"isResetApprovalWhenNewPush,omitempty"`
	// example:
	//
	// 1
	MinimumApproval *int32 `json:"minimumApproval,omitempty" xml:"minimumApproval,omitempty"`
	// example:
	//
	// general
	MrMode *string `json:"mrMode,omitempty" xml:"mrMode,omitempty"`
	// example:
	//
	// **.java
	WhiteList *string `json:"whiteList,omitempty" xml:"whiteList,omitempty"`
}

func (s UpdateProtectedBranchesRequestMergeRequestSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesRequestMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) GetAllowMergeRequestRoles() []*int32 {
	return s.AllowMergeRequestRoles
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) GetDefaultAssignees() []*string {
	return s.DefaultAssignees
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) GetIsAllowSelfApproval() *bool {
	return s.IsAllowSelfApproval
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) GetIsRequireDiscussionProcessed() *bool {
	return s.IsRequireDiscussionProcessed
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) GetIsResetApprovalWhenNewPush() *bool {
	return s.IsResetApprovalWhenNewPush
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) GetMinimumApproval() *int32 {
	return s.MinimumApproval
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) GetMrMode() *string {
	return s.MrMode
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) GetWhiteList() *string {
	return s.WhiteList
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetDefaultAssignees(v []*string) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetIsAllowSelfApproval(v bool) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.IsAllowSelfApproval = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetIsRequired(v bool) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetIsResetApprovalWhenNewPush(v bool) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.IsResetApprovalWhenNewPush = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetMinimumApproval(v int32) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.MinimumApproval = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetMrMode(v string) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.MrMode = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) SetWhiteList(v string) *UpdateProtectedBranchesRequestMergeRequestSetting {
	s.WhiteList = &v
	return s
}

func (s *UpdateProtectedBranchesRequestMergeRequestSetting) Validate() error {
	return dara.Validate(s)
}

type UpdateProtectedBranchesRequestTestSettingDTO struct {
	CheckConfig             *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig             `json:"checkConfig,omitempty" xml:"checkConfig,omitempty" type:"Struct"`
	CheckTaskQualityConfig  *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig  `json:"checkTaskQualityConfig,omitempty" xml:"checkTaskQualityConfig,omitempty" type:"Struct"`
	CodeGuidelinesDetection *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection `json:"codeGuidelinesDetection,omitempty" xml:"codeGuidelinesDetection,omitempty" type:"Struct"`
	// example:
	//
	// false
	IsRequired             *bool                                                               `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	SensitiveInfoDetection *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection `json:"sensitiveInfoDetection,omitempty" xml:"sensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTO) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTO) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) GetCheckConfig() *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig {
	return s.CheckConfig
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) GetCheckTaskQualityConfig() *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig {
	return s.CheckTaskQualityConfig
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) GetCodeGuidelinesDetection() *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection {
	return s.CodeGuidelinesDetection
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) GetSensitiveInfoDetection() *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection {
	return s.SensitiveInfoDetection
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetCheckConfig(v *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.CheckConfig = v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetCheckTaskQualityConfig(v *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.CheckTaskQualityConfig = v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetCodeGuidelinesDetection(v *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.CodeGuidelinesDetection = v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetIsRequired(v bool) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) SetSensitiveInfoDetection(v *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) *UpdateProtectedBranchesRequestTestSettingDTO {
	s.SensitiveInfoDetection = v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTO) Validate() error {
	if s.CheckConfig != nil {
		if err := s.CheckConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CheckTaskQualityConfig != nil {
		if err := s.CheckTaskQualityConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CodeGuidelinesDetection != nil {
		if err := s.CodeGuidelinesDetection.Validate(); err != nil {
			return err
		}
	}
	if s.SensitiveInfoDetection != nil {
		if err := s.SensitiveInfoDetection.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProtectedBranchesRequestTestSettingDTOCheckConfig struct {
	CheckItems []*UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems `json:"checkItems,omitempty" xml:"checkItems,omitempty" type:"Repeated"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) GetCheckItems() []*UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems {
	return s.CheckItems
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) SetCheckItems(v []*UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig {
	s.CheckItems = v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfig) Validate() error {
	if s.CheckItems != nil {
		for _, item := range s.CheckItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems struct {
	// example:
	//
	// false
	IsRequired *bool   `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) GetName() *string {
	return s.Name
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) SetIsRequired(v bool) *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) SetName(v string) *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems {
	s.Name = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckConfigCheckItems) Validate() error {
	return dara.Validate(s)
}

type UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig struct {
	// example:
	//
	// 123456
	BizNo *string `json:"bizNo,omitempty" xml:"bizNo,omitempty"`
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_task_quality
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// bz-task-quality
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) GetBizNo() *string {
	return s.BizNo
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) GetMessage() *string {
	return s.Message
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) SetBizNo(v string) *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig {
	s.BizNo = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) SetEnabled(v bool) *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) SetMessage(v string) *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig {
	s.Message = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) SetTaskName(v string) *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig {
	s.TaskName = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCheckTaskQualityConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_guide_lines
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) GetMessage() *string {
	return s.Message
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) SetEnabled(v bool) *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) SetMessage(v string) *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection {
	s.Message = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOCodeGuidelinesDetection) Validate() error {
	return dara.Validate(s)
}

type UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_sensitive_info
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) GetMessage() *string {
	return s.Message
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) SetEnabled(v bool) *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) SetMessage(v string) *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection {
	s.Message = &v
	return s
}

func (s *UpdateProtectedBranchesRequestTestSettingDTOSensitiveInfoDetection) Validate() error {
	return dara.Validate(s)
}
