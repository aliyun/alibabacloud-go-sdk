// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactRisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *ListArtifactRisksRequest
	GetArtifactId() *string
	SetArtifactVersion(v string) *ListArtifactRisksRequest
	GetArtifactVersion() *string
}

type ListArtifactRisksRequest struct {
	// Artifact ID.
	//
	// example:
	//
	// artifact-3fd95cdfdf0d4b1fa00c
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// Artifact version.
	//
	// example:
	//
	// 1
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
}

func (s ListArtifactRisksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactRisksRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactRisksRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ListArtifactRisksRequest) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *ListArtifactRisksRequest) SetArtifactId(v string) *ListArtifactRisksRequest {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactRisksRequest) SetArtifactVersion(v string) *ListArtifactRisksRequest {
	s.ArtifactVersion = &v
	return s
}

func (s *ListArtifactRisksRequest) Validate() error {
	return dara.Validate(s)
}
