// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigEffectActions interface {
	dara.Model
	String() string
	GoString() string
	SetConfigEffectAction(v string) *ConfigEffectActions
	GetConfigEffectAction() *string
	SetConfigFiles(v []*string) *ConfigEffectActions
	GetConfigFiles() []*string
}

type ConfigEffectActions struct {
	// 配置生效动作。
	//
	// example:
	//
	// restart
	ConfigEffectAction *string `json:"ConfigEffectAction,omitempty" xml:"ConfigEffectAction,omitempty"`
	// 配置生效配置文件。
	//
	// example:
	//
	// null
	ConfigFiles []*string `json:"ConfigFiles,omitempty" xml:"ConfigFiles,omitempty" type:"Repeated"`
}

func (s ConfigEffectActions) String() string {
	return dara.Prettify(s)
}

func (s ConfigEffectActions) GoString() string {
	return s.String()
}

func (s *ConfigEffectActions) GetConfigEffectAction() *string {
	return s.ConfigEffectAction
}

func (s *ConfigEffectActions) GetConfigFiles() []*string {
	return s.ConfigFiles
}

func (s *ConfigEffectActions) SetConfigEffectAction(v string) *ConfigEffectActions {
	s.ConfigEffectAction = &v
	return s
}

func (s *ConfigEffectActions) SetConfigFiles(v []*string) *ConfigEffectActions {
	s.ConfigFiles = v
	return s
}

func (s *ConfigEffectActions) Validate() error {
	return dara.Validate(s)
}
