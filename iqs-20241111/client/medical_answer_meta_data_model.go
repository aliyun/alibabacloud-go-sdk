// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAnswerMetaData interface {
	dara.Model
	String() string
	GoString() string
}

type MedicalAnswerMetaData struct {
}

func (s MedicalAnswerMetaData) String() string {
	return dara.Prettify(s)
}

func (s MedicalAnswerMetaData) GoString() string {
	return s.String()
}

func (s *MedicalAnswerMetaData) Validate() error {
	return dara.Validate(s)
}
