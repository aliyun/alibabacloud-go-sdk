// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtectdBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateProtectdBranchRequest
	GetAccessToken() *string
	SetAllowMergeRoles(v []*int32) *CreateProtectdBranchRequest
	GetAllowMergeRoles() []*int32
	SetAllowMergeUserIds(v []*string) *CreateProtectdBranchRequest
	GetAllowMergeUserIds() []*string
	SetAllowPushRoles(v []*int32) *CreateProtectdBranchRequest
	GetAllowPushRoles() []*int32
	SetAllowPushUserIds(v []*string) *CreateProtectdBranchRequest
	GetAllowPushUserIds() []*string
	SetBranch(v string) *CreateProtectdBranchRequest
	GetBranch() *string
	SetId(v int64) *CreateProtectdBranchRequest
	GetId() *int64
	SetMergeRequestSetting(v *CreateProtectdBranchRequestMergeRequestSetting) *CreateProtectdBranchRequest
	GetMergeRequestSetting() *CreateProtectdBranchRequestMergeRequestSetting
	SetTestSettingDTO(v *CreateProtectdBranchRequestTestSettingDTO) *CreateProtectdBranchRequest
	GetTestSettingDTO() *CreateProtectdBranchRequestTestSettingDTO
	SetOrganizationId(v string) *CreateProtectdBranchRequest
	GetOrganizationId() *string
}

type CreateProtectdBranchRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	AllowMergeRoles   []*int32  `json:"allowMergeRoles,omitempty" xml:"allowMergeRoles,omitempty" type:"Repeated"`
	AllowMergeUserIds []*string `json:"allowMergeUserIds,omitempty" xml:"allowMergeUserIds,omitempty" type:"Repeated"`
	AllowPushRoles    []*int32  `json:"allowPushRoles,omitempty" xml:"allowPushRoles,omitempty" type:"Repeated"`
	AllowPushUserIds  []*string `json:"allowPushUserIds,omitempty" xml:"allowPushUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// protectBranch
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// --
	Id                  *int64                                          `json:"id,omitempty" xml:"id,omitempty"`
	MergeRequestSetting *CreateProtectdBranchRequestMergeRequestSetting `json:"mergeRequestSetting,omitempty" xml:"mergeRequestSetting,omitempty" type:"Struct"`
	TestSettingDTO      *CreateProtectdBranchRequestTestSettingDTO      `json:"testSettingDTO,omitempty" xml:"testSettingDTO,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateProtectdBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchRequest) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateProtectdBranchRequest) GetAllowMergeRoles() []*int32 {
	return s.AllowMergeRoles
}

func (s *CreateProtectdBranchRequest) GetAllowMergeUserIds() []*string {
	return s.AllowMergeUserIds
}

func (s *CreateProtectdBranchRequest) GetAllowPushRoles() []*int32 {
	return s.AllowPushRoles
}

func (s *CreateProtectdBranchRequest) GetAllowPushUserIds() []*string {
	return s.AllowPushUserIds
}

func (s *CreateProtectdBranchRequest) GetBranch() *string {
	return s.Branch
}

func (s *CreateProtectdBranchRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateProtectdBranchRequest) GetMergeRequestSetting() *CreateProtectdBranchRequestMergeRequestSetting {
	return s.MergeRequestSetting
}

func (s *CreateProtectdBranchRequest) GetTestSettingDTO() *CreateProtectdBranchRequestTestSettingDTO {
	return s.TestSettingDTO
}

func (s *CreateProtectdBranchRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateProtectdBranchRequest) SetAccessToken(v string) *CreateProtectdBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateProtectdBranchRequest) SetAllowMergeRoles(v []*int32) *CreateProtectdBranchRequest {
	s.AllowMergeRoles = v
	return s
}

func (s *CreateProtectdBranchRequest) SetAllowMergeUserIds(v []*string) *CreateProtectdBranchRequest {
	s.AllowMergeUserIds = v
	return s
}

