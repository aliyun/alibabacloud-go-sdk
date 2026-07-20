// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRbacRolePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryRbacRolePermissionsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryRbacRolePermissionsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryRbacRolePermissionsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryRbacRolePermissionsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryRbacRolePermissionsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryRbacRolePermissionsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QueryRbacRolePermissionsResponseBodyModule) *QueryRbacRolePermissionsResponseBody
	GetModule() *QueryRbacRolePermissionsResponseBodyModule
	SetRequestId(v string) *QueryRbacRolePermissionsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryRbacRolePermissionsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryRbacRolePermissionsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QueryRbacRolePermissionsResponseBody
	GetSynchro() *bool
}

type QueryRbacRolePermissionsResponseBody struct {
	AccessDeniedDetail *string                                     `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AllowRetry         *bool                                       `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName            *string                                     `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DynamicCode        *string                                     `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs          []interface{}                               `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module             *QueryRbacRolePermissionsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	RequestId          *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootErrorCode      *string                                     `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg       *string                                     `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	Synchro            *bool                                       `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryRbacRolePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRolePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRbacRolePermissionsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryRbacRolePermissionsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryRbacRolePermissionsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryRbacRolePermissionsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryRbacRolePermissionsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryRbacRolePermissionsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryRbacRolePermissionsResponseBody) GetModule() *QueryRbacRolePermissionsResponseBodyModule {
	return s.Module
}

func (s *QueryRbacRolePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRbacRolePermissionsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryRbacRolePermissionsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryRbacRolePermissionsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryRbacRolePermissionsResponseBody) SetAccessDeniedDetail(v string) *QueryRbacRolePermissionsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetAllowRetry(v bool) *QueryRbacRolePermissionsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetAppName(v string) *QueryRbacRolePermissionsResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetDynamicCode(v string) *QueryRbacRolePermissionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetDynamicMessage(v string) *QueryRbacRolePermissionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetErrorArgs(v []interface{}) *QueryRbacRolePermissionsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetModule(v *QueryRbacRolePermissionsResponseBodyModule) *QueryRbacRolePermissionsResponseBody {
	s.Module = v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetRequestId(v string) *QueryRbacRolePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetRootErrorCode(v string) *QueryRbacRolePermissionsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetRootErrorMsg(v string) *QueryRbacRolePermissionsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) SetSynchro(v bool) *QueryRbacRolePermissionsResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRbacRolePermissionsResponseBodyModule struct {
	CurrentPageNum *int32                                            `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryRbacRolePermissionsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Next           *QueryRbacRolePermissionsResponseBodyModuleNext   `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	NextPage       *bool                                             `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                                             `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	ResultLimit    *bool                                             `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	TotalItemNum   *int32                                            `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                            `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryRbacRolePermissionsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRolePermissionsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryRbacRolePermissionsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryRbacRolePermissionsResponseBodyModule) GetData() []*QueryRbacRolePermissionsResponseBodyModuleData {
	return s.Data
}

func (s *QueryRbacRolePermissionsResponseBodyModule) GetNext() *QueryRbacRolePermissionsResponseBodyModuleNext {
	return s.Next
}

func (s *QueryRbacRolePermissionsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryRbacRolePermissionsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRbacRolePermissionsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryRbacRolePermissionsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *QueryRbacRolePermissionsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryRbacRolePermissionsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryRbacRolePermissionsResponseBodyModule) SetCurrentPageNum(v int32) *QueryRbacRolePermissionsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModule) SetData(v []*QueryRbacRolePermissionsResponseBodyModuleData) *QueryRbacRolePermissionsResponseBodyModule {
	s.Data = v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModule) SetNext(v *QueryRbacRolePermissionsResponseBodyModuleNext) *QueryRbacRolePermissionsResponseBodyModule {
	s.Next = v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModule) SetNextPage(v bool) *QueryRbacRolePermissionsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModule) SetPageSize(v int32) *QueryRbacRolePermissionsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModule) SetPrePage(v bool) *QueryRbacRolePermissionsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModule) SetResultLimit(v bool) *QueryRbacRolePermissionsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModule) SetTotalItemNum(v int32) *QueryRbacRolePermissionsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModule) SetTotalPageNum(v int32) *QueryRbacRolePermissionsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModule) Validate() error {
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

type QueryRbacRolePermissionsResponseBodyModuleData struct {
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Resource    *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s QueryRbacRolePermissionsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRolePermissionsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *QueryRbacRolePermissionsResponseBodyModuleData) GetAction() *string {
	return s.Action
}

func (s *QueryRbacRolePermissionsResponseBodyModuleData) GetDescription() *string {
	return s.Description
}

func (s *QueryRbacRolePermissionsResponseBodyModuleData) GetId() *string {
	return s.Id
}

func (s *QueryRbacRolePermissionsResponseBodyModuleData) GetResource() *string {
	return s.Resource
}

func (s *QueryRbacRolePermissionsResponseBodyModuleData) SetAction(v string) *QueryRbacRolePermissionsResponseBodyModuleData {
	s.Action = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModuleData) SetDescription(v string) *QueryRbacRolePermissionsResponseBodyModuleData {
	s.Description = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModuleData) SetId(v string) *QueryRbacRolePermissionsResponseBodyModuleData {
	s.Id = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModuleData) SetResource(v string) *QueryRbacRolePermissionsResponseBodyModuleData {
	s.Resource = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type QueryRbacRolePermissionsResponseBodyModuleNext struct {
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Resource    *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s QueryRbacRolePermissionsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRolePermissionsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *QueryRbacRolePermissionsResponseBodyModuleNext) GetAction() *string {
	return s.Action
}

func (s *QueryRbacRolePermissionsResponseBodyModuleNext) GetDescription() *string {
	return s.Description
}

func (s *QueryRbacRolePermissionsResponseBodyModuleNext) GetId() *string {
	return s.Id
}

func (s *QueryRbacRolePermissionsResponseBodyModuleNext) GetResource() *string {
	return s.Resource
}

func (s *QueryRbacRolePermissionsResponseBodyModuleNext) SetAction(v string) *QueryRbacRolePermissionsResponseBodyModuleNext {
	s.Action = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModuleNext) SetDescription(v string) *QueryRbacRolePermissionsResponseBodyModuleNext {
	s.Description = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModuleNext) SetId(v string) *QueryRbacRolePermissionsResponseBodyModuleNext {
	s.Id = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModuleNext) SetResource(v string) *QueryRbacRolePermissionsResponseBodyModuleNext {
	s.Resource = &v
	return s
}

func (s *QueryRbacRolePermissionsResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
