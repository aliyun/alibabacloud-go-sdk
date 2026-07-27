// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactUploadTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactPath(v string) *CreateArtifactUploadTokenRequest
	GetArtifactPath() *string
}

type CreateArtifactUploadTokenRequest struct {
	// example:
	//
	// upload/2026-05-25/
	ArtifactPath *string `json:"artifactPath,omitempty" xml:"artifactPath,omitempty"`
}

func (s CreateArtifactUploadTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactUploadTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactUploadTokenRequest) GetArtifactPath() *string {
	return s.ArtifactPath
}

func (s *CreateArtifactUploadTokenRequest) SetArtifactPath(v string) *CreateArtifactUploadTokenRequest {
	s.ArtifactPath = &v
	return s
}

func (s *CreateArtifactUploadTokenRequest) Validate() error {
	return dara.Validate(s)
}
