// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptPhoneNumRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EncryptPhoneNumRequest
  GetInstanceId() *string 
  SetPhoneNum(v string) *EncryptPhoneNumRequest
  GetPhoneNum() *string 
}

type EncryptPhoneNumRequest struct {
  // AICCS instance ID. You can obtain it from the Artificial Intelligence Cloud Call Service console.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ccc_xp_pre***
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The plaintext phone number.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 130****0000
  PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
}

func (s EncryptPhoneNumRequest) String() string {
  return dara.Prettify(s)
}

func (s EncryptPhoneNumRequest) GoString() string {
  return s.String()
}

func (s *EncryptPhoneNumRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EncryptPhoneNumRequest) GetPhoneNum() *string  {
  return s.PhoneNum
}

func (s *EncryptPhoneNumRequest) SetInstanceId(v string) *EncryptPhoneNumRequest {
  s.InstanceId = &v
  return s
}

func (s *EncryptPhoneNumRequest) SetPhoneNum(v string) *EncryptPhoneNumRequest {
  s.PhoneNum = &v
  return s
}

func (s *EncryptPhoneNumRequest) Validate() error {
  return dara.Validate(s)
}

