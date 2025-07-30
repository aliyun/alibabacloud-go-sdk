// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableUserRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableUserRequest
  GetInstanceId() *string 
  SetUserId(v string) *EnableUserRequest
  GetUserId() *string 
}

type EnableUserRequest struct {
  // The instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The account ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // user_d6sbsuumeta4h66ec3il7yxxxx
  UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s EnableUserRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableUserRequest) GoString() string {
  return s.String()
}

func (s *EnableUserRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableUserRequest) GetUserId() *string  {
  return s.UserId
}

func (s *EnableUserRequest) SetInstanceId(v string) *EnableUserRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableUserRequest) SetUserId(v string) *EnableUserRequest {
  s.UserId = &v
  return s
}

func (s *EnableUserRequest) Validate() error {
  return dara.Validate(s)
}

