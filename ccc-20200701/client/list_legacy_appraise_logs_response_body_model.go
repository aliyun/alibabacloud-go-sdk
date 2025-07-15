// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyAppraiseLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLegacyAppraiseLogsResponseBody
	GetCode() *string
	SetData(v *ListLegacyAppraiseLogsResponseBodyData) *ListLegacyAppraiseLogsResponseBody
	GetData() *ListLegacyAppraiseLogsResponseBodyData
	SetHttpStatusCode(v int32) *ListLegacyAppraiseLogsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListLegacyAppraiseLogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLegacyAppraiseLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLegacyAppraiseLogsResponseBody
	GetSuccess() *bool
}

type ListLegacyAppraiseLogsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListLegacyAppraiseLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A13BB835-94AA-4E55-8D9E-5EA585CE6555"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLegacyAppraiseLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAppraiseLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLegacyAppraiseLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLegacyAppraiseLogsResponseBody) GetData() *ListLegacyAppraiseLogsResponseBodyData {
	return s.Data
}

func (s *ListLegacyAppraiseLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListLegacyAppraiseLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLegacyAppraiseLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLegacyAppraiseLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLegacyAppraiseLogsResponseBody) SetCode(v string) *ListLegacyAppraiseLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBody) SetData(v *ListLegacyAppraiseLogsResponseBodyData) *ListLegacyAppraiseLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBody) SetHttpStatusCode(v int32) *ListLegacyAppraiseLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBody) SetMessage(v string) *ListLegacyAppraiseLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBody) SetRequestId(v string) *ListLegacyAppraiseLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBody) SetSuccess(v bool) *ListLegacyAppraiseLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLegacyAppraiseLogsResponseBodyData struct {
	List []*ListLegacyAppraiseLogsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 18
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLegacyAppraiseLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAppraiseLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLegacyAppraiseLogsResponseBodyData) GetList() []*ListLegacyAppraiseLogsResponseBodyDataList {
	return s.List
}

func (s *ListLegacyAppraiseLogsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLegacyAppraiseLogsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLegacyAppraiseLogsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLegacyAppraiseLogsResponseBodyData) SetList(v []*ListLegacyAppraiseLogsResponseBodyDataList) *ListLegacyAppraiseLogsResponseBodyData {
	s.List = v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyData) SetPageNumber(v int32) *ListLegacyAppraiseLogsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyData) SetPageSize(v int32) *ListLegacyAppraiseLogsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyData) SetTotalCount(v int32) *ListLegacyAppraiseLogsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListLegacyAppraiseLogsResponseBodyDataList struct {
	// example:
	//
	// 3786929579
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// example:
	//
	// Outbound
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// example:
	//
	// 10505
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// {}
	KeyMarkRelation *string `json:"KeyMarkRelation,omitempty" xml:"KeyMarkRelation,omitempty"`
	Note            *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ParentNote      *string `json:"ParentNote,omitempty" xml:"ParentNote,omitempty"`
	// example:
	//
	// 2
	PressKey *string `json:"PressKey,omitempty" xml:"PressKey,omitempty"`
	// example:
	//
	// 28036411123456****
	RamId *string `json:"RamId,omitempty" xml:"RamId,omitempty"`
	// example:
	//
	// test@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// 2021-12-03T10:15:30
	StatisticDate *string `json:"StatisticDate,omitempty" xml:"StatisticDate,omitempty"`
	// example:
	//
	// Launch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLegacyAppraiseLogsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAppraiseLogsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetAcid() *string {
	return s.Acid
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetContactType() *string {
	return s.ContactType
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetKeyMarkRelation() *string {
	return s.KeyMarkRelation
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetNote() *string {
	return s.Note
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetParentNote() *string {
	return s.ParentNote
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetPressKey() *string {
	return s.PressKey
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetRamId() *string {
	return s.RamId
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetStatisticDate() *string {
	return s.StatisticDate
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) GetType() *string {
	return s.Type
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetAcid(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.Acid = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetContactType(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.ContactType = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetId(v int64) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetInstanceId(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetKeyMarkRelation(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.KeyMarkRelation = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetNote(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.Note = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetParentNote(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.ParentNote = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetPressKey(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.PressKey = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetRamId(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.RamId = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetSkillGroupId(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetStatisticDate(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.StatisticDate = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) SetType(v string) *ListLegacyAppraiseLogsResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *ListLegacyAppraiseLogsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
