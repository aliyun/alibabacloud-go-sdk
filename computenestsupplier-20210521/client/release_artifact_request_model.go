// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *ReleaseArtifactRequest
	GetArtifactId() *string
	SetClientToken(v string) *ReleaseArtifactRequest
	GetClientToken() *string
}

type ReleaseArtifactRequest struct {
	// The ID of the artifact.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-9feded91880e4c78xxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ReleaseArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseArtifactRequest) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ReleaseArtifactRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReleaseArtifactRequest) SetArtifactId(v string) *ReleaseArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *ReleaseArtifactRequest) SetClientToken(v string) *ReleaseArtifactRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleaseArtifactRequest) Validate() error {
	return dara.Validate(s)
}
