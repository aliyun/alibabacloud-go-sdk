// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaqResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *DescribeFaqResponseBody
	GetCategoryId() *int64
	SetCreateTime(v string) *DescribeFaqResponseBody
	GetCreateTime() *string
	SetCreateUserName(v string) *DescribeFaqResponseBody
	GetCreateUserName() *string
	SetEffectStatus(v int32) *DescribeFaqResponseBody
	GetEffectStatus() *int32
	SetEndDate(v string) *DescribeFaqResponseBody
	GetEndDate() *string
	SetKnowledgeId(v int64) *DescribeFaqResponseBody
	GetKnowledgeId() *int64
	SetModifyTime(v string) *DescribeFaqResponseBody
	GetModifyTime() *string
	SetModifyUserName(v string) *DescribeFaqResponseBody
	GetModifyUserName() *string
	SetOutlines(v []*DescribeFaqResponseBodyOutlines) *DescribeFaqResponseBody
	GetOutlines() []*DescribeFaqResponseBodyOutlines
	SetRequestId(v string) *DescribeFaqResponseBody
	GetRequestId() *string
	SetSimQuestions(v []*DescribeFaqResponseBodySimQuestions) *DescribeFaqResponseBody
	GetSimQuestions() []*DescribeFaqResponseBodySimQuestions
	SetSolutions(v []*DescribeFaqResponseBodySolutions) *DescribeFaqResponseBody
	GetSolutions() []*DescribeFaqResponseBodySolutions
	SetStartDate(v string) *DescribeFaqResponseBody
	GetStartDate() *string
	SetStatus(v int32) *DescribeFaqResponseBody
	GetStatus() *int32
	SetTitle(v string) *DescribeFaqResponseBody
	GetTitle() *string
}

type DescribeFaqResponseBody struct {
	// example:
	//
	// 30000055617
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2020-11-30T03:03:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test01
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 20
	EffectStatus *int32 `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// example:
	//
	// 2023-04-27T06:08:54Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 2020-12-02T06:35:50Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test01
	ModifyUserName *string                            `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Outlines       []*DescribeFaqResponseBodyOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 8AD9FA10-7780-5E12-B701-13C928524F32
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SimQuestions []*DescribeFaqResponseBodySimQuestions `json:"SimQuestions,omitempty" xml:"SimQuestions,omitempty" type:"Repeated"`
	Solutions    []*DescribeFaqResponseBodySolutions    `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-04-27T07:04:39Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 3
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeFaqResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaqResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponseBody) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DescribeFaqResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFaqResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeFaqResponseBody) GetEffectStatus() *int32 {
	return s.EffectStatus
}

func (s *DescribeFaqResponseBody) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeFaqResponseBody) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *DescribeFaqResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeFaqResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeFaqResponseBody) GetOutlines() []*DescribeFaqResponseBodyOutlines {
	return s.Outlines
}

func (s *DescribeFaqResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFaqResponseBody) GetSimQuestions() []*DescribeFaqResponseBodySimQuestions {
	return s.SimQuestions
}

func (s *DescribeFaqResponseBody) GetSolutions() []*DescribeFaqResponseBodySolutions {
	return s.Solutions
}

func (s *DescribeFaqResponseBody) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeFaqResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeFaqResponseBody) GetTitle() *string {
	return s.Title
}

func (s *DescribeFaqResponseBody) SetCategoryId(v int64) *DescribeFaqResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeFaqResponseBody) SetCreateTime(v string) *DescribeFaqResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeFaqResponseBody) SetCreateUserName(v string) *DescribeFaqResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeFaqResponseBody) SetEffectStatus(v int32) *DescribeFaqResponseBody {
	s.EffectStatus = &v
	return s
}

func (s *DescribeFaqResponseBody) SetEndDate(v string) *DescribeFaqResponseBody {
	s.EndDate = &v
	return s
}

func (s *DescribeFaqResponseBody) SetKnowledgeId(v int64) *DescribeFaqResponseBody {
	s.KnowledgeId = &v
	return s
}

func (s *DescribeFaqResponseBody) SetModifyTime(v string) *DescribeFaqResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeFaqResponseBody) SetModifyUserName(v string) *DescribeFaqResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeFaqResponseBody) SetOutlines(v []*DescribeFaqResponseBodyOutlines) *DescribeFaqResponseBody {
	s.Outlines = v
	return s
}

func (s *DescribeFaqResponseBody) SetRequestId(v string) *DescribeFaqResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaqResponseBody) SetSimQuestions(v []*DescribeFaqResponseBodySimQuestions) *DescribeFaqResponseBody {
	s.SimQuestions = v
	return s
}

func (s *DescribeFaqResponseBody) SetSolutions(v []*DescribeFaqResponseBodySolutions) *DescribeFaqResponseBody {
	s.Solutions = v
	return s
}

func (s *DescribeFaqResponseBody) SetStartDate(v string) *DescribeFaqResponseBody {
	s.StartDate = &v
	return s
}

func (s *DescribeFaqResponseBody) SetStatus(v int32) *DescribeFaqResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFaqResponseBody) SetTitle(v string) *DescribeFaqResponseBody {
	s.Title = &v
	return s
}

func (s *DescribeFaqResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFaqResponseBodyOutlines struct {
	// example:
	//
	// 1000098002
	ConnQuestionId *int64 `json:"ConnQuestionId,omitempty" xml:"ConnQuestionId,omitempty"`
	// example:
	//
	// 2022-05-26T10:24:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-26T18:12:02Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 797
	OutlineId *int64  `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeFaqResponseBodyOutlines) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaqResponseBodyOutlines) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponseBodyOutlines) GetConnQuestionId() *int64 {
	return s.ConnQuestionId
}

