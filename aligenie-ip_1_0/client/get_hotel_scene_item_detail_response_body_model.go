// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSceneItemDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetHotelSceneItemDetailResponseBody
	GetCode() *int32
	SetMessage(v string) *GetHotelSceneItemDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelSceneItemDetailResponseBody
	GetRequestId() *string
	SetResult(v *GetHotelSceneItemDetailResponseBodyResult) *GetHotelSceneItemDetailResponseBody
	GetResult() *GetHotelSceneItemDetailResponseBodyResult
}

type GetHotelSceneItemDetailResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelSceneItemDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetHotelSceneItemDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSceneItemDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetHotelSceneItemDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelSceneItemDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelSceneItemDetailResponseBody) GetResult() *GetHotelSceneItemDetailResponseBodyResult {
	return s.Result
}

func (s *GetHotelSceneItemDetailResponseBody) SetCode(v int32) *GetHotelSceneItemDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBody) SetMessage(v string) *GetHotelSceneItemDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBody) SetRequestId(v string) *GetHotelSceneItemDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBody) SetResult(v *GetHotelSceneItemDetailResponseBodyResult) *GetHotelSceneItemDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelSceneItemDetailResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelSceneItemDetailResponseBodyResult struct {
	// example:
	//
	// 客用品类
	Category     *string                                                  `json:"Category,omitempty" xml:"Category,omitempty"`
	DialogueList []*GetHotelSceneItemDetailResponseBodyResultDialogueList `json:"DialogueList,omitempty" xml:"DialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/wupin/keyongpinlei/zhijin.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 10336
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 擦鞋布
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 170
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1666168828
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetHotelSceneItemDetailResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSceneItemDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailResponseBodyResult) GetCategory() *string {
	return s.Category
}

func (s *GetHotelSceneItemDetailResponseBodyResult) GetDialogueList() []*GetHotelSceneItemDetailResponseBodyResultDialogueList {
	return s.DialogueList
}

func (s *GetHotelSceneItemDetailResponseBodyResult) GetIcon() *string {
	return s.Icon
}

func (s *GetHotelSceneItemDetailResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetHotelSceneItemDetailResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetHotelSceneItemDetailResponseBodyResult) GetPrice() *int64 {
	return s.Price
}

func (s *GetHotelSceneItemDetailResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetHotelSceneItemDetailResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetHotelSceneItemDetailResponseBodyResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetCategory(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Category = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetDialogueList(v []*GetHotelSceneItemDetailResponseBodyResultDialogueList) *GetHotelSceneItemDetailResponseBodyResult {
	s.DialogueList = v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetIcon(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetId(v int64) *GetHotelSceneItemDetailResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetName(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetPrice(v int64) *GetHotelSceneItemDetailResponseBodyResult {
	s.Price = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetStatus(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetType(v string) *GetHotelSceneItemDetailResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) SetUpdateTime(v int64) *GetHotelSceneItemDetailResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResult) Validate() error {
	if s.DialogueList != nil {
		for _, item := range s.DialogueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotelSceneItemDetailResponseBodyResultDialogueList struct {
	// example:
	//
	// 1666164774
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 336
	DialogueId *string `json:"DialogueId,omitempty" xml:"DialogueId,omitempty"`
	NoAnswer   *string `json:"NoAnswer,omitempty" xml:"NoAnswer,omitempty"`
	// example:
	//
	// 4
	NoAnswerTemplate *string `json:"NoAnswerTemplate,omitempty" xml:"NoAnswerTemplate,omitempty"`
	// example:
	//
	// 0
	Process  *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// 10336
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 1666164774
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	YesAnswer  *string `json:"YesAnswer,omitempty" xml:"YesAnswer,omitempty"`
	// example:
	//
	// 4
	YesAnswerTemplate *string `json:"YesAnswerTemplate,omitempty" xml:"YesAnswerTemplate,omitempty"`
}

func (s GetHotelSceneItemDetailResponseBodyResultDialogueList) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSceneItemDetailResponseBodyResultDialogueList) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetDialogueId() *string {
	return s.DialogueId
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetNoAnswer() *string {
	return s.NoAnswer
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetNoAnswerTemplate() *string {
	return s.NoAnswerTemplate
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetProcess() *int32 {
	return s.Process
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetQuestion() *string {
	return s.Question
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetYesAnswer() *string {
	return s.YesAnswer
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) GetYesAnswerTemplate() *string {
	return s.YesAnswerTemplate
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetCreateTime(v int64) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.CreateTime = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetDialogueId(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.DialogueId = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetNoAnswer(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.NoAnswer = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetNoAnswerTemplate(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.NoAnswerTemplate = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetProcess(v int32) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.Process = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetQuestion(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.Question = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetServiceId(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.ServiceId = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetUpdateTime(v int64) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.UpdateTime = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetYesAnswer(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.YesAnswer = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) SetYesAnswerTemplate(v string) *GetHotelSceneItemDetailResponseBodyResultDialogueList {
	s.YesAnswerTemplate = &v
	return s
}

func (s *GetHotelSceneItemDetailResponseBodyResultDialogueList) Validate() error {
	return dara.Validate(s)
}
