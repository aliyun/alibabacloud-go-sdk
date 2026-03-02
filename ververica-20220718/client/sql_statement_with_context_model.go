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
	// The additional dependencies.
	AdditionalDependencies []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	// Specifies whether the deployment is a batch deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	BatchMode *bool `json:"batchMode,omitempty" xml:"batchMode,omitempty"`
	// The Realtime Compute for Apache Flink configuration.
	FlinkConfiguration map[string]interface{} `json:"flinkConfiguration,omitempty" xml:"flinkConfiguration,omitempty"`
	// The code of the deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREATE TEMPORARY TABLE datagen_source (
	//
	//     name VARCHAR,
	//
	//     score BIGINT
	//
	// ) WITH (
	//
	//    \\"connector\\" = \\"datagen\\"
	//
	// );
	//
	// CREATE TEMPORARY TABLE print_table (
	//
	//    name VARCHAR,
	//
	//    score BIGINT
	//
	// ) WITH (
	//
	//   \\"connector\\"=\\"print\\",
	//
	//   \\"logger\\"=\\"true\\"
	//
	// );
	//
	// INSERT INTO print_table
	//
	// select 	- from datagen_source;
	Statement *string `json:"statement,omitempty" xml:"statement,omitempty"`
	// The engine version.
	//
	// example:
	//
	// vvr-8.0.6-flink-1.17
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
