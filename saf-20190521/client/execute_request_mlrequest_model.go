// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRequestMLRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLang(v string) *ExecuteRequestMLRequest
  GetLang() *string 
  SetService(v string) *ExecuteRequestMLRequest
  GetService() *string 
  SetServiceParameters(v string) *ExecuteRequestMLRequest
  GetServiceParameters() *string 
}

type ExecuteRequestMLRequest struct {
  // The language.
  // 
  // example:
  // 
  // en
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The service product code. Set this parameter to saf_de because this operation supports only custom events.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // saf_de
  Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
  // The service input parameters in JSON format. You must include the eventCode property of the custom event in the JSON string.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // {"accountId": "1233***"}
  ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ExecuteRequestMLRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestMLRequest) GoString() string {
  return s.String()
}

func (s *ExecuteRequestMLRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExecuteRequestMLRequest) GetService() *string  {
  return s.Service
}

func (s *ExecuteRequestMLRequest) GetServiceParameters() *string  {
  return s.ServiceParameters
}

func (s *ExecuteRequestMLRequest) SetLang(v string) *ExecuteRequestMLRequest {
  s.Lang = &v
  return s
}

func (s *ExecuteRequestMLRequest) SetService(v string) *ExecuteRequestMLRequest {
  s.Service = &v
  return s
}

func (s *ExecuteRequestMLRequest) SetServiceParameters(v string) *ExecuteRequestMLRequest {
  s.ServiceParameters = &v
  return s
}

func (s *ExecuteRequestMLRequest) Validate() error {
  return dara.Validate(s)
}

