// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokePluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v map[string]interface{}) *InvokePluginRequest
	GetParams() map[string]interface{}
	SetPluginId(v string) *InvokePluginRequest
	GetPluginId() *string
}

type InvokePluginRequest struct {
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// 3mj87da7zr
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s InvokePluginRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokePluginRequest) GoString() string {
	return s.String()
}

func (s *InvokePluginRequest) GetParams() map[string]interface{} {
	return s.Params
}

func (s *InvokePluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *InvokePluginRequest) SetParams(v map[string]interface{}) *InvokePluginRequest {
	s.Params = v
	return s
}

func (s *InvokePluginRequest) SetPluginId(v string) *InvokePluginRequest {
	s.PluginId = &v
	return s
}

func (s *InvokePluginRequest) Validate() error {
	return dara.Validate(s)
}
