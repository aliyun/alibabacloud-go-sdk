// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppPluginUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *MobiPluginConfig) *NeuronAppPluginUpdateCmd
	GetConfig() *MobiPluginConfig
	SetId(v int64) *NeuronAppPluginUpdateCmd
	GetId() *int64
}

type NeuronAppPluginUpdateCmd struct {
	// This parameter is required.
	Config *MobiPluginConfig `json:"config,omitempty" xml:"config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s NeuronAppPluginUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppPluginUpdateCmd) GoString() string {
	return s.String()
}

func (s *NeuronAppPluginUpdateCmd) GetConfig() *MobiPluginConfig {
	return s.Config
}

func (s *NeuronAppPluginUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *NeuronAppPluginUpdateCmd) SetConfig(v *MobiPluginConfig) *NeuronAppPluginUpdateCmd {
	s.Config = v
	return s
}

func (s *NeuronAppPluginUpdateCmd) SetId(v int64) *NeuronAppPluginUpdateCmd {
	s.Id = &v
	return s
}

func (s *NeuronAppPluginUpdateCmd) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}
