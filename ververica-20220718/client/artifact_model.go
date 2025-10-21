// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifact interface {
	dara.Model
	String() string
	GoString() string
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
	JarArtifact *JarArtifact `json:"jarArtifact,omitempty" xml:"jarArtifact,omitempty"`
	// example:
	//
	// SQLSCRIPT
	Kind           *string         `json:"kind,omitempty" xml:"kind,omitempty"`
	PythonArtifact *PythonArtifact `json:"pythonArtifact,omitempty" xml:"pythonArtifact,omitempty"`
	SqlArtifact    *SqlArtifact    `json:"sqlArtifact,omitempty" xml:"sqlArtifact,omitempty"`
}

func (s Artifact) String() string {
	return dara.Prettify(s)
}

func (s Artifact) GoString() string {
	return s.String()
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
