// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtectdBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateProtectdBranchResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateProtectdBranchResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateProtectdBranchResponseBody
	GetRequestId() *string
	SetResult(v *CreateProtectdBranchResponseBodyResult) *CreateProtectdBranchResponseBody
	GetResult() *CreateProtectdBranchResponseBodyResult
	SetSuccess(v bool) *CreateProtectdBranchResponseBody
	GetSuccess() *bool
}

type CreateProtectdBranchResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// C2F153F6-BB43-50C4-9F4F-40593203E19A
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateProtectdBranchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateProtectdBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateProtectdBranchResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateProtectdBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProtectdBranchResponseBody) GetResult() *CreateProtectdBranchResponseBodyResult {
	return s.Result
}

func (s *CreateProtectdBranchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProtectdBranchResponseBody) SetErrorCode(v string) *CreateProtectdBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProtectdBranchResponseBody) SetErrorMessage(v string) *CreateProtectdBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateProtectdBranchResponseBody) SetRequestId(v string) *CreateProtectdBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProtectdBranchResponseBody) SetResult(v *CreateProtectdBranchResponseBodyResult) *CreateProtectdBranchResponseBody {
	s.Result = v
	return s
}

func (s *CreateProtectdBranchResponseBody) SetSuccess(v bool) *CreateProtectdBranchResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProtectdBranchResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProtectdBranchResponseBodyResult struct {
	AllowMergeRoles   []*int32 `json:"allowMergeRoles,omitempty" xml:"allowMergeRoles,omitempty" type:"Repeated"`
	AllowMergeUserIds []*int64 `json:"allowMergeUserIds,omitempty" xml:"allowMergeUserIds,omitempty" type:"Repeated"`
	AllowPushRoles    []*int32 `json:"allowPushRoles,omitempty" xml:"allowPushRoles,omitempty" type:"Repeated"`
	AllowPushUserIds  []*int64 `json:"allowPushUserIds,omitempty" xml:"allowPushUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// protectBranch
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 5240
	Id                  *int64                                                     `json:"id,omitempty" xml:"id,omitempty"`
	MergeRequestSetting *CreateProtectdBranchResponseBodyResultMergeRequestSetting `json:"mergeRequestSetting,omitempty" xml:"mergeRequestSetting,omitempty" type:"Struct"`
	TestSettingDTO      *CreateProtectdBranchResponseBodyResultTestSettingDTO      `json:"testSettingDTO,omitempty" xml:"testSettingDTO,omitempty" type:"Struct"`
}

func (s CreateProtectdBranchResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponseBodyResult) GetAllowMergeRoles() []*int32 {
	return s.AllowMergeRoles
}

func (s *CreateProtectdBranchResponseBodyResult) GetAllowMergeUserIds() []*int64 {
	return s.AllowMergeUserIds
}

func (s *CreateProtectdBranchResponseBodyResult) GetAllowPushRoles() []*int32 {
	return s.AllowPushRoles
}

func (s *CreateProtectdBranchResponseBodyResult) GetAllowPushUserIds() []*int64 {
	return s.AllowPushUserIds
}

func (s *CreateProtectdBranchResponseBodyResult) GetBranch() *string {
	return s.Branch
}

func (s *CreateProtectdBranchResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateProtectdBranchResponseBodyResult) GetMergeRequestSetting() *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	return s.MergeRequestSetting
}

func (s *CreateProtectdBranchResponseBodyResult) GetTestSettingDTO() *CreateProtectdBranchResponseBodyResultTestSettingDTO {
	return s.TestSettingDTO
}

func (s *CreateProtectdBranchResponseBodyResult) SetAllowMergeRoles(v []*int32) *CreateProtectdBranchResponseBodyResult {
	s.AllowMergeRoles = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResult) SetAllowMergeUserIds(v []*int64) *CreateProtectdBranchResponseBodyResult {
	s.AllowMergeUserIds = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResult) SetAllowPushRoles(v []*int32) *CreateProtectdBranchResponseBodyResult {
	s.AllowPushRoles = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResult) SetAllowPushUserIds(v []*int64) *CreateProtectdBranchResponseBodyResult {
	s.AllowPushUserIds = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResult) SetBranch(v string) *CreateProtectdBranchResponseBodyResult {
	s.Branch = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResult) SetId(v int64) *CreateProtectdBranchResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResult) SetMergeRequestSetting(v *CreateProtectdBranchResponseBodyResultMergeRequestSetting) *CreateProtectdBranchResponseBodyResult {
	s.MergeRequestSetting = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResult) SetTestSettingDTO(v *CreateProtectdBranchResponseBodyResultTestSettingDTO) *CreateProtectdBranchResponseBodyResult {
	s.TestSettingDTO = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResult) Validate() error {
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

type CreateProtectdBranchResponseBodyResultMergeRequestSetting struct {
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

func (s CreateProtectdBranchResponseBodyResultMergeRequestSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponseBodyResultMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) GetAllowMergeRequestRoles() []*int32 {
	return s.AllowMergeRequestRoles
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) GetDefaultAssignees() []*string {
	return s.DefaultAssignees
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) GetIsAllowSelfApproval() *bool {
	return s.IsAllowSelfApproval
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) GetIsRequireDiscussionProcessed() *bool {
	return s.IsRequireDiscussionProcessed
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) GetIsResetApprovalWhenNewPush() *bool {
	return s.IsResetApprovalWhenNewPush
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) GetMinimumApproval() *int32 {
	return s.MinimumApproval
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) GetMrMode() *string {
	return s.MrMode
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) GetWhiteList() *string {
	return s.WhiteList
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) SetDefaultAssignees(v []*string) *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) SetIsAllowSelfApproval(v bool) *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	s.IsAllowSelfApproval = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) SetIsRequired(v bool) *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	s.IsRequired = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) SetIsResetApprovalWhenNewPush(v bool) *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	s.IsResetApprovalWhenNewPush = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) SetMinimumApproval(v int32) *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	s.MinimumApproval = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) SetMrMode(v string) *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	s.MrMode = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) SetWhiteList(v string) *CreateProtectdBranchResponseBodyResultMergeRequestSetting {
	s.WhiteList = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultMergeRequestSetting) Validate() error {
	return dara.Validate(s)
}

type CreateProtectdBranchResponseBodyResultTestSettingDTO struct {
	CheckConfig             *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig             `json:"checkConfig,omitempty" xml:"checkConfig,omitempty" type:"Struct"`
	CheckTaskQualityConfig  *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig  `json:"checkTaskQualityConfig,omitempty" xml:"checkTaskQualityConfig,omitempty" type:"Struct"`
	CodeGuidelinesDetection *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection `json:"codeGuidelinesDetection,omitempty" xml:"codeGuidelinesDetection,omitempty" type:"Struct"`
	// example:
	//
	// false
	IsRequired             *bool                                                                       `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	SensitiveInfoDetection *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection `json:"sensitiveInfoDetection,omitempty" xml:"sensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTO) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTO) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) GetCheckConfig() *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig {
	return s.CheckConfig
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) GetCheckTaskQualityConfig() *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	return s.CheckTaskQualityConfig
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) GetCodeGuidelinesDetection() *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	return s.CodeGuidelinesDetection
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) GetSensitiveInfoDetection() *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	return s.SensitiveInfoDetection
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) SetCheckConfig(v *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig) *CreateProtectdBranchResponseBodyResultTestSettingDTO {
	s.CheckConfig = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) SetCheckTaskQualityConfig(v *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) *CreateProtectdBranchResponseBodyResultTestSettingDTO {
	s.CheckTaskQualityConfig = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) SetCodeGuidelinesDetection(v *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection) *CreateProtectdBranchResponseBodyResultTestSettingDTO {
	s.CodeGuidelinesDetection = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) SetIsRequired(v bool) *CreateProtectdBranchResponseBodyResultTestSettingDTO {
	s.IsRequired = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) SetSensitiveInfoDetection(v *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection) *CreateProtectdBranchResponseBodyResultTestSettingDTO {
	s.SensitiveInfoDetection = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTO) Validate() error {
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

type CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig struct {
	CheckItems []*CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems `json:"checkItems,omitempty" xml:"checkItems,omitempty" type:"Repeated"`
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig) GetCheckItems() []*CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	return s.CheckItems
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig) SetCheckItems(v []*CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems) *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig {
	s.CheckItems = v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfig) Validate() error {
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

type CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems struct {
	// example:
	//
	// false
	IsRequired *bool   `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems) GetName() *string {
	return s.Name
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetIsRequired(v bool) *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.IsRequired = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetName(v string) *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.Name = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckConfigCheckItems) Validate() error {
	return dara.Validate(s)
}

type CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig struct {
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

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetBizNo() *string {
	return s.BizNo
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetMessage() *string {
	return s.Message
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetBizNo(v string) *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.BizNo = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetEnabled(v bool) *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Enabled = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetMessage(v string) *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Message = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetTaskName(v string) *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.TaskName = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCheckTaskQualityConfig) Validate() error {
	return dara.Validate(s)
}

type CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_guide_lines
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GetMessage() *string {
	return s.Message
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetEnabled(v bool) *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetMessage(v string) *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Message = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOCodeGuidelinesDetection) Validate() error {
	return dara.Validate(s)
}

type CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_sensitive_info
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection) GetMessage() *string {
	return s.Message
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetEnabled(v bool) *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetMessage(v string) *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Message = &v
	return s
}

func (s *CreateProtectdBranchResponseBodyResultTestSettingDTOSensitiveInfoDetection) Validate() error {
	return dara.Validate(s)
}
