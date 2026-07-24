// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSqlExecutionBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *StartSqlExecutionBody
	GetDescription() *string
	SetSqlFileId(v string) *StartSqlExecutionBody
	GetSqlFileId() *string
	SetSqlScript(v string) *StartSqlExecutionBody
	GetSqlScript() *string
}

type StartSqlExecutionBody struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	SqlFileId   *string `json:"sqlFileId,omitempty" xml:"sqlFileId,omitempty"`
	SqlScript   *string `json:"sqlScript,omitempty" xml:"sqlScript,omitempty"`
}

func (s StartSqlExecutionBody) String() string {
	return dara.Prettify(s)
}

func (s StartSqlExecutionBody) GoString() string {
	return s.String()
}

func (s *StartSqlExecutionBody) GetDescription() *string {
	return s.Description
}

func (s *StartSqlExecutionBody) GetSqlFileId() *string {
	return s.SqlFileId
}

func (s *StartSqlExecutionBody) GetSqlScript() *string {
	return s.SqlScript
}

func (s *StartSqlExecutionBody) SetDescription(v string) *StartSqlExecutionBody {
	s.Description = &v
	return s
}

func (s *StartSqlExecutionBody) SetSqlFileId(v string) *StartSqlExecutionBody {
	s.SqlFileId = &v
	return s
}

func (s *StartSqlExecutionBody) SetSqlScript(v string) *StartSqlExecutionBody {
	s.SqlScript = &v
	return s
}

func (s *StartSqlExecutionBody) Validate() error {
	return dara.Validate(s)
}
