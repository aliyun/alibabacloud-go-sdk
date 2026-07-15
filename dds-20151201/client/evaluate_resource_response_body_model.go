// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateResourceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceAvailable(v string) *EvaluateResourceResponseBody
  GetDBInstanceAvailable() *string 
  SetEngine(v string) *EvaluateResourceResponseBody
  GetEngine() *string 
  SetEngineVersion(v string) *EvaluateResourceResponseBody
  GetEngineVersion() *string 
  SetRequestId(v string) *EvaluateResourceResponseBody
  GetRequestId() *string 
}

type EvaluateResourceResponseBody struct {
  // Indicates whether resources are available in the current region. Valid values:
  // 
  // - **1**: Resources are sufficient.
  // 
  // - **0**: Resources are insufficient.
  // 
  // example:
  // 
  // 1
  DBInstanceAvailable *string `json:"DBInstanceAvailable,omitempty" xml:"DBInstanceAvailable,omitempty"`
  // The database engine. The value is fixed to MongoDB.
  // 
  // example:
  // 
  // MongoDB
  Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
  // The database engine version.
  // 
  // example:
  // 
  // 4.0
  EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // AE2DE465-E45F-481F-ABD8-37D64173****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EvaluateResourceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EvaluateResourceResponseBody) GoString() string {
  return s.String()
}

func (s *EvaluateResourceResponseBody) GetDBInstanceAvailable() *string  {
  return s.DBInstanceAvailable
}

func (s *EvaluateResourceResponseBody) GetEngine() *string  {
  return s.Engine
}

func (s *EvaluateResourceResponseBody) GetEngineVersion() *string  {
  return s.EngineVersion
}

func (s *EvaluateResourceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EvaluateResourceResponseBody) SetDBInstanceAvailable(v string) *EvaluateResourceResponseBody {
  s.DBInstanceAvailable = &v
  return s
}

func (s *EvaluateResourceResponseBody) SetEngine(v string) *EvaluateResourceResponseBody {
  s.Engine = &v
  return s
}

func (s *EvaluateResourceResponseBody) SetEngineVersion(v string) *EvaluateResourceResponseBody {
  s.EngineVersion = &v
  return s
}

func (s *EvaluateResourceResponseBody) SetRequestId(v string) *EvaluateResourceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EvaluateResourceResponseBody) Validate() error {
  return dara.Validate(s)
}

