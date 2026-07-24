// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSqlExecution interface {
	dara.Model
	String() string
	GoString() string
	SetBatchMode(v bool) *SqlExecution
	GetBatchMode() *bool
	SetDescription(v string) *SqlExecution
	GetDescription() *string
	SetMessage(v string) *SqlExecution
	GetMessage() *string
	SetName(v string) *SqlExecution
	GetName() *string
	SetNamespace(v string) *SqlExecution
	GetNamespace() *string
	SetSessionClusterName(v string) *SqlExecution
	GetSessionClusterName() *string
	SetSqlExecutionId(v string) *SqlExecution
	GetSqlExecutionId() *string
	SetSqlFileId(v string) *SqlExecution
	GetSqlFileId() *string
	SetSqlScript(v string) *SqlExecution
	GetSqlScript() *string
	SetState(v string) *SqlExecution
	GetState() *string
	SetStatements(v []*SqlStatement) *SqlExecution
	GetStatements() []*SqlStatement
	SetWorkspace(v string) *SqlExecution
	GetWorkspace() *string
}

type SqlExecution struct {
	BatchMode   *bool   `json:"batchMode,omitempty" xml:"batchMode,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Message     *string `json:"message,omitempty" xml:"message,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace          *string         `json:"namespace,omitempty" xml:"namespace,omitempty"`
	SessionClusterName *string         `json:"sessionClusterName,omitempty" xml:"sessionClusterName,omitempty"`
	SqlExecutionId     *string         `json:"sqlExecutionId,omitempty" xml:"sqlExecutionId,omitempty"`
	SqlFileId          *string         `json:"sqlFileId,omitempty" xml:"sqlFileId,omitempty"`
	SqlScript          *string         `json:"sqlScript,omitempty" xml:"sqlScript,omitempty"`
	State              *string         `json:"state,omitempty" xml:"state,omitempty"`
	Statements         []*SqlStatement `json:"statements,omitempty" xml:"statements,omitempty" type:"Repeated"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s SqlExecution) String() string {
	return dara.Prettify(s)
}

func (s SqlExecution) GoString() string {
	return s.String()
}

func (s *SqlExecution) GetBatchMode() *bool {
	return s.BatchMode
}

func (s *SqlExecution) GetDescription() *string {
	return s.Description
}

func (s *SqlExecution) GetMessage() *string {
	return s.Message
}

func (s *SqlExecution) GetName() *string {
	return s.Name
}

func (s *SqlExecution) GetNamespace() *string {
	return s.Namespace
}

func (s *SqlExecution) GetSessionClusterName() *string {
	return s.SessionClusterName
}

func (s *SqlExecution) GetSqlExecutionId() *string {
	return s.SqlExecutionId
}

func (s *SqlExecution) GetSqlFileId() *string {
	return s.SqlFileId
}

func (s *SqlExecution) GetSqlScript() *string {
	return s.SqlScript
}

func (s *SqlExecution) GetState() *string {
	return s.State
}

func (s *SqlExecution) GetStatements() []*SqlStatement {
	return s.Statements
}

func (s *SqlExecution) GetWorkspace() *string {
	return s.Workspace
}

func (s *SqlExecution) SetBatchMode(v bool) *SqlExecution {
	s.BatchMode = &v
	return s
}

func (s *SqlExecution) SetDescription(v string) *SqlExecution {
	s.Description = &v
	return s
}

func (s *SqlExecution) SetMessage(v string) *SqlExecution {
	s.Message = &v
	return s
}

func (s *SqlExecution) SetName(v string) *SqlExecution {
	s.Name = &v
	return s
}

func (s *SqlExecution) SetNamespace(v string) *SqlExecution {
	s.Namespace = &v
	return s
}

func (s *SqlExecution) SetSessionClusterName(v string) *SqlExecution {
	s.SessionClusterName = &v
	return s
}

func (s *SqlExecution) SetSqlExecutionId(v string) *SqlExecution {
	s.SqlExecutionId = &v
	return s
}

func (s *SqlExecution) SetSqlFileId(v string) *SqlExecution {
	s.SqlFileId = &v
	return s
}

func (s *SqlExecution) SetSqlScript(v string) *SqlExecution {
	s.SqlScript = &v
	return s
}

func (s *SqlExecution) SetState(v string) *SqlExecution {
	s.State = &v
	return s
}

func (s *SqlExecution) SetStatements(v []*SqlStatement) *SqlExecution {
	s.Statements = v
	return s
}

func (s *SqlExecution) SetWorkspace(v string) *SqlExecution {
	s.Workspace = &v
	return s
}

func (s *SqlExecution) Validate() error {
	if s.Statements != nil {
		for _, item := range s.Statements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
