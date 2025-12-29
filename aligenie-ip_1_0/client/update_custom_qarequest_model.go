// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomQARequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswers(v []*string) *UpdateCustomQARequest
	GetAnswers() []*string
	SetCustomQAId(v string) *UpdateCustomQARequest
	GetCustomQAId() *string
	SetHotelId(v string) *UpdateCustomQARequest
	GetHotelId() *string
	SetKeyWords(v []*string) *UpdateCustomQARequest
	GetKeyWords() []*string
	SetMajorQuestion(v string) *UpdateCustomQARequest
	GetMajorQuestion() *string
	SetSupplementaryQuestions(v []*string) *UpdateCustomQARequest
	GetSupplementaryQuestions() []*string
}

type UpdateCustomQARequest struct {
	Answers []*string `json:"Answers,omitempty" xml:"Answers,omitempty" type:"Repeated"`
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
	HotelId  *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWords []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// example:
	//
	// ***
	MajorQuestion          *string   `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestions []*string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty" type:"Repeated"`
}

func (s UpdateCustomQARequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomQARequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomQARequest) GetAnswers() []*string {
	return s.Answers
}

func (s *UpdateCustomQARequest) GetCustomQAId() *string {
	return s.CustomQAId
}

func (s *UpdateCustomQARequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateCustomQARequest) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *UpdateCustomQARequest) GetMajorQuestion() *string {
	return s.MajorQuestion
}

func (s *UpdateCustomQARequest) GetSupplementaryQuestions() []*string {
	return s.SupplementaryQuestions
}

func (s *UpdateCustomQARequest) SetAnswers(v []*string) *UpdateCustomQARequest {
	s.Answers = v
	return s
}

func (s *UpdateCustomQARequest) SetCustomQAId(v string) *UpdateCustomQARequest {
	s.CustomQAId = &v
	return s
}

func (s *UpdateCustomQARequest) SetHotelId(v string) *UpdateCustomQARequest {
	s.HotelId = &v
	return s
}

func (s *UpdateCustomQARequest) SetKeyWords(v []*string) *UpdateCustomQARequest {
	s.KeyWords = v
	return s
}

func (s *UpdateCustomQARequest) SetMajorQuestion(v string) *UpdateCustomQARequest {
	s.MajorQuestion = &v
	return s
}

func (s *UpdateCustomQARequest) SetSupplementaryQuestions(v []*string) *UpdateCustomQARequest {
	s.SupplementaryQuestions = v
	return s
}

func (s *UpdateCustomQARequest) Validate() error {
	return dara.Validate(s)
}
