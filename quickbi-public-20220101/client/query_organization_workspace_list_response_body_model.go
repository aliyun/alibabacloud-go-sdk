// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrganizationWorkspaceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryOrganizationWorkspaceListResponseBody
	GetRequestId() *string
	SetResult(v *QueryOrganizationWorkspaceListResponseBodyResult) *QueryOrganizationWorkspaceListResponseBody
	GetResult() *QueryOrganizationWorkspaceListResponseBodyResult
	SetSuccess(v bool) *QueryOrganizationWorkspaceListResponseBody
	GetSuccess() *bool
}

type QueryOrganizationWorkspaceListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the paginated result of the workspace list, with detailed information about the workspaces stored in the Data parameter.
	Result *QueryOrganizationWorkspaceListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s QueryOrganizationWorkspaceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationWorkspaceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOrganizationWorkspaceListResponseBody) GetResult() *QueryOrganizationWorkspaceListResponseBodyResult {
	return s.Result
}

func (s *QueryOrganizationWorkspaceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryOrganizationWorkspaceListResponseBody) SetRequestId(v string) *QueryOrganizationWorkspaceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBody) SetResult(v *QueryOrganizationWorkspaceListResponseBodyResult) *QueryOrganizationWorkspaceListResponseBody {
	s.Result = v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBody) SetSuccess(v bool) *QueryOrganizationWorkspaceListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOrganizationWorkspaceListResponseBodyResult struct {
	// List of workspaces.
	Data []*QueryOrganizationWorkspaceListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of rows per page as set in the request.
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

func (s QueryOrganizationWorkspaceListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationWorkspaceListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) GetData() []*QueryOrganizationWorkspaceListResponseBodyResultData {
	return s.Data
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetData(v []*QueryOrganizationWorkspaceListResponseBodyResultData) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetPageNum(v int32) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetPageSize(v int32) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetTotalNum(v int32) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetTotalPages(v int32) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) Validate() error {
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

type QueryOrganizationWorkspaceListResponseBodyResultData struct {
	// Whether the work can be made public. Value range:
	//
	// - true: Public
	//
	// - false: Not public
	//
	// example:
	//
	// true
	AllowPublishOperation *bool `json:"AllowPublishOperation,omitempty" xml:"AllowPublishOperation,omitempty"`
	// Indicates whether the work can be authorized for sharing. Possible values:
	//
	// - true: Authorized
	//
	// - false: Not authorized
	//
	// example:
	//
	// true
	AllowShareOperation *bool `json:"AllowShareOperation,omitempty" xml:"AllowShareOperation,omitempty"`
	// Creation time of the workspace.
	//
	// example:
	//
	// 2020-11-10 17:51:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Quick BI user ID of the creator.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Aliyun account name of the creator.
	//
	// example:
	//
	// pop****@aliyun.com
	CreateUserAccountName *string `json:"CreateUserAccountName,omitempty" xml:"CreateUserAccountName,omitempty"`
	// Last modified time of the workspace.
	//
	// example:
	//
	// 2020-11-10 17:51:07
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// ID of the Quick BI user who modified the workspace.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// Aliyun account name of the modifier.
	//
	// example:
	//
	// pop****@aliyun.com
	ModifyUserAccountName *string `json:"ModifyUserAccountName,omitempty" xml:"ModifyUserAccountName,omitempty"`
	// ID of the organization to which the workspace belongs.
	//
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// Quick BI user ID of the workspace owner.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Aliyun account name of the workspace owner.
	//
	// example:
	//
	// pop****@aliyun.com
	OwnerAccountName *string `json:"OwnerAccountName,omitempty" xml:"OwnerAccountName,omitempty"`
	// example:
	//
	// test
	RealOwnerAccountName *string `json:"RealOwnerAccountName,omitempty" xml:"RealOwnerAccountName,omitempty"`
	// Workspace description.
	//
	// example:
	//
	// test
	WorkspaceDescription *string `json:"WorkspaceDescription,omitempty" xml:"WorkspaceDescription,omitempty"`
	// Workspace ID.
	//
	// example:
	//
	// 7350a155-0e94-4c6c-8620-57bbec38****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// Name of the workspace.
	//
	// example:
	//
	// test
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryOrganizationWorkspaceListResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationWorkspaceListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetAllowPublishOperation() *bool {
	return s.AllowPublishOperation
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetAllowShareOperation() *bool {
	return s.AllowShareOperation
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetCreateUserAccountName() *string {
	return s.CreateUserAccountName
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetModifyUserAccountName() *string {
	return s.ModifyUserAccountName
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetOwner() *string {
	return s.Owner
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetOwnerAccountName() *string {
	return s.OwnerAccountName
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetRealOwnerAccountName() *string {
	return s.RealOwnerAccountName
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetWorkspaceDescription() *string {
	return s.WorkspaceDescription
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetAllowPublishOperation(v bool) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.AllowPublishOperation = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetAllowShareOperation(v bool) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.AllowShareOperation = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetCreateTime(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.CreateTime = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetCreateUser(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.CreateUser = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetCreateUserAccountName(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.CreateUserAccountName = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetModifiedTime(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.ModifiedTime = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetModifyUser(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.ModifyUser = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetModifyUserAccountName(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.ModifyUserAccountName = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetOrganizationId(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.OrganizationId = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetOwner(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.Owner = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetOwnerAccountName(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.OwnerAccountName = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetRealOwnerAccountName(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.RealOwnerAccountName = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetWorkspaceDescription(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.WorkspaceDescription = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetWorkspaceId(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetWorkspaceName(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
