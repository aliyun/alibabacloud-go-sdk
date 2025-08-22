// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRegistryModuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *ExecuteRegistryModuleRequest
  GetClientToken() *string 
  SetParameters(v map[string]interface{}) *ExecuteRegistryModuleRequest
  GetParameters() map[string]interface{} 
}

type ExecuteRegistryModuleRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // a65451293e64979ba7a4b573950217fe
  ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
  // example:
  // 
  // {\\"region\\": \\"cn-hangzhou\\", \\"vpcId\\": \\"vpc-bp145sc90s26q0qbkfb6i\\", \\"functionName\\": \\"filemgr-cn-hangzhou-063fd4aead\\", \\"mountPointsServerAddr\\": \\"063fd4aead-dex50.cn-hangzhou.nas.aliyuncs.com\\"}
  Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s ExecuteRegistryModuleRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRegistryModuleRequest) GoString() string {
  return s.String()
}

func (s *ExecuteRegistryModuleRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteRegistryModuleRequest) GetParameters() map[string]interface{}  {
  return s.Parameters
}

func (s *ExecuteRegistryModuleRequest) SetClientToken(v string) *ExecuteRegistryModuleRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteRegistryModuleRequest) SetParameters(v map[string]interface{}) *ExecuteRegistryModuleRequest {
  s.Parameters = v
  return s
}

func (s *ExecuteRegistryModuleRequest) Validate() error {
  return dara.Validate(s)
}

