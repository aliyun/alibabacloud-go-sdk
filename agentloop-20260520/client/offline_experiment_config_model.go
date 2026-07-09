// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineExperimentConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *OfflineExperimentConfig
	GetDesc() *string
	SetLabel(v string) *OfflineExperimentConfig
	GetLabel() *string
	SetName(v string) *OfflineExperimentConfig
	GetName() *string
}

type OfflineExperimentConfig struct {
	Desc  *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OfflineExperimentConfig) String() string {
	return dara.Prettify(s)
}

func (s OfflineExperimentConfig) GoString() string {
	return s.String()
}

func (s *OfflineExperimentConfig) GetDesc() *string {
	return s.Desc
}

func (s *OfflineExperimentConfig) GetLabel() *string {
	return s.Label
}

func (s *OfflineExperimentConfig) GetName() *string {
	return s.Name
}

func (s *OfflineExperimentConfig) SetDesc(v string) *OfflineExperimentConfig {
	s.Desc = &v
	return s
}

func (s *OfflineExperimentConfig) SetLabel(v string) *OfflineExperimentConfig {
	s.Label = &v
	return s
}

func (s *OfflineExperimentConfig) SetName(v string) *OfflineExperimentConfig {
	s.Name = &v
	return s
}

func (s *OfflineExperimentConfig) Validate() error {
	return dara.Validate(s)
}
