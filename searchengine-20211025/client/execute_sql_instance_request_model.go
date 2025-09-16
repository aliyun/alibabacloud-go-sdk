// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSqlInstanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCombineParam(v map[string]interface{}) *ExecuteSqlInstanceRequest
  GetCombineParam() map[string]interface{} 
  SetContent(v string) *ExecuteSqlInstanceRequest
  GetContent() *string 
  SetDomain(v string) *ExecuteSqlInstanceRequest
  GetDomain() *string 
  SetDynamicParam(v map[string]interface{}) *ExecuteSqlInstanceRequest
  GetDynamicParam() map[string]interface{} 
  SetKvpair(v map[string]interface{}) *ExecuteSqlInstanceRequest
  GetKvpair() map[string]interface{} 
  SetParams(v map[string]interface{}) *ExecuteSqlInstanceRequest
  GetParams() map[string]interface{} 
  SetStaticParam(v map[string]interface{}) *ExecuteSqlInstanceRequest
  GetStaticParam() map[string]interface{} 
}

type ExecuteSqlInstanceRequest struct {
  CombineParam map[string]interface{} `json:"combineParam,omitempty" xml:"combineParam,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // select 	- from test
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // example:
  // 
  // vpc_hz_domain_1
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  DynamicParam map[string]interface{} `json:"dynamicParam,omitempty" xml:"dynamicParam,omitempty"`
  Kvpair map[string]interface{} `json:"kvpair,omitempty" xml:"kvpair,omitempty"`
  Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
  StaticParam map[string]interface{} `json:"staticParam,omitempty" xml:"staticParam,omitempty"`
}

func (s ExecuteSqlInstanceRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSqlInstanceRequest) GoString() string {
  return s.String()
}

func (s *ExecuteSqlInstanceRequest) GetCombineParam() map[string]interface{}  {
  return s.CombineParam
}

func (s *ExecuteSqlInstanceRequest) GetContent() *string  {
  return s.Content
}

func (s *ExecuteSqlInstanceRequest) GetDomain() *string  {
  return s.Domain
}

func (s *ExecuteSqlInstanceRequest) GetDynamicParam() map[string]interface{}  {
  return s.DynamicParam
}

func (s *ExecuteSqlInstanceRequest) GetKvpair() map[string]interface{}  {
  return s.Kvpair
}

func (s *ExecuteSqlInstanceRequest) GetParams() map[string]interface{}  {
  return s.Params
}

func (s *ExecuteSqlInstanceRequest) GetStaticParam() map[string]interface{}  {
  return s.StaticParam
}

func (s *ExecuteSqlInstanceRequest) SetCombineParam(v map[string]interface{}) *ExecuteSqlInstanceRequest {
  s.CombineParam = v
  return s
}

func (s *ExecuteSqlInstanceRequest) SetContent(v string) *ExecuteSqlInstanceRequest {
  s.Content = &v
  return s
}

func (s *ExecuteSqlInstanceRequest) SetDomain(v string) *ExecuteSqlInstanceRequest {
  s.Domain = &v
  return s
}

func (s *ExecuteSqlInstanceRequest) SetDynamicParam(v map[string]interface{}) *ExecuteSqlInstanceRequest {
  s.DynamicParam = v
  return s
}

func (s *ExecuteSqlInstanceRequest) SetKvpair(v map[string]interface{}) *ExecuteSqlInstanceRequest {
  s.Kvpair = v
  return s
}

func (s *ExecuteSqlInstanceRequest) SetParams(v map[string]interface{}) *ExecuteSqlInstanceRequest {
  s.Params = v
  return s
}

func (s *ExecuteSqlInstanceRequest) SetStaticParam(v map[string]interface{}) *ExecuteSqlInstanceRequest {
  s.StaticParam = v
  return s
}

func (s *ExecuteSqlInstanceRequest) Validate() error {
  return dara.Validate(s)
}

