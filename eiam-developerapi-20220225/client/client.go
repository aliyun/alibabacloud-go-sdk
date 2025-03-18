// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddUserToOrganizationalUnitsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddUserToOrganizationalUnitsHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddUserToOrganizationalUnitsHeaders) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsHeaders) SetCommonHeaders(v map[string]*string) *AddUserToOrganizationalUnitsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddUserToOrganizationalUnitsHeaders) SetAuthorization(v string) *AddUserToOrganizationalUnitsHeaders {
	s.Authorization = &v
	return s
}

type AddUserToOrganizationalUnitsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [ou_wovwffm62xifdziem7an7xxxxx]
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" xml:"organizationalUnitIds,omitempty" type:"Repeated"`
}

func (s AddUserToOrganizationalUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *AddUserToOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

type AddUserToOrganizationalUnitsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AddUserToOrganizationalUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *AddUserToOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *AddUserToOrganizationalUnitsResponse) SetStatusCode(v int32) *AddUserToOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

type AddUsersToGroupHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddUsersToGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupHeaders) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupHeaders) SetCommonHeaders(v map[string]*string) *AddUsersToGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddUsersToGroupHeaders) SetAuthorization(v string) *AddUsersToGroupHeaders {
	s.Authorization = &v
	return s
}

type AddUsersToGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [user_d6sbsuumeta4h66ec3il7yxxxx}
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddUsersToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupRequest) SetUserIds(v []*string) *AddUsersToGroupRequest {
	s.UserIds = v
	return s
}

type AddUsersToGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AddUsersToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponse) SetHeaders(v map[string]*string) *AddUsersToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUsersToGroupResponse) SetStatusCode(v int32) *AddUsersToGroupResponse {
	s.StatusCode = &v
	return s
}

type CreateGroupHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupHeaders) SetAuthorization(v string) *CreateGroupHeaders {
	s.Authorization = &v
	return s
}

type CreateGroupRequest struct {
	// example:
	//
	// group_2bo6lefcewdausyyxxxx
	GroupExternalId *string `json:"groupExternalId,omitempty" xml:"groupExternalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name001
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetGroupExternalId(v string) *CreateGroupRequest {
	s.GroupExternalId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

type CreateGroupResponseBody struct {
	// example:
	//
	// group_wovwffm62xifdziem7an7xxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetGroupId(v string) *CreateGroupResponseBody {
	s.GroupId = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateOrganizationalUnitHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *CreateOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrganizationalUnitHeaders) SetAuthorization(v string) *CreateOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

type CreateOrganizationalUnitRequest struct {
	// The description of the organizational unit.
	//
	// example:
	//
	// test organizational unit
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The external ID of the organizational unit. The external ID can be used to map external data to the data of the organizational unit in Employee Identity and Access Management (EIAM) of Identity as a Service (IDaaS). By default, the external ID is the organizational unit ID.
	//
	// For organizational units with the same source type and source ID, each organizational unit has a unique external ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	// The name of the organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	// The ID of the parent organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s CreateOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitRequest) SetDescription(v string) *CreateOrganizationalUnitRequest {
	s.Description = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetOrganizationalUnitExternalId(v string) *CreateOrganizationalUnitRequest {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetOrganizationalUnitName(v string) *CreateOrganizationalUnitRequest {
	s.OrganizationalUnitName = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetParentId(v string) *CreateOrganizationalUnitRequest {
	s.ParentId = &v
	return s
}

type CreateOrganizationalUnitResponseBody struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
}

func (s CreateOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitResponseBody) SetOrganizationalUnitId(v string) *CreateOrganizationalUnitResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

type CreateOrganizationalUnitResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitResponse) SetHeaders(v map[string]*string) *CreateOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *CreateOrganizationalUnitResponse) SetStatusCode(v int32) *CreateOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrganizationalUnitResponse) SetBody(v *CreateOrganizationalUnitResponseBody) *CreateOrganizationalUnitResponse {
	s.Body = v
	return s
}

type CreateUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUserHeaders) GoString() string {
	return s.String()
}

func (s *CreateUserHeaders) SetCommonHeaders(v map[string]*string) *CreateUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUserHeaders) SetAuthorization(v string) *CreateUserHeaders {
	s.Authorization = &v
	return s
}

type CreateUserRequest struct {
	// Custom fields
	CustomFields []*CreateUserRequestCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// The description of the account. The description can be up to 256 characters in length.
	//
	// example:
	//
	// test user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the account. The display name can be up to 64 characters in length.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address of the user who owns the account.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Indicates whether the email address is verified. This field is required if an email address is specified. If you have no special requirement, set this parameter to true.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	// The password of the account. For information about the password rules, go to the Create User panel in the Identity as a Service (IDaaS) console.
	//
	// example:
	//
	// xxxxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// Configure the initial password
	PasswordInitializationConfig *CreateUserRequestPasswordInitializationConfig `json:"passwordInitializationConfig,omitempty" xml:"passwordInitializationConfig,omitempty" type:"Struct"`
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// Indicates whether the mobile number is verified. This field is required if a mobile number is specified. If you have no special requirement, set this parameter to true.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +. This parameter is required if a mobile number is specified.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	// The ID of the primary organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	PrimaryOrganizationalUnitId *string `json:"primaryOrganizationalUnitId,omitempty" xml:"primaryOrganizationalUnitId,omitempty"`
	// The external ID of the account. The external ID can be used to map external data to the data of the account in EIAM of Identity as a Service (IDaaS). By default, the external ID is the account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	// The username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// name001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetCustomFields(v []*CreateUserRequestCustomFields) *CreateUserRequest {
	s.CustomFields = v
	return s
}

func (s *CreateUserRequest) SetDescription(v string) *CreateUserRequest {
	s.Description = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetEmailVerified(v bool) *CreateUserRequest {
	s.EmailVerified = &v
	return s
}

func (s *CreateUserRequest) SetPassword(v string) *CreateUserRequest {
	s.Password = &v
	return s
}

func (s *CreateUserRequest) SetPasswordInitializationConfig(v *CreateUserRequestPasswordInitializationConfig) *CreateUserRequest {
	s.PasswordInitializationConfig = v
	return s
}

func (s *CreateUserRequest) SetPhoneNumber(v string) *CreateUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateUserRequest) SetPhoneNumberVerified(v bool) *CreateUserRequest {
	s.PhoneNumberVerified = &v
	return s
}

func (s *CreateUserRequest) SetPhoneRegion(v string) *CreateUserRequest {
	s.PhoneRegion = &v
	return s
}

func (s *CreateUserRequest) SetPrimaryOrganizationalUnitId(v string) *CreateUserRequest {
	s.PrimaryOrganizationalUnitId = &v
	return s
}

func (s *CreateUserRequest) SetUserExternalId(v string) *CreateUserRequest {
	s.UserExternalId = &v
	return s
}

func (s *CreateUserRequest) SetUsername(v string) *CreateUserRequest {
	s.Username = &v
	return s
}

type CreateUserRequestCustomFields struct {
	// Field name
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Filed value
	//
	// example:
	//
	// fieldValue_001
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s CreateUserRequestCustomFields) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequestCustomFields) GoString() string {
	return s.String()
}

func (s *CreateUserRequestCustomFields) SetFieldName(v string) *CreateUserRequestCustomFields {
	s.FieldName = &v
	return s
}

func (s *CreateUserRequestCustomFields) SetFieldValue(v string) *CreateUserRequestCustomFields {
	s.FieldValue = &v
	return s
}

type CreateUserRequestPasswordInitializationConfig struct {
	// Password  forced update
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"passwordForcedUpdateStatus,omitempty" xml:"passwordForcedUpdateStatus,omitempty"`
	// Password policy
	//
	// example:
	//
	// global
	PasswordInitializationPolicyPriority *string `json:"passwordInitializationPolicyPriority,omitempty" xml:"passwordInitializationPolicyPriority,omitempty"`
	// Password Initialization Type
	//
	// example:
	//
	// random
	PasswordInitializationType *string `json:"passwordInitializationType,omitempty" xml:"passwordInitializationType,omitempty"`
	// User Notification Channels
	//
	// example:
	//
	// sms
	UserNotificationChannels []*string `json:"userNotificationChannels,omitempty" xml:"userNotificationChannels,omitempty" type:"Repeated"`
}

func (s CreateUserRequestPasswordInitializationConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequestPasswordInitializationConfig) GoString() string {
	return s.String()
}

func (s *CreateUserRequestPasswordInitializationConfig) SetPasswordForcedUpdateStatus(v string) *CreateUserRequestPasswordInitializationConfig {
	s.PasswordForcedUpdateStatus = &v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) SetPasswordInitializationPolicyPriority(v string) *CreateUserRequestPasswordInitializationConfig {
	s.PasswordInitializationPolicyPriority = &v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) SetPasswordInitializationType(v string) *CreateUserRequestPasswordInitializationConfig {
	s.PasswordInitializationType = &v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) SetUserNotificationChannels(v []*string) *CreateUserRequestPasswordInitializationConfig {
	s.UserNotificationChannels = v
	return s
}

type CreateUserResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetUserId(v string) *CreateUserResponseBody {
	s.UserId = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type DeleteGroupHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupHeaders) GoString() string {
	return s.String()
}

func (s *DeleteGroupHeaders) SetCommonHeaders(v map[string]*string) *DeleteGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteGroupHeaders) SetAuthorization(v string) *DeleteGroupHeaders {
	s.Authorization = &v
	return s
}

type DeleteGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetStatusCode(v int32) *DeleteGroupResponse {
	s.StatusCode = &v
	return s
}

type DeleteOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteOrganizationalUnitHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *DeleteOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteOrganizationalUnitHeaders) SetAuthorization(v string) *DeleteOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

type DeleteOrganizationalUnitResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitResponse) SetHeaders(v map[string]*string) *DeleteOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrganizationalUnitResponse) SetStatusCode(v int32) *DeleteOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

type DeleteUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUserHeaders) SetCommonHeaders(v map[string]*string) *DeleteUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUserHeaders) SetAuthorization(v string) *DeleteUserHeaders {
	s.Authorization = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

type DisableUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DisableUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s DisableUserHeaders) GoString() string {
	return s.String()
}

func (s *DisableUserHeaders) SetCommonHeaders(v map[string]*string) *DisableUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DisableUserHeaders) SetAuthorization(v string) *DisableUserHeaders {
	s.Authorization = &v
	return s
}

type DisableUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DisableUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableUserResponse) GoString() string {
	return s.String()
}

func (s *DisableUserResponse) SetHeaders(v map[string]*string) *DisableUserResponse {
	s.Headers = v
	return s
}

func (s *DisableUserResponse) SetStatusCode(v int32) *DisableUserResponse {
	s.StatusCode = &v
	return s
}

type EnableUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s EnableUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s EnableUserHeaders) GoString() string {
	return s.String()
}

func (s *EnableUserHeaders) SetCommonHeaders(v map[string]*string) *EnableUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EnableUserHeaders) SetAuthorization(v string) *EnableUserHeaders {
	s.Authorization = &v
	return s
}

type EnableUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s EnableUserResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableUserResponse) GoString() string {
	return s.String()
}

func (s *EnableUserResponse) SetHeaders(v map[string]*string) *EnableUserResponse {
	s.Headers = v
	return s
}

func (s *EnableUserResponse) SetStatusCode(v int32) *EnableUserResponse {
	s.StatusCode = &v
	return s
}

