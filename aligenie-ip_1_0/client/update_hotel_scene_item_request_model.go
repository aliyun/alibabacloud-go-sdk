// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *UpdateHotelSceneItemRequest
	GetHotelId() *string
	SetUpdateHotelSceneOperateReq(v *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) *UpdateHotelSceneItemRequest
	GetUpdateHotelSceneOperateReq() *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq
	SetUpdateHotelSceneReq(v *UpdateHotelSceneItemRequestUpdateHotelSceneReq) *UpdateHotelSceneItemRequest
	GetUpdateHotelSceneReq() *UpdateHotelSceneItemRequestUpdateHotelSceneReq
}

type UpdateHotelSceneItemRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// updateHotelSceneReq
	//
	// This parameter is required.
	UpdateHotelSceneOperateReq *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq `json:"UpdateHotelSceneOperateReq,omitempty" xml:"UpdateHotelSceneOperateReq,omitempty" type:"Struct"`
	// UpdateHotelSceneReq
	//
	// This parameter is required.
	UpdateHotelSceneReq *UpdateHotelSceneItemRequestUpdateHotelSceneReq `json:"UpdateHotelSceneReq,omitempty" xml:"UpdateHotelSceneReq,omitempty" type:"Struct"`
}

func (s UpdateHotelSceneItemRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateHotelSceneItemRequest) GetUpdateHotelSceneOperateReq() *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq {
	return s.UpdateHotelSceneOperateReq
}

func (s *UpdateHotelSceneItemRequest) GetUpdateHotelSceneReq() *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	return s.UpdateHotelSceneReq
}

func (s *UpdateHotelSceneItemRequest) SetHotelId(v string) *UpdateHotelSceneItemRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelSceneItemRequest) SetUpdateHotelSceneOperateReq(v *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) *UpdateHotelSceneItemRequest {
	s.UpdateHotelSceneOperateReq = v
	return s
}

func (s *UpdateHotelSceneItemRequest) SetUpdateHotelSceneReq(v *UpdateHotelSceneItemRequestUpdateHotelSceneReq) *UpdateHotelSceneItemRequest {
	s.UpdateHotelSceneReq = v
	return s
}

func (s *UpdateHotelSceneItemRequest) Validate() error {
	if s.UpdateHotelSceneOperateReq != nil {
		if err := s.UpdateHotelSceneOperateReq.Validate(); err != nil {
			return err
		}
	}
	if s.UpdateHotelSceneReq != nil {
		if err := s.UpdateHotelSceneReq.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	IsUseTemplateAnswer *bool `json:"IsUseTemplateAnswer,omitempty" xml:"IsUseTemplateAnswer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) GetIsUseTemplateAnswer() *bool {
	return s.IsUseTemplateAnswer
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) GetOperateType() *string {
	return s.OperateType
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) SetIsUseTemplateAnswer(v bool) *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq {
	s.IsUseTemplateAnswer = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) SetOperateType(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq {
	s.OperateType = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneOperateReq) Validate() error {
	return dara.Validate(s)
}

type UpdateHotelSceneItemRequestUpdateHotelSceneReq struct {
	BeyondLimitReply *string `json:"BeyondLimitReply,omitempty" xml:"BeyondLimitReply,omitempty"`
	DeliveryMethod   *string `json:"DeliveryMethod,omitempty" xml:"DeliveryMethod,omitempty"`
	// This parameter is required.
	DialogueList []*UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList `json:"DialogueList,omitempty" xml:"DialogueList,omitempty" type:"Repeated"`
	// icon
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/wupin/keyongpinlei/mianqian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// itemID
	//
	// example:
	//
	// 10337
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LimitNumber   *int64  `json:"LimitNumber,omitempty" xml:"LimitNumber,omitempty"`
	LimitSwitch   *int32  `json:"LimitSwitch,omitempty" xml:"LimitSwitch,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 165
	Price     *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneReq) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneReq) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetBeyondLimitReply() *string {
	return s.BeyondLimitReply
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetDeliveryMethod() *string {
	return s.DeliveryMethod
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetDialogueList() []*UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	return s.DialogueList
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetIcon() *string {
	return s.Icon
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetId() *int64 {
	return s.Id
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetLimitNumber() *int64 {
	return s.LimitNumber
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetLimitSwitch() *int32 {
	return s.LimitSwitch
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetName() *string {
	return s.Name
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetPaymentMethod() *string {
	return s.PaymentMethod
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetPrice() *int64 {
	return s.Price
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetRobotName() *string {
	return s.RobotName
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) GetStatus() *string {
	return s.Status
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetBeyondLimitReply(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.BeyondLimitReply = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetDeliveryMethod(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.DeliveryMethod = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetDialogueList(v []*UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.DialogueList = v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetIcon(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Icon = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetId(v int64) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Id = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetLimitNumber(v int64) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.LimitNumber = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetLimitSwitch(v int32) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.LimitSwitch = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetName(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Name = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetPaymentMethod(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.PaymentMethod = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetPrice(v int64) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Price = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetRobotName(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.RobotName = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) SetStatus(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReq {
	s.Status = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReq) Validate() error {
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

type UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList struct {
	// example:
	//
	// 335
	DialogueId *string `json:"DialogueId,omitempty" xml:"DialogueId,omitempty"`
	// example:
	//
	// 对不起，暂时不提供此物品
	NoAnswer *string `json:"NoAnswer,omitempty" xml:"NoAnswer,omitempty"`
	// example:
	//
	// 4
	NoAnswerTemplate *string `json:"NoAnswerTemplate,omitempty" xml:"NoAnswerTemplate,omitempty"`
	// example:
	//
	// 0
	Process  *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// itemId
	//
	// example:
	//
	// 10337
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 纸巾1.5元，请问需要么？
	YesAnswer *string `json:"YesAnswer,omitempty" xml:"YesAnswer,omitempty"`
	// example:
	//
	// 4
	YesAnswerTemplate *string `json:"YesAnswerTemplate,omitempty" xml:"YesAnswerTemplate,omitempty"`
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GetDialogueId() *string {
	return s.DialogueId
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GetNoAnswer() *string {
	return s.NoAnswer
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GetNoAnswerTemplate() *string {
	return s.NoAnswerTemplate
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GetProcess() *int32 {
	return s.Process
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GetQuestion() *string {
	return s.Question
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GetYesAnswer() *string {
	return s.YesAnswer
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) GetYesAnswerTemplate() *string {
	return s.YesAnswerTemplate
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetDialogueId(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.DialogueId = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetNoAnswer(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.NoAnswer = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetNoAnswerTemplate(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.NoAnswerTemplate = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetProcess(v int32) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.Process = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetQuestion(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.Question = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetServiceId(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.ServiceId = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetYesAnswer(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.YesAnswer = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) SetYesAnswerTemplate(v string) *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList {
	s.YesAnswerTemplate = &v
	return s
}

func (s *UpdateHotelSceneItemRequestUpdateHotelSceneReqDialogueList) Validate() error {
	return dara.Validate(s)
}
