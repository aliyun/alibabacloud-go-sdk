// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsightsLabel interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *InsightsLabel
	GetDescription() *string
	SetName(v string) *InsightsLabel
	GetName() *string
}

type InsightsLabel struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s InsightsLabel) String() string {
	return dara.Prettify(s)
}

func (s InsightsLabel) GoString() string {
	return s.String()
}

func (s *InsightsLabel) GetDescription() *string {
	return s.Description
}

func (s *InsightsLabel) GetName() *string {
	return s.Name
}

func (s *InsightsLabel) SetDescription(v string) *InsightsLabel {
	s.Description = &v
	return s
}

func (s *InsightsLabel) SetName(v string) *InsightsLabel {
	s.Name = &v
	return s
}

func (s *InsightsLabel) Validate() error {
	return dara.Validate(s)
}
