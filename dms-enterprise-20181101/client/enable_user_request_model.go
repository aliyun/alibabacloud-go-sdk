// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableUserRequest interface {
  dara.Model
  String() string
  GoString() string
  SetTid(v int64) *EnableUserRequest
  GetTid() *int64 
  SetUid(v string) *EnableUserRequest
  GetUid() *string 
}

type EnableUserRequest struct {
  // The ID of the tenant.
  // 
  // >  To obtain the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Tenant information](https://help.aliyun.com/document_detail/181330.html).
  // 
  // example:
  // 
  // -1
  Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
  // The UID of the Alibaba Cloud account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 12345
  Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EnableUserRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableUserRequest) GoString() string {
  return s.String()
}

func (s *EnableUserRequest) GetTid() *int64  {
  return s.Tid
}

func (s *EnableUserRequest) GetUid() *string  {
  return s.Uid
}

func (s *EnableUserRequest) SetTid(v int64) *EnableUserRequest {
  s.Tid = &v
  return s
}

func (s *EnableUserRequest) SetUid(v string) *EnableUserRequest {
  s.Uid = &v
  return s
}

func (s *EnableUserRequest) Validate() error {
  return dara.Validate(s)
}

