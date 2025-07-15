// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserLevelsOfSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserLevelsOfSkillGroupResponseBody
	GetCode() *string
	SetData(v *ListUserLevelsOfSkillGroupResponseBodyData) *ListUserLevelsOfSkillGroupResponseBody
	GetData() *ListUserLevelsOfSkillGroupResponseBodyData
	SetHttpStatusCode(v int32) *ListUserLevelsOfSkillGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListUserLevelsOfSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUserLevelsOfSkillGroupResponseBody
	GetRequestId() *string
}

type ListUserLevelsOfSkillGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListUserLevelsOfSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserLevelsOfSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserLevelsOfSkillGroupResponseBody) GetData() *ListUserLevelsOfSkillGroupResponseBodyData {
	return s.Data
}

func (s *ListUserLevelsOfSkillGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUserLevelsOfSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserLevelsOfSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetCode(v string) *ListUserLevelsOfSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetData(v *ListUserLevelsOfSkillGroupResponseBodyData) *ListUserLevelsOfSkillGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetHttpStatusCode(v int32) *ListUserLevelsOfSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetMessage(v string) *ListUserLevelsOfSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetRequestId(v string) *ListUserLevelsOfSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserLevelsOfSkillGroupResponseBodyData struct {
	List []*ListUserLevelsOfSkillGroupResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserLevelsOfSkillGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) GetList() []*ListUserLevelsOfSkillGroupResponseBodyDataList {
	return s.List
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) SetList(v []*ListUserLevelsOfSkillGroupResponseBodyDataList) *ListUserLevelsOfSkillGroupResponseBodyData {
	s.List = v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) SetPageNumber(v int32) *ListUserLevelsOfSkillGroupResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) SetPageSize(v int32) *ListUserLevelsOfSkillGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) SetTotalCount(v int32) *ListUserLevelsOfSkillGroupResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUserLevelsOfSkillGroupResponseBodyDataList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	RamId     *string `json:"RamId,omitempty" xml:"RamId,omitempty"`
	// example:
	//
	// Admin@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// Admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// skillgroup
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	// example:
	//
	// 5
	SkillLevel *int32 `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserLevelsOfSkillGroupResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) GetLoginName() *string {
	return s.LoginName
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) GetRamId() *string {
	return s.RamId
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) GetRoleId() *string {
	return s.RoleId
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) GetRoleName() *string {
	return s.RoleName
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) GetSkillLevel() *int32 {
	return s.SkillLevel
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) GetUserId() *string {
	return s.UserId
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetDisplayName(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetLoginName(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.LoginName = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetRamId(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.RamId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetRoleId(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.RoleId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetRoleName(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.RoleName = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetSkillGroupId(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetSkillGroupName(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetSkillLevel(v int32) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.SkillLevel = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetUserId(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
