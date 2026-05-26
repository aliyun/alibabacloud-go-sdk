// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartClusterConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFigure(v *FigureClusterConfig) *SmartClusterConfig
	GetFigure() *FigureClusterConfig
}

type SmartClusterConfig struct {
	Figure *FigureClusterConfig `json:"Figure,omitempty" xml:"Figure,omitempty"`
}

func (s SmartClusterConfig) String() string {
	return dara.Prettify(s)
}

func (s SmartClusterConfig) GoString() string {
	return s.String()
}

func (s *SmartClusterConfig) GetFigure() *FigureClusterConfig {
	return s.Figure
}

func (s *SmartClusterConfig) SetFigure(v *FigureClusterConfig) *SmartClusterConfig {
	s.Figure = v
	return s
}

func (s *SmartClusterConfig) Validate() error {
	if s.Figure != nil {
		if err := s.Figure.Validate(); err != nil {
			return err
		}
	}
	return nil
}
