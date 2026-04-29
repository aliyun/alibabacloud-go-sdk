// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolarClawPluginResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnablePolarClawPluginResponseBody
  GetApplicationId() *string 
  SetCode(v int32) *EnablePolarClawPluginResponseBody
  GetCode() *int32 
  SetMessage(v string) *EnablePolarClawPluginResponseBody
  GetMessage() *string 
  SetOk(v bool) *EnablePolarClawPluginResponseBody
  GetOk() *bool 
  SetPluginId(v string) *EnablePolarClawPluginResponseBody
  GetPluginId() *string 
  SetRequestId(v string) *EnablePolarClawPluginResponseBody
  GetRequestId() *string 
  SetRestarted(v bool) *EnablePolarClawPluginResponseBody
  GetRestarted() *bool 
}

type EnablePolarClawPluginResponseBody struct {
  // example:
  // 
  // pa-**************
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // true
  Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
  // example:
  // 
  // openclaw-lark
  PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 7F2007D3-7E74-4ECB-89A8-BF130D******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s EnablePolarClawPluginResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnablePolarClawPluginResponseBody) GoString() string {
  return s.String()
}

func (s *EnablePolarClawPluginResponseBody) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnablePolarClawPluginResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnablePolarClawPluginResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnablePolarClawPluginResponseBody) GetOk() *bool  {
  return s.Ok
}

func (s *EnablePolarClawPluginResponseBody) GetPluginId() *string  {
  return s.PluginId
}

func (s *EnablePolarClawPluginResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnablePolarClawPluginResponseBody) GetRestarted() *bool  {
  return s.Restarted
}

func (s *EnablePolarClawPluginResponseBody) SetApplicationId(v string) *EnablePolarClawPluginResponseBody {
  s.ApplicationId = &v
  return s
}

func (s *EnablePolarClawPluginResponseBody) SetCode(v int32) *EnablePolarClawPluginResponseBody {
  s.Code = &v
  return s
}

func (s *EnablePolarClawPluginResponseBody) SetMessage(v string) *EnablePolarClawPluginResponseBody {
  s.Message = &v
  return s
}

func (s *EnablePolarClawPluginResponseBody) SetOk(v bool) *EnablePolarClawPluginResponseBody {
  s.Ok = &v
  return s
}

func (s *EnablePolarClawPluginResponseBody) SetPluginId(v string) *EnablePolarClawPluginResponseBody {
  s.PluginId = &v
  return s
}

func (s *EnablePolarClawPluginResponseBody) SetRequestId(v string) *EnablePolarClawPluginResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnablePolarClawPluginResponseBody) SetRestarted(v bool) *EnablePolarClawPluginResponseBody {
  s.Restarted = &v
  return s
}

func (s *EnablePolarClawPluginResponseBody) Validate() error {
  return dara.Validate(s)
}

