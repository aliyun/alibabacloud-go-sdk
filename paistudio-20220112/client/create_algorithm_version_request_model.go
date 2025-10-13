// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlgorithmVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmSpec(v *AlgorithmSpec) *CreateAlgorithmVersionRequest
	GetAlgorithmSpec() *AlgorithmSpec
}

type CreateAlgorithmVersionRequest struct {
	AlgorithmSpec *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
}

func (s CreateAlgorithmVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlgorithmVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmVersionRequest) GetAlgorithmSpec() *AlgorithmSpec {
	return s.AlgorithmSpec
}

func (s *CreateAlgorithmVersionRequest) SetAlgorithmSpec(v *AlgorithmSpec) *CreateAlgorithmVersionRequest {
	s.AlgorithmSpec = v
	return s
}

func (s *CreateAlgorithmVersionRequest) Validate() error {
	if s.AlgorithmSpec != nil {
		if err := s.AlgorithmSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}
