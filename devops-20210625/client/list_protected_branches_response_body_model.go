// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectedBranchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListProtectedBranchesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListProtectedBranchesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListProtectedBranchesResponseBody
	GetRequestId() *string
	SetResult(v []*ListProtectedBranchesResponseBodyResult) *ListProtectedBranchesResponseBody
	GetResult() []*ListProtectedBranchesResponseBodyResult
	SetSuccess(v bool) *ListProtectedBranchesResponseBody
	GetSuccess() *bool
}

type ListProtectedBranchesResponseBody struct {
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
	// 313A1BF6-63B7-52D4-A098-952221A65254
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListProtectedBranchesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListProtectedBranchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListProtectedBranchesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListProtectedBranchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProtectedBranchesResponseBody) GetResult() []*ListProtectedBranchesResponseBodyResult {
	return s.Result
}

func (s *ListProtectedBranchesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProtectedBranchesResponseBody) SetErrorCode(v string) *ListProtectedBranchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProtectedBranchesResponseBody) SetErrorMessage(v string) *ListProtectedBranchesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListProtectedBranchesResponseBody) SetRequestId(v string) *ListProtectedBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProtectedBranchesResponseBody) SetResult(v []*ListProtectedBranchesResponseBodyResult) *ListProtectedBranchesResponseBody {
	s.Result = v
	return s
}

func (s *ListProtectedBranchesResponseBody) SetSuccess(v bool) *ListProtectedBranchesResponseBody {
	s.Success = &v
	return s
}

func (s *ListProtectedBranchesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProtectedBranchesResponseBodyResult struct {
	AllowMergeRoles   []*int32                                                  `json:"allowMergeRoles,omitempty" xml:"allowMergeRoles,omitempty" type:"Repeated"`
	AllowMergeUserIds []*int64                                                  `json:"allowMergeUserIds,omitempty" xml:"allowMergeUserIds,omitempty" type:"Repeated"`
	AllowMergeUsers   []*ListProtectedBranchesResponseBodyResultAllowMergeUsers `json:"allowMergeUsers,omitempty" xml:"allowMergeUsers,omitempty" type:"Repeated"`
	AllowPushRoles    []*int32                                                  `json:"allowPushRoles,omitempty" xml:"allowPushRoles,omitempty" type:"Repeated"`
	AllowPushUserIds  []*int64                                                  `json:"allowPushUserIds,omitempty" xml:"allowPushUserIds,omitempty" type:"Repeated"`
	AllowPushUsers    []*ListProtectedBranchesResponseBodyResultAllowPushUsers  `json:"allowPushUsers,omitempty" xml:"allowPushUsers,omitempty" type:"Repeated"`
	// example:
	//
	// protectedBranch
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 19285
	Id                  *int64                                                      `json:"id,omitempty" xml:"id,omitempty"`
	Matches             []*string                                                   `json:"matches,omitempty" xml:"matches,omitempty" type:"Repeated"`
	MergeRequestSetting *ListProtectedBranchesResponseBodyResultMergeRequestSetting `json:"mergeRequestSetting,omitempty" xml:"mergeRequestSetting,omitempty" type:"Struct"`
	TestSettingDTO      *ListProtectedBranchesResponseBodyResultTestSettingDTO      `json:"testSettingDTO,omitempty" xml:"testSettingDTO,omitempty" type:"Struct"`
	// example:
	//
	// 2022-03-18 14:24:54
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResult) GetAllowMergeRoles() []*int32 {
	return s.AllowMergeRoles
}

func (s *ListProtectedBranchesResponseBodyResult) GetAllowMergeUserIds() []*int64 {
	return s.AllowMergeUserIds
}

func (s *ListProtectedBranchesResponseBodyResult) GetAllowMergeUsers() []*ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	return s.AllowMergeUsers
}

func (s *ListProtectedBranchesResponseBodyResult) GetAllowPushRoles() []*int32 {
	return s.AllowPushRoles
}

func (s *ListProtectedBranchesResponseBodyResult) GetAllowPushUserIds() []*int64 {
	return s.AllowPushUserIds
}

func (s *ListProtectedBranchesResponseBodyResult) GetAllowPushUsers() []*ListProtectedBranchesResponseBodyResultAllowPushUsers {
	return s.AllowPushUsers
}

func (s *ListProtectedBranchesResponseBodyResult) GetBranch() *string {
	return s.Branch
}

func (s *ListProtectedBranchesResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListProtectedBranchesResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListProtectedBranchesResponseBodyResult) GetMatches() []*string {
	return s.Matches
}

