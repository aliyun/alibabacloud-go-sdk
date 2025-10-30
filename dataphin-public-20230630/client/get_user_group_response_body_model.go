// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetUserGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserGroupResponseBody
	GetSuccess() *bool
	SetUserGroupInfo(v *GetUserGroupResponseBodyUserGroupInfo) *GetUserGroupResponseBody
	GetUserGroupInfo() *GetUserGroupResponseBodyUserGroupInfo
}

type GetUserGroupResponseBody struct {
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success       *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	UserGroupInfo *GetUserGroupResponseBodyUserGroupInfo `json:"UserGroupInfo,omitempty" xml:"UserGroupInfo,omitempty" type:"Struct"`
}

func (s GetUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserGroupResponseBody) GetUserGroupInfo() *GetUserGroupResponseBodyUserGroupInfo {
	return s.UserGroupInfo
}

func (s *GetUserGroupResponseBody) SetCode(v string) *GetUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserGroupResponseBody) SetHttpStatusCode(v int32) *GetUserGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserGroupResponseBody) SetMessage(v string) *GetUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserGroupResponseBody) SetRequestId(v string) *GetUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGroupResponseBody) SetSuccess(v bool) *GetUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserGroupResponseBody) SetUserGroupInfo(v *GetUserGroupResponseBodyUserGroupInfo) *GetUserGroupResponseBody {
	s.UserGroupInfo = v
	return s
}

func (s *GetUserGroupResponseBody) Validate() error {
	if s.UserGroupInfo != nil {
		if err := s.UserGroupInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserGroupResponseBodyUserGroupInfo struct {
	// example:
	//
	// true
	Active    *bool                                             `json:"Active,omitempty" xml:"Active,omitempty"`
	AdminList []*GetUserGroupResponseBodyUserGroupInfoAdminList `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1313213
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	MyRole *string `json:"MyRole,omitempty" xml:"MyRole,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroupInfo) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroupInfo) GetActive() *bool {
	return s.Active
}

func (s *GetUserGroupResponseBodyUserGroupInfo) GetAdminList() []*GetUserGroupResponseBodyUserGroupInfoAdminList {
	return s.AdminList
}

func (s *GetUserGroupResponseBodyUserGroupInfo) GetDescription() *string {
	return s.Description
}

func (s *GetUserGroupResponseBodyUserGroupInfo) GetId() *string {
	return s.Id
}

func (s *GetUserGroupResponseBodyUserGroupInfo) GetMyRole() *string {
	return s.MyRole
}

func (s *GetUserGroupResponseBodyUserGroupInfo) GetName() *string {
	return s.Name
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetActive(v bool) *GetUserGroupResponseBodyUserGroupInfo {
	s.Active = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetAdminList(v []*GetUserGroupResponseBodyUserGroupInfoAdminList) *GetUserGroupResponseBodyUserGroupInfo {
	s.AdminList = v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetDescription(v string) *GetUserGroupResponseBodyUserGroupInfo {
	s.Description = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetId(v string) *GetUserGroupResponseBodyUserGroupInfo {
	s.Id = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetMyRole(v string) *GetUserGroupResponseBodyUserGroupInfo {
	s.MyRole = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetName(v string) *GetUserGroupResponseBodyUserGroupInfo {
	s.Name = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) Validate() error {
	if s.AdminList != nil {
		for _, item := range s.AdminList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserGroupResponseBodyUserGroupInfoAdminList struct {
	// example:
	//
	// xx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 131313
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroupInfoAdminList) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroupInfoAdminList) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) GetAccountName() *string {
	return s.AccountName
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) GetId() *string {
	return s.Id
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) SetAccountName(v string) *GetUserGroupResponseBodyUserGroupInfoAdminList {
	s.AccountName = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) SetDisplayName(v string) *GetUserGroupResponseBodyUserGroupInfoAdminList {
	s.DisplayName = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) SetId(v string) *GetUserGroupResponseBodyUserGroupInfoAdminList {
	s.Id = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) Validate() error {
	return dara.Validate(s)
}
