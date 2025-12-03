// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProtectedBranchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateProtectedBranchesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateProtectedBranchesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateProtectedBranchesResponseBody
	GetRequestId() *string
	SetResult(v *UpdateProtectedBranchesResponseBodyResult) *UpdateProtectedBranchesResponseBody
	GetResult() *UpdateProtectedBranchesResponseBodyResult
	SetSuccess(v bool) *UpdateProtectedBranchesResponseBody
	GetSuccess() *bool
}

type UpdateProtectedBranchesResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// A35869D5-BB29-5F84-A4DD-B09985EA2AFA
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateProtectedBranchesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProtectedBranchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateProtectedBranchesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateProtectedBranchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProtectedBranchesResponseBody) GetResult() *UpdateProtectedBranchesResponseBodyResult {
	return s.Result
}

func (s *UpdateProtectedBranchesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProtectedBranchesResponseBody) SetErrorCode(v string) *UpdateProtectedBranchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBody) SetErrorMessage(v string) *UpdateProtectedBranchesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBody) SetRequestId(v string) *UpdateProtectedBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBody) SetResult(v *UpdateProtectedBranchesResponseBodyResult) *UpdateProtectedBranchesResponseBody {
	s.Result = v
	return s
}

func (s *UpdateProtectedBranchesResponseBody) SetSuccess(v bool) *UpdateProtectedBranchesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProtectedBranchesResponseBodyResult struct {
	AllowMergeRoles   []*int32 `json:"allowMergeRoles,omitempty" xml:"allowMergeRoles,omitempty" type:"Repeated"`
	AllowMergeUserIds []*int64 `json:"allowMergeUserIds,omitempty" xml:"allowMergeUserIds,omitempty" type:"Repeated"`
	AllowPushRoles    []*int32 `json:"allowPushRoles,omitempty" xml:"allowPushRoles,omitempty" type:"Repeated"`
	AllowPushUserIds  []*int64 `json:"allowPushUserIds,omitempty" xml:"allowPushUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 123456
	Id                  *int64                                                        `json:"id,omitempty" xml:"id,omitempty"`
	MergeRequestSetting *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting `json:"mergeRequestSetting,omitempty" xml:"mergeRequestSetting,omitempty" type:"Struct"`
	TestSettingDTO      *UpdateProtectedBranchesResponseBodyResultTestSettingDTO      `json:"testSettingDTO,omitempty" xml:"testSettingDTO,omitempty" type:"Struct"`
}

func (s UpdateProtectedBranchesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResult) GetAllowMergeRoles() []*int32 {
	return s.AllowMergeRoles
}

func (s *UpdateProtectedBranchesResponseBodyResult) GetAllowMergeUserIds() []*int64 {
	return s.AllowMergeUserIds
}

func (s *UpdateProtectedBranchesResponseBodyResult) GetAllowPushRoles() []*int32 {
	return s.AllowPushRoles
}

func (s *UpdateProtectedBranchesResponseBodyResult) GetAllowPushUserIds() []*int64 {
	return s.AllowPushUserIds
}

func (s *UpdateProtectedBranchesResponseBodyResult) GetBranch() *string {
	return s.Branch
}

func (s *UpdateProtectedBranchesResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *UpdateProtectedBranchesResponseBodyResult) GetMergeRequestSetting() *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	return s.MergeRequestSetting
}

func (s *UpdateProtectedBranchesResponseBodyResult) GetTestSettingDTO() *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	return s.TestSettingDTO
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetAllowMergeRoles(v []*int32) *UpdateProtectedBranchesResponseBodyResult {
	s.AllowMergeRoles = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetAllowMergeUserIds(v []*int64) *UpdateProtectedBranchesResponseBodyResult {
	s.AllowMergeUserIds = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetAllowPushRoles(v []*int32) *UpdateProtectedBranchesResponseBodyResult {
	s.AllowPushRoles = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetAllowPushUserIds(v []*int64) *UpdateProtectedBranchesResponseBodyResult {
	s.AllowPushUserIds = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetBranch(v string) *UpdateProtectedBranchesResponseBodyResult {
	s.Branch = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetId(v int64) *UpdateProtectedBranchesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetMergeRequestSetting(v *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) *UpdateProtectedBranchesResponseBodyResult {
	s.MergeRequestSetting = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) SetTestSettingDTO(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) *UpdateProtectedBranchesResponseBodyResult {
	s.TestSettingDTO = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResult) Validate() error {
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

type UpdateProtectedBranchesResponseBodyResultMergeRequestSetting struct {
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

func (s UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GetAllowMergeRequestRoles() []*int32 {
	return s.AllowMergeRequestRoles
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GetDefaultAssignees() []*string {
	return s.DefaultAssignees
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GetIsAllowSelfApproval() *bool {
	return s.IsAllowSelfApproval
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GetIsRequireDiscussionProcessed() *bool {
	return s.IsRequireDiscussionProcessed
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GetIsResetApprovalWhenNewPush() *bool {
	return s.IsResetApprovalWhenNewPush
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GetMinimumApproval() *int32 {
	return s.MinimumApproval
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GetMrMode() *string {
	return s.MrMode
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) GetWhiteList() *string {
	return s.WhiteList
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetDefaultAssignees(v []*string) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsAllowSelfApproval(v bool) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsAllowSelfApproval = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsRequired(v bool) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsResetApprovalWhenNewPush(v bool) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsResetApprovalWhenNewPush = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetMinimumApproval(v int32) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.MinimumApproval = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetMrMode(v string) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.MrMode = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) SetWhiteList(v string) *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.WhiteList = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultMergeRequestSetting) Validate() error {
	return dara.Validate(s)
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTO struct {
	CheckConfig             *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig             `json:"checkConfig,omitempty" xml:"checkConfig,omitempty" type:"Struct"`
	CheckTaskQualityConfig  *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig  `json:"checkTaskQualityConfig,omitempty" xml:"checkTaskQualityConfig,omitempty" type:"Struct"`
	CodeGuidelinesDetection *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection `json:"codeGuidelinesDetection,omitempty" xml:"codeGuidelinesDetection,omitempty" type:"Struct"`
	// example:
	//
	// false
	IsRequired             *bool                                                                          `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	SensitiveInfoDetection *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection `json:"sensitiveInfoDetection,omitempty" xml:"sensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTO) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTO) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) GetCheckConfig() *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig {
	return s.CheckConfig
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) GetCheckTaskQualityConfig() *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	return s.CheckTaskQualityConfig
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) GetCodeGuidelinesDetection() *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	return s.CodeGuidelinesDetection
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) GetSensitiveInfoDetection() *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	return s.SensitiveInfoDetection
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetCheckConfig(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CheckConfig = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetCheckTaskQualityConfig(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CheckTaskQualityConfig = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetCodeGuidelinesDetection(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CodeGuidelinesDetection = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetIsRequired(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) SetSensitiveInfoDetection(v *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) *UpdateProtectedBranchesResponseBodyResultTestSettingDTO {
	s.SensitiveInfoDetection = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTO) Validate() error {
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

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig struct {
	CheckItems []*UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems `json:"checkItems,omitempty" xml:"checkItems,omitempty" type:"Repeated"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) GetCheckItems() []*UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	return s.CheckItems
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) SetCheckItems(v []*UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig {
	s.CheckItems = v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) Validate() error {
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

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems struct {
	// example:
	//
	// false
	IsRequired *bool   `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) GetName() *string {
	return s.Name
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetIsRequired(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.IsRequired = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetName(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.Name = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) Validate() error {
	return dara.Validate(s)
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig struct {
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
	// biz-task-quality
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetBizNo() *string {
	return s.BizNo
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetMessage() *string {
	return s.Message
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetBizNo(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.BizNo = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetEnabled(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetMessage(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Message = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetTaskName(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.TaskName = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_guide_lines
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GetMessage() *string {
	return s.Message
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetEnabled(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetMessage(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Message = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) Validate() error {
	return dara.Validate(s)
}

type UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_sensitive_info
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) GetMessage() *string {
	return s.Message
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetEnabled(v bool) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetMessage(v string) *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Message = &v
	return s
}

func (s *UpdateProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) Validate() error {
	return dara.Validate(s)
}
