// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *DeleteArtifactRequest
	GetArtifactId() *string
	SetArtifactVersion(v string) *DeleteArtifactRequest
	GetArtifactVersion() *string
	SetClientToken(v string) *DeleteArtifactRequest
	GetClientToken() *string
}

type DeleteArtifactRequest struct {
	// The ID of the artifact.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The version of the artifact.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteArtifactRequest) GoString() string {
	return s.String()
}

func (s *DeleteArtifactRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *DeleteArtifactRequest) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *DeleteArtifactRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteArtifactRequest) SetArtifactId(v string) *DeleteArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *DeleteArtifactRequest) SetArtifactVersion(v string) *DeleteArtifactRequest {
	s.ArtifactVersion = &v
	return s
}

func (s *DeleteArtifactRequest) SetClientToken(v string) *DeleteArtifactRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteArtifactRequest) Validate() error {
	return dara.Validate(s)
}
