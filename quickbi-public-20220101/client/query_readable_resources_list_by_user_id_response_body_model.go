// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReadableResourcesListByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryReadableResourcesListByUserIdResponseBody
	GetRequestId() *string
	SetResult(v []*QueryReadableResourcesListByUserIdResponseBodyResult) *QueryReadableResourcesListByUserIdResponseBody
	GetResult() []*QueryReadableResourcesListByUserIdResponseBodyResult
	SetSuccess(v bool) *QueryReadableResourcesListByUserIdResponseBody
	GetSuccess() *bool
}

type QueryReadableResourcesListByUserIdResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of works that the user has permission to view.
	Result []*QueryReadableResourcesListByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryReadableResourcesListByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryReadableResourcesListByUserIdResponseBody) GetResult() []*QueryReadableResourcesListByUserIdResponseBodyResult {
	return s.Result
}

func (s *QueryReadableResourcesListByUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryReadableResourcesListByUserIdResponseBody) SetRequestId(v string) *QueryReadableResourcesListByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBody) SetResult(v []*QueryReadableResourcesListByUserIdResponseBodyResult) *QueryReadableResourcesListByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBody) SetSuccess(v bool) *QueryReadableResourcesListByUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBody) Validate() error {
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

type QueryReadableResourcesListByUserIdResponseBodyResult struct {
	// The timestamp of the creation time in milliseconds.
	//
	// example:
	//
	// 1611023338000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Remarks on the work.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory structure in which the work is located.
	Directory *QueryReadableResourcesListByUserIdResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The name of the Alibaba Cloud account to which the modifier belongs.
	//
	// example:
	//
	// Li Si
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// The timestamp of the modification time in milliseconds.
	//
	// example:
	//
	// 1611023338000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The Quick BI UserID of the work owner.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud account name of the owner.
	//
	// example:
	//
	// Tom
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Security policies for collaborative authorization of works. Valid values:
	//
	// 	- 0: private
	//
	// 	- 12: Authorize specified members
	//
	// 	- 1 or 11: Authorize all workspace members
	//
	// >
	//
	// 	- If you use legacy permissions, the return value is 1.
	//
	// 	- If you use the new permissions, the return value is 11.
	//
	// example:
	//
	// 0
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The status of the report. Valid values:
	//
	// 	- 0: unpublished
	//
	// 	- 1: published
	//
	// 	- 2: modified but not published
	//
	// 	- 3: unpublished
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Third-party embedding status. Valid values:
	//
	// 	- 0: The embed service is not enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 1
	ThirdPartAuthFlag *int32 `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// The name of the work.
	//
	// example:
	//
	// Company Region Table
	WorkName *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	// The type of the work. Valid values:
	//
	// 	- DATAPRODUCT: BI portal
	//
	// 	- PAGE: Dashboard
	//
	// 	- FULLPAGE: full-screen dashboards
	//
	// 	- REPORT: workbook
	//
	// example:
	//
	// PAGE
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 03366b16-69ce-43c8-b782-56c2f6ec****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// The ID of the workspace to which the work belongs.
	//
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// Test Workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryReadableResourcesListByUserIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetDirectory() *QueryReadableResourcesListByUserIdResponseBodyResultDirectory {
	return s.Directory
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetModifyName() *string {
	return s.ModifyName
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetThirdPartAuthFlag() *int32 {
	return s.ThirdPartAuthFlag
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetWorkName() *string {
	return s.WorkName
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetWorkType() *string {
	return s.WorkType
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetWorksId() *string {
	return s.WorksId
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetCreateTime(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetDescription(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.Description = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetDirectory(v *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.Directory = v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetModifyName(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.ModifyName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetModifyTime(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetOwnerId(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetOwnerName(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetSecurityLevel(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.SecurityLevel = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetStatus(v int32) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetThirdPartAuthFlag(v int32) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorkName(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorkName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorkType(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorkType = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorksId(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorkspaceId(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorkspaceName(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryReadableResourcesListByUserIdResponseBodyResultDirectory struct {
	// The ID of the directory.
	//
	// example:
	//
	// e4276ea5-b232-4fb1-8f0f-efcee4a2****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the directory.
	//
	// example:
	//
	// Test directory
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The hierarchical structure of the directory ID, which is separated with \\"/\\".
	//
	// example:
	//
	// e4276ea5-b232-4fb1-8f0f-efcee4a2****
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The hierarchical structure of the directory name, which is separated with \\"/\\".
	//
	// example:
	//
	// Test directory
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryReadableResourcesListByUserIdResponseBodyResultDirectory) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdResponseBodyResultDirectory) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) GetId() *string {
	return s.Id
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) GetName() *string {
	return s.Name
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) GetPathId() *string {
	return s.PathId
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) GetPathName() *string {
	return s.PathName
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) SetId(v string) *QueryReadableResourcesListByUserIdResponseBodyResultDirectory {
	s.Id = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) SetName(v string) *QueryReadableResourcesListByUserIdResponseBodyResultDirectory {
	s.Name = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) SetPathId(v string) *QueryReadableResourcesListByUserIdResponseBodyResultDirectory {
	s.PathId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) SetPathName(v string) *QueryReadableResourcesListByUserIdResponseBodyResultDirectory {
	s.PathName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) Validate() error {
	return dara.Validate(s)
}