type GenerateDeviceCodeRequest struct {
	// The authorization scope.
	//
	// example:
	//
	// xxx
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s GenerateDeviceCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateDeviceCodeRequest) SetScope(v string) *GenerateDeviceCodeRequest {
	s.Scope = &v
	return s
}

type GenerateDeviceCodeResponseBody struct {
	// The device code.
	//
	// example:
	//
	// xxxxx
	DeviceCode *string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	// The time when the token expires. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1653288641
	ExpiresAt *int64 `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	// The remaining validity period of the device code. Unit: seconds.
	//
	// example:
	//
	// 1200
	ExpiresIn *int64 `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// The timeout period of the request token. Unit: seconds.
	//
	// example:
	//
	// 5
	Interval *int64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// The user authorization code.
	//
	// example:
	//
	// xxxxx
	UserCode *string `json:"user_code,omitempty" xml:"user_code,omitempty"`
	// The verification URI.
	//
	// example:
	//
	// https://example.com/authorize/device
	VerificationUri *string `json:"verification_uri,omitempty" xml:"verification_uri,omitempty"`
	// The complete verification URI.
	//
	// example:
	//
	// https://example.com/authorize/device?user_code=
	//
	// xxxx
	VerificationUriComplete *string `json:"verification_uri_complete,omitempty" xml:"verification_uri_complete,omitempty"`
}

func (s GenerateDeviceCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDeviceCodeResponseBody) SetDeviceCode(v string) *GenerateDeviceCodeResponseBody {
	s.DeviceCode = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetExpiresAt(v int64) *GenerateDeviceCodeResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetExpiresIn(v int64) *GenerateDeviceCodeResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetInterval(v int64) *GenerateDeviceCodeResponseBody {
	s.Interval = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetUserCode(v string) *GenerateDeviceCodeResponseBody {
	s.UserCode = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetVerificationUri(v string) *GenerateDeviceCodeResponseBody {
	s.VerificationUri = &v
	return s
}

func (s *GenerateDeviceCodeResponseBody) SetVerificationUriComplete(v string) *GenerateDeviceCodeResponseBody {
	s.VerificationUriComplete = &v
	return s
}

type GenerateDeviceCodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDeviceCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDeviceCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateDeviceCodeResponse) SetHeaders(v map[string]*string) *GenerateDeviceCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateDeviceCodeResponse) SetStatusCode(v int32) *GenerateDeviceCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDeviceCodeResponse) SetBody(v *GenerateDeviceCodeResponseBody) *GenerateDeviceCodeResponse {
	s.Body = v
	return s
}

type GenerateTokenRequest struct {
	// The client ID.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// The client secret. This parameter is required if grant_type is set to client_credentials.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	// The authorization code. This parameter is required if grant_type is set to authorization_code.
	//
	// example:
	//
	// xxxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The verification code.
	//
	// example:
	//
	// xxx
	CodeVerifier *string `json:"code_verifier,omitempty" xml:"code_verifier,omitempty"`
	// The device code. This parameter is required if grant_type is set to authorization_code.urn:ietf:params:oauth:grant-type:device_code.
	//
	// example:
	//
	// xxxx
	DeviceCode *string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	// The excluded tags.
	//
	// example:
	//
	// ATxxx
	ExclusiveTag *string `json:"exclusive_tag,omitempty" xml:"exclusive_tag,omitempty"`
	// The authorization type. Valid values:
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// <!---->
	//
	// 	- authorization_code
	//
	// 	- urn:ietf:params:oauth:grant-type:device_code
	//
	// 	- refresh_token
	//
	// 	- client_credentials: You must specify the client_id and client_secret parameters.
	//
	// 	- password: This option is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// client_credentials
	GrantType *string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	// The username. This parameter is required if grant_type is set to password. The password authentication type is not supported.
	//
	// example:
	//
	// xxxxxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The redirect URI. This parameter is required if grant_type is set to authorization_code. The value of this parameter must be the same as the redirect URI in the request to obtain the authorization code.
	//
	// example:
	//
	// xxx
	RedirectUri *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	// The refreshed token. This parameter is required if grant_type is set to refresh_token.
	//
	// example:
	//
	// ATxxx
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// The authorization scope. Valid values:
	//
	// 	- openid
	//
	// 	- email
	//
	// 	- phone
	//
	// 	- profile
	//
	// example:
	//
	// xxxx
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The username. This parameter is required if grant_type is set to password. The password authentication type is not supported.
	//
	// example:
	//
	// uesrname_001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GenerateTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateTokenRequest) SetClientId(v string) *GenerateTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GenerateTokenRequest) SetClientSecret(v string) *GenerateTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *GenerateTokenRequest) SetCode(v string) *GenerateTokenRequest {
	s.Code = &v
	return s
}

func (s *GenerateTokenRequest) SetCodeVerifier(v string) *GenerateTokenRequest {
	s.CodeVerifier = &v
	return s
}

func (s *GenerateTokenRequest) SetDeviceCode(v string) *GenerateTokenRequest {
	s.DeviceCode = &v
	return s
}

func (s *GenerateTokenRequest) SetExclusiveTag(v string) *GenerateTokenRequest {
	s.ExclusiveTag = &v
	return s
}

func (s *GenerateTokenRequest) SetGrantType(v string) *GenerateTokenRequest {
	s.GrantType = &v
	return s
}

func (s *GenerateTokenRequest) SetPassword(v string) *GenerateTokenRequest {
	s.Password = &v
	return s
}

func (s *GenerateTokenRequest) SetRedirectUri(v string) *GenerateTokenRequest {
	s.RedirectUri = &v
	return s
}

func (s *GenerateTokenRequest) SetRefreshToken(v string) *GenerateTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *GenerateTokenRequest) SetScope(v string) *GenerateTokenRequest {
	s.Scope = &v
	return s
}

func (s *GenerateTokenRequest) SetUsername(v string) *GenerateTokenRequest {
	s.Username = &v
	return s
}

type GenerateTokenResponseBody struct {
	// The access token.
	//
	// example:
	//
	// ATxxx
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// The time when the token expires. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1653288641
	ExpiresAt *int64 `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	// The remaining validity period of the token. Unit: seconds.
	//
	// example:
	//
	// 1200
	ExpiresIn *int64 `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// The ID token.
	//
	// example:
	//
	// xxxxx
	IdToken *string `json:"id_token,omitempty" xml:"id_token,omitempty"`
	// The refresh token.
	//
	// example:
	//
	// RTxxx
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// The type of the token. Valid values: Basic Bearer
	//
	// example:
	//
	// Bearer
	TokenType *string `json:"token_type,omitempty" xml:"token_type,omitempty"`
}

func (s GenerateTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTokenResponseBody) SetAccessToken(v string) *GenerateTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GenerateTokenResponseBody) SetExpiresAt(v int64) *GenerateTokenResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *GenerateTokenResponseBody) SetExpiresIn(v int64) *GenerateTokenResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GenerateTokenResponseBody) SetIdToken(v string) *GenerateTokenResponseBody {
	s.IdToken = &v
	return s
}

func (s *GenerateTokenResponseBody) SetRefreshToken(v string) *GenerateTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GenerateTokenResponseBody) SetTokenType(v string) *GenerateTokenResponseBody {
	s.TokenType = &v
	return s
}

type GenerateTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateTokenResponse) SetHeaders(v map[string]*string) *GenerateTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateTokenResponse) SetStatusCode(v int32) *GenerateTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTokenResponse) SetBody(v *GenerateTokenResponseBody) *GenerateTokenResponse {
	s.Body = v
	return s
}

type GetApplicationProvisioningScopeHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetApplicationProvisioningScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeHeaders) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeHeaders) SetCommonHeaders(v map[string]*string) *GetApplicationProvisioningScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetApplicationProvisioningScopeHeaders) SetAuthorization(v string) *GetApplicationProvisioningScopeHeaders {
	s.Authorization = &v
	return s
}

type GetApplicationProvisioningScopeResponseBody struct {
	GroupIds []*string `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	// The IDs of organizational units.
	//
	// example:
	//
	// [ou_xxx001]
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" xml:"organizationalUnitIds,omitempty" type:"Repeated"`
}

func (s GetApplicationProvisioningScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponseBody) SetGroupIds(v []*string) *GetApplicationProvisioningScopeResponseBody {
	s.GroupIds = v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBody) SetOrganizationalUnitIds(v []*string) *GetApplicationProvisioningScopeResponseBody {
	s.OrganizationalUnitIds = v
	return s
}

type GetApplicationProvisioningScopeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationProvisioningScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationProvisioningScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponse) SetHeaders(v map[string]*string) *GetApplicationProvisioningScopeResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationProvisioningScopeResponse) SetStatusCode(v int32) *GetApplicationProvisioningScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationProvisioningScopeResponse) SetBody(v *GetApplicationProvisioningScopeResponseBody) *GetApplicationProvisioningScopeResponse {
	s.Body = v
	return s
}

type GetGroupHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGroupHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupHeaders) SetCommonHeaders(v map[string]*string) *GetGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupHeaders) SetAuthorization(v string) *GetGroupHeaders {
	s.Authorization = &v
	return s
}

type GetGroupResponseBody struct {
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// description_demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupExternalId *string `json:"groupExternalId,omitempty" xml:"groupExternalId,omitempty"`
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// name_test
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupSourceId *string `json:"groupSourceId,omitempty" xml:"groupSourceId,omitempty"`
	// example:
	//
	// build_in
	GroupSourceType *string `json:"groupSourceType,omitempty" xml:"groupSourceType,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) SetCreateTime(v int64) *GetGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetGroupResponseBody) SetDescription(v string) *GetGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupExternalId(v string) *GetGroupResponseBody {
	s.GroupExternalId = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupId(v string) *GetGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupName(v string) *GetGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupSourceId(v string) *GetGroupResponseBody {
	s.GroupSourceId = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupSourceType(v string) *GetGroupResponseBody {
	s.GroupSourceType = &v
	return s
}

func (s *GetGroupResponseBody) SetInstanceId(v string) *GetGroupResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetGroupResponseBody) SetUpdateTime(v int64) *GetGroupResponseBody {
	s.UpdateTime = &v
	return s
}

type GetGroupResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGroupResponse) SetHeaders(v map[string]*string) *GetGroupResponse {
	s.Headers = v
	return s
}

func (s *GetGroupResponse) SetStatusCode(v int32) *GetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupResponse) SetBody(v *GetGroupResponseBody) *GetGroupResponse {
	s.Body = v
	return s
}

type GetOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetOrganizationalUnitHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationalUnitHeaders) SetAuthorization(v string) *GetOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

type GetOrganizationalUnitResponseBody struct {
	// The time when the organizational unit was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652083425923
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the organizational unit.
	//
	// example:
	//
	// xxxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The external ID of the organizational unit.
	//
	// example:
	//
	// id_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	// The source ID of the organizational unit.
	//
	// example:
	//
	// id_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitSourceId *string `json:"organizationalUnitSourceId,omitempty" xml:"organizationalUnitSourceId,omitempty"`
	// The source type of the organizational unit. Valid values:
	//
	// 	- build_in: The organizational unit was created in Identity as a Service (IDaaS).
	//
	// 	- ding_talk: The organizational unit was imported from DingTalk.
	//
	// 	- ad: The organizational unit was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The organizational unit was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	OrganizationalUnitSourceType *string `json:"organizationalUnitSourceType,omitempty" xml:"organizationalUnitSourceType,omitempty"`
	// The ID of the parent organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The time when the organizational unit was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652083425923
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponseBody) SetCreateTime(v int64) *GetOrganizationalUnitResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetDescription(v string) *GetOrganizationalUnitResponseBody {
	s.Description = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetInstanceId(v string) *GetOrganizationalUnitResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitId(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitName(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetParentId(v string) *GetOrganizationalUnitResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetUpdateTime(v int64) *GetOrganizationalUnitResponseBody {
	s.UpdateTime = &v
	return s
}

type GetOrganizationalUnitResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponse) SetHeaders(v map[string]*string) *GetOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationalUnitResponse) SetStatusCode(v int32) *GetOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationalUnitResponse) SetBody(v *GetOrganizationalUnitResponseBody) *GetOrganizationalUnitResponse {
	s.Body = v
	return s
}

type GetOrganizationalUnitIdByExternalIdHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationalUnitIdByExternalIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdHeaders) SetAuthorization(v string) *GetOrganizationalUnitIdByExternalIdHeaders {
	s.Authorization = &v
	return s
}

type GetOrganizationalUnitIdByExternalIdRequest struct {
	// The external ID of the organizational unit. The external ID can be used to map external data to the data of the organizational unit in Employee Identity and Access Management (EIAM) of Identity as a Service (IDaaS). By default, the external ID is the organizational unit ID.
	//
	// Note: For organizational units with the same source type and source ID, each organizational unit has a unique external ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	// The source ID of the organizational unit.
	//
	// If the organizational unit was created in IDaaS, its source ID is the ID of the IDaaS instance. If the organizational unit was imported, its source ID is the enterprise ID in the source. For example, if the organizational unit was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	OrganizationalUnitSourceId *string `json:"organizationalUnitSourceId,omitempty" xml:"organizationalUnitSourceId,omitempty"`
	// The source type of the organizational unit. Valid values:
	//
	// 	- build_in: The organizational unit was created in IDaaS.
	//
	// 	- ding_talk: The organizational unit was imported from DingTalk.
	//
	// 	- ad: The organizational unit was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The organizational unit was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// This parameter is required.
	//
	// example:
	//
	// build_in
	OrganizationalUnitSourceType *string `json:"organizationalUnitSourceType,omitempty" xml:"organizationalUnitSourceType,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitIdByExternalIdRequest {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitIdByExternalIdRequest {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitIdByExternalIdRequest {
	s.OrganizationalUnitSourceType = &v
	return s
}

