// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSqlFromNLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *GenerateSqlFromNLRequest
	GetDbId() *string
	SetDialect(v string) *GenerateSqlFromNLRequest
	GetDialect() *string
	SetKnowledge(v string) *GenerateSqlFromNLRequest
	GetKnowledge() *string
	SetLevel(v string) *GenerateSqlFromNLRequest
	GetLevel() *string
	SetModel(v string) *GenerateSqlFromNLRequest
	GetModel() *string
	SetQuestion(v string) *GenerateSqlFromNLRequest
	GetQuestion() *string
	SetTableNames(v string) *GenerateSqlFromNLRequest
	GetTableNames() *string
}

type GenerateSqlFromNLRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// MySQL
	Dialect *string `json:"Dialect,omitempty" xml:"Dialect,omitempty"`
	// example:
	//
	// state>0代表成功
	Knowledge *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	// example:
	//
	// base
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// users,orders
	TableNames *string `json:"TableNames,omitempty" xml:"TableNames,omitempty"`
}

func (s GenerateSqlFromNLRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlFromNLRequest) GoString() string {
	return s.String()
}

func (s *GenerateSqlFromNLRequest) GetDbId() *string {
	return s.DbId
}

func (s *GenerateSqlFromNLRequest) GetDialect() *string {
	return s.Dialect
}

func (s *GenerateSqlFromNLRequest) GetKnowledge() *string {
	return s.Knowledge
}

func (s *GenerateSqlFromNLRequest) GetLevel() *string {
	return s.Level
}

func (s *GenerateSqlFromNLRequest) GetModel() *string {
	return s.Model
}

func (s *GenerateSqlFromNLRequest) GetQuestion() *string {
	return s.Question
}

func (s *GenerateSqlFromNLRequest) GetTableNames() *string {
	return s.TableNames
}

func (s *GenerateSqlFromNLRequest) SetDbId(v string) *GenerateSqlFromNLRequest {
	s.DbId = &v
	return s
}

func (s *GenerateSqlFromNLRequest) SetDialect(v string) *GenerateSqlFromNLRequest {
	s.Dialect = &v
	return s
}

func (s *GenerateSqlFromNLRequest) SetKnowledge(v string) *GenerateSqlFromNLRequest {
	s.Knowledge = &v
	return s
}

func (s *GenerateSqlFromNLRequest) SetLevel(v string) *GenerateSqlFromNLRequest {
	s.Level = &v
	return s
}

func (s *GenerateSqlFromNLRequest) SetModel(v string) *GenerateSqlFromNLRequest {
	s.Model = &v
	return s
}

func (s *GenerateSqlFromNLRequest) SetQuestion(v string) *GenerateSqlFromNLRequest {
	s.Question = &v
	return s
}

func (s *GenerateSqlFromNLRequest) SetTableNames(v string) *GenerateSqlFromNLRequest {
	s.TableNames = &v
	return s
}

func (s *GenerateSqlFromNLRequest) Validate() error {
	return dara.Validate(s)
}
