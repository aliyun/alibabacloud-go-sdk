// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecSqlTransSingleScriptTranslateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetSourceDialect(v string) *ExecSqlTransSingleScriptTranslateRequest
  GetSourceDialect() *string 
  SetSourceSqlScript(v string) *ExecSqlTransSingleScriptTranslateRequest
  GetSourceSqlScript() *string 
  SetTableMapping(v []*string) *ExecSqlTransSingleScriptTranslateRequest
  GetTableMapping() []*string 
  SetTargetDialect(v string) *ExecSqlTransSingleScriptTranslateRequest
  GetTargetDialect() *string 
}

type ExecSqlTransSingleScriptTranslateRequest struct {
  // The source SQL dialect type.
  // 
  // example:
  // 
  // hive
  SourceDialect *string `json:"sourceDialect,omitempty" xml:"sourceDialect,omitempty"`
  // The source script content. It must be Base64-encoded before being passed in. The server decodes the content before performing the conversion.
  // 
  // example:
  // 
  // SELECT 	- FROM t;
  SourceSqlScript *string `json:"sourceSqlScript,omitempty" xml:"sourceSqlScript,omitempty"`
  // The table name mapping. In string format, the source table and target table are separated by a comma (,).
  TableMapping []*string `json:"tableMapping,omitempty" xml:"tableMapping,omitempty" type:"Repeated"`
  // The target SQL dialect type.
  // 
  // example:
  // 
  // hive
  TargetDialect *string `json:"targetDialect,omitempty" xml:"targetDialect,omitempty"`
}

func (s ExecSqlTransSingleScriptTranslateRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecSqlTransSingleScriptTranslateRequest) GoString() string {
  return s.String()
}

func (s *ExecSqlTransSingleScriptTranslateRequest) GetSourceDialect() *string  {
  return s.SourceDialect
}

func (s *ExecSqlTransSingleScriptTranslateRequest) GetSourceSqlScript() *string  {
  return s.SourceSqlScript
}

func (s *ExecSqlTransSingleScriptTranslateRequest) GetTableMapping() []*string  {
  return s.TableMapping
}

func (s *ExecSqlTransSingleScriptTranslateRequest) GetTargetDialect() *string  {
  return s.TargetDialect
}

func (s *ExecSqlTransSingleScriptTranslateRequest) SetSourceDialect(v string) *ExecSqlTransSingleScriptTranslateRequest {
  s.SourceDialect = &v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateRequest) SetSourceSqlScript(v string) *ExecSqlTransSingleScriptTranslateRequest {
  s.SourceSqlScript = &v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateRequest) SetTableMapping(v []*string) *ExecSqlTransSingleScriptTranslateRequest {
  s.TableMapping = v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateRequest) SetTargetDialect(v string) *ExecSqlTransSingleScriptTranslateRequest {
  s.TargetDialect = &v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateRequest) Validate() error {
  return dara.Validate(s)
}

