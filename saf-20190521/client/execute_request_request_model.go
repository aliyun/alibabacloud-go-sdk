// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRequestRequest interface {
  dara.Model
  String() string
  GoString() string
  SetService(v string) *ExecuteRequestRequest
  GetService() *string 
  SetServiceParameters(v string) *ExecuteRequestRequest
  GetServiceParameters() *string 
}

type ExecuteRequestRequest struct {
  // The service product code. For custom events, set this parameter to saf_de.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // account_abuse_pro
  Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
  // The service input parameters in JSON format. This must be a complete JSON string. For custom events, you must include the eventCode property.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // {"accountId": "1233***"}
  ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ExecuteRequestRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestRequest) GoString() string {
  return s.String()
}

func (s *ExecuteRequestRequest) GetService() *string  {
  return s.Service
}

func (s *ExecuteRequestRequest) GetServiceParameters() *string  {
  return s.ServiceParameters
}

func (s *ExecuteRequestRequest) SetService(v string) *ExecuteRequestRequest {
  s.Service = &v
  return s
}

func (s *ExecuteRequestRequest) SetServiceParameters(v string) *ExecuteRequestRequest {
  s.ServiceParameters = &v
  return s
}

func (s *ExecuteRequestRequest) Validate() error {
  return dara.Validate(s)
}

