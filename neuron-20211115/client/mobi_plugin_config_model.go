// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobiPluginConfig interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *MobiPluginConfig
	GetContent() *string
	SetMappings(v []*AppConfigMapping) *MobiPluginConfig
	GetMappings() []*AppConfigMapping
}

type MobiPluginConfig struct {
	// example:
	//
	// {}
	Content  *string             `json:"content,omitempty" xml:"content,omitempty"`
	Mappings []*AppConfigMapping `json:"mappings,omitempty" xml:"mappings,omitempty" type:"Repeated"`
}

func (s MobiPluginConfig) String() string {
	return dara.Prettify(s)
}

func (s MobiPluginConfig) GoString() string {
	return s.String()
}

func (s *MobiPluginConfig) GetContent() *string {
	return s.Content
}

func (s *MobiPluginConfig) GetMappings() []*AppConfigMapping {
	return s.Mappings
}

func (s *MobiPluginConfig) SetContent(v string) *MobiPluginConfig {
	s.Content = &v
	return s
}

func (s *MobiPluginConfig) SetMappings(v []*AppConfigMapping) *MobiPluginConfig {
	s.Mappings = v
	return s
}

func (s *MobiPluginConfig) Validate() error {
	if s.Mappings != nil {
		for _, item := range s.Mappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
