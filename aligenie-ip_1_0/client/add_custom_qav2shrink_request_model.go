// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQAV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswersShrink(v string) *AddCustomQAV2ShrinkRequest
	GetAnswersShrink() *string
	SetHotelId(v string) *AddCustomQAV2ShrinkRequest
	GetHotelId() *string
	SetKeyWordsShrink(v string) *AddCustomQAV2ShrinkRequest
	GetKeyWordsShrink() *string
	SetMajorQuestion(v string) *AddCustomQAV2ShrinkRequest
	GetMajorQuestion() *string
	SetSupplementaryQuestionsShrink(v string) *AddCustomQAV2ShrinkRequest
	GetSupplementaryQuestionsShrink() *string
}

type AddCustomQAV2ShrinkRequest struct {
	// This parameter is required.
	AnswersShrink *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId                      *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWordsShrink               *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	MajorQuestion                *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestionsShrink *string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty"`
}

func (s AddCustomQAV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2ShrinkRequest) GetAnswersShrink() *string {
	return s.AnswersShrink
}

func (s *AddCustomQAV2ShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddCustomQAV2ShrinkRequest) GetKeyWordsShrink() *string {
	return s.KeyWordsShrink
}

func (s *AddCustomQAV2ShrinkRequest) GetMajorQuestion() *string {
	return s.MajorQuestion
}

func (s *AddCustomQAV2ShrinkRequest) GetSupplementaryQuestionsShrink() *string {
	return s.SupplementaryQuestionsShrink
}

func (s *AddCustomQAV2ShrinkRequest) SetAnswersShrink(v string) *AddCustomQAV2ShrinkRequest {
	s.AnswersShrink = &v
	return s
}

func (s *AddCustomQAV2ShrinkRequest) SetHotelId(v string) *AddCustomQAV2ShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *AddCustomQAV2ShrinkRequest) SetKeyWordsShrink(v string) *AddCustomQAV2ShrinkRequest {
	s.KeyWordsShrink = &v
	return s
}

func (s *AddCustomQAV2ShrinkRequest) SetMajorQuestion(v string) *AddCustomQAV2ShrinkRequest {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQAV2ShrinkRequest) SetSupplementaryQuestionsShrink(v string) *AddCustomQAV2ShrinkRequest {
	s.SupplementaryQuestionsShrink = &v
	return s
}

func (s *AddCustomQAV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
