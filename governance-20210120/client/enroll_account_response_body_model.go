// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnrollAccountResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccountUid(v int64) *EnrollAccountResponseBody
  GetAccountUid() *int64 
  SetRequestId(v string) *EnrollAccountResponseBody
  GetRequestId() *string 
}

type EnrollAccountResponseBody struct {
  // The account ID.
  // 
  // example:
  // 
  // 143165363236****
  AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 7071E5FA-515E-5F53-B335-B87D619C6A66
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnrollAccountResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnrollAccountResponseBody) GoString() string {
  return s.String()
}

func (s *EnrollAccountResponseBody) GetAccountUid() *int64  {
  return s.AccountUid
}

func (s *EnrollAccountResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnrollAccountResponseBody) SetAccountUid(v int64) *EnrollAccountResponseBody {
  s.AccountUid = &v
  return s
}

func (s *EnrollAccountResponseBody) SetRequestId(v string) *EnrollAccountResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnrollAccountResponseBody) Validate() error {
  return dara.Validate(s)
}

