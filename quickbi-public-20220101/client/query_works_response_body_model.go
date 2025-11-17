// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryWorksResponseBody
	GetRequestId() *string
	SetResult(v *QueryWorksResponseBodyResult) *QueryWorksResponseBody
	GetResult() *QueryWorksResponseBodyResult
	SetSuccess(v bool) *QueryWorksResponseBody
	GetSuccess() *bool
}

type QueryWorksResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the work.
	Result *QueryWorksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s QueryWorksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWorksResponseBody) GetResult() *QueryWorksResponseBodyResult {
	return s.Result
}

func (s *QueryWorksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryWorksResponseBody) SetRequestId(v string) *QueryWorksResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorksResponseBody) SetResult(v *QueryWorksResponseBodyResult) *QueryWorksResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorksResponseBody) SetSuccess(v bool) *QueryWorksResponseBody {
	s.Success = &v
	return s
}

func (s *QueryWorksResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWorksResponseBodyResult struct {
	// Third-party embedding status. Valid values:
	//
	// 	- 0: The embed service is not enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 0
	Auth3rdFlag *int32 `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	// Remarks on the work.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory to which the work belongs.
	Directory *QueryWorksResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The timestamp of the creation of the work in milliseconds.
	//
	// example:
	//
	// 1496651577000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp of the modification of the work in milliseconds.
	//
	// example:
	//
	// 1496651577000
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// The Alibaba Cloud account name of the person who modified the work.
	//
	// example:
	//
	// Tom
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// The user ID of the work owner in the Quick BI.
	//
	// example:
	//
	// 9187a612aa474e2a8ac1414d5529****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud account name of the work owner.
	//
	// example:
	//
	// Tom
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Is it public
	//
	// example:
	//
	// true
	PublicFlag *bool `json:"PublicFlag,omitempty" xml:"PublicFlag,omitempty"`
	// Deadline for the public release of the report
	//
	// example:
	//
	// 1721366354000
	PublicInvalidTime *int64 `json:"PublicInvalidTime,omitempty" xml:"PublicInvalidTime,omitempty"`
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
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the work.
	//
	// example:
	//
	// Test report
	WorkName *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	// Queries the types of works. Fill in the blanks to query all types. Valid values:
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
	// PAGE
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// The ID of the workspace to which the work belongs.
	//
	// example:
	//
	// 87c6b145-090c-43e1-9426-8f93be23****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// Test Space
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryWorksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorksResponseBodyResult) GetAuth3rdFlag() *int32 {
	return s.Auth3rdFlag
}

func (s *QueryWorksResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *QueryWorksResponseBodyResult) GetDirectory() *QueryWorksResponseBodyResultDirectory {
	return s.Directory
}

func (s *QueryWorksResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryWorksResponseBodyResult) GetGmtModify() *string {
	return s.GmtModify
}

func (s *QueryWorksResponseBodyResult) GetModifyName() *string {
	return s.ModifyName
}

func (s *QueryWorksResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryWorksResponseBodyResult) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QueryWorksResponseBodyResult) GetPublicFlag() *bool {
	return s.PublicFlag
}

func (s *QueryWorksResponseBodyResult) GetPublicInvalidTime() *int64 {
	return s.PublicInvalidTime
}

func (s *QueryWorksResponseBodyResult) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *QueryWorksResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *QueryWorksResponseBodyResult) GetWorkName() *string {
	return s.WorkName
}

func (s *QueryWorksResponseBodyResult) GetWorkType() *string {
	return s.WorkType
}

func (s *QueryWorksResponseBodyResult) GetWorksId() *string {
	return s.WorksId
}

func (s *QueryWorksResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryWorksResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryWorksResponseBodyResult) SetAuth3rdFlag(v int32) *QueryWorksResponseBodyResult {
	s.Auth3rdFlag = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetDescription(v string) *QueryWorksResponseBodyResult {
	s.Description = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetDirectory(v *QueryWorksResponseBodyResultDirectory) *QueryWorksResponseBodyResult {
	s.Directory = v
	return s
}

func (s *QueryWorksResponseBodyResult) SetGmtCreate(v string) *QueryWorksResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetGmtModify(v string) *QueryWorksResponseBodyResult {
	s.GmtModify = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetModifyName(v string) *QueryWorksResponseBodyResult {
	s.ModifyName = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetOwnerId(v string) *QueryWorksResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetOwnerName(v string) *QueryWorksResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetPublicFlag(v bool) *QueryWorksResponseBodyResult {
	s.PublicFlag = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetPublicInvalidTime(v int64) *QueryWorksResponseBodyResult {
	s.PublicInvalidTime = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetSecurityLevel(v string) *QueryWorksResponseBodyResult {
	s.SecurityLevel = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetStatus(v int32) *QueryWorksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorkName(v string) *QueryWorksResponseBodyResult {
	s.WorkName = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorkType(v string) *QueryWorksResponseBodyResult {
	s.WorkType = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorksId(v string) *QueryWorksResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorkspaceId(v string) *QueryWorksResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorkspaceName(v string) *QueryWorksResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *QueryWorksResponseBodyResult) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWorksResponseBodyResultDirectory struct {
	// The ID of the directory.
	//
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the directory.
	//
	// example:
	//
	// Test directory
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The hierarchical structure of the directory ID to which the directory belongs. Separate the hierarchical structure with a /.
	//
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The hierarchical structure of the directory to which the directory belongs. Separate the hierarchical structure with a (/).
	//
	// example:
	//
	// Test directory
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryWorksResponseBodyResultDirectory) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksResponseBodyResultDirectory) GoString() string {
	return s.String()
}

func (s *QueryWorksResponseBodyResultDirectory) GetId() *string {
	return s.Id
}

func (s *QueryWorksResponseBodyResultDirectory) GetName() *string {
	return s.Name
}

func (s *QueryWorksResponseBodyResultDirectory) GetPathId() *string {
	return s.PathId
}

func (s *QueryWorksResponseBodyResultDirectory) GetPathName() *string {
	return s.PathName
}

func (s *QueryWorksResponseBodyResultDirectory) SetId(v string) *QueryWorksResponseBodyResultDirectory {
	s.Id = &v
	return s
}

func (s *QueryWorksResponseBodyResultDirectory) SetName(v string) *QueryWorksResponseBodyResultDirectory {
	s.Name = &v
	return s
}

func (s *QueryWorksResponseBodyResultDirectory) SetPathId(v string) *QueryWorksResponseBodyResultDirectory {
	s.PathId = &v
	return s
}

func (s *QueryWorksResponseBodyResultDirectory) SetPathName(v string) *QueryWorksResponseBodyResultDirectory {
	s.PathName = &v
	return s
}

func (s *QueryWorksResponseBodyResultDirectory) Validate() error {
	return dara.Validate(s)
}
