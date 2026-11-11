// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudPermissionTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudPermissionTemplate(v *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) *GetCloudPermissionTemplateResponseBody
	GetCloudPermissionTemplate() *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate
	SetRequestId(v string) *GetCloudPermissionTemplateResponseBody
	GetRequestId() *string
}

type GetCloudPermissionTemplateResponseBody struct {
	// The permission template.
	CloudPermissionTemplate *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate `json:"CloudPermissionTemplate,omitempty" xml:"CloudPermissionTemplate,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudPermissionTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudPermissionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudPermissionTemplateResponseBody) GetCloudPermissionTemplate() *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	return s.CloudPermissionTemplate
}

func (s *GetCloudPermissionTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudPermissionTemplateResponseBody) SetCloudPermissionTemplate(v *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) *GetCloudPermissionTemplateResponseBody {
	s.CloudPermissionTemplate = v
	return s
}

func (s *GetCloudPermissionTemplateResponseBody) SetRequestId(v string) *GetCloudPermissionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBody) Validate() error {
	if s.CloudPermissionTemplate != nil {
		if err := s.CloudPermissionTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate struct {
	// Specifies whether users are allowed to apply for the permission template.
	//
	// example:
	//
	// true
	AllowRequest *bool `json:"AllowRequest,omitempty" xml:"AllowRequest,omitempty"`
	// The applicable scope for applying for the permission template.
	ApplicableScope *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope `json:"ApplicableScope,omitempty" xml:"ApplicableScope,omitempty" type:"Struct"`
	// The ID of the cloud account to which the permission template belongs.
	//
	// example:
	//
	// ca_01kq1us6frlbo3rr8m216vj3bkak2
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The information about the creator of the permission template.
	CloudPermissionTemplateCreator *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator `json:"CloudPermissionTemplateCreator,omitempty" xml:"CloudPermissionTemplateCreator,omitempty" type:"Struct"`
	// The unique identifier of the permission template.
	//
	// example:
	//
	// cpt_01l8dagf2iablhigj75l1n99xxxxx
	CloudPermissionTemplateId *string `json:"CloudPermissionTemplateId,omitempty" xml:"CloudPermissionTemplateId,omitempty"`
	// The initialization result. This field is returned only when the status is init_failed.
	CloudPermissionTemplateInitResult *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult `json:"CloudPermissionTemplateInitResult,omitempty" xml:"CloudPermissionTemplateInitResult,omitempty" type:"Struct"`
	// The name of the permission template.
	//
	// example:
	//
	// testName
	CloudPermissionTemplateName *string `json:"CloudPermissionTemplateName,omitempty" xml:"CloudPermissionTemplateName,omitempty"`
	// The ID of the application that owns the exclusive template. This field is empty for shared templates.
	//
	// example:
	//
	// user_01kexamplexxxx
	CloudPermissionTemplateOwnerId *string `json:"CloudPermissionTemplateOwnerId,omitempty" xml:"CloudPermissionTemplateOwnerId,omitempty"`
	// The owner type of the exclusive template. Valid values:
	//
	// - application: M2M application.
	//
	// - user: Human user.
	//
	// This field is empty for shared templates.
	//
	// example:
	//
	// user
	CloudPermissionTemplateOwnerType *string `json:"CloudPermissionTemplateOwnerType,omitempty" xml:"CloudPermissionTemplateOwnerType,omitempty"`
	// The sharing mode of the permission template. Valid values:
	//
	// - shared: Shared.
	//
	// - exclusive: Exclusive.
	//
	// example:
	//
	// shared
	CloudPermissionTemplateSharingMode *string `json:"CloudPermissionTemplateSharingMode,omitempty" xml:"CloudPermissionTemplateSharingMode,omitempty"`
	// The usage type of the permission template. Valid values:
	//
	// - user: Privilege escalation for human users.
	//
	// - application: Privilege escalation for machine-to-machine (M2M) applications.
	//
	// example:
	//
	// user
	CloudPermissionTemplateUsageType *string `json:"CloudPermissionTemplateUsageType,omitempty" xml:"CloudPermissionTemplateUsageType,omitempty"`
	// The time when the permission template was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1776239357000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the permission template.
	//
	// example:
	//
	// this is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ki6hd7ihir4ybawogqk6kqsfjm
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The external unique identifier of the associated cloud role, which is the ARN of the role on the cloud provider side.
	//
	// example:
	//
	// acs:ram::123456:role/idaas-test
	PermissionEntityExternalId *string `json:"PermissionEntityExternalId,omitempty" xml:"PermissionEntityExternalId,omitempty"`
	// The ID of the associated CloudAccountRole.
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	PermissionEntityId *string `json:"PermissionEntityId,omitempty" xml:"PermissionEntityId,omitempty"`
	// The type of the permission entity. Fixed value: cloud_account_role.
	//
	// example:
	//
	// cloud_account_role
	PermissionEntityType *string `json:"PermissionEntityType,omitempty" xml:"PermissionEntityType,omitempty"`
	// The status of the permission template. Valid values:
	//
	// - initializing: The template is being initialized.
	//
	// - init_failed: The template initialization failed.
	//
	// - enabled: The template is enabled.
	//
	// - disabled: The template is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the permission template was last updated. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1776239357000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GoString() string {
	return s.String()
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetAllowRequest() *bool {
	return s.AllowRequest
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetApplicableScope() *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope {
	return s.ApplicableScope
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCloudPermissionTemplateCreator() *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator {
	return s.CloudPermissionTemplateCreator
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCloudPermissionTemplateId() *string {
	return s.CloudPermissionTemplateId
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCloudPermissionTemplateInitResult() *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult {
	return s.CloudPermissionTemplateInitResult
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCloudPermissionTemplateName() *string {
	return s.CloudPermissionTemplateName
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCloudPermissionTemplateOwnerId() *string {
	return s.CloudPermissionTemplateOwnerId
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCloudPermissionTemplateOwnerType() *string {
	return s.CloudPermissionTemplateOwnerType
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCloudPermissionTemplateSharingMode() *string {
	return s.CloudPermissionTemplateSharingMode
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCloudPermissionTemplateUsageType() *string {
	return s.CloudPermissionTemplateUsageType
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetDescription() *string {
	return s.Description
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetPermissionEntityExternalId() *string {
	return s.PermissionEntityExternalId
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetPermissionEntityId() *string {
	return s.PermissionEntityId
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetPermissionEntityType() *string {
	return s.PermissionEntityType
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetStatus() *string {
	return s.Status
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetAllowRequest(v bool) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.AllowRequest = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetApplicableScope(v *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.ApplicableScope = v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCloudAccountId(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CloudAccountId = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCloudPermissionTemplateCreator(v *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CloudPermissionTemplateCreator = v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCloudPermissionTemplateId(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CloudPermissionTemplateId = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCloudPermissionTemplateInitResult(v *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CloudPermissionTemplateInitResult = v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCloudPermissionTemplateName(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CloudPermissionTemplateName = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCloudPermissionTemplateOwnerId(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CloudPermissionTemplateOwnerId = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCloudPermissionTemplateOwnerType(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CloudPermissionTemplateOwnerType = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCloudPermissionTemplateSharingMode(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CloudPermissionTemplateSharingMode = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCloudPermissionTemplateUsageType(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CloudPermissionTemplateUsageType = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetCreateTime(v int64) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.CreateTime = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetDescription(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.Description = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetInstanceId(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.InstanceId = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetPermissionEntityExternalId(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.PermissionEntityExternalId = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetPermissionEntityId(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.PermissionEntityId = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetPermissionEntityType(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.PermissionEntityType = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetStatus(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.Status = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) SetUpdateTime(v int64) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate {
	s.UpdateTime = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplate) Validate() error {
	if s.ApplicableScope != nil {
		if err := s.ApplicableScope.Validate(); err != nil {
			return err
		}
	}
	if s.CloudPermissionTemplateCreator != nil {
		if err := s.CloudPermissionTemplateCreator.Validate(); err != nil {
			return err
		}
	}
	if s.CloudPermissionTemplateInitResult != nil {
		if err := s.CloudPermissionTemplateInitResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope struct {
	// The list of specified group IDs. This field is returned only when ScopeType is set to custom.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The list of specified organizational unit IDs. This field is returned only when ScopeType is set to custom.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The type of the applicable scope. Valid values:
	//
	// - global: Global.
	//
	// - custom: Custom.
	//
	// example:
	//
	// global
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The list of specified user IDs. This field is returned only when ScopeType is set to custom.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) String() string {
	return dara.Prettify(s)
}

func (s GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) GoString() string {
	return s.String()
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) GetScopeType() *string {
	return s.ScopeType
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) SetGroupIds(v []*string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope {
	s.GroupIds = v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) SetOrganizationalUnitIds(v []*string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope {
	s.OrganizationalUnitIds = v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) SetScopeType(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope {
	s.ScopeType = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) SetUserIds(v []*string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope {
	s.UserIds = v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateApplicableScope) Validate() error {
	return dara.Validate(s)
}

type GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator struct {
	// The identity identifier of the creator.
	//
	// example:
	//
	// 8645321
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// root
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The type of the creator. Valid values:
	//
	// - idaas_admin: IDaaS administrator.
	//
	// - idaas_user: IDaaS user.
	//
	// example:
	//
	// idaas_user
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
}

func (s GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) String() string {
	return dara.Prettify(s)
}

func (s GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) GoString() string {
	return s.String()
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) GetCreatorType() *string {
	return s.CreatorType
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) SetCreatorId(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator {
	s.CreatorId = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) SetCreatorName(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator {
	s.CreatorName = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) SetCreatorType(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator {
	s.CreatorType = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateCreator) Validate() error {
	return dara.Validate(s)
}

type GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult struct {
	// The error code returned when the initialization fails.
	//
	// example:
	//
	// CloudPermissionTemplate.AsyncInitError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the initialization fails.
	//
	// example:
	//
	// server error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The initialization result. Valid values:
	//
	// - success: The initialization succeeded.
	//
	// - failed: The initialization failed.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) String() string {
	return dara.Prettify(s)
}

func (s GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) GoString() string {
	return s.String()
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) GetResult() *string {
	return s.Result
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) SetErrorCode(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult {
	s.ErrorCode = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) SetErrorMessage(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) SetResult(v string) *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult {
	s.Result = &v
	return s
}

func (s *GetCloudPermissionTemplateResponseBodyCloudPermissionTemplateCloudPermissionTemplateInitResult) Validate() error {
	return dara.Validate(s)
}
