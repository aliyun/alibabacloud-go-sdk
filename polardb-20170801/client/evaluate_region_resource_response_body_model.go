// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateRegionResourceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceAvailable(v string) *EvaluateRegionResourceResponseBody
  GetDBInstanceAvailable() *string 
  SetDBType(v string) *EvaluateRegionResourceResponseBody
  GetDBType() *string 
  SetDBVersion(v string) *EvaluateRegionResourceResponseBody
  GetDBVersion() *string 
  SetRequestId(v string) *EvaluateRegionResourceResponseBody
  GetRequestId() *string 
}

type EvaluateRegionResourceResponseBody struct {
  // Indicates whether the resources are sufficient.
  // 
  // - **true**: The resources are sufficient.
  // 
  // - **false**: The resources are insufficient.
  // 
  // example:
  // 
  // true
  DBInstanceAvailable *string `json:"DBInstanceAvailable,omitempty" xml:"DBInstanceAvailable,omitempty"`
  // The database engine type. Valid values:
  // 
  // - **MySQL**
  // 
  // - **PostgreSQL**
  // 
  // - **Oracle**
  // 
  // example:
  // 
  // MySQL
  DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
  // The version of the database engine.
  // 
  // - Valid values for MySQL:
  // 
  //   - **5.6**
  // 
  //   - **5.7**
  // 
  //   - **8.0**
  // 
  // - Valid values for PostgreSQL:
  // 
  //   - **11**
  // 
  //   - **14**
  // 
  // - Valid value for Oracle: **11**.
  // 
  // example:
  // 
  // 8.0
  DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 69A85BAF-1089-4CDF-A82F-0A140F******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EvaluateRegionResourceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EvaluateRegionResourceResponseBody) GoString() string {
  return s.String()
}

func (s *EvaluateRegionResourceResponseBody) GetDBInstanceAvailable() *string  {
  return s.DBInstanceAvailable
}

func (s *EvaluateRegionResourceResponseBody) GetDBType() *string  {
  return s.DBType
}

func (s *EvaluateRegionResourceResponseBody) GetDBVersion() *string  {
  return s.DBVersion
}

func (s *EvaluateRegionResourceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EvaluateRegionResourceResponseBody) SetDBInstanceAvailable(v string) *EvaluateRegionResourceResponseBody {
  s.DBInstanceAvailable = &v
  return s
}

func (s *EvaluateRegionResourceResponseBody) SetDBType(v string) *EvaluateRegionResourceResponseBody {
  s.DBType = &v
  return s
}

func (s *EvaluateRegionResourceResponseBody) SetDBVersion(v string) *EvaluateRegionResourceResponseBody {
  s.DBVersion = &v
  return s
}

func (s *EvaluateRegionResourceResponseBody) SetRequestId(v string) *EvaluateRegionResourceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EvaluateRegionResourceResponseBody) Validate() error {
  return dara.Validate(s)
}

