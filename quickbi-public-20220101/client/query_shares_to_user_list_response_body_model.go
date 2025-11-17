// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySharesToUserListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QuerySharesToUserListResponseBody
	GetRequestId() *string
	SetResult(v []*QuerySharesToUserListResponseBodyResult) *QuerySharesToUserListResponseBody
	GetResult() []*QuerySharesToUserListResponseBodyResult
	SetSuccess(v bool) *QuerySharesToUserListResponseBody
	GetSuccess() *bool
}

type QuerySharesToUserListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns a list of works authorized to the user.
	Result []*QuerySharesToUserListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s QuerySharesToUserListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySharesToUserListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySharesToUserListResponseBody) GetResult() []*QuerySharesToUserListResponseBodyResult {
	return s.Result
}

func (s *QuerySharesToUserListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySharesToUserListResponseBody) SetRequestId(v string) *QuerySharesToUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySharesToUserListResponseBody) SetResult(v []*QuerySharesToUserListResponseBodyResult) *QuerySharesToUserListResponseBody {
	s.Result = v
	return s
}

func (s *QuerySharesToUserListResponseBody) SetSuccess(v bool) *QuerySharesToUserListResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySharesToUserListResponseBody) Validate() error {
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

type QuerySharesToUserListResponseBodyResult struct {
	// The timestamp of the creation time in milliseconds.
	//
	// example:
	//
	// 1530078690000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Remarks on the work.
	//
	// example:
	//
	// Description of the test report
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Information about the directory where the work is located.
	Directory *QuerySharesToUserListResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The name of the Alibaba Cloud account to which the modifier belongs.
	//
	// example:
	//
	// 13855265****@163.com
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// The timestamp of the modification time in milliseconds.
	//
	// example:
	//
	// 1530078690000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The UserID of the work owner in Quickbi.
	//
	// example:
	//
	// 74f5527216d14e9892245320ebf2****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud account name of the work owner.
	//
	// example:
	//
	// w****@aliyun.com
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
	// The publishing status of the report. Valid values:
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
	// 	- 0: No embedding is enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 0
	ThirdPartAuthFlag *int32 `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// The name of the report.
	//
	// example:
	//
	// Test report
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
	// 	- dashboardOfflineQuery: self-service data retrieval
	//
	// example:
	//
	// DATAFORM
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// The ID of the operations report.
	//
	// example:
	//
	// 97f7f4c1-543a-4069-8e8d-a56cfcd6****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// The ID of the workspace to which the report belongs.
	//
	// example:
	//
	// c5f86ad2-ef53-4c51-8720-162ecfdb****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the report belongs.
	//
	// example:
	//
	// Return to Professional Edition
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QuerySharesToUserListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QuerySharesToUserListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QuerySharesToUserListResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *QuerySharesToUserListResponseBodyResult) GetDirectory() *QuerySharesToUserListResponseBodyResultDirectory {
	return s.Directory
}

func (s *QuerySharesToUserListResponseBodyResult) GetModifyName() *string {
	return s.ModifyName
}

func (s *QuerySharesToUserListResponseBodyResult) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *QuerySharesToUserListResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QuerySharesToUserListResponseBodyResult) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QuerySharesToUserListResponseBodyResult) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *QuerySharesToUserListResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *QuerySharesToUserListResponseBodyResult) GetThirdPartAuthFlag() *int32 {
	return s.ThirdPartAuthFlag
}

func (s *QuerySharesToUserListResponseBodyResult) GetWorkName() *string {
	return s.WorkName
}

func (s *QuerySharesToUserListResponseBodyResult) GetWorkType() *string {
	return s.WorkType
}

func (s *QuerySharesToUserListResponseBodyResult) GetWorksId() *string {
	return s.WorksId
}

func (s *QuerySharesToUserListResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QuerySharesToUserListResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QuerySharesToUserListResponseBodyResult) SetCreateTime(v string) *QuerySharesToUserListResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetDescription(v string) *QuerySharesToUserListResponseBodyResult {
	s.Description = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetDirectory(v *QuerySharesToUserListResponseBodyResultDirectory) *QuerySharesToUserListResponseBodyResult {
	s.Directory = v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetModifyName(v string) *QuerySharesToUserListResponseBodyResult {
	s.ModifyName = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetModifyTime(v string) *QuerySharesToUserListResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetOwnerId(v string) *QuerySharesToUserListResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetOwnerName(v string) *QuerySharesToUserListResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetSecurityLevel(v string) *QuerySharesToUserListResponseBodyResult {
	s.SecurityLevel = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetStatus(v int32) *QuerySharesToUserListResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetThirdPartAuthFlag(v int32) *QuerySharesToUserListResponseBodyResult {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorkName(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorkName = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorkType(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorkType = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorksId(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorkspaceId(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorkspaceName(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySharesToUserListResponseBodyResultDirectory struct {
	// The ID of the directory where the resource is located.
	//
	// example:
	//
	// f7f6e22b-83be-47fd-b49d-9ca686a9****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// Chart Report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path ID of the directory where the resource is located.
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The path name of the directory where the resource is located.
	//
	// example:
	//
	// Level -1 Directory /Level -2 Directory
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QuerySharesToUserListResponseBodyResultDirectory) String() string {
	return dara.Prettify(s)
}

func (s QuerySharesToUserListResponseBodyResultDirectory) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) GetId() *string {
	return s.Id
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) GetName() *string {
	return s.Name
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) GetPathId() *string {
	return s.PathId
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) GetPathName() *string {
	return s.PathName
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) SetId(v string) *QuerySharesToUserListResponseBodyResultDirectory {
	s.Id = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) SetName(v string) *QuerySharesToUserListResponseBodyResultDirectory {
	s.Name = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) SetPathId(v string) *QuerySharesToUserListResponseBodyResultDirectory {
	s.PathId = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) SetPathName(v string) *QuerySharesToUserListResponseBodyResultDirectory {
	s.PathName = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) Validate() error {
	return dara.Validate(s)
}
