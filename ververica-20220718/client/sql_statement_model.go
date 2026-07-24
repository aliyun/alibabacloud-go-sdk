// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSqlStatement interface {
	dara.Model
	String() string
	GoString() string
	SetIndex(v int32) *SqlStatement
	GetIndex() *int32
	SetMessage(v string) *SqlStatement
	GetMessage() *string
	SetSqlScript(v string) *SqlStatement
	GetSqlScript() *string
	SetStatusState(v string) *SqlStatement
	GetStatusState() *string
	SetType(v string) *SqlStatement
	GetType() *string
}

type SqlStatement struct {
	Index       *int32  `json:"index,omitempty" xml:"index,omitempty"`
	Message     *string `json:"message,omitempty" xml:"message,omitempty"`
	SqlScript   *string `json:"sqlScript,omitempty" xml:"sqlScript,omitempty"`
	StatusState *string `json:"statusState,omitempty" xml:"statusState,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SqlStatement) String() string {
	return dara.Prettify(s)
}

func (s SqlStatement) GoString() string {
	return s.String()
}

func (s *SqlStatement) GetIndex() *int32 {
	return s.Index
}

func (s *SqlStatement) GetMessage() *string {
	return s.Message
}

func (s *SqlStatement) GetSqlScript() *string {
	return s.SqlScript
}

func (s *SqlStatement) GetStatusState() *string {
	return s.StatusState
}

func (s *SqlStatement) GetType() *string {
	return s.Type
}

func (s *SqlStatement) SetIndex(v int32) *SqlStatement {
	s.Index = &v
	return s
}

func (s *SqlStatement) SetMessage(v string) *SqlStatement {
	s.Message = &v
	return s
}

func (s *SqlStatement) SetSqlScript(v string) *SqlStatement {
	s.SqlScript = &v
	return s
}

func (s *SqlStatement) SetStatusState(v string) *SqlStatement {
	s.StatusState = &v
	return s
}

func (s *SqlStatement) SetType(v string) *SqlStatement {
	s.Type = &v
	return s
}

func (s *SqlStatement) Validate() error {
	return dara.Validate(s)
}
