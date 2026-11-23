// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalModelInferenceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type MedicalModelInferenceRequest struct {
}

func (s MedicalModelInferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceRequest) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceRequest) Validate() error {
	return dara.Validate(s)
}