func (s *DescribeFaqResponseBodyOutlines) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFaqResponseBodyOutlines) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeFaqResponseBodyOutlines) GetOutlineId() *int64 {
	return s.OutlineId
}

func (s *DescribeFaqResponseBodyOutlines) GetTitle() *string {
	return s.Title
}

func (s *DescribeFaqResponseBodyOutlines) SetConnQuestionId(v int64) *DescribeFaqResponseBodyOutlines {
	s.ConnQuestionId = &v
	return s
}

func (s *DescribeFaqResponseBodyOutlines) SetCreateTime(v string) *DescribeFaqResponseBodyOutlines {
	s.CreateTime = &v
	return s
}

func (s *DescribeFaqResponseBodyOutlines) SetModifyTime(v string) *DescribeFaqResponseBodyOutlines {
	s.ModifyTime = &v
	return s
}

func (s *DescribeFaqResponseBodyOutlines) SetOutlineId(v int64) *DescribeFaqResponseBodyOutlines {
	s.OutlineId = &v
	return s
}

func (s *DescribeFaqResponseBodyOutlines) SetTitle(v string) *DescribeFaqResponseBodyOutlines {
	s.Title = &v
	return s
}

func (s *DescribeFaqResponseBodyOutlines) Validate() error {
	return dara.Validate(s)
}

type DescribeFaqResponseBodySimQuestions struct {
	// example:
	//
	// 2022-05-26T10:24:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-29T03:55:07Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 10000000581
	SimQuestionId *int64  `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeFaqResponseBodySimQuestions) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaqResponseBodySimQuestions) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponseBodySimQuestions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFaqResponseBodySimQuestions) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeFaqResponseBodySimQuestions) GetSimQuestionId() *int64 {
	return s.SimQuestionId
}

func (s *DescribeFaqResponseBodySimQuestions) GetTitle() *string {
	return s.Title
}

func (s *DescribeFaqResponseBodySimQuestions) SetCreateTime(v string) *DescribeFaqResponseBodySimQuestions {
	s.CreateTime = &v
	return s
}

func (s *DescribeFaqResponseBodySimQuestions) SetModifyTime(v string) *DescribeFaqResponseBodySimQuestions {
	s.ModifyTime = &v
	return s
}

func (s *DescribeFaqResponseBodySimQuestions) SetSimQuestionId(v int64) *DescribeFaqResponseBodySimQuestions {
	s.SimQuestionId = &v
	return s
}

func (s *DescribeFaqResponseBodySimQuestions) SetTitle(v string) *DescribeFaqResponseBodySimQuestions {
	s.Title = &v
	return s
}

func (s *DescribeFaqResponseBodySimQuestions) Validate() error {
	return dara.Validate(s)
}

type DescribeFaqResponseBodySolutions struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 0
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 2022-05-26T10:24:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-29T07:07:13Z
	ModifyTime       *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	PerspectiveCodes []*string `json:"PerspectiveCodes,omitempty" xml:"PerspectiveCodes,omitempty" type:"Repeated"`
	PlainText        *string   `json:"PlainText,omitempty" xml:"PlainText,omitempty"`
	// example:
	//
	// 10000003071
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s DescribeFaqResponseBodySolutions) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaqResponseBodySolutions) GoString() string {
	return s.String()
}

func (s *DescribeFaqResponseBodySolutions) GetContent() *string {
	return s.Content
}

func (s *DescribeFaqResponseBodySolutions) GetContentType() *int32 {
	return s.ContentType
}

func (s *DescribeFaqResponseBodySolutions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFaqResponseBodySolutions) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeFaqResponseBodySolutions) GetPerspectiveCodes() []*string {
	return s.PerspectiveCodes
}

func (s *DescribeFaqResponseBodySolutions) GetPlainText() *string {
	return s.PlainText
}

func (s *DescribeFaqResponseBodySolutions) GetSolutionId() *int64 {
	return s.SolutionId
}

func (s *DescribeFaqResponseBodySolutions) SetContent(v string) *DescribeFaqResponseBodySolutions {
	s.Content = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetContentType(v int32) *DescribeFaqResponseBodySolutions {
	s.ContentType = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetCreateTime(v string) *DescribeFaqResponseBodySolutions {
	s.CreateTime = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetModifyTime(v string) *DescribeFaqResponseBodySolutions {
	s.ModifyTime = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetPerspectiveCodes(v []*string) *DescribeFaqResponseBodySolutions {
	s.PerspectiveCodes = v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetPlainText(v string) *DescribeFaqResponseBodySolutions {
	s.PlainText = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) SetSolutionId(v int64) *DescribeFaqResponseBodySolutions {
	s.SolutionId = &v
	return s
}

func (s *DescribeFaqResponseBodySolutions) Validate() error {
	return dara.Validate(s)
}