func (s *CreateProtectdBranchRequest) SetAllowPushRoles(v []*int32) *CreateProtectdBranchRequest {
	s.AllowPushRoles = v
	return s
}

func (s *CreateProtectdBranchRequest) SetAllowPushUserIds(v []*string) *CreateProtectdBranchRequest {
	s.AllowPushUserIds = v
	return s
}

func (s *CreateProtectdBranchRequest) SetBranch(v string) *CreateProtectdBranchRequest {
	s.Branch = &v
	return s
}

func (s *CreateProtectdBranchRequest) SetId(v int64) *CreateProtectdBranchRequest {
	s.Id = &v
	return s
}

func (s *CreateProtectdBranchRequest) SetMergeRequestSetting(v *CreateProtectdBranchRequestMergeRequestSetting) *CreateProtectdBranchRequest {
	s.MergeRequestSetting = v
	return s
}

func (s *CreateProtectdBranchRequest) SetTestSettingDTO(v *CreateProtectdBranchRequestTestSettingDTO) *CreateProtectdBranchRequest {
	s.TestSettingDTO = v
	return s
}

func (s *CreateProtectdBranchRequest) SetOrganizationId(v string) *CreateProtectdBranchRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateProtectdBranchRequest) Validate() error {
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

type CreateProtectdBranchRequestMergeRequestSetting struct {
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

func (s CreateProtectdBranchRequestMergeRequestSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchRequestMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) GetAllowMergeRequestRoles() []*int32 {
	return s.AllowMergeRequestRoles
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) GetDefaultAssignees() []*string {
	return s.DefaultAssignees
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) GetIsAllowSelfApproval() *bool {
	return s.IsAllowSelfApproval
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) GetIsRequireDiscussionProcessed() *bool {
	return s.IsRequireDiscussionProcessed
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) GetIsResetApprovalWhenNewPush() *bool {
	return s.IsResetApprovalWhenNewPush
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) GetMinimumApproval() *int32 {
	return s.MinimumApproval
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) GetMrMode() *string {
	return s.MrMode
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) GetWhiteList() *string {
	return s.WhiteList
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *CreateProtectdBranchRequestMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) SetDefaultAssignees(v []*string) *CreateProtectdBranchRequestMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) SetIsAllowSelfApproval(v bool) *CreateProtectdBranchRequestMergeRequestSetting {
	s.IsAllowSelfApproval = &v
	return s
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *CreateProtectdBranchRequestMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) SetIsRequired(v bool) *CreateProtectdBranchRequestMergeRequestSetting {
	s.IsRequired = &v
	return s
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) SetIsResetApprovalWhenNewPush(v bool) *CreateProtectdBranchRequestMergeRequestSetting {
	s.IsResetApprovalWhenNewPush = &v
	return s
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) SetMinimumApproval(v int32) *CreateProtectdBranchRequestMergeRequestSetting {
	s.MinimumApproval = &v
	return s
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) SetMrMode(v string) *CreateProtectdBranchRequestMergeRequestSetting {
	s.MrMode = &v
	return s
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) SetWhiteList(v string) *CreateProtectdBranchRequestMergeRequestSetting {
	s.WhiteList = &v
	return s
}

func (s *CreateProtectdBranchRequestMergeRequestSetting) Validate() error {
	return dara.Validate(s)
}

type CreateProtectdBranchRequestTestSettingDTO struct {
	CheckConfig             *CreateProtectdBranchRequestTestSettingDTOCheckConfig             `json:"checkConfig,omitempty" xml:"checkConfig,omitempty" type:"Struct"`
	CheckTaskQualityConfig  *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig  `json:"checkTaskQualityConfig,omitempty" xml:"checkTaskQualityConfig,omitempty" type:"Struct"`
	CodeGuidelinesDetection *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection `json:"codeGuidelinesDetection,omitempty" xml:"codeGuidelinesDetection,omitempty" type:"Struct"`
	// example:
	//
	// false
	IsRequired             *bool                                                            `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	SensitiveInfoDetection *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection `json:"sensitiveInfoDetection,omitempty" xml:"sensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s CreateProtectdBranchRequestTestSettingDTO) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchRequestTestSettingDTO) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchRequestTestSettingDTO) GetCheckConfig() *CreateProtectdBranchRequestTestSettingDTOCheckConfig {
	return s.CheckConfig
}

func (s *CreateProtectdBranchRequestTestSettingDTO) GetCheckTaskQualityConfig() *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig {
	return s.CheckTaskQualityConfig
}

func (s *CreateProtectdBranchRequestTestSettingDTO) GetCodeGuidelinesDetection() *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection {
	return s.CodeGuidelinesDetection
}

func (s *CreateProtectdBranchRequestTestSettingDTO) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *CreateProtectdBranchRequestTestSettingDTO) GetSensitiveInfoDetection() *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection {
	return s.SensitiveInfoDetection
}

