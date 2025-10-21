// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSqlStatementWithContext interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalDependencies(v []*string) *SqlStatementWithContext
	GetAdditionalDependencies() []*string
	SetBatchMode(v bool) *SqlStatementWithContext
	GetBatchMode() *bool
	SetFlinkConfiguration(v map[string]interface{}) *SqlStatementWithContext
	GetFlinkConfiguration() map[string]interface{}
	SetStatement(v string) *SqlStatementWithContext
	GetStatement() *string
	SetVersionName(v string) *SqlStatementWithContext
	GetVersionName() *string
}

type SqlStatementWithContext struct {
	AdditionalDependencies []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	// This parameter is required.
	BatchMode          *bool                  `json:"batchMode,omitempty" xml:"batchMode,omitempty"`
	FlinkConfiguration map[string]interface{} `json:"flinkConfiguration,omitempty" xml:"flinkConfiguration,omitempty"`
	// This parameter is required.
	Statement   *string `json:"statement,omitempty" xml:"statement,omitempty"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty"`
}

func (s SqlStatementWithContext) String() string {
	return dara.Prettify(s)
}

func (s SqlStatementWithContext) GoString() string {
	return s.String()
}

func (s *SqlStatementWithContext) GetAdditionalDependencies() []*string {
	return s.AdditionalDependencies
}

func (s *SqlStatementWithContext) GetBatchMode() *bool {
	return s.BatchMode
}

func (s *SqlStatementWithContext) GetFlinkConfiguration() map[string]interface{} {
	return s.FlinkConfiguration
}

func (s *SqlStatementWithContext) GetStatement() *string {
	return s.Statement
}

func (s *SqlStatementWithContext) GetVersionName() *string {
	return s.VersionName
}

func (s *SqlStatementWithContext) SetAdditionalDependencies(v []*string) *SqlStatementWithContext {
	s.AdditionalDependencies = v
	return s
}

func (s *SqlStatementWithContext) SetBatchMode(v bool) *SqlStatementWithContext {
	s.BatchMode = &v
	return s
}

func (s *SqlStatementWithContext) SetFlinkConfiguration(v map[string]interface{}) *SqlStatementWithContext {
	s.FlinkConfiguration = v
	return s
}

func (s *SqlStatementWithContext) SetStatement(v string) *SqlStatementWithContext {
	s.Statement = &v
	return s
}

func (s *SqlStatementWithContext) SetVersionName(v string) *SqlStatementWithContext {
	s.VersionName = &v
	return s
}

func (s *SqlStatementWithContext) Validate() error {
	return dara.Validate(s)
}
