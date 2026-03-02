// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterprise interface {
  dara.Model
  String() string
  GoString() string
  SetAccountId(v string) *Enterprise
  GetAccountId() *string 
  SetDescription(v string) *Enterprise
  GetDescription() *string 
  SetId(v int64) *Enterprise
  GetId() *int64 
  SetIdentification(v string) *Enterprise
  GetIdentification() *string 
  SetName(v string) *Enterprise
  GetName() *string 
  SetRequestId(v string) *Enterprise
  GetRequestId() *string 
  SetType(v string) *Enterprise
  GetType() *string 
}

type Enterprise struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 121321412341
  AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
  // example:
  // 
  // 互联网企业
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // This parameter is required.
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // linkedmall
  Identification *string `json:"identification,omitempty" xml:"identification,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 阿里
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 内部
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Enterprise) String() string {
  return dara.Prettify(s)
}

func (s Enterprise) GoString() string {
  return s.String()
}

func (s *Enterprise) GetAccountId() *string  {
  return s.AccountId
}

func (s *Enterprise) GetDescription() *string  {
  return s.Description
}

func (s *Enterprise) GetId() *int64  {
  return s.Id
}

func (s *Enterprise) GetIdentification() *string  {
  return s.Identification
}

func (s *Enterprise) GetName() *string  {
  return s.Name
}

func (s *Enterprise) GetRequestId() *string  {
  return s.RequestId
}

func (s *Enterprise) GetType() *string  {
  return s.Type
}

func (s *Enterprise) SetAccountId(v string) *Enterprise {
  s.AccountId = &v
  return s
}

func (s *Enterprise) SetDescription(v string) *Enterprise {
  s.Description = &v
  return s
}

func (s *Enterprise) SetId(v int64) *Enterprise {
  s.Id = &v
  return s
}

func (s *Enterprise) SetIdentification(v string) *Enterprise {
  s.Identification = &v
  return s
}

func (s *Enterprise) SetName(v string) *Enterprise {
  s.Name = &v
  return s
}

func (s *Enterprise) SetRequestId(v string) *Enterprise {
  s.RequestId = &v
  return s
}

func (s *Enterprise) SetType(v string) *Enterprise {
  s.Type = &v
  return s
}

func (s *Enterprise) Validate() error {
  return dara.Validate(s)
}

