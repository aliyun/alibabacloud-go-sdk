// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateLocalExtendDiskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAvailable(v string) *EvaluateLocalExtendDiskResponseBody
  GetAvailable() *string 
  SetDBInstanceId(v string) *EvaluateLocalExtendDiskResponseBody
  GetDBInstanceId() *string 
  SetDBInstanceTransType(v string) *EvaluateLocalExtendDiskResponseBody
  GetDBInstanceTransType() *string 
  SetLocalUpgradeDiskLimit(v int64) *EvaluateLocalExtendDiskResponseBody
  GetLocalUpgradeDiskLimit() *int64 
  SetRequestId(v string) *EvaluateLocalExtendDiskResponseBody
  GetRequestId() *string 
}

type EvaluateLocalExtendDiskResponseBody struct {
  // Indicates whether the instance is available. Valid values: true and false.
  // 
  // example:
  // 
  // True
  Available *string `json:"Available,omitempty" xml:"Available,omitempty"`
  // The instance ID.
  // 
  // example:
  // 
  // rm-bp1375i66nd******
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  // The data transfer type supported by the instance.
  // 
  // example:
  // 
  // 0
  DBInstanceTransType *string `json:"DBInstanceTransType,omitempty" xml:"DBInstanceTransType,omitempty"`
  // The maximum value of the local disk. Unit: GB.
  // 
  // example:
  // 
  // 100
  LocalUpgradeDiskLimit *int64 `json:"LocalUpgradeDiskLimit,omitempty" xml:"LocalUpgradeDiskLimit,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // A4C4D26F-E5CE-5A28-8C54-46A6FB318223
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EvaluateLocalExtendDiskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EvaluateLocalExtendDiskResponseBody) GoString() string {
  return s.String()
}

func (s *EvaluateLocalExtendDiskResponseBody) GetAvailable() *string  {
  return s.Available
}

func (s *EvaluateLocalExtendDiskResponseBody) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *EvaluateLocalExtendDiskResponseBody) GetDBInstanceTransType() *string  {
  return s.DBInstanceTransType
}

func (s *EvaluateLocalExtendDiskResponseBody) GetLocalUpgradeDiskLimit() *int64  {
  return s.LocalUpgradeDiskLimit
}

func (s *EvaluateLocalExtendDiskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EvaluateLocalExtendDiskResponseBody) SetAvailable(v string) *EvaluateLocalExtendDiskResponseBody {
  s.Available = &v
  return s
}

func (s *EvaluateLocalExtendDiskResponseBody) SetDBInstanceId(v string) *EvaluateLocalExtendDiskResponseBody {
  s.DBInstanceId = &v
  return s
}

func (s *EvaluateLocalExtendDiskResponseBody) SetDBInstanceTransType(v string) *EvaluateLocalExtendDiskResponseBody {
  s.DBInstanceTransType = &v
  return s
}

func (s *EvaluateLocalExtendDiskResponseBody) SetLocalUpgradeDiskLimit(v int64) *EvaluateLocalExtendDiskResponseBody {
  s.LocalUpgradeDiskLimit = &v
  return s
}

func (s *EvaluateLocalExtendDiskResponseBody) SetRequestId(v string) *EvaluateLocalExtendDiskResponseBody {
  s.RequestId = &v
  return s
}

func (s *EvaluateLocalExtendDiskResponseBody) Validate() error {
  return dara.Validate(s)
}

