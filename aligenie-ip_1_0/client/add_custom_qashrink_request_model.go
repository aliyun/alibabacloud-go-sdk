// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQAShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswersShrink(v string) *AddCustomQAShrinkRequest
	GetAnswersShrink() *string
	SetHotelId(v string) *AddCustomQAShrinkRequest
	GetHotelId() *string
	SetKeyWordsShrink(v string) *AddCustomQAShrinkRequest
	GetKeyWordsShrink() *string
	SetMajorQuestion(v string) *AddCustomQAShrinkRequest
	GetMajorQuestion() *string
	SetSupplementaryQuestionsShrink(v string) *AddCustomQAShrinkRequest
	GetSupplementaryQuestionsShrink() *string
}

type AddCustomQAShrinkRequest struct {
	AnswersShrink *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId        *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWordsShrink *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// example:
	//
	// ***
	MajorQuestion                *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestionsShrink *string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty"`
}

func (s AddCustomQAShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddCustomQAShrinkRequest) GetAnswersShrink() *string {
	return s.AnswersShrink
}

func (s *AddCustomQAShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddCustomQAShrinkRequest) GetKeyWordsShrink() *string {
	return s.KeyWordsShrink
}

func (s *AddCustomQAShrinkRequest) GetMajorQuestion() *string {
	return s.MajorQuestion
}

func (s *AddCustomQAShrinkRequest) GetSupplementaryQuestionsShrink() *string {
	return s.SupplementaryQuestionsShrink
}

func (s *AddCustomQAShrinkRequest) SetAnswersShrink(v string) *AddCustomQAShrinkRequest {
	s.AnswersShrink = &v
	return s
}

func (s *AddCustomQAShrinkRequest) SetHotelId(v string) *AddCustomQAShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *AddCustomQAShrinkRequest) SetKeyWordsShrink(v string) *AddCustomQAShrinkRequest {
	s.KeyWordsShrink = &v
	return s
}

func (s *AddCustomQAShrinkRequest) SetMajorQuestion(v string) *AddCustomQAShrinkRequest {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQAShrinkRequest) SetSupplementaryQuestionsShrink(v string) *AddCustomQAShrinkRequest {
	s.SupplementaryQuestionsShrink = &v
	return s
}

func (s *AddCustomQAShrinkRequest) Validate() error {
	return dara.Validate(s)
}
