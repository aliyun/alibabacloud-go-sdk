// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBriefSkillGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBriefSkillGroupsResponseBody
	GetCode() *string
	SetData(v *ListBriefSkillGroupsResponseBodyData) *ListBriefSkillGroupsResponseBody
	GetData() *ListBriefSkillGroupsResponseBodyData
	SetHttpStatusCode(v int32) *ListBriefSkillGroupsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBriefSkillGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBriefSkillGroupsResponseBody
	GetRequestId() *string
}

type ListBriefSkillGroupsResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListBriefSkillGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3969FC68-CEC2-4398-B76A-60D2F7EDEBAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBriefSkillGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBriefSkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBriefSkillGroupsResponseBody) GetData() *ListBriefSkillGroupsResponseBodyData {
	return s.Data
}

func (s *ListBriefSkillGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBriefSkillGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBriefSkillGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBriefSkillGroupsResponseBody) SetCode(v string) *ListBriefSkillGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBody) SetData(v *ListBriefSkillGroupsResponseBodyData) *ListBriefSkillGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListBriefSkillGroupsResponseBody) SetHttpStatusCode(v int32) *ListBriefSkillGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBody) SetMessage(v string) *ListBriefSkillGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBody) SetRequestId(v string) *ListBriefSkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBriefSkillGroupsResponseBodyData struct {
	List []*ListBriefSkillGroupsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBriefSkillGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBriefSkillGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsResponseBodyData) GetList() []*ListBriefSkillGroupsResponseBodyDataList {
	return s.List
}

func (s *ListBriefSkillGroupsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBriefSkillGroupsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBriefSkillGroupsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBriefSkillGroupsResponseBodyData) SetList(v []*ListBriefSkillGroupsResponseBodyDataList) *ListBriefSkillGroupsResponseBodyData {
	s.List = v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyData) SetPageNumber(v int32) *ListBriefSkillGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyData) SetPageSize(v int32) *ListBriefSkillGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyData) SetTotalCount(v int32) *ListBriefSkillGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListBriefSkillGroupsResponseBodyDataList struct {
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
	// 10
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s ListBriefSkillGroupsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListBriefSkillGroupsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListBriefSkillGroupsResponseBodyDataList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListBriefSkillGroupsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBriefSkillGroupsResponseBodyDataList) GetMediaType() *string {
	return s.MediaType
}

func (s *ListBriefSkillGroupsResponseBodyDataList) GetPhoneNumberCount() *int32 {
	return s.PhoneNumberCount
}

func (s *ListBriefSkillGroupsResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListBriefSkillGroupsResponseBodyDataList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListBriefSkillGroupsResponseBodyDataList) GetUserCount() *int32 {
	return s.UserCount
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetDescription(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetDisplayName(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetInstanceId(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetMediaType(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.MediaType = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetPhoneNumberCount(v int32) *ListBriefSkillGroupsResponseBodyDataList {
	s.PhoneNumberCount = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetSkillGroupId(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetSkillGroupName(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetUserCount(v int32) *ListBriefSkillGroupsResponseBodyDataList {
	s.UserCount = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
