// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolarClawPluginRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnablePolarClawPluginRequest
  GetApplicationId() *string 
  SetPluginId(v string) *EnablePolarClawPluginRequest
  GetPluginId() *string 
  SetRestart(v bool) *EnablePolarClawPluginRequest
  GetRestart() *bool 
}

type EnablePolarClawPluginRequest struct {
  // The ID of the application.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pa-**************
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // The ID of the plugin to enable.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // openclaw-lark
  PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
  // Indicates whether to restart the gateway after the plugin is enabled. Default value: `true`.
  // 
  // example:
  // 
  // true
  Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s EnablePolarClawPluginRequest) String() string {
  return dara.Prettify(s)
}

func (s EnablePolarClawPluginRequest) GoString() string {
  return s.String()
}

func (s *EnablePolarClawPluginRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnablePolarClawPluginRequest) GetPluginId() *string  {
  return s.PluginId
}

func (s *EnablePolarClawPluginRequest) GetRestart() *bool  {
  return s.Restart
}

func (s *EnablePolarClawPluginRequest) SetApplicationId(v string) *EnablePolarClawPluginRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnablePolarClawPluginRequest) SetPluginId(v string) *EnablePolarClawPluginRequest {
  s.PluginId = &v
  return s
}

func (s *EnablePolarClawPluginRequest) SetRestart(v bool) *EnablePolarClawPluginRequest {
  s.Restart = &v
  return s
}

func (s *EnablePolarClawPluginRequest) Validate() error {
  return dara.Validate(s)
}

