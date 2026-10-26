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
  // The operation parameter.
  OperationParam *ExecuteMigrationOperationRequestOperationParam `json:"operationParam,omitempty" xml:"operationParam,omitempty" type:"Struct"`
  // The instance ID.
  // 
  // example:
  // 
  // rmq-cn-x0r37kel***
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
  // The parameter information. The parameter type is a JSON string. See the examples below.
  // 
  // The valid values for the action field are as follows:
  // 
  // - MIGRATE_FORWARD (migrate forward)
  // 
  // - MIGRATE_BACKWARD (migrate backward)
  // 
  // The valid values for the targetStage field are as follows:
  // 
  //   - MIGRATE_WAIT (read from and write to the source cluster)
  // 
  //   - DOUBLE_READ (write to the source cluster with redundant reads)
  // 
  //   - SWITCH_WRITE (write to the destination cluster with redundant reads)
  // 
  //   - MIGRATE_FINISH (read from and write to the destination cluster)
  // 
  // 
  // > 
  // 
  // > - For migration forward, you can only advance one stage at a time in order.
  // 
  // > - For migration rollback, you can roll back to any previous stage.
  // 
  // example:
  // 
  // Migrate forward: {"action":"MIGRATE_FORWARD","targetStage":"MIGRATE_FINISH"}
  // 
  // Migrate backward: {"action":"MIGRATE_BACKWARD","targetStage":"MIGRATE_WAIT"}
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

