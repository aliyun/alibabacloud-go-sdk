// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgAuthorizedWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKgAuthorizedWorkspacesResponseBody
	GetCode() *string
	SetData(v *GetKgAuthorizedWorkspacesResponseBodyData) *GetKgAuthorizedWorkspacesResponseBody
	GetData() *GetKgAuthorizedWorkspacesResponseBodyData
	SetHttpStatusCode(v int32) *GetKgAuthorizedWorkspacesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetKgAuthorizedWorkspacesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKgAuthorizedWorkspacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKgAuthorizedWorkspacesResponseBody
	GetSuccess() *bool
}

type GetKgAuthorizedWorkspacesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request result.
	Data *GetKgAuthorizedWorkspacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKgAuthorizedWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKgAuthorizedWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *GetKgAuthorizedWorkspacesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKgAuthorizedWorkspacesResponseBody) GetData() *GetKgAuthorizedWorkspacesResponseBodyData {
	return s.Data
}

func (s *GetKgAuthorizedWorkspacesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetKgAuthorizedWorkspacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKgAuthorizedWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKgAuthorizedWorkspacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKgAuthorizedWorkspacesResponseBody) SetCode(v string) *GetKgAuthorizedWorkspacesResponseBody {
	s.Code = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBody) SetData(v *GetKgAuthorizedWorkspacesResponseBodyData) *GetKgAuthorizedWorkspacesResponseBody {
	s.Data = v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBody) SetHttpStatusCode(v int32) *GetKgAuthorizedWorkspacesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBody) SetMessage(v string) *GetKgAuthorizedWorkspacesResponseBody {
	s.Message = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBody) SetRequestId(v string) *GetKgAuthorizedWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBody) SetSuccess(v bool) *GetKgAuthorizedWorkspacesResponseBody {
	s.Success = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKgAuthorizedWorkspacesResponseBodyData struct {
	// The total number of knowledge graph workspaces that the user has permissions on.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of knowledge graph workspaces that the user has permissions on.
	WorkspaceList []*GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList `json:"WorkspaceList,omitempty" xml:"WorkspaceList,omitempty" type:"Repeated"`
}

func (s GetKgAuthorizedWorkspacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKgAuthorizedWorkspacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKgAuthorizedWorkspacesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetKgAuthorizedWorkspacesResponseBodyData) GetWorkspaceList() []*GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList {
	return s.WorkspaceList
}

func (s *GetKgAuthorizedWorkspacesResponseBodyData) SetTotalCount(v int32) *GetKgAuthorizedWorkspacesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyData) SetWorkspaceList(v []*GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) *GetKgAuthorizedWorkspacesResponseBodyData {
	s.WorkspaceList = v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyData) Validate() error {
	if s.WorkspaceList != nil {
		for _, item := range s.WorkspaceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList struct {
	// The description of the knowledge graph workspace.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time of the knowledge graph workspace.
	//
	// example:
	//
	// 2026-08-25 12:34:56
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The latest publish time of the knowledge graph workspace. This value is empty if the workspace has never been published successfully.
	//
	// example:
	//
	// 2026-08-25 12:34:56
	LastPublishTime *string `json:"LastPublishTime,omitempty" xml:"LastPublishTime,omitempty"`
	// The latest publish version number of the knowledge graph workspace. This value is empty if the workspace has never been published successfully.
	//
	// example:
	//
	// 10
	LastPublishVersion *int32 `json:"LastPublishVersion,omitempty" xml:"LastPublishVersion,omitempty"`
	// The name of the knowledge graph workspace.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of roles assigned to the specified user in the workspace. This is an empty list if the user is not a member of the workspace.
	RoleList []*GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// The ID of the knowledge graph workspace.
	//
	// example:
	//
	// abc1011
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) String() string {
	return dara.Prettify(s)
}

func (s GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) GoString() string {
	return s.String()
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) GetDescription() *string {
	return s.Description
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) GetLastPublishTime() *string {
	return s.LastPublishTime
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) GetLastPublishVersion() *int32 {
	return s.LastPublishVersion
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) GetName() *string {
	return s.Name
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) GetRoleList() []*GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList {
	return s.RoleList
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) SetDescription(v string) *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList {
	s.Description = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) SetGmtCreate(v string) *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList {
	s.GmtCreate = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) SetLastPublishTime(v string) *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList {
	s.LastPublishTime = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) SetLastPublishVersion(v int32) *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList {
	s.LastPublishVersion = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) SetName(v string) *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList {
	s.Name = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) SetRoleList(v []*GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList) *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList {
	s.RoleList = v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) SetWorkspaceId(v string) *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList {
	s.WorkspaceId = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceList) Validate() error {
	if s.RoleList != nil {
		for _, item := range s.RoleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList struct {
	// The code of the workspace role.
	//
	// example:
	//
	// WORKSPACE_ADMIN
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the workspace role.
	//
	// example:
	//
	// Storage management
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList) String() string {
	return dara.Prettify(s)
}

func (s GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList) GoString() string {
	return s.String()
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList) GetCode() *string {
	return s.Code
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList) GetName() *string {
	return s.Name
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList) SetCode(v string) *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList {
	s.Code = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList) SetName(v string) *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList {
	s.Name = &v
	return s
}

func (s *GetKgAuthorizedWorkspacesResponseBodyDataWorkspaceListRoleList) Validate() error {
	return dara.Validate(s)
}
