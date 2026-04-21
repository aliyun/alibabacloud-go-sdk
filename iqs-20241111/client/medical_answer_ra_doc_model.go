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
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	RawUrl  *string `json:"rawUrl,omitempty" xml:"rawUrl,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
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
