// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifact interface {
	dara.Model
	String() string
	GoString() string
	SetCdcYamlArtifact(v *CdcYamlArtifact) *Artifact
	GetCdcYamlArtifact() *CdcYamlArtifact
	SetJarArtifact(v *JarArtifact) *Artifact
	GetJarArtifact() *JarArtifact
	SetKind(v string) *Artifact
	GetKind() *string
	SetPythonArtifact(v *PythonArtifact) *Artifact
	GetPythonArtifact() *PythonArtifact
	SetSqlArtifact(v *SqlArtifact) *Artifact
	GetSqlArtifact() *SqlArtifact
}

type Artifact struct {
	// Required for a data ingestion job.
	CdcYamlArtifact *CdcYamlArtifact `json:"cdcYamlArtifact,omitempty" xml:"cdcYamlArtifact,omitempty"`
	// Required for a JAR job.
	JarArtifact *JarArtifact `json:"jarArtifact,omitempty" xml:"jarArtifact,omitempty"`
	// Specifies the kind of job. This field is required and cannot be changed after creation.
	//
	// - SQLSCRIPT: An SQL job.
	//
	// - JAR: A JAR job.
	//
	// - PYTHON: A Python job.
	//
	// - CDCYAML: A CDC data ingestion job.
	//
	// example:
	//
	// SQLSCRIPT
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Required for a Python job.
	PythonArtifact *PythonArtifact `json:"pythonArtifact,omitempty" xml:"pythonArtifact,omitempty"`
	// Required for an SQL job.
	SqlArtifact *SqlArtifact `json:"sqlArtifact,omitempty" xml:"sqlArtifact,omitempty"`
}

func (s Artifact) String() string {
	return dara.Prettify(s)
}

func (s Artifact) GoString() string {
	return s.String()
}

func (s *Artifact) GetCdcYamlArtifact() *CdcYamlArtifact {
	return s.CdcYamlArtifact
}

func (s *Artifact) GetJarArtifact() *JarArtifact {
	return s.JarArtifact
}

func (s *Artifact) GetKind() *string {
	return s.Kind
}

func (s *Artifact) GetPythonArtifact() *PythonArtifact {
	return s.PythonArtifact
}

func (s *Artifact) GetSqlArtifact() *SqlArtifact {
	return s.SqlArtifact
}

func (s *Artifact) SetCdcYamlArtifact(v *CdcYamlArtifact) *Artifact {
	s.CdcYamlArtifact = v
	return s
}

func (s *Artifact) SetJarArtifact(v *JarArtifact) *Artifact {
	s.JarArtifact = v
	return s
}

func (s *Artifact) SetKind(v string) *Artifact {
	s.Kind = &v
	return s
}

func (s *Artifact) SetPythonArtifact(v *PythonArtifact) *Artifact {
	s.PythonArtifact = v
	return s
}

func (s *Artifact) SetSqlArtifact(v *SqlArtifact) *Artifact {
	s.SqlArtifact = v
	return s
}

func (s *Artifact) Validate() error {
	if s.CdcYamlArtifact != nil {
		if err := s.CdcYamlArtifact.Validate(); err != nil {
			return err
		}
	}
	if s.JarArtifact != nil {
		if err := s.JarArtifact.Validate(); err != nil {
			return err
		}
	}
	if s.PythonArtifact != nil {
		if err := s.PythonArtifact.Validate(); err != nil {
			return err
		}
	}
	if s.SqlArtifact != nil {
		if err := s.SqlArtifact.Validate(); err != nil {
			return err
		}
	}
	return nil
}
