// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomQAShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswersShrink(v string) *UpdateCustomQAShrinkRequest
	GetAnswersShrink() *string
	SetCustomQAId(v string) *UpdateCustomQAShrinkRequest
	GetCustomQAId() *string
	SetHotelId(v string) *UpdateCustomQAShrinkRequest
	GetHotelId() *string
	SetKeyWordsShrink(v string) *UpdateCustomQAShrinkRequest
	GetKeyWordsShrink() *string
	SetMajorQuestion(v string) *UpdateCustomQAShrinkRequest
	GetMajorQuestion() *string
	SetSupplementaryQuestionsShrink(v string) *UpdateCustomQAShrinkRequest
	GetSupplementaryQuestionsShrink() *string
}

type UpdateCustomQAShrinkRequest struct {
	AnswersShrink *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomQAId *string `json:"CustomQAId,omitempty" xml:"CustomQAId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId        *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWordsShrink *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// example:
	//
	// ***
	MajorQuestion                *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestionsShrink *string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty"`
}

func (s UpdateCustomQAShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomQAShrinkRequest) GetAnswersShrink() *string {
	return s.AnswersShrink
}

func (s *UpdateCustomQAShrinkRequest) GetCustomQAId() *string {
	return s.CustomQAId
}

func (s *UpdateCustomQAShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateCustomQAShrinkRequest) GetKeyWordsShrink() *string {
	return s.KeyWordsShrink
}

func (s *UpdateCustomQAShrinkRequest) GetMajorQuestion() *string {
	return s.MajorQuestion
}

func (s *UpdateCustomQAShrinkRequest) GetSupplementaryQuestionsShrink() *string {
	return s.SupplementaryQuestionsShrink
}

func (s *UpdateCustomQAShrinkRequest) SetAnswersShrink(v string) *UpdateCustomQAShrinkRequest {
	s.AnswersShrink = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetCustomQAId(v string) *UpdateCustomQAShrinkRequest {
	s.CustomQAId = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetHotelId(v string) *UpdateCustomQAShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetKeyWordsShrink(v string) *UpdateCustomQAShrinkRequest {
	s.KeyWordsShrink = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetMajorQuestion(v string) *UpdateCustomQAShrinkRequest {
	s.MajorQuestion = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) SetSupplementaryQuestionsShrink(v string) *UpdateCustomQAShrinkRequest {
	s.SupplementaryQuestionsShrink = &v
	return s
}

func (s *UpdateCustomQAShrinkRequest) Validate() error {
	return dara.Validate(s)
}