type GetOrganizationalUnitIdByExternalIdResponseBody struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdResponseBody) SetOrganizationalUnitId(v string) *GetOrganizationalUnitIdByExternalIdResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

type GetOrganizationalUnitIdByExternalIdResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrganizationalUnitIdByExternalIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) SetHeaders(v map[string]*string) *GetOrganizationalUnitIdByExternalIdResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) SetStatusCode(v int32) *GetOrganizationalUnitIdByExternalIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdResponse) SetBody(v *GetOrganizationalUnitIdByExternalIdResponseBody) *GetOrganizationalUnitIdByExternalIdResponse {
	s.Body = v
	return s
}

type GetUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHeaders) SetCommonHeaders(v map[string]*string) *GetUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHeaders) SetAuthorization(v string) *GetUserHeaders {
	s.Authorization = &v
	return s
}

type GetUserResponseBody struct {
	// The time when the account expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	AccountExpireTime *int64 `json:"accountExpireTime,omitempty" xml:"accountExpireTime,omitempty"`
	// The time when the account was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The extended fields of the account.
	CustomFields []*GetUserResponseBodyCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// The description of the account.
	//
	// example:
	//
	// xxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Indicates whether the email address has been verified. A value of true indicates that the email address has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the email address has not been verified.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	// 账户所属组列表。
	Groups []*GetUserResponseBodyGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The time when the account lock expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	LockExpireTime *int64 `json:"lockExpireTime,omitempty" xml:"lockExpireTime,omitempty"`
	// The organizational units to which the account belongs.
	OrganizationalUnits []*GetUserResponseBodyOrganizationalUnits `json:"organizationalUnits,omitempty" xml:"organizationalUnits,omitempty" type:"Repeated"`
	// Indicates whether the password is set.
	//
	// example:
	//
	// true
	PasswordSet *bool `json:"passwordSet,omitempty" xml:"passwordSet,omitempty"`
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// Indicates whether the mobile number has been verified. A value of true indicates that the mobile number has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the mobile number has not been verified.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	// The ID of the primary organizational unit of the account.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	PrimaryOrganizationalUnitId *string `json:"primaryOrganizationalUnitId,omitempty" xml:"primaryOrganizationalUnitId,omitempty"`
	// The time when the account was registered. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	RegisterTime *int64 `json:"registerTime,omitempty" xml:"registerTime,omitempty"`
	// The status of the account. Valid values: enabled disabled
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the account was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The external ID of the account. The external ID can be used to map external data to the data of the account in EIAM of Identity as a Service (IDaaS). By default, the external ID is the account ID.
	//
	// Note: For accounts with the same source type and source ID, each account has a unique external ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The source ID of the account.
	//
	// If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"userSourceId,omitempty" xml:"userSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// 	- build_in: The account was created in IDaaS.
	//
	// 	- ding_talk: The account was imported from DingTalk.
	//
	// 	- ad: The account was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The account was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"userSourceType,omitempty" xml:"userSourceType,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// name001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetAccountExpireTime(v int64) *GetUserResponseBody {
	s.AccountExpireTime = &v
	return s
}

func (s *GetUserResponseBody) SetCreateTime(v int64) *GetUserResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBody) SetCustomFields(v []*GetUserResponseBodyCustomFields) *GetUserResponseBody {
	s.CustomFields = v
	return s
}

func (s *GetUserResponseBody) SetDescription(v string) *GetUserResponseBody {
	s.Description = &v
	return s
}

func (s *GetUserResponseBody) SetDisplayName(v string) *GetUserResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBody) SetEmail(v string) *GetUserResponseBody {
	s.Email = &v
	return s
}

func (s *GetUserResponseBody) SetEmailVerified(v bool) *GetUserResponseBody {
	s.EmailVerified = &v
	return s
}

func (s *GetUserResponseBody) SetGroups(v []*GetUserResponseBodyGroups) *GetUserResponseBody {
	s.Groups = v
	return s
}

func (s *GetUserResponseBody) SetInstanceId(v string) *GetUserResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetUserResponseBody) SetLockExpireTime(v int64) *GetUserResponseBody {
	s.LockExpireTime = &v
	return s
}

func (s *GetUserResponseBody) SetOrganizationalUnits(v []*GetUserResponseBodyOrganizationalUnits) *GetUserResponseBody {
	s.OrganizationalUnits = v
	return s
}

func (s *GetUserResponseBody) SetPasswordSet(v bool) *GetUserResponseBody {
	s.PasswordSet = &v
	return s
}

func (s *GetUserResponseBody) SetPhoneNumber(v string) *GetUserResponseBody {
	s.PhoneNumber = &v
	return s
}

func (s *GetUserResponseBody) SetPhoneNumberVerified(v bool) *GetUserResponseBody {
	s.PhoneNumberVerified = &v
	return s
}

func (s *GetUserResponseBody) SetPhoneRegion(v string) *GetUserResponseBody {
	s.PhoneRegion = &v
	return s
}

func (s *GetUserResponseBody) SetPrimaryOrganizationalUnitId(v string) *GetUserResponseBody {
	s.PrimaryOrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBody) SetRegisterTime(v int64) *GetUserResponseBody {
	s.RegisterTime = &v
	return s
}

func (s *GetUserResponseBody) SetStatus(v string) *GetUserResponseBody {
	s.Status = &v
	return s
}

func (s *GetUserResponseBody) SetUpdateTime(v int64) *GetUserResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetUserResponseBody) SetUserExternalId(v string) *GetUserResponseBody {
	s.UserExternalId = &v
	return s
}

func (s *GetUserResponseBody) SetUserId(v string) *GetUserResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBody) SetUserSourceId(v string) *GetUserResponseBody {
	s.UserSourceId = &v
	return s
}

func (s *GetUserResponseBody) SetUserSourceType(v string) *GetUserResponseBody {
	s.UserSourceType = &v
	return s
}

func (s *GetUserResponseBody) SetUsername(v string) *GetUserResponseBody {
	s.Username = &v
	return s
}

type GetUserResponseBodyCustomFields struct {
	// The name of the extended field.
	//
	// example:
	//
	// xxxx
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// The value of the extended field. Field values are returned as strings regardless of the data types of the fields. For example, true or false is returned for a Boolean field.
	//
	// example:
	//
	// 字段数据值
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s GetUserResponseBodyCustomFields) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyCustomFields) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyCustomFields) SetFieldName(v string) *GetUserResponseBodyCustomFields {
	s.FieldName = &v
	return s
}

func (s *GetUserResponseBodyCustomFields) SetFieldValue(v string) *GetUserResponseBodyCustomFields {
	s.FieldValue = &v
	return s
}

type GetUserResponseBodyGroups struct {
	// 组描述。
	//
	// example:
	//
	// description_demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 组ID。
	//
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 组名称。
	//
	// example:
	//
	// name_test
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s GetUserResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyGroups) SetDescription(v string) *GetUserResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *GetUserResponseBodyGroups) SetGroupId(v string) *GetUserResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *GetUserResponseBodyGroups) SetGroupName(v string) *GetUserResponseBodyGroups {
	s.GroupName = &v
	return s
}

type GetUserResponseBodyOrganizationalUnits struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	// Indicates whether the organizational unit is the primary organizational unit.
	//
	// example:
	//
	// true
	Primary *bool `json:"primary,omitempty" xml:"primary,omitempty"`
}

func (s GetUserResponseBodyOrganizationalUnits) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyOrganizationalUnits) SetOrganizationalUnitId(v string) *GetUserResponseBodyOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBodyOrganizationalUnits) SetOrganizationalUnitName(v string) *GetUserResponseBodyOrganizationalUnits {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetUserResponseBodyOrganizationalUnits) SetPrimary(v bool) *GetUserResponseBodyOrganizationalUnits {
	s.Primary = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetUserIdByEmailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserIdByEmailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByEmailHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByEmailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByEmailHeaders) SetAuthorization(v string) *GetUserIdByEmailHeaders {
	s.Authorization = &v
	return s
}

type GetUserIdByEmailRequest struct {
	// The email address of the user who owns the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s GetUserIdByEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByEmailRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailRequest) SetEmail(v string) *GetUserIdByEmailRequest {
	s.Email = &v
	return s
}

type GetUserIdByEmailResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByEmailResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailResponseBody) SetUserId(v string) *GetUserIdByEmailResponseBody {
	s.UserId = &v
	return s
}

type GetUserIdByEmailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByEmailResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailResponse) SetHeaders(v map[string]*string) *GetUserIdByEmailResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByEmailResponse) SetStatusCode(v int32) *GetUserIdByEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByEmailResponse) SetBody(v *GetUserIdByEmailResponseBody) *GetUserIdByEmailResponse {
	s.Body = v
	return s
}

type GetUserIdByPhoneNumberHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserIdByPhoneNumberHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByPhoneNumberHeaders) SetAuthorization(v string) *GetUserIdByPhoneNumberHeaders {
	s.Authorization = &v
	return s
}

type GetUserIdByPhoneNumberRequest struct {
	// The mobile number of the user who owns the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

func (s GetUserIdByPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberRequest) SetPhoneNumber(v string) *GetUserIdByPhoneNumberRequest {
	s.PhoneNumber = &v
	return s
}

type GetUserIdByPhoneNumberResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberResponseBody) SetUserId(v string) *GetUserIdByPhoneNumberResponseBody {
	s.UserId = &v
	return s
}

type GetUserIdByPhoneNumberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberResponse) SetHeaders(v map[string]*string) *GetUserIdByPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByPhoneNumberResponse) SetStatusCode(v int32) *GetUserIdByPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByPhoneNumberResponse) SetBody(v *GetUserIdByPhoneNumberResponseBody) *GetUserIdByPhoneNumberResponse {
	s.Body = v
	return s
}

type GetUserIdByUserExternalIdHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserIdByUserExternalIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUserExternalIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByUserExternalIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByUserExternalIdHeaders) SetAuthorization(v string) *GetUserIdByUserExternalIdHeaders {
	s.Authorization = &v
	return s
}

type GetUserIdByUserExternalIdRequest struct {
	// The external ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx001
	UserExternalId *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	// The source ID of the account. If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"userSourceId,omitempty" xml:"userSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// 	- build_in: The account was created in Identity as a Service (IDaaS).
	//
	// 	- ding_talk: The account was imported from DingTalk.
	//
	// 	- ad: The account was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The account was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// This parameter is required.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"userSourceType,omitempty" xml:"userSourceType,omitempty"`
}

func (s GetUserIdByUserExternalIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUserExternalIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdRequest) SetUserExternalId(v string) *GetUserIdByUserExternalIdRequest {
	s.UserExternalId = &v
	return s
}

func (s *GetUserIdByUserExternalIdRequest) SetUserSourceId(v string) *GetUserIdByUserExternalIdRequest {
	s.UserSourceId = &v
	return s
}

func (s *GetUserIdByUserExternalIdRequest) SetUserSourceType(v string) *GetUserIdByUserExternalIdRequest {
	s.UserSourceType = &v
	return s
}

type GetUserIdByUserExternalIdResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByUserExternalIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUserExternalIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdResponseBody) SetUserId(v string) *GetUserIdByUserExternalIdResponseBody {
	s.UserId = &v
	return s
}

type GetUserIdByUserExternalIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByUserExternalIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByUserExternalIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUserExternalIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdResponse) SetHeaders(v map[string]*string) *GetUserIdByUserExternalIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByUserExternalIdResponse) SetStatusCode(v int32) *GetUserIdByUserExternalIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByUserExternalIdResponse) SetBody(v *GetUserIdByUserExternalIdResponseBody) *GetUserIdByUserExternalIdResponse {
	s.Body = v
	return s
}

type GetUserIdByUsernameHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserIdByUsernameHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUsernameHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByUsernameHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByUsernameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByUsernameHeaders) SetAuthorization(v string) *GetUserIdByUsernameHeaders {
	s.Authorization = &v
	return s
}

type GetUserIdByUsernameRequest struct {
	// The username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// username_001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetUserIdByUsernameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUsernameRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByUsernameRequest) SetUsername(v string) *GetUserIdByUsernameRequest {
	s.Username = &v
	return s
}

type GetUserIdByUsernameResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserIdByUsernameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUsernameResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdByUsernameResponseBody) SetUserId(v string) *GetUserIdByUsernameResponseBody {
	s.UserId = &v
	return s
}

type GetUserIdByUsernameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByUsernameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByUsernameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserIdByUsernameResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByUsernameResponse) SetHeaders(v map[string]*string) *GetUserIdByUsernameResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByUsernameResponse) SetStatusCode(v int32) *GetUserIdByUsernameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByUsernameResponse) SetBody(v *GetUserIdByUsernameResponseBody) *GetUserIdByUsernameResponse {
	s.Body = v
	return s
}

type GetUserInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetUserInfoHeaders) SetCommonHeaders(v map[string]*string) *GetUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserInfoHeaders) SetAuthorization(v string) *GetUserInfoHeaders {
	s.Authorization = &v
	return s
}

type GetUserInfoResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponse) SetHeaders(v map[string]*string) *GetUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserInfoResponse) SetStatusCode(v int32) *GetUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInfoResponse) SetBody(v map[string]interface{}) *GetUserInfoResponse {
	s.Body = v
	return s
}

type ListGroupsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsHeaders) GoString() string {
	return s.String()
}

func (s *ListGroupsHeaders) SetCommonHeaders(v map[string]*string) *ListGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListGroupsHeaders) SetAuthorization(v string) *ListGroupsHeaders {
	s.Authorization = &v
	return s
}

type ListGroupsRequest struct {
	// example:
	//
	// group_xxx
	GroupNameStartWith *string `json:"groupNameStartWith,omitempty" xml:"groupNameStartWith,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetGroupNameStartWith(v string) *ListGroupsRequest {
	s.GroupNameStartWith = &v
	return s
}

func (s *ListGroupsRequest) SetMaxResults(v int32) *ListGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsRequest) SetNextToken(v string) *ListGroupsRequest {
	s.NextToken = &v
	return s
}

type ListGroupsResponseBody struct {
	Data []*ListGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetData(v []*ListGroupsResponseBodyData) *ListGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupsResponseBody) SetMaxResults(v int32) *ListGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsResponseBody) SetNextToken(v string) *ListGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotalCount(v int64) *ListGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupsResponseBodyData struct {
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// description_demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupExternalId *string `json:"groupExternalId,omitempty" xml:"groupExternalId,omitempty"`
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// name_test
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupSourceId *string `json:"groupSourceId,omitempty" xml:"groupSourceId,omitempty"`
	// example:
	//
	// build_in
	GroupSourceType *string `json:"groupSourceType,omitempty" xml:"groupSourceType,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyData) SetCreateTime(v int64) *ListGroupsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetDescription(v string) *ListGroupsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupExternalId(v string) *ListGroupsResponseBodyData {
	s.GroupExternalId = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupId(v string) *ListGroupsResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupName(v string) *ListGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupSourceId(v string) *ListGroupsResponseBodyData {
	s.GroupSourceId = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetGroupSourceType(v string) *ListGroupsResponseBodyData {
	s.GroupSourceType = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetInstanceId(v string) *ListGroupsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsResponseBodyData) SetUpdateTime(v int64) *ListGroupsResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListGroupsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsResponse) SetHeaders(v map[string]*string) *ListGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsResponse) SetStatusCode(v int32) *ListGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListGroupsForUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListGroupsForUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserHeaders) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserHeaders) SetCommonHeaders(v map[string]*string) *ListGroupsForUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListGroupsForUserHeaders) SetAuthorization(v string) *ListGroupsForUserHeaders {
	s.Authorization = &v
	return s
}

type ListGroupsForUserRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListGroupsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserRequest) SetMaxResults(v int32) *ListGroupsForUserRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsForUserRequest) SetNextToken(v string) *ListGroupsForUserRequest {
	s.NextToken = &v
	return s
}

