// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRbacUserRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryRbacUserRolesResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryRbacUserRolesResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryRbacUserRolesResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryRbacUserRolesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryRbacUserRolesResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryRbacUserRolesResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QueryRbacUserRolesResponseBodyModule) *QueryRbacUserRolesResponseBody
	GetModule() *QueryRbacUserRolesResponseBodyModule
	SetRequestId(v string) *QueryRbacUserRolesResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryRbacUserRolesResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryRbacUserRolesResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QueryRbacUserRolesResponseBody
	GetSynchro() *bool
}

type QueryRbacUserRolesResponseBody struct {
	AccessDeniedDetail *string                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AllowRetry         *bool                                 `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName            *string                               `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DynamicCode        *string                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs          []interface{}                         `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module             *QueryRbacUserRolesResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	RequestId          *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootErrorCode      *string                               `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg       *string                               `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	Synchro            *bool                                 `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryRbacUserRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRbacUserRolesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryRbacUserRolesResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryRbacUserRolesResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryRbacUserRolesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryRbacUserRolesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryRbacUserRolesResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryRbacUserRolesResponseBody) GetModule() *QueryRbacUserRolesResponseBodyModule {
	return s.Module
}

func (s *QueryRbacUserRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRbacUserRolesResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryRbacUserRolesResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryRbacUserRolesResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryRbacUserRolesResponseBody) SetAccessDeniedDetail(v string) *QueryRbacUserRolesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetAllowRetry(v bool) *QueryRbacUserRolesResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetAppName(v string) *QueryRbacUserRolesResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetDynamicCode(v string) *QueryRbacUserRolesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetDynamicMessage(v string) *QueryRbacUserRolesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetErrorArgs(v []interface{}) *QueryRbacUserRolesResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetModule(v *QueryRbacUserRolesResponseBodyModule) *QueryRbacUserRolesResponseBody {
	s.Module = v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetRequestId(v string) *QueryRbacUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetRootErrorCode(v string) *QueryRbacUserRolesResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetRootErrorMsg(v string) *QueryRbacUserRolesResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryRbacUserRolesResponseBody) SetSynchro(v bool) *QueryRbacUserRolesResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryRbacUserRolesResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRbacUserRolesResponseBodyModule struct {
	CurrentPageNum *int32                                      `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryRbacUserRolesResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Next           *QueryRbacUserRolesResponseBodyModuleNext   `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	NextPage       *bool                                       `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                                       `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	ResultLimit    *bool                                       `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	TotalItemNum   *int32                                      `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                      `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryRbacUserRolesResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacUserRolesResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryRbacUserRolesResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryRbacUserRolesResponseBodyModule) GetData() []*QueryRbacUserRolesResponseBodyModuleData {
	return s.Data
}

func (s *QueryRbacUserRolesResponseBodyModule) GetNext() *QueryRbacUserRolesResponseBodyModuleNext {
	return s.Next
}

func (s *QueryRbacUserRolesResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryRbacUserRolesResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRbacUserRolesResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryRbacUserRolesResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *QueryRbacUserRolesResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryRbacUserRolesResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryRbacUserRolesResponseBodyModule) SetCurrentPageNum(v int32) *QueryRbacUserRolesResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModule) SetData(v []*QueryRbacUserRolesResponseBodyModuleData) *QueryRbacUserRolesResponseBodyModule {
	s.Data = v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModule) SetNext(v *QueryRbacUserRolesResponseBodyModuleNext) *QueryRbacUserRolesResponseBodyModule {
	s.Next = v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModule) SetNextPage(v bool) *QueryRbacUserRolesResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModule) SetPageSize(v int32) *QueryRbacUserRolesResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModule) SetPrePage(v bool) *QueryRbacUserRolesResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModule) SetResultLimit(v bool) *QueryRbacUserRolesResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModule) SetTotalItemNum(v int32) *QueryRbacUserRolesResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModule) SetTotalPageNum(v int32) *QueryRbacUserRolesResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Next != nil {
		if err := s.Next.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRbacUserRolesResponseBodyModuleData struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	RoleId    *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryRbacUserRolesResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacUserRolesResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *QueryRbacUserRolesResponseBodyModuleData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *QueryRbacUserRolesResponseBodyModuleData) GetId() *string {
	return s.Id
}

func (s *QueryRbacUserRolesResponseBodyModuleData) GetOrgId() *string {
	return s.OrgId
}

func (s *QueryRbacUserRolesResponseBodyModuleData) GetRoleId() *string {
	return s.RoleId
}

func (s *QueryRbacUserRolesResponseBodyModuleData) GetUserId() *string {
	return s.UserId
}

func (s *QueryRbacUserRolesResponseBodyModuleData) SetCreatedAt(v string) *QueryRbacUserRolesResponseBodyModuleData {
	s.CreatedAt = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleData) SetId(v string) *QueryRbacUserRolesResponseBodyModuleData {
	s.Id = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleData) SetOrgId(v string) *QueryRbacUserRolesResponseBodyModuleData {
	s.OrgId = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleData) SetRoleId(v string) *QueryRbacUserRolesResponseBodyModuleData {
	s.RoleId = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleData) SetUserId(v string) *QueryRbacUserRolesResponseBodyModuleData {
	s.UserId = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type QueryRbacUserRolesResponseBodyModuleNext struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	RoleId    *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryRbacUserRolesResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacUserRolesResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) GetId() *string {
	return s.Id
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) GetOrgId() *string {
	return s.OrgId
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) GetRoleId() *string {
	return s.RoleId
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) GetUserId() *string {
	return s.UserId
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) SetCreatedAt(v string) *QueryRbacUserRolesResponseBodyModuleNext {
	s.CreatedAt = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) SetId(v string) *QueryRbacUserRolesResponseBodyModuleNext {
	s.Id = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) SetOrgId(v string) *QueryRbacUserRolesResponseBodyModuleNext {
	s.OrgId = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) SetRoleId(v string) *QueryRbacUserRolesResponseBodyModuleNext {
	s.RoleId = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) SetUserId(v string) *QueryRbacUserRolesResponseBodyModuleNext {
	s.UserId = &v
	return s
}

func (s *QueryRbacUserRolesResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
