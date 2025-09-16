// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSqlInstanceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteSqlInstanceResponseBody
  GetRequestId() *string 
  SetResult(v *ExecuteSqlInstanceResponseBodyResult) *ExecuteSqlInstanceResponseBody
  GetResult() *ExecuteSqlInstanceResponseBodyResult 
}

type ExecuteSqlInstanceResponseBody struct {
  // id of request
  // 
  // example:
  // 
  // FE03180A-0E29-5474-8A86-33F0683294A4
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // NodeVO
  Result *ExecuteSqlInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteSqlInstanceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSqlInstanceResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteSqlInstanceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteSqlInstanceResponseBody) GetResult() *ExecuteSqlInstanceResponseBodyResult  {
  return s.Result
}

func (s *ExecuteSqlInstanceResponseBody) SetRequestId(v string) *ExecuteSqlInstanceResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBody) SetResult(v *ExecuteSqlInstanceResponseBodyResult) *ExecuteSqlInstanceResponseBody {
  s.Result = v
  return s
}

func (s *ExecuteSqlInstanceResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExecuteSqlInstanceResponseBodyResult struct {
  // example:
  // 
  // 1719221186114
  GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1719220182844
  GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 22
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // ha-cn-pl32rf0****
  InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
  // example:
  // 
  // true
  IsDir *int32 `json:"isDir,omitempty" xml:"isDir,omitempty"`
  // example:
  // 
  // test
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // example:
  // 
  // -1
  Parent *int64 `json:"parent,omitempty" xml:"parent,omitempty"`
  // example:
  // 
  // 1
  TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
  // table, instance, template, function
  // 
  // example:
  // 
  // instance
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ExecuteSqlInstanceResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSqlInstanceResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExecuteSqlInstanceResponseBodyResult) GetGmtCreate() *string  {
  return s.GmtCreate
}

func (s *ExecuteSqlInstanceResponseBodyResult) GetGmtModified() *string  {
  return s.GmtModified
}

func (s *ExecuteSqlInstanceResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExecuteSqlInstanceResponseBodyResult) GetInstanceId() *int64  {
  return s.InstanceId
}

func (s *ExecuteSqlInstanceResponseBodyResult) GetIsDir() *int32  {
  return s.IsDir
}

func (s *ExecuteSqlInstanceResponseBodyResult) GetName() *string  {
  return s.Name
}

func (s *ExecuteSqlInstanceResponseBodyResult) GetParent() *int64  {
  return s.Parent
}

func (s *ExecuteSqlInstanceResponseBodyResult) GetTemplateId() *int64  {
  return s.TemplateId
}

func (s *ExecuteSqlInstanceResponseBodyResult) GetType() *string  {
  return s.Type
}

func (s *ExecuteSqlInstanceResponseBodyResult) SetGmtCreate(v string) *ExecuteSqlInstanceResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBodyResult) SetGmtModified(v string) *ExecuteSqlInstanceResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBodyResult) SetId(v int64) *ExecuteSqlInstanceResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBodyResult) SetInstanceId(v int64) *ExecuteSqlInstanceResponseBodyResult {
  s.InstanceId = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBodyResult) SetIsDir(v int32) *ExecuteSqlInstanceResponseBodyResult {
  s.IsDir = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBodyResult) SetName(v string) *ExecuteSqlInstanceResponseBodyResult {
  s.Name = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBodyResult) SetParent(v int64) *ExecuteSqlInstanceResponseBodyResult {
  s.Parent = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBodyResult) SetTemplateId(v int64) *ExecuteSqlInstanceResponseBodyResult {
  s.TemplateId = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBodyResult) SetType(v string) *ExecuteSqlInstanceResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExecuteSqlInstanceResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