func (s *CreateProtectdBranchRequestTestSettingDTO) SetCheckConfig(v *CreateProtectdBranchRequestTestSettingDTOCheckConfig) *CreateProtectdBranchRequestTestSettingDTO {
	s.CheckConfig = v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTO) SetCheckTaskQualityConfig(v *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) *CreateProtectdBranchRequestTestSettingDTO {
	s.CheckTaskQualityConfig = v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTO) SetCodeGuidelinesDetection(v *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection) *CreateProtectdBranchRequestTestSettingDTO {
	s.CodeGuidelinesDetection = v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTO) SetIsRequired(v bool) *CreateProtectdBranchRequestTestSettingDTO {
	s.IsRequired = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTO) SetSensitiveInfoDetection(v *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection) *CreateProtectdBranchRequestTestSettingDTO {
	s.SensitiveInfoDetection = v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTO) Validate() error {
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

type CreateProtectdBranchRequestTestSettingDTOCheckConfig struct {
	CheckItems []*CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems `json:"checkItems,omitempty" xml:"checkItems,omitempty" type:"Repeated"`
}

func (s CreateProtectdBranchRequestTestSettingDTOCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchRequestTestSettingDTOCheckConfig) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckConfig) GetCheckItems() []*CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems {
	return s.CheckItems
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckConfig) SetCheckItems(v []*CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems) *CreateProtectdBranchRequestTestSettingDTOCheckConfig {
	s.CheckItems = v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckConfig) Validate() error {
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

type CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems struct {
	// example:
	//
	// false
	IsRequired *bool   `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems) GetName() *string {
	return s.Name
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems) SetIsRequired(v bool) *CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems {
	s.IsRequired = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems) SetName(v string) *CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems {
	s.Name = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckConfigCheckItems) Validate() error {
	return dara.Validate(s)
}

type CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig struct {
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

func (s CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) GetBizNo() *string {
	return s.BizNo
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) GetMessage() *string {
	return s.Message
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) SetBizNo(v string) *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig {
	s.BizNo = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) SetEnabled(v bool) *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig {
	s.Enabled = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) SetMessage(v string) *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig {
	s.Message = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) SetTaskName(v string) *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig {
	s.TaskName = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOCheckTaskQualityConfig) Validate() error {
	return dara.Validate(s)
}

type CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_guide_lines
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection) GetMessage() *string {
	return s.Message
}

func (s *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection) SetEnabled(v bool) *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection) SetMessage(v string) *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection {
	s.Message = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOCodeGuidelinesDetection) Validate() error {
	return dara.Validate(s)
}

type CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_sensitive_info
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection) GetMessage() *string {
	return s.Message
}

func (s *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection) SetEnabled(v bool) *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection) SetMessage(v string) *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection {
	s.Message = &v
	return s
}

func (s *CreateProtectdBranchRequestTestSettingDTOSensitiveInfoDetection) Validate() error {
	return dara.Validate(s)
}
