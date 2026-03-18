// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryInferenceJobInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInferenceJobId(v string) *Personalizedtxt2imgQueryInferenceJobInfoRequest
	GetInferenceJobId() *string
}

type Personalizedtxt2imgQueryInferenceJobInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 180
	InferenceJobId *string `json:"inferenceJobId,omitempty" xml:"inferenceJobId,omitempty"`
}

func (s Personalizedtxt2imgQueryInferenceJobInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoRequest) GetInferenceJobId() *string {
	return s.InferenceJobId
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoRequest) SetInferenceJobId(v string) *Personalizedtxt2imgQueryInferenceJobInfoRequest {
	s.InferenceJobId = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoRequest) Validate() error {
	return dara.Validate(s)
}
