// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifactMeta interface {
	dara.Model
	String() string
	GoString() string
	SetChecksum(v string) *ArtifactMeta
	GetChecksum() *string
	SetName(v string) *ArtifactMeta
	GetName() *string
}

type ArtifactMeta struct {
	// example:
	//
	// CRC-64 code
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// example:
	//
	// my-artifact
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ArtifactMeta) String() string {
	return dara.Prettify(s)
}

func (s ArtifactMeta) GoString() string {
	return s.String()
}

func (s *ArtifactMeta) GetChecksum() *string {
	return s.Checksum
}

func (s *ArtifactMeta) GetName() *string {
	return s.Name
}

func (s *ArtifactMeta) SetChecksum(v string) *ArtifactMeta {
	s.Checksum = &v
	return s
}

func (s *ArtifactMeta) SetName(v string) *ArtifactMeta {
	s.Name = &v
	return s
}

func (s *ArtifactMeta) Validate() error {
	return dara.Validate(s)
}
