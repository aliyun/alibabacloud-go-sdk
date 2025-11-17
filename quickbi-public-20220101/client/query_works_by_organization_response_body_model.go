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
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns a list of all works under the organization that meet the request criteria.
	Result *QueryWorksByOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
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
	// Details of the work list.
	Data []*QueryWorksByOrganizationResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page set when requesting the interface.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of rows.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// Total number of pages.
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
	// Third-party embedding status. Possible values:
	//
	// - 0: Embedding not enabled
	//
	// - 1: Embedding enabled
	//
	// example:
	//
	// 1
	Auth3rdFlag *int32 `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	// Notes for the work.
	//
	// example:
	//
	// Attention
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Directory to which the work belongs.
	Directory *QueryWorksByOrganizationResponseBodyResultDataDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// Timestamp (in milliseconds) when the work was created.
	//
	// example:
	//
	// 1496651577000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Millisecond-level timestamp of the work modification.
	//
	// example:
	//
	// 1572334870000
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// The Alibaba Cloud account name of the work modifier.
	//
	// example:
	//
	// test
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// The UserID of the work\\"s owner in Quick BI.
	//
	// example:
	//
	// test
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud account name of the work\\"s owner.
	//
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Whether it is public
	//
	// example:
	//
	// true
	PublicFlag *bool `json:"PublicFlag,omitempty" xml:"PublicFlag,omitempty"`
	// The expiration date for the report to be made public
	//
	// example:
	//
	// 1721366354000
	PublicInvalidTime *int64 `json:"PublicInvalidTime,omitempty" xml:"PublicInvalidTime,omitempty"`
	// The security policy for collaborative authorization of the work. Values:
	//
	// - 0: Private
	//
	// - 12: Authorize specific members
	//
	// - 1 or 11: Authorize all space members
	//
	// >- If you are using the old version of permissions, the returned value is 1.
	//
	// >- If you are using the new version of permissions, the returned value is 11.
	//
	// example:
	//
	// 1
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The status of the report. Value range:
	//
	// - 0: Unpublished
	//
	// - 1: Published
	//
	// - 2: Modified but not published
	//
	// - 3: Offline
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the work.
	//
	// example:
	//
	// test
	WorkName *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	// The type of the work. Value range:
	//
	// - DATAPRODUCT: Data portal
	//
	// - PAGE: Dashboard
	//
	// - REPORT: Spreadsheet
	//
	// - dashboardOfflineQuery: Self-service data retrieval
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
	// test
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// test
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
	// ID of the directory to which it belongs.
	//
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Name of the directory to which it belongs.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Hierarchical structure of the directory ID, separated by『/』.
	//
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// Hierarchical structure of the directory name, separated by『/』.
	//
	// example:
	//
	// Attention
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
