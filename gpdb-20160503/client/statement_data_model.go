// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatementData interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *StatementData
	GetCreatedAt() *string
	SetDatabase(v string) *StatementData
	GetDatabase() *string
	SetId(v string) *StatementData
	GetId() *string
	SetParameters(v []*string) *StatementData
	GetParameters() []*string
	SetSecretArn(v string) *StatementData
	GetSecretArn() *string
	SetSql(v string) *StatementData
	GetSql() *string
	SetSqls(v []*string) *StatementData
	GetSqls() []*string
	SetStatus(v string) *StatementData
	GetStatus() *string
	SetUpdatedAt(v string) *StatementData
	GetUpdatedAt() *string
}

type StatementData struct {
	CreatedAt  *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Database   *string   `json:"Database,omitempty" xml:"Database,omitempty"`
	Id         *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	Parameters []*string `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	SecretArn  *string   `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	Sql        *string   `json:"Sql,omitempty" xml:"Sql,omitempty"`
	Sqls       []*string `json:"Sqls,omitempty" xml:"Sqls,omitempty" type:"Repeated"`
	Status     *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdatedAt  *string   `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s StatementData) String() string {
	return dara.Prettify(s)
}

func (s StatementData) GoString() string {
	return s.String()
}

func (s *StatementData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *StatementData) GetDatabase() *string {
	return s.Database
}

func (s *StatementData) GetId() *string {
	return s.Id
}

func (s *StatementData) GetParameters() []*string {
	return s.Parameters
}

func (s *StatementData) GetSecretArn() *string {
	return s.SecretArn
}

func (s *StatementData) GetSql() *string {
	return s.Sql
}

func (s *StatementData) GetSqls() []*string {
	return s.Sqls
}

func (s *StatementData) GetStatus() *string {
	return s.Status
}

func (s *StatementData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *StatementData) SetCreatedAt(v string) *StatementData {
	s.CreatedAt = &v
	return s
}

func (s *StatementData) SetDatabase(v string) *StatementData {
	s.Database = &v
	return s
}

func (s *StatementData) SetId(v string) *StatementData {
	s.Id = &v
	return s
}

func (s *StatementData) SetParameters(v []*string) *StatementData {
	s.Parameters = v
	return s
}

func (s *StatementData) SetSecretArn(v string) *StatementData {
	s.SecretArn = &v
	return s
}

func (s *StatementData) SetSql(v string) *StatementData {
	s.Sql = &v
	return s
}

func (s *StatementData) SetSqls(v []*string) *StatementData {
	s.Sqls = v
	return s
}

func (s *StatementData) SetStatus(v string) *StatementData {
	s.Status = &v
	return s
}

func (s *StatementData) SetUpdatedAt(v string) *StatementData {
	s.UpdatedAt = &v
	return s
}

func (s *StatementData) Validate() error {
	return dara.Validate(s)
}
