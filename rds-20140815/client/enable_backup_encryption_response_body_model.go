// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupEncryptionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceId(v string) *EnableBackupEncryptionResponseBody
  GetDBInstanceId() *string 
  SetRequestId(v string) *EnableBackupEncryptionResponseBody
  GetRequestId() *string 
}

type EnableBackupEncryptionResponseBody struct {
  // example:
  // 
  // rm-wz951f7f******
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  // example:
  // 
  // FCA65FA6-658A-5C43-96F4-D************
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableBackupEncryptionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupEncryptionResponseBody) GoString() string {
  return s.String()
}

func (s *EnableBackupEncryptionResponseBody) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *EnableBackupEncryptionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableBackupEncryptionResponseBody) SetDBInstanceId(v string) *EnableBackupEncryptionResponseBody {
  s.DBInstanceId = &v
  return s
}

func (s *EnableBackupEncryptionResponseBody) SetRequestId(v string) *EnableBackupEncryptionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableBackupEncryptionResponseBody) Validate() error {
  return dara.Validate(s)
}

