// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillLevelsOfUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSkillLevelsOfUserResponseBody
	GetCode() *string
	SetData(v *ListSkillLevelsOfUserResponseBodyData) *ListSkillLevelsOfUserResponseBody
	GetData() *ListSkillLevelsOfUserResponseBodyData
	SetHttpStatusCode(v int32) *ListSkillLevelsOfUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSkillLevelsOfUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSkillLevelsOfUserResponseBody
	GetRequestId() *string
}

type ListSkillLevelsOfUserResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListSkillLevelsOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListSkillLevelsOfUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillLevelsOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSkillLevelsOfUserResponseBody) GetData() *ListSkillLevelsOfUserResponseBodyData {
	return s.Data
}

func (s *ListSkillLevelsOfUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSkillLevelsOfUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSkillLevelsOfUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillLevelsOfUserResponseBody) SetCode(v string) *ListSkillLevelsOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBody) SetData(v *ListSkillLevelsOfUserResponseBodyData) *ListSkillLevelsOfUserResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillLevelsOfUserResponseBody) SetHttpStatusCode(v int32) *ListSkillLevelsOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBody) SetMessage(v string) *ListSkillLevelsOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBody) SetRequestId(v string) *ListSkillLevelsOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSkillLevelsOfUserResponseBodyData struct {
	List []*ListSkillLevelsOfUserResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s ListSkillLevelsOfUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillLevelsOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserResponseBodyData) GetList() []*ListSkillLevelsOfUserResponseBodyDataList {
	return s.List
}

func (s *ListSkillLevelsOfUserResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillLevelsOfUserResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillLevelsOfUserResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillLevelsOfUserResponseBodyData) SetList(v []*ListSkillLevelsOfUserResponseBodyDataList) *ListSkillLevelsOfUserResponseBodyData {
	s.List = v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyData) SetPageNumber(v int32) *ListSkillLevelsOfUserResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyData) SetPageSize(v int32) *ListSkillLevelsOfUserResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyData) SetTotalCount(v int32) *ListSkillLevelsOfUserResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListSkillLevelsOfUserResponseBodyDataList struct {
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
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
	SkillLevel *string `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
}

func (s ListSkillLevelsOfUserResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListSkillLevelsOfUserResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) GetMediaType() *string {
	return s.MediaType
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) GetSkillLevel() *string {
	return s.SkillLevel
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) SetMediaType(v string) *ListSkillLevelsOfUserResponseBodyDataList {
	s.MediaType = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) SetSkillGroupId(v string) *ListSkillLevelsOfUserResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) SetSkillGroupName(v string) *ListSkillLevelsOfUserResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) SetSkillLevel(v string) *ListSkillLevelsOfUserResponseBodyDataList {
	s.SkillLevel = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