func (s *ListProtectedBranchesResponseBodyResult) GetMergeRequestSetting() *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	return s.MergeRequestSetting
}

func (s *ListProtectedBranchesResponseBodyResult) GetTestSettingDTO() *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	return s.TestSettingDTO
}

func (s *ListProtectedBranchesResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowMergeRoles(v []*int32) *ListProtectedBranchesResponseBodyResult {
	s.AllowMergeRoles = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowMergeUserIds(v []*int64) *ListProtectedBranchesResponseBodyResult {
	s.AllowMergeUserIds = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowMergeUsers(v []*ListProtectedBranchesResponseBodyResultAllowMergeUsers) *ListProtectedBranchesResponseBodyResult {
	s.AllowMergeUsers = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowPushRoles(v []*int32) *ListProtectedBranchesResponseBodyResult {
	s.AllowPushRoles = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowPushUserIds(v []*int64) *ListProtectedBranchesResponseBodyResult {
	s.AllowPushUserIds = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetAllowPushUsers(v []*ListProtectedBranchesResponseBodyResultAllowPushUsers) *ListProtectedBranchesResponseBodyResult {
	s.AllowPushUsers = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetBranch(v string) *ListProtectedBranchesResponseBodyResult {
	s.Branch = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetCreatedAt(v string) *ListProtectedBranchesResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetId(v int64) *ListProtectedBranchesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetMatches(v []*string) *ListProtectedBranchesResponseBodyResult {
	s.Matches = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetMergeRequestSetting(v *ListProtectedBranchesResponseBodyResultMergeRequestSetting) *ListProtectedBranchesResponseBodyResult {
	s.MergeRequestSetting = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetTestSettingDTO(v *ListProtectedBranchesResponseBodyResultTestSettingDTO) *ListProtectedBranchesResponseBodyResult {
	s.TestSettingDTO = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) SetUpdatedAt(v string) *ListProtectedBranchesResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResult) Validate() error {
	if s.AllowMergeUsers != nil {
		for _, item := range s.AllowMergeUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AllowPushUsers != nil {
		for _, item := range s.AllowPushUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListProtectedBranchesResponseBodyResultAllowMergeUsers struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 19238
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// codeup-test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultAllowMergeUsers) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultAllowMergeUsers) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) GetAvatar() *string {
	return s.Avatar
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) GetEmail() *string {
	return s.Email
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) GetId() *int64 {
	return s.Id
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) GetName() *string {
	return s.Name
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) GetUsername() *string {
	return s.Username
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetAvatar(v string) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Avatar = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetEmail(v string) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Email = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetId(v int64) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Id = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetName(v string) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Name = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) SetUsername(v string) *ListProtectedBranchesResponseBodyResultAllowMergeUsers {
	s.Username = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowMergeUsers) Validate() error {
	return dara.Validate(s)
}

type ListProtectedBranchesResponseBodyResultAllowPushUsers struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 19238
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// codeup-test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultAllowPushUsers) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultAllowPushUsers) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) GetAvatar() *string {
	return s.Avatar
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) GetEmail() *string {
	return s.Email
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) GetId() *int64 {
	return s.Id
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) GetName() *string {
	return s.Name
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) GetUsername() *string {
	return s.Username
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetAvatar(v string) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Avatar = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetEmail(v string) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Email = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetId(v int64) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Id = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetName(v string) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Name = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) SetUsername(v string) *ListProtectedBranchesResponseBodyResultAllowPushUsers {
	s.Username = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultAllowPushUsers) Validate() error {
	return dara.Validate(s)
}

type ListProtectedBranchesResponseBodyResultMergeRequestSetting struct {
	AllowMergeRequestRoles []*int32                                                                      `json:"allowMergeRequestRoles,omitempty" xml:"allowMergeRequestRoles,omitempty" type:"Repeated"`
	DefaultAssignees       []*ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees `json:"defaultAssignees,omitempty" xml:"defaultAssignees,omitempty" type:"Repeated"`
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

func (s ListProtectedBranchesResponseBodyResultMergeRequestSetting) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) GetAllowMergeRequestRoles() []*int32 {
	return s.AllowMergeRequestRoles
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) GetDefaultAssignees() []*ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	return s.DefaultAssignees
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) GetIsAllowSelfApproval() *bool {
	return s.IsAllowSelfApproval
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) GetIsRequireDiscussionProcessed() *bool {
	return s.IsRequireDiscussionProcessed
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) GetIsResetApprovalWhenNewPush() *bool {
	return s.IsResetApprovalWhenNewPush
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) GetMinimumApproval() *int32 {
	return s.MinimumApproval
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) GetMrMode() *string {
	return s.MrMode
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) GetWhiteList() *string {
	return s.WhiteList
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetDefaultAssignees(v []*ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsAllowSelfApproval(v bool) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsAllowSelfApproval = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsRequired(v bool) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsRequired = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetIsResetApprovalWhenNewPush(v bool) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.IsResetApprovalWhenNewPush = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetMinimumApproval(v int32) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.MinimumApproval = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetMrMode(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.MrMode = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) SetWhiteList(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSetting {
	s.WhiteList = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSetting) Validate() error {
	if s.DefaultAssignees != nil {
		for _, item := range s.DefaultAssignees {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 19238
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// codeup-test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) GetAvatar() *string {
	return s.Avatar
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) GetEmail() *string {
	return s.Email
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) GetId() *int64 {
	return s.Id
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) GetName() *string {
	return s.Name
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) GetUsername() *string {
	return s.Username
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetAvatar(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Avatar = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetEmail(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Email = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetId(v int64) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Id = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetName(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Name = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) SetUsername(v string) *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Username = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultMergeRequestSettingDefaultAssignees) Validate() error {
	return dara.Validate(s)
}

type ListProtectedBranchesResponseBodyResultTestSettingDTO struct {
	CheckConfig             *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig             `json:"checkConfig,omitempty" xml:"checkConfig,omitempty" type:"Struct"`
	CheckTaskQualityConfig  *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig  `json:"checkTaskQualityConfig,omitempty" xml:"checkTaskQualityConfig,omitempty" type:"Struct"`
	CodeGuidelinesDetection *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection `json:"codeGuidelinesDetection,omitempty" xml:"codeGuidelinesDetection,omitempty" type:"Struct"`
	// example:
	//
	// false
	IsRequired             *bool                                                                        `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	SensitiveInfoDetection *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection `json:"sensitiveInfoDetection,omitempty" xml:"sensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTO) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTO) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) GetCheckConfig() *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig {
	return s.CheckConfig
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) GetCheckTaskQualityConfig() *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	return s.CheckTaskQualityConfig
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) GetCodeGuidelinesDetection() *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	return s.CodeGuidelinesDetection
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) GetSensitiveInfoDetection() *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	return s.SensitiveInfoDetection
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetCheckConfig(v *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CheckConfig = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetCheckTaskQualityConfig(v *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CheckTaskQualityConfig = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetCodeGuidelinesDetection(v *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.CodeGuidelinesDetection = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetIsRequired(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.IsRequired = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) SetSensitiveInfoDetection(v *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) *ListProtectedBranchesResponseBodyResultTestSettingDTO {
	s.SensitiveInfoDetection = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTO) Validate() error {
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

type ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig struct {
	CheckItems []*ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems `json:"checkItems,omitempty" xml:"checkItems,omitempty" type:"Repeated"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) GetCheckItems() []*ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	return s.CheckItems
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) SetCheckItems(v []*ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig {
	s.CheckItems = v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfig) Validate() error {
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

type ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems struct {
	// example:
	//
	// false
	IsRequired *bool   `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) GetName() *string {
	return s.Name
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetIsRequired(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.IsRequired = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) SetName(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems {
	s.Name = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckConfigCheckItems) Validate() error {
	return dara.Validate(s)
}

type ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig struct {
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

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetBizNo() *string {
	return s.BizNo
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetMessage() *string {
	return s.Message
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) GetTaskName() *string {
	return s.TaskName
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetBizNo(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.BizNo = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetEnabled(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Enabled = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetMessage(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.Message = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) SetTaskName(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig {
	s.TaskName = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCheckTaskQualityConfig) Validate() error {
	return dara.Validate(s)
}

type ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_guide_lines
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) GetMessage() *string {
	return s.Message
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetEnabled(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) SetMessage(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection {
	s.Message = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOCodeGuidelinesDetection) Validate() error {
	return dara.Validate(s)
}

type ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection struct {
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// test_code_sensitive_info
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) GetMessage() *string {
	return s.Message
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetEnabled(v bool) *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) SetMessage(v string) *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection {
	s.Message = &v
	return s
}

func (s *ListProtectedBranchesResponseBodyResultTestSettingDTOSensitiveInfoDetection) Validate() error {
	return dara.Validate(s)
}
