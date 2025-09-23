// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptySlsLogstoreRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLang(v string) *EmptySlsLogstoreRequest
  GetLang() *string 
  SetResourceGroupId(v string) *EmptySlsLogstoreRequest
  GetResourceGroupId() *string 
  SetSourceIp(v string) *EmptySlsLogstoreRequest
  GetSourceIp() *string 
}

type EmptySlsLogstoreRequest struct {
  // example:
  // 
  // cn
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // example:
  // 
  // xx
  ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
  // example:
  // 
  // 1.1.1.1
  SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s EmptySlsLogstoreRequest) String() string {
  return dara.Prettify(s)
}

func (s EmptySlsLogstoreRequest) GoString() string {
  return s.String()
}

func (s *EmptySlsLogstoreRequest) GetLang() *string  {
  return s.Lang
}

func (s *EmptySlsLogstoreRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EmptySlsLogstoreRequest) GetSourceIp() *string  {
  return s.SourceIp
}

func (s *EmptySlsLogstoreRequest) SetLang(v string) *EmptySlsLogstoreRequest {
  s.Lang = &v
  return s
}

func (s *EmptySlsLogstoreRequest) SetResourceGroupId(v string) *EmptySlsLogstoreRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *EmptySlsLogstoreRequest) SetSourceIp(v string) *EmptySlsLogstoreRequest {
  s.SourceIp = &v
  return s
}

func (s *EmptySlsLogstoreRequest) Validate() error {
  return dara.Validate(s)
}

