// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedTextToImageQueryPreModelInferenceJobInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInferenceJobId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest
	GetInferenceJobId() *string
}

type PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// girl-xxxx-xxxx-xxxx-xxxx
	InferenceJobId *string `json:"inferenceJobId,omitempty" xml:"inferenceJobId,omitempty"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) GetInferenceJobId() *string {
	return s.InferenceJobId
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) SetInferenceJobId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest {
	s.InferenceJobId = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) Validate() error {
	return dara.Validate(s)
}
