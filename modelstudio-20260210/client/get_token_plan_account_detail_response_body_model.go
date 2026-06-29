// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenPlanAccountDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTokenPlanAccountDetailResponseBody
	GetCode() *string
	SetData(v *GetTokenPlanAccountDetailResponseBodyData) *GetTokenPlanAccountDetailResponseBody
	GetData() *GetTokenPlanAccountDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetTokenPlanAccountDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTokenPlanAccountDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTokenPlanAccountDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTokenPlanAccountDetailResponseBody
	GetSuccess() *bool
}

type GetTokenPlanAccountDetailResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *GetTokenPlanAccountDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// C0DC05D9-C506-519B-AFF3-2B00165176E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTokenPlanAccountDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanAccountDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenPlanAccountDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTokenPlanAccountDetailResponseBody) GetData() *GetTokenPlanAccountDetailResponseBodyData {
	return s.Data
}

func (s *GetTokenPlanAccountDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTokenPlanAccountDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTokenPlanAccountDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenPlanAccountDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTokenPlanAccountDetailResponseBody) SetCode(v string) *GetTokenPlanAccountDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBody) SetData(v *GetTokenPlanAccountDetailResponseBodyData) *GetTokenPlanAccountDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBody) SetHttpStatusCode(v int32) *GetTokenPlanAccountDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBody) SetMessage(v string) *GetTokenPlanAccountDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBody) SetRequestId(v string) *GetTokenPlanAccountDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBody) SetSuccess(v bool) *GetTokenPlanAccountDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTokenPlanAccountDetailResponseBodyData struct {
	// The account ID.
	//
	// example:
	//
	// acc_123456789
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The account type. Valid values:
	//
	// - ALIYUN
	//
	// - SSO
	//
	// - SA
	//
	// example:
	//
	// ALIYUN
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The Alibaba Cloud UID. This parameter applies only to accounts of the ALIYUN type.
	//
	// example:
	//
	// 1122334455
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The time when the account was created.
	//
	// example:
	//
	// Thu May 28 14:33:52 CST 2026
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The email address. This parameter applies only to accounts of the SSO type.
	//
	// example:
	//
	// test@email.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The global status of the account. A value of 0 indicates Normal. A value of 1 indicates Frozen.
	//
	// example:
	//
	// 0
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of organization memberships in a tree structure (organization → workspace).
	OrgMemberships []*GetTokenPlanAccountDetailResponseBodyDataOrgMemberships `json:"OrgMemberships,omitempty" xml:"OrgMemberships,omitempty" type:"Repeated"`
}

func (s GetTokenPlanAccountDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanAccountDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenPlanAccountDetailResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *GetTokenPlanAccountDetailResponseBodyData) GetAccountType() *string {
	return s.AccountType
}

func (s *GetTokenPlanAccountDetailResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *GetTokenPlanAccountDetailResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetTokenPlanAccountDetailResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *GetTokenPlanAccountDetailResponseBodyData) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *GetTokenPlanAccountDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTokenPlanAccountDetailResponseBodyData) GetOrgMemberships() []*GetTokenPlanAccountDetailResponseBodyDataOrgMemberships {
	return s.OrgMemberships
}

func (s *GetTokenPlanAccountDetailResponseBodyData) SetAccountId(v string) *GetTokenPlanAccountDetailResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyData) SetAccountType(v string) *GetTokenPlanAccountDetailResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyData) SetAliyunUid(v string) *GetTokenPlanAccountDetailResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyData) SetCreatedAt(v string) *GetTokenPlanAccountDetailResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyData) SetEmail(v string) *GetTokenPlanAccountDetailResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyData) SetIsDeleted(v bool) *GetTokenPlanAccountDetailResponseBodyData {
	s.IsDeleted = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyData) SetName(v string) *GetTokenPlanAccountDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyData) SetOrgMemberships(v []*GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) *GetTokenPlanAccountDetailResponseBodyData {
	s.OrgMemberships = v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyData) Validate() error {
	if s.OrgMemberships != nil {
		for _, item := range s.OrgMemberships {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTokenPlanAccountDetailResponseBodyDataOrgMemberships struct {
	// The organization member status. Valid values:
	//
	// - ACTIVE
	//
	// - INITIAL
	//
	// - FROZEN
	//
	// example:
	//
	// ENABLE
	MemberStatus *string `json:"MemberStatus,omitempty" xml:"MemberStatus,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// org_123456789
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organization role code. Valid values:
	//
	// - ORG_OWNER
	//
	// - ORG_ADMIN
	//
	// - ORG_MEMBER
	//
	// example:
	//
	// ORG_MEMBER
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// The organization role ID.
	//
	// example:
	//
	// SYSTEM_ROLE_ORG_OWNER
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The list of workspaces that the account has joined under the organization.
	Workspaces []*GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) GoString() string {
	return s.String()
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) GetMemberStatus() *string {
	return s.MemberStatus
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) GetOrgId() *string {
	return s.OrgId
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) GetRoleCode() *string {
	return s.RoleCode
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) GetRoleId() *string {
	return s.RoleId
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) GetWorkspaces() []*GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces {
	return s.Workspaces
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) SetMemberStatus(v string) *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships {
	s.MemberStatus = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) SetOrgId(v string) *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships {
	s.OrgId = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) SetRoleCode(v string) *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships {
	s.RoleCode = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) SetRoleId(v string) *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships {
	s.RoleId = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) SetWorkspaces(v []*GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships {
	s.Workspaces = v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMemberships) Validate() error {
	if s.Workspaces != nil {
		for _, item := range s.Workspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces struct {
	// The member status. Valid values:
	//
	// - ACTIVE
	//
	// - FROZEN
	//
	// example:
	//
	// ACTIVE
	MemberStatus *string `json:"MemberStatus,omitempty" xml:"MemberStatus,omitempty"`
	// The workspace role code. Valid values:
	//
	// - WS_ADMIN
	//
	// - WS_MEMBER
	//
	// example:
	//
	// WS_ADMIN
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// The workspace role ID.
	//
	// example:
	//
	// SYSTEM_ROLE_WS_ADMIN
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws_123456789
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) GoString() string {
	return s.String()
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) GetMemberStatus() *string {
	return s.MemberStatus
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) GetRoleCode() *string {
	return s.RoleCode
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) GetRoleId() *string {
	return s.RoleId
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) SetMemberStatus(v string) *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces {
	s.MemberStatus = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) SetRoleCode(v string) *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces {
	s.RoleCode = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) SetRoleId(v string) *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces {
	s.RoleId = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) SetWorkspaceId(v string) *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponseBodyDataOrgMembershipsWorkspaces) Validate() error {
	return dara.Validate(s)
}
