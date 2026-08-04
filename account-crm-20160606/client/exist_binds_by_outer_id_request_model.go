// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExistBindsByOuterIdRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *ExistBindsByOuterIdRequest
  GetAppName() *string 
  SetMinorOuterId(v string) *ExistBindsByOuterIdRequest
  GetMinorOuterId() *string 
  SetOuterId(v string) *ExistBindsByOuterIdRequest
  GetOuterId() *string 
  SetTenantId(v string) *ExistBindsByOuterIdRequest
  GetTenantId() *string 
}

type ExistBindsByOuterIdRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  MinorOuterId *string `json:"MinorOuterId,omitempty" xml:"MinorOuterId,omitempty"`
  OuterId *string `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
  TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ExistBindsByOuterIdRequest) String() string {
  return dara.Prettify(s)
}

func (s ExistBindsByOuterIdRequest) GoString() string {
  return s.String()
}

func (s *ExistBindsByOuterIdRequest) GetAppName() *string  {
  return s.AppName
}

func (s *ExistBindsByOuterIdRequest) GetMinorOuterId() *string  {
  return s.MinorOuterId
}

func (s *ExistBindsByOuterIdRequest) GetOuterId() *string  {
  return s.OuterId
}

func (s *ExistBindsByOuterIdRequest) GetTenantId() *string  {
  return s.TenantId
}

func (s *ExistBindsByOuterIdRequest) SetAppName(v string) *ExistBindsByOuterIdRequest {
  s.AppName = &v
  return s
}

func (s *ExistBindsByOuterIdRequest) SetMinorOuterId(v string) *ExistBindsByOuterIdRequest {
  s.MinorOuterId = &v
  return s
}

func (s *ExistBindsByOuterIdRequest) SetOuterId(v string) *ExistBindsByOuterIdRequest {
  s.OuterId = &v
  return s
}

func (s *ExistBindsByOuterIdRequest) SetTenantId(v string) *ExistBindsByOuterIdRequest {
  s.TenantId = &v
  return s
}

func (s *ExistBindsByOuterIdRequest) Validate() error {
  return dara.Validate(s)
}

