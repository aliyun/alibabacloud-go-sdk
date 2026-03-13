// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricDefinition interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *MetricDefinition
	GetDescription() *string
	SetName(v string) *MetricDefinition
	GetName() *string
	SetRegex(v string) *MetricDefinition
	GetRegex() *string
}

type MetricDefinition struct {
	// example:
	//
	// train dataset oob score
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// train:oob_score
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// .*train:oob_score=([-+]?[0-9]*\\.?[0-9]+(?:[eE][-+]?[0-9]+)?).*
	Regex *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
}

func (s MetricDefinition) String() string {
	return dara.Prettify(s)
}

func (s MetricDefinition) GoString() string {
	return s.String()
}

func (s *MetricDefinition) GetDescription() *string {
	return s.Description
}

func (s *MetricDefinition) GetName() *string {
	return s.Name
}

func (s *MetricDefinition) GetRegex() *string {
	return s.Regex
}

func (s *MetricDefinition) SetDescription(v string) *MetricDefinition {
	s.Description = &v
	return s
}

func (s *MetricDefinition) SetName(v string) *MetricDefinition {
	s.Name = &v
	return s
}

func (s *MetricDefinition) SetRegex(v string) *MetricDefinition {
	s.Regex = &v
	return s
}

func (s *MetricDefinition) Validate() error {
	return dara.Validate(s)
}
