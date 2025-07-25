// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlgorithmVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmSpecShrink(v string) *UpdateAlgorithmVersionShrinkRequest
	GetAlgorithmSpecShrink() *string
}

type UpdateAlgorithmVersionShrinkRequest struct {
	AlgorithmSpecShrink *string `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
}

func (s UpdateAlgorithmVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlgorithmVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmVersionShrinkRequest) GetAlgorithmSpecShrink() *string {
	return s.AlgorithmSpecShrink
}

func (s *UpdateAlgorithmVersionShrinkRequest) SetAlgorithmSpecShrink(v string) *UpdateAlgorithmVersionShrinkRequest {
	s.AlgorithmSpecShrink = &v
	return s
}

func (s *UpdateAlgorithmVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
