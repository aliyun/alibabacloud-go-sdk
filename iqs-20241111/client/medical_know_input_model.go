// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalKnowInput interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *MedicalKnowInput
	GetQuery() *string
	SetType(v string) *MedicalKnowInput
	GetType() *string
}

type MedicalKnowInput struct {
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MedicalKnowInput) String() string {
	return dara.Prettify(s)
}

func (s MedicalKnowInput) GoString() string {
	return s.String()
}

func (s *MedicalKnowInput) GetQuery() *string {
	return s.Query
}

func (s *MedicalKnowInput) GetType() *string {
	return s.Type
}

func (s *MedicalKnowInput) SetQuery(v string) *MedicalKnowInput {
	s.Query = &v
	return s
}

func (s *MedicalKnowInput) SetType(v string) *MedicalKnowInput {
	s.Type = &v
	return s
}

func (s *MedicalKnowInput) Validate() error {
	return dara.Validate(s)
}
