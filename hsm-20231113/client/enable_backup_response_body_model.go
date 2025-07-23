// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableBackupResponseBody
  GetRequestId() *string 
}

type EnableBackupResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 4C467B38-3910-447D-87BC-AC049166F216
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableBackupResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupResponseBody) GoString() string {
  return s.String()
}

func (s *EnableBackupResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableBackupResponseBody) SetRequestId(v string) *EnableBackupResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableBackupResponseBody) Validate() error {
  return dara.Validate(s)
}

