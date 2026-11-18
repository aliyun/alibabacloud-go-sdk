// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAnswerRaDoc interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *MedicalAnswerRaDoc
	GetContent() *string
	SetRawUrl(v string) *MedicalAnswerRaDoc
	GetRawUrl() *string
	SetTitle(v string) *MedicalAnswerRaDoc
	GetTitle() *string
	SetType(v string) *MedicalAnswerRaDoc
	GetType() *string
}

type MedicalAnswerRaDoc struct {
	// The content summary of the reference source. Summaries can be returned for medical web pages and drug instructions.
	//
	// example:
	//
	// Stomach pain can be caused by various factors, including gastrointestinal spasms, inflammation, urinary tract stones, and menstrual pain. It can be relieved by applying a warm compress to the abdomen or treated by taking pain-relieving medication
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The URL of the reference source. This can be an empty string for sources such as papers, clinical guidelines, and books.
	//
	// example:
	//
	// https://www.example.com/article/4CD364Mm8hh.html
	RawUrl *string `json:"rawUrl,omitempty" xml:"rawUrl,omitempty"`
	// The title of the reference source.
	//
	// example:
	//
	// How to quickly and easily relieve stomach pain
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The type of the reference source, such as medical web pages, medical books, papers, or clinical guidelines.
	//
	// example:
	//
	// med_web_source
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalAnswerRaDoc) String() string {
	return dara.Prettify(s)
}

func (s MedicalAnswerRaDoc) GoString() string {
	return s.String()
}

func (s *MedicalAnswerRaDoc) GetContent() *string {
	return s.Content
}

func (s *MedicalAnswerRaDoc) GetRawUrl() *string {
	return s.RawUrl
}

func (s *MedicalAnswerRaDoc) GetTitle() *string {
	return s.Title
}

func (s *MedicalAnswerRaDoc) GetType() *string {
	return s.Type
}

func (s *MedicalAnswerRaDoc) SetContent(v string) *MedicalAnswerRaDoc {
	s.Content = &v
	return s
}

func (s *MedicalAnswerRaDoc) SetRawUrl(v string) *MedicalAnswerRaDoc {
	s.RawUrl = &v
	return s
}

func (s *MedicalAnswerRaDoc) SetTitle(v string) *MedicalAnswerRaDoc {
	s.Title = &v
	return s
}

func (s *MedicalAnswerRaDoc) SetType(v string) *MedicalAnswerRaDoc {
	s.Type = &v
	return s
}

func (s *MedicalAnswerRaDoc) Validate() error {
	return dara.Validate(s)
}
