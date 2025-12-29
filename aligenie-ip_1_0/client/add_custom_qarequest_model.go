// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQARequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswers(v []*string) *AddCustomQARequest
	GetAnswers() []*string
	SetHotelId(v string) *AddCustomQARequest
	GetHotelId() *string
	SetKeyWords(v []*string) *AddCustomQARequest
	GetKeyWords() []*string
	SetMajorQuestion(v string) *AddCustomQARequest
	GetMajorQuestion() *string
	SetSupplementaryQuestions(v []*string) *AddCustomQARequest
	GetSupplementaryQuestions() []*string
}

type AddCustomQARequest struct {
	Answers []*string `json:"Answers,omitempty" xml:"Answers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId  *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWords []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// example:
	//
	// ***
	MajorQuestion          *string   `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestions []*string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty" type:"Repeated"`
}

func (s AddCustomQARequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQARequest) GoString() string {
	return s.String()
}

func (s *AddCustomQARequest) GetAnswers() []*string {
	return s.Answers
}

func (s *AddCustomQARequest) GetHotelId() *string {
	return s.HotelId
}

func (s *AddCustomQARequest) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *AddCustomQARequest) GetMajorQuestion() *string {
	return s.MajorQuestion
}

func (s *AddCustomQARequest) GetSupplementaryQuestions() []*string {
	return s.SupplementaryQuestions
}

func (s *AddCustomQARequest) SetAnswers(v []*string) *AddCustomQARequest {
	s.Answers = v
	return s
}

func (s *AddCustomQARequest) SetHotelId(v string) *AddCustomQARequest {
	s.HotelId = &v
	return s
}

func (s *AddCustomQARequest) SetKeyWords(v []*string) *AddCustomQARequest {
	s.KeyWords = v
	return s
}

func (s *AddCustomQARequest) SetMajorQuestion(v string) *AddCustomQARequest {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQARequest) SetSupplementaryQuestions(v []*string) *AddCustomQARequest {
	s.SupplementaryQuestions = v
	return s
}

func (s *AddCustomQARequest) Validate() error {
	return dara.Validate(s)
}
