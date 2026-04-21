// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAnswerMessage interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *MedicalAnswerMessage
	GetContent() *string
	SetMetaData(v *MedicalAnswerMetaData) *MedicalAnswerMessage
	GetMetaData() *MedicalAnswerMetaData
}

type MedicalAnswerMessage struct {
	Content  *string                `json:"content,omitempty" xml:"content,omitempty"`
	MetaData *MedicalAnswerMetaData `json:"metaData,omitempty" xml:"metaData,omitempty"`
}

func (s MedicalAnswerMessage) String() string {
	return dara.Prettify(s)
}

func (s MedicalAnswerMessage) GoString() string {
	return s.String()
}

func (s *MedicalAnswerMessage) GetContent() *string {
	return s.Content
}

func (s *MedicalAnswerMessage) GetMetaData() *MedicalAnswerMetaData {
	return s.MetaData
}

func (s *MedicalAnswerMessage) SetContent(v string) *MedicalAnswerMessage {
	s.Content = &v
	return s
}

func (s *MedicalAnswerMessage) SetMetaData(v *MedicalAnswerMetaData) *MedicalAnswerMessage {
	s.MetaData = v
	return s
}

func (s *MedicalAnswerMessage) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}
