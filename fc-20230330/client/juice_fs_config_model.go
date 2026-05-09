// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJuiceFsConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnvs(v map[string]*string) *JuiceFsConfig
	GetEnvs() map[string]*string
	SetMountPoints(v []*JuiceFsMountConfig) *JuiceFsConfig
	GetMountPoints() []*JuiceFsMountConfig
}

type JuiceFsConfig struct {
	Envs        map[string]*string    `json:"envs" xml:"envs"`
	MountPoints []*JuiceFsMountConfig `json:"mountPoints" xml:"mountPoints" type:"Repeated"`
}

func (s JuiceFsConfig) String() string {
	return dara.Prettify(s)
}

func (s JuiceFsConfig) GoString() string {
	return s.String()
}

func (s *JuiceFsConfig) GetEnvs() map[string]*string {
	return s.Envs
}

func (s *JuiceFsConfig) GetMountPoints() []*JuiceFsMountConfig {
	return s.MountPoints
}

func (s *JuiceFsConfig) SetEnvs(v map[string]*string) *JuiceFsConfig {
	s.Envs = v
	return s
}

func (s *JuiceFsConfig) SetMountPoints(v []*JuiceFsMountConfig) *JuiceFsConfig {
	s.MountPoints = v
	return s
}

func (s *JuiceFsConfig) Validate() error {
	if s.MountPoints != nil {
		for _, item := range s.MountPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
