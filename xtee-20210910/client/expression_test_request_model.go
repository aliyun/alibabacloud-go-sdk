// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpressionTestRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLang(v string) *ExpressionTestRequest
  GetLang() *string 
  SetExpression(v string) *ExpressionTestRequest
  GetExpression() *string 
  SetExpressionVariable(v string) *ExpressionTestRequest
  GetExpressionVariable() *string 
  SetExpressionVariableIds(v string) *ExpressionTestRequest
  GetExpressionVariableIds() *string 
  SetId(v int64) *ExpressionTestRequest
  GetId() *int64 
  SetRegId(v string) *ExpressionTestRequest
  GetRegId() *string 
  SetScene(v string) *ExpressionTestRequest
  GetScene() *string 
}

type ExpressionTestRequest struct {
  // Sets the language type for requests and received messages, with a default value of **zh**. Values: 
  // 
  // - **zh**: Chinese
  // 
  // - **en**: English
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // Test expression.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // @ex_GX9rrlTq4b67 + 1001
  Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
  // Calculation expression variable
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // [{\\"name\\":\\"age\\",\\"code\\":\\"age\\",\\"fieldType\\":\\"INT\\",\\"id\\":44809,\\"value\\":\\"1\\"}]
  ExpressionVariable *string `json:"expressionVariable,omitempty" xml:"expressionVariable,omitempty"`
  // Associated variable ID
  // 
  // example:
  // 
  // [44659]
  ExpressionVariableIds *string `json:"expressionVariableIds,omitempty" xml:"expressionVariableIds,omitempty"`
  // Variable ID
  // 
  // example:
  // 
  // 3144
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // Region code
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
  // Scene
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // EXPRESSION
  Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s ExpressionTestRequest) String() string {
  return dara.Prettify(s)
}

func (s ExpressionTestRequest) GoString() string {
  return s.String()
}

func (s *ExpressionTestRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExpressionTestRequest) GetExpression() *string  {
  return s.Expression
}

func (s *ExpressionTestRequest) GetExpressionVariable() *string  {
  return s.ExpressionVariable
}

func (s *ExpressionTestRequest) GetExpressionVariableIds() *string  {
  return s.ExpressionVariableIds
}

func (s *ExpressionTestRequest) GetId() *int64  {
  return s.Id
}

func (s *ExpressionTestRequest) GetRegId() *string  {
  return s.RegId
}

func (s *ExpressionTestRequest) GetScene() *string  {
  return s.Scene
}

func (s *ExpressionTestRequest) SetLang(v string) *ExpressionTestRequest {
  s.Lang = &v
  return s
}

func (s *ExpressionTestRequest) SetExpression(v string) *ExpressionTestRequest {
  s.Expression = &v
  return s
}

func (s *ExpressionTestRequest) SetExpressionVariable(v string) *ExpressionTestRequest {
  s.ExpressionVariable = &v
  return s
}

func (s *ExpressionTestRequest) SetExpressionVariableIds(v string) *ExpressionTestRequest {
  s.ExpressionVariableIds = &v
  return s
}

func (s *ExpressionTestRequest) SetId(v int64) *ExpressionTestRequest {
  s.Id = &v
  return s
}

func (s *ExpressionTestRequest) SetRegId(v string) *ExpressionTestRequest {
  s.RegId = &v
  return s
}

func (s *ExpressionTestRequest) SetScene(v string) *ExpressionTestRequest {
  s.Scene = &v
  return s
}

func (s *ExpressionTestRequest) Validate() error {
  return dara.Validate(s)
}

