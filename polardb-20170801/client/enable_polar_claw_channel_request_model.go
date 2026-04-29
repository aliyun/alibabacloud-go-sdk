// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolarClawChannelRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnablePolarClawChannelRequest
  GetApplicationId() *string 
  SetChannelId(v string) *EnablePolarClawChannelRequest
  GetChannelId() *string 
  SetRestart(v bool) *EnablePolarClawChannelRequest
  GetRestart() *bool 
}

type EnablePolarClawChannelRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // pa-**************
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // feishu
  ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
  // example:
  // 
  // true
  Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s EnablePolarClawChannelRequest) String() string {
  return dara.Prettify(s)
}

func (s EnablePolarClawChannelRequest) GoString() string {
  return s.String()
}

func (s *EnablePolarClawChannelRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnablePolarClawChannelRequest) GetChannelId() *string  {
  return s.ChannelId
}

func (s *EnablePolarClawChannelRequest) GetRestart() *bool  {
  return s.Restart
}

func (s *EnablePolarClawChannelRequest) SetApplicationId(v string) *EnablePolarClawChannelRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnablePolarClawChannelRequest) SetChannelId(v string) *EnablePolarClawChannelRequest {
  s.ChannelId = &v
  return s
}

func (s *EnablePolarClawChannelRequest) SetRestart(v bool) *EnablePolarClawChannelRequest {
  s.Restart = &v
  return s
}

func (s *EnablePolarClawChannelRequest) Validate() error {
  return dara.Validate(s)
}

