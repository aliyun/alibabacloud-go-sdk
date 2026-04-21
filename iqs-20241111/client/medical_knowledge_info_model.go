// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalKnowledgeInfo interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *MedicalKnowledgeInfo
	GetContent() *string
	SetUrl(v string) *MedicalKnowledgeInfo
	GetUrl() *string
}

type MedicalKnowledgeInfo struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Url     *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MedicalKnowledgeInfo) String() string {
	return dara.Prettify(s)
}

func (s MedicalKnowledgeInfo) GoString() string {
	return s.String()
}

func (s *MedicalKnowledgeInfo) GetContent() *string {
	return s.Content
}

func (s *MedicalKnowledgeInfo) GetUrl() *string {
	return s.Url
}

func (s *MedicalKnowledgeInfo) SetContent(v string) *MedicalKnowledgeInfo {
	s.Content = &v
	return s
}

func (s *MedicalKnowledgeInfo) SetUrl(v string) *MedicalKnowledgeInfo {
	s.Url = &v
	return s
}

func (s *MedicalKnowledgeInfo) Validate() error {
	return dara.Validate(s)
}
