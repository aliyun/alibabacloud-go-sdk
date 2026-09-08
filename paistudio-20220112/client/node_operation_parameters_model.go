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
	SetResizeDiskParameters(v *ResizeDiskParameters) *NodeOperationParameters
	GetResizeDiskParameters() *ResizeDiskParameters
	SetUncordonParameters(v *NodeUncordonParameters) *NodeOperationParameters
	GetUncordonParameters() *NodeUncordonParameters
}

type NodeOperationParameters struct {
	// The parameter settings for disabling node scheduling.
	CordonParameters *NodeCordonParameters `json:"CordonParameters,omitempty" xml:"CordonParameters,omitempty"`
	// The parameter settings for draining task instances from a node.
	DrainParameters *NodeDrainParameters `json:"DrainParameters,omitempty" xml:"DrainParameters,omitempty"`
	// The parameters for changing disk capacity.
	ResizeDiskParameters *ResizeDiskParameters `json:"ResizeDiskParameters,omitempty" xml:"ResizeDiskParameters,omitempty"`
	// The parameter settings for enabling node scheduling.
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

func (s *NodeOperationParameters) GetResizeDiskParameters() *ResizeDiskParameters {
	return s.ResizeDiskParameters
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

func (s *NodeOperationParameters) SetResizeDiskParameters(v *ResizeDiskParameters) *NodeOperationParameters {
	s.ResizeDiskParameters = v
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
	if s.ResizeDiskParameters != nil {
		if err := s.ResizeDiskParameters.Validate(); err != nil {
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
