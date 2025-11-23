// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListWorkspacesResponseBodyData) *ListWorkspacesResponseBody
	GetData() *ListWorkspacesResponseBodyData
	SetErrorCode(v string) *ListWorkspacesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListWorkspacesResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListWorkspacesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkspacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkspacesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListWorkspacesResponseBody
	GetTotalCount() *int64
}

type ListWorkspacesResponseBody struct {
	// The dataset.
	Data *ListWorkspacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UserNotExist
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The total number of workspaces that meet the condition, which is the same as the TotalCount parameter.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// NextToken does not take effect.
	//
	// example:
	//
	// token-xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE214ECD-4330-503A-82F0-FFB03975****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**: The request succeeded.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of workspaces that meet the conditions.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) GetData() *ListWorkspacesResponseBodyData {
	return s.Data
}

func (s *ListWorkspacesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkspacesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListWorkspacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkspacesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWorkspacesResponseBody) SetData(v *ListWorkspacesResponseBodyData) *ListWorkspacesResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkspacesResponseBody) SetErrorCode(v string) *ListWorkspacesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetErrorMessage(v string) *ListWorkspacesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetSuccess(v bool) *ListWorkspacesResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int64) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkspacesResponseBodyData struct {
	BaseWorkspaces []*ListWorkspacesResponseBodyDataBaseWorkspaces `json:"BaseWorkspaces,omitempty" xml:"BaseWorkspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyData) GetBaseWorkspaces() []*ListWorkspacesResponseBodyDataBaseWorkspaces {
	return s.BaseWorkspaces
}

func (s *ListWorkspacesResponseBodyData) SetBaseWorkspaces(v []*ListWorkspacesResponseBodyDataBaseWorkspaces) *ListWorkspacesResponseBodyData {
	s.BaseWorkspaces = v
	return s
}

func (s *ListWorkspacesResponseBodyData) Validate() error {
	if s.BaseWorkspaces != nil {
		for _, item := range s.BaseWorkspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspacesResponseBodyDataBaseWorkspaces struct {
	// Whether the current user has joined the workspace.
	//
	// example:
	//
	// true
	AlreadyJoined *bool `json:"AlreadyJoined,omitempty" xml:"AlreadyJoined,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// 123
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The nickname of the creator.
	//
	// example:
	//
	// work*****
	CreatorNickName *string `json:"CreatorNickName,omitempty" xml:"CreatorNickName,omitempty"`
	// The Alibaba Cloud account UID of the creator.
	//
	// example:
	//
	// 1344****
	CreatorUid *string `json:"CreatorUid,omitempty" xml:"CreatorUid,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-01-01 00:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-01-01 00:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// workspace-xxxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner ID.
	//
	// example:
	//
	// 123****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The nickname of the owner.
	//
	// example:
	//
	// hel****
	OwnerNickName *string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty"`
	// The Alibaba Cloud UID of the owner.
	//
	// example:
	//
	// 15608564799****
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the service account.
	//
	// example:
	//
	// 12345
	ServiceAccountId *int64 `json:"ServiceAccountId,omitempty" xml:"ServiceAccountId,omitempty"`
	// The nickname of the service account.
	//
	// example:
	//
	// testname
	ServiceAccountNickName *string `json:"ServiceAccountNickName,omitempty" xml:"ServiceAccountNickName,omitempty"`
	// The Alibaba Cloud account UID of the service account.
	//
	// example:
	//
	// 1422****
	ServiceAccountUid *string `json:"ServiceAccountUid,omitempty" xml:"ServiceAccountUid,omitempty"`
	// The ID of the tenant to which the workspace belongs.
	//
	// example:
	//
	// 23456
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-25fl3qjqb****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 8652340494****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// test-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListWorkspacesResponseBodyDataBaseWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyDataBaseWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetAlreadyJoined() *bool {
	return s.AlreadyJoined
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetCreatorNickName() *string {
	return s.CreatorNickName
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetCreatorUid() *string {
	return s.CreatorUid
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetDescription() *string {
	return s.Description
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetName() *string {
	return s.Name
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetOwnerNickName() *string {
	return s.OwnerNickName
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetOwnerUid() *string {
	return s.OwnerUid
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetRegion() *string {
	return s.Region
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetServiceAccountId() *int64 {
	return s.ServiceAccountId
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetServiceAccountNickName() *string {
	return s.ServiceAccountNickName
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetServiceAccountUid() *string {
	return s.ServiceAccountUid
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetVpcId() *string {
	return s.VpcId
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetAlreadyJoined(v bool) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.AlreadyJoined = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetCreatorId(v int64) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.CreatorId = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetCreatorNickName(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.CreatorNickName = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetCreatorUid(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.CreatorUid = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetDescription(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetGmtCreate(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetGmtModified(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.GmtModified = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetName(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.Name = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetOwnerId(v int64) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.OwnerId = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetOwnerNickName(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.OwnerNickName = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetOwnerUid(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.OwnerUid = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetRegion(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.Region = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetServiceAccountId(v int64) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.ServiceAccountId = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetServiceAccountNickName(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.ServiceAccountNickName = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetServiceAccountUid(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.ServiceAccountUid = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetTenantId(v int64) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.TenantId = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetVpcId(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.VpcId = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetWorkspaceId(v int64) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) SetWorkspaceName(v string) *ListWorkspacesResponseBodyDataBaseWorkspaces {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesResponseBodyDataBaseWorkspaces) Validate() error {
	return dara.Validate(s)
}
