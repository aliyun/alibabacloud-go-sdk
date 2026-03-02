// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSqlArtifact interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalDependencies(v []*string) *SqlArtifact
	GetAdditionalDependencies() []*string
	SetSqlScript(v string) *SqlArtifact
	GetSqlScript() *string
}

type SqlArtifact struct {
	// The additional dependency files. If you want to use dependencies such as UDFs, connectors, and formats that are not registered on Ververica Platform (VVP), you need to configure this parameter. You do not need to configure this parameter for dependencies that are registered on VVP.
	AdditionalDependencies []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	// The script of the SQL deployment.
	//
	// example:
	//
	// CREATE TEMPORARY TABLE datagen_source(   name VARCHAR ) WITH (   \\"connector\\" = \\"datagen\\" ); CREATE TEMPORARY TABLE blackhole_sink(   name  VARCHAR ) with (   \\"connector\\" = \\"blackhole\\" ); INSERT INTO blackhole_sink SELECT name from datagen_source;
	SqlScript *string `json:"sqlScript,omitempty" xml:"sqlScript,omitempty"`
}

func (s SqlArtifact) String() string {
	return dara.Prettify(s)
}

func (s SqlArtifact) GoString() string {
	return s.String()
}

func (s *SqlArtifact) GetAdditionalDependencies() []*string {
	return s.AdditionalDependencies
}

func (s *SqlArtifact) GetSqlScript() *string {
	return s.SqlScript
}

func (s *SqlArtifact) SetAdditionalDependencies(v []*string) *SqlArtifact {
	s.AdditionalDependencies = v
	return s
}

func (s *SqlArtifact) SetSqlScript(v string) *SqlArtifact {
	s.SqlScript = &v
	return s
}

func (s *SqlArtifact) Validate() error {
	return dara.Validate(s)
}
