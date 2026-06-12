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
	// Call [ListArtifacts](https://help.aliyun.com/document_detail/469993.html) to obtain the artifact ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The version of the artifact.
	//
	// Call [ListArtifactVersions](https://help.aliyun.com/document_detail/469995.html) to obtain the artifact version.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// A client-generated token that ensures the idempotence of the request. Make sure that the token is unique for each request. **ClientToken*	- supports only ASCII characters and must be no more than 64 characters long.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
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
