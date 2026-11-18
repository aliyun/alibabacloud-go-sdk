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
	// The structured text composed of the entity name and its properties.
	//
	// example:
	//
	// Entity: Amoxicillin capsules
	//
	// - Property[Exact active ingredient]: Amoxicillin
	//
	// - Property[Manufactured dosage form]: Oral capsule
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The source link of the entity property content.
	//
	// example:
	//
	// https://p.example.cn/0387f345/index?struct_id=VdB%2BLV7w7lHHEp
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
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
