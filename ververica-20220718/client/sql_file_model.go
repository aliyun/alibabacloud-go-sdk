// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSqlFile interface {
	dara.Model
	String() string
	GoString() string
	SetBatchMode(v string) *SqlFile
	GetBatchMode() *string
	SetDescription(v string) *SqlFile
	GetDescription() *string
	SetName(v string) *SqlFile
	GetName() *string
	SetNamespace(v string) *SqlFile
	GetNamespace() *string
	SetParentId(v string) *SqlFile
	GetParentId() *string
	SetSessionClusterName(v string) *SqlFile
	GetSessionClusterName() *string
	SetSqlFileId(v string) *SqlFile
	GetSqlFileId() *string
	SetSqlScript(v string) *SqlFile
	GetSqlScript() *string
	SetWorkspace(v string) *SqlFile
	GetWorkspace() *string
}

type SqlFile struct {
	// The batch mode.
	BatchMode *string `json:"batchMode,omitempty" xml:"batchMode,omitempty"`
	// The description of the SQL file.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the SQL file.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The namespace.
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The ID of the parent SQL file.
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The name of the session cluster.
	SessionClusterName *string `json:"sessionClusterName,omitempty" xml:"sessionClusterName,omitempty"`
	// The SQL file ID.
	SqlFileId *string `json:"sqlFileId,omitempty" xml:"sqlFileId,omitempty"`
	// The SQL script content.
	SqlScript *string `json:"sqlScript,omitempty" xml:"sqlScript,omitempty"`
	// The workspace ID.
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s SqlFile) String() string {
	return dara.Prettify(s)
}

func (s SqlFile) GoString() string {
	return s.String()
}

func (s *SqlFile) GetBatchMode() *string {
	return s.BatchMode
}

func (s *SqlFile) GetDescription() *string {
	return s.Description
}

func (s *SqlFile) GetName() *string {
	return s.Name
}

func (s *SqlFile) GetNamespace() *string {
	return s.Namespace
}

func (s *SqlFile) GetParentId() *string {
	return s.ParentId
}

func (s *SqlFile) GetSessionClusterName() *string {
	return s.SessionClusterName
}

func (s *SqlFile) GetSqlFileId() *string {
	return s.SqlFileId
}

func (s *SqlFile) GetSqlScript() *string {
	return s.SqlScript
}

func (s *SqlFile) GetWorkspace() *string {
	return s.Workspace
}

func (s *SqlFile) SetBatchMode(v string) *SqlFile {
	s.BatchMode = &v
	return s
}

func (s *SqlFile) SetDescription(v string) *SqlFile {
	s.Description = &v
	return s
}

func (s *SqlFile) SetName(v string) *SqlFile {
	s.Name = &v
	return s
}

func (s *SqlFile) SetNamespace(v string) *SqlFile {
	s.Namespace = &v
	return s
}

func (s *SqlFile) SetParentId(v string) *SqlFile {
	s.ParentId = &v
	return s
}

func (s *SqlFile) SetSessionClusterName(v string) *SqlFile {
	s.SessionClusterName = &v
	return s
}

func (s *SqlFile) SetSqlFileId(v string) *SqlFile {
	s.SqlFileId = &v
	return s
}

func (s *SqlFile) SetSqlScript(v string) *SqlFile {
	s.SqlScript = &v
	return s
}

func (s *SqlFile) SetWorkspace(v string) *SqlFile {
	s.Workspace = &v
	return s
}

func (s *SqlFile) Validate() error {
	return dara.Validate(s)
}
