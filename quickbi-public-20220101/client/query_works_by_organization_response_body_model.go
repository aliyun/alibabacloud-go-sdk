// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksByOrganizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryWorksByOrganizationResponseBody
	GetRequestId() *string
	SetResult(v *QueryWorksByOrganizationResponseBodyResult) *QueryWorksByOrganizationResponseBody
	GetResult() *QueryWorksByOrganizationResponseBodyResult
	SetSuccess(v bool) *QueryWorksByOrganizationResponseBody
	GetSuccess() *bool
}

type QueryWorksByOrganizationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of all works in the organization that match the request conditions.
	Result *QueryWorksByOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWorksByOrganizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWorksByOrganizationResponseBody) GetResult() *QueryWorksByOrganizationResponseBodyResult {
	return s.Result
}

func (s *QueryWorksByOrganizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryWorksByOrganizationResponseBody) SetRequestId(v string) *QueryWorksByOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBody) SetResult(v *QueryWorksByOrganizationResponseBodyResult) *QueryWorksByOrganizationResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorksByOrganizationResponseBody) SetSuccess(v bool) *QueryWorksByOrganizationResponseBody {
	s.Success = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWorksByOrganizationResponseBodyResult struct {
	// The details of the works list.
	Data []*QueryWorksByOrganizationResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page specified in the request.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of rows.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryWorksByOrganizationResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponseBodyResult) GetData() []*QueryWorksByOrganizationResponseBodyResultData {
	return s.Data
}

func (s *QueryWorksByOrganizationResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryWorksByOrganizationResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryWorksByOrganizationResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *QueryWorksByOrganizationResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetData(v []*QueryWorksByOrganizationResponseBodyResultData) *QueryWorksByOrganizationResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetPageNum(v int32) *QueryWorksByOrganizationResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetPageSize(v int32) *QueryWorksByOrganizationResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetTotalNum(v int32) *QueryWorksByOrganizationResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetTotalPages(v int32) *QueryWorksByOrganizationResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResult) Validate() error {
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

type QueryWorksByOrganizationResponseBodyResultData struct {
	// The third-party embedding status. Valid values:
	//
	// - 0: Embedding is not enabled.
	//
	// - 1: Embedding is enabled.
	//
	// example:
	//
	// 1
	Auth3rdFlag *int32 `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	// The description of the work.
	//
	// example:
	//
	// Remarks
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory to which the work belongs.
	Directory *QueryWorksByOrganizationResponseBodyResultDataDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The timestamp when the work was created, in milliseconds.
	//
	// example:
	//
	// 1496651577000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the work was last modified, in milliseconds.
	//
	// example:
	//
	// 1572334870000
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// The Alibaba Cloud account name of the user who last modified the work.
	//
	// example:
	//
	// John
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// The Quick BI user ID of the work owner.
	//
	// example:
	//
	// The name of the workspace to which the work belongs
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud account name of the work owner.
	//
	// example:
	//
	// Jane
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Indicates whether the work is public.
	//
	// example:
	//
	// true
	PublicFlag *bool `json:"PublicFlag,omitempty" xml:"PublicFlag,omitempty"`
	// The expiration date for the public report, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1721366354000
	PublicInvalidTime *int64 `json:"PublicInvalidTime,omitempty" xml:"PublicInvalidTime,omitempty"`
	// The security policy for collaborative authorization of the work. Valid values:
	//
	// - 0: Private.
	//
	// - 12: Authorized to specified members.
	//
	// - 1 or 11: Authorized to all workspace members.
	//
	// >- If you use the legacy permissions, the return value is 1.
	//
	// >- If you use the new permissions, the return value is 11.
	//
	// example:
	//
	// 1
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The report status. Valid values:
	//
	// - 0: Unpublished.
	//
	// - 1: Published.
	//
	// - 2: Modified but not published.
	//
	// - 3: Offline.
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
	// - DATAPRODUCT: data portal
	//
	// - PAGE: dashboard
	//
	// - REPORT: workbook
	//
	// - dashboardOfflineQuery: custom query
	//
	// example:
	//
	// PAGE
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 897ce25e-****-****-af84-d13c5610****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// The ID of the workspace to which the work belongs.
	//
	// example:
	//
	// The Alibaba Cloud account name of the user who last modified the work
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// Test workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryWorksByOrganizationResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByOrganizationResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetAuth3rdFlag() *int32 {
	return s.Auth3rdFlag
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetDescription() *string {
	return s.Description
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetDirectory() *QueryWorksByOrganizationResponseBodyResultDataDirectory {
	return s.Directory
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetGmtModify() *string {
	return s.GmtModify
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetModifyName() *string {
	return s.ModifyName
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetPublicFlag() *bool {
	return s.PublicFlag
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetPublicInvalidTime() *int64 {
	return s.PublicInvalidTime
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetStatus() *int32 {
	return s.Status
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetWorkName() *string {
	return s.WorkName
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetWorkType() *string {
	return s.WorkType
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetWorksId() *string {
	return s.WorksId
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryWorksByOrganizationResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetAuth3rdFlag(v int32) *QueryWorksByOrganizationResponseBodyResultData {
	s.Auth3rdFlag = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetDescription(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.Description = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetDirectory(v *QueryWorksByOrganizationResponseBodyResultDataDirectory) *QueryWorksByOrganizationResponseBodyResultData {
	s.Directory = v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetGmtCreate(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetGmtModify(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.GmtModify = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetModifyName(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.ModifyName = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetOwnerId(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.OwnerId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetOwnerName(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetPublicFlag(v bool) *QueryWorksByOrganizationResponseBodyResultData {
	s.PublicFlag = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetPublicInvalidTime(v int64) *QueryWorksByOrganizationResponseBodyResultData {
	s.PublicInvalidTime = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetSecurityLevel(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.SecurityLevel = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetStatus(v int32) *QueryWorksByOrganizationResponseBodyResultData {
	s.Status = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorkName(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorkName = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorkType(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorkType = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorksId(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorksId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorkspaceId(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorkspaceName(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWorksByOrganizationResponseBodyResultDataDirectory struct {
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
	// The hierarchical structure of directory IDs, separated by forward slashes (/).
	//
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The hierarchical structure of directory names, separated by forward slashes (/).
	//
	// example:
	//
	// Test directory
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryWorksByOrganizationResponseBodyResultDataDirectory) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByOrganizationResponseBodyResultDataDirectory) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) GetId() *string {
	return s.Id
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) GetName() *string {
	return s.Name
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) GetPathId() *string {
	return s.PathId
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) GetPathName() *string {
	return s.PathName
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) SetId(v string) *QueryWorksByOrganizationResponseBodyResultDataDirectory {
	s.Id = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) SetName(v string) *QueryWorksByOrganizationResponseBodyResultDataDirectory {
	s.Name = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) SetPathId(v string) *QueryWorksByOrganizationResponseBodyResultDataDirectory {
	s.PathId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) SetPathName(v string) *QueryWorksByOrganizationResponseBodyResultDataDirectory {
	s.PathName = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) Validate() error {
	return dara.Validate(s)
}