type ListGroupsForUserResponseBody struct {
	Data []*ListGroupsForUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListGroupsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBody) SetData(v []*ListGroupsForUserResponseBodyData) *ListGroupsForUserResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupsForUserResponseBody) SetMaxResults(v int64) *ListGroupsForUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsForUserResponseBody) SetNextToken(v string) *ListGroupsForUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupsForUserResponseBody) SetTotalCount(v int64) *ListGroupsForUserResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupsForUserResponseBodyData struct {
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupMemberRelationSourceId *string `json:"groupMemberRelationSourceId,omitempty" xml:"groupMemberRelationSourceId,omitempty"`
	// example:
	//
	// build_in
	GroupMemberRelationSourceType *string `json:"groupMemberRelationSourceType,omitempty" xml:"groupMemberRelationSourceType,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListGroupsForUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyData) SetGroupId(v string) *ListGroupsForUserResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyData) SetGroupMemberRelationSourceId(v string) *ListGroupsForUserResponseBodyData {
	s.GroupMemberRelationSourceId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyData) SetGroupMemberRelationSourceType(v string) *ListGroupsForUserResponseBodyData {
	s.GroupMemberRelationSourceType = &v
	return s
}

func (s *ListGroupsForUserResponseBodyData) SetInstanceId(v string) *ListGroupsForUserResponseBodyData {
	s.InstanceId = &v
	return s
}

type ListGroupsForUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponse) SetHeaders(v map[string]*string) *ListGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsForUserResponse) SetStatusCode(v int32) *ListGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsForUserResponse) SetBody(v *ListGroupsForUserResponseBody) *ListGroupsForUserResponse {
	s.Body = v
	return s
}

type ListOrganizationalUnitParentIdsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListOrganizationalUnitParentIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsHeaders) SetCommonHeaders(v map[string]*string) *ListOrganizationalUnitParentIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOrganizationalUnitParentIdsHeaders) SetAuthorization(v string) *ListOrganizationalUnitParentIdsHeaders {
	s.Authorization = &v
	return s
}

type ListOrganizationalUnitParentIdsResponseBody struct {
	// The IDs of the parent organizational units. The IDs of the organizational unit are ordered based on their levels from high to low. Only the IDs of the organizational units within the authorization scope are displayed.
	//
	// example:
	//
	// [ou_xxx001]
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
}

func (s ListOrganizationalUnitParentIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsResponseBody) SetParentIds(v []*string) *ListOrganizationalUnitParentIdsResponseBody {
	s.ParentIds = v
	return s
}

type ListOrganizationalUnitParentIdsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitParentIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitParentIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitParentIdsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitParentIdsResponse) SetStatusCode(v int32) *ListOrganizationalUnitParentIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitParentIdsResponse) SetBody(v *ListOrganizationalUnitParentIdsResponseBody) *ListOrganizationalUnitParentIdsResponse {
	s.Body = v
	return s
}

type ListOrganizationalUnitsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListOrganizationalUnitsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsHeaders) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsHeaders) SetCommonHeaders(v map[string]*string) *ListOrganizationalUnitsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOrganizationalUnitsHeaders) SetAuthorization(v string) *ListOrganizationalUnitsHeaders {
	s.Authorization = &v
	return s
}

type ListOrganizationalUnitsRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the parent organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s ListOrganizationalUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsRequest) SetPageNumber(v int32) *ListOrganizationalUnitsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetPageSize(v int32) *ListOrganizationalUnitsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetParentId(v string) *ListOrganizationalUnitsRequest {
	s.ParentId = &v
	return s
}

type ListOrganizationalUnitsResponseBody struct {
	// The queried organizational units.
	Data []*ListOrganizationalUnitsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListOrganizationalUnitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBody) SetData(v []*ListOrganizationalUnitsResponseBodyData) *ListOrganizationalUnitsResponseBody {
	s.Data = v
	return s
}

func (s *ListOrganizationalUnitsResponseBody) SetTotalCount(v int64) *ListOrganizationalUnitsResponseBody {
	s.TotalCount = &v
	return s
}

type ListOrganizationalUnitsResponseBodyData struct {
	// The time when the organizational unit was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652083425923
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the organizational unit.
	//
	// example:
	//
	// test organizational unit
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The external ID of the organizational unit. The external ID can be used to map external data to the data of the organizational unit in EIAM of Identity as a Service (IDaaS). By default, the external ID is the organizational unit ID.
	//
	// Note: For organizational units with the same source type and source ID, each organizational unit has a unique external ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	// The source ID of the organizational unit.
	//
	// If the organizational unit was created in IDaaS, its source ID is the ID of the IDaaS instance. If the organizational unit was imported, its source ID is the enterprise ID in the source. For example, if the organizational unit was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	OrganizationalUnitSourceId *string `json:"organizationalUnitSourceId,omitempty" xml:"organizationalUnitSourceId,omitempty"`
	// The source type of the organizational unit. Valid values:
	//
	// 	- build_in: The organizational unit was created in IDaaS.
	//
	// 	- ding_talk: The organizational unit was imported from DingTalk.
	//
	// 	- ad: The organizational unit was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The organizational unit was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	OrganizationalUnitSourceType *string `json:"organizationalUnitSourceType,omitempty" xml:"organizationalUnitSourceType,omitempty"`
	// The ID of the parent organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The time when the organizational unit was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652083425923
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListOrganizationalUnitsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBodyData) SetCreateTime(v int64) *ListOrganizationalUnitsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetDescription(v string) *ListOrganizationalUnitsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetInstanceId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitExternalId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitName(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitName = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitSourceId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitSourceType(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetParentId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetUpdateTime(v int64) *ListOrganizationalUnitsResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListOrganizationalUnitsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitsResponse) SetStatusCode(v int32) *ListOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitsResponse) SetBody(v *ListOrganizationalUnitsResponseBody) *ListOrganizationalUnitsResponse {
	s.Body = v
	return s
}

type ListUsersHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListUsersHeaders) SetCommonHeaders(v map[string]*string) *ListUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUsersHeaders) SetAuthorization(v string) *ListUsersHeaders {
	s.Authorization = &v
	return s
}

type ListUsersRequest struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetOrganizationalUnitId(v string) *ListUsersRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	// The queried EIAM accounts.
	Data []*ListUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetData(v []*ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

type ListUsersResponseBodyData struct {
	// The time when the account expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	AccountExpireTime *int64 `json:"accountExpireTime,omitempty" xml:"accountExpireTime,omitempty"`
	// The time when the account was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// xxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address of the user who owns the account.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Indicates whether the email address has been verified. A value of true indicates that the email address has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the email address has not been verified.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The time when the account lock expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	LockExpireTime *int64 `json:"lockExpireTime,omitempty" xml:"lockExpireTime,omitempty"`
	// 密码是否已设置
	//
	// example:
	//
	// true
	PasswordSet *bool `json:"passwordSet,omitempty" xml:"passwordSet,omitempty"`
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// Indicates whether the mobile number has been verified. A value of true indicates that the mobile number has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the mobile number has not been verified.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	// The time when the account was registered. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	RegisterTime *int64 `json:"registerTime,omitempty" xml:"registerTime,omitempty"`
	// The status of the account. Valid values: enabled disabled
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the account was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The external ID of the account. The external ID can be used to map external data to the data of the account in EIAM of Identity as a Service (IDaaS). By default, the external ID is the account ID.
	//
	// Note: For accounts with the same source type and source ID, each account has a unique external ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The source ID of the account.
	//
	// If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"userSourceId,omitempty" xml:"userSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// 	- build_in: The account was created in IDaaS.
	//
	// 	- ding_talk: The account was imported from DingTalk.
	//
	// 	- ad: The account was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The account was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"userSourceType,omitempty" xml:"userSourceType,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// name001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetAccountExpireTime(v int64) *ListUsersResponseBodyData {
	s.AccountExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetCreateTime(v int64) *ListUsersResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetDescription(v string) *ListUsersResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyData) SetDisplayName(v string) *ListUsersResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyData) SetEmail(v string) *ListUsersResponseBodyData {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyData) SetEmailVerified(v bool) *ListUsersResponseBodyData {
	s.EmailVerified = &v
	return s
}

func (s *ListUsersResponseBodyData) SetInstanceId(v string) *ListUsersResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetLockExpireTime(v int64) *ListUsersResponseBodyData {
	s.LockExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPasswordSet(v bool) *ListUsersResponseBodyData {
	s.PasswordSet = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhoneNumber(v string) *ListUsersResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhoneNumberVerified(v bool) *ListUsersResponseBodyData {
	s.PhoneNumberVerified = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPhoneRegion(v string) *ListUsersResponseBodyData {
	s.PhoneRegion = &v
	return s
}

func (s *ListUsersResponseBodyData) SetRegisterTime(v int64) *ListUsersResponseBodyData {
	s.RegisterTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetStatus(v string) *ListUsersResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUpdateTime(v int64) *ListUsersResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserExternalId(v string) *ListUsersResponseBodyData {
	s.UserExternalId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserId(v string) *ListUsersResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserSourceId(v string) *ListUsersResponseBodyData {
	s.UserSourceId = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUserSourceType(v string) *ListUsersResponseBodyData {
	s.UserSourceType = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUsername(v string) *ListUsersResponseBodyData {
	s.Username = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListUsersForGroupHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListUsersForGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupHeaders) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupHeaders) SetCommonHeaders(v map[string]*string) *ListUsersForGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUsersForGroupHeaders) SetAuthorization(v string) *ListUsersForGroupHeaders {
	s.Authorization = &v
	return s
}

type ListUsersForGroupRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListUsersForGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupRequest) SetMaxResults(v int32) *ListUsersForGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersForGroupRequest) SetNextToken(v string) *ListUsersForGroupRequest {
	s.NextToken = &v
	return s
}

type ListUsersForGroupResponseBody struct {
	Data []*ListUsersForGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUsersForGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBody) SetData(v []*ListUsersForGroupResponseBodyData) *ListUsersForGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersForGroupResponseBody) SetMaxResults(v int32) *ListUsersForGroupResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetNextToken(v string) *ListUsersForGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetTotalCount(v int64) *ListUsersForGroupResponseBody {
	s.TotalCount = &v
	return s
}

type ListUsersForGroupResponseBodyData struct {
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// user_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListUsersForGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyData) SetInstanceId(v string) *ListUsersForGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForGroupResponseBodyData) SetUserId(v string) *ListUsersForGroupResponseBodyData {
	s.UserId = &v
	return s
}

type ListUsersForGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersForGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponse) SetHeaders(v map[string]*string) *ListUsersForGroupResponse {
	s.Headers = v
	return s
}

func (s *ListUsersForGroupResponse) SetStatusCode(v int32) *ListUsersForGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersForGroupResponse) SetBody(v *ListUsersForGroupResponseBody) *ListUsersForGroupResponse {
	s.Body = v
	return s
}

type PatchGroupHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PatchGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s PatchGroupHeaders) GoString() string {
	return s.String()
}

func (s *PatchGroupHeaders) SetCommonHeaders(v map[string]*string) *PatchGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchGroupHeaders) SetAuthorization(v string) *PatchGroupHeaders {
	s.Authorization = &v
	return s
}

type PatchGroupRequest struct {
	// example:
	//
	// name001
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s PatchGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s PatchGroupRequest) GoString() string {
	return s.String()
}

func (s *PatchGroupRequest) SetGroupName(v string) *PatchGroupRequest {
	s.GroupName = &v
	return s
}

type PatchGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PatchGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s PatchGroupResponse) GoString() string {
	return s.String()
}

func (s *PatchGroupResponse) SetHeaders(v map[string]*string) *PatchGroupResponse {
	s.Headers = v
	return s
}

func (s *PatchGroupResponse) SetStatusCode(v int32) *PatchGroupResponse {
	s.StatusCode = &v
	return s
}

type PatchOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PatchOrganizationalUnitHeaders) String() string {
	return tea.Prettify(s)
}

func (s PatchOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *PatchOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchOrganizationalUnitHeaders) SetAuthorization(v string) *PatchOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

type PatchOrganizationalUnitRequest struct {
	// The description of the organizational unit.
	//
	// example:
	//
	// test organizational unit
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
}

func (s PatchOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s PatchOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitRequest) SetDescription(v string) *PatchOrganizationalUnitRequest {
	s.Description = &v
	return s
}

func (s *PatchOrganizationalUnitRequest) SetOrganizationalUnitName(v string) *PatchOrganizationalUnitRequest {
	s.OrganizationalUnitName = &v
	return s
}

type PatchOrganizationalUnitResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PatchOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s PatchOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitResponse) SetHeaders(v map[string]*string) *PatchOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *PatchOrganizationalUnitResponse) SetStatusCode(v int32) *PatchOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

type PatchUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PatchUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s PatchUserHeaders) GoString() string {
	return s.String()
}

func (s *PatchUserHeaders) SetCommonHeaders(v map[string]*string) *PatchUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchUserHeaders) SetAuthorization(v string) *PatchUserHeaders {
	s.Authorization = &v
	return s
}

type PatchUserRequest struct {
	CustomFields []*PatchUserRequestCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// The display name of the account.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address of the user who owns the account.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Indicates whether the email address is verified. This field is required if an email address is specified. If you have no special requirement, set this parameter to true.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// Indicates whether the mobile number is verified. This field is required if a mobile number is specified. If you have no special requirement, set this parameter to true.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +. This parameter is required if a mobile number is specified.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// name001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s PatchUserRequest) String() string {
	return tea.Prettify(s)
}

func (s PatchUserRequest) GoString() string {
	return s.String()
}

func (s *PatchUserRequest) SetCustomFields(v []*PatchUserRequestCustomFields) *PatchUserRequest {
	s.CustomFields = v
	return s
}

func (s *PatchUserRequest) SetDisplayName(v string) *PatchUserRequest {
	s.DisplayName = &v
	return s
}

func (s *PatchUserRequest) SetEmail(v string) *PatchUserRequest {
	s.Email = &v
	return s
}

func (s *PatchUserRequest) SetEmailVerified(v bool) *PatchUserRequest {
	s.EmailVerified = &v
	return s
}

func (s *PatchUserRequest) SetPhoneNumber(v string) *PatchUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *PatchUserRequest) SetPhoneNumberVerified(v bool) *PatchUserRequest {
	s.PhoneNumberVerified = &v
	return s
}

func (s *PatchUserRequest) SetPhoneRegion(v string) *PatchUserRequest {
	s.PhoneRegion = &v
	return s
}

func (s *PatchUserRequest) SetUsername(v string) *PatchUserRequest {
	s.Username = &v
	return s
}

type PatchUserRequestCustomFields struct {
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// example:
	//
	// test_value
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// 字段操作类型，取值可选范围：
	//
	// - add：添加。
	//
	// - replace：替换。若对应扩展字段无设置值，会转换为add操作。
	//
	// - remove：移除。
	//
	// example:
	//
	// replace
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// Deprecated
	//
	// example:
	//
	// replace
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s PatchUserRequestCustomFields) String() string {
	return tea.Prettify(s)
}

func (s PatchUserRequestCustomFields) GoString() string {
	return s.String()
}

func (s *PatchUserRequestCustomFields) SetFieldName(v string) *PatchUserRequestCustomFields {
	s.FieldName = &v
	return s
}

func (s *PatchUserRequestCustomFields) SetFieldValue(v string) *PatchUserRequestCustomFields {
	s.FieldValue = &v
	return s
}

func (s *PatchUserRequestCustomFields) SetOperation(v string) *PatchUserRequestCustomFields {
	s.Operation = &v
	return s
}

func (s *PatchUserRequestCustomFields) SetOperator(v string) *PatchUserRequestCustomFields {
	s.Operator = &v
	return s
}

type PatchUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PatchUserResponse) String() string {
	return tea.Prettify(s)
}

func (s PatchUserResponse) GoString() string {
	return s.String()
}

func (s *PatchUserResponse) SetHeaders(v map[string]*string) *PatchUserResponse {
	s.Headers = v
	return s
}

func (s *PatchUserResponse) SetStatusCode(v int32) *PatchUserResponse {
	s.StatusCode = &v
	return s
}

type RemoveUserFromOrganizationalUnitsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RemoveUserFromOrganizationalUnitsHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsHeaders) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsHeaders) SetCommonHeaders(v map[string]*string) *RemoveUserFromOrganizationalUnitsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsHeaders) SetAuthorization(v string) *RemoveUserFromOrganizationalUnitsHeaders {
	s.Authorization = &v
	return s
}

type RemoveUserFromOrganizationalUnitsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [ou_wovwffm62xifdziem7an7xxxxx]
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" xml:"organizationalUnitIds,omitempty" type:"Repeated"`
}

func (s RemoveUserFromOrganizationalUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *RemoveUserFromOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

type RemoveUserFromOrganizationalUnitsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RemoveUserFromOrganizationalUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *RemoveUserFromOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsResponse) SetStatusCode(v int32) *RemoveUserFromOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

type RemoveUsersFromGroupHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RemoveUsersFromGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupHeaders) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupHeaders) SetCommonHeaders(v map[string]*string) *RemoveUsersFromGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveUsersFromGroupHeaders) SetAuthorization(v string) *RemoveUsersFromGroupHeaders {
	s.Authorization = &v
	return s
}

type RemoveUsersFromGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [user_d6sbsuumeta4h66ec3il7yxxxx}
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RemoveUsersFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupRequest) SetUserIds(v []*string) *RemoveUsersFromGroupRequest {
	s.UserIds = v
	return s
}

type RemoveUsersFromGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RemoveUsersFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponse) SetHeaders(v map[string]*string) *RemoveUsersFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersFromGroupResponse) SetStatusCode(v int32) *RemoveUsersFromGroupResponse {
	s.StatusCode = &v
	return s
}

type RevokeTokenRequest struct {
	// The client ID.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// The client secret.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	// The token to be revoked.
	//
	// This parameter is required.
	//
	// example:
	//
	// ATxxxx
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The type of the token. Valid values: access_token refresh_token
	//
	// example:
	//
	// access_token
	TokenTypeHint *string `json:"token_type_hint,omitempty" xml:"token_type_hint,omitempty"`
}

func (s RevokeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeTokenRequest) GoString() string {
	return s.String()
}

func (s *RevokeTokenRequest) SetClientId(v string) *RevokeTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RevokeTokenRequest) SetClientSecret(v string) *RevokeTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *RevokeTokenRequest) SetToken(v string) *RevokeTokenRequest {
	s.Token = &v
	return s
}

func (s *RevokeTokenRequest) SetTokenTypeHint(v string) *RevokeTokenRequest {
	s.TokenTypeHint = &v
	return s
}

type RevokeTokenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeTokenResponse) GoString() string {
	return s.String()
}

func (s *RevokeTokenResponse) SetHeaders(v map[string]*string) *RevokeTokenResponse {
	s.Headers = v
	return s
}

func (s *RevokeTokenResponse) SetStatusCode(v int32) *RevokeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeTokenResponse) SetBody(v map[string]interface{}) *RevokeTokenResponse {
	s.Body = v
	return s
}

type SetUserPrimaryOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *SetUserPrimaryOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitHeaders) SetAuthorization(v string) *SetUserPrimaryOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

type SetUserPrimaryOrganizationalUnitRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *SetUserPrimaryOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

type SetUserPrimaryOrganizationalUnitResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitResponse) SetHeaders(v map[string]*string) *SetUserPrimaryOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitResponse) SetStatusCode(v int32) *SetUserPrimaryOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

type UpdateUserPasswordHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateUserPasswordHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserPasswordHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserPasswordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserPasswordHeaders) SetAuthorization(v string) *UpdateUserPasswordHeaders {
	s.Authorization = &v
	return s
}

type UpdateUserPasswordRequest struct {
	// example:
	//
	// xxxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s UpdateUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordRequest) SetPassword(v string) *UpdateUserPasswordRequest {
	s.Password = &v
	return s
}

type UpdateUserPasswordResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordResponse) SetHeaders(v map[string]*string) *UpdateUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserPasswordResponse) SetStatusCode(v int32) *UpdateUserPasswordResponse {
	s.StatusCode = &v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eiam-developerapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将账户加入多个组织
//
// @param request - AddUserToOrganizationalUnitsRequest
//
// @param headers - AddUserToOrganizationalUnitsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToOrganizationalUnitsResponse
func (client *Client) AddUserToOrganizationalUnitsWithOptions(instanceId *string, applicationId *string, userId *string, request *AddUserToOrganizationalUnitsRequest, headers *AddUserToOrganizationalUnitsHeaders, runtime *util.RuntimeOptions) (_result *AddUserToOrganizationalUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		body["organizationalUnitIds"] = request.OrganizationalUnitIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserToOrganizationalUnits"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId)) + "/actions/addUserToOrganizationalUnits"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AddUserToOrganizationalUnitsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AddUserToOrganizationalUnitsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 将账户加入多个组织
//
// @param request - AddUserToOrganizationalUnitsRequest
//
// @return AddUserToOrganizationalUnitsResponse
func (client *Client) AddUserToOrganizationalUnits(instanceId *string, applicationId *string, userId *string, request *AddUserToOrganizationalUnitsRequest) (_result *AddUserToOrganizationalUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddUserToOrganizationalUnitsHeaders{}
	_result = &AddUserToOrganizationalUnitsResponse{}
	_body, _err := client.AddUserToOrganizationalUnitsWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为指定组批量关联账户
//
// @param request - AddUsersToGroupRequest
//
// @param headers - AddUsersToGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroupWithOptions(instanceId *string, applicationId *string, groupId *string, request *AddUsersToGroupRequest, headers *AddUsersToGroupHeaders, runtime *util.RuntimeOptions) (_result *AddUsersToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUsersToGroup"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/actions/addUsersToGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AddUsersToGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AddUsersToGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 为指定组批量关联账户
//
// @param request - AddUsersToGroupRequest
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroup(instanceId *string, applicationId *string, groupId *string, request *AddUsersToGroupRequest) (_result *AddUsersToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddUsersToGroupHeaders{}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.AddUsersToGroupWithOptions(instanceId, applicationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个EIAM组
//
// @param request - CreateGroupRequest
//
// @param headers - CreateGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(instanceId *string, applicationId *string, request *CreateGroupRequest, headers *CreateGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupExternalId)) {
		body["groupExternalId"] = request.GroupExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建一个EIAM组
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(instanceId *string, applicationId *string, request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupHeaders{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an organizational unit.
//
// @param request - CreateOrganizationalUnitRequest
//
// @param headers - CreateOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrganizationalUnitResponse
func (client *Client) CreateOrganizationalUnitWithOptions(instanceId *string, applicationId *string, request *CreateOrganizationalUnitRequest, headers *CreateOrganizationalUnitHeaders, runtime *util.RuntimeOptions) (_result *CreateOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitExternalId)) {
		body["organizationalUnitExternalId"] = request.OrganizationalUnitExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitName)) {
		body["organizationalUnitName"] = request.OrganizationalUnitName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["parentId"] = request.ParentId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrganizationalUnit"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/organizationalUnits"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateOrganizationalUnitResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateOrganizationalUnitResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an organizational unit.
//
// @param request - CreateOrganizationalUnitRequest
//
// @return CreateOrganizationalUnitResponse
func (client *Client) CreateOrganizationalUnit(instanceId *string, applicationId *string, request *CreateOrganizationalUnitRequest) (_result *CreateOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrganizationalUnitHeaders{}
	_result = &CreateOrganizationalUnitResponse{}
	_body, _err := client.CreateOrganizationalUnitWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Employee Identity and Access Management (EIAM) account in an organizational unit.
//
// @param request - CreateUserRequest
//
// @param headers - CreateUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(instanceId *string, applicationId *string, request *CreateUserRequest, headers *CreateUserHeaders, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		body["customFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EmailVerified)) {
		body["emailVerified"] = request.EmailVerified
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordInitializationConfig)) {
		body["passwordInitializationConfig"] = request.PasswordInitializationConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["phoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberVerified)) {
		body["phoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneRegion)) {
		body["phoneRegion"] = request.PhoneRegion
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryOrganizationalUnitId)) {
		body["primaryOrganizationalUnitId"] = request.PrimaryOrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.UserExternalId)) {
		body["userExternalId"] = request.UserExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an Employee Identity and Access Management (EIAM) account in an organizational unit.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(instanceId *string, applicationId *string, request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUserHeaders{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定组
//
// @param headers - DeleteGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(instanceId *string, applicationId *string, groupId *string, headers *DeleteGroupHeaders, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除指定组
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(instanceId *string, applicationId *string, groupId *string) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteGroupHeaders{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(instanceId, applicationId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an organizational unit.
//
// @param headers - DeleteOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOrganizationalUnitResponse
func (client *Client) DeleteOrganizationalUnitWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, headers *DeleteOrganizationalUnitHeaders, runtime *util.RuntimeOptions) (_result *DeleteOrganizationalUnitResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOrganizationalUnit"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/organizationalUnits/" + tea.StringValue(openapiutil.GetEncodeParam(organizationalUnitId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteOrganizationalUnitResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteOrganizationalUnitResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an organizational unit.
//
// @return DeleteOrganizationalUnitResponse
func (client *Client) DeleteOrganizationalUnit(instanceId *string, applicationId *string, organizationalUnitId *string) (_result *DeleteOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteOrganizationalUnitHeaders{}
	_result = &DeleteOrganizationalUnitResponse{}
	_body, _err := client.DeleteOrganizationalUnitWithOptions(instanceId, applicationId, organizationalUnitId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) account.
//
// @param headers - DeleteUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *DeleteUserHeaders, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) account.
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(instanceId *string, applicationId *string, userId *string) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUserHeaders{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an Employee Identity and Access Management (EIAM) account.
//
// @param headers - DisableUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableUserResponse
func (client *Client) DisableUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *DisableUserHeaders, runtime *util.RuntimeOptions) (_result *DisableUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("DisableUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId)) + "/actions/disable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DisableUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DisableUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Disables an Employee Identity and Access Management (EIAM) account.
//
// @return DisableUserResponse
func (client *Client) DisableUser(instanceId *string, applicationId *string, userId *string) (_result *DisableUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DisableUserHeaders{}
	_result = &DisableUserResponse{}
	_body, _err := client.DisableUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an Employee Identity and Access Management (EIAM) account.
//
// @param headers - EnableUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableUserResponse
func (client *Client) EnableUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *EnableUserHeaders, runtime *util.RuntimeOptions) (_result *EnableUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("EnableUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId)) + "/actions/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnableUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnableUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Enables an Employee Identity and Access Management (EIAM) account.
//
// @return EnableUserResponse
func (client *Client) EnableUser(instanceId *string, applicationId *string, userId *string) (_result *EnableUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EnableUserHeaders{}
	_result = &EnableUserResponse{}
	_body, _err := client.EnableUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a device code.
//
// @param request - GenerateDeviceCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDeviceCodeResponse
func (client *Client) GenerateDeviceCodeWithOptions(instanceId *string, applicationId *string, request *GenerateDeviceCodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateDeviceCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateDeviceCode"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/oauth2/device/code"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateDeviceCodeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateDeviceCodeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Generates a device code.
//
// @param request - GenerateDeviceCodeRequest
//
// @return GenerateDeviceCodeResponse
func (client *Client) GenerateDeviceCode(instanceId *string, applicationId *string, request *GenerateDeviceCodeRequest) (_result *GenerateDeviceCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateDeviceCodeResponse{}
	_body, _err := client.GenerateDeviceCodeWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a token for accessing an application in an instance.
//
// Description:
//
// >
//
// 	- The following authorization types are supported: authorization code, device code, refresh token, and client credentials.
//
// @param request - GenerateTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTokenResponse
func (client *Client) GenerateTokenWithOptions(instanceId *string, applicationId *string, request *GenerateTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["client_id"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSecret)) {
		query["client_secret"] = request.ClientSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.CodeVerifier)) {
		query["code_verifier"] = request.CodeVerifier
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		query["device_code"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExclusiveTag)) {
		query["exclusive_tag"] = request.ExclusiveTag
	}

	if !tea.BoolValue(util.IsUnset(request.GrantType)) {
		query["grant_type"] = request.GrantType
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectUri)) {
		query["redirect_uri"] = request.RedirectUri
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshToken)) {
		query["refresh_token"] = request.RefreshToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateToken"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/oauth2/token"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateTokenResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateTokenResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Generates a token for accessing an application in an instance.
//
// Description:
//
// >
//
// 	- The following authorization types are supported: authorization code, device code, refresh token, and client credentials.
//
// @param request - GenerateTokenRequest
//
// @return GenerateTokenResponse
func (client *Client) GenerateToken(instanceId *string, applicationId *string, request *GenerateTokenRequest) (_result *GenerateTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateTokenResponse{}
	_body, _err := client.GenerateTokenWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the synchronization scope of an application in an instance.
//
// Description:
//
// >
//
// 	- You can go to the Applications page in the IDaaS console to set the synchronization scope. After an application is created, the application has the permission to call this operation by default.
//
// @param headers - GetApplicationProvisioningScopeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisioningScopeResponse
func (client *Client) GetApplicationProvisioningScopeWithOptions(instanceId *string, applicationId *string, headers *GetApplicationProvisioningScopeHeaders, runtime *util.RuntimeOptions) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationProvisioningScope"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/provisioningScope"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetApplicationProvisioningScopeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetApplicationProvisioningScopeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the synchronization scope of an application in an instance.
//
// Description:
//
// >
//
// 	- You can go to the Applications page in the IDaaS console to set the synchronization scope. After an application is created, the application has the permission to call this operation by default.
//
// @return GetApplicationProvisioningScopeResponse
func (client *Client) GetApplicationProvisioningScope(instanceId *string, applicationId *string) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetApplicationProvisioningScopeHeaders{}
	_result = &GetApplicationProvisioningScopeResponse{}
	_body, _err := client.GetApplicationProvisioningScopeWithOptions(instanceId, applicationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询一个EIAM组信息
//
// @param headers - GetGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(instanceId *string, applicationId *string, groupId *string, headers *GetGroupHeaders, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroup"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询一个EIAM组信息
//
// @return GetGroupResponse
func (client *Client) GetGroup(instanceId *string, applicationId *string, groupId *string) (_result *GetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGroupHeaders{}
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(instanceId, applicationId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an organizational unit.
//
// @param headers - GetOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationalUnitResponse
func (client *Client) GetOrganizationalUnitWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, headers *GetOrganizationalUnitHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationalUnitResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrganizationalUnit"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/organizationalUnits/" + tea.StringValue(openapiutil.GetEncodeParam(organizationalUnitId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetOrganizationalUnitResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetOrganizationalUnitResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information of an organizational unit.
//
// @return GetOrganizationalUnitResponse
func (client *Client) GetOrganizationalUnit(instanceId *string, applicationId *string, organizationalUnitId *string) (_result *GetOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrganizationalUnitHeaders{}
	_result = &GetOrganizationalUnitResponse{}
	_body, _err := client.GetOrganizationalUnitWithOptions(instanceId, applicationId, organizationalUnitId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the ID of an organizational unit based on the external ID
//
// @param request - GetOrganizationalUnitIdByExternalIdRequest
//
// @param headers - GetOrganizationalUnitIdByExternalIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationalUnitIdByExternalIdResponse
func (client *Client) GetOrganizationalUnitIdByExternalIdWithOptions(instanceId *string, applicationId *string, request *GetOrganizationalUnitIdByExternalIdRequest, headers *GetOrganizationalUnitIdByExternalIdHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationalUnitIdByExternalIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitExternalId)) {
		body["organizationalUnitExternalId"] = request.OrganizationalUnitExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitSourceId)) {
		body["organizationalUnitSourceId"] = request.OrganizationalUnitSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitSourceType)) {
		body["organizationalUnitSourceType"] = request.OrganizationalUnitSourceType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrganizationalUnitIdByExternalId"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/organizationalUnits/_/actions/getOrganizationalUnitIdByExternalId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetOrganizationalUnitIdByExternalIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetOrganizationalUnitIdByExternalIdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the ID of an organizational unit based on the external ID
//
// @param request - GetOrganizationalUnitIdByExternalIdRequest
//
// @return GetOrganizationalUnitIdByExternalIdResponse
func (client *Client) GetOrganizationalUnitIdByExternalId(instanceId *string, applicationId *string, request *GetOrganizationalUnitIdByExternalIdRequest) (_result *GetOrganizationalUnitIdByExternalIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrganizationalUnitIdByExternalIdHeaders{}
	_result = &GetOrganizationalUnitIdByExternalIdResponse{}
	_body, _err := client.GetOrganizationalUnitIdByExternalIdWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an Employee Identity and Access Management (EIAM) account.
//
// @param headers - GetUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *GetUserHeaders, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information of an Employee Identity and Access Management (EIAM) account.
//
// @return GetUserResponse
func (client *Client) GetUser(instanceId *string, applicationId *string, userId *string) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account by email address.
//
// @param request - GetUserIdByEmailRequest
//
// @param headers - GetUserIdByEmailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdByEmailResponse
func (client *Client) GetUserIdByEmailWithOptions(instanceId *string, applicationId *string, request *GetUserIdByEmailRequest, headers *GetUserIdByEmailHeaders, runtime *util.RuntimeOptions) (_result *GetUserIdByEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserIdByEmail"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/_/actions/getUserIdByEmail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetUserIdByEmailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetUserIdByEmailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account by email address.
//
// @param request - GetUserIdByEmailRequest
//
// @return GetUserIdByEmailResponse
func (client *Client) GetUserIdByEmail(instanceId *string, applicationId *string, request *GetUserIdByEmailRequest) (_result *GetUserIdByEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserIdByEmailHeaders{}
	_result = &GetUserIdByEmailResponse{}
	_body, _err := client.GetUserIdByEmailWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the mobile number.
//
// @param request - GetUserIdByPhoneNumberRequest
//
// @param headers - GetUserIdByPhoneNumberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdByPhoneNumberResponse
func (client *Client) GetUserIdByPhoneNumberWithOptions(instanceId *string, applicationId *string, request *GetUserIdByPhoneNumberRequest, headers *GetUserIdByPhoneNumberHeaders, runtime *util.RuntimeOptions) (_result *GetUserIdByPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["phoneNumber"] = request.PhoneNumber
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserIdByPhoneNumber"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/_/actions/getUserIdByPhoneNumber"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetUserIdByPhoneNumberResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetUserIdByPhoneNumberResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the mobile number.
//
// @param request - GetUserIdByPhoneNumberRequest
//
// @return GetUserIdByPhoneNumberResponse
func (client *Client) GetUserIdByPhoneNumber(instanceId *string, applicationId *string, request *GetUserIdByPhoneNumberRequest) (_result *GetUserIdByPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserIdByPhoneNumberHeaders{}
	_result = &GetUserIdByPhoneNumberResponse{}
	_body, _err := client.GetUserIdByPhoneNumberWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the external ID.
//
// @param request - GetUserIdByUserExternalIdRequest
//
// @param headers - GetUserIdByUserExternalIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdByUserExternalIdResponse
func (client *Client) GetUserIdByUserExternalIdWithOptions(instanceId *string, applicationId *string, request *GetUserIdByUserExternalIdRequest, headers *GetUserIdByUserExternalIdHeaders, runtime *util.RuntimeOptions) (_result *GetUserIdByUserExternalIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserExternalId)) {
		body["userExternalId"] = request.UserExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSourceId)) {
		body["userSourceId"] = request.UserSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSourceType)) {
		body["userSourceType"] = request.UserSourceType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserIdByUserExternalId"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/_/actions/getUserIdByExternalId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetUserIdByUserExternalIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetUserIdByUserExternalIdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the external ID.
//
// @param request - GetUserIdByUserExternalIdRequest
//
// @return GetUserIdByUserExternalIdResponse
func (client *Client) GetUserIdByUserExternalId(instanceId *string, applicationId *string, request *GetUserIdByUserExternalIdRequest) (_result *GetUserIdByUserExternalIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserIdByUserExternalIdHeaders{}
	_result = &GetUserIdByUserExternalIdResponse{}
	_body, _err := client.GetUserIdByUserExternalIdWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the username.
//
// @param request - GetUserIdByUsernameRequest
//
// @param headers - GetUserIdByUsernameHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdByUsernameResponse
func (client *Client) GetUserIdByUsernameWithOptions(instanceId *string, applicationId *string, request *GetUserIdByUsernameRequest, headers *GetUserIdByUsernameHeaders, runtime *util.RuntimeOptions) (_result *GetUserIdByUsernameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserIdByUsername"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/_/actions/getUserIdByUsername"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetUserIdByUsernameResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetUserIdByUsernameResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the username.
//
// @param request - GetUserIdByUsernameRequest
//
// @return GetUserIdByUsernameResponse
func (client *Client) GetUserIdByUsername(instanceId *string, applicationId *string, request *GetUserIdByUsernameRequest) (_result *GetUserIdByUsernameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserIdByUsernameHeaders{}
	_result = &GetUserIdByUsernameResponse{}
	_body, _err := client.GetUserIdByUsernameWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a user by using the user token.
//
// @param headers - GetUserInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInfoResponse
func (client *Client) GetUserInfoWithOptions(instanceId *string, applicationId *string, headers *GetUserInfoHeaders, runtime *util.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserInfo"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/oauth2/userinfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetUserInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetUserInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information of a user by using the user token.
//
// @return GetUserInfoResponse
func (client *Client) GetUserInfo(instanceId *string, applicationId *string) (_result *GetUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserInfoHeaders{}
	_result = &GetUserInfoResponse{}
	_body, _err := client.GetUserInfoWithOptions(instanceId, applicationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询组列表
//
// @param request - ListGroupsRequest
//
// @param headers - ListGroupsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(instanceId *string, applicationId *string, request *ListGroupsRequest, headers *ListGroupsHeaders, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupNameStartWith)) {
		query["groupNameStartWith"] = request.GroupNameStartWith
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询组列表
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
func (client *Client) ListGroups(instanceId *string, applicationId *string, request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListGroupsHeaders{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取账户关联组列表
//
// @param request - ListGroupsForUserRequest
//
// @param headers - ListGroupsForUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUserWithOptions(instanceId *string, applicationId *string, userId *string, request *ListGroupsForUserRequest, headers *ListGroupsForUserHeaders, runtime *util.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupsForUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId)) + "/actions/listGroupsForUser"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListGroupsForUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListGroupsForUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取账户关联组列表
//
// @param request - ListGroupsForUserRequest
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUser(instanceId *string, applicationId *string, userId *string, request *ListGroupsForUserRequest) (_result *ListGroupsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListGroupsForUserHeaders{}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.ListGroupsForUserWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of all the parent organizational units of an organizational unit.
//
// @param headers - ListOrganizationalUnitParentIdsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitParentIdsResponse
func (client *Client) ListOrganizationalUnitParentIdsWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, headers *ListOrganizationalUnitParentIdsHeaders, runtime *util.RuntimeOptions) (_result *ListOrganizationalUnitParentIdsResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationalUnitParentIds"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/organizationalUnits/" + tea.StringValue(openapiutil.GetEncodeParam(organizationalUnitId)) + "/parentIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListOrganizationalUnitParentIdsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListOrganizationalUnitParentIdsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information of all the parent organizational units of an organizational unit.
//
// @return ListOrganizationalUnitParentIdsResponse
func (client *Client) ListOrganizationalUnitParentIds(instanceId *string, applicationId *string, organizationalUnitId *string) (_result *ListOrganizationalUnitParentIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOrganizationalUnitParentIdsHeaders{}
	_result = &ListOrganizationalUnitParentIdsResponse{}
	_body, _err := client.ListOrganizationalUnitParentIdsWithOptions(instanceId, applicationId, organizationalUnitId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of Employee Identity and Access Management (EIAM) organizational units by page.
//
// @param request - ListOrganizationalUnitsRequest
//
// @param headers - ListOrganizationalUnitsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitsResponse
func (client *Client) ListOrganizationalUnitsWithOptions(instanceId *string, applicationId *string, request *ListOrganizationalUnitsRequest, headers *ListOrganizationalUnitsHeaders, runtime *util.RuntimeOptions) (_result *ListOrganizationalUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["parentId"] = request.ParentId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationalUnits"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/organizationalUnits"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListOrganizationalUnitsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListOrganizationalUnitsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information of Employee Identity and Access Management (EIAM) organizational units by page.
//
// @param request - ListOrganizationalUnitsRequest
//
// @return ListOrganizationalUnitsResponse
func (client *Client) ListOrganizationalUnits(instanceId *string, applicationId *string, request *ListOrganizationalUnitsRequest) (_result *ListOrganizationalUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOrganizationalUnitsHeaders{}
	_result = &ListOrganizationalUnitsResponse{}
	_body, _err := client.ListOrganizationalUnitsWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of Employee Identity and Access Management (EIAM) accounts by page.
//
// @param request - ListUsersRequest
//
// @param headers - ListUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(instanceId *string, applicationId *string, request *ListUsersRequest, headers *ListUsersHeaders, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["organizationalUnitId"] = request.OrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListUsersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListUsersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information of Employee Identity and Access Management (EIAM) accounts by page.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(instanceId *string, applicationId *string, request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUsersHeaders{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定组下账户IDS
//
// @param request - ListUsersForGroupRequest
//
// @param headers - ListUsersForGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroupWithOptions(instanceId *string, applicationId *string, groupId *string, request *ListUsersForGroupRequest, headers *ListUsersForGroupHeaders, runtime *util.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsersForGroup"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/actions/listUsersForGroup"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListUsersForGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListUsersForGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询指定组下账户IDS
//
// @param request - ListUsersForGroupRequest
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroup(instanceId *string, applicationId *string, groupId *string, request *ListUsersForGroupRequest) (_result *ListUsersForGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUsersForGroupHeaders{}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.ListUsersForGroupWithOptions(instanceId, applicationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新组信息
//
// @param request - PatchGroupRequest
//
// @param headers - PatchGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchGroupResponse
func (client *Client) PatchGroupWithOptions(instanceId *string, applicationId *string, groupId *string, request *PatchGroupRequest, headers *PatchGroupHeaders, runtime *util.RuntimeOptions) (_result *PatchGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PatchGroup"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PatchGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PatchGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新组信息
//
// @param request - PatchGroupRequest
//
// @return PatchGroupResponse
func (client *Client) PatchGroup(instanceId *string, applicationId *string, groupId *string, request *PatchGroupRequest) (_result *PatchGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PatchGroupHeaders{}
	_result = &PatchGroupResponse{}
	_body, _err := client.PatchGroupWithOptions(instanceId, applicationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an EIAM organizational unit.
//
// Description:
//
// The operation conforms to the HTTP PATCH request method. The value of a parameter is modified only if the parameter is specified in the request.
//
// @param request - PatchOrganizationalUnitRequest
//
// @param headers - PatchOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchOrganizationalUnitResponse
func (client *Client) PatchOrganizationalUnitWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, request *PatchOrganizationalUnitRequest, headers *PatchOrganizationalUnitHeaders, runtime *util.RuntimeOptions) (_result *PatchOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitName)) {
		body["organizationalUnitName"] = request.OrganizationalUnitName
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PatchOrganizationalUnit"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/organizationalUnits/" + tea.StringValue(openapiutil.GetEncodeParam(organizationalUnitId))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PatchOrganizationalUnitResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PatchOrganizationalUnitResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies an EIAM organizational unit.
//
// Description:
//
// The operation conforms to the HTTP PATCH request method. The value of a parameter is modified only if the parameter is specified in the request.
//
// @param request - PatchOrganizationalUnitRequest
//
// @return PatchOrganizationalUnitResponse
func (client *Client) PatchOrganizationalUnit(instanceId *string, applicationId *string, organizationalUnitId *string, request *PatchOrganizationalUnitRequest) (_result *PatchOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PatchOrganizationalUnitHeaders{}
	_result = &PatchOrganizationalUnitResponse{}
	_body, _err := client.PatchOrganizationalUnitWithOptions(instanceId, applicationId, organizationalUnitId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an Employee Identity and Access Management (EIAM) account.
//
// Description:
//
// The operation conforms to the HTTP PATCH request method. The value of a parameter is modified only if the parameter is specified in the request.
//
// @param request - PatchUserRequest
//
// @param headers - PatchUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchUserResponse
func (client *Client) PatchUserWithOptions(instanceId *string, applicationId *string, userId *string, request *PatchUserRequest, headers *PatchUserHeaders, runtime *util.RuntimeOptions) (_result *PatchUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		body["customFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EmailVerified)) {
		body["emailVerified"] = request.EmailVerified
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		body["phoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberVerified)) {
		body["phoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneRegion)) {
		body["phoneRegion"] = request.PhoneRegion
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PatchUser"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PatchUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PatchUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies an Employee Identity and Access Management (EIAM) account.
//
// Description:
//
// The operation conforms to the HTTP PATCH request method. The value of a parameter is modified only if the parameter is specified in the request.
//
// @param request - PatchUserRequest
//
// @return PatchUserResponse
func (client *Client) PatchUser(instanceId *string, applicationId *string, userId *string, request *PatchUserRequest) (_result *PatchUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PatchUserHeaders{}
	_result = &PatchUserResponse{}
	_body, _err := client.PatchUserWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将账户从多个组织移除【不支持移除主组织】
//
// @param request - RemoveUserFromOrganizationalUnitsRequest
//
// @param headers - RemoveUserFromOrganizationalUnitsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromOrganizationalUnitsResponse
func (client *Client) RemoveUserFromOrganizationalUnitsWithOptions(instanceId *string, applicationId *string, userId *string, request *RemoveUserFromOrganizationalUnitsRequest, headers *RemoveUserFromOrganizationalUnitsHeaders, runtime *util.RuntimeOptions) (_result *RemoveUserFromOrganizationalUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		body["organizationalUnitIds"] = request.OrganizationalUnitIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserFromOrganizationalUnits"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId)) + "/actions/removeUserFromOrganizationalUnits"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RemoveUserFromOrganizationalUnitsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RemoveUserFromOrganizationalUnitsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 将账户从多个组织移除【不支持移除主组织】
//
// @param request - RemoveUserFromOrganizationalUnitsRequest
//
// @return RemoveUserFromOrganizationalUnitsResponse
func (client *Client) RemoveUserFromOrganizationalUnits(instanceId *string, applicationId *string, userId *string, request *RemoveUserFromOrganizationalUnitsRequest) (_result *RemoveUserFromOrganizationalUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveUserFromOrganizationalUnitsHeaders{}
	_result = &RemoveUserFromOrganizationalUnitsResponse{}
	_body, _err := client.RemoveUserFromOrganizationalUnitsWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为指定组批量解除账户关联
//
// @param request - RemoveUsersFromGroupRequest
//
// @param headers - RemoveUsersFromGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroupWithOptions(instanceId *string, applicationId *string, groupId *string, request *RemoveUsersFromGroupRequest, headers *RemoveUsersFromGroupHeaders, runtime *util.RuntimeOptions) (_result *RemoveUsersFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUsersFromGroup"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/actions/removeUsersFromGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RemoveUsersFromGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RemoveUsersFromGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 为指定组批量解除账户关联
//
// @param request - RemoveUsersFromGroupRequest
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroup(instanceId *string, applicationId *string, groupId *string, request *RemoveUsersFromGroupRequest) (_result *RemoveUsersFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveUsersFromGroupHeaders{}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.RemoveUsersFromGroupWithOptions(instanceId, applicationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes an access token or refresh token.
//
// @param request - RevokeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeTokenResponse
func (client *Client) RevokeTokenWithOptions(instanceId *string, applicationId *string, request *RevokeTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RevokeTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["client_id"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSecret)) {
		query["client_secret"] = request.ClientSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.TokenTypeHint)) {
		query["token_type_hint"] = request.TokenTypeHint
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeToken"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/oauth2/revoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RevokeTokenResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RevokeTokenResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Revokes an access token or refresh token.
//
// @param request - RevokeTokenRequest
//
// @return RevokeTokenResponse
func (client *Client) RevokeToken(instanceId *string, applicationId *string, request *RevokeTokenRequest) (_result *RevokeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeTokenResponse{}
	_body, _err := client.RevokeTokenWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将指定组织设置为账户主组织，移除旧主组织，加入新主组织。
//
// @param request - SetUserPrimaryOrganizationalUnitRequest
//
// @param headers - SetUserPrimaryOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserPrimaryOrganizationalUnitResponse
func (client *Client) SetUserPrimaryOrganizationalUnitWithOptions(instanceId *string, applicationId *string, userId *string, request *SetUserPrimaryOrganizationalUnitRequest, headers *SetUserPrimaryOrganizationalUnitHeaders, runtime *util.RuntimeOptions) (_result *SetUserPrimaryOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		body["organizationalUnitId"] = request.OrganizationalUnitId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetUserPrimaryOrganizationalUnit"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId)) + "/actions/setUserPrimaryOrganizationalUnit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SetUserPrimaryOrganizationalUnitResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SetUserPrimaryOrganizationalUnitResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 将指定组织设置为账户主组织，移除旧主组织，加入新主组织。
//
// @param request - SetUserPrimaryOrganizationalUnitRequest
//
// @return SetUserPrimaryOrganizationalUnitResponse
func (client *Client) SetUserPrimaryOrganizationalUnit(instanceId *string, applicationId *string, userId *string, request *SetUserPrimaryOrganizationalUnitRequest) (_result *SetUserPrimaryOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetUserPrimaryOrganizationalUnitHeaders{}
	_result = &SetUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.SetUserPrimaryOrganizationalUnitWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新账户密码
//
// @param request - UpdateUserPasswordRequest
//
// @param headers - UpdateUserPasswordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPasswordResponse
func (client *Client) UpdateUserPasswordWithOptions(instanceId *string, applicationId *string, userId *string, request *UpdateUserPasswordRequest, headers *UpdateUserPasswordHeaders, runtime *util.RuntimeOptions) (_result *UpdateUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["password"] = request.Password
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserPassword"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(applicationId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(userId)) + "/actions/updateUserPassword"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateUserPasswordResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateUserPasswordResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新账户密码
//
// @param request - UpdateUserPasswordRequest
//
// @return UpdateUserPasswordResponse
func (client *Client) UpdateUserPassword(instanceId *string, applicationId *string, userId *string, request *UpdateUserPasswordRequest) (_result *UpdateUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUserPasswordHeaders{}
	_result = &UpdateUserPasswordResponse{}
	_body, _err := client.UpdateUserPasswordWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
