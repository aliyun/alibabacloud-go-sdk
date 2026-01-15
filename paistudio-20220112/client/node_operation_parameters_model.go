// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeOperationParameters interface {
	dara.Model
	String() string
	GoString() string
	SetCordonParameters(v *NodeCordonParameters) *NodeOperationParameters
	GetCordonParameters() *NodeCordonParameters
	SetDrainParameters(v *NodeDrainParameters) *NodeOperationParameters
	GetDrainParameters() *NodeDrainParameters
	SetUncordonParameters(v *NodeUncordonParameters) *NodeOperationParameters
	GetUncordonParameters() *NodeUncordonParameters
}

type NodeOperationParameters struct {
	CordonParameters   *NodeCordonParameters   `json:"CordonParameters,omitempty" xml:"CordonParameters,omitempty"`
	DrainParameters    *NodeDrainParameters    `json:"DrainParameters,omitempty" xml:"DrainParameters,omitempty"`
	UncordonParameters *NodeUncordonParameters `json:"UncordonParameters,omitempty" xml:"UncordonParameters,omitempty"`
}

func (s NodeOperationParameters) String() string {
	return dara.Prettify(s)
}

func (s NodeOperationParameters) GoString() string {
	return s.String()
}

func (s *NodeOperationParameters) GetCordonParameters() *NodeCordonParameters {
	return s.CordonParameters
}

func (s *NodeOperationParameters) GetDrainParameters() *NodeDrainParameters {
	return s.DrainParameters
}

func (s *NodeOperationParameters) GetUncordonParameters() *NodeUncordonParameters {
	return s.UncordonParameters
}

func (s *NodeOperationParameters) SetCordonParameters(v *NodeCordonParameters) *NodeOperationParameters {
	s.CordonParameters = v
	return s
}

func (s *NodeOperationParameters) SetDrainParameters(v *NodeDrainParameters) *NodeOperationParameters {
	s.DrainParameters = v
	return s
}

func (s *NodeOperationParameters) SetUncordonParameters(v *NodeUncordonParameters) *NodeOperationParameters {
	s.UncordonParameters = v
	return s
}

func (s *NodeOperationParameters) Validate() error {
	if s.CordonParameters != nil {
		if err := s.CordonParameters.Validate(); err != nil {
			return err
		}
	}
	if s.DrainParameters != nil {
		if err := s.DrainParameters.Validate(); err != nil {
			return err
		}
	}
	if s.UncordonParameters != nil {
		if err := s.UncordonParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}
