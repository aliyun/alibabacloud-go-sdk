// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlgorithmVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmSpecShrink(v string) *CreateAlgorithmVersionShrinkRequest
	GetAlgorithmSpecShrink() *string
}

type CreateAlgorithmVersionShrinkRequest struct {
	AlgorithmSpecShrink *string `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
}

func (s CreateAlgorithmVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlgorithmVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmVersionShrinkRequest) GetAlgorithmSpecShrink() *string {
	return s.AlgorithmSpecShrink
}

func (s *CreateAlgorithmVersionShrinkRequest) SetAlgorithmSpecShrink(v string) *CreateAlgorithmVersionShrinkRequest {
	s.AlgorithmSpecShrink = &v
	return s
}

func (s *CreateAlgorithmVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
