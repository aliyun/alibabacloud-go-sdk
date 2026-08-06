// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWorkspacesResponseBody
	GetRequestId() *string
	SetResourceLimits(v map[string]interface{}) *ListWorkspacesResponseBody
	GetResourceLimits() map[string]interface{}
	SetTotalCount(v int64) *ListWorkspacesResponseBody
	GetTotalCount() *int64
	SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody
	GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces
}

type ListWorkspacesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D7B2E70-F770-505B-A672-09F1D8F2EC1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource types and quantity limits that a user can activate within a workspace. This list is returned when Option is set to GetResourceLimits.
	//
	// Currently supported resource types include:
	//
	// 	- MaxCompute_share: MaxCompute pay-as-you-go.
	//
	// 	- MaxCompute_isolate: MaxCompute subscription.
	//
	// 	- DLC_share: DLC pay-as-you-go.
	//
	// 	- PAI_isolate: PAI subscription.
	//
	// 	- PAI_share: PAI pay-as-you-go.
	//
	// 	- DataWorks_isolate: DataWorks subscription.
	//
	// 	- DataWorks_share: DataWorks pay-as-you-go.
	//
	// example:
	//
	// {
	//
	//    "MaxCompute_share": 1,
	//
	//    "MaxCompute_isolate": 1,
	//
	//    "DLC_share": 1
	//
	// }
	ResourceLimits map[string]interface{} `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty"`
	// The total number of workspaces that match the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of workspace details. This list is returned when Option is set to GetWorkspaces.
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResponseBody) GetResourceLimits() map[string]interface{} {
	return s.ResourceLimits
}

func (s *ListWorkspacesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWorkspacesResponseBody) GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces {
	return s.Workspaces
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetResourceLimits(v map[string]interface{}) *ListWorkspacesResponseBody {
	s.ResourceLimits = v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int64) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

func (s *ListWorkspacesResponseBody) Validate() error {
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

type ListWorkspacesResponseBodyWorkspaces struct {
	// The list of administrator account names.
	AdminNames []*string `json:"AdminNames,omitempty" xml:"AdminNames,omitempty" type:"Repeated"`
	// The user ID of the creator.
	//
	// example:
	//
	// 122424353535
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The workspace description.
	//
	// example:
	//
	// workspace description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the workspace.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The list of environments contained in the workspace.
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// The extended information. Currently includes TenantId, which indicates the tenant ID.
	//
	// example:
	//
	// {"TenantId": "4286******98"}
	ExtraInfos map[string]interface{} `json:"ExtraInfos,omitempty" xml:"ExtraInfos,omitempty"`
	// The time when the workspace was created. The time is in the ISO 8601 standard in UTC+0. Format: yyyy-MM-ddTHH:mm:ss.SSSZ.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the workspace was last modified. The time is in the ISO 8601 standard in UTC+0. Format: yyyy-MM-ddTHH:mmZ.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Indicates whether the workspace is the default workspace.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The workspace status.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// workspace-example
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmwp7rky****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetAdminNames() []*string {
	return s.AdminNames
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetCreator() *string {
	return s.Creator
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetDescription() *string {
	return s.Description
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetEnvTypes() []*string {
	return s.EnvTypes
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetExtraInfos() map[string]interface{} {
	return s.ExtraInfos
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetStatus() *string {
	return s.Status
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetAdminNames(v []*string) *ListWorkspacesResponseBodyWorkspaces {
	s.AdminNames = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreator(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Creator = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDescription(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDisplayName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.DisplayName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetEnvTypes(v []*string) *ListWorkspacesResponseBodyWorkspaces {
	s.EnvTypes = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetExtraInfos(v map[string]interface{}) *ListWorkspacesResponseBodyWorkspaces {
	s.ExtraInfos = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetGmtCreateTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.GmtCreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetGmtModifiedTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetIsDefault(v bool) *ListWorkspacesResponseBodyWorkspaces {
	s.IsDefault = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStatus(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Status = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetResourceGroupId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ResourceGroupId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) Validate() error {
	return dara.Validate(s)
}
