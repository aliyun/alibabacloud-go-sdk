// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRequestSGRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLang(v string) *ExecuteRequestSGRequest
  GetLang() *string 
  SetService(v string) *ExecuteRequestSGRequest
  GetService() *string 
  SetServiceParameters(v string) *ExecuteRequestSGRequest
  GetServiceParameters() *string 
}

type ExecuteRequestSGRequest struct {
  // The language.
  // 
  // example:
  // 
  // en
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The service product code. For custom events, set this parameter to saf_de.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // account_abuse_intl_pro
  Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
  // The service input parameters in JSON format. For custom events, you must include the eventCode property in the JSON string.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // {"accountId": "1233***"}
  ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ExecuteRequestSGRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestSGRequest) GoString() string {
  return s.String()
}

func (s *ExecuteRequestSGRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExecuteRequestSGRequest) GetService() *string  {
  return s.Service
}

func (s *ExecuteRequestSGRequest) GetServiceParameters() *string  {
  return s.ServiceParameters
}

func (s *ExecuteRequestSGRequest) SetLang(v string) *ExecuteRequestSGRequest {
  s.Lang = &v
  return s
}

func (s *ExecuteRequestSGRequest) SetService(v string) *ExecuteRequestSGRequest {
  s.Service = &v
  return s
}

func (s *ExecuteRequestSGRequest) SetServiceParameters(v string) *ExecuteRequestSGRequest {
  s.ServiceParameters = &v
  return s
}

func (s *ExecuteRequestSGRequest) Validate() error {
  return dara.Validate(s)
}

