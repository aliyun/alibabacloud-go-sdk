// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentSpec interface {
	dara.Model
	String() string
	GoString() string
	SetCodeDir(v *Location) *ComponentSpec
	GetCodeDir() *Location
	SetCommand(v string) *ComponentSpec
	GetCommand() *string
	SetHyperParameters(v []*HyperParameterDefinition) *ComponentSpec
	GetHyperParameters() []*HyperParameterDefinition
	SetImage(v string) *ComponentSpec
	GetImage() *string
	SetInputChannels(v []*Channel) *ComponentSpec
	GetInputChannels() []*Channel
	SetJobType(v string) *ComponentSpec
	GetJobType() *string
	SetMetricDefinitions(v []*MetricDefinition) *ComponentSpec
	GetMetricDefinitions() []*MetricDefinition
	SetOutputChannels(v []*Channel) *ComponentSpec
	GetOutputChannels() []*Channel
	SetResourceRequirements(v []*ConditionExpression) *ComponentSpec
	GetResourceRequirements() []*ConditionExpression
}

type ComponentSpec struct {
	CodeDir *Location `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
	// This parameter is required.
	Command         *string                     `json:"Command,omitempty" xml:"Command,omitempty"`
	HyperParameters []*HyperParameterDefinition `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	// This parameter is required.
	Image         *string    `json:"Image,omitempty" xml:"Image,omitempty"`
	InputChannels []*Channel `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	// This parameter is required.
	JobType              *string                `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MetricDefinitions    []*MetricDefinition    `json:"MetricDefinitions,omitempty" xml:"MetricDefinitions,omitempty" type:"Repeated"`
	OutputChannels       []*Channel             `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	ResourceRequirements []*ConditionExpression `json:"ResourceRequirements,omitempty" xml:"ResourceRequirements,omitempty" type:"Repeated"`
}

func (s ComponentSpec) String() string {
	return dara.Prettify(s)
}

func (s ComponentSpec) GoString() string {
	return s.String()
}

func (s *ComponentSpec) GetCodeDir() *Location {
	return s.CodeDir
}

func (s *ComponentSpec) GetCommand() *string {
	return s.Command
}

func (s *ComponentSpec) GetHyperParameters() []*HyperParameterDefinition {
	return s.HyperParameters
}

func (s *ComponentSpec) GetImage() *string {
	return s.Image
}

func (s *ComponentSpec) GetInputChannels() []*Channel {
	return s.InputChannels
}

func (s *ComponentSpec) GetJobType() *string {
	return s.JobType
}

func (s *ComponentSpec) GetMetricDefinitions() []*MetricDefinition {
	return s.MetricDefinitions
}

func (s *ComponentSpec) GetOutputChannels() []*Channel {
	return s.OutputChannels
}

func (s *ComponentSpec) GetResourceRequirements() []*ConditionExpression {
	return s.ResourceRequirements
}

func (s *ComponentSpec) SetCodeDir(v *Location) *ComponentSpec {
	s.CodeDir = v
	return s
}

func (s *ComponentSpec) SetCommand(v string) *ComponentSpec {
	s.Command = &v
	return s
}

func (s *ComponentSpec) SetHyperParameters(v []*HyperParameterDefinition) *ComponentSpec {
	s.HyperParameters = v
	return s
}

func (s *ComponentSpec) SetImage(v string) *ComponentSpec {
	s.Image = &v
	return s
}

func (s *ComponentSpec) SetInputChannels(v []*Channel) *ComponentSpec {
	s.InputChannels = v
	return s
}

func (s *ComponentSpec) SetJobType(v string) *ComponentSpec {
	s.JobType = &v
	return s
}

func (s *ComponentSpec) SetMetricDefinitions(v []*MetricDefinition) *ComponentSpec {
	s.MetricDefinitions = v
	return s
}

func (s *ComponentSpec) SetOutputChannels(v []*Channel) *ComponentSpec {
	s.OutputChannels = v
	return s
}

func (s *ComponentSpec) SetResourceRequirements(v []*ConditionExpression) *ComponentSpec {
	s.ResourceRequirements = v
	return s
}

func (s *ComponentSpec) Validate() error {
	if s.CodeDir != nil {
		if err := s.CodeDir.Validate(); err != nil {
			return err
		}
	}
	if s.HyperParameters != nil {
		for _, item := range s.HyperParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InputChannels != nil {
		for _, item := range s.InputChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MetricDefinitions != nil {
		for _, item := range s.MetricDefinitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputChannels != nil {
		for _, item := range s.OutputChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceRequirements != nil {
		for _, item := range s.ResourceRequirements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
