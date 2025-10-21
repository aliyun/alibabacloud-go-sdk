// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPythonArtifact interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalDependencies(v []*string) *PythonArtifact
	GetAdditionalDependencies() []*string
	SetAdditionalPythonArchives(v []*string) *PythonArtifact
	GetAdditionalPythonArchives() []*string
	SetAdditionalPythonLibraries(v []*string) *PythonArtifact
	GetAdditionalPythonLibraries() []*string
	SetEntryModule(v string) *PythonArtifact
	GetEntryModule() *string
	SetMainArgs(v string) *PythonArtifact
	GetMainArgs() *string
	SetPythonArtifactUri(v string) *PythonArtifact
	GetPythonArtifactUri() *string
}

type PythonArtifact struct {
	AdditionalDependencies    []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	AdditionalPythonArchives  []*string `json:"additionalPythonArchives,omitempty" xml:"additionalPythonArchives,omitempty" type:"Repeated"`
	AdditionalPythonLibraries []*string `json:"additionalPythonLibraries,omitempty" xml:"additionalPythonLibraries,omitempty" type:"Repeated"`
	EntryModule               *string   `json:"entryModule,omitempty" xml:"entryModule,omitempty"`
	MainArgs                  *string   `json:"mainArgs,omitempty" xml:"mainArgs,omitempty"`
	// example:
	//
	// https://oss//bucket//test.py
	PythonArtifactUri *string `json:"pythonArtifactUri,omitempty" xml:"pythonArtifactUri,omitempty"`
}

func (s PythonArtifact) String() string {
	return dara.Prettify(s)
}

func (s PythonArtifact) GoString() string {
	return s.String()
}

func (s *PythonArtifact) GetAdditionalDependencies() []*string {
	return s.AdditionalDependencies
}

func (s *PythonArtifact) GetAdditionalPythonArchives() []*string {
	return s.AdditionalPythonArchives
}

func (s *PythonArtifact) GetAdditionalPythonLibraries() []*string {
	return s.AdditionalPythonLibraries
}

func (s *PythonArtifact) GetEntryModule() *string {
	return s.EntryModule
}

func (s *PythonArtifact) GetMainArgs() *string {
	return s.MainArgs
}

func (s *PythonArtifact) GetPythonArtifactUri() *string {
	return s.PythonArtifactUri
}

func (s *PythonArtifact) SetAdditionalDependencies(v []*string) *PythonArtifact {
	s.AdditionalDependencies = v
	return s
}

func (s *PythonArtifact) SetAdditionalPythonArchives(v []*string) *PythonArtifact {
	s.AdditionalPythonArchives = v
	return s
}

func (s *PythonArtifact) SetAdditionalPythonLibraries(v []*string) *PythonArtifact {
	s.AdditionalPythonLibraries = v
	return s
}

func (s *PythonArtifact) SetEntryModule(v string) *PythonArtifact {
	s.EntryModule = &v
	return s
}

func (s *PythonArtifact) SetMainArgs(v string) *PythonArtifact {
	s.MainArgs = &v
	return s
}

func (s *PythonArtifact) SetPythonArtifactUri(v string) *PythonArtifact {
	s.PythonArtifactUri = &v
	return s
}

func (s *PythonArtifact) Validate() error {
	return dara.Validate(s)
}
