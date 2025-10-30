// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAuthorizedAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataServiceAuthorizedAppsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListDataServiceAuthorizedAppsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataServiceAuthorizedAppsResponseBody
	GetMessage() *string
	SetPageResult(v *ListDataServiceAuthorizedAppsResponseBodyPageResult) *ListDataServiceAuthorizedAppsResponseBody
	GetPageResult() *ListDataServiceAuthorizedAppsResponseBodyPageResult
	SetRequestId(v string) *ListDataServiceAuthorizedAppsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceAuthorizedAppsResponseBody
	GetSuccess() *bool
}

type ListDataServiceAuthorizedAppsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message    *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListDataServiceAuthorizedAppsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServiceAuthorizedAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedAppsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataServiceAuthorizedAppsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceAuthorizedAppsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataServiceAuthorizedAppsResponseBody) GetPageResult() *ListDataServiceAuthorizedAppsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListDataServiceAuthorizedAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceAuthorizedAppsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceAuthorizedAppsResponseBody) SetCode(v string) *ListDataServiceAuthorizedAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBody) SetHttpStatusCode(v int32) *ListDataServiceAuthorizedAppsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBody) SetMessage(v string) *ListDataServiceAuthorizedAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBody) SetPageResult(v *ListDataServiceAuthorizedAppsResponseBodyPageResult) *ListDataServiceAuthorizedAppsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBody) SetRequestId(v string) *ListDataServiceAuthorizedAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBody) SetSuccess(v bool) *ListDataServiceAuthorizedAppsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServiceAuthorizedAppsResponseBodyPageResult struct {
	AuthorizedAppList []*ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList `json:"AuthorizedAppList,omitempty" xml:"AuthorizedAppList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceAuthorizedAppsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedAppsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResult) GetAuthorizedAppList() []*ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	return s.AuthorizedAppList
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResult) SetAuthorizedAppList(v []*ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) *ListDataServiceAuthorizedAppsResponseBodyPageResult {
	s.AuthorizedAppList = v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResult) SetTotalCount(v int32) *ListDataServiceAuthorizedAppsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResult) Validate() error {
	if s.AuthorizedAppList != nil {
		for _, item := range s.AuthorizedAppList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList struct {
	// example:
	//
	// 1022
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 30012101
	ApplyUserId   *string `json:"ApplyUserId,omitempty" xml:"ApplyUserId,omitempty"`
	ApplyUserName *string `json:"ApplyUserName,omitempty" xml:"ApplyUserName,omitempty"`
	// example:
	//
	// 2025-06-30
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// 1121
	Id               *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsProjectManager *bool  `json:"IsProjectManager,omitempty" xml:"IsProjectManager,omitempty"`
	// example:
	//
	// 1121
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// test
	OwnerUserName *string `json:"OwnerUserName,omitempty" xml:"OwnerUserName,omitempty"`
	// example:
	//
	// 0
	PrivilegeAccount *int32 `json:"PrivilegeAccount,omitempty" xml:"PrivilegeAccount,omitempty"`
	// example:
	//
	// 1
	PrivilegeType *int32 `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// example:
	//
	// 102122
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName           *string                                                                                   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RealHasOwnerPrivilege *bool                                                                                     `json:"RealHasOwnerPrivilege,omitempty" xml:"RealHasOwnerPrivilege,omitempty"`
	RealHasPrivilege      *bool                                                                                     `json:"RealHasPrivilege,omitempty" xml:"RealHasPrivilege,omitempty"`
	RemarkForDebugList    []*ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList `json:"RemarkForDebugList,omitempty" xml:"RemarkForDebugList,omitempty" type:"Repeated"`
	Revocable             *bool                                                                                     `json:"Revocable,omitempty" xml:"Revocable,omitempty"`
	// example:
	//
	// -1
	RevocableDetail *int32 `json:"RevocableDetail,omitempty" xml:"RevocableDetail,omitempty"`
}

func (s ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetAppId() *int32 {
	return s.AppId
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetAppName() *string {
	return s.AppName
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetApplyUserId() *string {
	return s.ApplyUserId
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetApplyUserName() *string {
	return s.ApplyUserName
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetId() *int32 {
	return s.Id
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetIsProjectManager() *bool {
	return s.IsProjectManager
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetOwner() *string {
	return s.Owner
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetOwnerUserName() *string {
	return s.OwnerUserName
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetPrivilegeAccount() *int32 {
	return s.PrivilegeAccount
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetPrivilegeType() *int32 {
	return s.PrivilegeType
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetRealHasOwnerPrivilege() *bool {
	return s.RealHasOwnerPrivilege
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetRealHasPrivilege() *bool {
	return s.RealHasPrivilege
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetRemarkForDebugList() []*ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList {
	return s.RemarkForDebugList
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetRevocable() *bool {
	return s.Revocable
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) GetRevocableDetail() *int32 {
	return s.RevocableDetail
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetAppId(v int32) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.AppId = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetAppName(v string) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.AppName = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetApplyUserId(v string) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.ApplyUserId = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetApplyUserName(v string) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.ApplyUserName = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetExpireDate(v string) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.ExpireDate = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetId(v int32) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.Id = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetIsProjectManager(v bool) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.IsProjectManager = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetOwner(v string) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.Owner = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetOwnerUserName(v string) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.OwnerUserName = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetPrivilegeAccount(v int32) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.PrivilegeAccount = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetPrivilegeType(v int32) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.PrivilegeType = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetProjectId(v int32) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetProjectName(v string) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.ProjectName = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetRealHasOwnerPrivilege(v bool) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.RealHasOwnerPrivilege = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetRealHasPrivilege(v bool) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.RealHasPrivilege = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetRemarkForDebugList(v []*ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.RemarkForDebugList = v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetRevocable(v bool) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.Revocable = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) SetRevocableDetail(v int32) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList {
	s.RevocableDetail = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppList) Validate() error {
	if s.RemarkForDebugList != nil {
		for _, item := range s.RemarkForDebugList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList struct {
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList) GetKey() *string {
	return s.Key
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList) GetValue() *string {
	return s.Value
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList) SetKey(v string) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList {
	s.Key = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList) SetValue(v string) *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList {
	s.Value = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponseBodyPageResultAuthorizedAppListRemarkForDebugList) Validate() error {
	return dara.Validate(s)
}
