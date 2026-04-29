// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolarClawChannelResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnablePolarClawChannelResponseBody
  GetApplicationId() *string 
  SetChannelId(v string) *EnablePolarClawChannelResponseBody
  GetChannelId() *string 
  SetCode(v int32) *EnablePolarClawChannelResponseBody
  GetCode() *int32 
  SetMessage(v string) *EnablePolarClawChannelResponseBody
  GetMessage() *string 
  SetOk(v bool) *EnablePolarClawChannelResponseBody
  GetOk() *bool 
  SetRequestId(v string) *EnablePolarClawChannelResponseBody
  GetRequestId() *string 
  SetRestarted(v bool) *EnablePolarClawChannelResponseBody
  GetRestarted() *bool 
}

type EnablePolarClawChannelResponseBody struct {
  // example:
  // 
  // pa-**************
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // example:
  // 
  // feishu
  ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
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
  // Id of the request
  // 
  // example:
  // 
  // 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
}

func (s EnablePolarClawChannelResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnablePolarClawChannelResponseBody) GoString() string {
  return s.String()
}

func (s *EnablePolarClawChannelResponseBody) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnablePolarClawChannelResponseBody) GetChannelId() *string  {
  return s.ChannelId
}

func (s *EnablePolarClawChannelResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnablePolarClawChannelResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnablePolarClawChannelResponseBody) GetOk() *bool  {
  return s.Ok
}

func (s *EnablePolarClawChannelResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnablePolarClawChannelResponseBody) GetRestarted() *bool  {
  return s.Restarted
}

func (s *EnablePolarClawChannelResponseBody) SetApplicationId(v string) *EnablePolarClawChannelResponseBody {
  s.ApplicationId = &v
  return s
}

func (s *EnablePolarClawChannelResponseBody) SetChannelId(v string) *EnablePolarClawChannelResponseBody {
  s.ChannelId = &v
  return s
}

func (s *EnablePolarClawChannelResponseBody) SetCode(v int32) *EnablePolarClawChannelResponseBody {
  s.Code = &v
  return s
}

func (s *EnablePolarClawChannelResponseBody) SetMessage(v string) *EnablePolarClawChannelResponseBody {
  s.Message = &v
  return s
}

func (s *EnablePolarClawChannelResponseBody) SetOk(v bool) *EnablePolarClawChannelResponseBody {
  s.Ok = &v
  return s
}

func (s *EnablePolarClawChannelResponseBody) SetRequestId(v string) *EnablePolarClawChannelResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnablePolarClawChannelResponseBody) SetRestarted(v bool) *EnablePolarClawChannelResponseBody {
  s.Restarted = &v
  return s
}

func (s *EnablePolarClawChannelResponseBody) Validate() error {
  return dara.Validate(s)
}

