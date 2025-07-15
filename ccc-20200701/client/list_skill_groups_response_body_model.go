// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSkillGroupsResponseBody
	GetCode() *string
	SetData(v *ListSkillGroupsResponseBodyData) *ListSkillGroupsResponseBody
	GetData() *ListSkillGroupsResponseBodyData
	SetHttpStatusCode(v int32) *ListSkillGroupsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSkillGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSkillGroupsResponseBody
	GetRequestId() *string
}

type ListSkillGroupsResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListSkillGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSkillGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSkillGroupsResponseBody) GetData() *ListSkillGroupsResponseBodyData {
	return s.Data
}

func (s *ListSkillGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSkillGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSkillGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillGroupsResponseBody) SetCode(v string) *ListSkillGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupsResponseBody) SetData(v *ListSkillGroupsResponseBodyData) *ListSkillGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillGroupsResponseBody) SetHttpStatusCode(v int32) *ListSkillGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSkillGroupsResponseBody) SetMessage(v string) *ListSkillGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupsResponseBody) SetRequestId(v string) *ListSkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupsResponseBodyData struct {
	List []*ListSkillGroupsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListSkillGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsResponseBodyData) GetList() []*ListSkillGroupsResponseBodyDataList {
	return s.List
}

func (s *ListSkillGroupsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillGroupsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillGroupsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillGroupsResponseBodyData) SetList(v []*ListSkillGroupsResponseBodyDataList) *ListSkillGroupsResponseBodyData {
	s.List = v
	return s
}

func (s *ListSkillGroupsResponseBodyData) SetPageNumber(v int32) *ListSkillGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSkillGroupsResponseBodyData) SetPageSize(v int32) *ListSkillGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSkillGroupsResponseBodyData) SetTotalCount(v int32) *ListSkillGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSkillGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupsResponseBodyDataList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MediaType  *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 1
	PhoneNumberCount *int32 `json:"PhoneNumberCount,omitempty" xml:"PhoneNumberCount,omitempty"`
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
	// 2
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s ListSkillGroupsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListSkillGroupsResponseBodyDataList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListSkillGroupsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSkillGroupsResponseBodyDataList) GetMediaType() *string {
	return s.MediaType
}

func (s *ListSkillGroupsResponseBodyDataList) GetPhoneNumberCount() *int32 {
	return s.PhoneNumberCount
}

func (s *ListSkillGroupsResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListSkillGroupsResponseBodyDataList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListSkillGroupsResponseBodyDataList) GetUserCount() *int32 {
	return s.UserCount
}

func (s *ListSkillGroupsResponseBodyDataList) SetDescription(v string) *ListSkillGroupsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetDisplayName(v string) *ListSkillGroupsResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetInstanceId(v string) *ListSkillGroupsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetMediaType(v string) *ListSkillGroupsResponseBodyDataList {
	s.MediaType = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetPhoneNumberCount(v int32) *ListSkillGroupsResponseBodyDataList {
	s.PhoneNumberCount = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetSkillGroupId(v string) *ListSkillGroupsResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetSkillGroupName(v string) *ListSkillGroupsResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetUserCount(v int32) *ListSkillGroupsResponseBodyDataList {
	s.UserCount = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
