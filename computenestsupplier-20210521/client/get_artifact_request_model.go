// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *GetArtifactRequest
	GetArtifactId() *string
	SetArtifactName(v string) *GetArtifactRequest
	GetArtifactName() *string
	SetArtifactVersion(v string) *GetArtifactRequest
	GetArtifactVersion() *string
}

type GetArtifactRequest struct {
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The name of the deployment package.
	//
	// example:
	//
	// gpu-test
	ArtifactName *string `json:"ArtifactName,omitempty" xml:"ArtifactName,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// 1
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
}

func (s GetArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *GetArtifactRequest) GetArtifactName() *string {
	return s.ArtifactName
}

func (s *GetArtifactRequest) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *GetArtifactRequest) SetArtifactId(v string) *GetArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *GetArtifactRequest) SetArtifactName(v string) *GetArtifactRequest {
	s.ArtifactName = &v
	return s
}

func (s *GetArtifactRequest) SetArtifactVersion(v string) *GetArtifactRequest {
	s.ArtifactVersion = &v
	return s
}

func (s *GetArtifactRequest) Validate() error {
	return dara.Validate(s)
}
