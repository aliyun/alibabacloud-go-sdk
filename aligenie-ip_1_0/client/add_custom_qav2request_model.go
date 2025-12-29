// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQAV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAnswers(v []*string) *AddCustomQAV2Request
	GetAnswers() []*string
	SetHotelId(v string) *AddCustomQAV2Request
	GetHotelId() *string
	SetKeyWords(v []*string) *AddCustomQAV2Request
	GetKeyWords() []*string
	SetMajorQuestion(v string) *AddCustomQAV2Request
	GetMajorQuestion() *string
	SetSupplementaryQuestions(v []*string) *AddCustomQAV2Request
	GetSupplementaryQuestions() []*string
}

type AddCustomQAV2Request struct {
	// This parameter is required.
	Answers []*string `json:"Answers,omitempty" xml:"Answers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId                *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWords               []*string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	MajorQuestion          *string   `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	SupplementaryQuestions []*string `json:"SupplementaryQuestions,omitempty" xml:"SupplementaryQuestions,omitempty" type:"Repeated"`
}

func (s AddCustomQAV2Request) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAV2Request) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2Request) GetAnswers() []*string {
	return s.Answers
}

func (s *AddCustomQAV2Request) GetHotelId() *string {
	return s.HotelId
}

func (s *AddCustomQAV2Request) GetKeyWords() []*string {
	return s.KeyWords
}

func (s *AddCustomQAV2Request) GetMajorQuestion() *string {
	return s.MajorQuestion
}

func (s *AddCustomQAV2Request) GetSupplementaryQuestions() []*string {
	return s.SupplementaryQuestions
}

func (s *AddCustomQAV2Request) SetAnswers(v []*string) *AddCustomQAV2Request {
	s.Answers = v
	return s
}

func (s *AddCustomQAV2Request) SetHotelId(v string) *AddCustomQAV2Request {
	s.HotelId = &v
	return s
}

func (s *AddCustomQAV2Request) SetKeyWords(v []*string) *AddCustomQAV2Request {
	s.KeyWords = v
	return s
}

func (s *AddCustomQAV2Request) SetMajorQuestion(v string) *AddCustomQAV2Request {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQAV2Request) SetSupplementaryQuestions(v []*string) *AddCustomQAV2Request {
	s.SupplementaryQuestions = v
	return s
}

func (s *AddCustomQAV2Request) Validate() error {
	return dara.Validate(s)
}
