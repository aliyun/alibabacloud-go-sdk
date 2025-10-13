// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlgorithmVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmSpec(v *AlgorithmSpec) *UpdateAlgorithmVersionRequest
	GetAlgorithmSpec() *AlgorithmSpec
}

type UpdateAlgorithmVersionRequest struct {
	AlgorithmSpec *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
}

func (s UpdateAlgorithmVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlgorithmVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmVersionRequest) GetAlgorithmSpec() *AlgorithmSpec {
	return s.AlgorithmSpec
}

func (s *UpdateAlgorithmVersionRequest) SetAlgorithmSpec(v *AlgorithmSpec) *UpdateAlgorithmVersionRequest {
	s.AlgorithmSpec = v
	return s
}

func (s *UpdateAlgorithmVersionRequest) Validate() error {
	if s.AlgorithmSpec != nil {
		if err := s.AlgorithmSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}
