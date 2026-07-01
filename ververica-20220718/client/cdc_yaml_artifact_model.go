// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCdcYamlArtifact interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalDependencies(v []*string) *CdcYamlArtifact
	GetAdditionalDependencies() []*string
	SetCdcYaml(v string) *CdcYamlArtifact
	GetCdcYaml() *string
}

type CdcYamlArtifact struct {
	// Full URL paths to the additional dependencies.
	AdditionalDependencies []*string `json:"additionalDependencies,omitempty" xml:"additionalDependencies,omitempty" type:"Repeated"`
	// Code for the data ingestion job.
	CdcYaml *string `json:"cdcYaml,omitempty" xml:"cdcYaml,omitempty"`
}

func (s CdcYamlArtifact) String() string {
	return dara.Prettify(s)
}

func (s CdcYamlArtifact) GoString() string {
	return s.String()
}

func (s *CdcYamlArtifact) GetAdditionalDependencies() []*string {
	return s.AdditionalDependencies
}

func (s *CdcYamlArtifact) GetCdcYaml() *string {
	return s.CdcYaml
}

func (s *CdcYamlArtifact) SetAdditionalDependencies(v []*string) *CdcYamlArtifact {
	s.AdditionalDependencies = v
	return s
}

func (s *CdcYamlArtifact) SetCdcYaml(v string) *CdcYamlArtifact {
	s.CdcYaml = &v
	return s
}

func (s *CdcYamlArtifact) Validate() error {
	return dara.Validate(s)
}
