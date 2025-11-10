// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteMigrationOperationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOperationParam(v *ExecuteMigrationOperationRequestOperationParam) *ExecuteMigrationOperationRequest
  GetOperationParam() *ExecuteMigrationOperationRequestOperationParam 
  SetInstanceId(v string) *ExecuteMigrationOperationRequest
  GetInstanceId() *string 
}

type ExecuteMigrationOperationRequest struct {
  OperationParam *ExecuteMigrationOperationRequestOperationParam `json:"operationParam,omitempty" xml:"operationParam,omitempty" type:"Struct"`
  // example:
  // 
  // rmq-cn-x0r37kelk0o
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ExecuteMigrationOperationRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMigrationOperationRequest) GoString() string {
  return s.String()
}

func (s *ExecuteMigrationOperationRequest) GetOperationParam() *ExecuteMigrationOperationRequestOperationParam  {
  return s.OperationParam
}

func (s *ExecuteMigrationOperationRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExecuteMigrationOperationRequest) SetOperationParam(v *ExecuteMigrationOperationRequestOperationParam) *ExecuteMigrationOperationRequest {
  s.OperationParam = v
  return s
}

func (s *ExecuteMigrationOperationRequest) SetInstanceId(v string) *ExecuteMigrationOperationRequest {
  s.InstanceId = &v
  return s
}

func (s *ExecuteMigrationOperationRequest) Validate() error {
  if s.OperationParam != nil {
    if err := s.OperationParam.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteMigrationOperationRequestOperationParam struct {
  // example:
  // 
  // {}
  ParamData interface{} `json:"paramData,omitempty" xml:"paramData,omitempty"`
}

func (s ExecuteMigrationOperationRequestOperationParam) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMigrationOperationRequestOperationParam) GoString() string {
  return s.String()
}

func (s *ExecuteMigrationOperationRequestOperationParam) GetParamData() interface{}  {
  return s.ParamData
}

func (s *ExecuteMigrationOperationRequestOperationParam) SetParamData(v interface{}) *ExecuteMigrationOperationRequestOperationParam {
  s.ParamData = v
  return s
}

func (s *ExecuteMigrationOperationRequestOperationParam) Validate() error {
  return dara.Validate(s)
}

