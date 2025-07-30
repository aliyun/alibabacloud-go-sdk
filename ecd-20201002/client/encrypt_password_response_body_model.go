// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptPasswordResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetEncryptedPassword(v string) *EncryptPasswordResponseBody
  GetEncryptedPassword() *string 
  SetRequestId(v string) *EncryptPasswordResponseBody
  GetRequestId() *string 
}

type EncryptPasswordResponseBody struct {
  // The encrypted password.
  // 
  // example:
  // 
  // d34601bc-e6b1-4433-b0cc-8f6c5e52;n4apvGub3OBoj4Grwg==;thhO4UEomJfdvwnwlA==
  EncryptedPassword *string `json:"EncryptedPassword,omitempty" xml:"EncryptedPassword,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // AF538DA8-FFC6-52DA-8FF8-7B92579F****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EncryptPasswordResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EncryptPasswordResponseBody) GoString() string {
  return s.String()
}

func (s *EncryptPasswordResponseBody) GetEncryptedPassword() *string  {
  return s.EncryptedPassword
}

func (s *EncryptPasswordResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EncryptPasswordResponseBody) SetEncryptedPassword(v string) *EncryptPasswordResponseBody {
  s.EncryptedPassword = &v
  return s
}

func (s *EncryptPasswordResponseBody) SetRequestId(v string) *EncryptPasswordResponseBody {
  s.RequestId = &v
  return s
}

func (s *EncryptPasswordResponseBody) Validate() error {
  return dara.Validate(s)
}

