// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServicePluginStep interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v map[string]interface{}) *ServicePluginStep
	GetArgs() map[string]interface{}
	SetPlugin(v string) *ServicePluginStep
	GetPlugin() *string
}

type ServicePluginStep struct {
	// example:
	//
	// {"key":"value"}
	Args map[string]interface{} `json:"args,omitempty" xml:"args,omitempty"`
	// example:
	//
	// dingding-robot
	Plugin *string `json:"plugin,omitempty" xml:"plugin,omitempty"`
}

func (s ServicePluginStep) String() string {
	return dara.Prettify(s)
}

func (s ServicePluginStep) GoString() string {
	return s.String()
}

func (s *ServicePluginStep) GetArgs() map[string]interface{} {
	return s.Args
}

func (s *ServicePluginStep) GetPlugin() *string {
	return s.Plugin
}

func (s *ServicePluginStep) SetArgs(v map[string]interface{}) *ServicePluginStep {
	s.Args = v
	return s
}

func (s *ServicePluginStep) SetPlugin(v string) *ServicePluginStep {
	s.Plugin = &v
	return s
}

func (s *ServicePluginStep) Validate() error {
	return dara.Validate(s)
}
