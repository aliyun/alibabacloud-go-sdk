// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAnswerInput interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *MedicalAnswerInput
	GetQuery() *string
}

type MedicalAnswerInput struct {
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s MedicalAnswerInput) String() string {
	return dara.Prettify(s)
}

func (s MedicalAnswerInput) GoString() string {
	return s.String()
}

func (s *MedicalAnswerInput) GetQuery() *string {
	return s.Query
}

func (s *MedicalAnswerInput) SetQuery(v string) *MedicalAnswerInput {
	s.Query = &v
	return s
}

func (s *MedicalAnswerInput) Validate() error {
	return dara.Validate(s)
}
