// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListRolesResponseBodyData) *ListRolesResponseBody
	GetData() []*ListRolesResponseBodyData
	SetHttpStatusCode(v int32) *ListRolesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRolesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRolesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRolesResponseBody
	GetSuccess() *bool
}

type ListRolesResponseBody struct {
	Data []*ListRolesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) GetData() []*ListRolesResponseBodyData {
	return s.Data
}

func (s *ListRolesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRolesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRolesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRolesResponseBody) SetData(v []*ListRolesResponseBodyData) *ListRolesResponseBody {
	s.Data = v
	return s
}

func (s *ListRolesResponseBody) SetHttpStatusCode(v int32) *ListRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRolesResponseBody) SetMessage(v string) *ListRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetSuccess(v bool) *ListRolesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRolesResponseBody) Validate() error {
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

type ListRolesResponseBodyData struct {
	// example:
	//
	// 1
	BuId *int64 `json:"BuId,omitempty" xml:"BuId,omitempty"`
	// example:
	//
	// admin
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2020-01-03T20:25:33Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// Admin
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 0
	RoleGroupId   *int64  `json:"RoleGroupId,omitempty" xml:"RoleGroupId,omitempty"`
	RoleGroupName *string `json:"RoleGroupName,omitempty" xml:"RoleGroupName,omitempty"`
	// example:
	//
	// 1
	RoleId *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListRolesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyData) GetBuId() *int64 {
	return s.BuId
}

func (s *ListRolesResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ListRolesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRolesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListRolesResponseBodyData) GetRoleGroupId() *int64 {
	return s.RoleGroupId
}

func (s *ListRolesResponseBodyData) GetRoleGroupName() *string {
	return s.RoleGroupName
}

func (s *ListRolesResponseBodyData) GetRoleId() *int64 {
	return s.RoleId
}

func (s *ListRolesResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListRolesResponseBodyData) SetBuId(v int64) *ListRolesResponseBodyData {
	s.BuId = &v
	return s
}

func (s *ListRolesResponseBodyData) SetCode(v string) *ListRolesResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListRolesResponseBodyData) SetCreateTime(v string) *ListRolesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListRolesResponseBodyData) SetDescription(v string) *ListRolesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListRolesResponseBodyData) SetRoleGroupId(v int64) *ListRolesResponseBodyData {
	s.RoleGroupId = &v
	return s
}

func (s *ListRolesResponseBodyData) SetRoleGroupName(v string) *ListRolesResponseBodyData {
	s.RoleGroupName = &v
	return s
}

func (s *ListRolesResponseBodyData) SetRoleId(v int64) *ListRolesResponseBodyData {
	s.RoleId = &v
	return s
}

func (s *ListRolesResponseBodyData) SetTitle(v string) *ListRolesResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListRolesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
