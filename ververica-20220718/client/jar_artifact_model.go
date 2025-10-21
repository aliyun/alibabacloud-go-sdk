// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJarArtifact interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalDependencies(v []*string) *JarArtifact
	GetAdditionalDependencies() []*string
	SetEntryClass(v string) *JarArtifact
	GetEntryClass() *string
	SetJarUri(v string) *JarArtifact
	GetJarUri() *string
	SetMainArgs(v string) *JarArtifact
	GetMainArgs() *string
}

type JarArtifact struct {
	AdditionalDependencies []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	// example:
	//
	// org.apapche.flink.test
	EntryClass *string `json:"entryClass,omitempty" xml:"entryClass,omitempty"`
	// example:
	//
	// https://oss//bucket//test.jar
	JarUri   *string `json:"jarUri,omitempty" xml:"jarUri,omitempty"`
	MainArgs *string `json:"mainArgs,omitempty" xml:"mainArgs,omitempty"`
}

func (s JarArtifact) String() string {
	return dara.Prettify(s)
}

func (s JarArtifact) GoString() string {
	return s.String()
}

func (s *JarArtifact) GetAdditionalDependencies() []*string {
	return s.AdditionalDependencies
}

func (s *JarArtifact) GetEntryClass() *string {
	return s.EntryClass
}

func (s *JarArtifact) GetJarUri() *string {
	return s.JarUri
}

func (s *JarArtifact) GetMainArgs() *string {
	return s.MainArgs
}

func (s *JarArtifact) SetAdditionalDependencies(v []*string) *JarArtifact {
	s.AdditionalDependencies = v
	return s
}

func (s *JarArtifact) SetEntryClass(v string) *JarArtifact {
	s.EntryClass = &v
	return s
}

func (s *JarArtifact) SetJarUri(v string) *JarArtifact {
	s.JarUri = &v
	return s
}

func (s *JarArtifact) SetMainArgs(v string) *JarArtifact {
	s.MainArgs = &v
	return s
}

func (s *JarArtifact) Validate() error {
	return dara.Validate(s)
}
