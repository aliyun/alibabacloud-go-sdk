// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudPermissionTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudPermissionTemplates(v []*ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) *ListCloudPermissionTemplatesResponseBody
	GetCloudPermissionTemplates() []*ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates
	SetMaxResults(v int64) *ListCloudPermissionTemplatesResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListCloudPermissionTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCloudPermissionTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCloudPermissionTemplatesResponseBody
	GetTotalCount() *int64
}

type ListCloudPermissionTemplatesResponseBody struct {
	// The list of cloud permission templates.
	CloudPermissionTemplates []*ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates `json:"CloudPermissionTemplates,omitempty" xml:"CloudPermissionTemplates,omitempty" type:"Repeated"`
	// The maximum number of items per page for a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by the current call for the next query.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudPermissionTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudPermissionTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudPermissionTemplatesResponseBody) GetCloudPermissionTemplates() []*ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	return s.CloudPermissionTemplates
}

func (s *ListCloudPermissionTemplatesResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCloudPermissionTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudPermissionTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudPermissionTemplatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCloudPermissionTemplatesResponseBody) SetCloudPermissionTemplates(v []*ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) *ListCloudPermissionTemplatesResponseBody {
	s.CloudPermissionTemplates = v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBody) SetMaxResults(v int64) *ListCloudPermissionTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBody) SetNextToken(v string) *ListCloudPermissionTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBody) SetRequestId(v string) *ListCloudPermissionTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBody) SetTotalCount(v int64) *ListCloudPermissionTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBody) Validate() error {
	if s.CloudPermissionTemplates != nil {
		for _, item := range s.CloudPermissionTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates struct {
	// Specifies whether users are allowed to apply for the permission template.
	AllowRequest *bool `json:"AllowRequest,omitempty" xml:"AllowRequest,omitempty"`
	// The applicable scope of the permission template.
	ApplicableScope *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope `json:"ApplicableScope,omitempty" xml:"ApplicableScope,omitempty" type:"Struct"`
	// The ID of the cloud account to which the template belongs.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The creator information of the permission template.
	CloudPermissionTemplateCreator *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator `json:"CloudPermissionTemplateCreator,omitempty" xml:"CloudPermissionTemplateCreator,omitempty" type:"Struct"`
	// The unique identifier of the permission template.
	//
	// example:
	//
	// cpt_01l8dagf2iablhigj75l1n99xxxxx
	CloudPermissionTemplateId *string `json:"CloudPermissionTemplateId,omitempty" xml:"CloudPermissionTemplateId,omitempty"`
	// The initialization result. This parameter is returned only when the status is init_failed.
	CloudPermissionTemplateInitResult *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult `json:"CloudPermissionTemplateInitResult,omitempty" xml:"CloudPermissionTemplateInitResult,omitempty" type:"Struct"`
	// The name of the permission template.
	//
	// example:
	//
	// AdminAccessTemplate
	CloudPermissionTemplateName *string `json:"CloudPermissionTemplateName,omitempty" xml:"CloudPermissionTemplateName,omitempty"`
	// The ID of the application to which the exclusive template belongs. This parameter is empty for shared templates.
	//
	// example:
	//
	// user_01kexamplexxxx
	CloudPermissionTemplateOwnerId *string `json:"CloudPermissionTemplateOwnerId,omitempty" xml:"CloudPermissionTemplateOwnerId,omitempty"`
	// The owner type of the exclusive template. Valid values: application (M2M applications) and user (human users). This parameter is empty for shared templates.
	//
	// example:
	//
	// user
	CloudPermissionTemplateOwnerType *string `json:"CloudPermissionTemplateOwnerType,omitempty" xml:"CloudPermissionTemplateOwnerType,omitempty"`
	// The sharing mode of the permission template. Valid values: shared and exclusive.
	//
	// example:
	//
	// shared
	CloudPermissionTemplateSharingMode *string `json:"CloudPermissionTemplateSharingMode,omitempty" xml:"CloudPermissionTemplateSharingMode,omitempty"`
	// The usage type of the permission template. Valid values: user (privilege escalation for human users) and application (privilege escalation for M2M applications).
	//
	// example:
	//
	// user
	CloudPermissionTemplateUsageType *string `json:"CloudPermissionTemplateUsageType,omitempty" xml:"CloudPermissionTemplateUsageType,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the permission template.
	//
	// example:
	//
	// Administrator permission template for managing cloud accounts
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The external unique identifier of the associated cloud role, which is the role ARN on the cloud provider side.
	//
	// example:
	//
	// ReadOnlyAccess
	PermissionEntityExternalId *string `json:"PermissionEntityExternalId,omitempty" xml:"PermissionEntityExternalId,omitempty"`
	// The ID of the associated CloudAccountRole.
	//
	// example:
	//
	// carole_01kmek49aq4r8gxxxx
	PermissionEntityId *string `json:"PermissionEntityId,omitempty" xml:"PermissionEntityId,omitempty"`
	// The type of the permission entity. Fixed value: cloud_account_role.
	//
	// example:
	//
	// cloud_account_role
	PermissionEntityType *string `json:"PermissionEntityType,omitempty" xml:"PermissionEntityType,omitempty"`
	// The status of the permission template. Valid values: initializing, init_failed, enabled, and disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830227000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GoString() string {
	return s.String()
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetAllowRequest() *bool {
	return s.AllowRequest
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetApplicableScope() *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope {
	return s.ApplicableScope
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCloudPermissionTemplateCreator() *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator {
	return s.CloudPermissionTemplateCreator
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCloudPermissionTemplateId() *string {
	return s.CloudPermissionTemplateId
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCloudPermissionTemplateInitResult() *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult {
	return s.CloudPermissionTemplateInitResult
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCloudPermissionTemplateName() *string {
	return s.CloudPermissionTemplateName
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCloudPermissionTemplateOwnerId() *string {
	return s.CloudPermissionTemplateOwnerId
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCloudPermissionTemplateOwnerType() *string {
	return s.CloudPermissionTemplateOwnerType
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCloudPermissionTemplateSharingMode() *string {
	return s.CloudPermissionTemplateSharingMode
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCloudPermissionTemplateUsageType() *string {
	return s.CloudPermissionTemplateUsageType
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetPermissionEntityExternalId() *string {
	return s.PermissionEntityExternalId
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetPermissionEntityId() *string {
	return s.PermissionEntityId
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetPermissionEntityType() *string {
	return s.PermissionEntityType
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetStatus() *string {
	return s.Status
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetAllowRequest(v bool) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.AllowRequest = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetApplicableScope(v *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.ApplicableScope = v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCloudAccountId(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CloudAccountId = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCloudPermissionTemplateCreator(v *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CloudPermissionTemplateCreator = v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCloudPermissionTemplateId(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CloudPermissionTemplateId = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCloudPermissionTemplateInitResult(v *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CloudPermissionTemplateInitResult = v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCloudPermissionTemplateName(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CloudPermissionTemplateName = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCloudPermissionTemplateOwnerId(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CloudPermissionTemplateOwnerId = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCloudPermissionTemplateOwnerType(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CloudPermissionTemplateOwnerType = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCloudPermissionTemplateSharingMode(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CloudPermissionTemplateSharingMode = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCloudPermissionTemplateUsageType(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CloudPermissionTemplateUsageType = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetCreateTime(v int64) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetDescription(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.Description = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetInstanceId(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.InstanceId = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetPermissionEntityExternalId(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.PermissionEntityExternalId = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetPermissionEntityId(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.PermissionEntityId = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetPermissionEntityType(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.PermissionEntityType = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetStatus(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.Status = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) SetUpdateTime(v int64) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplates) Validate() error {
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

type ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope struct {
	// The list of specified group IDs. This parameter has a value only when ScopeType is set to custom.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The list of specified organizational unit IDs. This parameter has a value only when ScopeType is set to custom.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The type of the applicable scope. Valid values: global and custom.
	//
	// example:
	//
	// custom
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The list of specified user IDs. This parameter has a value only when ScopeType is set to custom.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) String() string {
	return dara.Prettify(s)
}

func (s ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) GoString() string {
	return s.String()
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) SetGroupIds(v []*string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope {
	s.GroupIds = v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) SetOrganizationalUnitIds(v []*string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope {
	s.OrganizationalUnitIds = v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) SetScopeType(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope {
	s.ScopeType = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) SetUserIds(v []*string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope {
	s.UserIds = v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesApplicableScope) Validate() error {
	return dara.Validate(s)
}

type ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator struct {
	// The identity identifier of the creator.
	//
	// example:
	//
	// user_01kmek49aq4r8gxxxx
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// Zhang San
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The type of the creator. Valid values: idaas_admin and idaas_user.
	//
	// example:
	//
	// idaas_user
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
}

func (s ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) String() string {
	return dara.Prettify(s)
}

func (s ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) GoString() string {
	return s.String()
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) GetCreatorType() *string {
	return s.CreatorType
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) SetCreatorId(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator {
	s.CreatorId = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) SetCreatorName(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator {
	s.CreatorName = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) SetCreatorType(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator {
	s.CreatorType = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateCreator) Validate() error {
	return dara.Validate(s)
}

type ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult struct {
	// The error code returned when initialization fails.
	//
	// example:
	//
	// CloudPermissionTemplate.AsyncInitError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when initialization fails.
	//
	// example:
	//
	// Permission template initialization failed. Try again later
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The initialization result. Valid values: success and failed.
	//
	// example:
	//
	// failed
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) String() string {
	return dara.Prettify(s)
}

func (s ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) GoString() string {
	return s.String()
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) GetResult() *string {
	return s.Result
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) SetErrorCode(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult {
	s.ErrorCode = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) SetErrorMessage(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult {
	s.ErrorMessage = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) SetResult(v string) *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult {
	s.Result = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponseBodyCloudPermissionTemplatesCloudPermissionTemplateInitResult) Validate() error {
	return dara.Validate(s)
}
