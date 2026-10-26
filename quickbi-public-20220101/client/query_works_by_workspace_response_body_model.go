// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksByWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryWorksByWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v *QueryWorksByWorkspaceResponseBodyResult) *QueryWorksByWorkspaceResponseBody
	GetResult() *QueryWorksByWorkspaceResponseBodyResult
	SetSuccess(v bool) *QueryWorksByWorkspaceResponseBody
	GetSuccess() *bool
}

type QueryWorksByWorkspaceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns a list of all works in the organization workspace that meet the requested criteria.
	Result *QueryWorksByWorkspaceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s QueryWorksByWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWorksByWorkspaceResponseBody) GetResult() *QueryWorksByWorkspaceResponseBodyResult {
	return s.Result
}

func (s *QueryWorksByWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryWorksByWorkspaceResponseBody) SetRequestId(v string) *QueryWorksByWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBody) SetResult(v *QueryWorksByWorkspaceResponseBodyResult) *QueryWorksByWorkspaceResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorksByWorkspaceResponseBody) SetSuccess(v bool) *QueryWorksByWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWorksByWorkspaceResponseBodyResult struct {
	// The details of the list of works.
	Data []*QueryWorksByWorkspaceResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page set when the interface is requested.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of rows in the table.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryWorksByWorkspaceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByWorkspaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponseBodyResult) GetData() []*QueryWorksByWorkspaceResponseBodyResultData {
	return s.Data
}

func (s *QueryWorksByWorkspaceResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryWorksByWorkspaceResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryWorksByWorkspaceResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *QueryWorksByWorkspaceResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetData(v []*QueryWorksByWorkspaceResponseBodyResultData) *QueryWorksByWorkspaceResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetPageNum(v int32) *QueryWorksByWorkspaceResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetPageSize(v int32) *QueryWorksByWorkspaceResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetTotalNum(v int32) *QueryWorksByWorkspaceResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetTotalPages(v int32) *QueryWorksByWorkspaceResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryWorksByWorkspaceResponseBodyResultData struct {
	// Third-party embedding status. Valid values:
	//
	// 	- 0: The embed service is not enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 1
	Auth3rdFlag *int32 `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	// Remarks on the work.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory to which the work belongs.
	Directory *QueryWorksByWorkspaceResponseBodyResultDataDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
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
	// 1572334870000
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// Nickname of the work modifier.
	//
	// example:
	//
	// Tom
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// The user ID of the work owner in the Quick BI.
	//
	// example:
	//
	// The name of the workspace to which the work belongs.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The nickname of the work owner.
	//
	// example:
	//
	// Li Si
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
	// Status of dashboards, full-screen dashboards, spreadsheets. The default value of other work types is 1. Valid values:
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
	// The name of the work.
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
	// 	- Analysis: Ad hoc analysis
	//
	// 	- DATAFORM: form filling
	//
	// example:
	//
	// PAGE
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 897ce25e-f993-4abd-af84-d13c5610****
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
	// Test Workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryWorksByWorkspaceResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByWorkspaceResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetAuth3rdFlag() *int32 {
	return s.Auth3rdFlag
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetDescription() *string {
	return s.Description
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetDirectory() *QueryWorksByWorkspaceResponseBodyResultDataDirectory {
	return s.Directory
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetGmtModify() *string {
	return s.GmtModify
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetModifyName() *string {
	return s.ModifyName
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetPublicFlag() *bool {
	return s.PublicFlag
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetPublicInvalidTime() *int64 {
	return s.PublicInvalidTime
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetStatus() *int32 {
	return s.Status
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetWorkName() *string {
	return s.WorkName
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetWorkType() *string {
	return s.WorkType
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetWorksId() *string {
	return s.WorksId
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetAuth3rdFlag(v int32) *QueryWorksByWorkspaceResponseBodyResultData {
	s.Auth3rdFlag = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetDescription(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.Description = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetDirectory(v *QueryWorksByWorkspaceResponseBodyResultDataDirectory) *QueryWorksByWorkspaceResponseBodyResultData {
	s.Directory = v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetGmtCreate(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetGmtModify(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.GmtModify = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetModifyName(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.ModifyName = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetOwnerId(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.OwnerId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetOwnerName(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetPublicFlag(v bool) *QueryWorksByWorkspaceResponseBodyResultData {
	s.PublicFlag = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetPublicInvalidTime(v int64) *QueryWorksByWorkspaceResponseBodyResultData {
	s.PublicInvalidTime = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetSecurityLevel(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.SecurityLevel = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetStatus(v int32) *QueryWorksByWorkspaceResponseBodyResultData {
	s.Status = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorkName(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorkName = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorkType(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorkType = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorksId(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorksId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorkspaceId(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorkspaceName(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWorksByWorkspaceResponseBodyResultDataDirectory struct {
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
	// The name of the directory.
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

func (s QueryWorksByWorkspaceResponseBodyResultDataDirectory) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByWorkspaceResponseBodyResultDataDirectory) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) GetId() *string {
	return s.Id
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) GetName() *string {
	return s.Name
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) GetPathId() *string {
	return s.PathId
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) GetPathName() *string {
	return s.PathName
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) SetId(v string) *QueryWorksByWorkspaceResponseBodyResultDataDirectory {
	s.Id = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) SetName(v string) *QueryWorksByWorkspaceResponseBodyResultDataDirectory {
	s.Name = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) SetPathId(v string) *QueryWorksByWorkspaceResponseBodyResultDataDirectory {
	s.PathId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) SetPathName(v string) *QueryWorksByWorkspaceResponseBodyResultDataDirectory {
	s.PathName = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) Validate() error {
	return dara.Validate(s)
}
